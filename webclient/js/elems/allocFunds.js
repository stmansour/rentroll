/*global
    $,w2ui,app,console,refreshUnallocAmtSummaries,unallocAmountRemaining,_unAllocRcpts,
    parseFloat,getPayorFund,number_format, computeAmountRemaining

*/
"use strict";
window.buildAllocFundsGrid = function () {
    //------------------------------------------------------------------------
    //          allocfundsGrid
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'allocfundsGrid',
        url: '/v1/allocfunds',
        multiSelect: false,
        title: 'Unpaid assesments for ',
        show: {
            header: false,
            toolbar: true,
            footer: true,
            searches: true,
            lineNumbers: false,
            selectColumn: false,
            expandColumn: false
        },
        columns: [
            {field: 'recid', type: 'int', caption: 'recid', hidden: true, size: '40px', sortable: true},
            {field: 'Name', type: 'text', caption: 'Payor Name', size: '90%', sortable: true},
            {field: 'BID', type: 'int', caption: 'BID', hidden: true},
            {field: 'TCID', type: 'int', caption: 'TCID', hidden: true},
        ],
        onRefresh: function(event) {
            event.onComplete = function() {
                var sel_recid = parseInt(this.last.sel_recid);
                if (app.active_grid == this.name && sel_recid > -1) {
                    if (app.new_form_rec) {
                        this.selectNone();
                    }
                    else{
                        this.select(sel_recid);
                    }
                }
            };
        },
        onClick:  function(event) {
            event.onComplete = function() {
                var rec = this.get(event.recid);

                app.TmpTCID = rec.TCID; // store here in case it is deselected by the time the Save function needs it
                getPayorFund(rec.BID, rec.TCID)
                .done(function(data) {
                    app.payor_fund = data.record.fund;  // store fund in app variable
                    w2ui.unpaidASMsGrid.load('/v1/unpaidasms/'+rec.BID+'/'+rec.TCID);   // load grid data

                    var top = _unAllocRcpts.layoutPanels.top(app.payor_fund, rec.Name/*, rec.TCID*/),
                        bottom = _unAllocRcpts.layoutPanels.bottom();

                    w2ui.allocfundsLayout.content('top', top);
                    w2ui.allocfundsLayout.content('main', w2ui.unpaidASMsGrid);
                    w2ui.allocfundsLayout.content('bottom', bottom);

                    w2ui.toplayout.show('right', true);
                    w2ui.toplayout.content('right', w2ui.allocfundsLayout);
                    w2ui.toplayout.sizeTo('right', 800);
                    w2ui.toplayout.render();
                })
                .fail(function(/*data*/) {
                    console.log('ERROR, unable to get payor fund for TCID', rec.TCID);
                });
            };
        },
    });

    //------------------------------------------------------------------------
    //          allocate fund layout
    //------------------------------------------------------------------------
    $().w2layout({
        name: 'allocfundsLayout',
        panels: [
            { type: "top", size: 140, style: 'border: 1px solid #cfcfcf; padding: 5px;', content: 'Allocate Fund Top Panel',
                toolbar: {
                    name: 'unallocfund_toolbar',
                    items: [
                        { id: 'btnNotes', type: 'button', icon: 'far fa-sticky-note' },
                        { id: 'bt3', type: 'spacer' },
                        { id: 'btnClose', type: 'button', icon: 'fas fa-times' },
                    ],
                    onClick: function(event) {
                        switch(event.target) {
                            case 'btnClose':
                                w2ui.toplayout.hide('right',true);
                                w2ui.allocfundsGrid.render();
                                break;
                        }
                    }
                }
            },
            { type: "left", hidden: true },
            { type: "main", size: "90%", style: app.pstyle, content: 'Allocate Fund Main Panel' },
            { type: "right", hidden: true },
            { type: "bottom", size: 60, style: app.pstyle, content: 'Allocate Fund Bottom Panel' },
        ]
    });


    //------------------------------------------------------------------------------
    //   UNPAID ASSESSMENTS GRID
    //------------------------------------------------------------------------------

    $().w2grid({
        name: 'unpaidASMsGrid',
        header: 'Unpaid Assessments',
        show: {
            toolbar: false,
            header: true,
            footer: false,
            toolbarReload   : false,
            toolbarColumns  : false,
            toolbarSearch   : false,
            toolbarAdd: false,
            toolbarDelete: false,
            toolbarInput: false,
            searchAll: false,
            toolbarSave: false,
            toolbarEdit: false,
            searches: false,
            lineNumbers: false,
            selectColumn: false,
            expandColumn: false
        },
        columns: [
            { field: 'recid', type: 'int', caption: 'recid', size: '40px', hidden: true},
            { field: 'Date', render: 'date', caption: 'Assessment<br>Date', style: 'text-align: right', size: '80px' },
            { field: 'Assessment', caption: 'Assessment', size: '10%' },
            { field: 'Amount', render: 'money', caption: 'Assessment<br>Amount', size: '100px' },
            { field: 'AmountPaid', render: 'money', caption: 'Amount Paid', size: '100px' },
            { field: 'AmountOwed', render: 'money', caption: 'Amount Owed', size: '100px' },
            { field: 'ARID', hidden: true },
            { field: 'ASMID', hidden: true },
            // INDEX 8 is Dt. If this changes, then the id reference to it must change in the onChange handler below
            { field: 'Dt', render: 'date', caption: 'Payment Date', size: '100px', style: 'text-align: right', editable: {type: 'date'}},
            { field: 'Allocate', render: 'money', caption: 'Allocate Amount', size: '110px', editable: {type: 'float'}},
        ],
        onRefresh: function(/*event*/) {
            unallocAmountRemaining();
            refreshUnallocAmtSummaries();
        },
        onLoad: function(event) {
            event.done(function () {
                if (w2ui.unpaidASMsGrid.summary.length === 0) {
                    w2ui.unpaidASMsGrid.add({recid: 's-1', Date: '', Assessment: '', Amount: 0, AmountPaid: 0, AmountOwed: 0, Allocate: 0, w2ui: {summary: true}});
                }
            });
        },
        onRender: function(event) {
            event.done(function() {
                refreshUnallocAmtSummaries();
                unallocAmountRemaining();
            });
        },
        onChange: function(event) {
            event.onComplete = function() {
            	var i;
                //----------------------------------------------------------------
                // if they just changed the date, take the change and return...
                //----------------------------------------------------------------
                if (event.column == 8) { // if the date changes, don't mess up anything else...
                    for ( i = 0; i < w2ui.unpaidASMsGrid.records.length; i++) {
                        if ( event.recid == w2ui.unpaidASMsGrid.records[i].recid ) {
                            w2ui.unpaidASMsGrid.records[i].Dt = event.value_new;
                            return;
                        }
                    }
                    return;
                }
                var tgrid = w2ui.unpaidASMsGrid;
                var c = w2ui[event.target].getChanges();
                var total_fund = parseFloat($("#total_fund_amount").attr("data-fund")).toFixed(2);
                for ( i = 0; i < c.length; i++) {
                    var rec_index = tgrid.get(c[i].recid, true);         // get record index in `records` array
                    if (c[i].Allocate > 0) {                             // Is an amount present?
                        var fundsAllocated = 0;
                        var fundsRemaining = 0;
                        for (var j=0; j < tgrid.records.length; j++) {   // how much has been allocated everywhere else?
                            if (rec_index != tgrid.records[j].recid) {
                                fundsAllocated += tgrid.records[j].Allocate;
                            }
                        }
                        fundsRemaining = total_fund - fundsAllocated;    // how much remains?
                        var owed = tgrid.records[rec_index].AmountOwed;  // here's how much is still owed
                        var amtToPay = c[i].Allocate;                    // start by paying what the user asked
                        if (amtToPay > fundsRemaining) {
                            amtToPay = fundsRemaining;                   // adjust if it's more than what is available
                        }
                        if (amtToPay > owed) {
                            amtToPay = owed;                             // adjust if it's more than what is owed
                        }
                        tgrid.records[rec_index].Allocate = amtToPay;
                    } else {                                             // if Allocated amount is not greater than zero
                        tgrid.records[rec_index].Allocate = 0;
                    }
                    tgrid.records[rec_index].w2ui = {};                  // clear w2ui object in record
                }
                tgrid.save();
                refreshUnallocAmtSummaries();
                unallocAmountRemaining();
            };
        }
    });

};

