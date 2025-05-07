package flags

import (
	"flag"
	"os"
)

type option struct {
	DB   bool
	File string
	Menu string
	Type string
}

var FlagOptions option

func init() {
	flag.StringVar(&FlagOptions.File, "f", "settings.yaml", "配置文件")
	flag.BoolVar(&FlagOptions.DB, "db", false, "数据库迁移")
	flag.StringVar(&FlagOptions.Menu, "m", "", "菜单")
	flag.StringVar(&FlagOptions.Type, "t", "", "类型")
	flag.Parse()
}

func Run() {
	if FlagOptions.DB {
		AutoMigrate()
		os.Exit(0)
	}
	switch FlagOptions.Menu {
	case "user":
		var u User
		switch FlagOptions.Type {
		case "create":
			u.Create()
			os.Exit(0)
		}
	}
}
