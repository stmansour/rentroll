/*global
    $, console, w2ui, w2uiDateControlString, app, plural, dateFmtStr,
    getCurrentBusiness, getBUDfromBID, w2popup, w2utils, w2confirm, form_dirty_alert,
    delete_confirm_options, formRefreshCallBack, getFormSubmitData, addDateNavToToolbar,
    ridRentablePickerRender,ridRentableDropRender,ridRentableCompare,tcidRUserPickerRender,
    tcidPickerDropRender, tcidPickerCompare, tcidRAPayorPickerRender, calcRarGridContractRent, popupRidRentablePicker,
    popupTcidRAPayorPicker, popupTcidRUserPicker, saveNewRAPayor, saveNewRUser, saveNewRARentable
*/

"use strict";
//-----------------------------------------------------------------------------
// setToRAForm -  enable the Rental Agreement form in toplayout.  Also, set
//                the forms url and request data from the server
// @params
//   bid = business id (or the BUD)
//  raid = Rental Agreement ID
//     d = date to use for time sensitive data
//-----------------------------------------------------------------------------
window.setToRAForm = function (bid, raid, d) {

    if (raid > 0) {
        var f = w2ui.rentalagrForm;
        w2ui.toplayout.content('right', w2ui.raLayout);
        w2ui.toplayout.show('right', true);
        w2ui.toplayout.sizeTo('right', app.WidestFormWidth);
        f.url = '/v1/rentalagr/' + bid + '/' + raid;
        f.request();
        w2ui.toplayout.render();

        // mark this flag as is this new record
        // record created already
        app.new_form_rec = false;

        // as new content will be loaded for this form
        // mark form dirty flag as false
        app.form_is_dirty = false;

        // click on first tab
        if (typeof f.tabs.name == "string") {
            f.tabs.click('tab1');
        }
    }

    //----------------------------------------------------------------
    // Get the associated Rentables...
    //      /v1/rar/bid/raid[?d1=2017-02-1]
    //      if no date is specified, today's date is used as the default.
    //----------------------------------------------------------------
    w2ui.rarGrid.url = '/v1/rar/' + bid + '/' + raid;
    // console.log('rar url = ' + w2ui.rarGrid.url);
    w2ui.rarGrid.request();
    w2ui.rarGrid.header = plural(app.sRentable) + ' as of ' + dateFmtStr(d);
    w2ui.rarGrid.show.toolbarSearch = false;

    //----------------------------------------------------------------
    // Get the associated Payors...
    //      /v1/rapeople/bid/raid[?type=payor&d1=2017-02-1]
    //      if no date is specified, today's date is used as the default.
    //      if no person type is provided, payor is assumed
    //----------------------------------------------------------------
    w2ui.rapGrid.url = '/v1/rapayor/' + bid + '/' + raid;
    // console.log('rapGrid url = ' + w2ui.rapGrid.url);
    w2ui.rapGrid.request();
    w2ui.rapGrid.header = plural(app.sPayor) + ' as of ' + dateFmtStr(d);

    //----------------------------------------------------------------
    // Get the associated Users...
    //      /v1/ruser/bid/raid[?&d1=2017-02-1]
    //      if no date is specified, today's date is used as the default.
    //----------------------------------------------------------------
    w2ui.rauGrid.url = '/v1/ruser/' + bid + '/' + raid;
    // console.log('rauGrid url = ' + w2ui.rauGrid.url);
    w2ui.rauGrid.request();
    w2ui.rauGrid.header = plural(app.sUser) + ' as of ' + dateFmtStr(d);

    //----------------------------------------------------------------
    // Get the associated Pets...
    //      /v1/xrapets/bid/raid
    //----------------------------------------------------------------
    w2ui.raPetGrid.url = '/v1/rapets/' + bid + '/' + raid;
    // console.log('xrapets url = ' + w2ui.rarGrid.url);
    w2ui.raPetGrid.request();
    w2ui.raPetGrid.header = 'Pets as of ' + dateFmtStr(d);
};

window.saveNewRUser = function () {
    var myRID = -1;
    var rname = w2ui.tcidRUserPicker.record.RentableName.text;
    var url = '/v1/ruser/';
    for (var i = 0; i < app.TcidRUserPicker.RAR.records.length; i++) {
        if (app.TcidRUserPicker.RAR.records[i].RentableName == rname) {
            myRID = app.TcidRUserPicker.RAR.records[i].RID;
            url += '' + app.TcidRUserPicker.RAR.records[i].BID + '/' + myRID;
        }
    }
    if (myRID < 1) {
        console.log('RID = ' + myRID);
        return;
    }
    var rec = {
        recid: w2ui.rapGrid.records.length,               // + 1
        BID: w2ui.rentalagrForm.record.BID,               // BID of RID we're editing
        TCID: w2ui.tcidRUserPicker.record.TCID,           // The TCID from the person picked
        RID: myRID,                                       // the selected RID
        RentableName: rname,                              // the name of the rentable
        FirstName: w2ui.tcidRUserPicker.record.FirstName, // ignored
        MiddleName: "",                                   // ignored
        LastName: w2ui.tcidRUserPicker.record.LastName,   // ignored
        IsCompany: w2ui.tcidRUserPicker.record.IsCompany,
        CompanyName: w2ui.tcidRUserPicker.record.CompanyName,
        DtStart: w2ui.tcidRUserPicker.record.DtStart,     // start
        DtStop: w2ui.tcidRUserPicker.record.DtStop,       // end
    };
    var params = {cmd: 'save', formname: 'tcidRUserPicker', record: rec };
    var dat = JSON.stringify(params);

    // save the record to the server, insert into grid if successful
    $.post(url, dat, null, "json")
    .done(function(data) {
        if (data.status == 'success') {
            rec.recid = data.recid; // success reply returns RUID of record just added
            w2ui.rauGrid.add(rec,true);
        } else {
            w2ui.rauGrid.message(data.message);
        }
    })
    .fail(function(data) {
        console.log('data = ' + data);
    });
};