//-----------------------------------------------------------------------------
// computeAmountRemaining - based on the amounts allocated to receipts in the
// unpaid receipts list, compute the amount of funds remaining to be allocated
// @params
// @return  the amount of funds remaining; a number
//-----------------------------------------------------------------------------
window.computeAmountRemaining = function () {
    var totalFunds = app.payor_fund; // must already be set to total unallocated receipt funds
    for (var i=0; i < w2ui.unpaidASMsGrid.records.length; i++) {
        totalFunds -= w2ui.unpaidASMsGrid.records[i].Allocate;
    }
    return totalFunds;
};


//-----------------------------------------------------------------------------
// unallocAmountRemaining - display the funds remaining to be allocated
// @params
// @return
//-----------------------------------------------------------------------------
window.unallocAmountRemaining = function () {
    var dispAmt = number_format(computeAmountRemaining(), 2, '.', ',');
    var x = document.getElementById("total_fund_amount");
    if (x !== null) {
        x.innerHTML = dispAmt;
    }
};

//-----------------------------------------------------------------------------
// refreshUnallocAmtSummaries - This routine totals the summary columns for the
// unpaid assessments grid.
// @params
// @return
//-----------------------------------------------------------------------------
window.refreshUnallocAmtSummaries = function () {
    if (w2ui.unpaidASMsGrid.records.length === 0 ) { return; }
    var amt = 0;
    var amtPaid = 0;
    var amtOwed = 0;
    var alloc = 0;
    for (var i=0; i < w2ui.unpaidASMsGrid.records.length; i++) {
        amt += w2ui.unpaidASMsGrid.records[i].Amount;
        amtPaid += w2ui.unpaidASMsGrid.records[i].AmountPaid;
        amtOwed += w2ui.unpaidASMsGrid.records[i].AmountOwed;
        alloc += w2ui.unpaidASMsGrid.records[i].Allocate;
    }
    w2ui.unpaidASMsGrid.set('s-1', {Amount: amt, AmountPaid: amtPaid, AmountOwed: amtOwed, Allocate: alloc});
};

