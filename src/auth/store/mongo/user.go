package mongo

import (
	"auth/store"
	"context"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"strings"
)

var WireMongoUserStore = wire.NewSet(NewUserStore, wire.Bind(new(store.UserStore), new(*userStore)))

func NewUserStore() *userStore {
	return &userStore{
		mc:      initMongoStore(),
		Context: context.Background(),
	}
}

type userStore struct {
	mc *mongoCollection
	context.Context
}

func (m userStore) CreateUser(user store.User) error {
	_, err := m.mc.articleCollection.InsertOne(m.Context, user)
	return err
}

func (m userStore) GetUser(user store.User) (*store.User, error) {
	result := new(store.User)
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
	return result, m.mc.articleCollection.FindOne(m.Context, query).Decode(result)
}

func (m userStore) UpdateUser(user *store.User, field ...string) error {
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
	_, err := m.mc.articleCollection.UpdateByID(m.Context, user.Id, bson.M{
		"$set": user,
	})
	return err
}
