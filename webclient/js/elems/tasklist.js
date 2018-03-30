"use strict";
/*global
    w2ui, $, app, console, w2utils,
    form_dirty_alert, addDateNavToToolbar, getFormSubmitData,
    dtTextRender, dateFromString, 
*/

function buildTaskListElements() {
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
            {field: 'DtDue',     hidden: false, caption: 'Due',               size: '150px', sortable: true,
                render: function (record, index, col_index) {
                    if (typeof record === "undefined") {
                        return '';
                    }
                    return taskDateRender(record.DtDue);
                }
            },
            {field: 'DtDone',    hidden: false, caption: 'Due completed',     size: '150px', sortable: true,
                render: function (record, index, col_index) { if (typeof record === "undefined") {return '';} return taskDateRender(record.DtDone); }
            },
            {field: 'DtPreDue',  hidden: false, caption: 'Pre Due',           size: '150px', sortable: true,
                render: function (record, index, col_index) { if (typeof record === "undefined") {return '';} return taskDateRender(record.DtPreDue); }
            },
            {field: 'DtPreDone', hidden: false, caption: 'Pre Due completed', size: '150px', sortable: true,
                render: function (record, index, col_index) { if (typeof record === "undefined") {return '';} return taskDateRender(record.DtPreDone); }
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
            { field: 'recid',        caption: 'recid',       type: 'int',       required: false },
            { field: 'TLID',         caption: 'TLID',        type: 'int',       required: false },
            { field: 'BID',          caption: 'BID',         type: 'int',       required: false },
            { field: 'BUD',          caption: 'BUD',         type: 'list',      required: true, options: {items: app.businesses} },
            { field: 'Name',         caption: 'Name',        type: 'text',      required: true },
            { field: 'Cycle',        caption: 'Cycle',       type: 'list',      required: true, options: {items: app.w2ui.listItems.cycleFreq}, },
            { field: 'DtDue',        caption: 'DtDue',       type: 'date',      required: false },
            { field: 'DtPreDue',     caption: 'DtPreDue',    type: 'date',      required: false },
            { field: 'DtDone',       caption: 'DtDone',      type: 'date',      required: false },
            { field: 'DtPreDone',    caption: 'DtPreDone',   type: 'date',      required: false },
            { field: 'FLAGS',        caption: 'FLAGS',       type: 'int',       required: false },
            { field: 'DoneUID',      caption: 'DoneUID',     type: 'int',       required: false },
            { field: 'PreDoneUID',   caption: 'PreDoneUID',  type: 'int',       required: false },
            { field: 'Comment',      caption: 'Comment',     type: 'text',      required: false },
            { field: 'CreateTS',     caption: 'CreateTS',    type: 'date',      required: false },
            { field: 'CreateBy',     caption: 'CreateBy',    type: 'int',       required: false },
            { field: 'LastModTime',  caption: 'LastModTime', type: 'date',      required: false },
            { field: 'LastModBy',    caption: 'LastModBy',   type: 'int',       required: false },
            { field: 'ChkDtDue',     caption: 'ChkDtDue',    type: 'checkbox',  required: false },
            { field: 'ChkDtDone',    caption: 'ChkDtDone',   type: 'checkbox',  required: false },
            { field: 'ChkDtPreDue',  caption: 'ChkDtPreDue', type: 'checkbox',  required: false },
            { field: 'ChkDtPreDone', caption: 'ChkDtPreDone',type: 'checkbox',  required: false },
        ],
        onRefresh: function(event) {
            // var f = this;
            event.onComplete = function(event) {
                var r = w2ui.tlsInfoForm.record;
                var y;
                var s = '';
                var dt;
                var now = new Date();
                if (typeof r.DtPreDue === "undefined") {
                    return;
                }
                if (r.DtPreDue !== null && r.DtPreDue.length > 0) {
                    y = dateFromString(r.DtPreDue);
                    r.ChkDtPreDue = y.getFullYear() >= 2000;
                    if (r.ChkDtPreDue) {
                        s = taskDateRender(r.DtPreDue); 
                    } else {
                        s = 'No pre-due date';
                    }
                    document.getElementById("sDtPreDue").innerHTML = s;
                }
                if (r.DtPreDone !== null && r.DtPreDone.length > 0) {
                    y = dateFromString(r.DtPreDone);
                    r.ChkDtPreDone = y.getFullYear() >= 2000;
                    if (r.ChkDtPreDone) {
                        s = taskDateRender(r.DtPreDone); 
                    } else {
                        dt = dateFromString(r.DtPreDone);
                        if (now > dt) {
                            s = '<span style="color:#FC0D1B;">overdue</span>';
                        }
                    }
                    document.getElementById("sDtPreDone").innerHTML = s;
                }
                if (r.DtDue !== null && r.DtDue.length > 0) {
                    y = dateFromString(r.DtDue);
                    r.ChkDtDue = y.getFullYear() >= 2000;
                    if (r.ChkDtDue) {
                        s = taskDateRender(r.DtDue); 
                    } else {
                        s = 'No due date';
                    }
                    document.getElementById("sDtDue").innerHTML = s;
                }
                if (r.DtDone !== null && r.DtDone.length > 0) {
                    y = dateFromString(r.DtDone);
                    r.ChkDtDone = y.getFullYear() >= 2000;
                    if (r.ChkDtDone) {
                        s = taskDateRender(r.DtDone); 
                    } else {
                        dt = dateFromString(r.DtDone);
                        if (now > dt) {
                            s = '<span style="color:#FC0D1B;">overdue</span>';
                        }
                    }
                    document.getElementById("sDtDone").innerHTML = s;
                }
            };
        },
        onChange: function(event) {
            event.onComplete = function() {
                var s = '';
                if (event.target === "ChkDtPreDone") {
                    if (event.value_new) { // marked as complete?
                        s = '<span style="color:blue;">will mark as completed when Save is clicked</span>';
                    } else {
                        s = '<span style="color:blue;">will mark as not completed when Save is clicked</span>';
                    }
                    document.getElementById("sDtPreDone").innerHTML = s;
                } else if (event.target === "ChkDtDone") {
                    if (event.value_new) { // marked as complete?
                        s = '<span style="color:blue;">will mark as completed when Save is clicked</span>';
                    } else {
                        s = '<span style="color:blue;">will mark as not completed when Save is clicked</span>';
                    }
                    document.getElementById("sDtDone").innerHTML = s;
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
                var r = w2ui.tlsDetailGrid.records[event.recid];
                console.log( 'detail clicked: v1/tasks/' + r.BID + '/'+ r.TID);
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
                var tl = {
                    cmd: "save",
                    record: w2ui.tlsInfoForm.record
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
}

function finishTaskListForm() {
    w2ui.tlLayout.content('top',   w2ui.tlsInfoForm);
    w2ui.tlLayout.content('main',  w2ui.tlsDetailGrid);
    w2ui.tlLayout.content('bottom',w2ui.tlsCloseForm);
}

//-----------------------------------------------------------------------------
// setToTLForm -  enable the Statement form in toplayout.  Also, set
//                the forms url and request data from the server
// @params
//   bid = business id (or the BUD)
//    id = Task List TLID
// d1,d2 = date range to use
//-----------------------------------------------------------------------------
function setToTLForm(bid, id, d1,d2) {
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
}

//-----------------------------------------------------------------------------
// taskDateRender - If the date is less than year 2000 then return a blank
//                otherwise return the date as a string.
// @params
//   y - the date to be printed
// @returns
//   the string to print
//-----------------------------------------------------------------------------
function taskDateRender(x) {
    if (x === null) {
        return '';
    }
    var y = dateFromString(x);
    var yr = y.getFullYear();
    if ( yr < 2000) {
        return '';
    }
    return dtTextRender(x,0,0);
}