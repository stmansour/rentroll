"use strict";

const GRID = "accountsGrid";
const SIDEBAR_ID = "accounts";
const FORM = "accountForm";
const MODULE = "account";

// Below configurations are in use while performing tests via roller_spec.js for AIR Roller application
// For Module: Charts of Account
export let conf = {
    grid: GRID,
    form: FORM,
    sidebarID: SIDEBAR_ID,
    module: MODULE,
    capture: "chartofAccountsGridRequest.png",
    endPoint: "/{0}/accounts/{1}",
    methodType: 'POST',
    requestData: JSON.stringify({"cmd": "get", "selected": [], "limit": 100, "offset": 0}),
    excludeGridColumns: [],
    buttonNamesInForm: ["save", "saveadd"],
    notVisibleButtonNamesInForm: ["close"],
    buttonNamesInDetailForm: ["save", "saveadd", "delete"],
    skipColumns: [],
    skipFields: [],
    primaryId: "LID"
};
