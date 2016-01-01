package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type AppConf struct {
	Port    int
	LogPath string
}

var str = `
port = 1234
logPath = "/home/webroot/xiaoju/"
`

func main() {
	var conf AppConf
	if _, err := toml.Decode(str, &conf); err != nil {

	} else {
		fmt.Println(conf.LogPath)
	}
}
