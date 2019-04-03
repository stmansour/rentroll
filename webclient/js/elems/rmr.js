/*global
    w2ui, app, $, console, form_dirty_alert, formRefreshCallBack, formRecDiffer,
    getFormSubmitData, w2confirm, delete_confirm_options, getBUDfromBID, getCurrentBusiness,
    addDateNavToToolbar, setRTLayout, getRTInitRecord, getRentASMARList, showForm,
    RTEdits, addRentableMarketRate, w2uiDateTimeControlString,
*/
/*jshint esversion: 6 */

"use strict";
window.buildRentableMarketRateElements = function() {
    //------------------------------------------------------------------------
    //          rentable Market Rates Grid
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'rmrGrid',
        style: 'padding: 0px',
        show: {
            header: false,
            toolbar: true,
            toolbarReload: false,
            toolbarColumns: false,
            toolbarSearch: true,
            toolbarAdd: true,
            toolbarDelete: false,
            toolbarSave: false,
            searchAll       : true,
            footer: true,
            lineNumbers: false,
            selectColumn: false,
            expandColumn: false
        },
        columns: [
            {field: 'recid', caption: 'recid', hidden: true},
            {field: 'RMRID', caption: 'RMRID', size: '150px', sortable: true},
            {field: 'RTID', caption: 'RTID', size: '150px', hidden: true},
            {field: 'BID', caption: 'BID', hidden: true},
            {field: 'BUD', caption: 'BUD', hidden: true},
            {field: 'MarketRate',  caption: 'MarketRate',  size: '100px', sortable: true, render: 'money', editable: {type: 'money'} },
            {field: 'DtStart',     caption: 'DtStart', size: "50%",    sortable: true, style: 'text-align: right', editable: {type: 'date'} },
            {field: 'DtStop',      caption: 'DtStop', size: "50%",    sortable: true, style: 'text-align: right', editable: {type: 'date'} },
        ],
        onLoad: function(event) {
            //------------------------------------------------------------------------
            // We need the URL to be active in case virtual scrolling is needed.
            // We will turn off the url when a local save is completed.
            //------------------------------------------------------------------------
            event.onComplete = function () {
                var BID = getCurrentBID();
                var RTID = w2ui.rtForm.record.RTID;
                w2ui.rmrGrid.url = '/v1/rmr/' + BID + '/' + RTID;
                RTEdits.MarketRateDataLoaded = true;
            };
        },
        onAdd: function(event) {
            addRentableMarketRate();
        },
        onSave: function(event) {
            // sman: I don't know if this is correct...
            // I'm commenting it out for now
            // event.changes = this.records;

            //------------------------------------------------------------------
            // Grid changes are saved locally. So we want no url when this funct
            // gets called due to a grid change. After the local save is complete
            // we put the url back so that if the virtual scrolling needs to call
            // the server it can.  The full save to disk is done when the user
            // presses the Save button, which is handled by a different function.
            //------------------------------------------------------------------
            this.url = '';  // save is done locally here...
            event.onComplete = function() {
                var BID = getCurrentBID();
                var RTID = w2ui.rtForm.record.RTID;
                this.url = '/v1/rmr/' + BID + '/' + RTID;
            };
        },
        onChange: function(event) {
            event.preventDefault();
            var g = this,
                field = g.columns[event.column].field,
                chgRec = g.get(event.recid),
                changeIsValid = true;

            if ( field === "MarketRate" ) { // if field is MarketRate
                if (event.value_new < 0) {
                    changeIsValid = false;
                }
            }

            //------------------------------------
            // Put any validation checks here...
            //------------------------------------
            if (event.value_new == "" && (g.columns[event.column].field == "DtStop" || g.columns[event.column].field == "DtStart")) {
                changeIsValid = false;
            }

            //------------------------------------
            // DtStart must be prior to DtStop...
            //------------------------------------
            var DtStart, DtStop;
            if (field == "DtStop") {
                DtStart = new Date(g.records[event.index].DtStart);
                DtStop = new Date(event.value_new);
                if (DtStart > DtStop) {
                    changeIsValid = false;
                    g.error("DtStop date must be after DtStart. DtStop has been reset to its previous value.");
                }
            } else {
                DtStart = new Date(event.value_new);
                DtStop = new Date(g.records[event.index].DtStop);
                if (DtStart > DtStop) {
                    changeIsValid = false;
                    g.error("DtStart date must be before DtStop. DtStart has been reset to its previous value.");
                }
            }

            event.isCancelled = !changeIsValid;

            event.onComplete = function() {
                if (!event.isCancelled) { // if event not cancelled then invoke save method
                    RTEdits.MarketRateChgList.push({index: event.index, ID: g.records[event.index].RMRID});
                    g.url = '';
                    this.save();  // save automatically locally
                }
            };
        },
        // onDelete: function(event) {
        //     var selected = this.getSelection(),
        //         RMRIDList = [],
        //         grid = this;
        //
        //     // if not selected then return
        //     if (selected.length < 0) {
        //         return;
        //     }
        //     // collect RMRID
        //     selected.forEach(function(id) {
        //         RMRIDList.push(grid.get(id).RMRID);
        //     });
        //
        //     event.onComplete = function() {
        //         var x = getCurrentBusiness(),
        //             BID=parseInt(x.value),
        //             BUD = getBUDfromBID(BID),
        //             RTID = w2ui.rtForm.record.RTID;
        //
        //         var payload = { "cmd": "delete", "RMRIDList": RMRIDList };
        //         $.ajax({
        //             type: "POST",
        //             url: "/v1/rmr/" + BID + "/" + RTID,
        //             data: JSON.stringify(payload),
        //             contentType: "application/json",
        //             dataType: "json",
        //             success: function(data) {
        //                 grid.reload();
        //             },
        //         });
        //     };
        // },
    });

};
// saveRentableMarketRate - creates a list of MarketRate entries that have
// been changed, then calls the webservice to save them.
//---------------------------------------------------------------------------
window.saveRentableMarketRate = function(BID,RTID) {
    var list = [];
    var i;

    //------------------------------------------------
    // Build a list of IDs. that were edited...
    //------------------------------------------------
    for (i = 0; i < RTEdits.MarketRateChgList.length ; i++) {
        list[i] = RTEdits.MarketRateChgList[i].ID;
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
    var grid = w2ui.rmrGrid;
    var index = -1;
    for ( i = 0; i < reclist.length; i++) {
        //------------------------------------------------------------
        // reclist[i] is the id of the element we wnat to find...
        //------------------------------------------------------------
        for (var j = 0; j < grid.records.length; j++ ) {
            if (grid.records[j].RMRID == reclist[i]) {
                index = j;
                break;
            }
        }
        //------------------------------------------------------------
        // if the ID could not be found resolve promis with an
        // error message.
        //------------------------------------------------------------
        if (index < 0) {
            var s='ERROR: could not find RMRID = '+reclist[i];
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
        if (nrec.RMRID < 1 ) {
            nrec.RMRID = 0;  // server wants new records to have RMRID=0
        }
        chgrec.push(nrec);
    }

    var params = {
        cmd: "save",
        selected: [],
        limit: 0,
        offset: 0,
        changes: chgrec,
        RTID: RTID
    };

    var dat = JSON.stringify(params);
    var url = '/v1/rmr/' + BID + '/' + RTID;

    return $.post(url, dat, null, "json")
    .done(function(data) {
        if (data.status === "success") {
            //------------------------------------------------------------------
            // Now that the save is complete, we can add the URL back to the
            // the grid so it can call the server to get updated rows. The
            // onLoad handler will reset the url to '' after the load completes
            // so that changes are done locally to gthe grid until the
            // rtForm save button is clicked.
            //------------------------------------------------------------------
            RTEdits.MarketRateChgList = []; // reset the change list now, because we've saved them
            w2ui.rmrGrid.url = url;
            w2ui.toplayout.hide('right', true);
        } else {
            w2ui.rentablesGrid.error('saveRentableMarketRate: ' + data.message);
        }
    })
    .fail(function(data){
        w2ui.rentablesGrid.error("Save RentableMarketRate failed. " + data);
    });
};

// addRentableMarketRate - creates a new RentableMarketRate entry and adds it
// to the grid.
//
// @params
//
// @return
//---------------------------------------------------------------------------
window.addRentableMarketRate = function() {
    var x = getCurrentBusiness();
    var BID = parseInt(x.value);
    var BUD = getBUDfromBID(BID);
    var g = w2ui.rmrGrid;
    var ndStart;
    var basedOnListDate = false;

    // get lastest date among all market rate object's stopDate for new MR's StartDate
    if (g.records.length === 0) {
        ndStart = new Date();
    } else {
        g.records.forEach(function (rec) {
            if (ndStart === undefined) {
                ndStart = new Date(rec.DtStop);
            }
            if (rec.DtStop) {
                var rdStop = new Date(rec.DtStop);
                if (ndStart < rdStop) {
                    ndStart = rdStop;
                    basedOnListDate = true;
                }
            }
        });
    }

    var newRec = {
        recid: g.records.length,
        BID: BID,
        BUD: BUD,
        RTID: w2ui.rtForm.record.RTID,
        MarketRate: 0,
        RMRID: RTEdits.RMRID,
        DtStart: dateFmtStr(ndStart),
        DtStop: "12/31/9999"
    };
    --RTEdits.RMRID;
    if (EDIEnabledForBUD(BUD)) {
        if (basedOnListDate) {
            var d = ndStart;
            d.setDate(d.getDate()+1);
            newRec.DtStart = dateFmtStr(d);
        }
        newRec.DtStop = "12/30/9999";
    }
    var d1 = new Date(newRec.DtStart);
    var d2 = new Date(newRec.DtStop);
    if (d1 > d2) {
        newRec.DtStart = dateFmtStr(d1);
    }
    RTEdits.MarketRateChgList.push({index: 0, ID: newRec.RMRID});
    g.add(newRec,true);
};
