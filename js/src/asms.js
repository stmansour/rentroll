//------------------------------------------------------------------------
//          asmsGrid
//------------------------------------------------------------------------
$().w2grid({
    name: 'asmsGrid',
    url: '/v1/asms',
    multiSelect: false,
    show: {
        toolbar         : true,
        footer          : true,
        toolbarAdd      : true,   // indicates if toolbar add new button is visible
        toolbarDelete   : false,   // indicates if toolbar delete button is visible
        toolbarSave     : false,   // indicates if toolbar save button is visible
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
        {field: 'recid', hidden: true, caption: 'recid',  size: '40px', sortable: true},
        {field: 'epoch', size: '10px', style: 'text-align: center', sortable: true,
                render: function (record /*, index, col_index*/) {
                    if (typeof record === "undefined") {
                        return;
                    }
                    if (record.RentCycle !== 0 && record.PASMID === 0) { // if epoch then make row bold
                        if (record.w2ui === undefined) {
                            record.w2ui = {class:""};
                        }
                        record.w2ui.class = "asmEpochRow";
                        return '<i class="fa fa-refresh" title="epoch" aria-hidden="true"></i>';
                    } else if (record.RentCycle) { // if recurring assessment then put refresh icon
                        if (record.w2ui === undefined) {
                            record.w2ui = {class:""};
                        }
                        record.w2ui.class = "asmInstRow";
                        return '<i class="fa fa-refresh" title="recurring" aria-hidden="true"></i>';
                    }
                    return ''; // if non-recur assessment then do nothing
                },
        },
        {field: 'reversed', size: '10px', style: 'text-align: center', sortable: true,
                render: function (record /*, index, col_index*/) {
                    if (typeof record === "undefined") {
                        return;
                    }
                    if ( (record.FLAGS & app.asmFLAGS.ASMREVERSED) !== 0 ) { // if reversed then
                        return '<i class="fa fa-exclamation-triangle" title="reversed" aria-hidden="true" style="color: #FFA500;"></i>';
                    }
                    return '';
                },
        },
        {field: 'ASMID', caption: 'ASMID',  size: '60px', style: 'text-align: right', sortable: true},
        {field: 'Invoice', caption: 'Invoice', size: '80px', sortable: true, style: 'text-align: right'},
        {field: 'ARID', caption: 'ARID',  hidden: true, sortable: false},
        {field: 'Start', caption: 'Start Date', size: '80px', sortable: true, style: 'text-align: right'},
        {field: 'Stop', caption: 'Stop Date', size: '80px', sortable: true, style: 'text-align: right'},
        {field: 'RentCycle', caption: 'RentCycle',  hidden: true, sortable: false},
        {field: 'AcctRule', caption: 'Account Rule',  size: '200px', style: 'text-align: left', sortable: true},
        {field: 'Amount', caption: 'Amount', size: '100px', sortable: true, render: 'money', style: 'text-align: right'},
        {field: 'BID', hidden: true, caption: 'BUD', size: '40px', sortable: false},
        {field: 'PASMID', hidden: true, caption: 'PMTID', size: '40px', sortable: false},
        {field: 'RAID', caption: app.sRentalAgreement,  size: '125px', style: 'text-align: right', sortable: true},
        {field: 'RID', caption: 'RID',  size: '40px', hidden: true, sortable: false},
        {field: 'Rentable', caption: app.sRentable,  size: '150px', style: 'text-align: right', sortable: true},
        // {field: 'ATypeLID', caption: 'Type', size: '100px', sortable: true, style: 'text-align: right'},
        // {field: 'RentCycle', caption: app.sRentCycle,  size: '60px', style: 'text-align: right', sortable: true},
        // {field: 'ProrationCycle', caption: sProrationCycle,  size: '60px', style: 'text-align: right', sortable: true},
    ],
    searches : [
        { field: 'Amount', caption: 'Amount', type: 'string' },
        // { field: 'Invoice', caption: 'Invoice Number', type: 'string' },
        { field: 'AcctRule', caption: 'Account Rule', type: 'string' },
        { field: 'Rentable', caption: app.sRentable, type: 'string' },
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
                    var form = (rec.RentCycle !== 0 && rec.PASMID === 0) ? "asmEpochForm" : "asmInstForm";
                    var myurl = '/v1/asm/' + rec.BID + '/' + rec.ASMID;
                    console.log( 'calling setToForm( '+form+', ' + myurl + ')');

                    // before setting to the form, get the list of AcctRules...
                    var x = getCurrentBusiness();
                    var Bid = x.value;
                    var Bud = getBUDfromBID(Bid);
                    $.get('/v1/uival/' + x.value + '/app.Assessments' )
                    .done( function(data) {
                        if (typeof data == 'string') {  // it's weird, a successful data add gets parsed as an object, an error message does not
                            app.Assessments = JSON.parse(data);
                            w2ui[form].get('ARID').options.items = app.Assessments[Bud];
                            w2ui[form].refresh();

                            setToForm(form, myurl, 400, true);
                        }
                        if (data.status != 'success') {
                            w2ui.asmsGrid.message(data.message);
                        }
                    })
                    .fail( function() {
                        console.log('Error getting /v1/uival/' + x.value + '/app.Assessments');
                     });
                };

                // warn user if form content has been changed
                form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
        };
    },
    onAdd: function (/*event*/) {
        var yes_args = [this],
            no_callBack = function() { return false; },
            yes_callBack = function(grid) {
                // reset it
                app.last.grid_sel_recid = -1;
                grid.selectNone();

                var x = getCurrentBusiness();
                $.get('/v1/uival/' + x.value + '/app.Assessments' )
                .done( function(data) {
                    if (typeof data == 'string') {  // it's weird, a successful data add gets parsed as an object, an error message does not
                        app.Assessments = JSON.parse(data);

                        // Insert an empty record...
                        var BID=parseInt(x.value);
                        var BUD = getBUDfromBID(BID);
                        app.ridRentablePicker.BID = BID; // needed by typedown

                        var y = new Date();
                        var y1 = new Date(new Date().setFullYear(new Date().getFullYear() + 1));
                        var record = {
                            ARID: 0,
                            recid: 0,
                            RID: 0,
                            ASMID: 0,
                            PASMID: 0,
                            ATypeLID: 0,
                            InvoiceNo: 0,
                            RAID: 0,
                            BID: BID,
                            BUD: BUD,
                            Start: w2uiDateControlString(y),
                            Stop: w2uiDateControlString(y1),
                            RentCycle: 'Monthly',
                            ProrationCycle: 'Daily',
                            TCID: 0,
                            Amount: 0,
                            AcctRule: '',
                            Comment: '',
                            LastModTime: y.toISOString(),
                            LastModBy: 0,
                            CreateTS: y.toISOString(),
                            CreateBy: 0,
                            ExpandPastInst: true,
                            FLAGS: 0,
                            Mode: 0,
                        };
                        // w2ui.asmEpochForm.fields[5].options.url = '/v1/rentablestd/' + app.ridRentablePicker.BID;
                        w2ui.asmEpochForm.fields[0].options.items = app.Assessments[BUD];
                        w2ui.asmEpochForm.record = record;
                        w2ui.asmEpochForm.refresh();

                        setToForm('asmEpochForm', '/v1/asm/' + BID + '/0', 400);
                    }
                    if (data.status != 'success') {
                        w2ui.asmEpochForm.message(data.message);
                    }
                })
                .fail( function() {
                    console.log('Error getting /v1/uival/'+x.value+'/app.Assessments');
                 });
            };

            // warn user if form content has been changed
            form_dirty_alert(yes_callBack, no_callBack, yes_args);
    },
    onRequest: function(/*event*/) {
        w2ui.asmsGrid.postData = {searchDtStart: app.D1, searchDtStop: app.D2};
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
                setDateControlsInToolbar('asms');
                w2ui.asmsGrid.postData = {searchDtStart: app.D1, searchDtStop: app.D2};
            }
        };
    }
});

