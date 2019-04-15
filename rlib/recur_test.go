package rlib

import (
	"fmt"
	"testing"
	"time"
)

const RPDATEFMT0 = "Mon, Jan _2, 2006 03:04 pm"

type recurtest struct {
	cycle    int64
	eyear    int        // epoch,this is the date(year) of the assessments occurring every cycle
	emonth   time.Month // epoch,this is the date(month) of the assessments occurring every cycle
	eday     int        // epoch,this is the date(day) of the assessments occurring every cycle
	ehour    int        // epoch,this is the date(hour) of the assessments occurring every cycle
	eminute  int
	tyear    int        // target,this is the date(year) of amending the original RA
	tmonth   time.Month // target,this is the date(month) of amending the original RA
	tday     int        // target,this is the date(day) of amending the original RA
	thour    int        // target,this is the date(hour) of amending the original RA
	tminute  int
	expyear  int        // expected result
	expmonth time.Month // expected result
	expday   int        // expected result
	exhour   int        //
	exminute int
}

//{
// 	{RECURWEEKLY, 2018, time.January, 1, 2018, time.August, 1, 2018, time.July, 30},

//		{RECURMONTHLY, 2018, time.January, 1, 2018, time.August, 1, 2018, time.August, 1},
//		{RECURMONTHLY, 2018, time.January, 1, 2018, time.August, 28, 2018, time.August, 1},
//		{RECURMONTHLY, 2018, time.February, 28, 2020, time.March, 8, 2020, time.February, 29},
//		{RECURMONTHLY, 2018, time.February, 28, 2020, time.February, 8, 2020, time.January, 31},

//		{RECURQUARTERLY, 2018, time.January, 1, 2018, time.August, 1, 2018, time.June, 1},
//		{RECURQUARTERLY, 2018, time.February, 14, 2018, time.August, 1, 2018, time.July, 14},
//	}

func ShowResult(cycle int64, epoch, dt, instance *time.Time) {
	fmt.Printf("Cycle = %d, epoch = %s, dt = %s, instance = %s\n", cycle, epoch.Format(RPDATEFMT0), dt.Format(RPDATEFMT0), instance.Format(RPDATEFMT0))
}

//-----------------------------invoke InstanceDateCoveringDate----------------------
func runDateCoveringTest(t *testing.T, name string, recurcase []recurtest) {

	verbose := false // should only be true when debugging
	passed := 0
	for i := 0; i < len(recurcase); i++ {

		epoch := time.Date(recurcase[i].eyear, recurcase[i].emonth, recurcase[i].eday, recurcase[i].ehour, recurcase[i].eminute, 0, 0, time.UTC)
		target := time.Date(recurcase[i].tyear, recurcase[i].tmonth, recurcase[i].tday, recurcase[i].thour, recurcase[i].tminute, 0, 0, time.UTC)
		expected := time.Date(recurcase[i].expyear, recurcase[i].expmonth, recurcase[i].expday, recurcase[i].exhour, recurcase[i].exminute, 0, 0, time.UTC)
		result := InstanceDateCoveringDate(&epoch, &target, recurcase[i].cycle)
		if verbose {
			ShowResult(recurcase[i].cycle, &epoch, &target, &result)
		}

		if !result.Equal(expected) {
			fmt.Printf("%s: Error on data row %d\n", name, i)
			fmt.Printf(result.String())
			fmt.Printf("\n")
		} else {
			passed++
		}
	}

	if verbose {
		fmt.Printf("%s : \n", name)
		fmt.Printf("Total tests: %d\n", len(recurcase))
		fmt.Printf("     Passed: %d\n", passed)
		fmt.Printf("     Failed: %d\n", len(recurcase)-passed)
	}
}

//---------------------Occur Monthly-----------------------------
func TestDateCoveringMonthly(t *testing.T) {

	var recurcase = []recurtest{
		{RECURMONTHLY, 2018, time.January, 1, 8, 0, 2018, time.August, 1, 8, 0, 2018, time.August, 1, 8, 0},         // 1/1
		{RECURMONTHLY, 2018, time.January, 1, 8, 0, 2018, time.August, 28, 8, 0, 2018, time.August, 1, 8, 0},        //1/1
		{RECURMONTHLY, 2018, time.February, 28, 8, 0, 2020, time.March, 8, 8, 0, 2020, time.February, 29, 8, 0},     //2/28
		{RECURMONTHLY, 2018, time.February, 28, 8, 0, 2020, time.February, 8, 8, 0, 2020, time.January, 31, 8, 0},   //2/28
		{RECURMONTHLY, 2018, time.January, 13, 8, 0, 2018, time.December, 31, 8, 0, 2018, time.December, 13, 8, 0},  //1/13
		{RECURMONTHLY, 2018, time.June, 29, 8, 0, 2020, time.February, 29, 8, 0, 2020, time.February, 29, 8, 0},     //6/29
		{RECURMONTHLY, 2018, time.February, 27, 8, 0, 2018, time.December, 31, 8, 0, 2018, time.December, 27, 8, 0}, //2/27
		{RECURMONTHLY, 2020, time.February, 29, 8, 0, 2020, time.December, 31, 8, 0, 2020, time.December, 31, 8, 0}, //2/29
	}
	runDateCoveringTest(t, "Test DateCoveringDate Monthly", recurcase[0:])
}

