package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
}

// NewConfig 实例化一个 Config 并设置相关信息
func NewConfig() (cfg *Config, err error) {
	cfg = &Config{Viper: viper.New()} // 实例化

	// TODO: 设置默认值
	// cfg.SetDefault("Server", &ServerConfigS{
	// 	RunMode:      "debug",
	// 	Port:         8080,
	// 	ReadTimeOut:  5,
	// 	WriteTimeOut: 5,
	// })

	// cfg.SetDefault("Database", &DBConfigS{
	// 	DBtype:   "rethinkdb",
	// 	Username: "",
	// 	Password: "",
	// 	Host:     "119.91.204.226:32769",
	// 	DBName:   "chat",
	// 	Timeout:  5,
	// })
	// 设置默认配置信息
	// cfg.SetDefault("Server.RunMode", "debug")
	// cfg.SetDefault("Server.port", 8080)
	// cfg.SetDefault("Server.ReadTimeOut", 60)
	// cfg.SetDefault("Server.WriteTimeOut", 60)
	// cfg.SetDefault("Database.DBtype", "rethinkdb")
	// cfg.SetDefault("Database.Username", "root")
	// cfg.SetDefault("Database.Password", "123")
	// cfg.SetDefault("Database.Host", "119.91.204.226:32769")
	// cfg.SetDefault("Database.DBName", "chat")
	// cfg.SetDefault("Database.Timeout", "5s")

	// 从配置文件中读取配置信息
	cfg.SetConfigName("config")   // 设置配置名
	cfg.SetConfigType("yaml")     // 设置配置类型
	cfg.AddConfigPath("./config") // 添加配置路径(当前工作目录)
	err = cfg.ReadInConfig()      // 读入配置
	if err != nil {               // 找到并读取配置文件
		return nil, err
	}
	// 配置文件发生改变 重新读取
	cfg.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
	})
	cfg.WatchConfig()
	return cfg, nil
}

// ReadSection
func (cfg *Config) ReadSection(k string, v any) (err error) {
	err = cfg.UnmarshalKey(k, v) // 接收一个单件并转换为结构体
	if err != nil {
		return err
	}
	return nil
}
