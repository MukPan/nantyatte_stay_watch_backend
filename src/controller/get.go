package controller

import (
	"github.com/gin-gonic/gin"
	"nantyatte_stay_watch/cmd"
	"net/http"
)

func Get(c *gin.Context) {
	//登録済みのIPアドレスリストを取得
	registeredIpAddrList := cmd.GetIpAddrList()

	//接続中のIPアドレスを格納するマップを取得
	connectingIpAddrMap := cmd.SendPingAll(registeredIpAddrList)
	cmd.PrintConnectingIpAddrMap(connectingIpAddrMap)

	//HTMLを返す
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
