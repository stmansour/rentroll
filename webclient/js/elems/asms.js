"use strict";
function getAsmsInitRecord(BID, BUD){
    var y = new Date();
    var y1 = new Date(new Date().setFullYear(new Date().getFullYear() + 1));
    return {
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
}

function renderReversalIcon(record /*, index, col_index*/) {
    if (typeof record === "undefined") {
        return;
    }
    if ( (record.FLAGS & app.asmFLAGS.ASMREVERSED) !== 0 ) { // if reversed then
        return '<i class="fa fa-exclamation-triangle" title="reversed" aria-hidden="true" style="color: #FFA500;"></i>';
    }
    return '';
}

function buildAssessmentElements() {
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
                    render: renderReversalIcon,
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
                        $.get('/v1/uival/' + x.value + '/app.AssessmentRules' )
                        .done( function(data) {
                            if (typeof data == 'string') {  // it's weird, a successful data add gets parsed as an object, an error message does not
                                app.AssessmentRules = JSON.parse(data);
                                w2ui[form].get('ARID').options.items = app.AssessmentRules[Bud];
                                w2ui[form].refresh();

                                setToForm(form, myurl, 400, true);
                            }
                            if (data.status != 'success') {
                                w2ui.asmsGrid.message(data.message);
                            }
                        })
                        .fail( function() {
                            console.log('Error getting /v1/uival/' + x.value + '/app.AssessmentRules');
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
                    $.get('/v1/uival/' + x.value + '/app.AssessmentRules' )
                    .done( function(data) {
                        if (typeof data == 'string') {  // it's weird, a successful data add gets parsed as an object, an error message does not
                            app.AssessmentRules = JSON.parse(data);

                            // Insert an empty record...
                            var BID=parseInt(x.value);
                            var BUD = getBUDfromBID(BID);
                            app.ridRentablePicker.BID = BID; // needed by typedown

                            var record = getAsmsInitRecord(BID, BUD);
                            // w2ui.asmEpochForm.fields[5].options.url = '/v1/rentablestd/' + app.ridRentablePicker.BID;
                            w2ui.asmEpochForm.fields[0].options.items = app.AssessmentRules[BUD];
                            w2ui.asmEpochForm.record = record;
                            w2ui.asmEpochForm.refresh();

                            setToForm('asmEpochForm', '/v1/asm/' + BID + '/0', 400);
                        }
                        if (data.status != 'success') {
                            w2ui.asmEpochForm.message(data.message);
                        }
                    })
                    .fail( function() {
                        console.log('Error getting /v1/uival/'+x.value+'/app.AssessmentRules');
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

    //---------------------------------------------------------------------------------
    //          asmEpochForm  -  assessment epoch - this is for recurring assessments
    //---------------------------------------------------------------------------------
    $().w2form({
        name: 'asmEpochForm',
        style: 'border: 0px; background-color: transparent;',
        header: app.sAssessment + ' Detail',
        url: '/v1/asm',
        formURL: '/webclient/html/formasmepoch.html',
        fields: [
            { field: 'ARID',          type: 'list',     required: true, options: { items: app.AssessmentRules }},
            { field: 'recid',         type: 'int',      required: false },
            { field: 'ASMID',         type: 'int',      required: false },
            { field: 'BID',           type: 'int',     required: true },
            { field: 'BUD',           type: 'list',     required: true, options: {items: app.businesses} },
            { field: 'PASMID',        type: 'w2int',    required: false },
            { field: 'Rentable',      type: 'text',     required: false },
            { field: 'InvoiceNo',     type: 'int',      required: false },
            { field: 'RID',           type: 'int',      required: false },
            { field: 'ATypeLID',      type: 'int',      required: false },
            { field: 'RAID',          type: 'int',      required: true },
            { field: 'Amount',        type: 'money',    required: true },
            { field: 'Start',         type: 'date',     required: true },
            { field: 'Stop',          type: 'date',     required: true },
            { field: 'RentCycle',     type: 'list',     required: true, options: {items: app.cycleFreq}, },
            { field: 'ProrationCycle',type: 'list',     required: true, options: {items: app.cycleFreq}, },
            { field: 'Comment',       type: 'text',     required: false },
            { field: 'LastModTime',   type: 'hidden',   required: false },
            { field: 'LastModBy',     type: 'hidden',   required: false },
            { field: 'CreateTS',   type: 'hidden',   required: false },
            { field: 'CreateBy',     type: 'hidden',   required: false },
            { field: 'ExpandPastInst',type: 'checkbox', required: false },
            { field: 'FLAGS',         type: 'w2int',    required: false },
            { field: 'Mode',          type: 'w2int',    required: false },
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
                            w2ui.asmsGrid.render();
                        };
                    form_dirty_alert(yes_callBack, no_callBack);
                    break;
                }
            },
        },
        onValidate: function (event) {
            if (this.record.ARID.id === 0) {
                event.errors.push({
                    field: this.get('ARID'),
                    error: 'The Account Rule needs to be set'
                });
            }
            if (this.record.Amount < 0.01) {
                event.errors.push({
                    field: this.get('Amount'),
                    error: 'Amount must be at least $0.01'
                });
            }
        },
        actions: {
            close: function() {
                var no_callBack = function() { return false; },
                    yes_callBack = function() {
                        w2ui.toplayout.hide('right',true);
                        w2ui.asmsGrid.render();
                    };
                form_dirty_alert(yes_callBack, no_callBack);
            },
            saveadd: function() {
                var f = this,
                    grid = w2ui.asmsGrid,
                    x = getCurrentBusiness(),
                    BID=parseInt(x.value),
                    BUD = getBUDfromBID(BID);

                // clean dirty flag of form
                app.form_is_dirty = false;
                // clear the grid select recid
                app.last.grid_sel_recid  =-1;

                // unselect the record
                grid.selectNone();

                // first save the record
                f.save({}, function (data) {
                    if (data.status == 'error') {
                        console.log('ERROR: '+ data.message);
                        return;
                    }
                    // render the grid only
                    grid.render();

                    $.get('/v1/uival/' + BID + '/app.AssessmentRules' )
                    .done( function(data) {
                        if (typeof data == 'string') {  // it's weird, a successful data add gets parsed as an object, an error message does not
                            app.AssessmentRules = JSON.parse(data);
                            app.ridRentablePicker.BID = BID; // needed by typedown

                            var record = getAsmsInitRecord(BID, BUD);

                            // f.fields[5].options.url = '/v1/rentablestd/' + app.ridRentablePicker.BID;
                            f.fields[0].options.items = app.AssessmentRules[BUD];
                            f.record = record;
                            f.header = "Edit Assessment (new)"; // have to provide header here, otherwise have to call refresh method twice to get this change in form
                            f.url  = "/v1/asm/" + BID + "/0";
                            f.refresh();
                        }
                        if (data.status != 'success') {
                            f.message(data.message);
                        }
                    })
                    .fail( function() {
                        console.log('Error getting /v1/uival/'+x.value+'/app.AssessmentRules');
                     }); //get assessment UI val done

                });
            },
            save: function () {
                var x = getCurrentBusiness(),
                    grid = w2ui.asmsGrid;

                grid.selectNone();
                w2ui.asmEpochForm.url = '/v1/asm/' + x.value + '/' + w2ui.asmEpochForm.record.ASMID;

                this.save({}, function (data) {
                    if (data.status == 'error') {
                        console.log('ERROR: '+ data.message);
                        return;
                    }
                    w2ui.toplayout.hide('right',true);
                    grid.render();
                });
            },
            reverse: function() {
                popupAsmRevMode(2, w2ui.asmEpochForm);
            },
        },
        onRefresh: function(event) {
            event.onComplete = function() {
                var f = this,
                    r = f.record,
                    header = "Edit Assessment ({0})";

                formRefreshCallBack(f, "ASMID", header);

                // ===========================
                // SPECIAL CASE
                // ===========================

                if (r.ASMID === 0) { // if new record then do not worry about reversed thing
                    $("#"+f.name).find("button[name=reverse]").addClass("hidden");
                    $("#"+f.name).find("button[name=save]").removeClass("hidden");
                    $("#"+f.name).find("button[name=saveadd]").removeClass("hidden");
                    $("#"+f.name).find("button[name=close]").addClass("hidden");
                    $("#"+f.name).find("#FLAGReport").addClass("hidden");
                    $("#"+f.name).find("#AssessmentInfo").addClass("hidden");

                    // ENABLE ALL INPUTS IF ALL OF THOSE HAVE BEEN DISABLED FOR REVERSED PREVIOUSLY
                    $("#"+f.name).find('input,button').prop("disabled", false);

                    return;
                } else {
                    $("#"+f.name).find("#FLAGReport").removeClass("hidden");
                    $("#"+f.name).find("#AssessmentInfo").removeClass("hidden");
                }

                // Assessment Info at the top of form in white box
                var info = '<p><i class="fa fa-refresh" style="margin-right: 5px;"></i> Repeating Assessment Series Definition</p>'.format(r.ASMID);
                $("#"+f.name).find("#AssessmentInfo").html(info);

                // FLAG reports
                var flag = r.FLAGS,
                    flagHTML = "";

                // check if it is reversed or not
                if ( (flag & app.asmFLAGS.ASMREVERSED) !== 0 ) { // if reversed then
                    flagHTML += "<p style='margin-bottom: 5px;'><strong>{0}</strong> ({1})</p>".format("REVERSED", r.Comment);
                    // reversed indication icon
                    flagHTML += "<div class='reverseIconContainer'><i class='fa fa-exclamation-triangle fa-2x reverseIcon' aria-hidden='true'></i></div>";
                    // if reversed then do not show reverse, save button
                    $("#"+f.name).find("button[name=reverse]").addClass("hidden");
                    $("#"+f.name).find("button[name=save]").addClass("hidden");
                    $("#"+f.name).find("button[name=saveadd]").addClass("hidden");
                    // if reversed then we need to show close button at the bottom of form
                    $("#"+f.name).find("button[name=close]").removeClass("hidden");

                    // ****************************************
                    // IF REVERSED THEN DISABLE ALL INPUTS, BUTTONS  EXCEPT close button
                    // ****************************************
                    $("#"+f.name).find('input,button:not([name=close])').prop("disabled", true);

                } else {
                    // IF NOT REVERSED THEN ONLY SHOW PAID STATUS IN FOOTER
                    // unpaid, partial paid or fully paid
                    if ( (flag | app.asmFLAGS.ASMUNPAID) === 0 || (flag & (app.asmFLAGS.ASMPARTIALPAID | app.asmFLAGS.ASMFULLPAID)) === 0 ) {
                        flagHTML += "<p style='margin-bottom: 5px;'><strong>{0}</strong></p>".format("Unpaid");
                    }
                    else if ( (flag & app.asmFLAGS.ASMPARTIALPAID) !== 0 ) {
                        flagHTML += "<p style='margin-bottom: 5px;'><strong>{0}</strong></p>".format("Partially paid");
                    }
                    else if ( (flag & app.asmFLAGS.ASMFULLPAID) !== 0 ) {
                        flagHTML += "<p style='margin-bottom: 5px;'><strong>{0}</strong></p>".format("Fully paid");
                    }

                    // show reverse, save button, hide close button
                    $("#"+f.name).find("button[name=reverse]").removeClass("hidden");
                    $("#"+f.name).find("button[name=save]").removeClass("hidden");
                    $("#"+f.name).find("button[name=saveadd]").removeClass("hidden");
                    $("#"+f.name).find("button[name=close]").addClass("hidden");

                    // ****************************************
                    // IF not REVERSED THEN ENABLE ALL INPUTS
                    // ****************************************
                    $("#"+f.name).find('input,button').prop("disabled", false);
                }

                // finally append
                flagHTML += "<p style='margin-bottom: 5px;'>Last Update: {0} by {1}</p>".format(r.LastModTime, r.LastModBy);
                flagHTML += "<p>CreateTS: {0} by {1}</p>".format(r.CreateTS, r.CreateBy);
                $("#"+f.name).find("#FLAGReport").html(flagHTML);
            };
        },
        onChange: function(event) {
            event.onComplete = function() {
                var f = this,
                    r = f.record;

                if (event.target == "Start") {
                    var x = document.getElementsByName('ExpandPastInst')[0];
                    if (r.RentCycle.text != "Norecur") {
                        var DtStart = dateFromString(event.value_new);
                        var y = new Date();
                        x.checked = (DtStart <  y);
                    } else {
                        // if Start date has been changed, in rentcycle with norecur mode
                        // then we need to set stop date same value of start date
                        r.Stop = r.Start;
                        x.checked = false;
                    }
                }
                if (event.target == "RentCycle") {
                    if (event.value_new.text == "Norecur") {
                        r.RentCycle = event.value_new;
                        r.ProrationCycle = "Norecur";
                        r.Stop = r.Start;
                        // disable stop date control
                        $("#"+f.name).find("input[name=Stop]").prop("disabled", true);
                    } else {
                        // enable stop date control
                        $("#"+f.name).find("input[name=Stop]").prop("disabled", false);
                    }
                }

                // copy original record temporary and reset it back after refresh event
                var temp = $.extend(true, {}, app.active_form_original);
                // finally refresh the form
                f.refresh();
                // now reset app original form record
                app.active_form_original = $.extend(true, {}, temp);

                // formRecDiffer: 1=current record, 2=original record, 3=diff object
                var diff = formRecDiffer(r, app.active_form_original, {});
                // if diff == {} then make dirty flag as false, else true
                if ($.isPlainObject(diff) && $.isEmptyObject(diff)) {
                    app.form_is_dirty = false;
                } else {
                    app.form_is_dirty = true;
                }
            };
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
        onSubmit: function(target, data) {
            delete data.postData.record.LastModTime;
            delete data.postData.record.LastModBy;
            delete data.postData.record.CreateTS;
            delete data.postData.record.CreateBy;
            // modify form data for server request
            getFormSubmitData(data.postData.record);
            console.log(data.postData.record);
            app.cycleFreq.forEach(function(item, index) {
                if (item == data.postData.record.RentCycle) {
                    data.postData.record.RentCycle = index;
                }
                if (item == data.postData.record.ProrationCycle) {
                    data.postData.record.ProrationCycle = index;
                }
            });
        },
    });

    //----------------------------------------------------------------------------------------------
    //          asmInstForm  -  assessment instance - instances only, not for recurring assessments
    //----------------------------------------------------------------------------------------------
    $().w2form({
        name: 'asmInstForm',
        style: 'border: 0px; background-color: transparent;',
        header: app.sAssessment + ' Detail',
        url: '/v1/asm',
        formURL: '/webclient/html/formasminst.html',
        fields: [
            { field: 'ARID',          type: 'list',   required: true, options: { items: app.AssessmentRules } },
            { field: 'recid',         type: 'int',    required: false },
            { field: 'ASMID',         type: 'int',    required: false },
            { field: 'BUD',           type: 'list',   options:  {items: app.businesses}, required: false },
            { field: 'BID',         type: 'int',    required: true },
            { field: 'PASMID',        type: 'w2int',  required: false },
            { field: 'RID',           type: 'w2int',  hidden:   true },
            { field: 'Rentable',      type: 'text',   required: false },
            { field: 'RAID',          type: 'w2int',  required: false },
            { field: 'Amount',        type: 'money',  required: true },
            { field: 'Start',         type: 'date',   required: true },
            { field: 'Stop',          type: 'date',   required: true },
            { field: 'RentCycle',     type: 'list',   options:  {items: app.cycleFreq}, required: true },
            { field: 'ProrationCycle',type: 'list',   options:  {items: app.cycleFreq}, required: true },
            { field: 'Comment',       type: 'text',   required: false },
            { field: 'LastModTime',   type: 'hidden', required: false },
            { field: 'LastModBy',     type: 'hidden', required: false },
            { field: 'CreateTS',   type: 'hidden', required: false },
            { field: 'CreateBy',     type: 'hidden', required: false },
            { field: 'FLAGS',         type: 'w2int',  required: false },
            { field: 'Mode',          type: 'w2int',  required: false },
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
                            w2ui.asmsGrid.render();
                        };
                    form_dirty_alert(yes_callBack, no_callBack);
                    break;
                }
            },
        },
        onError: function(event) {
            console.log('onError handler called. event - '+event);
        },
        actions: {
            close: function() {
                var no_callBack = function() { return false; },
                    yes_callBack = function() {
                        w2ui.toplayout.hide('right',true);
                        w2ui.asmsGrid.render();
                    };
                form_dirty_alert(yes_callBack, no_callBack);
            },
            saveadd: function() {
                var f = this,
                    epochForm = w2ui.asmEpochForm,
                    grid = w2ui.asmsGrid,
                    x = getCurrentBusiness(),
                    BID=parseInt(x.value),
                    BUD = getBUDfromBID(BID);

                // clean dirty flag of form
                app.form_is_dirty = false;
                // clear the grid select recid
                app.last.grid_sel_recid  =-1;

                // unselect the record
                grid.selectNone();

                // first save the record
                f.save({}, function (data) {
                    if (data.status == 'error') {
                        console.log('ERROR: '+ data.message);
                        return;
                    }

                    // render the grid only
                    grid.render();

                    $.get('/v1/uival/' + BID + '/app.AssessmentRules' )
                    .done( function(data) {
                        if (typeof data == 'string') {  // it's weird, a successful data add gets parsed as an object, an error message does not
                            app.AssessmentRules = JSON.parse(data);
                            app.ridRentablePicker.BID = BID; // needed by typedown

                            var record = getAsmsInitRecord(BID, BUD);

                            // epochForm.fields[5].options.url = '/v1/rentablestd/' + app.ridRentablePicker.BID;
                            epochForm.fields[0].options.items = app.AssessmentRules[BUD];
                            epochForm.record = record;
                            // epochForm.header = "Edit Assessment (new)"; // have to provide header here, otherwise have to call refresh method twice to get this change in form
                            // f.refresh();
                            setToForm(epochForm.name, '/v1/asm/' + BID + '/0', 400);
                        }
                        if (data.status != 'success') {
                            f.message(data.message);
                        }
                    })
                    .fail( function() {
                        console.log('Error getting /v1/uival/'+x.value+'/app.AssessmentRules');
                     }); //get assessment UI val done

                });
            },
            save: function () {
                var f = this,
                    grid = w2ui.asmsGrid;

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
                popupAsmRevMode(0,w2ui.asmInstForm);
            },
        },
        onLoad: function(event) {
            event.onComplete = function() {
                var f = this;
                var r = f.record;
                var x = getCurrentBusiness();

                var params = {"cmd":"get","recid":0,"name":"asmInstForm"};
                var dat = JSON.stringify(params);
                $.post('/v1/asm/' + x.value + '/' + r.PASMID, dat)
                .done( function(data) {
                    if (data.status != 'success') {
                        f.message(data.message);
                        f.pasmStart = "";
                        f.pasmStop = "";
                    }
                    data = JSON.parse(data);
                    // get parent assessment dates and store it in form
                    f.pasmStart = data.record.Start;
                    f.pasmStop = data.record.Stop;
                })
                .fail( function() {
                    console.log('Error getting /v1/asm/' + x.value + '/' + r.PASMID);
                 });
            };
        },
        onRefresh: function(event) {
            event.onComplete = function() {
                var f = this,
                    r = f.record,
                    header = "Edit Assessment ({0})",
                    info = "";

                formRefreshCallBack(f, "ASMID", header);

                // ==============================
                // SPECIAL CASE
                // ==============================

                if (r.ASMID === 0) { // if new record then do not worry about reversed thing
                    $("#"+f.name).find("button[name=reverse]").addClass("hidden");
                    $("#"+f.name).find("button[name=save]").removeClass("hidden");
                    $("#"+f.name).find("button[name=saveadd]").removeClass("hidden");
                    $("#"+f.name).find("button[name=close]").addClass("hidden");
                    $("#"+f.name).find("#FLAGReport").addClass("hidden");
                    $("#"+f.name).find("#AssessmentInfo").addClass("hidden");

                    // ENABLE ALL INPUTS IF ALL OF THOSE HAVE BEEN DISABLED FOR REVERSED PREVIOUSLY
                    $("#"+f.name).find('input,button').prop("disabled", false);

                    return;
                } else {
                    $("#"+f.name).find("#FLAGReport").removeClass("hidden");
                    $("#"+f.name).find("#AssessmentInfo").removeClass("hidden");
                }

                // Assessment Info at the top of form
                // r.epoch = app.epochInstance[  (r.RentCycle !== 'Norecur' && r.PASMID === 0) ? 0 : 1 ];
                if (typeof r.RentCycle !== "object") { return; }
                if (r.RentCycle.text == 'Norecur' && r.PASMID === 0) {
                    // Non-recurred instance
                    info = '<p style="margin-bottom: 0px;">Single Instance Assessment</p>'.format(r.ASMID);
                } else {
                    // INSTANCE has 4 variables: ParentASM, RentCycle, Start, Stop
                    info = app.asmInstanceHeader.format(''+r.PASMID, r.RentCycle.text, f.pasmStart, f.pasmStop);
                }
                $("#"+f.name).find("#AssessmentInfo").html(info);

                // FLAG reports
                var flag = r.FLAGS,
                    flagHTML = "";

                // check if it is reversed or not
                if ( (flag & app.asmFLAGS.ASMREVERSED) !== 0 ) { // if reversed then
                    flagHTML += "<p style='margin-bottom: 5px;'><strong>{0}</strong> ({1})</p>".format("REVERSED", r.Comment);
                    // reversed indication icon
                    flagHTML += "<div class='reverseIconContainer'><i class='fa fa-exclamation-triangle fa-2x reverseIcon' aria-hidden='true'></i></div>";
                    // if reversed then do not show reverse, save button in form
                    $("#"+f.name).find("button[name=reverse]").addClass("hidden");
                    $("#"+f.name).find("button[name=save]").addClass("hidden");
                    $("#"+f.name).find("button[name=saveadd]").addClass("hidden");
                    // if reversed then we need to show close button at the bottom
                    $("#"+f.name).find("button[name=close]").removeClass("hidden");

                    // ****************************************
                    // IF REVERSED THEN DISABLE ALL INPUTS, BUTTONS  EXCEPT close button
                    // ****************************************
                    $("#"+f.name).find('input,button:not([name=close])').prop("disabled", true);

                } else {
                    // IF NOT REVERSED THEN ONLY SHOW PAID STATUS IN FOOTER
                    // unpaid, partial paid or fully paid
                    if ( (flag | app.asmFLAGS.ASMUNPAID) === 0 || (flag & (app.asmFLAGS.ASMPARTIALPAID | app.asmFLAGS.ASMFULLPAID)) === 0 ) {
                        flagHTML += "<p style='margin-bottom: 5px;'><strong>{0}</strong></p>".format("Unpaid");
                    }
                    else if ( (flag & app.asmFLAGS.ASMPARTIALPAID) !== 0 ) {
                        flagHTML += "<p style='margin-bottom: 5px;'><strong>{0}</strong></p>".format("Partially paid");
                    }
                    else if ( (flag & app.asmFLAGS.ASMFULLPAID) !== 0 ) {
                        flagHTML += "<p style='margin-bottom: 5px;'><strong>{0}</strong></p>".format("Fully paid");
                    }

                    // show reverse, save button, hide close button
                    $("#"+f.name).find("button[name=reverse]").removeClass("hidden");
                    $("#"+f.name).find("button[name=save]").removeClass("hidden");
                    $("#"+f.name).find("button[name=saveadd]").removeClass("hidden");
                    $("#"+f.name).find("button[name=close]").addClass("hidden");

                    // ****************************************
                    // IF not REVERSED THEN ENABLE ALL INPUTS
                    // ****************************************
                    $("#"+f.name).find('input,button').prop("disabled", false);
                }

                // finally append
                flagHTML += "<p style='margin-bottom: 5px;'>Last Update: {0} by {1}</p>".format(r.LastModTime, r.LastModBy);
                flagHTML += "<p>CreateTS: {0} by {1}</p>".format(r.CreateTS, r.CreateBy);
                $("#"+f.name).find("#FLAGReport").html(flagHTML);
            };
        },
        onChange: function(event) {
            event.onComplete = function() {
                var f = this,
                    r = f.record;

                if (event.target == "Start") {
                    if (r.RentCycle.text == "Norecur") {
                        // if Start date has been changed, in rentcycle with norecur mode
                        // then we need to set stop date same value of start date
                        r.Stop = r.Start;
                    }
                }
                if (event.target == "RentCycle") {
                    if (event.value_new.text == "Norecur") {
                        r.RentCycle = event.value_new;
                        r.ProrationCycle = "Norecur";
                        r.Stop = r.Start;
                        // disable stop date control
                        $("#"+f.name).find("input[name=Stop]").prop("disabled", true);
                    } else {
                        // enable stop date control
                        $("#"+f.name).find("input[name=Stop]").prop("disabled", false);
                    }
                }

                // copy original record temporary and reset it back after refresh event
                var temp = $.extend(true, {}, app.active_form_original);
                // finally refresh the form
                f.refresh();
                // now reset app original form record
                app.active_form_original = $.extend(true, {}, temp);

                // formRecDiffer: 1=current record, 2=original record, 3=diff object
                var diff = formRecDiffer(r, app.active_form_original, {});
                // if diff == {} then make dirty flag as false, else true
                if ($.isPlainObject(diff) && $.isEmptyObject(diff)) {
                    app.form_is_dirty = false;
                } else {
                    app.form_is_dirty = true;
                }
            };
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
        onSubmit: function(target, data) {
            delete data.postData.record.LastModTime;
            delete data.postData.record.LastModBy;
            delete data.postData.record.CreateTS;
            delete data.postData.record.CreateBy;
            // modify form data for server request
            getFormSubmitData(data.postData.record);
            app.cycleFreq.forEach(function(item, index) {
                if (item == data.postData.record.RentCycle) {
                    data.postData.record.RentCycle = index;
                }
                if (item == data.postData.record.ProrationCycle) {
                    data.postData.record.ProrationCycle = index;
                }
            });
        },
    });


    //------------------------------------------------------------------------
    //          asmsReverseMode
    //------------------------------------------------------------------------
    $().w2form({
        name: 'reverseMode',
        style: 'border: 0px; background-color: transparent;',
        formURL: '/webclient/html/formasmrev.html',
        focus  : 0,
        fields: [
            { field: 'ReverseMode', type: 'list',
              options: { items: app.asmsRevMode},
             required: true },
        ],
        actions: {
            reverse: function () {
                // var form = w2ui.asmInstForm;
                var form = app.AsmtModeCallerForm;
                var tgrid = w2ui.asmsGrid;

                console.log('asmsReverseMode: Mode = ' + w2ui.reverseMode.record.ReverseMode.id);
                w2popup.close();
                w2ui.toplayout.hide('right',true);
                tgrid.refresh();

                var params = {
                    cmd:         'delete',
                    formname:    form.name,
                    ASMID:       form.record.ASMID,
                    ReverseMode: w2ui.reverseMode.record.ReverseMode.id,
                };
                var dat = JSON.stringify(params);

                // They chose the delete / reverse button
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
                    form.error("Reverse Assessment instance failed.");
                    return;
                });
            },
            cancel: function() {
                w2popup.close();
            },
        },
    });
}