addDateNavToToolbar('asms');

// bind onchange event for date input control for assessments
$(document).on("keypress", "input[name=asmsD1]", function(e) {
    // do not procedd further untill user press the Enter key
    if (e.which != 13) {
        return;
    }
    var xd1 = document.getElementsByName('asmsD1')[0].value;
    var xd2 = document.getElementsByName('asmsD2')[0].value;
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
    w2ui.asmsGrid.postData = {searchDtStart: app.D1, searchDtStop: app.D2};
    w2ui.asmsGrid.load(w2ui.asmsGrid.url, function() {
        document.getElementsByName('asmsD1')[0].focus();
        w2ui.asmsGrid.refresh();
    });
}).on("keypress", "input[name=asmsD2]", function(e) {
    // do not procedd further untill user press the Enter key
    if (e.which != 13) {
        return;
    }
    var xd1 = document.getElementsByName('asmsD1')[0].value;
    var xd2 = document.getElementsByName('asmsD2')[0].value;
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
    w2ui.asmsGrid.postData = {searchDtStart: app.D1, searchDtStop: app.D2};
    w2ui.asmsGrid.load(w2ui.asmsGrid.url, function() {
        document.getElementsByName('asmsD2')[0].focus();
        w2ui.asmsGrid.refresh();
    });
});
