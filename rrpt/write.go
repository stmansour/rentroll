package rrpt

import (
	"fmt"
	"gotable"
	"io"
	"rentroll/rlib"
	"strings"
)

// WritePDFReport writes the report to the supplied io.writer
//
// INPUTS:
//    w   - the file to write to
//    tsh - info about the needs of this particular report
//          set this value to nil to use the defaults
//    ri  - Report formatting info
//
// RETURNS:
//    any error encountered
//-----------------------------------------------------------------------------
func WritePDFReport(w io.Writer, tsh *SingleTableReportHandler, ri *ReporterInfo, rctx *ReportContext, tbl *gotable.Table) error {
	// custom template if available
	tfname := tsh.HTMLTemplate
	// rlib.Console("report.go:  tfname = %s\n", tfname)
	if len(tfname) > 0 {
		var err error
		bud := rlib.GetBUDFromBIDList(ri.Bid)
		tfname = "webclient/html/rpt-templates/" + strings.ToUpper(string(bud)) + "/" + tfname

		// cwd, err := os.Getwd()
		// rlib.Console("report.go:  cwd = %s\n", cwd)
		// rlib.Console("report.go:  setHTMLTemplate to %s\n", tfname)
		err = tbl.SetHTMLTemplate(tfname)
		if err != nil {
			s := fmt.Sprintf("Error in CSVprintTable: %s\n", err.Error())
			fmt.Print(s)
			fmt.Fprintf(w, "%s\n", s)
		}
	}

	// pdf props title
	var pdfProps []*gotable.PDFProperty
	if tsh.PDFprops != nil {
		pdfProps = tsh.PDFprops
	} else {
		pdfProps = GetReportPDFProps()
	}

	// get page size and orientation, set title
	if tsh.NeedsPDFTitle {
		// NOTE: There is no support of custom title right now
		pdfProps = SetPDFOption(pdfProps, "--header-center", tbl.Title)
	}

	// if custom dimension needy, then get those from client side
	if tsh.NeedsCustomPDFDimension {
		// pdf page width from the UI
		pdfPageWidth := rlib.Float64ToString(rctx.PDFPageWidth) + rctx.PDFPageSizeUnit
		pdfProps = SetPDFOption(pdfProps, "--page-width", pdfPageWidth)

		// pdf page height from the UI
		pdfPageHeight := rlib.Float64ToString(rctx.PDFPageHeight) + rctx.PDFPageSizeUnit
		pdfProps = SetPDFOption(pdfProps, "--page-height", pdfPageHeight)
	}

	err := tbl.PDFprintTable(w, pdfProps)
	if err != nil {
		s := fmt.Sprintf("Error in PDFprintTable: %s\n", err.Error())
		rlib.Ulog(s)
		fmt.Fprintf(w, "%s\n", s)
	}

	return err
}
