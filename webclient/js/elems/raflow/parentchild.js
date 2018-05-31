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
                        var html = '',
                            items = parentRentableW2UIItems;

                        for (var s in items) {
                            if (items[s].id == this.getCellValue(index, col_index)) html = items[s].text;
                        }
                        return html;
                    }
                },
            ],
            onChange: function(event) {
                event.onComplete = function() {
                    this.save();
                };
            }
        });
    }

    // prepare parent and child rentable list based on rentables section data
    var rentableCompData = getRAFlowCompData("rentables", app.raflow.activeFlowID) || [],
        compData = getRAFlowCompData("parentchild", app.raflow.activeFlowID) || [],
        recidCounter = 1, // always starts with 1
        BID = getCurrentBID(),
        gridRecords = [],
        parentRentableW2UIItems = [];

    // always render data from latest modified rentable comp data
    rentableCompData.forEach(function(rentableItem) {
        if ( (rentableItem.RTFLAGS & (1 << app.rtFLAGS.IsChildRentable)) != 0) { // 1 means this is child rentable

            // check that this record does exist in loca comp data of parentchild
            var RID = rentableItem.RID;
            var PRID = 0;
            var cRentable = getChildRentableLocalData(RID);

            // parent Rentable ID found then
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
        } else { // 0 = parent rentable
            parentRentableW2UIItems.push({id: rentableItem.RID, text: rentableItem.RentableName});
        }
    });
    console.debug(parentRentableW2UIItems);

    // if there is only one parent rentable then pre-select it for all child rentable
    // otherwise built drop down menu
    if (parentRentableW2UIItems.length == 0) {
        gridRecords.forEach(function(rec) {
            rec.PRID = 0;
            rec.ParentRentableName = 0;
        });
        parentRentableW2UIItems.unshift({id: 0, text: " -- select parent rentables -- "});
    } else if (parentRentableW2UIItems.length == 1) {
        // re-assign PRID
        gridRecords.forEach(function(rec) {
            rec.PRID = parentRentableW2UIItems[0].id;
            rec.ParentRentableName = parentRentableW2UIItems[0].id;
        });
    } else {
        parentRentableW2UIItems.unshift({id: 0, text: " -- select parent rentables -- "});
    }

    // now load grid in division
    $('#ra-form #parentchild .grid-container').w2render(w2ui.RAParentChildGrid);

    // load the existing data in rentables component
    setTimeout(function () {
        var compData = getRAFlowCompData("parentchild", app.raflow.activeFlowID);
        var grid = w2ui.RAParentChildGrid;
        grid.records = gridRecords;
        grid.refresh();

        // assign item prepared earlier for parent rentable list
        grid.getColumn("ParentRentableName").editable.items = parentRentableW2UIItems;
        grid.getColumn("ParentRentableName").render();

        // if there are any difference between server data and local data at this step
        // then save the modified data on the server via API

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
