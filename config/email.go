package config

type Email struct {
	Enable       bool   `yaml:"enable" json:"enable"`
	Domain       string `yaml:"domain" json:"domain"`
	Port         int    `yaml:"port" json:"port"`
	SendEmail    string `yaml:"send_email" json:"sendEmail"`
	AuthCode     string `yaml:"auth_code" json:"authCode"`
	SendNickname string `yaml:"send_nickname" json:"sendNickname"`
}
