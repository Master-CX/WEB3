package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
	Port      int    `mapstructure:"port"`

	*LogConfig   `mapstructure:"log"`
	*BscConfig   `mapstructure:"Bsc"`
	*EthConfig   `mapstructure:"Eth"`
	*AvaxConfig  `mapstructure:"Avax"`
	*MaticConfig `mapstructure:"Matic"`
}
type BscConfig struct {
	Name    string `mapstructure:"Name"`
	Url     string `mapstructure:"Url"`
	ChainID string `mapstructure:"ChainId"`
	Symbol  string `mapstructure:"Symbol"`
}
type EthConfig struct {
	Name    string `mapstructure:"Name"`
	Url     string `mapstructure:"Url"`
	ChainID string `mapstructure:"ChainId"`
	Symbol  string `mapstructure:"Symbol"`
}
type AvaxConfig struct {
	Name    string `mapstructure:"Name"`
	Url     string `mapstructure:"Url"`
	ChainID string `mapstructure:"ChainId"`
	Symbol  string `mapstructure:"Symbol"`
}
type MaticConfig struct {
	Name    string `mapstructure:"Name"`
	Url     string `mapstructure:"Url"`
	ChainID string `mapstructure:"ChainId"`
	Symbol  string `mapstructure:"Symbol"`
}
type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

func InitConfig() (err error) {
	// 加载配置
	viper.SetConfigFile("./config/config.yaml")
	// 读取配置信息
	viper.ReadInConfig()

	// 把读取到的配置信息反序列化到 Conf 变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	}
	//观察配置信息是否改变
	viper.WatchConfig()

	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		// 反序列化到Conf变量中
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})
	return err
}
