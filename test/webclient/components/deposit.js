"use strict";

var GRID = "depositGrid";
var SIDEBAR_ID = "deposit";
var FORM = "depositForm";
var common = require("../common.js");

// Below configurations are in use while performing tests via gridRecords.js
exports.gridConf = {
    grid: GRID,
    sidebarID: SIDEBAR_ID,
    capture: "depositGridRequest.png"
};

// Below configurations are in use while performing tests via form.js
exports.formConf = {
    grid: GRID,
    form: "depositForm",
    sidebarID: SIDEBAR_ID,
    row: "0",
    capture: "depositFormRequest.png",
    captureAfterClosingForm: "depositFormRequestAfterClosingForm.png",
    buttonName: ["save", "saveadd", "delete"],
    testCount: 5
};

// Below configurations are in use while performing tests via addNew.js
exports.addNewConf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    capture: "depositAddNewButton.png",
    buttonName: ["save", "saveadd"],
    disableFields: [],
    tabs: [],
    testCount: 13
};
