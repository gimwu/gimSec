package server

import (
	"context"
	"gimSec/api"
	"gimSec/basic/logging"
	"gimSec/model"
)

type UserProvider struct {
	api.UnimplementedUserServiceServer
}

func (p *UserProvider) GetUserById(ctx context.Context, in *api.IdMessage) (*api.User, error) {
	logging.Info("server GetUserById")
	user, err := model.GetUser(in.GetId())
	if err != nil {
		logging.Error(err.Error())
	}
	return &api.User{
		Id:            user.Id,
		Name:          user.Name,
		Telephone:     user.Telephone,
		Password:      user.Password,
		UserType:      0,
		LastTimeLogin: user.LastTimeLogin.String(),
	}, nil
}

func (p *UserProvider) GetUserPage(page *api.PageInfo, svr api.UserService_GetUserPageServer) error {
	return nil
}
