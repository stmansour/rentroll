"use strict";

const GRID = "depGrid";
const SIDEBAR_ID = "dep";
const FORM = "depositoryForm";
const MODULE = "dep";

// Below configurations are in use while performing tests via roller_spec.js for AIR Roller application
// For Module: Deposit accounts
export let conf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    module: MODULE,
    endPoint: "/{0}/dep/{1}",
    methodType: 'POST',
    excludeGridColumns: [],
    buttonNamesInForm: ["save", "saveadd"],
    notVisibleButtonNamesInForm: ["close"],
    buttonNamesInDetailForm: ["save", "saveadd", "delete"],
    skipColumns: [],
    skipFields: [],
    primaryId: "DEPID",
    fixtureFile: "depositoryAccounts.json"
};
