package configs

import (
	"depeche/pkg/connector"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type Config struct {
	Host           string                `yaml:"host"`
	Port           int                   `yaml:"port"`
	SessionStorage connector.RedisConfig `yaml:"session"`
}

func InitCfg(config *Config) error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	rPass := os.Getenv("REDIS_PASS")

	filename, err := filepath.Abs("./configs/config.yml")

	if err != nil {
		return err
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	//fmt.Println(config)
	err = yaml.Unmarshal(file, config)
	if err != nil {
		return err
	}

	config.SessionStorage.Pass = rPass
	return nil
}
