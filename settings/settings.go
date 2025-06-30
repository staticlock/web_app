package settings

import (
	"fmt"

	"github.com/spf13/pflag"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 所有的结构体都要首字母大写，不然读取配置读不到
type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
}
type MysqlConfig struct {
	Host        string `mapstructure:"host"`
	Port        string `mapstructure:"port"`
	User        string `mapstructure:"user"`
	PassWord    string `mapstructure:"password"`
	DbName      string `mapstructure:"dbname"`
	MaxIdleConn int    `mapstructure:"MaxIdleConn"`
	MaxOpenConn int    `mapstructure:"MaxOpenConn"`
}
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	PassWord string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}
type AppConfig struct {
	Name string `mapstructure:"name"`
	Port string `mapstructure:"port"`
}
type config struct {
	Name        string `mapstructure:"name"`
	Mode        string `mapstructure:"mode"`
	Port        string `mapstructure:"port"`
	Version     string `mapstructure:"version"`
	LogConfig   `mapstructure:"log"`
	MysqlConfig `mapstructure:"mysql"`
	RedisConfig `mapstructure:"redis"`
}

// 全局配置变量
var Config = new(config)

// 指出返回值为err，里面return会默认返回err
func Init() error {
	// 1. 使用pflag（兼容flag标准库）
	// 默认值设置为空
	pflag.StringP("config", "c", "", "config file path")
	pflag.IntP("port", "p", 8080, "server port")
	pflag.Parse()
	// 2. 绑定pflag到viper
	viper.BindPFlags(pflag.CommandLine)
	// 3. 加载配置文件
	if configFile := viper.GetString("config"); configFile == "" {
		//直接指定配置文件路径和名称类型
		viper.SetConfigFile("./settings/config.yaml")
	} else {
		viper.SetConfigFile(configFile)
	}
	//viper.SetConfigName("config")     //指定配置文件名称
	//viper.AddConfigPath("./settings") //指定搜索路径
	//指定配置文件类型,这个函数配合远程配置中心使用,告诉viper用什么格式解析
	//viper.SetConfigType("yaml")
	//读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("读取配置文件错误:%v\n", err)
		return err
	}
	if err := viper.Unmarshal(Config); err != nil {
		fmt.Printf("配置无法解码为结构体:%v\n", err)
		return err
	}
	//配置文件热加载
	viper.OnConfigChange(func(e fsnotify.Event) {
		//fmt.Println("配置文件发生变化...")
		fmt.Println(Config)
		if err := viper.Unmarshal(Config); err != nil {
			fmt.Printf("配置无法解码为结构体:%v\n", err)
		}
		fmt.Println(Config)
	})
	viper.WatchConfig()
	return nil
}
