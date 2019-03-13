"use strict";
/*global
    w2ui, $, app, console, w2utils,
    form_dirty_alert, addDateNavToToolbar, w2uiDateControlString,
    dateFromString, taskDateRender, setToTLForm,
    taskFormDueDate,taskCompletionChange,taskFormDoneDate,
    openTaskForm,setInnerHTML,w2popup,ensureSession,dtFormatISOToW2ui,
    createNewBusiness, getBUDfromBID, exportItemReportPDF, exportItemReportCSV,
    popupNewBusinessForm, getTLDs, getCurrentBID, getNewBusinessRecord,
    closeTaskForm, setTaskButtonsState, renderTaskGridDate, localtimeToUTC, TLD,
    taskFormDueDate1, finishBizForm, setToBizForm, renderReversalIcon,
    getGridReversalSymbolHTML, openBizForm, createNewBusiness, getBusinessInitRecord,
    tlPickerRender, tlPickerDropRender, tlPickerCompare, getTLName,
*/

//-----------------------------------------------------------------------------
// getBusinessInitRecord - the default new record values
//
// @params
//   bid = business id (or the BUD)
// d1,d2 = date range to use
//-----------------------------------------------------------------------------
window.getBusinessInitRecord = function() {
    var y = {
        Name: "",
        BUD: "",
        BID: 0,
        DefaultRentCycle: 6,
        DefaultProrationCycle: 4,
        DefaultGSRPC: 4,
        FLAGS: 1,
        EDIenabled: true,
        AllowBackdatedRA: false,
        Disabled: false,
    };
    return y;
};

//-----------------------------------------------------------------------------
// renderReversalIcon - Show the reversal icon if FLAGS bit 2 is set
//
// @params
//   bid = business id (or the BUD)
// d1,d2 = date range to use
//-----------------------------------------------------------------------------
window.renderReversalIcon = function (record /*, index, col_index*/) {
    if (typeof record === "undefined") {
        return;
    }
    if ( (record.FLAGS & (1<<2)) !== 0 ) { // if reversed then
        return getGridReversalSymbolHTML();
    }
    return '';
};

