package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Telmate/proxmox-api-go/proxmox"
)

func main() {
	fmt.Println("wow proxctl!")

	prox_api_url := os.Getenv("PROXMOX_API_URL")
	prox_user := os.Getenv("PROXMOX_USERNAME")
	prox_pw := os.Getenv("PROXMOX_PASSWORD")

	if prox_api_url == "" || prox_user == "" || prox_pw == "" {
		log.Fatal("PROXMOX_API_URL, PROXMOX_USERNAME, PROXMOX_PASSWORD environment variables are not set")
	}

	// 証明書検証無効にする
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: tr}

	client, err := proxmox.NewClient(prox_api_url, httpClient, "", nil, "", 60)
	if err != nil {
		log.Fatal("Failed to create Proxmox client: ", err)
	}

	err = client.Login(prox_user, prox_pw, "")
	if err != nil {
		log.Fatal("Failed to log in to Proxmox: ", err)
	}

	// show vm list
	vmList, err := client.GetVmList()
	if err != nil {
		log.Fatal("Failed to get VM list: ", err)
	}

	fmt.Println("VM List:")
	for key, value := range vmList {
		fmt.Printf("Key: %s, Value: %+v\n", key, value)
	}

}
