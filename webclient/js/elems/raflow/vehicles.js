/* global
    RACompConfig, reassignGridRecids,
    hideSliderContent, showSliderContentW2UIComp,
    saveActiveCompData, getRAFlowCompData,
    lockOnGrid,
    getVehicleFormInitRecord, setVehicleLocalData, getVehicleLocalData,
    AssignVehiclesGridRecords, saveVehiclesCompData,
    SetRAVehicleLayoutContent,
    getVehicleFeeLocalData, setVehicleFeeLocalData,
    AssignVehicleFeesGridRecords,
    SetRAVehicleFormRecordFromLocalData,
    SetlocalDataFromRAVehicleFormRecord,
    getAllARsWithAmount
*/

"use strict";

// -------------------------------------------------------------------------------
// Rental Agreement - Vehicles Grid
// -------------------------------------------------------------------------------
window.getVehicleFormInitRecord = function (previousFormRecord) {
    var BID = getCurrentBID();

    var t = new Date(),
        nyd = new Date(new Date().setFullYear(new Date().getFullYear() + 1));

    var defaultFormData = {
        recid:                  w2ui.RAVehiclesGrid.records.length + 1,
        TMPVID:                 0,
        VID:                    0,
        BID:                    BID,
        TMPTCID:                0,
        VIN:                    "",
        VehicleType:            "",
        VehicleMake:            "",
        VehicleModel:           "",
        VehicleColor:           "",
        VehicleYear:            0,
        LicensePlateState:      "",
        LicensePlateNumber:     "",
        ParkingPermitNumber:    "",
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

// -------------------------------------------------------------
// SetlocalDataFromRAVehicleFormRecord
// ==================================
// will update the data from the record
// it will only update the field defined in fields list in
// form definition
// -------------------------------------------------------------
window.SetlocalDataFromRAVehicleFormRecord = function(TMPVID) {
    var form            = w2ui.RAVehicleForm,
        fields          = form.fields || [],
        vehicleFormData = getFormSubmitData(form.record, true);

    // local vehicle data
    var localVehicleData;

    if (TMPVID === 0) {
        localVehicleData = vehicleFormData;
    } else {
        // get data from form field's TMPVID
        localVehicleData = getVehicleLocalData(vehicleFormData.TMPVID);

        // loop over form fields
        fields.forEach(function(fieldItem) {
            localVehicleData[fieldItem.field] = vehicleFormData[fieldItem.field];
        });
    }

    // if not Fees then assign in vehicle data
    if (!localVehicleData.hasOwnProperty("Fees")) {
        localVehicleData.Fees = [];
    }
    localVehicleData.Fees = w2ui.RAVehicleFeesGrid.records;

    // set this modified data back
    setVehicleLocalData(TMPVID, localVehicleData);
};

// -------------------------------------------------------------
// SetRAVehicleFormRecordFromLocalData
// ================================
// will set the data in the form record
// from local vehicle data
// -------------------------------------------------------------
window.SetRAVehicleFormRecordFromLocalData = function(TMPVID) {
    var form        = w2ui.RAVehicleForm,
        fields      = form.fields || [];

    if (TMPVID === 0) {
        form.record = getVehicleFormInitRecord(null);
    } else {
        // get data from form field's TMPVID
        var localVehicleData = getVehicleLocalData(TMPVID);

        // loop over form fields
        fields.forEach(function(fieldItem) {
            form.record[fieldItem.field] = localVehicleData[fieldItem.field];
        });
    }

    // refresh the form after setting the record
    form.refresh();
    form.refresh();
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
                { field: 'DtStart',             type: 'date',   required: false,    html: { caption: 'DtStart', page: 0, column: 0 } },
                { field: 'DtStop',              type: 'date',   required: false,    html: { caption: 'DtStop', page: 0, column: 0 } },
                { field: 'LastModTime',         type: 'time',   required: false,    html: { caption: 'LastModTime', page: 0, column: 0 } },
                { field: 'LastModBy',           type: 'int',    required: false,    html: { caption: 'LastModBy', page: 0, column: 0 } },
            ],
            actions: {
                reset: function () {
                    w2ui.RAVehicleForm.clear();
                }
            },
            onRefresh: function(event) {
                event.onComplete = function() {
                    var f = w2ui.RAVehicleForm,
                        header = "Edit Rental Agreement Vehicles ({0})";

                    // there is NO VID actually, so have to work around with TMPVID key
                    formRefreshCallBack(f, "TMPVID", header);

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
                    if (f.record.TMPVID === 0) {
                        $(f.box).find("button[name=delete]").addClass("hidden");
                    } else {
                        $(f.box).find("button[name=delete]").removeClass("hidden");
                    }

                    // get RAID for active flow
                    var RAID = app.raflow.data[app.raflow.activeFlowID].ID;
                    if (RAID > 0) {
                        $(f.box).find("input[name=ParkingPermitNumber]").prop("disabled", false);
                    } else {
                        // if RAID is not available then disable
                        $(f.box).find("input[name=ParkingPermitNumber]").prop("disabled", true);
                    }
                };
            },
            onChange: function(event) {
                event.onComplete = function() {
                    // formRecDiffer: 1=current record, 2=original record, 3=diff object
                    var diff = formRecDiffer(this.record, app.active_form_original, {});
                    // if diff == {} then make dirty flag as false, else true
                    if ($.isPlainObject(diff) && $.isEmptyObject(diff)) {
                        app.form_is_dirty = false;
                    } else {
                        app.form_is_dirty = true;
                    }
                };
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
                    field: 'TMPVID',
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
                {
                    field: 'VehicleType',
                    caption: 'Type',
                    size: '80px',
                    editable: { type: 'text' }
                },
                {
                    field: 'VIN',
                    caption: 'VIN',
                    size: '80px'
                },
                {
                    field: 'VehicleMake',
                    caption: 'Make',
                    size: '80px'
                },
                {
                    field: 'VehicleModel',
                    caption: 'Model',
                    size: '80px'
                },
                {
                    field: 'VehicleColor',
                    caption: 'Color',
                    size: '80px'
                },
                {
                    field: 'VehicleYear',
                    caption: 'Year',
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
                    $("#RAVehiclesGrid_checkbox")[0].checked = app.raflow.data[app.raflow.activeFlowID].Data.meta.HaveVehicles;
                    $("#RAVehiclesGrid_checkbox")[0].disabled = app.raflow.data[app.raflow.activeFlowID].Data.meta.HaveVehicles;
                    lockOnGrid("RAVehiclesGrid");
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

                            // get TMPVID from grid
                            var TMPVID = grid.get(app.last.grid_sel_recid).TMPVID;

                            // render layout in the slider
                            showSliderContentW2UIComp(w2ui.RAVehicleLayout, RACompConfig.vehicles.sliderWidth);

                            // load pet fees grid
                            setTimeout(function() {
                                // fill layout with components and with data
                                SetRAVehicleLayoutContent(TMPVID);
                            }, 0);
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

                        // render the layout in slider
                        showSliderContentW2UIComp(w2ui.RAVehicleLayout, RACompConfig.vehicles.sliderWidth);

                        // load pet fees grid
                        setTimeout(function() {
                            // fill layout with components
                            SetRAVehicleLayoutContent(0);
                        }, 0);
                    };

                // warn user if form content has been changed
                form_dirty_alert(yes_callBack, no_callBack, yes_args);
            }
        });

        // vehicle fees grid
        $().w2grid({
            name: 'RAVehicleFeesGrid',
            header: 'Vehicle Fees',
            show: {
                toolbar:        true,
                header:         false,
                toolbarSearch:  false,
                toolbarAdd:     true,
                toolbarReload:  false,
                toolbarInput:   false,
                toolbarColumns: true,
                footer:         false,
            },
            multiSelect: false,
            style: 'border: 1px solid silver;',
            columns: [
                {
                    field: 'recid',
                    hidden: true
                },
                {
                    field: 'TMPVID',
                    caption: 'TMPVID',
                    hidden: true
                },
                {
                    field: 'BID',
                    caption: 'BID',
                    hidden: true
                },
                {
                    field: 'ASMID',
                    caption: 'ASMID',
                    hidden: true
                },
                {
                    field: 'ARID',
                    caption: 'ARID',
                    hidden: true
                },
                {
                    field: 'ARName',
                    caption: 'Name',
                    size: '70%',
                    editable: {
                        type: 'select',
                        items: [],
                    },
                    render: function (record/*, index, col_index*/) {
                        var html = '';

                        if (record) {
                            var items = app.raflow.arW2UIItems;
                            for (var s in items) {
                                if (items[s].id == record.ARID) html = items[s].text;
                            }
                        }
                        return html;
                    }
                },
                {
                    field: 'Amount',
                    caption: 'Amount',
                    size: '100px',
                    render: 'money',
                    editable: {
                        type: 'money',
                    }
                },
                {
                    field: 'RemoveRec',
                    caption: "Remove Vehicle Fee",
                    size: '100px',
                    style: 'text-align: center;',
                    render: function (record/*, index, col_index*/) {
                        var html = "";
                        if (record && record.ARID > -1) {
                            html = '<i class="fas fa-minus-circle" style="color: #DC3545; cursor: pointer;" title="remove rentable"></i>';
                        }
                        return html;
                    },
                }
            ],
            onChange: function (event) {
                var grid = this;
                event.onComplete = function () {
                    var BID = getCurrentBID(),
                        record = grid.get(event.recid),
                        localVehicleFeeData = getVehicleFeeLocalData(record.TMPVID, record.ARID);

                    switch(event.column) {
                    case grid.getColumn("Amount", true):
                        // update data in local and grid record
                        localVehicleFeeData.Amount = record.Amount = parseFloat(event.value_new);
                        // set data
                        grid.set(event.recid, record);
                        setVehicleFeeLocalData(record.TMPVID, record.ARID, localVehicleFeeData);
                        break;
                    case grid.getColumn("ARName", true):
                        var arItem = {};
                        // get aritem
                        app.raflow.arList[BID].forEach(function(item) {
                            if (parseInt(event.value_new) == item.ARID) {
                                arItem = item;
                                return false;
                            }
                        });

                        // change the values
                        localVehicleFeeData.Amount = record.Amount = arItem.DefaultAmount;
                        localVehicleFeeData.ARID = record.ARID = parseInt(event.value_new);
                        // set data
                        grid.set(event.recid, record);
                        // grid.getColumn("ARName").render();
                        setVehicleFeeLocalData(record.TMPVID, record.ARID, localVehicleFeeData);
                        break;
                    }

                    w2ui.RAVehicleFeesGrid.save();
                    // TODO(Sudip): we still need to update data locally
                };
            },
            onClick: function(event) {
                event.onComplete = function() {
                    // if it's remove column then remove the record
                    // maybe confirm dialog will be added
                    if(w2ui.RAVehicleFeesGrid.getColumn("RemoveRec", true) == event.column) {

                        // remove entry from local data
                        var rec = w2ui.RAVehicleFeesGrid.get(event.recid);
                        var vehicleData = getVehicleLocalData(rec.TMPVID);

                        if (vehicleData.hasOwnProperty("Fees")) {
                            var feeIndex = getVehicleFeeLocalData(rec.TMPVID, rec.ARID, true);

                            // also manage local data
                            vehicleData.Fees.splice(feeIndex, 1);

                            // set modified vehicleData back in local
                            setVehicleLocalData(rec.TMPVID, vehicleData);

                            // save the data on server data
                            saveVehiclesCompData()
                            .done(function(data) {
                                if (data.status === "success") {
                                    // remove from grid
                                    w2ui.RAVehicleFeesGrid.remove(event.recid);
                                }
                            });
                        } else {
                            // simple remove record from grid
                            w2ui.RAVehicleFeesGrid.remove(event.recid);
                        }

                        return;
                    }
                };
            },
            onAdd: function(/*event*/) {
                var grid = w2ui.RAVehicleFeesGrid;
                var rec     = {
                    recid:  grid.records.length + 1,
                    TMPVID: 0,
                    BID:    0,
                    ASMID:  0,
                    ARID:   0,
                    ARName: "",
                    Amount: 0.0,
                };
                grid.add(rec);
                grid.select(rec.recid);
                grid.getColumn("ARName").editable.items = app.raflow.arW2UIItems;
                grid.getColumn("ARName").render();
                grid.refresh();
            },
        });

        //------------------------------------------------------------------------
        //          Vehicle Form Buttons
        //------------------------------------------------------------------------
        $().w2form({
            name: 'RAVehicleFormBtns',
            style: 'border: none; background-color: transparent;',
            formURL: '/webclient/html/formravehiclebtns.html',
            url: '',
            fields: [],
            actions: {
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
                            f.record = getVehicleFormInitRecord(f.record);
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

        //------------------------------------------------------------------------
        //  vehicleLayout - The layout to contain the vehicleForm and vehicleFees grid
        //              top  -      vehicleForm
        //              main -      vehicleFeesGrid
        //              bottom -    action buttons form
        //------------------------------------------------------------------------
        $().w2layout({
            name: 'RAVehicleLayout',
            padding: 0,
            panels: [
                { type: 'left',    size: 0,     hidden: true },
                { type: 'top',     size: '60%', hidden: false, content: 'top',  resizable: true, style: app.pstyle },
                { type: 'main',    size: '40%', hidden: false, content: 'main', resizable: true, style: app.pstyle },
                { type: 'preview', size: 0,     hidden: true,  content: 'PREVIEW'  },
                { type: 'bottom',  size: 50,    hidden: false, content: 'bottom', resizable: false, style: app.pstyle },
                { type: 'right',   size: 0,     hidden: true }
            ]
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

// fill rental agreement vehicle layout with all forms, grids
window.SetRAVehicleLayoutContent = function(TMPVID) {
    w2ui.RAVehicleLayout.content('bottom',  w2ui.RAVehicleFormBtns);
    w2ui.RAVehicleLayout.content('top',     w2ui.RAVehicleForm);
    w2ui.RAVehicleLayout.content('main',    w2ui.RAVehicleFeesGrid);

    // after 0 ms set the record
    setTimeout(function() {
        // set vehicle form record
        SetRAVehicleFormRecordFromLocalData(TMPVID);

        // assign vehicle fees grid
        var BID = getCurrentBID();
        getAllARsWithAmount(BID)
        .done(function() {
            AssignVehicleFeesGridRecords(TMPVID);
        });
    }, 0);
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

    // clear the grid
    grid.clear();

    compData.forEach(function(vehicleData) {
        var gridRec = {};

        // for each grid column
        grid.columns.forEach(function(gridColumn) {
            gridRec[gridColumn.field] = vehicleData[gridColumn.field];
        });

        // push the record in grid
        grid.records.push(gridRec);

        // assign record in grid
        reassignGridRecids(grid.name);
    });
};

//------------------------------------------------------------------------------
// saveVehiclesCompData - saves the data on server side
//------------------------------------------------------------------------------
window.saveVehiclesCompData = function() {
    var compData = getRAFlowCompData("vehicles", app.raflow.activeFlowID);
    return saveActiveCompData(compData, "vehicles");
};

//-----------------------------------------------------------------------------
// getVehicleFeeLocalData - returns the clone of vehicle fee data for requested
//                      TMPVID and ARID
//-----------------------------------------------------------------------------
window.getVehicleFeeLocalData = function(TMPVID, ARID, returnIndex) {
    var cloneData = {};
    var foundIndex = -1;
    var compData = getRAFlowCompData("vehicles", app.raflow.activeFlowID) || [];
    compData.forEach(function(item, index) {
        if (item.TMPVID == TMPVID) {
            var feesData = item.Fees || [];
            feesData.forEach(function(feeItem, index) {
                if (feeItem.ARID == ARID) {
                    if (returnIndex) {
                        foundIndex = index;
                    } else {
                        cloneData = $.extend(true, {}, feeItem);
                    }
                }
            });
            return false;
        }
    });
    if (returnIndex) {
        return foundIndex;
    }
    return cloneData;
};


//-----------------------------------------------------------------------------
// setVehicleFeeLocalData - save the data for requested a TMPVID, ARID
//                   in local data
//-----------------------------------------------------------------------------
window.setVehicleFeeLocalData = function(TMPVID, ARID, vehicleFeeData) {
    var compData = getRAFlowCompData("vehicles", app.raflow.activeFlowID);
    var pIndex = -1,
        fIndex = -1;

    compData.forEach(function(item, itemIndex) {
        if (item.TMPVID == TMPVID) {
            var feesData = item.Fees || [];
            feesData.forEach(function(feeItem, feeItemIndex) {
                if (feeItem.ARID == ARID) {
                    fIndex = feeItemIndex;
                }
                return false;
            });
            pIndex = itemIndex;
            return false;
        }
    });

    // only if rentable found then
    if (pIndex > -1) {
        if (fIndex > -1) {
            compData[pIndex].Fees[fIndex] = vehicleFeeData;
        } else {
            compData[pIndex].Fees.push(vehicleFeeData);
        }
    }
};

//-----------------------------------------------------------------------------
// AssignVehicleFeesGridRecords - will set the vehicle fees grid records from local
//                            copy of vehicle fees data again
//-----------------------------------------------------------------------------
window.AssignVehicleFeesGridRecords = function(TMPVID) {
    var grid    = w2ui.RAVehicleFeesGrid,
        BID     = getCurrentBID();

    // clear the grid
    grid.clear();

    // list of fees
    var vehicleFeesData = [];

    // SPECIAL CASE if adding new vehicle //
    if (TMPVID === 0) {
        // get initial fees from bizProps
        app.vehicleFees[BID].forEach(function(bizPropVehicleFee) {
            vehicleFeesData.push($.extend(true, {TMPVID: 0, ASMID: 0}, bizPropVehicleFee));
        });
    } else {
        var vehicleData = getVehicleLocalData(TMPVID);
        vehicleFeesData = vehicleData.Fees;
    }

    // vehicle fees data
    vehicleFeesData.forEach(function(fee) {
        var gridRec = {};

        // for each grid column
        grid.columns.forEach(function(gridColumn) {
            gridRec[gridColumn.field] = fee[gridColumn.field];
        });

        // push the record in grid
        grid.records.push(gridRec);

        // assign recid again
        reassignGridRecids(grid.name);

        // assign item prepared earlier for parent rentable list
        grid.getColumn("ARName").editable.items = app.raflow.arW2UIItems;
        grid.getColumn("ARName").render();
    });
};

