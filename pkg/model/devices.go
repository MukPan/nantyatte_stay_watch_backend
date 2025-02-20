package model

import (
	"fmt"
	"nantyatte_stay_watch/internal/db"
	"nantyatte_stay_watch/pkg/command"
	"strconv"
	"strings"
	"sync"
)

type Devices []Device

// Print 全てのデバイス情報を出力する関数
func (devices Devices) Print() {
	fmt.Println("-- 全てのデバイス情報 --")
	for _, device := range devices {
		device.Print()
	}
}

// SearchByIpAddr IPアドレスからDeviceを取得する関数
func (devices Devices) SearchByIpAddr(ipAddr string) Device {
	for _, device := range devices {
		//一致するIPを見つけたらMACアドレスを返す
		if device.IpAddr == ipAddr {
			return device
		}
	}
	return Device{}
}

// SearchByMacAddr MACアドレスを検索しDeviceを取得する関数
func (devices Devices) SearchByMacAddr(macAddr string) Device {
	for _, device := range devices {
		//一致するIPを見つけたらMACアドレスを返す
		if device.MacAddr == macAddr {
			return device
		}
	}
	return Device{}
}

// GetNowDevices arpコマンドを用いて、現在のローカル内の全てのデバイス情報を取得する関数
func GetNowDevices() Devices {
	//arpリクエストをローカル全体に送りつける
	arpOut := command.ExecArp()
	arpOutLines := strings.Split(arpOut, "\n")

	// デバイス情報を格納するスライスを作成
	devices := make(Devices, 0, 50)

	//すべての行に対して処理
	for i, line := range arpOutLines { //i, v
		//IPアドレスとMACアドレスを抽出
		ipAddr, macAddr := getIpMacAddr(line)
		if ipAddr == "" || macAddr == "" {
			continue
		}
		//デバイス情報を新規作成してリストに追加
		//TODO: MACアドレスを用いてDBにアクセスし、デバイス名を取得する
		devices = append(devices, Device{
			IpAddr:  ipAddr,
			MacAddr: macAddr,
			Name:    "デバイス" + strconv.Itoa(i),
		})
	}

	return devices
}

// 出力文の1行からIPアドレスとMACアドレスを抽出する関数
func getIpMacAddr(arpOutLines string) (ipAddr string, macAddr string) {
	elems := strings.Split(arpOutLines, " ")
	//要素がないとき
	if len(elems) < 4 {
		return "", ""
	}

	//IPアドレスとMACアドレスを抽出
	ipAddr = strings.Trim(elems[1], "()")
	macAddr = elems[3]
	return ipAddr, macAddr
}

// GetRegisteredDevices dbから取得した登録済みMACアドレスからデバイス情報を取得する関数
func GetRegisteredDevices() Devices {
	//登録済みMACアドレスを取得
	//TODO: ここではIPアドレス以外の情報を取得する予定
	macAddrList := db.RegisteredMacAddrList

	//現在のローカルDevicesを取得
	nowDevices := GetNowDevices()

	//登録済みdeviceリスト
	registeredDevices := make(Devices, 0, 50)

	for _, macAddr := range macAddrList {
		//一致するMACアドレスを持つDeviceが周囲に存在するか捜索する
		device := nowDevices.SearchByMacAddr(macAddr)
		if device.MacAddr != "" {
			registeredDevices = append(registeredDevices, device)
		}
	}

	return registeredDevices

}

// SendPingAll 指定した全IPアドレスにpingを送信する関数
func (devices Devices) SendPingAll() (connectingDeviceMap map[Device]bool) {
	//MACアドレスが接続しているか否かを格納するマップx
	connectingDeviceMap = make(map[Device]bool)

	//データ送受信用チャンネルを作成
	var wg = sync.WaitGroup{}
	wg.Add(len(devices)) //IPアドレスの数だけgoroutineを立てる

	//アクセスするごとにgoroutineを立てpingを送信
	for _, device := range devices {
		//goroutine実行
		go func() {
			defer wg.Done()
			connectingDeviceMap[device] = command.SendPing(device.IpAddr)
		}()
	}

	//全てのgoroutineが終了するまで待機
	wg.Wait()

	return connectingDeviceMap
}
