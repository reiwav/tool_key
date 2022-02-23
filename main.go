package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/reiwav/tool_key/config"

	"github.com/golang/glog"
)

func main() {
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
