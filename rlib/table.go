package rlib

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"
)

// Table is a simple skeletal row-column "class" for go that implements a few
// useful methods for building, maintaining, and printing tables of data.
// To use this table, you must add all the columns first. Then call the AddRow
// method and begin adding, modifying, getting things at a row,col.
//
// You can insert rows, append rows, sort all or selected rows by their column
// values, and put lines

// COLJUSTIFYLEFT et. al. are the constants used in the Table class
const (
	COLJUSTIFYLEFT  = 1
	COLJUSTIFYRIGHT = 2

	CELLINT      = 1
	CELLFLOAT    = 2
	CELLSTRING   = 3
	CELLDATE     = 4
	CELLDATETIME = 5

	TABLEOUTTEXT = 1
	TABLEOUTHTML = 2
	TABLEOUTPDF  = 3
	TABLEOUTCSV  = 4
)

var (
	// ErrHeaders error
	ErrHeaders = errors.New("No Headers found in table")
	// ErrRows error
	ErrRows = errors.New("No Rows found in table")
	// ErrUnKnownFmt error
	ErrUnKnownFmt = errors.New("Unrecognized format")
	// ErrPDF error
	ErrPDF = errors.New("PDF output format yet not supported")
)

// Cell is the basic data value type for the Table class.
type Cell struct {
	Type int       // int, float, or string enumeration
	Ival int64     // integer value
	Fval float64   // float value
	Sval string    // string value
	Dval time.Time // datetime value
}

// ColumnDef defines a Table column -- a column title, justification, and formatting
// information for cells in the column.
type ColumnDef struct {
	ColTitle  string   // the column title
	Width     int      // column width for TEXT
	Justify   int      // justification
	Pfmt      string   // printf-style formatting information for values in this column
	CellType  int      // type of data in this column
	Hdr       []string // multiple lines of column headers as needed -- based on width and Title
	Fdecimals int      // the number of decimal digits for floating point numbers. The default is 2
	HTMLWidth int
}

// Colset defines a set of Cells
type Colset struct {
	Col    []Cell // 1 row's worth of Cells, contains len(Col) number of Cells
	Height int    // height of row
}

// Rowset defines a set of rows to be operated on at a later time.
type Rowset struct {
	R []int // the row numbers of interest
}

// Table is a structure that defines a spreadsheet-like grid of cells and the
// operations that can be performed.
type Table struct {
	Title        string      // table title
	Section1     string      // another section for the title, different font
	Section2     string      // a third section for the title, different font
	ColDefs      []ColumnDef // table's column definitions, ordered 0..n left to right
	Row          []Colset    // Each Colset forms a row
	TextColSpace int         // space between text columns
	maxHdrRows   int         // maximum number of header rows across all ColDefs
	DateFmt      string      // format for printing dates
	DateTimeFmt  string      // format for datetime values
	LineAfter    []int       // array of row numbers that have a horizontal line after they are printed
	LineBefore   []int       // array of row numbers that have a horizontal line before they are printed
	RS           []Rowset    // a list of rowsets
}

// SetTitle sets the table's Title string to the supplied value
func (t *Table) SetTitle(s string) {
	t.Title = s
}

// GetTitle sets the table's Title string to the supplied value
func (t *Table) GetTitle() string {
	return t.Title
}

// SetSection1 sets the table's Section1 string to the supplied value
func (t *Table) SetSection1(s string) {
	t.Section1 = s
}

// GetSection1 sets the table's Section1 string to the supplied value
func (t *Table) GetSection1() string {
	return t.Section1
}

// SetSection2 sets the table's Section2 string to the supplied value
func (t *Table) SetSection2(s string) {
	t.Section2 = s
}

// GetSection2 sets the table's Section2 string to the supplied value
func (t *Table) GetSection2() string {
	return t.Section2
}

// RowCount returns the number of rows in the table
func (t *Table) RowCount() int {
	return len(t.Row)
}

// TypeToString returns a string describing the data type of the cell.
func (c *Cell) TypeToString() string {
	switch c.Type {
	case CELLSTRING:
		return "string"
	case CELLINT:
		return "int"
	case CELLFLOAT:
		return "float"
	case CELLDATE:
		return "date"
	case CELLDATETIME:
		return "datetime"
	}
	return "unknown"
}

