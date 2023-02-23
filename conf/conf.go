package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"loginsystem/Log"
)

type Config struct {
	App *App    `yaml:"app"` //数据库信息
	Log *LogLog `yaml:"log"`
}

type App struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type LogLog struct {
	Suffix  string `yaml:"suffix"`
	MaxSize int    `yaml:"maxSize"`
}

var TanConfig *Config

func InitConfig() {
	yamlFile, err := ioutil.ReadFile("./conf/config.yml") //Go 1.17之后放到os里面了
	if err != nil {
		Log.ErrorLog.Printf(err.Error())
		return
	}
	err = yaml.Unmarshal(yamlFile, &TanConfig)
	if err != nil {
		Log.ErrorLog.Printf(err.Error())
		return
	}
}
