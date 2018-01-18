"use strict";

// Below configurations are in use while performing tests via gridRecords.js
exports.gridConf = {
    grid: "allocfundsGrid",
    sidebarID: "allocfunds",
    capture: "allocfundsGridRequest.png"
};

// Below configurations are in use while performing tests via form.js
exports.formConf = {
    grid: "allocfundsGrid",
    form: "allocfundsForm",
    sidebarID: "allocfunds",
    row: "0",
    capture: "allocfundsFormRequest.png",
    captureAfterClosingForm: "allocfundsFormRequestAfterClosingForm.png",
    buttonName: ["save", "saveadd", "delete"],
    testCount: 5
};