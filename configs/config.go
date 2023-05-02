package configs

import (
	"depeche/authorization_ms/config"
	"depeche/pkg/connector"
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
	"os"
	"os/exec"
	"path/filepath"
)

type Config struct {
	Host           string                   `yaml:"host"`
	Port           int                      `yaml:"port"`
	SessionStorage connector.RedisConfig    `yaml:"session"`
	DB             connector.PostgresConfig `yaml:"db"`
	DBMSName       string                   `yaml:"dbms_name"`
	AuthMs         config.AuthMsConfig      `yaml:"auth_ms"`
}

func InitCfg(config *Config) error {

	cmd := exec.Command("ls", "-a")
	stdout, _ := cmd.Output()
	fmt.Println(string(stdout))

	err := godotenv.Load(".env/backend", ".env/postgres", ".env/redis")
	if err != nil {
		return err
	}

	filename, err := filepath.Abs("configs/config.yml")

	if err != nil {
		return err
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		return err
	}
	rPass := os.Getenv("REDIS_PASS")
	config.SessionStorage.Pass = rPass

	pgUser := os.Getenv("POSTGRES_USER")
	pgPass := os.Getenv("POSTGRES_PASSWORD")
	config.DB.User = pgUser
	config.DB.Password = pgPass
	config.DBMSName = "postgres"
	return nil
}
