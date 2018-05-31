/* global
    RACompConfig, sliderContentDivLength, reassignGridRecids,
    hideSliderContent, appendNewSlider, showSliderContentW2UIComp,
    loadTargetSection, requiredFieldsFulFilled, getRAFlowPartTypeIndex, initRAFlowAjax,
    getRAFlowAllParts, saveActiveCompData, toggleHaveCheckBoxDisablity, getRAFlowCompData,
    lockOnGrid,
    getVehicleFormInitalRecord
*/

"use strict";

// -------------------------------------------------------------------------------
// Rental Agreement - Vehicles Grid
// -------------------------------------------------------------------------------
window.getVehicleFormInitalRecord = function (BID, BUD, previousFormRecord) {
    var t = new Date(),
        nyd = new Date(new Date().setFullYear(new Date().getFullYear() + 1));

    var defaultFormData = {
        recid: 0,
        VID: 0,
        BID: BID,
        TCID: 0,
        VIN: "",
        Type: "",
        Make: "",
        Model: "",
        Color: "",
        LicensePlateState: "",
        LicensePlateNumber: "",
        ParkingPermitNumber: "",
        ParkingPermitFee: 0,
        DtStart: w2uiDateControlString(t),
        DtStop: w2uiDateControlString(nyd)
    };

    // if it called after 'save and add another' action there previous form record is passed as Object
    // else it is null
    if ( previousFormRecord ) {
        defaultFormData = setDefaultFormFieldAsPreviousRecord(
            [ 'Type', 'Make', 'Model', 'Color', 'Year', 'LicensePlateState', 'LicensePlateNumber', 'VIN',
                'ParkingPermitNumber', 'ParkingPermitFee'], // Fields to Reset
            defaultFormData,
            previousFormRecord
        );
    }

    return defaultFormData;
};

