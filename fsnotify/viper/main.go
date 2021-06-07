package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Name    string
	Type    string
	Content string
}

func main() {
	viper.SetConfigFile("./fsnotify/test.json")
	viper.SetConfigType("json")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	vip := viper.GetViper()
	conf := Config{}
	if err := vip.Unmarshal(&conf); err != nil {
		panic(err)
	}
	fmt.Println(conf)

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		c := Config{}
		if err := vip.Unmarshal(&c); err != nil {
			panic(err)
		}
		fmt.Println(c, conf)
		fmt.Println("Config file changed:", e.Name, vip.Get("content"))
	})

	select {}
}
