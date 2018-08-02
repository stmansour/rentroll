/* global
    RACompConfig, reassignGridRecids,
    HideSliderContent, ShowSliderContentW2UIComp,
    saveActiveCompData, getRAFlowCompData,
    lockOnGrid,
    GetPetFormInitRecord, GetPetLocalData, SetPetLocalData,
    AssignPetsGridRecords, SavePetsCompData,
    SetRAPetLayoutContent,
    GetPetFeeLocalData, SetPetFeeLocalData,
    AssignPetFeesGridRecords,
    SetRAPetFormRecordFromLocalData,
    SetlocalDataFromRAPetFormRecord,
    GetAllARForFeeForm, SetDataFromFormRecord, SetFormRecordFromData,
    GetFeeGridColumns, GetFeeFormFields, GetFeeFormToolbar,
    SetFeeDataFromFeeFormRecord,
    GetFeeFormInitRecord,
    FeeFormOnChangeHandler, FeeFormOnRefreshHandler,
    SliderContentDivLength, SetFeeFormRecordFromFeeData,
    RenderPetFeesGridSummary, RAFlowNewPetAJAX, updateFlowData,
    GetFeeAccountRulesW2UIListItems, RenderFeesGridSummary, getRecIDFromTMPASMID,
    GetTiePeopleLocalData, RecalculatePetFees, displayRAPetsGridError, getRecIDFromTMPPETID, displayRAPetFeesGridError
*/

"use strict";

//-----------------------------------------------------------------------------
// RAFlowNewPetAJAX - Request to create new entry for a pet in raflow json
//-----------------------------------------------------------------------------
window.RAFlowNewPetAJAX = function() {
    var BID = getCurrentBID();
    var data = {"cmd": "new", "FlowID": app.raflow.activeFlowID};

    return $.ajax({
        url: '/v1/raflow-pets/' + BID.toString() + "/" + app.raflow.activeFlowID.toString() + "/",
        method: "POST",
        data: JSON.stringify(data),
        contentType: "application/json",
        dataType: "json"
    })
    .done(function(data) {
        if (data.status === "success") {
            // Update flow local copy and green checks
            updateFlowData(data);
            // mark new TMPPETID from meta
            app.raflow.last.TMPPETID = data.record.Flow.Data.meta.LastTMPPETID;
        }
    });
};

