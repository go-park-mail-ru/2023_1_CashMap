package config

type StaticMsConfig struct {
	Host      string `yaml:"host"`
	Port      uint   `yaml:"port"`
	ColorPort uint   `yaml:"color_port"`
}
