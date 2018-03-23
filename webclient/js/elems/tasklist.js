"use strict";
/*global
    GridMoneyFormat, number_format, w2ui, $, app, console,setToStmtForm,
    form_dirty_alert, addDateNavToToolbar
*/

function buildTaskListElements() {
    //------------------------------------------------------------------------
    //          tlsGrid  -  TASK LISTS in the date range
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'tlsGrid',
        url: '/v1/tls',
        multiSelect: false,
        postData: {searchDtStart: app.D1, searchDtStop: app.D2},
        show: {
            toolbar         : true,
            footer          : true,
            toolbarAdd      : true,    // indicates if toolbar add new button is visible
            toolbarDelete   : false,   // indicates if toolbar delete button is visible
            toolbarSave     : false,   // indicates if toolbar save button is visible
            selectColumn    : false,
            expandColumn    : false,
            toolbarEdit     : false,
            toolbarSearch   : false,
            toolbarInput    : false,
            searchAll       : false,
            toolbarReload   : true,
            toolbarColumns  : true,
        },
        columns: [
            {field: 'recid',     hidden: true,  caption: 'recid',             size: '40px',  sortable: true},
            {field: 'BID',       hidden: true,  caption: 'BID',               size: '40px',  sortable: true},
            {field: 'Name',      hidden: false, caption: 'Name',              size: '110px', sortable: true},
            {field: 'DtDue',     hidden: false, caption: 'Due',               size: '110px', sortable: true},
            {field: 'DtDone',    hidden: false, caption: 'Due completed',     size: '110px', sortable: true},
            {field: 'DtPreDue',  hidden: false, caption: 'Pre Due',           size: '110px', sortable: true},
            {field: 'DtPreDone', hidden: false, caption: 'Pre Due completed', size: '110px', sortable: true},
        ],
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
                        console.log( 'BID = ' + rec.BID + ',   TLID = ' + rec.TLID);
                        setToTLForm(rec.BID, rec.TLID, app.D1, app.D2);
                    };

                // warn user if form content has been changed
                form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
            };
        },
    });

    addDateNavToToolbar('tls'); // "Grid" is appended to the 

    //------------------------------------------------------------------------
    //  tlsInfoForm
    //------------------------------------------------------------------------
    $().w2form({
        name: 'tlsInfoForm',
        style: 'border: 0px; background-color: transparent;',
        header: 'Task List',
        url: '/v1/tl',
        formURL: '/webclient/html/formtl.html',
        toolbar: {
            items: [
                { id: 'btnNotes', type: 'button', icon: 'fa fa-sticky-note-o' },
                { id: 'btnClose', type: 'button', icon: 'fa fa-times' },
            ],
            onClick: function (event) {
                event.onComplete = function() {
                    var g = w2ui.tlsDetailGrid;
                    var r = w2ui.tlsInfoForm.record;
                    var d1, d2, url;
                    switch(event.target) {
                    case 'btnClose':
                        var no_callBack = function() { return false; },
                            yes_callBack = function() {
                                w2ui.toplayout.hide('right',true);
                                w2ui.payorstmtGrid.render();
                            };
                        form_dirty_alert(yes_callBack, no_callBack);
                        break;
                    // case 'payorstmtint':
                    //     app.PayorStmtExt = false;
                    //     g.postData.Bool1 = false;
                    //     g.url = '/v1/payorstmt/' + r.BID + '/' + r.TCID;
                    //     g.reload();
                    //     break;
                    // case 'payorstmtext':
                    //     app.PayorStmtExt = true;
                    //     g.postData.Bool1 = true;
                    //     g.url = '/v1/payorstmt/' + r.BID + '/' + r.TCID;
                    //     g.reload();
                    //     break;
                    // case 'csvexport':
                    //     d1 = document.getElementsByName("tlsDetailD1")[0].value;
                    //     d2 = document.getElementsByName("tlsDetailD2")[0].value;
                    //     url = exportReportCSV("RPTpayorstmt", d1, d2, true);
                    //     url += "&internal=" + !app.PayorStmtExt;
                    //     url += "&tcid=" + r.TCID;
                    //     downloadMediaFromURL(url);
                    //     break;
                    // case 'pdfexport':
                    //     d1 = document.getElementsByName("tlsDetailD1")[0].value;
                    //     d2 = document.getElementsByName("tlsDetailD2")[0].value;
                    //     url = exportReportPDF("RPTpayorstmt", d1, d2, true);
                    //     url += "&internal=" + !app.PayorStmtExt;
                    //     url += "&tcid=" + r.TCID;
                    //     downloadMediaFromURL(url);
                    //     break;
                    }
                };
            },
        },
        fields: [
            { field: 'recid',      type: 'int',  required: false },
            { field: 'TLID',       type: 'int',  required: false },
            { field: 'BID',        type: 'int',  required: false },
            { field: 'Name',       type: 'text', required: false },
            { field: 'Cycle',      type: 'int',  required: false },
            { field: 'DtDue',      type: 'date', required: false },
            { field: 'DtPreDue',   type: 'date', required: false },
            { field: 'DtDone',     type: 'date', required: false },
            { field: 'DtPreDone',  type: 'date', required: false },
            { field: 'FLAGS',      type: 'int',  required: false },
            { field: 'DoneUID',    type: 'int',  required: false },
            { field: 'PreDoneUID', type: 'int',  required: false },
            { field: 'Comment',    type: 'text', required: false },
            { field: 'CreateTS',   type: 'date', required: false },
            { field: 'CreateBy',   type: 'int',  required: false },
            { field: 'LastModTime',type: 'date', required: false },
            { field: 'LastModBy',  type: 'int',  required: false },
        ],
        onRefresh: function(event) {
            // var f = this;
            event.onComplete = function() {
                // var r = f.record;
                // var x = document.getElementById("bannerTCID");
                // if (x !== null) {
                //     var title;
                //     if (r.PayorIsCompany) {
                //         title = r.CompanyName;
                //     } else {
                //         title = r.FirstName + ' ';
                //         if (typeof r.MiddleName == "string") {
                //             if (r.MiddleName.length > 0 ) {
                //                 title += r.MiddleName + ' ';
                //             }
                //         }
                //         title += r.LastName + ' ';
                //     }
                //     x.innerHTML = title;
                // }
                // x = document.getElementById("payorstmtaddr");
                // if (x !== null) {
                //     x.innerHTML = '' + r.Address;
                // }
                // x = document.getElementById("RentalAgreementDates");
                // if (x !== null) {
                //     x.innerHTML = '' + this.record.AgreementStart + ' - ' + this.record.AgreementStop;
                // }
                // x = document.getElementById("PossessionDates");
                // if (x !== null) {
                //     x.innerHTML = '' + this.record.PossessionStart + ' - ' + this.record.PossessionStop;
                // }
                // x = document.getElementById("RentDates");
                // if (x !== null) {
                //     x.innerHTML = '' + this.record.RentStart + ' - ' + this.record.RentStop;
                // }
                // x = document.getElementById("CurrentStatementBalance");
                // if (x !== null) {
                //     x.innerHTML = '$ ' + number_format(this.record.Balance ,2);
                // }
                // x = document.getElementById("payorunalloc");
                // if (x !== null) {
                //     x.innerHTML = '' + this.record.PayorUnalloc;
                // }
            };
        },

    });

    // addDateNavToToolbar('tlsInfoForm');

    //------------------------------------------------------------------------
    //  tlsDetailGrid  -  lists all the assessments and receipts for
    //                     the selected Rental Agreement from the stmtGrid
    //------------------------------------------------------------------------
    // $().w2grid({
    //     name: 'tlsDetailGrid',
    //     url: '/v1/payorstmt',
    //     multiSelect: false,
    //     postData: {searchDtStart: app.D1, searchDtStop: app.D2, Bool1: app.PayorStmtExt},
    //     show: {
    //         toolbar         : true,
    //         footer          : true,
    //         toolbarAdd      : false,   // indicates if toolbar add new button is visible
    //         toolbarDelete   : false,   // indicates if toolbar delete button is visible
    //         toolbarSave     : false,   // indicates if toolbar save button is visible
    //         selectColumn    : false,
    //         expandColumn    : false,
    //         toolbarEdit     : false,
    //         toolbarSearch   : false,
    //         toolbarInput    : false,   // the text area for searches
    //         searchAll       : false,
    //         toolbarReload   : false,
    //         toolbarColumns  : false,
    //     },
    //     columns: [
    //         {field: 'recid',           caption: 'recid',           size: '35px',  sortable: true, hidden: true},
    //         {field: 'Date',            caption: 'Date',            size: '70px',  sortable: true, render: function(rec) {return renderPayorStmtDate(rec.Date); }},
    //         {field: 'Reverse',         caption: ' ',               size: '12px',  sortable: true, render: renderPayorStmtReversal },
    //         {field: 'Payor',           caption: 'Payor',           size: '100px', sortable: true},
    //         {field: 'TCID',            caption: 'TCID',            size: '80px',  sortable: true, hidden: true},
    //         {field: 'RAID',            caption: 'RAID',            size: '45px',  sortable: true },
    //         {field: 'ASMID',           caption: 'ASMID',           size: '55px',  sortable: true },
    //         {field: 'RCPTID',          caption: 'RCPTID',          size: '60px',  sortable: true },
    //         {field: 'RentableName',    caption: app.sRentable,     size: '30%',   sortable: true},
    //         {field: 'Description',     caption: 'Description',     size: '60%',   sortable: true},
    //         {field: 'UnappliedAmount', caption: 'Unapplied Funds', size: '90px',  sortable: true, style: 'text-align: right',
    //                 render: function (record,index,col_index) { return payorstmtRenderHandler(record,index,col_index,record.UnappliedAmount,true); },
    //         },
    //         {field: 'AppliedAmount',   caption: 'Applied Funds',   size: '95px', sortable: true, style: 'text-align: right',
    //                 render: function (record,index,col_index) { return payorstmtRenderHandler(record,index,col_index,record.AppliedAmount,true); },
    //         },
    //         {field: 'Assessment',      caption: 'Assessment',      size: '90px',  sortable: true, style: 'text-align: right',
    //                 render: function (record,index,col_index) { return payorstmtRenderHandler(record,index,col_index,record.Assessment,true); },
    //         },
    //         {field: 'Balance',         caption: 'Balance',         size: '90px', sortable: true, style: 'text-align: right',
    //                 render: function (record,index,col_index) { return payorstmtRenderHandler(record,index,col_index,record.Balance,false); },
    //         },
    //         {field: 'spacer',          caption: ' ',               size: '7px'},

    //     ],
    // });


    //------------------------------------------------------------------------
    //  payorstmtlayout - The layout to contain the stmtForm and tlsDetailGrid
    //               top  - stmtForm
    //               main - tlsDetailGrid
    //------------------------------------------------------------------------
    $().w2layout({
        name: 'tlLayout',
        padding: 0,
        panels: [
            { type: 'left',    size: 0,     hidden: true },
            { type: 'top',     size: '40%', hidden: false, content: 'top',  resizable: true, style: app.pstyle },
            { type: 'main',    size: '70%', hidden: false, content: 'main', resizable: true, style: app.pstyle },
            { type: 'preview', size: 0,     hidden: true,  content: 'PREVIEW'  },
            { type: 'bottom',  size: 0,     hidden: true },
            { type: 'right',   size: 0,     hidden: true }
        ]
    });
}

