/* global
    RACompConfig, reassignGridRecids,
    HideSliderContent, ShowSliderContentW2UIComp,
    saveActiveCompData, getRAFlowCompData,
    lockOnGrid, displayRAVehicleFeesGridError,
    GetVehicleFormInitRecord, SetVehicleLocalData, GetVehicleLocalData,
    AssignVehiclesGridRecords, SaveVehiclesCompData,
    SetRAVehicleLayoutContent,
    GetVehicleFeeLocalData, SetVehicleFeeLocalData,
    AssignVehicleFeesGridRecords,
    SetRAVehicleFormRecordFromLocalData,
    SetlocalDataFromRAVehicleFormRecord,
    GetAllARForFeeForm, SetDataFromFormRecord, SetFormRecordFromData,
    GetFeeGridColumns, GetFeeFormFields, GetFeeFormToolbar,
    SetFeeDataFromFeeFormRecord,
    GetFeeFormInitRecord, getRecIDFromTMPASMID, getFeeIndex,
    FeeFormOnChangeHandler, FeeFormOnRefreshHandler,
    SliderContentDivLength, SetFeeFormRecordFromFeeData,
    displayRAVehicleFeeFormError, RenderFeesGridSummary, displayRAVehicleFormError,
    displayFormFieldsError, getVehicleIndex,
    RenderVehicleFeesGridSummary, RAFlowNewVehicleAJAX,
    GetFeeAccountRulesW2UIListItems, RenderFeesGridSummary,
    GetVehicleIdentity, updateFlowData, GetTiePeopleLocalData,
    getRecIDFromTMPVID, dispalyRAVehiclesGridError, GetCurrentFlowID,
    EnableDisableRAFlowVersionInputs, ShowHideGridToolbarAddButton,
    HideAllSliderContent
*/

"use strict";

//-----------------------------------------------------------------------------
// RAFlowNewVehicleAJAX - Request to create new vehicle in raflow json
//-----------------------------------------------------------------------------
window.RAFlowNewVehicleAJAX = function() {
    var BID = getCurrentBID();
    var FlowID = GetCurrentFlowID();
    var data = {"cmd": "new", "FlowID": FlowID};

    return $.ajax({
        url: '/v1/raflow-vehicles/' + BID.toString() + "/" + FlowID.toString() + "/",
        method: "POST",
        data: JSON.stringify(data),
        contentType: "application/json",
        dataType: "json"
    })
    .done(function(data) {
        if (data.status === "success") {
            // Update flow local copy and green checks
            updateFlowData(data);

            // reassign records
            AssignVehiclesGridRecords();

            // mark new TMPVID from meta
            app.raflow.last.TMPVID = data.record.Flow.Data.meta.LastTMPVID;
        }
    });
};

// -------------------------------------------------------------------------------
// Rental Agreement - Vehicles Grid
// -------------------------------------------------------------------------------
window.GetVehicleFormInitRecord = function (previousFormRecord) {
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
    var form            = w2ui.RAVehicleForm;

    // get data from form field's TMPVID
    var localVehicleData = GetVehicleLocalData(TMPVID);

    // set data from form
    var vehicleData = SetDataFromFormRecord(TMPVID, form, localVehicleData);

    // KEEP VEHICLE YEAR IN NUMBERIC
    vehicleData.VehicleYear = parseInt(vehicleData.VehicleYear);

    // if not Fees then assign in vehicle data
    if (!vehicleData.hasOwnProperty("Fees")) {
        vehicleData.Fees = [];
    }
    vehicleData.Fees = w2ui.RAVehicleFeesGrid.records;

    // set this modified data back
    SetVehicleLocalData(TMPVID, localVehicleData);
};

// -------------------------------------------------------------
// SetRAVehicleFormRecordFromLocalData
// ================================
// will set the data in the form record
// from local vehicle data
// -------------------------------------------------------------
window.SetRAVehicleFormRecordFromLocalData = function(TMPVID) {
    var form = w2ui.RAVehicleForm;

    // get data from form field's TMPVID
    var localVehicleData = GetVehicleLocalData(TMPVID);

    // set form record from data
    SetFormRecordFromData(form, localVehicleData);

    // refresh the form after setting the record
    form.refresh();
    form.refresh();
};

