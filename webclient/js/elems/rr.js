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
        postData: {searchDtStart: app.D1, searchDtStop: app.D2, limit: 20},
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
            {field: 'IsSubTotalRow',    caption: 'Is SubTotal Row',                           sortable: false, hidden: true},
            {field: 'IsBlankRow',       caption: 'Is Blank Row',                              sortable: false, hidden: true},
            {field: 'IsRentableMainRow',       caption: 'Is Rentable Main Row',               sortable: false, hidden: true},
            {field: 'recid',            caption: 'recid',                      size: '35px',  sortable: true, hidden: true},
            {field: 'BID',              caption: 'BID',                        size: '75px',  sortable: true, hidden: true},
            {field: 'RID',              caption: 'RID',                        size: '75px',  sortable: true, hidden: true},
            {field: 'RentableName',     caption: app.sRentable,                size: '110px', sortable: true},
            {field: 'RTID',             caption: 'RTID',                       size: '75px',  sortable: true, hidden: true},
            {field: 'RentableType',     caption: 'Rentable Type',              size: '100px', sortable: true},
            {field: 'Sqft',             caption: 'Sqft',                       size:  '50px', sortable: true, style: 'text-align: right'},
            {field: 'Description',      caption: 'Description',                size: '150px', sortable: true},
            {field: 'Users',            caption: 'Users',                      size: '150px', sortable: true},
            {field: 'Payors',           caption: 'Payors',                     size: '150px', sortable: true},
            {field: 'RAID',             caption: app.sRentalAgreement,         size: '85px',  sortable: true,
                render: function(record/*,index, col_index*/) {
                    if (typeof record === undefined) {
                        return;
                    }
                    if (record.RAID) {
                        return "RA-" + record.RAID; // return ID with "RA-"
                    }
                }
            },
            {field: 'UsePeriod',        caption: 'Use Period',                 size: '85px',  sortable: true, style: 'text-align: right'},
            {field: 'PossessionStart',  caption: 'PossessionStart',            size: '80px',  sortable: true, render: 'date', style: 'text-align: right', hidden: true},
            {field: 'PossessionStop',   caption: 'PossessionStop',             size: '80px',  sortable: true, render: 'date', style: 'text-align: right', hidden: true},
            {field: 'RentPeriod',       caption: 'Rent<br>Period',             size: '85px',  sortable: true, style: 'text-align: right'},
            {field: 'RentStart',        caption: 'RentStart',                  size: '80px',  sortable: true, render: 'date', style: 'text-align: right', hidden: true},
            {field: 'RentStop',         caption: 'RentStop',                   size: '80px',  sortable: true, render: 'date', style: 'text-align: right', hidden: true},
            {field: 'Agreement',        caption: 'Agreement Period',           size: '200px', sortable: true, style: 'text-align: right', hidden: true},
            {field: 'AgreementStart',   caption: 'AgreementStart',             size: '80px',  sortable: true, render: 'date', style: 'text-align: right', hidden: true},
            {field: 'AgreementStop',    caption: 'AgreementStop',              size: '80px',  sortable: true, render: 'date', style: 'text-align: right', hidden: true},
            {field: 'RentCycle',        caption: 'Rent Cycle',                 size: '75px',  sortable: true,
                render: function(record/*, index, col_index*/) {
                    if (typeof record === undefined) {
                        return;
                    }
                    return app.cycleFreq[record.RentCycle];
                }
            },
            {field: 'GSR',              caption: 'GSR',                              size: '85px',  sortable: true,  render: 'float:2'},
            {field: 'PeriodGSR',        caption: 'Period<br>GSR',                    size: '85px',  sortable: true,  render: 'float:2'},
            {field: 'IncomeOffsets',    caption: 'Income<br>Offsets',                size: '85px',  sortable: true,  render: 'float:2'},
            {field: 'AmountDue',        caption: 'Amount<br>Due',                    size: '85px',  sortable: true,  render: 'float:2'},
            {field: 'PaymentsApplied',  caption: 'Payments<br>Applied',              size: '85px',  sortable: true,  render: 'float:2'},
            {field: 'BeginningRcv',	    caption: 'Beginning<br>Receivable',          size: '100px', sortable: false, render: 'float:2'},
            {field: 'ChangeInRcv',	    caption: 'Change in<br>Receivable',          size: '100px', sortable: false, render: 'float:2'},
            {field: 'EndingRcv',	    caption: 'Ending<br>Receivable',             size: '100px', sortable: false, render: 'float:2'},
            {field: 'BeginningSecDep',	caption: 'Beginning<br>Security<br>Deposit', size: '100px', sortable: false, render: 'float:2'},
            {field: 'ChangeInSecDep',	caption: 'Change in<br>Security<br>Deposit', size: '100px', sortable: false, render: 'float:2'},
            {field: 'EndingSecDep',	    caption: 'Ending<br>Security<br>Deposit',    size: '100px', sortable: false, render: 'float:2'},
        ],
        onLoad: function(event) {
            event.onComplete = function() {
                var g = this;
                if (!("_rt_offset" in g.last)) {
                    g.last._rt_offset = 0;
                }
                var data = JSON.parse(event.xhr.responseText);
                if (data.records) {
                    for (var i = data.records.length - 1; i >= 0; i--) {
                        // get record from grid to apply css
                        var record = g.records[data.records[i].recid];
                        if(record.IsRentableMainRow) {
                            g.last._rt_offset++;
                        }
                        if (!("w2ui" in record)) {
                            record.w2ui = {}; // init w2ui if not present
                        }
                        if (!("class" in record.w2ui)) {
                            record.w2ui.class = ""; // init class string
                        }
                        if (!("style" in record.w2ui)) {
                            record.w2ui.style = {}; // init style object
                        }
                        // var g = w2ui.rrGrid;
                        if (record.IsSubTotalRow) {
                            record.w2ui.class = "subTotalRow";
                        }
                        else if (record.IsBlankRow) {
                            record.w2ui.class = "blankRow";
                        } else {
                            // apply greyish cell backgroud color to some cells
                            for (var j = 0; j < grey_fields.length; j++) {
                                var colIndex = g.getColumn(grey_fields[j], true);
                                record.w2ui.style[colIndex] = "background-color: #CCC;";
                            }
                        }
                    }
                    // everytime you have to assign limit here, otherwise you'll get alert message of differed count
                    // see: https://github.com/vitmalina/w2ui/blob/master/src/w2grid.js#L2488
                    g.limit = data.records.length;
                }

                // stop request if all rows have been loaded
                if(g.total <= g.records.length) {
                    g.last.pull_more = false;
                }

                // need to redraw grid after loading data
                setTimeout(function() {
                    g.refresh();
                }, 0);
            };
        },
        onRequest: function(event) {
            var g = this;
            if (g.records.length == 0) { // if grid is empty then reset all flags
                g.last._rt_offset = 0;
            }
            event.postData.rentableOffset = g.last._rt_offset;
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
