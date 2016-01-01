package main

import (
	"comm/base"
	"comm/strus"
	"fmt"

	"github.com/yaml"
)

func main() {
	fmt.Println(base.Env)

	appConf := strus.App{}
	bytes, _ := base.Buffers("app.yml")
	yaml.Unmarshal(bytes, &appConf)

	fmt.Println(appConf)
}
