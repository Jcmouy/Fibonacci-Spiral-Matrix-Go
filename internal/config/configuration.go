package config

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	// log only debug mode
	if gin.Mode() != gin.DebugMode {
		log.SetOutput(ioutil.Discard)
	}
}
