/*global
    GridMoneyFormat, number_format, w2ui, $, app, console,
    form_dirty_alert, addDateNavToToolbar, renderPayorStmtReversal, payorstmtRenderHandler,
    dateFromString, exportReportCSV, exportReportPDF,getGridReversalSymbolHTML, setToPayorStmtForm, renderPayorStmtDate
*/
"use strict";

window.buildPayorStatementElements = function () {
    //------------------------------------------------------------------------
    //  payorstmt  -  lists all the assessments and receipts for
    //                     the selected Payors
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'payorstmtGrid',
        url: '/v1/payorstmt',
        multiSelect: false,
        postData: {searchDtStart: app.D1, searchDtStop: app.D2},
        show: {
            toolbar         : true,
            footer          : true,
            toolbarAdd      : false,   // indicates if toolbar add new button is visible
            toolbarDelete   : false,   // indicates if toolbar delete button is visible
            toolbarSave     : false,   // indicates if toolbar save button is visible
            selectColumn    : false,
            expandColumn    : false,
            toolbarEdit     : false,
            toolbarSearch   : false,
            toolbarInput    : true,    // the text area for searches
            searchAll       : false,
            toolbarReload   : false,
            toolbarColumns  : false,
        },
        columns: [
            {field: 'IsSubTotalRow', caption: 'Is SubTotal Row',                sortable: false, hidden: true},
            {field: 'recid',         caption: 'recid',           size: '35px',  sortable: true,  hidden: true},
            {field: 'TCID',          caption: 'TCID',            size: '75px',  sortable: true},
            {field: 'FirstName',     caption: 'FirstName',       size: '150px', sortable: true},
            {field: 'LastName',      caption: 'LastName',        size: '150px', sortable: true},
            {field: 'CompanyName',   caption: 'CompanyName',     size: '200px', sortable: true},
            {field: 'IsCompany',     caption: 'IsCompany',       size: '200px', sortable: true,  hidden: true },
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
                        console.log( 'BID = ' + rec.BID + ',   RAID = ' + rec.RAID);
                        setToPayorStmtForm(rec.BID, rec.TCID, app.D1, app.D2);
                    };

                // warn user if form content has been changed
                form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
            };
        },
    });

    addDateNavToToolbar('payorstmt');

    //------------------------------------------------------------------------
    //  payorStmtInfoForm
    //------------------------------------------------------------------------
    $().w2form({
        name: 'payorStmtInfoForm',
        style: 'border: 0px; background-color: transparent;',
        header: 'Payor Statement',
        url: '/v1/payorstmtinfo',
        formURL: '/webclient/html/formpayorstmtdet.html',
        toolbar: {
            items: [
                { id: 'btnNotes',     type: 'button', icon: 'far fa-sticky-note' },
                {                     type: 'break' },
                { id: 'payorstmtint', type: 'radio', group: '1', text: 'Internal', icon: 'far fa-file-alt', checked: true },
                { id: 'payorstmtext', type: 'radio', group: '1', text: 'External', icon: 'far fa-file' },
                {                     type: 'break' },
                { id: 'csvexport',    type: 'button', icon: 'fas fa-table', tooltip: 'export to CSV' },
                { id: 'pdfexport',    type: 'button', icon: 'far fa-file-pdf', tooltip: 'export to PDF' },
                {                     type: 'spacer' },
                { id: 'btnClose',     type: 'button', icon: 'fas fa-times' },
            ],
            onClick: function (event) {
                event.onComplete = function() {
                    var g = w2ui.payorStmtDetailGrid;
                    var r = w2ui.payorStmtInfoForm.record;
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
                    case 'payorstmtint':
                        app.PayorStmtExt = false;
                        g.postData.Bool1 = false;
                        g.url = '/v1/payorstmt/' + r.BID + '/' + r.TCID;
                        g.reload();
                        break;
                    case 'payorstmtext':
                        app.PayorStmtExt = true;
                        g.postData.Bool1 = true;
                        g.url = '/v1/payorstmt/' + r.BID + '/' + r.TCID;
                        g.reload();
                        break;
                    case 'csvexport':
                        d1 = document.getElementsByName("payorStmtDetailD1")[0].value;
                        d2 = document.getElementsByName("payorStmtDetailD2")[0].value;
                        url = exportReportCSV("RPTpayorstmt", d1, d2, true);
                        url += "&internal=" + !app.PayorStmtExt;
                        url += "&tcid=" + r.TCID;
                        downloadMediaFromURL(url);
                        break;
                    case 'pdfexport':
                        d1 = document.getElementsByName("payorStmtDetailD1")[0].value;
                        d2 = document.getElementsByName("payorStmtDetailD2")[0].value;
                        url = exportReportPDF("RPTpayorstmt", d1, d2, true);
                        url += "&internal=" + !app.PayorStmtExt;
                        url += "&tcid=" + r.TCID;
                        downloadMediaFromURL(url);
                        break;
                    }
                };
            },
        },
        fields: [
            { field: 'recid', type: 'int', required: false, html: {page: 0, column: 0 } },
            { field: 'RAID', type: 'int', required: false, html: {  page: 0, column: 0 } },
            { field: 'TCID', type: 'int', required: false, html: {  page: 0, column: 0 } },
            { field: 'BID', type: 'int', required: false, html: { page: 0, column: 0 } },
            { field: 'FirstName', type: 'text', required: false, html: { page: 0, column: 0 } },
            { field: 'MiddleName', type: 'text', required: false, html: { page: 0, column: 0 } },
            { field: 'LastName', type: 'text', required: false, html: { page: 0, column: 0 } },
            { field: 'PayorIsCompany', type: 'checkbox', required: false, html: { page: 0, column: 0 } },
            { field: 'CompanyName', type: 'text', required: false, html: { page: 0, column: 0 } },
            { field: 'Address', type: 'text', required: false, html: { page: 0, column: 0 } },
        ],
        onRefresh: function(event) {
            var f = this;
            event.onComplete = function() {
                var r = f.record;
                var x = document.getElementById("bannerTCID");
                if (x !== null) {
                    var title;
                    if (r.PayorIsCompany) {
                        title = r.CompanyName;
                    } else {
                        title = r.FirstName + ' ';
                        if (typeof r.MiddleName == "string") {
                            if (r.MiddleName.length > 0 ) {
                                title += r.MiddleName + ' ';
                            }
                        }
                        title += r.LastName + ' ';
                    }
                    x.innerHTML = title;
                }
                x = document.getElementById("payorstmtaddr");
                if (x !== null) {
                    x.innerHTML = '' + r.Address;
                }
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

    //------------------------------------------------------------------------
    //  payorStmtDetailGrid  -  lists all the assessments and receipts for
    //                     the selected Rental Agreement from the stmtGrid
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'payorStmtDetailGrid',
        url: '/v1/payorstmt',
        multiSelect: false,
        postData: {searchDtStart: app.D1, searchDtStop: app.D2, Bool1: app.PayorStmtExt},
        show: {
            toolbar         : true,
            footer          : true,
            toolbarAdd      : false,   // indicates if toolbar add new button is visible
            toolbarDelete   : false,   // indicates if toolbar delete button is visible
            toolbarSave     : false,   // indicates if toolbar save button is visible
            selectColumn    : false,
            expandColumn    : false,
            toolbarEdit     : false,
            toolbarSearch   : false,
            toolbarInput    : false,   // the text area for searches
            searchAll       : false,
            toolbarReload   : false,
            toolbarColumns  : false,
        },
        columns: [
            {field: 'recid',           caption: 'recid',           size: '35px',  sortable: true, hidden: true},
            {field: 'Date',            caption: 'Date',            size: '70px',  sortable: true, render: function(rec) {return renderPayorStmtDate(rec.Date); }},
            {field: 'Reverse',         caption: ' ',               size: '12px',  sortable: true, render: renderPayorStmtReversal },
            {field: 'Payor',           caption: 'Payor',           size: '100px', sortable: true},
            {field: 'TCID',            caption: 'TCID',            size: '80px',  sortable: true, hidden: true},
            {field: 'RAID',            caption: 'RAID',            size: '45px',  sortable: true },
            {field: 'ASMID',           caption: 'ASMID',           size: '55px',  sortable: true },
            {field: 'RCPTID',          caption: 'RCPTID',          size: '60px',  sortable: true },
            {field: 'RentableName',    caption: app.sRentable,     size: '30%',   sortable: true},
            {field: 'Description',     caption: 'Description',     size: '60%',   sortable: true},
            {field: 'UnappliedAmount', caption: 'Unapplied Funds', size: '90px',  sortable: true, style: 'text-align: right',
                    render: function (record,index,col_index) { return payorstmtRenderHandler(record,index,col_index,record.UnappliedAmount,true); },
            },
            {field: 'AppliedAmount',   caption: 'Applied Funds',   size: '95px', sortable: true, style: 'text-align: right',
                    render: function (record,index,col_index) { return payorstmtRenderHandler(record,index,col_index,record.AppliedAmount,true); },
            },
            {field: 'Assessment',      caption: 'Assessment',      size: '90px',  sortable: true, style: 'text-align: right',
                    render: function (record,index,col_index) { return payorstmtRenderHandler(record,index,col_index,record.Assessment,true); },
            },
            {field: 'Balance',         caption: 'Balance',         size: '90px', sortable: true, style: 'text-align: right',
                    render: function (record,index,col_index) { return payorstmtRenderHandler(record,index,col_index,record.Balance,false); },
            },
            {field: 'spacer',          caption: ' ',               size: '7px'},

        ],
    });

    addDateNavToToolbar('payorStmtDetail');

    //------------------------------------------------------------------------
    //  payorstmtlayout - The layout to contain the stmtForm and payorStmtDetailGrid
    //               top  - stmtForm
    //               main - payorStmtDetailGrid
    //------------------------------------------------------------------------
    $().w2layout({
        name: 'payorstmtLayout',
        padding: 0,
        panels: [
            { type: 'left',    size: '30%', hidden: true },
            { type: 'top',     size: 170,   hidden: false, content: 'top',  resizable: true, style: app.pstyle },
            { type: 'main',    size: '70%', hidden: false, content: 'main', resizable: true, style: app.pstyle },
            { type: 'preview', size: 0,     hidden: true,  content: 'PREVIEW'  },
            { type: 'bottom',  size: 0,     hidden: true },
            { type: 'right',   size: 0,     hidden: true }
        ]
    });
};

