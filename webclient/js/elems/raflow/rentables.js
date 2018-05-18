/* global
    RACompConfig, sliderContentDivLength, reassignGridRecids,
    hideSliderContent, appendNewSlider, showSliderContentW2UIComp,
    loadTargetSection, requiredFieldsFulFilled, getRAFlowPartTypeIndex, initRAFlowAJAX,
    getRAFlowAllParts, saveActiveCompData, toggleHaveCheckBoxDisablity, getRAFlowPartData,
    lockOnGrid,
    getRentableFeeFormInitalRecord, getRentablesGridInitalRecord, getInitialRentableFeesData,
    getRentableLocalData, setRentableLocalData, getAllARsWithAmount,
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


window.getRentableFeeFormInitalRecord = function () {
    var BID = getCurrentBID(),
        BUD = getBUDfromBID(BID);

    var t = new Date(),
        nyd = new Date(new Date().setFullYear(new Date().getFullYear() + 1));

    return {
        recid: 0,
        RID: 0,
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
                toolbar: false,
                footer: true,
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
                }
            ],
            onClick: function (event) {
                event.onComplete = function () {
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
                            if (localRData.Fees.length == 0) {
                                var BID = getCurrentBID();
                                getInitialRentableFeesData(BID, rec.RID)
                                .done(function(data) {
                                    // assign records in the grid and then render it
                                    w2ui.RARentableFeesGrid.records = data.records || [];
                                    reassignGridRecids(w2ui.RARentableFeesGrid.name);

                                    // save fees record
                                    localRData.Fees = w2ui.RARentableFeesGrid.records;
                                    setRentableLocalData(rec.RID, localRData);
                                });
                            } else {
                                w2ui.RARentableFeesGrid.records = localRData.Fees;
                                reassignGridRecids(w2ui.RARentableFeesGrid.name);
                            }
                            showSliderContentW2UIComp(w2ui.RARentableFeesGrid, RACompConfig.rentables.sliderWidth);
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

                            var BID = getCurrentBID();
                            getAllARsWithAmount(BID)
                            .done(function(data) {
                                var arid_items = [];
                                app.raflow.arList[BID].forEach(function(item) {
                                    arid_items.push({id: item.ARID, text: item.Name});
                                });
                                w2ui.RARentableFeesForm.get("ARID").options.items = arid_items;
                                w2ui.RARentableFeesForm.record = getRentableFeeFormInitalRecord();
                                w2ui.RARentableFeesForm.record.recid = w2ui.RARentableFeesGrid.records.length;
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
                    caption: 'Account Rule',
                    size: '150px',
                },
                {
                    field: 'Amount',
                    caption: 'Amount',
                    size: '100px',
                    render: 'money',
                },
                {
                    field: 'RentCycle',
                    caption: 'Cycle',
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
                                .w2render(w2ui.RARentableFeesForm);

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
                                w2ui.RARentableFeesForm.get("ARID").options.items = arid_items;
                                w2ui.RARentableFeesForm.record = $.extend(true, {}, grid.get(app.last.grid_sel_recid));
                                showSliderContentW2UIComp(w2ui.RARentableFeesForm, sliderContentDivLength, sliderID);
                                w2ui.RARentableFeesForm.refresh(); // need to refresh for header changes
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
                {name: 'ARID',              type: 'list',   required: true, html: {page: 0, column: 0}, options: {items: [], selected: {}}},
                {name: 'Amount',            type: 'money',  required: true, html: {page: 0, column: 0}},
                {name: 'RentCycle',         type: 'list',   required: true, html: {page: 0, column: 0}, options: {items: app.cycleFreq}},
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
                    this.clear();
                },
                save: function() {
                    var f = this;
                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // sync this info in local data
                    w2ui.RARentableFeesGrid.set(f.record.recid, f.record);
                    w2ui.RARentableFeesGrid.refresh();

                    f.save({}, function(data) {
                        if (data.status === 'error') {
                            f.message(data.message);
                            return;
                        }
                        hideSliderContent(2);
                    });
                },
                saveadd: function() {
                    var BID = getCurrentBID(),
                        BUD = getBUDfromBID(BID),
                        f = this;

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    f.save({}, function(data) {
                        if (data.status === 'error') {
                            f.message(data.message);
                            return;
                        }

                        // f.record = getRAAddRentableFormInitRec(BID, BUD, f.record);
                        f.refresh();
                    });
                },
            },
            onChange: function(event) {
                var f = this;
                event.onComplete = function() {
                    switch(event.target) {
                        case "RentCycle":
                            if (event.value_new) {
                                app.cycleFreq.forEach(function(itemText, itemIndex) {
                                    if (event.value_new.text == itemText) {
                                        f.record.RentCycle = itemIndex;
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
        var i = getRAFlowPartTypeIndex(app.raFlowPartTypes.rentables);
        if (i >= 0 && app.raflow.data[app.raflow.activeFlowID][i].Data) {
            w2ui.RARentablesGrid.records = app.raflow.data[app.raflow.activeFlowID][i].Data;
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
    var rec = getRentablesGridInitalRecord();
    rec.RID = w2ui.RARentableForm.record.RID;
    rec.RentableName = w2ui.RARentableForm.record.RentableName;
    rec.recid = w2ui.RARentablesGrid.records.length;
    w2ui.RARentablesGrid.add(rec);
    w2ui.RARentableForm.actions.reset(); // clear the form

    // also manage local data
    var partTypeIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.rentables);
    var rentableData = $.extend(true, {"Fees": []}, rec);
    app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data.push(rentableData);
};

//-----------------------------------------------------------------------------
// getRentableLocalData - returns the clone of rentable data for requested
//                        RID
//-----------------------------------------------------------------------------
window.getRentableLocalData = function(RID) {
    var cloneData = {};
    var partTypeIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.rentables);
    var data = app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data;
    data.forEach(function(item) {
        if (item.RID == RID) {
            cloneData = $.extend(true, {}, item);
            return false;
        }
    });
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
        return true;
    }
    return false;
};

