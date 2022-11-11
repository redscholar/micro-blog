package store

import "time"

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

type ArticleStore interface {
	CreateArticle(*Article) error
	PageArticle(page, limit int64, id, keyword string) ([]*Article, int64, error)
}