//-----------------------------------------------------------------------------
// getPayorFund - get payor fund
// @params
// @return  the jquery promise
//-----------------------------------------------------------------------------
window.getPayorFund = function (BID, TCID) {
    return jQuery.ajax({
        type: "GET",
        url: '/v1/payorfund/'+BID+'/'+TCID,
        dataType: "json",
    });
};

//-----------------------------------------------------------------------------
// Auto Allocate amount for each unpaid assessment
// returns false
//-----------------------------------------------------------------------------
jQuery(document).on('click', '#auto_allocate_btn', function(/*event*/) {

    // var fund = app.payor_fund;
    var fund = computeAmountRemaining();
    var grid = w2ui.unpaidASMsGrid;

    for (var i = 0; i < grid.records.length; i++) {
        if (fund <= 0) {
            break;
        }

        //--------------------------------------------
        // if it already has an amount, skip it...
        //--------------------------------------------
        if (grid.records[i].Allocate > 0 ) {
            continue;
        }

        // //------------------------------------------------------------------
        // // if this row has been fully paid then move on to the next row
        // //------------------------------------------------------------------
        // if (grid.records[i].Amount - grid.records[i].AmountPaid <= 0) {
        //     continue;
        // }

        // check if fully paid or not
        if (grid.records[i].Amount - grid.records[i].AmountPaid <= fund){
            grid.records[i].Allocate = grid.records[i].Amount - grid.records[i].AmountPaid;
            grid.set(grid.records[i].recid, grid.records[i]);
        } else {
            grid.records[i].Allocate = fund;
            grid.set(grid.records[i].recid, grid.records[i]);
        }

        // decrement fund value by whatever the amount allocated for each record
        fund = fund - grid.records[i].Allocate;
    }
    refreshUnallocAmtSummaries();
    unallocAmountRemaining();
    return false;
});

jQuery(document).on('click', '#alloc_fund_save_btn', function(/*event*/) {

    var tgrid = w2ui.allocfundsGrid;
    var rec = tgrid.getSelection();
    if (rec.length < 0) {
        return;
    }

    // rec = tgrid.get(rec[0]);
    var tcid = app.TmpTCID,
        x = getCurrentBusiness();
    var bid = parseInt(x.value,10);


    var params = {cmd: 'save', TCID: tcid, BID: bid, records: w2ui.unpaidASMsGrid.records };
    var dat = JSON.stringify(params);

    // submit request
    $.post('/v1/allocfunds/'+bid+'/', dat, null, "json")
    .done(function(data) {
        if (data.status != "success") {
            return;
        }
        w2ui.toplayout.hide('right',true);
        w2ui.toplayout.render();
        tgrid.reload();
    })
    .fail(function(/*data*/){
        console.log("Payor Fund Allocation failed.");
    });
});

