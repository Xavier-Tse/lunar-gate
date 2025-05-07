package config

type Jwt struct {
	Expires int    `yaml:"expires"` // 单位：小时
	Issuer  string `yaml:"issuer"`
	Secret  string `yaml:"secret"`
}
