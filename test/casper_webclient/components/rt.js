"use strict";

var GRID = "rtGrid";
var FORM = "rtForm";
var SIDEBAR_ID = "rt";
var common = require("../common.js");

// Below configurations are in use while performing tests via gridRecords.js
exports.gridConf = {
    grid: GRID,
    sidebarID: SIDEBAR_ID,
    capture: "rtGridRequest.png",
    endPoint: common.apiBaseURL + "/{0}/rt/{1}",
    methodType: "POST",
    requestData: JSON.stringify({
        "cmd": "get",
        "selected": [],
        "limit": 100,
        "offset": 0,
        "searchDtStart": "11/30/2017",
        "searchDtStop": "12/31/2017"
    }),
    excludeGridColumns: {
        Active: "rtActiveFLAGS", // TODO(Sudip): "rtActiveFLAGS" has been removed
        RentCycle: "cycleFreq",
        Proration: "cycleFreq",
        GSRPC: "cycleFreq",
        ManageToBudget: "manageToBudgetList" // TODO(Sudip): "manageToBBudgetList" has been removed
    },
    testCount: 91
};

// Below configurations are in use while performing tests via form.js
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

// Below configurations are in use while performing tests via addNew.js
exports.addNewConf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    capture: "rtAddNewButton.png",
    buttonName: ["save", "saveadd"],
    disableFields: [],
    tabs: [],
    testCount: 23
};