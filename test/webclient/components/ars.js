"use strict";

var GRID = "arsGrid";
var SIDEBAR_ID = "ars";
var FORM = "arsForm";
var common = require("../common.js");

// Below configurations are in use while performing tests via gridRecords.js
exports.gridConf = {
    grid: GRID,
    sidebarID: SIDEBAR_ID,
    tableName: "AR",
    capture: "arsGridRequest.png",
    endPoint: common.apiBaseURL + "/{0}/ars/{1}",
    methodType: "POST",
    requestData: JSON.stringify({
        'cmd': 'get', 'selected': [], 'limit': 100, 'offset': 0
    }),
    excludeGridColumns: {ARType: "ARTypes"},
    testCount: 40
};

// Below configurations are in use while performing tests via form.js
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

// Below configurations are in use while performing tests via addNew.js
exports.addNewConf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    capture: "arsAddNewButton.png",
    buttonName: ["save", "saveadd"],
    disableFields: ["RAIDrqd"],
    tabs: [],
    testCount: 24
};

