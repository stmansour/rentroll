package rlib

// This file manages the mapping between the programming data types
// and the way those types are displayed in the user interface.
//
// IF YOU ADD SOMETHING HERE, DO NOT FORGET TO UPDATE assignmap
import (
	"fmt"
	"reflect"
	"strings"
)

// W2uiHTMLSelect is a struct that covers the way w2ui sends back the
// selection from a dropdown list.
type W2uiHTMLSelect struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

var assignmap = []struct {
	a      string
	b      string
	mapper func(a, b *reflect.Value) error
}{
	{a: "XJSONAssignmentTime", b: "int64", mapper: AssignmentTime2Int64},
	{a: "int64", b: "XJSONAssignmentTime", mapper: Int642AssignmentTime},
	{a: "int", b: "JSONbool", mapper: Int2Bool},
	{a: "JSONbool", b: "int", mapper: Bool2Int},
	{a: "int64", b: "JSONbool", mapper: Int642Bool},
	{a: "JSONbool", b: "int64", mapper: Bool2Int64},
	{a: "XJSONYesNo", b: "int", mapper: yesNoStr2Int64},
	{a: "int", b: "XJSONYesNo", mapper: int642YesNo},
	{a: "XJSONBud", b: "int", mapper: bud2Int64},
	{a: "int", b: "XJSONBud", mapper: int642Bud},
}

// XJSONAssignmentTime is a UI converter: backend int64, Front End string
type XJSONAssignmentTime string

// AssignmentTimeMap is the mapping for Rentable - AssignmentTime
var AssignmentTimeMap = Str2Int64Map{
	"unset":        0,
	"Pre-Assign":   1,
	"Commencement": 2,
}

// XJSONBud is a UI converter: back-end int, UI: string
type XJSONBud string

// bud2Int64 converter
// a must be *XJSONBud
// b must be *int
func bud2Int64(a, b *reflect.Value) error {
	bud := (*a).Interface()
	s := fmt.Sprintf("%v", bud)
	bid, ok := RRdb.BUDlist[s]
	if !ok {
		bid = int64(0)
	}
	(*b).Set(reflect.ValueOf(bid))
	return nil
}

// int642Bud converter
// a must be *int
// b must be *XJSONBud
func int642Bud(a, b *reflect.Value) error {
	s, err := RRdb.BUDlist.ReverseMap((*a).Interface().(int64))
	if err != nil {
		return err
	}
	(*b).Set(reflect.ValueOf(XJSONBud(s)))
	return nil
}

// XJSONYesNo is a UI converter: back-end int, UI: string
type XJSONYesNo string

// YesNoMap is the mapping for no = 0, 1 = yes
var YesNoMap = map[string]int64{
	"no":  int64(0),
	"yes": int64(1),
}

// yesNoStr2Int64 converter
// a must be *XJSONYesNo
// b must be *int
func yesNoStr2Int64(a, b *reflect.Value) error {
	s1 := (*a).Interface()
	s := fmt.Sprintf("%v", s1)
	yn := int64(0)
	if strings.ToLower(s) == "yes" {
		yn = int64(1)
	}
	(*b).Set(reflect.ValueOf(yn))
	return nil
}

// int642YesNo converter
// a must be *int
// b must be *XJSONYesNo
func int642YesNo(a, b *reflect.Value) error {
	i := fmt.Sprintf("%v", (*a).Interface().(int64))
	s := "no"
	if "1" == i {
		s = "yes"
	}
	(*b).Set(reflect.ValueOf(XJSONYesNo(s)))
	return nil
}

// // AssignmentTimeSL is the inverse of AssignmentTimeMap; maps an int to a string
// var AssignmentTimeSL = []string{"unset", "Pre-Assign", "Commencement"}

// AssignmentTime2Int64 converter
// a must be *XJSONAssignmentTime
// b must be *int64
func AssignmentTime2Int64(a, b *reflect.Value) error {
	s1 := (*a).Interface()
	s := fmt.Sprintf("%v", s1)
	y := int64(0)
	var ok bool
	y, ok = AssignmentTimeMap[s]
	if !ok {
		return fmt.Errorf("AssignmentTime2Int64: index %q not found\n", s)
	}
	(*b).Set(reflect.ValueOf(y))
	return nil
}

// Int642AssignmentTime converter
// a must be *int64
// b must be *Int642AssignmentTime
func Int642AssignmentTime(a, b *reflect.Value) error {
	// var s XJSONAssignmentTime
	s, err := AssignmentTimeMap.ReverseMap((*a).Interface().(int64))
	if err != nil {
		return err
	}
	(*b).Set(reflect.ValueOf(XJSONAssignmentTime(s)))
	return nil
}

// Int2Bool copies an int into a bool value as follows
// if the int is 0, the bool value is false
// for any other value of the int the bool is true
// a must point to an int
// b must point to a bool
func Int2Bool(a, b *reflect.Value) error {
	(*b).Set(reflect.ValueOf(0 != (*a).Interface().(int)))
	return nil
}

// Bool2Int is the exact inverse of Int2Bool
// a must point to a bool
// b must point to an int
func Bool2Int(a, b *reflect.Value) error {
	i := 0
	if false != (*a).Interface().(bool) {
		i = 1
	}
	(*b).Set(reflect.ValueOf(i))
	return nil
}

// Int642Bool copies an int into a bool value as follows
// if the int is 0, the bool value is false
// for any other value of the int the bool is true
// a must point to an int
// b must point to a bool
func Int642Bool(a, b *reflect.Value) error {
	(*b).Set(reflect.ValueOf(0 != (*a).Interface().(int64)))
	return nil
}

// Bool2Int64 is the exact inverse of Int642Bool
// a must point to a bool
// b must point to an int
func Bool2Int64(a, b *reflect.Value) error {
	i := int64(0)
	if false != (*a).Interface().(bool) {
		i = int64(1)
	}
	(*b).Set(reflect.ValueOf(i))
	return nil
}
