package mongo

// Datastore -
type Datastore interface {
	DeleteUser(id string) error
	GetUserByID(id string) error
	GetUser(email string, user interface{}) error
	AddUser(user interface{}) (id string, err error)
	UpdateUser(user interface{}) error
	ListUsers(users interface{}) error
}
