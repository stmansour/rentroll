//------------------------------------------------------------------------
//          receiptsGrid
//------------------------------------------------------------------------
$().w2grid({
    name: 'receiptsGrid',
    url: '/v1/receipts',
    multiSelect: false,
    postData: {searchDtStart: app.D1, searchDtStop: app.D2},
    show: {
        toolbar        : true,
        footer         : true,
        toolbarAdd     : true,   // indicates if toolbar add new button is visible
        toolbarDelete  : false,   // indicates if toolbar delete button is visible
        toolbarSave    : false,   // indicates if toolbar save button is visible
        selectColumn    : false,
        expandColumn    : false,
        toolbarEdit     : false,
        toolbarSearch   : false,
        toolbarInput    : true,
        searchAll       : false,
        toolbarReload   : true,
        toolbarColumns  : true,
    },
    columns: [
        {field: 'recid', hidden: true, caption: 'recid',  size: '40px', sortable: true},                                    // 0
        {field: 'reversed', size: '10px', style: 'text-align: center', sortable: true,
                render: function (record /*, index, col_index*/) {
                    if (typeof record === "undefined") {
                        return;
                    }
                    if ( (record.FLAGS & app.rcptFLAGS.RCPTREVERSED) !== 0 ) { // if reversed then
                        return '<i class="fa fa-exclamation-triangle" title="reversed" aria-hidden="true" style="color: #FFA500;"></i>';
                    }
                    return '';
                },
        },
        {field: 'RCPTID', caption: 'Receipt ID',  size: '80px', style: 'text-align: right', sortable: true},                // 1
        {field: 'Dt', caption: 'Date', size: '80px', sortable: true, style: 'text-align: right'},                           // 2
        {field: 'ARID', caption: 'ARID',  size: '150px', hidden: true, sortable: false}, // 3
        {field: 'AcctRule', caption: 'Account Rule',  size: '150px', sortable: true}, // 4
        {field: 'Amount', caption: 'Amount', size: '100px', sortable: true, render: 'money', style: 'text-align: right'},   // 5
        {field: 'BID', hidden: true, caption: 'BUD', size: '40px', sortable: false},                                        // 6
        {field: 'TCID', hidden: true, caption: 'TCID', size: '40px', sortable: false},                                     // 7
        {field: 'PMTID', hidden: true, caption: 'PMTID', sortable: false},              // 8 - if this changes, update switchToGrid()
        {field: 'PmtTypeName', caption: 'Payment Type', size: '100px', sortable: true},
        {field: 'DocNo', caption: 'Document Number',  size: '150px', style: 'text-align: right', sortable: true},
        {field: 'Payor', caption: 'Payor', size: '150px', sortable: true},                                                  // 3
    ],
    searches : [
        { field: 'Amount', caption: 'Amount', type: 'string' },
        // { field: 'DocNo', caption: 'Document Number', type: 'string' },
        { field: 'Payor', caption: 'Payor', type: 'string' },
        { field: 'PmtTypeName', caption: 'Payment Type', type: 'string' },
        { field: 'AcctRule', caption: 'Account Rule', type: 'string' },
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
                    var x = getCurrentBusiness();
                    var Bid = x.value;
                    var Bud = getBUDfromBID(Bid);
                    $.get('/v1/uival/' + x.value + '/app.Receipts' )
                    .done( function(data) {
                        if (typeof data == 'string') {  // it's weird, a successful data add gets parsed as an object, an error message does not
                            app.Receipts = JSON.parse(data);
                            w2ui.receiptForm.get('ARID').options.items = app.Receipts[Bud];
                            w2ui.receiptForm.refresh();
                            setToForm('receiptForm', '/v1/receipt/' + rec.BID + '/' + rec.RCPTID, 400, true);
                        }
                        if (data.status != 'success') {
                            w2ui.receiptForm.message(data.message);
                        }
                    })
                    .fail( function() {
                        console.log('Error getting /v1/uival/' + x.value + '/app.Receipts');
                     });
                };

            // warn user if form content has been changed
            form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
        };
    },
    onRequest: function(/*event*/) {
        w2ui.receiptsGrid.postData = {searchDtStart: app.D1, searchDtStop: app.D2};
    },
    onRefresh: function(event) {
        event.onComplete = function() {
            if (app.active_grid == this.name) {
                if (app.new_form_rec) {
                    this.selectNone();
                }
                else{
                    this.select(app.last.grid_sel_recid);
                }
            }

            if (event.target == 'monthfwd') {  // we do these tasks after monthfwd is refreshed so we know that the 2 date controls exist
                setDateControlsInToolbar('receipts');
                w2ui.receiptsGrid.postData = {searchDtStart: app.D1, searchDtStop: app.D2};
            }
        };
    },
    onAdd: function (/*event*/) {
        var yes_args = [this],
            no_callBack = function() { return false; },
            yes_callBack = function(grid) {
                // reset it
                app.last.grid_sel_recid = -1;
                grid.selectNone();

                // Insert an empty record...
                var x = getCurrentBusiness();
                var BID=parseInt(x.value);
                var BUD = getBUDfromBID(BID);
                $.get('/v1/uival/' + x.value + '/app.Receipts' )
                .done( function(data) {
                    if (typeof data == 'string') {  // it's weird, a successful data add gets parsed as an object, an error message does not
                        app.Receipts = JSON.parse(data);

                        var y = new Date();
                        var pmt_options = buildPaymentTypeSelectList(BUD);
                        var ptInit = (pmt_options.length > 0) ? pmt_options[0] : '';
                        var record = {
                            recid: 0,
                            RCPTID: 0,
                            PRCPTID: 0,
                            ARID: 0,
                            BID: BID,
                            BUD: BUD,
                            PMTID: 0,
                            PmtTypeName: ptInit,
                            Dt: w2uiDateControlString(y),
                            DocNo: '',
                            Payor: '',
                            TCID: 0,
                            Amount: 0,
                            Comment: '',
                            OtherPayorName: '',
                            FLAGS: 0,
                            LastModTime: y.toISOString(),
                            LastModBy: 0,
                            CreateTS: y.toISOString(),
                            CreateBy: 0
                        };
                        w2ui.receiptForm.fields[0].options.items = pmt_options;
                        w2ui.receiptForm.fields[1].options.items = app.Receipts[BUD];
                        w2ui.receiptForm.record = record;
                        w2ui.receiptForm.refresh();
                        setToForm('receiptForm', '/v1/receipt/' + BID + '/0', 400);
                    }
                    if (data.status != 'success') {
                        w2ui.receiptForm.message(data.message);
                    }
                })
                .fail( function() {
                    console.log('Error getting /v1/uival/'+BUD+'/app.Receipts');
                });
            };

        // warn user if form content has been changed
        form_dirty_alert(yes_callBack, no_callBack, yes_args);
    },
});

