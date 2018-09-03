"use strict";

const GRID = "receiptsGrid";
const SIDEBAR_ID = "receipts";
const FORM = "receiptForm";
const MODULE = "receipt";

// Below configurations are in use while performing tests via receipt_2_spec.js for AIR Receipt application
export let conf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    module: MODULE,
    endPoint: "/{0}/receipts/{1}",
    methodType: 'POST',
    excludeGridColumns: [],
    buttonNamesInForm: ["save", "saveadd"],
    notVisibleButtonNamesInForm: ["close"],
    buttonNamesInDetailForm: ["save", "saveadd", "reverse"],
    skipColumns: ["reversed"],
    skipFields: ["Payor"], // TODO(Jay): Payor field is of enum type, hence it is skipped for now.
    primaryId: "RCPTID",
    haveDateValue: true,
    fromDate: new Date(2018, 0, 1), // year, month-1, day : 1st Jan 2018
    toDate: new Date(2018, 1, 1) // 1st Feb 2018
};

// TODO(Akshay): Green allocated section
