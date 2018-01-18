"use strict";

// Below configurations are in use while performing tests via gridRecords.js
exports.gridConf = {
    grid: "rentalagrsGrid",
    sidebarID: "rentalagrs",
    capture: "rentalagrsGridRequest.png"
};

// Below configurations are in use while performing tests via form.js
exports.formConf = {
    grid: "rentalagrsGrid",
    form: "rentalagrsForm",
    sidebarID: "rentalagrs",
    row: "0",
    capture: "rentalagrsFormRequest.png",
    captureAfterClosingForm: "rentalagrsFormRequestAfterClosingForm.png",
    buttonName: ["save", "saveadd", "delete"],
    testCount: 5
};