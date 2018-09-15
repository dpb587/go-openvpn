package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dpb587/go-openvpn/ovpn"
	"github.com/dpb587/go-openvpn/ovpn/onc"
)

func main() {
	ovpnBytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	profile, err := ovpn.Parse(ovpnBytes)
	if err != nil {
		panic(err)
	}

	config, err := onc.Encode(profile)
	if err != nil {
		panic(err)
	}

	bytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", bytes)
}
