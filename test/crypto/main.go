// The purpose of this test is to validate the db update routines.
package main

import (
	"context"
	"database/sql"
	"encoding/hex"
	"extres"
	"flag"
	"fmt"
	"log"
	"os"
	"rentroll/rlib"

	_ "github.com/go-sql-driver/mysql"
)

// App is the global application structure
var App struct {
	dbdir   *sql.DB        // phonebook db
	dbrr    *sql.DB        //rentroll db
	DBDir   string         // phonebook database
	DBRR    string         //rentroll database
	DBUser  string         // user for all databases
	PortRR  int            // rentroll port
	Bud     string         // Biz Unit Descriptor
	Xbiz    rlib.XBusiness // lots of info about this biz
	NoAuth  bool
	Verbose bool
}

func readCommandLineArgs() {
	dbuPtr := flag.String("B", "ec2-user", "database user name")
	dbnmPtr := flag.String("N", "accord", "directory database (accord)")
	dbrrPtr := flag.String("M", "rentroll", "database name (rentroll)")
	pBud := flag.String("b", "REX", "Business Unit Identifier (Bud)")
	portPtr := flag.Int("p", 8270, "port on which RentRoll server listens")
	noauth := flag.Bool("noauth", false, "if specified, inhibit authentication")
	verb := flag.Bool("v", false, "verbose output - shows the ciphertext")

	flag.Parse()

	App.DBDir = *dbnmPtr
	App.DBRR = *dbrrPtr
	App.DBUser = *dbuPtr
	App.PortRR = *portPtr
	App.Bud = *pBud
	App.NoAuth = *noauth
	App.Verbose = *verb
}

func main() {
	var err error
	readCommandLineArgs()
	rlib.RRReadConfig()

	//----------------------------
	// Open RentRoll database
	//----------------------------
	if err = rlib.RRReadConfig(); err != nil {
		fmt.Printf("sql.Open for database=%s, dbuser=%s: Error = %v\n", rlib.AppConfig.RRDbname, rlib.AppConfig.RRDbuser, err)
		os.Exit(1)
	}

	s := extres.GetSQLOpenString(rlib.AppConfig.RRDbname, &rlib.AppConfig)
	App.dbrr, err = sql.Open("mysql", s)
	if nil != err {
		fmt.Printf("sql.Open for database=%s, dbuser=%s: Error = %v\n", rlib.AppConfig.RRDbname, rlib.AppConfig.RRDbuser, err)
		os.Exit(1)
	}
	defer App.dbrr.Close()
	err = App.dbrr.Ping()
	if nil != err {
		fmt.Printf("DBRR.Ping for database=%s, dbuser=%s: Error = %v\n", rlib.AppConfig.RRDbname, rlib.AppConfig.RRDbuser, err)
		os.Exit(1)
	}

	//----------------------------
	// Open Phonebook database
	//----------------------------
	s = extres.GetSQLOpenString(rlib.AppConfig.Dbname, &rlib.AppConfig)
	App.dbdir, err = sql.Open("mysql", s)
	if nil != err {
		fmt.Printf("sql.Open: Error = %v\n", err)
		os.Exit(1)
	}
	err = App.dbdir.Ping()
	if nil != err {
		fmt.Printf("dbdir.Ping: Error = %v\n", err)
		os.Exit(1)
	}

	rlib.RpnInit()
	rlib.InitDBHelpers(App.dbrr, App.dbdir)
	rlib.SetAuthFlag(App.NoAuth)

	// create background context
	ctx := context.Background()

	DoTest(ctx)
}

// DoTest does the account balance checks for Rental Agreements
func DoTest(ctx context.Context) {
	rawssn := "012-34-5678"               // likely the way people will enter it
	ssn := rlib.Stripchars(rawssn, "- .") // remove everything except the digits
	encssn, err := rlib.Encrypt(ssn)      // simple encryption that will be used on the db
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	decssn, err := rlib.Decrypt(encssn) // decrypted as the db routines would do after pulling from db
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// We do not print the encssn, it is different every time, even with the same input string and key
	fmt.Printf("raw ssn:          %s\n", rawssn)
	if App.Verbose {
		fmt.Printf("ciphertext:       %x\n", encssn)
	}
	fmt.Printf("ssn:              %s\n", ssn)
	fmt.Printf("decrypted ssn:    %s\n", decssn)

	// Here are some pre-saved encrypted sstrings that should all decode to the same thing...
	ciphers := []string{
		"fae8d6d0fb6c7cc07e83320b9365f39d010df23cee8c72903c80e8ae9b6a71d7d93911eed6",
		"491657366c81500837275ba11b31ec0443da7dc0760448070633d04c17c7916ff76702c726",
		"00798a295c2f4b544fd3e8a244165edb8af5d2bc1593d7167b87e8b3a7bab748648da79210",
		"5b67481ca0edc40b07a8fb74e284710da988eabd6e8215fe00e53efddb6217562e07458055",
		"e513e0e59baaf52647a4f30bbd069a500948a8ce6bf2e968efbb661e342819ee6cb66e848b",
		"5e526cb389650e7a4edf0c0349130e33e5d4c4a94f7a22f7312f79cbbe8bd4f0c80cfa4e2b",
		"fae8d6d0fb6c7cc07e83320b9365f39d010df23cee8c72903c80e8ae9b6a71d7d93911eed6",
		"4fe875cd81267c0f2cd1754f295a95ea2b341c3c66e25f865f9297ac963a562b82cbefac81",
	}

	for i := 0; i < len(ciphers); i++ {
		fmt.Printf("presaved chars:   %s\n", ciphers[i])
		b, err := hex.DecodeString(ciphers[i])
		fmt.Printf("decoded hexbytes: %x\n", b)
		dssn, err := rlib.Decrypt(b)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		fmt.Printf("decrypted:        %s\n", dssn)

	}
}
