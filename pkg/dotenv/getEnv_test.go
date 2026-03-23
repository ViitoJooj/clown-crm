package dotenv_test

import (
	"crm/pkg/dotenv"
	"testing"
)

func TestGetEnv(t *testing.T) {
	dotenv.GetEnv()

	if dotenv.PgUrl == "" {
		t.Errorf("Error, PgUrl is null")
	}
}
