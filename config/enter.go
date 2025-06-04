package config

type Config struct {
	System `yaml:"system"`
	DB     `yaml:"db"`
	Redis  `yaml:"redis"`
	Jwt    `yaml:"jwt"`
	Router `yaml:"router"`
	Email  `yaml:"email"`
	Info   `yaml:"info"`
}
