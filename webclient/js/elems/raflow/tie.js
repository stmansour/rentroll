/* global
    loadRATieSection,
    getRAFlowCompData, reassignGridRecids,
    setVehiclesTieGridRecords, setPetsTieGridRecords
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
                        active: 'pets-tie',
                        tabs: [
                            { id: 'pets-tie', caption: 'Pets' },
                            { id: 'vehicles-tie', caption: 'Vehicles' },
                            { id: 'payors-tie', caption: 'Payors' },
                        ],
                        onClick: function (event) {
                            switch(event.target) {
                            case "pets-tie":
                                w2ui.RATieLayout.html('main', w2ui.RAPetsTieGrid);
                                setPetsTieGridRecords();
                                break;
                            case "vehicles-tie":
                                w2ui.RATieLayout.html('main', w2ui.RAVehiclesTieGrid);
                                setVehiclesTieGridRecords();
                                break;
                            case "payors-tie":
                                // w2ui.RATieLayout.html('main', w2ui.RAVehiclesTieGrid);
                                // setVehiclesTieGridRecords();
                                // break;
                            }
                        }
                    }
                },
            ],
        });

        // Pets tie grid
        $().w2grid({
            name: 'RAPetsTieGrid',
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
                        record.PRID = parseInt(event.value_new);
                        record.ParentRentableName = parseInt(event.value_new);
                        grid.set(event.recid, record);
                    }

                    // save grid changes
                    this.save();
                };
            }
        });

        // Vehicles tie grid
        $().w2grid({
            name: 'RAVehiclesTieGrid',
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
                        record.PRID = parseInt(event.value_new);
                        record.ParentRentableName = parseInt(event.value_new);
                        grid.set(event.recid, record);
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
        w2ui.RATieLayout.get("main").tabs.click("pets-tie");
    }, 0);
};

// -------------------------------------------------------------------------------
// setPetsTieGridRecords - set records in pets tie grid from "pets" comp data
// -------------------------------------------------------------------------------
window.setPetsTieGridRecords = function() {
    var BID = getCurrentBID();
    var petsCompData = getRAFlowCompData("pets", app.raflow.activeFlowID) || [];
    var grid = w2ui.RAPetsTieGrid,
        tieGridRecords = [];

    petsCompData.forEach(function(petData) {
        var record = {
            recid:              0,
            BID:                BID,
            PRID:               0,
            Name:               petData.Name,
            Breed:              petData.Breed,
            Type:               petData.Type,
            ParentRentableName: 0,
        };
        tieGridRecords.push(record);
    });

    if (tieGridRecords.length > 0) {
        grid.records = tieGridRecords;
        reassignGridRecids(grid.name);

        // assign item prepared earlier for parent rentable list
        grid.getColumn("ParentRentableName").editable.items = app.raflow.parentRentableW2UIItems;
        grid.getColumn("ParentRentableName").render();
    } else {
        grid.clear();
    }
};

// -------------------------------------------------------------------------------
// setVehiclesTieGridRecords - set records in vehicles tie grid from "vehicles" comp data
// -------------------------------------------------------------------------------
window.setVehiclesTieGridRecords = function() {
    var BID = getCurrentBID();
    var vehiclesCompData = getRAFlowCompData("vehicles", app.raflow.activeFlowID) || [];
    var grid = w2ui.RAVehiclesTieGrid,
        tieGridRecords = [];

    vehiclesCompData.forEach(function(vehicleData) {
        var record = {
            recid:              0,
            BID:                BID,
            PRID:               0,
            Type:               vehicleData.Type,
            VIN:                vehicleData.VIN,
            Make:               vehicleData.Make,
            Model:              vehicleData.Model,
            Color:              vehicleData.Color,
            Year:               vehicleData.Year,
            ParentRentableName: 0,
        };
        tieGridRecords.push(record);
    });

    if (tieGridRecords.length > 0) {
        grid.records = tieGridRecords;
        reassignGridRecids(grid.name);

        // assign item prepared earlier for parent rentable list
        grid.getColumn("ParentRentableName").editable.items = app.raflow.parentRentableW2UIItems;
        grid.getColumn("ParentRentableName").render();
    } else {
        grid.clear();
    }
};


