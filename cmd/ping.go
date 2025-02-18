package cmd

import (
	"fmt"
	"os/exec"
	"sync"
)

func SendPingAll(ipAddrList []string) (connectingIpAddrMap map[string]bool) {
	//IPアドレスが接続しているか否かを格納するマップ
	connectingIpAddrMap = make(map[string]bool)

	//データ送受信用チャンネルを作成
	var wg = sync.WaitGroup{}
	wg.Add(len(ipAddrList)) //IPアドレスの数だけgoroutineを立てる

	//アクセスするごとにgoroutineを立てpingを送信
	for _, ipAddr := range ipAddrList {
		go sendPing(ipAddr, connectingIpAddrMap, &wg)
	}

	//全てのgoroutineが終了するまで待機
	wg.Wait()

	return connectingIpAddrMap
}

// pingの送信が成功したかを返す関数
func sendPing(targetIpAddr string, connectingIpAddrMap map[string]bool, wg *sync.WaitGroup) {
	pingCmd := exec.Command("sh", "-c",
		fmt.Sprintf("ping %s -o -c 1", targetIpAddr))
	out, err := pingCmd.CombinedOutput()

	//確認用
	fmt.Println(string(out))

	//結果をマップに代入
	isConnecting := err == nil
	connectingIpAddrMap[targetIpAddr] = isConnecting

	//終了報告
	wg.Done()
}

func PrintConnectingIpAddrMap(connectingIpAddrMap map[string]bool) {
	fmt.Println("--接続中のIPアドレス--")
	for ipAddr, isConnecting := range connectingIpAddrMap {
		fmt.Println("IP:", ipAddr, "接続中:", isConnecting)
	}
}
