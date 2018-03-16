"use strict";

const GRID = "asmsGrid";
const SIDEBAR_ID = "asms";
const FORM = "asmEpochForm";
const MODULE = "asm";

// Below configurations are in use while performing tests via roller_spec.js for AIR Roller application
// For Module: Assess Charges
export let conf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    module: MODULE,
    endPoint: "/{0}/asms/{1}",
    methodType: 'POST',
    excludeGridColumns: [],
    buttonNamesInForm: ["save", "saveadd"],
    notVisibleButtonNamesInForm: ["close"],
    buttonNamesInDetailForm: ["save", "saveadd", "reverse"],
    skipColumns: ["epoch", "reversed", "Invoice"],
    skipFields: [],
    primaryId: "ASMID",
    haveDateValue: true,
    fromDate: new Date(2018, 1, 1), // year, month-1, day : 1st Feb 2018
    toDate: new Date(2018, 2, 1) // 1st March 2018
};

//TODO(Akshay): Find button in form

