"use strict";

const GRID = "tlsGrid";
const SIDEBAR_ID = "tls";
const FORM = "tlsInfoForm";
const MODULE = "tls";

// Below configurations are in use while performing tests via roller_spec.js for AIR Roller application
// For Module: Task List Definitions
export let conf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    module: MODULE,
    endPoint: "/{0}/tls/{1}",
    methodType: 'POST',
    excludeGridColumns: [],
    skipColumns: ["DtPreDone","DtDone"],
    skipFields: [],
    primaryId: "TLID",
    haveDateValue: true,
    fromDate: new Date(2018, 5, 1), // year, month-1, day : 1st June 2018
    toDate: new Date(2018, 5, 30), // 30th June 2018
    gridInForm: "tlsDetailGrid",
    formInPopUp: "newTaskListForm",
    notVisibleButtonNamesInForm: [],
    buttonNamesInDetailForm: ["save", "delete"]
};
