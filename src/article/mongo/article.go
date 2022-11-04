package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Article struct {
	Id        string        `bson:"_id"`
	Title     string        `bson:"title"`
	Content   string        `bson:"content"`
	Image     string        `bson:"image"`
	CreatedAt time.Time     `bson:"created_at"`
	Author    ArticleAuthor `bson:"author"`
}

type ArticleAuthor struct {
	Id       string `bson:"_id"`
	Username string `bson:"username"`
}

func NewArticleStore() *ArticleStore {
	return &ArticleStore{
		articleCollection: mongoStore.ArticleCollection,
		Context:           context.Background(),
	}
}

type ArticleStore struct {
	articleCollection *mongo.Collection
	context.Context
}

func (a ArticleStore) CreateArticle(article *Article) error {
	_, err := a.articleCollection.InsertOne(a.Context, article)
	return err
}

func (a ArticleStore) PageArticle(page, limit int64, id, keyword string) ([]*Article, int64, error) {
	skip := (page - 1) * limit
	filer := bson.M{
		"_id": bson.M{"$gt": id},
		"$or": bson.M{
			"title":   bson.M{"$regex": fmt.Sprintf("/%v/", keyword)},
			"content": bson.M{"$regex": fmt.Sprintf("/%v/", keyword)},
		},
	}
	count, err := a.articleCollection.CountDocuments(a.Context, filer)
	if err != nil {
		return nil, 0, err
	}
	cur, err := a.articleCollection.Find(a.Context, &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
		Sort:  "_id",
	})
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(a.Context)
	results := make([]*Article, 0)
	for cur.Next(a.Context) {
		r := new(Article)
		if err := cur.Decode(r); err != nil {
			return nil, 0, err
		}
		results = append(results, r)
	}
	return results, count, nil
}
