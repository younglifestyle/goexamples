package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

const cmdRoot = "core"

func main() {
	//viper.SetEnvPrefix(cmdRoot)
	//viper.AutomaticEnv()
	//replacer := strings.NewReplacer(".", "_")
	//viper.SetEnvKeyReplacer(replacer)
	//
	//viper.SetConfigName(cmdRoot)
	//viper.AddConfigPath(`.`)
	//
	//err := viper.ReadInConfig()
	//if err != nil {
	//	fmt.Println(fmt.Errorf("Fatal error when reading %s config file:%s", cmdRoot, err))
	//	os.Exit(1)
	//}
	//
	//environment := viper.GetBool("security.enabled")
	//fmt.Println("security.enabled:", environment)
	//
	//environment1 := viper.GetBool("security.enableds")
	//fmt.Println("security.enableds:", environment1)
	//
	//fullstate := viper.GetString("statetransfer.timeout.fullstate")
	//fmt.Println("statetransfer.timeout.fullstate:", fullstate)
	//
	//abcdValue := viper.GetString("peer.abcd")
	//fmt.Println("peer.abcd:", abcdValue)

	watchRemoteTest()

	for i := 0; i < 1000; i++ {
		go func() {
			for {
				value, ok := syncMap.Load("config")
				if ok {
					bytes, _ := json.Marshal(value.(*ConfigS))
					fmt.Println("required data :", string(bytes),
						value.(*ConfigS).LogLevel,
						value.(*ConfigS).UserCfg.Doamin,
						value.(*ConfigS).UserCfg.Port)
				}

				time.Sleep(time.Millisecond * 20)
			}
		}()
	}

	select {}
}

type HttpConfigs struct {
	Doamin string `mapstructure:"domain"`
	Port   string `mapstructure:"port"`
}
type ConfigS struct {
	LogLevel string      `mapstructure:"log_level"`
	UserCfg  HttpConfigs `mapstructure:"user"`
	MailCfg  string      `mapstructure:"mail"`
}

var syncMap sync.Map

func watchRemoteTest() {
	// alternatively, you can create a new viper instance.
	var runtime_viper = viper.New()
	var runtimeConf ConfigS

	runtime_viper.AddRemoteProvider("etcd",
		"http://mps.longsys.com:2379",
		"/config/Elrond/172.16.9.229/common.json")
	runtime_viper.SetConfigType("json") // because there is no file extension in a stream of bytes, supported extensions are "json", "toml", "yaml", "yml", "properties", "props", "prop"

	// read from remote config the first time.
	err := runtime_viper.ReadRemoteConfig()
	if err != nil {
		log.Fatal("error read from viper!", err)
	}

	// unmarshal config
	runtime_viper.Unmarshal(&runtimeConf)

	syncMap.Store("config", &runtimeConf)

	// open a goroutine to watch remote changes forever
	go func() {
		for {
			time.Sleep(time.Millisecond * 10) // delay after each request

			// currently, only tested with etcd support
			err := runtime_viper.WatchRemoteConfig()
			if err != nil {
				log.Println("unable to read remote config: ", err)
				continue
			}

			runtime_viper.Unmarshal(&runtimeConf)
			syncMap.Store("config", &runtimeConf)
		}
	}()
}

/*
//core.yaml
statetransfer:
    recoverdamage: true
    blocksperrequest: 20
    maxdeltas: 200
    timeout:
        singleblock: 2s
        singlestatedelta: 2s
        fullstate: 60s
peer:
    abcd:   3322d
*/
