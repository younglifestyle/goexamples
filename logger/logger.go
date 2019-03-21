package main

import log "github.com/Sirupsen/logrus"

func InitLog(level string) {
	switch level {
	case "info":
		log.SetLevel(log.InfoLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	default:
		log.Fatal("log conf only allow [info, debug, warn], please check your confguire")
	}
	return
}

func main() {
	InitLog("warn")

	log.Debug("debug")
	log.Info("info")
	log.Println("print")
	log.Error("error")
}
