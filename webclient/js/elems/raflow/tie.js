/* global
    loadRATieSection
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
                        ],
                        onClick: function (event) {
                            if (event.target === "pets-tie") {
                                w2ui.RATieLayout.html('main', w2ui.RAPetsTieGrid);
                            }
                            if (event.target === "vehicles-tie") {
                                w2ui.RATieLayout.html('main', w2ui.RAVehiclesTieGrid);
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
                    field: 'CRID',
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
                    field: 'CRID',
                    hidden: true
                },
                {
                    field: 'PRID',
                    hidden: true
                },
                {
                    field: 'Type',
                    caption: 'Type',
                    size: '120px',
                },
                {
                    field: 'VIN',
                    caption: 'VIN',
                    size: '120px',
                },
                {
                    field: 'Make',
                    caption: 'Make',
                    size: '120px',
                },
                {
                    field: 'Model',
                    caption: 'Model',
                    size: '120px',
                },
                {
                    field: 'Color',
                    caption: 'Color',
                    size: '120px',
                },
                {
                    field: 'Year',
                    caption: 'Year',
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
