/*global
    parseInt, w2ui, getDepMeth, getDepository
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
                var yes_args = [this, event.recid];
                var no_args = [this];
                var no_callBack = function(grid) {
                        grid.select(app.last.grid_sel_recid);
                        return false;
                    };
                var yes_callBack = function(grid, recid) {
                        var x = getCurrentBusiness();
                        var Bid = x.value;
                        var Bud = getBUDfromBID(Bid);

                        var getUIInfo = function(bid,x) {
                            return $.get('/v1/uival/' + bid + x );
                        };

                        $.when( getUIInfo(Bid,"/app.depmeth"),
                                getUIInfo(Bid,"/app.Depositories"))
                        .done( function(dpmArgs,depArgs) {
                            if (typeof dpmArgs[0] == 'string') {
                                app.depmeth = JSON.parse(dpmArgs[0]);
                                w2ui.depositForm.get('DPMName').options.items = app.depmeth[Bud];
                            } else if (dpmArgs[0].status != 'success') {
                                w2ui.depositForm.message(dpmArgs[0].message);
                            }

                            if (typeof depArgs[0] == 'string') {
                                app.Depositories = JSON.parse(depArgs[0]);
                                w2ui.depositForm.get('DEPName').options.items = app.Depositories[Bud];
                            } else if (depArgs[0].status != 'success') {
                                w2ui.depositForm.message(depArgs[0].message);
                            }

                            w2ui.depositForm.refresh();
                            app.last.grid_sel_recid = parseInt(recid);
                            grid.select(app.last.grid_sel_recid); // keep highlighting current row in any case
                            var rec = grid.get(recid);
                            var myurl = '/v1/deposit/' + rec.BID + '/' + rec.DID;
                            var urlgrid = '/v1/depositlist/' + rec.BID + '/' + rec.DID;
                            setToDepositForm("depositLayout","depositForm",myurl,urlgrid,700,true);
                        })
                        .fail( function() { console.log('Error getting /v1/uival/' + x.value + '/{app.depmeth | app.Depositories}'); });
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
                var getUIInfo = function(bid,x) {
                    return $.get('/v1/uival/' + bid + x );
                };

                $.when( getUIInfo(BID,"/app.depmeth"),
                        getUIInfo(BID,"/app.Depositories"))
                .done( function(dpmArgs,depArgs) {
                    if (typeof dpmArgs[0] == 'string') {
                        app.depmeth = JSON.parse(dpmArgs[0]);
                        w2ui.depositForm.get('DPMName').options.items = app.depmeth[BUD];
                    } else if (dpmArgs[0].status != 'success') {
                        w2ui.depositForm.message(dpmArgs[0].message);
                    }

                    if (typeof depArgs[0] == 'string') {
                        app.Depositories = JSON.parse(depArgs[0]);
                        w2ui.depositForm.get('DEPName').options.items = app.Depositories[BUD];
                    } else if (depArgs[0].status != 'success') {
                        w2ui.depositForm.message(depArgs[0].message);
                    }
                })
                .fail( function() { console.log('Error getting /v1/uival/' + x.value + '/{app.depmeth | app.Depositories}'); });

                f.refresh();
                setToDepositForm('depositLayout', 'depositForm', '/v1/deposit/' + BID + '/0','', 700);
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
            { field: 'DPMID',         type: 'int',   required: true, html: { page: 0, column: 0 } },
            { field: 'FLAGS',         type: 'int',   required: false, html: { page: 0, column: 0 } },
            { field: 'BUD',           type: 'list',  required: true, options: {items: app.businesses}, html: { page: 0, column: 0 } },
            { field: 'Dt',            type: 'date',  required: true },
            { field: 'DEPName',       type: 'list',  required: true, options:  {items: [], selected: {}} },
            { field: 'DPMName',       type: 'list',  required: true, options:  {items: [], selected: {}} },
            { field: 'Amount',        type: 'money', required: true,  html: { page: 0, column: 0 } },
            { field: 'ClearedAmount', type: 'money', required: false, html: { page: 0, column: 0 } },
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
                var r = f.record;
                var header = "Edit Deposit ({0})";
                var bud = r.BUD.text;
                var dpmid = r.DPMID;
                var depid = r.DEPID;

                f.get("DPMName").options.selected = getDepMeth(bud, dpmid);
                f.get("DEPName").options.selected = getDepository(bud, depid);
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
    //  depositListGrid - For new deposits, it lists all that are not 
    //                    currently part of a deposit. Any combination can be
    //                    selected to be part of the new deposit.
    //
    //                    For existing deposits, it lists the receipts that
    //                    belong to the deposit.
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'depositListGrid',
        url: '/v1/depositlist',
        multiSelect: false,
        show: {
            toolbar        : false,
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
            {field: 'recid',    caption: 'recid',        hidden: true,  size: '40px',  sortable: true  },
            {field: 'RCPTID',   caption: 'Receipt ID',   hidden: false, size: '80px',  sortable: true, style: 'text-align: right'},
            {field: 'Dt',       caption: 'Date',         hidden: false, size: '80px',  sortable: true, style: 'text-align: right'},
            {field: 'ARID',     caption: 'ARID',         hidden: true,  size: '150px', sortable: false },
            {field: 'AcctRule', caption: 'Account Rule', hidden: true,  size: '150px', sortable: true  },
            {field: 'Amount',   caption: 'Amount',       hidden: false, size: '100px', sortable: true, style: 'text-align: right', render: 'money'},
            {field: 'BID',      caption: 'BUD',          hidden: true,  size: '40px',  sortable: false },
            {field: 'TCID',     caption: 'TCID',         hidden: true,  size: '40px',  sortable: false },
            {field: 'PMTID',    caption: 'PMTID',        hidden: true,                 sortable: false },
            {field: 'PMTName',  caption: 'Payment Type', hidden: false, size: '100px', sortable: true, style: 'text-align: center' },
            {field: 'DocNo',    caption: 'Document No.', hidden: false, size: '100px', sortable: true, style: 'text-align: right'},
            {field: 'Payors',   caption: 'Payors',       hidden: false, size: '200px', sortable: true  },
        ],
    });    

    //-------------------------------------------------------------------------------
    //  depositLayout - The layout to contain the depositForm and depositDetailGrid
    //-------------------------------------------------------------------------------
    $().w2layout({
        name: 'depositLayout',
        padding: 0,
        panels: [
            { type: 'left',    size: '30%', hidden: true },
            { type: 'top',     size: 375,   hidden: false, content: 'top',  resizable: true, style: app.pstyle },
            { type: 'main',    size: '70%', hidden: false, content: 'main', resizable: true, style: app.pstyle },
            { type: 'preview', size: 0,     hidden: true,  content: 'PREVIEW'  },
            { type: 'bottom',  size: 0,     hidden: true },
            { type: 'right',   size: 0,     hidden: true }
        ]
    });
}

//-----------------------------------------------------------------------------
// createDepositForm - add the grid and form to the statement layout.  I'm not
//      sure why this is necessary. But if I put this grid and form directly
//      into the layout when it gets created, they do not work correctly.
// @params
//-----------------------------------------------------------------------------
function createDepositForm() {
    w2ui.depositLayout.content('top',w2ui.depositForm);
    w2ui.depositLayout.content('main',w2ui.depositListGrid);
}

//-----------------------------------------------------------------------------
// setToDepositForm - set to the Deposit Form - puts the depositLayout in
//                    toplayout's right content area. Didn't use the general
//                    call in rutil.js because this form requires the layout
//                    and has multiple parts.
// @params
//   sform   = name of the form
//   url     = request URL for the form
//   [width] = optional, if specified it is the width of the form
//   doRequest = 
//-----------------------------------------------------------------------------
function setToDepositForm(slayout, sform, url, urlgrid, width, doRequest) {
    // if not url defined then return
    var url_len=url.length > 0;
    if (!url_len) {
        return false;
    }

    // if form not found then return
    var f = w2ui[sform];
    if (!f) {
        return false;
    }

    // if current grid not found then return
    var g = w2ui[app.active_grid];
    if (!g) {
        return false;
    }

    // if doRequest is defined then take false as default one
    if (!doRequest) {
        doRequest = false;
    }
    f.url = url;
    if (typeof f.tabs.name == "string") {
        f.tabs.click('tab1');
    }
    app.new_form_rec = !doRequest;
    app.form_is_dirty = false;

    var right_panel_content = w2ui.toplayout.get("right").content;
    var fc = w2ui[slayout]; // in this case, we're putting the layout into the content area
    w2ui.depositListGrid.url = urlgrid;
    var showForm = function() {
        // if the same content is there, then no need to render toplayout again
        if ( fc !== right_panel_content) {
            w2ui.toplayout.content('right', fc);
            w2ui.toplayout.sizeTo('right', width);
            w2ui.toplayout.render();
        } else {
            fc.refresh();
        }
        $().w2tag();
        w2ui.toplayout.show('right', true);
    };

    if (doRequest) {
        f.request(function(event) {
            if (event.status === "success") {
                showForm();
                return true;
            }
            else {
                showForm();
                f.message("Could not get form data from server...!!");
                return false;
            }
        });
    } else {
        var sel_recid = parseInt(g.last.sel_recid);
        if (sel_recid > -1) {
            g.unselect(g.last.sel_recid); // if new record is being added then unselect {{the selected record}} from the grid
        }
        showForm();
        return true;
    }
}
