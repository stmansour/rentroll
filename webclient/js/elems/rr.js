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
            {field: 'RAID',             caption: app.sRentalAgreement,         size: '150px', sortable: true,
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
            {field: 'RentPeriod',       caption: 'Rent Period',                size: '85px',  sortable: true, style: 'text-align: right'},
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
        onLoad: function(event) {
            var g = this,
                records_length = g.records.length,
                data = JSON.parse(event.xhr.responseText);

            if (event.status != "success") {
                return;
            }

            if (!("_rt_offset" in g.last)) {
                g.last._rt_offset = 0; // rentable offset
            }
            if (!("_rrIndexMap" in g.last)) {
                g.last._rrIndexMap = {};
            }
            if (g.last.scrollTop == 0 && records_length == 0) {
                g.last._rt_offset = 0; // rentable offset
                g.last._rrIndexMap = {};
            }

            event.onComplete = function() {
                for (var i = records_length; i < g.records.length; i++) {
                    var record = g.records[i];
                    record.w2ui.class = "";
                    record.w2ui.style = {};

                    // if it is subtotal row then add class to "tr" tag
                    if (record.IsRentableMainRow) {
                        // always keep rows expanded, if it is main row then
                        g.expand(record.recid);

                        g.last._rrIndexMap[i] = g.last._rt_offset;
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
                g.total = data.total;
                setTimeout(function() {
                    calculateRRPagination(); // at last we need to execute this custom pagination mechanism
                }, 0);
            };
        },
        onRequest: function(event) {
            event.postData.rentableOffset = this.last._rt_offset;
        },
        onRefresh: function(event) {
            var g = this;
            event.onComplete = function() {
                $("#grid_"+g.name+"_records").on("scroll", function() {
                    calculateRRPagination();
                });
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

function calculateRRPagination() {
    // perform virtual scroll
    var g = w2ui.rrGrid;
    var url  = (typeof g.url != 'object' ? g.url : g.url.get);
    var records = $("#grid_" + g.name + "_records");
    var buffered = g.records.length;
    if (g.searchData.length != 0 && !url) buffered = g.last.searchIds.length;
    if (buffered === 0 || records.length === 0 || records.height() === 0) return;
    if (buffered > g.vs_start) g.last.show_extra = g.vs_extra; else g.last.show_extra = g.vs_start;
    // need this to enable scrolling when g.limit < then a screen can fit
    if (records.height() < buffered * g.recordHeight && records.css('overflow-y') == 'hidden') {
        // TODO: is this needed?
        // if (g.total > 0) g.refresh();
        return;
    }
    // update footer
    var t1 = Math.round(records[0].scrollTop / g.recordHeight);
    var t2 = t1 + (Math.round(records.height() / g.recordHeight));
    if (t1 > buffered) t1 = buffered - 1;
    if (t2 >= buffered) t2 = buffered - 1;
    // custom pagination number start - stop for rentroll report
    var startPageRec = 0, endPageRec = 0, i;
    for (i = t1; i >= 0; i--) {
        if(g.records[i].IsRentableMainRow){
            startPageRec = i;
            break;
        }
    }

    for (i = t2; i >= t1; i--) {
        if(g.records[i].IsRentableMainRow){
            endPageRec = i;
            break;
        }
    }

    var startPageNo = g.last._rrIndexMap[startPageRec] + 1;
    var endPageNo = g.last._rrIndexMap[endPageRec] + 1;

    $('#grid_'+ g.name + '_footer .w2ui-footer-right').html(
        (g.show.statusRange ? w2utils.formatNumber(startPageNo) + '-' + w2utils.formatNumber(endPageNo) +
                (g.total != -1 ? ' ' + w2utils.lang('of') + ' ' +    w2utils.formatNumber(g.total) : '') : '') +
        (url && g.show.statusBuffered ? ' ('+ w2utils.lang('buffered') + ' '+ w2utils.formatNumber(buffered) +
                (g.offset > 0 ? ', skip ' + w2utils.formatNumber(g.offset) : '') + ')' : '')
    );
}

