"use strict";

/*======================================
    THIS MODULE GIVE LOT OF UTILITIES REGARDING W2UI
    such as which are ids w2ui assigns to grid, grid records, form, form record,
    and you can get easily. Analysed from w2ui source code.
======================================*/

exports.getGridRecordsDivID = function(gridName) {
    return "grid_" + gridName + "_records";
};

exports.getSidebarID = function(id) {
    return "node_" + id;
};

exports.getGridTableRowsLength = function(gridRecordsContainerID) {
    // find rows (tr) which has `recid` attribute (which indicates that grid row record loaded in this tr)
    return $("#"+gridRecordsContainerID).find("tr[recid]").length;
};

exports.getGridRecordsLength = function(grid) {
    return w2ui[grid].records.length;
};


