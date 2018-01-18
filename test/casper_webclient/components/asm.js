"use strict";

var GRID = "asmsGrid";
var SIDEBAR_ID = "asms";
var FORM = "asmEpochForm";
var common = require("../common.js");

// Below configurations are in use while performing tests via gridRecords.js
exports.gridConf = {
    grid: GRID,
    sidebarID: SIDEBAR_ID,
    capture: "asmsGridRequest.png",
    endPoint: common.apiBaseURL + "/{0}/asms/{1}",
    methodType: "POST",
    requestData: JSON.stringify({
        "cmd": "get",
        "selected": [],
        "limit": 100,
        "offset": 0,
        "searchDtStart": "3/5/2018",
        "searchDtStop": "4/5/2018"
    }),
    excludeGridColumns: [],
    testCount: 23
};

// Below configurations are in use while performing tests via form.js
exports.formConf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    row: "0",
    capture: "asmsFormRequest.png",
    captureAfterClosingForm: "asmsFormRequestAfterClosingForm.png",
    buttonName: ["save", "saveadd", "reverse"],
    testCount: 5
};

// Below configurations are in use while performing tests via addNew.js
exports.addNewConf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    capture: "asmsAddNewButton.png",
    buttonName: ["save", "saveadd"],
    disableFields: ["RAID", "Rentable"],
    tabs: [],
    testCount: 31
};
