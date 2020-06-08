package main

import (
	"os"
	"fmt"
	"strings"
	"encoding/json"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/layers"
)

// type Result struct {
// 	Msg    string  `json:"msg"`
// 	Success  bool   `json:"success"`
// }

func validate(f string)(bool, error) {
	_, err := pcap.NewBPF(layers.LinkTypeEthernet, 1500, f)
	if err != nil {
		return false, err
	}
	return true, err  
}

func main(){
	if len(os.Args) >= 2 {
		f := strings.Join(os.Args[1:], " ")
		valid, err := validate(f)
		msg := ""
		if err != nil {
			msg = err.Error()
		}
		res := map[string]interface{}{"msg":msg,"success":valid}
		data, err := json.Marshal(res)
        if err != nil {
                fmt.Println("An error occured: %v", err)
                os.Exit(1)
        }
		switch valid {
		case true:
			fmt.Println(string(data))
			os.Exit(0)
		case false:
			fmt.Println(string(data))
			os.Exit(1)
		}
	} else {
		fmt.Println("Incorrect number of arguments. Try again.")
		os.Exit(1)
	}
}