"use strict";

var GRID = "asmsGrid";
var SIDEBAR_ID = "asms";
var FORM = "asmEpochForm";
var common = require("../common.js");

// Below configurations are in use while performing tests via gridRecords.js
exports.gridConf = {
    grid: GRID,
    sidebarID: SIDEBAR_ID,
    capture: "asmsGridRequest.png"
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
    testCount: 33
};