window.saveNewRARentable = function () {
    var rec = {
        recid: 0,                                                   // new
        BUI: w2ui.rentalagrForm.record.BUD.text,                         // BID of RID we're editing
        RAID: w2ui.rentalagrForm.record.RAID,                       // The RAID from the form being edited
        RID: w2ui.ridRentablePicker.record.RID,                     // the selected RID
        RentableName: w2ui.ridRentablePicker.record.RentableName[0].RentableName,   // the name of the rentable
        RARDtStart: w2ui.ridRentablePicker.record.DtStart,          // start
        RARDtStop: w2ui.ridRentablePicker.record.DtStop,            // end
        ContractRent: w2ui.ridRentablePicker.record.Amount,         // end
    };
    var params = {cmd: 'save', formname: 'ridRentablePicker', record: rec };
    var dat = JSON.stringify(params);

    // save the record to the server, insert into grid. We need the
    // RARID of the record just saved so that we can store it along
    // with the rest of the record information that we insert into the grid.
    // We pick this up in the "success" return information
    var url = '/v1/rar/' + rec.BUI + '/' + rec.RAID;
    $.post(url, dat, null, "json")
    .done(function(data) {
        if (data.status == 'success') {
            rec.recid = data.recid; // success reply returns RARID of record just added
            w2ui.rarGrid.add(rec,true);
        } else {
            w2ui.rarGrid.message(data.message);
        }
    })
    .fail(function(data) {
        console.log('data = ' + data);
    });
};

