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
        postData: {searchDtStart: app.D1, searchDtStop: app.D2, limit: 15},
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
            toolbarInput    : false,    // the text area for searches
            searchAll       : false,
            toolbarReload   : false,
            toolbarColumns  : false,
        },
        columns: [
            {field: 'IsSubTotalRow',            caption: 'Is SubTotal Row',                                 sortable: false, hidden: true},
            {field: 'IsBlankRow',               caption: 'Is Blank Row',                                    sortable: false, hidden: true},
            {field: 'IsMainRow',                caption: 'Is Main Row',                                     sortable: false, hidden: true},
            {field: 'IsRentableSectionMainRow', caption: 'Is Rentable Section Main Row',                    sortable: false, hidden: true},
            {field: 'IsNoRentableSectionRow',   caption: 'Is No Rentalbe Section Row',                      sortable: false, hidden: true},
            {field: 'recid',                    caption: 'recid',                            size: '35px',  sortable: true,  hidden: true},
            {field: 'BID',                      caption: 'BID',                              size: '75px',  sortable: true,  hidden: true},
            {field: 'RID',                      caption: 'RID',                              size: '75px',  sortable: true,  hidden: true},
            {field: 'RentableName',             caption: app.sRentable,                      size: '110px', sortable: true},
            {field: 'RTID',                     caption: 'RTID',                             size: '75px',  sortable: true,  hidden: true},
            {field: 'RentableType',             caption: 'Rentable Type',                    size: '100px', sortable: true},
            {field: 'Sqft',                     caption: 'Sqft',                             size:  '50px', sortable: true,                                    style: 'text-align: right'},
            {field: 'Description',              caption: 'Description',                      size: '150px', sortable: true},
            {field: 'Users',                    caption: 'Users',                            size: '150px', sortable: true},
            {field: 'Payors',                   caption: 'Payors',                           size: '150px', sortable: true},
            {field: 'RAID',                     caption: app.sRentalAgreement,               size: '85px',  sortable: true,  hidden: true},
            {field: 'RAIDStr',                  caption: app.sRentalAgreement,               size: '85px',  sortable: true},
            {field: 'UsePeriod',                caption: 'Use Period',                       size: '85px',  sortable: true,                                    style: 'text-align: right'},
            {field: 'PossessionStart',          caption: 'PossessionStart',                  size: '80px',  sortable: true,  hidden: true, render: 'date',     style: 'text-align: right'},
            {field: 'PossessionStop',           caption: 'PossessionStop',                   size: '80px',  sortable: true,  hidden: true, render: 'date',     style: 'text-align: right'},
            {field: 'RentPeriod',               caption: 'Rent<br>Period',                   size: '85px',  sortable: true,                                    style: 'text-align: right'},
            {field: 'RentStart',                caption: 'RentStart',                        size: '80px',  sortable: true,  hidden: true, render: 'date',     style: 'text-align: right'},
            {field: 'RentStop',                 caption: 'RentStop',                         size: '80px',  sortable: true,  hidden: true, render: 'date',     style: 'text-align: right'},
            {field: 'AgreementPeriod',          caption: 'Agreement Period',                 size: '200px', sortable: true,  hidden: true,                     style: 'text-align: right'},
            {field: 'AgreementStart',           caption: 'AgreementStart',                   size: '80px',  sortable: true,  hidden: true, render: 'date',     style: 'text-align: right'},
            {field: 'AgreementStop',            caption: 'AgreementStop',                    size: '80px',  sortable: true,  hidden: true, render: 'date',     style: 'text-align: right'},
            {field: 'RentCycle',                caption: 'Rent Cycle',                       size: '75px',  sortable: true,  hidden: true},
            {field: 'RentCycleStr',             caption: 'Rent Cycle',                       size: '85px',  sortable: true},
            {field: 'GSR',                      caption: 'GSR',                              size: '85px',  sortable: true,                render: 'float:2'},
            {field: 'PeriodGSR',                caption: 'Period<br>GSR',                    size: '85px',  sortable: true,                render: 'float:2'},
            {field: 'IncomeOffsets',            caption: 'Income<br>Offsets',                size: '85px',  sortable: true,                render: 'float:2'},
            {field: 'AmountDue',                caption: 'Amount<br>Due',                    size: '85px',  sortable: true,                render: 'float:2'},
            {field: 'PaymentsApplied',          caption: 'Payments<br>Applied',              size: '85px',  sortable: true,                render: 'float:2'},
            {field: 'BeginningRcv',             caption: 'Beginning<br>Receivable',          size: '100px', sortable: false,               render: 'float:2'},
            {field: 'ChangeInRcv',              caption: 'Change in<br>Receivable',          size: '100px', sortable: false,               render: 'float:2'},
            {field: 'EndingRcv',                caption: 'Ending<br>Receivable',             size: '100px', sortable: false,               render: 'float:2'},
            {field: 'BeginningSecDep',          caption: 'Beginning<br>Security<br>Deposit', size: '100px', sortable: false,               render: 'float:2'},
            {field: 'ChangeInSecDep',           caption: 'Change in<br>Security<br>Deposit', size: '100px', sortable: false,               render: 'float:2'},
            {field: 'EndingSecDep',             caption: 'Ending<br>Security<br>Deposit',    size: '100px', sortable: false,               render: 'float:2'},
        ],
        onLoad: function(event) {
            var g = this;
            var data = JSON.parse(event.xhr.responseText);
            if (!("_total_main_rows" in g)) {
                g._total_main_rows = 0;
            }
            if (data.total_main_rows) {
                g._total_main_rows = data.total_main_rows;
            }

            // everytime you have to assign limit here, otherwise you'll get alert message of differed count
            // see: https://github.com/vitmalina/w2ui/blob/master/src/w2grid.js#L2488 and 2481
            if (data.records) {
                g.limit = data.records.length;
            }

            event.onComplete = function() {
                if (!("_main_rows_offset" in g.last)) {
                    g.last._main_rows_offset = 0;
                }
                if (!("_rrIndexMap" in g.last)) {
                    g.last._rrIndexMap = {};
                }
                if (!("_rows_offset" in g.last)) {
                    g.last._rows_offset = 0;
                }

                var record, i;
                if (data.records) {
                    for (i = 0; i < data.records.length; i++) {

                        // get record from grid to apply css
                        record = g.get(data.records[i].recid);

                        if (!("w2ui" in record)) {
                            record.w2ui = {}; // init w2ui if not present
                        }
                        if (!("class" in record.w2ui)) {
                            record.w2ui.class = ""; // init class string
                        }
                        if (!("style" in record.w2ui)) {
                            record.w2ui.style = {}; // init style object
                        }

                        if(record.IsMainRow) {
                            var rec_index = g.get(record.recid, true);
                            g.last._rrIndexMap[rec_index] = g.last._main_rows_offset;
                            g.last._main_rows_offset++;
                        }

                        if (record.IsRentRollViewRow) {
                            // apply greyish cell backgroud color to some cells
                            for (var j = 0; j < grey_fields.length; j++) {
                                var colIndex = g.getColumn(grey_fields[j], true);
                                record.w2ui.style[colIndex] = "background-color: #CCC;";
                            }
                            g.last._rows_offset++;
                        }

                        if (record.IsSubTotalRow) {
                            record.w2ui.class = "subTotalRow";
                        }

                        if (record.IsBlankRow) {
                            record.w2ui.class = "blankRow";
                        }

                        // redraw row
                        g.refreshRow(data.records[i].recid);
                    }
                }

                // summary rows
                if (data.summary) {
                    for (i = 0; i < data.summary.length; i++) {

                        // get record from grid to apply css
                        record = g.get(data.summary[i].recid);

                        if (!("w2ui" in record)) {
                            record.w2ui = {}; // init w2ui if not present
                        }
                        if (!("class" in record.w2ui)) {
                            record.w2ui.class = ""; // init class string
                        }
                        if (!("style" in record.w2ui)) {
                            record.w2ui.style = {}; // init style object
                        }

                        if (record.IsGrandTotalRow) {
                            record.w2ui.class = "grandTotalRow";
                        }

                        // redraw row
                        g.refreshRow(data.summary[i].recid);
                    }
                }

                // stop request if all rows have been loaded
                if(g.total <= g.records.length) {
                    g.last.pull_more = false;
                }

                // need to redraw grid after loading data
                setTimeout(function() {
                    calculateRRPagination();
                }, 0);
            };
        },
        onRefresh: function(event) {
            var g = this;
            event.onComplete = function() {
                $("#grid_"+g.name+"_records").on("scroll", function() {
                    calculateRRPagination();
                });
            };
        },
        onRequest: function(event) {
            var g = this;
            if (g.records.length == 0) { // if grid is empty then reset all flags
                g._total_main_rows = 0;
                g.last._main_rows_offset = 0;
                g.last._rows_offset = 0;
                g.last._rrIndexMap = {};
            }
            event.postData.rows_offset = g.last._rows_offset;
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

    // now add items for csv/pdf export report options
    w2ui.rrGrid.toolbar.add([
        { type: 'spacer',},
        { type: 'button', id: 'csvexport', icon: 'fa fa-table', tooltip: 'export to CSV' },
        { type: 'button', id: 'printreport', icon: 'fa fa-file-pdf-o', tooltip: 'export to PDF' },
        { type: 'break', id: 'break2' },
        { type: 'menu-radio', id: 'page_size', icon: 'fa fa-print',
            tooltip: 'exported PDF page size',
            text: function (item) {
            //var text = item.selected;
            var el   = this.get('page_size:' + item.selected);
            if (item.selected == "Custom") {
                popupPDFCustomDimensions();
            }
            return 'Page Size: ' + el.text;
            },
            selected: 'USLetter',
            items: [
                { id: 'USLetter', text: 'US Letter (8.5 x 11 in)'},
                { id: 'Legal', text: 'Legal (8.5 x 14 in)'},
                { id: 'Ledger', text: 'Ledger (11 x 17 in)'},
                { id: 'Custom', text: 'Custom'},
            ]
        },
        { type: 'menu-radio', id: 'orientation', icon: 'fa fa-clone',
            tooltip: 'exported PDF orientation',
            text: function (item) {
            //var text = item.selected;
            var el   = this.get('orientation:' + item.selected);
            var pageSize = w2ui.reportstoolbar.get('page_size').selected;
            if (pageSize != "Custom" && item.selected == "Portrait") {
                app.pdfPageWidth = app.pageSizes[pageSize].w;
                app.pdfPageHeight = app.pageSizes[pageSize].h;
            }
            else if (pageSize != "Custom" && item.selected == "LandScape") {
                app.pdfPageWidth = app.pageSizes[pageSize].h;
                app.pdfPageHeight = app.pageSizes[pageSize].w;
            }
            return 'Orientation: ' + el.text;
            },
            selected: 'LandScape',
            items: [
                { id: 'LandScape', text: 'LandScape'},
                { id: 'Portrait', text: 'Portrait'},
            ]
        },
    ]);

    // handle pdf/csv report download actions
    w2ui.rrGrid.toolbar.on('click', function(event) {
        var d1, d2; // start date, stop date

        if (event.target == "csvexport") {
            d1 = document.getElementsByName("rrD1")[0].value;
            app.D1 = d1;
            d2 = document.getElementsByName("rrD2")[0].value;
            app.D2 = d2;

            // now call to export csv report function with start and stop date
            exportReportCSV("RPTrr", d1, d2);
        }
        else if (event.target == "printreport") {
            d1 = document.getElementsByName("rrD1")[0].value;
            app.D1 = d1;
            d2 = document.getElementsByName("rrD2")[0].value;
            app.D2 = d2;

            // call to export pdf report function with start and stop date
            exportReportPDF("RPTrr", d1, d2);
        }
    });
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
    var t1 = Math.round(records[0].scrollTop / g.recordHeight + 1);
    var t2 = t1 + (Math.round(records.height() / g.recordHeight) - 1);
    if (t1 > buffered) t1 = buffered - 1;
    if (t2 >= buffered) t2 = buffered - 1;
    // custom pagination number start - stop for rentroll report
    var startPageRec = 0, endPageRec = 0, i;
    for (i = t1; i >= 0; i--) {
        if(g.records[i].IsMainRow){
            startPageRec = i;
            break;
        }
    }

    for (i = t2; i >= t1; i--) {
        if(g.records[i].IsMainRow){
            endPageRec = i;
            break;
        }
    }

    var startPageNo = g.last._rrIndexMap[startPageRec] + 1;
    var endPageNo = g.last._rrIndexMap[endPageRec] + 1;

    $('#grid_'+ g.name + '_footer .w2ui-footer-right').html(
        (g.show.statusRange ? w2utils.formatNumber(startPageNo) + '-' + w2utils.formatNumber(endPageNo) +
        (g._total_main_rows != -1 ? ' ' + w2utils.lang('of') + ' ' +    w2utils.formatNumber(g._total_main_rows) : '') : '')
    );
}