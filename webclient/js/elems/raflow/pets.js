/* global
    RACompConfig, reassignGridRecids, RAFlowAJAX,
    HideSliderContent, ShowSliderContentW2UIComp,
    SaveCompDataAJAX, GetRAFlowCompLocalData,
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
    RenderPetFeesGridSummary, RAFlowNewPetAJAX, UpdateRAFlowLocalData,
    GetFeeAccountRulesW2UIListItems, RenderFeesGridSummary, getRecIDFromTMPASMID,
    displayRAPetFormError, getPetIndex, displayFormFieldsError, displayRAPetFeeFormError, getFeeIndex,
    GetTiePeopleLocalData, displayRAPetsGridError, getRecIDFromTMPPETID, displayRAPetFeesGridError,
    GetCurrentFlowID, EnableDisableRAFlowVersionInputs, ShowHideGridToolbarAddButton,
    HideAllSliderContent
*/

"use strict";

//-----------------------------------------------------------------------------
// RAFlowNewPetAJAX - Request to create new entry for a pet in raflow json
//-----------------------------------------------------------------------------
window.RAFlowNewPetAJAX = function() {
    var BID = getCurrentBID();
    var FlowID = GetCurrentFlowID();

    var url = '/v1/raflow-pets/' + BID.toString() + "/" + FlowID.toString() + "/";
    var data = {
        "cmd": "new",
        "FlowID": FlowID
    };

    return RAFlowAJAX(url, "POST", data, true)
    .done(function(data) {
        if (data.status !== "error") {
            // reassign records
            AssignPetsGridRecords();

            // mark new TMPPETID from meta
            app.raflow.last.TMPPETID = data.record.Flow.Data.meta.LastTMPPETID;
        }
    });
};

window.GetPetFormInitRecord = function (previousFormRecord){
    var defaultFormData = {
        recid:                  w2ui.RAPetsGrid.records.length + 1,
        TMPPETID:               0,
        PETID:                  0,
        TMPTCID:                0,
        Name:                   "",
        Breed:                  "",
        Type:                   "",
        Color:                  "",
        Weight:                 0,
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
                toolbarReload: false,
                toolbarInput: false,
                toolbarColumns: false,
                footer: true,
            },
            multiSelect: false,
            style: 'border: none; display: block;',
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
                    field: 'haveError',
                    size: '30px',
                    hidden: false,
                    render: function (record) {
                        var haveError = false;
                        if (app.raflow.validationErrors.pets) {
                            var pets = app.raflow.validationCheck.errors.pets.errors;
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
                    size: '120px'
                },
                {
                    field: 'Breed',
                    caption: 'Breed',
                    size: '150px'
                },
                {
                    field: 'Color',
                    caption: 'Color',
                    size: '100px'
                },
                {
                    field: 'Weight',
                    caption: 'Weight<br>(pounds)',
                    size: '100px'
                },
            ],
            onRefresh: function (event) {
                var grid = this;
                event.onComplete = function (){
                    $("#RAPetsGrid_checkbox")[0].checked = app.raflow.Flow.Data.meta.HavePets;
                    $("#RAPetsGrid_checkbox")[0].disabled = app.raflow.Flow.Data.meta.HavePets;
                    lockOnGrid("RAPetsGrid");

                    ShowHideGridToolbarAddButton(grid.name);
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

                            // load form fields error
                            setTimeout(function () {
                                displayRAPetFormError();
                            }, 500);

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
                            w2ui.RAPetsGrid.selectNone();
                            break;
                    }
                }
            },
            fields  : [
                { field: 'recid',                   type: 'int',    required: false,     html: { caption: 'recid', page: 0, column: 0 } },
                { field: 'TMPPETID',                type: 'int',    required: false  },
                { field: 'PETID',                   type: 'int',    required: false,     html: { caption: 'PETID', page: 0, column: 0 } },
                { field: 'TMPTCID',                 type: 'list',   required: false,     options: {items: [], selected: {}} },
                { field: 'Name',                    type: 'text',   required: true   },
                { field: 'Breed',                   type: 'text',   required: false  },
                { field: 'Type',                    type: 'text',   required: false  },
                { field: 'Color',                   type: 'text',   required: false  },
                { field: 'Weight',                  type: 'float',  required: false  },
            ],
            actions: {
                reset: function() {
                    w2ui.RAPetForm.clear();
                }
            },
            onRefresh: function(event) {
                event.onComplete = function() {
                    var form = w2ui.RAPetForm,
                        header = "Edit Rental Agreement Pets ({0})";

                    // there is NO PETID actually, so have to work around with recid key
                    formRefreshCallBack(form, "TMPPETID", header);

                    // selection of contact person
                    var TMPTCIDSel = {};
                    app.raflow.peopleW2UIItems.forEach(function(item) {
                        if (item.id === form.record.TMPTCID) {
                            $.extend(TMPTCIDSel, item);
                        }
                    });
                    form.get("TMPTCID").options.items = app.raflow.peopleW2UIItems;
                    form.get("TMPTCID").options.selected = TMPTCIDSel;

                    // hide delete button if it is NewRecord
                    if (form.record.TMPPETID === 0) {
                        $("#RAPetFormBtns").find("button[name=delete]").addClass("hidden");
                    } else {
                        $("#RAPetFormBtns").find("button[name=delete]").removeClass("hidden");
                    }

                    // format header
                    var petIdentity = form.record.Name,
                        petString   = "<em>new</em>";

                    if (form.record.PETID > 0) {
                        petString = petIdentity;
                    } else if (petIdentity) {
                        petString = "<em>new</em> - {0}".format(petIdentity);
                    }
                    form.header = "Edit Pet (<strong>{0}</strong>)".format(petString);

                    // FREEZE THE INPUTS IF VERSION IS RAID
                    EnableDisableRAFlowVersionInputs(form);
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

                            // get new entry for pet
                            RAFlowNewPetAJAX()
                            .done(function(data) {
                                // IT'S MANAGED IN AJAX API
                                var TMPPETID = app.raflow.last.TMPPETID;

                                // add new formatted record to current form
                                f.actions.reset();
                                f.record = GetPetLocalData(TMPPETID);
                                f.refresh();
                                f.refresh();

                                // re-assign records in grid
                                AssignPetsGridRecords();
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
                delete: function() {
                    var f = w2ui.RAPetForm;

                    // get local data from TMPPETID
                    var compData = GetRAFlowCompLocalData("pets") || [];
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

                                // Unselect all selected record from the grid
                                w2ui.RAPetsGrid.selectNone();
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
                toolbarColumns: false,
                footer:         false
            },
            multiSelect: false,
            style: 'border-color: silver; border-style: solid; border-width: 1px 0 1px 0;',
            columns: GetFeeGridColumns('RAPetFeesGrid'),
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

                                // displat form field error if it have
                                setTimeout(function(){
                                    displayRAPetFeeFormError(w2ui.RAPetForm.record.TMPPETID);
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
            onRefresh: function(event) {
                var grid = this;
                event.onComplete = function() {
                    ShowHideGridToolbarAddButton(grid.name);
                };
            }
        });

        // -----------------------------------------------------------
        //      ***** PET ***** FEE ***** FORM *****
        // -----------------------------------------------------------
        $().w2form({
            name: 'RAPetFeeForm',
            header: 'Add New Pet Fee',
            style: 'border: none; display: block;',
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
                        TMPASMID    = feeForm.record.TMPASMID;

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // FRESH NEW FEE THEN JUST RETURN WITH CLOSING THE FORM
                    if (TMPASMID === 0) {
                        // reset form
                        feeForm.actions.reset();

                        // close the form
                        HideSliderContent(2);

                        return;
                    }

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
                                // reset form
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

                    // minimum actions need to be taken care in refres event for fee form
                    FeeFormOnRefreshHandler(feeForm);
                };
            }
        });
    }

    // now load grid in division
    $('#ra-form #pets .grid-container').w2render(w2ui.RAPetsGrid);
    HideAllSliderContent();

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
    var compData = GetRAFlowCompLocalData("pets") || [];
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
    var compData = GetRAFlowCompLocalData("pets") || [];
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
    var compData = GetRAFlowCompLocalData("pets");
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
    var compData = GetRAFlowCompLocalData("pets");
    return SaveCompDataAJAX(compData, "pets");
};

