package setting

import (
	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
)

type Server struct {
	RunMode string
	Port int
}

type Database struct {
	Port int
}

type Log struct {
	LogSavePath string
	LogPrefix string
	LogFileExtension string
	TimeFormat string
}

var ServerSetting = &Server{}
var DatabaseSetting = &Database{}
var LogSetting = &Log{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("../conf/app.ini")
	if err != nil {
		logrus.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	MapTo("server", ServerSetting)
	MapTo("database", DatabaseSetting)
	MapTo("log", LogSetting)
}

// Map to map section
func MapTo(section string, v interface{}) {
	var err error
	err = cfg.Section(section).MapTo(v)
	if err != nil {
		logrus.Fatalf("Cfg.MapTo %s err: %v", section, err)

	}
}