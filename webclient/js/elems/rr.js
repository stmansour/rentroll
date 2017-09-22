/*global
    w2ui,console,app,form_dirty_alert,addDateNavToToolbar,
*/

//
// NOTE:  w2ui.grid.recordHeight default height is 24.  Change it to 40 or something to show 2 lines of text in a single cell
//

"use strict";

var grey_fields = [ "BeginningRcv","ChangeInRcv","EndingRcv","BeginningSecDep","ChangeInSecDep","EndingSecDep"];

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
        limit: 20,
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
            {field: 'IsSubTotalRow',     caption: 'Is SubTotal Row',                           sortable: false, hidden: true},
            {field: 'IsBlankRow',        caption: 'Is Blank Row',                              sortable: false, hidden: true},
            {field: 'IsRentableMainRow', caption: 'Is Rentable Main Row',                      sortable: false, hidden: true},
            {field: 'recid',             caption: 'recid',                      size: '35px',  sortable: true,  hidden: true},
            {field: 'BID',               caption: 'BID',                        size: '75px',  sortable: true,  hidden: true},
            {field: 'RID',               caption: 'RID',                        size: '75px',  sortable: true,  hidden: true},
            {field: 'RTID',              caption: 'RTID',                       size: '75px',  sortable: true,  hidden: true},
            {field: 'RentableName',      caption: app.sRentable,                size: '110px', sortable: true},
            {field: 'RentableType',      caption: 'Rentable Type',              size: '100px', sortable: true},
            {field: 'Sqft',              caption: 'Sqft',                       size:  '50px', sortable: true, style: 'text-align: right'},
            {field: 'Description',       caption: 'Description',                size: '150px', sortable: true},
            {field: 'Users',             caption: 'Users',                      size: '150px', sortable: true},
            {field: 'Payors',            caption: 'Payors',                     size: '150px', sortable: true},
            {field: 'RAID',              caption: app.sRentalAgreement,         size: '150px', sortable: true,
                render: function(record/*,index, col_index*/) {
                    if (typeof record === undefined) {
                        return;
                    }
                    if (record.RAID) {
                        return "RA-" + record.RAID; // return ID with "RA-"
                    }
                }
            },
            {field: 'UsePeriod',         caption: 'Use Period',                 size: '85px',  sortable: true, style: 'text-align: right'},
            {field: 'PossessionStart',   caption: 'PossessionStart',            size: '80px',  sortable: true, render: 'date', style: 'text-align: right', hidden: true},
            {field: 'PossessionStop',    caption: 'PossessionStop',             size: '80px',  sortable: true, render: 'date', style: 'text-align: right', hidden: true},
            {field: 'RentPeriod',        caption: 'Rent Period',                size: '85px',  sortable: true, style: 'text-align: right'},
            {field: 'RentStart',         caption: 'RentStart',                  size: '80px',  sortable: true, render: 'date', style: 'text-align: right', hidden: true},
            {field: 'RentStop',          caption: 'RentStop',                   size: '80px',  sortable: true, render: 'date', style: 'text-align: right', hidden: true},
            {field: 'Agreement',         caption: 'Agreement Period',           size: '200px', sortable: true, style: 'text-align: right', hidden: true},
            {field: 'AgreementStart',    caption: 'AgreementStart',             size: '80px',  sortable: true, render: 'date', style: 'text-align: right', hidden: true},
            {field: 'AgreementStop',     caption: 'AgreementStop',              size: '80px',  sortable: true, render: 'date', style: 'text-align: right', hidden: true},
            {field: 'RentCycle',         caption: 'Rent Cycle',                 size: '75px',  sortable: true,
                render: function(record/*, index, col_index*/) {
                    if (typeof record === undefined) {
                        return;
                    }
                    return app.cycleFreq[record.RentCycle];
                }
            },
            {field: 'GSR',               caption: 'GSR',                        size: '200px', sortable: true, render: 'money'},
            {field: 'PeriodGSR',         caption: 'PeriodGSR',                  size: '200px', sortable: true, render: 'money'},
            {field: 'IncomeOffsets',     caption: 'IncomeOffsets',              size: '200px', sortable: true, render: 'money'},
            {field: 'AmountDue',         caption: 'AmountDue',                  size: '200px', sortable: true, render: 'money'},
            {field: 'PaymentsApplied',   caption: 'Payments Applied',           size: '200px', sortable: true, render: 'money'},
            {field: 'BeginningRcv',	     caption: 'Beginning Receivable',       size: '100px', sortable: false, render: 'money'},
            {field: 'ChangeInRcv',	     caption: 'Change in Receivable',       size: '100px', sortable: false, render: 'money'},
            {field: 'EndingRcv',	     caption: 'Ending Receivable',          size: '100px', sortable: false, render: 'money'},
            {field: 'BeginningSecDep',	 caption: 'Beginning Security Deposit', size: '100px', sortable: false, render: 'money'},
            {field: 'ChangeInSecDep',	 caption: 'Change in Security Deposit', size: '100px', sortable: false, render: 'money'},
            {field: 'EndingSecDep',	     caption: 'Ending Security Deposit',    size: '100px', sortable: false, render: 'money'},
        ],
        onLoad: function(event) {
            var g = this;
            event.onComplete = function() {
                if (!("_rt_offset" in g.last)) {
                    g.last._rt_offset = 0; // rentable offset
                }
                g.last._total = g.total; // total
                for (var i = 0; i < g.records.length; i++) {
                    var record = g.records[i];
                    record.w2ui.class = "";
                    record.w2ui.style = {};

                    // always keep rows expanded
                    g.expand(record.recid);

                    // if it is subtotal row then add class to "tr" tag
                    if (record.IsRentableMainRow) {
                        g.last._rt_offset++;
                    }
                    else if (record.IsSubTotalRow) {
                        record.w2ui.class = "subTotalRow";
                    }
                    else if (record.IsBlankRow) {
                        record.w2ui.class = "blankRow";
                    } else {
                        // apply greyish cell backgroud color to some cells
                        for (var j = 0; j < grey_fields.length; j++) {
                            var colIndex = g.getColumn(grey_fields[j], true);
                            record.w2ui.style[colIndex] = "background-color: grey;";
                        }
                    }
                }
                g.refresh();
                g.total = g.last._total;
                g.offset += g.last._rt_offset;
            };
        },
        /*onRequest: function(event) {
            event.postData.offset = this.last._rt_offset;
        },*/
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

    addDateNavToToolbar('rr');
}

function getRentRollReportData() {
    var x = getCurrentBusiness(),
        BID=parseInt(x.value);

    var g = w2ui.rrGrid,
        rt_offset = g.last._rt_offset;

    return $.get("http://localhost:8270/v1/rrREST/2/", {
        "searchDtStart": "2017-09-01",
        "searchDtStop": "2017-10-01",
        "rt_offset": rt_offset,
        "offset": g.offset,
    }, null, "json")
    .done(function(data){
        if (data.status) {
            g.records = data.records;
            g.render();
        }
    });
}
