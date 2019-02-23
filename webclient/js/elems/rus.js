/*global
    setDefaultFormFieldAsPreviousRecord, w2uiDateControlString, $, w2ui, app, getCurrentBusiness, parseInt, getBUDfromBID,
    getRentableTypes, setToForm, form_dirty_alert, console, getFormSubmitData, addDateNavToToolbar, setRentableLayout,
    getRentableInitRecord, saveRentableLeaseStatus, RentableEdits
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
            event.onComplete = function () {
                this.url = '';
            };
        },
        onAdd: function (/*event*/) {
            var x = getCurrentBusiness(),
                BID = parseInt(x.value),
                BUD = getBUDfromBID(BID),
                fr = w2ui.rentableForm.record,
                g = this,
                ndStart;

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
                RSID: 0,
                UseStatus: 0,
                LeaseStatus: 0,
                DtStart: dateFmtStr(ndStart),
                DtStop: "12/31/9999"
            };
            g.add(newRec);
            RentableEdits.UseStatusChgList.push(newRec.recid);
        },
        onSave: function (event) {
            // if url is set then only take further actions, for local save just ignore those
            if (this.url === "") {
                return false;
            }

            // TODO(Sudip): validation on values before sending these to server

            // get "Unknown" status value from the map, as well as for "Inactive" from Use Status items
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

            // // get "Unknown" status value from the map, as well as for "Inactive" from Use Status items
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

            // this.records.forEach(function (item, index, arr) {
            //     arr[index].UseStatus = parseInt(arr[index].UseStatus);
            //     // arr[index].LeaseStatus = parseInt(arr[index].LeaseStatus);
            //
            //     if (arr[index].UseStatus === UseUnknownStatus && arr[index].LeaseStatus === LeaseUnknownStatus) {
            //         // if UseStatus and LeaseStatus both kept as "unknown" then it doesn't
            //         // make sense to send this entry to server, remove it
            //         arr.splice(index, 1);
            //     } else if (arr[index].UseStatus === UseInactiveStatus || arr[index].LeaseStatus === LeaseInactiveStatus) {
            //         // if "Inactive" set in any of UseStatus, LeaseStatus, then set "Inactive"
            //         // in both status field
            //         arr[index].UseStatus = UseInactiveStatus;
            //         // arr[index].LeaseStatus = LeaseInactiveStatus;
            //     }
            // });
            event.changes = this.records;
        },
        onDelete: function (event) {
            var selected = this.getSelection(),
                RSIDList = [],
                grid = this;

            // if not selected then return
            if (selected.length < 0) {
                return;
            }
            // collect RMRID
            selected.forEach(function (id) {
                RSIDList.push(grid.get(id).RSID);
            });

            event.onComplete = function () {
                var x = getCurrentBusiness(),
                    BID = parseInt(x.value),
                    BUD = getBUDfromBID(BID),
                    RID = w2ui.rentableForm.record.RID;

                var payload = {"cmd": "delete", "RSIDList": RSIDList};
                $.ajax({
                    type: "POST",
                    url: "/v1/rentableusestatus/" + BID + "/" + RID,
                    data: JSON.stringify(payload),
                    contentType: "application/json",
                    dataType: "json",
                    success: function (data) {
                        grid.reload();
                    },
                });
            };
        },
        onChange: function (event) {
            event.preventDefault();
            var g = this;
            var field = g.columns[event.column].field;
            var chgRec = g.get(event.recid);
            var changeIsValid = true;

            RentableEdits.UseStatusChgList.push(chgRec.recid);

            switch (field) {
                case "UseStatus":
                    // in local save check if use status is unknown for existing instance
                    // if yes, then don't allow that change
                    app.RSUseStatusItems.forEach(function (status) {
                        switch (status.text) {
                            case "Unknown":
                                if (chgRec.RSID > 0) { // only for existing instance
                                    changeIsValid = false;
                                }
                        }
                    });
                    break;
                case "LeaseStatus":
                    // in local save check if lease status is unknown for existing instance
                    // if yes, then don't allow that change
                    app.RSLeaseStatusItems.forEach(function (status) {
                        switch (status.text) {
                            case "Unknown":
                                if (chgRec.RSID > 0) { // only for existing instance
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
                        // make sure date values don't overlap with other market rate dates
                        for (var i = 0; i< g.records.length; i++) {
                            var rec = g.records[i];
                            if (rec.recid >= chgRec.recid) { // if same record then continue to next one
                                continue;
                            }

                            var rDStart = new Date(rec.DtStart),
                                rDStop = new Date(rec.DtStop);

                            // return if changed record startDate falls in other MR time span
                            if (rDStart < chgDStart && chgDStart < rDStop) {
                                changeIsValid = false;
                            } else if (rDStart < chgDStop && chgDStop < rDStop) {
                                changeIsValid = false;
                            } else if (chgDStart < rDStart && rDStop < chgDStop) {
                                changeIsValid = false;
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

            event.onComplete = function () {
                if (!event.isCancelled) { // if event not cancelled then invoke save method
                    // save automatically locally
                    this.save();
                }
            };
        }
    });

};

// saveRentableUseStatus - creates a list of UseStatus entries that have
// been changed, then calls the webservice to save them.
//---------------------------------------------------------------------------
window.saveRentableUseStatus = function(BID,RID) {
    var reclist = Array.from(new Set(RentableEdits.UseStatusChgList));

    if (reclist.length == 0) {
        return Promise.resolve('{"status": "success"}');
    }

    var chgrec = [];
    for (var i = 0; i < reclist.length; i++) {
        var nrec =  w2ui.rentableUseStatusGrid.get(reclist[i]);
        if (typeof nrec.UseStatus == "string") {
            var ls = parseInt(nrec.UseStatus,10);
            nrec.UseStatus = ls;
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
    var url = '/v1/rentableusestatus/' + BID + '/' + w2ui.rentableForm.record.RID;
    return $.post(url, dat, null, "json")
    .done(function(data) {
        if (data.status === "success") {
            RentableEdits.UseStatusChgList = []; // reset the change list now, because we've saved them
            w2ui.toplayout.hide('right', true);
            w2ui.rentablesGrid.render();
        }
    })
    .fail(function(data){
        console.log("Save RentableUseStatus failed.");
    });
};
