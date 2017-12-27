"use strict";

var GRID = "rentablesGrid";
var SIDEBAR_ID = "rentables";
var FORM = "rentableForm";
var common = require("../common.js");

exports.gridConf = {
    grid: GRID,
    sidebarID: SIDEBAR_ID,
    capture: "rentablesGridRequest.png",
    endPoint: common.apiBaseURL + "/"+ common.apiVersion + "/rentables/" + common.BID,
    methodType: "POST",
    requestData: JSON.stringify({
        'cmd': 'get', 'selected': [], 'limit': 100, 'offset': 0
    }),
    excludeGridColumns: [],
    testCount: 58
};

exports.formConf = {
    grid: GRID,
    form: "rentableForm",
    sidebarID: SIDEBAR_ID,
    row: "0",
    capture: "rentablesFormRequest.png",
    captureAfterClosingForm: "rentablesFormRequestAfterClosingForm.png",
    buttonName: ["save", "saveadd", "delete"],
    testCount: 5
};

exports.addNewConf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    capture: "rentablesAddNewButton.png",
    buttonName: ["save", "saveadd"],
    disableFields: ["RAID", "RARDtStart", "RARDtStop"],
    testCount: 16
};

