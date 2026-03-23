package database_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ViitoJooj/clown-crm/pkg/database"
	"github.com/ViitoJooj/clown-crm/pkg/dotenv"
)

func TestConn(t *testing.T) {
	dotenv.GetEnv()
	database.Conn()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := database.DB.Ping(ctx)
	if err != nil {
		t.Errorf("Database ping failed: %v\n", err)
	}

	fmt.Println("Database connection successful!")
}
