/* global
    RACompConfig, sliderContentDivLength, reassignGridRecids,
    hideSliderContent, appendNewSlider, showSliderContentW2UIComp,
    loadTargetSection, requiredFieldsFulFilled, getRAFlowPartTypeIndex, initRAFlowAjax,
    getRAFlowAllParts, saveActiveCompData, toggleHaveCheckBoxDisablity, getRAFlowCompData,
    lockOnGrid,
    getPetFormInitRecord, getPetLocalData, setPetLocalData,
    AssignPetsGridRecords, savePetsCompData
*/

"use strict";

window.getPetFormInitRecord = function (previousFormRecord){
    var BID = getCurrentBID();

    var t = new Date(),
        nyd = new Date(new Date().setFullYear(new Date().getFullYear() + 1));

    var defaultFormData = {
        recid:                  0,
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
        NonRefundablePetFee:    0,
        RefundablePetDeposit:   0,
        RecurringPetFee:        0,
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

window.loadRAPetsGrid = function () {

    // if form is loaded then return
    if (!("RAPetsGrid" in w2ui)) {

        // pet form
        $().w2form({
            name    : 'RAPetForm',
            header  : 'Add Pet information',
            style   : 'border: 0px; background-color: transparent; display: block;',
            formURL : '/webclient/html/formrapets.html',
            toolbar : {
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
                { field: 'recid',                   type: 'int',    required: false,    html: { caption: 'recid', page: 0, column: 0 } },
                { field: 'TMPPETID',                type: 'int',    required: true  },
                { field: 'BID',                     type: 'int',    required: true,     html: { caption: 'BID', page: 0, column: 0 } },
                { field: 'PETID',                   type: 'int',    required: true,     html: { caption: 'PETID', page: 0, column: 0 } },
                { field: 'TMPTCID',                 type: 'list',   required: true,     options: {items: [], selected: {}} },
                { field: 'Name',                    type: 'text',   required: true  },
                { field: 'Breed',                   type: 'text',   required: true  },
                { field: 'Type',                    type: 'text',   required: true  },
                { field: 'Color',                   type: 'text',   required: true  },
                { field: 'Weight',                  type: 'int',    required: true  },
                { field: 'NonRefundablePetFee',     type: 'money',  required: true  },
                { field: 'RefundablePetDeposit',    type: 'money',  required: true  },
                { field: 'RecurringPetFee',         type: 'money',  required: true  },
                { field: 'DtStart',                 type: 'date',   required: true,     html: { caption: 'DtStart', page: 0, column: 0 } },
                { field: 'DtStop',                  type: 'date',   required: true,     html: { caption: 'DtStop', page: 0, column: 0 } },
                { field: 'LastModTime',             type: 'time',   required: false,    html: { caption: 'LastModTime', page: 0, column: 0 } },
                { field: 'LastModBy',               type: 'int',    required: false,    html: { caption: 'LastModBy', page: 0, column: 0 } },
            ],
            onRefresh: function(event) {
                event.onComplete = function() {
                    var f = w2ui.RAPetForm,
                        header = "Edit Rental Agreement Pets ({0})";

                    // there is NO PETID actually, so have to work around with recid key
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
                    var isNewRecord = (w2ui.RAPetsGrid.get(f.record.recid, true) === null);
                    if (isNewRecord) {
                        $(f.box).find("button[name=delete]").addClass("hidden");
                    } else {
                        $(f.box).find("button[name=delete]").removeClass("hidden");
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
            },
            actions: {
                reset: function() {
                    w2ui.RAPetForm.clear();
                },
                save: function() {
                    var f = w2ui.RAPetForm,
                        grid = w2ui.RAPetsGrid,
                        TMPPETID = f.record.TMPPETID;

                    // validate form
                    var errors = f.validate();
                    if (errors.length > 0) return;

                    // sync this info in local data
                    var petData = getFormSubmitData(f.record, true);

                    // set data locally
                    setPetLocalData(TMPPETID, petData);

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // save this records in json Data
                    savePetsCompData()
                    .done(function(data) {
                        if (data.status === 'success') {
                            // re-assign records in grid
                            AssignPetsGridRecords();

                            // Disable "have pets?" checkbox if there is any record.
                            toggleHaveCheckBoxDisablity('RAPetsGrid');

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
                saveadd: function() {
                    var f = w2ui.RAPetForm,
                        grid = w2ui.RAPetsGrid,
                        TMPPETID = f.record.TMPPETID;

                    // validate form
                    var errors = f.validate();
                    if (errors.length > 0) return;

                    // sync this info in local data
                    var petData = getFormSubmitData(f.record, true);

                    // set data locally
                    setPetLocalData(TMPPETID, petData);

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // save this records in json Data
                    savePetsCompData()
                    .done(function(data) {
                        if (data.status === 'success') {
                            // add new formatted record to current form
                            f.record = getPetFormInitRecord(f.record);
                            // set record id
                            f.record.recid = grid.records.length + 1;
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
                    var itemIndex = getPetLocalData(f.record.TMPPETID, true);
                    compData.splice(itemIndex, 1);

                    // save this records in json Data
                    savePetsCompData()
                    .done(function(data) {
                        if (data.status === 'success') {
                            // reset form
                            f.actions.reset();

                            // Disable "have pets?" checkbox if there is any record.
                            toggleHaveCheckBoxDisablity('RAPetsGrid');

                            // reassign grid records
                            AssignPetsGridRecords();

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
            },
        });

        // pets grid
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
                    hidden: true
                },
                {
                    field: 'TMPPETID',
                    hidden: true
                },
                {
                    field: 'PETID',
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
                    caption: 'Weight',
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
                },
                {
                    field: 'NonRefundablePetFee',
                    caption: 'NonRefundable<br>PetFee',
                    size: '70px',
                    render: 'money'
                },
                {
                    field: 'RefundablePetDeposit',
                    caption: 'Refundable<br>PetDeposit',
                    size: '70px',
                    render: 'money'
                },
                {
                    field: 'RecurringPetFee',
                    caption: 'Recurring<br>PetFee',
                    size: '70px',
                    render: 'money'
                }
            ],
            onChange: function (event) {
                event.onComplete = function () {
                    this.save();
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
                            w2ui.RAPetForm.record = $.extend(true, {}, grid.get(app.last.grid_sel_recid));

                            showSliderContentW2UIComp(w2ui.RAPetForm, RACompConfig.pets.sliderWidth);
                            w2ui.RAPetForm.refresh(); // need to refresh for header changes
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

                        var BID = getCurrentBID(),
                            BUD = getBUDfromBID(BID);

                        w2ui.RAPetForm.record = getPetFormInitRecord(null);
                        // set record id
                        w2ui.RAPetForm.record.recid = w2ui.RAPetsGrid.records.length + 1;
                        showSliderContentW2UIComp(w2ui.RAPetForm, RACompConfig.pets.sliderWidth);
                        w2ui.RAPetForm.refresh();
                    };

                // warn user if form content has been changed
                form_dirty_alert(yes_callBack, no_callBack, yes_args);
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

//-----------------------------------------------------------------------------
// getPetLocalData - returns the clone of pet data for requested TMPPETID
//-----------------------------------------------------------------------------
window.getPetLocalData = function(TMPPETID, returnIndex) {
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
// setPetLocalData - save the data for requested a TMPPETID in local data
//-----------------------------------------------------------------------------
window.setPetLocalData = function(TMPPETID, petData) {
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
    // select none
    grid.selectNone();

    if (compData) {
        grid.records = compData;
        reassignGridRecids(grid.name);

        // lock the grid until "Have pets?" checkbox checked.
        lockOnGrid(grid.name);

        // Operation on RAPetForm
        w2ui.RAPetForm.refresh();
    } else {
        // clear the grid
        grid.clear();
        // Operation on RAPetForm
        w2ui.RAPetForm.actions.reset();
    }
};

//------------------------------------------------------------------------------
// savePetsCompData - saves the data on server side
//------------------------------------------------------------------------------
window.savePetsCompData = function() {
    var compData = getRAFlowCompData("pets", app.raflow.activeFlowID);
    return saveActiveCompData(compData, "pets");
};