addDateNavToToolbar('receipts');

// bind onchange event for date input control for receipts
$(document).on("keypress", "input[name=receiptsD1]", function(e) {
    // do not procedd further untill user press the Enter key
    if (e.which != 13) {
        return;
    }
    var xd1 = document.getElementsByName('receiptsD1')[0].value;
    var xd2 = document.getElementsByName('receiptsD2')[0].value;
    var d1 = dateFromString(xd1);
    var d2 = dateFromString(xd2);
    // check that it is valid or not
    if (isNaN(Date.parse(xd1))) {
        return;
    }
    // check that year is not behind 2000
    if (d1.getFullYear() < 2000) {
        return;
    }
    // check that from date does not have value greater then To date
    if (d1.getTime() >= d2.getTime()) {
        d1 = new Date(d2.getTime() - 24 * 60 * 60 * 1000); //one day back from To date
    }
    app.D1 = dateControlString(d1);
    w2ui.receiptsGrid.postData = {searchDtStart: app.D1, searchDtStop: app.D2};
    w2ui.receiptsGrid.load(w2ui.receiptsGrid.url, function() {
        document.getElementsByName('receiptsD1')[0].focus();
        w2ui.receiptsGrid.refresh();
    });
}).on("keypress", "input[name=receiptsD2]", function(e) {
    // do not procedd further untill user press the Enter key
    if (e.which != 13) {
        return;
    }
    var xd1 = document.getElementsByName('receiptsD1')[0].value;
    var xd2 = document.getElementsByName('receiptsD2')[0].value;
    var d1 = dateFromString(xd1);
    var d2 = dateFromString(xd2);
    // check that it is valid or not
    if (isNaN(Date.parse(xd2))) {
        return;
    }
    // check that year is not behind 2000
    if (d2.getFullYear() < 2000) {
        return;
    }
    // check that from date does not have value greater then To date
    if (d2.getTime() <= d1.getTime()) {
        d2 = new Date(d1.getTime() + 24 * 60 * 60 * 1000); //one day forward from From date
    }
    app.D2 = dateControlString(d2);
    w2ui.receiptsGrid.postData = {searchDtStart: app.D1, searchDtStop: app.D2};
    w2ui.receiptsGrid.load(w2ui.receiptsGrid.url, function() {
        document.getElementsByName('receiptsD2')[0].focus();
        w2ui.receiptsGrid.refresh();
    });
});
