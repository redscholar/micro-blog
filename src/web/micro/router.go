package micro

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GinRouter() {
	router := gin.Default()
	router.LoadHTMLGlob("template/**/*")
	router.StaticFS("static/", http.Dir("./static"))
	router.Run(":5001")
}
