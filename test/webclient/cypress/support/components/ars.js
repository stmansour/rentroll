"use strict";

const GRID = "arsGrid";
const SIDEBAR_ID = "ars";
const FORM = "arsForm";
const MODULE = "ar";

// Below configurations are in use while performing tests via roller_spec.js for AIR Roller application
// For Module: Deposit accounts
export let conf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    module: MODULE,
    capture: "depositMethodsGridRequest.png",
    endPoint: "/{0}/ars/{1}",
    methodType: 'POST',
    requestData: JSON.stringify({"cmd": "get", "selected": [], "limit": 100, "offset": 0}),
    excludeGridColumns: [],
    buttonNamesInForm: ["save", "saveadd"],
    notVisibleButtonNamesInForm: ["close"],
    buttonNamesInDetailForm: ["save", "saveadd", "delete"],
    skipColumns: [],
    skipFields: ["ApplyRcvAccts", "RAIDrqd", "PriorToRAStart", "PriorToRAStop"], //TODO(Akshay): Write UI test for checkboxes
    primaryId: "ARID"
};
