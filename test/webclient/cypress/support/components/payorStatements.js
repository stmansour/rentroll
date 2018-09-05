"use strict";

const GRID = "payorstmtGrid";
const SIDEBAR_ID = "payorstmt";
const MODULE = "payorstmtinfo";
const FORM = "payorStmtInfoForm";

// Below configurations are in use while performing tests via roller_spec.js for AIR Roller application
// For Module: Payor Statements
export let conf = {
    grid: GRID,
    sidebarID: SIDEBAR_ID,
    module: MODULE,
    form: FORM,
    endPoint: "/{0}/payorstmt/{1}",
    methodType: 'POST',
    excludeGridColumns: [],
    skipColumns: [],
    skipFields: [],
    primaryId: "TCID",
    haveDateValue: true,
    fromDate: new Date(2018, 7, 1), // year, month-1, day : 1st August 2018
    toDate: new Date(2018, 8, 1), // 1st September 2018
    gridInForm: "payorStmtDetailGrid"
};
