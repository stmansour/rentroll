"use strict";

var GRID = "pmtsGrid";
var SIDEBAR_ID = "pmts";
var FORM = "pmtForm";

exports.gridConf = {
    grid: "pmtsGrid",
    sidebarID: "pmts",
    capture: "pmtsGridRequest.png"
};

exports.formConf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    row: "0",
    capture: "pmtsFormRequest.png",
    captureAfterClosingForm: "pmtsFormRequestAfterClosingForm.png",
    buttonName: ["save", "saveadd", "delete"]
};

exports.addNewConf = {
  grid: GRID,
  form: FORM,
  sidebarID: SIDEBAR_ID,
  capture: "pmtsAddNewButton.png",
  inputField: ["Name", "Description"],
  buttonName: ["save", "saveadd"],
  inputSelectField: []
};



