package rrpt

import (
	"context"
	"fmt"
	"gotable"
	"rentroll/rlib"
)

// ReportNoteTypeToText returns a string representation of the chart of accts
func ReportNoteTypeToText(p *rlib.NoteType) string {
	return fmt.Sprintf("NT%08d  B%08d  %-50s\n",
		p.NTID, p.BID, p.Name)
}

// RRreportNoteTypes generates a report of all rlib.GLAccount accounts
func RRreportNoteTypes(ctx context.Context, ri *ReporterInfo) string {
	m, err := rlib.GetAllNoteTypes(ctx, ri.Bid)
	if err != nil {
		return err.Error()
	}

	s := fmt.Sprintf("%-10s  %-9s  %-50s\n", "NTID", "BID", "Name")
	for i := 0; i < len(m); i++ {
		switch ri.OutputFormat {
		case gotable.TABLEOUTTEXT:
			s += ReportNoteTypeToText(&m[i])
		case gotable.TABLEOUTHTML:
			fmt.Printf("UNIMPLEMENTED\n")
		default:
			fmt.Printf("RRreportNoteTypes: unrecognized print format: %d\n", ri.OutputFormat)
			return ""
		}
	}
	return s
}
