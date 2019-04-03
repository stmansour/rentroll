/*global
    setDefaultFormFieldAsPreviousRecord, w2uiDateControlString, $, w2ui, app, getCurrentBusiness, parseInt, getBUDfromBID,
    getRentableTypes, setToForm, form_dirty_alert, console, getFormSubmitData, addDateNavToToolbar, setRentableLayout,
    getRentableInitRecord, saveRentableLeaseStatus, RentableEdits, addRentableTypeRef,
    w2uiDateTimeControlString,
*/
/*jshint esversion: 6 */

"use strict";

window.buildRentableTypeRefElements = function () {
    //------------------------------------------------------------------------
    //          rentable Type Ref Grid
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'rentableTypeRefGrid',
        style: 'padding: 0px',
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
            {field: 'RTRID', caption: 'RTRID', size: '50px'},
            {
                field: 'RTID', caption: 'Rentable Type', size: '150px',
                editable: {type: 'select', align: 'left', items: []},
                render: function (record, index, col_index) {
                    var html = '';
                    var BID = getCurrentBID(),
                        BUD = getBUDfromBID(BID);
                    for (var rt in app.rt_list[BUD]) {
                        if (app.rt_list[BUD][rt].id == this.getCellValue(index, col_index)) {
                            html = app.rt_list[BUD][rt].text;
                        }
                    }
                    return html;
                },
            },
            {
                field: 'OverrideRentCycle', caption: 'OverrideRentCycle', size: "150px",
                editable: {type: 'select', align: 'left', items: app.cycleFreqItems},
                render: function (record, index, col_index) {
                    var html = '';
                    for (var f in app.cycleFreqItems) {
                        if (app.cycleFreqItems[f].id == this.getCellValue(index, col_index)) {
                            html = app.cycleFreqItems[f].text;
                        }
                    }
                    return html;
                },
            },
            {
                field: 'OverrideProrationCycle', caption: 'OverrideProrationCycle', size: "150px",
                editable: {type: 'select', align: 'left', items: app.cycleFreqItems},
                render: function (record, index, col_index) {
                    var html = '';
                    for (var f in app.cycleFreqItems) {
                        if (app.cycleFreqItems[f].id == this.getCellValue(index, col_index)) {
                            html = app.cycleFreqItems[f].text;
                        }
                    }
                    return html;
                },
            },
            {field: 'DtStart',   caption: 'DtStart', size: "50%", sortable: true, style: 'text-align: right', editable: {type: 'date'} },
            {field: 'DtStop',    caption: 'DtStop', size: "50%", sortable: true, style: 'text-align: right', editable: {type: 'date'} },
            {field: 'CreateBy',  caption: 'CreateBy', hidden: true},
            {field: 'LastModBy', caption: 'LastModBy', hidden: true},
        ],
        onAdd: function (event) {
            addRentableTypeRef();
        },
        onLoad: function (event) {
            event.onComplete = function () {
                //------------------------------------------------------------------------
                // We need the URL to be active in case virtual scrolling is needed.
                // We will turn off the url when a local save is completed.
                //------------------------------------------------------------------------
                var BID = getCurrentBID();
                var RID = w2ui.rentableForm.record.RID;
                this.url = '/v1/rentabletyperef/'+BID+'/'+RID;
                RentableEdits.RTRDataLoaded = true;
            };
        },
        onSave: function (event) {
            // TODO:
            // // sman:  3/12/2019 -- I'm not sure if this is correct here it may
            // // need to go in the save button handler.
            // //-----------------------------------------------------------------
            // this.records.forEach(function (item, index, arr) {
            //     arr[index].OverrideRentCycle = parseInt(arr[index].OverrideRentCycle);
            //     arr[index].OverrideProrationCycle = parseInt(arr[index].OverrideProrationCycle);
            //     arr[index].RTID = parseInt(arr[index].RTID);
            // });

            //------------------------------------------------------------------
            // Grid changes are saved locally. So we want no url when this funct
            // gets called due to a grid change. After the local save is complete
            // we put the url back so that if the virtual scrolling needs to call
            // the server it can.  The full save to disk is done when the user
            // presses the Save button, which is handled by a different function.
            //------------------------------------------------------------------
            this.url = '';  // we're only doing a local save in the grid
            event.onComplete = function() {  // see the onLoad comment
                var BID = getCurrentBID();
                var RID = w2ui.rentableForm.record.RID;
                this.url = '/v1/rentabletyperef/'+BID+'/'+RID;
            };
        },
        onChange: function (event) {
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

            //---------------------------------------------------
            // Inform w2ui if the change is cancelled or not...
            //---------------------------------------------------
            event.isCancelled = !changeIsValid;

            event.onComplete = function () {
                //---------------------------------------------------------------
                // 2/19/2019 sman - This save is used to save the data into the
                // grid's records.  We need to ensure that the grids URL is ''
                //---------------------------------------------------------------
                if (!event.isCancelled) { // if event not cancelled then invoke save method
                    RentableEdits.RTRChgList.push({index: event.index, ID: g.records[event.index].RTRID});
                    g.url = '';  // just ensure that no server service is called
                    this.save(); // save automatically locally
                }
            };
        },
    });
};

