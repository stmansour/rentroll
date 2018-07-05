/* global
    RACompConfig, sliderContentDivLength, reassignGridRecids,
    hideSliderContent, appendNewSlider, showSliderContentW2UIComp,
    saveActiveCompData, getRAFlowCompData,
    GetFeeFormInitRecord, getInitialRentableFeesData,
    getRentableLocalData, setRentableLocalData, getAllARsWithAmount, GetRentableIndexInGridRecords,
    saveRentableCompData, setRentableFeeLocalData, getRentableFeeLocalData,
    ridRentablePickerRender, ridRentableDropRender, ridRentableCompare,
    AssignRentableGridRecords, AssignRentableFeesGridRecords,
    SetRentableAmountsFromFees, manageParentRentableW2UIItems,
    RenderRentablesGridSummary, GetFeeFormFields, GetFeeGridColumns
*/

"use strict";

// -------------------------------------------------------------------------------
// getAllARsWithAmount - pull down all account rules with amount, flags info
// -------------------------------------------------------------------------------
window.getAllARsWithAmount = function(BID) {
    var data = {"type": "ALL"};
    return $.ajax({
        url: '/v1/arslist/' + BID.toString() + "/",
        method: "POST",
        data: JSON.stringify(data),
        contentType: "application/json",
        dataType: "json"
    })
    .done(function(data) {
        if (data.success !== "error") {
            app.raflow.arList[BID] = data.records || [];
            app.raflow.arW2UIItems = [{id: 0, text: " -- select account rule -- " }];
            app.raflow.arList[BID].forEach(function(arItem) {
                app.raflow.arW2UIItems.push({id: arItem.ARID, text: arItem.Name});
            });
        }
    });
};


// -------------------------------------------------------------------------------
// getInitialRentableFeesData - pull down all fees records for the requested RID
// @params - RID
// -------------------------------------------------------------------------------
window.getInitialRentableFeesData = function(BID, RID, FlowID) {
    var data = {"RID": RID, "FlowID": FlowID};
    return $.ajax({
        url: "/v1/raflow-rentable-fees/" + BID.toString() + "/" + FlowID.toString(),
        method: "POST",
        contentType: "application/json",
        data: JSON.stringify(data)
    }).done(function(data) {
        if (data.status === "success") {
            // update the local copy of flow for the active one
            app.raflow.data[data.record.FlowID] = data.record;

            // set the rentable grid records again
            AssignRentableGridRecords();
        }
    });
};

