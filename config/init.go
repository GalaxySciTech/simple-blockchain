package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var Settings struct {
	Port string `json:"port"`
}

func Init() {
	viper.SetConfigName("config") //设置配置文件的名字
	viper.AddConfigPath(".")      //添加配置文件所在的路径
	viper.SetConfigType("json")   //设置配置文件类型，可选
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file error: %s\n", err)
		os.Exit(1)
	}
	Settings.Port = viper.GetString("port")
}
