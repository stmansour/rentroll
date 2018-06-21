"use strict";
/*global
    w2ui, $, app, console,setToTLDForm, w2alert,
    form_dirty_alert, addDateNavToToolbar, w2utils, 
    openTaskDescForm, ensureSession, dtFormatISOToW2ui,
    dtFormatISOToW2ui, localtimeToUTC, setDefaultFormFieldAsPreviousRecord,
    getTLDInitRecord, getCurrentBID, getTDInitRecord, saveTaskListDefinition,
    closeTaskDescForm, setTaskDescButtonsState, newDateKeepOldTime,
    w2uiDateTimeControlString,
*/

// Temporary storage for when a date is toggled off
var TaskDescData = {
    sEpochDue: '',
    sEpochPreDue: '',
};

var TLData = {
    sEpoch: '',
    sEpochDue: '',
    sEpochPreDue: '',
};

var TLD = {
    FormWidth: 450,
    TaskDescWidth: 400,
    formBtnsDisabled: false,
    TIME0: '1/1/1970',
};

window.getTLDInitRecord = function (BID, previousFormRecord){
    var y = new Date();
    var y1 = new Date( y.getFullYear(), y.getMonth(), 1);
    var month = (y.getMonth() + 1) % 12;
    var epochPreDue = new Date(y.getFullYear(), y.getMonth(), 20);
    var epochDue = new Date(y.getFullYear(), month, 0); 

    // var Cycle;
    // for (var i = 0; i < app.cycleFreq.length; i++) {
    //     if (app.cycleFreq[i] === "Monthly") {
    //         Cycle = i;
    //     }
    // }

    var defaultFormData = {
        TLDID: 0,
        BID: BID,
        Name: '',
        Cycle: 6,
        ChkEpochDue: true,
        ChkEpochPreDue: true,
        Epoch: dtFormatISOToW2ui(y1.toString()),
        EpochDue: dtFormatISOToW2ui(epochDue.toString()),
        EpochPreDue: dtFormatISOToW2ui(epochPreDue.toString()),
        DurWait: 86400000000000,
        FLAGS: 0,
        Comment: '',
        CreateTS: y.toString(),
        CreateBy: app.uid,
        LastModTime: y.toString(),
        LastModBy: app.uid,
        EmailList: '',
    };

    // if called after 'save and add another' action there previous form record is passed as Object
    // else it is null
    if ( previousFormRecord ) {
        defaultFormData = setDefaultFormFieldAsPreviousRecord(
            [ 'ChkEpochDue', 'ChkEpochPreDue', 'EpochDue', 'EpochPreDue'], // Fields to Reset
            defaultFormData,
            previousFormRecord
        );
    }
    return defaultFormData;
};

window.getTDInitRecord = function (BID, TDID, previousFormRecord){
    var y = new Date();
    var y1 = new Date( y.getFullYear(), y.getMonth(), 1);
    var month = (y.getMonth() + 1) % 12;
    var epochPreDue = new Date(y.getFullYear(), y.getMonth(), 20);
    var epochDue = new Date(y.getFullYear(), month, 0); 

    var defaultFormData = {
        TDID: 0,
        Worker: '',
        lstWorker: '',
        DoneUID: 0,
        PreDoneUID: 0,
        TLDID: w2ui.tldsInfoForm.record.TLDID,
        BID: BID,
        Name: '',
        Cycle: 6,
        ChkEpochDue: true,
        ChkEpochPreDue: true,
        Epoch: dtFormatISOToW2ui(y1.toString()),
        EpochDue: dtFormatISOToW2ui(epochDue.toString()),
        EpochPreDue: dtFormatISOToW2ui(epochPreDue.toString()),
        FLAGS: 0,
        Comment: '',
        CreateTS: y.toString(),
        CreateBy: app.uid,
        LastModTime: y.toString(),
        LastModBy: app.uid,
    };

    // if called after 'save and add another' action there previous form record is passed as Object
    // else it is null
    if ( previousFormRecord ) {
        defaultFormData = setDefaultFormFieldAsPreviousRecord(
            [ 'ChkEpochDue', 'ChkEpochPreDue', 'EpochDue', 'EpochPreDue'], // Fields to Reset
            defaultFormData,
            previousFormRecord
        );
    }
    return defaultFormData;
};