window.loadRARentablesGrid = function () {

    // if form is loaded then return
    if (!("RARentablesGrid" in w2ui)) {

        // -----------------------------------------------------------
        // RENTABLE SEARCH FORM
        // -----------------------------------------------------------
        $().w2form({
            name: 'RARentableSearchForm',
            header: 'Rentable',
            style: 'display: block; border: none;',
            formURL: '/webclient/html/raflow/formra-rentable.html',
            focus: -1,
            fields: [
                {name: 'Rentable',      type: 'enum',   required: true,
                    options: {
                        url:           '/v1/rentablestd/' + app.raflow.BID,
                        max:           1,
                        cacheMax:      50,
                        maxDropHeight: 350,
                        renderItem:    function(item) {
                            // Enable Accept button
                            $(w2ui.RARentableSearchForm.box).find("button[name=accept]").prop("disabled", false);
                            w2ui.RARentableSearchForm.record.RID = item.RID;
                            w2ui.RARentableSearchForm.record.RentableName = item.RentableName;
                            return item.RentableName + '  (RID: ' + item.RID + ')';
                        },
                        renderDrop:    ridRentableDropRender,
                        compare:       ridRentableCompare,
                        onNew:         function (event) {
                            //console.log('++ New Item: Do not forget to submit it to the server too', event);
                            $.extend(event.item, { RentableName : event.item.text });
                        },
                        onRemove: function(event) {
                            event.onComplete = function() {
                                w2ui.RARentableSearchForm.actions.reset();
                            };
                        }
                    },
                },
                {name: 'BID',           type: 'int',    required: true, html: {caption: "BID"}},
                {name: 'RID',           type: 'int',    required: true, html: {caption: "RID"}},
                {name: 'RentableName',  type: 'int',    required: true, html: {caption: "RentableName"}},
            ],
            actions: {
                reset: function () {
                    w2ui.RARentableSearchForm.clear();
                    $(w2ui.RARentableSearchForm.box).find("button[name=accept]").prop("disabled", true);
                }
            },
            onRefresh: function (event) {
                var f = this;
                event.onComplete = function () {
                    var BID = getCurrentBID(),
                        BUD = getBUDfromBID(BID);

                    f.record.BID = BID;
                };
            }
        });

        // -----------------------------------------------------------
        // RENTABLES GRID
        // -----------------------------------------------------------
        $().w2grid({
            name: 'RARentablesGrid',
            header: 'Rentables',
            show: {
                toolbar:    false,
                footer:     true,
            },
            multiSelect: false,
            style: 'display: block;',
            columns: [
                {
                    field: 'recid',
                    hidden: true
                },
                {
                    field: 'RID',
                    hidden: true
                },
                {
                    field: 'BID',
                    hidden: true
                },
                {
                    field: 'RTID',
                    hidden: true
                },
                {
                    field: 'RTFLAGS',
                    hidden: true
                },
                {
                    field: 'RentableName',
                    caption: 'Rentable',
                    size: '160px'
                },
                {
                    field: 'RentCycle',
                    hidden: true
                },
                {
                    field: 'RentCycleText',
                    caption: 'RentCycle',
                    size: '100px',
                    render: function (record) {
                        return app.cycleFreq[record.RentCycle];
                    }
                },
                {
                    field: 'AtSigningPreTax',
                    caption: 'At Signing<br>(pre-tax)',
                    size: '90px',
                    render: 'money'
                },
                {
                    field: 'SalesTax',
                    caption: 'Sales Tax',
                    size: '100px',
                    render: 'money'
                },
                /*{ // FUTURE RELEASE
                    field: 'SalesTaxAmt',
                    caption: 'Sales Tax Amt',
                    size: '100px',
                    render: 'money'
                },*/
                {
                    field: 'TransOccTax',
                    caption: 'Trans Occ Tax',
                    size: '100px',
                    render: 'money'
                },
                /*{ // FUTURE RELEASE
                    field: 'TransOccAmt',
                    caption: 'Trans Occ Amt',
                    size: '100px',
                    render: 'money'
                },*/
                {
                    field: 'RemoveRec',
                    caption: "Remove Rentable",
                    size: '100%',
                    render: function (record/*, index, col_index*/) {
                        var html = "";
                        if (record.RID && record.RID > 0) {
                            html = '<i class="fas fa-minus-circle" style="color: #DC3545; cursor: pointer;" title="remove rentable"></i>';
                        }
                        return html;
                    },
                }
            ],
            onClick: function (event) {
                event.onComplete = function () {
                    // if it's remove column then remove the record
                    // maybe confirm dialog will be added
                    if(w2ui.RARentablesGrid.getColumn("RemoveRec", true) == event.column) {
                        // remove entry from local data
                        var rec = w2ui.RARentablesGrid.get(event.recid);
                        var index = getRentableLocalData(rec.RID, true);

                        // also manage local data
                        var compData = getRAFlowCompData("rentables", app.raflow.activeFlowID);
                        compData.splice(index, 1);

                        // save data on server
                        saveRentableCompData()
                        .done(function(data) {
                            if (data.status === "success") {
                                // after saving rentable comp data re-calculate parent rentable
                                // w2ui items list
                                manageParentRentableW2UIItems();

                                // remove from grid
                                w2ui.RARentablesGrid.remove(event.recid);
                            }
                        });
                        return;
                    }

                    var yes_args = [w2ui.RARentablesGrid, event.recid],
                        no_args = [w2ui.RARentablesGrid],
                        no_callBack = function(grid) {
                            grid.select(app.last.grid_sel_recid);
                            return false;
                        },
                        yes_callBack = function(grid, recid) {
                            app.last.grid_sel_recid = parseInt(recid);

                            // keep highlighting current row in any case
                            grid.select(app.last.grid_sel_recid);

                            // get auto populated to new RA account rules
                            var rec = grid.get(recid);
                            var localRData = getRentableLocalData(rec.RID);

                            // just render the record from local data if fees are available
                            if(localRData.hasOwnProperty("Fees") && localRData.Fees.length > 0) {
                                // set fees grid records
                                AssignRentableFeesGridRecords(rec.RID);

                                // show slider content
                                showSliderContentW2UIComp(w2ui.RARentableFeesGrid, RACompConfig.rentables.sliderWidth);
                            } else {
                                // pull fees in case it's empty
                                var BID = getCurrentBID();
                                getInitialRentableFeesData(BID, rec.RID, app.raflow.activeFlowID)
                                .done(function(data) {
                                    if (data.status === "success") {
                                        // re-render fees grid records
                                        AssignRentableFeesGridRecords(rec.RID);

                                        // show the slider content
                                        showSliderContentW2UIComp(w2ui.RARentableFeesGrid, RACompConfig.rentables.sliderWidth);
                                    }
                                })
                                .fail(function(data) {
                                    console.log("ERROR from fees data: " + data);
                                });
                            }
                        };

                    // warn user if content has been changed
                    form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
                };
            }
        });

        // -----------------------------------------------------------
        // RENTABLE ***** FEES ***** GRID
        // -----------------------------------------------------------
        var feeGridColumns = GetFeeGridColumns();

        $().w2grid({
            name: 'RARentableFeesGrid',
            header: 'Rentables Fees',
            show: {
                toolbar: true,
                footer: true,
                header: true,
                searchAll: false,
                toolbarSearch   : false,
                toolbarInput    : false,
            },
            style: 'border: 2px solid white; display: block;',
            toolbar: {
                items: [
                    {id: 'add', type: 'button', caption: 'Add Record', icon: 'w2ui-icon-plus'},
                    {id: 'bt3', type: 'spacer'},
                    {id: 'btnClose', type: 'button', icon: 'fas fa-times'},
                ],
                onClick: function (event) {
                    switch(event.target) {
                        case "add":
                            var sliderID = 2;
                            appendNewSlider(sliderID);
                            $("#raflow-container")
                                .find(".slider[data-slider-id="+sliderID+"]")
                                .find(".slider-content")
                                .width(400)
                                .w2render(w2ui.RARentableFeeForm);

                            // keep highlighting current row in any case
                            w2ui.RARentableFeesGrid.selectNone();

                            var RID = w2ui.RARentablesGrid.get(w2ui.RARentablesGrid.getSelection()[0]).RID;

                            var BID = getCurrentBID();
                            getAllARsWithAmount(BID)
                            .done(function(data) {
                                var arid_items = [];
                                app.raflow.arList[BID].forEach(function(item) {
                                    arid_items.push({id: item.ARID, text: item.Name});
                                });
                                w2ui.RARentableFeeForm.get("ARID").options.items = arid_items;
                                w2ui.RARentableFeeForm.record = GetFeeFormInitRecord();
                                w2ui.RARentableFeeForm.record.recid = w2ui.RARentableFeesGrid.records.length + 1;

                                showSliderContentW2UIComp(w2ui.RARentableFeeForm, sliderContentDivLength, sliderID);
                                w2ui.RARentableFeeForm.refresh();
                            })
                            .fail(function(data) {
                                console.log("failure" + data);
                            });
                            break;
                        case "btnClose":
                            hideSliderContent();
                            break;
                    }
                }
            },
            columns: GetFeeGridColumns(),
            onChange: function (event) {
                event.onComplete = function () {
                    this.save();
                };
            },
            onClick: function(event) {
                event.onComplete = function() {
                    var form = w2ui.RARentableFeeForm;
                    var yes_args = [this, event.recid],
                        no_args = [this],
                        no_callBack = function(grid) {
                            grid.select(app.last.grid_sel_recid);
                            return false;
                        },
                        yes_callBack = function(grid, recid) {
                            var sliderID = 2;
                            appendNewSlider(sliderID);
                            $("#raflow-container")
                                .find(".slider[data-slider-id="+sliderID+"]")
                                .find(".slider-content")
                                .width(400)
                                .w2render(form);

                            app.last.grid_sel_recid = parseInt(recid);

                            // keep highlighting current row in any case
                            grid.select(app.last.grid_sel_recid);

                            var BID = getCurrentBID();
                            getAllARsWithAmount(BID)
                            .done(function(data) {
                                var arid_items = [];
                                app.raflow.arList[BID].forEach(function(item) {
                                    arid_items.push({id: item.ARID, text: item.Name});
                                });
                                form.get("ARID").options.items = arid_items;
                                form.record = $.extend(true, {}, grid.get(app.last.grid_sel_recid));

                                form.record.RentCycleText = app.cycleFreq[form.record.RentCycle];

                                showSliderContentW2UIComp(form, sliderContentDivLength, sliderID);
                                form.refresh(); // need to refresh for header changes

                                // When RentCycle is Norecur then disable the RentCycle list field.
                                var isDisabled = form.record.RentCycleText.text === app.cycleFreq[0];
                                $("#RentCycleText").prop("disabled", isDisabled);
                            })
                            .fail(function(data) {
                                console.log("failure" + data);
                            });
                        };

                    // warn user if form content has been changed
                    form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
                };
            }
        });

        // -----------------------------------------------------------
        //  ***** RENTABLE ***** FEE ***** FORM *****
        // -----------------------------------------------------------
        var feeFormFields = GetFeeFormFields();

        $().w2form({
            name: 'RARentableFeeForm',
            header: 'Add New Rentable Fee',
            style: 'display: block;',
            formURL: '/webclient/html/raflow/formra-fee.html',
            focus: -1,
            fields: feeFormFields,
            toolbar : {
                items: [
                    { id: 'btnNotes', type: 'button', icon: 'far fa-sticky-note' },
                    { id: 'bt3', type: 'spacer' },
                    { id: 'btnClose', type: 'button', icon: 'fas fa-times'}
                ],
                onClick: function (event) {
                    switch (event.target){
                        case 'btnClose':
                            hideSliderContent(2);
                            break;
                    }
                }
            },
            actions: {
                reset: function () {
                    w2ui.RARentableFeeForm.clear();
                },
                save: function() {
                    var f           = this,
                        grid        = w2ui.RARentableFeesGrid,
                        TMPASMID    = f.record.TMPASMID;

                    // get RID based on rentable grid selection
                    var rentableGridRecid = w2ui.RARentablesGrid.getSelection()[0];
                    var gridRec = w2ui.RARentablesGrid.get(rentableGridRecid);
                    var RID = gridRec.RID;

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // sync this info in local data
                    var rentableFeeData = getFormSubmitData(f.record, true);

                    // set data locally
                    setRentableFeeLocalData(RID, TMPASMID, rentableFeeData);

                    // re-calculate amounts for rentable
                    SetRentableAmountsFromFees(RID);

                    saveRentableCompData()
                    .done(function (data) {
                        if (data.status === 'success') {
                            // Re render the fees grid records
                            AssignRentableFeesGridRecords(RID);

                            // reset the form
                            f.actions.reset();

                            // close the form
                            hideSliderContent(2);
                        } else {
                            f.message(data.message);
                        }
                    })
                    .fail(function (data) {
                        console.log("failure " + data);
                    });
                },
                saveadd: function() {
                    var f           = this,
                        grid        = w2ui.RARentableFeesGrid,
                        TMPASMID    = f.record.TMPASMID;

                    // get RID based on rentable grid selection
                    var rentableGridRecid = w2ui.RARentablesGrid.getSelection()[0];
                    var gridRec = w2ui.RARentablesGrid.get(rentableGridRecid);
                    var RID = gridRec.RID;

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // sync this info in local data
                    var rentableFeeData = getFormSubmitData(f.record, true);

                    // set data locally
                    setRentableFeeLocalData(RID, TMPASMID, rentableFeeData);

                    // re-calculate amounts for rentable
                    SetRentableAmountsFromFees(RID);

                    saveRentableCompData()
                    .done(function (data) {
                        if (data.status === 'success') {
                            f.actions.reset();
                            f.record = GetFeeFormInitRecord();
                            f.record.recid = grid.records.length + 1;
                            f.refresh();

                            // enable this field
                            $(f.box).find("#RentCycleText").prop("disabled", false);

                            // Re render the fees grid records
                            AssignRentableFeesGridRecords(RID);
                        } else {
                            f.message(data.message);
                        }
                    })
                    .fail(function (data) {
                        console.log("failure " + data);
                    });
                },
                delete: function() {
                    var f           = this,
                        grid        = w2ui.RARentableFeesGrid,
                        TMPASMID    = f.record.TMPASMID;

                    // get RID based on rentable grid selection
                    var rentableGridRecid = w2ui.RARentablesGrid.getSelection()[0];
                    var gridRec = w2ui.RARentablesGrid.get(rentableGridRecid);
                    var RID = gridRec.RID;

                    var localRData = getRentableLocalData(RID);
                    if (localRData.Fees.length > 0) {
                        var itemIndex = getRentableFeeLocalData(RID, TMPASMID, true);

                        // remove fee item
                        localRData.Fees.splice(itemIndex, 1);

                        // set this modified local rentable data to back
                        setRentableLocalData(RID, localRData);

                        // re-calculate amounts for rentable
                        SetRentableAmountsFromFees(RID);

                        // sync data on backend side
                        saveRentableCompData()
                        .done(function (data) {
                            if (data.status === 'success') {
                                // reset form as well as remove record from the grid
                                grid.remove(RID);
                                grid.refresh();
                                f.actions.reset();

                                // // Re render the fees grid records
                                AssignRentableFeesGridRecords(RID);

                                // close the form
                                hideSliderContent(2);
                            } else {
                                f.message(data.message);
                            }
                        })
                        .fail(function (data) {
                            console.log("failure " + data);
                        });
                    }
                }
            },
            onChange: function(event) {
                var f = this;
                event.onComplete = function() {
                    switch(event.target) {
                        case "RentCycleText":
                            if (event.value_new) {
                                app.cycleFreq.forEach(function(itemText, itemIndex) {
                                    if (event.value_new.text == itemText) {
                                        f.record.RentCycle = itemIndex;
                                        return false;
                                    }
                                });
                                f.refresh();
                            }
                            break;
                        case "ARID":
                            if (event.value_new) {
                                var BID = getCurrentBID();

                                // find account rules
                                var arItem = {};
                                app.raflow.arList[BID].forEach(function(item) {
                                    if (event.value_new.id == item.ARID) {
                                        arItem = item;
                                        return false;
                                    }
                                });

                                // update form record based on selected account rules item
                                f.record.ContractAmount = arItem.DefaultAmount;
                                f.record.ARName = event.value_new.text;

                                // check for non-recurring cycle flag
                                if (arItem.FLAGS&0x40 != 0) { // then it is set to non-recur flag
                                    // It indicates that rule follow non recur charge
                                    // f.record.RentCycleText = app.cycleFreq[0];
                                    f.record.RentCycle = 0;
                                } else {
                                    var rentableGrid = w2ui.RARentablesGrid;
                                    var rentableGridRecid = rentableGrid.getSelection()[0];
                                    var gridRec = rentableGrid.get(rentableGridRecid);
                                    // f.record.RentCycleText = app.cycleFreq[record.RentCycle];
                                    f.record.RentCycle = gridRec.RentCycle;
                                }

                                // select rentcycle as well
                                var selectedRentCycle = app.cycleFreq[f.record.RentCycle];
                                var rentCycleW2UISel = { id: selectedRentCycle, text: selectedRentCycle };
                                f.get("RentCycleText").options.selected = rentCycleW2UISel;
                                f.record.RentCycleText = rentCycleW2UISel;
                                f.refresh();

                                // When RentCycle is Norecur then disable the RentCycle list field.
                                var isDisabled = f.record.RentCycleText.text === app.cycleFreq[0];
                                $(f.box).find("#RentCycleText").prop("disabled", isDisabled);
                            }
                            break;
                    }

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
                var f = this;
                event.onComplete = function() {
                    // ARID value
                    var ARIDSel = {};
                    f.get("ARID").options.items.forEach(function(item) {
                        if (item.id == f.record.ARID) {
                            ARIDSel = {id: item.id, text: item.text};
                        }
                    });
                    f.get("ARID").options.selected = ARIDSel;

                    // rent cycle text value
                    var selectedRentCycle = app.cycleFreq[f.record.RentCycle];
                    var RentCycleTextSel = { id: selectedRentCycle, text: selectedRentCycle };
                    f.get("RentCycleText").options.selected = RentCycleTextSel;
                    f.record.RentCycleText = RentCycleTextSel;

                    // if RentCycle is 0=nonrecur then disable Stop date field
                    // and value should be same as Start
                    if (f.record.RentCycle === 0) {
                        $(f.box).find("input[name=Stop]").prop("disabled", true);
                        $(f.box).find("input[name=Stop]").w2field().set(f.record.Start);
                        f.record.Stop = f.record.Start;
                    } else {
                        $(f.box).find("input[name=Stop]").prop("disabled", false);
                    }

                    // there is NO PETID actually, so have to work around with recid key
                    formRefreshCallBack(f);

                    // set header
                    var rentable_sel = w2ui.RARentablesGrid.getSelection();
                    if (rentable_sel.length > 0) {
                        var rec = w2ui.RARentablesGrid.get(rentable_sel[0]);
                        var header = "Fee ({0}) for {1}";
                        if (f.record.ARID > 0) {
                            f.header = header.format(f.record.ARName, rec.RentableName);
                        } else {
                            f.header = header.format("new", rec.RentableName);
                        }
                    }
                };
            }
        });
    }

    // now load grid in division
    $('#ra-form #rentables .grid-container').w2render(w2ui.RARentablesGrid);
    $('#ra-form #rentables .form-container').w2render(w2ui.RARentableSearchForm);

    // load the existing data in rentables component
    setTimeout(function () {
        AssignRentableGridRecords();
    }, 500);
};

//-----------------------------------------------------------------------------
// AssignRentableGridRecords - will set the rentable grid records from local
//                               copy of flow data
//-----------------------------------------------------------------------------
window.AssignRentableGridRecords = function() {
    var compData = getRAFlowCompData("rentables", app.raflow.activeFlowID);
    if (compData) {
        w2ui.RARentablesGrid.records = compData;
        reassignGridRecids(w2ui.RARentablesGrid.name);

        // Operation on RARentableSearchForm
        w2ui.RARentableSearchForm.refresh();

        // manage parent rentables list
        manageParentRentableW2UIItems();

        // Render RentableGrid Summary
        RenderRentablesGridSummary();

    } else {
        w2ui.RARentablesGrid.clear();
        // Operation on RARentableSearchForm
        w2ui.RARentableSearchForm.actions.reset();
    }
};

//-----------------------------------------------------------------------------
// SetRentableAmountsFromFees - set the all amounts to rentable locally
//                              calculated from fees list associated with
//                              requested RID
//
// @params
//    RID = RentableID
//-----------------------------------------------------------------------------
window.SetRentableAmountsFromFees = function(RID) {
    // get the local data again after new data has been set
    var localRData = getRentableLocalData(RID);

    if (!localRData) {
        console.error("local rentable data not found for RID: ", RID);
        return;
    }

    // temp variable to hold the summing up figure for all amounts
    var amountsSum = {
        AtSigningPreTax:    0.0,
        SalesTax:           0.0,
        // SalesTaxAmt:        0.0,
        TransOccTax:        0.0,
        // TransOccAmt:        0.0,
    };

    // iterate over each Fees record
    localRData.Fees.forEach(function(feeData) {
        amountsSum.AtSigningPreTax += feeData.AtSigningPreTax;
        amountsSum.SalesTax += feeData.SalesTax;
        // amountsSum.SalesTaxAmt += feeData.SalesTaxAmt;
        amountsSum.TransOccTax += feeData.TransOccTax;
        // amountsSum.TransOccAmt += feeData.TransOccAmt;
    });

    // set the amount to rentable
    localRData.AtSigningPreTax = amountsSum.AtSigningPreTax;
    localRData.SalesTax = amountsSum.SalesTax;
    // localRData.SalesTaxAmt = amountsSum.SalesTaxAmt;
    localRData.TransOccTax = amountsSum.TransOccTax;
    // localRData.TransOccAmt = amountsSum.TransOccAmt;

    // save this modified rentable data
    setRentableLocalData(RID, localRData);
};

//-----------------------------------------------------------------------------
// RenderRentablesGridSummary - will render grid summary row from rentable
//                             comp data
//-----------------------------------------------------------------------------
window.RenderRentablesGridSummary = function() {
    var compData = getRAFlowCompData("rentables", app.raflow.activeFlowID) || [];
    var grid = w2ui.RARentablesGrid;

    // summary record in fees grid
    var summaryRec = {
        recid:              0,
        RentableName:       "Grand Total",
        AtSigningPreTax:    0.0,
        SalesTax:           0.0,
        // SalesTaxAmt:     0.0,
        TransOccTax:        0.0,
        // TransOccAmt:     0.0,
        RemoveRec:          null,
    };

    compData.forEach(function(rentableItem) {
        summaryRec.AtSigningPreTax += rentableItem.AtSigningPreTax;
        summaryRec.SalesTax += rentableItem.SalesTax;
        // summaryRec.SalesTaxAmt += rentableItem.SalesTaxAmt;
        summaryRec.TransOccTax += rentableItem.TransOccTax;
        // summaryRec.TransOccAmt += rentableItem.TransOccAmt;
    });

    // set the summary rec in summary array of grid
    grid.summary = [summaryRec];

    // refresh the grid
    grid.refresh();
};

//-----------------------------------------------------------------------------
// AssignRentableFeesGridRecords - will set the rentable Fees grid records
//                                   from local copy of flow data again and
//                                   set the summary row record in the grid
//
// @params
//    RID = RentableID
//-----------------------------------------------------------------------------
window.AssignRentableFeesGridRecords = function(RID) {
    var grid = w2ui.RARentableFeesGrid;

    // get the local data again after new data has been set
    var localRData = getRentableLocalData(RID);

    // set the records list
    grid.records = localRData.Fees || [];

    // set the header as well
    grid.header = "Fees for " + localRData.RentableName;

    // summary record in fees grid
    var summaryRec = {
        recid:              0,
        ARName:             "Grand Total",
        ContractAmount:     0.0,
        AtSigningPreTax:    0.0,
        SalesTax:           0.0,
        // SalesTaxAmt:        0.0,
        TransOccTax:        0.0,
        // TransOccAmt:        0.0,
    };

    // calculate amount for summary row
    grid.records.forEach(function(item) {
        // summaryRec.ContractAmount += item.ContractAmount;
        summaryRec.AtSigningPreTax += item.AtSigningPreTax;
        summaryRec.SalesTax += item.SalesTax;
        // summaryRec.SalesTaxAmt += item.SalesTaxAmt;
        summaryRec.TransOccTax += item.TransOccTax;
        // summaryRec.TransOccAmt += item.TransOccAmt;
    });

    // set the summary rec in summary array of grid
    grid.summary = [summaryRec];

    // reassign records id in feees grid and refresh it
    reassignGridRecids(grid.name);

    // set the summarized value in rentable grid too
    var rentablesGridRecords = w2ui.RARentablesGrid.records || [];
    var foundRIDIndex = -1;
    rentablesGridRecords.forEach(function(gRec, index) {
        if (gRec.RID == RID) {
            foundRIDIndex = index + 1; // we've reassigned recid which starts from 1, not 0
            return false;
        }
    });

    if (foundRIDIndex > -1) {
        var rentableGridRec = w2ui.RARentablesGrid.get(foundRIDIndex);
        rentableGridRec.AtSigningPreTax = summaryRec.AtSigningPreTax;
        rentableGridRec.SalesTax = summaryRec.SalesTax;
        // rentableGridRec.SalesTaxAmt = summaryRec.SalesTaxAmt;
        rentableGridRec.TransOccTax = summaryRec.TransOccTax;
        // rentableGridRec.TransOccAmt = summaryRec.TransOccAmt;

        // set the modified data in grid back
        w2ui.RARentablesGrid.set(foundRIDIndex, rentableGridRec);
        w2ui.RARentablesGrid.refresh();
    }

    // render rentable grid summary record
    RenderRentablesGridSummary();
};

//-----------------------------------------------------------------------------
// acceptRentable - add Rentable to the list rentables grid records
//-----------------------------------------------------------------------------
window.acceptRentable = function () {
    var recIndex = GetRentableIndexInGridRecords(w2ui.RARentableSearchForm.record.RID);
    var BID = getCurrentBID();

    if(recIndex > -1 ) {
        w2ui.RARentablesGrid.select(recIndex); // highlight the existing record
    } else {
        var fRec = w2ui.RARentableSearchForm.record;
        getInitialRentableFeesData(BID, fRec.RID, app.raflow.activeFlowID)
        .done(function(data) {

            if (data.status === "success") {
                // reset the form
                w2ui.RARentableSearchForm.actions.reset();
            }
        })
        .fail(function(data) {
            console.log("ERROR from fees data: " + data);
        });
    }
};

//------------------------------------------------------------------------------
// manageParentRentableW2UIItems - maintain parent rentable w2ui items list
//------------------------------------------------------------------------------
window.manageParentRentableW2UIItems = function() {

    // reset it first
    app.raflow.parentRentableW2UIItems = [];

    // inner function to push item in "app.raflow.parentRentableW2UIItems"
    var pushItem = function(rentableItem, atIndex) {
        var found = false;
        app.raflow.parentRentableW2UIItems.forEach(function(item) {
            if (item.id === rentableItem.id) {
                found = true;
                return false;
            }
        });

        // if not found the push item in app.raflow.parentRentableW2UIItems
        if (!found) {
            app.raflow.parentRentableW2UIItems.splice(atIndex, 0, rentableItem);
        }
    };

    // get comp data
    var rentableCompData = getRAFlowCompData("rentables", app.raflow.activeFlowID) || [];

    // first build the list of parent rentables and sort it out in asc order of RID
    rentableCompData.forEach(function(rentableItem) {
        var RID = rentableItem.RID,
            RentableName = rentableItem.RentableName;

        var childRentableFLAG = (rentableItem.RTFLAGS & (1 << app.rtFLAGS.IsChildRentable));

        if ( childRentableFLAG === 0) { // 0 means it is not child, it is parent
            var item = {id: RID, text: RentableName};
            pushItem(item, app.raflow.parentRentableW2UIItems.length);
        }
    });

    // sort it out in asc order of RID value
    app.raflow.parentRentableW2UIItems.sort(function(a, b) {
        return a.id - b.id;
    });

    // if there is only one parent rentable then pre-select it for all child rentable
    // otherwise built drop down menu
    if (app.raflow.parentRentableW2UIItems.length != 1) {
        var item = {id: 0, text: " -- select rentable -- "};
        pushItem(item, 0);
    } else {
        app.raflow.parentRentableW2UIItems.forEach(function(item, index) {
            if (item.id === 0) {
                app.raflow.parentRentableW2UIItems.splice(index, 1);
            }
        });
    }
};

//------------------------------------------------------------------------------
// saveRentableCompData - saves the data on server side
//------------------------------------------------------------------------------
window.saveRentableCompData = function() {
    var compData = getRAFlowCompData("rentables", app.raflow.activeFlowID);
    return saveActiveCompData(compData, "rentables");
};

//------------------------------------------------------------------------------
// GetRentableIndexInGridRecords - returns record index in grid records list
//------------------------------------------------------------------------------
window.GetRentableIndexInGridRecords = function(RID) {
    var found = -1;
    w2ui.RARentablesGrid.records.forEach(function(rec, index) {
        if (RID == rec.RID) {
            found = index;
            return false;
        }
    });
    return found;
};

//-----------------------------------------------------------------------------
// getRentableLocalData - returns the clone of rentable data for requested
//                        RID
//-----------------------------------------------------------------------------
window.getRentableLocalData = function(RID, returnIndex) {
    var cloneData = {};
    var foundIndex = -1;
    var compData = getRAFlowCompData("rentables", app.raflow.activeFlowID);
    compData.forEach(function(item, index) {
        if (item.RID == RID) {
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
// setRentableLocalData - save the data for requested RID in local data
//-----------------------------------------------------------------------------
window.setRentableLocalData = function(RID, rentableData) {
    var compData = getRAFlowCompData("rentables", app.raflow.activeFlowID);
    var dataIndex = -1;
    compData.forEach(function(item, index) {
        if (item.RID == RID) {
            dataIndex = index;
            return false;
        }
    });

    // if rentable has no property of Fees then
    if (!rentableData.hasOwnProperty("Fees")) {
        rentableData.Fees = [];
    }

    if (dataIndex > -1) {
        compData[dataIndex] = rentableData;
    } else {
        compData.push(rentableData);
    }
};

//-----------------------------------------------------------------------------
// getRentableFeeLocalData - returns the clone of rentable fee data for requested
//                           RID, TMPASMID from "Fees" list of a rentable
//-----------------------------------------------------------------------------
window.getRentableFeeLocalData = function(RID, TMPASMID, returnIndex) {
    var cloneData = {};
    var foundIndex = -1;
    var compData = getRAFlowCompData("rentables", app.raflow.activeFlowID);
    compData.forEach(function(item) {
        if (item.RID == RID) {
            var feesData = item.Fees || [];
            feesData.forEach(function(feeItem, index) {
                if (feeItem.TMPASMID == TMPASMID) {
                    if (returnIndex) {
                        foundIndex = index;
                    } else {
                        cloneData = $.extend(true, {}, feeItem);
                    }
                }
                return false;
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
// setRentableFeeLocalData - save the data for Fee with
//                           requested RID in local data
//-----------------------------------------------------------------------------
window.setRentableFeeLocalData = function(RID, TMPASMID, rentableFeeData) {
    console.debug(RID, TMPASMID, rentableFeeData);

    var compData = getRAFlowCompData("rentables", app.raflow.activeFlowID);
    var rIndex = -1,
        fIndex = -1;

    compData.forEach(function(item, itemIndex) {
        if (item.RID == RID) {
            var feesData = item.Fees || [];
            feesData.forEach(function(feeItem, feeItemIndex) {
                if (feeItem.TMPASMID == TMPASMID) {
                    fIndex = feeItemIndex;
                }
                return false;
            });
            rIndex = itemIndex;
            return false;
        }
    });

    // only if rentable found then
    if (rIndex > -1) {
        if (fIndex > -1) {
            compData[rIndex].Fees[fIndex] = rentableFeeData;
        } else {
            compData[rIndex].Fees.push(rentableFeeData);
        }
    }
};

