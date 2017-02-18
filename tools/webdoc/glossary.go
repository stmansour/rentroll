package main

import (
	"fmt"
	"rentroll/rlib"
	"strings"
)

// This package reads in the RentRoll glossary and builds a map
// so that it can be indexed by abbreviations.

// GlossaryDef defines an individual glossary term data structure
type GlossaryDef struct {
	Term         string // full term
	Abbreviation string // abbreviated version of the term
	Definition   string // definition of the term
	Selections   string // list of possible values if the term is an enumerated list of choices
	Example      string // example of the term
	Module       string // modules where this term applies
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
		g.Definition = ta[Definition]
		g.Selections = ta[Selections]
		g.Example = ta[Example]
		g.Module = ta[Module]
		termIndex := strings.ToLower(rlib.Stripchars(g.Term, ". "))
		Glossary = append(Glossary, &g)
		if len(g.Abbreviation) > 0 {
			GlossaryAbbr[strings.ToLower(g.Abbreviation)] = &g
		}
		GlossaryTerm[termIndex] = &g
	}
	return nil
}
