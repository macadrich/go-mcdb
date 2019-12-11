package postgres

import (
	"log"
	"testing"
	"time"
)

func TestPostgresConnection(t *testing.T) {
	db, err := ConnectDB("localhost", 5432, "adriel", "postgresAD32", "people", "disable")
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
		db, err := ConnectDB("localhost", 5432, "adriel", "postgresAD32", "people", "disable")
		if err != nil {
			t.Errorf("Error connecting postgresql: %s", err.Error())
			return
		}
		log.Println("Connection success:", db)

		// (age, email, first_name, last_name)

		_, err = db.Add("ken@gmail.com", "ken", "password567", time.Now().Unix(), time.Now().Unix())
		if err != nil {
			t.Error(err)
		}
	})

}
