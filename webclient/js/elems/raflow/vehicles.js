/* global
    RACompConfig, sliderContentDivLength, reassignGridRecids,
    hideSliderContent, appendNewSlider, showSliderContentW2UIComp,
    loadTargetSection, requiredFieldsFulFilled, getRAFlowPartTypeIndex, initRAFlowAjax,
    getRAFlowAllParts, saveActiveCompData, toggleHaveCheckBoxDisablity, getRAFlowCompData,
    lockOnGrid,
    getVehicleFormInitalRecord, setVehicleLocalData, getVehicleLocalData,
    AssignVehiclesGridRecords, saveVehiclesCompData
*/

"use strict";

// -------------------------------------------------------------------------------
// Rental Agreement - Vehicles Grid
// -------------------------------------------------------------------------------
window.getVehicleFormInitalRecord = function (previousFormRecord) {
    var BID = getCurrentBID();

    var t = new Date(),
        nyd = new Date(new Date().setFullYear(new Date().getFullYear() + 1));

    var defaultFormData = {
        recid:                   0,
        TMPVID:                  0,
        VID:                     0,
        BID:                   BID,
        TMPTCID:                 0,
        VIN:                    "",
        VehicleType:            "",
        VehicleMake:            "",
        VehicleModel:           "",
        VehicleColor:           "",
        VehicleYear:             0,
        LicensePlateState:      "",
        LicensePlateNumber:     "",
        ParkingPermitNumber:    "",
        ParkingPermitFee:        0,
        DtStart:                w2uiDateControlString(t),
        DtStop:                 w2uiDateControlString(nyd)
    };

    // if it called after 'save and add another' action there previous form record is passed as Object
    // else it is null
    if ( previousFormRecord ) {
        defaultFormData = setDefaultFormFieldAsPreviousRecord(
            ['*'], // Fields to Reset
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
                { field: 'recid',               type: 'int',    required: false,    html: { caption: 'recid', page: 0, column: 0 } },
                { field: 'TMPVID',              type: 'int',    required: true  },
                { field: 'TMPTCID',             type: 'list',   required: true,     options: {items: [], selected: {}} },
                { field: 'VehicleType',         type: 'text',   required: true  },
                { field: 'VehicleMake',         type: 'text',   required: true  },
                { field: 'VehicleModel',        type: 'text',   required: true  },
                { field: 'VehicleColor',        type: 'text',   required: false },
                { field: 'VehicleYear',         type: 'int',    required: false },
                { field: 'LicensePlateState',   type: 'text',   required: false },
                { field: 'LicensePlateNumber',  type: 'text',   required: false },
                { field: 'VIN',                 type: 'text',   required: false },
                { field: 'ParkingPermitNumber', type: 'text',   required: false },
                { field: 'ParkingPermitFee',    type: 'money',  required: false },
                { field: 'DtStart',             type: 'date',   required: false,    html: { caption: 'DtStart', page: 0, column: 0 } },
                { field: 'DtStop',              type: 'date',   required: false,    html: { caption: 'DtStop', page: 0, column: 0 } },
                { field: 'LastModTime',         type: 'time',   required: false,    html: { caption: 'LastModTime', page: 0, column: 0 } },
                { field: 'LastModBy',           type: 'int',    required: false,    html: { caption: 'LastModBy', page: 0, column: 0 } },
            ],
            onRefresh: function(event) {
                event.onComplete = function() {
                    var f = w2ui.RAVehicleForm,
                        header = "Edit Rental Agreement Vehicles ({0})";

                    // there is NO VID actually, so have to work around with recid key
                    formRefreshCallBack(f, "recid", header);

                    // selection of contact person
                    var TMPTCIDSel = {};
                    app.raflow.peopleW2UIItems.forEach(function(item) {
                        if (item.id === f.record.TMPTCID) {
                            $.extend(TMPTCIDSel, item);
                        }
                    });
                    f.get("TMPTCID").options.items = app.raflow.peopleW2UIItems;
                    f.get("TMPTCID").options.selected = TMPTCIDSel;

                    // hide delete button if it is NewRecord
                    var isNewRecord = (w2ui.RAVehiclesGrid.get(f.record.recid, true) === null);
                    if (isNewRecord) {
                        $(f.box).find("button[name=delete]").addClass("hidden");
                    } else {
                        $(f.box).find("button[name=delete]").removeClass("hidden");
                    }

                    // get RAID for active flow
                    var RAID = app.raflow.data[app.raflow.activeFlowID].ID;
                    console.debug("RAID", RAID);
                    if (RAID > 0) {
                        $(f.box).find("input[name=ParkingPermitNumber]").prop("disabled", false);
                        // $(f.box).find("input[name=ParkingPermitFee]").prop("disabled", false);
                    } else {
                        // if RAID is not available then disable
                        $(f.box).find("input[name=ParkingPermitNumber]").prop("disabled", true);
                        // $(f.box).find("input[name=ParkingPermitFee]").prop("disabled", true);
                    }
                };
            },
            actions: {
                reset: function () {
                    w2ui.RAVehicleForm.clear();
                },
                save: function () {
                    var f =     w2ui.RAVehicleForm,
                        grid =  w2ui.RAVehiclesGrid,
                        TMPVID = f.record.TMPVID;

                    // validate form record
                    var errors = f.validate();
                    if (errors.length > 0) return;

                    // sync this info in local data
                    var vehicleData = getFormSubmitData(f.record, true);

                    // set data locally
                    setVehicleLocalData(TMPVID, vehicleData);

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // save this records in json Data
                    saveVehiclesCompData()
                    .done(function(data) {
                        if (data.status === 'success') {
                            // re-assign records in grid
                            AssignVehiclesGridRecords();

                            // reset the form
                            f.actions.reset();

                            // Disable "have vehicles?" checkbox if there is any record.
                            toggleHaveCheckBoxDisablity('RAVehiclesGrid');

                            // close the form
                            hideSliderContent();
                        } else {
                            f.message(data.message);
                        }
                    })
                    .fail(function(data) {
                        console.log("failure " + data);
                    });
                },
                saveadd: function () {
                    var f =     w2ui.RAVehicleForm,
                        grid =  w2ui.RAVehiclesGrid,
                        TMPVID = f.record.TMPVID;

                    // validate the form first
                    var errors = f.validate();
                    if (errors.length > 0) return;


                    // sync this info in local data
                    var vehicleData = getFormSubmitData(f.record, true);

                    // set data locally
                    setVehicleLocalData(TMPVID, vehicleData);

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // save this records in json Data
                    saveVehiclesCompData()
                    .done(function(data) {
                        if (data.status === 'success') {
                            // reset form
                            f.actions.reset();
                            f.record = getVehicleFormInitalRecord(f.record);
                            f.record.recid =grid.records.length + 1;
                            f.refresh();
                            f.refresh();

                            // re-assign records in grid
                            AssignVehiclesGridRecords();
                        } else {
                            f.message(data.message);
                        }
                    })
                    .fail(function(data) {
                        console.log("failure " + data);
                    });
                },
                delete: function () {
                    var f = w2ui.RAVehicleForm;

                    // get local data from TMPVID
                    var compData = getRAFlowCompData("vehicles", app.raflow.activeFlowID) || [];
                    var itemIndex = getVehicleLocalData(f.record.TMPVID, true);
                    compData.splice(itemIndex, 1);

                    // save this records in json Data
                    saveVehiclesCompData()
                    .done(function(data) {
                        if (data.status === 'success') {
                            // reset form
                            f.actions.reset();

                            // Disable "have vehicles?" checkbox if there is any record.
                            toggleHaveCheckBoxDisablity('RAVehiclesGrid');

                            // re-assign records in grid
                            AssignVehiclesGridRecords();

                            // close the form
                            hideSliderContent();
                        } else {
                            f.message(data.message);
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
                {field: 'recid', hidden: true },
                {field: 'TMPVID', hidden: true },
                {field: 'VID', hidden: true },
                {field: 'BID', hidden: true },
                {
                    field: 'TMPTCID',
                    caption: 'Contact<br>Person',
                    size: '150px',
                    render: function (record/*, index, col_index*/) {
                        var html = '';
                        if (record) {
                            var items = app.raflow.peopleW2UIItems;
                            for (var s in items) {
                                if (items[s].id == record.TMPTCID) html = items[s].text;
                            }
                        }
                        return html;
                    }
                },
                {field: 'VehicleType',         caption: 'Type',    size: '80px', editable: {type: 'text'} },
                {field: 'VIN',                 caption: 'VIN',     size: '80px'},
                {field: 'VehicleMake',         caption: 'Make',    size: '80px'},
                {field: 'VehicleModel',        caption: 'Model',   size: '80px'},
                {field: 'VehicleColor',        caption: 'Color',   size: '80px'},
                {field: 'VehicleYear',         caption: 'Year',    size: '80px'},
                {field: 'LicensePlateState',   caption: 'License Plate<br>State',    size: '100px'},
                {field: 'LicensePlateNumber',  caption: 'License Plate<br>Number',   size: '100px'},
                {field: 'ParkingPermitNumber', caption: 'Parking Permit <br>Number', size: '100px'},
                {field: 'ParkingPermitFee',    caption: 'Parking Permit <br>Fee',    size: '100px', render: 'money'},
                {field: 'DtStart',             caption: 'DtStart', size: '100px'},
                {field: 'DtStop',              caption: 'DtStop',  size: '100px'}
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

                        w2ui.RAVehicleForm.record = getVehicleFormInitalRecord(null);
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
        // assign grid records
        AssignVehiclesGridRecords();
    }, 500);
};

//-----------------------------------------------------------------------------
// getVehicleLocalData - returns the clone of vehicle data for requested TMPVID
//-----------------------------------------------------------------------------
window.getVehicleLocalData = function(TMPVID, returnIndex) {
    var cloneData = {};
    var foundIndex = -1;
    var compData = getRAFlowCompData("vehicles", app.raflow.activeFlowID) || [];
    compData.forEach(function(item, index) {
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
// setVehicleLocalData - save the data for requested a TMPVID in local data
//-----------------------------------------------------------------------------
window.setVehicleLocalData = function(TMPVID, vehicleData) {
    var compData = getRAFlowCompData("vehicles", app.raflow.activeFlowID) || [];
    var dataIndex = -1;
    compData.forEach(function(item, index) {
        if (item.TMPVID == TMPVID) {
            dataIndex = index;
            return false;
        }
    });
    if (dataIndex > -1) {
        compData[dataIndex] = vehicleData;
    } else {
        compData.push(vehicleData);
    }
};

//-----------------------------------------------------------------------------
// AssignVehiclesGridRecords - will set the vehicles grid records from local
//                               copy of flow data again
//-----------------------------------------------------------------------------
window.AssignVehiclesGridRecords = function() {
    var compData = getRAFlowCompData("vehicles", app.raflow.activeFlowID);
    var grid = w2ui.RAVehiclesGrid;

    // reset last sel recid
    app.last.grid_sel_recid  =-1;
    // select none
    grid.selectNone();

    if (compData) {
        grid.records = compData;
        reassignGridRecids(grid.name);

        // lock the grid until "Have vehicles?" checkbox checked.
        lockOnGrid(grid.name);

        // Operation on RAVehicleForm
        w2ui.RAVehicleForm.refresh();
    } else {
        // clear the grid
        grid.clear();
        // Operation on RAVehicleForm
        w2ui.RAVehicleForm.actions.reset();
    }
};

//------------------------------------------------------------------------------
// saveVehiclesCompData - saves the data on server side
//------------------------------------------------------------------------------
window.saveVehiclesCompData = function() {
    var compData = getRAFlowCompData("vehicles", app.raflow.activeFlowID);
    return saveActiveCompData(compData, "vehicles");
};

