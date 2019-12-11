package mongo

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	// MongoDBErrTableName table name not set return error
	MongoDBErrTableName = "mongodb: table not set"
)

// TableName holds table name type as string
type TableName string

func (tn TableName) length() int {
	return len(tn)
}

// DB hold fields data that need in connection for database
type DB struct {
	Conn         *mgo.Session
	DatabaseName string
	tablename    TableName
	Collection   map[TableName]*mgo.Collection
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
		Conn:         conn,
		DatabaseName: database,
		Collection:   make(map[TableName]*mgo.Collection),
	}, nil
}

// Close mongo database.
func (db *DB) Close() {
	db.Conn.Close()
}

// AddTable add table in mongo database
func (db *DB) AddTable(tableName string) {
	tb := TableName(tableName)
	db.Collection[tb] = db.Conn.DB(db.DatabaseName).C(string(tb))
}

// Table query specific table
func (db *DB) Table(name string) *DB {
	db.tablename = TableName(name)
	return db
}

// Delete delete item in database
func (db *DB) Delete(key string, value string) error {
	return db.Collection[db.tablename].Remove(bson.D{{Name: key, Value: key}})
}

// Get get item using by id
// @param key search key
// @param item object to cast
func (db *DB) Get(key string, value string, item interface{}) error {
	if db.tablename.length() <= 0 {
		return errors.New(MongoDBErrTableName)
	}
	if err := db.Collection[db.tablename].Find(bson.D{{Name: key, Value: value}}).One(item); err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// Add saves a given item, assigning it a new ID.
func (db *DB) Add(item interface{}) (id string, err error) {
	if db.tablename.length() <= 0 {
		return "", errors.New(MongoDBErrTableName)
	}
	if err := db.Collection[db.tablename].Insert(item); err != nil {
		return "", errors.New(err.Error())
	}
	return id, nil
}

// Update update fields in database
func (db *DB) Update(key string, value string, item interface{}) error {
	if db.tablename.length() <= 0 {
		return errors.New(MongoDBErrTableName)
	}
	return db.Collection[db.tablename].Update(bson.D{{Name: key, Value: value}}, item)
}

// List view list of data
func (db *DB) List(items interface{}) error {
	if db.tablename.length() <= 0 {
		return errors.New(MongoDBErrTableName)
	}
	if err := db.Collection[db.tablename].Find(nil).All(items); err != nil {
		return err
	}
	return nil
}
