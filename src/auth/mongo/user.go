package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"strings"
)

type User struct {
	Id       string `bson:"_id,"`
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
	result := new(User)
	query := bson.M{}
	t := reflect.TypeOf(user)
	v := reflect.ValueOf(user)
	for i := 0; i < t.NumField(); i++ {
		if v.Field(i).String() != "" {
			bname := strings.Split(t.Field(i).Tag.Get("bson"), ",")[0]
			if bname == "" {
				bname = strings.ToLower(t.Field(i).Name[:1]) + t.Field(i).Name[1:]
			}
			query[bname] = v.Field(i).String()
		}
	}
	return result, m.userCollection.FindOne(m.Context, query).Decode(result)
}

func (m UserStore) UpdateUser(user *User, field ...string) error {
	set := bson.M{}
	t := reflect.TypeOf(*user)
	v := reflect.ValueOf(*user)
	for _, s := range field {
		if f, e := t.FieldByName(s); e {
			bname := strings.Split(f.Tag.Get("bson"), ",")[0]
			if bname == "" {
				bname = strings.ToLower(f.Name[:1]) + f.Name[1:]
			}
			set[bname] = v.FieldByName(s).String()
		}
	}
	_, err := m.userCollection.UpdateByID(m.Context, user.Id, bson.M{
		"$set": user,
	})
	return err
}
