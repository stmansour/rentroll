"use strict";

var GRID = "receiptsGrid";
var SIDEBAR_ID = "receipts";
var FORM = "receiptForm";
var common = require("../common.js");

// Below configurations are in use while performing tests via gridRecords.js
exports.gridConf = {
    grid: GRID,
    sidebarID: SIDEBAR_ID,
    capture: "receiptsGridRequest.png",
    endPoint: common.apiBaseURL + "/{0}/receipts/{1}",
    methodType: "POST",
    requestData: JSON.stringify({
        "cmd": "get",
        "selected": [],
        "limit": 100,
        "offset": 0,
        "searchDtStart": "10/1/2017",
        "searchDtStop": "11/1/2017"
    }),
    excludeGridColumns: [],
    testCount: 23
};

// Below configurations are in use while performing tests via addNew.js
exports.addNewConf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    capture: "receiptsAddNewButton.png",
    buttonName: ["save", "saveadd"],
    disableFields: [],
    tabs: [],
    testCount: 29
};

