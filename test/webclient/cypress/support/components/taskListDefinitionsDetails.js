"use strict";

const GRID = "tldsDetailGrid";
const SIDEBAR_ID = "tlds";
const FORM = "taskDescForm";
const MODULE = "td";

// Below configurations are in use while performing Record Detail Form tests via roller_spec.js for AIR Roller application
// For Module: Task List Definitions
export let conf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    module: MODULE,
    endPoint: "/{0}/td/{1}",
    methodType: 'POST',
    excludeGridColumns: [],
    skipColumns: [],
    skipFields: [],
    primaryId: "TDID",
    haveDateValue: false,
    notVisibleButtonNamesInForm: [],
    buttonNamesInForm: ["save", "delete"],
    buttonNamesInDetailForm: ["save", "delete"]
};
