package core

import (
	"flag"
	"fmt"
	"gin-react-admin/global"
	"gin-react-admin/utils"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

//viper 是一个配置解决方案，拥有丰富的特性

/*
	-支持 JSON/TOML/YAML/HCL/envfile/Java properties 等多种格式的配置文件
	-可以设置监听配置文件的修改，修改时自动加载新的配置
	-从环境变量、命令行选项和io.Reader中读取配置
	-从远程配置系统中读取和监听修改，如 etcd/Consul
	-代码逻辑中显示设置键值
*/

func Viper(path ...string) *viper.Viper {
	var config string
	fmt.Println(len(path))

	if len(path) == 0 { //当没有传入path的时候
		// 从命令行中获取 -c 的值
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		fmt.Println(config)
		if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(utils.ConfigEnv); configEnv == "" {
				config = utils.ConfigFile
				fmt.Printf("您正在使用config的默认值,config的路径为%v\n", utils.ConfigFile)
			} else {
				config = configEnv
				fmt.Printf("您正在使用GVA_CONFIG环境变量,config的路径为%v\n", config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}

	v := viper.New()
	// 设置配置文件路径
	v.SetConfigFile(config)
	// 设置配置文件类型
	v.SetConfigType("yaml")
	// 读取配置文件并捕获异常
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 监听配置文件修改
	v.WatchConfig()

	// 配置文件修改的钩子函数
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		// 将配置文件赋值给global.GRA_CONFIG
		if err := v.Unmarshal(&global.GRA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	// 将配置文件赋值给global.GRA_CONFIG
	if err := v.Unmarshal(&global.GRA_CONFIG); err != nil {
		fmt.Println(err)
	}

	fmt.Println(global.GRA_CONFIG.Mysql)

	return v

}
