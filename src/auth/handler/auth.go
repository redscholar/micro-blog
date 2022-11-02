package handler

import (
	"auth/micro"
	"auth/mongo"
	pb "auth/proto"
	"context"
	"errors"
	"github.com/google/uuid"
	"util"
)

type Auth struct {
	*mongo.UserStore
}

func (a Auth) ChangePwd(ctx context.Context, request *pb.AuthChangePwdRequest, _ *pb.AuthChangePwdResponse) error {
	account := util.GetAccount(ctx)
	user, err := a.UserStore.GetUser(mongo.User{Id: account.ID})
	if err != nil {
		util.LoggerHelper(ctx).Errorf("get user records from store error:%v", err)
		return err
	}
	if user.Password != request.OldPwd {
		return errors.New("旧密码错误")
	}

	// 替换密码
	user.Password = request.NewPwd
	err = a.UserStore.UpdateUser(user)
	if err != nil {
		util.LoggerHelper(ctx).Errorf("update user to error:%v", err)
		return err
	}
	return nil
}

func (a Auth) SignIn(ctx context.Context, request *pb.AuthSignInRequest, response *pb.AuthSignInResponse) error {
	// store中认证
	user, err := a.UserStore.GetUser(mongo.User{Username: request.Username})
	if err != nil {
		util.LoggerHelper(ctx).Errorf("get record from store error:%v", err)
		return err
	}
	if user.Password != request.Password {
		return errors.New("password is error")
	}
	// 单位秒，token超时时间
	response.Token, err = util.GenToken(micro.Service, user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (a Auth) SignUp(ctx context.Context, request *pb.AuthSignUpRequest, response *pb.AuthSignUpResponse) error {
	// 检查账号是否重复
	_, err := a.UserStore.GetUser(mongo.User{Username: request.Username})
	if err != nil && err.Error() != "mongo: no documents in result" {
		util.LoggerHelper(ctx).Errorf("get record from store error:%v", err)
		return err
	}
	userId := uuid.New().String()
	err = a.UserStore.CreateUser(mongo.User{
		Id:       userId,
		Username: request.Username,
		Password: request.Password,
	})
	if err != nil {
		util.LoggerHelper(ctx).Errorf("create user  error:%v", err)
		return err
	}
	if response.Token, err = util.GenToken(micro.Service, userId); err != nil {
		util.LoggerHelper(ctx).Errorf("generate new token error:%v", err)
		return err
	}
	return nil
}

func (a Auth) Info(ctx context.Context, _ *pb.AuthInfoRequest, response *pb.AuthInfoResponse) error {
	account := util.GetAccount(ctx)
	user, err := a.UserStore.GetUser(mongo.User{Id: account.ID})
	if err != nil {
		util.LoggerHelper(ctx).Errorf("get user records from store error:%v", err)
		return err
	}
	response.Id = user.Id
	response.Username = user.Username
	return nil
}
