package config

type Email struct {
	Enable       bool   `yaml:"enable"`
	Domain       string `yaml:"domain"`
	Port         int    `yaml:"port"`
	SendEmail    string `yaml:"send_email"`
	AuthCode     string `yaml:"auth_code"`
	SendNickname string `yaml:"send_nickname"`
}
