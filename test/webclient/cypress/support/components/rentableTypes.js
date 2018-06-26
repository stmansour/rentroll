"use strict";

const GRID = "rtGrid";
const SIDEBAR_ID = "rt";
const FORM = "rtForm";
const MODULE = "rt";

// Below configurations are in use while performing tests via roller_spec.js for AIR Roller application
// For Module: Rentable Types
export let conf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    module: MODULE,
    endPoint: "/{0}/rt/{1}",
    methodType: 'POST',
    excludeGridColumns: [],
    buttonNamesInForm: ["save", "saveadd", "deactivate"],
    notVisibleButtonNamesInForm: ["reactivate"],
    buttonNamesInDetailForm: ["save", "saveadd", "deactivate"],
    skipColumns: ["Active"],
    skipFields: [],
    primaryId: "RTID",
    haveDateValue: true,
    fromDate: new Date(2018, 2, 1), // year, month-1, day : 1st March 2018
    toDate: new Date(2018, 2, 31), // 1st April 2018
    fixtureFile: "rentableTypes.json"
};
