package model

import "fmt"

type Device struct {
	IpAddr  string
	MacAddr string
	Name    string
}

// Print デバイス情報を出力する関数
func (d Device) Print() {
	fmt.Println("IP:", d.IpAddr, "MAC:", d.MacAddr, "Name:", d.Name)
}
