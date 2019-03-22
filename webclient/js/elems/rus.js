/*global
    setDefaultFormFieldAsPreviousRecord, w2uiDateControlString, $, w2ui, app, getCurrentBusiness, parseInt, getBUDfromBID,
    getRentableTypes, setToForm, form_dirty_alert, console, getFormSubmitData, addDateNavToToolbar, setRentableLayout,
    getRentableInitRecord, saveRentableLeaseStatus, RentableEdits, dtFormatISOToW2ui, addRentableUseStatus, datetimeFmtStr,
*/
/*jshint esversion: 6 */

"use strict";

window.buildRentableUseStatusElements = function () {

    //------------------------------------------------------------------------
    //          rentable Use Status Grid
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'rentableUseStatusGrid',
        style: 'padding: 0px',
        // url: '/v1/rentableusestatus',
        url: '',
        show: {
            header: false,
            toolbar: true,
            toolbarReload: false,
            toolbarColumns: false,
            toolbarSearch: true,
            toolbarAdd: true,
            toolbarDelete: false,
            toolbarSave: false,
            searchAll: true,
            footer: true,
            lineNumbers: false,
            selectColumn: false,
            expandColumn: false
        },
        columns: [
            {field: 'recid', caption: 'recid', hidden: true},
            {field: 'RID', caption: 'RID', hidden: true},
            {field: 'BID', caption: 'BID', hidden: true},
            {field: 'BUD', caption: 'BUD', hidden: true},
            {field: 'RSID', caption: 'RSID', size: '50px'},
            {
                field: 'UseStatus', caption: 'UseStatus', size: '150px',
                editable: {type: 'select', align: 'left', items: app.RSUseStatusItems},
                render: function (record, index, col_index) {
                    var html = '';
                    for (var s in app.RSUseStatusItems) {
                        if (app.RSUseStatusItems[s].id == this.getCellValue(index, col_index)) {
                            html = app.RSUseStatusItems[s].text;
                        }
                    }
                    return html;
                },
            },
            {
                field: 'DtStart',
                caption: 'DtStart',
                size: "50%",
                sortable: true,
                style: 'text-align: right',
                editable: {type: 'datetime'},
                render: function(rec,row,col) {
                    return '' + dtFormatISOToW2ui(rec.DtStart);
                },
            },
            {
                field: 'DtStop',
                caption: 'DtStop',
                size: "50%",
                sortable: true,
                style: 'text-align: right',
                editable: {type: 'datetime'},
                render: function(rec,row,col) {
                    return '' + dtFormatISOToW2ui(rec.DtStop);
                },
            },
            {field: 'CreateBy', caption: 'CreateBy', hidden: true},
            {field: 'LastModBy', caption: 'LastModBy', hidden: true},
        ],
        onAdd: function (/*event*/) {
            addRentableUseStatus();
        },
        onLoad: function (event) {
            //------------------------------------------------------------------------
            // We only want the grids to request server data on their initial load
            // and on a RentableForm Save.  So, we will clear them after the
            // grids complete their loading or after a save completes.
            //------------------------------------------------------------------------
            event.onComplete = function () {
                // var BID = getCurrentBID();
                // var RID = w2ui.rentableForm.record.RID;
                // this.url = '/v1/rentableusestatus/'+BID+'/'+RID;
                this.url = '';

                //---------------------------------------------------------------------
                // every datetime value needs to be converted to a localtime string...
                //---------------------------------------------------------------------
                for (var i = 0; i < this.records.length; i++) {
                    var d = new Date(this.records[i].DtStart);
                    this.records[i].DtStart = datetimeFmtStr(d);
                }
            };
        },
        onSave: function (event) {
            this.url = '';
            var UseUnknownStatus, UseInactiveStatus;
            app.RSUseStatusItems.forEach(function (status) {
                switch (status.text) {
                    case "Unknown":
                        UseUnknownStatus = status.id;
                        break;
                    case "Inactive":
                        UseInactiveStatus = status.id;
                        break;
                }
            });
            event.changes = this.records;
            event.onComplete = function() {
                // var BID = getCurrentBID();
                // var RID = w2ui.rentableForm.record.RID;
                // w2ui.rentableUseStatusGrid.url = "/v1/rentableusestatus/" + BID + "/" + RID;
                this.url = '';
            };
        },
        // onDelete: function (event) {
        //     var selected = this.getSelection();
        //     var RSIDList = [];
        //     var grid = this;
        //
        //     // if not selected then return
        //     if (selected.length < 0) {
        //         return;
        //     }
        //     // collect the RSIDs to remove
        //     selected.forEach(function (id) {
        //         var r = grid.get(id);
        //         if ( 0 != r.RSID ) {
        //             RSIDList.push(r.RSID);  // only save the ones already in the server db
        //         }
        //     });
        //
        //     event.onComplete = function () {
        //         //-----------------------------------------------------------
        //         // If the record has not yet been saved to the server, then
        //         // just remove it from the grid and we're done with it.
        //         //-----------------------------------------------------------
        //         var grid = this;
        //         var Unselect = [];
        //         for (var i = 0; i < selected.length; i++) {
        //             var r = grid.get(selected[i]);
        //             if (0 == r.RSID) {
        //                 Unselect.push(selected[i]);
        //             }
        //         }
        //
        //         grid.selectNone();
        //         grid.select.apply(Unselect);
        //         grid.delete(true);  // get rid of them
        //
        //         if (RSIDList.length == 0 ) {
        //             RentableEdits.rusDeleteInProgress = false;
        //             return;
        //         }
        //
        //         var BID = getCurrentBID();
        //         var RID = w2ui.rentableForm.record.RID;
        //
        //         var tgrid = w2ui.rentableUseStatusGrid;
        //         var url = "/v1/rentableusestatus/" + BID + "/" + RID;
        //         var params = {"cmd": "delete", "RSIDList": RSIDList};
        //         var dat = JSON.stringify(params);
        //
        //         $.post(grid.url, dat, null, "json")
        //         .done(function(data) {
        //             if (data.status === "error") {
        //                 grid.error('onDelete: '+w2utils.lang(data.message));
        //                 return;
        //             }
        //             grid.reload();
        //         })
        //         .fail(function(/*data*/){
        //             grid.error("Delete RentableUseStatus failed.");
        //             return;
        //         });
        //     };
        // },
        onChange: function (event) {
            event.preventDefault();
            var g = this;
            var field = g.columns[event.column].field;
            var chgRec = g.get(event.recid);
            var changeIsValid = true;

            //------------------------------------
            // Put any validation checks here...
            //------------------------------------
            if (event.value_new == "" && (g.columns[event.column].field == "DtStop" || g.columns[event.column].field == "DtStart")) {
                changeIsValid = false;
            }

            //---------------------------------------------------
            // Inform w2ui if the change is cancelled or not...
            //---------------------------------------------------
            event.isCancelled = !changeIsValid;

            //---------------------------------------------------------------
            // 2/19/2019 sman - This save is used to save the data into the
            // grid's records.  We need to ensure that the grids URL is ''
            //---------------------------------------------------------------
            event.onComplete = function () {
                var dt;
                if (!event.isCancelled) { // if event not cancelled then invoke save method
                    RentableEdits.UseStatusChgList.push(chgRec.recid);
                    g.url = '';  // just ensure that no server service is called
                    if (g.columns[event.column].field == "DtStop") {
                        dt=new Date(event.value_new);
                        g.records[event.recid].DtStop = datetimeFmtStr(dt);
                    } else if (g.columns[event.column].field == "DtStart") {
                        dt=new Date(event.value_new);
                        g.records[event.recid].DtStart = datetimeFmtStr(dt);
                    }
                    this.save(); // save automatically locally
                    // var BID = getCurrentBusiness();
                    // var RID = w2ui.rentableForm.record.RID;
                    // g.url = '/v1/rentableusestatus/' + BID + '/' + RID;
                }
            };
        }
    });
};

