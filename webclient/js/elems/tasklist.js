"use strict";
/*global
    w2ui, $, app, console, w2utils,
    form_dirty_alert, addDateNavToToolbar, w2uiDateControlString,
    dateFromString, taskDateRender, setToTLForm,
    taskFormDueDate,taskCompletionChange,taskFormDoneDate,
    openTaskForm,setInnerHTML,w2popup,ensureSession,dtFormatISOToW2ui,
    createNewTaskList, getBUDfromBID, exportItemReportPDF, exportItemReportCSV,
    popupNewTaskListForm, getTLDs, getCurrentBID, getNewTaskListRecord,
    closeTaskForm, setTaskButtonsState, renderTaskGridDate, localtimeToUTC, TLD,
*/

var TL = {
    TaskWidth: 500,             // width of task form over tasklist grid
    formBtnsDisabled: false,    // indicates whether the tasklist save/delete buttons should be on/off
    TIME0: '1/1/1970',          // value of time if date control returns an empty string
};

window.getNewTaskListRecord = function (bid) {
    var rec = {
        BID: bid,
        TLDID: 0,
        Name: '',
        Pivot: w2uiDateControlString(new Date()),
    };
    return rec;
};

window.buildTaskListElements = function () {
    //------------------------------------------------------------------------
    //          tlsGrid  -  TASK LISTS in the date range
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'tlsGrid',
        url: '/v1/tls',
        multiSelect: false,
        postData: {searchDtStart: app.D1, searchDtStop: app.D2},
        show: {
            toolbar         : true,
            footer          : true,
            toolbarAdd      : true,    // indicates if toolbar add new button is visible
            toolbarDelete   : false,   // indicates if toolbar delete button is visible
            toolbarSave     : false,   // indicates if toolbar save button is visible
            selectColumn    : false,
            expandColumn    : false,
            toolbarEdit     : false,
            toolbarSearch   : false,
            toolbarInput    : false,
            searchAll       : false,
            toolbarReload   : true,
            toolbarColumns  : true,
        },
        columns: [
            {field: 'recid',     hidden: true,  caption: 'recid',             size: '40px',  sortable: true},
            {field: 'BID',       hidden: true,  caption: 'BID',               size: '40px',  sortable: true},
            {field: 'Name',      hidden: false, caption: 'Name',              size: '110px', sortable: true},
            {field: 'DtPreDue',  hidden: false, caption: 'Pre Due',           size: '150px', sortable: true,
                render: function (record, index, col_index) { if (typeof record === "undefined") {return '';} return taskDateRender(record.DtPreDue); }
            },
            {field: 'DtPreDone', hidden: false, caption: 'Pre Due completed', size: '150px', sortable: true,
                render: function (record, index, col_index) { if (typeof record === "undefined") {return '';} return taskDateRender(record.DtPreDone); }
            },
            {field: 'DtDue',     hidden: false, caption: 'Due',               size: '150px', sortable: true,
                render: function (record, index, col_index) {if (typeof record === "undefined") {return ''; } return taskDateRender(record.DtDue); }
            },
            {field: 'DtDone',    hidden: false, caption: 'Due completed',     size: '150px', sortable: true,
                render: function (record, index, col_index) { if (typeof record === "undefined") {return '';} return taskDateRender(record.DtDone); }
            },
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
                        console.log( 'BID = ' + rec.BID + ',   TLID = ' + rec.TLID);
                        setToTLForm(rec.BID, rec.TLID, app.D1, app.D2);
                    };
                form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
            };
        },
        onAdd: function(event) {
            event.onComplete = function () {
                var bid = getCurrentBID();
                createNewTaskList(bid);
            };
        },
    });

    addDateNavToToolbar('tls'); // "Grid" is appended to the 

    //------------------------------------------------------------------------
    //  tlsInfoForm
    //------------------------------------------------------------------------
    $().w2form({
        name: 'tlsInfoForm',
        style: 'border: 0px; background-color: transparent;',
        header: 'Task List',
        url: '/v1/tl',
        formURL: '/webclient/html/formtl.html',
        toolbar: {
            items: [
                // { id: 'btnNotes', type: 'button', icon: 'far fa-sticky-note' },

                { type: 'button', id: 'csvexport', icon: 'fas fa-table',    tooltip: 'export to CSV' },
                { type: 'button', id: 'pdfexport', icon: 'far fa-file-pdf', tooltip: 'export to PDF' },
                { type: 'spacer', id: 'bt3'  },
                { type: 'button', id: 'btnClose',  icon: 'fas fa-times' },
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
                                closeTaskForm();
                                w2ui.toplayout.hide('right',true);
                                w2ui.tlsGrid.render();
                            };
                        form_dirty_alert(yes_callBack, no_callBack);
                        break;
                    case 'csvexport':
                        exportItemReportCSV("RPTtl", w2ui.tlsInfoForm.record.TLID, app.D1, app.D2);
                        break;
                    case 'pdfexport':
                        exportItemReportPDF("RPTtl", w2ui.tlsInfoForm.record.TLID, app.D1, app.D2);
                        break;
                    }
                };
            },
        },
        fields: [
            { field: 'recid',        type: 'int',       required: false },
            { field: 'TLID',         type: 'int',       required: false },
            { field: 'PTLID',        type: 'int',       required: false },
            { field: 'BID',          type: 'int',       required: false },
            { field: 'BUD',          type: 'list',      required: true, options: {items: app.businesses} },
            { field: 'Name',         type: 'text',      required: true },
            { field: 'Cycle',        type: 'list',      required: true, options: {items: app.w2ui.listItems.cycleFreq}, },
            { field: 'DtDue',        type: 'datetime',  required: false },
            { field: 'DtPreDue',     type: 'datetime',  required: false },
            { field: 'DtDone',       type: 'datetime',  required: false },
            { field: 'DtPreDone',    type: 'datetime',  required: false },
            { field: 'FLAGS',        type: 'int',       required: false },
            { field: 'DoneUID',      type: 'int',       required: false },
            { field: 'DoneName',     type: 'text',      required: false },
            { field: 'PreDoneUID',   type: 'int',       required: false },
            { field: 'PreDoneName',  type: 'text',      required: false },
            { field: 'EmailList',    type: 'text',      required: false },
            { field: 'Comment',      type: 'text',      required: false },
            { field: 'CreateTS',     type: 'date',      required: false },
            { field: 'CreateBy',     type: 'int',       required: false },
            { field: 'LastModTime',  type: 'date',      required: false },
            { field: 'LastModBy',    type: 'int',       required: false },
            { field: 'TZOffset',     type: 'int',       required: false },
            { field: 'ChkDtDue',     type: 'checkbox',  required: false },
            { field: 'ChkDtDone',    type: 'checkbox',  required: false },
            { field: 'ChkDtPreDue',  type: 'checkbox',  required: false },
            { field: 'ChkDtPreDone', type: 'checkbox',  required: false },
            { field: 'TZOffset',     type: 'int',       required: false },
        ],
        onRefresh: function(event) {
            event.onComplete = function(event) {
                var r = w2ui.tlsInfoForm.record;
                if (typeof r.DtPreDue === "undefined") {
                    return;
                }
                r.ChkDtPreDue  = taskFormDueDate(r.DtPreDue,  r.ChkDtPreDue,'sDtPreDue','no pre-due date');
                r.ChkDtDue     = taskFormDueDate(r.DtDue,     r.ChkDtDue,   'sDtDue',   'no due date');
                r.ChkDtDone    = taskFormDoneDate(r.ChkDtDone,    r.DtDone,   r.DtDue,      r.ChkDtDone,    r.DoneUID,    r.DoneName,    'sDtDone',   'tlDoneName',    'tlOverdue');
                r.ChkDtPreDone = taskFormDoneDate(r.ChkDtPreDone, r.DtPreDone,r.DtPreDue,   r.ChkDtPreDone, r.PreDoneUID, r.PreDoneName, 'sDtPreDone','tlPreDoneName', 'tlPreOverdue');
            };
        },
        onLoad: function(event) {
            event.onComplete = function(event) {
                var f = w2ui.tlsInfoForm;
                var r = f.record;

                // translate dates into a format that w2ui understands
                r.DtPreDue  = dtFormatISOToW2ui(r.DtPreDue);
                r.DtDue     = dtFormatISOToW2ui(r.DtDue);
                r.DtPreDone = dtFormatISOToW2ui(r.DtPreDone);
                r.DtDone    = dtFormatISOToW2ui(r.DtDone);

                if (r.DtDone === "") {
                    r.DtDone = new Date();
                }
                if (r.DtPreDone === "") {
                    r.DtPreDone = new Date();
                }

                // now enable/disable as needed
                $(f.box).find("input[name=DtDue]").prop( "disabled", !r.ChkDtDue );
                $(f.box).find("input[name=DtPreDue]").prop( "disabled", !r.ChkDtPreDue );
            };
        },
        onChange: function(event) {
            event.onComplete = function() {
                var f = w2ui.tlsInfoForm;
                var r = f.record;
                var s = '';
                if (event.target === "ChkDtPreDone") {
                    taskCompletionChange(event.value_new,"sDtPreDone");
                    r.DtPreDone = new Date();
                } else if (event.target === "ChkDtDone") {
                    taskCompletionChange(event.value_new,"sDtDone");
                    r.DtDone = new Date();
                }
            };
        },

    });

    // addDateNavToToolbar('tlsInfoForm');

    //------------------------------------------------------------------------
    //  tlsTaskGrid  -  lists all the assessments and receipts for
    //                  the selected Rental Agreement from the stmtGrid
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'tlsDetailGrid',
        url: '/v1/tasks/',
        multiSelect: false,
        postData: {searchDtStart: app.D1, searchDtStop: app.D2, Bool1: app.PayorStmtExt},
        columns: [
            {field: 'recid',        caption: 'recid',       size: '35px', sortable: true, hidden: true},
            { field: 'TID',         caption: 'TID',         size: '35px', sotrable: true, hidden: true},
            { field: 'BID',         caption: 'BID',         size: '35px', sotrable: true, hidden: true},
            { field: 'TLID',        caption: 'TLID',        size: '35px', sotrable: true, hidden: true},
            { field: 'TaskName',    caption: 'Name',        size: '200px',sotrable: true, hidden: false},
            { field: 'Worker',      caption: 'Worker',      size: '75px', sotrable: true, hidden: false},
            { field: 'DtPreDue',    caption: 'DtPreDue',    size: '80px', sotrable: true, hidden: false,
                render: function(record, index, col_index) {return renderTaskGridDate(record.DtPreDue); }
            },
            { field: 'DtPreDone',   caption: 'DtPreDone',   size: '80px', sotrable: true, hidden: false,
                render: function(record, index, col_index) {return renderTaskGridDate(record.DtPreDone); }
            },
            { field: 'DtDue',       caption: 'DtDue',       size: '80px', sotrable: true, hidden: false,
                render: function(record, index, col_index) {return renderTaskGridDate(record.DtDue); }
            },
            { field: 'DtDone',      caption: 'DtDone',      size: '80px', sotrable: true, hidden: false,
                render: function(record, index, col_index) {
                    return renderTaskGridDate(record.DtDone); 
                }
            },
            { field: 'FLAGS',       caption: 'FLAGS',       size: '35px', sotrable: true, hidden: true},
            { field: 'DoneUID',     caption: 'DoneUID',     size: '35px', sotrable: true, hidden: true},
            { field: 'PreDoneUID',  caption: 'PreDoneUID',  size: '35px', sotrable: true, hidden: true},
            { field: 'TaskComment', caption: 'Comment',     size: '35px', sotrable: true, hidden: true},
            { field: 'LastModTime', caption: 'LastModTime', size: '35px', sotrable: true, hidden: true},
            { field: 'LastModBy',   caption: 'LastModBy',   size: '35px', sotrable: true, hidden: true},
            { field: 'CreateTS',    caption: 'CreateTS',    size: '35px', sotrable: true, hidden: true},
            { field: 'CreateBy',    caption: 'CreateBy',    size: '35px', sotrable: true, hidden: true},
        ],
        onClick: function(event) {
            event.onComplete = function (event) {
                var r = w2ui.tlsDetailGrid.records[event.recid];
                console.log( 'detail clicked: v1/tasks/' + r.BID + '/'+ r.TID);
                openTaskForm(r.BID,r.TID);
            };
        },
        onRender: function (event) {
            event.onComplete = function (event) {
                setTaskButtonsState();
            };
        },
        onReload: function(event) {
            event.onComplete = function (event) {
                setTaskButtonsState();
            };
        },
    });

    //------------------------------------------------------------------------
    //  taskForm
    //------------------------------------------------------------------------
    $().w2form({
        name: 'taskForm',
        style: 'border: 0px; background-color: transparent;',
        formURL: '/webclient/html/formtask.html',
        url: '/v1/task',
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
                                closeTaskForm();
                                w2ui.tlLayout.render();
                            };
                        form_dirty_alert(yes_callBack, no_callBack);
                        break;
                    }
                };
            },
        },
        fields: [
            { field: 'recid',        type: 'text',     required: false },
            { field: 'TID',          type: 'text',     required: false },
            { field: 'BID',          type: 'text',     required: false },
            { field: 'TLID',         type: 'text',     required: false },
            { field: 'TaskName',     type: 'text',     required: true  },
            { field: 'Worker',       type: 'text',     required: false },
            { field: 'DtDue',        type: 'text',     required: false },
            { field: 'DtPreDue',     type: 'text',     required: false },
            { field: 'DtDone',       type: 'text',     required: false },
            { field: 'DtPreDone',    type: 'text',     required: false },
            { field: 'FLAGS',        type: 'text',     required: false },
            { field: 'DoneUID',      type: 'text',     required: false },
            { field: 'PreDoneUID',   type: 'text',     required: false },
            { field: 'TaskComment',  type: 'text',     required: false },
            { field: 'LastModTime',  type: 'date',     required: false },
            { field: 'LastModBy',    type: 'int',      required: false },
            { field: 'CreateTS',     type: 'date',     required: false },
            { field: 'CreateBy',     type: 'int',      required: false },
            { field: 'TZOffset',     type: 'int',      required: false },
            { field: 'ChkDtDue',     type: 'checkbox', required: false },
            { field: 'ChkDtDone',    type: 'checkbox', required: false },
            { field: 'ChkDtPreDue',  type: 'checkbox', required: false },
            { field: 'ChkDtPreDone', type: 'checkbox', required: false },
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
                var f = w2ui.taskForm;
                var r = f.record;
                var d = {cmd: "save", record: r};
                r.DtDone = localtimeToUTC(r.DtDone);
                r.DtPreDone = localtimeToUTC(r.DtPreDone);
                if (r.DtDone.length === 0) {
                    r.DtDone = TLD.TIME0;
                }
                if (r.DtPreDone.length === 0) {
                    r.DtPreDone = TLD.TIME0;
                }
                r.TZOffset = app.TZOffset;
                var dat=JSON.stringify(d);
                f.url = '/v1/task/' + r.BID + '/' + r.TID;
                $.post(f.url,dat)
                .done(function(data) {
                    if (data.status === "error") {
                        f.error(w2utils.lang(data.message));
                        return;
                    }
                    //w2ui.tlsDetailGrid.url = '/v1/tl/'
                    w2ui.tlsDetailGrid.reload();
                    // w2popup.close();
                    closeTaskForm();
                    setTaskButtonsState();
                })
                .fail(function(/*data*/){
                    f.error("Save Tasklist failed.");
                    return;
                });
            },
        },
       onRefresh: function(event) {
            // var f = this;
            event.onComplete = function(event) {
                var r = w2ui.taskForm.record;
                if (typeof r.DtPreDue === "undefined") {
                    return;
                }
                r.ChkDtPreDue  = taskFormDueDate(r.DtPreDue,      r.ChkDtPreDue,'tskDtPreDue',  'no pre-due date'               );
                r.ChkDtDue     = taskFormDueDate(r.DtDue,         r.ChkDtDue,   'tskDtDue',     'no due date'                   );
                r.ChkDtPreDone = taskFormDoneDate(r.ChkDtPreDone, r.DtPreDone,  r.DtPreDue,   r.ChkDtPreDone, r.PreDoneUID, r.TaskPreDoneName, 'tskDtPreDone', 'tskPreDoneName', 'tskPreOverdue');
                r.ChkDtDone    = taskFormDoneDate(r.ChkDtDone,    r.DtDone,     r.DtDue,      r.ChkDtDone,    r.DoneUID,    r.TaskDoneName,    'tskDtDone',    'tskDoneName',    'tskOverdue'   );
            };
        },
        onChange: function(event) {
            event.onComplete = function() {
                var s = '';
                if (event.target === "ChkDtPreDone") {
                    taskCompletionChange(event.value_new,"tskDtPreDone");
                } else if (event.target === "ChkDtDone") {
                    taskCompletionChange(event.value_new,"tskDtDone");
                }
            };
        },
        onRender: function(event) {
            event.onComplete = function(event) {
                setTaskButtonsState();
            };
        },
        onLoad: function(event) {
            event.onComplete = function(event) {
                var f = w2ui.taskForm;
                var r = f.record;

                // translate dates into a format that w2ui understands
                r.DtPreDue  = dtFormatISOToW2ui(r.DtPreDue);
                r.DtDue     = dtFormatISOToW2ui(r.DtDue);
                r.DtPreDone = dtFormatISOToW2ui(r.DtPreDone);
                r.DtDone    = dtFormatISOToW2ui(r.DtDone);

                if (r.DtDone === "") {
                    r.DtDone = new Date();
                }
                if (r.DtPreDone === "") {
                    r.DtPreDone = new Date();
                }
            };
        },
    });

    //------------------------------------------------------------------------
    //  tlsCloseForm
    //------------------------------------------------------------------------
    $().w2form({
        name: 'tlsCloseForm',
        style: 'border: 0px; background-color: transparent;',
        formURL: '/webclient/html/formtlclose.html',
        url: '',
        fields: [],
        actions: {
            save: function(target, data){
                // getFormSubmitData(data.postData.record);
                var r = w2ui.tlsInfoForm.record;

                //------------------------------------------------
                // due and done times do not matter, server looks
                // at the check values and sets time based on its
                // own local clock. We do not accept these times
                // from a client.
                //------------------------------------------------
                r.DtDone = TLD.TIME0;
                r.DtPreDone = TLD.TIME0;
                r.TZOffset = app.TZOffset;
                r.DtDue = localtimeToUTC(r.DtDue);
                r.DtPreDue = localtimeToUTC(r.DtPreDue);
                var cycle = r.Cycle.id;
                r.Cycle = cycle;

                var tl = {
                    cmd: "save",
                    record: r,
                };
                var dat=JSON.stringify(tl);
                var url='/v1/tl/' + w2ui.tlsInfoForm.record.BID + '/' + w2ui.tlsInfoForm.record.TLID;
                $.post(url,dat)
                .done(function(data) {
                    if (data.status === "error") {
                        w2ui.tlsInfoForm.error(w2utils.lang(data.message));
                        return;
                    }
                    w2ui.toplayout.hide('right',true);
                    w2ui.tlsGrid.render();
                })
                .fail(function(/*data*/){
                    w2ui.tlsInfoForm.error("Save Tasklist failed.");
                    return;
                });
            },
            delete: function(target,data) {
                var tl = {
                    cmd: "delete",
                };
                var dat=JSON.stringify(tl);
                var url='/v1/tl/' + w2ui.tlsInfoForm.record.BID + '/' + w2ui.tlsInfoForm.record.TLID;
                $.post(url,dat)
                .done(function(data) {
                    if (data.status === "error") {
                        w2ui.tlsInfoForm.error(w2utils.lang(data.message));
                        return;
                    }
                    w2ui.toplayout.hide('right',true);
                    w2ui.tlsGrid.render();
                })
                .fail(function(/*data*/){
                    w2ui.tlsInfoForm.error("Save Tasklist failed.");
                    return;
                });
            },
        },
    });

    //------------------------------------------------------------------------
    //  payorstmtlayout - The layout to contain the stmtForm and tlsDetailGrid
    //               top  - stmtForm
    //               main - tlsDetailGrid
    //------------------------------------------------------------------------
    $().w2layout({
        name: 'tlLayout',
        padding: 0,
        panels: [
            { type: 'left',    size: 0,     hidden: true },
            { type: 'top',     size: '35%', hidden: false, content: 'top',  resizable: true, style: app.pstyle },
            { type: 'main',    size: '65%', hidden: false, content: 'main', resizable: true, style: app.pstyle },
            { type: 'preview', size: 0,     hidden: true,  content: 'PREVIEW'  },
            { type: 'bottom',  size: 50,    hidden: false, content: 'bottom', resizable: false, style: app.pstyle },
            { type: 'right',   size: 0,     hidden: true }
        ]
    });

    //------------------------------------------------------------------------
    //  newTaskListForm
    //------------------------------------------------------------------------
    $().w2form({
        name: 'newTaskListForm',
        style: 'border: 0px; background-color: transparent;',
        formURL: '/webclient/html/formnewtl.html',
        url: '/v1/tl',
        fields: [
            { field: 'BID',      type: 'int',  required: false },
            { field: 'TLDID',    type: 'int',  required: false },
            { field: 'Name',     type: 'list', required: true, options:  {items: [], selected: {}},  },
            { field: 'Pivot',    type: 'date', required: true },
        ],
        actions: {
            save: function(target, data){
                var f = w2ui.newTaskListForm;
                var r = f.record;
                f.url = '/v1/tl/' + r.BID + '/0';
                var s = r.Name.text;
                r.TLDID = r.Name.id;
                r.Name = s;
                r.Pivot = localtimeToUTC(r.Pivot);
                r.Timezone = app.timezone;
                var params = {cmd: 'save', formname: f.name, record: r };

                var dat = JSON.stringify(params);
                var BID = r.BID;

                // submit request
                $.post(f.url, dat, null, "json")
                .done(function(data) {
                    if (data.status != "success") {
                        return;
                    }
                    w2ui.tlsGrid.reload();
                    var tlid = data.recid;
                    setToTLForm(BID, tlid, app.D1, app.D2);                    
                    w2popup.close();
                })
                .fail(function(/*data*/){
                    console.log("Payor Fund Allocation failed.");
                });

            },
        },
        // onLoad: function(event) {
        //     event.onComplete = function(event) {

        //     };
        // },
        onRefresh: function(event) {
            // var f = this;
            // event.onComplete = function(event) {
            // };
        },
        onChange: function(event) {
            // event.onComplete = function() {
            // };
        },
    });

};

