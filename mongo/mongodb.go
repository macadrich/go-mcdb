package mongo

import (
	"crypto/rand"
	"errors"
	"go-serverless-example/model"
	"math/big"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DB -
type DB struct {
	conn *mgo.Session
	user *mgo.Collection
}

var maxRand = big.NewInt(1<<63 - 1)

// NewMongoDB initialize mongodb with credential and source
func NewMongoDB(host, username, password, database, source string) (*mgo.Session, error) {
	Host := []string{host}
	conn, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    Host,
		Username: username,
		Password: password,
		Database: database,
		Source:   source,
	})
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return conn, nil
}

// Close mongo database.
func (db *DB) Close() {
	db.conn.Close()
}

func randomID() (int64, error) {
	// Get a random number within the range [0, 1<<63-1)
	n, err := rand.Int(rand.Reader, maxRand)
	if err != nil {
		return 0, err
	}
	// Don't assign 0.
	return n.Int64() + 1, nil
}

// DeleteUser -
func (db *DB) DeleteUser(id string) error {
	return db.user.Remove(bson.D{{Name: "id", Value: id}})
}

// GetUserByID -
func (db *DB) GetUserByID(id string) (*model.User, error) {
	user := &model.User{}
	if err := db.user.Find(bson.D{{Name: "id", Value: id}}).One(user); err != nil {
		return user, errors.New(err.Error())
	}
	return user, nil
}

// GetUser -
func (db *DB) GetUser(email string, user interface{}) error {
	if err := db.user.Find(bson.D{{Name: "email", Value: email}}).One(user); err != nil {
		return err
	}
	return nil
}

// AddUser saves a given user, assigning it a new ID.
func (db *DB) AddUser(obj interface{}) (id string, err error) {
	if err := db.user.Insert(obj); err != nil {
		return "", errors.New(err.Error())
	}
	return id, nil
}

// UpdateUser -
func (db *DB) UpdateUser(id string, user interface{}) error {
	return db.user.Update(bson.D{{Name: "id", Value: id}}, user)
}

// ListUsers -
func (db *DB) ListUsers(user interface{}) error {
	if err := db.user.Find(nil).All(user); err != nil {
		return err
	}
	return nil
}
