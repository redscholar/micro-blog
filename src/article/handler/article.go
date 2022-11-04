package handler

import (
	"article/micro"
	"article/mongo"
	pbarticle "article/proto/article"
	pbauth "article/proto/auth"
	"context"
	"github.com/google/uuid"
	"time"
	"util"
)

type Article struct {
	*mongo.ArticleStore
}

func (a Article) List(ctx context.Context, request *pbarticle.ArticleListRequest, response *pbarticle.ArticleListResponse) error {
	articles, count, err := a.ArticleStore.PageArticle(request.Page, request.Limit, request.LastId, request.Keyword)
	if err != nil {
		util.LoggerHelper(ctx).Errorf("article List error:%v", err)
		return err
	}
	response = &pbarticle.ArticleListResponse{
		Total: count,
		Data:  make([]*pbarticle.ArticleListResponse_Data, 0),
	}
	for i, article := range articles {
		response.Data[i] = &pbarticle.ArticleListResponse_Data{
			Id:        article.Id,
			Title:     article.Title,
			Content:   article.Content,
			Image:     article.Image,
			CreatedAt: article.CreatedAt.Format(time.RFC3339),
			Author: &pbarticle.ArticleListResponse_Data_Author{
				Id:       article.Author.Id,
				Username: article.Author.Username,
			},
		}
	}
	return nil
}

func (a Article) Create(ctx context.Context, request *pbarticle.ArticleCreateRequest, _ *pbarticle.ArticleCreateResponse) error {
	infoResp := &pbauth.AuthInfoResponse{}
	err := micro.Service.Client().Call(ctx, micro.Service.Client().NewRequest("auth", "Auth.Info", &pbauth.AuthInfoRequest{}), infoResp)
	if err != nil {
		util.LoggerHelper(ctx).Errorf("article Create get current user error:%v", err)
		return err
	}
	err = a.ArticleStore.CreateArticle(&mongo.Article{
		Id:        uuid.New().String(),
		Title:     request.Title,
		Content:   request.Content,
		Image:     request.Image,
		CreatedAt: time.Now(),
		Author: mongo.ArticleAuthor{
			Id:       infoResp.Id,
			Username: infoResp.Username,
		},
	})
	if err != nil {
		util.LoggerHelper(ctx).Errorf("article Create error:%v", err)
		return err
	}
	return nil
}
