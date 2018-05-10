/* global
    w2ui,console,$,setDefaultFormFieldAsPreviousRecord,w2uiDateControlString,
    app,getBUDfromBID,getPostAccounts,setToForm,getCurrentBusiness,form_dirty_alert,
    w2confirm,w2utils,getFormSubmitData,int_to_bool,formRefreshCallBack,formRecDiffer,
    getARRulesInitRecord
*/
"use strict";
window.getARRulesInitRecord = function (BID, BUD, post_accounts_pre_selected, previousFormRecord){
    var y1 = new Date();
    var y = new Date(y1.getFullYear(), 0, 1, 0,0,0);
    var ny = new Date(9999, 11, 31, 0, 0, 0);

    var defaultFormData = {
        recid: 0,
        BID: BID,
        BUD: BUD,
        ARID: 0,
        ARType: -1,
        DebitLID: post_accounts_pre_selected,
        CreditLID: post_accounts_pre_selected,
        Name: '',
        Description: '',
        DtStart: w2uiDateControlString(y),
        DtStop: w2uiDateControlString(ny),
        PriorToRAStart: true,
        PriorToRAStop: true,
        ApplyRcvAccts: false,
        RAIDrqd: false,
        AutoPopulateToNewRA: false,
        IsRentAR: false,
        DefaultAmount: 0.0,
    };


    // if it called after 'save and add another' action there previous form record is passed as Object
    // else it is null
    if ( previousFormRecord ) {
        defaultFormData = setDefaultFormFieldAsPreviousRecord(
            [ 'Name', 'Description'], // Fields to Reset
            defaultFormData,
            previousFormRecord
        );
    }

    return defaultFormData;
};