window.buildBusinessElements = function () {
    //------------------------------------------------------------------------
    //          tlsGrid  -  TASK LISTS in the date range
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'businessGrid',
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
            {field: 'recid',                    hidden: true,  size: '40px', sortable: false, caption: 'recid' },
            {field: 'reversed',                 hidden: false, size: '10px', style: 'text-align: center', sortable: true,
                    render: renderReversalIcon,
            },
            {field: 'BID',                      hidden: false, size: '40px', sortable: false, caption: 'BID' },
            {field: 'BUD',                      hidden: false, size: '40px', sortable: false, caption: 'BUD' },
            {field: 'Name',                     hidden: false, size: '90%', sortable: false, caption: 'Name' },
            {field: 'DefaultRentCycle',         hidden: true,  size: '40px', sortable: false, caption: 'DefaultRentCycle' },
            {field: 'DefaultProrationCycle',    hidden: true,  size: '40px', sortable: false, caption: 'DefaultProrationCycle' },
            {field: 'DefaultGSRPC',             hidden: true,  size: '40px', sortable: false, caption: 'DefaultGSRPC' },
            {field: 'ClosePeriodTLID',          hidden: true,  size: '40px', sortable: false, caption: 'ClosePeriodTLID' },
            {field: 'FLAGS',                    hidden: true,  size: '40px', sortable: false, caption: 'FLAGS' },
            {field: 'EDIenabled',               hidden: true,  size: '40px', sortable: false, caption: 'EDIenabled' },
            {field: 'AllowBackdatedRA',         hidden: true,  size: '40px', sortable: false, caption: 'AllowBackdatedRA' },
            {field: 'Disabled',                 hidden: true,  size: '40px', sortable: false, caption: 'Disabled' },
            {field: 'LastModTime',              hidden: true,  size: '40px', sortable: false, caption: 'LastModTime' },
            {field: 'LastModBy',                hidden: true,  size: '40px', sortable: false, caption: 'LastModBy' },
            {field: 'CreateTS',                 hidden: true,  size: '40px', sortable: false, caption: 'CreateTS' },
            {field: 'CreateBy',                 hidden: true,  size: '40px', sortable: false, caption: 'CreateBy' },
        ],
        onLoad: function(event) {
            event.onComplete = function(event) {
                //------------------------------------------------
                // rebuild app bizmap and biz dropdown menu...
                //------------------------------------------------
                var html = '<select name="BusinessSelect" onchange="ChangeBusiness();">';
                var BizMap = [];
                for (var i = 0; i < this.records.length; i++) {
                    var BUD = this.records[i].BUD;
                    var BID = this.records[i].BID;
                    html += '<option value="' + BID + '" name="' + BUD + '">' + BUD + '</option>';
                    BizMap.push({BID: BID, BUD: BUD});
                }
                html += '</select>';
                document.getElementById("bizdropdown").innerHTML = html;
                app.BizMap = BizMap;
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
                        setToBizForm(rec.BID, app.D1, app.D2);
                    };
                form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
            };
        },
        onAdd: function(event) {
            event.onComplete = function () {
                var bid = getCurrentBID();
                createNewBusiness(bid);
            };
        },
    });

    addDateNavToToolbar('business'); // "Grid" is appended to the supplied string

    //------------------------------------------------------------------------
    //  bizLayout - The layout to contain the tabbed form and buttons
    //               main - bizDetailForm
    //               bottom - bizCloseForm
    //------------------------------------------------------------------------
    $().w2layout({
        name: 'bizLayout',
        padding: 0,
        panels: [
            { type: 'left',    size: 0,     hidden: true },
            { type: 'top',     size: 0,     hidden: true, content: 'top',  resizable: true, style: app.pstyle },
            { type: 'main',    size: '60%', hidden: false, content: 'main', resizable: true, style: app.pstyle },
            { type: 'preview', size: 0,     hidden: true,  content: 'PREVIEW'  },
            { type: 'bottom',  size: 50,    hidden: false, content: 'bottom', resizable: false, style: app.pstyle },
            { type: 'right',   size: 0,     hidden: true }
        ]
    });

    //------------------------------------------------------------------------
    //  bizDetailForm
    //------------------------------------------------------------------------
    $().w2form({
        name: 'bizDetailForm',
        style: 'border: 0px; background-color: transparent;',
        formURL: '/webclient/html/formbizdetail.html',
        url: '/v1/business',
        toolbar: {
            items: [
                { id: 'btnNotes', type: 'button', icon: 'far fa-sticky-note' },
                { id: 'bt3', type: 'spacer' },
                { id: 'btnClose', type: 'button', icon: 'fas fa-times' },
            ],
            onClick: function (event) {
                if (event.target == 'btnClose') {
                    var no_callBack = function() { return false; },
                        yes_callBack = function() {
                            w2ui.toplayout.hide('right',true);
                            w2ui.businessGrid.render();
                        };
                    form_dirty_alert(yes_callBack, no_callBack);
                }
                if (event.target == 'btnNotes') {
                    notesPopUp();
                }
            },
        },
        fields: [
            { field: 'BID',                     type: 'int',      required: false },
            { field: 'BUD',                     type: 'text',     required: false },
            { field: 'Name',                    type: 'text',     required: false },
            { field: 'DefaultRentCycle',        type: 'list',     required: false, options: {items: app.w2ui.listItems.cycleFreq}, },
            { field: 'DefaultProrationCycle',   type: 'list',     required: false, options: {items: app.w2ui.listItems.cycleFreq}, },
            { field: 'DefaultGSRPC',            type: 'list',     required: false, options: {items: app.w2ui.listItems.cycleFreq}, },
            { field: 'FLAGS',                   type: 'hidden',   required: false },
            { field: 'EDIenabled',              type: 'checkbox', required: false },
            { field: 'AllowBackdatedRA',        type: 'checkbox', required: false },
            { field: 'Disabled',                type: 'checkbox', required: false },
            { field: 'CPTLID',                  type: 'int',      required: false },
            { field: 'CPTLName',                type: 'enum',     required: false,
                options: {
                    url:            '/v1/tltd/',
                    max:            1,
                    items:          [],
                    openOnFocus:    true,
                    maxDropWidth:   350,
                    maxDropHeight:  350,
                    renderItem:     tlPickerRender,
                    renderDrop:     tlPickerDropRender,
                    compare:        tlPickerCompare,
                    onNew: function (event) {
                        console.log('++ New Item: Do not forget to submit it to the server too', event);
                        //$.extend(event.item, { FirstName: '', LastName : event.item.text });
                    }
                },
            },

            { field: 'LastModTime',             type: 'hidden',   required: false },
            { field: 'LastModBy',               type: 'hidden',   required: false },
            { field: 'CreateTS',                type: 'hidden',   required: false },
            { field: 'CreateBy',                type: 'hidden',   required: false },
        ],
        actions: {
            // save: function(target, data){
            //     var f = w2ui.bizDetailForm;
            //     var r = f.record;
            //     f.url = '/v1/business/' + r.BID ;
            //     var s = r.Name.text;
            //     r.TLDID = r.Name.id;
            //     r.Name = s;
            //     r.Pivot = localtimeToUTC(r.Pivot);
            //     r.Timezone = app.timezone;
            //     var params = {cmd: 'save', formname: f.name, record: r };
            //
            //     var dat = JSON.stringify(params);
            //     var BID = r.BID;
            //
            //     // submit request
            //     $.post(f.url, dat, null, "json")
            //     .done(function(data) {
            //         if (data.status != "success") {
            //             return;
            //         }
            //         w2ui.tlsGrid.reload();
            //         var tlid = data.recid;
            //         setToTLForm(BID, tlid, app.D1, app.D2);
            //         w2popup.close();
            //     })
            //     .fail(function(/*data*/){
            //         console.log("Payor Fund Allocation failed.");
            //     });
            //
            // },
        },
        // onLoad: function(event) {
        //     // event.onComplete = function(event) {
        //     // };
        // },
        onRefresh: function(event) {
            var f = this;
            event.onComplete = function(event) {
                var f = w2ui.bizDetailForm;
                var r = f.record;
                // if (typeof r.CPTLName === "undefined") {return;}
                if (r.ClosePeriodTLID === 0 || typeof r.CPTLName === "undefined") {
                    return;
                }
                var cp = {
                    recid: r.ClosePeriodTLID,
                    TLID: r.ClosePeriodTLID,
                    Name: r.CPTLName,
                };
                if ($(f.box).find("input[name=CPTLName]").length > 0) {
                    console.log('initialized CPTLName to '+cp.Name+ ' ' + cp.TLID);
                    $(f.box).find("input[name=CPTLName]").data('selected', [cp]).data('w2field').refresh();
                    //f.refresh();
                }
            };
        },
        onChange: function(event) {
            // event.onComplete = function() {
            // };
        },
    });

    //------------------------------------------------------------------------
    //  bizCloseForm
    //------------------------------------------------------------------------
    $().w2form({
        name: 'bizCloseForm',
        style: 'border: 0px; background-color: transparent;',
        formURL: '/webclient/html/formbizclose.html',
        url: '',
        fields: [],
        actions: {
            save: function(target, data){
                var r = w2ui.bizDetailForm.record;
                var business = {
                    cmd: "save",
                    record: r,
                };

                //---------------------------------------------
                // Change a few objects to integer values...
                //---------------------------------------------
                var x = r.DefaultRentCycle.id;
                business.record.DefaultRentCycle = x;
                x = r.DefaultProrationCycle.id;
                business.record.DefaultProrationCycle = x;
                x = r.DefaultGSRPC.id;
                business.record.DefaultGSRPC = x;

                var dat=JSON.stringify(business);
                var url='/v1/business/' + w2ui.bizDetailForm.record.BID;
                $.post(url,dat)
                .done(function(data) {
                    if (data.status === "error") {
                        w2ui.bizDetailForm.error(w2utils.lang(data.message));
                        return;
                    }
                    w2ui.toplayout.hide('right',true);
                    w2ui.businessGrid.render();
                })
                .fail(function(/*data*/){
                    w2ui.bizDetailForm.error("Save Business failed.");
                    return;
                });
            },

            delete: function(target,data) {
                // var business = {
                //     cmd: "delete",
                // };
                // var dat=JSON.stringify(business);
                // var url='/v1/business/' + w2ui.bizDetailForm.record.BID + '/' + w2ui.bizDetailForm.record.TLID;
                // $.post(url,dat)
                // .done(function(data) {
                //     if (data.status === "error") {
                //         w2ui.bizDetailForm.error(w2utils.lang(data.message));
                //         return;
                //     }
                //     w2ui.toplayout.hide('right',true);
                //     w2ui.tlsGrid.render();
                // })
                // .fail(function(/*data*/){
                //     w2ui.bizDetailForm.error("Save Business failed.");
                //     return;
                // });
            },
        },
    });
};

