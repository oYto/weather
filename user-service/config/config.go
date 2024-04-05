package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var globalConfig = new(GlobalConfig)

type GlobalConfig struct {
	*SvrConfig    `mapstructure:"svr_config"`
	*DbConfig     `mapstructure:"mysql"`
	*LogConfig    `mapstructure:"log"`
	*ConsulConfig `mapstructure:"consul"`
}

type SvrConfig struct {
	Name        string `mapstructure:"name"` // 服务name
	Host        string `mapstructure:"host"` // 服务host
	Port        int    `mapstructure:"port"`
	UserSvrName string `mapstructure:"user_svr_name"` // 用户服务name
}

type DbConfig struct {
	Host        string `mapstructure:"host"`
	Port        string `mapstructure:"port"`
	Database    string `mapstructure:"database"`
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
	MaxIdleConn int    `mapstructure:"max_idle_conn"`  // 最大空闲连接数
	MaxOpenConn int    ` mapstructure:"max_open_conn"` // 最大打开的连接数
	MaxIdleTime int64  ` mapstructure:"max_idle_time"` // 连接最大空闲时间
}
type LogConfig struct {
	FileName   string `mapstructure:"file_name"`
	Level      string `mapstructure:"level"`
	LogPath    string `mapstructure:"log_path"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type ConsulConfig struct {
	Host string   `mapstructure:"host" json:"host" yaml:"host"`
	Port int      `mapstructure:"port" json:"port" yaml:"port"`
	Tags []string `mapstructure:"tags" json:"tags" yaml:"tags"`
}

// Init 读取配置文件
func Init() (err error) {
	//configFile := GetRootDir() + "/config/config.yaml"
	//viper.SetConfigFile(configFile)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml") // 远程配置文件传输 确定配置文件的格式
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config") // 指定配置文件的一个寻找路径
	viper.AddConfigPath("./../../config")
	viper.AddConfigPath("./../config")
	err = viper.ReadInConfig() // 读取配置信息
	//fmt.Println("------------------")

	if err != nil {
		// 读取配置信息错误
		fmt.Printf("viper.ReadInConfig() failed: %v\n", err)
		return fmt.Errorf("viper.ReadInConfig() failed: %v\n", err)
	}

	// 把读取到的信息反序列化到 Conf 变量中
	if err = viper.Unmarshal(globalConfig); err != nil {
		fmt.Printf("viper.Unmarshal failed: %v\n", err)
		return fmt.Errorf("viper.Unmarshal failed: %v\n", err)
	}

	viper.WatchConfig() // 实时监控配置文件（热加载）
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改...")
		// 当配置文件信息发生变化 就修改 Conf 变量
		if err := viper.Unmarshal(globalConfig); err != nil {
			fmt.Printf("viper.Unmarshal failed: %v\n", err)
		}
	})
	return nil
}

func GetGlobalConfig() *GlobalConfig {
	return globalConfig
}
