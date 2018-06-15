/* global
    loadRATieSection,
    getRAFlowCompData, reassignGridRecids,
    getTiePeopleLocalData, setTiePeopleLocalData, AssignTiePeopleGridRecords, saveTiePeopleData,
    getFullName
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
                    style: 'background-color: white; border: 1px solid silver; padding: 0px;',
                    tabs: {
                        style: "padding-top: 10px;",
                        active: 'tie-pets',
                        tabs: [
                            { id: 'tie-people', caption: 'People' },
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
                    field: 'TMPTCID',
                    hidden: true
                },
                {
                    field: 'PRID',
                    hidden: true
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
            onChange: function(event) {
                var grid = this;
                event.onComplete = function() {
                    // parent rentable name column index
                    var PRNCI = grid.getColumn("ParentRentableName", true);
                    if (PRNCI === event.column) {
                        var record = grid.get(event.recid);
                        var localTiePeopleData = getTiePeopleLocalData(record.TMPTCID);

                        localTiePeopleData.PRID = record.PRID = parseInt(event.value_new);
                        record.ParentRentableName = parseInt(event.value_new);

                        // set data
                        grid.set(event.recid, record);
                        setTiePeopleLocalData(record.TMPTCID, localTiePeopleData);
                        console.debug(getRAFlowCompData("tie", app.raflow.activeFlowID));
                    }

                    // save grid changes
                    this.save();
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
// getTiePeopleLocalData - returns the clone of people data for requested TMPTCID
//                      from tie comp data
//-----------------------------------------------------------------------------
window.getTiePeopleLocalData = function(TMPTCID, returnIndex) {
    var cloneData = {};
    var foundIndex = -1;

    var compData = getRAFlowCompData("tie", app.raflow.activeFlowID);
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
// setTiePeopleLocalData - set the modified tie people data locally
//                      for requested TMPTCID by matching TMPTCID
//-----------------------------------------------------------------------------
window.setTiePeopleLocalData = function(TMPTCID, data) {
    var compData = getRAFlowCompData("tie", app.raflow.activeFlowID);
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
    var BID = getCurrentBID();
    var peopleCompData = getRAFlowCompData("people", app.raflow.activeFlowID) || [];
    var grid = w2ui.RATiePeopleGrid,
        tieGridRecords = [];

    peopleCompData.forEach(function(peopleData) {
        var PRID = 0;
        var tiePeople = getTiePeopleLocalData(peopleData.TMPTCID);

        // NOTE: list down every single person added in people section

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
            BID:                BID,
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

    } else {
        grid.clear();
    }

    // save the data if it's been modified
    saveTiePeopleData();
};

//-----------------------------------------------------------------------------
// saveTiePeopleData -   if there are any difference between server data
//                       and local data at this step then save the
//                       modified data on the server via API
//-----------------------------------------------------------------------------
window.saveTiePeopleData = function() {
    var compData = getRAFlowCompData("tie", app.raflow.activeFlowID),
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
        var BID = getCurrentBID(),
            modTiePeopleData = [];

        gridRecords.forEach(function(rec) {
            modTiePeopleData.push({BID: BID, TMPTCID: rec.TMPTCID, PRID: rec.PRID});
        });

        // set this to it's position
        compData.people = modTiePeopleData;

        // now hit the server API to save
        saveActiveCompData(compData, "tie");
    }
};

