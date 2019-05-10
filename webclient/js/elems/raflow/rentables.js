/* global
    RAFlowAJAX,
    RACompConfig, SliderContentDivLength, reassignGridRecids,
    HideSliderContent, appendNewSlider, ShowSliderContentW2UIComp,
    SaveCompDataAJAX, GetRAFlowCompLocalData,
    GetFeeFormInitRecord, SaveRAFlowRentableAJAX,
    GetRentableLocalData, SetRentableLocalData, GetAllARForFeeForm,
    SaveRentableCompData, SetRentableFeeLocalData, GetRentableFeeLocalData,
    ridRentablePickerRender, ridRentableDropRender, ridRentableCompare,
    AssignRentableGridRecords, AssignRentableFeesGridRecords,
    SetRentableAmountsFromFees, manageParentRentableW2UIItems, getRecIDFromTMPASMID,
    RenderRentablesGridSummary, GetFeeFormFields, GetFeeGridColumns, getFeeIndex,
    SetFeeDataFromFeeFormRecord, SetFeeFormRecordFromFeeData, displayRARentableFeesGridError,
    FeeFormOnChangeHandler, GetFeeFormToolbar, FeeFormOnRefreshHandler, getRecIDFromRID, displayFormFieldsError,
    GetFeeAccountRulesW2UIListItems, RenderFeesGridSummary, updateFlowData, dispalyRARentablesGridError,
    displayRARentableFeeFormError, getRentableIndex, displayFormFieldsError,
    GetCurrentFlowID, EnableDisableRAFlowVersionInputs, ShowHideGridToolbarAddButton,
    HideAllSliderContent, displayNonFieldsError, RemoveRAFlowRentableAJAX
*/

"use strict";

// -------------------------------------------------------------------------------
// SaveRAFlowRentableAJAX - pull down all fees records for the requested RID
// @params - RID
// -------------------------------------------------------------------------------
window.SaveRAFlowRentableAJAX = function(RID) {
    var BID = getCurrentBID(),
        FlowID = GetCurrentFlowID();

    var url = "/v1/raflow-rentable/" + BID.toString() + "/" + FlowID.toString() + "/";
    var data = {
        "cmd": "save",
        "RID": RID,
        "FlowID": FlowID
    };

    return RAFlowAJAX(url, "POST", data, true)
    .done(function(data) {
        if (data.status !== "error") {
            // set the rentable grid records again
            AssignRentableGridRecords();
        }
    });
};