// Init sets internal formatting controls to their default values
func (t *Table) Init() {
	t.TextColSpace = 2
	t.DateFmt = "01/02/2006"
	t.DateTimeFmt = "01/02/2006 15:04:00 MST"
}

// AddLineAfter keeps track of the row numbers after which a line will be printed
func (t *Table) AddLineAfter(row int) {
	t.LineAfter = append(t.LineAfter, row)
	sort.Ints(t.LineAfter)
}

// AddLineBefore keeps track of the row numbers before which a line will be printed
func (t *Table) AddLineBefore(row int) {
	t.LineBefore = append(t.LineBefore, row)
	sort.Ints(t.LineBefore)
}

// CreateRowset creates a new rowset. You can add row indeces to it.  You can process the rows at those indeces later.
// The return value is the Rowset identifier; rsid.  Use it to refer to this rowset.
func (t *Table) CreateRowset() int {
	var a Rowset
	t.RS = append(t.RS, a)
	return len(t.RS) - 1
}

// AppendToRowset adds a new row index to the rowset rsid
func (t *Table) AppendToRowset(rsid, row int) {
	t.RS[rsid].R = append(t.RS[rsid].R, row)
}

// SumRowset computes the sum of the rows in rowset[rs] at the specified column index. It returns a Cell with the sum
func (t *Table) SumRowset(rsid, col int) Cell {
	var c Cell
	for i := 0; i < len(t.RS[rsid].R); i++ {
		row := t.RS[rsid].R[i]
		switch t.Row[row].Col[col].Type {
		case CELLINT:
			c.Type = CELLINT
			c.Ival += t.Row[row].Col[col].Ival
		case CELLFLOAT:
			c.Type = CELLFLOAT
			c.Fval += t.Row[row].Col[col].Fval
		}
	}
	return c
}

// InsertSumRowsetCols sums the values for the specified rowset and appends it at the specified row
func (t *Table) InsertSumRowsetCols(rsid, row int, cols []int) {
	t.InsertRow(row)
	for i := 0; i < len(cols); i++ {
		c := t.SumRowset(rsid, cols[i])
		t.Put(row, cols[i], c)
	}
}

// AdjustFormatString can be called when the format string is null or when the column width changes
// to set a proper formatting string
func (t *Table) AdjustFormatString(cd *ColumnDef) {
	lft := ""
	if cd.Justify == COLJUSTIFYLEFT {
		lft += "-"
	}
	switch cd.CellType {
	case CELLINT:
		cd.Pfmt = fmt.Sprintf("%%%s%dd", lft, cd.Width)
	case CELLFLOAT:
		cd.Pfmt = fmt.Sprintf("%%%d.%ds", cd.Width, cd.Width)
	case CELLSTRING:
		cd.Pfmt = fmt.Sprintf("%%%s%d.%ds", lft, cd.Width, cd.Width)
	}
}

// AddColumn adds a new ColumnDef to the table
func (t *Table) AddColumn(title string, width, celltype int, justification int) {
	var cd = ColumnDef{
		ColTitle: title, Width: width, CellType: celltype,
		Justify: justification, Fdecimals: 2, HTMLWidth: -1,
	}
	t.AdjustColumnHeader(&cd)
	t.AdjustFormatString(&cd)
	t.ColDefs = append(t.ColDefs, cd)
}

// AdjustColumnHeader will break up the header into multiple lines if necessary to
// make the title fit.  If necessary, it will force the width of the column to be
// wide enough to fit the longest word in the title.
func (t *Table) AdjustColumnHeader(cd *ColumnDef) {
	a, maxColWidth := t.getMultiLineText(cd.ColTitle, cd.Width)
	if maxColWidth > cd.Width { // if the length of the column title is greater than the user-specified width
		cd.Width = maxColWidth //increase the column width to hold the column title
	}
	cd.Hdr = a
}

