/*global
    parseInt, w2ui, getDepMeth
*/

"use strict";

function getDepoInitRecord(BID, BUD){
    var y = new Date();
    return {
        recid: 0,
        DID: 0,
        BID: BID,
        BUD: BUD,
        DEPID: 0,
        DEPName: "",
        DPMID: 0,
        DPMName: "",
        Dt: y.toISOString(),
        FLAGS: 0,
        Amount: 0.0,
        ClearedAmount: 0.0,
        LastModTime: y.toISOString(),
        LastModBy: 0,
        CreateTS: y.toISOString(),
        CreateBy: 0,
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
        onClick: function(event) {
            event.onComplete = function () {
                var yes_args = [this, event.recid],
                    no_args = [this],
                    no_callBack = function(grid) {
                        grid.select(app.last.grid_sel_recid);
                        return false;
                    },
                    yes_callBack = function(grid, recid) {
                        var x = getCurrentBusiness();
                        var Bid = x.value;
                        var Bud = getBUDfromBID(Bid);
                        $.get('/v1/uival/' + x.value + '/app.depmeth' )
                        .done( function(data) {
                            if (typeof data == 'string') {  // it's weird, a successful data add gets parsed as an object, an error message does not
                                app.depmeth = JSON.parse(data);
                                w2ui.depositForm.get('DPMName').options.items = app.depmeth[Bud];
                                w2ui.depositForm.refresh();
                                app.last.grid_sel_recid = parseInt(recid);
                                grid.select(app.last.grid_sel_recid); // keep highlighting current row in any case
                                var rec = grid.get(recid);
                                var myurl = '/v1/deposit/' + rec.BID + '/' + rec.DID;
                                setToForm("depositForm",myurl,400,true);
                            }
                            if (data.status != 'success') {w2ui.depositForm.message(data.message);}
                        })
                        .fail( function() { console.log('Error getting /v1/uival/' + x.value + '/app.depmeth'); });
                    };

                    form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args); // warn user if form content has been changed
            };
        },
        onAdd: function (/*event*/) {
            var yes_args = [this];
            var no_callBack = function() { return false; };
            var yes_callBack = function(grid) {
                // reset it
                app.last.grid_sel_recid = -1;
                grid.selectNone();

                // Insert an empty record...
                var x = getCurrentBusiness(),
                BID=parseInt(x.value),
                BUD = getBUDfromBID(BID),
                f = w2ui.depositForm;
                var record = getDepoInitRecord(BID, BUD);
                f.record = record;
                f.refresh();
                setToForm('depositForm', '/v1/deposit/' + BID + '/0', 500);
            };

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
        url: '/v1/deposit',
        formURL: '/html/formdeposit.html',

        fields: [
            { field: 'recid',         type: 'int',   required: false, html: { page: 0, column: 0 } },
            { field: 'DID',           type: 'int',   required: false, html: { page: 0, column: 0 } },
            { field: 'BID',           type: 'int',   required: false, html: { page: 0, column: 0 } },
            { field: 'DEPID',         type: 'int',   required: false, html: { page: 0, column: 0 } },
            { field: 'DPMID',         type: 'int',  required: true, html: { page: 0, column: 0 } },
            { field: 'FLAGS',         type: 'int',   required: false, html: { page: 0, column: 0 } },
            { field: 'BUD',           type: 'list',  required: true, options: {items: app.businesses}, html: { page: 0, column: 0 } },
            { field: 'Dt',            type: 'date',  required: true },
            { field: 'DEPName',       type: 'list',  required: true, options:  {items: [], selected: {}} },
            { field: 'DPMName',       type: 'list',  required: true, options:  {items: [], selected: {}} },
            { field: 'Amount',        type: 'float', required: true,  html: { page: 0, column: 0 } },
            { field: 'ClearedAmount', type: 'float', required: false, html: { page: 0, column: 0 } },
            { field: 'LastModTime',   type: 'time',  required: false, html: { page: 0, column: 0 } },
            { field: 'LastModBy',     type: 'int',   required: false, html: { page: 0, column: 0 } },
            { field: 'CreateTS',      type: 'time',  required: false, html: { page: 0, column: 0 } },
            { field: 'CreateBy',      type: 'int',   required: false, html: { page: 0, column: 0 } },
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
                            w2ui.depositGrid.render();
                        };
                    form_dirty_alert(yes_callBack, no_callBack);
                    break;
                }
            },
        },
        actions: {
            save: function (/*target, data*/) {
                var f = this,
                    tgrid = w2ui.depositGrid;

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
                var f = this;
                var header = "Edit Deposit ({0})";
                var bud = f.record.BUD.text;
                var dpmid = f.record.DPMID; 

                f.get("DPMName").options.selected = getDepMeth(bud, dpmid);
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
