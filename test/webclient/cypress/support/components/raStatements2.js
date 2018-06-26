"use strict";

const GRID = "stmtDetailGrid";
const FORM = "stmtDetailForm";

// Below configurations are in use while performing Record Detail Form tests via roller_spec.js for AIR Roller application
// For Module: RA Statements
export let conf = {
    grid: GRID,
    form: FORM,
    excludeGridColumns: [],
    skipColumns: ["Reverse","dummy"],
    skipFields: [],
    haveDateValue: false
};
