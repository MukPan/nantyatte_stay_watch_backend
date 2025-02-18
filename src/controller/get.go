package controller

import (
	"github.com/gin-gonic/gin"
	"nantyatte_stay_watch/cmd"
	"net/http"
)

func Get(c *gin.Context) {
	//登録済みのIPアドレスリストを取得
	registerdDeviceInfos := cmd.GetRegisterdDeviceInfos()

	//接続中のIPアドレスを格納するマップを取得
	connectingDeviceInfosMap := cmd.SendPingAll(registerdDeviceInfos)
	cmd.PrintConnectingDeviceInfosMap(connectingDeviceInfosMap)

	//HTMLを返す
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
