/* global
    RACompConfig, reassignGridRecids,
    hideSliderContent, showSliderContentW2UIComp,
    saveActiveCompData, getRAFlowCompData,
    lockOnGrid,
    getPetFormInitRecord, getPetLocalData, setPetLocalData,
    AssignPetsGridRecords, savePetsCompData,
    SetRAPetLayoutContent,
    getPetFeeLocalData, setPetFeeLocalData,
    AssignPetFeesGridRecords,
    SetRAPetFormRecordFromLocalData,
    SetlocalDataFromRAPetFormRecord,
    getAllARsWithAmount
*/

"use strict";

window.getPetFormInitRecord = function (previousFormRecord){
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
        fields      = form.fields || [],
        petFormData = getFormSubmitData(form.record, true);

    // local pet data
    var localPetData;

    if (TMPPETID === 0) {
        localPetData = petFormData;
    } else {
        // get data from form field's TMPPETID
        localPetData = getPetLocalData(petFormData.TMPPETID);

        // loop over form fields
        fields.forEach(function(fieldItem) {
            localPetData[fieldItem.field] = petFormData[fieldItem.field];
        });
    }

    // if not Fees then assign in pet data
    if (!localPetData.hasOwnProperty("Fees")) {
        localPetData.Fees = [];
    }
    localPetData.Fees = w2ui.RAPetFeesGrid.records;

    // set this modified data back
    setPetLocalData(TMPPETID, localPetData);
};

// -------------------------------------------------------------
// SetRAPetFormRecordFromLocalData
// ================================
// will set the data in the form record
// from local pet data
// -------------------------------------------------------------
window.SetRAPetFormRecordFromLocalData = function(TMPPETID) {
    var form        = w2ui.RAPetForm,
        fields      = form.fields || [];

    if (TMPPETID === 0) {
        form.record = getPetFormInitRecord(null);
    } else {
        // get data from form field's TMPPETID
        var localPetData = getPetLocalData(TMPPETID);

        // loop over form fields
        fields.forEach(function(fieldItem) {
            form.record[fieldItem.field] = localPetData[fieldItem.field];
        });
    }

    // refresh the form after setting the record
    form.refresh();
    form.refresh();
};