// AdjustAllColumnHeaders formats the column names for printing. It will attempt to break up the column headers
// into multiple lines if necessary.
func (t *Table) AdjustAllColumnHeaders() {
	//----------------------------------
	// Which column has the most rows?
	//----------------------------------
	t.maxHdrRows = 0
	for i := 0; i < len(t.ColDefs); i++ {
		j := len(t.ColDefs[i].Hdr)
		if j > t.maxHdrRows {
			t.maxHdrRows = j
		}
	}

	//---------------------------------------------
	// Set all columns to that number of rows...
	//---------------------------------------------
	for i := 0; i < len(t.ColDefs); i++ {
		n := make([]string, t.maxHdrRows)
		lenOrig := len(t.ColDefs[i].Hdr)
		iStart := t.maxHdrRows - lenOrig
		// Create a new Hdr array, n.
		// Add any initial blank lines...
		if iStart > 0 {
			for j := 0; j < iStart; j++ {
				n[j] = ""
			}
		}
		// now add the remaining strings
		for j := iStart; j < t.maxHdrRows; j++ {
			n[j] = t.ColDefs[i].Hdr[j-iStart]
		}
		t.ColDefs[i].Hdr = n // replace the old hdr with the new one
	}
}

// AddRow appends a new Row to the table. Initially, all cells are empty
func (t *Table) AddRow() {
	var c Colset
	for i := 0; i < len(t.ColDefs); i++ {
		var cell Cell
		c.Col = append(c.Col, cell)
	}
	c.Height = 1
	t.Row = append(t.Row, c)
}

// Get returns the cell at the supplied row,col.  If the supplied
// row or col is outside the table's boundaries, then an empty cell
// is returned
func (t *Table) Get(row, col int) Cell {
	if row >= len(t.Row) || col >= len(t.ColDefs) {
		var c Cell
		return c
	}
	return t.Row[row].Col[col]
}

// Geti returns the int at the supplied row,col.  If the supplied
// row or col is outside the table's boundaries, then 0 is returned
func (t *Table) Geti(row, col int) int64 {
	if row >= len(t.Row) || col >= len(t.ColDefs) {
		return int64(0)
	}
	return t.Row[row].Col[col].Ival
}

// Getf returns the floatval at the supplied row,col.  If the supplied
// row or col is outside the table's boundaries, then 0
// is returned
func (t *Table) Getf(row, col int) float64 {
	if row >= len(t.Row) || col >= len(t.ColDefs) {
		return float64(0)
	}
	return t.Row[row].Col[col].Fval
}

// Gets returns the strinb value at the supplied row,col.  If the supplied
// row or col is outside the table's boundaries, then ""
// is returned
func (t *Table) Gets(row, col int) string {
	if row >= len(t.Row) || col >= len(t.ColDefs) {
		return ""
	}
	return t.Row[row].Col[col].Sval
}

// Getd returns the date at the supplied row,col.  If the supplied
// row or col is outside the table's boundaries, then a 0 date
func (t *Table) Getd(row, col int) time.Time {
	if row >= len(t.Row) || col >= len(t.ColDefs) {
		return time.Date(0, time.January, 0, 0, 0, 0, 0, time.UTC)
	}
	return t.Row[row].Col[col].Dval
}

// Type returns the data type for the cell at the supplied row,col.
// If the supplied row or col is outside the table's boundaries, then 0
// is returned
func (t *Table) Type(row, col int) int {
	if row >= len(t.Row) || col >= len(t.ColDefs) {
		return 0
	}
	return t.Row[row].Col[col].Type
}

// Puti updates the Cell at row,col with the int64 value v
// and sets its type to CELLINT. If row or col is out of
// bounds the return value is false. Otherwise, the return
// value is true
func (t *Table) Puti(row, col int, v int64) bool {
	if row >= len(t.Row) || col >= len(t.ColDefs) {
		return false
	}
	if row < 0 {
		row = len(t.Row) - 1
	}
	t.Row[row].Col[col].Type = CELLINT
	t.Row[row].Col[col].Ival = v
	return true
}

// Putf updates the Cell at row,col with the float64 value v
// and sets its type to CELLFLOAT.
// if row < 0 then row is set to the last row of the table.
// If row or col is out of
// bounds the return value is false. Otherwise, the return
// value is true.
func (t *Table) Putf(row, col int, v float64) bool {
	if row >= len(t.Row) || col >= len(t.ColDefs) {
		return false
	}
	if row < 0 {
		row = len(t.Row) - 1
	}
	t.Row[row].Col[col].Type = CELLFLOAT
	t.Row[row].Col[col].Fval = v
	return true
}

