package mongo

import (
	"article/store"
	"context"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var WireMongoArticleStore = wire.NewSet(NewArticleStore, wire.Bind(new(store.ArticleStore), new(*articleStore)))

func NewArticleStore() *articleStore {
	return &articleStore{
		mongoCollection: initMongoStore(),
		Context:         context.Background(),
	}
}

type articleStore struct {
	*mongoCollection
	context.Context
}

func (a articleStore) CreateArticle(article *store.Article) error {
	_, err := a.mongoCollection.articleCollection.InsertOne(a.Context, article)
	return err
}

func (a articleStore) PageArticle(page, limit int64, id, keyword string) ([]*store.Article, int64, error) {
	skip := (page - 1) * limit
	print(skip)
	filer := bson.M{
		"_id": bson.M{"$gt": id},
		"$or": []bson.M{
			{"title": bson.M{"$regex": keyword}},
			{"content": bson.M{"$regex": keyword}},
		},
	}
	count, err := a.mongoCollection.articleCollection.CountDocuments(a.Context, filer)
	if err != nil {
		return nil, 0, err
	}
	cur, err := a.mongoCollection.articleCollection.Find(a.Context, filer, &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
		Sort:  bson.M{"_id": 1},
	})
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(a.Context)
	results := make([]*store.Article, 0)
	for cur.Next(a.Context) {
		r := new(store.Article)
		if err := cur.Decode(r); err != nil {
			return nil, 0, err
		}
		results = append(results, r)
	}
	return results, count, nil
}
