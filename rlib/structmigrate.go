package rlib

import (
	"fmt"
	"reflect"
	"strings"
)

var xjson = string("XJSON")

// XJSONAssignmentTime is a UI converter: backend int64, Front End string
type XJSONAssignmentTime string

var assignmap = []struct {
	a      string
	b      string
	mapper func(a, b *reflect.Value)
}{
	{a: "XJSONAssignmentTime", b: "int64", mapper: AssignmentTime2Int64},
	{a: "int64", b: "XJSONAssignmentTime", mapper: Int642AssignmentTime},
}

// XJSONprocess attempts to map a to b.
func XJSONprocess(a, b *reflect.Value) {
	at := (*a).Type().String()
	bt := (*b).Type().String()
	for i := 0; i < len(assignmap); i++ {
		if strings.Index(at, assignmap[i].a) >= 0 && strings.Index(bt, assignmap[i].b) >= 0 {
			assignmap[i].mapper(a, b)
			return
		}
	}
	s := fmt.Sprintf("XJSONmap - no conversion between: %s and %s\n", at, bt)
	fmt.Printf(s)
	panic(s)
}

// AssignmentTime2Int64 converter
// a must be *AssignmentTime2Int64
// b must be *int64
func AssignmentTime2Int64(a, b *reflect.Value) {
	s1 := (*a).Interface()
	s := fmt.Sprintf("%v", s1)
	var y int64
	if s == "Pre-Assign" {
		y = 1
	} else if s == "Commencement" {
		y = 2
	} else {
		y = 0
	}
	(*b).Set(reflect.ValueOf(y))
}

// Int642AssignmentTime converter
// a must be *int64
// b must be *Int642AssignmentTime
func Int642AssignmentTime(a, b *reflect.Value) {
	var s XJSONAssignmentTime
	switch (*a).Interface().(int64) {
	case int64(1):
		s = XJSONAssignmentTime("Pre-Assign")
	case int64(2):
		s = XJSONAssignmentTime("Commencement")
	default:
		s = XJSONAssignmentTime("unset")
	}
	(*b).Set(reflect.ValueOf(s))
}

// BuildFieldMap creates a map so that we can find
// a field's index using its name as the map index
//--------------------------------------------------------------------
func BuildFieldMap(p interface{}) map[string]int {
	var fmap = map[string]int{}
	v := reflect.ValueOf(p).Elem()
	for j := 0; j < v.NumField(); j++ {
		n := v.Type().Field(j).Name
		fmap[n] = j
	}
	return fmap
}

// MigrateStructVals copies values from pa to pb where the field
// names for whatever pa is pointing to matches the field name in pb
// There is a basic assumption that the data will either copy directly
// or convert cleanly from one struct to another.
//--------------------------------------------------------------------
func MigrateStructVals(pa interface{}, pb interface{}) error {
	m := BuildFieldMap(pb)
	ar := reflect.ValueOf(pa).Elem()
	for i := 0; i < ar.NumField(); i++ {
		fa := ar.Field(i)
		afldname := ar.Type().Field(i).Name
		if !fa.IsValid() {
			continue
		}
		bdx, ok := m[afldname]
		if !ok {
			continue
		}
		br := reflect.ValueOf(pb).Elem()
		fb := br.Field(bdx)
		if !fb.CanSet() {
			continue
		}
		if fa.Type() == fb.Type() {
			fb.Set(reflect.ValueOf(fa.Interface()))
		} else {
			ta := fa.Type().String()
			tb := fb.Type().String()
			if strings.Index(ta, xjson) >= 0 || strings.Index(tb, xjson) >= 0 {
				XJSONprocess(&fa, &fb)
			} else {
				val := reflect.ValueOf(fa.Interface())
				fb.Set(val.Convert(fb.Type()))
			}
		}
	}
	return nil
}