function popupAsmRevMode(mode,form) {
    w2ui.reverseMode.record.ReverseMode = mode;
    app.AsmtModeCallerForm = form;
    $().w2popup('open', {
        title   : 'Reverse Repeating Assessment',
        body    : '<div id="form" style="width: 100%; height: 100%;"></div>',
        style   : 'padding: 15px 0px 0px 0px',
        width   : 400,
        height  : 225,
        showMax : true,
        onToggle: function (event) {
            $(w2ui.reverseMode.box).hide();
            event.onComplete = function () {
                $(w2ui.reverseMode.box).show();
                w2ui.reverseMode.resize();
            };
        },
        onOpen: function (event) {
            event.onComplete = function () {
                $('#w2ui-popup #form').w2render('reverseMode');
            };
        }
    });
}

// popupRentalAgrPicker comes up when the user clicks on the Find... button
// while creating an assessment. It is used to locate a rental agreement by payor.
//----------------------------------------------------------------------------------
function popupRentalAgrPicker() {
    var x = getCurrentBusiness();
    app.RentalAgrFinder = {BID: x.value, RAID: 0, TCID: 0, RID: 0, FirstName: '', LastName: '', CompanyName: '', IsCompany: false, RAR: [], RARentablesNames: []};
    app.RentalAgrFinder.RARentablesNames = [{id: 0, text:" "}];
    w2ui.rentalAgrFinder.fields[2].options.items = app.RentalAgrFinder.RARentablesNames;
    w2ui.rentalAgrFinder.record.TCID = -1;
    w2ui.rentalAgrFinder.record.RAID = -1;
    w2ui.rentalAgrFinder.record.PayorName = '';
    w2ui.rentalAgrFinder.record.IsCompany = -1;
    w2ui.rentalAgrFinder.record.CompanyName = '';
    w2ui.rentalAgrFinder.record.FirstName = '';
    w2ui.rentalAgrFinder.record.LastName = '';
    w2ui.rentalAgrFinder.refresh();

    $().w2popup('open', {
        title   : 'Find Rental Agreement',
        body    : '<div id="form" style="width: 100%; height: 100%;"></div>',
        style   : 'padding: 15px 0px 0px 0px',
        width   : 400,
        height  : 250,
        showMax : true,
        onToggle: function (event) {
            $(w2ui.rentalAgrFinder.box).hide();
            event.onComplete = function () {
                $(w2ui.rentalAgrFinder.box).show();
                w2ui.rentalAgrFinder.resize();
            };
        },
        onOpen: function (event) {
            event.onComplete = function () {
                // specifying an onOpen handler instead would be equivalent to specifying
                // an onBeforeOpen handler, which would make this code execute too
                // early and hence not deliver.
                $('#w2ui-popup #form').w2render('rentalAgrFinder');
            };
        }
    });
}
