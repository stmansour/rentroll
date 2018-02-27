"use strict";

const GRID = "pmtsGrid";
const SIDEBAR_ID = "pmts";
const FORM = "pmtForm";
const MODULE = "pmt";

// Below configurations are in use while performing tests via roller_spec.js for AIR Roller application
// For Module: Charts of Account
export let conf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    module: MODULE,
    capture: "paymentTypesGridRequest.png",
    endPoint: "/{0}/pmts/{1}",
    methodType: 'POST',
    requestData: JSON.stringify({"cmd": "get", "selected": [], "limit": 100, "offset": 0}),
    excludeGridColumns: [],
    buttonNamesInForm: ["save", "saveadd"],
    notVisibleButtonNamesInForm: ["close"],
    buttonNamesInDetailForm: ["save", "saveadd", "delete"],
    skipColumns: [],
    skipFields: [],
    primaryId: ""
};
