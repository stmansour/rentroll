"use strict";

// return selector for close button of form
export function getFormCloseButtonSelector() {
    return '[class="fas fa-times"]';
}

// return selector for print button of form
export function getFormPrintButtonSelector() {
    return '[class="fas fa-print"]';
}

// return selector for print receipt popup
export function getPrintReceiptPopUpSelector() {
    return '#w2ui-popup';
}

// return selector for permanent resident radio button in print receipt popup
export function getPermanentResidentRadioButtonSelector() {
    return '[name="report_type"][type="radio"][value="permanent_resident"]';
}

// return selector for hotel radio button in print receipt popup
export function getHotelRadioButtonSelector() {
 return '[name="report_type"][type="radio"][value="hotel"]';
}

// return selector for close button in w2ui popup
export function getClosePopupButtonSelector() {
    return '[class="w2ui-popup-button w2ui-popup-close"]';
}

// return selector for print receipt pop up title
export function getPrintReceiptPopUpTitleSelector() {
    return '[class="w2ui-popup-title"]';
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

// return row selector
export function getGridRecordsSelector(gridName, rowNo){
    // return '#grid_' + gridName + '_rec_' + rowNo;
    return '#grid_' + gridName + '_records';
}

// return selector for first record in grid
export function getFirstRecordInGridSelector(gridName) {
    return '#grid_' + gridName + '_rec_0';
}

// return selector for second record in grid
export function getSecondRecordInGridSelector(gridName) {
    return '#grid_' + gridName + '_rec_1';
}

// return selector for formName
export function getFormSelector(formName) {
    return 'div[name=' + formName + ']';
}

// return selector for field
export function getFieldSelector(fieldId) {
    return '#' + fieldId;
}

// return selector for buttons
export function getButtonSelector(buttonName) {
    return 'button[name=' + buttonName + ']';
}

// return selector for Business Unit
export function getBUDSelector() {
    return '#BUD';
}

// return selector for Unallocated section in detail form
export function getAllocatedSectionSelector() {
    return '#FLAGReport';
}

// return selector for export csv button in grid toolbar
export function getExportCSVButtonSelector(gridname) {
    return '#tb_' + gridname + '_toolbar_item_csvexport';
}

// return selector for export pdf button in grid toolbar
export function getExportPDFButtonSelector(gridname) {
    return '#tb_'+ gridname + '_toolbar_item_printreport';
}

// return selector for dropdown value field
export function getDropDownValueFieldSelector(fieldValue) {
    return '#w2ui-overlay tr[index=' + fieldValue + ']';
}