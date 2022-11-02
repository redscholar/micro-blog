package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id       string `bson:"_id"`
	Username string `bson:"username"`
	Password string `bson:"password"`
}

func NewUserStore() *UserStore {
	return &UserStore{
		userCollection: mongoStore.UserCollection,
		Context:        context.Background(),
	}
}

type UserStore struct {
	userCollection *mongo.Collection
	context.Context
}

func (m UserStore) CreateUser(user User) error {
	_, err := m.userCollection.InsertOne(m.Context, user)
	if err != nil {
		return err
	}
	return nil
}

func (m UserStore) GetUser(user User) (*User, error) {
	//filter, err := bson.Marshal(user)
	//if err != nil {
	//	return nil, err
	//}
	result := new(User)
	return result, m.userCollection.FindOne(m.Context, bson.D{{"username", user.Username}}).Decode(result)

}

func (m UserStore) UpdateUser(user *User) error {
	_, err := m.userCollection.UpdateByID(m.Context, user.Id, user)
	return err
}