function finishTaskListForm() {
    w2ui.tlLayout.content('top',   w2ui.tlsInfoForm);
    // w2ui.tlLayout.content('main',  w2ui.depositListGrid);
    // w2ui.tlLayout.content('bottom',w2ui.depositFormBtns);
}

//-----------------------------------------------------------------------------
// setToTLForm -  enable the Statement form in toplayout.  Also, set
//                the forms url and request data from the server
// @params
//   bid = business id (or the BUD)
//    id = Task List TLID
// d1,d2 = date range to use
//-----------------------------------------------------------------------------
function setToTLForm(bid, id, d1,d2) {
    if (id > 0) {
        w2ui.tlsGrid.url = '/v1/tl/' + bid + '/' + id;
        w2ui.tlsInfoForm.url = '/v1/tl/' + bid + '/' + id;
        w2ui.tlsInfoForm.postData = {
            searchDtStart: d1,
            searchDtStop: d2,
        };
        w2ui.tlsInfoForm.header = 'Task List ' + id;
        w2ui.tlsInfoForm.request();

        w2ui.toplayout.content('right', w2ui.tlLayout);
        w2ui.toplayout.show('right', true);
        w2ui.toplayout.sizeTo('right', 500);
        w2ui.toplayout.render();
        app.new_form_rec = false;  // mark as record exists
        app.form_is_dirty = false; // mark as no changes yet
    }
}
