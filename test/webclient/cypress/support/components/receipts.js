"use strict";

const GRID = "receiptsGrid";
const SIDEBAR_ID = "receipts";
const FORM = "receiptForm";

// Below configurations are in use while performing tests via roller_spec.js
export let conf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    capture: "receiptsGridRequest.png",
    endPoint: "/{0}/receipts/{1}",
    methodType: 'POST',
    requestData: JSON.stringify({
        "cmd": "get",
        "selected": [],
        "limit": 100,
        "offset": 0,
        "searchDtStart": "10/1/2017",
        "searchDtStop": "11/1/2017"
    }),
    excludeGridColumns: [],
    buttonNamesInForm: ["save", "saveprint"],
    notVisibleButtonNamesInForm: ["close"],
    buttonNamesInDetailForm: ["save", "saveprint", "reverse"],
    skipColumns: ["reversed"],
    skipFields: ["ERentableName"],
    primaryId: "RCPTID"
};
