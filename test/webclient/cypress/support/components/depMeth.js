"use strict";

const GRID = "depmethGrid";
const SIDEBAR_ID = "depmeth";
const FORM = "depmethForm";
const MODULE = "depmeth";

// Below configurations are in use while performing tests via roller_spec.js for AIR Roller application
// For Module: Deposit Methods
export let conf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    module: MODULE,
    endPoint: "/{0}/depmeth/{1}",
    methodType: 'POST',
    excludeGridColumns: [],
    buttonNamesInForm: ["save", "saveadd"],
    notVisibleButtonNamesInForm: ["close"],
    buttonNamesInDetailForm: ["save", "saveadd", "delete"],
    skipColumns: [],
    skipFields: [],
    primaryId: "DPMID",
    fixtureFile: "depositMethods.json"
};
