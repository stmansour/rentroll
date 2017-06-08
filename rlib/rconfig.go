package rlib

import (
	"extres"
	"fmt"
	"log"
	"time"

	"github.com/kardianos/osext"
)

// AppConfig is the shared struct of configuration values
var AppConfig extres.ExternalResources

// RRReadConfig will read the configuration file "config.json" if
// it exists in the current directory
func RRReadConfig() error {
	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}
	fname := folderPath + "/config.json"
	err = extres.ReadConfig(fname, &AppConfig)
	if err != nil {
		log.Fatal(err)
	}
	RRdb.Zone, err = time.LoadLocation(AppConfig.Timezone)
	if err != nil {
		fmt.Printf("Error loading timezone %s : %s\n", AppConfig.Timezone, err.Error())
		Ulog("Error loading timezone %s : %s", AppConfig.Timezone, err.Error())
	}
	return err
}