window.loadRAVehiclesGrid = function () {

    // if form is loaded then return
    if (!("RAVehiclesGrid" in w2ui)) {

        // Add vehicle information form
        $().w2form({
            name    : 'RAVehicleForm',
            header  : 'Add Vehicle form',
            formURL : '/webclient/html/formravehicles.html',
            toolbar :{
                items: [
                    { id: 'bt3', type: 'spacer' },
                    { id: 'btnClose', type: 'button', icon: 'fas fa-times'}
                ],
                onClick: function (event) {
                    switch (event.target){
                        case 'btnClose':
                            hideSliderContent();
                            break;
                    }
                }
            },
            fields  : [
                { field: 'recid', type: 'int', required: false, html: { caption: 'recid', page: 0, column: 0 } },
                { field: 'Type', type: 'text', required: true},
                { field: 'Make', type: 'text', required: true},
                { field: 'Model', type: 'text', required: true},
                { field: 'Color', type: 'text', required: true},
                { field: 'Year', type: 'text', required: true},
                { field: 'LicensePlateState', type: 'text', required: true},
                { field: 'LicensePlateNumber', type: 'text', required: true},
                { field: 'VIN', type: 'text', required: true},
                { field: 'ParkingPermitNumber', type: 'text', required: true},
                { field: 'ParkingPermitFee', type: 'money', required: true},
                { field: 'DtStart', type: 'date', required: false, html: { caption: 'DtStart', page: 0, column: 0 } },
                { field: 'DtStop', type: 'date', required: false, html: { caption: 'DtStop', page: 0, column: 0 } },
                { field: 'LastModTime', type: 'time', required: false, html: { caption: 'LastModTime', page: 0, column: 0 } },
                { field: 'LastModBy', type: 'int', required: false, html: { caption: 'LastModBy', page: 0, column: 0 } },
            ],
            onRefresh: function(event) {
                event.onComplete = function() {
                    var f = w2ui.RAVehicleForm,
                        header = "Edit Rental Agreement Vehicles ({0})";

                    // there is NO PETID actually, so have to work around with recid key
                    formRefreshCallBack(f, "recid", header);

                    // hide delete button if it is NewRecord
                    var isNewRecord = (w2ui.RAVehiclesGrid.get(f.record.recid, true) === null);
                    if (isNewRecord) {
                        $(f.box).find("button[name=delete]").addClass("hidden");
                    } else {
                        $(f.box).find("button[name=delete]").removeClass("hidden");
                    }
                };
            },
            actions : {
                save: function () {
                    var form = this;
                    var grid = w2ui.RAVehiclesGrid;
                    var errors = form.validate();
                    if (errors.length > 0) return;
                    var record = $.extend(true, { recid: grid.records.length + 1 }, form.record);
                    var recordsData = $.extend(true, [], grid.records);
                    var isNewRecord = (grid.get(record.recid, true) === null);

                    // if it doesn't exist then only push
                    if (isNewRecord) {
                        recordsData.push(record);
                    }

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // save this records in json Data
                    saveActiveCompData(recordsData, "vehicles")
                        .done(function(data) {
                            if (data.status === 'success') {
                                // if null
                                if(isNewRecord) {
                                    grid.add(record);
                                }else {
                                    grid.set(record.recid, record);
                                }
                                form.clear();

                                // Disable "have vehicles?" checkbox if there is any record.
                                toggleHaveCheckBoxDisablity('RAVehiclesGrid');

                                // close the form
                                hideSliderContent();
                            } else {
                                form.message(data.message);
                            }
                        })
                        .fail(function(data) {
                            console.log("failure " + data);
                        });
                },
                saveadd: function () {
                    var BID = getCurrentBID(),
                        BUD = getBUDfromBID(BID);

                    var form = this;
                    var grid = w2ui.RAVehiclesGrid;
                    var errors = form.validate();
                    if (errors.length > 0) return;
                    var record = $.extend(true, { recid: grid.records.length + 1 }, form.record);
                    var recordsData = $.extend(true, [], grid.records);
                    var isNewRecord = (grid.get(record.recid, true) === null);

                    if (isNewRecord) {
                        recordsData.push(record);
                    }

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // save this records in json Data
                    saveActiveCompData(recordsData, "vehicles")
                        .done(function(data) {
                            if (data.status === 'success') {
                                // clear the grid select recid
                                app.last.grid_sel_recid  = -1;
                                // selectNone
                                grid.selectNone();

                                if (isNewRecord) {
                                    grid.add(record);
                                } else {
                                    grid.set(record.recid, record);
                                }
                                form.record = getVehicleFormInitalRecord(BID, BUD, form.record);
                                form.record.recid =grid.records.length + 1;
                                form.refresh();
                                form.refresh();
                            } else {
                                form.message(data.message);
                            }
                        })
                        .fail(function(data) {
                            console.log("failure " + data);
                        });
                },
                delete: function () {
                    var form = this;
                    var grid = w2ui.RAVehiclesGrid;

                    // backup the records
                    var records = $.extend(true, [], grid.records);
                    for (var i = 0; i < records.length; i++) {
                        if(records[i].recid == form.record.recid) {
                            records.splice(i, 1);
                        }
                    }

                    // save this records in json Data
                    saveActiveCompData(records, "vehicles")
                        .done(function(data) {
                            if (data.status === 'success') {
                                // clear the grid select recid
                                app.last.grid_sel_recid  =-1;
                                // selectNone
                                grid.selectNone();

                                grid.remove(form.record.recid);
                                form.clear();

                                // Disable "have vehicles?" checkbox if there is any record.
                                toggleHaveCheckBoxDisablity('RAVehiclesGrid');

                                // need to refresh the grid as it will re-assign new recid
                                reassignGridRecids(grid.name);

                                // close the form
                                hideSliderContent();
                            } else {
                                form.message(data.message);
                            }
                        })
                        .fail(function(data) {
                            console.log("failure " + data);
                        });

                }
            }
        });

        // vehicles grid
        $().w2grid({
            name    : 'RAVehiclesGrid',
            header  : 'Vehicles',
            show    : {
                toolbar         : true,
                toolbarSearch   : false,
                toolbarReload   : true,
                toolbarInput    : false,
                toolbarColumns  : false,
                footer          : true,
                toolbarAdd      : true   // indicates if toolbar add new button is visible
            },
            multiSelect: false,
            style   : 'border: 0px solid black; display: block;',
            columns : [
                {
                    field: 'recid',
                    hidden: true
                },
                {
                    field: 'VID',
                    hidden: true
                },
                {
                    field: 'BID',
                    hidden: true
                },
                {
                    field: 'TCID',
                    hidden: true
                },
                {
                    field: 'Type',
                    caption: 'Type',
                    size: '80px',
                    editable: {type: 'text'}
                },
                {
                    field: 'VIN',
                    caption: 'VIN',
                    size: '80px'
                },
                {
                    field: 'Make',
                    caption: 'Make',
                    size: '80px'
                },
                {
                    field: 'Model',
                    caption: 'Model',
                    size: '80px'
                },
                {
                    field: 'Color',
                    caption: 'Color',
                    size: '80px'
                },
                {
                    field: 'LicensePlateState',
                    caption: 'License Plate<br>State',
                    size: '100px'
                },
                {
                    field: 'LicensePlateNumber',
                    caption: 'License Plate<br>Number',
                    size: '100px'
                },
                {
                    field: 'ParkingPermitNumber',
                    caption: 'Parking Permit <br>Number',
                    size: '100px'
                },
                {
                    field: 'ParkingPermitFee',
                    caption: 'Parking Permit <br>Fee',
                    size: '100px',
                    render: 'money'
                },
                {
                    field: 'DtStart',
                    caption: 'DtStart',
                    size: '100px'
                },
                {
                    field: 'DtStop',
                    caption: 'DtStop',
                    size: '100px'
                }
            ],
            onChange: function (event) {
                event.onComplete = function () {
                    this.save();
                };
            },
            onRefresh: function(event) {
                // have to manage recid on every refresh of this grid
                event.onComplete = function() {
                    for (var j = 0; j < w2ui.RAVehiclesGrid.records.length; j++) {
                        w2ui.RAVehiclesGrid.records[j].recid = j + 1;
                    }
                };
            },
            onClick : function (event){
                event.onComplete = function () {
                    var yes_args = [this, event.recid],
                        no_args = [this],
                        no_callBack = function(grid) {
                            grid.select(app.last.grid_sel_recid);
                            return false;
                        },
                        yes_callBack = function (grid, recid) {
                            app.last.grid_sel_recid = parseInt(recid);

                            // keep highlighting current row in any case
                            grid.select(app.last.grid_sel_recid);

                            w2ui.RAVehicleForm.record = $.extend(true, {}, grid.get(app.last.grid_sel_recid));
                            showSliderContentW2UIComp(w2ui.RAVehicleForm, RACompConfig.vehicles.sliderWidth);
                            w2ui.RAVehicleForm.refresh();

                        };

                    // warn user if form content has been changed
                    form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
                };
            },
            onAdd   : function (/*event*/) {
                var yes_args = [this],
                    no_callBack = function() {
                        return false;
                    },
                    yes_callBack = function(grid) {
                        app.last.grid_sel_recid = -1;
                        grid.selectNone();

                        var BID = getCurrentBID(),
                            BUD = getBUDfromBID(BID);

                        w2ui.RAVehicleForm.record = getVehicleFormInitalRecord(BID, BUD, null);
                        w2ui.RAVehicleForm.record.recid = w2ui.RAVehiclesGrid.records.length + 1;
                        showSliderContentW2UIComp(w2ui.RAVehicleForm, RACompConfig.vehicles.sliderWidth);
                        w2ui.RAVehicleForm.refresh();
                    };

                // warn user if form content has been changed
                form_dirty_alert(yes_callBack, no_callBack, yes_args);
            }
        });
    }

    // now load grid in target division
    $('#ra-form #vehicles .grid-container').w2render(w2ui.RAVehiclesGrid);

    // load the existing data in vehicles component
    setTimeout(function () {
        var compData = getRAFlowCompData("vehicles", app.raflow.activeFlowID);
        var grid = w2ui.RAVehiclesGrid;

        if (compData) {
            grid.records = compData;
            reassignGridRecids(grid.name);

            // lock the grid until "Have vehicles?" checkbox checked.
            lockOnGrid(grid.name);
        } else {
            grid.clear();
        }
    }, 500);
};
