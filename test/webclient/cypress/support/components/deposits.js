"use strict";

const GRID = "depositGrid";
const SIDEBAR_ID = "deposit";
const FORM = "depositForm";
const MODULE = "deposit";

// Below configurations are in use while performing tests via roller_spec.js for AIR Roller application
// For Module: Deposits
export let conf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    module: MODULE,
    endPoint: "/{0}/deposit/{1}",
    methodType: 'POST',
    excludeGridColumns: [],
    buttonNamesInForm: ["save", "saveadd"],
    notVisibleButtonNamesInForm: ["close"],
    buttonNamesInDetailForm: ["save", "saveadd", "delete"],
    skipColumns: ["Check", "reversed"],
    skipFields: [],
    primaryId: "DID",
    haveDateValue: true,
    fromDate: new Date(2018, 2, 1), // year, month-1, day : 1st March 2018
    toDate: new Date(2018, 3, 1), // 1st April 2018
    gridInForm: 'depositlist'
};