//----------------Occur Weekly-----------------------------------------
func TestDateCoveringWeekly(t *testing.T) {

	var recurcase = []recurtest{
		{RECURWEEKLY, 2018, time.January, 1, 8, 0, 2018, time.August, 1, 8, 0, 2018, time.July, 30, 8, 0}, // 1
		{RECURWEEKLY, 2018, time.January, 1, 8, 0, 2018, time.October, 28, 8, 0, 2018, time.October, 22, 8, 0},
		{RECURWEEKLY, 2018, time.February, 28, 8, 0, 2018, time.March, 2, 8, 0, 2018, time.February, 28, 8, 0},
		{RECURWEEKLY, 2018, time.December, 28, 8, 0, 2019, time.January, 3, 8, 0, 2018, time.December, 28, 8, 0},
	}
	runDateCoveringTest(t, "Test DateCoveringDate Weekly", recurcase[0:])
}

//----------------Occur Quarterly--------------------------------
func TestDateCoveringQuarterly(t *testing.T) {

	var recurcase = []recurtest{
		{RECURQUARTERLY, 2018, time.January, 1, 8, 0, 2018, time.August, 1, 8, 0, 2018, time.July, 1, 8, 0},  // 1/1
		{RECURQUARTERLY, 2018, time.January, 1, 8, 0, 2018, time.August, 28, 8, 0, 2018, time.July, 1, 8, 0}, //1/1
		{RECURQUARTERLY, 2018, time.February, 28, 8, 0, 2020, time.August, 1, 8, 0, 2020, time.May, 31, 8, 0},
		{RECURQUARTERLY, 2018, time.February, 28, 8, 0, 2020, time.November, 8, 8, 0, 2020, time.August, 31, 8, 0},
		{RECURQUARTERLY, 2018, time.January, 27, 8, 0, 2020, time.November, 8, 8, 0, 2020, time.October, 27, 8, 0},
		{RECURQUARTERLY, 2018, time.March, 27, 8, 0, 2020, time.November, 8, 8, 0, 2020, time.September, 27, 8, 0},
		{RECURQUARTERLY, 2018, time.December, 15, 8, 0, 2020, time.January, 8, 8, 0, 2019, time.December, 15, 8, 0},
	}
	runDateCoveringTest(t, "Test DateCoveringDate Quarterly", recurcase[0:])
}

//----------------Occur Yearly--------------------------------
func TestDateCoveringYearly(t *testing.T) {

	var recurcase = []recurtest{
		{RECURYEARLY, 2018, time.January, 1, 8, 0, 2019, time.August, 1, 8, 0, 2019, time.January, 1, 8, 0}, // 1
		{RECURYEARLY, 2018, time.January, 1, 8, 0, 2018, time.August, 28, 8, 0, 2018, time.January, 1, 8, 0},
		{RECURYEARLY, 2020, time.February, 29, 8, 0, 2022, time.March, 8, 8, 0, 2022, time.February, 28, 8, 0}, /// special day of 2/29/2020
		{RECURYEARLY, 2018, time.December, 26, 8, 0, 2020, time.February, 8, 8, 0, 2019, time.December, 26, 8, 0},
		{RECURYEARLY, 2018, time.July, 30, 8, 0, 2021, time.March, 8, 8, 0, 2020, time.July, 31, 8, 0},
	}
	runDateCoveringTest(t, "Test DateCoveringDate Yearly", recurcase[0:])
}

//----------------Occur Daily--------------------------------
func TestDateCoveringDaily(t *testing.T) {

	var recurcase = []recurtest{
		{RECURDAILY, 2018, time.January, 28, 10, 0, 2018, time.February, 2, 8, 0, 2018, time.February, 1, 10, 0},   //
		{RECURDAILY, 2018, time.December, 27, 22, 0, 2019, time.January, 28, 8, 0, 2019, time.January, 27, 22, 0},  //
		{RECURDAILY, 2018, time.September, 28, 10, 0, 2018, time.February, 2, 8, 0, 2018, time.February, 1, 10, 0}, //
		{RECURDAILY, 2018, time.May, 22, 22, 0, 2018, time.June, 8, 8, 0, 2018, time.June, 7, 22, 0},               //
	}
	runDateCoveringTest(t, "Test DateCoveringDate Daily", recurcase[0:])
}

//----------------Occur Hourly--------------------------------
func TestDateCoveringHourly(t *testing.T) {

	var recurcase = []recurtest{
		{RECURHOURLY, 2018, time.February, 3, 8, 1, 2018, time.February, 3, 20, 5, 2018, time.February, 3, 20, 1},
		{RECURHOURLY, 2018, time.July, 31, 13, 0, 2018, time.August, 8, 23, 31, 2018, time.August, 8, 23, 0},
		{RECURHOURLY, 2018, time.January, 1, 8, 30, 2018, time.January, 8, 20, 55, 2018, time.January, 8, 20, 30},
		{RECURHOURLY, 2018, time.July, 31, 13, 28, 2018, time.August, 8, 20, 45, 2018, time.August, 8, 20, 28},
		{RECURHOURLY, 2018, time.November, 30, 5, 0, 2018, time.December, 2, 20, 20, 2018, time.December, 2, 20, 0},
		{RECURHOURLY, 2018, time.November, 30, 5, 30, 2018, time.December, 2, 20, 30, 2018, time.December, 2, 20, 30},
	}
	runDateCoveringTest(t, "Test DateCoveringDate Hourly", recurcase[0:])
}
