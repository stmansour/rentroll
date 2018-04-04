"use strict";
/*global
    GridMoneyFormat, number_format, w2ui, $, app, console,setToTLDForm,
    form_dirty_alert, addDateNavToToolbar, w2utils, w2popup, taskDateRender,
    popupTaskDescForm
*/

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
            {field: 'Name',      hidden: false, caption: 'Name',                    size: '110px', sortable: true},
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
            { field: 'recid',       type: 'int',  required: false },
            { field: 'TLDID',       type: 'int',  required: false },
            { field: 'BID',         type: 'int',  required: false },
            { field: 'BUD',         type: 'list', required: true, options: {items: app.businesses} },
            { field: 'Name',        type: 'text', required: true },
            { field: 'Cycle',       type: 'list', required: true, options: {items: app.w2ui.listItems.cycleFreq}, },
            { field: 'Epoch',       type: 'date', required: false },
            { field: 'EpochDue',    type: 'date', required: false },
            { field: 'EpochPreDue', type: 'date', required: false },
            { field: 'FLAGS',       type: 'int',  required: false },
            { field: 'CreateTS',    type: 'date', required: false },
            { field: 'CreateBy',    type: 'int',  required: false },
            { field: 'LastModTime', type: 'date', required: false },
            { field: 'LastModBy',   type: 'int',  required: false },
        ],
        onRefresh: function(event) {
            // var f = this;
            event.onComplete = function(event) {
                var r = w2ui.tldsInfoForm.record;
                if (typeof r.EpochPreDue === "undefined") {
                    return;
                }
            };
        },
        onChange: function(event) {
            event.onComplete = function() {
            };
        },

    });

    //------------------------------------------------------------------------
    //  tldsTaskGrid  -  lists all the assessments and receipts for
    //                  the selected Rental Agreement from the stmtGrid
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'tldsDetailGrid',
        url: '/v1/tds/',
        multiSelect: false,
        postData: {searchDtStart: app.D1, searchDtStop: app.D2, Bool1: app.PayorStmtExt},
        columns: [
            {field: 'recid',        caption: 'recid',       size: '35px', sortable: true, hidden: true},
            { field: 'TDID',        caption: 'TDID',        size: '35px', sotrable: true, hidden: true},
            { field: 'BID',         caption: 'BID',         size: '35px', sotrable: true, hidden: true},
            { field: 'TLDID',       caption: 'TLDID',       size: '35px', sotrable: true, hidden: true},
            { field: 'Name',        caption: 'Name',        size: '120px', sotrable: true, hidden: false},
            { field: 'Worker',      caption: 'Worker',      size: '75px', sotrable: true, hidden: false},
            { field: 'DtPreDue',    caption: 'DtPreDue',    size: '80px', sotrable: true, hidden: false},
            { field: 'DtPreDone',   caption: 'DtPreDone',   size: '80px', sotrable: true, hidden: false,
                render: function (record, index, col_index) { if (typeof record == "undefined") {return '';} return taskDateRender(record.DtPreDone); }
            },
            { field: 'DtDue',       caption: 'DtDue',       size: '80px', sotrable: true, hidden: false},
            { field: 'DtDone',      caption: 'DtDone',      size: '80px', sotrable: true, hidden: false,
                render: function (record, index, col_index) { if (typeof record == "undefined") {return '';} return taskDateRender(record.DtDone); }
            },
            { field: 'FLAGS',       caption: 'FLAGS',       size: '35px', sotrable: true, hidden: true},
            { field: 'DoneUID',     caption: 'DoneUID',     size: '35px', sotrable: true, hidden: true},
            { field: 'PreDoneUID',  caption: 'PreDoneUID',  size: '35px', sotrable: true, hidden: true},
            { field: 'Comment',     caption: 'Comment',     size: '35px', sotrable: true, hidden: true},
            { field: 'LastModTime', caption: 'LastModTime', size: '35px', sotrable: true, hidden: true},
            { field: 'LastModBy',   caption: 'LastModBy',   size: '35px', sotrable: true, hidden: true},
            { field: 'CreateTS',    caption: 'CreateTS',    size: '35px', sotrable: true, hidden: true},
            { field: 'CreateBy',    caption: 'CreateBy',    size: '35px', sotrable: true, hidden: true},
        ],
        onClick: function(event) {
            event.onComplete = function (event) {
                var r = w2ui.tldsDetailGrid.records[event.recid];
                console.log( 'detail clicked: v1/tasks/' + r.BID + '/'+ r.TDID);
                popupTaskDescForm(r.BID,r.TDID);
            };
        },
    });

    //------------------------------------------------------------------------
    //  taskDescForm
    //------------------------------------------------------------------------
    $().w2form({
        name: 'taskDescForm',
        style: 'border: 0px; background-color: transparent;',
        formURL: '/webclient/html/formtd.html',
        url: '/v1/td',
        fields: [
            { field: 'recid',          type: 'int',         required: false },
            { field: 'TDID',           type: 'int',         required: false },
            { field: 'BID',            type: 'int',         required: false },
            { field: 'TLID',           type: 'int',         required: false },
            { field: 'Name',           type: 'text',        required: true  },
            { field: 'Worker',         type: 'text',        required: false },
            { field: 'lstWorker',      type: 'list',        required: false, options: {items: app.workers}, },
            { field: 'EpochDue',       type: 'datetime',    required: false },
            { field: 'EpochPreDue',    type: 'datetime',    required: false },
            { field: 'Epoch',          type: 'datetime',    required: false },
            { field: 'ChkEpoch',       type: 'checkbox',    required: false },
            { field: 'ChkEpochDue',    type: 'checkbox',    required: false },
            { field: 'ChkEpochPreDue', type: 'checkbox',    required: false },
            { field: 'FLAGS',          type: 'text',        required: false },
            { field: 'DoneUID',        type: 'int',         required: false },
            { field: 'PreDoneUID',     type: 'int',         required: false },
            { field: 'Comment',        type: 'text',        required: false },
            { field: 'LastModTime',    type: 'date',        required: false },
            { field: 'LastModBy',      type: 'date',        required: false },
            { field: 'CreateTS',       type: 'date',        required: false },
            { field: 'CreateBy',       type: 'date',        required: false },
        ],
        actions: {
            save: function(target, data){
                // console.log("save task");
                var f = w2ui.taskDescForm;
                var r = f.record;
                var d = {cmd: "save", record: r};
                var dat=JSON.stringify(d);
                f.url = '/v1/td/' + r.BID + '/' + r.TDID;
                $.post(f.url,dat)
                .done(function(data) {
                    if (data.status === "error") {
                        f.error(w2utils.lang(data.message));
                        return;
                    }
                    //w2ui.tlsDetailGrid.url = '/v1/tl/'
                    w2ui.tlsDetailGrid.reload();
                    w2popup.close();
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
            };
        },
        onChange: function(event) {
            event.onComplete = function() {
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
                // getFormSubmitData(data.postData.record);
                var tl = {
                    cmd: "save",
                    record: w2ui.tldsInfoForm.record
                };
                var dat=JSON.stringify(tl);
                var url='/v1/tld/' + w2ui.tldsInfoForm.record.BID + '/' + w2ui.tldsInfoForm.record.TLDID;
                console.log( 'url = ' + url + '  dat = ' + dat);
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
                    w2ui.tldsInfoForm.error("Save Tasklist failed.");
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
        name: 'tldLayout',
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
};

window.finishTLDForm = function () {
    w2ui.tldLayout.content('top',   w2ui.tldsInfoForm);
    w2ui.tldLayout.content('main',  w2ui.tldsDetailGrid);
    w2ui.tldLayout.content('bottom',w2ui.tldsCloseForm);
};

//-----------------------------------------------------------------------------
// popupTaskDescForm - Bring up the task descriptor edit form
// 
// @params
//     bid = business id
//     tdid = task descriptor id
//  
// @returns
//  
//-----------------------------------------------------------------------------
window.popupTaskDescForm = function (bid,tdid) {
    w2ui.taskDescForm.url = '/v1/td/' + bid + '/' + tdid;
    w2ui.taskDescForm.request();
    $().w2popup('open', {
        title   : 'Task',
        body    : '<div id="form" style="width: 100%; height: 100%;"></div>',
        style   : 'padding: 15px 0px 0px 0px',
        width   : 600,
        height  : 450,
        showMax : true,
        onToggle: function (event) {
            $(w2ui.taskDescForm.box).hide();
            event.onComplete = function () {
                $(w2ui.taskDescForm.box).show();
                w2ui.taskDescForm.resize();
            };
        },
        onOpen: function (event) {
            event.onComplete = function () {
                $('#w2ui-popup #form').w2render('taskDescForm');
            };
        }
    });
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

        w2ui.toplayout.content('right', w2ui.tldLayout);
        w2ui.toplayout.show('right', true);
        w2ui.toplayout.sizeTo('right', 600);
        w2ui.toplayout.render();
        app.new_form_rec = false;  // mark as record exists
        app.form_is_dirty = false; // mark as no changes yet
    }
};

