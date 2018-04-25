/*global
    popupRentalAgrPicker, $, console, w2ui, w2uiDateControlString, app,
    getCurrentBusiness, getBUDfromBID, w2popup, w2utils, rafinder, get2XReversalSymbolHTML,
    getGridReversalSymbolHTML, setDefaultFormFieldAsPreviousRecord, isDatePriorToCurrentDate,
    form_dirty_alert,setToForm,addDateNavToToolbar,getCurrentBID,formRefreshCallBack, renderReversalIcon,
    getBusinessAssessmentRules, getAsmsInitRecord, popupAsmRevMode, asmFormRASelect

*/
"use strict";

window.getAsmsInitRecord = function (BID, BUD, previousFormRecord){
    var y = new Date();
    var y1 = new Date(new Date().setFullYear(new Date().getFullYear() + 1));

    var initRentCycle, initProrationCycle;
    for (var i = 0; i < app.cycleFreq.length; i++) {
        if (app.cycleFreq[i] === "Monthly") {
            initRentCycle = i;
        }
        if (app.cycleFreq[i] === "Daily") {
            initProrationCycle = i;
        }
    }

    var defaultFormData = {
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
        RentCycle: initRentCycle,
        ProrationCycle: initProrationCycle,
        TCID: 0,
        Amount: 0,
        Rentable: '',
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

    // if it called after 'save and add another' action there previous form record is passed as Object
    // else it is null
    if ( previousFormRecord ) {
        defaultFormData = setDefaultFormFieldAsPreviousRecord(
            [ 'Amount', 'Comment', 'RAID', 'Rentable'], // Fields to Reset
            defaultFormData,
            previousFormRecord
        );
    }

    return defaultFormData;
};

//-----------------------------------------------------------------------------
// getBusinessAssessmentRules - return the promise object of request to get latest
//                              assessment rules for given BID.
//                              It updates the "app.AssessmentRules" variable for requested BUD
// @params  - BID : Business ID (expected current one)
//          - BUD : Business Unit Designation
// @return  - promise object from $.get
//-----------------------------------------------------------------------------
window.getBusinessAssessmentRules = function (BID, BUD) {
    // if not BUD in app.AssessmentRules then initialize it with blank list
    if (!(BUD in app.AssessmentRules)) {
        app.AssessmentRules[BUD] = [];
    }

    // return promise
    return $.get("/v1/uival/" + BID + "/app.AssessmentRules", null, null, "json").done(function(data) {
            // if it doesn't meet this condition, then save the data
            if (!('status' in data && data.status !== "success")) {
                app.AssessmentRules[BUD] = data[BUD];
            }
        });
};

window.renderReversalIcon = function (record /*, index, col_index*/) {
    if (typeof record === "undefined") {
        return;
    }
    if ( (record.FLAGS & app.asmFLAGS.REVERSED) !== 0 ) { // if reversed then
        return getGridReversalSymbolHTML();
    }
    return '';
};

window.buildAssessmentElements = function () {
    //------------------------------------------------------------------------
    //          asmsGrid
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'asmsGrid',
        url: '/v1/asms',
        multiSelect: false,
        postData: {searchDtStart: app.D1, searchDtStop: app.D2},
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
                            return '<i class="fas fa-sync-alt" title="epoch" aria-hidden="true"></i>';
                        } else if (record.RentCycle) { // if recurring assessment then put refresh icon
                            if (record.w2ui === undefined) {
                                record.w2ui = {class:""};
                            }
                            record.w2ui.class = "asmInstRow";
                            return '<i class="fas fa-sync-alt" title="recurring" aria-hidden="true"></i>';
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
            {field: 'PASMID', hidden: true, caption: 'PASMID', size: '40px', sortable: false},
            {field: 'RAID', caption: app.sRentalAgreement,  size: '125px', style: 'text-align: right', sortable: true},
            {field: 'RID', caption: 'RID',  size: '40px', hidden: true, sortable: false},
            {field: 'Rentable', caption: app.sRentable,  size: '150px', sortable: true},
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
                        var BID = getCurrentBID(),
                            BUD = getBUDfromBID(BID);

                        app.last.grid_sel_recid = parseInt(recid);

                        // keep highlighting current row in any case
                        grid.select(app.last.grid_sel_recid);

                        var rec = grid.get(recid);
                        var myurl = '/v1/asm/' + BID + '/' + rec.ASMID;
                        var formName = (rec.RentCycle !== 0 && rec.PASMID === 0) ? "asmEpochForm" : "asmInstForm";
                        var f = w2ui[formName];
                        console.log( 'calling setToForm( '+formName+', ' + myurl + ')');

                        // before setting to the form, get the list of AcctRules...
                        getBusinessAssessmentRules(BID, BUD)
                        .done( function(data) {
                            if ('status' in data && data.status !== 'success') {
                                f.message(data.message);
                            } else {
                                f.get('ARID').options.items = app.AssessmentRules[BUD];
                                f.refresh();
                                setToForm(f.name, myurl, 450, true);
                            }
                        })
                        .fail( function() {
                            console.log('Error getting /v1/uival/' + BID + '/app.AssessmentRules');
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
                    var BID = getCurrentBID(),
                        BUD = getBUDfromBID(BID);

                    // Always create epoch assessment
                    var f = w2ui.asmEpochForm;

                    // reset it
                    app.last.grid_sel_recid = -1;
                    grid.selectNone();

                    getBusinessAssessmentRules(BID, BUD)
                    .done( function(data) {
                        if ('status' in data && data.status !== 'success') {
                            f.message(data.message);
                        } else {
                            app.ridRentablePicker.BID = BID; // needed by typedown
                            // f.fields[5].options.url = '/v1/rentablestd/' + app.ridRentablePicker.BID;
                            f.get("ARID").options.items = app.AssessmentRules[BUD];
                            f.record = getAsmsInitRecord(BID, BUD, null);
                            f.refresh();

                            setToForm('asmEpochForm', '/v1/asm/' + BID + '/0', 450);
                        }
                    })
                    .fail( function() {
                        console.log('Error getting /v1/uival/'+BID+'/app.AssessmentRules');
                     });
                };

                // warn user if form content has been changed
                form_dirty_alert(yes_callBack, no_callBack, yes_args);
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
            { field: 'BID',           type: 'int',      required: true },
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
            { field: 'RentCycle',     type: 'list',     required: true, options: {items: app.w2ui.listItems.cycleFreq}, },
            { field: 'ProrationCycle',type: 'list',     required: true, options: {items: app.w2ui.listItems.cycleFreq}, },
            { field: 'Comment',       type: 'text',     required: false },
            { field: 'LastModTime',   type: 'hidden',   required: false },
            { field: 'LastModBy',     type: 'hidden',   required: false },
            { field: 'LastModByUser', type: 'hidden',   required: false },
            { field: 'CreateTS',      type: 'hidden',   required: false },
            { field: 'CreateBy',      type: 'hidden',   required: false },
            { field: 'CreateByUser',  type: 'hidden',   required: false },
            { field: 'ExpandPastInst',type: 'checkbox', required: false },
            { field: 'FLAGS',         type: 'w2int',    required: false },
            { field: 'Mode',          type: 'w2int',    required: false },
        ],
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
                    BID = getCurrentBID(),
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

                    getBusinessAssessmentRules(BID, BUD)
                    .done( function(data) {
                        if ('status' in data && data.status !== 'success') {
                            f.message(data.message);
                        } else {
                            app.ridRentablePicker.BID = BID; // needed by typedown
                            // f.fields[5].options.url = '/v1/rentablestd/' + app.ridRentablePicker.BID;
                            f.get("ARID").options.items = app.AssessmentRules[BUD];
                            f.record = getAsmsInitRecord(BID, BUD, f.record);
                            f.header = "Edit Assessment (new)"; // have to provide header here, otherwise have to call refresh method twice to get this change in form
                            f.url  = "/v1/asm/" + BID + "/0";
                            f.refresh();
                        }
                    })
                    .fail( function() {
                        console.log('Error getting /v1/uival/'+BID+'/app.AssessmentRules');
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
                    $(f.box).find("button[name=reverse]").addClass("hidden");
                    $(f.box).find("button[name=save]").removeClass("hidden");
                    $(f.box).find("button[name=saveadd]").removeClass("hidden");
                    $(f.box).find("button[name=close]").addClass("hidden");
                    $(f.box).find("#FLAGReport").addClass("hidden");
                    $(f.box).find("#AssessmentInfo").addClass("hidden");

                    // ENABLE ALL INPUTS IF ALL OF THOSE HAVE BEEN DISABLED FOR REVERSED PREVIOUSLY
                    // $(f.box).find('input,button').not('input[name=BUD]').prop("disabled", false);

                    return;
                } else {
                    $(f.box).find("#FLAGReport").removeClass("hidden");
                    $(f.box).find("#AssessmentInfo").removeClass("hidden");
                }

                // Assessment Info at the top of form in white box
                var info = '<p><i class="fas fa-sync-alt" style="margin-right: 5px;"></i> Repeating Assessment Series Definition</p>'.format(r.ASMID);
                $(f.box).find("#AssessmentInfo").html(info);

                // FLAG reports
                var flag = r.FLAGS,
                    flagHTML = "";

                // check if it is reversed or not
                if ( (flag & app.asmFLAGS.REVERSED) !== 0 ) { // if reversed then
                    flagHTML += "<p><strong>{0}</strong> ({1})</p>".format("REVERSED", r.Comment);
                    // reversed indication icon
                    flagHTML += get2XReversalSymbolHTML();
                    // if reversed then do not show reverse, save button
                    $(f.box).find("button[name=reverse]").addClass("hidden");
                    $(f.box).find("button[name=save]").addClass("hidden");
                    $(f.box).find("button[name=saveadd]").addClass("hidden");
                    // if reversed then we need to show close button at the bottom of form
                    $(f.box).find("button[name=close]").removeClass("hidden");

                    // ****************************************
                    // IF REVERSED THEN DISABLE ALL INPUTS, BUTTONS  EXCEPT close button
                    // ****************************************
                    $(f.box).find('input,button:not([name=close])').prop("disabled", true);

                } else {
                    // IF NOT REVERSED THEN ONLY SHOW PAID STATUS IN FOOTER
                    // unpaid, partial paid or fully paid
                    if ( (flag | app.asmFLAGS.UNPAID) === 0 || (flag & (app.asmFLAGS.PARTIALPAID | app.asmFLAGS.FULLYPAID)) === 0 ) {
                        flagHTML += "<p><strong>{0}</strong></p>".format("Unpaid");
                    }
                    else if ( (flag & app.asmFLAGS.PARTIALPAID) !== 0 ) {
                        flagHTML += "<p><strong>{0}</strong></p>".format("Partially paid");
                    }
                    else if ( (flag & app.asmFLAGS.FULLYPAID) !== 0 ) {
                        flagHTML += "<p><strong>{0}</strong></p>".format("Fully paid");
                    }

                    // show reverse, save button, hide close button
                    $(f.box).find("button[name=reverse]").removeClass("hidden");
                    $(f.box).find("button[name=save]").removeClass("hidden");
                    $(f.box).find("button[name=saveadd]").removeClass("hidden");
                    $(f.box).find("button[name=close]").addClass("hidden");

                    // ****************************************
                    // IF not REVERSED THEN ENABLE ALL INPUTS
                    // ****************************************
                    $(f.box).find('input,button').not('input[name=BUD]').prop("disabled", false);
                }

                // finally append
                flagHTML += "<p>Last Update: {0} by {1}</p>".format(r.LastModTime, r.LastModByUser);
                flagHTML += "<p>Created: {0} by {1}</p>".format(r.CreateTS, r.CreateByUser);
                $(f.box).find("#FLAGReport").html(flagHTML);
            };
        },
        onChange: function(event) {
            event.onComplete = function() {
                var f = this,
                    r = f.record;
                if (event.target == "Start") {
                    var x = document.getElementsByName('ExpandPastInst')[0];
                    var DtStart = dateFromString(event.value_new);
                    if (r.RentCycle.text != "Norecur") {
                        // create past instances is marked as true if startdate is prior to current date
                        f.record.ExpandPastInst = isDatePriorToCurrentDate(DtStart);
                        $(f.box).find("input[name=ExpandPastInst]").prop( "disabled", !isDatePriorToCurrentDate(DtStart) );
                    } else {
                        // if Start date has been changed, in rentcycle with norecur mode
                        // then we need to set stop date same value of start date
                        r.Stop = r.Start;
                        // Norecur then disable checkbox for "create past instances"
                        $(f.box).find("input[name=ExpandPastInst]").prop( "disabled", true);
                        f.record.ExpandPastInst = false;
                    }
                }
                if (event.target == "RentCycle") {
                    if (event.value_new.text == "Norecur") {
                        r.RentCycle = event.value_new;
                        r.ProrationCycle = "Norecur";
                        r.Stop = r.Start;
                        // disable stop date control
                        $(f.box).find("input[name=Stop]").prop( "disabled", true );
                        // Norecur then disable checkbox for "create past instances"
                        $(f.box).find("input[name=ExpandPastInst]").prop( "disabled", true);
                        f.record.ExpandPastInst = false;
                    } else {
                        // enable stop date control
                        $(f.box).find("input[name=Stop]").prop("disabled", false);
                        var startDate = $(f.box).find("input[name=Start]").val();
                        f.record.ExpandPastInst  = isDatePriorToCurrentDate(new Date(startDate));
                        $(f.box).find("input[name=ExpandPastInst]").prop( "disabled", !isDatePriorToCurrentDate(new Date(startDate)) );
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
            { field: 'RentCycle',     type: 'list',   options:  {items: app.w2ui.listItems.cycleFreq}, required: true },
            { field: 'ProrationCycle',type: 'list',   options:  {items: app.w2ui.listItems.cycleFreq}, required: true },
            { field: 'Comment',       type: 'text',   required: false },
            { field: 'LastModTime',   type: 'hidden', required: false },
            { field: 'LastModBy',     type: 'hidden', required: false },
            { field: 'LastModByUser', type: 'hidden', required: false },
            { field: 'CreateTS',      type: 'hidden', required: false },
            { field: 'CreateBy',      type: 'hidden', required: false },
            { field: 'CreateByUser',  type: 'hidden', required: false },
            { field: 'FLAGS',         type: 'w2int',  required: false },
            { field: 'Mode',          type: 'w2int',  required: false },
        ],
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
                    BID = getCurrentBID(),
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

                    getBusinessAssessmentRules(BID, BUD)
                    .done( function(data) {
                        if ('status' in data && data.status !== 'success') {
                            f.message(data.message);
                        } else {
                            app.ridRentablePicker.BID = BID; // needed by typedown

                            // epochForm.fields[5].options.url = '/v1/rentablestd/' + app.ridRentablePicker.BID;
                            epochForm.get("ARID").options.items = app.AssessmentRules[BUD];
                            epochForm.record =  getAsmsInitRecord(BID, BUD, null);
                            // epochForm.header = "Edit Assessment (new)"; // have to provide header here, otherwise have to call refresh method twice to get this change in form
                            // f.refresh();
                            setToForm(epochForm.name, '/v1/asm/' + BID + '/0', 400);
                        }
                    })
                    .fail( function() {
                        console.log('Error getting /v1/uival/'+BID+'/app.AssessmentRules');
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

                // if PASMID is 0 then return
                if (r.PASMID == 0) {
                    return;
                }

                var BID = getCurrentBID();
                var params = {"cmd":"get", "recid":0,"name":"asmInstForm"};
                var dat = JSON.stringify(params);
                $.post('/v1/asm/' + BID + '/' + r.PASMID, dat, null, "json")
                .done( function(data) {
                    if (data.status !== 'success') {
                        f.message(data.message);
                        f.pasmStart = "";
                        f.pasmStop = "";
                    } else {
                        // get parent assessment dates and store it in form
                        f.pasmStart = data.record.Start;
                        f.pasmStop = data.record.Stop;
                    }
                })
                .fail( function() {
                    console.log('Error getting /v1/asm/' + BID + '/' + r.PASMID);
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
                    $(f.box).find("button[name=reverse]").addClass("hidden");
                    $(f.box).find("button[name=save]").removeClass("hidden");
                    $(f.box).find("button[name=saveadd]").removeClass("hidden");
                    $(f.box).find("button[name=close]").addClass("hidden");
                    $(f.box).find("#FLAGReport").addClass("hidden");
                    $(f.box).find("#AssessmentInfo").addClass("hidden");

                    // ENABLE ALL INPUTS IF ALL OF THOSE HAVE BEEN DISABLED FOR REVERSED PREVIOUSLY
                    // $(f.box).find('input,button').not('input[name=BUD]').prop("disabled", false);

                    return;
                } else {
                    $(f.box).find("#FLAGReport").removeClass("hidden");
                    $(f.box).find("#AssessmentInfo").removeClass("hidden");
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
                $(f.box).find("#AssessmentInfo").html(info);

                // FLAG reports
                var flag = r.FLAGS,
                    flagHTML = "";

                // check if it is reversed or not
                if ( (flag & app.asmFLAGS.REVERSED) !== 0 ) { // if reversed then
                    flagHTML += "<p><strong>{0}</strong> ({1})</p>".format("REVERSED", r.Comment);
                    // reversed indication icon
                    flagHTML += get2XReversalSymbolHTML();
                    // if reversed then do not show reverse, save button in form
                    $(f.box).find("button[name=reverse]").addClass("hidden");
                    $(f.box).find("button[name=save]").addClass("hidden");
                    $(f.box).find("button[name=saveadd]").addClass("hidden");
                    // if reversed then we need to show close button at the bottom
                    $(f.box).find("button[name=close]").removeClass("hidden");

                    // ****************************************
                    // IF REVERSED THEN DISABLE ALL INPUTS, BUTTONS  EXCEPT close button
                    // ****************************************
                    $(f.box).find('input,button:not([name=close])').prop("disabled", true);

                } else {
                    // IF NOT REVERSED THEN ONLY SHOW PAID STATUS IN FOOTER
                    // unpaid, partial paid or fully paid
                    if ( (flag | app.asmFLAGS.UNPAID) === 0 || (flag & (app.asmFLAGS.PARTIALPAID | app.asmFLAGS.FULLYPAID)) === 0 ) {
                        flagHTML += "<p><strong>{0}</strong></p>".format("Unpaid");
                    }
                    else if ( (flag & app.asmFLAGS.PARTIALPAID) !== 0 ) {
                        flagHTML += "<p><strong>{0}</strong></p>".format("Partially paid");
                    }
                    else if ( (flag & app.asmFLAGS.FULLYPAID) !== 0 ) {
                        flagHTML += "<p><strong>{0}</strong></p>".format("Fully paid");
                    }

                    // show reverse, save button, hide close button
                    $(f.box).find("button[name=reverse]").removeClass("hidden");
                    $(f.box).find("button[name=save]").removeClass("hidden");
                    $(f.box).find("button[name=saveadd]").removeClass("hidden");
                    $(f.box).find("button[name=close]").addClass("hidden");

                    // ****************************************
                    // IF not REVERSED THEN ENABLE ALL INPUTS
                    // ****************************************
                    $(f.box).find('input,button').not('input[name=BUD]').prop("disabled", false);
                }

                // finally append
                flagHTML += "<p>Last Update: {0} by {1}</p>".format(r.LastModTime, r.LastModByUser);
                flagHTML += "<p>Created: {0} by {1}</p>".format(r.CreateTS, r.CreateByUser);
                $(f.box).find("#FLAGReport").html(flagHTML);
            };
        },
        onChange: function(event) {
            event.onComplete = function() {
                var f = this,
                    r = f.record;
                if (event.target == "Start") {
                    var x = document.getElementsByName('ExpandPastInst')[0];
                    var DtStart = dateFromString(event.value_new);
                    if (r.RentCycle.text != "Norecur") {
                        // create past instances is marked as true if startdate is prior to current date
                        f.record.ExpandPastInst = isDatePriorToCurrentDate(DtStart);
                        $(f.box).find("input[name=ExpandPastInst]").prop( "disabled", !isDatePriorToCurrentDate(DtStart) );
                    } else {
                        // if Start date has been changed, in rentcycle with norecur mode
                        // then we need to set stop date same value of start date
                        r.Stop = r.Start;
                        // Norecur then disable checkbox for "create past instances"
                        $(f.box).find("input[name=ExpandPastInst]").prop( "disabled", true);
                        f.record.ExpandPastInst = false;
                    }
                }
                if (event.target == "RentCycle") {
                    if (event.value_new.text == "Norecur") {
                        r.RentCycle = event.value_new;
                        r.ProrationCycle = "Norecur";
                        r.Stop = r.Start;
                        // disable stop date control
                        $(f.box).find("input[name=Stop]").prop( "disabled", true );
                        // Norecur then disable checkbox for "create past instances"
                        $(f.box).find("input[name=ExpandPastInst]").prop( "disabled", true);
                        f.record.ExpandPastInst = false;
                    } else {
                        // enable stop date control
                        $(f.box).find("input[name=Stop]").prop("disabled", false);
                        var startDate = $(f.box).find("input[name=Start]").val();
                        f.record.ExpandPastInst  = isDatePriorToCurrentDate(new Date(startDate));
                        $(f.box).find("input[name=ExpandPastInst]").prop( "disabled", !isDatePriorToCurrentDate(new Date(startDate)) );
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
};

window.asmOpenRASelect = function () {
    rafinder.cb = asmFormRASelect;
    popupRentalAgrPicker();
};

window.asmFormRASelect = function () {
    w2ui.asmEpochForm.record.RAID = w2ui.rentalAgrPicker.record.RAID;
    w2ui.asmEpochForm.record.Rentable = w2ui.rentalAgrPicker.record.RentableName.text;
    w2ui.asmEpochForm.record.RID = w2ui.rentalAgrPicker.record.RentableName.id;
    w2ui.asmEpochForm.refresh();
};

window.popupAsmRevMode = function (mode,form) {
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
};