//-----------------------------------------------------------------------------
// openBizForm - make the form visible and render it
//
// @params
//-----------------------------------------------------------------------------
window.openBizForm = function() {
    w2ui.toplayout.content('right', w2ui.bizLayout);
    w2ui.toplayout.show('right', true);
    w2ui.toplayout.sizeTo('right', 500);
    w2ui.toplayout.render();
    app.new_form_rec = false;  // mark as record exists
    app.form_is_dirty = false; // mark as no changes yet
};

//-----------------------------------------------------------------------------
// setToBizForm - edit the supplied business.
//
// @params
//   bid = business id (or the BUD)
// d1,d2 = date range to use
//-----------------------------------------------------------------------------
window.setToBizForm = function (bid, d1,d2) {
    if (bid > 0) {
        w2ui.bizDetailForm.url = '/v1/business/' + bid; // the tasks associated with the selected tasklist
        w2ui.bizDetailForm.postData = {
            searchDtStart: d1,
            searchDtStop: d2,
        };
        w2ui.bizDetailForm.header = 'Business ' + bid;
        w2ui.bizDetailForm.request();
        w2ui.bizDetailForm.fields[11].options.url = '/v1/tltd/' + bid;
        openBizForm();
    }
};

//-----------------------------------------------------------------------------
// createNewBusiness - Initialize the form's record to a new business and
//      open the form.
//
// @params
//-----------------------------------------------------------------------------
window.createNewBusiness = function() {
    w2ui.bizDetailForm.record = getBusinessInitRecord();
    w2ui.bizDetailForm.header = 'Business (new)';
    w2ui.bizDetailForm.fields[11].options.url = '/v1/tltd'; // no biz id
    openBizForm();
};

