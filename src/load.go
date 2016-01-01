package comm

import (
	"comm/base"
	"comm/strus"
	"io/ioutil"

	"github.com/yaml"
)

var Env string = "dev"

/**
 * 装载配置文件
 * @param {[type]} fileName string [description]
 * @param {[type]} entity interface{} [description]
 */
func Buffers(fileName string) ([]byte, error) {
	path := "conf/" + Env + "/" + fileName
	path2 := "conf/" + fileName

	filePath := ""
	if base.FileExists(path) {
		filePath = path
	}
	if base.FileExists(path2) {
		filePath = path2
	}
	return ioutil.ReadFile(filePath)
}

// 初始化配置
func init() {
	bytes, err := ioutil.ReadFile("../conf/env.yml")
	base.CheckError(err)

	env := strus.Env{}
	err2 := yaml.Unmarshal(bytes, &env)
	base.CheckError(err2)

	Env = env.Env
}
