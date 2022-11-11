package handler

import (
	pbarticle "article/proto/article"
	pbauth "article/proto/auth"
	"article/store"
	"context"
	"github.com/google/uuid"
	"go-micro.dev/v4"
	"time"
	"util"
)

func NewArticle(service micro.Service, store store.ArticleStore) *articleHandler {
	return &articleHandler{
		service,
		store,
	}
}

type articleHandler struct {
	svc micro.Service
	as  store.ArticleStore
}

func (a articleHandler) List(ctx context.Context, request *pbarticle.ArticleListRequest, response *pbarticle.ArticleListResponse) error {
	articles, count, err := a.as.PageArticle(request.Page, request.Limit, request.LastId, request.Keyword)
	if err != nil {
		util.LoggerHelper(ctx).Errorf("article List error:%v", err)
		return err
	}
	response.Total = count
	response.Data = make([]*pbarticle.ArticleListResponse_Data, len(articles))

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

func (a articleHandler) Create(ctx context.Context, request *pbarticle.ArticleCreateRequest, _ *pbarticle.ArticleCreateResponse) error {
	infoResp := &pbauth.AuthInfoResponse{}
	err := a.svc.Client().Call(ctx, a.svc.Client().NewRequest("auth", "Auth.Info", &pbauth.AuthInfoRequest{}), infoResp)
	if err != nil {
		util.LoggerHelper(ctx).Errorf("article Create get current user error:%v", err)
		return err
	}
	err = a.as.CreateArticle(&store.Article{
		Id:        uuid.New().String(),
		Title:     request.Title,
		Content:   request.Content,
		Image:     request.Image,
		CreatedAt: time.Now(),
		Author: store.ArticleAuthor{
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
