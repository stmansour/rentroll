package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"rentroll/rlib"
	"strconv"
)

// Alphabet contains caps of the alphabet
var Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Digits contains characters for 0 - 9
var Digits = "0123456789"

// CarInfo contains year, make, and model
type CarInfo struct {
	Year  int
	Make  string
	Model string
}

// IG is the struct containing info for doing Identity Generation
var IG struct {
	FirstNames []string   // array of first names
	LastNames  []string   // array of last names
	Streets    []string   // array of street names
	Cities     []string   // array of cities
	States     []string   // array of states
	Companies  []string   // array of random company names
	CarColors  []string   // array of colors
	Dogs       []string   // array of dog breeds
	Cats       []string   // array of cat breeds
	DogNames   []string   // array of dog names
	DogColors  []string   // array of dog colors
	CatNames   []string   // array of cat names
	CatColors  []string   // array of cat colors
	Cars       []CarInfo  // array of info about cars
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

func loadCars(fname string, c *[]CarInfo) {
	funcname := "loadCars"
	csvFile, _ := os.Open(fname)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			rlib.LogAndPrintError(funcname, err)
		}
		// rlib.Console("line: [0] = %s, [1] = %s, [2] = %s\n", line[0], line[1], line[2])
		var car CarInfo
		car.Year, err = strconv.Atoi(line[0])
		if err != nil {
			rlib.Console("line[0] = %q\n", line[0])
			rlib.LogAndPrintError(funcname, err)
		}
		car.Make = line[1]
		car.Model = line[2]
		*c = append(*c, car)
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
		{"./idgen/carcolors.txt", &IG.CarColors},
		{"./idgen/cats.txt", &IG.Cats},
		{"./idgen/catnames.txt", &IG.CatNames},
		{"./idgen/dogs.txt", &IG.Dogs},
		{"./idgen/dognames.txt", &IG.DogNames},
		{"./idgen/dogcolors.txt", &IG.DogColors},
		{"./idgen/catcolors.txt", &IG.CatColors},
	}

	loadCars("./idgen/cars.csv", &IG.Cars)
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
	// rlib.Console("CarInfo: %d\n", len(IG.Cars))
	// rlib.Console("CarColors: %d\n", len(IG.CarColors))
	// rlib.Console("Cats: %d\n", len(IG.Cats))
	// rlib.Console("Dogs: %d\n", len(IG.Dogs))
	// rlib.Console("DogNames: %d\n", len(IG.DogNames))
	// rlib.Console("CatNames: %d\n", len(IG.CatNames))
	// rlib.Console("DogColors: %d\n", len(IG.DogColors))
	// rlib.Console("CatColors: %d\n", len(IG.CatColors))
}

// GenerateRandomLicensePlate returns a string with a random license plate
// number according to california rules -- 7-digit plates, 3 letters,
// 4 numbers
//-----------------------------------------------------------------------------
func GenerateRandomLicensePlate() string {
	var l []byte
	for i := 0; i < 3; i++ {
		l = append(l, Alphabet[IG.Rand.Intn(26)])
	}
	for i := 3; i < 7; i++ {
		l = append(l, Digits[IG.Rand.Intn(10)])
	}
	for i := 0; i < 5; i++ {
		j := IG.Rand.Intn(7)
		k := IG.Rand.Intn(7)
		l[k], l[j] = l[j], l[k]
	}
	return string(l)
}

// GenerateRandomPhoneNumber returns a string with a random phone number
//-----------------------------------------------------------------------------
func GenerateRandomPhoneNumber() string {
	return fmt.Sprintf("(%d) %3d-%04d", 100+IG.Rand.Intn(899), 100+IG.Rand.Intn(899), IG.Rand.Intn(9999))
}

// GenerateRandomDog returns a string with a random dog breed
//-----------------------------------------------------------------------------
func GenerateRandomDog() string {
	return IG.Dogs[IG.Rand.Intn(len(IG.Dogs))]
}

// GenerateRandomCat returns a string with a random dog breed
//-----------------------------------------------------------------------------
func GenerateRandomCat() string {
	return IG.Cats[IG.Rand.Intn(len(IG.Cats))]
}

// GenerateRandomDogName returns a string with a random dog breed
//-----------------------------------------------------------------------------
func GenerateRandomDogName() string {
	return IG.DogNames[IG.Rand.Intn(len(IG.DogNames))]
}

// GenerateRandomCatName returns a string with a random dog breed
//-----------------------------------------------------------------------------
func GenerateRandomCatName() string {
	return IG.CatNames[IG.Rand.Intn(len(IG.CatNames))]
}

// GenerateRandomDogColor returns a string with a random dog breed
//-----------------------------------------------------------------------------
func GenerateRandomDogColor() string {
	return IG.DogColors[IG.Rand.Intn(len(IG.DogColors))]
}

// GenerateRandomCatColor returns a string with a random dog breed
//-----------------------------------------------------------------------------
func GenerateRandomCatColor() string {
	return IG.CatColors[IG.Rand.Intn(len(IG.CatColors))]
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
