package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env          string           `yaml:"env"`
	ServerConfig HTTPServerConfig `yaml:"http_server"`
	DBConfig     DBConfig
}
type HTTPServerConfig struct {
	Host        string        `yaml:"host" env:"HOST"`
	Port        string        `yaml:"port" env:"PORT"`
	Timeout     time.Duration `yaml:"timeout" env:"TIMEOUT"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"IDLE_TIMEOUT"`
}
type DBConfig struct {
	User    string `env:"FTODO_DB_USER"`
	Pass    string `env:"FTODO_DB_PASS"`
	Host    string `env:"FTODO_DB_HOST"`
	Port    string `env:"FTODO_DB_PORT"`
	DBName  string `env:"FTODO_DB_NAME"`
	SSLMode bool   `env:"FTODO_DB_SSLMODE"`
}

func MustLoad() *Config {
	path := os.Getenv("FTODO_CONFIG_PATH")
	if path == "" {
		log.Fatalf("Cannot find config path env")
		return nil
	}
	if _, err := os.Stat(path); err != nil {
		log.Fatal("Cannot find config file by path!", err)
		return nil
	}
	cfg := &Config{}
	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		log.Fatalf("Cannot read config! Error=%s", err.Error())
		return nil
	}
	return cfg
}