window.buildRAElements = function() {
    //------------------------------------------------------------------------
    //          rentalagrsGrid
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'rentalagrsGrid',
        url: '/v1/rentalagrs',
        multiSelect: false,
        show: {
            toolbar: true,
            footer: true,
            lineNumbers: false,
            selectColumn: false,
            expandColumn: false,
            toolbarAdd: true,
            toolbarDelete: false,
            toolbarSave: false,
            toolbarEdit: false,
            toolbarSearch: true,
            toolbarInput: true,
            searchAll: true,
            toolbarReload: true,
            toolbarColumns: false,
        },
        multiSearch: true,
        searches: [
            { field: 'RAID', caption: 'RAID', type: 'text' },
            { field: 'Payors', caption: 'Payor(s)', type: 'text' },
            { field: 'AgreementStart', caption: 'Agreement Start Date', type: 'date' },
            { field: 'AgreementStop', caption: 'Agreement Stop Date', type: 'date' },
        ],
        columns: [
            {field: 'recid', hidden: true, caption: 'recid',  size: '40px', sortable: true},
            {field: 'BID', hidden: true, caption: 'BID',  size: '40px', sortable: false},
            {field: 'RAID', caption: 'RAID',  size: '50px', sortable: true},
            {field: 'Payors', caption: 'Payor(s)', size: '250px', sortable: true,
                // render: function (record, index, col_index) {
                //     if (record) {
                //         var icon;
                //         if (record.PayorIsCompany) {
                //             icon = 'fa-handshake-o'
                //         } else {
                //             icon = 'fa-user-o'
                //         }
                //         return '<i class="fa '+icon+'"></i>&nbsp;<span>'+record.Payors+'</span>';
                //     }
                // },
            },
            {field: 'AgreementStart', caption: 'Agreement<br>Start', render: 'date', size: '80px', sortable: true, style: 'text-align: right'},
            {field: 'AgreementStop', caption: 'Agreement<br>Stop',  render: 'date', size: '80px', sortable: true, style: 'text-align: right'},
            {field: 'PayorIsCompany', hidden: true, caption: 'IsCompany',  size: '40px', sortable: false},
        ],
        onRefresh: function(event) {
            event.onComplete = function() {
                // var sel_recid = parseInt(this.last.sel_recid);
                if (app.active_grid == this.name) {
                    this.select(app.last.grid_sel_recid);
                    // This one is special case, you need to set last sel_recid when you're adding
                    // new record with help of onAdd event handler, so new record automatically
                    // will be selected

                    /*if (app.new_form_rec) {
                        this.unselect(sel_recid);
                    }
                    else{
                        this.select(sel_recid);
                    }*/
                }
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
                        var d = new Date();  // we'll use today for time-sensitive data
                        setToRAForm(rec.BID, rec.RAID, d);
                    };

                // warn user if form content has been changed
                form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
            };
        },
        onAdd: function (/*event*/) {
            var yes_args = [this],
                no_callBack = function() { return false; },
                yes_callBack = function(grid) {

                    // this record will be added in grid, so set last sel_recid to this record
                    // so it will be automatically highlighted
                    var newRecId = grid.records.length; // get length and assign recid to new record
                    grid.last.sel_recid = String(newRecId);

                    // NOTE: this one is special case, we're adding record which is not new one so select it anyways
                    app.last.grid_sel_recid = newRecId;

                    // Insert an empty record...
                    var x = getCurrentBusiness();
                    var BID=parseInt(x.value);
                    var BUD = getBUDfromBID(BID);
                    var y = new Date();
                    var y1 = new Date(new Date().setFullYear(new Date().getFullYear() + 1));
                    var y2 = new Date();
                    y2.setDate(1); // set it to the first of the month
                    var rec = {
                        recid: newRecId,
                        RAID: 0,
                        RATID: 0,
                        BID: BID,
                        BUD: BUD,
                        NLID: 0,
                        AgreementStart: w2uiDateControlString(y),
                        AgreementStop: w2uiDateControlString(y1),
                        PossessionStart: w2uiDateControlString(y),
                        PossessionStop: w2uiDateControlString(y1),
                        RentStart: w2uiDateControlString(y),
                        RentStop: w2uiDateControlString(y1),
                        RentCycleEpoch: w2uiDateControlString(y2),
                        UnspecifiedAdults: 0,
                        UnspecifiedChildren: 0,
                        Renewal: "month to month automatic renewal",
                        SpecialProvisions: '',
                        LastModTime: y.toISOString(),
                        LastModBy: 0,
                        LeaseType: 0,
                        ExpenseAdjustmentType: 0,
                        ExpensesStop: 0,
                        ExpenseStopCalculation: '',
                        BaseYearEnd: '1/1/1900',
                        ExpenseAdjustment: '1/1/1900',
                        EstimatedCharges: 0,
                        RateChange: 0,
                        NextRateChange: '1/1/1900',
                        PermittedUses: '',
                        ExclusiveUses: '',
                        ExtensionOption: '',
                        ExtensionOptionNotice: '1/1/1900',
                        ExpansionOption: '',
                        ExpansionOptionNotice: '1/1/1900',
                        RightOfFirstRefusal: '',
                    };
                    var url = '/v1/rentalagr/' + BUD + '/0';

                    var request={cmd:"save",recid:0,name:"rentalagrForm",record: rec};
                    $.post(url, JSON.stringify(request), null, "json")
                    .done(function(data) {
                        if (data.status == 'success') {
                            rec.recid = data.recid; // success reply returns RAPID of record just added
                            rec.RAID = data.recid;
                            setToRAForm(BID, rec.RAID, y);
                        } else {
                            w2ui.rentalagrsGrid.message(data.message);
                        }
                    });
                };

            // warn user if form content has been changed
            form_dirty_alert(yes_callBack, no_callBack, yes_args);
        },
    });
    addDateNavToToolbar('rentalagrs');


    //------------------------------------------------------------------------
    //          Rental Agreement Details
    //------------------------------------------------------------------------
    $().w2layout({
        name: 'raLayout',
        padding: 0,
        panels: [
            { type: 'left', size: '40%', resizable: true, style: app.pstyle, content: 'top' },
            { type: 'top', size: 0, hidden: true },
            { type: 'main', size: '60%', resizable: true, style: app.pstyle, content: 'main' },
            { type: 'preview', size: 0, maxSize: 1, content: 'PREVIEW',  hidden: true },
            { type: 'bottom', size: 0, minSize: 0, maxSize: 1, hidden: true },
            { type: 'right', size: 0, minSize: 0, maxSize: 1, hidden: true }
        ]
    });
    $().w2layout({
        name: 'raLayoutSub1',
        padding: 1,
        panels: [
            { type: 'top', size: '30%', resizable: true, style: app.pstyle4, content: 'Rentables' },
            { type: 'left', size: 0, hidden: true },
            { type: 'main', size: '20%', resizable: true, style: app.pstyle4, content: 'Payors' },
            { type: 'preview', size: '60%', resizable: true, style: app.pstyle4, content: 'Users' },
            { type: 'right', size: 0, hidden: true },
            { type: 'bottom', size: '25%', resizable: true, style: app.pstyle4, content: 'Pets' }
        ]
    });

    //------------------------------------------------------------------------
    //          rentalagrForm
    //------------------------------------------------------------------------
    $().w2form({
        name: 'rentalagrForm',
        style: 'border: 0px; background-color: transparent;',
        header: app.sRentalAgreement + ' Detail',
        //url: '/v1/rentalagrs',
        formURL: '/webclient/html/formra.html',
        toolbar: {
            items: [
                { id: 'btnNotes', type: 'button', icon: 'far fa-sticky-note' },
                { id: 'bt3', type: 'spacer' },
                { id: 'btnClose', type: 'button', icon: 'fas fa-times' },
            ],
            onClick: function (event) {
                switch(event.target) {
                case 'btnClose':
                    var no_callBack = function() { return false; },
                        yes_callBack = function() {
                            w2ui.toplayout.hide('right',true);
                            w2ui.rentalagrsGrid.render();
                        };
                    form_dirty_alert(yes_callBack, no_callBack);
                    break;
                }
            },
        },
        fields: [
            { field: 'recid', type: 'int', required: false, html: {page: 0, column: 0 } },
            { field: 'RAID', type: 'int', required: false, html: {  page: 0, column: 0 } },
            { field: 'RATID', type: 'int', required: false, html: { page: 0, column: 0 } },
            { field: 'BID', type: 'int', required: false, html: { page: 0, column: 0 } },
            { field: 'BUD', type: 'list', options: {items: app.businesses}, required: true, html: { page: 0, column: 0 } },
            { field: 'NLID', type: 'int', required: false, html: {page: 0, column: 0 } },
            { field: 'AgreementStart', type: 'date', required: true, html: { page: 0, column: 0 } },
            { field: 'AgreementStop', type: 'date', required: true, html: { page: 0, column: 0 } },
            { field: 'PossessionStart', type: 'date', required: true, html: { page: 0, column: 0 } },
            { field: 'PossessionStop', type: 'date', required: true, html: { page: 0, column: 0 } },
            { field: 'RentStart', type: 'date', required: true, html: { page: 0, column: 0 } },
            { field: 'RentStop', type: 'date', required: true, html: { page: 0, column: 0 } },
            { field: 'RentCycleEpoch', type: 'date', required: true, html: {  page: 0, column: 0 } },
            { field: 'UnspecifiedAdults', type: 'int', required: false, html: { page: 0, column: 0 } },
            { field: 'UnspecifiedChildren', type: 'int', required: false, html: { page: 0, column: 0 } },
            { field: 'Renewal', type: 'list', options: {items: app.renewalMap}, required: true, html: { page: 0, column: 0 } },
            { field: 'SpecialProvisions', type: 'text', required: false, html: { page: 0, column: 0 } },
            { field: 'LastModTime', type: 'time', required: false, html: { page: 0, column: 0 } },
            { field: 'LastModBy', type: 'int', required: false, html: { page: 0, column: 0 } },
            { field: 'CreateTS', type: 'time', required: false, html: { page: 0, column: 0 } },
            { field: 'CreateBy', type: 'int', required: false, html: { page: 0, column: 0 } },
            { field: 'LeaseType', type: 'w2int', required: false, html: {page: 1, column: 0}},
            { field: 'ExpenseAdjustmentType', type: 'w2int', required: false, html: {page: 1, column: 0}},
            { field: 'ExpensesStop', type: 'w2float', required: false, html: {page: 1, column: 0}},
            { field: 'ExpenseStopCalculation', type: 'text', required: false, html: {page: 1, column: 0}},
            { field: 'BaseYearEnd', type: 'date', required: false, html: {page: 1, column: 0}},
            { field: 'ExpenseAdjustment', type: 'date', required: false, html: {page: 1, column: 0}},
            { field: 'EstimatedCharges', type: 'w2float', required: false, html: {page: 1, column: 0}},
            { field: 'RateChange', type: 'w2float', required: false, html: {page: 1, column: 0}},
            { field: 'NextRateChange', type: 'date', required: false, html: {page: 1, column: 0}},
            { field: 'PermittedUses', type: 'text', required: false, html: {page: 1, column: 0}},
            { field: 'ExclusiveUses', type: 'text', required: false, html: {page: 1, column: 0}},
            { field: 'ExtensionOption', type: 'text', required: false, html: {page: 1, column: 0}},
            { field: 'ExtensionOptionNotice', type: 'date', required: false, html: {page: 1, column: 0}},
            { field: 'ExpansionOption', type: 'text', required: false, html: {page: 1, column: 0}},
            { field: 'ExpansionOptionNotice', type: 'date', required: false, html: {page: 1, column: 0}},
            { field: 'RightOfFirstRefusal', type: 'text', required: false, html: {page: 1, column: 0}},
        ],
        tabs: [
            { id: 'tab1', caption: app.sRAMainTab },
            { id: 'tab2', caption: app.sRACommercialTab },
        ],
        actions: {
            save: function () {
                var tgrid = w2ui.rentalagrsGrid;
                tgrid.selectNone();
                this.save({}, function (data) {
                    if (data.status == 'error') {
                        console.log('ERROR: '+ data.message);
                        return;
                    }
                    w2ui.toplayout.hide('right',true);
                    tgrid.render();
                });
            },
            delete: function(/*target, data*/) {

                var form = this;

                w2confirm(delete_confirm_options)
                .yes(function() {

                    var tgrid = w2ui.rentalagrsGrid;
                    var params = {cmd: 'delete', formname: form.name, RAID: form.record.RAID };
                    var dat = JSON.stringify(params);

                    // delete RentalAgreement request
                    $.post(form.url, dat, null, "json")
                    .done(function(data) {
                        if (data.status === "error") {
                            form.error(w2utils.lang(data.message));
                            return;
                        }
                        w2ui.toplayout.hide('right',true);
                        tgrid.remove(app.last.grid_sel_recid);
                        tgrid.render();
                    })
                    .fail(function(/*data*/){
                        form.error("Delete RentalAgreement failed.");
                        return;
                    });
                })
                .no(function() {
                    return;
                });
            },
        },
        onRefresh: function(event) {
            event.onComplete = function() {
                var f = this,
                    header = "Edit Rental Agreement ({0})";

                formRefreshCallBack(f, "RAID", header);
            };
        },
        onValidate: function(event) {
            if (this.record.AgreementStart === '') {
                event.errors.push({
                    field: this.get('AgreementStart'),
                    error: 'Agreement Start date cannot be blank'
                });
            }
            if (this.record.AgreementStop === '') {
                event.errors.push({
                    field: this.get('AgreementStop'),
                    error: 'Agreement Stop date cannot be blank'
                });
            }
            if (this.record.PossessionStart === '') {
                event.errors.push({
                    field: this.get('PossessionStart'),
                    error: 'Possession Start date cannot be blank'
                });
            }
            if (this.record.PossessionStop === '') {
                event.errors.push({
                    field: this.get('PossessionStop'),
                    error: 'Possession Stop date cannot be blank'
                });
            }
            if (this.record.RentStart === '') {
                event.errors.push({
                    field: this.get('RentStart'),
                    error: 'Rent Start date cannot be blank'
                });
            }
            if (this.record.RentStop === '') {
                event.errors.push({
                    field: this.get('RentStop'),
                    error: 'Rent Stop date cannot be blank'
                });
            }
             if (this.record.RentCycleEpoch === '') {
                event.errors.push({
                    field: this.get('RentCycleEpoch'),
                    error: 'Rent Cycle Epoch date cannot be blank'
                });
            }
        },
        onChange: function(event) {
            event.onComplete = function() {
                // formRecDiffer: 1=current record, 2=original record, 3=diff object
                var diff = formRecDiffer(this.record, app.active_form_original, {});
                // if diff == {} then make dirty flag as false, else true
                if ($.isPlainObject(diff) && $.isEmptyObject(diff)) {
                    app.form_is_dirty = false;
                } else {
                    app.form_is_dirty = true;
                }
            };
        },
        onSubmit: function(target, data) {
            delete data.postData.record.LastModTime;
            delete data.postData.record.LastModBy;
            delete data.postData.record.CreateTS;
            delete data.postData.record.CreateBy;
            // modify form data for server request
            getFormSubmitData(data.postData.record);
            app.renewalMap.forEach(function(item) {
                if (item.id == data.postData.record.Renewal) {
                    data.postData.record.Renewal = item.text;
                    return false;
                }
            });
        },
    });

    //------------------------------------------------------------------------
    //          rarGrid
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'rarGrid',
        show: {
            header: app.sRentable,
            toolbar: true,
            footer: false,
            lineNumbers: false,
            selectColumn: false,
            expandColumn: false,
            toolbarAdd: true,
            toolbarDelete: true,
            toolbarSave: true,
            toolbarEdit: false,
            toolbarSearch: false,
            toolbarInput: false,
            searchAll: false,
            toolbarReload: false,
            toolbarColumns: false,
        },
        url: '/v1/rar/',
        columns: [
            {field: 'recid', hidden: true, caption: 'recid',  size: '40px', sortable: true},
            {field: 'RAID', hidden: true, caption: 'RAID',  size: '40px', sortable: true},
            {field: 'RID', hidden: true, caption: app.sRentable,  size: '10%', sortable: true},
            {field: 'RentableName', caption: app.sRentable+' Name',  size: '10%', sortable: true},
            {field: 'ContractRent', caption: 'Rent', size: '10%', sortable: true, render: 'money', editable: { type: 'money' }},
            {field: 'RARDtStart', caption: 'Start', size: '10%', sortable: true, style: 'text-align: right', editable: { type: 'date' }},
            {field: 'RARDtStop', caption: 'Stop',  size: '10%', sortable: true, style: 'text-align: right', editable: { type: 'date' }},
        ],
        onAdd: function (/*event*/) {
            var x = getCurrentBusiness();
            popupRidRentablePicker(app.sRentable, x.value);
        },
        onChange: function(event) {
            var g = this;
            event.done(function () {
                calcRarGridContractRent(g);
                g.save();
            });
        },
        onLoad: function(event) {
            var g = this;
            event.done(function () {
                if (w2ui.rarGrid.summary.length === 0) {
                    w2ui.rarGrid.add({recid: 's-1', RAID: 0, RID: 0, RentableName: 'Total Rent', ContractRent: 0, RARDtStart: '', RARDtStop: '', w2ui: {summary: true}});
                }
                calcRarGridContractRent(g);
            });
        },
     });

    //------------------------------------------------------------------------
    //          rapGrid
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'rapGrid',
        show: {
            header: app.sPayor,
            toolbar: true,
            footer: false,
            lineNumbers: false,
            selectColumn: false,
            expandColumn: false,
            toolbarAdd: true,
            toolbarDelete: true,
            toolbarSave: true,
            toolbarEdit: false,
            toolbarSearch: false,
            toolbarInput: false,
            searchAll: false,
            toolbarReload: false,
            toolbarColumns: false,
        },
        url: '/v1/rapayor',
        columns: [
            {field: 'recid', hidden: true, caption: 'recid',  size: '40px', sortable: true},
            {field: 'TCID', hidden: true, caption: 'TCID',  size: '40px', sortable: true},
            {field: 'FirstName', caption: 'First Name',  size: '10%', sortable: true},
            {field: 'MiddleName', caption: 'MI',  size: '30px', sortable: true},
            {field: 'LastName', caption: 'Last Name',  size: '10%', sortable: true},
            {field: 'IsCompany', caption: 'IsCompany', hidden: true, sortable: false },
            {field: 'CompanyName', caption: 'CompanyName',  size: '10%', sortable: true, style: 'text-align: left',
                render: function (record/*, index, col_index*/) {
                    var html = '';
                    if (typeof record === "undefined") {
                        return;
                    }
                    if (record.IsCompany) {
                        html = record.CompanyName;
                    } else {
                        html = '';
                    }
                    return html;
                },
            },
            {field: 'DtStart', caption: 'Start', size: '10%', sortable: true, style: 'text-align: right', editable: { type: 'date' }},
            {field: 'DtStop', caption: 'Stop',  size: '10%', sortable: true, style: 'text-align: right', editable: { type: 'date' }},
        ],
        onAdd: function (/*event*/) {
            var x = getCurrentBusiness();
            popupTcidRAPayorPicker(app.sPayor, x.value );
        },
        onChange: function(event) {
            event.done(function () {
                w2ui[event.target].save();
            });
        },
        onSave: function() {
            console.log('save payors');
        },
        onDelete: function(/*event*/) {
            var sel = w2ui.rapGrid.getSelection(true); // get the record indeces rather than the recids
            w2ui.rapGrid.postData = {
                RAPID: w2ui.rapGrid,
                TCID: w2ui.rapGrid.records[sel[0]].TCID,
                DtStart: w2ui.rapGrid.records[sel[0]].DtStart,
                DtStop: w2ui.rapGrid.records[sel[0]].DtStop
            };
        }
     });

    //------------------------------------------------------------------------
    //          rauGrid
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'rauGrid',
        show: {
            header: app.sUser,
            toolbar: true,
            toolbarAdd: true,
            toolbarDelete: true,
            toolbarSave: true,
            footer: false,
            lineNumbers: false,
            selectColumn: false,
            expandColumn: false,
            toolbarEdit: false,
            toolbarSearch: false,
            toolbarInput: false,
            searchAll: false,
            toolbarReload: false,
            toolbarColumns: false,
        },
        url: '/v1/ruser',
        columns: [
            {field: 'recid', hidden: true, caption: 'recid',  size: '40px', sortable: true},
            {field: 'TCID', hidden: true, caption: 'TCID',  size: '40px', sortable: true},
            {field: 'RID', hidden: true, caption: 'RID',  size: '40px', sortable: true},
            {field: 'BID', hidden: true, caption: 'RID',  size: '40px', sortable: true},
            {field: 'RentableName', caption: app.sRentable,  size: '10%', sortable: true},
            {field: 'FirstName', caption: 'First Name',  size: '10%', sortable: true},
            {field: 'MiddleName', caption: 'MI',  size: '30px', sortable: true},
            {field: 'LastName', caption: 'Last Name',  size: '10%', sortable: true},
            {field: 'IsCompany', caption: 'IsCompany', hidden: true, sortable: false },
            {field: 'CompanyName', caption: 'CompanyName',  size: '10%', sortable: true, style: 'text-align: left',
                render: function (record/*, index, col_index*/) {
                    var html = '';
                    if (typeof record === "undefined") {
                        return;
                    }
                    if (record.IsCompany) {
                        html = record.CompanyName;
                    } else {
                        html = '';
                    }
                    return html;
                },
            },
            {field: 'DtStart', caption: 'Start', size: '10%', sortable: true, style: 'text-align: right', editable: { type: 'date' }},
            {field: 'DtStop', caption: 'Stop',  size: '10%', sortable: true, style: 'text-align: right', editable: { type: 'date' }},
        ],
        onAdd: function (/*event*/) {
            var x = getCurrentBusiness();
            console.log('rentalagrForm.record.RAID = ' + w2ui.rentalagrForm.record.RAID);
            popupTcidRUserPicker(app.sUser, x.value, w2ui.rentalagrForm.record.RAID );
        },
        onChange: function(event) {
            event.done(function () {
                w2ui[event.target].save();
            });
        },
        onSave: function() {
            console.log('save users');
        },
        onDelete: function(/*event*/) {
            var sel = w2ui.rauGrid.getSelection(true); // get the record indeces rather than the recids
            var x = getCurrentBusiness();
            w2ui.rauGrid.postData = {
                TCID: w2ui.rauGrid.records[sel[0]].TCID,
            };
            w2ui.rauGrid.url = '/v1/ruser/' + x.value + '/' + w2ui.rauGrid.records[sel[0]].RID;
        }
     });

    //------------------------------------------------------------------------
    //          raPetGrid
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'raPetGrid',
        show: {
            header: app.sUser,
            toolbar: true,
            footer: false,
            toolbarAdd: true,
            toolbarDelete: true,
            toolbarSave: true,
            toolbarEdit: false,
            toolbarSearch: false,
            toolbarInput: false,
            searchAll: false,
            toolbarReload: false,
            toolbarColumns: false,
        },
        url: '/v1/rapets',
        columns: [
            {field: 'recid', hidden: true, caption: 'recid',  size: '40px', sortable: true},
            {field: 'PETID', hidden: true, caption: 'PETID',  size: '40px', sortable: true},
            {field: 'Name',     caption: 'Name', size: '10%', sortable: true},
            {field: 'Type',     caption: 'Type', size: '10%', sortable: true},
            {field: 'Breed',    caption: 'Breed', size: '10%', sortable: true},
            {field: 'Color',    caption: 'Color', size: '10%', sortable: true},
            {field: 'Weight',   caption: 'Weight', size: '10%', sortable: true},
            {field: 'DtStart',  caption: 'DtStart', size: '10%', sortable: true, style: 'text-align: right'},
            {field: 'DtStop',   caption: 'DtStop', size: '10%', sortable: true, style: 'text-align: right'},
        ],
        onDelete: function(/*event*/) {
            var sel = w2ui.rPetGrid.getSelection(true); // get the record indeces rather than the recids
            w2ui.rPetGrid.postData = {
                PETID: w2ui.rPetGrid.records[sel[0]].PETID,
            };
        }

     });
};

