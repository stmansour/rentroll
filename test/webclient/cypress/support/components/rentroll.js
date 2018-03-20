"use strict";

const GRID = "rrGrid";
const SIDEBAR_ID = "rr";
const MODULE = "rr";

// Below configurations are in use while performing tests via roller_spec.js for AIR Roller application
// For Module: Rent Roll
export let conf = {
    grid: GRID,
    sidebarID: SIDEBAR_ID,
    module: MODULE,
    endPoint: "/{0}/rr/{1}",
    methodType: 'POST',
    excludeGridColumns: [],
    skipColumns: [],
    skipFields: [],
    haveDateValue: true,
    fromDate: new Date(2018, 2, 1), // year, month-1, day : 1st March 2018
    toDate: new Date(2018, 3, 1) // 1st April 2018
};
