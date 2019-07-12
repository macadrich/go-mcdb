package mongo

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DB -
type DB struct {
	Conn       *mgo.Session
	TableName  string
	Collection map[string]*mgo.Collection
}

// NewMongoDB initialize mongodb with credential and source
func NewMongoDB(host, username, password, database, source string) (*DB, error) {
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
	return &DB{
		Conn:       conn,
		TableName:  "",
		Collection: make(map[string]*mgo.Collection),
	}, nil
}

// Close mongo database.
func (db *DB) Close() {
	db.Conn.Close()
}

// DeleteUser delete user in database
func (db *DB) DeleteUser(id string) error {
	return db.Collection[db.TableName].Remove(bson.D{{Name: "id", Value: id}})
}

// GetUserByID get user using by id
// @param id search key
// @param user object to cast
func (db *DB) GetUserByID(id string, user interface{}) error {
	if err := db.Collection[db.TableName].Find(bson.D{{Name: "id", Value: id}}).One(user); err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// GetUser get user using by email
// @param email search key
// @param user object to cast
func (db *DB) GetUser(email string, user interface{}) error {
	if err := db.Collection[db.TableName].Find(bson.D{{Name: "email", Value: email}}).One(user); err != nil {
		return err
	}
	return nil
}

// AddUser saves a given user, assigning it a new ID.
func (db *DB) AddUser(obj interface{}) (id string, err error) {
	if err := db.Collection[db.TableName].Insert(obj); err != nil {
		return "", errors.New(err.Error())
	}
	return id, nil
}

// UpdateUser -
func (db *DB) UpdateUser(id string, user interface{}) error {
	return db.Collection[db.TableName].Update(bson.D{{Name: "id", Value: id}}, user)
}

// ListUsers -
func (db *DB) ListUsers(user interface{}) error {
	if err := db.Collection[db.TableName].Find(nil).All(user); err != nil {
		return err
	}
	return nil
}
