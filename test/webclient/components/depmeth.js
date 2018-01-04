"use strict";

var GRID = "depmethGrid";
var SIDEBAR_ID = "depmeth";
var FORM = "depmethForm";
var common = require("../common.js");

exports.gridConf = {
    grid: GRID,
    sidebarID: SIDEBAR_ID,
    capture: "depmethGridRequest.png",
    endPoint: common.apiBaseURL + "/{0}/depmeth/{1}",
    methodType: "POST",
    requestData: JSON.stringify({
        'cmd': 'get', 'selected': [], 'limit': 100, 'offset': 0
    }),
    excludeGridColumns: [],
    testCount: 11
};

exports.formConf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    row: "0",
    capture: "depmethFormRequest.png",
    captureAfterClosingForm: "depmethFormRequestAfterClosingForm.png",
    buttonName: ["save", "saveadd", "delete"],
    testCount: 5
};

exports.addNewConf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    capture: "depmethAddNewButton.png",
    buttonName: ["save", "saveadd"],
    disableFields: [],
    tabs: [],
    testCount: 11
};
