package rrpt

import (
	"fmt"
	"strings"
)

// TextReportColumn is a struct defining a column in a text report
type TextReportColumn struct {
	Name    []string // name - will be the column name
	Type    string   // printf type:  d, s, f, ...
	Width   int      // how wide to make the column
	Justify int      // 0 = left, 1 = right
	p       string   // the string to use in Printf-style routines for this column
}

// TextReport is a collection of columns
type TextReport struct {
	Cols       []TextReportColumn // defines all the columns
	Spacing    int                // space between columns
	Fmt        string             // printf fmt string
	Hdr        []string           // column header string
	maxHdrRows int                // number of strings in Hdr
	Line       string             // dash line "-"
	Length     int                // total width in columns
}

// AddColumn is a method of TextReport to add a column
func (t *TextReport) AddColumn(n, y string, w, just int) {
	var a TextReportColumn

	//--------------------------------------------
	// break the column header up into words
	//--------------------------------------------
	sa := strings.Split(n, " ") // break up the string at the spaces

	j := 0
	maxColWidth := 0
	for i := 0; i < len(sa); i++ { // for each substring
		if len(sa[i]) < w && i+1 < len(sa) { // if the width of the substring is less than the requested width, and we're not at the end of the list
			if len(sa[i])+len(sa[i+1])+1 < w { // is there enough room for the next word in the list?
				a.Name = append(a.Name, sa[i]+" "+sa[i+1])
				i++ // skip the next element of sa since we've already added it
			} else {
				a.Name = append(a.Name, sa[i])
			}
		} else {
			a.Name = append(a.Name, sa[i])
		}
		if len(a.Name[j]) > maxColWidth {
			maxColWidth = len(a.Name[j])
		}
		j++
	}

	a.Type = y
	a.Width = w
	if maxColWidth > w { // if the length of the column title is greater than the user-specified width
		a.Width = maxColWidth //increase the column width to hold the column title
	}
	a.Justify = just
	lr := ""
	if a.Justify == 0 {
		lr = "-"
	}
	precision := ""
	if a.Type == "s" {
		precision = fmt.Sprintf(".%d", a.Width)
	}
	a.p = fmt.Sprintf("%%%s%d%s%s", lr, a.Width, precision, a.Type)
	t.Cols = append(t.Cols, a)
	t.SetFormat()
}

// AdjustColHdr formats the column names for printing. It will attempt to break up the column headers
// into multiple lines if necessary.
func (t *TextReport) AdjustColHdr() {
	//----------------------------------
	// Which column has the most rows?
	//----------------------------------
	t.maxHdrRows = 0
	for i := 0; i < len(t.Cols); i++ {
		j := len(t.Cols[i].Name)
		if j > t.maxHdrRows {
			t.maxHdrRows = j
		}
	}

	//---------------------------------------------
	// Set all columns to that number of rows...
	//---------------------------------------------
	for i := 0; i < len(t.Cols); i++ {
		n := make([]string, t.maxHdrRows)
		lenOrig := len(t.Cols[i].Name)
		iStart := t.maxHdrRows - lenOrig
		if iStart > 0 {
			for j := 0; j < iStart; j++ {
				n[j] = ""
			}
		}
		for j := iStart; j < t.maxHdrRows; j++ {
			n[j] = t.Cols[i].Name[j-iStart]
		}
		t.Cols[i].Name = n
	}
}

// SetFormat is called when a column is added or changed. It regenerates
// the format string used for printf.
func (t *TextReport) SetFormat() {
	end := len(t.Cols) - 1
	t.AdjustColHdr()
	t.Fmt = ""
	hdr := make([]string, t.maxHdrRows)
	t.Hdr = hdr
	sp := ""
	for i := 0; i < t.Spacing; i++ {
		sp += " "
	}
	t.Length = 0
	for i := 0; i < len(t.Cols); i++ {
		t.Fmt += t.Cols[i].p // build up the column data printf string

		// build up the column header string
		lr := ""
		if t.Cols[i].Justify == 0 {
			lr = "-"
		}
		s := fmt.Sprintf("%%%s%d.%ds", lr, t.Cols[i].Width, t.Cols[i].Width)
		for j := 0; j < t.maxHdrRows; j++ {
			t.Hdr[j] += fmt.Sprintf(s, t.Cols[i].Name[j])
		}
		t.Length += t.Cols[i].Width

		// add spacing to next column, if necessary
		if i < end {
			t.Fmt += sp
			t.Length += t.Spacing
			for j := 0; j < t.maxHdrRows; j++ {
				t.Hdr[j] += sp
			}
		}
	}
	t.Fmt += "\n"
	for j := 0; j < t.maxHdrRows; j++ {
		t.Hdr[j] += "\n"
	}

	t.Line = ""
	for i := 0; i < t.Length; i++ {
		t.Line += "-"
	}
	t.Line += "\n"
}

// Printf works just like fmt.Printf only it applies the formatting already set up for each column.
// You must supply exactly the number of arguments as columns, and in the correct order, just like fmt.Printf
func (t *TextReport) Printf(a ...interface{}) {
	fmt.Printf(t.Fmt, a...)
}

// PrintColHdr prints the column headers for all columns of the table
func (t *TextReport) PrintColHdr() {
	for i := 0; i < t.maxHdrRows; i++ {
		fmt.Print(t.Hdr[i])
	}
}

// PrintLine prints a single dashed line across the entire width of the table.
func (t *TextReport) PrintLine() {
	fmt.Print(t.Line)
}
