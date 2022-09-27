package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"go-micro.dev/v4/store"
	"user/micro"
	pb "user/proto"
	sm "user/store"
	"util"
)

type Auth struct {
}

func (a Auth) ChangePwd(ctx context.Context, request *pb.AuthChangePwdRequest, _ *pb.AuthChangePwdResponse) error {
	account := util.GetAccount(ctx)
	records, err := micro.Service.Options().Store.Read(account.ID)
	if err != nil {
		util.LoggerHelper(ctx).Errorf("get records from store error:%v", err)
		return err
	}
	if len(records) == 0 {
		return errors.New("用户不存在")
	}
	cuser := new(sm.User)
	err = json.Unmarshal(records[0].Value, cuser)
	if err != nil {
		util.LoggerHelper(ctx).Errorf("convert record to user error:%v", err)
		return err
	}
	if cuser.Password != request.OldPwd {
		return errors.New("旧密码错误")
	}

	// 替换密码
	data, err := json.Marshal(&sm.User{
		Id:       cuser.Id,
		Username: cuser.Username,
		Password: request.NewPwd,
	})
	if err != nil {
		util.LoggerHelper(ctx).Errorf("convert user to json error:%v", err)
		return err
	}
	if err = micro.Service.Options().Store.Write(&store.Record{
		Key:      cuser.Id,
		Value:    data,
		Metadata: nil,
		Expiry:   -1,
	}); err != nil {
		util.LoggerHelper(ctx).Errorf("save user to store error:%v", err)
		return err
	}
	return nil
}

func (a Auth) SignIn(ctx context.Context, request *pb.AuthSignInRequest, response *pb.AuthSignInResponse) error {
	// store中认证
	records, err := micro.Service.Options().Store.Read("", store.ReadPrefix())
	if err != nil {
		util.LoggerHelper(ctx).Errorf("get record from store error:%v", err)
		return err
	}
	var cuser = new(sm.User)
	for _, record := range records {
		user := new(sm.User)
		err = json.Unmarshal(record.Value, user)
		if err != nil {
			util.LoggerHelper(ctx).Errorf("convert json record to user error:%v", err)
			return err
		}
		if user.Username == request.GetUsername() && user.Password == request.GetPassword() {
			cuser = user
			break
		}
	}
	if cuser.Id == "" {
		return errors.New("账号密码错误")
	}

	// 单位秒，token超时时间
	response.Token, err = util.GenToken(micro.Service, cuser.Id)
	if err != nil {
		return err
	}
	return nil
}

func (a Auth) SignUp(ctx context.Context, request *pb.AuthSignUpRequest, response *pb.AuthSignUpResponse) error {
	// 检查账号是否重复
	records, err := micro.Service.Options().Store.Read("", store.ReadPrefix())
	if err != nil {
		util.LoggerHelper(ctx).Errorf("get record from store error:%v", err)
		return err
	}
	for _, record := range records {
		user := new(sm.User)
		err = json.Unmarshal(record.Value, user)
		if err != nil {
			util.LoggerHelper(ctx).Errorf("convert json record to user error:%v", err)
			return err
		}
		if user.Username == request.GetUsername() {
			return errors.New("账号已注册")
		}
	}
	id := uuid.New().String()
	data, err := json.Marshal(&sm.User{
		Id:       id,
		Username: request.Username,
		Password: request.Password,
	})
	if err != nil {
		util.LoggerHelper(ctx).Errorf("convert user to json error:%v", err)
		return err
	}
	if err = micro.Service.Options().Store.Write(&store.Record{
		Key:      id,
		Value:    data,
		Metadata: nil,
		Expiry:   -1,
	}); err != nil {
		util.LoggerHelper(ctx).Errorf("save user to store error:%v", err)
		return err
	}
	if response.Token, err = util.GenToken(micro.Service, id); err != nil {
		util.LoggerHelper(ctx).Errorf("generate new token error:%v", err)
		return err
	}
	return nil
}

func (a Auth) Info(ctx context.Context, _ *pb.AuthInfoRequest, response *pb.AuthInfoResponse) error {
	account := util.GetAccount(ctx)
	records, err := micro.Service.Options().Store.Read(account.ID)
	if err != nil {
		util.LoggerHelper(ctx).Errorf("get records from store error:%v", err)
		return err
	}
	if len(records) == 0 {
		return errors.New("用户不存在")
	}
	if err := json.Unmarshal(records[0].Value, response); err != nil {
		util.LoggerHelper(ctx).Errorf("convert json record to info error:%v", err)
		return err
	}
	return nil
}
