package secret_test

import "testing"

func TestSecretModel(t *testing.T) {
	t.Run("Test example true", func(t *testing.T) {
		r := true
		if r != true {
			t.Errorf("r not true")
		}
	})
}
