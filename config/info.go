package config

type Info struct {
	Site    `yaml:"site" json:"site"`
	Project `yaml:"project" json:"project"`
	Login   `yaml:"login" json:"login"`
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
	Captcha `yaml:"captcha" json:"captcha"`
}

type Captcha struct {
	Enable bool   `yaml:"enable" json:"enable"`
	Type   string `yaml:"type" json:"type"`
}