// saveRentableUseStatus - creates a list of UseStatus entries that have
// been changed, then calls the webservice to save them. Note that every
// datetime value needs to be converted to UTC prior to saving to server.
//---------------------------------------------------------------------------
window.saveRentableUseStatus = function(BID,RID) {
    var reclist = Array.from(new Set(RentableEdits.UseStatusChgList));

    if (reclist.length == 0) {
        return Promise.resolve('{"status": "success"}');
    }

    var chgrec = [];
    var dt;
    for (var i = 0; i < reclist.length; i++) {
        var recid = reclist[i];
        var nrec =  w2ui.rentableUseStatusGrid.get(recid);
        if (typeof nrec.UseStatus == "string") {
            var ls = parseInt(nrec.UseStatus,10);
            nrec.UseStatus = ls;
        }
        //-----------------------------------------------------------
        // translate all localtimes to UTC before sending to server
        //-----------------------------------------------------------
        dt = new Date(nrec.DtStart);
        nrec.DtStart = dt.toUTCString();
        dt = new Date(nrec.DtStop);
        nrec.DtStop = dt.toUTCString();
        chgrec.push(nrec);
    }

    var params = {
        cmd: "save",
        selected: [],
        limit: 0,
        offset: 0,
        changes: chgrec,
        RID: w2ui.rentableForm.record.RID
    };

    var dat = JSON.stringify(params);
    var url = '/v1/rentableusestatus/' + BID + '/' + w2ui.rentableForm.record.RID;
    return $.post(url, dat, null, "json")
    .done(function(data) {
        if (data.status === "success") {
            RentableEdits.UseStatusChgList = []; // reset the change list now, because we've saved them
            w2ui.toplayout.hide('right', true);
            // w2ui.rentablesGrid.render();
        } else {
            w2ui.rentablesGrid.error('saveRentableUseStatus: '+data.status);
        }
    })
    .fail(function(data){
        console.log("Save RentableUseStatus failed.");
    });
};

