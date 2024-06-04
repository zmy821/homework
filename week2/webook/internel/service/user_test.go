package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
import "golang.org/x/crypto/bcrypt"

func TestPasswordEncrypt(t *testing.T) {
	password := []byte("123456#hello")
	encrypted, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	assert.NoError(t, err)
	println(string(encrypted))
	err = bcrypt.CompareHashAndPassword(encrypted, []byte("123456#hello"))
	assert.NotNil(t, err)
}
