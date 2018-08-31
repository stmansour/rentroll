"use strict";

const GRID = "RARentableFeesGrid";
const FORM = "RARentableFeeForm";

// Below configurations are in use while performing tests via roller_spec.js for AIR Roller application
// For Module: RARentableFee
export let conf = {
    grid: GRID,
    form: FORM,
    excludeGridColumns: [],
    buttonNamesInForm: ["save", "saveadd"],
    notVisibleButtonNamesInForm: [],
    buttonNamesInDetailForm: ["save", "saveadd", "delete"],
    skipColumns: ["haveError", "RemoveRec"],
    skipFields: [],
};


