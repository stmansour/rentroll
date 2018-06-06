/* global
    RACompConfig, sliderContentDivLength, reassignGridRecids,
    hideSliderContent, appendNewSlider, showSliderContentW2UIComp,
    loadTargetSection, requiredFieldsFulFilled, getRAFlowPartTypeIndex, initRAFlowAjax,
    getRAFlowAllParts, saveActiveCompData, toggleHaveCheckBoxDisablity, getRAFlowCompData,
    getChildRentableLocalData, setChildRentableLocalData, saveParentChildCompData
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
                    render: function (record/*, index, col_index*/) {
                        var html = '';

                        if (record) {
                            var items = app.raflow.parentRentableW2UIItems;
                            for (var s in items) {
                                if (items[s].id == record.ParentRentableName) html = items[s].text;
                            }
                        }
                        return html;
                    }
                },
            ],
            onChange: function(event) {
                var grid = this;
                event.onComplete = function() {
                    // parent rentable name column index
                    var PRNCI = grid.getColumn("ParentRentableName", true);
                    if (PRNCI === event.column) {
                        var record = grid.get(event.recid);
                        var localData = getChildRentableLocalData(record.CRID);

                        // set PRID locally as well
                        localData.PRID = record.PRID = parseInt(event.value_new);
                        record.ParentRentableName = parseInt(event.value_new);

                        // set modified data in grid and locally
                        grid.set(event.recid, record);
                        setChildRentableLocalData(record.CRID, localData);
                    }

                    // save grid changes
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
        gridRecords = [];

    // always render data from latest modified rentable comp data
    rentableCompData.forEach(function(rentableItem) {
        var RID = rentableItem.RID,
            RentableName = rentableItem.RentableName;

        // 1 means this is child rentable
        if ( (rentableItem.RTFLAGS & (1 << app.rtFLAGS.IsChildRentable)) != 0) {
            var PRID = 0;
            var cRentable = getChildRentableLocalData(RID);

            // parent Rentable ID found then
            if (cRentable.PRID) {
                // if it's found in parent rentable list then keep as it is
                // else assign 0 if not found
                var PRIDFound = false;
                app.raflow.parentRentableW2UIItems.forEach(function(parentRItem) {
                    if (parentRItem.id == cRentable.PRID) {
                        PRIDFound = true;
                        return false;
                    }
                });

                // if parent RID found
                if (PRIDFound) {
                    PRID = cRentable.PRID;
                } else {
                    PRID = 0;
                }
            }

            // prepare record struct for grid records list
            var rec = {
                recid:                  recidCounter,
                BID:                    BID,
                CRID:                   RID,
                PRID:                   PRID,
                ChildRentableName:      RentableName,
                ParentRentableName:     PRID, // grid's render will take care
            };
            recidCounter++;
            gridRecords.push(rec);
        }
    });

    // if there is only one parent rentable then pre-select it for all child rentable
    // otherwise built drop down menu
    if (app.raflow.parentRentableW2UIItems.length == 0) {
        gridRecords.forEach(function(rec) {
            rec.PRID = 0;
            rec.ParentRentableName = 0;
        });
    } else if (app.raflow.parentRentableW2UIItems.length == 1) {
        // re-assign PRID
        gridRecords.forEach(function(rec) {
            rec.PRID = app.raflow.parentRentableW2UIItems[0].id;
            rec.ParentRentableName = app.raflow.parentRentableW2UIItems[0].id;
        });
    }

    // now load grid in division
    $('#ra-form #parentchild .grid-container').w2render(w2ui.RAParentChildGrid);

    // load the existing data in rentables component
    setTimeout(function () {
        var grid = w2ui.RAParentChildGrid;

        // first clear the grid
        grid.clear();

        // assign calculated grid records and refresh it
        grid.records = gridRecords;
        grid.refresh();

        // assign item prepared earlier for parent rentable list
        grid.getColumn("ParentRentableName").editable.items = app.raflow.parentRentableW2UIItems;
        grid.getColumn("ParentRentableName").render();

        // save the data if it's been modified
        saveParentChildCompData();

    }, 500);
};


//-----------------------------------------------------------------------------
// saveParentChildCompData - if there are any difference between server data
//                           and local data at this step then save the
//                           modified data on the server via API
//-----------------------------------------------------------------------------
window.saveParentChildCompData = function() {
    var compData = getRAFlowCompData("parentchild", app.raflow.activeFlowID) || [],
        dataToSaveFlag = false,
        gridRecords = w2ui.RAParentChildGrid.records || [];

    // first check the length
    if (gridRecords.length !== compData.length) {
        dataToSaveFlag = true;
    } else {
        var ridExists = false;
        // scan for each record from grid with compData, if RID not found then hit the API to save data
        gridRecords.forEach(function(gridRec) {
            compData.forEach(function(dataItem) {
                if (gridRec.CRID === dataItem.CRID && gridRec.PRID === dataItem.PRID) {
                    ridExists = true;
                    return false;
                }
            });
            if (!ridExists) { // if not found then it means we have mismatch in data
                dataToSaveFlag = true;
                return false;
            }
        });
    }

    // if have to save the data then update the local copy
    if (dataToSaveFlag) {
        var BID = getCurrentBID(),
            modCompData = [];

        gridRecords.forEach(function(rec) {
            modCompData.push({BID: BID, CRID: rec.CRID, PRID: rec.PRID});
        });

        // set this to it's position
        app.raflow.data[app.raflow.activeFlowID].parentchild = modCompData;

        // now hit the server API to save
        saveActiveCompData(modCompData, "parentchild");
    }
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
// setChildRentableLocalData - set the modified rentable data locally
//                              for requested RID by matching CRID
//-----------------------------------------------------------------------------
window.setChildRentableLocalData = function(RID, data) {
    var compData = getRAFlowCompData("parentchild", app.raflow.activeFlowID);
    var dataIndex = -1;
    compData.forEach(function(item, index) {
        if (item.CRID == RID) {
            dataIndex = index;
            return false;
        }
    });
    if (dataIndex > -1) {
        compData[dataIndex] = data;
    } else {
        compData.push(data);
    }
};
