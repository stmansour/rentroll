/*global
    parseInt, w2ui, getDepMeth, getReservation, $, app, getBUDfromBID, getCurrentBusiness, console,
    saveReservationForm, switchToReservationss, finishReservationsForm, reservationsUpdateRTList,
    getReservationInitRecord, reservationSrch, daysBetweenDates, switchToBookRes,
    getBookResInitRecord, resSaveCB, setToForm, setResUpdateRecordForUI,
    showReservationRentable, checkRentableAvailability, cancelReservation, getRTName,
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
                render: function (record, index, col_index) {
                    if (record == undefined) {
                        return '';
                    }
                    var d = new Date(record.DtStart);
                    return dateFmtStr(d); // if non-recur assessment then do nothing
                },
            },
            {field: 'DtStop',           caption: 'DtStop',          size: '80px', hidden: false, sortable: true, style: 'text-align: center',
                render: function (record , index, col_index) {
                    if (record == undefined) {
                        return '';
                    }
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
                    w2ui.availabilityGrid.clear(); // remove any contents from prior checks
                    w2ui.availabilityGrid.url = '';
                    var url = '/v1/reservation/' + BID + '/' + rec.RLID;
                    setToForm('resUpdateForm',url,750,true,w2ui.resFormLayout);
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
    // resFormLayout
    //
    //  top    = resUpdateForm
    //  main   = availabilityGrid
    //  bottom = resUpFormBtns
    //------------------------------------------------------------------------
    $().w2layout({
        name: 'resFormLayout',
        padding: 0,
        panels: [
            { type: 'left',    size: 0,     hidden: true,  content: 'left'    },
            { type: 'top',     size: '50%', hidden: false, content: 'top',    resizable: true, style: app.pstyle },
            { type: 'main',    size: '50%', hidden: false,  content: 'main',   resizable: true, style: app.pstyle },
            { type: 'preview', size: 0,     hidden: true,  content: 'preview' },
            { type: 'bottom',  size: '50px',hidden: false,  content: 'bottom', resizable: false, style: app.pstyle },
            { type: 'right',   size: 0,     hidden: true,  content: 'right',  resizable: true, style: app.pstyle }
        ]
    });

    //---------------------------------------------------------------------------------
    // finishResUpForm - load the layout properly.
    //  top    = resUpdateForm
    //  main   = availabilityGrid
    //  bottom = resUpFormBtns
    //---------------------------------------------------------------------------------
    window.finishResUpForm = function () {
        w2ui.resFormLayout.content('top', w2ui.resUpdateForm);
        w2ui.resFormLayout.content('main', w2ui.availabilityGrid);
        w2ui.resFormLayout.content('bottom', w2ui.resUpFormBtns);
    };

    //------------------------------------------------------------------------
    //          resUpdateForm
    //
    //    >>>>>>>>>   top    = resUpdateForm       <<<<<<<<<<<
    //                main   = availabilityGrid
    //                bottom = resUpFormBtns
    //------------------------------------------------------------------------
    $().w2form({
        name: 'resUpdateForm',
        style: 'border: 0px; background-color: transparent;',
        header: 'Change A Reservation',
        url: '/v1/reservation/',
        formURL: '/webclient/html/formresup.html',
        fields: [
            { field: 'BUD',              type: 'list',     required: false,  options: {items: app.businesses} },
            { field: 'RLID',             type: 'int',      required: false,  },
            { field: 'RTRID',            type: 'int',      required: false,  },
            { field: 'RTID',             type: 'int',      required: false,  },
            { field: 'rdRTID',           type: 'list',     required: false,  },
            { field: 'RID',              type: 'int',      required: false,  },
            { field: 'rdBID',            type: 'int',      required: false,  },
            { field: 'ConfirmationCode', type: 'text',     required: false,  },
            { field: 'LeaseStatus',      type: 'int',      required: false,  },
            { field: 'DtStart',          type: 'date',     required: false,  },
            { field: 'DtStop',           type: 'date',     required: false,  },
            { field: 'Nights',           type: 'int',      required: false,  },
            { field: 'FirstName',        type: 'text',     required: true ,  },
            { field: 'LastName',         type: 'text',     required: true ,  },
            { field: 'Phone',            type: 'text',     required: false,  },
            { field: 'Email',            type: 'text',     required: false,  },
            { field: 'Street',           type: 'text',     required: false,  },
            { field: 'City',             type: 'text',     required: false,  },
            { field: 'State',            type: 'text',     required: false,  },
            { field: 'PostalCode',       type: 'text',     required: false,  },
            { field: 'CCName',           type: 'text',     required: false,  },
            { field: 'CCType',           type: 'text',     required: false,  },
            { field: 'CCNumber',         type: 'text',     required: false,  },
            { field: 'CCExpMonth',       type: 'text',     required: false,  },
            { field: 'CCExpYear',        type: 'text',     required: false,  },
            { field: 'Comment',          type: 'textarea', required: false,  },
            { field: 'LastModTime',      type: 'time',     required: false,  },
            { field: 'LastModBy',        type: 'int',      required: false,  },
            { field: 'CreateTS',         type: 'time',     required: false,  },
            { field: 'CreateBy',         type: 'int',      required: false,  },
        ],
        toolbar: {
            items: [
                { id: 'cancelReservation', type: 'button', caption: 'Cancel Reservation', icon: "fas fa-ban"},
                { id: 'bt3',               type: 'spacer' },
                { id: 'btnClose',          type: 'button', icon: 'fas fa-times' },
            ],
            onClick: function (event) {
                switch(event.target) {
                case 'btnClose':
                    var no_callBack = function() { return false; },
                        yes_callBack = function() {
                            w2ui.toplayout.hide('right',true);
                            w2ui.resUpdateGrid.render();
                        };
                    form_dirty_alert(yes_callBack, no_callBack);
                    break;
                case 'cancelReservation':
                    var f = w2ui.resUpdateForm;
                    var r = f.record;
                    cancelReservation(r.RLID);
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
                var y = new Date(r.DtStart);
                r.DtStart = dateFmtStr(y);
                var x = new Date(r.DtStop);
                r.DtStop = dateFmtStr(x);
                f.record.Nights = daysBetweenDates(x,y);
            };
        },
        onRender: function(event) {
            setResUpdateRecordForUI(this);
        },
        onRefresh: function(event) {
            setResUpdateRecordForUI(this);
            showReservationRentable(this.record);
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
    //------------------------------------------------------------------------
    //    availabilityGrid
    //
    //                top    = resUpdateForm
    //    >>>>>>>>>   main   = availabilityGrid    <<<<<<<<<<<
    //                bottom = resUpFormBtns
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'availabilityGrid',
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
            {field: 'RID',          caption: 'RID',                size: '45px',  hidden: false, sortable: true, style: 'text-align: right'},
            {field: 'RentableName', caption: app.sRentable,        size: '150px', hidden: false, sortable: true, style: 'text-align: left'},
            {field: 'RTID',         caption: 'RTID',               size: '150px', hidden: false,  sortable: false,
                render: function(record/*, index, col_index*/) {
                    return getRTName(record.RTID);
                },
            },
            {field: 'DtStart',      caption: 'DtStart',            size: '90px',  hidden: false, sortable: true, style: 'text-align: right',
                render: function(record/*, index, col_index*/) { return w2uiDateControlString(new Date(record.DtStart)); },
            },
            {field: 'DtStop',       caption: 'DtStop',             size: '90px',  hidden: false, sortable: true, style: 'text-align: right',
                render: function(record/*, index, col_index*/) { return w2uiDateControlString(new Date(record.DtStop)); },
            },
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
                var rec = w2ui.availabilityGrid.get(event.recid);
                // console.log('book RID = ' + rec.RID);
                showReservationRentable(rec);
                return;
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

    //------------------------------------------------------------------------
    //    resUpFormBtns
    //
    //                top    = resUpdateForm
    //                main   = availabilityGrid
    //    >>>>>>>>>   bottom = resUpFormBtns      <<<<<<<<<
    //------------------------------------------------------------------------
    $().w2form({
        name: 'resUpFormBtns',
        style: 'border: 0px; background-color: transparent;',
        formURL: '/webclient/html/formresupbtns.html',
        url: '',
        fields: [],
        actions: {
            save: function () {
                console.log("Book this res!");
                var BID = getCurrentBID();
                var BUD = getBUDfromBID(BID);
                var f = w2ui.resUpdateForm;
                var r = f.record;
                f.url = '/v1/reservation/' + BID;
                var rtid = r.rdRTID.id;
                r.LeaseStatus = 2; // 2 = Reserved
                r.rdRTID = rtid;
                if (typeof r.RID === "undefined") {
                    r.RID = 0;
                }
                //this.save({},resSaveCB);

                // w2ui.rentablesGrid.selectNone();
                // saveRentableCore(finishRentableSave);
            },

            saveadd: function () {
                // saveRentableCore(finishRentableSaveAdd);
            }
        },
    });
};


