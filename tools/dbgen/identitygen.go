package main

import (
	"bufio"
	"context"
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

// WMIInfo describes the world car manufacturers information
// in a VIN
type WMIInfo struct {
	Code         string
	Manufacturer string
}

// IG is the struct containing info for doing Identity Generation
var IG struct {
	FirstNames  []string        // array of first names
	LastNames   []string        // array of last names
	Streets     []string        // array of street names
	Cities      []string        // array of cities
	States      []string        // array of states
	Companies   []string        // array of random company names
	CarColors   []string        // array of colors
	Dogs        []string        // array of dog breeds
	Cats        []string        // array of cat breeds
	DogNames    []string        // array of dog names
	DogColors   []string        // array of dog colors
	CatNames    []string        // array of cat names
	CatColors   []string        // array of cat colors
	Occupations []string        // career occupations
	Industries  []string        // industry area of focus
	Cars        []CarInfo       // array of info about cars
	Mfgs        []WMIInfo       // array of auto manufacturers worldwide
	Rand        *rand.Rand      // random number generator to use
	WhyLeaving  rlib.StringList // strings for why leaving last residence
	HowFound    rlib.StringList // strings for how the applicant found out about the property
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

// loadWMI loads the list of WMI codes and mfg
//------------------------------------------------
func loadWMI(fname string, c *[]WMIInfo) {
	funcname := "loadWMI"
	csvFile, _ := os.Open(fname)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			rlib.LogAndPrintError(funcname, err)
		}
		var mfg WMIInfo
		mfg.Code = line[0] // this is the only thing we need.
		mfg.Manufacturer = line[1]
		*c = append(*c, mfg)
	}
}

// IGInit initializes the Identity Generation code
//-----------------------------------------------------------------------------
func IGInit(r *rand.Rand) {
	var err error
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
		{"./idgen/occupation.txt", &IG.Occupations},
		{"./idgen/industries.txt", &IG.Industries},
	}

	loadCars("./idgen/cars.csv", &IG.Cars)
	loadWMI("./idgen/wmi.csv", &IG.Mfgs)
	for i := 0; i < len(n); i++ {
		initOpen(n[i].FName, n[i].PM)
	}

	IG.Rand = r

	ctx := context.Background()
	if err = rlib.GetStringListByName(ctx, 1, "WhyLeaving", &IG.WhyLeaving); err != nil {
		rlib.Console("Error getting StringList: WhyLeaving: %s\n", err.Error())
		os.Exit(1)
	}
	if err = rlib.GetStringListByName(ctx, 1, "HowFound", &IG.HowFound); err != nil {
		rlib.Console("Error getting StringList: HowFound: %s\n", err.Error())
		os.Exit(1)
	}

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
	// rlib.Console("WMIs: %d\n", len(IG.Mfgs))
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

