/*global
	w2ui,console,app,
*/

"use strict";

function buildRentRollElements() {
    //------------------------------------------------------------------------
    //  rr  -  lists all the assessments and receipts for
    //                     the selected Payors
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'rrGrid',
        url: '/v1/rentroll',
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
            {field: 'recid',            caption: 'recid',                      size: '35px',  sortable: true, hidden: true},
            {field: 'RID',              caption: 'RID',                        size: '75px',  sortable: true, hidden: true},
            {field: 'RentableName',     caption: app.sRentable,                size: '150px', sortable: true},
            {field: 'RTID',             caption: 'RTID',                       size: '75px',  sortable: true, hidden: true},
            {field: 'RTName',           caption: 'RTName',                     size: '150px', sortable: true},
            {field: 'Sqft',             caption: 'Sqft',                       size: '150px', sortable: true},
            {field: 'Description',      caption: 'Description',                size: '150px', sortable: true},
            {field: 'Users',            caption: 'Users',                      size: '150px', sortable: true},
            {field: 'Payors',           caption: 'Payors',                     size: '150px', sortable: true},
            {field: 'RAID',             caption: app.sRentalAgreement,         size: '150px', sortable: true},
            {field: 'Use',              caption: 'Use',                        size: '200px', sortable: true},
            {field: 'Agreement',        caption: 'Agreement',                  size: '200px', sortable: true},
            {field: 'RentCycle',        caption: 'Rent Cycle',                 size: '200px', sortable: true},
            {field: 'GSR',              caption: 'GSR',                        size: '200px', sortable: true, render: 'money'},
            {field: 'PeriodGSR',        caption: 'PeriodGSR',                  size: '200px', sortable: true, render: 'money'},
            {field: 'IncomeOffsets',    caption: 'IncomeOffsets',              size: '200px', sortable: true, render: 'money'},
            {field: 'AmountDue',        caption: 'AmountDue',                  size: '200px', sortable: true, render: 'money'},
            {field: 'PaymentsApplied',  caption: 'Payments Applied',           size: '200px', sortable: true, render: 'money'},
			{field: 'BeginningRcv',	    caption: 'Beginning Receivable',       size: '100px', sortable: false, render: 'money'},
			{field: 'ChangeInRcv',	    caption: 'Change in Receivable',       size: '100px', sortable: false, render: 'money'},
			{field: 'EndingRcv',	    caption: 'Ending Receivable',          size: '100px', sortable: false, render: 'money'},
			{field: 'BeginningSecDep',	caption: 'Beginning Security Deposit', size: '100px', sortable: false, render: 'money'},
			{field: 'ChangeInSecDep',	caption: 'Change in Security Deposit', size: '100px', sortable: false, render: 'money'},
			{field: 'EndingSecDep',	    caption: 'Ending Security Deposit',    size: '100px', sortable: false, render: 'money'},
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
                        //setToForm('rr', '', 800);
                    };

                // warn user if form content has been changed
                form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
            };
        },
    });

    addDateNavToToolbar('payorstmt');
}