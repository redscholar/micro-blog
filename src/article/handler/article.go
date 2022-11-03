package handler

import (
	"article/mongo"
	pb "article/proto"
	"context"
	"github.com/google/uuid"
	"time"
	"util"
)

type Article struct {
	*mongo.ArticleStore
}

func (a Article) List(ctx context.Context, request *pb.ArticleListRequest, response *pb.ArticleListResponse) error {
	articles, count, err := a.ArticleStore.PageArticle(request.Page, request.Limit, request.LastId, request.Keyword)
	if err != nil {
		util.LoggerHelper(ctx).Errorf("article List error:%v", err)
		return err
	}
	response = &pb.ArticleListResponse{
		Total: count,
		Data:  make([]*pb.ArticleListResponse_Data, 0),
	}
	for i, article := range articles {
		response.Data[i] = &pb.ArticleListResponse_Data{
			Id:      article.Id,
			Title:   article.Title,
			Content: article.Content,
			Image:   article.Image,
		}
	}
	return nil
}

func (a Article) Create(ctx context.Context, request *pb.ArticleCreateRequest, _ *pb.ArticleCreateResponse) error {
	err := a.ArticleStore.CreateArticle(&mongo.Article{
		Id:       uuid.New().String(),
		Title:    request.Title,
		Content:  request.Content,
		Image:    request.Image,
		CreateAt: time.Now(),
	})
	if err != nil {
		util.LoggerHelper(ctx).Errorf("article Create error:%v", err)
		return err
	}
	return nil
}
