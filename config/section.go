// 说明：声明配置属性结构体
package config

import "time"

// 数据库配置信息
type DBConfigS struct {
	DBtype   string        `mapstructure:"dbtype"`
	Username string        `mapstructure:"username"`
	Password string        `mapstructure:"password"`
	Host     string        `mapstructure:"host"`
	DBName   string        `mapstructure:"dbname"`
	Timeout  time.Duration `mapstructure:"timeout"`
}

// 服务器配置信息
type ServerConfigS struct {
	RunMode      string        `mapstructure:"runmode"`
	Port         string        `mapstructure:"port"`
	ReadTimeOut  time.Duration `mapstructure:"read_timeout"`
	WriteTimeOut time.Duration `mapstructure:"write_timeout"`
}
