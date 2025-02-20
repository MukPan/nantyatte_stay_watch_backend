package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "nantyatte_stay_watch/pkg/model"
	"net/http"
)

func Get(c *gin.Context) {
	//登録済みのDeviceリストを取得
	registerdDevices := GetRegisteredDevices()

	//接続中のIPアドレスを格納するマップを取得
	connectingDeviceMap := registerdDevices.SendPingAll()

	//コンソールに出力
	for device, isConnecting := range connectingDeviceMap {
		device.Print()
		fmt.Println("isConn:", isConnecting)
	}

	//HTMLを返す
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
