/* global
    RACompConfig, sliderContentDivLength, reassignGridRecids,
    hideSliderContent, appendNewSlider, showSliderContentW2UIComp,
    loadTargetSection, requiredFieldsFulFilled, getRAFlowPartTypeIndex, initRAFlowAjax,
    getRAFlowAllParts, saveActiveCompData, toggleHaveCheckBoxDisablity, getRAFlowCompData,
    lockOnGrid,
    getRentableFeeFormInitialRecord, getRentablesGridInitalRecord, getInitialRentableFeesData,
    getRentableLocalData, setRentableLocalData, getAllARsWithAmount, GetRentableIndexInGridRecords,
    saveRentableCompData, setRentableFeeLocalData, getRentableFeeLocalData,
    ridRentablePickerRender, ridRentableDropRender, ridRentableCompare,
    AssignRentableGridRecords, AssignRentableFeesGridRecords,
    SetRentableAmountsFromFees, manageParentRentableW2UIItems,
    RenderRentablesGridSummary
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

/*// getRentablesGridInitalRecord returns grid initial record
// for the grid
window.getRentablesGridInitalRecord = function () {
    var BID = getCurrentBID(),
        BUD = getBUDfromBID(BID);

    return {
        recid: 0,
        RID: 0,
        BID: BID,
        RTID: 0,
        RentableName: "",
        RentCycle: 0,
        RentCycleText: "",
        AtSigningAmt: 0.0,
        ProrateAmt: 0.0,
        SalesTax: 0.0,
        TaxableAmt: 0.0,
        TransOcc: 0.0,
        Fees: []
    };
};*/


window.getRentableFeeFormInitialRecord = function (RID) {
    var BID = getCurrentBID(),
        BUD = getBUDfromBID(BID);

    var t = new Date(),
        nyd = new Date(new Date().setFullYear(new Date().getFullYear() + 1));

    return {
        recid: 0,
        RID: RID,
        ARID: 0,
        ARName: "",
        BID: BID,
        BUD: BUD,
        ContractAmount: 0.0,
        RentCycle: 0,
        RentCycleText: "",
        Epoch: 0,
        RentPeriodStart: w2uiDateControlString(t),
        RentPeriodStop: w2uiDateControlString(nyd),
        UsePeriodStart: w2uiDateControlString(t),
        UsePeriodStop: w2uiDateControlString(nyd),
        AtSigningAmt: 0.0,
        ProrateAmt: 0.0,
        SalesTaxAmt: 0.0,
        SalesTax: 0.0,
        TransOccAmt: 0.0,
        TransOcc: 0.0,
    };
};

