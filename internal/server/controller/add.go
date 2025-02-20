package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nantyatte_stay_watch/internal/db"
	. "nantyatte_stay_watch/pkg/model"
	"net/http"
)

// Add 自身のMACアドレスを追加
func Add(c *gin.Context) {
	//web経由でアクセスしてきたIPアドレスを取得
	ipFromWeb := c.ClientIP()

	//ローカル内の全てのデバイス情報を取得
	nowDevices := GetNowDevices()
	nowDevices.Print()

	//IPアドレスをもとにアクセスしてきた端末のDeviceを取得
	targetDevice := nowDevices.SearchByIpAddr(ipFromWeb)
	db.RegistMacAddrList(targetDevice.MacAddr)

	//自身のMACアドレスを表示
	fmt.Println("自身の IPアドレス:", targetDevice.IpAddr)
	fmt.Println("自身のMACアドレス:", targetDevice.MacAddr)

	//HTMLを返す
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
