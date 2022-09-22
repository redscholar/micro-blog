package handler

import (
	"context"
	"go-micro.dev/v4/auth"
	"go-micro.dev/v4/metadata"
	"strconv"
	"time"
	"user/micro"
	pb "user/proto"
)

type User struct{}

func (e *User) Info(ctx context.Context, _ *pb.InfoRequest, rsp *pb.InfoResponse) error {
	md, _ := metadata.FromContext(ctx)
	token, _ := md.Get(micro.AuthHeader)
	account, err := micro.Service.Options().Auth.Inspect(token)
	if err != nil {
		return err
	}
	// store中获取
	//records, err := micro.Service.Options().Store.Read(account.ID)
	//if err != nil {
	//	return err
	//}
	//user := new(store.User)
	//err = json.Unmarshal(records[0].Value, user)
	//if err != nil {
	//	return err
	//}
	rsp.Id = account.ID
	rsp.Username = account.ID
	return nil
}

func (e *User) Login(ctx context.Context, req *pb.LoginRequest, rsp *pb.LoginResponse) error {
	// store中认证
	//records, err := micro.Service.Options().Store.Read(req.GetUsername())
	//if err != nil {
	//	return err
	//}
	//user := new(store.User)
	//err = json.Unmarshal(records[0].Value, user)
	//if err != nil {
	//	return err
	//}
	//if user.Password != req.GetPassword() {
	//	return errors.New("账号密码错误")
	//}
	// 单位秒，token超时时间
	expireTime := micro.Service.Options().Config.Get("auth", "expireTime").Int(0)
	account, err := micro.Service.Options().Auth.Generate(req.GetUsername(), auth.WithType("user"), // todo 应该存id，等做完注册再做
		auth.WithMetadata(map[string]string{
			"createAt": strconv.FormatInt(time.Now().Unix(), 10),
			"expireAt": strconv.FormatInt(time.Now().Add(time.Second*time.Duration(expireTime)).Unix(), 10),
		}))
	if err != nil {
		return err
	}
	token, err := micro.Service.Options().Auth.Token(auth.WithExpiry(time.Second*time.Duration(expireTime)), auth.WithCredentials(account.ID, account.Secret))
	if err != nil {
		return err
	}
	rsp.Token = token.AccessToken
	return nil
}
