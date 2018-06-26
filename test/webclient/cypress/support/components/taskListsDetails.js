"use strict";

const GRID = "tlsDetailGrid";
const SIDEBAR_ID = "tls";
const FORM = "taskForm";
const MODULE = "task";

// Below configurations are in use while performing Record Detail Form tests via roller_spec.js for AIR Roller application
// For Module: Task List Definitions
export let conf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    module: MODULE,
    endPoint: "/{0}/task/{1}",
    methodType: 'POST',
    excludeGridColumns: [],
    skipColumns: ["DtPreDone","DtDone"],
    skipFields: [],
    primaryId: "TID",
    haveDateValue: false,
    formInPopUp: "newTaskListForm",
    notVisibleButtonNamesInForm: [],
    buttonNamesInForm: ["save"],
    buttonNamesInDetailForm: ["save", "delete"]
};
