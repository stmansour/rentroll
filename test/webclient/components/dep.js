"use strict";

var GRID = "depGrid";
var FORM = "depositoryForm";
var SIDEBAR_ID = "dep";
var common = require("../common.js");

exports.gridConf = {
    grid: "depGrid",
    sidebarID: "dep",
    capture: "depGridRequest.png",
    endPoint: common.apiBaseURL + "/{0}/dep/{1}",
    methodType: "POST",
    requestData: JSON.stringify({
        'cmd': 'get', 'selected': [], 'limit': 100, 'offset': 0
    }),
    excludeGridColumns: [],
    testCount: 19
};

exports.formConf = {
    grid: "depGrid",
    form: "depForm",
    sidebarID: "dep",
    row: "0",
    capture: "depFormRequest.png",
    captureAfterClosingForm: "depFormRequestAfterClosingForm.png",
    buttonName: ["save", "saveadd", "delete"],
    testCount: 5
};

exports.addNewConf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    capture: "depAddNewButton.png",
    buttonName: ["save", "saveadd"],
    disableFields: [],
    testCount: 15
};

