package db

// RegisteredMacAddrList 登録済MACアドレスリスト
var RegisteredMacAddrList = make([]string, 0, 30)

// RegistMacAddrList MACアドレスを新規登録する関数
func RegistMacAddrList(macAddr string) {
	RegisteredMacAddrList = append(RegisteredMacAddrList, macAddr)
}

// GetRegisteredMacAddrList 登録済MACアドレスリストを取得する関数