// addRentableUseStatus - creates a new RentableUseStatus entry and adds it
// to the grid.
//
// @params
//
// @return
//---------------------------------------------------------------------------
window.addRentableUseStatus = function() {
    var  x = getCurrentBusiness();
    var BID = parseInt(x.value);
    var BUD = getBUDfromBID(BID);
    var fr = w2ui.rentableForm.record;
    var g = w2ui.rentableUseStatusGrid;
    var ndStart = new Date();

    // get lastest date among all market rate object's stopDate for new MR's StartDate
    for (var i = 0; i < g.records.length; i++) {
        var rec = g.records[i];
        if (rec.DtStop) {
            var rdStop = new Date(rec.DtStop);
            if (ndStart < rdStop) {
                ndStart = rdStop;
            }
        }
    }

    var newRec = {
        recid: g.records.length,
        BID: BID,
        BUD: BUD,
        RID: fr.RID,
        RSID: 0,
        UseStatus: 0,
        LeaseStatus: 0,
        DtStart: datetimeFmtStr(ndStart),
        DtStop: "12/31/9999 12:00:00 am"
    };
    // EDI does not apply to Use Status -- which is a datetime.  EDIT applies
    // to date-only Fields
    //------------------------------------------------------------------------
    // if (EDIEnabledForBUD(BUD)) {
    //     var d = ndStart;
    //     d.setDate(d.getDate()+1);
    //     newRec.DtStart = datetimeFmtStr(d);
    //     newRec.DtStop = "12/30/9999 12:00:00 am";
    // }
    var d1 = new Date(newRec.DtStart);
    var d2 = new Date(newRec.DtStop);
    if (d1 > d2) {
        newRec.DtStart = datetimeFmtStr(d1);
    }
    g.add(newRec,true); // true forces the add to the beginning of the list
    RentableEdits.UseStatusChgList.push(newRec.recid);
};