window.popupTcidRAPayorPicker = function (sTitle,bid) {
    app.TcidRAPayorPicker = {BID: bid, Title: sTitle};  // used by tcidRAPayorPicker form
    var y = new Date();
    var ny = new Date(y.getFullYear() + 1, y.getMonth(), y.getDate(), 0, 0, 0);
    w2ui.tcidRAPayorPicker.record.DtStart = w2uiDateControlString(y);
    w2ui.tcidRAPayorPicker.record.DtStop = w2uiDateControlString(ny);
    w2ui.tcidRAPayorPicker.record.TCID = -1;
    w2ui.tcidRAPayorPicker.record.pickedName = '';
    w2ui.tcidRAPayorPicker.record.IsCompany = false;
    w2ui.tcidRAPayorPicker.record.CompanyName = '';
    w2ui.tcidRAPayorPicker.refresh();
    var raid = w2ui.rentalagrForm.record.RAID;
    $().w2popup('open', {
        title   : 'Add New ' + sTitle + ' to RA0'+raid,
        body    : '<div id="form" style="width: 100%; height: 100%;"></div>',
        style   : 'padding: 15px 0px 0px 0px',
        width   : 400,
        height  : 200,
        showMax : true,
        onToggle: function (event) {
            $(w2ui.tcidRAPayorPicker.box).hide();
            event.onComplete = function () {
                $(w2ui.tcidRAPayorPicker.box).show();
                w2ui.tcidRAPayorPicker.resize();
            };
        },
        onOpen: function (event) {
            event.onComplete = function () {
                // specifying an onOpen handler instead would be equivalent to specifying
                // an onBeforeOpen handler, which would make this code execute too
                // early and hence not deliver.
                $('#w2ui-popup #form').w2render('tcidRAPayorPicker');
            };
        }
    });
};

