package configs

import (
	"fmt"
	"testing"
)

func TestConnectDB(t *testing.T) {
	db, err := ConnectDB()
	if err != nil {
		t.Error("Error to connect with db")
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		msg := fmt.Sprintf("Ping DB error: %s", err)
		t.Error(msg)
	}
}

func TestConnectMO(t *testing.T) {
	session, err := ConnectMO()
	if err != nil {
		t.Error("Error to Connect with mongo")
	}
	defer session.Close()
}
