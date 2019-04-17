/*global
    parseInt, w2ui, getDepMeth, getReservation, $, app, getBUDfromBID, getCurrentBusiness, console,
    saveReservationForm, switchToReservationss, finishReservationsForm, reservationsUpdateRTList,
    getReservationInitRecord, reservationSrch, daysBetweenDates, switchToBookRes,
    getBookResInitRecord, resSaveCB,
*/

"use strict";

window.getReservationInitRecord = function (BID, BUD, previousFormRecord){
    var y = new Date();
    var y1 = new Date(y.getFullYear(),y.getMonth(),y.getDate()+1,0,0);
    var defaultFormData = {
        recid:          0,
        BID:            BID,
        BUD:            BUD,
        RTID:           0,
        Nights:         1,
        DtStart:        dateControlString(y),
        DtStop:         dateControlString(y1),
        PromoCode:      '',
        LastModTime:    y.toISOString(),
        LastModBy:      0,
        CreateTS:       y.toISOString(),
        CreateBy:       0,
    };

    if (typeof app.rt_list[BUD] != "undefined") {
        defaultFormData.RTID = app.rt_list[BUD][0].id;
    }

    // if it called after 'save and add another' action there previous form
    // record is passed as Object else it is null
    // if ( previousFormRecord ) {
    //     defaultFormData = setDefaultFormFieldAsPreviousRecord(
    //         [ 'FLAGS', 'Amount', 'ClearedAmount', 'LastModTime', 'Comment'], // Fields to Reset
    //         defaultFormData,
    //         previousFormRecord
    //     );
    // }

    return defaultFormData;
};

// Record used by bookResForm
window.getBookResInitRecord = function (BID, BUD, previousFormRecord){
    var f = w2ui.reservationForm;
    var r = f.record;

    var y = new Date(r.DtStart);
    var y1 = new Date(r.DtStop);

    var defaultFormData = {
        RLID:           0,
        rdBID:          BID,
        rdBUD:          BUD,
        rdRTID:         r.RTID,
        RID:            r.RID,
        LeaseStatus:    0,
        Nights:         r.Nights,
        DtStart:        dateControlString(y),
        DtStop:         dateControlString(y1),
        RentableName:   '',
        FirstName:      '',
        LastName:       '',
        Email:          '',
        Phone:          '',
        Street:         '',
        City:           '',
        Country:        '',
        State:          '',
        PostalCode:     '',
        CCName:         '',
        CCType:         '',
        CCNumber:       '',
        CCExpMonth:     '',
        CCExpYear:      '',
        Comment:        '',
    };


    // if it called after 'save and add another' action there previous form
    // record is passed as Object else it is null
    if ( previousFormRecord ) {
        defaultFormData = setDefaultFormFieldAsPreviousRecord(
            [ 'FLAGS', 'Amount', 'ClearedAmount', 'LastModTime', 'Comment'], // Fields to Reset
            defaultFormData,
            previousFormRecord
        );
    }

    return defaultFormData;
};

