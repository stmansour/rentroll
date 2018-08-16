/* global
    GetRAFlowCompLocalData, reassignGridRecids, SaveCompDataAJAX,
    GetTiePeopleLocalData, SetTiePeopleLocalData, AssignTiePeopleGridRecords, SaveTiePeopleData,
    getFullName, dispalyRATiePeopleGridError, getRecIDFromTMPTCID,
    EnableDisableRAFlowVersionGrid, SaveTieCompData
*/

"use strict";

//-----------------------------------------------------------------------
// loadRATieSection -
//-----------------------------------------------------------------------
window.loadRATieSection = function () {

    if (!("RATieLayout" in w2ui)) {

        //------------------------------------------------------------------------
        //          tieLayout which holds all grids
        //------------------------------------------------------------------------
        $().w2layout({
            name: 'RATieLayout',
            panels: [
                {
                    type: 'main',
                    overflow: "hidden",
                    style: 'border: none; background-color: white; padding: 0px;',
                    tabs: {
                        style: "padding-top: 10px;",
                        tabs: [
                            { id: 'tie-people', caption: 'Occupants' },
                        ],
                        onClick: function (event) {
                            switch(event.target) {
                            case "tie-people":
                                w2ui.RATieLayout.html('main', w2ui.RATiePeopleGrid);

                                // once it's loaded then set the grid records
                                // and render parentRentableName columns in grid
                                setTimeout(function() {
                                    AssignTiePeopleGridRecords();
                                }, 500);
                                break;
                            }
                        }
                    }
                },
            ],
        });

        // TIe People grid
        $().w2grid({
            name: 'RATiePeopleGrid',
            header: 'People Tie',
            show: {
                toolbar:    false,
                footer:     true
            },
            multiSelect: false,
            style: 'border: none; display: block;',
            columns: [
                {
                    field: 'recid',
                    hidden: true
                },
                {
                    field: 'TMPTCID',
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
                        if (app.raflow.validationErrors.tie) {
                            var tiePeople = app.raflow.validationCheck.errors.tie.people.errors;
                            for (var i = 0; i < tiePeople.length; i++) {
                                if (tiePeople[i].TMPTCID === record.TMPTCID && tiePeople[i].total > 0) {
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
                    field: 'FullName',
                    caption: 'Full Name',
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
                        var localTiePeopleData = GetTiePeopleLocalData(record.TMPTCID);

                        localTiePeopleData.PRID = record.PRID = parseInt(event.value_new);
                        record.ParentRentableName = parseInt(event.value_new);

                        // set data
                        grid.set(event.recid, record);
                        SetTiePeopleLocalData(record.TMPTCID, localTiePeopleData);

                        // SAVE DATA ON SERVER SIDE
                        SaveTieCompData()
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

    // now load layout in division
    $('#ra-form #tie .layout-container').w2render(w2ui.RATieLayout);
    setTimeout(function() {
        w2ui.RATieLayout.get("main").tabs.click("tie-people");
    }, 0);
};

//-----------------------------------------------------------------------------
// GetTiePeopleLocalData - returns the clone of people data for requested TMPTCID
//                      from tie comp data
//-----------------------------------------------------------------------------
window.GetTiePeopleLocalData = function(TMPTCID, returnIndex) {
    var cloneData = {};
    var foundIndex = -1;

    var compData = GetRAFlowCompLocalData("tie");
    var tiePeopleData = compData.people || [];

    tiePeopleData.forEach(function(item, index) {
        if (item.TMPTCID == TMPTCID) {
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
// SetTiePeopleLocalData - set the modified tie people data locally
//                      for requested TMPTCID by matching TMPTCID
//-----------------------------------------------------------------------------
window.SetTiePeopleLocalData = function(TMPTCID, data) {
    var compData = GetRAFlowCompLocalData("tie");
    var tiePeopleData = compData.people || [];

    var dataIndex = -1;
    tiePeopleData.forEach(function(item, index) {
        if (item.TMPTCID == TMPTCID) {
            dataIndex = index;
            return false;
        }
    });
    if (dataIndex > -1) {
        tiePeopleData[dataIndex] = data;
    } else {
        tiePeopleData.push(data);
    }

    // modified data
    compData.people = tiePeopleData;
};

// -------------------------------------------------------------------------------
// AssignTiePeopleGridRecords - assign calculated records in people tie grid
//                            from "people" comp data
// -------------------------------------------------------------------------------
window.AssignTiePeopleGridRecords = function() {
    var peopleCompData = GetRAFlowCompLocalData("people") || [];
    var grid = w2ui.RATiePeopleGrid,
        tieGridRecords = [];

    peopleCompData.forEach(function(peopleData) {

        // NOTE: list down only occupants in tie section
        if (!peopleData.IsOccupant) {
            return; // continue to next person
        }

        var PRID = 0;
        var tiePeople = GetTiePeopleLocalData(peopleData.TMPTCID);

        // parent Rentable ID found then for initial load in grid
        if (tiePeople.PRID) {
            // if it's found in parent rentable list then keep as it is
            // else assign 0 if not found
            app.raflow.parentRentableW2UIItems.forEach(function(parentRItem) {
                if (parentRItem.id == tiePeople.PRID) {
                    PRID = tiePeople.PRID;
                    return false;
                }
            });
        }

        // get full name from individual record
        var FullName;
        if (!peopleData.IsCompany) {
            FullName = getFullName(peopleData);
        } else {
            FullName = peopleData.Employer;
        }

        var record = {
            recid:              0,
            TMPTCID:           peopleData.TMPTCID,
            PRID:               PRID,
            ParentRentableName: PRID,
            FullName:           FullName,
        };
        tieGridRecords.push(record);
    });

    if (tieGridRecords.length > 0) {

        // if there is only one parent rentable then pre-select it for all child rentable
        // otherwise built drop down menu
        if (app.raflow.parentRentableW2UIItems.length == 0) {
            tieGridRecords.forEach(function(rec) {
                rec.PRID = 0;
                rec.ParentRentableName = 0;
            });
        } else if (app.raflow.parentRentableW2UIItems.length == 1) {
            // re-assign PRID
            tieGridRecords.forEach(function(rec) {
                rec.PRID = app.raflow.parentRentableW2UIItems[0].id;
                rec.ParentRentableName = app.raflow.parentRentableW2UIItems[0].id;
            });
        }

        // feed array of records to grid
        grid.records = tieGridRecords;
        reassignGridRecids(grid.name);

        // assign item prepared earlier for parent rentable list
        grid.getColumn("ParentRentableName").editable.items = app.raflow.parentRentableW2UIItems;
        grid.getColumn("ParentRentableName").render();

        // display row with light red background if it have error
        dispalyRATiePeopleGridError();

    } else {
        grid.clear();
    }

    // save the data if it's been modified
    SaveTiePeopleData();
};

//-----------------------------------------------------------------------------
// SaveTiePeopleData -   if there are any difference between server data
//                       and local data at this step then save the
//                       modified data on the server via API
//-----------------------------------------------------------------------------
window.SaveTiePeopleData = function() {
    var compData = GetRAFlowCompLocalData("tie"),
        tiePeopleData = compData.people || [],
        dataToSaveFlag = false,
        gridRecords = w2ui.RATiePeopleGrid.records || [];

    // first check the length
    if (gridRecords.length !== tiePeopleData.length) {
        dataToSaveFlag = true;
    } else {
        var tmpIDExists = false;
        // scan for each record from grid with compData, if RID not found then hit the API to save data
        gridRecords.forEach(function(gridRec) {
            tiePeopleData.forEach(function(peopleItem) {
                if (gridRec.TMPTCID === peopleItem.TMPTCID && gridRec.PRID === peopleItem.PRID) {
                    tmpIDExists = true;
                    return false;
                }
            });
            if (!tmpIDExists) { // if not found then it means we have mismatch in data
                dataToSaveFlag = true;
                return false;
            }
        });
    }

    // if have to save the data then update the local copy
    if (dataToSaveFlag) {
        var modTiePeopleData = [];

        gridRecords.forEach(function(rec) {
            modTiePeopleData.push({TMPTCID: rec.TMPTCID, PRID: rec.PRID});
        });

        // set this to it's position
        compData.people = modTiePeopleData;

        // now hit the server API to save
        SaveCompDataAJAX(compData, "tie");
    }
};

// dispalyRATiePeopleGridError
// It highlights grid's row if it have error
window.dispalyRATiePeopleGridError = function (){
    // load grid errors if any
    var g = w2ui.RATiePeopleGrid;
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

    if (app.raflow.validationErrors.tie) {
        var tie = app.raflow.validationCheck.errors.tie.people.errors;
        for (i = 0; i < tie.length; i++) {
            if (tie[i].total > 0) {
                var recid = getRecIDFromTMPTCID(g, tie[i].TMPTCID);
                g.get(recid).w2ui.style = "background-color: #EEB4B4";
                g.refreshRow(recid);
            }
        }
    }
};

// getRecIDFromRID It returns recid of grid record which matches TMPTCID
window.getRecIDFromTMPTCID = function(grid, TMPTCID){
    var recid;
    for (var i = 0; i < grid.records.length; i++) {
        if (grid.records[i].TMPTCID === TMPTCID) {
            recid = grid.records[i].recid;
        }
    }
    return recid;
};

//------------------------------------------------------------------------------
// SaveTieCompData - saves the data on server side
//------------------------------------------------------------------------------
window.SaveTieCompData = function() {
    var compData = GetRAFlowCompLocalData("tie");
    return SaveCompDataAJAX(compData, "tie");
};