func (t *Table) getMultiLineText(v string, colWidth int) ([]string, int) {
	var a []string

	// fit the content in one line whatever it is irrespective of column width
	if colWidth < 1 {
		a[0] = v
		return a, -1
	}

	// get multi line chunk in form of array
	sa := strings.Split(v, " ") // break up the string at the spaces
	j := 0
	maxColWidth := 0
	for i := 0; i < len(sa); i++ { // spin through all substrings
		if len(sa[i]) <= colWidth && i+1 < len(sa) { // if the width of this substring is less than the requested width, and we're not at the end of the list
			s := sa[i]                         // we know we're adding this one
			for k := i + 1; k < len(sa); k++ { // take as many as possible
				if len(s)+len(sa[k])+1 <= colWidth { // if it fits...
					s += " " + sa[k] // ...add it to the list...
					i = k            // ...and keep loop in sync
				} else {
					break // otherwise, add what we have and then go back to the outer loop
				}
			}
			a = append(a, s)
		} else {
			a = append(a, sa[i])
		}
		if len(a[j]) > maxColWidth { // if there's not enough room for the current string
			maxColWidth = len(a[j]) // then adjust the max column width we need
		}
		j++
	}
	return a, maxColWidth
}

// Puts updates the Cell at row,col with the string value v
// and sets its type to CELLSTRING. If row or col is out of
// bounds the return value is false. Otherwise, the return
// value is true
func (t *Table) Puts(row, col int, v string) bool {
	if row >= len(t.Row) || col >= len(t.ColDefs) {
		return false
	}
	if row < 0 {
		row = len(t.Row) - 1
	}
	t.Row[row].Col[col].Type = CELLSTRING
	t.Row[row].Col[col].Sval = v

	// Need to check width of column everytime when we adding new content
	// if it is updatable or not
	cd := t.ColDefs[col]
	_, cellWidth := t.getMultiLineText(v, cd.Width)
	if cellWidth > cd.Width { // if the length of the column title is greater than the user-specified width
		cd.Width = cellWidth //increase the column width to hold the column title
		t.AdjustFormatString(&cd)
		t.ColDefs[col] = cd
	}

	return true
}

// Putd updates the Cell at row,col with the date value v
// and sets its type to CELLDATE. If row or col is out of
// bounds the return value is false. Otherwise, the return
// value is true
func (t *Table) Putd(row, col int, v time.Time) bool {
	return t.putdint(row, col, v, CELLDATE)
}

// Putdt updates the Cell at row,col with the datetimv value v
// and sets its type to CELLDATETIME. If row or col is out of
// bounds the return value is false. Otherwise, the return
// value is true
func (t *Table) Putdt(row, col int, v time.Time) bool {
	return t.putdint(row, col, v, CELLDATETIME)
}

func (t *Table) putdint(row, col int, v time.Time, x int) bool {
	if row >= len(t.Row) || col >= len(t.ColDefs) {
		return false
	}
	if row < 0 {
		row = len(t.Row) - 1
	}
	t.Row[row].Col[col].Type = x
	t.Row[row].Col[col].Dval = v
	return true
}

// Put places Cell c at location row,col
func (t *Table) Put(row, col int, c Cell) {
	if row < 0 {
		row = len(t.Row) - 1
	}
	t.Row[row].Col[col] = c
}

// Rows returns the number of rows in the table
func (t *Table) Rows() int {
	return len(t.Row)
}

// Cols returns the number of columns in the table
func (t *Table) Cols() int {
	return len(t.ColDefs)
}

// String is the "stringer" method implementation for go so that you can simply
// print(t)
func (t Table) String() string {
	return t.Title + t.Section1 + t.Section2 + t.SprintTable(TABLEOUTTEXT)
}

// InsertRow adds a new Row at the specified index.
func (t *Table) InsertRow(row int) {
	if row >= len(t.Row) {
		t.AddRow()
		return
	}

	var c Colset
	for i := 0; i < len(t.ColDefs); i++ {
		var cell Cell
		c.Col = append(c.Col, cell)
	}

	t.Row = append(t.Row[:row+1], t.Row[row:]...)
	t.Row[row] = c
}

// Sum computes the sum of the rows at the specified column index. It returns a Cell
func (t *Table) Sum(col int) Cell {
	return t.SumRows(col, 0, len(t.Row)-1)
}