// buildReservationsElements creates the reservation form
//------------------------------------------------------------------------------
window.buildReservationsElements = function () {
    //------------------------------------------------------------------------
    //          reservation Form
    //------------------------------------------------------------------------
    $().w2form({
        name: 'reservationForm',
        style: 'border: 0px; background-color: transparent;',
        header: 'Make A Reservation',
        url: '/v1/available',
        formURL: '/webclient/html/formreservation.html',

        fields: [
            { field: 'recid',         type: 'int',   required: false, html: { page: 0, column: 0 } },
            { field: 'BID',           type: 'int',   required: false, html: { page: 0, column: 0 } },
            { field: 'BUD',           type: 'list',  required: true, options: {items: app.businesses}, html: { page: 0, column: 0 } },
            { field: 'Nights',        type: 'int',   required: false, html: { page: 0, column: 0 } },
            { field: 'RTID',          type: 'list',  required: true },
            { field: 'DtStart',       type: 'date',  required: true },
            { field: 'DtStop',        type: 'date',  required: true },
            { field: 'PromoCode',     type: 'text',  required: true },
            { field: 'LastModTime',   type: 'time',  required: false, html: { page: 0, column: 0 } },
            { field: 'LastModBy',     type: 'int',   required: false, html: { page: 0, column: 0 } },
            { field: 'CreateTS',      type: 'time',  required: false, html: { page: 0, column: 0 } },
            { field: 'CreateBy',      type: 'int',   required: false, html: { page: 0, column: 0 } },
        ],
        toolbar: {
            items: [
                // { id: 'btnNotes', type: 'button', icon: 'far fa-sticky-note' },
                //{ id: 'formSave', type: 'button', caption: 'Save', icon: 'fas fa-check'},
                // { id: 'bt3', type: 'spacer' },
                // { id: 'btnClose', type: 'button', icon: 'fas fa-times' },
            ],
            onClick: function (event) {
                switch(event.target) {
                // case 'btnClose':
                //     var no_callBack = function() { return false; },
                //         yes_callBack = function() {
                //             w2ui.toplayout.hide('right',true);
                //             w2ui.reservationGrid.render();
                //         };
                //     form_dirty_alert(yes_callBack, no_callBack);
                //     break;
                case 'formSave':
                    saveReservationForm();
                }
            },
        },
        onRender: function(event) {
            var BID = getCurrentBID();
            var BUD = getBUDfromBID(BID);
            if (typeof w2ui.reservationForm.get('RTID').options != "undefined") {
                w2ui.reservationForm.get('RTID').options.items = app.rt_list[BUD];
            }
            if (typeof w2ui.reservationForm.get('BUD').options != "undefined") {
                w2ui.reservationForm.get('BUD').options.items = app.businesses;
            }
        },
        onRefresh: function(event) {
            event.onComplete = function() {
                var f = this;
                // var r = f.record;
                // var x = getCurrentBusiness(),
                //     BID=parseInt(x.value),
                //     BUD = getBUDfromBID(BID);
                //
                // var header = "Edit Reservation ({0})";
                // f.get("DPMName").options.selected = getDepMeth(BUD, dpmid);
                // f.get("DEPName").options.selected = getReservation(BUD, depid);
                //
                // formRefreshCallBack(f, "DID", header);
                var BID = getCurrentBID();
                var BUD = getBUDfromBID(BID);
                w2ui.reservationForm.get('RTID').options.items = app.rt_list[BUD];
                w2ui.reservationForm.get('BUD').options.items = app.businesses;
            };
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

                //console.log('event = ' + event);
                // formRecDiffer: 1=current record, 2=original record, 3=diff object
                // var diff = formRecDiffer(this.record, app.active_form_original, {});
                // // if diff == {} then make dirty flag as false, else true
                // if ($.isPlainObject(diff) && $.isEmptyObject(diff)) {
                //     app.form_is_dirty = false;
                // } else {
                //     app.form_is_dirty = true;
                // }
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


    //------------------------------------------------------------------------
    //          reservationGrid
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'reservationGrid',
        multiSelect: false,
        show: {
            toolbar         : false,
            footer          : false,
            toolbarAdd      : false,   // indicates if toolbar add new button is visible
            toolbarDelete   : false,   // indicates if toolbar delete button is visible
            toolbarSave     : false,   // indicates if toolbar save button is visible
            selectColumn    : false,
            expandColumn    : false,
            toolbarEdit     : false,
            toolbarSearch   : false,
            toolbarInput    : false,
            searchAll       : false,
            toolbarReload   : false,
            toolbarColumns  : false,
        },
        columns: [
            {field: 'recid',        caption: 'recid',              size: '40px',  hidden: true,  sortable: true },
            {field: 'BID',          caption: 'RID',                size: '60px',  hidden: true,  sortable: true, style: 'text-align: right'},
            {field: 'RID',          caption: 'RID',                size: '60px',  hidden: false, sortable: true, style: 'text-align: right'},
            {field: 'RentableName', caption: app.sRentable,        size: '150px', hidden: false, sortable: true, style: 'text-align: left'},
            {field: 'RTID',         caption: 'RTID',                              hidden: true,  sortable: false },
            {field: 'DtStart',      caption: 'DtStart',            size: '90px',  hidden: false, sortable: true, style: 'text-align: right'},
            {field: 'DtStop',       caption: 'DtStop',             size: '90px',  hidden: false, sortable: true, style: 'text-align: right'},
            {
                field: 'Book',
                caption: "Book it",
                size: '55px',
                style: 'text-align: center',
                render: function (record/*, index, col_index*/) {
                    // SPECIAL CHECK FOR THIS REMOVE BUTTON
                    var html = "";
                    if (record.RID && record.RID > 0) {
                        html = '<i class="far fa-calendar-check fa-lg" style="color: #2A64A4; cursor: pointer;" title="make reservation"></i>';
                    }
                    return html;
                },
            }
        ],
        onClick: function(event) {
            event.onComplete = function () {
                if(w2ui.reservationGrid.getColumn("Book", true) == event.column) {
                    var rec = w2ui.reservationGrid.get(event.recid);
                    console.log('book RID = ' + rec.RID);
                    switchToBookRes(rec.RID,rec.RentableName);
                    return;
                }
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

                // if (event.target == 'monthfwd') {  // we do these tasks after monthfwd is refreshed so we know that the 2 date controls exist
                //     setDateControlsInToolbar('asms');
                //     w2ui.expenseGrid.postData = {searchDtStart: app.D1, searchDtStop: app.D2};
                // }
            };
        }
    });

    // addDateNavToToolbar('reservation');


    //------------------------------------------------------------------------
    //         bookResForm
    //------------------------------------------------------------------------
    $().w2form({
        name: 'bookResForm',
        style: 'border: 0px; background-color: transparent;',
        header: 'Book A Reservation',
        formURL: '/webclient/html/formbookres.html',

        fields: [
            { field: 'rdBID',           type: 'int',  required: false },
            { field: 'rdBUD',           type: 'text', required: false },
            { field: 'rdRTID',          type: 'int',  required: false },
            { field: 'RID',           type: 'int',  required: false },
            { field: 'RentableName',  type: 'text', required: false },
            { field: 'Nights',        type: 'int',  required: false },
            { field: 'DtStart',       type: 'date', required: false },
            { field: 'DtStop',        type: 'date', required: false },
            { field: 'LeaseStatus',   type: 'date', required: false },
            { field: 'RLID',          type: 'int',  required: false, html: { page: 0, column: 0 } },
            { field: 'FirstName',     type: 'text', required: true,  html: { page: 0, column: 0 } },
            { field: 'LastName',      type: 'text', required: true,  html: { page: 0, column: 0 } },
            { field: 'Email',         type: 'text', required: true },
            { field: 'Phone',         type: 'text', required: true },
            { field: 'Address',       type: 'text', required: true },
            { field: 'Address2',      type: 'text', required: false },
            { field: 'City',          type: 'text', required: false, html: { page: 0, column: 0 } },
            { field: 'Country',       type: 'text', required: false, html: { page: 0, column: 0 } },
            { field: 'State',         type: 'text', required: false, html: { page: 0, column: 0 } },
            { field: 'PostalCode',    type: 'text', required: false, html: { page: 0, column: 0 } },
            { field: 'CCName',        type: 'text', required: false, html: { page: 0, column: 0 } },
            { field: 'CCType',        type: 'text', required: false, html: { page: 0, column: 0 } },
            { field: 'CCNumber',      type: 'text', required: false, html: { page: 0, column: 0 } },
            { field: 'CCExpMonth',    type: 'text', required: false, html: { page: 0, column: 0 } },
            { field: 'CCExpYear',     type: 'text', required: false, html: { page: 0, column: 0 } },
            { field: 'Comment',       type: 'text', required: false, html: { page: 0, column: 0 } },
        ],
        toolbar: {
            items: [
                { id: 'bt3', type: 'spacer' },
                { id: 'btnClose', type: 'button', icon: 'fas fa-times' },
            ],
            onClick: function (event) {
                switch(event.target) {
                case 'btnClose':
                    w2ui.reservationLayout.hide('right');
                    break;
                }
            },
        },
        actions: {
                save: function() {
                    var BID = getCurrentBID();
                    var BUD = getBUDfromBID(BID);
                    var f = w2ui.bookResForm;
                    var r = f.record;
                    f.url = '/v1/reservation/' + BID;
                    var rtid = r.rdRTID.id;
                    r.LeaseStatus = 2; // 2 = Reserved
                    r.rdRTID = rtid;
                    if (typeof r.RID === "undefined") {
                        r.RID = 0;
                    }
                    this.save({},resSaveCB);
                }
        },
        onRender: function(event) {
            // var BID = getCurrentBID();
            // var BUD = getBUDfromBID(BID);
            var f = w2ui.bookResForm;
            var r = f.record;
            //$(f.box).find("input[name=Stop]").prop( "disabled", true );
        },
        onRefresh: function(event) {
            event.onComplete = function() {
                var f = this;
                var r = f.record;
                // var BID = getCurrentBID();
                // var BUD = getBUDfromBID(BID);
                if (typeof r.RLID == "undefined" || r.RLID == 0) {
                    $(f.box).find("button[name=delete]").addClass("hidden");
                } else {
                    $(f.box).find("button[name=delete]").removeClass("hidden");
                }
                document.getElementById("bookResDtStart").innerHTML = r.DtStart;
                document.getElementById("bookResDtStop").innerHTML  = r.DtStop;
                document.getElementById("bookResNights").innerHTML  = '' + r.Nights;
                if (r.rdRTID == null) {
                    r.rdRTID = w2ui.reservationForm.record.RTID;
                }
                if (r.rdRTID.text != "undefined") {
                    document.getElementById("bookResRType").innerHTML = r.rdRTID.text;
                }
                document.getElementById("bookResRName").innerHTML = ''+r.RentableName;
                document.getElementById("bookResBUD").innerHTML = r.rdBUD;
            };
        },
        onChange: function(event) {
            // event.onComplete = function() {
            // };
        },
        onSubmit: function(target, data) {
            getFormSubmitData(data.postData.record);
        },
    });

    $().w2layout({
        name: 'reservationLayout',
        padding: 0,
        panels: [
            { type: 'left',    size: 0,     hidden: true,  content: 'left'    },
            { type: 'top',     size: '25%', hidden: false, content: 'top',    resizable: true, style: app.pstyle },
            { type: 'main',    size: '75%', hidden: false, content: 'main',   resizable: true, style: app.pstyle },
            { type: 'preview', size: 0,     hidden: true,  content: 'preview' },
            { type: 'bottom',  size: 0,     hidden: true,  content: 'bottom', resizable: true, style: app.pstyle },
            { type: 'right',   size: 0,     hidden: true,  content: 'right',  resizable: true, style: app.pstyle }
        ]
    });

    //addDateNavToToolbar('reservation');
};

window.resSaveCB = function(data) {
    console.log('save callback status:  ' + data.status);
    w2ui.reservationLayout.hide('right');
};

window.saveReservationForm = function (BID, BUD, DtStart, DtStop, RTID) {
    console.log('saveReservationForm');
};

//---------------------------------------------------------------------------------
// reservationSrch - called when the user clicks the search butting in the
//                   reservationForm.  Collect all the data and ask the server
//                   for available rentables.
//
// @params
//
// @return
//
//---------------------------------------------------------------------------------
window.reservationSrch = function() {
    var BID = getCurrentBID();
    var BUD = getBUDfromBID(BID);
    var dtStart = dateFromString(w2ui.reservationForm.record.DtStart);
    var dtStop = dateFromString(w2ui.reservationForm.record.DtStop);
    var f = {
        recid:          0,
        BID:            BID,
        BUD:            BUD,
        RTID:           w2ui.reservationForm.record.RTID.id,
        Nights:         w2ui.reservationForm.record.Nights,
        DtStart:        dtStart.toUTCString(),
        DtStop:         dtStop.toUTCString(),
    };
    w2ui.reservationGrid.postData.record = f;
    w2ui.reservationGrid.url = '/v1/available/'+BID;
    w2ui.reservationGrid.reload();
};


//---------------------------------------------------------------------------------
// reservationsUpdateRTList - changes the main view of the program to the
//      Reservations form. Note that we build the list of RentableTypes to
//      only list those rentables with bit 3 set to 0.  Bit 3 controls whether
//      the rentable should be reserved in the future after a RentalAgreement
//      term has expired. For long term rentables, that bit is set to 1
//      indicating that it is held in reserve for the long term renter. For
//      rooms such as hotel rooms, the bit is 0 indicating that it is available
//      to be rented to anyone else immediately after a RentalAgreement has
//      expired.
//
// @params  BUD - business unit descriptor where we got the rentable types
//
// @return
//
//---------------------------------------------------------------------------------
window.reservationsUpdateRTList = function (BUD) {
    if (typeof w2ui.reservationForm.get('RTID').options != "undefined") {
        var rtlist = [];
        var i;
        for (i = 0; i < app.rt_list[BUD].length; i++) {
            if (0 === (app.rt_list[BUD][i].FLAGS & 8)) {
                rtlist.push( app.rt_list[BUD][i]);
            }
        }
        w2ui.reservationForm.get('RTID').options.items = rtlist;  // app.rt_list[BUD];
        if (0 === w2ui.reservationForm.record.RTID) {
            w2ui.reservationForm.record.RTID = rtlist[0].id; // app.rt_list[BUD][0].id;
        }
        w2ui.reservationForm.render();
    }
};

//---------------------------------------------------------------------------------
// switchToReservations - changes the main view of the program to the
//                        Reservations form
//
// @params
//          svcOverride = name of webservice to call if the name does not
//                match the name of the svc
//
// @return
//
//---------------------------------------------------------------------------------
window.switchToReservations = function (svcOverride) {
    // w2ui[grid].url = url;
    w2ui.reservationForm.last.sel_recid = null; // whenever switch grid, erase last selected record
    app.last.grid_sel_recid = -1;
    app.active_grid = "reservationForm"; // mark active grid in app.active_grid
    w2ui.toplayout.content('main', w2ui.reservationLayout);
    w2ui.toplayout.hide('right',true);

    var BID = getCurrentBID();
    var BUD = getBUDfromBID(BID);
    w2ui.reservationForm.record = getReservationInitRecord(BID,BUD);
    getRentableTypes(BUD,reservationsUpdateRTList);
};

//---------------------------------------------------------------------------------
// switchToBookRes - turns on the bookres form
//
// @params
//     RID - Rentable ID
//     n   - Rentable Name
// @return
//
//---------------------------------------------------------------------------------
window.switchToBookRes = function (RID,n) {
    var BID = getCurrentBID();
    var BUD = getBUDfromBID(BID);
    w2ui.reservationLayout.sizeTo('right',500);
    w2ui.reservationLayout.show('right');
    w2ui.bookResForm.record = getBookResInitRecord(BID,BUD);
    w2ui.bookResForm.record.RID = RID;
    w2ui.bookResForm.record.RentableName = n;
    w2ui.bookResForm.refresh();
};

//---------------------------------------------------------------------------------
// finishReservationsForm - load the layout properly.
// @params
//
// @return
//
//---------------------------------------------------------------------------------
window.finishReservationsForm = function () {
    w2ui.reservationLayout.content('top',w2ui.reservationForm);
    w2ui.reservationLayout.content('main', w2ui.reservationGrid);
    w2ui.reservationLayout.content('right', w2ui.bookResForm);
};
