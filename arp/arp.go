package main

import (
	"fmt"
	"os/exec"
	"strings"
)

type DeviceInfo struct {
	IpAddr  string
	MacAddr string
}

// デバイス情報を格納するスライスを作成
var deviceInfos = make([]DeviceInfo, 0, 30)

// go run arp.go
func main() {
	fmt.Println("--周辺のデバイス情報--")
	//arpリクエストをローカル全体に送りつける
	arpCmd := exec.Command("arp", "-a")
	arpOutLines := getCmdOutLines(arpCmd)

	for _, line := range arpOutLines { //i, v
		elems := strings.Split(line, " ")
		//要素がないとき
		if len(elems) < 4 {
			continue
		}

		//IPアドレスとMACアドレスを抽出
		ipAddr := strings.Trim(elems[1], "()")
		macAddr := elems[3]
		fmt.Println("IP:", ipAddr, "", "MAC:", macAddr)
		deviceInfos = append(deviceInfos, DeviceInfo{
			IpAddr:  ipAddr,
			MacAddr: macAddr,
		})
	}

	fmt.Println("\n--自身のデバイス情報--")

	//自身のIPアドレスとMACアドレスを取得
	confCmd := exec.Command("ifconfig")
	confOutLines := getCmdOutLines(confCmd)

	var isEn0Range bool
	for _, line := range confOutLines {
		elems := strings.Split(line, " ")

		if elems[0] == "en0:" { //"en0:" 区間開始
			isEn0Range = true

		} else if isEn0Range && line[0] != 9 { //"en0:" 区間終了
			isEn0Range = false
		}

		//"en0:"区間の情報抽出開始
		if !isEn0Range {
			continue
		}

		//MACアドレス(ether行)
		if elems[0] == "\tether" {
			fmt.Println("MAC:", elems[1])
		}
		//IPアドレス(inet行)
		if elems[0] == "\tinet" {
			fmt.Println("IP:", elems[1])
		}
		//fmt.Println(elems)
		//if isEn0Range &&
		//lineの最初の文字がスペース" "のとき
	}
}

// コマンドの出力を取得する関数
func getCmdOutLines(cmd *exec.Cmd) []string {
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}

	//結果を出力
	cmdOut := string(out)
	return strings.Split(cmdOut, "\n")
}
