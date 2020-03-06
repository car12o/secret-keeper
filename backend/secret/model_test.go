package secret_test

import "testing"

func TestSecretModel(t *testing.T) {
	t.Run("Test example true", func(t *testing.T) {
		r := true
		if r != true {
			t.Errorf("r not true")
		}
	})

	t.Run("Test example false", func(t *testing.T) {
		r := false
		if r != true {
			t.Errorf("r not true")
		}
	})
}
