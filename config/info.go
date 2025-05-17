package config

type Info struct {
	Site    `yaml:"site"`
	Project `yaml:"project"`
	Login   `yaml:"login"`
}

type Site struct {
	Title   string `yaml:"title"`
	EnTitle string `yaml:"en_title"`
	Slogan  string `yaml:"slogan"`
	Logo    string `yaml:"logo"`
	Icp     string `yaml:"icp"`
}

type Project struct {
	Title string `yaml:"title"`
	Icon  string `yaml:"icon"`
	Path  string `yaml:"path"`
}

type Login struct {
	Captcha `yaml:"captcha"`
	Email   `yaml:"email"`
}

type Captcha struct {
	Enable bool   `yaml:"enable"`
	Type   string `yaml:"type"`
}

type Email struct {
	Enable       bool   `yaml:"enable"`
	Domain       string `yaml:"domain"`
	Port         int    `yaml:"port"`
	SendEmail    string `yaml:"send_email"`
	AuthCode     string `yaml:"auth_code"`
	SendNickname string `yaml:"send_nickname"`
}