window.finishTaskListForm = function () {
    w2ui.tlLayout.content('top',   w2ui.tlsInfoForm);
    w2ui.tlLayout.content('main',  w2ui.tlsDetailGrid);
    w2ui.tlLayout.content('bottom',w2ui.tlsCloseForm);
};

//-----------------------------------------------------------------------------
// createNewTaskList - pop up dialog where the user can select one of the
//      defined TaskListDefinitions and set the Pivot date. Then create a 
//      new TaskList, update the grid, and bring it up in the edit form
// 
// @params
//  
// @returns 
//  
//-----------------------------------------------------------------------------
window.createNewTaskList = function (bid) {
    //-------------------------------------------------------
    // First get the latest list of TaskListDefinitions...
    //-------------------------------------------------------
    getTLDs(bid,popupNewTaskListForm);
};


//-----------------------------------------------------------------------------
// renderTaskGridDate - if the year is 1970 or less return '', otherwise
//      return the date string (ds).
// 
// @params - ds = date string
//  
// @return date string or ''
//  
//-----------------------------------------------------------------------------
window.renderTaskGridDate = function (ds) {
    var d = new Date(ds);
    if (d.getFullYear() > 1970) {
        return ds;
    } 
    return '';
};


//-----------------------------------------------------------------------------
// getTLDs - return the promise object of request to get latest
//           TaskListDefinitions for given BID.
//           It updates the "app.TaskListDefinitions" variable for requested BUD
//
// @params  - BID     : Business ID (expected current one)
//          - handler : Business Unit Designation
// @return  - promise object from $.get
//-----------------------------------------------------------------------------
window.getTLDs = function (BID,handler) {
    var BUD = getBUDfromBID(BID);

    // return promise
    return $.get("/v1/uival/" + BID + "/app.TaskListDefinitions", null, null, "json")
            .done(function(data) {
                // if it doesn't meet this condition, then save the data
                if (app.TaskListDefinitions === null ) {
                    app.TaskListDefinitions = [];
                }
                app.TaskListDefinitions[BUD] = data[BUD];
                var f = w2ui.newTaskListForm;
                f.get('Name').options.items = app.TaskListDefinitions[BUD];
                f.record = getNewTaskListRecord(BID);
                f.record.TLID = app.TaskListDefinitions[BUD][0].id;
                f.record.Name = app.TaskListDefinitions[BUD][0];
                f.refresh();
                handler(BID);
            });
};

