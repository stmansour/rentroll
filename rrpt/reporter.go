package rrpt

import "fmt"

// TextReportColumn is a struct defining a column in a text report
type TextReportColumn struct {
	Name    string // name - will be the column name
	Type    string // printf type:  d, s, f, ...
	Width   int    // how wide to make the column
	Justify int    // 0 = left, 1 = right
	p       string // the string to use in Printf-style routines for this column
}

// TextReport is a collection of columns
type TextReport struct {
	Cols    []TextReportColumn // defines all the columns
	Spacing int                // space between columns
	Fmt     string             // printf fmt string
	Hdr     string             // column header string
	Line    string             // dash line "-"
	Length  int                // total width in columns
}

// AddColumn is a method of TextReport to add a column
func (t *TextReport) AddColumn(n, y string, w, j int) {
	var a TextReportColumn
	a.Name = n
	a.Type = y
	a.Width = w
	m := len(a.Name) // if the length of the column title
	if m > w {       // is greater than the user-specified width
		a.Width = m //increase the column width to hold the column title
	}
	a.Justify = j
	lr := ""
	if a.Justify == 0 {
		lr = "-"
	}
	a.p = fmt.Sprintf("%%%s%d%s", lr, a.Width, a.Type)
	t.Cols = append(t.Cols, a)
	t.SetFormat()
}

// SetFormat is called when a column is added or changed. It regenerates
// the format string used for printf.
func (t *TextReport) SetFormat() {
	t.Fmt = ""
	t.Hdr = ""
	sp := ""
	for i := 0; i < t.Spacing; i++ {
		sp += " "
	}
	end := len(t.Cols) - 1
	t.Length = 0
	for i := 0; i < len(t.Cols); i++ {
		t.Fmt += t.Cols[i].p // build up the column data printf string

		// build up the column header string
		lr := ""
		if t.Cols[i].Justify == 0 {
			lr = "-"
		}
		s := fmt.Sprintf("%%%s%ds", lr, t.Cols[i].Width)
		t.Hdr += fmt.Sprintf(s, t.Cols[i].Name)
		t.Length += t.Cols[i].Width

		// add spacing to next column, if necessary
		if i < end {
			t.Fmt += sp
			t.Hdr += sp
			t.Length += t.Spacing
		}
	}
	t.Fmt += "\n"
	t.Hdr += "\n"

	t.Line = ""
	for i := 0; i < t.Length; i++ {
		t.Line += "-"
	}
	t.Line += "\n"

	// debug
	// fmt.Printf("SetFormat: Fmt = %s\n", t.Fmt)
	// fmt.Printf("           Hdr = %s\n", t.Hdr)
	// fmt.Printf("          Line = %s\n", t.Line)
}

// Printf works just like fmt.Printf only it applies the formatting already set up for each column.
// You must supply exactly the number of arguments as columns, and in the correct order, just like fmt.Printf
func (t *TextReport) Printf(a ...interface{}) {
	fmt.Printf(t.Fmt, a...)
}

// PrintColHdr prints all the column headers
func (t *TextReport) PrintColHdr() {
	fmt.Print(t.Hdr)
}

// PrintLine prints a single dashed line across the entire width of the table.
func (t *TextReport) PrintLine() {
	fmt.Print(t.Line)
}
