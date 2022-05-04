package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/yjlwuhen/yjlwuhen/modules"
	"os"
	"strings"
)

func main() {
	// animals struct
	var config modules.Config
	var animalflag bool
	var animal_type string
	//config init
	viper.New()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Errorf("Config file not found", err)
		} else {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}
	errumsh := viper.Unmarshal(&config)
	if errumsh != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
	}
	// 监控配置文件变化
	viper.WatchConfig()
	// 注意！！！配置文件发生变化后要同步到全局变量Conf
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件被人修改啦...")
		if err := viper.Unmarshal(&config); err != nil {
			panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
		}
	})
	fmt.Println("加载配置ok")
	modules.Engine, err = modules.Connection(config.Database.User, config.Database.Pass, config.Database.Host, config.Database.Name, config.Database.Char, config.Database.Port)
	if err != nil {
		fmt.Errorf("mysql connect fail", err)
		return
	}

	//migrate data
	//if !modules.Engine.Migrator().HasTable(&modules.Data{}) {
	//	modules.Engine.Migrator().CreateTable(&modules.Data{})
	//	fmt.Printf("success create data")
	//}
	//读取环境变量
	shout := os.Getenv("SHOUT")
	if len(os.Args[1:]) != 2 {
		fmt.Println("Example animals cow/bird/snake eat/move/speck")
		return
	}
	animal := os.Args[1:][0]
	action := os.Args[1:][1]
	animalflag = false
	for i := 0; i < len(config.Animals); i++ {
		if config.Animals[i].Name == animal {
			animalflag = true
			animal_type = config.Animals[i].Type

		}
	}
	if !animalflag {
		fmt.Printf("Animal %s doesn't exist! \n", string(animal))
		return
	}
	var data *modules.Data
	errmysql := modules.Engine.Where("animal_type=?", animal_type).First(&data).Error
	if err != nil {
		fmt.Errorf("data select fail", errmysql.Error())
	}

	switch action {
	case "eat":
		eat := data.Eat
		if shout == "True" {
			fmt.Println(strings.ToUpper(eat))
		} else {
			fmt.Println(eat)
		}
	case "move":
		move := data.Move
		if shout == "True" {
			fmt.Println(strings.ToUpper(move))
		} else {
			fmt.Println(move)
		}
	case "speak":
		speak := data.Speak
		if shout == "True" {
			fmt.Println(strings.ToUpper(speak))
		} else {
			fmt.Println(speak)
		}
	}
}
