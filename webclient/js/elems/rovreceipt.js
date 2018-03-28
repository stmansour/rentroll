/*global
    w2ui, $, app, w2confirm, getBUDfromBID, getCurrentBusiness, setToForm,
    console, form_dirty_alert, buildPaymentTypeSelectList, setDateControlsInToolbar,
    addDateNavToToolbar,
    getPersonDetailsByTCID, getPaymentType, formRefreshCallBack, w2utils, reverse_confirm_options,
    getFormSubmitData, w2uiDateControlString, getGridReversalSymbolHTML, get2XReversalSymbolHTML,
    setDefaultFormFieldAsPreviousRecord, formRecDiffer, getCurrentBID, getBUDfromBID,
    exportItemReportCSV, exportItemReportPDF, exportReportCSV, exportReportPDF, getROVReceiptInitRecord, popupReceiptPrintChoice,
    doRcptSave, receiptChoicePrint, loadReceiptChoiceForm
*/
"use strict";
window.getROVReceiptInitRecord = function (BID, BUD, ptInit, previousFormRecord){
    var y = new Date();
    var defaultFormData = {
        recid: 0,
        RCPTID: 0,
        PRCPTID: 0,
        ARID: 0,
        PMTID: 0,
        RAID: 0,
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
        RentableName: '',
        FLAGS: 0,
        LastModBy: 0,
        CreateBy: 0
    };

    // if it called after 'save and add another' action there previous form record is passed as Object
    // else it is null
    if ( previousFormRecord ) {
        defaultFormData = setDefaultFormFieldAsPreviousRecord(
            [ 'DocNo', 'Payor', 'Amount', 'OtherPayorName', 'Comment', 'RAID'], // Fields to Reset
            defaultFormData,
            previousFormRecord
        );
    }

    return defaultFormData;
};


