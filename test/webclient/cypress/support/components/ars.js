"use strict";

const GRID = "arsGrid";
const SIDEBAR_ID = "ars";
const FORM = "arsForm";
const MODULE = "ar";

// Below configurations are in use while performing tests via roller_spec.js for AIR Roller application
// For Module: Account Rules
export let conf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    module: MODULE,
    endPoint: "/{0}/ars/{1}",
    methodType: 'POST',
    excludeGridColumns: [],
    buttonNamesInForm: ["save", "saveadd"],
    notVisibleButtonNamesInForm: ["close"],
    buttonNamesInDetailForm: ["save", "saveadd", "delete"],
    skipColumns: [],
    skipFields: [],
    primaryId: "ARID",
    fixtureFile: "accountRules.json"
};