window.loadRARentablesGrid = function () {

    // if form is loaded then return
    if (!("RARentablesGrid" in w2ui)) {

        // people form
        $().w2form({
            name: 'RARentableForm',
            header: 'Rentable',
            style: 'display: block; border: none;',
            formURL: '/webclient/html/formrar.html',
            focus: -1,
            fields: [
                { field: 'Rentable', required: true,
                    type: 'enum',
                    options: {
                        url:           '/v1/rentablestd/' + app.raflow.BID,
                        max:           1,
                        cacheMax:      50,
                        maxDropHeight: 350,
                        renderItem:    function(item) {
                            // Enable Accept button
                            $(w2ui.RARentableForm.box).find("button[name=accept]").prop("disabled", false);
                            w2ui.RARentableForm.record.RID = item.RID;
                            w2ui.RARentableForm.record.RentableName = item.RentableName;
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
                                w2ui.RARentableForm.actions.reset();
                            };
                        }
                    },
                },
                {name: 'BID',           type: 'int', required: true, html: {caption: "BID"}},
                {name: 'RID',           type: 'int', required: true, html: {caption: "RID"}},
                {name: 'RentableName',  type: 'int', required: true, html: {caption: "RentableName"}},
            ],
            actions: {
                reset: function () {
                    w2ui.RARentableForm.clear();
                    $(w2ui.RARentableForm.box).find("button[name=accept]").prop("disabled", true);
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

        // rentables grid
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
                    hidden: true,
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
                    size: '160px',
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
                    field: 'AtSigningAmt',
                    caption: 'At Signing',
                    size: '90px',
                    render: 'money',
                },
                {
                    field: 'ProrateAmt',
                    caption: 'Prorate',
                    size: '90px',
                    render: 'money',
                },
                {
                    field: 'TaxableAmt',
                    caption: 'Taxable Amt',
                    size: '90px',
                    render: 'money',
                },
                {
                    field: 'SalesTax',
                    caption: 'Sales Tax',
                    size: '90px',
                    render: 'money',
                },
                {
                    field: 'TransOcc',
                    caption: 'Trans OCC',
                    size: '90px',
                    render: 'money',
                },
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
                    if(this.getColumn("RemoveRec", true) == event.column) {
                        // remove entry from local data
                        var rec = this.get(event.recid);
                        var index = GetRentableIndexInGridRecords(rec.RID);

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
                            }
                        });

                        // remove from grid
                        this.remove(event.recid);
                        return;
                    }

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

                            // get auto populated to new RA account rules
                            var rec = grid.get(recid);
                            var localRData = getRentableLocalData(rec.RID);

                            // just render the record from local data if fees are available
                            if(localRData.Fees.length > 0) {
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

        // rentables grid
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
                                .w2render(w2ui.RARentableFeesForm);

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
                                w2ui.RARentableFeesForm.get("ARID").options.items = arid_items;
                                w2ui.RARentableFeesForm.record = getRentableFeeFormInitialRecord(RID);
                                w2ui.RARentableFeesForm.record.recid = w2ui.RARentableFeesGrid.records.length + 1;

                                // mark current ARID in app last rentableFeeARID
                                app.raflow.last.rentableFeeARID = w2ui.RARentableFeesForm.record.ARID;

                                showSliderContentW2UIComp(w2ui.RARentableFeesForm, sliderContentDivLength, sliderID);
                                w2ui.RARentableFeesForm.refresh();
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
            columns: [
                {
                    field: 'recid',
                    hidden: true,
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
                    field: 'ARID',
                    hidden: true
                },
                {
                    field: 'ARName',
                    caption: 'Account Rule',
                    size: '150px'
                },
                {
                    field: 'ContractAmount',
                    caption: 'Contract<br>Amount',
                    size: '100px',
                    render: 'money',
                },
                {
                    field: 'RentCycleText',
                    caption: 'Rent Cycle',
                    size: '100px',
                    render: function (record/*, index, col_index*/) {
                        var text = '';
                        if (record) {
                            app.cycleFreq.forEach(function(itemText, itemIndex) {
                                if (record.RentCycle == itemIndex) {
                                    text = itemText;
                                    return false;
                                }
                            });
                        }
                        return text;
                    },
                },
                {
                    field: 'RentCycle',
                    caption: 'Rent Cycle Index',
                    hidden: true
                },
                {
                    field: 'Epoch',
                    caption: 'Epoch',
                    size: '100px',
                },
                {
                    field: 'RentPeriod',
                    caption: 'Rent Period',
                    size: '100px',
                    render: function(record) {
                        var html = "";
                        if (record) {
                            if (record.RentPeriodStart && record.RentPeriodStop) {
                                html = record.RentPeriodStart + " - <br>" + record.RentPeriodStop;
                            }
                        }
                        return html;
                    }
                },
                {
                    field: 'RentPeriodStart',
                    hidden: true,
                },
                {
                    field: 'RentPeriodStop',
                    hidden: true,
                },
                {
                    field: 'UsePeriod',
                    caption: 'Use Period',
                    size: '100px',
                    render: function(record) {
                        var html = "";
                        if (record) {
                            if (record.UsePeriodStart && record.UsePeriodStop) {
                                html = record.UsePeriodStart + " - <br>" + record.UsePeriodStop;
                            }
                        }
                        return html;
                    }
                },
                {
                    field: 'UsePeriodStart',
                    hidden: true,
                },
                {
                    field: 'UsePeriodStop',
                    hidden: true,
                },
                {
                    field: 'AtSigningAmt',
                    caption: 'At Signing',
                    size: '100px',
                    render: 'money',
                },
                {
                    field: 'ProrateAmt',
                    caption: 'Prorate',
                    size: '100px',
                    render: 'money',
                },
                {
                    field: 'SalesTaxAmt',
                    caption: 'Sales Tax Amt',
                    size: '100px',
                    render: 'money',
                },
                {
                    field: 'SalesTax',
                    caption: 'Sales Tax',
                    size: '100px',
                    render: 'money',
                },
                {
                    field: 'TransOccAmt',
                    caption: 'Trans Occ Amt',
                    size: '100px',
                    render: 'money',
                },
                {
                    field: 'TransOcc',
                    caption: 'Trans Occ',
                    size: '100px',
                    render: 'money',
                }
            ],
            onChange: function (event) {
                event.onComplete = function () {
                    this.save();
                };
            },
            onClick: function(event) {
                event.onComplete = function() {
                    var form = w2ui.RARentableFeesForm;
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

                                // mark current ARID in app last rentableFeeARID
                                app.raflow.last.rentableFeeARID = form.record.ARID;

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

        // new Rentable form especially for this RA flow
        $().w2form({
            name: 'RARentableFeesForm',
            header: 'Add New Rentable Fee',
            style: 'display: block;',
            formURL: '/webclient/html/formra-rentablefee.html',
            focus: -1,
            fields: [
                {name: 'recid',             type: 'int',    required: true, html: {page: 0, column: 0}},
                {name: 'BID',               type: 'int',    required: true, html: {page: 0, column: 0}},
                {name: 'BUD',               type: 'list',   required: true, html: {page: 0, column: 0}, options: {items: app.businesses}},
                {name: 'RID',               type: 'int',    required: true, html: {page: 0, column: 0}},
                {name: 'ARName',            type: 'text',   required: true, html: {page: 0, column: 0}},
                {name: 'ARID',              type: 'list',   required: true, html: {page: 0, column: 0}, options: {items: [], selected: {}}},
                {name: 'ContractAmount',    type: 'money',  required: true, html: {page: 0, column: 0}},
                {name: 'RentCycle',         type: 'int',    required: true, html: {page: 0, column: 0}},
                {name: 'RentCycleText',     type: 'list',   required: true, html: {page: 0, column: 0}, options: {items: app.cycleFreq}},
                {name: 'Epoch',             type: 'int',    required: true, html: {page: 0, column: 0}},
                {name: 'RentPeriodStart',   type: 'date',   required: true, html: {page: 0, column: 0}},
                {name: 'RentPeriodStop',    type: 'date',   required: true, html: {page: 0, column: 0}},
                {name: 'UsePeriodStart',    type: 'date',   required: true, html: {page: 0, column: 0}},
                {name: 'UsePeriodStop',     type: 'date',   required: true, html: {page: 0, column: 0}},
                {name: 'AtSigningAmt',      type: 'money',  required: true, html: {page: 0, column: 0}},
                {name: 'ProrateAmt',        type: 'money',  required: true, html: {page: 0, column: 0}},
                {name: 'SalesTaxAmt',       type: 'money',  required: true, html: {page: 0, column: 0}},
                {name: 'SalesTax',          type: 'money',  required: true, html: {page: 0, column: 0}},
                {name: 'TransOccAmt',       type: 'money',  required: true, html: {page: 0, column: 0}},
                {name: 'TransOcc',          type: 'money',  required: true, html: {page: 0, column: 0}},
            ],
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
                    w2ui.RARentableFeesForm.clear();
                },
                save: function() {
                    var f = this,
                        grid = w2ui.RARentableFeesGrid,
                        RID = f.record.RID;

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // sync this info in local data
                    var rentableFeeData = getFormSubmitData(f.record, true);

                    // set data locally
                    setRentableFeeLocalData(RID, app.raflow.last.rentableFeeARID, rentableFeeData);

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
                    var BID = getCurrentBID(),
                        BUD = getBUDfromBID(BID),
                        f = this,
                        grid = w2ui.RARentableFeesGrid,
                        RID = f.record.RID;

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // sync this info in local data
                    var rentableFeeData = getFormSubmitData(f.record, true);

                    // set data locally
                    setRentableFeeLocalData(RID, app.raflow.last.rentableFeeARID, rentableFeeData);

                    // re-calculate amounts for rentable
                    SetRentableAmountsFromFees(RID);

                    saveRentableCompData()
                    .done(function (data) {
                        if (data.status === 'success') {
                            f.actions.reset();
                            f.record = getRentableFeeFormInitialRecord(RID);
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
                    var f = w2ui.RARentableFeesForm,
                        RID = f.record.RID;

                    var localRData = getRentableLocalData(RID);
                    if (localRData.Fees.length > 0) {
                        var itemIndex = getRentableFeeLocalData(RID, app.raflow.last.rentableFeeARID, true);

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
                                w2ui.RARentableFeesGrid.remove(RID);
                                w2ui.RARentableFeesGrid.refresh();
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
                                // mark previous ARID in app last rentableFeeARID
                                app.raflow.last.rentableFeeARID = event.value_previous.id;

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
                    var BID = getCurrentBID(),
                        BUD = getBUDfromBID(BID);


                    var ARIDSel = {};
                    // select value for rentable type FLAGS
                    f.get("ARID").options.items.forEach(function(item) {
                        if (item.id == f.record.ARID) {
                            ARIDSel = {id: item.id, text: item.text};
                        }
                    });

                    f.record.BID = BID;
                    f.record.BUD = BUD;
                    f.get("ARID").options.selected = ARIDSel;

                    // there is NO PETID actually, so have to work around with recid key
                    formRefreshCallBack(f, "recid");
                };
            }
        });

    }

    // now load grid in division
    $('#ra-form #rentables .grid-container').w2render(w2ui.RARentablesGrid);
    $('#ra-form #rentables .form-container').w2render(w2ui.RARentableForm);

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

        // Operation on RARentableForm
        w2ui.RARentableForm.refresh();

        // manage parent rentables list
        manageParentRentableW2UIItems();

        // Render RentableGrid Summary
        RenderRentablesGridSummary();

    } else {
        w2ui.RARentablesGrid.clear();
        // Operation on RARentableForm
        w2ui.RARentableForm.actions.reset();
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

    // temp variable to hold the summing up figure for all amounts
    var amountsSum = {
        AtSigningAmt:   0.0,
        ProrateAmt:     0.0,
        SalesTaxAmt:    0.0,
        SalesTax:       0.0,
        TransOccAmt:    0.0,
        TransOcc:       0.0,
    };

    // iterate over each Fees record
    localRData.Fees.forEach(function(feeData) {
        amountsSum.AtSigningAmt += feeData.AtSigningAmt;
        amountsSum.ProrateAmt += feeData.ProrateAmt;
        // amountsSum.SalesTaxAmt += feeData.SalesTaxAmt;
        amountsSum.SalesTax += feeData.SalesTax;
        // amountsSum.TransOccAmt += feeData.TransOccAmt;
        amountsSum.TransOcc += feeData.TransOcc;
    });

    // set the amount to rentable
    localRData.AtSigningAmt = amountsSum.AtSigningAmt;
    localRData.ProrateAmt = amountsSum.ProrateAmt;
    // localRData.SalesTaxAmt = amountsSum.SalesTaxAmt;
    localRData.SalesTax = amountsSum.SalesTax;
    // localRData.TransOccAmt = amountsSum.TransOccAmt;
    localRData.TransOcc = amountsSum.TransOcc;

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
        recid:          0,
        RentableName:   "Grand Total",
        AtSigningAmt:   0.0,
        ProrateAmt:     0.0,
        // SalesTaxAmt:    0.0,
        SalesTax:       0.0,
        // TransOccAmt:    0.0,
        TransOcc:       0.0,
        RemoveRec:      null,
    };

    compData.forEach(function(rentableItem) {
        summaryRec.AtSigningAmt += rentableItem.AtSigningAmt;
        summaryRec.ProrateAmt += rentableItem.ProrateAmt;
        // summaryRec.SalesTaxAmt += rentableItem.SalesTaxAmt;
        summaryRec.SalesTax += rentableItem.SalesTax;
        // summaryRec.TransOccAmt += rentableItem.TransOccAmt;
        summaryRec.TransOcc += rentableItem.TransOcc;
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

    // summary record in fees grid
    var summaryRec = {
        recid:          0,
        ARName:         "Grand Total",
        ContractAmount: 0.0,
        AtSigningAmt:   0.0,
        ProrateAmt:     0.0,
        SalesTaxAmt:    0.0,
        SalesTax:       0.0,
        TransOccAmt:    0.0,
        TransOcc:       0.0
    };

    // calculate amount for summary row
    grid.records.forEach(function(item) {
        summaryRec.ContractAmount += item.ContractAmount;
        summaryRec.AtSigningAmt += item.AtSigningAmt;
        summaryRec.ProrateAmt += item.ProrateAmt;
        summaryRec.SalesTaxAmt += item.SalesTaxAmt;
        summaryRec.SalesTax += item.SalesTax;
        summaryRec.TransOccAmt += item.TransOccAmt;
        summaryRec.TransOcc += item.TransOcc;
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
        rentableGridRec.AtSigningAmt = summaryRec.AtSigningAmt;
        rentableGridRec.ProrateAmt = summaryRec.ProrateAmt;
        // rentableGridRec.SalesTaxAmt = summaryRec.SalesTaxAmt;
        rentableGridRec.SalesTax = summaryRec.SalesTax;
        // rentableGridRec.TransOccAmt = summaryRec.TransOccAmt;
        rentableGridRec.TransOcc = summaryRec.TransOcc;

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
    var recIndex = GetRentableIndexInGridRecords(w2ui.RARentableForm.record.RID);
    var BID = getCurrentBID();

    if(recIndex > -1 ) {
        w2ui.RARentablesGrid.select(recIndex); // highlight the existing record
    } else {
        var fRec = w2ui.RARentableForm.record;
        getInitialRentableFeesData(BID, fRec.RID, app.raflow.activeFlowID)
        .done(function(data) {

            if (data.status === "success") {
                // reset the form
                w2ui.RARentableForm.actions.reset();
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
    if (dataIndex > -1) {
        compData[dataIndex] = rentableData;
    } else {
        compData.push(rentableData);
    }
};

//-----------------------------------------------------------------------------
// getRentableFeeLocalData - returns the clone of rentable fee data for requested
//                        RID, ARID from "Fees" list of a rentable
//-----------------------------------------------------------------------------
window.getRentableFeeLocalData = function(RID, ARID, returnIndex) {
    var cloneData = {};
    var foundIndex = -1;
    var compData = getRAFlowCompData("rentables", app.raflow.activeFlowID);
    compData.forEach(function(item) {
        if (item.RID == RID) {
            var feesData = item.Fees || [];
            feesData.forEach(function(feeItem, index) {
                if (feeItem.ARID == ARID) {
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
// setRentableFeeLocalData - save the data for Fee requested RID in local data
//-----------------------------------------------------------------------------
window.setRentableFeeLocalData = function(RID, ARID, rentableFeeData) {
    var compData = getRAFlowCompData("rentables", app.raflow.activeFlowID);
    var rIndex = -1,
        fIndex = -1;

    compData.forEach(function(item, itemIndex) {
        if (item.RID == RID) {
            var feesData = item.Fees || [];
            feesData.forEach(function(feeItem, feeItemIndex) {
                if (feeItem.ARID == ARID) {
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