//-----------------------------------------------------------------------------
// popupNewTaskListForm - Bring up the task edit form
// 
// @params
//     bid = business id
//  
// @returns
//  
//-----------------------------------------------------------------------------
window.popupNewTaskListForm = function (bid) {
    w2ui.newTaskListForm.url = '/v1/tl/' + bid + '/0';
    // w2ui.newTaskListForm.request();
    $().w2popup('open', {
        title   : 'New Task List',
        body    : '<div id="form" style="width: 100%; height: 100%;"></div>',
        style   : 'padding: 15px 0px 0px 0px',
        width   : 450,
        height  : 220,
        showMax : true,
        onToggle: function (event) {
            $(w2ui.newTaskListForm.box).hide();
            event.onComplete = function () {
                $(w2ui.newTaskListForm.box).show();
                w2ui.newTaskListForm.resize();
            };
        },
        onOpen: function (event) {
            event.onComplete = function () {
                $('#w2ui-popup #form').w2render('newTaskListForm');
            };
        }
    });
};
//-----------------------------------------------------------------------------
// setToTLForm -  enable the Statement form in toplayout.  Also, set
//                the forms url and request data from the server
// @params
//   bid = business id (or the BUD)
//    id = Task List TLID
// d1,d2 = date range to use
//-----------------------------------------------------------------------------
window.setToTLForm = function (bid, id, d1,d2) {
    if (id > 0) {
        w2ui.tlsGrid.url = '/v1/tls/' + bid;                    // the grid of tasklists
        w2ui.tlsDetailGrid.url = '/v1/tasks/' + bid + '/' + id; // the tasks associated with the selected tasklist
        w2ui.tlsInfoForm.url = '/v1/tl/' + bid + '/' + id;      // the tasklist details
        w2ui.tlsInfoForm.postData = {
            searchDtStart: d1,
            searchDtStop: d2,
        };
        w2ui.tlsInfoForm.header = 'Task List ' + id;
        w2ui.tlsInfoForm.request();

        w2ui.toplayout.content('right', w2ui.tlLayout);
        w2ui.toplayout.show('right', true);
        w2ui.toplayout.sizeTo('right', 600);
        w2ui.toplayout.render();
        app.new_form_rec = false;  // mark as record exists
        app.form_is_dirty = false; // mark as no changes yet
    }
};

