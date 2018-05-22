/* global
    RACompConfig, sliderContentDivLength, reassignGridRecids,
    hideSliderContent, appendNewSlider, showSliderContentW2UIComp,
    loadTargetSection, requiredFieldsFulFilled, getRAFlowPartTypeIndex, initRAFlowAJAX,
    getRAFlowAllParts, saveActiveCompData, toggleHaveCheckBoxDisablity, getRAFlowPartData,
    lockOnGrid,
    getRentableFeeFormInitialRecord, getRentablesGridInitalRecord, getInitialRentableFeesData,
    getRentableLocalData, setRentableLocalData, getAllARsWithAmount, GetRentableIndexInGridRecords,
    saveRentableCompData, setRentableFeeLocalData, getRentableFeeLocalData,
    ridRentablePickerRender, ridRentableDropRender, ridRentableCompare
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
window.getInitialRentableFeesData = function(BID, RID) {
    var data = {"RID": RID};
    return $.ajax({
        url: "/v1/raflow-rentable-fees/" + BID.toString() + "/",
        method: "POST",
        contentType: "application/json",
        data: JSON.stringify(data),
    });
};

// getRentablesGridInitalRecord returns grid initial record
// for the grid
window.getRentablesGridInitalRecord = function () {
    var BID = getCurrentBID(),
        BUD = getBUDfromBID(BID);

    return {
        recid: 0,
        RID: 0,
        BID: BID,
        BUD: BUD,
        RentableName: "",
        ContractRent: 0.0,
        ProrateAmt: 0.0,
        SalesTax: 0.0,
        TaxableAmt: 0.0,
        TransOcc: 0.0,
    };
};


window.getRentableFeeFormInitialRecord = function (RID) {
    var BID = getCurrentBID(),
        BUD = getBUDfromBID(BID);

    var t = new Date(),
        nyd = new Date(new Date().setFullYear(new Date().getFullYear() + 1));

    return {
        recid: 0,
        RID: RID,
        ARID: 0,
        BID: BID,
        BUD: BUD,
        Amount: 0.0,
        RentCycle: "Daily",
        Epoch: 0,
        RentPeriodStart: w2uiDateControlString(t),
        RentPeriodStop: w2uiDateControlString(nyd),
        UsePeriodStart: w2uiDateControlString(t),
        UsePeriodStop: w2uiDateControlString(nyd),
        ContractRent: 0.0,
        ProrateAmt: 0.0,
        SalesTaxAmt: 0.0,
        SalesTax: 0.0,
        TransOccAmt: 0.0,
        TransOcc: 0.0,
    };
};

window.loadRARentablesGrid = function () {

    var partType = app.raFlowPartTypes.rentables;
    var partTypeIndex = getRAFlowPartTypeIndex(partType);

    if (partTypeIndex < 0){
        return;
    }

    // Fetch data from the server if there is any record available.
    getRAFlowPartData(partType)
        .done(function(data){
            if(data.status === 'success'){
                app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data = data.record.Data;
            }else {
                console.log(data.message);
            }
        })
        .fail(function(data){
            console.log("failure" + data);
        });

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
                    field: 'RentableName',
                    caption: 'Rentable',
                    size: '200px',
                },
                {
                    field: 'ContractRent',
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
                    field: 'TaxableAmt',
                    caption: 'Taxable Amt',
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
                    field: 'TransOcc',
                    caption: 'Trans OCC',
                    size: '100px',
                    render: 'money',
                },
                {
                    field: 'RemoveRec',
                    caption: "Remove Rentable",
                    size: '100%',
                    render: function (record/*, index, col_index*/) {
                        var html = "";
                        if (record) {
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
                        var partTypeIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.rentables);
                        app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data.splice(index, 1);

                        // save data on server
                        saveRentableCompData();

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

                            // pull fees in case it's empty
                            if(localRData.Fees.length > 0) {
                                // show slider content
                                w2ui.RARentableFeesGrid.records = localRData.Fees;
                                reassignGridRecids(w2ui.RARentableFeesGrid.name);
                                showSliderContentW2UIComp(w2ui.RARentableFeesGrid, RACompConfig.rentables.sliderWidth);
                            } else {
                                var BID = getCurrentBID();
                                getInitialRentableFeesData(BID, rec.RID)
                                .done(function(data) {
                                    if (data.status === "success") {
                                        // save fees record
                                        localRData.Fees = data.records || [];
                                        // also manage local data
                                        setRentableLocalData(rec.RID, localRData);

                                        // save data on the server
                                        saveRentableCompData();

                                        // show slider content
                                        w2ui.RARentableFeesGrid.records = localRData.Fees;
                                        reassignGridRecids(w2ui.RARentableFeesGrid.name);
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
                    caption: 'Account Rule ID',
                    size: '100px',
                    hidden: true
                },
                {
                    field: 'ARName',
                    caption: 'Account Rule',
                    size: '150px'
                },
                {
                    field: 'Amount',
                    caption: 'Amount',
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
                    field: 'ContractRent',
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
                                form.get("ARName").options.items = arid_items;
                                form.record = $.extend(true, {}, grid.get(app.last.grid_sel_recid));

                                form.record.RentCycleList = app.cycleFreq[form.record.RentCycle];

                                // mark current ARID in app last rentableFeeARID
                                app.raflow.last.rentableFeeARID = form.record.ARID;

                                showSliderContentW2UIComp(form, sliderContentDivLength, sliderID);
                                form.refresh(); // need to refresh for header changes
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
                {name: 'ARID',              type: 'int',   required: true, html: {page: 0, column: 0}},
                {name: 'ARName',            type: 'list',   required: true, html: {page: 0, column: 0}, options: {items: [], selected: {}}},
                {name: 'Amount',            type: 'money',  required: true, html: {page: 0, column: 0}},
                {name: 'RentCycle',         type: 'int',   required: true, html: {page: 0, column: 0}},
                {name: 'RentCycleList',     type: 'list',   required: true, html: {page: 0, column: 0}, options: {items: app.cycleFreq}},
                {name: 'Epoch',             type: 'int',    required: true, html: {page: 0, column: 0}},
                {name: 'RentPeriodStart',   type: 'date',   required: true, html: {page: 0, column: 0}},
                {name: 'RentPeriodStop',    type: 'date',   required: true, html: {page: 0, column: 0}},
                {name: 'UsePeriodStart',    type: 'date',   required: true, html: {page: 0, column: 0}},
                {name: 'UsePeriodStop',     type: 'date',   required: true, html: {page: 0, column: 0}},
                {name: 'ContractRent',      type: 'money',  required: true, html: {page: 0, column: 0}},
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
                        grid = w2ui.RARentableFeesGrid;

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    var isNewRecord = (grid.get(f.record.recid, true) === null);

                    // sync this info in local data
                    var rentableFeeData = getFormSubmitData(f.record, true);
                    if (isNewRecord) {
                        grid.add(rentableFeeData);
                    } else {
                        grid.set(f.record.recid, rentableFeeData);
                    }
                    grid.refresh();

                    // set data locally
                    setRentableFeeLocalData(f.record.RID, app.raflow.last.rentableFeeARID, rentableFeeData);

                    saveRentableCompData()
                    .done(function (data) {
                        if (data.status === 'success') {
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
                        grid = w2ui.RARentableFeesGrid;

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    var isNewRecord = (grid.get(f.record.recid, true) === null);

                    // sync this info in local data
                    var rentableFeeData = getFormSubmitData(f.record, true);
                    if (isNewRecord) {
                        grid.add(rentableFeeData);
                    } else {
                        grid.set(f.record.recid, rentableFeeData);
                    }
                    grid.refresh();

                    // set data locally
                    setRentableFeeLocalData(f.record.RID, app.raflow.last.rentableFeeARID, rentableFeeData);

                    // get current rentable
                    var RID = w2ui.RARentablesGrid.get(w2ui.RARentablesGrid.getSelection()[0]).RID;

                    saveRentableCompData()
                    .done(function (data) {
                        if (data.status === 'success') {
                            f.actions.reset();
                            f.record = getRentableFeeFormInitialRecord(RID);
                            f.record.recid = grid.records.length + 1;
                            f.refresh();
                        } else {
                            f.message(data.message);
                        }
                    })
                    .fail(function (data) {
                        console.log("failure " + data);
                    });
                },
                delete: function() {
                    var f = w2ui.RARentableFeesForm;
                    var localRData = getRentableLocalData(f.record.RID);
                    if (localRData.Fees.length > 0) {
                        var itemIndex = getRentableFeeLocalData(f.record.RID, app.raflow.last.rentableFeeARID, true);
                        // remove fee item
                        localRData.Fees.splice(itemIndex, 1);

                        // set this modified local rentable data to back
                        setRentableLocalData(localRData.RID, localRData);

                        // sync data on backend side
                        saveRentableCompData()
                        .done(function (data) {
                            if (data.status === 'success') {
                                // reset form as well as remove record from the grid
                                w2ui.RARentableFeesGrid.remove(f.record.recid);
                                w2ui.RARentableFeesGrid.refresh();
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
                    }
                }
            },
            onChange: function(event) {
                var f = this;
                event.onComplete = function() {
                    switch(event.target) {
                        case "RentCycleList":
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
                        case "ARName":
                            if (event.value_new) {
                                // mark previous ARID in app last rentableFeeARID
                                app.raflow.last.rentableFeeARID = event.value_previous.id;

                                var BID = getCurrentBID();
                                app.raflow.arList[BID].forEach(function(item) {
                                    if (event.value_new.id == item.ARID) {
                                        console.log(item);
                                        f.record.Amount = item.DefaultAmount;
                                        f.record.ARID = item.ARID;

                                        if(item.FLAGS === 66){
                                            f.record.RentCycle = app.cycleFreq[0];
                                        }

                                        f.refresh();
                                        return false;
                                    }
                                });
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
                    f.get("ARName").options.items.forEach(function(item) {
                        if (item.id == f.record.ARID) {
                            ARIDSel = {id: item.id, text: item.text};
                        }
                    });

                    f.record.BID = BID;
                    f.record.BUD = BUD;
                    f.get("ARName").options.selected = ARIDSel;

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
        var i = getRAFlowPartTypeIndex(app.raFlowPartTypes.rentables);
        if (i >= 0 && app.raflow.data[app.raflow.activeFlowID][i].Data) {
            w2ui.RARentablesGrid.records = app.raflow.data[app.raflow.activeFlowID][i].Data || [];
            reassignGridRecids(w2ui.RARentablesGrid.name);

            // Operation on RARentableForm
            w2ui.RARentableForm.refresh();
        } else {
            w2ui.RARentablesGrid.clear();
            // Operation on RARentableForm
            w2ui.RARentableForm.actions.reset();
        }
    }, 500);
};

//-----------------------------------------------------------------------------
// acceptRentable - add Rentable to the list rentables grid records
//-----------------------------------------------------------------------------
window.acceptRentable = function () {
    var recIndex = GetRentableIndexInGridRecords(w2ui.RARentableForm.record.RID);
    var rec;
    var BID = getCurrentBID();

    if(recIndex > -1 ) {
        w2ui.RARentablesGrid.select(recIndex); // highlight the existing record
    } else {
        rec = getRentablesGridInitalRecord();
        rec.RID = w2ui.RARentableForm.record.RID;
        rec.RentableName = w2ui.RARentableForm.record.RentableName;
        rec.recid = w2ui.RARentablesGrid.records.length + 1;
        w2ui.RARentablesGrid.add(rec);

        // get latest fees record
        var rentableData = $.extend(true, {"Fees": []}, rec);
        getInitialRentableFeesData(BID, rec.RID)
        .done(function(data) {

            if (data.status === "success") {
                // save fees record
                rentableData.Fees = data.records || [];

                // also manage local data
                setRentableLocalData(rec.RID, rentableData);
                var partTypeIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.rentables);
                saveActiveCompData(app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data,
                    app.raFlowPartTypes.rentables);
            }
        })
        .fail(function(data) {
            console.log("ERROR from fees data: " + data);
        });

        // select none
        setTimeout(function() {
            w2ui.RARentablesGrid.selectNone();
        }, 0);
    }

     // clear the form
    w2ui.RARentableForm.actions.reset();
};

//------------------------------------------------------------------------------
// saveRentableCompData - saves the data on
//------------------------------------------------------------------------------
window.saveRentableCompData = function() {
    var partTypeIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.rentables);
    return saveActiveCompData(app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data,
        app.raFlowPartTypes.rentables);
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
    var partTypeIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.rentables);
    var data = app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data;
    data.forEach(function(item, index) {
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
    var partTypeIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.rentables);
    var dataIndex = -1;
    app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data.forEach(function(item, index) {
        if (item.RID == RID) {
            dataIndex = index;
            return false;
        }
    });
    if (dataIndex > -1) {
        app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data[dataIndex] = rentableData;
    } else {
        app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data.push(rentableData);
    }
};

//-----------------------------------------------------------------------------
// getRentableFeeLocalData - returns the clone of rentable fee data for requested
//                        RID, ARID from "Fees" list of a rentable
//-----------------------------------------------------------------------------
window.getRentableFeeLocalData = function(RID, ARID, returnIndex) {
    var cloneData = {};
    var foundIndex = -1;
    var partTypeIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.rentables);
    var data = app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data;
    data.forEach(function(item) {
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
    var partTypeIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.rentables);
    var rIndex = -1,
        fIndex = -1;

    app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data.forEach(function(item, itemIndex) {
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
            app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data[rIndex].Fees[fIndex] = rentableFeeData;
        } else {
            app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data[rIndex].Fees.push(rentableFeeData);
        }
    }
};

