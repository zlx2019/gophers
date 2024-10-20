package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var Envs = setup()

// 配置
type Config struct {
	Host  string
	Port  string
	Store StoreConfig
}

func (c Config) Addr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

// 数据库配置
type StoreConfig struct {
	Username string
	Password string
	Addr     string
	DBName   string
}

// 初始化配置
func setup() Config {
	// 将 .env 配置载入环境变量
	godotenv.Load()
	storeConfig := StoreConfig{
		Username: getEnvVar("DB_USER", "root"),
		Password: getEnvVar("DB_PASSWORD", "rootROOT@1234"),
		DBName:   getEnvVar("DB_NAME", "ecom"),
		Addr:     fmt.Sprintf("%s:%s", getEnvVar("DB_HOST", "127.0.0.1"), getEnvVar("DB_PORT", "3306")),
	}
	c := Config{
		Host:  getEnvVar("SERVER_HOST", "127.0.0.1"),
		Port:  getEnvVar("SERVER_PORT", "9898"),
		Store: storeConfig,
	}
	return c
}

// 获取环境变量
func getEnvVar(key, defVal string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defVal
}