window.buildARElements = function () {

//------------------------------------------------------------------------
//          Account Rules Grid
//------------------------------------------------------------------------
$().w2grid({
    name: 'arsGrid',
    url: '/v1/ars',
    multiSelect: false,
    show: {
        toolbar        : true,
        footer         : true,
        toolbarAdd     : true,    // indicates if toolbar add new button is visible
        toolbarDelete  : false,   // indicates if toolbar delete button is visible
        toolbarSave    : false,   // indicates if toolbar save button is visible
        selectColumn   : false,
        expandColumn   : false,
        toolbarEdit    : false,
        toolbarSearch  : false,
        toolbarInput   : false,
        searchAll      : false,
        toolbarReload  : false,
        toolbarColumns : true,
    },
    columns: [
        {field: 'recid',  hidden: true,  caption: 'recid', size: '40px',  sortable: true},
        {field: 'ARID',   hidden: true,  caption: 'ARID', size: '75px', sortable: true, style: 'text-align: center'},
        {field: 'BID',    hidden: true,  caption: 'BID',   size: '150px', sortable: true, style: 'text-align: center'},
        {field: 'Name',   hidden: false, caption: 'Name',  size: '15%',   sortable: true, style: 'text-align: left'},
        {field: 'ARType', hidden: false, caption: 'ARType',size: '100px',  sortable: true, style: 'text-align: left',
            render: function (record, index, col_index) {
                return app.ARTypes[this.getCellValue(index, col_index)];
            }
        },
        {field: 'DebitLID',        hidden: true,  caption: 'DebitLID',   size: '50px', sortable: true},
        {field: 'DebitLedgerName', hidden: false, caption: 'Debit',      size: '200px',sortable: true, style: 'text-align: left'},
        {field: 'CreditLID',       hidden: true,  caption: 'CreditLID',  size: '50px', sortable: true},
        {field: 'CreditLedgerName',hidden: false, caption: 'Credit',     size: '200px',sortable: true, style: 'text-align: left'},
        {field: 'DtStart',                        caption: 'Start',      size: '80px', sortable: true, style: 'text-align: right'},
        {field: 'DtStop',                         caption: 'Stop',       size: '80px', sortable: true, style: 'text-align: right'},
        {field: 'Description',     hidden: false, caption: 'Description',size: '20%',  sortable: true, style: 'text-align: left'},
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

                    // multi select is false
                    var rec = grid.get(recid),
                        BUD = getBUDfromBID(rec.BID);

                    // set ARType DropDown list
                    var artype_selected,
                        artype_items = [];

                    Object.keys(app.ARTypes).forEach(function(id){
                        var artype_id = parseInt(id);
                        if (rec.ARType == artype_id) {
                            artype_selected = {id: artype_id, text: app.ARTypes[artype_id]};
                        }
                        artype_items.push({id: artype_id, text: app.ARTypes[artype_id]});
                    });

                    // get gl account list
                    getPostAccounts(rec.BID)
                    .done(function(/*data*/){
                        // set fields
                        w2ui.arsForm.get('ARType').options.items=artype_items;
                        w2ui.arsForm.get('ARType').options.selected=artype_selected;
                        w2ui.arsForm.get('DebitLID').options.items=app.post_accounts[BUD];
                        w2ui.arsForm.get('DebitLID').options.selected={id: rec.DebitLID, text: rec.DebitLedgerName};
                        w2ui.arsForm.get('CreditLID').options.items=app.post_accounts[BUD];
                        w2ui.arsForm.get('CreditLID').options.selected={id: rec.CreditLID, text: rec.CreditLedgerName};
                        setToForm('arsForm', '/v1/ar/' + rec.BID + '/' + rec.ARID, 400, true);
                    })
                    .fail(function(/*data*/){
                        console.log("Failed to get glAccountList");
                    });
                };

            // alert user if form content has been changed
            form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
        };
    },
    onAdd: function (/*event*/) {
        var yes_args = [this],
            no_callBack = function() { return false; },
            yes_callBack = function(grid) {
                // first reset grid sel recid
                app.last.grid_sel_recid = -1;
                grid.selectNone();

                // Insert an empty record...
                var x = getCurrentBusiness(),
                    BID=parseInt(x.value),
                    BUD = getBUDfromBID(BID);

                // get latest gl accounts first
                getPostAccounts(BID)
                .done(function(data) {
                    if (data.status != 'success') {
                        w2ui.arsForm.message(data.message);
                        w2ui.toplayout.content('right', w2ui.arsForm);
                        w2ui.toplayout.show('right', true);
                        w2ui.toplayout.sizeTo('right', 700);
                        return;
                    }
                    else {
                        var artype_pre_selected = {id: -1, text: " -- Select AR type -- "};
                        var artype_items = [artype_pre_selected];
                        Object.keys(app.ARTypes).forEach(function(id){
                            var artype_id = parseInt(id);
                            artype_items.push({id: artype_id, text: app.ARTypes[artype_id]});
                        });

                        var post_accounts_pre_selected = {id: 0, text: " -- Select GL Account -- "};
                        var post_accounts_items = [post_accounts_pre_selected];
                        post_accounts_items = post_accounts_items.concat(app.post_accounts[BUD]);

                        w2ui.arsForm.get('ARType').options.items = artype_items;
                        w2ui.arsForm.get('ARType').options.selected = artype_pre_selected;
                        w2ui.arsForm.get('DebitLID').options.items = post_accounts_items;
                        w2ui.arsForm.get('DebitLID').options.selected = post_accounts_pre_selected;
                        w2ui.arsForm.get('CreditLID').options.items = post_accounts_items;
                        w2ui.arsForm.get('CreditLID').options.selected = post_accounts_pre_selected;
                        // w2ui.arsForm.refresh();
                        var record = getARRulesInitRecord(BID, BUD, post_accounts_pre_selected, null);
                        w2ui.arsForm.record = record;
                        w2ui.arsForm.refresh();
                        setToForm('arsForm', '/v1/ar/' + BID + '/0', 400);
                    }
                })
                .fail( function() {
                    console.log('unable to get GLAccounts');
                    return;
                 });
            };

        // alert user if form content has been changed
        form_dirty_alert(yes_callBack, no_callBack, yes_args);
    },
});

    //------------------------------------------------------------------------
    //          Account Rules Form
    //------------------------------------------------------------------------
    $().w2form({
        name: 'arsForm',
        style: 'border: 0px; background-color: transparent;',
        header: 'Edit Account Rule',
        url: '/v1/ar',
        formURL: '/webclient/html/formar.html',
        fields: [
            { field: 'recid',                type: 'int',      required: false, html: { page: 0, column: 0 } },
            { field: 'ARID',                 type: 'int',      required: false, html: { page: 0, column: 0 } },
            { field: 'BID',                  type: 'int',      required: true,  html: { page: 0, column: 0 } },
            { field: 'BUD',                  type: 'list',     required: true,  html: { page: 0, column: 0 }, options: { items: app.businesses } },
            { field: 'Name',                 type: 'text',     required: true,  html: { page: 0, column: 0 } },
            { field: 'ARType',               type: 'list',     required: true,  html: { page: 0, column: 0 }, options: { items: [], selected: {}, maxDropHeight: 200 } },
            { field: 'DebitLID',             type: 'list',     required: true,  html: { page: 0, column: 0 }, options: { items: [], selected: {}, maxDropHeight: 200 } },
            { field: 'CreditLID',            type: 'list',     required: true,  html: { page: 0, column: 0 }, options: { items: [], selected: {}, maxDropHeight: 200 } },
            { field: 'Description',          type: 'text',     required: false, html: { page: 0, column: 0} },
            { field: 'DtStart',              type: 'date',     required: true,  html: { page: 0, column: 0 } },
            { field: 'DtStop',               type: 'date',     required: true,  html: { page: 0, column: 0 } },
            { field: 'PriorToRAStart',       type: 'checkbox', required: false, html: { page: 0, column: 0 } },
            { field: 'PriorToRAStop',        type: 'checkbox', required: false, html: { page: 0, column: 0 } },
            { field: 'ApplyRcvAccts',        type: 'checkbox', required: false, html: { page: 0, column: 0 } },
            { field: 'RAIDrqd',              type: 'checkbox', required: false, html: { page: 0, column: 0 } },
            { field: 'DefaultAmount',        type: 'money',    required: true,  html: { page: 0, column: 0 } },
            { field: 'AutoPopulateToNewRA',  type: 'checkbox', required: false, html: { page: 0, column: 0 } },
            { field: 'IsRentAR',             type: 'checkbox', required: false, html: { page: 0, column: 0 } },
            { field: "LastModTime",          type: 'time',     required: false, html: { caption: "LastModTime", page: 0, column: 0 } },
            { field: "LastModBy",            type: 'int',      required: false, html: { caption: "LastModBy", page: 0, column: 0 } },
            { field: "CreateTS",             type: 'time',     required: false, html: { caption: "CreateTS", page: 0, column: 0 } },
            { field: "CreateBy",             type: 'int',      required: false, html: { caption: "CreateBy", page: 0, column: 0 } },
        ],
        toolbar: {
            items: [
                { id: 'btnNotes', type: 'button', icon: 'far fa-sticky-note' },
                { id: 'bt3', type: 'spacer' },
                { id: 'btnClose', type: 'button', icon: 'fas fa-times' },
            ],
            onClick: function (event) {
                switch(event.target) {
                case 'btnClose':
                    var no_callBack = function() { return false; },
                        yes_callBack = function() {
                            w2ui.toplayout.hide('right',true);
                            w2ui.arsGrid.render();
                        };
                    form_dirty_alert(yes_callBack, no_callBack);
                    break;
                }
            },
        },
        onValidate: function(event) {
            if (this.record.DtStart === '') {
                event.errors.push({
                    field: this.get('DtStart'),
                    error: 'Start date cannot be blank'
                });
            }
            if (this.record.DtStop === '') {
                event.errors.push({
                    field: this.get('DtStop'),
                    error: 'Stop date cannot be blank'
                });
            }
        },
        actions: {
            saveadd: function() {
                var f = this,
                    grid = w2ui.arsGrid,
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

                    // JUST RENDER THE GRID ONLY
                    grid.render();

                    // add new empty record and just refresh the form, don't need to do CLEAR form
                    var artype_items = [];
                    Object.keys(app.ARTypes).forEach(function(id){
                        var artype_id = parseInt(id);
                        artype_items.push({id: artype_id, text: app.ARTypes[artype_id]});
                    });

                    var post_accounts_pre_selected = {id: 0, text: " -- Select GL Account -- "};
                    var post_accounts_items = [post_accounts_pre_selected];
                    post_accounts_items = post_accounts_items.concat(app.post_accounts[BUD]);

                    w2ui.arsForm.get('ARType').options.items = artype_items;
                    // w2ui.arsForm.get('ARType').options.selected = artype_pre_selected;
                    w2ui.arsForm.get('DebitLID').options.items = post_accounts_items;
                    w2ui.arsForm.get('DebitLID').options.selected = post_accounts_pre_selected;
                    w2ui.arsForm.get('CreditLID').options.items = post_accounts_items;
                    w2ui.arsForm.get('CreditLID').options.selected = post_accounts_pre_selected;
                    var record = getARRulesInitRecord(BID, BUD, post_accounts_pre_selected, f.record);
                    f.record = record;
                    f.header = "Edit Account Rule (new)"; // have to provide header here, otherwise have to call refresh method twice to get this change in form
                    f.url = '/v1/ar/' + BID+'/0';
                    f.refresh();
                });
            },
            delete: function() {
                var form = this;
                w2confirm(delete_confirm_options)
                .yes(function () {
                    var tgrid = w2ui.arsGrid;
                    var params = {cmd: 'delete', formname: form.name, ARID: form.record.ARID };
                    var dat = JSON.stringify(params);

                    // delete AR request
                    $.post(form.url, dat, null, "json")
                    .done(function(data) {
                        if (data.status === "error") {
                            form.error(w2utils.lang(data.message));
                            return;
                        }
                        w2ui.toplayout.hide('right',true);
                        tgrid.remove(app.last.grid_sel_recid);
                        tgrid.render();
                    })
                    .fail(function(/*data*/){
                        form.error("Delete AR failed.");
                        return;
                    });
                })
                .no(function () {
                    return;
                });
            },
            save: function () {
                //var obj = this;
                var tgrid = w2ui.arsGrid;
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
        },
        onSubmit: function(target, data){
            delete data.postData.record.LastModTime;
            delete data.postData.record.LastModBy;
            delete data.postData.record.CreateTS;
            delete data.postData.record.CreateBy;
            getFormSubmitData(data.postData.record);
            // object to value before submit to server
            data.postData.record.PriorToRAStart = int_to_bool(data.postData.record.PriorToRAStart);
            data.postData.record.PriorToRAStop = int_to_bool(data.postData.record.PriorToRAStop);
            data.postData.record.ApplyRcvAccts = int_to_bool(data.postData.record.ApplyRcvAccts);
            data.postData.record.RAIDrqd = int_to_bool(data.postData.record.RAIDrqd);
            data.postData.record.AutoPopulateToNewRA = int_to_bool(data.postData.record.AutoPopulateToNewRA);
            data.postData.record.IsRentAR = int_to_bool(data.postData.record.IsRentAR);
        },
        onRefresh: function(event) {
            event.onComplete = function() {
                var f = this,
                    header = "Edit Account Rule ({0})";
                formRefreshCallBack(f, "ARID", header);
                var b = ("Receipt" === f.record.ARType.text && f.record.ApplyRcvAccts);
                $(f.box).find("input[name=RAIDrqd]").prop( "disabled", !b);
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
                var f = this;
                var b = ("Receipt" === f.record.ARType.text && f.record.ApplyRcvAccts);
                $(f.box).find("input[name=RAIDrqd]").prop( "disabled", !b);
            };
        }
    });
};
