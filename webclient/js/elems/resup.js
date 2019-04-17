/*global
    parseInt, w2ui, getDepMeth, getReservation, $, app, getBUDfromBID, getCurrentBusiness, console,
    saveReservationForm, switchToReservationss, finishReservationsForm, reservationsUpdateRTList,
    getReservationInitRecord, reservationSrch, daysBetweenDates, switchToBookRes,
    getBookResInitRecord, resSaveCB,
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
            {field: 'Name',             caption: 'Name',            size: '90px', hidden: false, sortable: true },
                ],
        onClick: function(event) {
            event.onComplete = function () {
                if(w2ui.resUpdateGrid.getColumn("Book", true) == event.column) {
                    var rec = w2ui.resUpdateGrid.get(event.recid);
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
            };
        }
    });

    addDateNavToToolbar('resUpdate');
};

//---------------------------------------------------------------------------------
// switchToResUpdate - changes the main view of the program to the
//                        Reservations form
//
// @params
//          svcOverride = name of webservice to call if the name does not
//                match the name of the svc
//
// @return
//
//---------------------------------------------------------------------------------
window.switchToResUpdate = function (svcOverride) {
    var g = w2ui.resUpdateGrid;
    // w2ui[grid].url = url;
    w2ui.reservationForm.last.sel_recid = null; // whenever switch grid, erase last selected record
    app.last.grid_sel_recid = -1;
    app.active_grid = "resUpdateGrid"; // mark active grid in app.active_grid
    w2ui.toplayout.content('main',g);
    w2ui.toplayout.hide('right',true);

    var BID = getCurrentBID();
    g.url = '/v1/reservation/' + BID;
};
