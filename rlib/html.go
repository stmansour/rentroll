package rlib

import (
	"fmt"
	"strconv"
)

// SprintTableHTML prints the whole table in HTML form
func (t *Table) SprintTableHTML(f int) string {

	// get headers first
	s, err := t.SprintColHdrsHTML()
	if err != nil {
		return err.Error()
	}

	// then append table body
	rs, err := t.SprintRows(f)
	if err != nil {
		return err.Error()
	}
	s += rs

	// finally return HTML table layout
	return "<table>" + s + "</table>"
}

// SprintColHdrsHTML formats the requested row in HTML and returns the HTML as a string
func (t *Table) SprintColHdrsHTML() (string, error) {
	tHeader := ""
	for i := 0; i < len(t.ColDefs); i++ {
		cd := t.ColDefs[i]
		headerCell := t.ColDefs[i].ColTitle
		if cd.HTMLWidth != -1 {
			headerCell = "<th width=\"" + strconv.Itoa(cd.HTMLWidth) + "\">" + headerCell + "</th>"
		} else {
			headerCell = "<th>" + headerCell + "</th>"
		}
		tHeader += headerCell
	}
	return "<thead><tr>" + tHeader + "</tr></thead>", nil
}

// SprintRowsHTML returns all rows text string
func (t *Table) SprintRowsHTML(f int) (string, error) {
	rowsStr := ""
	for i := 0; i < t.Rows(); i++ {
		rowsStr += t.SprintRow(i, f)
	}
	return "<tbody>" + rowsStr + "</tbody>", nil
}

// SprintRowHTML formats the requested row in HTML and returns the HTML as a string
// REF: http://stackoverflow.com/questions/21033440/disable-automatic-change-of-width-in-table-tag
func (t *Table) SprintRowHTML(row int) string {

	tRow := ""

	// fill the content in rowTextList for the first line
	for i := 0; i < len(t.Row[row].Col); i++ {

		var rowCell string
		// append content in TD
		switch t.Row[row].Col[i].Type {
		case CELLFLOAT:
			rowCell = fmt.Sprintf(t.ColDefs[i].Pfmt, RRCommaf(t.Row[row].Col[i].Fval))
		case CELLINT:
			rowCell = fmt.Sprintf(t.ColDefs[i].Pfmt, t.Row[row].Col[i].Ival)
		case CELLSTRING:
			// ******************************************************
			// FOR HTML, APPEND FULL STRING, THERE ARE NO
			// MULTILINE TEXT IN THIS
			// ******************************************************
			rowCell = fmt.Sprintf("%s", t.Row[row].Col[i].Sval)
		case CELLDATE:
			rowCell = fmt.Sprintf("%*.*s", t.ColDefs[i].Width, t.ColDefs[i].Width, t.Row[row].Col[i].Dval.Format(t.DateFmt))
		case CELLDATETIME:
			rowCell = fmt.Sprintf("%*.*s", t.ColDefs[i].Width, t.ColDefs[i].Width, t.Row[row].Col[i].Dval.Format(t.DateTimeFmt))
		default:
			rowCell = Mkstr(t.ColDefs[i].Width, ' ')
		}

		// format td cell
		rowCell = "<td>" + rowCell + "</td>"
		tRow += rowCell
	}

	return "<tr>" + tRow + "</tr>"
}
