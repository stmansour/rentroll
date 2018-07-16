package ws

import (
	"context"
	"fmt"
)

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

// SaveRAFlowData saves data of raflow from client requested data
func SaveRAFlowData(ctx context.Context, sf SaveFlowRequest, flow *rlib.Flow) (err error) {
	var (
		jsBtData    []byte
		modMetaInfo []byte
	)

	// is valid part
	partType, ok := rlib.RAFlowPartsMap[sf.FlowPartKey]
	if !ok {
		err = fmt.Errorf("Unable to find part with key: %s for flowID: %d, flowType: %s, Error: %s",
			sf.FlowPartKey, sf.FlowID, sf.FlowType, err.Error())
	}

	modMetaInfo, jsBtData, err = rlib.GetRAFlowPartDataFromJSON(d.BID, sf.Data, int(partType), &flow)
	if err != nil {
		rlib.Console("Error while getting data from GetRAFlowPartDataFromJSON %s\n", err.Error())
		err = fmt.Errorf("Data is not in valid format for flowID: %d, flowType: %s, Error: %s",
			sf.FlowID, sf.FlowType, err.Error())
		return
	}

	// update data with given json data key
	err = rlib.UpdateFlowData(ctx, sf.FlowPartKey, jsBtData, &flow)
	if err != nil {
		return
	}

	// update data for modified meta data
	err = rlib.UpdateFlowData(ctx, "meta", modMetaInfo, &flow)
	if err != nil {
		return
	}
}
