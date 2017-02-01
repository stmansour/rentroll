package rlib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/kardianos/osext"
)

//==================================================================
// Read config information for the app. The information must
// be in a file named "config.json". It can be used for production
// information that is too sensitive to hardcode in binaries and
// store in source code.
//==================================================================
type rrconfig struct {
	Env      int    `json:"Env"`      // 0 = dev, 1 = prod, ...
	Dbuser   string `json:"Dbuser"`   // database user name
	Dbpass   string `json:"Dbpass"`   // database password
	Dbhost   string `json:"Dbhost"`   // tcp address of db host
	Dbport   int    `json:"Dbport"`   // tcp port on db host
	Dbtype   string `json:"Dbtype"`   // what type of database: mysql, ...
	RRDbuser string `json:"RRDbuser"` // database user name
	RRDbpass string `json:"RRDbpass"` // database password
	RRDbhost string `json:"RRDbhost"` // tcp address of db host
	RRDbport int    `json:"RRDbport"` // tcp port on db host
	RRDbtype string `json:"RRDbtype"` // what type of database: mysql, ...
}

// APPENVDEV et. al. are constants describing the environment where
// the app is running. It is set via the conf.json file
const (
	APPENVDEV  = 0
	APPENVPROD = 1
	APPENVQA   = 2
)

// AppConfig is the shared struct of configuration values
var AppConfig rrconfig

// RRReadConfig will read the configuration file "config.json" if
// it exists in the current directory
func RRReadConfig() {
	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("Executable folder = %s\n", folderPath)
	fname := folderPath + "/conf.json"
	_, err = os.Stat(fname)
	if nil != err {
		fmt.Printf("RRReadConfig: error reading %s: %v\n", fname, err)
		os.Exit(1)
	}
	content, err := ioutil.ReadFile(fname)
	Errcheck(err)
	Errcheck(json.Unmarshal(content, &AppConfig))
	// fmt.Printf("RRReadConfig: AppConfig = %#v\n", AppConfig)
}

// RRGetSQLOpenString builds the string to use for opening an sql database.
// Input string is the name of the database:  "accord" for phonebook, "rentroll" for RentRoll
// Returns:  a string to pass to sql.Open()
//=======================================================================================
func RRGetSQLOpenString(dbname string) string {
	s := ""
	switch strings.ToLower(dbname) {
	case "accord":
		switch AppConfig.Env {
		case 0: //dev
			s = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True",
				AppConfig.Dbuser, AppConfig.Dbpass, dbname)
		case 1: //production
			s = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
				AppConfig.Dbuser, AppConfig.Dbpass, AppConfig.Dbhost, AppConfig.Dbport, dbname)
		default:
			fmt.Printf("Unhandled configuration environment: %d\n", AppConfig.Env)
			os.Exit(1)
		}
	case "rentroll":
		switch AppConfig.Env {
		case 0: //dev
			s = fmt.Sprintf("%s:@/%s?charset=utf8&parseTime=True", AppConfig.RRDbuser, dbname)
		case 1: //production
			s = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
				AppConfig.RRDbuser, AppConfig.RRDbpass, AppConfig.RRDbhost, AppConfig.RRDbport, dbname)
		default:
			fmt.Printf("Unhandled configuration environment: %d\n", AppConfig.Env)
			os.Exit(1)
		}
	default:
		s = fmt.Sprintf("%s:@/%s?charset=utf8&parseTime=True", AppConfig.Dbuser, dbname)
	}
	return s
}
