package secret

import (
	"net/http"
	"time"

	"github.com/car12o/secret-keeper/database"
	"github.com/car12o/secret-keeper/metrics"
	"github.com/car12o/secret-keeper/response"

	"github.com/gin-gonic/gin"
)

type secretFormData struct {
	Secret           string `json:"secret" form:"secret" binding:"min=4,required"`
	ExpireAfterViews int32  `json:"expireAfterViews" form:"expireAfterViews" binding:"gte=0"`
	ExpireAfter      int32  `json:"expireAfter" form:"expireAfter" binding:"min=1,max=60,required"`
}

// Post ...
func Post(c *gin.Context) {
	defer metrics.RequestEnd(c)

	var body secretFormData
	if err := c.ShouldBind(&body); err != nil {
		response.Send(c, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dateFormat := "2006-01-02T15:04:05"
	s, err := NewModel(database.Connection).Create(Secret{
		CreatedAt:      time.Now().Format(dateFormat),
		ExpiresAt:      time.Now().Add(time.Minute * time.Duration(body.ExpireAfter)).Format(dateFormat),
		SecretText:     body.Secret,
		RemainingViews: body.ExpireAfterViews,
	})
	if err != nil {
		response.Send(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response.Send(c, http.StatusOK, s)
}

// Get ...
func Get(c *gin.Context) {
	defer metrics.RequestEnd(c)

	hash := c.Param("hash")

	secretModel := NewModel(database.Connection)
	s, err := secretModel.FindByID(hash)
	if err != nil {
		response.Send(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	secret, err := secretModel.UpdateAndCheckExpires(s)
	if err != nil {
		response.Send(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response.Send(c, http.StatusOK, secret)
}
