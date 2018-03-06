"use strict";

const GRID = "receiptsGrid";
const SIDEBAR_ID = "receipts";
const FORM = "receiptForm";
const MODULE = "receipt";

// Below configurations are in use while performing tests via receipt_2_spec.js for AIR Receipt application
export let conf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    module: MODULE,
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
    skipFields: [],
    printReceiptButtons: ["print", "close"],
    primaryId: "RCPTID",
    haveDateValue: true,
    fromDate: new Date(2018, 0, 1), // year, month-1, day : 1st Jan 2018
    toDate: new Date(2018, 1, 1) // 1st Feb 2018
};
