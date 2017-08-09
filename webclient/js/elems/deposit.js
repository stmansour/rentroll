/*global
    parseInt, w2ui, app
*/

"use strict";

function getDepoInitRecord(BID, BUD){
    return {
        recid: 0,
        DEPID: 0,
        BID: BID,
        BUD: BUD,
        LID: 0,
        Name: "",
        AccountNo: "",
    };
}


//---------------------------------------------------------------------------------
// buildDepositElements - changes the main view of the program to a grid with
//                variable name svc + 'Grid'
//
//---------------------------------------------------------------------------------
function buildDepositElements() {
    //------------------------------------------------------------------------
    //          deposit Grid
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'depositGrid',
        url: '/v1/deposit',
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
            toolbarReload  : true,
            toolbarColumns : true,
        },
        columns: [
            {field: 'recid',        hidden: true,  caption: 'recid',        size: '40px',  sortable: true},
            {field: 'DID',          hidden: false, caption: 'DID',          size: '50px',  sortable: true, style: 'text-align: center'},
            {field: 'BID',          hidden: false, caption: 'BID',          size: '50px',  sortable: true, style: 'text-align: center'},
            {field: 'BUD',          hidden: false, caption: 'BUD',          size: '50px',  sortable: true, style: 'text-align: center'},
            {field: 'DEPID',        hidden: true,  caption: 'DEPID',        size: '50px',  sortable: true, style: 'text-align: center'},
            {field: 'DEPName',      hidden: false, caption: 'Depository',   size: '80px',  sortable: true, style: 'text-align: center'},
            {field: 'DPMID',        hidden: true,  caption: 'DPMID',        size: '50px',  sortable: true, style: 'text-align: center'},
            {field: 'DPMName',      hidden: false, caption: 'Method',       size: '150px',  sortable: true, style: 'text-align: center'},
            {field: 'Dt',           hidden: false, caption: 'Date',         size: '100px', sortable: true, style: 'text-align: center'},
            {field: 'Amount',       hidden: false, caption: 'Amount',       size: '100px', sortable: true, style: 'text-align: right', render: 'money'},
            {field: 'ClearedAmount',hidden: false, caption: 'ClearedAmount',size: '100px', sortable: true, style: 'text-align: right', render: 'money'},
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
                        f = w2ui.depositForm;

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
                            setToForm('depositForm', '/v1/deposit/' + BID + '/0', 400);
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

    addDateNavToToolbar('deposit');


    //------------------------------------------------------------------------
    //          deposit Form
    //------------------------------------------------------------------------
    $().w2form({
        name: 'depositForm',
        style: 'border: 0px; background-color: transparent;',
        header: 'Deposit Detail',
        //url: '/v1/deposit',
        // formURL: '/webclient/html/formdep.html',
        fields: [
            { field: 'recid', type: 'int', required: false, html: { page: 0, column: 0 } },
            { field: 'DID', type: 'int', required: false, html: { page: 0, column: 0 } },
            { field: 'BID', type: 'int', required: false, html: { page: 0, column: 0 } },
            { field: 'BUD', type: 'list', required: true, options: {items: app.businesses}, html: { page: 0, column: 0 } },
            { field: 'LID', type: 'list', required: true, options: { items: [], selected: {}, maxDropHeight: 200 }, html: { page: 0, column: 0 } },
            // { field: 'Transmittal', type: 'list', required: true, options: {items: app.businesses}, html: { page: 0, column: 0 } },
            { field: 'WrittenAmount', type: 'Total For Deposit', required: true, options: { items: [], selected: {}, maxDropHeight: 200 }, html: { page: 0, column: 0 } },
            { field: 'ClearedAmount', type: 'Cleared Amount', required: false, html: { page: 0, column: 0 } },
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
                            w2ui.depGrid.render();
                        };
                    form_dirty_alert(yes_callBack, no_callBack);
                    break;
                }
            },
        },
        actions: {
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
         },
        onRefresh: function(event) {
            event.onComplete = function() {
                var f = this,
                    header = "Edit Deposit ({0})";

                formRefreshCallBack(f, "DID", header);
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

    //------------------------------------------------------------------------
    //  depositLayout - The layout to contain the depositForm and depositDetailGrid
    //               top  - depositForm
    //               main - depositDetailGrid
    //------------------------------------------------------------------------
    $().w2layout({
        name: 'depositLayout',
        padding: 0,
        panels: [
            { type: 'left',    size: '30%', hidden: true },
            { type: 'top',     size: 300,   hidden: false, content: 'top',  resizable: true, style: app.pstyle },
            { type: 'main',    size: '70%', hidden: false, content: 'main', resizable: true, style: app.pstyle },
            { type: 'preview', size: 0,     hidden: true,  content: 'PREVIEW'  },
            { type: 'bottom',  size: 0,     hidden: true },
            { type: 'right',   size: 0,     hidden: true }
        ]
    });
}