// -------------------------------------------------------------------------------
// RemoveRAFlowRentableAJAX - remove rentables
// @params - RID
// -------------------------------------------------------------------------------
window.RemoveRAFlowRentableAJAX = function (RID) {
    var BID = getCurrentBID(),
        FlowID = GetCurrentFlowID();

    var url = "/v1/raflow-rentable/" + BID.toString() + "/" + FlowID.toString() + "/";
    var data = {
        "cmd": "delete",
        "RID": RID,
        "FlowID": FlowID
    };

    return RAFlowAJAX(url, "POST", data, true)
    .done(function(data) {
        if (data.status !== "error") {
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
            formURL: '/webclient/html/raflow/formra-rentablesearch.html',
            focus: -1,
            fields: [
                {name: 'Rentable',      type: 'enum',   required: true,
                    options: {
                        url:           '/v1/rentablestd/' + getCurrentBID().toString(),
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

                    // FREEZE THE INPUTS IF VERSION IS RAID
                    EnableDisableRAFlowVersionInputs(f);
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
            style: 'border-color: silver; border-style: solid; border-width: 1px 0 0 0; display: block;',
            columns: [
                { field: 'recid', hidden: true },
                { field: 'RID', hidden: true },
                { field: 'RTID', hidden: true },
                { field: 'RTFLAGS', hidden: true },
                { field: 'haveError', size: '30px', hidden: false,
                    render: function (record) {
                        var haveError = false;
                        if (app.raflow.validationErrors.rentables) {
                            var rentables = app.raflow.validationCheck.errors.rentables.errors;
                            for (var i = 0; i < rentables.length; i++) {
                                if (rentables[i].RID === record.RID && rentables[i].total > 0) {
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
                { field: 'RentableName', caption: 'Rentable', size: '100%' },
                { field: 'RentCycle', hidden: true },
                { field: 'RentCycleText', caption: 'RentCycle', size: '100px',
                    render: function (record) {
                        return app.cycleFreq[record.RentCycle];
                    }
                },
                { field: 'AtSigningPreTax', caption: 'At Signing<br>(pre-tax)', size: '100px', render: 'money' },
                { field: 'SalesTax', caption: 'Sales Tax', size: '100px', render: 'money' },
                /*{ // FUTURE RELEASE field: 'SalesTaxAmt', caption: 'Sales Tax Amt', size: '100px', render: 'money' },*/
                { field: 'TransOccTax', caption: 'Trans Occ Tax', size: '100px', render: 'money' },
                /*{ // FUTURE RELEASE field: 'TransOccAmt', caption: 'Trans Occ Amt', size: '100px', render: 'money' },*/
                { field: 'RowTotal', caption: 'Grand Total', size: '100px', style: 'text-align: right',
                    render: function(record) {
                        var html = "";
                        var total = 0.0;
                        if (record) {
                            if (record.AtSigningPreTax) {
                                total += record.AtSigningPreTax;
                            }
                            if (record.SalesTax) {
                                total += record.SalesTax;
                            }
                            if (record.TransOccTax) {
                                total += record.TransOccTax;
                            }

                            // make it bold
                            html = "<strong>" + w2utils.formatters.money(total) + "</strong>";
                        }
                        return html;
                    }
                },
                {
                    field: 'RemoveRec',
                    caption: "Remove<br>Rentable",
                    size: '90px',
                    render: function (record/*, index, col_index*/) {
                        // SPECIAL CHECK FOR THIS REMOVE BUTTON
                        if (app.raflow.version === "raid") {
                            return;
                        }
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
                        var rec = w2ui.RARentablesGrid.get(event.recid);

                        RemoveRAFlowRentableAJAX(rec.RID)
                        .done(function(data) {
                            if (data.status === "success") {
                                // after removing rentable comp data re-calculate parent rentable
                                // w2ui items list
                                manageParentRentableW2UIItems();
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

                            // keep this clicked rentable in last object
                            app.raflow.last.RID = rec.RID;

                            // get local data of this rentable
                            var localRData = GetRentableLocalData(rec.RID);

                            // set fees grid records
                            AssignRentableFeesGridRecords(rec.RID);


                            // show slider content
                            ShowSliderContentW2UIComp(w2ui.RARentableFeesGrid, RACompConfig.rentables.sliderWidth);
                        };

                    // warn user if content has been changed
                    form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
                };
            }
        });

        // -----------------------------------------------------------
        // RENTABLE ***** FEES ***** GRID
        // -----------------------------------------------------------
        $().w2grid({
            name: 'RARentableFeesGrid',
            header: 'Rentables Fees',
            show: {
                toolbar:        true,
                header:         true,
                toolbarSearch:  false,
                toolbarAdd:     true,
                toolbarReload:  false,
                toolbarInput:   false,
                toolbarColumns: false,
                footer:         false
            },
            style: 'border: none; display: block;',
            columns: GetFeeGridColumns('RARentableFeesGrid'),
            toolbar: {
                items: [
                    {id: 'bt3', type: 'spacer'},
                    {id: 'btnClose', type: 'button', icon: 'fas fa-times'}
                ],
                onClick: function (event) {
                    switch(event.target) {
                        case "btnClose":
                            HideSliderContent();
                            // unselect selected record
                            w2ui.RARentablesGrid.selectNone();
                            AssignRentableGridRecords();
                            break;
                    }
                }
            },
            onAdd: function(/*event*/) {
                var feesGrid    = w2ui.RARentableFeesGrid,
                    feeForm     = w2ui.RARentableFeeForm;

                var sliderID = 2;
                appendNewSlider(sliderID);
                $("#raflow-container")
                    .find(".slider[data-slider-id="+sliderID+"]")
                    .find(".slider-content")
                    .width(400)
                    .w2render(feeForm);

                // new record so select none
                feesGrid.selectNone();

                // get RID from last of raflow
                var RID = app.raflow.last.RID;

                // get all account rules in fit those in form "ARID" field
                var BID = getCurrentBID();
                GetAllARForFeeForm(BID)
                .done(function(data) {
                    // get filtered account rules items
                    feeForm.get("ARID").options.items = GetFeeAccountRulesW2UIListItems(BID, "rentables");

                    // set form record
                    SetFeeFormRecordFromFeeData(RID, 0, "rentables");
                    feeForm.record.recid = feesGrid.records.length + 1;

                    // show form in the DOM
                    ShowSliderContentW2UIComp(feeForm, SliderContentDivLength, sliderID);
                    feeForm.refresh();
                })
                .fail(function(data) {
                    console.log("failure" + data);
                });
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
                            var feeForm = w2ui.RARentableFeeForm;

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

                            // get RID from last of raflow
                            var RID = app.raflow.last.RID;

                            // get TMPASMID from grid record
                            var TMPASMID = grid.get(recid).TMPASMID;

                            // get all account rules then
                            var BID = getCurrentBID();
                            GetAllARForFeeForm(BID)
                            .done(function(data) {
                                // get filtered account rules items
                                feeForm.get("ARID").options.items = GetFeeAccountRulesW2UIListItems(BID, "rentables");

                                // set record in form
                                SetFeeFormRecordFromFeeData(RID, TMPASMID, "rentables");
                                feeForm.record.RentCycleText = app.cycleFreq[feeForm.record.RentCycle];

                                ShowSliderContentW2UIComp(feeForm, SliderContentDivLength, sliderID);
                                feeForm.refresh(); // need to refresh for header changes

                                // When RentCycle is Norecur then disable the RentCycle list field.
                                var isDisabled = feeForm.record.RentCycleText.text === app.cycleFreq[0];
                                $("#RentCycleText").prop("disabled", isDisabled);

                                setTimeout(function () {
                                    displayRARentableFeeFormError(app.raflow.last.RID);
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
            onRefresh: function(event) {
                var grid = this;
                event.onComplete = function() {
                    ShowHideGridToolbarAddButton(grid.name);
                };
            }
        });

        // -----------------------------------------------------------
        //  ***** RENTABLE ***** FEE ***** FORM *****
        // -----------------------------------------------------------
        $().w2form({
            name: 'RARentableFeeForm',
            header: 'Add New Rentable Fee',
            style: 'border: none; display: block;',
            formURL: '/webclient/html/raflow/formra-fee.html',
            focus: -1,
            fields: GetFeeFormFields(),
            toolbar : GetFeeFormToolbar(),
            actions: {
                reset: function () {
                    w2ui.RARentableFeeForm.clear();
                },
                save: function() {
                    var feeForm     = w2ui.RARentableFeeForm,
                        TMPASMID    = feeForm.record.TMPASMID;

                    // get RID from last of raflow
                    var RID = app.raflow.last.RID;

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // set local fee data from fee form
                    SetFeeDataFromFeeFormRecord(RID, TMPASMID, "rentables");

                    // re-calculate amounts for rentable
                    SetRentableAmountsFromFees(RID);

                    SaveRentableCompData()
                    .done(function (data) {
                        if (data.status === 'success') {
                            // Re render the fees grid records
                            AssignRentableFeesGridRecords(RID);

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
                    var feeForm     = w2ui.RARentableFeeForm,
                        feesGrid    = w2ui.RARentableFeesGrid,
                        TMPASMID    = feeForm.record.TMPASMID;

                    // get RID from last of raflow
                    var RID = app.raflow.last.RID;

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // set local fee data from fee form
                    SetFeeDataFromFeeFormRecord(RID, TMPASMID, "rentables");

                    // re-calculate amounts for rentable
                    SetRentableAmountsFromFees(RID);

                    SaveRentableCompData()
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
                            AssignRentableFeesGridRecords(RID);

                        } else {
                            feeForm.message(data.message);
                        }
                    })
                    .fail(function (data) {
                        console.log("failure " + data);
                    });
                },
                delete: function() {
                    var feeForm     = w2ui.RARentableFeeForm,
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

                    // get RID from last of raflow
                    var RID = app.raflow.last.RID;

                    var localRData = GetRentableLocalData(RID);
                    if (localRData.Fees.length > 0) {
                        var itemIndex = GetRentableFeeLocalData(RID, TMPASMID, true);

                        // remove fee item
                        localRData.Fees.splice(itemIndex, 1);

                        // set this modified local rentable data to back
                        SetRentableLocalData(RID, localRData);

                        // re-calculate amounts for rentable
                        SetRentableAmountsFromFees(RID);

                        // sync data on backend side
                        SaveRentableCompData()
                        .done(function (data) {
                            if (data.status === 'success') {
                                // reset form
                                feeForm.actions.reset();

                                // // Re render the fees grid records
                                AssignRentableFeesGridRecords(RID);

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
                    var feeForm = w2ui.RARentableFeeForm;

                    // take action on change event for this form
                    FeeFormOnChangeHandler(feeForm, event.target, event.value_new);

                       // formRecDiffer: 1=current record, 2=original record, 3=diff object
                    var diff = formRecDiffer(this.record, app.active_form_original, {});
                    // if diff == {} then make dirty flag as false, else true
                    app.form_is_dirty = !($.isPlainObject(diff) && $.isEmptyObject(diff));
                };
            },
            onRefresh: function(event) {
                var feeForm = this;
                event.onComplete = function() {

                    // there is NO PETID actually, so have to work around with recid key
                    formRefreshCallBack(feeForm);

                    // set header
                    var RID             = app.raflow.last.RID,
                        localRData      = GetRentableLocalData(RID),
                        rentableName    = localRData.RentableName;

                    var header = "Fee (<strong>{0}</strong>) for <strong>{1}</strong>";
                    if (feeForm.record.ARName && feeForm.record.ARName.length > 0) {
                        feeForm.header = header.format(feeForm.record.ARName, rentableName);
                    } else {
                        feeForm.header = header.format("new", rentableName);
                    }

                    // minimum actions need to be taken care in refres event for fee form
                    FeeFormOnRefreshHandler(feeForm);
                };
            }
        });
    }

    // now load grid in division
    $('#ra-form #rentables .grid-container').w2render(w2ui.RARentablesGrid);
    $('#ra-form #rentables .form-container').w2render(w2ui.RARentableSearchForm);
    HideAllSliderContent();

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
    var compData = GetRAFlowCompLocalData("rentables");
    if (compData) {
        w2ui.RARentablesGrid.records = compData;
        reassignGridRecids(w2ui.RARentablesGrid.name);

        // Operation on RARentableSearchForm
        w2ui.RARentableSearchForm.refresh();

        // manage parent rentables list
        manageParentRentableW2UIItems();

        // Render RentableGrid Summary
        RenderRentablesGridSummary();

        // display row with light red background if it have error
        dispalyRARentablesGridError();

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
    var localRData = GetRentableLocalData(RID);

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
    SetRentableLocalData(RID, localRData);
};

//-----------------------------------------------------------------------------
// RenderRentablesGridSummary - will render grid summary row from rentable
//                             comp data
//-----------------------------------------------------------------------------
window.RenderRentablesGridSummary = function() {
    var compData = GetRAFlowCompLocalData("rentables") || [];
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
        RowTotal:           0.0,
        RemoveRec:          null,
    };

    compData.forEach(function(rentableItem) {
        summaryRec.AtSigningPreTax += rentableItem.AtSigningPreTax;
        summaryRec.SalesTax += rentableItem.SalesTax;
        // summaryRec.SalesTaxAmt += rentableItem.SalesTaxAmt;
        summaryRec.TransOccTax += rentableItem.TransOccTax;
        // summaryRec.TransOccAmt += rentableItem.TransOccAmt;
        summaryRec.RowTotal += rentableItem.RowTotal;
    });

    // set style of entire summary row
    summaryRec.w2ui = {style: "font-weight: bold"};

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
    var localRData = GetRentableLocalData(RID),
        Fees = localRData.Fees || [];

    grid.records = Fees;

    // set the header as well
    grid.header = "Fees for (<strong>{0}</strong>)".format(localRData.RentableName);

    // render fees amount summary
    RenderFeesGridSummary(grid, Fees);

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

    // It highlight row with light red color if it have error
    displayRARentableFeesGridError();

    if (foundRIDIndex > -1) {
        var rentableGridRec = w2ui.RARentablesGrid.get(foundRIDIndex);
        var summaryRec = grid.summary[0]; //only one summary we have

        // summing up total
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
// AcceptRentable - add Rentable to the list rentables grid records
//-----------------------------------------------------------------------------
window.AcceptRentable = function () {
    var RID = w2ui.RARentableSearchForm.record.RID;

    // find index of this RID in grid if it exists
    var gridRecIndex = -1;
    w2ui.RARentablesGrid.records.forEach(function(rec) {
        if (RID == rec.RID) {
            gridRecIndex = rec.recid;
            return false;
        }
    });

    if(gridRecIndex > -1 ) {
        w2ui.RARentablesGrid.select(gridRecIndex); // highlight the existing record
        w2ui.RARentableSearchForm.clear(); // clear the search rentable form
    } else {
        var fRec    = w2ui.RARentableSearchForm.record;

        SaveRAFlowRentableAJAX(fRec.RID)
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
    var rentableCompData = GetRAFlowCompLocalData("rentables") || [];

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
// SaveRentableCompData - saves the data on server side
//------------------------------------------------------------------------------
window.SaveRentableCompData = function() {
    var compData = GetRAFlowCompLocalData("rentables");
    return SaveCompDataAJAX(compData, "rentables");
};

//-----------------------------------------------------------------------------
// GetRentableLocalData - returns the clone of rentable data for requested
//                        RID
//-----------------------------------------------------------------------------
window.GetRentableLocalData = function(RID, returnIndex) {
    var cloneData = {};
    var foundIndex = -1;
    var compData = GetRAFlowCompLocalData("rentables");
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
// SetRentableLocalData - save the data for requested RID in local data
//-----------------------------------------------------------------------------
window.SetRentableLocalData = function(RID, rentableData) {
    var compData = GetRAFlowCompLocalData("rentables");
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
// GetRentableFeeLocalData - returns the clone of rentable fee data for requested
//                           RID, TMPASMID from "Fees" list of a rentable
//-----------------------------------------------------------------------------
window.GetRentableFeeLocalData = function(RID, TMPASMID, returnIndex) {
    var cloneData = {};
    var foundIndex = -1;
    var compData = GetRAFlowCompLocalData("rentables");
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
// SetRentableFeeLocalData - save the data for Fee with
//                           requested RID in local data
//-----------------------------------------------------------------------------
window.SetRentableFeeLocalData = function(RID, TMPASMID, rentableFeeData) {
    var compData    = GetRAFlowCompLocalData("rentables"),
        rIndex      = -1,
        fIndex      = -1;

    // find rentable and fee in it
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

// dispalyRARentablesGridError
// It highlights grid's row if it have error
window.dispalyRARentablesGridError = function (){
    // load grid errors if any
    var g = w2ui.RARentablesGrid;
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

    if (app.raflow.validationErrors.rentables) {
        var rentables = app.raflow.validationCheck.errors.rentables.errors;
        for (i = 0; i < rentables.length; i++) {
            if (rentables[i].total > 0) {
                var recid = getRecIDFromRID(g, rentables[i].RID);
                g.get(recid).w2ui.style = "background-color: #EEB4B4";
                g.refreshRow(recid);
            }
        }
    }
};

// displayRARentableFeesGridError It highlight row with light red color if it have error
window.displayRARentableFeesGridError = function () {
    // load grid errors if any
    var g = w2ui.RARentableFeesGrid;
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

    if (app.raflow.validationErrors.rentables) {
        var rentables = app.raflow.validationCheck.errors.rentables.errors;
        for (i = 0; i < rentables.length; i++) {
            for (var j = 0; j < rentables[i].fees.errors.length; j++) {
                if (rentables[i].fees.errors[j].total > 0) {
                    var recid = getRecIDFromTMPASMID(g, rentables[i].fees.errors[j].TMPASMID);
                    g.get(recid).w2ui.style = "background-color: #EEB4B4";
                    g.refreshRow(recid);
                }
            }
        }
    }
};

// getRecIDFromRID It returns recid of grid record which matches TMPTCID
window.getRecIDFromRID = function(grid, RID){
    // var g = w2ui.RAPeopleGrid;
    var recid;
    for (var i = 0; i < grid.records.length; i++) {
        if (grid.records[i].RID === RID) {
            recid = grid.records[i].recid;
        }
    }
    return recid;
};

// displayRARentableFeeFormError If form field have error than it highlight with red border and
window.displayRARentableFeeFormError = function(RID){

    // if pet section doesn't have error than return
    if(!app.raflow.validationErrors.rentables){
        return;
    }

    var form = w2ui.RARentableFeeForm;
    var record = form.record;

    // get list of pets
    var rentables = app.raflow.validationCheck.errors.rentables.errors;

    // get index of vehicle for whom form is opened
    var rentableIndex = getRentableIndex(RID, rentables);

    var index = getFeeIndex(record.TMPASMID, rentables[rentableIndex].fees.errors);

    if(index > -1){
        displayFormFieldsError(index, rentables[rentableIndex].fees.errors, "RARentableFeeForm");
    }
};

// getRentableIndex it return an index of rentable who have RID
window.getRentableIndex = function (RID, rentables) {

    var index = -1;

    for(var i = 0; i < rentables.length; i++){
        // If RID doesn't match iterate for next element
        if(rentables[i].RID === RID){
            index = i;
            break;
        }
    }

    return index;
};