window.buildTaskListDefElements = function () {
    //------------------------------------------------------------------------
    //          tldsGrid  -  THE LIST OF ALL Task List Definitions
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'tldsGrid',
        url: '/v1/tlds',
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
            {field: 'recid',     hidden: true,  caption: 'recid',                   size: '40px',  sortable: true},
            {field: 'BID',       hidden: true,  caption: 'BID',                     size: '40px',  sortable: true},
            {field: 'Name',      hidden: false, caption: 'Name',                    size: '250px', sortable: true},
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
                        console.log( 'BID = ' + rec.BID + ',   TLDID = ' + rec.TLDID);
                        setToTLDForm(rec.BID, rec.TLDID, app.D1, app.D2);
                    };

                // warn user if form content has been changed
                form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
            };
        },
        onAdd: function (/*event*/) {
            var yes_args = [this],
                no_callBack = function() { return false; },
                yes_callBack = function(grid) {
                    var BID = getCurrentBID();
                    // reset it
                    app.last.grid_sel_recid = -1;
                    grid.selectNone();
                    setToTLDForm(BID, 0, app.D1, app.D2);
                };

                // warn user if form content has been changed
                form_dirty_alert(yes_callBack, no_callBack, yes_args);
        },
    });

    addDateNavToToolbar('tlds'); // "Grid" is appended to the 

    //------------------------------------------------------------------------
    //  tldsInfoForm
    //------------------------------------------------------------------------
    $().w2form({
        name: 'tldsInfoForm',
        style: 'border: 0px; background-color: transparent;',
        header: 'Task List Definition',
        url: '/v1/tld',
        formURL: '/webclient/html/formtld.html',
        toolbar: {
            items: [
                { id: 'btnNotes', type: 'button', icon: 'far fa-sticky-note' },
                { id: 'bt3', type: 'spacer' },
                { id: 'btnClose', type: 'button', icon: 'fas fa-times' },
            ],
            onClick: function (event) {
                event.onComplete = function() {
                    // var g = w2ui.tlsDetailGrid;
                    // var r = w2ui.tlsInfoForm.record;
                    // var d1, d2, url;
                    switch(event.target) {
                    case 'btnClose':
                        var no_callBack = function() { return false; },
                            yes_callBack = function() {
                                closeTaskDescForm();
                                w2ui.toplayout.hide('right',true);
                                w2ui.tlsGrid.render();
                            };
                        form_dirty_alert(yes_callBack, no_callBack);
                        break;
                    }
                };
            },
        },
        fields: [
            { field: 'recid',          type: 'int',      required: false },
            { field: 'TLDID',          type: 'int',      required: false },
            { field: 'BID',            type: 'int',      required: false },
            { field: 'Name',           type: 'text',     required: true },
            { field: 'Cycle',          type: 'list',     required: true, options: {items: app.w2ui.listItems.cycleFreq}, },
            { field: 'ChkEpochDue',    type: 'checkbox', required: false },
            { field: 'ChkEpochPreDue', type: 'checkbox', required: false },
            { field: 'Epoch',          type: 'datetime', required: false },
            { field: 'EpochDue',       type: 'datetime', required: false },
            { field: 'EpochPreDue',    type: 'datetime', required: false },
            { field: 'FLAGS',          type: 'int',      required: false },
            { field: 'EmailList',      type: 'text',     required: false },
            { field: 'Comment',        type: 'text',     required: false },
            { field: 'DurWait',        type: 'int',      required: false },
            { field: 'CreateTS',       type: 'date',     required: false },
            { field: 'CreateBy',       type: 'int',      required: false },
            { field: 'TZOffset',       type: 'int',      required: false },
            { field: 'LastModTime',    type: 'date',     required: false },
            { field: 'LastModBy',      type: 'int',      required: false },
        ],
        onLoad: function(event) {
            event.onComplete = function(event) {
                var f = w2ui.tldsInfoForm;
                var r = f.record;
                if (typeof r.EpochPreDue === "undefined") {
                    return;
                }
                
                // translate dates into a format that w2ui understands
                r.EpochPreDue = dtFormatISOToW2ui(r.EpochPreDue);
                r.EpochDue    = dtFormatISOToW2ui(r.EpochDue);
                r.Epoch       = dtFormatISOToW2ui(r.Epoch);

                // now enable/disable as needed
                $(f.box).find("input[name=EpochDue]").prop( "disabled", !r.ChkEpochDue );
                $(f.box).find("input[name=EpochPreDue]").prop( "disabled", !r.ChkEpochPreDue );
                $(f.box).find("input[name=Epoch]").prop( "disabled", r.Cycle < 4);  // enable if recur is Daily, Weekly, Monthlhy, quarterly or yearly
            };
        },
        onChange: function(event) {
            event.onComplete = function() {
                var f = this;
                var r = f.record;
                var b;
                switch (event.target) {
                case "ChkEpochPreDue":
                    $(f.box).find("input[name=EpochPreDue]").prop( "disabled", !r.ChkEpochPreDue );
                    if (r.ChkEpochPreDue) {
                        if (r.EpochPreDue === "" && TLData.sEpochPreDue.length > 1) {
                            r.EpochPreDue = TLData.sEpochPreDue;
                        }
                    } else {
                        TLData.sEpochPreDue = r.EpochPreDue;
                        r.EpochPreDue = '';
                    }
                    f.refresh();
                    break;
                case "ChkEpochDue":
                    $(f.box).find("input[name=EpochDue]").prop( "disabled", !r.ChkEpochDue );
                    if (r.ChkEpochDue) {
                        if (r.EpochDue === "" && TLData.sEpochDue.length > 1) {
                            r.EpochDue = TLData.sEpochDue;
                        }
                    } else {
                        TLData.sEpochDue = r.EpochDue;
                        r.EpochDue = '';
                    }
                    f.refresh();
                    break;
                case "Cycle":
                    b = r.Cycle.id < 4; // 4 is daily
                    $(f.box).find("input[name=Epoch]").prop( "disabled", b);
                    if (b && event.value_previous.id >= 4) {  // change from need date to don't need date
                        TLData.sEpoch = r.Epoch;
                        r.Epoch = '';
                    } else if (!b && event.value_previous.id < 4 ) { // change from don't need date to need date
                        if (r.Epoch === "" && TLData.sEpoch.length > 1) {
                            r.Epoch = TLData.sEpoch;
                        }
                    }
                    f.refresh();
                    break;
                case "Epoch":
                case "EpochDue":
                case "EpochPreDue":
                    // all dates must be in sync if cycle > 0 and < 4
                    if (0 < r.Cycle.id && r.Cycle.id <= 4) {
                        if (event.value_new === "") {
                            r[event.target] = event.value_previous;
                            event.isCancelled = true;
                            f.refresh();
                            return;
                        }
                        var dt = new Date(r[event.target]);
                        var da = dt.getDate();
                        var mn = dt.getMonth();
                        var yr = dt.getFullYear();
                        r.Epoch       = w2uiDateTimeControlString(newDateKeepOldTime(r.Epoch,yr,mn,da));
                        r.EpochDue    = w2uiDateTimeControlString(newDateKeepOldTime(r.EpochDue,yr,mn,da));
                        r.EpochPreDue = w2uiDateTimeControlString(newDateKeepOldTime(r.EpochPreDue,yr,mn,da));
                    }
                    // we must always keep epoch at localtime 00:00
                    var ddt = new Date(r.Epoch);
                    var dd1 = new Date(ddt.getFullYear(), ddt.getMonth(), ddt.getDate(), 0, 0 );
                    r.Epoch = w2uiDateTimeControlString(dd1);
                    f.refresh();
                    break;
                }
            };
        },

    });

    //------------------------------------------------------------------------
    //  tldsTaskGrid  -  lists all the tasks associated with the task list
    //                   definition currently being edited.
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'tldsDetailGrid',
        url: '/v1/tds/',
        multiSelect: false,
        postData: {searchDtStart: app.D1, searchDtStop: app.D2, Bool1: app.PayorStmtExt},
        show: {
            toolbar         : true,
            footer          : false,
            toolbarAdd      : true,   // indicates if toolbar add new button is visible
            toolbarDelete   : false,   // indicates if toolbar delete button is visible
            toolbarSave     : false,   // indicates if toolbar save button is visible
            selectColumn    : false,
            expandColumn    : false,
            toolbarEdit     : false,
            toolbarSearch   : false,
            toolbarInput    : false,
            searchAll       : false,
            toolbarReload   : false,
            toolbarColumns  : true,
        },
        columns: [
            { field: 'recid',       caption: 'recid',       size: '35px',  sortable: true, hidden: true},
            { field: 'TDID',        caption: 'TDID',        size: '35px',  sotrable: true, hidden: true},
            { field: 'BID',         caption: 'BID',         size: '35px',  sotrable: true, hidden: true},
            { field: 'TLDID',       caption: 'TLDID',       size: '35px',  sotrable: true, hidden: true},
            { field: 'TDName',      caption: 'Name',        size: '360px', sotrable: true, hidden: false},
            { field: 'Worker',      caption: 'Worker',      size: '95px',  sotrable: true, hidden: false},
            { field: 'EpochPreDue', caption: 'Pre Due',     size: '130px', sotrable: true, hidden: false,
                render: function (rec, idx, col) {if (typeof rec === "undefined") {return ''; } return dtFormatISOToW2ui(rec.EpochPreDue); }
            },
            { field: 'EpochDue',    caption: 'Due',         size: '130px', sotrable: true, hidden: false,
                render: function (rec, idx, col) {if (typeof rec === "undefined") {return ''; } return dtFormatISOToW2ui(rec.EpochDue); }
            },
            { field: 'FLAGS',       caption: 'FLAGS',       size: '35px',  sotrable: true, hidden: true},
            { field: 'DoneUID',     caption: 'DoneUID',     size: '35px',  sotrable: true, hidden: true},
            { field: 'PreDoneUID',  caption: 'PreDoneUID',  size: '35px',  sotrable: true, hidden: true},
            { field: 'Comment',     caption: 'Comment',     size: '35px',  sotrable: true, hidden: true},
            { field: 'LastModTime', caption: 'LastModTime', size: '35px',  sotrable: true, hidden: true},
            { field: 'LastModBy',   caption: 'LastModBy',   size: '35px',  sotrable: true, hidden: true},
            { field: 'CreateTS',    caption: 'CreateTS',    size: '35px',  sotrable: true, hidden: true},
            { field: 'CreateBy',    caption: 'CreateBy',    size: '35px',  sotrable: true, hidden: true},
        ],
        onClick: function(event) {
            event.onComplete = function (event) {
                var r = w2ui.tldsDetailGrid.records[event.recid];
                openTaskDescForm(r.BID,r.TDID);
            };
        },
        onRender: function (event) {
            event.onComplete = function (event) {
                setTaskDescButtonsState();
            };
        },
        onAdd: function (event) {
            // make sure that task list definition itself has been saved, if not, we'll need to save it.
            var r = w2ui.tldsInfoForm.record;
            if (r.TLDID === 0) {
                if (saveTaskListDefinition(false,true) > 0) {
                    return;  // an error occurred.
                }
            }
            openTaskDescForm(w2ui.tldsInfoForm.record.BID,0);
        }
    });

    //------------------------------------------------------------------------
    //  taskDescForm
    //------------------------------------------------------------------------
    $().w2form({
        name: 'taskDescForm',
        style: 'border: 0px; background-color: transparent;',
        formURL: '/webclient/html/formtd.html',
        url: '/v1/td',
        toolbar: {
            items: [
                { id: 'btnNotes', type: 'button', icon: 'far fa-sticky-note' },
                { id: 'bt3', type: 'spacer' },
                { id: 'btnClose', type: 'button', icon: 'fas fa-times' },
            ],
            onClick: function (event) {
                event.onComplete = function() {
                    switch(event.target) {
                    case 'btnClose':
                        var no_callBack = function() { return false; },
                            yes_callBack = function() {
                                closeTaskDescForm();
                                w2ui.tldLayout.render();
                            };
                        form_dirty_alert(yes_callBack, no_callBack);
                        break;
                    }
                };
            },
        },
        fields: [
            { field: 'recid',          type: 'int',         required: false },
            { field: 'TDID',           type: 'int',         required: false },
            { field: 'BID',            type: 'int',         required: false },
            { field: 'TLDID',          type: 'int',         required: false },
            { field: 'TDName',         type: 'text',        required: true  },
            { field: 'Worker',         type: 'text',        required: false },
            { field: 'lstWorker',      type: 'list',        required: false, options: {items: app.workers}, },
            { field: 'EpochDue',       type: 'datetime',    required: false },
            { field: 'EpochPreDue',    type: 'datetime',    required: false },
            { field: 'ChkEpochDue',    type: 'checkbox',    required: false },
            { field: 'ChkEpochPreDue', type: 'checkbox',    required: false },
            { field: 'FLAGS',          type: 'text',        required: false },
            { field: 'DoneUID',        type: 'int',         required: false },
            { field: 'PreDoneUID',     type: 'int',         required: false },
            { field: 'TDComment',      type: 'text',        required: false },
            { field: 'LastModTime',    type: 'date',        required: false },
            { field: 'LastModBy',      type: 'int',         required: false },
            { field: 'CreateTS',       type: 'date',        required: false },
            { field: 'CreateBy',       type: 'int',         required: false },
            { field: 'TZOffset',       type: 'int',         required: false },
        ],
        actions: {
            save: function(target, data){
                //---------------------------------------------------------
                // When the w2popup is active, it suspends the operation
                // of things like setInterval() handling.  So the session
                // may have expired by the time we close this dialog. So,
                // we need to explicity call ensureSession to make sure
                // we have a session before proceeding.
                //---------------------------------------------------------
                ensureSession();

                //---------------------------
                // Now, on with the save...
                //---------------------------
                var f = w2ui.taskDescForm;
                var r = f.record;
                r.Worker = r.lstWorker.text;
                if (r.TLDID === 0) {
                    r.TLDID = w2ui.tldsInfoForm.record.TLDID;  // this should no longer be 
                }

                //------------------------------------------------
                // convert times to UTC before saving
                //------------------------------------------------
                r.EpochDue = localtimeToUTC(r.EpochDue);
                r.EpochPreDue = localtimeToUTC(r.EpochPreDue);
                if (r.EpochDue.length === 0) {
                    r.EpochDue = TLD.TIME0;
                }
                if (r.EpochPreDue.length === 0) {
                    r.EpochPreDue = TLD.TIME0;
                }
                r.TZOffset = app.TZOffset;

                var d = {cmd: "save", record: r};
                var dat=JSON.stringify(d);
                f.url = '/v1/td/' + r.BID + '/' + r.TDID;

                $.post(f.url,dat)
                .done(function(data) {
                    if (data.status === "error") {
                        f.error(w2utils.lang(data.message));
                        return;
                    }
                    w2ui.tldsDetailGrid.url='/v1/tds/'+w2ui.taskDescForm.record.BID+'/'+w2ui.taskDescForm.record.TLDID;
                    w2ui.tldsDetailGrid.reload();
                    // w2popup.close();
                    closeTaskDescForm();
                    setTaskDescButtonsState();
                })
                .fail(function(/*data*/){
                    f.error("Save TaskDescriptor failed.");
                    return;
                });
            },
            delete: function(target, data) {
                ensureSession();
                var f = w2ui.taskDescForm;
                var r = f.record;
                var d = {cmd: "delete"};
                var dat=JSON.stringify(d);
                f.url = '/v1/td/' + r.BID + '/' + r.TDID;
                $.post(f.url,dat)
                .done(function(data) {
                    if (data.status === "error") {
                        f.error(w2utils.lang(data.message));
                        return;
                    }
                    w2ui.tldsDetailGrid.reload();
                    // w2popup.close();
                    closeTaskDescForm();
                })
                .fail(function(/*data*/){
                    f.error("Delete TaskDescriptor failed.");
                    return;
                });
            },
        },
       onLoad: function(event) {
            event.onComplete = function(event) {
                var f = w2ui.taskDescForm;
                var r = f.record;
                if (typeof r.EpochPreDue === "undefined") {
                    return;
                }
                r.EpochPreDue = dtFormatISOToW2ui(r.EpochPreDue);
                r.EpochDue    = dtFormatISOToW2ui(r.EpochDue);
                $(f.box).find("input[name=EpochPreDue]").prop( "disabled", !r.ChkEpochPreDue );
                $(f.box).find("input[name=EpochDue]").prop( "disabled", !r.ChkEpochDue );
            };
        },
       onRender: function(event) {
            event.onComplete = function(event) {
                setTaskDescButtonsState();
            };
        },
        onChange: function(event) {
            event.onComplete = function() {
                var f = this;
                var r = f.record;
                switch (event.target) {
                case "ChkEpochPreDue":
                    $(f.box).find("input[name=EpochPreDue]").prop( "disabled", !r.ChkEpochPreDue );
                    if (r.ChkEpochPreDue) {
                        if (r.EpochPreDue === "" && TaskDescData.sEpochPreDue.length > 1) {
                            r.EpochPreDue = TaskDescData.sEpochPreDue;
                        }
                    } else {
                        TaskDescData.sEpochPreDue = r.EpochPreDue;
                        r.EpochPreDue = '';
                    }
                    f.refresh();
                    break;
                case "ChkEpochDue":
                    $(f.box).find("input[name=EpochDue]").prop( "disabled", !r.ChkEpochDue );
                    if (r.ChkEpochDue) {
                        if (r.EpochDue === "" && TaskDescData.sEpochDue.length > 1) {
                            r.EpochDue = TaskDescData.sEpochDue;
                        }
                    } else {
                        TaskDescData.sEpochDue = r.EpochDue;
                        r.EpochDue = '';
                    }
                    f.refresh();
                    break;
                }
            };
        },
    });


    //------------------------------------------------------------------------
    //  tldsCloseForm
    //------------------------------------------------------------------------
    $().w2form({
        name: 'tldsCloseForm',
        style: 'border: 0px; background-color: transparent;',
        formURL: '/webclient/html/formtldclose.html',
        url: '',
        fields: [],
        actions: {
            save: function(target, data){
                saveTaskListDefinition(true,false);
            },
            delete: function(target, data) {
                var tl = {
                    cmd: "delete",
                };
                var dat = JSON.stringify(tl);
                var url = '/v1/tld/' + w2ui.tldsInfoForm.record.BID + '/' + w2ui.tldsInfoForm.record.TLDID;
                $.post(url,dat)
                .done(function(data) {
                    if (data.status === "error") {
                        w2ui.tldsInfoForm.error(w2utils.lang(data.message));
                        return;
                    }
                    w2ui.toplayout.hide('right',true);
                    w2ui.tldsGrid.render();
                })
                .fail(function(/*data*/){
                    w2ui.tldsInfoForm.error("Delete Tasklist Definition failed.");
                    return;
                });
            }
        },
    });

    //------------------------------------------------------------------------
    //  payorstmtlayout - The layout to contain the stmtForm and tlsDetailGrid
    //               top  - stmtForm
    //               main - tlsDetailGrid
    //------------------------------------------------------------------------
    $().w2layout({
        name: 'tldLayout',
        padding: 0,
        panels: [
            { type: 'left',    size: 0,     hidden: true },
            { type: 'top',     size: '30%', hidden: false, content: 'top',  resizable: true, style: app.pstyle },
            { type: 'main',    size: '70%', hidden: false, content: 'main', resizable: true, style: app.pstyle },
            { type: 'preview', size: 0,     hidden: true,  content: 'PREVIEW'  },
            { type: 'bottom',  size: 50,    hidden: false, content: 'bottom', resizable: false, style: app.pstyle },
            { type: 'right',   size: 0,     hidden: true }
        ]
    });
};

