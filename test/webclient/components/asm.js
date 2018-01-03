"use strict";

var GRID = "asmsGrid";
var SIDEBAR_ID = "asms";
var FORM = "asmEpochForm";
var common = require("../common.js");

exports.gridConf = {
    grid: GRID,
    sidebarID: SIDEBAR_ID,
    capture: "asmsGridRequest.png"
};

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

exports.addNewConf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    capture: "asmsAddNewButton.png",
    buttonName: ["save", "saveadd"],
    disableFields: [],
    testCount: 27
};
