/* global
    loadRATieSection,
    getRAFlowCompData, reassignGridRecids,
    AssignTieVehiclesGridRecords, AssignTiePetsGridRecords, AssignTiePeopleGridRecords,
    getTiePetLocalData, getTieVehicleLocalData, getTiePeopleLocalData,
    saveTiePetsData, saveTieVehiclesData, saveTiePeopleData,
    setTiePetLocalData, setTieVehicleLocalData, setTiePeopleLocalData,
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
                            { id: 'tie-pets', caption: 'Pets' },
                            { id: 'tie-vehicles', caption: 'Vehicles' },
                            { id: 'tie-people', caption: 'People' },
                        ],
                        onClick: function (event) {
                            switch(event.target) {
                            case "tie-pets":
                                w2ui.RATieLayout.html('main', w2ui.RATiePetsGrid);

                                // once it's loaded then set the grid records
                                // and render parentRentableName columns in grid
                                setTimeout(function() {
                                    AssignTiePetsGridRecords();
                                }, 500);
                                break;
                            case "tie-vehicles":
                                w2ui.RATieLayout.html('main', w2ui.RATieVehiclesGrid);

                                // once it's loaded then set the grid records
                                // and render parentRentableName columns in grid
                                setTimeout(function() {
                                    AssignTieVehiclesGridRecords();
                                }, 500);
                                break;
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
                    field: 'TMPPETID',
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
                        var localTiePetData = getTiePetLocalData(record.TMPPETID);

                        localTiePetData.PRID = record.PRID = parseInt(event.value_new);
                        record.ParentRentableName = parseInt(event.value_new);

                        // set data
                        grid.set(event.recid, record);
                        setTiePetLocalData(record.TMPPETID, localTiePetData);
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
                    field: 'TMPVID',
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
                        var localTieVehicleData = getTieVehicleLocalData(record.TMPVID);

                        localTieVehicleData.PRID = record.PRID = parseInt(event.value_new);
                        record.ParentRentableName = parseInt(event.value_new);

                        // set data
                        grid.set(event.recid, record);
                        setTieVehicleLocalData(record.TMPVID, localTieVehicleData);
                    }

                    // save grid changes
                    this.save();
                };
            }
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
        w2ui.RATieLayout.get("main").tabs.click("tie-pets");
    }, 0);
};

//-----------------------------------------------------------------------------
// getTiePetLocalData - returns the clone of pet data for requested TMPPETID
//                      from tie comp data
//-----------------------------------------------------------------------------
window.getTiePetLocalData = function(TMPPETID, returnIndex) {
    var cloneData = {};
    var foundIndex = -1;

    var compData = getRAFlowCompData("tie", app.raflow.activeFlowID);
    var tiePetsData = compData.pets || [];

    tiePetsData.forEach(function(item, index) {
        if (item.TMPPETID == TMPPETID) {
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
//                      for requested TMPPETID by matching TMPPETID
//-----------------------------------------------------------------------------
window.setTiePetLocalData = function(TMPPETID, data) {
    var compData = getRAFlowCompData("tie", app.raflow.activeFlowID);
    var tiePetsData = compData.pets || [];

    var dataIndex = -1;
    tiePetsData.forEach(function(item, index) {
        if (item.TMPPETID == TMPPETID) {
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
        var tiePet = getTiePetLocalData(petData.TMPPETID);

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
            TMPPETID:           petData.TMPPETID,
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

    } else {
        grid.clear();
    }

    // save the data if it's been modified
    saveTiePetsData();
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
                if (gridRec.TMPPETID === petItem.TMPPETID && gridRec.PRID === petItem.PRID) {
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
            modTiePetsData.push({BID: BID, TMPPETID: rec.TMPPETID, PRID: rec.PRID});
        });

        // set this to it's position
        compData.pets = modTiePetsData;

        // now hit the server API to save
        saveActiveCompData(compData, "tie");
    }
};

//-----------------------------------------------------------------------------
// getTieVehicleLocalData - returns the clone of vehicle data for requested TMPVID
//                          from tie comp data
//-----------------------------------------------------------------------------
window.getTieVehicleLocalData = function(TMPVID, returnIndex) {
    var cloneData = {};
    var foundIndex = -1;

    var compData = getRAFlowCompData("tie", app.raflow.activeFlowID);
    var tieVehiclesData = compData.vehicles || [];

    tieVehiclesData.forEach(function(item, index) {
        if (item.TMPVID == TMPVID) {
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
//                          for requested TMPVID by matching TMPVID
//-----------------------------------------------------------------------------
window.setTieVehicleLocalData = function(TMPVID, data) {
    var compData = getRAFlowCompData("tie", app.raflow.activeFlowID);
    var tieVehiclesData = compData.vehicles || [];

    var dataIndex = -1;
    tieVehiclesData.forEach(function(item, index) {
        if (item.TMPVID == TMPVID) {
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
        var tieVehicle = getTieVehicleLocalData(vehicleData.TMPVID);

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
            TMPVID:           vehicleData.TMPVID,
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

    } else {
        grid.clear();
    }

    // save the data if it's been modified
    saveTieVehiclesData();

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
                if (gridRec.TMPVID === vehicleItem.TMPVID && gridRec.PRID === vehicleItem.PRID) {
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
            modTieVehiclesData.push({BID: BID, TMPVID: rec.TMPVID, PRID: rec.PRID});
        });

        // set this to it's position
        compData.vehicles = modTieVehiclesData;

        // now hit the server API to save
        saveActiveCompData(compData, "tie");
    }
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

        // if it's a payor then ignore to set in grid
        if (peopleData.IsRenter) {
            return; // return from forEach method
        }

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