//-----------------------------------------------------------------------------
// GetPetFeeLocalData - returns the clone of pet fee data for requested
//                      TMPPETID and TMPASMID
//-----------------------------------------------------------------------------
window.GetPetFeeLocalData = function(TMPPETID, TMPASMID, returnIndex) {
    var cloneData = {};
    var foundIndex = -1;
    var compData = GetRAFlowCompLocalData("pets") || [];
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
    var compData = GetRAFlowCompLocalData("pets");
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
    });

    // highlight row with light red color if it have error
    displayRAPetFeesGridError();

    // render pet fees grid summary
    RenderPetFeesGridSummary(TMPPETID);
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

    if (app.raflow.validationErrors.pets) {
        var pets = app.raflow.validationCheck.errors.pets.errors;
        for (i = 0; i < pets.length; i++) {
            var recid = getRecIDFromTMPPETID(g, pets[i].TMPPETID);
            if (pets[i].total > 0) {
                g.get(recid).w2ui.style = "background-color: #EEB4B4";
            }else{
                g.get(recid).w2ui.style = {};
            }
            g.refreshRow(recid);
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

    if (app.raflow.validationErrors.pets) {
        var pets = app.raflow.validationCheck.errors.pets.errors;
        for (i = 0; i < pets.length; i++) {
            for (var j = 0; j < pets[i].fees.errors.length; j++) {
                if (pets[i].fees.errors[j].total > 0) {
                    var recid = getRecIDFromTMPASMID(g, pets[i].fees.errors[j].TMPASMID);
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

// displayRAPetFormError If form field have error than it highlight with red border and
window.displayRAPetFormError = function(){

    // if pet section doesn't have error than return
    if(!app.raflow.validationErrors.pets){
        return;
    }

    var form = w2ui.RAPetForm;
    var record = form.record;

    // get list of pets
    var pets = app.raflow.validationCheck.errors.pets.errors;

    // get index of pet for whom form is opened
    var index = getPetIndex(record.TMPPETID, pets);

    if(index > -1){
        displayFormFieldsError(index, pets, "RAPetForm");
    }
};

// getPetIndex it return an index of pet who have TMPPETID
window.getPetIndex = function (TMPPETID, pets) {

    var index = -1;

    for(var i = 0; i < pets.length; i++){
        // If TMPPETID doesn't match iterate for next element
        if(pets[i].TMPPETID === TMPPETID){
            index = i;
            break;
        }
    }

    return index;
};


// displayRAPetFeeFormError If form field have error than it highlight with red border and
window.displayRAPetFeeFormError = function(TMPPETID){

    // if pet section doesn't have error than return
    if(!app.raflow.validationErrors.pets){
        return;
    }

    var form = w2ui.RAPetFeeForm;
    var record = form.record;

    // get list of pets
    var pets = app.raflow.validationCheck.errors.pets.errors;

    // get index of pet for whom form is opened
    var petIndex = getPetIndex(TMPPETID, pets);

    var index = getFeeIndex(record.TMPASMID, pets[petIndex].fees.errors);

    if(index > -1){
        displayFormFieldsError(index, pets[petIndex].fees.errors, "RAPetFeeForm");
    }
};
