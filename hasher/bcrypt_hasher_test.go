package hasher

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestBCryptHasher(t *testing.T) {
	t.Run("NewBCryptHasher", func(t *testing.T) {
		bCryptHasher, isBCryptHasher := NewBCryptHasher().(*bCryptHasherImpl)

		assert.True(t, isBCryptHasher)
		assert.Equal(t, BCryptPasswordCost, bCryptHasher.cost)
	})

	t.Run("Make", func(t *testing.T) {
		password := "pass0"
		bCryptHasher := &bCryptHasherImpl{cost: BCryptPasswordCost}
		hash, err := bCryptHasher.Make(password)

		assert.Nil(t, bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)))
		assert.Nil(t, err)
	})

	t.Run("Make:Error", func(t *testing.T) {
		bCryptHasher := &bCryptHasherImpl{cost: 100}
		hash, err := bCryptHasher.Make("")

		assert.Empty(t, hash)
		assert.Error(t, err)
	})

	t.Run("Check", func(t *testing.T) {
		password := "pass0"
		bCryptHasher := &bCryptHasherImpl{cost: BCryptPasswordCost}
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), BCryptPasswordCost)

		assert.Nil(t, bCryptHasher.Check(string(hash), password))
	})

	t.Run("Check:Error", func(t *testing.T) {
		bCryptHasher := new(bCryptHasherImpl)

		assert.Error(t, bCryptHasher.Check("", ""))
	})
}
