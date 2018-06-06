package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

// IG is the struct containing info for doing Identity Generation
var IG struct {
	FirstNames []string   // array of first names
	LastNames  []string   // array of last names
	Streets    []string   // array of street names
	Cities     []string   // array of cities
	States     []string   // array of states
	Companies  []string   // array of random company names
	Rand       *rand.Rand // random number generator to use
}

func initOpen(fname string, pm *[]string) {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Error opening file: %s - %s\n", fname, err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		*pm = append(*pm, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		log.Fatalf("Error with scanner: %s\n", err.Error())
	}
}

// IGInit initializes the Identity Generation code
//-----------------------------------------------------------------------------
func IGInit(r *rand.Rand) {
	var n = []struct {
		FName string
		PM    *[]string
	}{
		{"./idgen/firstnames.txt", &IG.FirstNames},
		{"./idgen/lastnames.txt", &IG.LastNames},
		{"./idgen/states.txt", &IG.States},
		{"./idgen/cities.txt", &IG.Cities},
		{"./idgen/streets.txt", &IG.Streets},
		{"./idgen/companies.txt", &IG.Companies},
	}

	for i := 0; i < len(n); i++ {
		initOpen(n[i].FName, n[i].PM)
	}
	IG.Rand = r
	// rlib.Console("FirstNames: %d\n", len(IG.FirstNames))
	// rlib.Console("LastNames: %d\n", len(IG.LastNames))
	// rlib.Console("Cities: %d\n", len(IG.Cities))
	// rlib.Console("States: %d\n", len(IG.States))
	// rlib.Console("Streets: %d\n", len(IG.Streets))
	// rlib.Console("Companies: %d\n", len(IG.Companies))
}

// GenerateRandomPhoneNumber returns a string with a random phone number
//-----------------------------------------------------------------------------
func GenerateRandomPhoneNumber() string {
	return fmt.Sprintf("(%d) %3d-%04d", 100+IG.Rand.Intn(899), 100+IG.Rand.Intn(899), IG.Rand.Intn(9999))
}

// GenerateRandomFirstName returns a string with a random first name
//-----------------------------------------------------------------------------
func GenerateRandomFirstName() string {
	return IG.FirstNames[IG.Rand.Intn(len(IG.FirstNames))]
}

// GenerateRandomLastName returns a string with a random last name
//-----------------------------------------------------------------------------
func GenerateRandomLastName() string {
	return IG.LastNames[IG.Rand.Intn(len(IG.LastNames))]
}

// GenerateRandomCity returns a string with a random city
//-----------------------------------------------------------------------------
func GenerateRandomCity() string {
	return IG.Cities[IG.Rand.Intn(len(IG.Cities))]
}

// GenerateRandomState returns a string with a random State
//-----------------------------------------------------------------------------
func GenerateRandomState() string {
	return IG.States[IG.Rand.Intn(len(IG.States))]
}

// GenerateRandomCompany returns a string with a random Company
//-----------------------------------------------------------------------------
func GenerateRandomCompany() string {
	return IG.Companies[IG.Rand.Intn(len(IG.Companies))]
}

// GenerateRandomEmail returns a string with a random email address
//-----------------------------------------------------------------------------
func GenerateRandomEmail(lastname string, firstname string) string {
	var providers = []string{"gmail.com", "yahoo.com", "comcast.net", "aol.com", "bdiddy.com", "hotmail.com", "abiz.com"}
	np := len(providers)
	n := IG.Rand.Intn(10)
	switch {
	case n < 4:
		return fmt.Sprintf("%s%s%d@%s", firstname[0:1], lastname, IG.Rand.Intn(10000), providers[IG.Rand.Intn(np)])
	case n > 6:
		return fmt.Sprintf("%s%s%d@%s", firstname, lastname[0:1], IG.Rand.Intn(10000), providers[IG.Rand.Intn(np)])
	default:
		return fmt.Sprintf("%s%s%d@%s", firstname, lastname, IG.Rand.Intn(1000), providers[IG.Rand.Intn(np)])
	}
}

// GenerateRandomAddress returns a string with a random address
//-----------------------------------------------------------------------------
func GenerateRandomAddress() string {
	return fmt.Sprintf("%d %s", IG.Rand.Intn(99999), IG.Streets[IG.Rand.Intn(len(IG.Streets))])
}
