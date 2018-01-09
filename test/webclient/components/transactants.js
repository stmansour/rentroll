"use strict";

var GRID = "transactantsGrid";
var SIDEBAR_ID = "transactants";
var FORM = "transactantForm";
var common = require("../common.js");

// Below configurations are in use while performing tests via gridRecords.js
exports.gridConf = {
    grid: GRID,
    sidebarID: SIDEBAR_ID,
    capture: "transactantsGridRequest.png"
};

// Below configurations are in use while performing tests via form.js
exports.formConf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    row: "0",
    capture: "transactantsFormRequest.png",
    captureAfterClosingForm: "transactantsFormRequestAfterClosingForm.png",
    buttonName: ["save", "saveadd", "delete"],
    testCount: 5
};

// Below configurations are in use while performing tests via addNew.js
exports.addNewConf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    capture: "transactantsAddNewButton.png",
    buttonName: ["save", "saveadd"],
    tabs: ["Transactant", "User", "Payor", "Prospect"], // Must provide in order
    disableFields: [],
    testCount: 83
};