//---------------------------------------------------------------------------------
// showReservationRentable - use the supplied record on the resUpdateForm
//     as the rentable to use for this reservation
//
// @params
//      f = the form
// @return
//
//---------------------------------------------------------------------------------
window.showReservationRentable = function(rec) {
    var f = w2ui.resUpdateForm;
    var r = f.record;
    r.RID = rec.RID;
    r.RTID = rec.RTID;
    if (r.RTID == undefined) {
        r.RTID = rec.rdRTID;
    }
    var s = rec.RentableName;
    if (typeof r.RTID  != undefined) {
        s += ' &nbsp;&nbsp(' + getRTName(r.RTID) + ')';
    }
    document.getElementById("reservationRentableName").innerHTML = s;
    var d1 = new Date(rec.DtStart);
    var d2 = new Date(rec.DtStop);
    s = 'available: ' + w2uiDateControlString(d1) + ' - ' + w2uiDateControlString(d2);
    document.getElementById("reservationRentableFreePeriod").innerHTML = s;
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

//---------------------------------------------------------------------------------
// checkRentableAvailability - pull together the info in the form, make a query
//     to the server to see what's available.
//
// @params
//     f = the form
// @return
//
//---------------------------------------------------------------------------------
window.checkRentableAvailability = function() {
    var BID = getCurrentBID();
    var BUD = getBUDfromBID(BID);
    var f = w2ui.resUpdateForm;
    var r = f.record;
    var dtStart = dateFromString(r.DtStart);
    var dtStop = dateFromString(r.DtStop);
    var req = {
        recid:          0,
        BID:            BID,
        BUD:            BUD,
        RTID:           r.rdRTID.id,
        Nights:         r.Nights,
        DtStart:        dtStart.toUTCString(),
        DtStop:         dtStop.toUTCString(),
    };
    w2ui.availabilityGrid.postData.record = req;
    w2ui.availabilityGrid.url = '/v1/available/'+BID;
    w2ui.availabilityGrid.reload();
};

//---------------------------------------------------------------------------------
// getRTName - search for and return the rentable type name. return an empty string
//     if not found.
//
// @params
//     RTID = the RenttableType ID
//
// @return
//     the rentable type name or '' if not found
//---------------------------------------------------------------------------------
window.getRTName = function(RTID) {
    var BUD = getBUDfromBID(getCurrentBID());
    for (var i = 0; i < app.rt_list[BUD].length; i++) {
        if (app.rt_list[BUD][i].id == RTID) {
            return app.rt_list[BUD][i].text;
        }
    }
    return '';
};

//---------------------------------------------------------------------------------
// cancelReservation - cancel the supplied RLID
//
// @params
//     RLID = the Reservation ID
//
// @return
//
//---------------------------------------------------------------------------------
window.cancelReservation = function(RLID) {
    console.log("cancel RLID = " + RLID);
};
