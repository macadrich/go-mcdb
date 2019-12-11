package postgres

import (
	"database/sql"
	"fmt"

	"errors"

	_ "github.com/lib/pq"
)

// Config -
type Config struct {
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
	SSLMode      string
}

// DB hold fields data that need in connection for database
type DB struct {
	Database *sql.DB
	Cfg      Config
}

// ConnectDB initialize with credentials
func ConnectDB(host string, port int, user string, password string, dbname string, sslmode string) (*DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DB{
		Database: db,
		Cfg: Config{
			Host:         host,
			Port:         port,
			User:         user,
			Password:     password,
			DatabaseName: dbname,
			SSLMode:      sslmode,
		},
	}, nil
}

// Close database and return if success, else error
func (db *DB) Close() (err error) {
	if db.Database == nil {
		return
	}

	if err = db.Database.Close(); err != nil {
		err = errors.New("Errored closing database connection")
	}
	return
}

// Add insert items in table
func (db *DB) Add(item ...interface{}) (id int64, err error) {
	insertStmt := `INSERT INTO users (email, username, password, created, updated) VALUES ($1, $2, $3, $4, $5)`
	result, err := db.Database.Exec(insertStmt, item...)
	if err != nil {
		return -1, err
	}
	id, err = result.RowsAffected()
	if err != nil {
		return -1, err
	}
	return id, nil
}
