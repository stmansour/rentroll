package rcsv

import (
	"fmt"
	"strings"
)

func expandBraceArray(tmpl string) ([]string, error) {
	var m []string
	var l = len(tmpl)
	// rlib.Console("Entering expandBraceArray, tmpl = %q\n", tmpl)
	if tmpl[0] != '{' {
		return m, fmt.Errorf("Supplied string is missing initial '{'")
	}
	if tmpl[l-1] != '}' {
		return m, fmt.Errorf("Supplied string is missing closing '}'")
	}
	m = strings.Split(tmpl[1:l-1], "|")
	return m, nil
}

func updateStringMatchArray(m []string, s string) []string {
	var n []string
	if len(m) == 0 {
		n = append(n, s)
		return n
	}
	for i := 0; i < len(m); i++ {
		n = append(n, m[i]+s)
	}
	return n
}

// getStrMatchArray parses the header string into an array of acceptable
// header strings.  Strings in braces ({}) are separated by the vertical
// bar (|) and any of the strings between the braces are acceptable. Strings
// between brackets ([]) are optional.
//
// Sample test program: https://play.golang.org/p/rk6ggQ-o6GU
//
// Examples:
//
// tmpl = "Use{Status|Type}"  returns ["UseStatus","UseType"]
// tmpl = "Type[Ref]" returns ["Type", "TypeRef"]
// tmpe = "Rentable"  returns ["Rentable"]
//
// INPUTS:
// tmpl = template string
//
// RETURNS
// explanded array of acceptable strings
// any errors encountered
//-----------------------------------------------------------------------------
func getStrMatchArray(tmpl string) ([]string, error) {
	var m []string

	// rlib.Console("Entering getStrMatchArray, tmpl = %q\n", tmpl)

	for len(tmpl) > 0 {
		iBrace := strings.Index(tmpl, "{")
		iBracket := strings.Index(tmpl, "[")

		//-------------------------------
		// Any more special characters?
		//-------------------------------
		if iBrace < 0 && iBracket < 0 {
			m = updateStringMatchArray(m, tmpl)
			return m, nil
		}

		//-------------------------------------------------------
		// Was it a bracket, or if both was the bracket first...
		//-------------------------------------------------------
		if iBracket >= 0 && (iBrace < 0 || iBracket < iBrace) {
			var j int
			if j = strings.Index(tmpl, "]"); j < 0 {
				return m, fmt.Errorf("Missing closing bracket")
			}
			s := tmpl[:iBracket]
			opt := tmpl[iBracket+1 : j]
			n := updateStringMatchArray(m, s)
			n1 := updateStringMatchArray(m, s+opt)
			m = n
			m = append(m, n1...)
			tmpl = tmpl[j+1:]
			continue
		}

		if iBrace >= 0 {
			var j int
			if j = strings.Index(tmpl, "}"); j < 0 {
				return m, fmt.Errorf("Missing closing brace")
			}
			if iBrace > 0 {
				s := tmpl[:iBrace]
				m = updateStringMatchArray(m, s)
			}
			opts := tmpl[iBrace : j+1]
			// rlib.Console("iBrace = %d, j = %d, opts = %q\n", iBrace, j, opts)
			tmpl = tmpl[j+1:]
			optsa, err := expandBraceArray(opts)
			if err != nil {
				return m, err
			}
			var n [][]string
			for i := 0; i < len(optsa); i++ {
				n = append(n, updateStringMatchArray(m, optsa[i]))
			}
			m = n[0]
			for i := 1; i < len(n); i++ {
				m = append(m, n[i]...)
			}
			continue
		}
		fmt.Printf("PANIC!!  how did we get here???\n")
	}
	return m, nil
}