// SumRows computes the sum of rows 0 thru row at the specified column index. It returns a Cell
func (t *Table) SumRows(col, from, to int) Cell {
	var c Cell
	if from < 0 {
		from = 0
	}
	if to >= len(t.Row) {
		to = len(t.Row) - 1
	}
	for i := from; i <= to; i++ {
		switch t.Row[i].Col[col].Type {
		case CELLINT:
			c.Type = CELLINT
			c.Ival += t.Row[i].Col[col].Ival
		case CELLFLOAT:
			c.Type = CELLFLOAT
			c.Fval += t.Row[i].Col[col].Fval
		}
	}
	return c
}

// InsertSumRow inserts a new Row at index row, it then sums the specified columns in the Row range: from,to
// and sets the newly inserted row values at the specified columns to the sums.
func (t *Table) InsertSumRow(row, from, to int, cols []int) {
	t.InsertRow(row)
	for i := 0; i < len(cols); i++ {
		c := t.SumRows(cols[i], from, to)
		t.Put(row, cols[i], c)
	}
}

// Sort sorts rows (from,to) by column col ascending
func (t *Table) Sort(from, to, col int) {
	// fmt.Printf("Table.Sort:  from = %d, to = %d, col = %d,  len(t.Row) = %d\n", from, to, col, len(t.Row))
	var swap bool
	for i := from; i < to; i++ {
		for j := i + 1; j <= to; j++ {
			switch t.Row[i].Col[col].Type {
			case CELLINT:
				swap = t.Row[i].Col[col].Ival > t.Row[j].Col[col].Ival
			case CELLFLOAT:
				swap = t.Row[i].Col[col].Fval > t.Row[j].Col[col].Fval
			case CELLSTRING:
				swap = t.Row[i].Col[col].Sval > t.Row[j].Col[col].Sval
			case CELLDATE, CELLDATETIME:
				swap = t.Row[i].Col[col].Dval.After(t.Row[j].Col[col].Dval)
			}
			if swap {
				t.Row[i], t.Row[j] = t.Row[j], t.Row[i]
			}
		}
	}
}

// DeleteRow removes the table row at the specified index. All rowsets and LineAfter sets are adjusted.
// Cleanup on LineAfter and RowSets does not work if row == 0. I was just too lazy at the time to add this
// code because I know how/where delete will be used and it will not affect row 0.
func (t *Table) DeleteRow(row int) {
	if row == 0 {
		t.Row = t.Row[1:]
	} else {
		n := t.Row[0:row]
		if len(t.Row) > row {
			n = append(n, t.Row[row+1:]...)
		}
		t.Row = n
	}
	// Clean up LineAfter
	for i := 0; i < len(t.LineAfter); i++ {
		if t.LineAfter[i] >= row {
			t.LineAfter[i]--
		}
	}
	// Clean up RowSets
	for i := 0; i < len(t.RS); i++ {
		for j := 0; j < len(t.RS[i].R); j++ {
			if t.RS[i].R[j] >= row {
				t.RS[i].R[j]--
			}
		}
	}
}

// TightenColumns goes through all values in STRING columns and determines the maximum length in characters (max).
// If this length is less than the column width the column width is reduced to max.  This is
// mostly useful for text formatting.
func (t *Table) TightenColumns() {
	for i := 0; i < len(t.ColDefs); i++ {
		if t.ColDefs[i].CellType != CELLSTRING {
			continue
		}
		max := 0
		for j := 0; j < len(t.ColDefs[i].Hdr); j++ { // first, find the max len of the col hdrs
			l := len(t.ColDefs[i].Hdr[j])
			if max < l {
				max = l
			}
		}
		for j := 0; j < len(t.Row); j++ { // continue by find the max width of cell values in this col
			if t.Row[j].Col[i].Type == CELLSTRING {
				l := len(t.Row[j].Col[i].Sval)
				if max < l {
					max = l
				}
			}
		}
		if max < t.ColDefs[i].Width { // if the max width is less than the column width, contract the column width
			t.ColDefs[i].Width = max
		}
		cd := t.ColDefs[i]
		t.AdjustFormatString(&cd)
		t.ColDefs[i] = cd
	}
}

// SprintTable renders the entire table to a string
func (t *Table) SprintTable(f int) string {
	switch f {
	case TABLEOUTTEXT:
		return t.SprintTableText(f)
	case TABLEOUTHTML:
		return t.SprintTableHTML(f)
	case TABLEOUTCSV:
		return t.SprintTableCSV(f)
	case TABLEOUTPDF:
		return t.SprintTablePDF(f)
	}
	Ulog("SprintTable: unrecognized format:  %d\n", f)
	return ErrUnKnownFmt.Error()

}

