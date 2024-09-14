package database_test

import (
	"plunger-beam/internal/database"
	"testing"
)

func TestConnect(t *testing.T) {
	mysqlOption := database.MySqlOption{
		Address:     "localhost",
		Username:    "root",
		Password:    "root",
		Port:        "3307",
		Database:    "chat-database",
		IsPopulated: false,
		IsMigrate:   true,
	}

	mysqlConn := database.NewMySqlConnection(mysqlOption)
	if err := mysqlConn.ConnectToDB(); err != nil {
		t.Error("Expected successful connection to database : ", err)
	}
}