window.payorstmtRenderHandler = function (record,index,col_index,amt,bRemoveZero) {
    var f = w2ui.payorStmtDetailGrid.columns[col_index];
    if (record.Reverse && f.field == "Balance") { return; }  // don't update balance if it's a reversal
    if (record.Description.includes("***") || record.Description.length == 0) {return;} // blank if it's a header or spacer
    if (Math.abs(amt) < 0.001) {
        if (record.Description.includes("Closing Balance") || !bRemoveZero) {
            return '$ 0.00';
        }
    }
    return GridMoneyFormat(amt);
};

window.renderPayorStmtReversal = function (record /*, index, col_index*/) {
    if (typeof record === "undefined") {
        return;
    }
    if ( record.Reverse ) { // if reversed then
        return getGridReversalSymbolHTML();
    }
    return '';
};

window.renderPayorStmtDate = function (s) {
    //var d = new Date(y);
    var d = dateFromString(s);
    if (d.getFullYear() < 1971) {
        return '';
    }
    return dateFmtStr(d);
};

//-----------------------------------------------------------------------------
// renderPayorStmtID - render the ID number for RAID, ASMID, and RCPTID.
//        If the ID is > 0 return the number, otherwise just return an
//        empty string.
// @params
//    record = current record being rendered
//     index = index within the record array
// col_index = column index within the record
//
// @returns
//      an empty string if the id is 0
//      the number if ID >= 1
//-----------------------------------------------------------------------------
window.renderPayorStmtID = function (record, index, col_index) {
    var f = w2ui.payorStmtDetailGrid.columns[col_index];
    var n = 0;
    switch ( f.field ) {
        case "RAID":   n = record.RAID; break;
        case "ASMID":  n = record.ASMID; break;
        case "RCPTID": n = record.RCPTID; break;
        default:
            return '';
    }
    if (n > 0) {
        return ''+n;
    }
    return '';
};

