// 说明：声明配置属性结构体
package config

import "time"

// DBConfigS 数据库配置信息
type DBConfigS struct {
	DBtype   string        `mapstructure:"dbtype"`
	Username string        `mapstructure:"username"`
	Password string        `mapstructure:"password"`
	Host     string        `mapstructure:"host"`
	DBName   string        `mapstructure:"dbname"`
	Timeout  time.Duration `mapstructure:"timeout"`
}

// ServerConfigS 服务器配置信息
type ServerConfigS struct {
	RunMode      string        `mapstructure:"runmode"`
	Port         string        `mapstructure:"port"`
	ReadTimeOut  time.Duration `mapstructure:"read_timeout"`
	WriteTimeOut time.Duration `mapstructure:"write_timeout"`
}

// Storage 存储配置
type StorageS struct {
	SavePath       string
	ServerUrl      string
	ImageAllowExts []string
}

// Email 邮箱配置 - 用于发送邮件
type EmailS struct {
	Host     string
	Port     int
	Username string
	Password string
}
