package util

import (
	"context"
	"go-micro.dev/v4"
	"go-micro.dev/v4/auth"
	"strconv"
	"time"
)

func GenToken(srv micro.Service, userId string) (string, error) {
	expireTime := srv.Options().Config.Get("auth", "expireTime").Int(0)
	newAccount, err := srv.Options().Auth.Generate(userId, auth.WithType("user"),
		auth.WithMetadata(map[string]string{
			"createAt": strconv.FormatInt(time.Now().Unix(), 10),
			"expireAt": strconv.FormatInt(time.Now().Add(time.Second*time.Duration(expireTime)).Unix(), 10),
		}))
	if err != nil {
		return "", err
	}
	newToken, err := srv.Options().Auth.Token(auth.WithExpiry(time.Second*time.Duration(expireTime)), auth.WithCredentials(newAccount.ID, newAccount.Secret))
	if err != nil {
		return "", err
	}
	return newToken.AccessToken, nil
}

type account struct{}

func SaveAccount(ctx context.Context, ac *auth.Account) context.Context {
	return context.WithValue(ctx, account{}, ac)
}

func GetAccount(ctx context.Context) *auth.Account {
	return ctx.Value(account{}).(*auth.Account)
}
