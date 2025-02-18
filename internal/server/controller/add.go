package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nantyatte_stay_watch/pkg/commands"
	"net/http"
)

// Add 自身のMACアドレスを追加
func Add(c *gin.Context) {
	//web経由でアクセスしてきたIPアドレスを取得
	ipFromWeb := c.ClientIP()

	//IPアドレスとMACアドレスの対応表を取得
	deviceInfos := commands.GetDeviceInfos()
	commands.PrintDeviceInfos(deviceInfos)

	//MACアドレスを取得
	macAddr := commands.SearchMacAddr(deviceInfos, ipFromWeb)
	commands.RegistMacAddrList(macAddr)

	//自身のMACアドレスを表示
	fmt.Println("自身の IPアドレス:", ipFromWeb)
	fmt.Println("自身のMACアドレス:", macAddr)

	//HTMLを返す
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
