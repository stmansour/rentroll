package main

// https://play.golang.org/p/KFpXRBJUJW

import (
	"encoding/json"
	"fmt"
	"html/template"
	"reflect"
	"strings"
)

// InitFields accepts an interface to a structure and generates zero-value
// instances for all fields so that it can be marshaled into a string
// showing all fields.
func InitFields(a interface{}) {
	v := reflect.ValueOf(a).Elem()
	for j := 0; j < v.NumField(); j++ {
		f := v.Field(j)
		t := f.Type().String()
		inst := t
		isArray := false
		if ix := strings.Index(t, "[]"); ix >= 0 {
			isArray = true
			inst = t[ix+2:]
		}
		strings.TrimPrefix(inst, "main.")

		// fmt.Printf("Name: %s  Kind: %s  Type: %s   inst: %s\n", v.Type().Field(j).Name, f.Kind(), t, inst)
		if inst != t {
			if isArray {
				x := reflect.MakeSlice(f.Type(), 1, 1)
				f.Set(x)
			}
		} else {
			switch f.Kind().String() {
			case "struct":
				// fmt.Printf("recurse on x, kind = struct, type = %s\n", inst)
				x, err := FactoryNew(inst)
				if err != nil {
					continue
				}
				InitFields(&x)
				y := reflect.ValueOf(x)
				f.Set(y)
			}
		}
	}
}

// GenExample produces a zero-valued formatted string that shows how the JSON will
// look for type t
func GenExample(t string) template.HTML {
	// fmt.Printf("GenExample:  t = %s\n", t)
	f, ok := WSTypeFactory[t]
	if !ok {
		// fmt.Printf("Returning empty string\n")
		return template.HTML("")
	}
	x := f()
	// fmt.Printf("Created x with type %s\n", t)
	InitFields(x)
	b, err := json.MarshalIndent(&x, "", "    ")
	if err != nil {
		return template.HTML(fmt.Sprintf("Error with json.MarshalIndent:  %s\n", err.Error()))
	}
	return template.HTML(b)
}
