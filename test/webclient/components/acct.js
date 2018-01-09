"use strict";

var GRID = "accountsGrid";
var SIDEBAR_ID = "accounts";
var FORM = "accountForm";
var common = require("../common.js");

// Below configurations are in use while performing tests via gridRecords.js
exports.gridConf = {
    grid: GRID,
    sidebarID: SIDEBAR_ID,
    capture: "accountsGridRequest.png"
};

// Below configurations are in use while performing tests via form.js
exports.formConf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    row: "0",
    capture: "accountFormRequest.png",
    captureAfterClosingForm: "accountFormRequestAfterClosingForm.png",
    buttonName: ["save", "saveadd", "delete"],
    testCount: 5
};

// Below configurations are in use while performing tests via addNew.js
exports.addNewConf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    capture: "accountAddNewButton.png",
    buttonName: ["save", "saveadd"],
    disableFields: [],
    tabs: [],
    testCount: 12
};
