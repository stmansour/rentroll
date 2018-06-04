/* global
    RACompConfig, sliderContentDivLength, reassignGridRecids,
    hideSliderContent, appendNewSlider, showSliderContentW2UIComp,
    loadTargetSection, requiredFieldsFulFilled, getRAFlowPartTypeIndex, initRAFlowAjax,
    getRAFlowAllParts, saveActiveCompData, toggleHaveCheckBoxDisablity, getRAFlowCompData,
    lockOnGrid,
    getPetFormInitRecord
*/

"use strict";

window.getPetFormInitRecord = function (BID, BUD, previousFormRecord){
    var t = new Date(),
        nyd = new Date(new Date().setFullYear(new Date().getFullYear() + 1));

    var defaultFormData = {
        recid: 0,
        PETID: 0,
        BID: BID,
        // BUD: BUD,
        Name: "",
        Breed: "",
        Type: "",
        Color: "",
        Weight: 0,
        DtStart: w2uiDateControlString(t),
        DtStop: w2uiDateControlString(nyd),
        NonRefundablePetFee: 0,
        RefundablePetDeposit: 0,
        RecurringPetFee: 0,
        LastModTime: t.toISOString(),
        LastModBy: 0,
    };

    // if it called after 'save and add another' action there previous form record is passed as Object
    // else it is null
    if ( previousFormRecord ) {
        defaultFormData = setDefaultFormFieldAsPreviousRecord(
            [ 'Name', 'Breed', 'Type', 'Color', 'Weight',
              'NonRefundablePetFee', 'RefundablePetDeposit', 'ReccurringPetFee' ], // Fields to Reset
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
                { field: 'recid', type: 'int', required: false, html: { caption: 'recid', page: 0, column: 0 } },
                { field: 'BID', type: 'int', hidden: true, html: { caption: 'BID', page: 0, column: 0 } },
                // { field: 'BUD', type: 'text', hidden: false, html: { caption: 'BUD', page: 0, column: 0 } },
                { field: 'PETID', type: 'int', hidden: false, html: { caption: 'PETID', page: 0, column: 0 } },
                { field: 'Name', type: 'text', required: true},
                { field: 'Breed', type: 'text', required: true},
                { field: 'Type', type: 'text', required: true},
                { field: 'Color', type: 'text', required: true},
                { field: 'Weight', type: 'int', required: true},
                { field: 'NonRefundablePetFee', type: 'money', required: false},
                { field: 'RefundablePetDeposit', type: 'money', required: false},
                { field: 'RecurringPetFee', type: 'money', required: false},
                { field: 'DtStart', type: 'date', required: false, html: { caption: 'DtStart', page: 0, column: 0 } },
                { field: 'DtStop', type: 'date', required: false, html: { caption: 'DtStop', page: 0, column: 0 } },
                { field: 'LastModTime', type: 'time', required: false, html: { caption: 'LastModTime', page: 0, column: 0 } },
                { field: 'LastModBy', type: 'int', required: false, html: { caption: 'LastModBy', page: 0, column: 0 } },
            ],
            onRefresh: function(event) {
                event.onComplete = function() {
                    var f = w2ui.RAPetForm,
                        header = "Edit Rental Agreement Pets ({0})";

                    // there is NO PETID actually, so have to work around with recid key
                    formRefreshCallBack(f, "recid", header);

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
                save: function() {
                    var form = this;
                    var grid = w2ui.RAPetsGrid;
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
                    saveActiveCompData(recordsData, "pets")
                    .done(function(data) {
                        if (data.status === 'success') {
                            // if null
                            if (isNewRecord) {
                                grid.add(record);
                            } else {
                                grid.set(record.recid, record);
                            }
                            form.clear();

                            // Disable "have pets?" checkbox if there is any record.
                            toggleHaveCheckBoxDisablity('RAPetsGrid');

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
                saveadd: function() {
                    var BID = getCurrentBID(),
                        BUD = getBUDfromBID(BID);

                    var form = this;
                    var grid = w2ui.RAPetsGrid;
                    var errors = form.validate();
                    if (errors.length > 0) return;
                    var record = $.extend(true, {}, form.record);
                    var recordsData = $.extend(true, [], grid.records);
                    var isNewRecord = (grid.get(record.recid, true) === null);

                    // if it doesn't exist then only push
                    if (isNewRecord) {
                        recordsData.push(record);
                    }

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // save this records in json Data
                    saveActiveCompData(recordsData, "pets")
                    .done(function(data) {
                        if (data.status === 'success') {
                            // clear the grid select recid
                            app.last.grid_sel_recid  =-1;
                            // selectNone
                            grid.selectNone();

                            // if null
                            if (isNewRecord) {
                                // add this record to grid
                                grid.add(record);
                            } else {
                                grid.set(record.recid, record);
                            }
                            // add new formatted record to current form
                            form.record = getPetFormInitRecord(BID, BUD, form.record);
                            // set record id
                            form.record.recid = grid.records.length + 1;
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
                delete: function() {
                    var form = this;
                    var grid = w2ui.RAPetsGrid;

                    // backup the records
                    var records = $.extend(true, [], grid.records);
                    for (var i = 0; i < records.length; i++) {
                        if(records[i].recid == form.record.recid) {
                            records.splice(i, 1);
                        }
                    }

                    // save this records in json Data
                    saveActiveCompData(records, "pets")
                    .done(function(data) {
                        if (data.status === 'success') {
                            // clear the grid select recid
                            app.last.grid_sel_recid  =-1;
                            // selectNone
                            grid.selectNone();

                            grid.remove(form.record.recid);
                            form.clear();

                            // Disable "have pets?" checkbox if there is any record.
                            toggleHaveCheckBoxDisablity('RAPetsGrid');

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
                    field: 'PETID',
                    hidden: true
                },
                {
                    field: 'BID',
                    hidden: true
                },
                /*{
                    field: 'BUD',
                    hidden: true
                },*/
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

                        w2ui.RAPetForm.record = getPetFormInitRecord(BID, BUD, null);
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
        var compData = getRAFlowCompData("pets", app.raflow.activeFlowID);
        var grid = w2ui.RAPetsGrid;

        if (compData) {
            grid.records = compData;
            reassignGridRecids(grid.name);

            // lock the grid until "Have pets?" checkbox checked.
            lockOnGrid(grid.name);
        } else {
            grid.clear();
        }
    }, 500);
};