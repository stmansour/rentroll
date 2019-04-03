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
            { field: 'DtStart', caption: 'DtStart', size: "50%", sortable: true, style: 'text-align: right', editable: {type: 'datetime'},
              render: function(rec,row,col) { return '' + dtFormatISOToW2ui(rec.DtStart); },
            },
            { field: 'DtStop', caption: 'DtStop', size: "50%", sortable: true, style: 'text-align: right', editable: {type: 'datetime'},
                render: function(rec,row,col) { return '' + dtFormatISOToW2ui(rec.DtStop); },
            },
            {field: 'CreateBy', caption: 'CreateBy', hidden: true},
            {field: 'LastModBy', caption: 'LastModBy', hidden: true},
        ],
        onAdd: function (/*event*/) {
            addRentableUseStatus();
        },
        onLoad: function (event) {
            //------------------------------------------------------------------------
            // We need the URL to be active in case virtual scrolling is needed.
            // We will turn off the url when a local save is completed.
            //------------------------------------------------------------------------
            event.onComplete = function () {
                var BID = getCurrentBID();
                var RID = w2ui.rentableForm.record.RID;
                this.url = '/v1/rentableusestatus/'+BID+'/'+RID;
                RentableEdits.UseStatusDataLoaded = true;

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
            // NOTE:
            // I don't know if the code block below is necessary
            //--------------
            // var UseUnknownStatus, UseInactiveStatus;
            // app.RSUseStatusItems.forEach(function (status) {
            //     switch (status.text) {
            //         case "Unknown":
            //             UseUnknownStatus = status.id;
            //             break;
            //         case "Inactive":
            //             UseInactiveStatus = status.id;
            //             break;
            //     }
            // });
            // event.changes = this.records;

            //------------------------------------------------------------------
            // Grid changes are saved locally. So we want no url when this funct
            // gets called due to a grid change. After the local save is complete
            // we put the url back so that if the virtual scrolling needs to call
            // the server it can.  The full save to disk is done when the user
            // presses the Save button, which is handled by a different function.
            //------------------------------------------------------------------
            this.url = '';  // no url for local save in the grid
            event.onComplete = function() {  // restore the url to support virtual scrolling
                var BID = getCurrentBID();
                var RID = w2ui.rentableForm.record.RID;
                this.url = "/v1/rentableusestatus/" + BID + "/" + RID;
            };
        },
        onChange: function (event) {
            // event.preventDefault();
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

            event.onComplete = function () {
                //-------------------------------------------------------------------
                // 2/19/2019 sman - Store into the grid's records.  We need to
                // ensure that the grid's URL is ''
                //-------------------------------------------------------------------
                var dt;
                if (!event.isCancelled) { // if event not cancelled then invoke save method
                    RentableEdits.UseStatusChgList.push({index: event.index, ID: g.records[event.index].RSID});
                    if (g.columns[event.column].field == "DtStop") {
                        dt=new Date(event.value_new);
                        g.records[event.index].DtStop = datetimeFmtStr(dt);
                    } else if (g.columns[event.column].field == "DtStart") {
                        dt=new Date(event.value_new);
                        g.records[event.index].DtStart = datetimeFmtStr(dt);
                    }
                    g.url = '';  // just ensure that no server service is called
                    this.save(); // save automatically locally
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
    var list = [];
    var i;

    //------------------------------------------------
    // Build a list of IDs. that were edited...
    //------------------------------------------------
    for (i = 0; i < RentableEdits.UseStatusChgList.length ; i++) {
        list[i] = RentableEdits.UseStatusChgList[i].ID;
    }

    //------------------------------------------------
    // Filter the list to the unique members...
    //------------------------------------------------
    var reclist = Array.from(new Set(list));

    //------------------------------------------------
    // If there's nothing in the list, we're done.
    // Set the promise to resolved and return.
    //------------------------------------------------
    if (reclist.length == 0) {
        return Promise.resolve('{"status": "success"}');
    }

    //------------------------------------------------
    // Find and process the record for each ID
    //------------------------------------------------
    var chgrec = [];
    var dt;
    var index = -1;
    var grid = w2ui.rentableUseStatusGrid;
    for (i = 0; i < reclist.length; i++) {
        //------------------------------------------------------------
        // reclist[i] is the id of the element we wnat to find...
        //------------------------------------------------------------
        for (var j = 0; j < grid.records.length; j++ ) {
            if (grid.records[j].RSID == reclist[i]) {
                index = j;
                break;
            }
        }
        //------------------------------------------------------------
        // if the ID could not be found resolve promis with an
        // error message.
        //------------------------------------------------------------
        if (index < 0) {
            var s='ERROR: could not find RSID = '+reclist[i];
            w2ui.rentablesGrid.error(s);  // place an error where we will be sure to see it
            return Promise.resolve('{"status": "error", "message": s}');
        }

        //------------------------------------------------------------
        //  This is the record we need to save.  Make any last-min
        //  changes...
        //------------------------------------------------------------
        var nrec = grid.records[index];
        if (typeof nrec.UseStatus == "string") {
            var ls = parseInt(nrec.UseStatus,10);
            nrec.UseStatus = ls;
        }
        if (nrec.RSID < 0) {
            nrec.RSID = 0;  // server needs RSID = 0 for new records
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
            //------------------------------------------------------------------
            // Now that the save is complete, we can add the URL back to the
            // the grid so it can call the server to get updated rows. The
            // onLoad handler will reset the url to '' after the load completes
            // so that changes are done locally to gthe grid until the
            // rentableForm save button is clicked.
            //------------------------------------------------------------------
            RentableEdits.UseStatusChgList = []; // reset the change list now, because we've saved them
            w2ui.rentableUseStatusGrid.url = url;
        } else {
            w2ui.rentablesGrid.error('saveRentableUseStatus: '+data.message);
        }
    })
    .fail(function(data){
        w2ui.rentablesGrid.error("Save RentableUseStatus failed. " + data);
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
    var ndStart;

    // get lastest date among all market rate object's stopDate for new MR's StartDate
    for (var i = 0; i < g.records.length; i++) {
        var rec = g.records[i];
        if (typeof ndStart === "undefined") {
            ndStart = new Date(rec.DtStop);
        }
        if (rec.DtStop) {
            var rdStop = new Date(rec.DtStop);
            if (ndStart < rdStop) {
                ndStart = rdStop;
            }
        }
    }
    if (typeof ndStart === "undefined") {
        ndStart = new Date();
    }

    var newRec = {
        recid: g.records.length,
        BID: BID,
        BUD: BUD,
        RID: fr.RID,
        RSID: RentableEdits.RSID,
        UseStatus: 0,
        LeaseStatus: 0,
        DtStart: datetimeFmtStr(ndStart),
        DtStop: "12/31/9999 12:00:00 am"
    };

    RentableEdits.RSID--;

    //------------------------------------------------------------------------
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
    RentableEdits.UseStatusChgList.push({index: 0, ID: newRec.RSID});
    g.add(newRec,true); // true forces the add to the beginning of the list
};
