package config

type Captcha struct {
	Enable bool   `yaml:"enable"`
	Type   string `yaml:"type"`
}