window.loadRAVehiclesGrid = function () {

    // if form is loaded then return
    if (!("RAVehiclesGrid" in w2ui)) {

        // -----------------------------------------------------------
        //      VEHICLES GRID
        // -----------------------------------------------------------
        $().w2grid({
            name    : 'RAVehiclesGrid',
            header  : 'Vehicles',
            show    : {
                toolbar         : true,
                toolbarSearch   : false,
                toolbarReload   : false,
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
                    field: 'haveError',
                    size: '30px',
                    hidden: false,
                    render: function (record) {
                        var haveError = false;
                        if (app.raflow.validationErrors.vehicles) {
                            var vehicles = app.raflow.validationCheck.errors.vehicles;
                            for (var i = 0; i < vehicles.length; i++) {
                                if (vehicles[i].TMPVID === record.TMPVID && vehicles[i].total > 0) {
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
            onRefresh: function(event) {
                var grid = this;

                // have to manage recid on every refresh of this grid
                event.onComplete = function() {
                    $("#RAVehiclesGrid_checkbox")[0].checked = app.raflow.Flow.Data.meta.HaveVehicles;
                    $("#RAVehiclesGrid_checkbox")[0].disabled = app.raflow.Flow.Data.meta.HaveVehicles;
                    lockOnGrid("RAVehiclesGrid");

                    ShowHideGridToolbarAddButton(grid.name);
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
                            var TMPVID = grid.get(recid).TMPVID;

                            // keep this clicked TMPVID in last object
                            app.raflow.last.TMPVID = TMPVID;

                            // render layout in the slider
                            ShowSliderContentW2UIComp(w2ui.RAVehicleLayout, RACompConfig.vehicles.sliderWidth);

                            // load vehicle fees grid
                            setTimeout(function() {
                                // fill layout with components and with data
                                SetRAVehicleLayoutContent(TMPVID);
                            }, 0);

                            setTimeout(function () {
                                displayRAVehicleFormError();
                            }, 500);
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

                        // get new entry for vehicle
                        RAFlowNewVehicleAJAX()
                        .done(function(data) {
                            // get last clicked TMPVID
                            var TMPVID = app.raflow.last.TMPVID;

                            // render the layout in slider
                            ShowSliderContentW2UIComp(w2ui.RAVehicleLayout, RACompConfig.vehicles.sliderWidth);

                            // load vehicle fees grid
                            setTimeout(function() {
                                // fill layout with components
                                SetRAVehicleLayoutContent(TMPVID);
                            }, 0);
                        });
                    };

                // warn user if form content has been changed
                form_dirty_alert(yes_callBack, no_callBack, yes_args);
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
                { type: 'top',     size: '50%', hidden: false, content: 'top',  resizable: true, style: app.pstyle },
                { type: 'main',    size: '50%', hidden: false, content: 'main', resizable: true, style: app.pstyle },
                { type: 'preview', size: 0,     hidden: true,  content: 'PREVIEW'  },
                { type: 'bottom',  size: 50,    hidden: false, content: 'bottom', resizable: false, style: app.pstyle },
                { type: 'right',   size: 0,     hidden: true }
            ]
        });

        // -----------------------------------------------------------
        //      ***** VEHICLE FORM *****
        // -----------------------------------------------------------
        $().w2form({
            name    : 'RAVehicleForm',
            header  : 'Add Vehicle form',
            formURL : '/webclient/html/raflow/formra-vehicles.html',
            toolbar :{
                items: [
                    { id: 'bt3', type: 'spacer' },
                    { id: 'btnClose', type: 'button', icon: 'fas fa-times'}
                ],
                onClick: function (event) {
                    switch (event.target){
                        case 'btnClose':
                            HideSliderContent();
                            break;
                    }
                }
            },
            fields  : [
                { field: 'recid',               type: 'int',    required: false,     html: { caption: 'recid', page: 0, column: 0 } },
                { field: 'TMPVID',              type: 'int',    required: false  },
                { field: 'BID',                 type: 'int',    required: true,      html: { caption: 'BID', page: 0, column: 0 } },
                { field: 'VID',                 type: 'int',    required: false,     html: { caption: 'VID', page: 0, column: 0 } },
                { field: 'TMPTCID',             type: 'list',   required: false,     options: {items: [], selected: {}} },
                { field: 'VehicleType',         type: 'text',   required: true },
                { field: 'VehicleMake',         type: 'text',   required: false },
                { field: 'VehicleModel',        type: 'text',   required: false },
                { field: 'VehicleColor',        type: 'text',   required: false },
                { field: 'VehicleYear',         type: 'number', required: false },
                { field: 'LicensePlateState',   type: 'text',   required: false },
                { field: 'LicensePlateNumber',  type: 'text',   required: false },
                { field: 'VIN',                 type: 'text',   required: false },
                { field: 'ParkingPermitNumber', type: 'text',   required: false },
                { field: 'DtStart',             type: 'date',   required: false,    html: { caption: 'DtStart', page: 0, column: 0 } },
                { field: 'DtStop',              type: 'date',   required: false,    html: { caption: 'DtStop',  page: 0, column: 0 } },
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
                    var RAID = app.raflow.Flow.ID;
                    if (RAID > 0) {
                        $(f.box).find("input[name=ParkingPermitNumber]").prop("disabled", false);
                    } else {
                        // make it required false
                        f.get("ParkingPermitNumber").required = false;
                        // if RAID is not available then disable
                        $(f.box).find("input[name=ParkingPermitNumber]").prop("disabled", true);
                    }

                    // format header
                    var vehicleIdentity = GetVehicleIdentity(f.record),
                        vehicleString   = "<em>new</em>";

                    if (f.record.VID > 0) {
                        vehicleString = vehicleIdentity;
                    } else if (vehicleIdentity) {
                        vehicleString = "<em>new</em> - {0}".format(vehicleIdentity);
                    }
                    f.header = "Edit Vehicle (<strong>{0}</strong>)".format(vehicleString);

                    // FREEZE THE INPUTS IF VERSION IS RAID
                    EnableDisableRAFlowVersionInputs(f);
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

        //------------------------------------------------------------------------
        //      ***** VEHICLE ACTION FORM BUTTONS *****
        //------------------------------------------------------------------------
        $().w2form({
            name: 'RAVehicleFormBtns',
            style: 'border: none; background-color: transparent;',
            formURL: '/webclient/html/raflow/formra-vehiclebtns.html',
            url: '',
            fields: [],
            actions: {
                save: function () {
                    var f       = w2ui.RAVehicleForm,
                        TMPVID  = f.record.TMPVID;

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // validate form record
                    var errors = f.validate();
                    if (errors.length > 0) return;

                    // update the modified data
                    SetlocalDataFromRAVehicleFormRecord(TMPVID);

                    // save this records in json Data
                    SaveVehiclesCompData()
                    .done(function(data) {
                        if (data.status === 'success') {
                            // re-assign records in grid
                            AssignVehiclesGridRecords();

                            // reset the form
                            f.actions.reset();

                            // close the form
                            HideSliderContent();
                        } else {
                            f.message(data.message);
                        }
                    })
                    .fail(function(data) {
                        console.log("failure " + data);
                    });
                },
                saveadd: function () {
                    var f       = w2ui.RAVehicleForm,
                        grid    = w2ui.RAVehiclesGrid,
                        TMPVID  = f.record.TMPVID;

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // validate the form first
                    var errors = f.validate();
                    if (errors.length > 0) return;

                    // update local data from this form record
                    SetlocalDataFromRAVehicleFormRecord(TMPVID);

                    // save this records in json Data
                    SaveVehiclesCompData()
                    .done(function(data) {
                        if (data.status === 'success') {

                            // get new entry for vehicle
                            RAFlowNewVehicleAJAX()
                            .done(function(data) {
                                // IT'S MANAGED IN AJAX API
                                var TMPVID = app.raflow.last.TMPVID;

                                // reset form
                                f.actions.reset();
                                f.record = GetVehicleLocalData(TMPVID);
                                f.refresh();
                                f.refresh();

                                // re-assign records in grid
                                AssignVehiclesGridRecords();
                            })
                            .fail(function(data) {
                                f.message("failure " + data);
                            });
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
                    var compData = getRAFlowCompData("vehicles") || [];
                    var itemIndex = GetVehicleLocalData(f.record.TMPVID, true);

                    // if it exists then
                    if (itemIndex > -1) {

                        // remove locally
                        compData.splice(itemIndex, 1);

                        // save this records in json Data
                        SaveVehiclesCompData()
                        .done(function(data) {
                            if (data.status === 'success') {
                                // reset form
                                f.actions.reset();

                                // re-assign records in grid
                                AssignVehiclesGridRecords();

                                // close the form
                                HideSliderContent();
                            } else {
                                f.message(data.message);
                            }
                        })
                        .fail(function(data) {
                            console.log("failure " + data);
                        });
                    }
                },
            },
            onRefresh: function(event) {
                var form = this;
                event.onComplete = function() {
                    // FREEZE THE INPUTS IF VERSION IS RAID
                    EnableDisableRAFlowVersionInputs(form);
                };
            }
        });

        // -----------------------------------------------------------
        //      ***** VEHICLE ***** FEES ***** GRID *****
        // -----------------------------------------------------------
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
            columns: GetFeeGridColumns('RAVehicleFeesGrid'),
            onClick: function(event) {
                event.onComplete = function() {
                    var yes_args = [this, event.recid],
                        no_args = [this],
                        no_callBack = function(grid) {
                            grid.select(app.last.grid_sel_recid);
                            return false;
                        },
                        yes_callBack = function(grid, recid) {
                            var feeForm = w2ui.RAVehicleFeeForm;

                            var sliderID = 2;
                            appendNewSlider(sliderID);
                            $("#raflow-container")
                                .find(".slider[data-slider-id="+sliderID+"]")
                                .find(".slider-content")
                                .width(400)
                                .w2render(feeForm);

                            app.last.grid_sel_recid = parseInt(recid);

                            // keep highlighting current row in any case
                            grid.select(app.last.grid_sel_recid);

                            // get TMPVID from last of raflow
                            var TMPVID = app.raflow.last.TMPVID;

                            // get TMPASMID from grid record
                            var TMPASMID = grid.get(recid).TMPASMID;

                            // get all account rules then
                            var BID = getCurrentBID();
                            GetAllARForFeeForm(BID)
                            .done(function(data) {
                                // get filtered account rules items
                                feeForm.get("ARID").options.items = GetFeeAccountRulesW2UIListItems(BID, "vehicles");

                                // set record in form
                                SetFeeFormRecordFromFeeData(TMPVID, TMPASMID, "vehicles");
                                feeForm.record.RentCycleText = app.cycleFreq[feeForm.record.RentCycle];

                                ShowSliderContentW2UIComp(feeForm, SliderContentDivLength, sliderID);
                                feeForm.refresh(); // need to refresh for header changes

                                // When RentCycle is Norecur then disable the RentCycle list field.
                                var isDisabled = feeForm.record.RentCycleText.text === app.cycleFreq[0];
                                $("#RentCycleText").prop("disabled", isDisabled);

                                setTimeout(function () {
                                    displayRAVehicleFeeFormError(w2ui.RAVehicleForm.record.TMPVID);
                                }, 500);
                            })
                            .fail(function(data) {
                                console.log("failure" + data);
                            });
                        };

                    // warn user if form content has been changed
                    form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
                };
            },
            onAdd: function(/*event*/) {
                var feesGrid    = w2ui.RAVehicleFeesGrid,
                    feeForm     = w2ui.RAVehicleFeeForm;

                var sliderID = 2;
                appendNewSlider(sliderID);
                $("#raflow-container")
                    .find(".slider[data-slider-id="+sliderID+"]")
                    .find(".slider-content")
                    .width(400)
                    .w2render(feeForm);

                // new record so select none
                feesGrid.selectNone();

                var TMPVID = app.raflow.last.TMPVID;

                // get all account rules in fit those in form "ARID" field
                var BID = getCurrentBID();
                GetAllARForFeeForm(BID)
                .done(function(data) {
                    // get filtered account rules
                    feeForm.get("ARID").options.items = GetFeeAccountRulesW2UIListItems(BID, "vehicles");

                    // set form record
                    SetFeeFormRecordFromFeeData(TMPVID, 0, "vehicles");
                    feeForm.record.recid = feesGrid.records.length + 1;

                    // show form in the DOM
                    ShowSliderContentW2UIComp(feeForm, SliderContentDivLength, sliderID);
                    feeForm.refresh();
                })
                .fail(function(data) {
                    console.log("failure" + data);
                });
            },
            onRefresh: function(event) {
                var grid = this;
                event.onComplete = function() {
                    ShowHideGridToolbarAddButton(grid.name);
                };
            },
        });

        // -----------------------------------------------------------
        //      ***** VEHICLE ***** FEE ***** FORM *****
        // -----------------------------------------------------------
        $().w2form({
            name: 'RAVehicleFeeForm',
            header: 'Add New Vehicle Fee',
            style: 'display: block;',
            formURL: '/webclient/html/raflow/formra-fee.html',
            focus: -1,
            fields: GetFeeFormFields(),
            toolbar : GetFeeFormToolbar(),
            actions: {
                reset: function () {
                    w2ui.RAVehicleFeeForm.clear();
                },
                save: function() {
                    var feeForm     = w2ui.RAVehicleFeeForm,
                        TMPASMID    = feeForm.record.TMPASMID;

                    // get TMPVID from last of raflow
                    var TMPVID = app.raflow.last.TMPVID;

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // set local fee data from fee form
                    SetFeeDataFromFeeFormRecord(TMPVID, TMPASMID, "vehicles");

                    SaveVehiclesCompData()
                    .done(function (data) {
                        if (data.status === 'success') {
                            // Re render the fees grid records
                            AssignVehicleFeesGridRecords(TMPVID);

                            // reset the form
                            feeForm.actions.reset();

                            // close the form
                            HideSliderContent(2);
                        } else {
                            feeForm.message(data.message);
                        }
                    })
                    .fail(function (data) {
                        console.log("failure " + data);
                    });
                },
                saveadd: function() {
                    var feeForm     = w2ui.RAVehicleFeeForm,
                        feesGrid    = w2ui.RAVehicleFeesGrid,
                        TMPASMID    = feeForm.record.TMPASMID;

                    // get TMPVID from last of raflow
                    var TMPVID = app.raflow.last.TMPVID;

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // set local fee data from fee form
                    SetFeeDataFromFeeFormRecord(TMPVID, TMPASMID, "vehicles");

                    SaveVehiclesCompData()
                    .done(function (data) {
                        if (data.status === 'success') {

                            // reset the form
                            feeForm.actions.reset();

                            // set record in form
                            feeForm.record = GetFeeFormInitRecord();
                            feeForm.record.recid = feesGrid.records.length + 1;
                            feeForm.refresh();

                            // enable this field
                            $(feeForm.box).find("#RentCycleText").prop("disabled", false);

                            // Re render the fees grid records
                            AssignVehicleFeesGridRecords(TMPVID);

                        } else {
                            feeForm.message(data.message);
                        }
                    })
                    .fail(function (data) {
                        console.log("failure " + data);
                    });
                },
                delete: function() {
                    var feeForm     = w2ui.RAVehicleFeeForm,
                        feesGrid    = w2ui.RAVehicleFeesGrid,
                        TMPASMID    = feeForm.record.TMPASMID;

                    // get TMPVID from last of raflow
                    var TMPVID = app.raflow.last.TMPVID;

                    var localVehicleData = GetVehicleLocalData(TMPVID);
                    if (localVehicleData.Fees.length > 0) {
                        var itemIndex = GetVehicleFeeLocalData(TMPVID, TMPASMID, true);

                        // remove fee item
                        localVehicleData.Fees.splice(itemIndex, 1);

                        // set this modified local vehicle data to back
                        SetVehicleLocalData(TMPVID, localVehicleData);

                        // sync data on backend side
                        SaveVehiclesCompData()
                        .done(function (data) {
                            if (data.status === 'success') {
                                // reset form as well as remove record from the grid
                                feesGrid.remove(TMPVID);
                                feesGrid.refresh();
                                feeForm.actions.reset();

                                // // Re render the fees grid records
                                AssignVehicleFeesGridRecords(TMPVID);

                                // close the form
                                HideSliderContent(2);
                            } else {
                                feeForm.message(data.message);
                            }
                        })
                        .fail(function (data) {
                            console.log("failure " + data);
                        });
                    }
                }
            },
            onChange: function(event) {
                event.onComplete = function() {
                    var feeForm = w2ui.RAVehicleFeeForm;

                    // take action on change event for this form
                    FeeFormOnChangeHandler(feeForm, event.target, event.value_new);

                       // formRecDiffer: 1=current record, 2=original record, 3=diff object
                    var diff = formRecDiffer(this.record, app.active_form_original, {});
                    // if diff == {} then make dirty flag as false, else true
                    if ($.isPlainObject(diff) && $.isEmptyObject(diff)) {
                        app.form_is_dirty = false;
                    } else {
                        app.form_is_dirty = true;
                    }
                };
            },
            onRefresh: function(event) {
                var feeForm = this;
                event.onComplete = function() {

                    // there is NO VID actually, so have to work around with recid key
                    formRefreshCallBack(feeForm);

                    // set header
                    var header          = "Edit Fee (<strong>{0}</strong>) for Vehicle (<strong>{1}</strong>)".format(vehicleIdentity),
                        vehicleIdentity = GetVehicleIdentity(w2ui.RAVehicleForm.record),
                        vehicleString   = "<em>new</em>";

                    if (w2ui.RAVehicleForm.record.VID > 0) {
                        vehicleString = vehicleIdentity;
                    } else if (vehicleIdentity) {
                        vehicleString = "<em>new</em> - {0}".format(vehicleIdentity);
                    }

                    if (feeForm.record.ARName && feeForm.record.ARName.length > 0) {
                        feeForm.header = header.format(feeForm.record.ARName, vehicleString);
                    } else {
                        feeForm.header = header.format("new", vehicleString);
                    }

                    // FREEZE THE INPUTS IF VERSION IS RAID
                    EnableDisableRAFlowVersionInputs(feeForm);

                    // minimum actions need to be taken care in refres event for fee form
                    FeeFormOnRefreshHandler(feeForm);
                };
            }
        });
    }

    // now load grid in target division
    $('#ra-form #vehicles .grid-container').w2render(w2ui.RAVehiclesGrid);
    HideAllSliderContent();

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
        GetAllARForFeeForm(BID)
        .done(function() {
            AssignVehicleFeesGridRecords(TMPVID);
        });
    }, 0);
};

//-----------------------------------------------------------------------------
// GetVehicleLocalData - returns the clone of vehicle data for requested TMPVID
//-----------------------------------------------------------------------------
window.GetVehicleLocalData = function(TMPVID, returnIndex) {
    var cloneData = {};
    var foundIndex = -1;
    var compData = getRAFlowCompData("vehicles") || [];
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
// SetVehicleLocalData - save the data for requested a TMPVID in local data
//-----------------------------------------------------------------------------
window.SetVehicleLocalData = function(TMPVID, vehicleData) {
    var compData = getRAFlowCompData("vehicles") || [];
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
    var compData = getRAFlowCompData("vehicles");
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
    });

    // assign record in grid
    reassignGridRecids(grid.name);

    // display row error with light red background if it have
    dispalyRAVehiclesGridError();

    // lock the grid until "Have vehicles?" checkbox checked.
    lockOnGrid(grid.name);
};

//------------------------------------------------------------------------------
// SaveVehiclesCompData - saves the data on server side
//------------------------------------------------------------------------------
window.SaveVehiclesCompData = function() {
    var compData = getRAFlowCompData("vehicles");
    return saveActiveCompData(compData, "vehicles");
};

//-----------------------------------------------------------------------------
// GetVehicleFeeLocalData - returns the clone of vehicle fee data for requested
//                          TMPVID and TMPASMID
//-----------------------------------------------------------------------------
window.GetVehicleFeeLocalData = function(TMPVID, TMPASMID, returnIndex) {
    var cloneData = {};
    var foundIndex = -1;
    var compData = getRAFlowCompData("vehicles") || [];
    compData.forEach(function(item, index) {
        if (item.TMPVID == TMPVID) {
            var feesData = item.Fees || [];
            feesData.forEach(function(feeItem, index) {
                if (feeItem.TMPASMID == TMPASMID) {
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
// SetVehicleFeeLocalData - save the data for requested a TMPVID, TMPASMID
//                          in local data
//-----------------------------------------------------------------------------
window.SetVehicleFeeLocalData = function(TMPVID, TMPASMID, vehicleFeeData) {
    var compData = getRAFlowCompData("vehicles");
    var pIndex = -1,
        fIndex = -1;

    compData.forEach(function(item, itemIndex) {
        if (item.TMPVID == TMPVID) {
            var feesData = item.Fees || [];
            feesData.forEach(function(feeItem, feeItemIndex) {
                if (feeItem.TMPASMID == TMPASMID) {
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
// RenderVehicleFeesGridSummary - will render grid summary row from vehicle
//                                comp data
//-----------------------------------------------------------------------------
window.RenderVehicleFeesGridSummary = function(TMPVID) {
    var vehicleData = GetVehicleLocalData(TMPVID),
        grid = w2ui.RAVehicleFeesGrid,
        Fees = vehicleData.Fees || [];

    // render fees amount summary
    RenderFeesGridSummary(grid, Fees);
};

//-----------------------------------------------------------------------------
// AssignVehicleFeesGridRecords - will set the vehicle fees grid records from local
//                                copy of vehicle fees data again
//-----------------------------------------------------------------------------
window.AssignVehicleFeesGridRecords = function(TMPVID) {
    var grid    = w2ui.RAVehicleFeesGrid,
        BID     = getCurrentBID();

    // clear the grid
    grid.clear();

    // list of fees
    var vehicleData = GetVehicleLocalData(TMPVID),
        vehicleFeesData = vehicleData.Fees;

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
    });

    // highlight row with light red background if it have error
    displayRAVehicleFeesGridError();

    // render vehicle fees grid summary
    RenderVehicleFeesGridSummary(TMPVID);
};

//-----------------------------------------------------------------------------
// GetVehicleIdentity - return easily readable vehicle identity string
//-----------------------------------------------------------------------------
window.GetVehicleIdentity = function(record) {
    var year    = record.VehicleYear,
        make    = record.VehicleMake,
        model   = record.VehicleModel;

    if (year || make || model) {
        return "{0} {1} {2}".format(year, make, model);
    }

    return "";
};

// dispalyRAPeopleGridError
// It highlights grid's row if it have error
window.dispalyRAVehiclesGridError = function (){
    // load grid errors if any
    var g = w2ui.RAVehiclesGrid;
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

    if (app.raflow.validationErrors.vehicles) {
        var vehicles = app.raflow.validationCheck.errors.vehicles;
        for (i = 0; i < vehicles.length; i++) {
            if (vehicles[i].total > 0) {
                var recid = getRecIDFromTMPVID(g, vehicles[i].TMPVID);
                g.get(recid).w2ui.style = "background-color: #EEB4B4";
                g.refreshRow(recid);
            }
        }
    }
};

window.displayRAVehicleFeesGridError = function () {
    // load grid errors if any
    var g = w2ui.RAVehicleFeesGrid;
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

    if (app.raflow.validationErrors.vehicles) {
        var vehicles = app.raflow.validationCheck.errors.vehicles;
        for (i = 0; i < vehicles.length; i++) {
            for (var j = 0; j < vehicles[i].fees.length; j++) {
                if (vehicles[i].fees[j].total > 0) {
                    var recid = getRecIDFromTMPASMID(g, vehicles[i].fees[j].TMPASMID);
                    g.get(recid).w2ui.style = "background-color: #EEB4B4";
                    g.refreshRow(recid);
                }
            }
        }
    }
};

// getRecIDFromTMPVID It returns recid of grid record which matches TMPTCID
window.getRecIDFromTMPVID = function(grid, TMPVID){
    // var g = w2ui.RAPeopleGrid;
    var recid;
    for (var i = 0; i < grid.records.length; i++) {
        if (grid.records[i].TMPVID === TMPVID) {
            recid = grid.records[i].recid;
        }
    }
    return recid;
};

// displayRAVehicleFormError If form field have error than it highlight with red border and
window.displayRAVehicleFormError = function(){

    // if pet section doesn't have error than return
    if(!app.raflow.validationErrors.vehicles){
        return;
    }

    var form = w2ui.RAVehicleForm;
    var record = form.record;

    // get list of pets
    var vehicles = app.raflow.validationCheck.errors.vehicles;

    // get index of pet for whom form is opened
    var index = getVehicleIndex(record.TMPVID, vehicles);

    if(index > -1){
        displayFormFieldsError(index, vehicles, "RAVehicleForm");
    }
};

// getVehicleIndex it return an index of vehicle who have TMPVID
window.getVehicleIndex = function (TMPVID, vehicles) {

    var index = -1;

    for(var i = 0; i < vehicles.length; i++){
        // If TMPVID doesn't match iterate for next element
        if(vehicles[i].TMPVID === TMPVID){
            index = i;
            break;
        }
    }

    return index;
};

// displayRAVehicleFeeFormError If form field have error than it highlight with red border and
window.displayRAVehicleFeeFormError = function(TMPVID){

    // if pet section doesn't have error than return
    if(!app.raflow.validationErrors.vehicles){
        return;
    }

    var form = w2ui.RAVehicleFeeForm;
    var record = form.record;

    // get list of pets
    var vehicles = app.raflow.validationCheck.errors.vehicles;

    // get index of vehicle for whom form is opened
    var vehicleIndex = getVehicleIndex(TMPVID, vehicles);

    var index = getFeeIndex(record.TMPASMID, vehicles[vehicleIndex].fees);

    if(index > -1){
        displayFormFieldsError(index, vehicles[vehicleIndex].fees, "RAVehicleFeeForm");
    }
};
