package router

import (
	"github.com/gin-gonic/gin"
	controller2 "nantyatte_stay_watch/internal/server/controller"
)

func Init() {
	router := gin.Default()
	// htmlのディレクトリを指定
	router.LoadHTMLGlob("templates/*")

	// V1の設定
	v1 := router.Group("/api/v1/")

	//自身のMACアドレスを追加
	v1.GET("/add", controller2.Add)

	//全てのMACアドレスに対して応答確認
	v1.GET("/get", controller2.Get)

	//ローカルのIPアドレスを確認
	//v1.GET("/ip", controller.GetLocalIP)

	//10秒ごとにgoルーチンを実行
	//go execPing()

	//サーバ起動
	router.Run(":3000")
}
