/* global
    RACompConfig, sliderContentDivLength, reassignGridRecids,
    hideSliderContent, appendNewSlider, showSliderContentW2UIComp,
    loadTargetSection, requiredFieldsFulFilled, getRAFlowPartTypeIndex, initRAFlowAJAX,
    getRAFlowAllParts, saveActiveCompData, toggleHaveCheckBoxDisablity, getRAFlowPartData,
    lockOnGrid,
    getRentableFeeFormInitalRecord,
*/

"use strict";

// -------------------------------------------------------------------------------
// getAutoPopulateARs - pull down all account rules which are set to auto populate
//                      to new rental agreement
// -------------------------------------------------------------------------------
window.getAutoPopulateARs = function() {
    return $.ajax({
        url: ''
    });
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
        ARName: "",
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

        // rentables grid
        $().w2grid({
            name: 'RARentablesGrid',
            header: 'Rentables',
            show: {
                toolbar: true,
                footer: true,
            },
            style: 'border: 1px solid black; display: block;',
            toolbar: {
                items: [
                    {id: 'add', type: 'button', caption: 'Add Record', icon: 'w2ui-icon-plus'},
                ],
                onClick: function (event) {
                    switch(event.target) {
                        case "add":
                            // get auto populated to new RA account rules
                            var data = {
                              "type":"FLAGS",
                              "FLAGS": 1<<app.arFLAGS.PopulateOnRA
                            };
                            var BID = getCurrentBID();
                            $.ajax({
                                url: "/v1/raflow-rentable-fees/" + BID.toString() + "/",
                                method: "POST",
                                contentType: "application/json",
                                data: JSON.stringify(data),
                            }).done(function(data) {
                                // assign records in the grid and then render it
                                w2ui.RARentableFeesGrid.records = data.records;
                                reassignGridRecids(w2ui.RARentableFeesGrid.name);
                                showSliderContentW2UIComp(w2ui.RARentableFeesGrid, RACompConfig.rentables.sliderWidth);
                            });
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
                    field: 'RTID',
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
                    field: 'TransOCC',
                    caption: 'Trans OCC',
                    size: '100px',
                    render: 'money',
                }
            ],
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
                    {id: 'rentablestd', type: 'menu-radio',
                        text: function(item) {
                            var el   = this.get('rentablestd:' + item.selected);
                            if (el) {
                                return "<span>Rentable: " + el.text + "</span>";
                            }
                            return "<span>Rentable: ----</span>";
                        },
                        items: [
                            {id: 1, text: "Rentable001"},
                            {id: 2, text: "Rentable002"},
                            {id: 3, text: "Rentable003"},
                        ],
                        onRefresh: function(event) {
                            console.log(event);
                        }
                    },
                    { type: 'break' },
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

                            w2ui.RARentableFeesForm.record = getRentableFeeFormInitalRecord();
                            w2ui.RARentableFeesForm.record.recid = w2ui.RARentableFeesGrid.records.length;
                            w2ui.RARentableFeesForm.refresh();
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
                    field: 'ARID',
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
                    field: 'ARName',
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
                            w2ui.RARentableFeesForm.record = $.extend(true, {}, grid.get(app.last.grid_sel_recid));

                            // showSliderContentW2UIComp(w2ui.RARentableFeesForm, RACompConfig.rentables.sliderWidth);
                            showSliderContentW2UIComp(w2ui.RARentableFeesForm, sliderContentDivLength, sliderID);
                            w2ui.RARentableFeesForm.refresh(); // need to refresh for header changes
                        };

                    // warn user if form content has been changed
                    form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
                };
            }
        });

        // new transactant form especially for this RA flow
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
                {name: 'ARID',              type: 'int',    required: true, html: {page: 0, column: 0}},
                {name: 'ARName',            type: 'text',   required: true, html: {page: 0, column: 0}},
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

                        // f.record = getRAAddTransactantFormInitRec(BID, BUD, f.record);
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

                    f.record.BID = BID;
                    f.record.BUD = BUD;

                    // there is NO PETID actually, so have to work around with recid key
                    formRefreshCallBack(f, "recid");
                };
            }
        });

    }

    // now load grid in division
    $('#ra-form #rentables').w2render(w2ui.RARentablesGrid);

    // load the existing data in rentables component
    setTimeout(function () {
        var i = getRAFlowPartTypeIndex(app.raFlowPartTypes.rentables);
        if (i >= 0 && app.raflow.data[app.raflow.activeFlowID][i].Data) {
            w2ui.RARentablesGrid.records = app.raflow.data[app.raflow.activeFlowID][i].Data;
            w2ui.RARentablesGrid.refresh();
        } else {
            w2ui.RARentablesGrid.clear();
        }
    }, 500);
};