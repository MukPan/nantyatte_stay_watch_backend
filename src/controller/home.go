package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nantyatte_stay_watch/cmd"
	"net/http"
)

// 接続中のMACアドレスを格納するマップ
var connectingMacAddrMap = make(map[string]bool)

// Home ホーム画面(仮登録画面、アクセスしたら自動的にMACアドレス登録)
func Home(c *gin.Context) {
	//web経由でアクセスしてきたIPアドレスを取得
	ipFromWeb := c.ClientIP()

	//IPアドレスとMACアドレスの対応表を取得
	deviceInfos := cmd.GetDeviceInfos()
	cmd.PrintDeviceInfos(deviceInfos)

	//MACアドレスを取得
	macAddr := cmd.SearchMacAddr(deviceInfos, ipFromWeb)

	//自身のMACアドレスを表示
	fmt.Println("自身の IPアドレス:", ipFromWeb)
	fmt.Println("自身のMACアドレス:", macAddr)

	//HTMLを返す
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