window.finishTLDForm = function () {
    w2ui.tldLayout.content('top',   w2ui.tldsInfoForm);
    w2ui.tldLayout.content('main',  w2ui.tldsDetailGrid);
    w2ui.tldLayout.content('bottom',w2ui.tldsCloseForm);
};

//-----------------------------------------------------------------------------
// saveTaskListDefinition - save the task list definition, called from 
// multiple places.
// 
// @params
//     hide = boolean - indicates whether or not the dialog should close
//            after the save is complete
//     reloadTldsInfo = boolean - indicates if we need to explicitly reload
//            w2ui.tldsInfoForm.record
//  
// @returns
//     0 = no errors, just continue
//     1 = error message was popped up.
//  
//-----------------------------------------------------------------------------
window.saveTaskListDefinition = function (hide, reloadTldsInfo) {
    var tmp         = w2ui.tldsInfoForm.record;
    var y           = tmp.Cycle.id;
    tmp.Cycle       = y; // we don't want the struct, we just want the ID
    tmp.Epoch       = localtimeToUTC(tmp.Epoch);
    tmp.EpochDue    = localtimeToUTC(tmp.EpochDue);
    tmp.EpochPreDue = localtimeToUTC(tmp.EpochPreDue);
    if (tmp.Epoch.length === 0) {
        tmp.Epoch = TLD.TIME0;
    }
    if (tmp.EpochDue.length === 0) {
        tmp.EpochDue = TLD.TIME0;
    }
    if (tmp.EpochPreDue.length === 0) {
        tmp.EpochPreDue = TLD.TIME0;
    }
    if (tmp.Name.length === 0) {
        w2alert('Please name the task list definition, then try again.');
        return 1;
    }
    var tl = {
        cmd: "save",
        record: tmp,
    };
    var dat = JSON.stringify(tl);
    var url = '/v1/tld/' + w2ui.tldsInfoForm.record.BID + '/' + w2ui.tldsInfoForm.record.TLDID;
    $.post(url,dat)
    .done(function(data) {
        if (data.status === "error") {
            w2ui.tldsInfoForm.error(w2utils.lang(data.message));
            return;
        }
        if (hide) {
            w2ui.toplayout.hide('right',true);
        }
        w2ui.tldsGrid.render();
        w2ui.tldsInfoForm.record.TLDID = data.recid;
        if (reloadTldsInfo) {
            w2ui.tldsInfoForm.reload();
        }
    })
    .fail(function(/*data*/){
        w2ui.tldsInfoForm.error("Save Tasklist failed.");
        return;
    });
    return 0;
};

