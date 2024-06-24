package utils

import (
	log "github.com/sirupsen/logrus"
)

func LogInfo(fields interface{}) {
	log.Info(fields)
}

func LogFatal(fields interface{}) {
	log.Fatal(fields)
}

func LogWarn(fields interface{}) {
	log.Warn(fields)
}