window.loadRAPetsGrid = function () {

    // if form is loaded then return
    if (!("RAPetsGrid" in w2ui)) {

        // pet form
        $().w2form({
            name    : 'RAPetForm',
            header  : 'Add Pet information',
            style   : 'border: 0px; background-color: transparent; display: block;',
            formURL : '/webclient/html/raflow/formrapets.html',
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
                { field: 'DtStart',                 type: 'date',   required: true,     html: { caption: 'DtStart', page: 0, column: 0 } },
                { field: 'DtStop',                  type: 'date',   required: true,     html: { caption: 'DtStop', page: 0, column: 0 } },
                { field: 'LastModTime',             type: 'time',   required: false,    html: { caption: 'LastModTime', page: 0, column: 0 } },
                { field: 'LastModBy',               type: 'int',    required: false,    html: { caption: 'LastModBy', page: 0, column: 0 } },
            ],
            actions: {
                reset: function() {
                    w2ui.RAPetForm.clear();
                },
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
                }
            ],
            onChange: function (event) {
                event.onComplete = function () {
                    this.save();
                };
            },
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
                            var TMPPETID = grid.get(app.last.grid_sel_recid).TMPPETID;

                            // render layout in the slider
                            showSliderContentW2UIComp(w2ui.RAPetLayout, RACompConfig.pets.sliderWidth);

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

                        // render the layout in slider
                        showSliderContentW2UIComp(w2ui.RAPetLayout, RACompConfig.pets.sliderWidth);

                        // load pet fees grid
                        setTimeout(function() {
                            // fill layout with components
                            SetRAPetLayoutContent(0);
                        }, 0);
                    };

                // warn user if form content has been changed
                form_dirty_alert(yes_callBack, no_callBack, yes_args);
            }
        });

        // pet fees grid
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
            columns: [
                {
                    field: 'recid',
                    hidden: true
                },
                {
                    field: 'TMPPETID',
                    caption: 'TMPPETID',
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
                    caption: "Remove Pet Fee",
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
                        localPetFeeData = getPetFeeLocalData(record.TMPPETID, record.ARID);

                    switch(event.column) {
                    case grid.getColumn("Amount", true):
                        // update data in local and grid record
                        localPetFeeData.Amount = record.Amount = parseFloat(event.value_new);
                        // set data
                        grid.set(event.recid, record);
                        setPetFeeLocalData(record.TMPPETID, record.ARID, localPetFeeData);
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
                        localPetFeeData.Amount = record.Amount = arItem.DefaultAmount;
                        localPetFeeData.ARID = record.ARID = parseInt(event.value_new);
                        // set data
                        grid.set(event.recid, record);
                        // grid.getColumn("ARName").render();
                        setPetFeeLocalData(record.TMPPETID, record.ARID, localPetFeeData);
                        break;
                    }

                    w2ui.RAPetFeesGrid.save();
                    // TODO(Sudip): we still need to update data locally
                };
            },
            onClick: function(event) {
                event.onComplete = function() {
                    // if it's remove column then remove the record
                    // maybe confirm dialog will be added
                    if(w2ui.RAPetFeesGrid.getColumn("RemoveRec", true) == event.column) {

                        // remove entry from local data
                        var rec = w2ui.RAPetFeesGrid.get(event.recid);
                        var petData = getPetLocalData(rec.TMPPETID);

                        if (petData.hasOwnProperty("Fees")) {
                            var feeIndex = getPetFeeLocalData(rec.TMPPETID, rec.ARID, true);

                            // also manage local data
                            petData.Fees.splice(feeIndex, 1);

                            // set modified petData back in local
                            setPetLocalData(rec.TMPPETID, petData);

                            // save the data on server data
                            savePetsCompData()
                            .done(function(data) {
                                if (data.status === "success") {
                                    // remove from grid
                                    w2ui.RAPetFeesGrid.remove(event.recid);
                                }
                            });
                        } else {
                            // simple remove record from grid
                            w2ui.RAPetFeesGrid.remove(event.recid);
                        }

                        return;
                    }
                };
            },
            onAdd: function(/*event*/) {
                var grid = w2ui.RAPetFeesGrid;
                var rec     = {
                    recid:      grid.records.length + 1,
                    TMPPETID:   0,
                    BID:        0,
                    ASMID:      0,
                    ARID:       0,
                    ARName:     "",
                    Amount:     0.0,
                };
                grid.add(rec);
                grid.select(rec.recid);
                grid.getColumn("ARName").editable.items = app.raflow.arW2UIItems;
                grid.getColumn("ARName").render();
                grid.refresh();
            },
        });

        //------------------------------------------------------------------------
        //          Pet Form Buttons
        //------------------------------------------------------------------------
        $().w2form({
            name: 'RAPetFormBtns',
            style: 'border: none; background-color: transparent;',
            formURL: '/webclient/html/raflow/formrapetbtns.html',
            url: '',
            fields: [],
            actions: {
                save: function() {
                    var f = w2ui.RAPetForm,
                        TMPPETID = f.record.TMPPETID;

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // validate form
                    var errors = f.validate();
                    if (errors.length > 0) return;

                    // update the modified data
                    SetlocalDataFromRAPetFormRecord(TMPPETID);

                    // save this records in json Data
                    savePetsCompData()
                    .done(function(data) {
                        if (data.status === 'success') {
                            // re-assign records in grid
                            AssignPetsGridRecords();

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
                saveadd: function() {
                    var f = w2ui.RAPetForm,
                        grid = w2ui.RAPetsGrid,
                        TMPPETID = f.record.TMPPETID;

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // validate form
                    var errors = f.validate();
                    if (errors.length > 0) return;

                    // update local data from this form record
                    SetlocalDataFromRAPetFormRecord(TMPPETID);

                    // save this records in json Data
                    savePetsCompData()
                    .done(function(data) {
                        if (data.status === 'success') {
                            // add new formatted record to current form
                            f.record = getPetFormInitRecord(f.record);
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

                    // if it exists then
                    if (itemIndex > -1) {

                        // remove locally
                        compData.splice(itemIndex, 1);

                        // save this records in json Data
                        savePetsCompData()
                        .done(function(data) {
                            if (data.status === 'success') {
                                // reset form
                                f.actions.reset();

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
                    }
                },
            },
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
                { type: 'top',     size: '60%', hidden: false, content: 'top',  resizable: true, style: app.pstyle },
                { type: 'main',    size: '40%', hidden: false, content: 'main', resizable: true, style: app.pstyle },
                { type: 'preview', size: 0,     hidden: true,  content: 'PREVIEW'  },
                { type: 'bottom',  size: 50,    hidden: false, content: 'bottom', resizable: false, style: app.pstyle },
                { type: 'right',   size: 0,     hidden: true }
            ]
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
        getAllARsWithAmount(BID)
        .done(function() {
            AssignPetFeesGridRecords(TMPPETID);
        });
    }, 0);
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

        // assign record in grid
        reassignGridRecids(grid.name);
    });

    // lock the grid until "Have pets?" checkbox checked.
    lockOnGrid(grid.name);
};

//------------------------------------------------------------------------------
// savePetsCompData - saves the data on server side
//------------------------------------------------------------------------------
window.savePetsCompData = function() {
    var compData = getRAFlowCompData("pets", app.raflow.activeFlowID);
    return saveActiveCompData(compData, "pets");
};

//-----------------------------------------------------------------------------
// getPetFeeLocalData - returns the clone of pet fee data for requested
//                      TMPPETID and ARID
//-----------------------------------------------------------------------------
window.getPetFeeLocalData = function(TMPPETID, ARID, returnIndex) {
    var cloneData = {};
    var foundIndex = -1;
    var compData = getRAFlowCompData("pets", app.raflow.activeFlowID) || [];
    compData.forEach(function(item, index) {
        if (item.TMPPETID == TMPPETID) {
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
// setPetFeeLocalData - save the data for requested a TMPPETID, ARID
//                   in local data
//-----------------------------------------------------------------------------
window.setPetFeeLocalData = function(TMPPETID, ARID, petFeeData) {
    var compData = getRAFlowCompData("pets", app.raflow.activeFlowID);
    var pIndex = -1,
        fIndex = -1;

    compData.forEach(function(item, itemIndex) {
        if (item.TMPPETID == TMPPETID) {
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
            compData[pIndex].Fees[fIndex] = petFeeData;
        } else {
            compData[pIndex].Fees.push(petFeeData);
        }
    }
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
    var petFeesData = [];

    // SPECIAL CASE if adding new pet //
    if (TMPPETID === 0) {
        // get initial fees from bizProps
        app.petFees[BID].forEach(function(bizPropPetFee) {
            petFeesData.push($.extend(true, {TMPPETID: 0, ASMID: 0}, bizPropPetFee));
        });
    } else {
        var petData = getPetLocalData(TMPPETID);
        petFeesData = petData.Fees;
    }

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

        // assign item prepared earlier for parent rentable list
        grid.getColumn("ARName").editable.items = app.raflow.arW2UIItems;
        grid.getColumn("ARName").render();
    });
};

