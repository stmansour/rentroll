"use strict";

/*======================================
    THIS MODULE GIVE LOT OF UTILITIES REGARDING W2UI
    such as which are ids w2ui assigns to grid, grid records, form, form record,
    and you can get easily. Analysed from w2ui source code.
======================================*/

exports.getGridRecordsDivID = function (gridName) {
    return "grid_" + gridName + "_records";
};

exports.getSidebarID = function (id) {
    return "node_" + id;
};

exports.getGridRecordID = function (id, gridName) {
    return "grid_" + gridName + "_rec_" + id;
};

exports.getGridTableRowsLength = function (gridRecordsContainerID) {
    // find rows (tr) which has `recid` attribute (which indicates that grid row record loaded in this tr)
    return $("#" + gridRecordsContainerID).find("tr[recid]").length;
};

exports.getGridRecordsLength = function (grid) {
    return w2ui[grid].records.length;
};

exports.getRightPanelID = function () {
    return "layout_toplayout_panel_right";
};

exports.getW2UIButtonReferanceSelector = function (btnName) {
    return 'button[class=w2ui-btn][name=' + btnName + ']';
};

exports.getCloseButtonSelector = function () {
    return '[class="fa fa-times"]';
};

exports.getInputFieldSelector = function (inputFieldID) {
    return 'input[class="w2ui-input"][id="{0}"]'.format(inputFieldID);
};

exports.getBUDSelector = function () {
    return 'input#BUD.w2ui-input.w2ui-select.w2field';
};


exports.getInputSelectFieldSelector = function (inputSelectFieldID) {
    return 'input#{0}.w2ui-input.w2ui-select.w2field'.format(inputSelectFieldID);
};


