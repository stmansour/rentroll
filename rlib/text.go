package rlib

import (
	"fmt"
	"sort"
)

// SprintTableText prints the whole table in text form
func (t *Table) SprintTableText(f int) string {
	// get headers first
	s, err := t.SprintColumnHeaders(f)
	if err != nil {
		return err.Error()
	}

	// then append strings of rows
	rs, err := t.SprintRows(f)
	if err != nil {
		return err.Error()
	}
	s += rs

	return s
}

// SprintColHdrsText formats the column headers as text and returns the string
func (t *Table) SprintColHdrsText() (string, error) {
	t.AdjustAllColumnHeaders()
	s := ""
	for j := 0; j < len(t.ColDefs[0].Hdr); j++ {
		for i := 0; i < len(t.ColDefs); i++ {
			sf := ""
			lft := ""
			if t.ColDefs[i].Justify == COLJUSTIFYLEFT {
				lft += "-"
			}
			sf += fmt.Sprintf("%%%s%ds", lft, t.ColDefs[i].Width)
			s += fmt.Sprintf(sf, t.ColDefs[i].Hdr[j])
			if i < len(t.ColDefs)-1 {
				s += Mkstr(t.TextColSpace, ' ')
			}
		}
		s += "\n"
	}
	for i := 0; i < len(t.ColDefs); i++ {
		s += fmt.Sprintf("%s", Mkstr(t.ColDefs[i].Width, '-'))
		if i < len(t.ColDefs)-1 {
			s += Mkstr(t.TextColSpace, ' ')
		}
	}
	s += "\n"
	return s, nil
}

// SprintRowsText returns all rows text string
func (t *Table) SprintRowsText(f int) (string, error) {
	rowsStr := ""
	for i := 0; i < t.Rows(); i++ {
		rowsStr += t.SprintRow(i, f)
	}
	return rowsStr, nil
}

// SprintRowText formats the requested row as text in a string and returns the string
func (t *Table) SprintRowText(row int) string {

	s := ""
	if len(t.LineBefore) > 0 {
		j := sort.SearchInts(t.LineBefore, row)
		if j < len(t.LineBefore) && row == t.LineBefore[j] {
			s += t.SprintLineText()
		}
	}

	rowColumns := len(t.Row[row].Col)

	// columns string chunk map, each column holds list of string
	// that fits in one line at best
	colStringChunkMap := map[int][]string{}

	// get Height of row that require to fit the content of max cell string content
	// by default table has no all the data in string format, so that we need to add
	// logic here only, to support multi line functionality
	for i := 0; i < rowColumns; i++ {
		if t.Row[row].Col[i].Type == CELLSTRING {
			cd := t.ColDefs[i]
			a, _ := t.getMultiLineText(t.Row[row].Col[i].Sval, cd.Width)

			colStringChunkMap[i] = a

			if len(a) > t.Row[row].Height {
				t.Row[row].Height = len(a)
			}
		}
	}

	// rowTextList holds the 2D array, containing data for each block
	// to achieve multiline row
	rowTextList := [][]string{}

	// init rowTextList with empty values
	for k := 0; k < t.Row[row].Height; k++ {
		temp := make([]string, rowColumns)
		for i := 0; i < rowColumns; i++ {
			// assign default string with length of column width
			temp = append(temp, Mkstr(t.ColDefs[i].Width, ' '))
		}
		rowTextList = append(rowTextList, temp)
	}

	// fill the content in rowTextList for the first line
	for i := 0; i < rowColumns; i++ {
		switch t.Row[row].Col[i].Type {
		case CELLFLOAT:
			rowTextList[0][i] = fmt.Sprintf(t.ColDefs[i].Pfmt, RRCommaf(t.Row[row].Col[i].Fval))
		case CELLINT:
			rowTextList[0][i] = fmt.Sprintf(t.ColDefs[i].Pfmt, t.Row[row].Col[i].Ival)
		case CELLSTRING:
			rowTextList[0][i] = fmt.Sprintf(t.ColDefs[i].Pfmt, colStringChunkMap[i][0])
		case CELLDATE:
			rowTextList[0][i] = fmt.Sprintf("%*.*s", t.ColDefs[i].Width, t.ColDefs[i].Width, t.Row[row].Col[i].Dval.Format(t.DateFmt))
		case CELLDATETIME:
			rowTextList[0][i] = fmt.Sprintf("%*.*s", t.ColDefs[i].Width, t.ColDefs[i].Width, t.Row[row].Col[i].Dval.Format(t.DateTimeFmt))
		default:
			rowTextList[0][i] = Mkstr(t.ColDefs[i].Width, ' ')
		}
	}

	// rowTextList to string
	for k := 0; k < t.Row[row].Height; k++ {
		for i := 0; i < rowColumns; i++ {

			// if not first row then process multi line format
			if k > 0 {
				if t.Row[row].Col[i].Type == CELLSTRING {
					if k >= len(colStringChunkMap[i]) {
						rowTextList[k][i] = fmt.Sprintf(t.ColDefs[i].Pfmt, "")
					} else {
						rowTextList[k][i] = fmt.Sprintf(t.ColDefs[i].Pfmt, colStringChunkMap[i][k])
					}
				}
			}

			// if blank then append string of column width with blank
			if rowTextList[k][i] == "" {
				rowTextList[k][i] = Mkstr(t.ColDefs[i].Width, ' ')
			}
			s += rowTextList[k][i]

			// if it is not last block then
			if i < (rowColumns - 1) {
				s += Mkstr(t.TextColSpace, ' ')
			}
		}
		s += "\n"
	}

	if len(t.LineAfter) > 0 {
		j := sort.SearchInts(t.LineAfter, row)
		if j < len(t.LineAfter) && row == t.LineAfter[j] {
			s += t.SprintLineText()
		}
	}
	return s
}

// SprintLineText returns a line across all rows in the table
func (t *Table) SprintLineText() string {
	s := ""
	for i := 0; i < len(t.ColDefs); i++ {
		s += Mkstr(t.ColDefs[i].Width, '-')
		if i < len(t.ColDefs)-1 {
			s += Mkstr(t.TextColSpace, ' ')
		}
	}
	return s + "\n"
}
