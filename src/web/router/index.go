package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/metadata"
	"net/http"
	"web/micro"
	"web/proto/user"
)

//IndexGET 首页
func IndexGET(c *gin.Context) {
	token, _ := c.Get(micro.AuthHeader)
	// 获取登录用户信息
	infoRsp := &user.InfoResponse{}
	if err := micro.Service.Options().Client.Call(metadata.NewContext(context.Background(), map[string]string{micro.AuthHeader: token.(string)}), micro.Service.Options().Client.NewRequest("user", "User.Info", &user.InfoRequest{}), infoRsp); err != nil {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
	// 获取用户文章

	c.Writer.Header().Set(micro.AuthHeader, token.(string))
	//返回数据到HTML
	c.HTML(200, "index/index.html", gin.H{
		"currentUser": infoRsp.Username,
		"msg":         infoRsp.Username + " ,欢迎来到Blog!",
		"style":       "alert alert-success alert-dismissable",
		"token":       token,
		//"articles":    articles,
		//"messages":    messagesNew,
		"currentPage": 1,
	})

}

//
////IndexGETNextPage 首页翻页
//func IndexGETNextPage(ctx *gin.Context) {
//	//获取redis连接
//	//rdb := common.GetRedis()
//	db := common.GetDB()
//	session := sessions.Default(ctx)
//	//获取当前登录用户
//	userinfo := session.Get("currentUser")
//	if userinfo == nil {
//		ctx.Redirect(http.StatusMovedPermanently, "/login")
//		return
//	}
//	userinfoNew := userinfo.(UserInfo)
//	//判断UserInfo数据是否为空
//	if userinfoNew.UserName == "" || userinfoNew.ExpirationTime == "" {
//		ctx.Redirect(http.StatusMovedPermanently, "/login")
//		return
//	}
//	//判断session id中的时间是否过期
//	ExpirationTime := userinfoNew.ExpirationTime
//	CurrentTime := time.Now().Format("2006-01-02 15:04:05")
//	//先把时间字符串格式化成相同的时间类型
//	t1, err1 := time.Parse("2006-01-02 15:04:05", ExpirationTime)
//	t2, err2 := time.Parse("2006-01-02 15:04:05", CurrentTime)
//	if err1 == nil && err2 == nil && t1.Before(t2) {
//		//session失效，清空session，UserInfo 重定向到login页面
//		session.Delete("currentUser")
//		session.Save()
//		ctx.Redirect(http.StatusTemporaryRedirect, "/login")
//		return
//	}
//	//获取页码参数
//	pageNumber := ctx.Query("pageNumber")
//	//将pageNumber转换为int
//	pageNumberInt, err := strconv.Atoi(pageNumber)
//	if err != nil {
//		fmt.Println("参数错误:" + err.Error())
//		return
//	}
//	//定义接收mysql article，messages数据的slice
//	var articles []models.Articles
//	var messages []models.MessageBoard
//	//查询数据库
//	db.Limit(5).Offset((pageNumberInt - 1) * 5).Find(&articles)
//	db.Limit(5).Offset(0).Order("created_at desc").Find(&messages)
//
//	//根据messages切片中的IfAnonymous 字段是否为true,然后复制到新切片
//	var messagesNew []models.MessageBoard
//	for _, val := range messages {
//		if val.IfAnonymous == true {
//			val.PostUser = "****"
//		}
//		messagesNew = append(messagesNew, val)
//	}
//	//返回数据到HTML
//	ctx.HTML(200, "index/index.html", gin.H{
//		"currentUser": userinfoNew.UserName,
//		"msg":         userinfoNew.UserName + " ,欢迎来到Blog!",
//		"style":       "alert alert-success alert-dismissable",
//		"articles":    articles,
//		"messages":    messagesNew,
//		"currentPage": pageNumberInt,
//	})
//
//}
//
//// IndexMessageBoard 首页留言板
//func IndexMessageBoard(ctx *gin.Context) {
//	db := common.GetDB()
//	var SendMessage index.SendMessageBoard
//	//获取post参数
//	err := ctx.ShouldBind(&SendMessage)
//	if err != nil {
//		fmt.Println(err.Error())
//		ctx.Redirect(http.StatusMovedPermanently, "/index")
//		return
//	}
//	//获取留言用户，留言信息，是否匿名参数
//	username := SendMessage.Username
//	content := SendMessage.MessageContent
//	anonymous := SendMessage.IfAnonymous
//	//判断数据长度 5<username<=20 5<password<=20 phone = 11
//	if len(username) <= 5 && len(username) > 20 {
//		fmt.Println("用户名长度范围为:5-20!")
//		ctx.Redirect(http.StatusMovedPermanently, "/index")
//		return
//	}
//	if len(content) > 600 {
//		fmt.Println("留言信息长度最大为:200!")
//		ctx.Redirect(http.StatusMovedPermanently, "/index")
//		return
//	}
//	//fmt.Println(anonymous)
//	//写入数据库
//	var Value bool
//	if anonymous == "on" {
//		Value = true
//	} else {
//		Value = false
//	}
//	//写入数据库
//	MessageInfo := models.MessageBoard{
//		PostUser:    username,
//		Content:     content,
//		IfAnonymous: Value}
//	db.Create(&MessageInfo)
//	ctx.Redirect(http.StatusMovedPermanently, "/index")
//}
//
//// IndexMessageDelete 首页留言板删除留言信息
//func IndexMessageDelete(ctx *gin.Context) {
//	db := common.GetDB()
//	//获取get请求参数，ID
//	messageId := ctx.Query("id")
//	//将id 由string转为int
//	messageIdInt, err := strconv.Atoi(messageId)
//	if err != nil {
//		fmt.Println("转换ID数据类型错误: " + err.Error())
//		return
//	}
//	//查询留言信息
//	var MessageInfo models.MessageBoard
//	db.First(&MessageInfo, messageIdInt)
//	if MessageInfo.ID == 0 {
//		fmt.Println("未查询到留言信息,无法删除!")
//		return
//	}
//	//执行删除操作
//	db.Delete(&MessageInfo, MessageInfo.ID)
//	ctx.Redirect(http.StatusTemporaryRedirect, "/index")
//}
