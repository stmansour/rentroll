package rlib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

// AppConfig is the shared struct of configuration values
var AppConfig rrconfig

// RRReadConfig will read the configuration file "config.json" if
// it exists in the current directory
func RRReadConfig() {
	fname := "config.json"
	if _, err := os.Stat(fname); err == nil {
		content, err := ioutil.ReadFile(fname)
		Errcheck(err)
		Errcheck(json.Unmarshal(content, &AppConfig))
	}
}

// RRGetSQLOpenString builds the string to use for opening an sql database.
// If the configuration file is not present, it uses the supplied default information.
// Returns:  a string to pass to sql.Open()
//=======================================================================================
func RRGetSQLOpenString(defaultUser, defaultName string) string {
	s := ""
	switch AppConfig.Env {
	case 0: //dev
		s = fmt.Sprintf("%s:@/%s?charset=utf8&parseTime=True", defaultUser, defaultName)
	case 1: //production
		s = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
			AppConfig.RRDbuser, AppConfig.RRDbpass, AppConfig.RRDbhost, AppConfig.RRDbport, defaultName)
	default:
		fmt.Printf("Unhandled configuration environment: %d\n", AppConfig.Env)
		os.Exit(1)
	}
	return s
}
