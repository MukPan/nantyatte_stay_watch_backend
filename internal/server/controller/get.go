package controller

import (
	"github.com/gin-gonic/gin"
	"nantyatte_stay_watch/pkg/commands"
	"net/http"
)

func Get(c *gin.Context) {
	//登録済みのIPアドレスリストを取得
	registerdDeviceInfos := commands.GetRegisterdDeviceInfos()

	//接続中のIPアドレスを格納するマップを取得
	connectingDeviceInfosMap := commands.SendPingAll(registerdDeviceInfos)
	commands.PrintConnectingDeviceInfosMap(connectingDeviceInfosMap)

	//HTMLを返す
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
