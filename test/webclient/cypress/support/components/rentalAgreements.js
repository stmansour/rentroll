"use strict";

const GRID = "raflowsGrid";
const SIDEBAR_ID = "raflows";
const FORM = "tlsInfoForm";
const MODULE = "raflows";

// Below configurations are in use while performing tests via roller_spec.js for AIR Roller application
// For Module: Task List Definitions
export let conf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    module: MODULE,
    endPoint: "/{0}/flow/{1}",
    methodType: 'POST',
    excludeGridColumns: [],
    skipColumns: [],
    skipFields: [],
    primaryId: "RAID",
    haveDateValue: true,
    fromDate: new Date(2018, 7, 1), // year, month-1, day : 1st August 2018
    toDate: new Date(2018, 7, 31), // 31st August 2018
    gridInForm: "tlsDetailGrid",
    formInPopUp: "newTaskListForm",
    notVisibleButtonNamesInForm: [],
    buttonNamesInForm: ["save"],
    buttonNamesInDetailForm: ["save", "delete"]
};
