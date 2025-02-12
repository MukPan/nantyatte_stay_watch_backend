package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"sort"
	"sync"
)

// 接続中のIPアドレスを格納するマップ
var connectingIpAddrMap = make(map[string]bool)

// Home ホーム画面
func Home(c *gin.Context) {
	//接続時に自IPアドレスを取得
	myIp := c.ClientIP()

	//接続してきた新規IPアドレスをリストに登録
	connectingIpAddrMap[myIp] = true //自分自身は接続中(多分falseでもいい)

	//データ送受信用チャンネルを作成
	var wg = sync.WaitGroup{}
	wg.Add(len(connectingIpAddrMap)) //IPアドレスの数だけgoroutineを立てる

	//アクセスするごとにgoroutineを立てpingを送信
	for ipAddr, _ := range connectingIpAddrMap {
		go sendPing(ipAddr, &wg)
	}

	//全てのgoroutineが終了するまで待機
	wg.Wait()
	//出力用resultを作成
	printMap()
	result := sprintIpAddrMap()

	//HTMLを返す
	c.HTML(http.StatusOK, "index.html", gin.H{
		"myIp":   myIp,
		"result": result,
	})
}

// pingの送信が成功したか
func sendPing(targetIpAddr string, wg *sync.WaitGroup) {
	pingCmd := exec.Command("sh", "-c",
		fmt.Sprintf("ping %s -o -c 3", targetIpAddr))
	_, err := pingCmd.CombinedOutput()

	//結果をマップに代入
	isConnecting := err == nil
	connectingIpAddrMap[targetIpAddr] = isConnecting

	//終了報告
	wg.Done()
}

// マップを文字列に変換
func sprintIpAddrMap() string {
	var result string
	//keyをソートする
	keys := make([]string, 0, len(connectingIpAddrMap))
	for key := range connectingIpAddrMap {
		keys = append(keys, key)
	}
	//逆ソート
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))

	//すべてのIPアドレスのRowを作成
	for _, key := range keys {
		//状態変数
		var isConnectingStr string
		//接続中か未接続かを判定
		if connectingIpAddrMap[key] {
			isConnectingStr = "接続中"
		} else {
			isConnectingStr = "未接続"
		}
		//結果を追加
		result += fmt.Sprintln("IPアドレス→", key, "| 状態→", isConnectingStr, "<br>")
	}
	return result
}

// マップの表示
func printMap() {
	for key, value := range connectingIpAddrMap {
		fmt.Println("key:", key, "value:", value)
	}
}
