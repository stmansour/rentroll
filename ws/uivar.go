package ws

import (
	"fmt"
	"io"
	"net/http"
	"rentroll/rlib"
)

func bidToBud(bid int64) string {
	var b rlib.Business
	rlib.GetBusiness(bid, &b)
	return b.Designation
}

// GetAssessmentList returns all assessments for the supplied business
func GetAssessmentList(bid int64) string {
	s := ""
	m := rlib.GetARsByType(bid, rlib.ARASSESSMENT)
	s += fmt.Sprintf("app.Assessments['%s']=[", bidToBud(bid))
	for i := 0; i < len(m); i++ {
		s += fmt.Sprintf("{id:%d,text:%q},", m[i].ARType, m[i].Name)
	}
	return s + "];\n"
}

// GetReceiptList returns all assessments for the supplied business
func GetReceiptList(bid int64) string {
	s := ""
	m := rlib.GetARsByType(bid, rlib.ARRECEIPT)
	s += fmt.Sprintf("app.Receipts['%s']=[", bidToBud(bid))
	for i := 0; i < len(m); i++ {
		s += fmt.Sprintf("{id:%d,text:%q},", m[i].ARType, m[i].Name)
	}
	return s + "];\n"
}

// SvcUIVal returns the requested variable in JSON form
//
// wsdoc {
//  @Title  Get UI Value
//	@URL /v1/uival/:BID/varname
//  @Method  GET
//	@Synopsis Return JSON representing the UI Value
//  @Desc Return data can be processed by eval() to create the string lists used in the UI.
//	@Input
//  @Response string
// wsdoc }
func SvcUIVal(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	funcname := "SvcUIVar"
	fmt.Printf("Entered %s\n", funcname)
	switch d.DetVal {
	case "app.Assessments":
		s := GetAssessmentList(d.BID)
		io.WriteString(w, s)
	case "app.Receipts":
		s := GetReceiptList(d.BID)
		io.WriteString(w, s)
	default:
		e := fmt.Errorf("Unknown variable requested: %s", d.DetVal)
		SvcGridErrorReturn(w, e)
		return
	}
}
