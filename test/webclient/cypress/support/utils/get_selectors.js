"use strict";

// return selector for close button of form
export function getFormCloseButtonSelector() {
    return '[class="fa fa-times"]';
}

// return selector for node in left side panel
export function getNodeSelector(nodeName) {
    return '#node_' + nodeName;
}

// return selector for grid
export function getGridSelector(gridName) {
    return '#grid_' + gridName + '_records';
}

// return
export function getRowsInGridSelector(gridName) {
    return '#grid_' + gridName + '_records table tr[recid]';
}

// return cell selector
export function getCellSelector(gridName, rowNo, columnNo) {
    return '#grid_' + gridName + '_data_' + rowNo + '_' + columnNo;
}

// return selector for first record in grid
export function getFirstRecordInGridSelector(gridName) {
    return '#grid_' + gridName + '_rec_0';
}

// return selector for formName
export function getFormSelector(formName) {
    return 'div[name=' + formName + ']';
}

// return selector for field
export function getFieldSelector(fieldId) {
    return '#' + fieldId;
}

export function getButtonSelector(buttonName) {
    return 'button[name=' + buttonName + ']';
}