// SprintColumnHeaders returns a string with the column headers formatted as type f
func (t *Table) SprintColumnHeaders(f int) (string, error) {

	// first check if there are any headers
	if len(t.ColDefs) < 1 {
		Ulog("there are no headers\n")
		return "", ErrHeaders
	}

	switch f {
	case TABLEOUTTEXT:
		return t.SprintColHdrsText()
	case TABLEOUTHTML:
		return t.SprintColHdrsHTML()
	case TABLEOUTCSV:
		return t.SprintColHdrsCSV()
	case TABLEOUTPDF:
		return t.SprintColHdrsPDF()
	}
	Ulog("SprintColumnHeaders unrecognized format:  %d\n", f)
	return "", ErrUnKnownFmt
}

// SprintRows returns a string formatted for all rows
func (t *Table) SprintRows(f int) (string, error) {

	// if there are no rows in table
	if t.Rows() < 1 {
		Ulog("SprintRows: there are no rows\n")
		return "", ErrRows
	}

	switch f {
	case TABLEOUTTEXT:
		return t.SprintRowsText(f)
	case TABLEOUTHTML:
		return t.SprintRowsHTML(f)
	case TABLEOUTCSV:
		return t.SprintRowsCSV(f)
	case TABLEOUTPDF:
		return t.SprintRowsPDF(f)
	}
	Ulog("SprintRows unrecognized format:  %d\n", f)
	return "", ErrUnKnownFmt
}

// SprintRow returns a string formatted for output type f with the information in row
func (t *Table) SprintRow(row, f int) string {

	if row < 0 {
		Ulog("SprintRow: row number is less than zero , row: %d\n", row)
		return ""
	}
	if row >= len(t.Row) {
		Ulog("SprintRow: row number > rows in table, row: %d\n", row)
		return ""
	}

	switch f {
	case TABLEOUTTEXT:
		return t.SprintRowText(row)
	case TABLEOUTHTML:
		return t.SprintRowHTML(row)
	case TABLEOUTCSV:
		return t.SprintRowCSV(row)
	case TABLEOUTPDF:
		return t.SprintRowPDF(row)
	}
	Ulog("SprintRow unrecognized format:  %d\n", f)
	return ErrUnKnownFmt.Error()
}

// func main() {
// var t Table
// dt := time.Date(2016, time.February, 14, 0, 0, 0, 0, time.UTC)
// dt1 := time.Date(2014, time.January, 3, 0, 0, 0, 0, time.UTC)
// dt2 := time.Date(2016, time.October, 23, 0, 0, 0, 0, time.UTC)
// t.Init()
// t.AddColumn("NAME", 20, COLJUSTIFYLEFT, CELLSTRING)
// t.AddColumn("YEARS OF AGE", 3, COLJUSTIFYRIGHT, CELLINT)
// t.AddColumn("WEIGHT", 8, COLJUSTIFYRIGHT, CELLFLOAT)
// t.AddColumn("CITY", 15, COLJUSTIFYLEFT, CELLSTRING)
// t.AddColumn("SPECIAL DAY", 10, COLJUSTIFYLEFT, CELLDATE)
// t.AddRow()
// t.Puts(0, 0, "Cletus")
// t.Puti(0, 1, 37)
// t.Putf(0, 2, 97.23)
// t.Puts(0, 3, "Springfield")
// t.Putd(0, 4, dt)
// t.AddRow()
// t.Puts(1, 0, "Dumbo")
// t.Puti(1, 1, 21)
// t.Putf(1, 2, 2957.8)
// t.Puts(1, 3, "Congo")
// t.Putd(1, 4, dt1)
// t.InsertRow(1)
// t.Puts(1, 0, "Bugs")
// t.Puti(1, 1, 7)
// t.Putf(1, 2, 3.4)
// t.Puts(1, 3, "El Segundo")
// t.Putd(1, 4, dt2)
// //fmt.Printf("\n\n%s", t)

// t.AddLineAfter(2)

// t.InsertSumRow(3, 0, 2, []int{1, 2})
// fmt.Printf("%s\n", t)

// t.Sort(0, 2, 4)
// fmt.Println(t)
// }
