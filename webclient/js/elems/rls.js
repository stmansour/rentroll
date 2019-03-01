/*global
    setDefaultFormFieldAsPreviousRecord, w2uiDateControlString, $, w2ui, app, getCurrentBusiness, parseInt, getBUDfromBID,
    getRentableTypes, setToForm, form_dirty_alert, console, getFormSubmitData, addDateNavToToolbar, setRentableLayout,
    getRentableInitRecord, saveRentableLeaseStatus, RentableEdits,
*/
/*jshint esversion: 6 */

"use strict";

window.buildRentableLeaseStatusElements = function () {
    //------------------------------------------------------------------------
    //          rentable Lease Status Grid
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'rentableLeaseStatusGrid',
        style: 'padding: 0px',
        url: '/v1/rentableleasestatus',
        show: {
            header: false,
            toolbar: true,
            toolbarReload: false,
            toolbarColumns: false,
            toolbarSearch: true,
            toolbarAdd: true,
            toolbarDelete: true,
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
            {field: 'RLID', caption: 'RLID', size: '50px'},
            {
                field: 'LeaseStatus', caption: 'LeaseStatus', size: '150px',
                editable: {type: 'select', align: 'left', items: app.RSLeaseStatusItems},
                render: function (record, index, col_index) {
                    var html = '';
                    for (var s in app.RSLeaseStatusItems) {
                        if (app.RSLeaseStatusItems[s].id == this.getCellValue(index, col_index)) {
                            html = app.RSLeaseStatusItems[s].text;
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
                editable: {type: 'date'}
            },
            {
                field: 'DtStop',
                caption: 'DtStop',
                size: "50%",
                sortable: true,
                style: 'text-align: right',
                editable: {type: 'date'}
            },
            {field: 'CreateBy', caption: 'CreateBy', hidden: true},
            {field: 'LastModBy', caption: 'LastModBy', hidden: true},
        ],
        onLoad: function (event) {
            var BID = getCurrentBID();
            var RID = w2ui.rentableForm.record.RID;
            event.onComplete = function () {
                this.url = '/v1/rentableleasestatus/'+BID+'/'+RID;
            };
        },
        onAdd: function (/*event*/) {
            var BID = getCurrentBID();
            var BUD = getBUDfromBID(BID);
            var fr = w2ui.rentableForm.record;
            var g = this;
            var ndStart;

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
                        }
                    }
                });
            }

            var newRec = {
                recid: g.records.length,
                BID: BID,
                BUD: BUD,
                RID: fr.RID,
                RLID: 0,
                LeaseStatus: 0,
                DtStart: dateFmtStr(ndStart),
                DtStop: "12/31/9999"
            };
            if (EDIEnabledForBUD(BUD)) {
                var d = ndStart;
                d.setDate(d.getDate()+1);
                newRec.DtStart = dateFmtStr(d);
                newRec.DtStop = "12/30/9999";
            }
            RentableEdits.LeaseStatusChgList.push(newRec.recid);
            g.add(newRec,true); // the boolean forces the new row to be added at the top of the grid
        },
        onSave: function (event) {
            // if url is set then only take further actions, for local save just ignore those
            this.url = '';
            event.onComplete = function() {
                var BID = getCurrentBID();
                var RID = w2ui.rentableForm.record.RID;
                w2ui.rentableLeaseStatusGrid.url = "/v1/rentableleasestatus/" + BID + "/" + RID;
            };
        },
        onDelete: function (event) {
            if (RentableEdits.rlsDeleteInProgress) {
                return;
            }

            var selected = this.getSelection();
            var RLIDList = [];
            var grid = this;
            RentableEdits.rlsDeleteInProgress = true;

            //-------------------------------------
            // if nothing is selected then return
            //-------------------------------------
            if (selected.length < 0) {
                return;
            }

            //-------------------------------------
            // collect the selected ids...
            //-------------------------------------
            for (var id = 0; id < selected.length; id++) {
                var RLID = grid.get(id).RLID;
                if (RLID > 0) {
                    RLIDList.push(RLID);
                }
            }

            event.onComplete = function () {
                var grid = this;
                var BID = getCurrentBID();
                var RID = w2ui.rentableForm.record.RID;

                var tgrid = w2ui.rentableLeaseStatusGrid;
                var url = "/v1/rentableusestatus/" + BID + "/" + RID;
                var params = {"cmd": "delete", "RLIDList": RLIDList};
                var dat = JSON.stringify(params);

                //-------------------------------------------------------------
                // If there are any other pending changes, we'll need to save
                // those first, then we proceed with the delete.
                // The save function implements the Promise interface so
                // we just need to supply the pass/fail functions...
                //-------------------------------------------------------------
                $.when(
                    saveRentableLeaseStatus(BID,RID)
                )
                .done(function(){
                    $.post(grid.url, dat, null, "json")
                    .done(function(data) {
                        if (data.status === "error") {
                            grid.error('onDelete: '+w2utils.lang(data.message));
                            return;
                        }
                        grid.reload();
                    })
                    .fail(function(/*data*/){
                        grid.error("Delete RentableLeaseStatus failed.");
                        return;
                    });
                })
                .fail(function(){
                    console.log('RentableLeaseSave: when failed.');
                });

                RentableEdits.rlsDeleteInProgress = false;
                // $.ajax({
                //     type: "POST",
                //     url: "/v1/rentableusestatus/" + BID + "/" + RID,
                //     data: JSON.stringify(payload),
                //     contentType: "application/json",
                //     dataType: "json",
                //     success: function (data) {
                //         grid.reload();
                //     },
                // });
            };
        },

        onChange: function (event) {
            // event.preventDefault();   // not sure what this does
            var g = this;
            var field = g.columns[event.column].field;
            var chgRec = g.get(event.recid);
            var changeIsValid = true;

            RentableEdits.LeaseStatusChgList.push(chgRec.recid);

            switch (field) {
                case "LeaseStatus":
                    // in local save check if lease status is unknown for existing instance
                    // if yes, then don't allow that change
                    app.RSLeaseStatusItems.forEach(function (status) {
                        switch (status.text) {
                            case "Unknown":
                                if (chgRec.RLID > 0) { // only for existing instance
                                    changeIsValid = false;
                                }
                        }
                    });
                    break;
                case "DtStart":
                case "DtStop":
                    // get the changed value if field, otherwise take the record saved date value
                    var chgDStart = field === "DtStart" ? new Date(event.value_new) : new Date(chgRec.DtStart),
                        chgDStop = field === "DtStop" ? new Date(event.value_new) : new Date(chgRec.DtStop);

                    // Stop date should not before Start Date
                    if (chgDStop <= chgDStart) {
                        changeIsValid = false;
                    } else {
                        // TODO: This date verification code needs to be rewritten!!!
                        // make sure date values don't overlap with other market rate dates
                        // for (var i = 0; i< g.records.length; i++) {
                        //     var rec = g.records[i];
                        //     if (rec.recid === chgRec.recid) { // if same record then continue to next one
                        //         i = g.records.length;
                        //         continue;
                        //     }
                        //
                        //     var rDStart = new Date(rec.DtStart),
                        //         rDStop = new Date(rec.DtStop);
                        //
                        //     // return if changed record startDate falls in other MR time span
                        //     if (rDStart < chgDStart && chgDStart < rDStop) {
                        //         changeIsValid = false;
                        //     } else if (rDStart < chgDStop && chgDStop < rDStop) {
                        //         changeIsValid = false;
                        //     } else if (chgDStart < rDStart && rDStop < chgDStop) {
                        //         changeIsValid = false;
                        //     }
                        //
                        // }
                        if (changeIsValid) {
                            // for some reason, there are cases where changeIsValid is true yet the value is not changed.
                            // this is a temporary fix to
                            if (field == "DtStart") {
                                w2ui.rentableLeaseStatusGrid.records[event.recid].DtStart = event.value_new;
                            } else {
                                w2ui.rentableLeaseStatusGrid.records[event.recid].DtStop = event.value_new;
                            }
                        }
                    }
                    break;
            }

            if (changeIsValid) {
                // if everything is ok, then mark this as false
                event.isCancelled = false;
            } else {
                event.isCancelled = true;
            }

            // 2/19/2019 sman - This save is used to save the data into the
            // grid's records.  We need to ensure that the grids URL is ''
            //-------------------------------------------------------------------
            event.onComplete = function () {
                if (!event.isCancelled) { // if event not cancelled then invoke save method
                    g.url = '';  // just ensure that no server service is called
                    this.save(); // save automatically locally
                    var BID = getCurrentBID();
                    var RID = w2ui.rentableForm.record.RID;
                    g.url = '/v1/rentableleasestatus/' + BID + '/' + RID;
                }
            };
        }
    });
};