// saveRentableTypeRef - creates a list of RentableTypeRef entries that have
// been changed, then calls the webservice to save them.
//
// @params
//     BID = business id
//     BUD = business designator
//
// @return
//     a Promise object:
//           if there are no changes to the Rentable's TypeRefs the return a resolved Promise
//           if we need to call the server, return the $.post() Promise
//---------------------------------------------------------------------------
window.saveRentableTypeRef = function(BID,RID) {
    var list = [];
    var i;

    //------------------------------------------------
    // Build a list of IDs. that were edited...
    //------------------------------------------------
    for (i = 0; i < RentableEdits.RTRChgList.length; i++) {
        list[i] = RentableEdits.RTRChgList[i].ID;
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
    // Find the records for each ID
    //------------------------------------------------
    var chgrec = [];
    var index = -1;
    var grid = w2ui.rentableTypeRefGrid;
    for (i = 0; i < reclist.length; i++) {
        //------------------------------------------------------------
        // reclist[i] is the id of the element we wnat to find...
        //------------------------------------------------------------
        for (var j = 0; j < grid.records.length; j++ ) {
            if (grid.records[j].RTRID == reclist[i]) {
                index = j;
                break;
            }
        }
        //------------------------------------------------------------
        // if the ID could not be found resolve promis with an
        // error message.
        //------------------------------------------------------------
        if (index < 0) {
            var s='ERROR: could not find RTRID = '+reclist[i];
            w2ui.rentablesGrid.error(s);  // place an error where we will be sure to see it
            return Promise.resolve('{"status": "error", "message": s}');
        }

        //------------------------------------------------------------
        //  This is the record we need to save.  Make any last-min
        //  changes...
        //------------------------------------------------------------
        var nrec = grid.records[index];
        if (typeof nrec.RTID == "string") {
            var ls = parseInt(nrec.RTID,10);
            nrec.RTID = ls;
        }
        if (nrec.RTRID < 0) {
            nrec.RTRID = 0;  // server needs RTRID = 0 for new records
        }
        chgrec.push(nrec);
    }

    //------------------------------------------------------------
    //  Save the list of chgrecs to the server...
    //------------------------------------------------------------
    var params = {
        cmd: "save",
        selected: [],
        limit: 0,
        offset: 0,
        changes: chgrec,
        RID: w2ui.rentableForm.record.RID
    };

    var dat = JSON.stringify(params);
    var url = '/v1/rentabletyperef/' + BID + '/' + w2ui.rentableForm.record.RID;

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
            RentableEdits.RTRChgList = []; // reset the change list now, because we've saved them
            w2ui.rentableTypeRefGrid.url = url;
            // w2ui.rentablesGrid.render();  // this maay need changing
        } else {
            w2ui.rentablesGrid.error("save failed: " + data);
        }
    })
    .fail(function(data){
        w2ui.rentablesGrid.error("Save RentableLeaseStatus failed. " + data);
    });
};

// addRentableTypeRef - creates a new RentableTypeRef entry and adds it
// to the grid.
//
// @params
//
// @return
//---------------------------------------------------------------------------
window.addRentableTypeRef = function() {
    var x = getCurrentBusiness();
    var BID = parseInt(x.value);
    var BUD = getBUDfromBID(BID);
    var g = w2ui.rentableTypeRefGrid;
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
        RID: w2ui.rentableForm.record.RID,
        RTID: app.rt_list[BUD][0].id,  // initialize to the first rentable type available
        RTRID: RentableEdits.RTRID,
        OverrideRentCycle: 0,
        OverrideProrationCycle: 0,
        DtStart: dateFmtStr(ndStart),
        DtStop: "12/31/9999"
    };
    --RentableEdits.RTRID;
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
    RentableEdits.RTRChgList.push({index: 0, ID: newRec.RTRID});
    g.add(newRec,true);
};
