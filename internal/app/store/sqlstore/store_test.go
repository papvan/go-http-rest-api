package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M)  {
	databaseURL = os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		databaseURL = "host=localhost dbname=restapi_test sslmode=disable user=db_user password=db_user_pass"
	}

	os.Exit(m.Run())
}