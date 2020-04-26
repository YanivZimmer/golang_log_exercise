package main

import (
	"runtime"

	"github.com/natefinch/lumberjack.v2"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	if runtime.GOOS == "windows" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   "log_yaniv.log",
			MaxSize:    20, // 20MB,
			MaxAge:     7,  // 7 days tops
			MaxBackups: 3,  // no more than 3 files
		})
	}
	log.Debug("Hi i am the first line")
	log.Info("HI this is a line at the end of the program")
}
