"use strict";

const GRID = "allocfundsGrid";
const SIDEBAR_ID = "allocfunds";
const FORM = "allocfundsForm";
const MODULE = "allocfunds";

// Below configurations are in use while performing tests via roller_spec.js for AIR Roller application
// For Module: Apply Receipts
export let conf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    module: MODULE,
    endPoint: "/{0}/allocfunds/{1}",
    methodType: 'POST',
    excludeGridColumns: [],
    buttonNamesInForm: [],
    notVisibleButtonNamesInForm: ["close"],
    buttonNamesInDetailForm: ["save"],
    skipColumns: [],
    skipFields: [],
    primaryId: "TCID",
    haveDateValue: false,
    gridInForm: 'unpaidasms'
};

