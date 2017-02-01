package main

// There are unit test for these conversions. The problem is if the unit tests fail
// it is hard to debug Go tests -- stdout is piped someplace but it doesn't go to
// the screen.  So this test program can be used to debug something in the normal
// fashion.

import (
	"fmt"
	"reflect"
	"rentroll/rlib"
)

// Data conversion tests

type foo1 struct {
	I1 int
}

type bar1 struct {
	I1 rlib.XJSONYesNo
}

type testDataConv1 struct {
	I1     int
	expect rlib.XJSONYesNo
}

var td1 = []testDataConv1{
	{0, "no"},
	{1, "yes"},
	{99, "no"},
}

// TestConversion1 tests data migration of XJSONYesNo
func TestConversion1() {
	for i := 0; i < len(td1); i++ {
		var a foo1
		var b bar1
		b.I1 = rlib.XJSONYesNo("no")
		a.I1 = td1[i].I1
		rlib.MigrateStructVals(&a, &b)
		if b.I1 != td1[i].expect {
			fmt.Printf("int2YestNo( %d )  expect %s, got %s\n", td1[i].I1, td1[i].expect, b.I1)
		}
	}
}

// TestConversion2 tests data migration using Str2Int64 maps
func TestConversion2() {
	var BUDlist = rlib.Str2Int64Map{
		"REX": 1,
		"ISO": 2,
		"CCC": 3,
	}
	var foo struct {
		BID rlib.XJSONBud
	}
	var bar struct {
		BID int64
	}
	foo.BID = "ISO"

	ar := reflect.ValueOf(&foo).Elem()
	fa := ar.Field(0)
	br := reflect.ValueOf(&bar).Elem()
	fb := br.Field(0)
	rlib.MigrateStrToInt64(&fa, &fb, &BUDlist)
	if bar.BID != 2 {
		fmt.Printf("Error: bar.BID != 2 after migration: bar = %#v, foo = %#v, BUDlist = %#v\n", bar, foo, BUDlist)
	}
	foo.BID = ""
	bar.BID = int64(1)
	ar = reflect.ValueOf(&foo).Elem()
	fa = ar.Field(0)
	br = reflect.ValueOf(&bar).Elem()
	fb = br.Field(0)
	rlib.MigrateInt64ToString(&fb, &fa, &BUDlist)
	if foo.BID != "REX" {
		fmt.Printf("Error: foo.BID != REX after migration: bar = %#v, foo = %#v, BUDlist = %#v\n", bar, foo, BUDlist)
	}
	fmt.Printf("After migration, foo.BID = %s\n", foo.BID)
	fmt.Printf("SUCCESS!\n")
}

func main() {
	TestConversion1()
	TestConversion2()
}
