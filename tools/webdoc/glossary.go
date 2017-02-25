package main

import (
	"fmt"
	"html/template"
	"rentroll/rlib"
	"strings"
)

// This package reads in the RentRoll glossary and builds a map
// so that it can be indexed by abbreviations.

// GlossaryDef defines an individual glossary term data structure
type GlossaryDef struct {
	Term         string        // full term
	Abbreviation string        // abbreviated version of the term
	Definition   template.HTML // definition of the term
	Selections   string        // list of possible values if the term is an enumerated list of choices
	Example      string        // example of the term
	Module       string        // modules where this term applies
}

// Term et all are constants indicating the column in which the data can be found
const (
	Term         = 0
	Abbreviation = iota
	Definition   = iota
	Selections   = iota
	Example      = iota
	Module       = iota
)

// Glossary is the slice with all terms from the glossary
var Glossary []*GlossaryDef

// GlossaryAbbr is a map that can be indexed by abbreviations.
// The index string should be lower case characters
var GlossaryAbbr = map[string]*GlossaryDef{}

// GlossaryTerm is a map that can be indexed by the term.
// The index string should be lower case characters and all
// spaces should be removed
var GlossaryTerm = map[string]*GlossaryDef{}

// LoadGlossary loads the supplied file name into memory
// and creates a map
func LoadGlossary(fname string) error {
	var k int
	t := rlib.LoadCSV(fname)
	for k = 0; k < len(t); k++ {
		if strings.ToLower(t[k][Term]) == "term" {
			break
		}
	}
	if k >= len(t) {
		return fmt.Errorf("Could not find Column Header starting with Term")
	}
	for i := k + 1; i < len(t); i++ {
		var g GlossaryDef
		ta := t[i]
		g.Term = ta[Term]
		g.Abbreviation = ta[Abbreviation]
		g.Definition = template.HTML(ta[Definition])
		g.Selections = ta[Selections]
		g.Example = ta[Example]
		g.Module = ta[Module]
		termIndex := strings.ToLower(rlib.Stripchars(g.Term, " "))
		Glossary = append(Glossary, &g)
		if len(g.Abbreviation) > 0 {
			GlossaryAbbr[strings.ToLower(g.Abbreviation)] = &g
		}
		GlossaryTerm[termIndex] = &g
	}
	return nil
}

// getTermDefinition searches the glossary for the exact string it is provided.
// It returns the definition it finds or a null string if no match is found.
func getTermDefinition(f string) template.HTML {
	fp, ok := GlossaryAbbr[f]
	if ok {
		return template.HTML((*fp).Definition)
	}
	fp, ok = GlossaryTerm[f]
	if ok {
		return template.HTML((*fp).Definition)
	}
	return template.HTML("")
}

// getDefinition looks for term in the glossary maps. It first searches
// for the string that is the result of downshifting the characters and removing
// any spaces. So, a string like "Gen Search .Field" becomes "gensearch.field"
// If this search fails, then it will strip the struct prefix and search for the
// term after the dot.  That is, "gensearch.field" becomes "field".
// It will return the definition if it finds one. Otherwise
// it returns an empty string
func getDefinition(term string) template.HTML {
	f := strings.ToLower(rlib.Stripchars(term, " ")) // remove spaces and downshift chars
	for strings.Contains(f, "&nbsp;") {              // remove non-breaking spaces
		f = f[6:]
	}
	// fmt.Printf("getDefinition: search for %s... ", f)
	if def := getTermDefinition(f); len(def) > 0 {
		// fmt.Printf("%t!\n", len(def) > 0)
		return def
	}

	// see if we can identified a field name.  If so, remove everything up through
	// and including the last dot and check the glossary again.
	for strings.Contains(f, ".") {
		if i := strings.Index(f, "."); i >= 0 && i+1 < len(f) {
			f = f[i+1:]
			if def := getTermDefinition(f); len(def) > 0 {
				// fmt.Printf("%t!\n", len(def) > 0)
				return def
			}
		}
	}
	// fmt.Printf("FALSE!\n")
	return template.HTML("")
}

// IsGlossaryTerm checks to see if the supplied string can be found in
// the glossary as either a term or an abbreviation.  It returns true
// if the term was found, false otherwise.
func IsGlossaryTerm(t string) bool {
	s := strings.ToLower(t)
	_, ok := GlossaryAbbr[s]
	if ok {
		return true
	}
	_, ok = GlossaryTerm[s]
	return ok
}