//-----------------------------------------------------------------------------
// taskDateRender - If the date is less than year 2000 then return a blank
//                otherwise return the date as a string.
// @params
//   y - the date to be printed
// @returns
//   the string to print
//-----------------------------------------------------------------------------
window.taskDateRender = function (x) {
    if (x === null) {
        return '';
    }
    var y;
    if (typeof x == "string"){
        y = dateFromString(x);
    }
    if (typeof x == "object") {
        y = x;
    }
    var yr = y.getFullYear();
    if ( yr <= 1970) {
        return '';
    }
    // return dtTextRender(x,0,0);
    return dtFormatISOToW2ui(x);
};

//-----------------------------------------------------------------------------
// openTaskForm - Bring up the task edit form
// 
// @params
//     bid = business id
//     tid = task id
//  
// @returns
//  
//-----------------------------------------------------------------------------
window.openTaskForm = function (bid,tid) {
    TL.formBtnsDisabled = true;
    w2ui.taskForm.url = '/v1/task/' + bid + '/' + tid;
    w2ui.taskForm.request();
    var n = '' + tid;
    w2ui.taskForm.header = 'Task ('+ (n === '0' ? 'new':n)  + ')';
    w2ui.tlLayout.content('right', w2ui.taskForm);
    w2ui.tlLayout.sizeTo('right', TL.TaskWidth);
    w2ui.tlLayout.show('right');
    w2ui.tlLayout.render();
};

