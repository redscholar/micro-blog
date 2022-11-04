package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	pb "article/proto/article"
	"context"
	"web/graph/generated"
	"web/graph/model"
	"web/micro"
)

// CreateArticle is the resolver for the createArticle field.
func (r *mutationResolver) CreateArticle(ctx context.Context, request *model.CreateArticleRequest) (*model.CreateArticleResponse, error) {
	microReq := micro.Service.Client().NewRequest(articleService, "Article.Create", &pb.ArticleCreateRequest{
		Title:   request.Title,
		Content: request.Content,
		Image:   request.Image,
	})
	microResp := &pb.ArticleCreateResponse{}
	clientCtx, _ := r.Context.Get(micro.ClientCtx)
	err := micro.Service.Client().Call(clientCtx.(context.Context), microReq, microResp)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// ListArticle is the resolver for the listArticle field.
func (r *queryResolver) ListArticle(ctx context.Context, request *model.ListArticleRequest) (*model.ListArticleResponse, error) {
	microReq := micro.Service.Client().NewRequest(articleService, "Article.List", &pb.ArticleListRequest{
		Keyword: request.Keyword,
		LastId:  request.LastID,
		Page:    int64(request.Pagination.Page),
		Limit:   int64(request.Pagination.Limit),
	})
	microResp := &pb.ArticleListResponse{}
	clientCtx, _ := r.Context.Get(micro.ClientCtx)
	err := micro.Service.Client().Call(clientCtx.(context.Context), microReq, microResp)
	if err != nil {
		return nil, err
	}
	resp := &model.ListArticleResponseData{
		Total:    int(microResp.Total),
		Articles: make([]*model.Article, 0),
	}
	for i, datum := range microResp.Data {
		resp.Articles[i] = &model.Article{
			ID:      datum.Id,
			Title:   &datum.Title,
			Content: &datum.Content,
			Image:   &datum.Image,
		}
	}
	return nil, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
const articleService = "article"
