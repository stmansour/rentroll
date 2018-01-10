"use strict";

var GRID = "rentablesGrid";
var SIDEBAR_ID = "rentables";
var FORM = "rentableForm";
var common = require("../common.js");

// Below configurations are in use while performing tests via gridRecords.js
exports.gridConf = {
    grid: GRID,
    sidebarID: SIDEBAR_ID,
    capture: "rentablesGridRequest.png",
    endPoint: common.apiBaseURL + "/{0}/rentables/{1}",
    methodType: "POST",
    requestData: JSON.stringify({
        'cmd': 'get', 'selected': [], 'limit': 100, 'offset': 0
    }),
    excludeGridColumns: [],
    testCount: 59
};

// Below configurations are in use while performing tests via form.js
exports.formConf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    row: "0",
    capture: "rentablesFormRequest.png",
    captureAfterClosingForm: "rentablesFormRequestAfterClosingForm.png",
    buttonName: ["save", "saveadd", "delete"],
    testCount: 5
};

// Below configurations are in use while performing tests via addNew.js
exports.addNewConf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    capture: "rentablesAddNewButton.png",
    buttonName: ["save", "saveadd"],
    disableFields: ["RAID", "RARDtStart", "RARDtStop"],
    tabs: [],
    testCount: 31
};

