/*global
    w2ui, $, app, w2confirm, getBUDfromBID, getCurrentBusiness, setToForm,
    console, form_dirty_alert, buildPaymentTypeSelectList, setDateControlsInToolbar,
    addDateNavToToolbar, tcidReceiptPayorPickerRender, tcidPickerDropRender, tcidPickerCompare,
    getPersonDetailsByTCID, getPaymentType, formRefreshCallBack, w2utils, reverse_confirm_options,
    getFormSubmitData, w2uiDateControlString, getGridReversalSymbolHTML, get2XReversalSymbolHTML,
*/
"use strict";
function getReceiptInitRecord(BID, BUD, ptInit){
    var y = new Date();
    return {
        recid: 0,
        RCPTID: 0,
        PRCPTID: 0,
        ARID: 0,
        PMTID: 0,
        PmtTypeName: ptInit,
        BID: BID,
        BUD: BUD,
        DID: 0,
        Dt: w2uiDateControlString(y),
        LastModTime: y.toISOString(),
        CreateTS: y.toISOString(),
        DocNo: '',
        Payor: '',
        TCID: 0,
        Amount: 0,
        Comment: '',
        OtherPayorName: '',
        FLAGS: 0,
        LastModBy: 0,
        CreateBy: 0
    };
}

function buildReceiptElements() {
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
            toolbarAdd     : true,    // indicates if toolbar add new button is visible
            toolbarDelete  : false,   // indicates if toolbar delete button is visible
            toolbarSave    : false,   // indicates if toolbar save button is visible
            selectColumn   : false,
            expandColumn   : false,
            toolbarEdit    : false,
            toolbarSearch  : false,
            toolbarInput   : true,
            searchAll      : false,
            toolbarReload  : true,
            toolbarColumns : true,
        },
        columns: [
            {field: 'recid', hidden: true, caption: 'recid',  size: '40px', sortable: true},
            {field: 'reversed', size: '10px', style: 'text-align: center', sortable: true,
                    render: function (record /*, index, col_index*/) {
                        if (typeof record === "undefined") {
                            return;
                        }
                        if ( (record.FLAGS & app.rcptFLAGS.RCPTREVERSED) !== 0 ) { // if reversed then
                            return getGridReversalSymbolHTML();
                        }
                        return '';
                    },
            },
            {field: 'RCPTID',      caption: 'Receipt ID',     size: '80px',  hidden: false, sortable: true, style: 'text-align: right'},
            {field: 'Dt',          caption: 'Date',           size: '80px',  hidden: false, sortable: true, style: 'text-align: right'},
            {field: 'ARID',        caption: 'ARID',           size: '150px', hidden: true,  sortable: false},
            {field: 'DID',         caption: 'DID',            size: '150px', hidden: false, sortable: false},
            {field: 'AcctRule',    caption: 'Account Rule',   size: '150px', hidden: false, sortable: true},
            {field: 'Amount',      caption: 'Amount',         size: '100px', hidden: false, sortable: true, render: 'money', style: 'text-align: right'},
            {field: 'BID',         caption: 'BUD',            size: '40px',  hidden: true,  sortable: false},
            {field: 'TCID',        caption: 'TCID',           size: '40px',  hidden: true,  sortable: false},
            {field: 'PMTID',       caption: 'PMTID',                         hidden: true,  sortable: false},
            {field: 'PmtTypeName', caption: 'Payment Type',   size: '100px', hidden: false, sortable: true},
            {field: 'DocNo',       caption: 'Document Number',size: '150px', hidden: false, sortable: true, style: 'text-align: right'},
            {field: 'Payor',       caption: 'Payor',          size: '150px', hidden: false, sortable: true},
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
                        grid.select(app.last.grid_sel_recid);// keep highlighting current row in any case

                        var rec = grid.get(recid);
                        var x = getCurrentBusiness();
                        var Bid = x.value;
                        var Bud = getBUDfromBID(Bid);
                        $.get('/v1/uival/' + x.value + '/app.ReceiptRules' )
                        .done( function(data) {
                            if (typeof data == 'string') {  // it's weird, a successful data add gets parsed as an object, an error message does not
                                app.ReceiptRules = JSON.parse(data);
                                w2ui.receiptForm.get('ARID').options.items = app.ReceiptRules[Bud];
                                w2ui.receiptForm.refresh();
                                setToForm('receiptForm', '/v1/receipt/' + rec.BID + '/' + rec.RCPTID, 400, true);
                            }
                            if (data.status != 'success') {
                                w2ui.receiptForm.message(data.message);
                            }
                        })
                        .fail( function() {
                            console.log('Error getting /v1/uival/' + x.value + '/app.ReceiptRules');
                         });
                    };
                form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);  // warn user if form content has been changed
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
                    $.get('/v1/uival/' + x.value + '/app.ReceiptRules' )
                    .done( function(data) {
                        if (typeof data == 'string') {  // it's weird, a successful data add gets parsed as an object, an error message does not
                            app.ReceiptRules = JSON.parse(data);

                            var pmt_options = buildPaymentTypeSelectList(BUD);
                            var ptInit = (pmt_options.length > 0) ? pmt_options[0] : '';
                            var record = getReceiptInitRecord(BID, BUD, ptInit);
                            w2ui.receiptForm.fields[0].options.items = pmt_options;
                            w2ui.receiptForm.fields[1].options.items = app.ReceiptRules[BUD];
                            w2ui.receiptForm.record = record;
                            w2ui.receiptForm.refresh();
                            setToForm('receiptForm', '/v1/receipt/' + BID + '/0', 400);
                        }
                        if (data.status != 'success') {
                            w2ui.receiptForm.message(data.message);
                        }
                    })
                    .fail( function() {
                        console.log('Error getting /v1/uival/'+BUD+'/app.ReceiptRules');
                    });
                };

            // warn user if form content has been changed
            form_dirty_alert(yes_callBack, no_callBack, yes_args);
        },
    });

    addDateNavToToolbar('receipts');

    //------------------------------------------------------------------------
    //          receiptForm
    //------------------------------------------------------------------------
    $().w2form({
        name: 'receiptForm',
        style: 'border: 0px; background-color: transparent;',
        header: app.sRentable + ' Detail',
        url: '/v1/receipt',
        formURL: '/webclient/html/formrcpt.html',
        fields: [
            { field: 'PmtTypeName',    type: 'list', required: true, options: { items: [], selected: {} }, html: { caption: "BUD", page: 0 } }, // keep this at position 0 as the list changes and we need to update it
            { field: 'ARID',           type: 'list',  required: true, options:  {items: app.ReceiptRules} },  // 1
            { field: 'recid',          type: 'int',   required: false },                                     // 2
            { field: "BID", required: false, type: 'int', html: { caption: "BID", page: 0, column: 0 } },
            { field: "BUD", required: true, options: { items: app.businesses, maxDropHeight: 350 }, type: 'list', html: { caption: "BUD", page: 0, column: 0 } },
            { field: 'RCPTID',         type: 'int',   required: true },     // 4
            { field: 'PRCPTID',        type: 'int',   required: false },    // 5
            { field: 'PMTID',          type: 'int',   required: false },    // 6
            { field: 'Dt',             type: 'date',  required: true },     // 7
            { field: 'DocNo',          type: 'text',  required: false },    // 8
            { field: 'Payor', required: true,                               // 9   <<<<<<********
                type: 'enum',
                options: {
                    url:        '/v1/transactantstd/',
                    max:        1,
                    renderItem: tcidReceiptPayorPickerRender,
                    renderDrop: tcidPickerDropRender,
                    compare:    tcidPickerCompare,
                    onNew: function (event) {
                        //console.log('++ New Item: Do not forget to submit it to the server too', event);
                        $.extend(event.item, { FirstName: '', LastName : event.item.text });
                    },
                    onRemove: function(event) {
                        event.onComplete = function() {
                            var f = w2ui.receiptForm;
                            // reset payor field related data when removed
                            f.record.TCID = 0;
                            f.record.Payor = "";

                            // NOTE: have to trigger manually, b'coz we manually change the record,
                            // otherwise it triggers the change event but it won't get change (Object: {})
                            var event = f.trigger({ phase: 'before', target: f.name, type: 'change', event: event }); // event before
                            if (event.cancelled === true) return false;
                            f.trigger($.extend(event, { phase: 'after' })); // event after
                        };
                    }
                },
            },
            { field: 'TCID',           type: 'int64',  required: false },
            { field: 'Amount',         type: 'money',  required: true },
            { field: 'Comment',        type: 'text',   required: false },
            { field: 'OtherPayorName', type: 'text',   required: false },
            { field: 'FLAGS',          type: 'w2int',  required: false },
            { field: 'DID',            type: 'int',    required: false },
            { field: 'LastModTime',    type: 'hidden', required: false },
            { field: 'LastModBy',      type: 'hidden', required: false },
            { field: 'CreateTS',       type: 'hidden', required: false },
            { field: 'CreateBy',       type: 'hidden', required: false },
        ],
        toolbar: {
            items: [
                { id: 'btnNotes', type: 'button', icon: 'fa fa-sticky-note-o' },
                { id: 'bt3', type: 'spacer' },
                { id: 'btnClose', type: 'button', icon: 'fa fa-times' },
            ],
            onClick: function (event) {
                switch(event.target) {
                case 'btnClose':
                    var no_callBack = function() { return false; },
                        yes_callBack = function() {
                            w2ui.toplayout.hide('right',true);
                            w2ui.receiptsGrid.render();
                        };
                    form_dirty_alert(yes_callBack, no_callBack);
                    break;
                }
            },
        },
        onRender: function(event) {
            event.onComplete = function() {
                var f = this,
                    r = f.record,
                    x = getCurrentBusiness(),
                    BID=parseInt(x.value);

                if (r.TCID < 1) { // if TCID is not greater than 0 then return
                    return;
                }

                var record = {};
                getPersonDetailsByTCID(BID, r.TCID)
                .done(function(data) {
                    record = JSON.parse(data).record;
                    var item = {
                        CompanyName: record.CompanyName,
                        IsCompany: record.IsCompany,
                        FirstName: record.FirstName,
                        LastName: record.LastName,
                        MiddleName: record.MiddleName,
                        TCID: record.TCID,
                        recid: 0,
                    };
                    if ($("#receiptForm").find("input[name=Payor]").length > 0) {
                        $("#receiptForm").find("input[name=Payor]").data('selected', [item]).data('w2field').refresh();
                    }
                })
                .fail(function() {
                    f.message("Couldn't get person details for TCID: ", r.TCID);
                    console.log("couldn't get person details for TCID: ", r.TCID);
                });
            };
        },
        onLoad: function(event) {
            event.onComplete = function() {
                var f = this,
                    r = f.record,
                    x = getCurrentBusiness(),
                    BID=parseInt(x.value);

                if (r.TCID < 1) { // if TCID is not greater than 0 then return
                    return;
                }

                var record = {};
                getPersonDetailsByTCID(BID, r.TCID)
                .done(function(data) {
                    record = JSON.parse(data).record;
                    var item = {
                        CompanyName: record.CompanyName,
                        IsCompany: record.IsCompany,
                        FirstName: record.FirstName,
                        LastName: record.LastName,
                        MiddleName: record.MiddleName,
                        TCID: record.TCID,
                        recid: 0,
                    };
                    if ($("#receiptForm").find("input[name=Payor]").length > 0) {
                        $("#receiptForm").find("input[name=Payor]").data('selected', [item]).data('w2field').refresh();
                    }
                })
                .fail(function() {
                    f.message("Couldn't get person details for TCID: ", r.TCID);
                    console.log("couldn't get person details for TCID: ", r.TCID);
                });
            };
        },
        onRefresh: function(event) {
            event.onComplete = function() {
                var f = this,
                    r = f.record,
                    x = getCurrentBusiness(),
                    BID=parseInt(x.value),
                    BUD = getBUDfromBID(BID),
                    header = "Edit Receipt ({0})";

                f.get("PmtTypeName").options.items = buildPaymentTypeSelectList( BUD );
                f.get("PmtTypeName").options.selected = getPaymentType(BUD, r.PMTID);
                f.get("ARID").options.items = app.ReceiptRules[BUD];
                f.get("Payor").options.url = '/v1/transactantstd/'+ BUD;
                // $("#receiptForm").find("input[name=Dt]").prop("disabled", r.RCPTID !== 0);

                formRefreshCallBack(f, "RCPTID", header);

                // ==================================
                // SPECIAL CASE
                // ==================================
                if (r.RCPTID === 0) { // if new record then do not worry about reversed thing
                    $("#"+f.name).find("button[name=reverse]").addClass("hidden");
                    $("#"+f.name).find("button[name=save]").removeClass("hidden");
                    $("#"+f.name).find("button[name=saveadd]").removeClass("hidden");
                    $("#"+f.name).find("button[name=close]").addClass("hidden");
                    $("#"+f.name).find("#FLAGReport").addClass("hidden");

                    // ENABLE ALL INPUTS IF ALL OF THOSE HAVE BEEN DISABLED FOR REVERSED PREVIOUSLY
                    $("#"+f.name).find('input,button').prop("disabled", false);
                    return;
                } else {
                    $("#"+f.name).find("#FLAGReport").removeClass("hidden");
                }
                // this one is a special case, where also have to take care of reverse button
                // FLAG reports
                var flag = r.FLAGS,
                    flagHTML = "";

                // check if it is reversed or not
                if ( (flag & app.rcptFLAGS.RCPTREVERSED) !== 0 ) { // if reversed then
                    flagHTML += "<p style='margin-bottom: 5px;'><strong>{0}</strong> ({1})</p>".format("REVERSED", r.Comment);
                    // reversed indication icon
                    flagHTML += get2XReversalSymbolHTML();
                    // if reversed then do not show reverse, save, saveadd button
                    $("#"+f.name).find("button[name=reverse]").addClass("hidden");
                    $("#"+f.name).find("button[name=save]").addClass("hidden");
                    $("#"+f.name).find("button[name=saveadd]").addClass("hidden");
                    // if reversed then we need to show close button
                    $("#"+f.name).find("button[name=close]").removeClass("hidden");

                    // ********************************************************
                    // IF REVERSED THEN DISABLE ALL INPUTS, BUTTONS EXCEPT close button
                    // ********************************************************
                    $("#"+f.name).find('input,button:not([name=close])').prop("disabled", true);

                } else {
                    // IF NOT REVERSED THEN ONLY SHOW PAID STATUS IN FOOTER
                    // unpaid, partial paid or fully paid
                    if ( (flag | app.rcptFLAGS.RCPTUNALLOCATED) === 0 || (flag & (app.rcptFLAGS.RCPTPARTIALALLOCATED | app.rcptFLAGS.RCPTFULLALLOCATED)) === 0 ) {
                        flagHTML += "<p style='margin-bottom: 5px;'><strong>{0}</strong></p>".format("Unallocated");
                    }
                    else if ( (flag & app.rcptFLAGS.RCPTPARTIALALLOCATED) !== 0 ) {
                        flagHTML += "<p style='margin-bottom: 5px;'><strong>{0}</strong></p>".format("Partially allocated");
                    }
                    else if ( (flag & app.rcptFLAGS.RCPTFULLALLOCATED) !== 0 ) {
                        flagHTML += "<p style='margin-bottom: 5px;'><strong>{0}</strong></p>".format("Fully allocated");
                    }

                    // show save, saveadd, reverse button, hide close button
                    $("#"+f.name).find("button[name=reverse]").removeClass("hidden");
                    $("#"+f.name).find("button[name=save]").removeClass("hidden");
                    $("#"+f.name).find("button[name=saveadd]").removeClass("hidden");
                    $("#"+f.name).find("button[name=close]").addClass("hidden");

                    // ********************************************************
                    // IF not REVERSED THEN ENABLE ALL INPUTS, BUTTONS
                    // ********************************************************
                    $("#"+f.name).find('input,button').prop("disabled", false);
                }

                // finally append
                flagHTML += "<p style='margin-bottom: 5px;'>Last Update: {0} by {1}</p>".format(r.LastModTime, r.LastModBy);
                flagHTML += "<p>CreateTS: {0} by {1}</p>".format(r.CreateTS, r.CreateBy);
                $("#"+f.name).find("#FLAGReport").html(flagHTML);
            };
        },
        onValidate: function (event) {
            if (this.record.Amount === 0.0) {
                event.errors.push({
                    field: this.get('Amount'),
                    error: 'Amount must be something other than $0.00'
                });
            }
            if (this.record.PMTID === 0) {
                event.errors.push({
                    field: this.get('PmtTypeName'),
                    error: 'Please select the payment type'
                });
            }
            if (this.record.ARID.id === 0) {
                event.errors.push({
                    field: this.get('ARID'),
                    error: 'Please select the receipt rule'
                });
            }
        },
        onSubmit: function(target, data) {
            delete data.postData.record.LastModTime;
            delete data.postData.record.LastModBy;
            delete data.postData.record.CreateTS;
            delete data.postData.record.CreateBy;
            // modify form data for server request
            getFormSubmitData(data.postData.record);
        },
        onChange: function(event) {
            event.onComplete = function() {
                var f = this,
                    r = f.record;

                if (event.target == "PmtTypeName") {
                    r.PMTID = event.value_new.id;
                    console.log('record.PMTID set to ' + r.PMTID);
                }

                // formRecDiffer: 1=current record, 2=original record, 3=diff object
                var diff = formRecDiffer(f.record, app.active_form_original, {});
                // if diff == {} then make dirty flag as false, else true
                if ($.isPlainObject(diff) && $.isEmptyObject(diff)) {
                    app.form_is_dirty = false;
                } else {
                    app.form_is_dirty = true;
                }
            };
        },
        actions: {
            close: function() {
                var no_callBack = function() { return false; },
                    yes_callBack = function() {
                        w2ui.toplayout.hide('right',true);
                        w2ui.receiptsGrid.render();
                    };
                form_dirty_alert(yes_callBack, no_callBack);
            },
            saveadd: function() {
                var f = this,
                    grid = w2ui.receiptsGrid,
                    x = getCurrentBusiness(),
                    BID=parseInt(x.value),
                    BUD = getBUDfromBID(BID);

                // clean dirty flag of form
                app.form_is_dirty = false;
                // clear the grid select recid
                app.last.grid_sel_recid  =-1;

                // select none if you're going to add new record
                grid.selectNone();

                f.save({}, function (data) {
                    if (data.status == 'error') {
                        console.log('ERROR: '+ data.message);
                        return;
                    }

                    // JUST RENDER THE GRID ONLY
                    grid.render();

                    // add new empty record and just refresh the form, don't need to do CLEAR form
                    var pmt_options = buildPaymentTypeSelectList(BUD);
                    var ptInit = (pmt_options.length > 0) ? pmt_options[0] : '';
                    var record = getReceiptInitRecord(BID, BUD, ptInit);
                    f.fields[0].options.items = pmt_options;
                    f.fields[1].options.items = app.ReceiptRules[BUD];
                    f.record = record;
                    f.header = "Edit Receipt (new)"; // have to provide header here, otherwise have to call refresh method twice to get this change in form
                    f.url = '/v1/receipt/' + BID + '/0';
                    f.refresh();
                });
            },
            save: function () {
                var f = this,
                    // x = getCurrentBusiness(),
                    // BID=parseInt(x.value),
                    grid = w2ui.receiptsGrid;

                grid.selectNone();

                f.save({}, function (data) {
                    if (data.status == 'error') {
                        console.log('ERROR: '+ data.message);
                        return;
                    }
                    w2ui.toplayout.hide('right',true);
                    grid.render();
                });
            },
            reverse: function() {
                var form = this;

                w2confirm(reverse_confirm_options)
                .yes(function() {
                    var tgrid = w2ui.receiptsGrid;
                    var params = {cmd: 'delete', formname: form.name, RCPTID: form.record.RCPTID };
                    var dat = JSON.stringify(params);
                    // Reverse receipt request
                    $.post(form.url, dat, null, "json")
                    .done(function(data) {
                        if (data.status === "error") {
                            form.error(w2utils.lang(data.message));
                            return;
                        }
                        w2ui.toplayout.hide('right',true);
                        // reversed items should not be deleted!
                        tgrid.render();
                    })
                    .fail(function(/*data*/){
                        form.error("Reverse Receipt failed.");
                        return;
                    });
                })
                .no(function() {
                    return;
                });
            },
        },
        onResize: function(event) {
            event.onComplete = function() {
                // HACK: set the height of right panel of toplayout box div and form's box div
                // this is how w2ui set the content inside box of toplayout panel, and form's main('div.w2ui-form-box')
                var h = w2ui.toplayout.get("right").height;
                $(w2ui.toplayout.get("right").content.box).height(h);
                $(this.box).find("div.w2ui-form-box").height(h);
            };
        },
   });
}