window.popupRidRentablePicker = function (sTitle,bid) {
    app.ridRentablePicker = {BID: bid, Title: sTitle};  // used by RidRentablePicker form
    var y = new Date();
    var ny = new Date(y.getFullYear() + 1, y.getMonth(), y.getDate(), 0, 0, 0);
    w2ui.ridRentablePicker.record.DtStart = w2uiDateControlString(y);
    w2ui.ridRentablePicker.record.DtStop = w2uiDateControlString(ny);
    w2ui.ridRentablePicker.record.TCID = -1;
    w2ui.ridRentablePicker.record.RentableName = '';
    w2ui.ridRentablePicker.record.Amount = '';
    w2ui.ridRentablePicker.refresh();
    var raid = w2ui.rentalagrForm.record.RAID;
    $().w2popup('open', {
        title   : 'Add New ' + sTitle + ' to RA0' + raid,
        body    : '<div id="form" style="width: 100%; height: 100%;"></div>',
        style   : 'padding: 15px 0px 0px 0px',
        width   : 400,
        height  : 240,
        showMax : true,
        onToggle: function (event) {
            $(w2ui.ridRentablePicker.box).hide();
            event.onComplete = function () {
                $(w2ui.ridRentablePicker.box).show();
                w2ui.ridRentablePicker.resize();
            };
        },
        onOpen: function (event) {
            event.onComplete = function () {
                // specifying an onOpen handler instead would be equivalent to specifying
                // an onBeforeOpen handler, which would make this code execute too
                // early and hence not deliver.
                $('#w2ui-popup #form').w2render('ridRentablePicker');
            };
        }
    });
};

