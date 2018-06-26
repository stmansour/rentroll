"use strict";

const GRID = "pmtsGrid";
const SIDEBAR_ID = "pmts";
const FORM = "pmtForm";
const MODULE = "pmts";

// Below configurations are in use while performing tests via roller_spec.js for AIR Roller application
// For Module: Payment Types
export let conf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    module: MODULE,
    endPoint: "/{0}/pmts/{1}",
    methodType: 'POST',
    excludeGridColumns: [],
    buttonNamesInForm: ["save", "saveadd"],
    notVisibleButtonNamesInForm: ["close"],
    buttonNamesInDetailForm: ["save", "saveadd", "delete"],
    skipColumns: [],
    skipFields: [],
    primaryId: "PMTID",
    fixtureFile: "paymentTypes.json"
};
