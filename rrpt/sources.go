package rrpt

import (
	"context"
	"fmt"
	"gotable"
	"rentroll/rlib"
)

// RRreportSources generates a report of all rlib.GLAccount accounts
func RRreportSources(ctx context.Context, ri *ReporterInfo) string {
	m, err := rlib.GetAllDemandSources(ctx, ri.Bid)
	if err != nil {
		return err.Error()
	}

	s := fmt.Sprintf("%-9s  %-9s  %-35s  %-35s\n", "SourceSLSID", "BID", "Name", "Industry")
	for i := 0; i < len(m); i++ {
		switch ri.OutputFormat {
		case gotable.TABLEOUTTEXT:
			s += fmt.Sprintf("S%08d  B%08d  %-35s  %-35s\n", m[i].SourceSLSID, m[i].BID, m[i].Name, m[i].Industry)
		case gotable.TABLEOUTHTML:
			fmt.Printf("UNIMPLEMENTED\n")
		default:
			fmt.Printf("RRreportSources: unrecognized print format: %d\n", ri.OutputFormat)
			return ""
		}
	}
	return s
}
