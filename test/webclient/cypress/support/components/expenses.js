"use strict";

const GRID = "expenseGrid";
const SIDEBAR_ID = "expense";
const FORM = "expenseForm";
const MODULE = "expense";

// Below configurations are in use while performing tests via roller_spec.js for AIR Roller application
// For Module: Expense
export let conf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    module: MODULE,
    endPoint: "/{0}/expense/{1}",
    methodType: 'POST',
    excludeGridColumns: [],
    buttonNamesInForm: ["save", "saveadd"],
    notVisibleButtonNamesInForm: ["close"],
    buttonNamesInDetailForm: ["save", "saveadd", "reverse"],
    skipColumns: ["Reversed"],
    skipFields: [],
    primaryId: "EXPID",
    haveDateValue: true,
    fromDate: new Date(2018, 7, 1), // year, month-1, day : 1st August 2018
    toDate: new Date(2018, 8, 1) // 1st September 2018
};

//TODO(Akshay): Find button in form
