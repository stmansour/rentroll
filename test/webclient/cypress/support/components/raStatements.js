"use strict";

const GRID = "stmtGrid";
const SIDEBAR_ID = "stmt";
const MODULE = "stmtinfo";
const FORM = "stmtDetailForm";

// Below configurations are in use while performing tests via roller_spec.js for AIR Roller application
// For Module: Deposit accounts
export let conf = {
    grid: GRID,
    sidebarID: SIDEBAR_ID,
    module: MODULE,
    form: FORM,
    endPoint: "/{0}/stmt/{1}",
    methodType: 'POST',
    excludeGridColumns: [],
    skipColumns: [],
    skipFields: [],
    primaryId: "RAID",
    haveDateValue: true,
    fromDate: new Date(2018, 2, 1), // year, month-1, day : 1st March 2018
    toDate: new Date(2018, 3, 1), // 1st April 2018
    gridInForm: "stmtDetailGrid"
};
