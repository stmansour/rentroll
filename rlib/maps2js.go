package rlib

import "fmt"

// StrToInt64MapList is a list of all StrToInt64MapList structures for
// which we need to build lists in the front end. The front end may
// need to build a dropdown list with these values
type StrToInt64MapList struct {
	JSname string
	t      Str2Int64Map
}

var maps = []StrToInt64MapList{
	{"AssignmentTimeList", AssignmentTimeMap},
}

// MapsToJS writes to stdout the javascript code that creates string arrays
// for all Str2Int64Map values in maps
func MapsToJS() {
	for i := 0; i < len(maps); i++ {
		fmt.Printf("%s = [ ", maps[i].JSname)
		L := len(maps[i].t)
		j := 0
		for k := range maps[i].t {
			j++
			fmt.Printf("'%s'", k)
			if j < L {
				fmt.Printf(", ")
			}
		}
		fmt.Printf("];\n")
	}
}
