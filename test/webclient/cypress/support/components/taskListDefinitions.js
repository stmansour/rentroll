"use strict";

const GRID = "tldsGrid";
const SIDEBAR_ID = "tlds";
const FORM = "tldsInfoForm";
const MODULE = "tlds";

// Below configurations are in use while performing tests via roller_spec.js for AIR Roller application
// For Module: Task List Definitions
export let conf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    module: MODULE,
    endPoint: "/{0}/tlds/{1}",
    methodType: 'POST',
    excludeGridColumns: [],
    skipColumns: [],
    skipFields: [],
    primaryId: "TLDID",
    haveDateValue: true,
    fromDate: new Date(2018, 5, 1), // year, month-1, day : 1st June 2018
    toDate: new Date(2018, 5, 30), // 30th June 2018
    gridInForm: "tldsDetailGrid",
    notVisibleButtonNamesInForm: [],
    buttonNamesInDetailForm: ["save", "delete"]
};
