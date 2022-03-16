package main

import (
	"fmt"
	"os"

	"github.com/erick-otenyo/go-webservice-starter/internal/conf"
	"github.com/erick-otenyo/go-webservice-starter/internal/service"

	"github.com/pborman/getopt/v2"
	log "github.com/sirupsen/logrus"
)

var flagDebugOn bool
var flagVersion bool
var flagConfigFilename string

func init() {
	initCommandOptions()
}

func initCommandOptions() {
	getopt.FlagLong(&flagConfigFilename, "config", 'c', "", "config file name")
	getopt.FlagLong(&flagDebugOn, "debug", 'd', "Set logging level to TRACE")
	getopt.FlagLong(&flagVersion, "version", 'v', "Output the version information")
}

func main() {
	getopt.Parse()

	if flagVersion {
		fmt.Printf("%s %s\n", conf.AppConfig.Name, conf.AppConfig.Version)
		os.Exit(1)
	}

	log.Infof("----  %s - Version %s ----------\n", conf.AppConfig.Name, conf.AppConfig.Version)

	conf.InitConfig(flagConfigFilename)

	if flagDebugOn {
		log.SetLevel(log.TraceLevel)
		log.Debugf("Log level = DEBUG\n")
	}

	// start the web service
	service.Serve()
}