// GenerateRandomVIN returns a string with a random Vehicle Identification
// Number.  A VIN is a 17-digit alpha-numeric string constructed as follows:
//
//                 11111111
// Digit: 12345678901234567
// VIN:   1G6AF5SX6D0125409
//                 |      |<-- 10 : 17 = Vehicle Identifier Section
//            |   |<----------  4 :  9 = Vehicle Descriptor Section
//        |  |<---------------  1 :  3 = World Manufacturer Identifier
//
// World Manufacturer Identifier  (1:3)
//     1. Position one represents the nation of origin, or the final point of
//        assembly. For instance, cars made in the U.S. start with 1,4 or 5,
//        Canada is 2, Mexico is 3, Japan is J, South Korea is K, England is S,
//        Germany is W, and Sweden or Finland is Y.  So, it can be:
//     2. Manufacturer. In some cases, it's the letter that begins the
//        manufacturer's name. For example, A is for Audi, B is for BMW, G is
//        for General Motors, L is for Lincoln and N is for Nissan. But that
//        "A" can also stand for Jaguar or Mitsubishi and an "R" can also mean
//        Audi. It may sound confusing, but the next digit ties it all
//        together.
//      3. Position three, when combined with the first two digits, indicates
//         the vehicle's type or manufacturing division. In our example, 1G6
//         means a Cadillac passenger car. 1G1 means Chevrolet passenger cars
//         and 1GC means Chevrolet trucks. There have been many variations on
//         the World Manufacturer Identifier as brands have come and gone.
//         This Wikipedia page has a list of WMI codes:
//         https://en.wikibooks.org/wiki/Vehicle_Identification_Numbers_(VIN_codes)/World_Manufacturer_Identifier_(WMI)
//
// Vehicle Descriptor Section (4:9)
//      1. Positions 4-8 describe the car with such information as the model,
//         body type, restraint system, transmission type and engine code.
//
//      2. Position 9, the "check" digit, is used to detect invalid VINs,
//         based on a mathematical formula that was developed by the
//         Department of Transportation.  This code will just choose a random
//         digit or letter.
//
// Vehicle Identifier Section (10:17)
//      Varies by Manufacturer.  As an example, here's the info for Cadillac:
//
//         Position 10 indicates the model year. The letters from B-Y
//         correspond to the model years 1981-2000. There is no I, O, Q, U or
//         Z. From 2001-'09, the numbers one through nine were used in place
//         of numbers. The alphabet started over from A in 2010 and will
//         continue until 2030. The letter or number in position 11 indicates
//         the manufacturing plant in which the vehicle was assembled. Each
//         automaker has its own set of plant codes. The last 6 digits
//         (positions 12 through 17) are the production sequence numbers. This
//         is the number each car receives on the assembly line. In the case
//         of our Cadillac ATS, it was the 125,409th car to roll off the
//         assembly line in Lansing, Michigan.
//-----------------------------------------------------------------------------
func GenerateRandomVIN() string {
	wmi := IG.Mfgs[IG.Rand.Intn(len(IG.Mfgs))].Code
	vds := randomAlphaNumeric(5)
	vis := randomAlphaNumeric(8)
	return wmi + vds + vis
}

func randomAlphaNumeric(n int) string {
	var b []byte
	for i := 0; i < n; i++ {
		j := IG.Rand.Intn(36)
		if j < 10 {
			b = append(b, Digits[j])
			continue
		}
		b = append(b, Alphabet[j-10])
	}
	return string(b)
}

// GenerateRandomOccupation returns a random career occupation
//-----------------------------------------------------------------------------
func GenerateRandomOccupation() string {
	return IG.Occupations[IG.Rand.Intn(len(IG.Occupations))]
}

// GenerateRandomIndustry returns a random career industry
//-----------------------------------------------------------------------------
func GenerateRandomIndustry() string {
	return IG.Industries[IG.Rand.Intn(len(IG.Industries))]
}

// GenerateRandomDurationString returns a random duration
//-----------------------------------------------------------------------------
func GenerateRandomDurationString() string {
	if IG.Rand.Intn(100) > 90 {
		return fmt.Sprintf("%d months", 2+IG.Rand.Intn(10))
	}
	return fmt.Sprintf("%d years %d months", 1+IG.Rand.Intn(10), 2+IG.Rand.Intn(10))
}

// GenerateRandomOneLineAddress returns a random full address in a single line
//-----------------------------------------------------------------------------
func GenerateRandomOneLineAddress() string {
	return GenerateRandomAddress() + ", " + GenerateRandomCity() + ", " + GenerateRandomState() + " " + fmt.Sprintf("%05d", rand.Intn(100000))
}

// GenerateRandomSSN returns a random social security number
//-----------------------------------------------------------------------------
func GenerateRandomSSN() string {
	return fmt.Sprintf("%03d-%02d-%04d", IG.Rand.Intn(1000), IG.Rand.Intn(100), IG.Rand.Intn(10000))
}

// GenerateRandomCarColor returns a random social security number
//-----------------------------------------------------------------------------
func GenerateRandomCarColor() string {
	return IG.CarColors[IG.Rand.Intn(len(IG.CarColors))]
}

// GenerateRandomDriversLicense returns a random drivers license number
//-----------------------------------------------------------------------------
func GenerateRandomDriversLicense() string {
	return fmt.Sprintf("%c%07d", Alphabet[IG.Rand.Intn(26)], IG.Rand.Intn(10000000))
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

// GenerateRandomName returns a string with a random first and last name
//-----------------------------------------------------------------------------
func GenerateRandomName() string {
	return GenerateRandomFirstName() + " " + GenerateRandomLastName()
}

// GenerateRandomFullName returns a string with a random first, middle, and
// last name
//-----------------------------------------------------------------------------
func GenerateRandomFullName() string {
	return GenerateRandomFirstName() + " " + GenerateRandomFirstName() + " " + GenerateRandomLastName()
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
