package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dpb587/go-openvpn/ovpn"
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

	dump := map[string]interface{}{}

	for _, e := range profile.Elements {
		switch pe := e.(type) {
		case ovpn.DirectiveProfileElement:
			d := pe.Directive()

			switch d {
			case "remote":
				if _, exists := dump[d]; !exists {
					dump[d] = []interface{}{}
				}

				dump[d] = append(dump[d].([]interface{}), pe.Args())
			default:
				dump[d] = pe.Args()
			}
		case ovpn.EmbeddedProfileElement:
			dump[pe.Embed()] = pe.Data()
		}
	}

	dumpBytes, err := json.MarshalIndent(dump, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", dumpBytes)
}
