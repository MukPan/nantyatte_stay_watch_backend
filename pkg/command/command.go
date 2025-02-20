package command

import (
	"fmt"
	"os/exec"
)

func ExecArp() string {
	arpCmd := exec.Command("arp", "-a")
	out, err := arpCmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}

	//結果を出力
	return string(out)
}

// SendPing pingの送信が成功したかを返す関数
func SendPing(ipAddr string) bool {
	pingCmd := exec.Command("sh", "-c",
		fmt.Sprintf("ping %s -o -c 1", ipAddr))
	out, err := pingCmd.CombinedOutput()

	//確認用
	fmt.Println(string(out))

	//結果をマップに代入
	return err == nil
}
