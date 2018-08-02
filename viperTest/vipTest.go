package main

import (
	"strings"

	"fmt"
	"os"

	"github.com/spf13/viper"
)

const cmdRoot = "core"

func main() {
	viper.SetEnvPrefix(cmdRoot)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.SetConfigName(cmdRoot)
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("Fatal error when reading %s config file:%s", cmdRoot, err))
		os.Exit(1)
	}

	environment := viper.GetBool("security.enabled")
	fmt.Println("security.enabled:", environment)

	fullstate := viper.GetString("statetransfer.timeout.fullstate")
	fmt.Println("statetransfer.timeout.fullstate:", fullstate)

	abcdValue := viper.GetString("peer.abcd")
	fmt.Println("peer.abcd:", abcdValue)
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
