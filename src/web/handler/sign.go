package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"web/micro"
	"web/proto/user"
)

//SignInReq 登录表单结构体
type SignInReq struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func SignInPOST(c *gin.Context) {
	var signIn SignInReq
	//获取登录参数
	err := c.ShouldBind(&signIn)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	//获取用户名,密码
	username := signIn.Username
	password := signIn.Password
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
	loginRsp := &user.AuthSignInResponse{}
	ctx, _ := c.Get("clientCtx")
	if err := micro.Service.Options().Client.Call(ctx.(context.Context), micro.Service.Options().Client.NewRequest("user", "Auth.SignIn", &user.AuthSignInRequest{Username: username, Password: password}), loginRsp); err != nil {
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

//SignUpReq 登录表单结构体
type SignUpReq struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func SignUpPOST(c *gin.Context) {
	var signUp SignUpReq
	//获取登录参数
	err := c.ShouldBind(&signUp)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	//获取用户名,密码
	username := signUp.Username
	password := signUp.Password
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
	loginRsp := &user.AuthSignUpResponse{}
	ctx, _ := c.Get("clientCtx")
	if err := micro.Service.Options().Client.Call(ctx.(context.Context), micro.Service.Options().Client.NewRequest("user", "Auth.SignUp", &user.AuthSignUpRequest{Username: username, Password: password}), loginRsp); err != nil {
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