window.popupTcidRUserPicker = function (sTitle,bid,raid) {
    app.TcidRUserPicker = {BID: bid, Title: sTitle, RAID: raid, RARentablesNames: [], RAR: []};  // used by tcidRUserPicker form
    var y = new Date();
    var ny = new Date(y.getFullYear() + 1, y.getMonth(), y.getDate(), 0, 0, 0);
    w2ui.tcidRUserPicker.record.DtStart = w2uiDateControlString(y);
    w2ui.tcidRUserPicker.record.DtStop = w2uiDateControlString(ny);
    w2ui.tcidRUserPicker.record.TCID = -1;
    w2ui.tcidRUserPicker.record.pickedName = '';
    w2ui.tcidRUserPicker.record.IsCompany = false;
    w2ui.tcidRUserPicker.record.CompanyName = '';
    w2ui.tcidRUserPicker.refresh();

    var url = '/v1/rar/' + bid + '/' + raid;
    $.get(url, null, null, "json")
    .done(function(data) {
        if (data.status != "success") {
            return;
        }

        app.TcidRUserPicker.RAR = data;
        for (var i = 0; i < app.TcidRUserPicker.RAR.records.length; i++) {
            app.TcidRUserPicker.RARentablesNames.push(app.TcidRUserPicker.RAR.records[i].RentableName);
        }
    });

    $().w2popup('open', {
        title   : 'Add A New ' + sTitle,
        body    : '<div id="form" style="width: 100%; height: 100%;"></div>',
        style   : 'padding: 15px 0px 0px 0px',
        width   : 400,
        height  : 250,
        showMax : true,
        onToggle: function (event) {
            $(w2ui.tcidRUserPicker.box).hide();
            event.onComplete = function () {
                $(w2ui.tcidRUserPicker.box).show();
                w2ui.tcidRUserPicker.resize();
            };
        },
        onOpen: function (event) {
            event.onComplete = function () {
                // specifying an onOpen handler instead would be equivalent to specifying
                // an onBeforeOpen handler, which would make this code execute too
                // early and hence not deliver.
                $('#w2ui-popup #form').w2render('tcidRUserPicker');
            };
        }
    });
};

