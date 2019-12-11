package postgres

// Datastore -
type Datastore interface {
	AddTable(tableName string)
	Table(name string) Datastore
	Delete(key string, value string) error
	Get(key string, value string, item interface{}) error
	Add(item interface{}) (id string, err error)
	Update(key string, value string, item interface{}) error
	List(item interface{}) error
}