//-----------------------------------------------------------------------------
// openTaskDescForm - Bring up the task descriptor edit form
// 
// @params
//     bid = business id
//     tdid = task descriptor id
//  
// @returns
//  
//-----------------------------------------------------------------------------
window.openTaskDescForm = function (bid,tdid) {
    TLD.formBtnsDisabled = true;
    if (tdid > 0) {
        w2ui.taskDescForm.url = '/v1/td/' + bid + '/' + tdid;
        w2ui.taskDescForm.request();
    } else {
        w2ui.taskDescForm.url = '';
        w2ui.taskDescForm.record = getTDInitRecord(bid, tdid, null);
    }
    var n = '' + tdid;
    w2ui.taskDescForm.header = 'Task Descriptor  ('+ (n === '0' ? 'new':n)  + ')';
    w2ui.tldLayout.content('right', w2ui.taskDescForm);
    w2ui.tldLayout.sizeTo('right', TLD.TaskDescWidth);
    w2ui.tldLayout.show('right');
    w2ui.tldLayout.render();
};

//-----------------------------------------------------------------------------
// closeTaskDescForm - Close the task descriptor edit form
// 
// @params
//     bid = business id
//     tdid = task descriptor id
//  
// @returns
//  
//-----------------------------------------------------------------------------
window.closeTaskDescForm = function (bid,tdid) {
    w2ui.tldLayout.hide('right');
    w2ui.tldLayout.sizeTo('right', 0);
    w2ui.tldsDetailGrid.render();
    TLD.formBtnsDisabled = false;
};

