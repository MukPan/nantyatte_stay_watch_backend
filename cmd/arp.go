package cmd

import (
	"fmt"
	"os/exec"
	"strings"
)

type DeviceInfo struct {
	IpAddr  string
	MacAddr string
}

// GetDeviceInfos 2つのデバイス情報を取得する関数
func GetDeviceInfos() []DeviceInfo {
	//arpリクエストをローカル全体に送りつける
	arpCmd := exec.Command("arp", "-a")
	arpOut := execCmd(arpCmd)
	arpOutLines := strings.Split(arpOut, "\n")

	// デバイス情報を格納するスライスを作成
	deviceInfos := make([]DeviceInfo, 0, 30)

	//すべての行に対して処理
	for _, line := range arpOutLines { //i, v
		//IPアドレスとMACアドレスを抽出
		ipAddr, macAddr := getIpMacAddr(line)
		if ipAddr == "" || macAddr == "" {
			continue
		}
		//リストに追加
		deviceInfos = append(deviceInfos, DeviceInfo{
			IpAddr:  ipAddr,
			MacAddr: macAddr,
		})
	}

	return deviceInfos
}

// PrintDeviceInfos デバイス情報を出力する関数
func PrintDeviceInfos(deviceInfos []DeviceInfo) {
	fmt.Println("--周辺のデバイス情報--")
	for _, deviceInfo := range deviceInfos {
		fmt.Println("IP:", deviceInfo.IpAddr, "MAC:", deviceInfo.MacAddr)
	}
}

// SearchMacAddr IPアドレスを検索しMACアドレスを取得する関数
func SearchMacAddr(deviceInfos []DeviceInfo, ipAddr string) string {
	for _, deviceInfo := range deviceInfos {
		//一致するIPを見つけたらMACアドレスを返す
		if deviceInfo.IpAddr == ipAddr {
			return deviceInfo.MacAddr
		}
	}
	return ""
}

// コマンドの出力を取得する関数
func execCmd(cmd *exec.Cmd) string {
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}

	//結果を出力
	return string(out)
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
