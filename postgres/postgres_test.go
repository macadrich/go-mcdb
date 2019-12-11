package postgres

import (
	"log"
	"testing"
	"time"
)

const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "user123"
	PASSWORD = "password123"
	DATABASE = "records"
	SSLMODE  = "disable"
)

func TestPostgresConnection(t *testing.T) {
	db, err := ConnectDB(HOST, PORT, USER, PASSWORD, DATABASE, SSLMODE)
	t.Run("postgres connection", func(t *testing.T) {
		if err != nil {
			t.Errorf("Error connecting postgresql: %s", err.Error())
			return
		}
		log.Println("-------------------")
		log.Println("Connect success:", db)
		log.Println("-------------------")
	})
}

func TestAddItem(t *testing.T) {
	t.Run("connect to postgresql database", func(t *testing.T) {
		db, err := ConnectDB(HOST, PORT, USER, PASSWORD, DATABASE, SSLMODE)
		if err != nil {
			t.Errorf("Error connecting postgresql: %s", err.Error())
			return
		}
		_, err = db.Add("test3@gmail.com", "test3", "password1235", time.Now(), time.Now())
		if err != nil {
			t.Error(err)
		}
		log.Println("Add item success:", db)
	})
}
