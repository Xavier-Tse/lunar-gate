package config

type Info struct {
	Site    `yaml:"site"`
	Project `yaml:"project"`
	Login   `yaml:"login"`
}

type Site struct {
	Title   string `yaml:"title" json:"title"`
	EnTitle string `yaml:"en_title" json:"enTitle"`
	Slogan  string `yaml:"slogan" json:"slogan"`
	Logo    string `yaml:"logo" json:"logo"`
	Icp     string `yaml:"icp" json:"icp"`
}

type Project struct {
	Title string `yaml:"title" json:"title"`
	Icon  string `yaml:"icon" json:"icon"`
	Path  string `yaml:"path" json:"path"`
}

type Login struct {
	Captcha `yaml:"captcha"`
	Email   `yaml:"email"`
}

type Captcha struct {
	Enable bool   `yaml:"enable" json:"enable"`
	Type   string `yaml:"type" json:"type"`
}

type Email struct {
	Enable       bool   `yaml:"enable" json:"enable"`
	Domain       string `yaml:"domain" json:"domain"`
	Port         int    `yaml:"port" json:"port"`
	SendEmail    string `yaml:"send_email" json:"sendEmail"`
	AuthCode     string `yaml:"auth_code" json:"authCode"`
	SendNickname string `yaml:"send_nickname" json:"sendNickname"`
}
