/* global
    loadRATieSection,
    getRAFlowCompData, reassignGridRecids,
    AssignTieVehiclesGridRecords, AssignTiePetsGridRecords,
    getTiePetLocalData, getTieVehicleLocalData,
    saveTiePetsData, saveTieVehiclesData,
    setTiePetLocalData, setTieVehicleLocalData
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
                            { id: 'tie-pets', caption: 'Pets' },
                            { id: 'tie-vehicles', caption: 'Vehicles' },
                            { id: 'tie-people', caption: 'People' },
                        ],
                        onClick: function (event) {
                            var grid;
                            switch(event.target) {
                            case "tie-pets":
                                grid = w2ui.RATiePetsGrid;
                                w2ui.RATieLayout.html('main', grid);

                                // once it's loaded then set the grid records
                                // and render parentRentableName columns in grid
                                setTimeout(function() {
                                    AssignTiePetsGridRecords();
                                    grid.getColumn("ParentRentableName").editable.items = app.raflow.parentRentableW2UIItems;
                                    grid.getColumn("ParentRentableName").render();
                                }, 500);
                                break;
                            case "tie-vehicles":
                                grid = w2ui.RATieVehiclesGrid;
                                w2ui.RATieLayout.html('main', grid);

                                // once it's loaded then set the grid records
                                // and render parentRentableName columns in grid
                                setTimeout(function() {
                                    AssignTieVehiclesGridRecords();
                                    grid.getColumn("ParentRentableName").editable.items = app.raflow.parentRentableW2UIItems;
                                    grid.getColumn("ParentRentableName").render();
                                }, 500);
                                break;
                            // case "tie-people":
                            //     grid = w2ui.RAPeopleTieGrid;
                            //     w2ui.RATieLayout.html('main', grid);

                            //     // once it's loaded then set the grid records
                            //     // and render parentRentableName columns in grid
                            //     setTimeout(function() {
                            //         setPeopleTieGridRecords();
                            //         grid.getColumn("ParentRentableName").editable.items = app.raflow.parentRentableW2UIItems;
                            //         grid.getColumn("ParentRentableName").render();
                            //     }, 500);
                            //     break;
                            }
                        }
                    }
                },
            ],
        });

        // Pets tie grid
        $().w2grid({
            name: 'RATiePetsGrid',
            header: 'Pets Tie',
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
                    field: 'TMPREFID',
                    hidden: true
                },
                {
                    field: 'PRID',
                    hidden: true
                },
                {
                    field: 'Name',
                    caption: 'Pet Name',
                    size: '140px',
                },
                {
                    field: 'Breed',
                    caption: 'Breed',
                    size: '100px',
                },
                {
                    field: 'Type',
                    caption: 'Type',
                    size: '120px',
                },
                {
                    field: 'ParentRentableName',
                    caption: 'Assign To',
                    size: '40%',
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
                        var localTiePetData = getTiePetLocalData(record.TMPREFID);

                        localTiePetData.PRID = record.PRID = parseInt(event.value_new);
                        record.ParentRentableName = parseInt(event.value_new);

                        // set data
                        grid.set(event.recid, record);
                        setTiePetLocalData(record.TMPREFID, localTiePetData);
                    }

                    // save grid changes
                    this.save();
                };
            }
        });

        // Vehicles tie grid
        $().w2grid({
            name: 'RATieVehiclesGrid',
            header: 'Vehicles Tie',
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
                    field: 'TMPREFID',
                    hidden: true
                },
                {
                    field: 'PRID',
                    hidden: true
                },
                {
                    field: 'Type',
                    caption: 'Type',
                    size: '90px',
                },
                {
                    field: 'VIN',
                    caption: 'VIN',
                    size: '90px',
                },
                {
                    field: 'Make',
                    caption: 'Make',
                    size: '90px',
                },
                {
                    field: 'Model',
                    caption: 'Model',
                    size: '90px',
                },
                {
                    field: 'Color',
                    caption: 'Color',
                    size: '90px',
                },
                {
                    field: 'Year',
                    caption: 'Year',
                    size: '90px',
                },
                {
                    field: 'ParentRentableName',
                    caption: 'Assign To',
                    size: '40%',
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
                        var localTieVehicleData = getTieVehicleLocalData(record.TMPREFID);

                        localTieVehicleData.PRID = record.PRID = parseInt(event.value_new);
                        record.ParentRentableName = parseInt(event.value_new);

                        // set data
                        grid.set(event.recid, record);
                        setTieVehicleLocalData(record.TMPREFID, localTieVehicleData);
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
        w2ui.RATieLayout.get("main").tabs.click("tie-pets");
    }, 0);
};

//-----------------------------------------------------------------------------
// getTiePetLocalData - returns the clone of pet data for requested TMPID
//                      from tie comp data
//-----------------------------------------------------------------------------
window.getTiePetLocalData = function(TMPID, returnIndex) {
    var cloneData = {};
    var foundIndex = -1;

    var compData = getRAFlowCompData("tie", app.raflow.activeFlowID);
    var tiePetsData = compData.pets || [];

    tiePetsData.forEach(function(item, index) {
        if (item.TMPREFID == TMPID) {
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
// setTiePetLocalData - set the modified tie pet data locally
//                      for requested TMPREFID by matching TMPREFID
//-----------------------------------------------------------------------------
window.setTiePetLocalData = function(TMPREFID, data) {
    var compData = getRAFlowCompData("tie", app.raflow.activeFlowID);
    var tiePetsData = compData.pets || [];

    var dataIndex = -1;
    tiePetsData.forEach(function(item, index) {
        if (item.TMPREFID == TMPREFID) {
            dataIndex = index;
            return false;
        }
    });
    if (dataIndex > -1) {
        tiePetsData[dataIndex] = data;
    } else {
        tiePetsData.push(data);
    }

    // modified data
    compData.pets = tiePetsData;
};

// -------------------------------------------------------------------------------
// AssignTiePetsGridRecords - assign calculated records in pets tie grid
//                            from "pets" comp data
// -------------------------------------------------------------------------------
window.AssignTiePetsGridRecords = function() {
    var BID = getCurrentBID();
    var petsCompData = getRAFlowCompData("pets", app.raflow.activeFlowID) || [];
    var grid = w2ui.RATiePetsGrid,
        tieGridRecords = [];

    petsCompData.forEach(function(petData) {
        var PRID = 0;
        var tiePet = getTiePetLocalData(petData.TMPID);

        // parent Rentable ID found then for initial load in grid
        if (tiePet.PRID) {
            // if it's found in parent rentable list then keep as it is
            // else assign 0 if not found
            app.raflow.parentRentableW2UIItems.forEach(function(parentRItem) {
                if (parentRItem.id == tiePet.PRID) {
                    PRID = tiePet.PRID;
                    return false;
                }
            });
        }

        var record = {
            recid:              0,
            BID:                BID,
            TMPREFID:           petData.TMPID,
            PRID:               PRID,
            ParentRentableName: PRID,
            Name:               petData.Name,
            Breed:              petData.Breed,
            Type:               petData.Type,
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

        // save the data if it's been modified
        saveTiePetsData();

    } else {
        grid.clear();
    }
};

//-----------------------------------------------------------------------------
// saveTiePetsData -     if there are any difference between server data
//                       and local data at this step then save the
//                       modified data on the server via API
//-----------------------------------------------------------------------------
window.saveTiePetsData = function() {
    var compData = getRAFlowCompData("tie", app.raflow.activeFlowID),
        tiePetsData = compData.pets || [],
        dataToSaveFlag = false,
        gridRecords = w2ui.RATiePetsGrid.records || [];

    // first check the length
    if (gridRecords.length !== tiePetsData.length) {
        dataToSaveFlag = true;
    } else {
        var tmpIDExists = false;
        // scan for each record from grid with compData, if RID not found then hit the API to save data
        gridRecords.forEach(function(gridRec) {
            tiePetsData.forEach(function(petItem) {
                if (gridRec.TMPREFID === petItem.TMPREFID && gridRec.PRID === petItem.PRID) {
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
            modTiePetsData = [];

        gridRecords.forEach(function(rec) {
            modTiePetsData.push({BID: BID, TMPREFID: rec.TMPREFID, PRID: rec.PRID});
        });

        // set this to it's position
        compData.pets = modTiePetsData;

        // now hit the server API to save
        saveActiveCompData(compData, "tie");
    }
};

//-----------------------------------------------------------------------------
// getTieVehicleLocalData - returns the clone of vehicle data for requested TMPID
//                          from tie comp data
//-----------------------------------------------------------------------------
window.getTieVehicleLocalData = function(TMPID, returnIndex) {
    var cloneData = {};
    var foundIndex = -1;

    var compData = getRAFlowCompData("tie", app.raflow.activeFlowID);
    var tieVehiclesData = compData.vehicles || [];

    tieVehiclesData.forEach(function(item, index) {
        if (item.TMPREFID == TMPID) {
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
// setTieVehicleLocalData - set the modified tie vehicle data locally
//                          for requested TMPREFID by matching TMPREFID
//-----------------------------------------------------------------------------
window.setTieVehicleLocalData = function(TMPREFID, data) {
    var compData = getRAFlowCompData("tie", app.raflow.activeFlowID);
    var tieVehiclesData = compData.vehicles || [];

    var dataIndex = -1;
    tieVehiclesData.forEach(function(item, index) {
        if (item.TMPREFID == TMPREFID) {
            dataIndex = index;
            return false;
        }
    });
    if (dataIndex > -1) {
        tieVehiclesData[dataIndex] = data;
    } else {
        tieVehiclesData.push(data);
    }

    // modified data
    compData.vehicles = tieVehiclesData;
};

// -------------------------------------------------------------------------------
// AssignTieVehiclesGridRecords - assign calculated records in vehicles tie grid
//                                from "vehicles" comp data
// -------------------------------------------------------------------------------
window.AssignTieVehiclesGridRecords = function() {
    var BID = getCurrentBID();
    var vehiclesCompData = getRAFlowCompData("vehicles", app.raflow.activeFlowID) || [];
    var grid = w2ui.RATieVehiclesGrid,
        tieGridRecords = [];

    vehiclesCompData.forEach(function(vehicleData) {
        var PRID = 0;
        var tieVehicle = getTieVehicleLocalData(vehicleData.TMPID);

        // parent Rentable ID found then for initial load in grid
        if (tieVehicle.PRID) {
            // if it's found in parent rentable list then keep as it is
            // else assign 0 if not found
            app.raflow.parentRentableW2UIItems.forEach(function(parentRItem) {
                if (parentRItem.id == tieVehicle.PRID) {
                    PRID = tieVehicle.PRID;
                    return false;
                }
            });
        }

        var record = {
            recid:              0,
            BID:                BID,
            TMPREFID:           vehicleData.TMPID,
            PRID:               PRID,
            ParentRentableName: PRID,
            Type:               vehicleData.Type,
            VIN:                vehicleData.VIN,
            Make:               vehicleData.Make,
            Model:              vehicleData.Model,
            Color:              vehicleData.Color,
            Year:               vehicleData.Year,
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

        grid.records = tieGridRecords;
        reassignGridRecids(grid.name);

        // assign item prepared earlier for parent rentable list
        grid.getColumn("ParentRentableName").editable.items = app.raflow.parentRentableW2UIItems;
        grid.getColumn("ParentRentableName").render();

        // save the data if it's been modified
        saveTieVehiclesData();

    } else {
        grid.clear();
    }
};

//-----------------------------------------------------------------------------
// saveTieVehiclesData - if there are any difference between server data
//                       and local data at this step then save the
//                       modified data on the server via API
//-----------------------------------------------------------------------------
window.saveTieVehiclesData = function() {
    var compData = getRAFlowCompData("tie", app.raflow.activeFlowID),
        tieVehiclesData = compData.vehicles || [],
        dataToSaveFlag = false,
        gridRecords = w2ui.RATieVehiclesGrid.records || [];

    // first check the length
    if (gridRecords.length !== tieVehiclesData.length) {
        dataToSaveFlag = true;
    } else {
        var tmpIDExists = false;
        // scan for each record from grid with compData, if RID not found then hit the API to save data
        gridRecords.forEach(function(gridRec) {
            tieVehiclesData.forEach(function(vehicleItem) {
                if (gridRec.TMPREFID === vehicleItem.TMPREFID && gridRec.PRID === vehicleItem.PRID) {
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
            modTieVehiclesData = [];

        gridRecords.forEach(function(rec) {
            modTieVehiclesData.push({BID: BID, TMPREFID: rec.TMPREFID, PRID: rec.PRID});
        });

        // set this to it's position
        compData.vehicles = modTieVehiclesData;

        // now hit the server API to save
        saveActiveCompData(compData, "tie");
    }
};


