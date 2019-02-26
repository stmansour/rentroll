/*global
    setDefaultFormFieldAsPreviousRecord, w2uiDateControlString, $, w2ui, app, getCurrentBusiness, parseInt, getBUDfromBID,
    getRentableTypes, setToForm, form_dirty_alert, console, getFormSubmitData, addDateNavToToolbar, setRentableLayout,
    getRentableInitRecord, saveRentableLeaseStatus, RentableEdits
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
        onLoad: function (event) {
            event.onComplete = function () {
                this.url = '';
            };
        },
        onAdd: function (event) {
            var x = getCurrentBusiness(),
                BID = parseInt(x.value),
                BUD = getBUDfromBID(BID),
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
                RID: w2ui.rentableForm.record.RID,
                RTID: 0,
                RTRID: 0,
                OverrideRentCycle: 0,
                OverrideProrationCycle: 0,
                DtStart: dateFmtStr(ndStart),
                DtStop: "12/31/9999"
            };
            RentableEdits.RTRChgList.push(newRec.recid);
            g.add(newRec);
        },
        onSave: function (event) {
            // TODO(Sudip): validation on values before sending these to server

            this.records.forEach(function (item, index, arr) {
                arr[index].OverrideRentCycle = parseInt(arr[index].OverrideRentCycle);
                arr[index].OverrideProrationCycle = parseInt(arr[index].OverrideProrationCycle);
                arr[index].RTID = parseInt(arr[index].RTID);
            });
            event.changes = this.records;
        },
        onDelete: function (event) {
            var selected = this.getSelection(),
                RTRIDList = [],
                grid = this;

            // if not selected then return
            if (selected.length < 0) {
                return;
            }
            // collect RTRID
            selected.forEach(function (id) {
                RTRIDList.push(grid.get(id).RTRID);
            });

            event.onComplete = function () {
                var x = getCurrentBusiness(),
                    BID = parseInt(x.value),
                    BUD = getBUDfromBID(BID),
                    RID = w2ui.rentableForm.record.RID;

                var payload = {"cmd": "delete", "RTRIDList": RTRIDList};
                $.ajax({
                    type: "POST",
                    url: "/v1/rentabletyperef/" + BID + "/" + RID,
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

            RentableEdits.RTRChgList.push(chgRec.recid);

            // if fields are DtStart or DtStop
            if (field === "DtStart" || field === "DtStop") {

                var chgDStart = field === "DtStart" ? new Date(event.value_new) : new Date(chgRec.DtStart),
                    chgDStop = field === "DtStop" ? new Date(event.value_new) : new Date(chgRec.DtStop);

                // Stop date should not before Start Date
                if (chgDStop <= chgDStart) {
                    changeIsValid = false;
                } else {
                    // make sure date values don't overlap with other market rate dates
                    for (var i in g.records) {
                        var rec = g.records[i];
                        if (rec.recid === chgRec.recid) { // if same record then continue to next one
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
            }

            if (changeIsValid) {
                event.isCancelled = false; // if everything is ok, then don't cancel
            } else {
                event.isCancelled = true;
            }

            event.onComplete = function () {
                if (!event.isCancelled) { // if event not cancelled then invoke save method
                    this.save();          // save automatically locally
                }
            };
        }
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
    var reclist = Array.from(new Set(RentableEdits.RTRChgList));

    if (reclist.length == 0) {
        return Promise.resolve('{"status": "success"}');
    }

    var chgrec = [];
    for (var i = 0; i < reclist.length; i++) {
        var nrec =  w2ui.rentableTypeRefGrid.get(reclist[i]);
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
    var url = '/v1/rentabletyperef/' + BID + '/' + w2ui.rentableForm.record.RID;
    return $.post(url, dat, null, "json")
    .done(function(data) {
        if (data.status === "success") {
            RentableEdits.RTRChgList = []; // reset the change list now, because we've saved them
            // w2ui.toplayout.hide('right', true);
            // w2ui.rentablesGrid.render();
            w2ui.rentableTypeRefGrid.url = '/v1/rentabletyperef/' + BID + '/' + w2ui.rentableForm.record.RID;
            w2ui.rentableTypeRefGrid.reload();
            w2ui.rentablesGrid.render();  // this maay need changing
        }
    })
    .fail(function(data){
        console.log("Save RentableLeaseStatus failed.");
    });
};
