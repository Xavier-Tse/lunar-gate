package config

type Config struct {
	System  `yaml:"system"`
	DB      `yaml:"db"`
	Redis   `yaml:"redis"`
	Jwt     `yaml:"jwt"`
	Captcha `yaml:"captcha"`
	Email   `yaml:"email"`
}