//-----------------------------------------------------------------------------
// closeTaskForm - Close the task descriptor edit form
// 
// @params
//     bid = business id
//     tdid = task descriptor id
//  
// @returns
//  
//-----------------------------------------------------------------------------
window.closeTaskForm = function (bid,tdid) {
    w2ui.tlLayout.hide('right');
    w2ui.tlLayout.sizeTo('right', 0);
    w2ui.tlsDetailGrid.render();
    TL.formBtnsDisabled = false;
};

//-----------------------------------------------------------------------------
// setInnerHTML - form formatting.  saves a few lines by handling the null case.
// 
// @params
//      id  = html element id for string update
//      s   = string for no date value 
//  
// @returns 
//  
//-----------------------------------------------------------------------------
window.setInnerHTML = function (id,s) {
    var e = document.getElementById(id);
    if (e != null) {
        e.innerHTML = s;
    }
};

//-----------------------------------------------------------------------------
// taskFormDueDate - form formatting
// 
// @params
//       dt = datetime string
//       b  = boolean check box value (false = unchecked)
//      id  = html element id for string update
//      txt = string for no date value 
//  
// @returns 
//      updated value for ChkDt...  true if year >= 2000
//  
//-----------------------------------------------------------------------------
window.taskFormDueDate = function (dt,b,id,txt) {
    if (dt !== null && dt.length > 0) {
        var y = dateFromString(dt);
        var s = '';
        b = y.getFullYear() >= 2000;
        if (b) {
            s = taskDateRender(dt); 
        } else {
            s = txt;
        }
        setInnerHTML(id,s);
    }
    return b;
};

