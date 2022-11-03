package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"web/micro"
	"web/proto/auth"
)

const (
	authService = "auth"
)

func SignInPOST(c *gin.Context) {
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
	ctx, _ := c.Get(micro.ClientCtx)
	loginRsp := &auth.AuthSignInResponse{}
	if err := micro.Service.Options().Client.Call(ctx.(context.Context), micro.Service.Options().Client.NewRequest(authService, "Auth.SignIn", &auth.AuthSignInRequest{Username: username, Password: password}), loginRsp); err != nil {
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
}

func SignUpPOST(c *gin.Context) {
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
	ctx, _ := c.Get(micro.ClientCtx)
	signUpRsp := &auth.AuthSignUpResponse{}
	if err := micro.Service.Options().Client.Call(ctx.(context.Context), micro.Service.Options().Client.NewRequest(authService, "Auth.SignUp", &auth.AuthSignUpRequest{Username: username, Password: password}), signUpRsp); err != nil {
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
}

func InfoPOST(c *gin.Context) {
	ctx, _ := c.Get(micro.ClientCtx)
	infoRsp := new(auth.AuthInfoResponse)
	if err := micro.Service.Options().Client.Call(ctx.(context.Context), micro.Service.Options().Client.NewRequest(authService, "Auth.Info", new(auth.AuthInfoRequest)), infoRsp); err != nil {
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
}

func ChangePwdPOST(c *gin.Context) {
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
	ctx, _ := c.Get(micro.ClientCtx)
	changePwdRsp := &auth.AuthChangePwdResponse{}
	if err := micro.Service.Options().Client.Call(ctx.(context.Context), micro.Service.Options().Client.NewRequest(authService, "Auth.ChangePwd", &auth.AuthChangePwdRequest{OldPwd: cpReq.OldPwd, NewPwd: cpReq.NewPwd}), changePwdRsp); err != nil {
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
}
