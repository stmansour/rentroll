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
    updateAppBusinesses, setToBiz, rebuildBusinessSelect,BUDPickerRender,
    BUDPickerDropRender, BUDPickerCompare, updateBUDLink,setBUDLink,BUDHandler,
    setBUDSpinner,
*/
//-----------------------------------------------------------------------------
// updateBUDFormList updates the dropdown list of BUDs form interfaces. The
// contents may have been updated based on the user's previous actions...
//
// INPUTS
// f = name of form with a field that uses a dropdown list for BUD
//-----------------------------------------------------------------------------
window.updateBUDFormList = function(f) {
    var idx = f.get("BUD", true);
    f.fields[idx].options.items = app.businesses;
};

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

// // buildBizDropdownHTML creates the HTML for a dropdown menu of businesses
// //
// // INPUTS
// //     id = the id string used for the <span> containing the <select>
// //
// // RETURNS the html for a dropdown menu with all businesses.
// //----------------------------------------------------------------------------
// window.buildBizDropdownHTML = function(id) {
//     var html = '<span id="' +id + '"><select name="BusinessSelect" onchange="ChangeBusiness();">';
//     for (var i = 0; i < this.records.length; i++) {
//         var BUD = this.records[i].BUD;
//         var BID = this.records[i].BID;
//         html += '<option value="' + BID + '" name="' + BUD + '">' + BUD + '</option>';
//     }
//     html += '</select></span>';
//     return html;
// };

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
                var r = this.records;
                var BizMap = [];
                var BID = getCurrentBID();
                var BUD = getBUDfromBID(BID);
                for (var i = 0; i < this.records.length; i++) {
                    BizMap.push({BID: r[i].BID, BUD: r[i].BUD});
                }
                app.BizMap = BizMap;
                updateAppBusinesses();
                rebuildBusinessSelect(BUD);
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
            { field: 'BUD',                     type: 'enum',     required: false,
                options: {
                    url:            app.Directory + 'v1/butd/',
                    max:            1,
                    items:          [],
                    openOnFocus:    false,
                    maxDropWidth:   350,
                    maxDropHeight:  350,
                    renderItem:     BUDPickerRender,
                    renderDrop:     BUDPickerDropRender,
                    compare:        BUDPickerCompare,
                    onNew: function (event) {
                        $.extend(event.item, { Designation: event.item.text, Name: '', ClassCode: 0, CoCode: 0 } );
                    },
                    onRemove: function(event) {
                        // event.onComplete = function() {
                        //     var f = w2ui.bizDetailForm;
                        //     // reset BUD field related data when removed
                        //     f.record.ClassCode = 0;
                        //     f.record.CoCode = 0;
                        //     f.record.BUD = "";
                        //
                        //     // NOTE: have to trigger manually, b'coz we manually change the record,
                        //     // otherwise it triggers the change event but it won't get change (Object: {})
                        //     var event = f.trigger({ phase: 'before', target: f.name, type: 'change', event: event }); // event before
                        //     if (event.cancelled === true) return false;
                        //     f.trigger($.extend(event, { phase: 'after' })); // event after
                        // };
                    },
                },
            },
            { field: 'DefaultRentCycle',        type: 'list',     required: false, options: {items: app.w2ui.listItems.cycleFreq}, },
            { field: 'Name',                    type: 'text',     required: false },
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
            { field: 'ClassCode',               type: 'hidden',   required: false },
            { field: 'CoCode',                  type: 'hidden',   required: false },
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
        onLoad: function(event) {
            event.onComplete = function(event) {
                setBUDSpinner();
                setTimeout(BUDHandler, 500);
                // BUDHandler();
            };
        },
        // onRender: function(event) {
        //     if (this.record.BID > 0) {
        //         setBUDSpinner();
        //         setTimeout(BUDHandler, 750);
        //     }
        // },
        onRefresh: function(event) {
            var f = this;
            event.onComplete = function(event) {
                var f = w2ui.bizDetailForm;
                var r = f.record;
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
                }
                BUDHandler();
            };
        },
        onChange: function(event) {
            event.onComplete = function() {
                if (event.target == "BUD") {
                    BUDHandler();
                }
            };
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
                var newBID = w2ui.bizDetailForm.record.BID;
                var dat=JSON.stringify(business);
                var url='/v1/business/' + newBID;
                var newBUD = business.record.BUD;
                var BID = getCurrentBID();
                var BUD = getBUDfromBID(BID);


                $.post(url,dat)
                .done(function(data) {
                    if (data.status === "error") {
                        w2ui.bizDetailForm.error(w2utils.lang(data.message));
                        return;
                    }
                    setToBiz(BID,BUD,newBUD);
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
        var idx = w2ui.bizDetailForm.get("CPTLName",true);
        w2ui.bizDetailForm.fields[idx].options.url = '/v1/tltd/' + bid;
        var x = document.getElementById("CPTLName");
        if (x != null) {
            x.disabled = false;
        }
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
    var f = w2ui.bizDetailForm;
    f.record = getBusinessInitRecord();
    f.header = 'Business (new)';

    //------------------------------------------------------------------
    // for a new business, there will not yet be any Task Lists.  So we
    // need to make this control insensitive.
    //------------------------------------------------------------------
    // var idx = f.get("CPTLName",1);
    // f.fields[11].options.url = '/v1/tltd'; // no biz id
    var x = document.getElementById("CPTLName");
    if (x != null) {
        x.disabled = true;
    }
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
    // console.log('entered tlPickerCompare');
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
    // console.log('entered tlPickerDropRender');
    return getTLName(item);
};

//-----------------------------------------------------------------------------
// tlPickerRender - renders a name during typedown.
// @params
//   item = an object assumed to have a FirstName and LastName
// @return - true if the names match, false otherwise
//-----------------------------------------------------------------------------
window.tlPickerRender = function (item) {
    // console.log('entered tlPickerRender.  item.Name = ' + item.Name + '  item.TLID = ' + item.TLID);
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


//-----------------------------------------------------------------------------
// setToBiz - sets internal BizMap data to the latest business list from
// the server. It udates app.BizMap and app.businesses first.
//
// @params
//  newBUD - new business to switch to if we're not already on it...
// @return
//-----------------------------------------------------------------------------
window.setToBiz = function(BID,BUD,newBUD) {
    var url = "/v1/uival/" + BID + "/app.BizMap";
    $.get(url, null, null, "json")
    .done(function(data) {
        if (data == null) {return;}
        if (typeof data == "object" && typeof data.status == "string" && data.status == "error") {
            w2popup.open({
                title   : 'PBUDoblem Loading BizMap',
                buttons : '<BUDtton class="w2ui-btn" onclick="w2popup.close();">Close</button> ',
                width   : 500,
                height  : 300,
                body    : '<div class="w2ui-centered"><div style="padding: 10px;">Error updating BizMap. data = ' + data.message + '</div></div>'
            });
            return;
        }
        app.BizMap = data;
        updateAppBusinesses();
        rebuildBusinessSelect(newBUD);
    })
    .fail(function(data) {
        w2popup.open({
            title   : 'Problem Loading BizMap',
            buttons : '<button class="w2ui-btn" onclick="w2popup.close();">Close</button> ',
            width   : 500,
            height  : 300,
            body    : '<div class="w2ui-centered"><div style="padding: 10px;">Error updating BizMap. data = ' + data + '</div></div>'
        });
    });
};

window.updateAppBusinesses = function() {
    var b = [];
    for (var i = 0; i < app.BizMap.length; i++) {
        b.push(app.BizMap[i].BUD);
    }
    app.businesses = b;
};

// rebuildBusinessSelect rebuilds the business select dropdown menu on the top
// toolbar.  It builds it based on the contents of app.BizMap, so be sure to
// update app.BizMap prior to calling this routine.
//
// INPUTS:
//   selBUD = bud to be selected on newly rebuilt dropdown menu
//-----------------------------------------------------------------------------
window.rebuildBusinessSelect = function(selBUD) {
    var html = '<select name="BusinessSelect" onchange="ChangeBusiness();">';
    for (var i = 0; i < app.BizMap.length; i++) {
        var BUD = app.BizMap[i].BUD;
        var selected = (BUD == selBUD) ? " selected " : "";
        html += '<option' + selected + ' value="' + app.BizMap[i].BID + '" name="' + BUD + '">' + BUD + '</option>';
    }
    html += '</select>';
    document.getElementById("bizdropdown").innerHTML = html;
};

//-----------------------------------------------------------------------------
// BUDPickerCompare - Compare item to the search string. Verify that the
//          supplied search string can be found in item.  If it's already
//          listed we don't want to list it again.
// @params
//   item = an object assumed to have a Name and TLID field
// @return - true if the search string is found, false otherwise
//-----------------------------------------------------------------------------
window.BUDPickerCompare = function (item, search) {
    var s = item.Designation;
    s = s.toLowerCase();
    var srch = search.toLowerCase();
    var match = (s.indexOf(srch) >= 0);
    return match;
};

//-----------------------------------------------------------------------------
// BUDPickerDropRender - renders a name during typedown.
// @params
//   item = an object assumed to have a FirstName and LastName
// @return - the name to render
//-----------------------------------------------------------------------------
window.BUDPickerDropRender = function (item) {
    return item.Designation;
};

//-----------------------------------------------------------------------------
// BUDPickerRender - renders a name during typedown.
// @params
//   item = an object assumed to have a FirstName and LastName
// @return - true if the names match, false otherwise
//-----------------------------------------------------------------------------
window.BUDPickerRender = function (item) {
    var r = w2ui.bizDetailForm.record;
    if (item.text != undefined && item.text != null) {
        r.BUD = item.text; // doesn't exist in Directory. It's OK, just exit now
        return item.text;
    }
    r.BUD = item.Designation;
    if (r.BUD == undefined) {return 'rBUD.undefined';}
    if (r.BUD == null ) { return 'rBUD.null'; }
    if (item.Name.length > 0 ) {
        r.Name = item.Name;
        var x = document.getElementById("Name");
        if (x == undefined || x == null) { return r.BUD; }
        x.value = item.Name;
    }
    return r.BUD;
};

window.setBUDLink = function(x,s) {
    if (x != null) {
        x.innerHTML = s;
    }
};

window.setBUDSpinner = function() {
    var s = '<i class="fas fa-spinner fa-pulse fa-lg fa-fw" style="color:#0000CC;"></i>';
    var x = document.getElementById("BUDlink");
    setBUDLink(x,s);
};

window.updateBUDLink = function() {
    var f = w2ui.bizDetailForm;
    var r = f.record;
    var x = document.getElementById("BUDlink");
    var u = '<i class="far fa-circle fa-lg" style="color:#FF8C00;"></i>';

    if (r.BUD == undefined || r.BUD.length == 0) {
        setBUDLink(x,u);
        return;
    }

    var l = '<i class="fas fa-link fa-lg" style="color:#00CC00;"></i>';
    setBUDSpinner();

    //---------------------------------------------------------
    // filter r.BUD object down to bud string if necessary...
    //---------------------------------------------------------
    if (typeof r.BUD == "object") {
        var BUD = r.BUD[0].Designation;
        r.BUD = BUD;
    }

    var url = app.Directory + 'v1/bud?request=' + encodeURIComponent(JSON.stringify({search: r.BUD}));
    $.get(url, null, null, "json")
    .done(function(data) {
        var x = document.getElementById("BUDlink");
        if (data == null) {return;}
        if (typeof data == "object" && typeof data.status == "string" && data.status == "error") {
            setBUDLink(x,u);
            if ( ! data.message.includes("Not Found:")) {
                f.error('Error retrieving BUD info: ' + data.message);
            }
            return;
        }
        setBUDLink(x,l);
        w2ui.bizDetailForm.record.Name = data.record.Name;
        var y = document.getElementById("Name");
        if (y != undefined && y != null) {
            y.value = data.record.Name;
        }
    })
    .fail(function(data) {
        f.error('Error retrieving BUD info ' + data);
    });

};

window.BUDHandler = function() {
    var f = w2ui.bizDetailForm;
    var r = f.record;
    updateBUDLink();
    if (r.BUD != undefined && r.BUD != null) {
        var item = { Designation: r.BUD, Name: '', ClassCode: 0, CoCode: 0 };
        if ($(f.box).find("input[name=BUD]").length > 0) {
            $(f.box).find("input[name=BUD]").data('selected', [item]).data('w2field').refresh();
        }
    }
};
