package methods

import (
	"context"
	"ozon/internal/models"
	"ozon/internal/usecase"
	"ozon/pkg/api"
)

type Server struct{}

func (s *Server) Get(ctx context.Context, in *api.Request) (*api.Response, error) {
	var (
		link api.Response
		err  models.OwnError
	)
	if models.DB == "redis" {
		link.ResultLink, err = usecase.GetterRedis(in.GetStartLink())
	} else {
		link.ResultLink, err = usecase.GetFullLink(in.GetStartLink())
	}
	if err.Err != nil {
		return &api.Response{ResultLink: err.Message}, err.Err

	}
	return &link, nil
}

func (s *Server) Post(ctx context.Context, in *api.Request) (*api.Response, error) {
	var (
		link api.Response
		err  models.OwnError
	)
	if models.DB == "redis" {
		link.ResultLink, err = usecase.CreationRedis(in.GetStartLink())
	} else {
		link.ResultLink, err = usecase.CreateShortLink(in.GetStartLink())
	}
	if err.Err != nil {
		return &api.Response{ResultLink: err.Message}, err.Err
	}
	return &link, nil
}
