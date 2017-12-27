"use strict";

var GRID = "arsGrid";
var SIDEBAR_ID = "ars";
var FORM = "arsForm";
var common = require("../common.js");

exports.gridConf = {
    grid: GRID,
    sidebarID: SIDEBAR_ID,
    tableName: "AR",
    capture: "arsGridRequest.png",
    endPoint: common.apiBaseURL + "/"+ common.apiVersion + "/ars/" + common.BID,
    methodType: "POST",
    requestData: JSON.stringify({
        'cmd': 'get', 'selected': [], 'limit': 100, 'offset': 0
    }),
    excludeGridColumns: {ARType: "ARTypes"},
    testCount: 40
};

exports.formConf = {
    grid: "arsGrid",
    form: "arsForm",
    sidebarID: "ars",
    row: "0",
    capture: "arsFormRequest.png",
    captureAfterClosingForm: "arsFormRequestAfterClosingForm.png",
    buttonName: ["save", "saveadd", "delete"],
    testCount: 5
};

exports.addNewConf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    capture: "arsAddNewButton.png",
    buttonName: ["save", "saveadd"],
    testCount: 24
};

