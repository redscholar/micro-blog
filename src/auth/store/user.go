package store

type User struct {
	Id       string `bson:"_id,"`
	Username string `bson:"username"`
	Password string `bson:"password"`
}

type UserStore interface {
	CreateUser(user User) error
	GetUser(user User) (*User, error)
	UpdateUser(user *User, field ...string) error
}
