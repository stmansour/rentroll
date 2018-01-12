"use strict";

var GRID = "expenseGrid";
var SIDEBAR_ID = "expense";
var FORM = "expenseForm";
var common = require("../common.js");

exports.addNewConf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    capture: "expensesAddNewButton.png",
    buttonName: ["save", "saveadd"],
    disableFields: [],
    tabs: [],
    testCount: 13
};

exports.gridConf = {
    grid: GRID,
    sidebarID: SIDEBAR_ID,
    capture: "expenseGridRequest.png",
    endPoint: common.apiBaseURL + "/{0}/expense/{1}",
    methodType: "POST",
    requestData: JSON.stringify({
        "cmd": "get",
        "selected": [],
        "limit": 100,
        "offset": 0,
        "searchDtStart": "10/1/2017",
        "searchDtStop": "11/1/2017"
    }),
    excludeGridColumns: [],
    testCount: 23
};