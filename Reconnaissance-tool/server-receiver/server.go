package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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

func collectHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost { // если это не метод отправки
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var receivedInfo SystemInfo
	err := json.NewDecoder(r.Body).Decode(&receivedInfo) // &receivedInfo передаем с указателем
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("--- New Data Received ---")
	fmt.Printf("User: %s\n", receivedInfo.UserName)
	fmt.Printf("User Group: %s\n", receivedInfo.UserGroup)
	fmt.Printf("Host Name: %s\n", receivedInfo.HostName)
	fmt.Printf("Home Dir: %s\n", receivedInfo.HomeDir)
	fmt.Printf("OS: %s\n", receivedInfo.OS)
	fmt.Printf("Executable: %s\n", receivedInfo.Executable)
	fmt.Printf("Architecture: %s\n", receivedInfo.Architecture)

	for _, addr := range receivedInfo.IPs {
		fmt.Println("- IP Address:", addr)
	}

	fmt.Println("----------------------------------------------------")
}

func main() {
	http.HandleFunc("/collect", collectHandler) // всё что приходит на адрес /collect должно обрабатываться методом collectHandler()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server start error: ", err.Error())
	}
}
