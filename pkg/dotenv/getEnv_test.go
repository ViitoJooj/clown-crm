package dotenv_test

import (
	"testing"

	"github.com/ViitoJooj/clown-crm/pkg/dotenv"
)

func TestGetEnv(t *testing.T) {
	dotenv.GetEnv()

	if dotenv.PgUrl == "" {
		t.Errorf("Error, PgUrl is null")
	}
}