window.GetPetFormInitRecord = function (previousFormRecord){
    var BID = getCurrentBID();

    var t = new Date(),
        nyd = new Date(new Date().setFullYear(new Date().getFullYear() + 1));

    var defaultFormData = {
        recid:                  w2ui.RAPetsGrid.records.length + 1,
        TMPPETID:               0,
        PETID:                  0,
        TMPTCID:                0,
        BID:                    BID,
        Name:                   "",
        Breed:                  "",
        Type:                   "",
        Color:                  "",
        Weight:                 0,
        DtStart:                w2uiDateControlString(t),
        DtStop:                 w2uiDateControlString(nyd),
        LastModTime:            t.toISOString(),
        LastModBy:              0,
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
// SetlocalDataFromRAPetFormRecord
// ==================================
// will update the data from the record
// it will only update the field defined in fields list in
// form definition
// -------------------------------------------------------------
window.SetlocalDataFromRAPetFormRecord = function(TMPPETID) {
    var form        = w2ui.RAPetForm,
        petFormData = getFormSubmitData(form.record, true);

    // get data from form field's TMPPETID
    var localPetData = GetPetLocalData(TMPPETID);

    // set data from form
    var petData = SetDataFromFormRecord(TMPPETID, form, localPetData);

    // if not Fees then assign in pet data
    if (!petData.hasOwnProperty("Fees")) {
        petData.Fees = [];
    }
    petData.Fees = w2ui.RAPetFeesGrid.records;

    // set this modified data back
    SetPetLocalData(TMPPETID, localPetData);
};

// -------------------------------------------------------------
// SetRAPetFormRecordFromLocalData
// ================================
// will set the data in the form record
// from local pet data
// -------------------------------------------------------------
window.SetRAPetFormRecordFromLocalData = function(TMPPETID) {
    var form = w2ui.RAPetForm;

    // get data from form field's TMPPETID
    var localPetData = GetPetLocalData(TMPPETID);

    // set form record from data
    SetFormRecordFromData(form, localPetData);

    // refresh the form after setting the record
    form.refresh();
    form.refresh();
};

window.loadRAPetsGrid = function () {

    // if form is loaded then return
    if (!("RAPetsGrid" in w2ui)) {

        // -----------------------------------------------------------
        //      ***** PET GRID *****
        // -----------------------------------------------------------
        $().w2grid({
            name: 'RAPetsGrid',
            header: 'Pets',
            show: {
                toolbar: true,
                toolbarSearch: false,
                toolbarAdd: true,
                toolbarReload: true,
                toolbarInput: false,
                toolbarColumns: false,
                footer: true,
            },
            multiSelect: false,
            style: 'border: 0px solid black; display: block;',
            columns: [
                {
                    field: 'recid',
                    caption: 'recid',
                    hidden: true
                },
                {
                    field: 'TMPPETID',
                    caption: 'TMPPETID',
                    hidden: true
                },
                {
                    field: 'PETID',
                    caption: 'PETID',
                    hidden: true
                },
                {
                    field: 'BID',
                    caption: 'BID',
                    hidden: true
                },
                {
                    field: 'haveError',
                    size: '30px',
                    hidden: false,
                    render: function (record) {
                        var haveError = false;
                        var flowID = app.raflow.activeFlowID;
                        if (app.raflow.validationErrors[flowID].pets) {
                            var pets = app.raflow.validationCheck[flowID].errors.pets;
                            for (var i = 0; i < pets.length; i++) {
                                if (pets[i].TMPPETID === record.TMPPETID && pets[i].total > 0) {
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
                    field: 'Name',
                    caption: 'Name',
                    size: '150px'
                },
                {
                    field: 'Type',
                    caption: 'Type',
                    size: '80px'
                },
                {
                    field: 'Breed',
                    caption: 'Breed',
                    size: '80px'
                },
                {
                    field: 'Color',
                    caption: 'Color',
                    size: '80px'
                },
                {
                    field: 'Weight',
                    caption: 'Weight<br>(pounds)',
                    size: '80px'
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
            onRefresh: function (event) {
                event.onComplete = function (){
                    $("#RAPetsGrid_checkbox")[0].checked = app.raflow.data[app.raflow.activeFlowID].Data.meta.HavePets;
                    $("#RAPetsGrid_checkbox")[0].disabled = app.raflow.data[app.raflow.activeFlowID].Data.meta.HavePets;
                    lockOnGrid("RAPetsGrid");
                };
            },
            onClick: function(event) {
                event.onComplete = function() {
                    var yes_args = [this, event.recid],
                        no_args = [this],
                        no_callBack = function(grid) {
                            grid.select(app.last.grid_sel_recid);
                            return false;
                        },
                        yes_callBack = function(grid, recid) {
                            app.last.grid_sel_recid = parseInt(recid);

                            // keep highlighting current row in any case
                            grid.select(app.last.grid_sel_recid);

                            // get TMPPETID from grid
                            var TMPPETID = grid.get(recid).TMPPETID;

                            // keep this clicked TMPPETID in last object
                            app.raflow.last.TMPPETID = TMPPETID;

                            // render layout in the slider
                            ShowSliderContentW2UIComp(w2ui.RAPetLayout, RACompConfig.pets.sliderWidth);

                            // load pet fees grid
                            setTimeout(function() {
                                // fill layout with components and with data
                                SetRAPetLayoutContent(TMPPETID);
                            }, 0);
                        };

                    // warn user if form content has been changed
                    form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
                };
            },
            onAdd: function (/*event*/) {
                var yes_args = [this],
                    no_callBack = function() {
                        return false;
                    },
                    yes_callBack = function(grid) {
                        app.last.grid_sel_recid = -1;
                        grid.selectNone();

                        // get new entry for pet
                        RAFlowNewPetAJAX()
                        .done(function(data) {
                            // get last clicked TMPPETID
                            var TMPPETID = app.raflow.last.TMPPETID;

                            // render the layout in slider
                            ShowSliderContentW2UIComp(w2ui.RAPetLayout, RACompConfig.pets.sliderWidth);

                            // load pet fees grid
                            setTimeout(function() {
                                // fill layout with components
                                SetRAPetLayoutContent(TMPPETID);
                            }, 0);
                        });
                    };

                // warn user if form content has been changed
                form_dirty_alert(yes_callBack, no_callBack, yes_args);
            }
        });

        //------------------------------------------------------------------------
        //  petLayout - The layout to contain the petForm and petFees grid
        //              top  -      petForm
        //              main -      petFeesGrid
        //              bottom -    action buttons form
        //------------------------------------------------------------------------
        $().w2layout({
            name: 'RAPetLayout',
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
        //      ***** PET FORM *****
        // -----------------------------------------------------------
        $().w2form({
            name    : 'RAPetForm',
            header  : 'Add Pet information',
            style   : 'border: 0px; background-color: transparent; display: block;',
            formURL : '/webclient/html/raflow/formra-pets.html',
            toolbar : {
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
                { field: 'recid',                   type: 'int',    required: false,     html: { caption: 'recid', page: 0, column: 0 } },
                { field: 'TMPPETID',                type: 'int',    required: false  },
                { field: 'BID',                     type: 'int',    required: true,      html: { caption: 'BID', page: 0, column: 0 } },
                { field: 'PETID',                   type: 'int',    required: false,     html: { caption: 'PETID', page: 0, column: 0 } },
                { field: 'TMPTCID',                 type: 'list',   required: false,     options: {items: [], selected: {}} },
                { field: 'Name',                    type: 'text',   required: true   },
                { field: 'Breed',                   type: 'text',   required: false  },
                { field: 'Type',                    type: 'text',   required: false  },
                { field: 'Color',                   type: 'text',   required: false  },
                { field: 'Weight',                  type: 'float',  required: false  },
                { field: 'DtStart',                 type: 'date',   required: false,     html: { caption: 'DtStart', page: 0, column: 0 } },
                { field: 'DtStop',                  type: 'date',   required: false,     html: { caption: 'DtStop', page: 0, column: 0 } }
            ],
            actions: {
                reset: function() {
                    w2ui.RAPetForm.clear();
                }
            },
            onRefresh: function(event) {
                event.onComplete = function() {
                    var f = w2ui.RAPetForm,
                        header = "Edit Rental Agreement Pets ({0})";

                    // there is NO PETID actually, so have to work around with recid key
                    formRefreshCallBack(f, "TMPPETID", header);

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
                    if (f.record.TMPPETID === 0) {
                        $("#RAPetFormBtns").find("button[name=delete]").addClass("hidden");
                    } else {
                        $("#RAPetFormBtns").find("button[name=delete]").removeClass("hidden");
                    }

                    // format header
                    var petIdentity = f.record.Name,
                        petString   = "<em>new</em>";

                    if (f.record.PETID > 0) {
                        petString = petIdentity;
                    } else if (petIdentity) {
                        petString = "<em>new</em> - {0}".format(petIdentity);
                    }
                    f.header = "Edit Pet (<strong>{0}</strong>)".format(petString);
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
        //     ***** PET ACTION FORM BUTTONS *****
        //------------------------------------------------------------------------
        $().w2form({
            name: 'RAPetFormBtns',
            style: 'border: none; background-color: transparent;',
            formURL: '/webclient/html/raflow/formra-petbtns.html',
            url: '',
            fields: [],
            actions: {
                save: function() {
                    var f           = w2ui.RAPetForm,
                        TMPPETID    = f.record.TMPPETID;

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // validate form
                    var errors = f.validate();
                    if (errors.length > 0) return;

                    // update the modified data
                    SetlocalDataFromRAPetFormRecord(TMPPETID);

                    // save this records in json Data
                    SavePetsCompData()
                    .done(function(data) {
                        if (data.status === 'success') {
                            // re-assign records in grid
                            AssignPetsGridRecords();

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
                saveadd: function() {
                    var f           = w2ui.RAPetForm,
                        grid        = w2ui.RAPetsGrid,
                        TMPPETID    = f.record.TMPPETID;

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // validate form
                    var errors = f.validate();
                    if (errors.length > 0) return;

                    // update local data from this form record
                    SetlocalDataFromRAPetFormRecord(TMPPETID);

                    // save this records in json Data
                    SavePetsCompData()
                    .done(function(data) {
                        if (data.status === 'success') {
                            // add new formatted record to current form
                            f.actions.reset();
                            f.record = GetPetFormInitRecord(f.record);
                            f.refresh();
                            f.refresh();

                            // re-assign records in grid
                            AssignPetsGridRecords();
                        } else {
                            f.message(data.message);
                        }
                    })
                    .fail(function(data) {
                        console.log("failure " + data);
                    });
                },
                delete: function() {
                    var f = w2ui.RAPetForm;

                    // get local data from TMPPETID
                    var compData = getRAFlowCompData("pets", app.raflow.activeFlowID) || [];
                    var itemIndex = GetPetLocalData(f.record.TMPPETID, true);

                    // if it exists then
                    if (itemIndex > -1) {

                        // remove locally
                        compData.splice(itemIndex, 1);

                        // save this records in json Data
                        SavePetsCompData()
                        .done(function(data) {
                            if (data.status === 'success') {
                                // reset form
                                f.actions.reset();

                                // reassign grid records
                                AssignPetsGridRecords();

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
        });

        // -----------------------------------------------------------
        //      ***** PET ***** FEES ***** GRID *****
        // -----------------------------------------------------------
        $().w2grid({
            name: 'RAPetFeesGrid',
            header: 'Pet Fees',
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
            columns: GetFeeGridColumns(),
            onClick: function(event) {
                event.onComplete = function() {
                    var yes_args = [this, event.recid],
                        no_args = [this],
                        no_callBack = function(grid) {
                            grid.select(app.last.grid_sel_recid);
                            return false;
                        },
                        yes_callBack = function(grid, recid) {
                            var feeForm = w2ui.RAPetFeeForm;

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

                            // get TMPPETID from last of raflow
                            var TMPPETID = app.raflow.last.TMPPETID;

                            // get TMPASMID from grid record
                            var TMPASMID = grid.get(recid).TMPASMID;

                            // get all account rules then
                            var BID = getCurrentBID();
                            GetAllARForFeeForm(BID)
                            .done(function(data) {
                                // get filtered account rules items
                                feeForm.get("ARID").options.items = GetFeeAccountRulesW2UIListItems(BID, "pets");

                                // set record in form
                                SetFeeFormRecordFromFeeData(TMPPETID, TMPASMID, "pets");
                                feeForm.record.RentCycleText = app.cycleFreq[feeForm.record.RentCycle];

                                ShowSliderContentW2UIComp(feeForm, SliderContentDivLength, sliderID);
                                feeForm.refresh(); // need to refresh for header changes

                                // When RentCycle is Norecur then disable the RentCycle list field.
                                var isDisabled = feeForm.record.RentCycleText.text === app.cycleFreq[0];
                                $("#RentCycleText").prop("disabled", isDisabled);
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
                var feesGrid    = w2ui.RAPetFeesGrid,
                    feeForm     = w2ui.RAPetFeeForm;

                var sliderID = 2;
                appendNewSlider(sliderID);
                $("#raflow-container")
                    .find(".slider[data-slider-id="+sliderID+"]")
                    .find(".slider-content")
                    .width(400)
                    .w2render(feeForm);

                // new record so select none
                feesGrid.selectNone();

                var TMPPETID = app.raflow.last.TMPPETID;

                // get all account rules in fit those in form "ARID" field
                var BID = getCurrentBID();
                GetAllARForFeeForm(BID)
                .done(function(data) {
                    // get filtered account rules items
                    feeForm.get("ARID").options.items = GetFeeAccountRulesW2UIListItems(BID, "pets");

                    // set form record
                    SetFeeFormRecordFromFeeData(TMPPETID, 0, "pets");
                    feeForm.record.recid = feesGrid.records.length + 1;

                    // show form in the DOM
                    ShowSliderContentW2UIComp(feeForm, SliderContentDivLength, sliderID);
                    feeForm.refresh();
                })
                .fail(function(data) {
                    console.log("failure" + data);
                });
            },
        });

        // -----------------------------------------------------------
        //      ***** PET ***** FEE ***** FORM *****
        // -----------------------------------------------------------
        $().w2form({
            name: 'RAPetFeeForm',
            header: 'Add New Pet Fee',
            style: 'display: block;',
            formURL: '/webclient/html/raflow/formra-fee.html',
            focus: -1,
            fields: GetFeeFormFields(),
            toolbar : GetFeeFormToolbar(),
            actions: {
                reset: function () {
                    w2ui.RAPetFeeForm.clear();
                },
                save: function() {
                    var feeForm     = w2ui.RAPetFeeForm,
                        TMPASMID    = feeForm.record.TMPASMID;

                    // get TMPPETID from last of raflow
                    var TMPPETID = app.raflow.last.TMPPETID;

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // set local fee data from fee form
                    SetFeeDataFromFeeFormRecord(TMPPETID, TMPASMID, "pets");

                    SavePetsCompData()
                    .done(function (data) {
                        if (data.status === 'success') {
                            // Re render the fees grid records
                            AssignPetFeesGridRecords(TMPPETID);

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
                    var feeForm     = w2ui.RAPetFeeForm,
                        feesGrid    = w2ui.RAPetFeesGrid,
                        TMPASMID    = feeForm.record.TMPASMID;

                    // get TMPPETID from last of raflow
                    var TMPPETID = app.raflow.last.TMPPETID;

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // set local fee data from fee form
                    SetFeeDataFromFeeFormRecord(TMPPETID, TMPASMID, "pets");

                    SavePetsCompData()
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
                            AssignPetFeesGridRecords(TMPPETID);

                        } else {
                            feeForm.message(data.message);
                        }
                    })
                    .fail(function (data) {
                        console.log("failure " + data);
                    });
                },
                delete: function() {
                    var feeForm     = w2ui.RAPetFeeForm,
                        feesGrid    = w2ui.RAPetFeesGrid,
                        TMPASMID    = feeForm.record.TMPASMID;

                    // get TMPPETID from last of raflow
                    var TMPPETID = app.raflow.last.TMPPETID;

                    var localPetData = GetPetLocalData(TMPPETID);
                    if (localPetData.Fees.length > 0) {
                        var itemIndex = GetPetFeeLocalData(TMPPETID, TMPASMID, true);

                        // remove fee item
                        localPetData.Fees.splice(itemIndex, 1);

                        // set this modified local pet data to back
                        SetPetLocalData(TMPPETID, localPetData);

                        // sync data on backend side
                        SavePetsCompData()
                        .done(function (data) {
                            if (data.status === 'success') {
                                // reset form as well as remove record from the grid
                                feesGrid.remove(TMPPETID);
                                feesGrid.refresh();
                                feeForm.actions.reset();

                                // // Re render the fees grid records
                                AssignPetFeesGridRecords(TMPPETID);

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
                    var feeForm = w2ui.RAPetFeeForm;

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

                    // minimum actions need to be taken care in refres event for fee form
                    FeeFormOnRefreshHandler(feeForm);

                    // there is NO PETID actually, so have to work around with recid key
                    formRefreshCallBack(feeForm);

                    // set header
                    var header      = "Edit Fee (<strong>{0}</strong>) for Pet (<strong>{1}</strong>)",
                        petIdentity = w2ui.RAPetForm.record.Name,
                        petString   = "<em>new</em>";

                    if (w2ui.RAPetForm.record.PETID > 0) {
                        petString = petIdentity;
                    } else if (petIdentity) {
                        petString = "<em>new</em> - {0}".format(petIdentity);
                    }

                    if (feeForm.record.ARName && feeForm.record.ARName.length > 0) {
                        feeForm.header = header.format(feeForm.record.ARName, petString);
                    } else {
                        feeForm.header = header.format("new", petString);
                    }
                };
            }
        });
    }

    // now load grid in division
    $('#ra-form #pets .grid-container').w2render(w2ui.RAPetsGrid);

    // load the existing data in pets component
    setTimeout(function () {
        // assign grid records
        AssignPetsGridRecords();
    }, 500);
};

// fill rental agreement pet layout with all forms, grids
window.SetRAPetLayoutContent = function(TMPPETID) {
    w2ui.RAPetLayout.content('bottom',  w2ui.RAPetFormBtns);
    w2ui.RAPetLayout.content('top',     w2ui.RAPetForm);
    w2ui.RAPetLayout.content('main',    w2ui.RAPetFeesGrid);

    // after 0 ms set the record
    setTimeout(function() {
        // set pet form record
        SetRAPetFormRecordFromLocalData(TMPPETID);

        // assign pet fees grid
        var BID = getCurrentBID();
        GetAllARForFeeForm(BID)
        .done(function() {
            AssignPetFeesGridRecords(TMPPETID);
        });
    }, 0);
};

//-----------------------------------------------------------------------------
// GetPetLocalData - returns the clone of pet data for requested TMPPETID
//-----------------------------------------------------------------------------
window.GetPetLocalData = function(TMPPETID, returnIndex) {
    var cloneData = {};
    var foundIndex = -1;
    var compData = getRAFlowCompData("pets", app.raflow.activeFlowID) || [];
    compData.forEach(function(item, index) {
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
// SetPetLocalData - save the data for requested a TMPPETID in local data
//-----------------------------------------------------------------------------
window.SetPetLocalData = function(TMPPETID, petData) {
    var compData = getRAFlowCompData("pets", app.raflow.activeFlowID) || [];
    var dataIndex = -1;
    compData.forEach(function(item, index) {
        if (item.TMPPETID == TMPPETID) {
            dataIndex = index;
            return false;
        }
    });
    if (dataIndex > -1) {
        compData[dataIndex] = petData;
    } else {
        compData.push(petData);
    }
};

//-----------------------------------------------------------------------------
// AssignPetsGridRecords - will set the pets grid records from local
//                               copy of flow data again
//-----------------------------------------------------------------------------
window.AssignPetsGridRecords = function() {
    var compData = getRAFlowCompData("pets", app.raflow.activeFlowID);
    var grid = w2ui.RAPetsGrid;

    // reset last sel recid
    app.last.grid_sel_recid  =-1;

    // clear the grid
    grid.clear();

    compData.forEach(function(petData) {
        var gridRec = {};

        // for each grid column
        grid.columns.forEach(function(gridColumn) {
            gridRec[gridColumn.field] = petData[gridColumn.field];
        });

        // push the record in grid
        grid.records.push(gridRec);
    });

    // assign record in grid
    reassignGridRecids(grid.name);

    // Display row with light red background if it have error
    displayRAPetsGridError();

    // lock the grid until "Have pets?" checkbox checked.
    lockOnGrid(grid.name);
};

//------------------------------------------------------------------------------
// SavePetsCompData - saves the data on server side
//------------------------------------------------------------------------------
window.SavePetsCompData = function() {
    var compData = getRAFlowCompData("pets", app.raflow.activeFlowID);
    return saveActiveCompData(compData, "pets");
};

//-----------------------------------------------------------------------------
// GetPetFeeLocalData - returns the clone of pet fee data for requested
//                      TMPPETID and TMPASMID
//-----------------------------------------------------------------------------
window.GetPetFeeLocalData = function(TMPPETID, TMPASMID, returnIndex) {
    var cloneData = {};
    var foundIndex = -1;
    var compData = getRAFlowCompData("pets", app.raflow.activeFlowID) || [];
    compData.forEach(function(item, index) {
        if (item.TMPPETID == TMPPETID) {
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
// SetPetFeeLocalData - save the data for requested a TMPPETID, TMPASMID
//                   in local data
//-----------------------------------------------------------------------------
window.SetPetFeeLocalData = function(TMPPETID, TMPASMID, petFeeData) {
    var compData = getRAFlowCompData("pets", app.raflow.activeFlowID);
    var pIndex = -1,
        fIndex = -1;

    compData.forEach(function(item, itemIndex) {
        if (item.TMPPETID == TMPPETID) {
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

    // only if pet found then
    if (pIndex > -1) {
        if (fIndex > -1) {
            compData[pIndex].Fees[fIndex] = petFeeData;
        } else {
            compData[pIndex].Fees.push(petFeeData);
        }
    }
};

//-----------------------------------------------------------------------------
// RenderPetFeesGridSummary - will render grid summary row from pet
//                            comp data
//-----------------------------------------------------------------------------
window.RenderPetFeesGridSummary = function(TMPPETID) {
    var petData = GetPetLocalData(TMPPETID),
        grid = w2ui.RAPetFeesGrid,
        Fees = petData.Fees || [];

    // render fees amount summary
    RenderFeesGridSummary(grid, Fees);
};

//-----------------------------------------------------------------------------
// AssignPetFeesGridRecords - will set the pet fees grid records from local
//                            copy of pet fees data again
//-----------------------------------------------------------------------------
window.AssignPetFeesGridRecords = function(TMPPETID) {
    var grid    = w2ui.RAPetFeesGrid,
        BID     = getCurrentBID();

    // clear the grid
    grid.clear();

    // list of fees
    var petData = GetPetLocalData(TMPPETID),
        petFeesData = petData.Fees;

    // pet fees data
    petFeesData.forEach(function(fee) {
        var gridRec = {};

        // for each grid column
        grid.columns.forEach(function(gridColumn) {
            gridRec[gridColumn.field] = fee[gridColumn.field];
        });

        // push the record in grid
        grid.records.push(gridRec);

        // assign recid again
        reassignGridRecids(grid.name);

        // highlight row with light red color if it have error
        displayRAPetFeesGridError();
    });

    // render pet fees grid summary
    RenderPetFeesGridSummary(TMPPETID);
};

//-----------------------------------------------------------------------------
// RecalculatePetFees - will determine if recalcuation needed for pet fees
//                      If needed, it will hit the server to get the latest
//                      new collection of fees for that.
//-----------------------------------------------------------------------------
window.RecalculatePetFees = function (TMPPETID, TMPTCID) {
    var BID = getCurrentBID();
    var tiePerson = GetTiePeopleLocalData(TMPTCID);

    // if no tied rentable then return
    var RID = tiePerson.PRID;
    if (!RID) {
        return;
    }

    var data = {
        "cmd":          "recalculate",
        "FlowID":       app.raflow.activeFlowID,
        "TMPPETID":     TMPPETID,
        "RID":          RID,
    };

    return $.ajax({
        url: "/v1/petfees/" + BID.toString() + "/" + app.raflow.activeFlowID.toString(),
        method: "POST",
        contentType: "application/json",
        dataType: "json",
        data: JSON.stringify(data),
        success: function (data) {
            if (data.status !== "error") {
                // get the last tmpasmid of fees
                var oldLastTMPASMID = app.raflow.data[app.raflow.activeFlowID].Data.meta.LastTMPASMID;

                // Update flow local copy and green checks
                updateFlowData(data);

                // re-assign fees grid records if modifiec
                if (oldLastTMPASMID !== data.record.Flow.Data.meta.LastTMPASMID) {
                    AssignPetFeesGridRecords(TMPPETID);
                }
            }
        },
        error: function (data) {
            console.error(data);
        }
    });
};

// dispalyRAPeopleGridError
// It highlights grid's row if it have error
window.displayRAPetsGridError = function (){
    // load grid errors if any
    var g = w2ui.RAPetsGrid;
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

    // If biz error than highlight grid row
    var flowID = app.raflow.activeFlowID;
    if (app.raflow.validationErrors[flowID].pets) {
        var pets = app.raflow.validationCheck[flowID].errors.pets;
        for (i = 0; i < pets.length; i++) {
            var recid = getRecIDFromTMPPETID(g, pets[i].TMPPETID);
            if (pets[i].total > 0) {
                g.get(recid).w2ui.style = "background-color: #EEB4B4";
                g.refreshRow(recid);
            }else{
                g.get(recid).w2ui.style = {};
                g.refreshRow(recid);
            }
        }
    }
};

window.displayRAPetFeesGridError = function () {
    // load grid errors if any
    var g = w2ui.RAPetFeesGrid;
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

    // If biz error than highlight grid row
    var flowID = app.raflow.activeFlowID;
    if (app.raflow.validationErrors[flowID].pets) {
        var pets = app.raflow.validationCheck[flowID].errors.pets;
        for (i = 0; i < pets.length; i++) {
            for (var j = 0; j < pets[i].fees.length; j++) {
                if (pets[i].fees[j].total > 0) {
                    var recid = getRecIDFromTMPASMID(g, pets[i].fees[j].TMPASMID);
                    g.get(recid).w2ui.style = "background-color: #EEB4B4";
                    g.refreshRow(recid);
                }
            }
        }
    }
};

// getRecIDFromTMPPETID It returns recid of grid record which matches TMPTCID
window.getRecIDFromTMPPETID = function(grid, TMPPETID){
    // var g = w2ui.RAPeopleGrid;
    var recid;
    for (var i = 0; i < grid.records.length; i++) {
        if (grid.records[i].TMPPETID === TMPPETID) {
            recid = grid.records[i].recid;
        }
    }
    return recid;
};
