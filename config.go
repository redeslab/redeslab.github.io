package main

import (
	"encoding/json"
	"fmt"
)

type NodeName struct {
	IPAddr string `json:"ip_addr"`
	Name   string `json:"name"`
}

func main() {
	nn := make(map[string]*NodeName)
	nn["aa"] = &NodeName{
		IPAddr: "192.168.1.2",
		Name:   "Tokyo",
	}
	nn["bb"] = &NodeName{
		IPAddr: "192.168.1.1",
		Name:   "souel",
	}
	bt, _ := json.MarshalIndent(nn, "", "\t")
	fmt.Println(string(bt))
}
