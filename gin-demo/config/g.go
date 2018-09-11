package config

import (
	"log"
	"runtime"

	"github.com/spf13/viper"
)

// change log:
const (
	VERSION = "0.0.1"
)

func SetReadCfgEnv() error {
	viper.AddConfigPath("./cfgfile/")
	viper.SetConfigName("cfg")

	return viper.ReadInConfig()
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	if err := SetReadCfgEnv(); err != nil {
		log.Panic(`please put the config file in the "./cfgfile/"`)
	}
}
