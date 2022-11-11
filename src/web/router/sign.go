package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
	"net/http"
	"web/option"
	"web/proto/auth"
	pbauth "web/proto/auth"
)

const (
	authService = "auth"
)

func signRoute(r *gin.Engine, svc micro.Service) *gin.Engine {

	// 登录
	r.POST("/signIn", func(c *gin.Context) {
		var siReq struct {
			Username string `form:"username" binding:"required"`
			Password string `form:"password" binding:"required"`
		}
		//获取登录参数
		err := c.ShouldBind(&siReq)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  err.Error(),
			})
			return
		}
		//获取用户名,密码
		username := siReq.Username
		password := siReq.Password
		//判断数据长度 5<username<=20 5<password<=20 phone = 11
		if len(username) <= 5 && len(username) > 20 {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  err.Error(),
			})
			return
		}
		if len(password) <= 5 && len(password) > 20 {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  err.Error(),
			})
			return
		}
		ctx, _ := c.Get(option.ClientCtx)
		loginRsp := &pbauth.AuthSignInResponse{}
		if err := svc.Options().Client.Call(ctx.(context.Context), svc.Options().Client.NewRequest(authService, "Auth.SignIn", &auth.AuthSignInRequest{Username: username, Password: password}), loginRsp); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": loginRsp.GetToken(),
		})
	})
	r.POST("/signUp", func(c *gin.Context) {
		var suReq struct {
			Username string `form:"username" binding:"required"`
			Password string `form:"password" binding:"required"`
		}
		//获取登录参数
		err := c.ShouldBind(&suReq)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  err.Error(),
			})
			return
		}
		//获取用户名,密码
		username := suReq.Username
		password := suReq.Password
		//判断数据长度 5<username<=20 5<password<=20 phone = 11
		if len(username) <= 5 && len(username) > 20 {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  err.Error(),
			})
			return
		}
		if len(password) <= 5 && len(password) > 20 {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  err.Error(),
			})
			return
		}
		ctx, _ := c.Get(option.ClientCtx)
		signUpRsp := &auth.AuthSignUpResponse{}
		if err := svc.Options().Client.Call(ctx.(context.Context), svc.Options().Client.NewRequest(authService, "Auth.SignUp", &auth.AuthSignUpRequest{Username: username, Password: password}), signUpRsp); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": signUpRsp.GetToken(),
		})
	})
	r.POST("/info", func(c *gin.Context) {
		ctx, _ := c.Get(option.ClientCtx)
		infoRsp := new(auth.AuthInfoResponse)
		if err := svc.Options().Client.Call(ctx.(context.Context), svc.Options().Client.NewRequest(authService, "Auth.Info", new(auth.AuthInfoRequest)), infoRsp); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": map[string]string{
				"id":       infoRsp.Id,
				"username": infoRsp.Username,
			},
		})
	})
	r.POST("/changePwd", func(c *gin.Context) {
		var cpReq struct {
			OldPwd string `binding:"required"`
			NewPwd string `binding:"required"`
		}
		err := c.ShouldBindJSON(&cpReq)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  err.Error(),
			})
			return
		}
		ctx, _ := c.Get(option.ClientCtx)
		changePwdRsp := &auth.AuthChangePwdResponse{}
		if err := svc.Options().Client.Call(ctx.(context.Context), svc.Options().Client.NewRequest(authService, "Auth.ChangePwd", &auth.AuthChangePwdRequest{OldPwd: cpReq.OldPwd, NewPwd: cpReq.NewPwd}), changePwdRsp); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
		})
	})

	return r
}
