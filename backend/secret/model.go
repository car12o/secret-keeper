package secret

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Secret ...
type Secret struct {
	Hash           interface{} `json:"hash" xml:"hash" bson:"_id,omitempty"`
	SecretText     string      `json:"secretText" xml:"secretText" bson:"secretText"`
	CreatedAt      string      `json:"createdAt" xml:"createdAt" bson:"createdAt"`
	ExpiresAt      string      `json:"expiresAt" xml:"expiresAt" bson:"expiresAt"`
	RemainingViews int32       `json:"remainingViews" xml:"remainingViews" bson:"remainingViews"`
}

// Model ...
type Model struct {
	database *mongo.Collection
}

// NewModel ...
func NewModel(dbConnection *mongo.Database) *Model {
	return &Model{
		database: dbConnection.Collection("secrets"),
	}
}

// FindByID ...
func (m *Model) FindByID(id string) (Secret, error) {
	var s Secret

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Secret{}, err
	}

	res := m.database.FindOne(context.Background(), bson.M{"_id": objectID})
	if err := res.Decode(&s); err != nil {
		return Secret{}, err
	}
	s.Hash = s.Hash.(primitive.ObjectID).Hex()

	decryptedSecret, err := decrypt(os.Getenv("ENCRYPT_SECRET"), s.SecretText)
	if err != nil {
		return Secret{}, err
	}
	s.SecretText = decryptedSecret

	return s, nil
}

// Create ...
func (m *Model) Create(s Secret) (Secret, error) {
	secretText := s.SecretText

	encryptedSecret, err := encrypt(os.Getenv("ENCRYPT_SECRET"), s.SecretText)
	if err != nil {
		return Secret{}, err
	}
	s.SecretText = encryptedSecret

	res, err := m.database.InsertOne(context.Background(), s)
	if err != nil {
		return Secret{}, err
	}
	s.Hash = res.InsertedID.(primitive.ObjectID).Hex()
	s.SecretText = secretText

	return s, nil
}

// Update ...
func (m *Model) Update(s Secret) (Secret, error) {
	secretText := s.SecretText

	encryptedSecret, err := encrypt(os.Getenv("ENCRYPT_SECRET"), s.SecretText)
	if err != nil {
		return Secret{}, err
	}
	s.SecretText = encryptedSecret

	objectID, err := primitive.ObjectIDFromHex(s.Hash.(string))
	if err != nil {
		return Secret{}, err
	}
	s.Hash = objectID

	if res := m.database.FindOneAndReplace(context.Background(), bson.M{"_id": s.Hash}, s); res.Err() != nil {
		return Secret{}, res.Err()
	}
	s.Hash = s.Hash.(primitive.ObjectID).Hex()
	s.SecretText = secretText

	return s, nil
}

// UpdateAndCheckExpires ...
func (m *Model) UpdateAndCheckExpires(s Secret) (Secret, error) {
	dateFormat := "2006-01-02T15:04:05"

	date, err := time.Parse(dateFormat, s.ExpiresAt)
	if err != nil {
		return Secret{}, err
	}

	now, err := time.Parse(dateFormat, time.Now().Format(dateFormat))
	if err != nil {
		return Secret{}, err
	}

	if date.Before(now) {
		return Secret{}, fmt.Errorf("Secret already expired")
	}

	s.RemainingViews = s.RemainingViews - 1
	if s.RemainingViews < 0 {
		return Secret{}, fmt.Errorf("Secret has no remaining views")
	}

	secret, err := m.Update(s)
	if err != nil {
		return Secret{}, err
	}

	return secret, nil
}

func encrypt(key, secret string) (string, error) {
	byteKey := []byte(key)
	byteText := []byte(secret)

	block, err := aes.NewCipher(byteKey)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(byteText))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], byteText)

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func decrypt(key, cryptoText string) (string, error) {
	byteKey := []byte(key)
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(byteKey)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext), nil
}
