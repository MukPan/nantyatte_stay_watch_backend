package commands

import (
	"fmt"
	"os/exec"
	"sync"
)

// SendPingAll 指定した全IPアドレスにpingを送信する関数
func SendPingAll(deviceInfos []DeviceInfo) (connectingDeviceInfosMap map[DeviceInfo]bool) {
	//MACアドレスが接続しているか否かを格納するマップ
	connectingDeviceInfosMap = make(map[DeviceInfo]bool)

	//データ送受信用チャンネルを作成
	var wg = sync.WaitGroup{}
	wg.Add(len(deviceInfos)) //IPアドレスの数だけgoroutineを立てる

	//アクセスするごとにgoroutineを立てpingを送信
	for _, deviceInfo := range deviceInfos {
		go sendPing(deviceInfo, connectingDeviceInfosMap, &wg)
	}

	//全てのgoroutineが終了するまで待機
	wg.Wait()

	return connectingDeviceInfosMap
}

// pingの送信が成功したかを返す関数
func sendPing(targetDeviceInfo DeviceInfo, connectingDeviceInfosMap map[DeviceInfo]bool, wg *sync.WaitGroup) {
	pingCmd := exec.Command("sh", "-c",
		fmt.Sprintf("ping %s -o -c 1", targetDeviceInfo.IpAddr))
	out, err := pingCmd.CombinedOutput()

	//確認用
	fmt.Println(string(out))

	//結果をマップに代入
	isConnecting := err == nil
	connectingDeviceInfosMap[targetDeviceInfo] = isConnecting

	//終了報告
	wg.Done()
}

func PrintConnectingDeviceInfosMap(connectingDeviceInfosMap map[DeviceInfo]bool) {
	fmt.Println("--接続中のIPアドレス--")
	for deviceInfo, isConnecting := range connectingDeviceInfosMap {
		fmt.Println(
			"IP:", deviceInfo.IpAddr,
			"Mac:", deviceInfo.MacAddr,
			"接続中:", isConnecting,
		)
	}
}