//-----------------------------------------------------------------------------
// taskFormDoneDate - form formatting
//       r.ChkDtPreDone = taskFormDoneDate(r.ChkDtPreDone, r.DtPreDone,r.DtPreDue,   r.ChkDtPreDone, r.PreDoneUID, r.PreDoneName, 'sDtPreDone','tlPreDoneName', 'tlPreOverdue');
// @params
//      bDone = boolean indicates whether or not a done date has been supplied.
//              This value may change during editing. When the user checks
//              the box, the server will supply the done date.  If the box is
//              unchecked, the server will mark that no done date has been
//              supplied (thus the task is not completed)
//      sDtDone  = datetime string when the task was completed
//      sDtDue  = due datetime string - indicates when the task was due
//      
//      uid  = uid of user who marked this as done
//      name = name associated with uid
//      id   = html element id for string update
//      id2  = html area for name
//      id3  = string to indicate late
//  
// @returns 
//      updated value for ChkDt...  true if year >= 2000
//  
//-----------------------------------------------------------------------------
window.taskFormDoneDate = function (bDone,sDtDone,sDtDue,b,uid,name,id,id2,id3) {
    var strDoneDate = ""; // string for sDtDone

    if (typeof sDtDone== "string") {
        strDoneDate = sDtDone;
    } else if (typeof sDtDone == "object") {
        strDoneDate = sDtDone.toISOString();
    }

    if (strDoneDate !== null && strDoneDate.length > 0) {
        //--------------------------
        // id: date
        //--------------------------
        var y = new Date(strDoneDate);
        var s = '';
        if (bDone) {
            bDone = y.getFullYear() >= 2000;
            if (bDone) {
                s = taskDateRender(sDtDone);
            }
        }
        setInnerHTML(id,s);
        
        //--------------------------
        // id2: name indicator
        //--------------------------
        setInnerHTML(id2, (uid > 0) ? '('+name+')' : '');

        //--------------------------
        // id3: late indicator:
        //      if a Done date has been supplied, see if it's past the due date
        //      if no Done date has been supplied, see if the current time is past the due date
        //--------------------------
        var dtDue = dateFromString(sDtDue);
        var now = new Date();
        var dtDone = dateFromString(strDoneDate);
        s = '';
        if ( (bDone && dtDone > dtDue) || (!bDone && now > dtDue)) {
            if (dtDue.getFullYear() > 1970) {
                s = '<span style="color:#FC0D1B;">&nbsp;LATE</span>';
            }
        }
        if ( bDone && dtDone <= dtDue ) {
            s = '&#9989;';
        }
        setInnerHTML(id3,s);
    }
    return bDone;
};

//-----------------------------------------------------------------------------
// taskCompletionChange - form formatting
// 
// @params
//       b  = boolean check box value (false = unchecked)
//      id  = html element id for string update
//  
// @returns 
//      updated value for ChkDt...  true if year >= 2000
//  
//-----------------------------------------------------------------------------
window.taskCompletionChange = function (b,id) {
    var s;
    if (b) { // marked as complete?
        s = '<span style="color:blue;">mark as completed when Save is clicked</span>';
    } else {
        s = '<span style="color:blue;">mark as not completed when Save is clicked</span>';
    }
    setInnerHTML(id,s);
};

//-----------------------------------------------------------------------------
// setTaskButtonsState - set the form Save / Delete button state to 
//                       the value in TL.
// 
// @params
//  
// @returns 
//  
//-----------------------------------------------------------------------------
window.setTaskButtonsState = function() {
    $(w2ui.tlsCloseForm.box).find("button[name=save]").prop( "disabled", TL.formBtnsDisabled );
    $(w2ui.tlsCloseForm.box).find("button[name=delete]").prop( "disabled", TL.formBtnsDisabled );
};