//-----------------------------------------------------------------------------
// setToPayorStmtForm -  enable the Statement form in toplayout.  Also, set
//                the forms url and request data from the server
// @params
//   bid = business id (or the BUD)
//  tcid = Payor's TCID
// d1,d2 = date range to use
//-----------------------------------------------------------------------------
window.setToPayorStmtForm = function (bid, tcid, d1,d2) {
    if (tcid > 0) {
        w2ui.payorStmtDetailGrid.url = '/v1/payorstmt/' + bid + '/' + tcid;
        w2ui.payorStmtInfoForm.url = '/v1/payorstmtinfo/' + bid + '/' + tcid;
        w2ui.payorStmtInfoForm.postData = {
            searchDtStart: d1,
            searchDtStop: d2,
            Bool1: app.PayorStmtExt,
        };

        // ==================
        // INTERNAL FUNCTION
        // ==================
        var showForm = function() {
            w2ui.toplayout.content('right', w2ui.payorstmtLayout);
            w2ui.toplayout.show('right', true);
            w2ui.toplayout.sizeTo('right', 1000);
            w2ui.toplayout.render();
            app.new_form_rec = false;  // mark as record exists
            app.form_is_dirty = false; // mark as no changes yet
            // NOTE: remove any error tags bound to field from previous form
            $().w2tag();
            // SHOW the right panel now
            w2ui.toplayout.show('right', true);
        };

        w2ui.payorStmtInfoForm.header = 'Payor Statement for TCID ' + tcid;
        w2ui.payorStmtInfoForm.request(function(event) {
            if (event.status === "success") {
                showForm();
                return true;
            } else {
                showForm();
                w2ui.payorStmtInfoForm.message("Could not get form data from server...!!");
                return false;
            }
        });
    }
};

//-----------------------------------------------------------------------------
// createPayorStmtForm - add the grid and form to the statement layout.  I'm not
//      sure why this is necessary. But if I put this grid and form directly
//      into the layout when it gets created, they do not work correctly.
// @params
//-----------------------------------------------------------------------------
window.createPayorStmtForm = function () {
    w2ui.payorstmtLayout.content('top',w2ui.payorStmtInfoForm);
    w2ui.payorstmtLayout.content('main',w2ui.payorStmtDetailGrid);
};