//-----------------------------------------------------------------------------
// finishBizForm - Add graphical elements to the layout
//
// @params
//   bid = business id (or the BUD)
// d1,d2 = date range to use
//-----------------------------------------------------------------------------
window.finishBizForm = function (bid, d1,d2) {
    w2ui.bizLayout.content('main',  w2ui.bizDetailForm);
    w2ui.bizLayout.content('bottom',w2ui.bizCloseForm);
};

//-----------------------------------------------------------------------------
// tlPickerCompare - Compare item to the search string. Verify that the
//          supplied search string can be found in item.  If it's already
//          listed we don't want to list it again.
// @params
//   item = an object assumed to have a Name and TLID field
// @return - true if the search string is found, false otherwise
//-----------------------------------------------------------------------------
window.tlPickerCompare = function (item, search) {
    console.log('entered tlPickerCompare');
    var s = getTLName(item);
    s = s.toLowerCase();
    var srch = search.toLowerCase();
    var match = (s.indexOf(srch) >= 0);
    return match;
};

//-----------------------------------------------------------------------------
// tlPickerDropRender - renders a name during typedown.
// @params
//   item = an object assumed to have a FirstName and LastName
// @return - the name to render
//-----------------------------------------------------------------------------
window.tlPickerDropRender = function (item) {
    console.log('entered tlPickerDropRender');
    return getTLName(item);
};

//-----------------------------------------------------------------------------
// tlPickerRender - renders a name during typedown.
// @params
//   item = an object assumed to have a FirstName and LastName
// @return - true if the names match, false otherwise
//-----------------------------------------------------------------------------
window.tlPickerRender = function (item) {
    console.log('entered tlPickerRender.  item.Name = ' + item.Name + '  item.TLID = ' + item.TLID);
    w2ui.bizDetailForm.record.CPTLName = item.Name;
    w2ui.bizDetailForm.record.CPTLID = item.TLID;
    return getTLName(item);
};

//-----------------------------------------------------------------------------
// getTLName - returns an appropriate name for the supplied item.
//
// @params
//   item = an object assumed to have a Name and TLID
// @return - the name to render
//-----------------------------------------------------------------------------
window.getTLName = function (item) {
    var s = item.Name;
    if (item.TLID > 0) {
        s += ' ('+ item.TLID +')';
    }
    return s;
};
