/*global
    parseInt, w2ui, getDepMeth, getReservation, $, app, getBUDfromBID, getCurrentBusiness, console,
    saveReservationForm, switchToReservationss, finishReservationsForm, reservationsUpdateRTList,
    getReservationInitRecord, reservationSrch, daysBetweenDates, switchToBookRes,
    getBookResInitRecord, resSaveCB, setToForm, setResUpdateRecordForUI,
*/

"use strict";

// buildResUpdateElements creates the rid and reservation form to find
//------------------------------------------------------------------------------
window.buildResUpdateElements = function () {
    //------------------------------------------------------------------------
    //          resUpdateGrid
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'resUpdateGrid',
        url: '/v1/reservation',
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
            toolbarSearch   : true,
            toolbarInput    : true,
            searchAll       : true,
            toolbarReload   : false,
            toolbarColumns  : false,
        },
        columns: [
            {field: 'recid',            caption: 'recid',           size: '60px', hidden: true, sortable: true },
            {field: 'RLID',             caption: 'RLID',            size: '60px', hidden: true, sortable: true, style: 'text-align: right' },
            {field: 'RID',              caption: 'RID',             size: '60px', hidden: true, sortable: true, style: 'text-align: right' },
            {field: 'ConfirmationCode', caption: 'ConfirmationCode',size: '165px', hidden: false, sortable: true },
            {field: 'DtStart',          caption: 'DtStart',         size: '80px', hidden: false, sortable: true, style: 'text-align: center',
                render: function (record , index, col_index) {
                    var d = new Date(record.DtStart);
                    return dateFmtStr(d); // if non-recur assessment then do nothing
                },
            },
            {field: 'DtStop',           caption: 'DtStop',          size: '80px', hidden: false, sortable: true, style: 'text-align: center',
                render: function (record , index, col_index) {
                    var d = new Date(record.DtStop);
                    return dateFmtStr(d); // if non-recur assessment then do nothing
                },
            },
            {field: 'FirstName',        caption: 'FirstName',       size: '90px', hidden: false, sortable: true },
            {field: 'LastName',         caption: 'LastName',        size: '90px', hidden: false, sortable: true },
            {field: 'Email',            caption: 'Email',           size: '175px', hidden: false, sortable: true },
            {field: 'Phone',            caption: 'Phone',           size: '100px', hidden: false, sortable: true },
            {field: 'RentableName',     caption: 'RentableName',    size: '100px', hidden: false, sortable: true },
            {field: 'Name',             caption: 'Name',            size: '5%', hidden: false, sortable: true },
                ],
        onClick: function(event) {
            event.onComplete = function () {
                var rec = w2ui.resUpdateGrid.get(event.recid);
                console.log('book RLID = ' + rec.RLID);
                // switchToResUpdate(rec.RLID);
                var BID = getCurrentBID();
                var BUD = getBUDfromBID(BID);
                getRentableTypes(BUD, function() {
                    var url = '/v1/reservation/' + BID + '/' + rec.RLID;
                    setToForm('resUpdateForm',url,500,true);
                });
            };
        },

        onRequest: function(/*event*/) {
            w2ui.expenseGrid.postData = {searchDtStart: app.D1, searchDtStop: app.D2};
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
            };
        }
    });

    addDateNavToToolbar('resUpdate');

    //------------------------------------------------------------------------
    //          reservation Update Form
    //------------------------------------------------------------------------
    $().w2form({
        name: 'resUpdateForm',
        style: 'border: 0px; background-color: transparent;',
        header: 'Change A Reservation',
        url: '/v1/reservation/',
        formURL: '/webclient/html/formresup.html',
        fields: [
            { field: 'BUD',              type: 'list',  required: false, options: {items: app.businesses} },
            { field: 'RLID',             type: 'int',   required: false },
            { field: 'RTRID',            type: 'int',   required: false },
            { field: 'rdRTID',           type: 'list',   required: false },
            { field: 'RID',              type: 'int',   required: false },
            { field: 'rdBID',            type: 'int',   required: false },
            { field: 'ConfirmationCode', type: 'text',  required: false },
            { field: 'LeaseStatus',      type: 'int',   required: false },
            { field: 'DtStart',          type: 'date',  required: true },
            { field: 'DtStop',           type: 'date',  required: true },
            { field: 'Nights',           type: 'int',   required: false },
            { field: 'FirstName',        type: 'text',  required: true },
            { field: 'LastName',         type: 'text',  required: true },
            { field: 'Phone',            type: 'text',  required: false },
            { field: 'Email',            type: 'text',  required: false },
            { field: 'Street',           type: 'text',  required: false },
            { field: 'City',             type: 'text',  required: false },
            { field: 'State',            type: 'text',  required: false },
            { field: 'PostalCode',       type: 'text',  required: false },
            { field: 'CCName',           type: 'text',  required: false },
            { field: 'CCType',           type: 'text',  required: false },
            { field: 'CCNumber',         type: 'text',  required: false },
            { field: 'CCExpMonth',       type: 'text',  required: false },
            { field: 'CCExpYear',        type: 'text',  required: false },
            { field: 'Comment',          type: 'text',  required: false },
            { field: 'LastModTime',      type: 'time',  required: false },
            { field: 'LastModBy',        type: 'int',   required: false },
            { field: 'CreateTS',         type: 'time',  required: false },
            { field: 'CreateBy',         type: 'int',   required: false },
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
                            w2ui.reservationGrid.render();
                        };
                    form_dirty_alert(yes_callBack, no_callBack);
                    break;
                }
            },
        },
        onLoad: function(event) {
            // Need to add the BUD value...
            event.onComplete = function() {
                var f = this;
                var r = this.record;
                r.BUD = getBUDfromBID(r.rdBID);
                var d = new Date(r.DtStart);
                r.DtStart = dateFmtStr(d);
                d = new Date(r.DtStop);
                r.DtStop = dateFmtStr(d);
            };
        },
        onRender: function(event) {
            setResUpdateRecordForUI(this);
        },
        onRefresh: function(event) {
            setResUpdateRecordForUI(this);
        },
        onChange: function(event) {
            event.onComplete = function() {
                var x,y;
                var f = this;
                var draw=false;
                switch (event.target) {
                case "DtStart":
                    x = dateFromString(event.value_new);
                    y = dateFromString(f.record.DtStop);
                    x.setDate(x.getDate() + f.record.Nights);
                    f.record.DtStop = w2uiDateControlString(x);
                    draw = true;
                    break;
                case "DtStop":
                    x = dateFromString(event.value_new);
                    y = dateFromString(f.record.DtStart);
                    if (x <= y) {
                        x.setDate(x.getDate() - f.record.Nights);
                        f.record.DtStart = w2uiDateControlString(x);
                    }
                    x = dateFromString(f.record.DtStart);
                    y = dateFromString(f.record.DtStop);
                    f.record.Nights = daysBetweenDates(x,y);
                    draw = true;
                    break;
                case "Nights":
                    x = dateFromString(f.record.DtStart);
                    y = dateFromString(f.record.DtStop);
                    x.setDate(x.getDate() + event.value_new);
                    f.record.DtStop = w2uiDateControlString(x);
                    draw = true;
                    break;
                }
                if (draw) { f.refresh(); }
            };
        },
        onSubmit: function(target, data) {
            delete data.postData.record.LastModTime;
            delete data.postData.record.LastModBy;
            delete data.postData.record.CreateTS;
            delete data.postData.record.CreateBy;
            // modify form data for server request
            //getFormSubmitData(data.postData.record);
        },
    });
};

//---------------------------------------------------------------------------------
// setResUpdateRecordForUI - changes the main view of the program to the
//                        Reservations form
//
// @params
//      f = the form 
// @return
//
//---------------------------------------------------------------------------------
window.setResUpdateRecordForUI = function (f) {
    var r = f.record;
    var BID = getCurrentBID();
    var BUD = getBUDfromBID(BID);
    if (typeof f.get('rdRTID').options != "undefined") {
        f.get('rdRTID').options.items = app.rt_list[BUD];
    }
    if (typeof f.get('BUD').options != "undefined") {
        f.get('BUD').options.items = app.businesses;
    }
};