window.buildROVReceiptElements = function () {
    //------------------------------------------------------------------------
    //          receiptsGrid
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'receiptsGrid',
        url: '/v1/receipts',
        multiSelect: false,
        postData: {searchDtStart: app.D1, searchDtStop: app.D2, client: app.client},
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
                        if ( (record.FLAGS & app.rcptFLAGS.REVERSED) !== 0 ) { // if reversed then
                            return getGridReversalSymbolHTML();
                        }
                        return '';
                    },
            },
            {field: 'RCPTID',        caption: 'Receipt ID',      size: '80px',  hidden: false, sortable: true, style: 'text-align: right'},
            {field: 'Dt',            caption: 'Date',            size: '80px',  hidden: false, sortable: true, style: 'text-align: right'},
            {field: 'ARID',          caption: 'ARID',            size: '150px', hidden: true,  sortable: false},
            {field: 'DID',           caption: 'DID',             size: '150px', hidden: true,  sortable: false},
            {field: 'AcctRule',      caption: 'Account Rule',    size: '150px', hidden: true,  sortable: true},
            {field: 'Amount',        caption: 'Amount',          size: '100px', hidden: false, sortable: true, render: 'money', style: 'text-align: right'},
            {field: 'BID',           caption: 'BUD',             size: '40px',  hidden: true,  sortable: false},
            {field: 'TCID',          caption: 'TCID',            size: '40px',  hidden: true,  sortable: false},
            {field: 'PMTID',         caption: 'PMTID',                          hidden: true,  sortable: false},
            {field: 'PmtTypeName',   caption: 'Payment Type',    size: '100px', hidden: false, sortable: true},
            {field: 'DocNo',         caption: 'Document Number', size: '150px', hidden: false, sortable: true, style: 'text-align: right'},
            {field: 'OtherPayorName',caption: 'Payor',           size: '150px', hidden: false, sortable: true},
            {field: 'RentableName',  caption: 'Resident Address',size: '150px', hidden: false, sortable: true},
            {field: 'Comment',       caption: 'Comment',         size: '150px', hidden: false, sortable: true},
        ],
        toolbar: {
            onClick: function (event) {
                switch(event.target) {
                case 'csvexport':
                    exportReportCSV("RPTrcptlist", app.D1, app.D2);
                    break;
                case 'printreport':
                    exportReportPDF("RPTrcptlist", app.D1, app.D2);
                    break;
                }
            },
        },
        searches : [
            { field: 'Amount', caption: 'Amount', type: 'string' },
            // { field: 'DocNo', caption: 'Document Number', type: 'string' },
            { field: 'OtherPayorName', caption: 'Payor', type: 'string' },
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
                        var BID = getCurrentBID();
                            // var BUD = getBUDfromBID(BID);

                        app.last.grid_sel_recid = parseInt(recid);
                        grid.select(app.last.grid_sel_recid);// keep highlighting current row in any case

                        var rec = grid.get(recid);
                        var f = w2ui.receiptForm;
                        var j = f.get('ERentableName',true); // index of the enumerated RentableName field
                        f.fields[j].options.url = '/v1/rentablestd/' + BID;
                        f.postData = {searchDtStart: app.D1, searchDtStop: app.D2, client: app.client};
                        setToForm('receiptForm', '/v1/receipt/' + BID + '/' + rec.RCPTID, 400, true);
                    };
                form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);  // warn user if form content has been changed
            };
        },
        onRequest: function(/*event*/) {
            w2ui.receiptsGrid.postData = {searchDtStart: app.D1, searchDtStop: app.D2, client: app.client};
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
                    w2ui.receiptsGrid.postData = {searchDtStart: app.D1, searchDtStop: app.D2, client: app.client};
                }
            };
        },
        onAdd: function (/*event*/) {
            var BID = getCurrentBID();
            var BUD = getBUDfromBID(BID);

            // reset it
            app.last.grid_sel_recid = -1;
            // grid.selectNone();

            var f = w2ui.receiptForm;
            var j = f.get('ERentableName',true); // index of the enumerated RentableName field
            f.fields[j].options.url = '/v1/rentablestd/' + BID;
            var pmt_options = buildPaymentTypeSelectList(BUD);
            var ptInit = (pmt_options.length > 0) ? pmt_options[0] : '';
            f.record = getROVReceiptInitRecord(BID, BUD, ptInit, null);
            f.header = "Edit Receipt (new)";
            f.postData = {client: app.client};
            f.refresh();
            setToForm('receiptForm', '/v1/receipt/' + BID + '/0', 400);
        },
    });

    addDateNavToToolbar('receipts');
    w2ui.receiptsGrid.toolbar.add([
        { type: 'spacer',},
        { type: 'button', id: 'csvexport', icon: 'fas fa-table', tooltip: 'export to CSV' },
        { type: 'button', id: 'printreport', icon: 'far fa-file-pdf', tooltip: 'export to PDF' },
        ]);

    //------------------------------------------------------------------------
    //          receiptForm
    //------------------------------------------------------------------------
    $().w2form({
        name: 'receiptForm',
        style: 'border: 0px; background-color: transparent;',
        header: 'Receipt Detail',
        url: '/v1/receipt',
        formURL: '/webclient/html/formrovrcpt.html',
        fields: [
            { field: 'PmtTypeName',type: 'list', required: true, options: { items: [], selected: {} }, html: { caption: "BUD", page: 0 } },
            // { field: 'ARID',    type: 'list', required: true, options:  {items: app.ReceiptRules} },
            { field: 'recid',      type: 'int',  required: false },
            { field: "BID",        type: 'int',  required: false, html: { caption: "BID", page: 0, column: 0 } },
            { field: "BUD",        type: 'list', required: true, options: { items: app.businesses, maxDropHeight: 350 }, html: { caption: "BUD", page: 0, column: 0 } },
            { field: 'RCPTID',     type: 'int',  required: true },
            { field: 'PRCPTID',    type: 'int',  required: false },
            { field: 'PMTID',      type: 'int',  required: false },
            { field: 'Dt',         type: 'date', required: true },
            { field: 'DocNo',      type: 'text', required: true },
            { field: 'ERentableName',
                type: 'combo',
                options: {
                    url:           '/v1/rentablestd/',
                    max:           1,
                    cacheMax:      50,
                    maxDropHeight: 350,
                    maxDropWidth:  300,
                    compare:       ridRentableCompare,
                    recId:         'recid',
                    recText:       'RentableName',
                },
            },
            { field: 'ARID',           type: 'hidden', required: false },
            { field: 'Payor',          type: 'hidden', required: false },
            { field: 'TCID',           type: 'hidden', required: false },
            { field: 'RAID',           type: 'hidden', required: false },
            { field: 'Amount',         type: 'money',  required: true },
            { field: 'Comment',        type: 'text',   required: false },
            { field: 'OtherPayorName', type: 'text',   required: false },
            { field: 'FLAGS',          type: 'w2int',  required: false },
            { field: 'DID',            type: 'hidden', required: false },
            { field: 'LastModTime',    type: 'hidden', required: false },
            { field: 'LastModBy',      type: 'hidden', required: false },
            { field: 'LastModByUser',  type: 'hidden', required: false },
            { field: 'CreateTS',       type: 'hidden', required: false },
            { field: 'CreateBy',       type: 'hidden', required: false },
            { field: 'CreateByUser',   type: 'hidden', required: false },
            { field: 'RentableName',   type: 'hidden', required: false },
        ],
        toolbar: {
            items: [
                // { id: 'btnNotes',    type: 'button', icon: 'far fa-sticky-note' },
                { id: 'print',       type: 'button', icon: 'fas fa-print',        tooltip: 'print receipt' },
                { id: 'bt3',         type: 'spacer' },
                { id: 'btnClose',    type: 'button', icon: 'fas fa-times' },
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
                case "print":
                    if (w2ui.receiptForm.record.RCPTID === 0) {
                        return;
                    }
                    popupReceiptPrintChoice();
                    break;
                }
            },
        },
        onRender: function(event) { // when form is loaded first time in toplayout right panel
            event.onComplete = function() {
                var f = this;
                var r = f.record;

                // Create the RentableName element if rentable name exists...
                if (r.RentableName.length > 0) {
                    r.ERentableName = r.RentableName;
                }
            };
        },
        onRefresh: function(event) {
            //w2ui.ridRentablePicker.fields[0].options.url = '/v1/rentablestd/' + app.ridRentablePicker.BID;
            event.onComplete = function() {
                var f      = this;
                var r      = f.record;
                var x      = getCurrentBusiness();
                var BID    = parseInt(x.value);
                var BUD    = getBUDfromBID(BID);
                var header = "Edit Receipt ({0})";

                if (f.record.RCPTID === 0) {
                    header = "Edit Receipt (new)";
                }

                if (r.RentableName.length > 0) {
                    r.ERentableName = r.RentableName;
                }

                f.get("PmtTypeName").options.items = buildPaymentTypeSelectList( BUD );
                f.get("PmtTypeName").options.selected = getPaymentType(BUD, r.PMTID);

                formRefreshCallBack(f, "RCPTID", header);

                // ==================================
                // SPECIAL CASE
                // ==================================
                if (r.RCPTID === 0) { // if new record then do not worry about reversed thing
                    $(f.box).find("button[name=reverse]").addClass("hidden");
                    $(f.box).find("button[name=save]").removeClass("hidden");
                    $(f.box).find("button[name=saveadd]").removeClass("hidden");
                    $(f.box).find("button[name=close]").addClass("hidden");
                    $(f.box).find("#FLAGReport").addClass("hidden");

                    // ENABLE ALL INPUTS IF ALL OF THOSE HAVE BEEN DISABLED FOR REVERSED PREVIOUSLY
                    $(f.box).find('input,button').not('input[name=BUD]').prop("disabled", false);
                    return;
                } else {
                    $(f.box).find("#FLAGReport").removeClass("hidden");
                }
                // this one is a special case, where also have to take care of reverse button
                // FLAG reports
                var flag = r.FLAGS,
                    flagHTML = "";

                // check if it is reversed or not
                if ( (flag & app.rcptFLAGS.REVERSED) !== 0 ) { // if reversed then
                    flagHTML += "<p><strong>{0}</strong> ({1})</p>".format("REVERSED", r.Comment);
                    // reversed indication icon
                    flagHTML += get2XReversalSymbolHTML();
                    // if reversed then do not show reverse, save, saveadd button
                    $(f.box).find("button[name=reverse]").addClass("hidden");
                    $(f.box).find("button[name=save]").addClass("hidden");
                    $(f.box).find("button[name=saveadd]").addClass("hidden");
                    // if reversed then we need to show close button
                    $(f.box).find("button[name=close]").removeClass("hidden");

                    // ********************************************************
                    // IF REVERSED THEN DISABLE ALL INPUTS, BUTTONS EXCEPT close button
                    // ********************************************************
                    $(f.box).find('input,button:not([name=close])').prop("disabled", true);

                } else {
                    // IF NOT REVERSED THEN ONLY SHOW PAID STATUS IN FOOTER
                    // unpaid, partial paid or fully paid
                    if ( (flag | app.rcptFLAGS.UNALLOCATED) === 0 || (flag & (app.rcptFLAGS.PARTIALALLOCATED | app.rcptFLAGS.FULLYALLOCATED)) === 0 ) {
                        flagHTML += "<p><strong>{0}</strong></p>".format("Unallocated");
                    }
                    else if ( (flag & app.rcptFLAGS.PARTIALALLOCATED) !== 0 ) {
                        flagHTML += "<p><strong>{0}</strong></p>".format("Partially allocated");
                    }
                    else if ( (flag & app.rcptFLAGS.FULLYALLOCATED) !== 0 ) {
                        flagHTML += "<p><strong>{0}</strong></p>".format("Fully allocated");
                    }

                    // show save, saveadd, reverse button, hide close button
                    $(f.box).find("button[name=reverse]").removeClass("hidden");
                    $(f.box).find("button[name=save]").removeClass("hidden");
                    // $(f.box).find("button[name=saveadd]").removeClass("hidden");
                    $(f.box).find("button[name=close]").addClass("hidden");

                    // ********************************************************
                    // IF not REVERSED THEN ENABLE ALL INPUTS, BUTTONS
                    // ********************************************************
                    $(f.box).find('input,button').not('input[name=BUD]').prop("disabled", false);
                }

                // finally append
                flagHTML += "<p>Last Update: {0} by {1}</p>".format(r.LastModTime, r.LastModByUser);
                flagHTML += "<p>Created: {0} by {1}</p>".format(r.CreateTS, r.CreateByUser);
                $(f.box).find("#FLAGReport").html(flagHTML);
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
            w2ui.receiptForm.postData = {client: app.client};
            getFormSubmitData(data.postData.record);
        },
        onChange: function(event) {
            event.onComplete = function() {
                var f = this,
                    r = f.record;
                switch (event.target) {
                case "PmtTypeName":
                    r.PMTID = event.value_new.id;
                    break;
                case "ERentableName":
                    // if ERentableName changed then also update RentableName field
                    r.RentableName = r.ERentableName;
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
            saveprint: function() {
                doRcptSave(this,true);
            },
            save: function () {
                doRcptSave(this,false);
            },
            reverse: function() {
                var form = this;

                w2confirm(reverse_confirm_options)
                .yes(function() {
                    var tgrid = w2ui.receiptsGrid;
                    var params = {cmd: 'delete', formname: form.name, RCPTID: form.record.RCPTID, client: app.client };
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
                // so that footer content get placed at correct position
                var h = w2ui.toplayout.get("right").height;
                $(w2ui.toplayout.get("right").content.box).height(h);
                $(this.box).find("div.w2ui-form-box").height(h);
            };
        },

   });
};

window.doRcptSave = function (f, prnt) {
    var r = f.record;
    var grid = w2ui.receiptsGrid;

    if (typeof r.RAID === "string") {
        r.RAID = parseInt(r.RAID);
    }
    grid.selectNone();

    f.postData = {client: app.client, RentableName: f.record.RentableName};
    f.save(null, function (data) {
        if (data.status == 'error') {
            console.log('ERROR: '+ data.message);
            return;
        }

        // save the id in record, in case form record is new
        f.record.RCPTID = data.recid;

        w2ui.toplayout.hide('right',true);
        grid.render();
        if (prnt) {
            popupReceiptPrintChoice();
        }
    });
};

//--------------------------------------------------------------------------------
// This contains the definition part of receiptChoiceForm, loads the form on call
//--------------------------------------------------------------------------------
window.loadReceiptChoiceForm = function () {
    $().w2form({
        name: 'receiptChoiceForm',
        style: 'border: 0px; background-color: transparent;',
        focus: -1,
        formHTML:
            '<div class="w2ui-page page-0">'+
            '    <div class="w2ui-field">'+
            '        <label>Format:</label>'+
            '        <div>'+
            '           <input name="report_format" type="list" />'+
            '        </div>'+
            '    </div>'+
            '    <div class="w2ui-field">'+
            '        <label>Type:</label>'+
            '        <div>'+
            '           <label>'+
            '               <input name="report_type" type="radio" value="permanent_resident" /> Permanent Resident'+
            '           </label>'+
            '           </br>'+
            '           <label>'+
            '               <input name="report_type" type="radio" value="hotel" /> Hotel'+
            '           </label>'+
            '        </div>'+
            '    </div>'+
            '</div>'+
            '<div class="w2ui-buttons">'+
            '    <button class="w2ui-btn" name="print" >Print</button>'+
            '    <button class="w2ui-btn" name="close">Close</button>'+
            '</div>',
        fields: [
            { field: 'report_format', type: 'list' , options: { items: ['PDF', 'CSV'] } },
            { field: 'report_type'  , type: 'radio' },
        ],
        record: {
            report_format : "PDF",
            report_type   : "permanent_resident",
        },
        actions: {
            close: function () {
                w2popup.close();
            },
            print: function() {
                receiptChoicePrint();
            },
        }
    });
};

//--------------------------------------------------------------------------------
// Pops up dialog to get print choice for the receipt (permanent resident / hotel)
//--------------------------------------------------------------------------------
window.popupReceiptPrintChoice = function () {

    // if receipt form is not loaded then load it first
    if (!w2ui.receiptChoiceForm) {
        loadReceiptChoiceForm();
    }

    w2popup.open({
        title     : 'Print Receipt',
        body      : '<div id="form" style="width: 100%; height: 100%;"></div>',
        style     : 'padding: 15px 0px 0px 0px',
        width     : 360,
        height    : 220,
        overflow  : 'hidden',
        color     : '#333',
        speed     : '0.3',
        opacity   : '0.5',
        modal     : true,
        showClose : true,
        onOpen    : function(event) {
            event.onComplete = function() {
                // specifying an onOpen handler instead is equivalent to specifying an onBeforeOpen handler, which would make this code execute too early and hence not deliver.
                $('#w2ui-popup #form').w2render('receiptChoiceForm');
            };
        }
    });
};

//--------------------------------------------------------------------------------------------
// Sends the request to print receipt based upon a choice by user (permanent resident / hotel)
//--------------------------------------------------------------------------------------------
window.receiptChoicePrint = function () {
    // decide function call based on format first
    var exportFormatFunc;
    switch(w2ui.receiptChoiceForm.record.report_format.id) {
    case "PDF":
        exportFormatFunc = exportItemReportPDF;
        break;
    case "CSV":
        exportFormatFunc = exportItemReportCSV;
        break;
    default:
        alert("Invalid export format for receipt print");
        return false;
    }

    // choose type of report based on user selection
    switch(w2ui.receiptChoiceForm.record.report_type){
    case "permanent_resident":
        exportFormatFunc("RPTrcpt", w2ui.receiptForm.record.RCPTID, app.D1, app.D2);
        break;
    case "hotel":
        exportFormatFunc("RPTrcpthotel", w2ui.receiptForm.record.RCPTID, app.D1, app.D2);
        break;
    }

    // close the dialog after 500ms
    setTimeout(function() {
        w2popup.close();
        // TODO(Sudip): should we close the right panel after save/print succeed?
    }, 500);
};

