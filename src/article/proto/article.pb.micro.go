// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/article.proto

package article

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Article service

func NewArticleEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Article service

type ArticleService interface {
	List(ctx context.Context, in *ArticleListRequest, opts ...client.CallOption) (*ArticleListResponse, error)
	Create(ctx context.Context, in *ArticleCreateRequest, opts ...client.CallOption) (*ArticleCreateResponse, error)
}

type articleService struct {
	c    client.Client
	name string
}

func NewArticleService(name string, c client.Client) ArticleService {
	return &articleService{
		c:    c,
		name: name,
	}
}

func (c *articleService) List(ctx context.Context, in *ArticleListRequest, opts ...client.CallOption) (*ArticleListResponse, error) {
	req := c.c.NewRequest(c.name, "Article.List", in)
	out := new(ArticleListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *articleService) Create(ctx context.Context, in *ArticleCreateRequest, opts ...client.CallOption) (*ArticleCreateResponse, error) {
	req := c.c.NewRequest(c.name, "Article.Create", in)
	out := new(ArticleCreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Article service

type ArticleHandler interface {
	List(context.Context, *ArticleListRequest, *ArticleListResponse) error
	Create(context.Context, *ArticleCreateRequest, *ArticleCreateResponse) error
}

func RegisterArticleHandler(s server.Server, hdlr ArticleHandler, opts ...server.HandlerOption) error {
	type article interface {
		List(ctx context.Context, in *ArticleListRequest, out *ArticleListResponse) error
		Create(ctx context.Context, in *ArticleCreateRequest, out *ArticleCreateResponse) error
	}
	type Article struct {
		article
	}
	h := &articleHandler{hdlr}
	return s.Handle(s.NewHandler(&Article{h}, opts...))
}

type articleHandler struct {
	ArticleHandler
}

func (h *articleHandler) List(ctx context.Context, in *ArticleListRequest, out *ArticleListResponse) error {
	return h.ArticleHandler.List(ctx, in, out)
}

func (h *articleHandler) Create(ctx context.Context, in *ArticleCreateRequest, out *ArticleCreateResponse) error {
	return h.ArticleHandler.Create(ctx, in, out)
}
