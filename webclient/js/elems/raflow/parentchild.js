/* global
    RACompConfig, sliderContentDivLength, reassignGridRecids,
    hideSliderContent, appendNewSlider, showSliderContentW2UIComp,
    loadTargetSection, requiredFieldsFulFilled, getRAFlowPartTypeIndex, initRAFlowAjax,
    getRAFlowAllParts, saveActiveCompData, toggleHaveCheckBoxDisablity, getRAFlowCompData,
    getChildRentableLocalData, getParentRentableLocalData
*/

"use strict";

//-----------------------------------------------------------------------
// loadRAPeopleChildSection -
//-----------------------------------------------------------------------
window.loadRAPeopleChildSection = function () {

    if (!("RAParentChildGrid" in w2ui)) {

        // rentables grid
        $().w2grid({
            name: 'RAParentChildGrid',
            header: 'Parent Child Rentables Relation',
            show: {
                toolbar:    false,
                footer:     true,
            },
            multiSelect: false,
            style: 'display: block;',
            columns: [
                {
                    field: 'recid',
                    hidden: true,
                },
                {
                    field: 'BID',
                    hidden: true
                },
                {
                    field: 'CRID',
                    hidden: true
                },
                {
                    field: 'PRID',
                    hidden: true
                },
                {
                    field: 'ChildRentableName',
                    caption: 'Rentable',
                    size: '40%',
                },
                {
                    field: 'ParentRentableName',
                    caption: 'Assign To',
                    size: '60%',
                    editable: {
                        type: 'select',
                        items: [],
                    },
                    render: function (record, index, col_index) {
                        var html = '';
                        var items = parentRentableW2UIItems;
                        for (var s in items) {
                            if (items[s].id == this.getCellValue(index, col_index)) html = items[s].text;
                        }
                        return html;
                    }
                },
            ],
        });
    }

    // prepare parent and child rentable list based on rentables section data
    var rentableCompData = getRAFlowCompData("rentables", app.raflow.activeFlowID) || [],
        compData = getRAFlowCompData("parentchild", app.raflow.activeFlowID) || [],
        tempCompData = [],
        recidCounter = 0,
        BID = getCurrentBID(),
        gridRecords = [],
        parentRentableW2UIItems = [{id: 0, text: " -- select parent rentables -- "}];

    // always render data from latest modified rentable comp data
    rentableCompData.forEach(function(rentableItem) {
        if ( (rentableItem.RTFLAGS & (1 << app.rtFLAGS.IsChildRentable)) != 0) { // 1 means this is child rentable

            // check that this record does exist in loca comp data of parentchild
            var RID = rentableItem.RID;
            var PRID = 0;
            var cRentable = getChildRentableLocalData(RID);

            // parent Rentable ID
            if (cRentable.PRID) {
                PRID = cRentable.PRID;
            }

            // prepare record struct for grid records list
            var rec = {
                recid:                  recidCounter,
                BID:                    BID,
                CRID:                   RID,
                PRID:                   PRID,
                ChildRentableName:      rentableItem.RentableName,
                ParentRentableName:     PRID, // grid's render will take care
            };
            recidCounter++;
            gridRecords.push(rec);

            // push the data temporarily
            tempCompData.push({BID: rec.BID, CRID: rec.CRID, PRID: rec.PRID});
        } else { // 0 = parent rentable
            parentRentableW2UIItems.push({id: rentableItem.RID, text: rentableItem.RentableName});
        }
    });

    console.debug(tempCompData);
    console.debug(compData);

    // now load grid in division
    $('#ra-form #parentchild .grid-container').w2render(w2ui.RAParentChildGrid);

    // load the existing data in rentables component
    setTimeout(function () {
        var compData = getRAFlowCompData("parentchild", app.raflow.activeFlowID);
        var grid = w2ui.RAParentChildGrid;
        // assign item prepared earlier for parent rentable list
        grid.getColumn("ParentRentableName").editable.items = parentRentableW2UIItems;
        grid.records = gridRecords;
        grid.refresh();

        if (compData) {
            // grid.records = compData;
            reassignGridRecids(grid.name);
        } else {
            grid.clear();
        }
    }, 500);
};

//-----------------------------------------------------------------------------
// getChildRentableLocalData -  returns the clone of child rentable data
//                              for requested RID by matching CRID
//-----------------------------------------------------------------------------
window.getChildRentableLocalData = function(RID, returnIndex) {
    var cloneData = {};
    var foundIndex = -1;
    var compData = getRAFlowCompData("parentchild", app.raflow.activeFlowID);
    compData.forEach(function(item, index) {
        if (item.CRID == RID) {
            if (returnIndex) {
                foundIndex = index;
            } else {
                cloneData = $.extend(true, {}, item);
            }
            return false;
        }
    });
    if (returnIndex) {
        return foundIndex;
    }
    return cloneData;
};

//-----------------------------------------------------------------------------
// getParentRentableLocalData - returns the clone of parent rentable data
//                              for requested RID by matching PRID
//-----------------------------------------------------------------------------
window.getParentRentableLocalData = function(RID, returnIndex) {
    var cloneData = {};
    var foundIndex = -1;
    var compData = getRAFlowCompData("parentchild", app.raflow.activeFlowID);
    compData.forEach(function(item, index) {
        if (item.PRID == RID) {
            if (returnIndex) {
                foundIndex = index;
            } else {
                cloneData = $.extend(true, {}, item);
            }
            return false;
        }
    });
    if (returnIndex) {
        return foundIndex;
    }
    return cloneData;
};