// saveRentableLeaseStatus - creates a list of LeaseStatus entries that have
// been changed, then calls the webservice to save them.
//---------------------------------------------------------------------------
window.saveRentableLeaseStatus = function(BID,RID) {
    var reclist = Array.from(new Set(RentableEdits.LeaseStatusChgList));

    if (reclist.length == 0) {
        return Promise.resolve('{"status": "success"}');
    }

    var chgrec = [];
    for (var i = 0; i < reclist.length; i++) {
        var nrec =  w2ui.rentableLeaseStatusGrid.get(reclist[i]);
        if (typeof nrec.LeaseStatus == "string") {
            var ls = parseInt(nrec.LeaseStatus,10);
            nrec.LeaseStatus = ls;
        }
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
    var url = '/v1/rentableleasestatus/' + BID + '/' + w2ui.rentableForm.record.RID;

    return $.post(url, dat, null, "json")
    .done(function(data) {
        if (data.status === "success") {
            RentableEdits.LeaseStatusChgList = []; // reset the change list now, because we've saved them
            w2ui.toplayout.hide('right', true);
            w2ui.rentablesGrid.render();
        } else {
            w2ui.rentablesGrid.error('saveRentableLeaseStatus: ' + data.message);
        }
    })
    .fail(function(data){
        console.log("Save RentableLeaseStatus failed.");
    });
};
