package uconfig

import (
	"github.com/micro/go-config"
	"github.com/sirupsen/logrus"
)

type (
	LogConfig struct {
		Path     string
		Level    string
		FileSize int
	}

)

var (
	DefaultLogConf LogConfig
)

func Init() {
	// 加载mongo配置
	err := config.Get("mongodb").Scan(&DefaultMgoConf)
	if err != nil {
		logrus.Fatalf("get mgo config error: %s", err)
	}

	if len(DefaultMgoConf.Addr) == 0 {
		logrus.Fatalf("invalid mgo addr")
	}

}