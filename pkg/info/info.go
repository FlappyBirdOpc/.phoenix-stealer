package user

import (
	"fmt"
	"os/user"
	"strings"
)

func GetUsername() (string, error) {
	Uname, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("failed to get current user: %w", err)
	}
	parts := strings.Split(Uname.Username, "\\")
	username := parts[len(parts)-1]
	return username, nil
}

func AnotherFunction() {
	username, err := GetUsername()
	if err != nil {
		fmt.Println("[-] Error:", err)
		return
	}
	fmt.Println("[+] Username in AnotherFunction:", username)
}
