package main

import (
	"context"
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"goexamples/my-log-agent/conf"
	"goexamples/my-log-agent/httpstream"
	"os"
)

const AppVersion = "0.0.1"

type Agent struct {
	httpstream *httpstream.HttpStream
}

var cfg = conf.Config{
	HttpAddr: httpstream.Config{
		Addr: ":8010",
	},
}

func main() {
	var (
		err    error
		ctx, _ = context.WithCancel(context.Background())
		//cancel context.CancelFunc
	)

	version := flag.Bool("v", false, "show version and exit")
	flag.Parse()
	if *version {
		fmt.Println(AppVersion)
		os.Exit(0)
	}

	agent := new(Agent)

	// set context
	ctx = context.WithValue(ctx, "GlobalConfig", cfg)

	//err = conf.ParseConfig("cfg.json")
	//if err != nil {
	//	log.Error("parse config error :", err)
	//	return
	//}

	// httpstream
	if agent.httpstream, err = httpstream.NewHttpStream(&cfg.HttpAddr); err != nil {
		log.Warn("httpstream disabled: %s", err)
	}

	select {}
}