//-----------------------------------------------------------------------------
// setToTLDForm - enable the Task List Definition form.  Also, set
//                the forms url and request data from the server
// @params
//   bid = business id (or the BUD)
//    id = Task List TLID
// d1,d2 = date range to use
//-----------------------------------------------------------------------------
window.setToTLDForm = function (bid, id, d1,d2) {
    if (id > 0) {
        w2ui.tldsGrid.url = '/v1/tlds/' + bid;                    // the grid of tasklist Defs
        w2ui.tldsDetailGrid.url = '/v1/tds/' + bid + '/' + id; // the tasks associated with the selected tasklistDefinition
        w2ui.tldsInfoForm.url = '/v1/tld/' + bid + '/' + id;      // the tasklist def details
        w2ui.tldsInfoForm.postData = {
            searchDtStart: d1,
            searchDtStop: d2,
        };
        w2ui.tldsInfoForm.header = 'Task List Definition ' + id;
        w2ui.tldsInfoForm.request();

    } else {
        w2ui.tldsDetailGrid.url = '';
        w2ui.tldsDetailGrid.records = [];
        w2ui.tldsInfoForm.header = 'Task List Definition (new)';
        w2ui.tldsInfoForm.record = getTLDInitRecord(bid, null);
    }
    w2ui.toplayout.content('right', w2ui.tldLayout);
    w2ui.toplayout.show('right', true);
    w2ui.toplayout.sizeTo('right', 700);
    w2ui.toplayout.render();
    app.new_form_rec = false;  // mark as record exists
    app.form_is_dirty = false; // mark as no changes yet
};
//-----------------------------------------------------------------------------
// setTaskDescButtonsState - set the form Save / Delete button state to 
//                       the value in TL.
// 
// @params
//  
// @returns 
//  
//-----------------------------------------------------------------------------
window.setTaskDescButtonsState = function() {
    $(w2ui.tldsCloseForm.box).find("button[name=save]").prop( "disabled", TLD.formBtnsDisabled );
    $(w2ui.tldsCloseForm.box).find("button[name=delete]").prop( "disabled", TLD.formBtnsDisabled );
};
