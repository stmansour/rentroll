/* global
    SaveCompDataAJAX, GetRAFlowCompLocalData, getRecIDFromCRID, dispalyRAParentChildGridError,
    getChildRentableLocalData, SetChildRentableLocalData, saveParentChildCompData,
    EnableDisableRAFlowVersionGrid, SaveParentChildCompData
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
            style: 'border: none; display: block;',
            columns: [
                {
                    field: 'recid',
                    hidden: true
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
                    field: 'haveError',
                    size: '30px',
                    hidden: false,
                    render: function (record) {
                        var haveError = false;
                        if (app.raflow.validationErrors.parentchild) {
                            var parentchild = app.raflow.validationCheck.errors.parentchild.errors;
                            for (var i = 0; i < parentchild.length; i++) {
                                if (parentchild[i].PRID === record.PRID && parentchild[i].CRID === record.CRID && parentchild[i].total > 0) {
                                    haveError = true;
                                    break;
                                }
                            }
                        }
                        if (haveError) {
                            return '<i class="fas fa-exclamation-triangle" title="error"></i>';
                        } else {
                            return "";
                        }
                    }
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
            onRefresh: function(event) {
                var grid = this;
                event.onComplete = function() {
                    EnableDisableRAFlowVersionGrid(grid);
                };
            },
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
                        SetChildRentableLocalData(record.CRID, localData);

                        // SAVE DATA ON SERVER SIDE
                        SaveParentChildCompData()
                        .done(function(data) {
                            if (data.status === 'success') {
                                // save grid changes
                                grid.save();
                            } else {
                                grid.message(data.message);
                            }
                        })
                        .fail(function(data) {
                            console.log("failure " + data);
                        });
                    }
                };
            }
        });
    }

    // prepare parent and child rentable list based on rentables section data
    var rentableCompData = GetRAFlowCompLocalData("rentables") || [],
        compData = GetRAFlowCompLocalData("parentchild") || [],
        recidCounter = 1, // always starts with 1
        BID = getCurrentBID(),
        gridRecords = [];

    // always render data from latest modified rentable comp data
    rentableCompData.forEach(function(rentableItem) {
        var RID = rentableItem.RID,
            RentableName = rentableItem.RentableName;

        // 1 means this is child rentable
        if ( (rentableItem.RTFLAGS & (1 << app.rtFLAGS.IsChildRentable) ) != 0) {
            var PRID = 0;
            var cRentable = getChildRentableLocalData(RID);

            // parent Rentable ID found then for initial load in grid
            if (cRentable.PRID) {
                // if it's found in parent rentable list then keep as it is
                // else assign 0 if not found
                app.raflow.parentRentableW2UIItems.forEach(function(parentRItem) {
                    if (parentRItem.id == cRentable.PRID) {
                        PRID = cRentable.PRID;
                        return false;
                    }
                });
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

        // display row with light red background if it have error
        dispalyRAParentChildGridError();

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
    var compData = GetRAFlowCompLocalData("parentchild") || [],
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
        app.raflow.Flow.parentchild = modCompData;

        // now hit the server API to save
        SaveCompDataAJAX(modCompData, "parentchild");
    }
};

//-----------------------------------------------------------------------------
// getChildRentableLocalData -  returns the clone of child rentable data
//                              for requested RID by matching CRID
//-----------------------------------------------------------------------------
window.getChildRentableLocalData = function(RID, returnIndex) {
    var cloneData = {};
    var foundIndex = -1;
    var compData = GetRAFlowCompLocalData("parentchild");
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
// SetChildRentableLocalData - set the modified rentable data locally
//                              for requested RID by matching CRID
//-----------------------------------------------------------------------------
window.SetChildRentableLocalData = function(RID, data) {
    var compData = GetRAFlowCompLocalData("parentchild");
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

// dispalyRARentablesGridError
// It highlights grid's row if it have error
window.dispalyRAParentChildGridError = function (){
    // load grid errors if any
    var g = w2ui.RAParentChildGrid;
    var record, i;
    for (i = 0; i < g.records.length; i++) {
        // get record from grid to apply css
        record = g.get(g.records[i].recid);

        if (!("w2ui" in record)) {
            record.w2ui = {}; // init w2ui if not present
        }
        if (!("class" in record.w2ui)) {
            record.w2ui.class = ""; // init class string
        }
        if (!("style" in record.w2ui)) {
            record.w2ui.style = {}; // init style object
        }
    }

    if (app.raflow.validationErrors.parentchild) {
        var parentchild = app.raflow.validationCheck.errors.parentchild.errors;
        for (i = 0; i < parentchild.length; i++) {
            if (parentchild[i].total > 0) {
                var recid = getRecIDFromCRID(g, parentchild[i].CRID);
                g.get(recid).w2ui.style = "background-color: #EEB4B4";
                g.refreshRow(recid);
            }
        }
    }
};

// getRecIDFromRID It returns recid of grid record which matches TMPTCID
window.getRecIDFromCRID = function(grid, CRID){
    var recid;
    for (var i = 0; i < grid.records.length; i++) {
        if (grid.records[i].CRID === CRID) {
            recid = grid.records[i].recid;
        }
    }
    return recid;
};

//------------------------------------------------------------------------------
// SaveParentChildCompData - saves the data on server side
//------------------------------------------------------------------------------
window.SaveParentChildCompData = function() {
    var compData = GetRAFlowCompLocalData("parentchild");
    return SaveCompDataAJAX(compData, "parentchild");
};
