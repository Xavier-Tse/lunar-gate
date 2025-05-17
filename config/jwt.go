package config

type Jwt struct {
	Expires int    `yaml:"expires" json:"expires"` // 单位：小时
	Issuer  string `yaml:"issuer" json:"issuer"`
	Secret  string `yaml:"secret" json:"secret"`
}
