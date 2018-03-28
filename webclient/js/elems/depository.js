/*global
    getDepoInitRecord
*/
"use strict";
window.getDepoInitRecord = function (BID, BUD){
    return {
        recid: 0,
        DEPID: 0,
        BID: BID,
        BUD: BUD,
        LID: 0,
        Name: "",
        AccountNo: "",
    };
};

window.buildDepositoryElements = function() {

    //------------------------------------------------------------------------
    //          depository Grid
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'depGrid',
        url: '/v1/dep',
        multiSelect: false,
        show: {
            toolbar        : true,
            footer         : true,
            toolbarAdd     : true,   // indicates if toolbar add new button is visible
            toolbarDelete  : false,   // indicates if toolbar delete button is visible
            toolbarSave    : false,   // indicates if toolbar save button is visible
            selectColumn   : false,
            expandColumn   : false,
            toolbarEdit    : false,
            toolbarSearch  : false,
            toolbarInput   : false,
            searchAll      : false,
            toolbarReload  : false,
            toolbarColumns : false,
        },
        columns: [
            {field: 'recid', hidden: true, caption: 'recid', size: '40px', sortable: true},
            {field: 'DEPID', hidden: true, caption: 'DEPID', size: '150px', sortable: true, style: 'text-align: center'},
            {field: 'BID', hidden: true, caption: 'BID', size: '150px', sortable: true, style: 'text-align: center'},
            {field: 'BUD', hidden: true, caption: 'BUD', size: '150px', sortable: true, style: 'text-align: center'},
            {field: 'LID', hidden: true, caption: 'LID', size: '150px', sortable: true, style: 'text-align: center'},
            {field: 'AccountNo', hidden: false, caption: 'Account Number', size: '150px', sortable: true, style: 'text-align: right'},
            {field: 'Name', hidden: false, caption: 'Name', size: '20%', sortable: true, style: 'text-align: left'},
            {field: 'LdgrName', hidden: false, caption: 'GL Acct Name', size: '20%', sortable: true, style: 'text-align: left'},
            {field: 'GLNumber', hidden: false, caption: 'GL Number', size: '150px', sortable: true, style: 'text-align: right'},
        ],
        onRefresh: function(event) {
            event.onComplete = function() {
                var sel_recid = parseInt(this.last.sel_recid);
                if (app.active_grid == this.name && sel_recid > -1) {
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

                        // get selected element as multiSelect is false
                        var rec = grid.get(recid),
                            x = getCurrentBusiness(),
                            BID=parseInt(x.value),
                            BUD = getBUDfromBID(BID);

                        // get gl account list
                        getAccountsList(rec.BID)
                        .done(function(data){
                            if (data.status != 'success') {
                                w2ui.depositoryForm.message(data.message);
                            } else {
                                var gl_selected = {id: rec.LID, text: rec.GLNumber+" "+"(" + rec.LdgrName + ")"};
                                // get gl accounts for BUD
                                w2ui.depositoryForm.get('LID').options.items = app.gl_accounts[BUD];
                                w2ui.depositoryForm.get('LID').options.selected = gl_selected;
                            }
                            setToForm('depositoryForm', '/v1/dep/' + rec.BID + '/' + rec.DEPID, 400, true);
                        })
                        .fail(function(/*data*/){
                            console.log("Failed to get glAccountList");
                        });
                    };

                // warn user if form content has been changed
                form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
            };
        },
        onAdd: function (/*event*/) {
            var yes_args = [this],
                no_callBack = function() { return false; },
                yes_callBack = function(grid) {
                    // reset it
                    app.last.grid_sel_recid = -1;
                    grid.selectNone();

                    // Insert an empty record...
                    var x = getCurrentBusiness(),
                        BID=parseInt(x.value),
                        BUD = getBUDfromBID(BID),
                        f = w2ui.depositoryForm;

                    // get latest gl accounts first
                    getAccountsList(BID)
                    .done(function(data) {

                        if (data.status != 'success') {
                            f.message(data.message);
                            w2ui.toplayout.content('right', f);
                            w2ui.toplayout.show('right', true);
                            w2ui.toplayout.sizeTo('right', 700);
                            return;
                        } else {
                            var record = getDepoInitRecord(BID, BUD);

                            var gl_accounts_pre_selected = {id: 0, text: " -- Select GL Account -- "};
                            var gl_accounts_items = [gl_accounts_pre_selected];
                            // get gl account list for BUD from `gl_accounts` key of `app`
                            gl_accounts_items = gl_accounts_items.concat(app.gl_accounts[BUD]);

                            f.get('LID').options.items = gl_accounts_items;
                            f.get('LID').options.selected = gl_accounts_pre_selected;
                            f.record = record;
                            f.refresh();
                            setToForm('depositoryForm', '/v1/dep/' + BID + '/0', 400);
                        }
                    })
                    .fail( function() {
                        console.log('unable to get GLAccounts');
                        return;
                     });
                };  // yes callback

            // warn user if form content has been changed
            form_dirty_alert(yes_callBack, no_callBack, yes_args);
        },
    });

    //------------------------------------------------------------------------
    //          depository Form
    //------------------------------------------------------------------------
    $().w2form({
        name: 'depositoryForm',
        style: 'border: 0px; background-color: transparent;',
        header: 'Deposit Account Detail',
        url: '/v1/dep',
        formURL: '/webclient/html/formdepository.html',
        fields: [
            { field: 'recid', type: 'int', required: false, html: { page: 0, column: 0 } },
            { field: 'DEPID', type: 'int', required: false, html: { page: 0, column: 0 } },
            { field: 'BID', type: 'int', required: false, html: { page: 0, column: 0 } },
            { field: 'BUD', type: 'list', required: true, options: {items: app.businesses}, html: { page: 0, column: 0 } },
            { field: 'LID', type: 'list', required: true, options: { items: [], selected: {}, maxDropHeight: 200 }, html: { page: 0, column: 0 } },
            { field: 'Name', type: 'text', required: true, html: { page: 0, column: 0 } },
            { field: 'AccountNo', type: 'text', required: false, html: { page: 0, column: 0 } },
            { field: 'LastModTime', type: 'time', required: false, html: { page: 0, column: 0 } },
            { field: 'LastModBy', type: 'int', required: false, html: { page: 0, column: 0 } },
            { field: 'CreateTS', type: 'time', required: false, html: { page: 0, column: 0 } },
            { field: 'CreateBy', type: 'int', required: false, html: { page: 0, column: 0 } },
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
                            w2ui.depGrid.render();
                        };
                    form_dirty_alert(yes_callBack, no_callBack);
                    break;
                }
            },
        },
        actions: {
            saveadd: function() {
                var f = this,
                    grid = w2ui.depGrid,
                    x = getCurrentBusiness(),
                    BID=parseInt(x.value),
                    BUD = getBUDfromBID(BID);

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

                    var gl_accounts_pre_selected = {id: 0, text: " -- Select GL Account -- "};
                    var gl_accounts_items = [gl_accounts_pre_selected];
                    // get gl account list for BUD from `gl_accounts` key of `app`
                    gl_accounts_items = gl_accounts_items.concat(app.gl_accounts[BUD]);

                    f.get('LID').options.items = gl_accounts_items;
                    f.get('LID').options.selected = gl_accounts_pre_selected;

                    // add new empty record and just refresh the form, don't need to do CLEAR form
                    var record = getDepoInitRecord(BID, BUD);

                    f.record = record;
                    f.header = "Edit Depository Account (new)"; // have to provide header here, otherwise have to call refresh method twice to get this change in form
                    f.url = '/v1/dep/' + BID + '/0';
                    f.refresh();
                });
            },
            save: function (/*target, data*/) {
                var f = this,
                    tgrid = w2ui.depGrid;

                f.save({}, function (data) {
                    if (data.status == 'error') {
                        console.log('ERROR: '+ data.message);
                        return;
                    }
                    w2ui.toplayout.hide('right',true);
                    tgrid.render();
                });
            },
            delete: function(/*target, data*/) {

                var form = this;

                w2confirm(delete_confirm_options)
                .yes(function() {

                    var tgrid = w2ui.depGrid;
                    var params = {cmd: 'delete', formname: form.name, ID: form.record.DEPID };
                    var dat = JSON.stringify(params);

                    // delete Depository request
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
                        form.error("Delete Depository Account failed.");
                        return;
                    });
                })
                .no(function() {
                    return;
                });
            },
        },
        onRefresh: function(event) {
            event.onComplete = function() {
                var f = this,
                    header = "Edit Depository Account ({0})";

                formRefreshCallBack(f, "DEPID", header);
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
        onSubmit: function(target, data) {
            delete data.postData.record.LastModTime;
            delete data.postData.record.LastModBy;
            delete data.postData.record.CreateTS;
            delete data.postData.record.CreateBy;
            // modify form data for server request
            getFormSubmitData(data.postData.record);
        },
    });
};
