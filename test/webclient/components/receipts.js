"use strict";

var GRID = "receiptsGrid";
var SIDEBAR_ID = "receipts";
var FORM = "receiptForm";
var common = require("../common.js");

exports.gridConf = {
    grid: GRID,
    sidebarID: SIDEBAR_ID,
    capture: "receiptsGridRequest.png"
};

exports.addNewConf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    capture: "receiptsAddNewButton.png",
    buttonName: ["save", "saveadd"],
    disableFields: [],
    testCount: 25
};

