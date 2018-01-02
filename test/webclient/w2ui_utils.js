"use strict";

/*======================================
    THIS MODULE GIVE LOT OF UTILITIES REGARDING W2UI
    such as which are ids w2ui assigns to grid, grid records, form, form record,
    and you can get easily. Analysed from w2ui source code.
======================================*/

var common = require("./common.js");

// getGridRecordsDivID
exports.getGridRecordsDivID = function (gridName) {
    return "grid_" + gridName + "_records";
};

// getSidebarID Return id of sidebar node
exports.getSidebarID = function (id) {
    return "node_" + id;
};

// getGridRecordID Return record id of record in the w2ui grid
exports.getGridRecordID = function (id, gridName) {
    return "grid_" + gridName + "_rec_" + id;
};

// getGridTableRowsLength
exports.getGridTableRowsLength = function (gridRecordsContainerID) {
    // find rows (tr) which has `recid` attribute (which indicates that grid row record loaded in this tr)
    return $("#" + gridRecordsContainerID).find("tr[recid]").length;
};

// getGridRecordsLength Return total number of record in w2ui grid
exports.getGridRecordsLength = function (grid) {
    return w2ui[grid].records.length;
};

// getRightPanelID ID of the right hand side panel in application
exports.getRightPanelID = function () {
    return "layout_toplayout_panel_right";
};

// getW2UIButtonReferanceSelector Query selector for the button in the form
exports.getW2UIButtonReferanceSelector = function (btnName) {
    return 'button[class=w2ui-btn][name=' + btnName + ']';
};

// getCloseButtonSelector Query selector for close button of the form
exports.getCloseButtonSelector = function () {
    return '[class="fa fa-times"]';
};

// getInputFieldSelector Query selector for input box in w2ui form field
exports.getInputFieldSelector = function (inputFieldID) {
    // return 'input[class="w2ui-input"][id="{0}"]'.format(inputFieldID);
    return '#{0}'.format(inputFieldID);
};

// getBUDSelector Query selector for input Business Unit Description list
exports.getBUDSelector = function () {
    return 'input#BUD.w2ui-input.w2ui-select.w2field';
};

// getInputSelectFieldSelector Query selector for drop down in w2ui form field
exports.getInputSelectFieldSelector = function (inputSelectFieldID) {
    return 'input#{0}.w2ui-input.w2ui-select.w2field'.format(inputSelectFieldID);
};

// getCheckBoxSelector Query selector for checkbox in w2ui form field
exports.getCheckBoxSelector = function (checkboxID) {
    return 'input#{0}.w2ui-input'.format(checkboxID);
};

// getCheckBoxSelector Query selector for checkbox in w2ui form field
exports.getDisableFieldSelector = function (disableField) {
    return 'input#{0}'.format(disableField);
};

// getDateFieldSelector Query selector for date field in w2ui form field
exports.getDateFieldSelector = function (dateFieldId) {
    return 'input#{0}.w2ui-input'.format(dateFieldId);
};

// getRowColumnDataSelector Query selector for cell in the grid
exports.getRowColumnDataSelector = function (gridName, rowNo, columnNo) {
    return "#grid_{0}_data_{1}_{2}".format(gridName, rowNo, columnNo);
};

// getVisibleColumnName
exports.getVisibleColumnName = function (column) {

    if (!column.hidden) {
        // check column is in excludeGridColumns
        // 'this' represents callback argument. That is here excludeGridColumns
        if(!(column.field in this)){
            return column;
        }
    }
};

// getVisibleExcludedColumnName
exports.getVisibleExcludedColumnName = function (column) {

    if (!column.hidden) {
        // check column is in excludeGridColumns
        // 'this' represents callback argument. That is here excludeGridColumns
        if(column.field in this){
            return column;
        }
    }
};

// getCheckBoxW2UIFields Return a w2ui form field  object with type 'checkbox'
exports.getCheckBoxW2UIFields = function (inputField) {
    if (inputField.type === "checkbox") {
        return inputField;
    }
};

// getCheckBoxW2UIFields Return a w2ui form field  object with type 'checkbox'
exports.getDateW2UIFields = function (inputField) {
    if (inputField.type === "date") {
        return inputField;
    }
};

// getInputListW2UIFields Return a w2ui form field object with type 'list'
exports.getInputListW2UIFields = function (inputField) {
    if (inputField.type === "list") {
        return inputField;
    }
};

// getTextTypeW2UIFields Return a w2ui form fields  object with type 'text' or 'enum' or 'money'
exports.getTextTypeW2UIFields = function (inputField) {
    if (inputField.type === "text" || inputField.type === "enum" || inputField.type === "money") {
        return inputField;
    }
};

//
exports.getRecordsParentDivSelector = function (gridName) {
    return "#grid_{0}_records".format(gridName);
};