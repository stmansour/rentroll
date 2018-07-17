package ws

// GridRAFlowResponse is a struct to hold info for rental agreement for the grid response
type GridRAFlowResponse struct {
	Recid     int64 `json:"recid"`
	BID       int64
	BUD       string
	FlowID    int64
	UserRefNo string
}

/*// RAFlowDiffRequest struct
type RAFlowDiffRequest struct {
	RAID int64
}

// SvcRAFlowDiffHandler to tell diff in raflow data for existing RA
// wsdoc {
//  @Title raflow data difference
//	@URL /v1/raflow-diff/:BUI/:RAID
//	@Method POST
//	@Synopsis yes/no if there is any difference in raflow data
//  @Description This service tells whether if there is any difference between raflow json and permanent data
//  @Input RAFlowDiffRequest
//  @Response SvcStatusResponse
// wsdoc }
func SvcRAFlowDiffHandler(w http.ResponseWriter, r *http.Request, d *ServiceData) {
	const funcname = "SvcRAFlowDiffHandler"
	fmt.Printf("Entered in %s\n", funcname)

	var (
		err error
		foo RAFlowDiffRequest
		g   struct {
			Diff   bool
			Status string
		}
	)

	// http method check
	if r.Method != "POST" {
		err = fmt.Errorf("Only POST method is allowed")
		return
	}

	// unmarshal data into request data struct
	if err = json.Unmarshal([]byte(d.data), &foo); err != nil {
		return
	}

	// get diff for raflow
	var diff bool
	diff, err = RAFlowDataDiff(r.Context(), foo.RAID)
	if err != nil {
		SvcErrorReturn(w, err, funcname)
		return
	}

	// set the response
	g.Diff = diff
	g.Status = "success"
	SvcWriteResponse(d.BID, &g, w)
}
*/
