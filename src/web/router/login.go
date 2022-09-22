package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"web/micro"
	user "web/proto/user"
)

//LoginForm 登录表单结构体
type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

//UserInfo session中的用户信息结构体
type UserInfo struct {
	UserName       string
	ExpirationTime string
}

//LoginGET 登录页面GET请求
func LoginGET(ctx *gin.Context) {
	ctx.HTML(200, "user/login.html", gin.H{
		"message": "welcome!",
	})

}

//LoginPOST 登录页面POST请求
func LoginPOST(c *gin.Context) {
	//session := sessions.Default(c)
	//db := common.GetDB()
	var login LoginForm
	//获取登录参数
	err := c.ShouldBind(&login)
	if err != nil {
		c.HTML(422, "user/login.html", gin.H{
			"message": err.Error(),
		})
		return
	}
	//获取用户名,密码
	username := login.Username
	password := login.Password
	//判断数据长度 5<username<=20 5<password<=20 phone = 11
	if len(username) <= 5 && len(username) > 20 {
		c.HTML(422, "user/login.html", gin.H{
			"message": "用户名范围为:5-20!",
		})
		return
	}
	if len(password) <= 5 && len(password) > 20 {
		c.HTML(422, "user/login.html", gin.H{
			"message": "密码长度范围为:5-20!",
		})
		return
	}
	loginRsp := &user.LoginResponse{}
	if err := micro.Service.Options().Client.Call(context.Background(), micro.Service.Options().Client.NewRequest("user", "User.Login", &user.LoginRequest{Username: username, Password: password}), loginRsp); err != nil {
		c.HTML(422, "user/login.html", gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"token": loginRsp.GetToken(),
	})
	//c.HTML(20,"user/login.html",gin.H{})
	//c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/index?token=%v", loginRsp.GetToken()))
}
