package rlib

// This file manages the mapping between the programming data types
// and the way those types are displayed in the user interface.
//
// IF YOU ADD SOMETHING HERE, DO NOT FORGET TO UPDATE assignmap
import (
	"fmt"
	"reflect"
)

var assignmap = []struct {
	a      string
	b      string
	mapper func(a, b *reflect.Value) error
}{
	{a: "XJSONAssignmentTime", b: "int64", mapper: AssignmentTime2Int64},
	{a: "int64", b: "XJSONAssignmentTime", mapper: Int642AssignmentTime},
	{a: "int", b: "JSONbool", mapper: Int2Bool},
	{a: "JSONbool", b: "int", mapper: Int2Bool},
}

// XJSONAssignmentTime is a UI converter: backend int64, Front End string
type XJSONAssignmentTime string

// AssignmentTimeMap is the mapping for Rentable - AssignmentTime
var AssignmentTimeMap = Str2Int64Map{
	"unset":        0,
	"Pre-Assign":   1,
	"Commencement": 2,
}

// AssignmentTime2Int64 converter
// a must be *AssignmentTime2Int64
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