window.saveNewRAPayor = function () {

    var rec = {
        recid: w2ui.rapGrid.records.length,                 // + 1
        BID: w2ui.rentalagrForm.record.BID,                 // BID of RAID we're editing
        TCID: w2ui.tcidRAPayorPicker.record.TCID,           // The TCID from the person picked
        RAID: w2ui.rentalagrForm.record.RAID,               // the RAID we're editing
        FirstName: w2ui.tcidRAPayorPicker.record.FirstName, // ignored
        MiddleName: "",                                     // ignored
        LastName: w2ui.tcidRAPayorPicker.record.LastName,   // ignored
        IsCompany: int_to_bool(w2ui.tcidRAPayorPicker.record.IsCompany),
        CompanyName: w2ui.tcidRAPayorPicker.record.CompanyName,   // ignored
        DtStart: w2ui.tcidRAPayorPicker.record.DtStart, // start
        DtStop: w2ui.tcidRAPayorPicker.record.DtStop,                  // end
    };
    var params = {cmd: 'save', formname: 'tcidRAPayorPicker', record: rec };
    var dat = JSON.stringify(params);

    $.post(w2ui.rapGrid.url, dat, null, "json")
    .done(function(data) {
        if (data.status == 'success') {
            rec.recid = data.recid; // success reply returns RAPID of record just added
            w2ui.rapGrid.add(rec,true);
        } else {
            w2ui.rapGrid.message(data.message);
        }
    })
    .fail(function(data) {
        console.log('data = ' + data);
    });

};

window.buildRAPayorPicker = function (){
    //------------------------------------------------------------------------
    //          tcidRAPayorPicker
    //------------------------------------------------------------------------
    $().w2form({
        name: 'tcidRAPayorPicker',
        style: 'border: 0px; background-color: transparent;',
        formURL: '/webclient/html/tcidrapayorpicker.html',
        focus  : 0,
        fields: [
            { field: 'TCID', type: 'w2int', required: true },
            { field: 'pickedName', required: true,
                type: 'enum',
                options: {
                    url:           '/v1/transactantstd/' + app.TcidRAPayorPicker.BID,
                    max:           1,
                    maxDropHeight: 350,
                    renderItem:    tcidRAPayorPickerRender,
                    renderDrop:    tcidPickerDropRender,
                    compare:       tcidPickerCompare,
                    onNew: function (event) {
                        //console.log('++ New Item: Do not forget to submit it to the server too', event);
                        $.extend(event.item, { FirstName: '', LastName : event.item.text });
                    }
                },
            },
            { field: 'DtStart', type: 'date', required: true },
            { field: 'DtStop', type: 'date', required: true },
            { field: 'FirstName', type: 'text', required: false },
            { field: 'LastName', type: 'text', required: false },
            { field: 'CompanyName', type: 'text', required: false },
            { field: 'IsCompany', type: 'checkbox', required: false },
        ],
        onRefresh: function(/*event*/) {
            w2ui.tcidRAPayorPicker.fields[1].options.url = '/v1/transactantstd/' + app.TcidRAPayorPicker.BID;
            console.log('SAVE tcidRAPayorPicker: TCID = ' + w2ui.tcidRAPayorPicker.record.TCID + '  DtStart = ' + w2ui.tcidRAPayorPicker.record.DtStart + '  DtStop = ' + w2ui.tcidRAPayorPicker.record.DtStop);
        },
        actions: {
            save: function () {
                var errs = w2ui.tcidRAPayorPicker.validate(true);
                if (errs.length > 0) {
                    return;
                }
                w2popup.close();
                saveNewRAPayor();
            }
        },
        onSubmit: function(target, data){
            // server request form data
            getFormSubmitData(data.postData.record);
        }
    });
};

