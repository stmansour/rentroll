package main

import (
	"fmt"
	"strings"
	"time"
)

// RECURNONE - RECURLAST are the allowed recurrence types
const (
	RECURNONE      = 0
	RECURHOURLY    = 1
	RECURDAILY     = 2
	RECURWEEKLY    = 3
	RECURMONTHLY   = 4
	RECURQUARTERLY = 5
	RECURYEARLY    = 6
	RECURLAST      = RECURYEARLY
)

func recurStringToInt(s string) int {
	var i int
	s = strings.ToUpper(s)
	s = strings.Replace(s, " ", "", -1)
	switch {
	case s == "NONE":
		i = RECURNONE
	case s == "HOURLY":
		i = RECURHOURLY
	case s == "DAILY":
		i = RECURDAILY
	case s == "WEEKLY":
		i = RECURWEEKLY
	case s == "MONTHLY":
		i = RECURMONTHLY
	case s == "QUARTERLY":
		i = RECURQUARTERLY
	case s == "YEARLY":
		i = RECURYEARLY
	default:
		fmt.Printf("Unknown recurrence type: %s\n", s)
		i = RECURNONE
	}
	return i
}

func recurIntToString(i int) string {
	var s string
	switch {
	case i == RECURNONE:
		s = "None"
	case i == RECURHOURLY:
		s = "Hourly"
	case i == RECURDAILY:
		s = "Daily"
	case i == RECURWEEKLY:
		s = "Weekly"
	case i == RECURMONTHLY:
		s = "Monthly"
	case i == RECURQUARTERLY:
		s = "Quarterly"
	case i == RECURYEARLY:
		s = "Yearly"
	default:
		fmt.Printf("Unknown acceptance value: %d\n", i)
		s = "None"
	}
	return s
}

func monthToInt(m time.Month) int {
	switch m {
	case time.January:
		return 1
	case time.February:
		return 2
	case time.March:
		return 3
	case time.April:
		return 4
	case time.May:
		return 5
	case time.June:
		return 6
	case time.July:
		return 7
	case time.August:
		return 8
	case time.September:
		return 9
	case time.October:
		return 10
	case time.November:
		return 11
	case time.December:
		return 12
	}
	return 0 // should never happen
}

func incMonths(m time.Month, n int) (time.Month, int) {
	y := 0
	mo := monthToInt(m) + n - 1
	y += mo / 12
	mo = mo % 12
	switch mo {
	case 0:
		m = time.January
	case 1:
		m = time.February
	case 2:
		m = time.March
	case 3:
		m = time.April
	case 4:
		m = time.May
	case 5:
		m = time.June
	case 6:
		m = time.July
	case 7:
		m = time.August
	case 8:
		m = time.September
	case 9:
		m = time.October
	case 10:
		m = time.November
	case 11:
		m = time.December
	}
	return m, y
}
