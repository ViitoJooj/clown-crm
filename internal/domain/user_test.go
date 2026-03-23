package domain_test

import (
	"crm/internal/domain"
	"fmt"
	"testing"
)

func TestNewUser(t *testing.T) {
	first_name := "João Vitor"
	last_name := "Santana Oqueres"
	email := "joaovitor819@gmail.com"
	password := "123123123Mj4!"

	user, err := domain.NewUser(first_name, last_name, email, password)
	if err != nil {
		t.Errorf("%s", err)
	}

	fmt.Println(user)
}
