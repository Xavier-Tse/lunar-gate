package flags

import (
	"flag"
	"os"
)

type option struct {
	DB   bool
	File string
}

var FlagOptions option

func init() {
	flag.StringVar(&FlagOptions.File, "f", "settings.yaml", "配置文件")
	flag.BoolVar(&FlagOptions.DB, "db", false, "数据库迁移")
	flag.Parse()
}

func Run() {
	if FlagOptions.DB {
		AutoMigrate()
		os.Exit(0)
	}
}
