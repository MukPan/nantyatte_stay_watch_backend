package router

import (
	"github.com/gin-gonic/gin"
	"nantyatte_stay_watch/src/controller"
)

func Init() {
	engine := gin.Default()
	// htmlのディレクトリを指定
	engine.LoadHTMLGlob("templates/*")
	//home
	engine.GET("/", controller.Home)

	//10秒ごとにgoルーチンを実行
	//go execPing()

	//サーバ起動
	engine.Run(":3000")
}
