/*global
    w2ui, app, $, console, setToForm, form_dirty_alert, formRefreshCallBack, formRecDiffer,
    getFormSubmitData, w2confirm, delete_confirm_options, getBUDfromBID, getCurrentBusiness, 
    addDateNavToToolbar, 
*/
"use strict";
function getRTInitRecord(BID, BUD){
    var y = new Date();
    return {
        recid: 0,
        BID: BID,
        BUD: BUD,
        RTID: 0,
        Style: "",
        Name: "",
        RentCycle: 0,
        Proration: 0,
        GSRPC: 0,
        ManageToBudget: 0,
        RMRID: 0,
        MarketRate: 0.0,
        LastModTime: y.toISOString(),
        LastModBy: 0,
    };
}

function buildRentableTypeElements() {

    //------------------------------------------------------------------------
    //          rentable types Grid
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'rtGrid',
        url: '/v1/rt',
        multiSelect: false,
        show: {
            header: false,
            toolbar: true,
            toolbarReload: false,
            toolbarColumns: false,
            toolbarSearch: true,
            toolbarAdd: true,
            toolbarDelete: false,
            toolbarSave: false,
            searchAll       : true,
            footer: true,
            lineNumbers: false,
            selectColumn: false,
            expandColumn: false
        },
        columns: [
            {field: 'recid', caption: 'recid', hidden: true},
            {field: 'RTID', caption: 'RTID', size: '50px', sortable: true},
            {field: 'Name', caption: 'Name', size: '150px', sortable: true},
            {field: 'Style', caption: 'Style', size: '100px', sortable: true},
            {field: 'BID', caption: 'BID', hidden: true},
            {field: 'BUD', caption: 'BUD', hidden: true},
            {field: 'RentCycle', caption: 'RentCycle', size: '75px', sortable: true,
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
            {field: 'Proration', caption: 'Proration', size: '90px', sortable: true,
                render: function (record/*, index, col_index*/) {
                    var text = '';
                    if (record) {
                        app.cycleFreq.forEach(function(itemText, itemIndex) {
                            if (record.Proration == itemIndex) {
                                text = itemText;
                                return false;
                            }
                        });
                    }
                    return text;
                },
            },
            {field: 'GSRPC', caption: 'GSRPC', size: '65px', sortable: true,
                render: function (record/*, index, col_index*/) {
                    var text = '';
                    if (record) {
                        app.cycleFreq.forEach(function(itemText, itemIndex) {
                            if (record.GSRPC == itemIndex) {
                                text = itemText;
                                return false;
                            }
                        });
                    }
                    return text;
                },
            },
            {field: 'ManageToBudget', caption: 'ManageToBudget', size: '200px', sortable: true,
                render: function (record/*, index, col_index*/) {
                    var text = '';
                    if (record) {
                        app.manageToBudgetList.forEach(function(item) {
                            if (record.ManageToBudget == item.id) {
                                text = item.text;
                                return false;
                            }
                        });
                    }
                    return text;
                },
            },
            {field: 'LastModTime', caption: 'LastModTime', hidden: true},
            {field: 'LastModBy', caption: 'LastModBy', hidden: true},
            {field: 'RMRID', caption: 'RMRID', hidden: true},
            {field: 'MarketRate', caption: 'MarketRate', size: '100px', sortable: true, render: 'money'},
        ],
        onRefresh: function(event) {
            event.onComplete = function() {
                if (app.active_grid == this.name) {
                    if (app.new_form_rec) {
                        this.selectNone();
                    }
                    else{
                        this.select(app.last.grid_sel_recid);
                    }
                }
            };
        },
        onClick: function(event) {
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

                        var rec = grid.get(recid);
                        console.log('rentableType form url: ' + '/v1/rt/' + rec.BID + '/' + rec.RTID);
                        setToForm('rentableTypeForm', '/v1/rt/' + rec.BID + '/' + rec.RTID, 400, true);
                    };

                // warn user if content has been changed
                form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
            };
        },
        onAdd: function(/*event*/) {
            var yes_args = [this],
                no_callBack = function() { return false; },
                yes_callBack = function(grid) {
                    // reset it
                    app.last.grid_sel_recid = -1;
                    grid.selectNone();

                    var x = getCurrentBusiness(),
                        BID=parseInt(x.value),
                        BUD = getBUDfromBID(BID);

                    var record = getRTInitRecord(BID, BUD);
                    w2ui.rentableTypeForm.record = record;
                    w2ui.rentableTypeForm.refresh();
                    setToForm('rentableTypeForm', '/v1/rt/' + BID + '/0', 400);
                };

            // warn user if content has been changed of form
            form_dirty_alert(yes_callBack, no_callBack, yes_args);
        },
    });

    addDateNavToToolbar('rt');

    //------------------------------------------------------------------------
    //          rentable Type Form
    //------------------------------------------------------------------------
    $().w2form({
        name: 'rentableTypeForm',
        style: 'border: 0px; background-color: transparent;',
        header: app.sRentableType + ' Detail',
        url: '/v1/rentabletypes',
        formURL: '/webclient/html/formrt.html',
        fields: [
            { field: 'recid', type: 'int', required: false, html: { page: 0, column: 0 } },
            { field: 'RTID', type: 'int', required: true, html: { page: 0, column: 0 } },
            { field: 'BID', type: 'int', required: true, html: { page: 0, column: 0 } },
            { field: 'BUD', type: 'list', options: {items: app.businesses}, required: true, html: { page: 0, column: 0 } },
            { field: 'Style', type: 'text', required: true, html: { page: 0, column: 0 } },
            { field: 'Name', type: 'text', required: true, html: { page: 0, column: 0 } },
            { field: 'RentCycle', type: 'list', options: {items: app.cycleFreq, selected: {}}, required: true, html: { page: 0, column: 0 } },
            { field: 'Proration', type: 'list', options: {items: app.cycleFreq, selected: {}}, required: true, html: { page: 0, column: 0 } },
            { field: 'GSRPC', type: 'list', options: {items: app.cycleFreq, selected: {}}, required: true, html: { page: 0, column: 0 } },
            { field: 'ManageToBudget', type: 'list', options: {items: app.manageToBudgetList, selected: {}}, required: true, html: { page: 0, column: 0 } },
            { field: 'RMRID', type: 'int', required: true, html: { page: 0, column: 0 } },
            { field: 'MarketRate', type: 'money', required: false, html: { page: 0, column: 0 } },
            { field: 'LastModTime', type: 'time', required: false, html: { page: 0, column: 0 } },
            { field: 'LastModBy', type: 'int', required: false, html: { page: 0, column: 0 } },
            { field: 'CreateTS', type: 'time', required: false, html: { page: 0, column: 0 } },
            { field: 'CreateBy', type: 'int', required: false, html: { page: 0, column: 0 } },
        ],
        toolbar: {
            items: [
                { id: 'btnNotes', type: 'button', icon: 'fa fa-sticky-note-o' },
                { id: 'bt3', type: 'spacer' },
            { id: 'btnClose', type: 'button', icon: 'fa fa-times' },
        ],
        onClick: function (event) {
                switch(event.target) {
                case 'btnClose':
                    var no_callBack = function() { return false; },
                        yes_callBack = function() {
                            w2ui.toplayout.hide('right',true);
                            w2ui.rtGrid.render();
                        };
                    form_dirty_alert(yes_callBack, no_callBack);
                    break;
                }
            },
        },
        onValidate: function(event) {
            if (this.record.ManageToBudget.id === 1 && this.record.MarketRate === 0) {
                event.errors.push({
                    field: this.get('MarketRate'),
                    error: 'MarketRate cannot be blank when Mange To Budget is Yes'
                });
            }
            if (this.record.Style === "") {
                event.errors.push({
                    field: this.get('Style'),
                    error: 'Style cannot be blank'
                });
            }
            if (this.record.Name === "") {
                event.errors.push({
                    field: this.get('Name'),
                    error: 'Name cannot be blank'
                });
            }
        },
        actions: {
            saveadd: function() {
                var f = this,
                    grid = w2ui.rtGrid,
                    r = f.record,
                    x = getCurrentBusiness(),
                    BID=parseInt(x.value),
                    BUD=getBUDfromBID(BID);

                // clean dirty flag of form
                app.form_is_dirty = false;
                // clear the grid select recid
                app.last.grid_sel_recid  =-1;

                // select none if you're going to add new record
                grid.selectNone();

                f.save({}, function (data) {
                    if (data.status == 'error') {
                        console.log('ERROR: '+ data.message);
                        return;
                    }

                    // dropdown list items and selected variables
                    var rentCycleSel = {}, prorationSel = {}, gsrpcSel = {},
                        manageToBudgetSel = {}, cycleFreqItems = [];

                    // select value for rentcycle, proration, gsrpc
                    app.cycleFreq.forEach(function(itemText, itemIndex) {
                        if (itemIndex == r.RentCycle) {
                            rentCycleSel = { id: itemIndex, text: itemText };
                        }
                        if (itemIndex == r.Proration) {
                            prorationSel = { id: itemIndex, text: itemText };
                        }
                        if (itemIndex == r.GSRPC) {
                            gsrpcSel = { id: itemIndex, text: itemText };
                        }
                        cycleFreqItems.push({ id: itemIndex, text: itemText });
                    });

                    // select value for manage to budget
                    app.manageToBudgetList.forEach(function(item) {
                        if (item.id == r.ManageToBudget) {
                            manageToBudgetSel = {id: item.id, text: item.text};
                        }
                    });

                    f.get("ManageToBudget").options.items = app.manageToBudgetList;
                    f.get("ManageToBudget").options.selected = manageToBudgetSel[0];
                    f.get("RentCycle").options.items = cycleFreqItems;
                    f.get("RentCycle").options.selected = rentCycleSel[0];
                    f.get("Proration").options.items = cycleFreqItems;
                    f.get("Proration").options.selected = prorationSel[0];
                    f.get("GSRPC").options.items = cycleFreqItems;
                    f.get("GSRPC").options.selected = gsrpcSel[0];

                    // JUST RENDER THE GRID ONLY
                    grid.render();

                    var record = getRTInitRecord(BID, BUD);
                    f.record = record;
                    f.header = "Edit Rentable Type (new)"; // have to provide header here, otherwise have to call refresh method twice to get this change in form
                    f.url = '/v1/rt/' + BID+'/0';
                    f.refresh();
                });
            },
            save: function () {
                //var obj = this;
                var tgrid = w2ui.rtGrid;
                tgrid.selectNone();

                this.save({}, function (data) {
                    if (data.status == 'error') {
                        console.log('ERROR: '+ data.message);
                        return;
                    }
                    w2ui.toplayout.hide('right',true);
                    tgrid.render();
                });
            },
            delete: function() {

                var form = this;

                w2confirm(delete_confirm_options)
                .yes(function() {
                    var tgrid = w2ui.rtGrid;
                    var params = {cmd: 'delete', formname: form.name, ID: form.record.RTID };
                    var dat = JSON.stringify(params);

                    // delete Depository request
                    $.post(form.url, dat, null, "json")
                    .done(function(data) {
                        if (data.status === "error") {
                            return;
                        }

                        w2ui.toplayout.hide('right',true);
                        tgrid.remove(app.last.grid_sel_recid);
                        tgrid.render();
                    })
                    .fail(function(/*data*/){
                        form.error("Delete Payment failed.");
                        return;
                    });
                })
                .no(function() {
                    return;
                });
            },
        },
        onSubmit: function(target, data){
            delete data.postData.record.LastModTime;
            delete data.postData.record.LastModBy;
            delete data.postData.record.CreateTS;
            delete data.postData.record.CreateBy;
            // server request form data
            getFormSubmitData(data.postData.record);
        },
        onRefresh: function(event) {
            event.onComplete = function() {
                var f = this,
                    r = f.record,
                    header = "Edit Rentable Type ({0})";

                // dropdown list items and selected variables
                var rentCycleSel = {}, prorationSel = {}, gsrpcSel = {},
                    manageToBudgetSel = {}, cycleFreqItems = [];

                // select value for rentcycle, proration, gsrpc
                app.cycleFreq.forEach(function(itemText, itemIndex) {
                    if (itemIndex == r.RentCycle) {
                        rentCycleSel = { id: itemIndex, text: itemText };
                    }
                    if (itemIndex == r.Proration) {
                        prorationSel = { id: itemIndex, text: itemText };
                    }
                    if (itemIndex == r.GSRPC) {
                        gsrpcSel = { id: itemIndex, text: itemText };
                    }
                    cycleFreqItems.push({ id: itemIndex, text: itemText });
                });

                // select value for manage to budget
                app.manageToBudgetList.forEach(function(item) {
                    if (item.id == r.ManageToBudget) {
                        manageToBudgetSel = {id: item.id, text: item.text};
                    }
                });

                // fill the field with values
                f.get("RentCycle").options.items = cycleFreqItems;
                f.get("RentCycle").options.selected = rentCycleSel;
                f.get("Proration").options.items = cycleFreqItems;
                f.get("Proration").options.selected = prorationSel;
                f.get("GSRPC").options.items = cycleFreqItems;
                f.get("GSRPC").options.selected = gsrpcSel;
                f.get("ManageToBudget").options.items = app.manageToBudgetList;
                f.get("ManageToBudget").options.selected = manageToBudgetSel;

                formRefreshCallBack(f, "RTID", header);
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
        onResize: function(event) {
            event.onComplete = function() {
                // HACK: set the height of right panel of toplayout box div and form's box div
                // this is how w2ui set the content inside box of toplayout panel, and form's main('div.w2ui-form-box')
                var h = w2ui.toplayout.get("right").height;
                $(w2ui.toplayout.get("right").content.box).height(h);
                $(this.box).find("div.w2ui-form-box").height(h);
            };
        }
    });
}
