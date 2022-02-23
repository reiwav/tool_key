package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/hyperboloide/lk"
	"github.com/reiwav/tool_key/config"

	"github.com/golang/glog"
)

func main() {
	privateKey, err := lk.NewPrivateKey()
	if err != nil {
		log.Fatal(err)

	}
	publicKey := privateKey.GetPublicKey()
	key1, _ := privateKey.ToB64String()
	fmt.Println("=== 1 === ", key1)
	fmt.Println("=== 2 === ", publicKey.ToB64String())
	var conf = config.Read()
	switch conf.Type {
	case "READ":
		fallthrough
	case "USE":
		var ls = config.Decode()
		ls.Println()
	case "NEW":
		var err = conf.GenerateLicense()
		if err != nil {
			fmt.Println("ERROR: ", err)
			glog.Errorf("ERROR: ", err)
		}
		fmt.Println("==== GEN OK =====")
		conf.Println()
	}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}
