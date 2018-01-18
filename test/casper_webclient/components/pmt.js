"use strict";

var GRID = "pmtsGrid";
var SIDEBAR_ID = "pmts";
var FORM = "pmtForm";
var common = require("../common.js");

// Below configurations are in use while performing tests via gridRecords.js
exports.gridConf = {
    grid: GRID,
    sidebarID: SIDEBAR_ID,
    tableName: "PaymentType",
    capture: "pmtsGridRequest.png",
    endPoint: common.apiBaseURL + "/{0}/pmts/{1}",
    methodType: "POST",
    requestData: JSON.stringify({
        'cmd': 'get', 'selected': [], 'limit': 100, 'offset': 0
    }),
    excludeGridColumns: [],
    testCount: 23
};

// Below configurations are in use while performing tests via form.js
exports.formConf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    row: "0",
    capture: "pmtsFormRequest.png",
    captureAfterClosingForm: "pmtsFormRequestAfterClosingForm.png",
    buttonName: ["save", "saveadd", "delete"],
    testCount: 5
};

// Below configurations are in use while performing tests via addNew.js
exports.addNewConf = {
  grid: GRID,
  form: FORM,
  sidebarID: SIDEBAR_ID,
  capture: "pmtsAddNewButton.png",
  buttonName: ["save", "saveadd"],
  disableFields: [],
  tabs: [],
  testCount: 13
};



