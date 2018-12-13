package common

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestBCryptHasher(t *testing.T) {
	hasher := NewBCryptHasher()
	testPass := "pass0"

	t.Run("Make(string)", func(t *testing.T) {
		hash, err := hasher.Make(testPass)

		assert.Nil(t, bcrypt.CompareHashAndPassword([]byte(hash), []byte(testPass)))
		assert.Nil(t, err)
	})

	t.Run("Make(string) Error", func(t *testing.T) {
		hasher2 := &BCryptHasher{cost: 100}
		hash, err := hasher2.Make("")

		assert.Empty(t, hash)
		assert.Error(t, err)
	})

	t.Run("Check(string,string)", func(t *testing.T) {
		hash, _ := bcrypt.GenerateFromPassword([]byte(testPass), hasher.cost)
		err := hasher.Check(string(hash), testPass)

		assert.Nil(t, err)
	})

	t.Run("Check(string,string)", func(t *testing.T) {
		err := hasher.Check(testPass, testPass)

		assert.Error(t, err)
	})
}