window.buildRUserPicker = function (){
    //------------------------------------------------------------------------
    //          tcidRUserPicker
    //------------------------------------------------------------------------
    $().w2form({
        name: 'tcidRUserPicker',
        style: 'border: 0px; background-color: transparent;',
        formURL: '/webclient/html/tcidruserpicker.html',
        focus  : 0,
        fields: [
            { field: 'TCID', type: 'w2int', required: true },
            { field: 'pickedName', required: true,
                type: 'enum',
                options: {
                    url:           '/v1/transactantstd/' + app.TcidRUserPicker.BID,
                    max:           1,
                    maxDropHeight: 350,
                    renderItem:    tcidRUserPickerRender,
                    renderDrop:    tcidPickerDropRender,
                    compare:       tcidPickerCompare,
                    onNew: function (event) {
                        //console.log('++ New Item: Do not forget to submit it to the server too', event);
                        $.extend(event.item, { FirstName: '', LastName : event.item.text });
                    }
                },
            },
            { field: 'RentableName', type: 'list', required: true, options: { items: [] } },
            { field: 'DtStart', type: 'date', required: true },
            { field: 'DtStop', type: 'date', required: true },
            { field: 'FirstName', type: 'text', required: false },
            { field: 'LastName', type: 'text', required: false },
            { field: 'CompanyName', type: 'text', required: false },
            { field: 'IsCompany', type: 'checkbox', required: false },
        ],
        onRefresh: function(event) {
            event.onComplete = function() {
                w2ui.tcidRUserPicker.fields[1].options.url = '/v1/transactantstd/' + app.TcidRUserPicker.BID;
                w2ui.tcidRUserPicker.fields[2].options.items = app.TcidRUserPicker.RARentablesNames;
                if (app.TcidRUserPicker.RARentablesNames.length === 1) {
                    w2ui.tcidRUserPicker.record.RentableName = app.TcidRUserPicker.RARentablesNames[0];
                }
            };
        },
        actions: {
            save: function () {
                var errs = w2ui.tcidRUserPicker.validate(true);
                if (errs.length > 0) { return; }
                console.log('SAVE tcidRUserPicker: TCID = ' + w2ui.tcidRUserPicker.record.TCID + '  DtStart = ' + w2ui.tcidRUserPicker.record.DtStart + '  DtStop = ' + w2ui.tcidRUserPicker.record.DtStop);
                w2popup.close();
                saveNewRUser();
            }
        },
        onSubmit: function(target, data){
            // server request form data
            getFormSubmitData(data.postData.record);
        }
    });
};

window.buildRentablePicker = function (){
    //------------------------------------------------------------------------
    //          ridRentablePicker
    //------------------------------------------------------------------------
    $().w2form({
        name: 'ridRentablePicker',
        style: 'border: 0px; background-color: transparent;',
        formURL: '/webclient/html/ridrentablepicker.html',
        focus  : 0,
        openOnFocus: true,
        fields: [
            { field: 'RentableName', required: true,
                type: 'enum',
                options: {
                    url:           '/v1/rentablestd/' + app.ridRentablePicker.BID,
                    max:           1,
                    cacheMax:      50,
                    maxDropHeight: 350,
                    renderItem:    ridRentablePickerRender,
                    renderDrop:    ridRentableDropRender,
                    compare:       ridRentableCompare,
                    onNew:         function (event) {
                        //console.log('++ New Item: Do not forget to submit it to the server too', event);
                        $.extend(event.item, { RentableName : event.item.text });
                    }
                },
            },
            { field: 'DtStart', type: 'date',  required: true },
            { field: 'DtStop',  type: 'date',  required: true },
            { field: 'Amount',  type: 'money', required: true },
            { field: 'RID',     type: 'wsint' },
        ],
        onRefresh: function(/*event*/) {
            w2ui.ridRentablePicker.fields[0].options.url = '/v1/rentablestd/' + app.ridRentablePicker.BID;
        },
        actions: {
            save: function () {
                var errs = w2ui.ridRentablePicker.validate(true);
                if (errs.length > 0) { return; }
                console.log('SAVE ridRentablePicker: TCID = ' + w2ui.ridRentablePicker.record.TCID + '  DtStart = ' + w2ui.ridRentablePicker.record.DtStart + '  DtStop = ' + w2ui.ridRentablePicker.record.DtStop);
                w2popup.close();
                saveNewRARentable();
            },
        },
        onSubmit: function(target, data){
            // server request form data
            getFormSubmitData(data.postData.record);
        },
    });
};

window.createRentalAgreementForm = function () {
    w2ui.raLayout.content('left', w2ui.rentalagrForm);
    w2ui.raLayout.content('main', w2ui.raLayoutSub1);

    w2ui.raLayoutSub1.content('top', w2ui.rarGrid);
    w2ui.rarGrid.header = plural(app.sRentable);
    w2ui.raLayoutSub1.content('main', w2ui.rapGrid);
    w2ui.rapGrid.header = plural(app.sPayor);

    w2ui.raLayoutSub1.content('preview', w2ui.rauGrid);
    w2ui.rauGrid.header = plural(app.sUser);
    w2ui.raLayoutSub1.content('bottom', w2ui.raPetGrid);
    w2ui.raPetGrid.header = "Pets";
};
