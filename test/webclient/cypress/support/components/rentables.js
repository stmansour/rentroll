"use strict";

const GRID = "rentablesGrid";
const SIDEBAR_ID = "rentables";
const FORM = "rentableForm";
const MODULE = "rentable";

// Below configurations are in use while performing tests via roller_spec.js for AIR Roller application
// For Module: Rentables
export let conf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    module: MODULE,
    endPoint: "/{0}/rentables/{1}",
    methodType: 'POST',
    excludeGridColumns: [],
    buttonNamesInForm: ["save", "saveadd"],
    notVisibleButtonNamesInForm: [],
    buttonNamesInDetailForm: ["save", "saveadd"],
    skipColumns: ["RTID"],
    skipFields: [],
    primaryId: "RID",
    haveDateValue: true,
    fromDate: new Date(2018, 2, 1), // year, month-1, day : 1st March 2018
    toDate: new Date(2018, 4, 1) // 1st April 2018
};


