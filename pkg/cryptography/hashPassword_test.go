package cryptography_test

import (
	"fmt"
	"testing"

	"github.com/ViitoJooj/clown-crm/pkg/cryptography"
)

func TestHashPassword(t *testing.T) {
	input := "ILoveGolang"

	hash, err := cryptography.HashPassword(input)
	if err != nil {
		t.Errorf("error on hash password, error: %s", err)
	}

	fmt.Println("Input: " + input)
	fmt.Println("Output: " + hash)

	PasswordIsValid := cryptography.CheckPasswordHash(input, hash)
	if !PasswordIsValid {
		t.Errorf("error, invalid hash/password")
	}
}
