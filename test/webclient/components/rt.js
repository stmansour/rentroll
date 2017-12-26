"use strict";

var GRID = "rtGrid";
var FORM = "rtForm";
var SIDEBAR_ID = "rt";
var common = require("../common.js");

exports.gridConf = {
    grid: GRID,
    sidebarID: SIDEBAR_ID,
    capture: "rtGridRequest.png",
    endPoint: common.apiBaseURL + "/" + common.apiVersion + "/rt/" + common.BID,
    methodType: "POST",
    requestData: JSON.stringify({
        "cmd": "get",
        "selected": [],
        "limit": 100,
        "offset": 0,
        "searchDtStart": "11/30/2017",
        "searchDtStop": "12/31/2017"
    }),
    excludeGridColumns: ["Active", "RentCycle", "Proration", "GSRPC", "ManageToBudget"],
    testCount: 50
};

exports.formConf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    row: "0",
    capture: "rtFormRequest.png",
    captureAfterClosingForm: "rtFormRequestAfterClosingForm.png",
    buttonName: ["save", "saveadd", "delete"],
    testCount: 5
};