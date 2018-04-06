/*global
    setDefaultFormFieldAsPreviousRecord, w2uiDateControlString, $, w2ui, app, getCurrentBusiness, parseInt, getBUDfromBID,
    getRentableTypes, setToForm, form_dirty_alert, console, getFormSubmitData, addDateNavToToolbar, setRentableLayout,
    getRentableInitRecord
*/
"use strict";
window.getRentableInitRecord = function (BID, BUD, previousFormRecord) {
    var y = new Date();
    var defaultFormData = {
        recid: 0,
        BID: BID,
        BUD: BUD,
        RID: 0,
        RentableName: "",
        RARID: 0,
        RAID: 0,
        RARDtStart: w2uiDateControlString(y),
        RARDtStop: "1/1/9999",
        RTID: {id: 0, text: ''},
        RTRID: 0,
        RTRefDtStart: w2uiDateControlString(y),
        RTRefDtStop: "1/1/9999",
        RSID: 0,
        RentableStatus: "unknown",
        RSDtStart: w2uiDateControlString(y),
        RSDtStop: "1/1/9999",
        AssignmentTime: 0,
        Comment: ""
    };

    // if it called after 'save and add another' action there previous form record is passed as Object
    // else it is null
    if (previousFormRecord) {
        defaultFormData = setDefaultFormFieldAsPreviousRecord(
            ['RentableName'], // Fields to Reset
            defaultFormData,
            previousFormRecord
        );
    }

    return defaultFormData;
};

//-----------------------------------------------------------------------------
// getRentableTypes - return the RentableTypes list with respect of BUD
// @params
//      - BUD: current business designation
// @return  the Rentable Types List
//-----------------------------------------------------------------------------
window.getRentableTypes = function (BUD) {
    return jQuery.ajax({
        type: "GET",
        url: "/v1/rtlist/" + BUD,
        dataType: "json",
    }).done(function (data) {
        if (data.status == "success") {
            if (data.records) {
                app.rt_list[BUD] = data.records;
            } else {
                app.rt_list[BUD] = [];
            }
        }
    });
};

window.buildRentableElements = function () {
    // inside rentable part, we need this items
    app.cycleFreqItems = []; // cycle freq items
    app.cycleFreq.forEach(function (item, index) {
        app.cycleFreqItems.push({id: index, text: item});
    });

    app.RSUseStatusItems = []; // rentable use status items
    app.RSUseStatus.forEach(function (item, index) {
        app.RSUseStatusItems.push({id: index, text: item});
    });

    app.RSLeaseStatusItems = []; // rentable lease status items
    app.RSLeaseStatus.forEach(function (item, index) {
        app.RSLeaseStatusItems.push({id: index, text: item});
    });

    //------------------------------------------------------------------------
    //          rentablesGrid
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'rentablesGrid',
        url: '/v1/rentable',
        multiSelect: false,
        show: {
            header: false,
            toolbar: true,
            toolbarAdd: true,
            searches: true,
            footer: true,
            lineNumbers: false,
            selectColumn: false,
            expandColumn: false
        },
        columns: [
            {field: 'recid', caption: 'recid', size: '50px', hidden: true, sortable: true},
            {field: 'RID', caption: 'RID', size: '50px', sortable: true},
            {field: 'RentableName', caption: 'Rentable Name', size: '150px', sortable: true},
            {field: 'RTRID', caption: 'RTRID', hidden: true, sortable: true},
            {field: 'RTID', caption: 'Rentable Type ID', hidden: true, sortable: true},
            {field: 'RentableType', caption: 'Rentable Type', size: '200px', sortable: true},
            {field: 'RSID', caption: 'RSID', hidden: true, sortable: true},
            {
                field: 'UseStatus', caption: 'Rentable <br>Use Status', size: '100px', sortable: true,
                render: function (record/*, index, col_index*/) {
                    var text = '';
                    if (record) {
                        app.RSUseStatus.forEach(function (item, index) {
                            if (record.UseStatus == index) {
                                text = item;
                                return false;
                            }
                        });
                    }
                    return text;
                },
            },
            {
                field: 'LeaseStatus', caption: 'Rentable <br>Lease Status', size: '100px', sortable: true,
                render: function (record/*, index, col_index*/) {
                    var text = '';
                    if (record) {
                        app.RSLeaseStatus.forEach(function (item, index) {
                            if (record.LeaseStatus == index) {
                                text = item;
                                return false;
                            }
                        });
                    }
                    return text;
                },
            },
            {field: 'RARID', caption: 'RARID', hidden: true, sortable: true},
            {field: 'RAID', caption: 'RAID', size: '70px', sortable: true},
            {field: 'RentalAgreementStart', caption: 'Rental Agreement <br>Start', size: '120px', sortable: true},
            {field: 'RentalAgreementStop', caption: 'Rental Agreement <br>Stop', size: '120px', sortable: true},
        ],
        onRefresh: function (event) {
            event.onComplete = function () {
                var sel_recid = parseInt(this.last.sel_recid);
                if (app.active_grid == this.name && sel_recid > -1) {
                    if (app.new_form_rec) {
                        this.selectNone();
                    }
                    else {
                        this.select(app.last.grid_sel_recid);
                    }
                }
            };
        },
        onClick: function (event) {
            event.onComplete = function () {
                var yes_args = [this, event.recid],
                    no_args = [this],
                    no_callBack = function (grid) {
                        grid.select(app.last.grid_sel_recid);
                        return false;
                    },
                    yes_callBack = function (grid, recid) {
                        app.last.grid_sel_recid = parseInt(recid);

                        // keep highlighting current row in any case
                        grid.select(app.last.grid_sel_recid);

                        var rec = grid.get(recid),
                            x = getCurrentBusiness(),
                            BID = parseInt(x.value),
                            BUD = getBUDfromBID(BID);

                        getRentableTypes(BUD)
                            .done(function (data) {
                                if ('status' in data && data.status !== 'success') {
                                    w2ui.rentableForm.message(data.message);
                                } else {
                                    // get "RTID" column index and set rentable types list in editable items
                                    var RTIDColIndex = w2ui.rentableTypeRefGrid.getColumn("RTID", true);
                                    w2ui.rentableTypeRefGrid.columns[RTIDColIndex].editable.items = app.rt_list[BUD];
                                    setRentableLayout(BID, rec.RID);
                                }
                            })
                            .fail(function () {
                                console.log('Error getting /v1/uival/' + BID + '/app.ReceiptRules');
                            });
                    };

                // warn user if form content has been chagned
                form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
            };
        },
        onAdd: function (/*event*/) {
            var yes_args = [this],
                no_callBack = function () {
                    return false;
                },
                yes_callBack = function (grid) {
                    // reset it
                    app.last.grid_sel_recid = -1;
                    grid.selectNone();

                    var x = getCurrentBusiness(),
                        BID = parseInt(x.value),
                        BUD = getBUDfromBID(BID);

                    w2ui.rentableForm.record = getRentableInitRecord(BID, BUD, null);
                    w2ui.rentableForm.refresh();

                    getRentableTypes(BUD)
                        .done(function (data) {
                            if ('status' in data && data.status !== 'success') {
                                w2ui.rentableForm.message(data.message);
                            } else {
                                // get "RTID" column index and set rentable types list in editable items
                                var RTIDColIndex = w2ui.rentableTypeRefGrid.getColumn("RTID", true);
                                w2ui.rentableTypeRefGrid.columns[RTIDColIndex].editable.items = app.rt_list[BUD];
                                setRentableLayout(BID, 0);
                            }
                        })
                        .fail(function () {
                            console.log('Error getting /v1/uival/' + BID + '/app.ReceiptRules');
                        });
                };

            // warn user if form content has been chagned
            form_dirty_alert(yes_callBack, no_callBack, yes_args);
        },
    });
    addDateNavToToolbar('rentables');

    //------------------------------------------------------------------------
    //          rentable detailed layout with form in main panel
    //------------------------------------------------------------------------
    $().w2layout({
        name: 'rentableDetailLayout',
        panels: [
            {
                type: 'top',
                size: 35,
                style: 'border: 1px solid silver;',
                content: "",
                toolbar: {
                    style: "height: 35px; background-color: #eee; border: 0px;",
                    items: [
                        {id: 'btnNotes', type: 'button', icon: 'far fa-sticky-note'},
                        {id: 'bt3', type: 'spacer'},
                        {id: 'btnClose', type: 'button', icon: 'fas fa-times'},
                    ],
                    onClick: function (event) {
                        switch (event.target) {
                            case 'btnClose':
                                var no_callBack = function () {
                                        return false;
                                    },
                                    yes_callBack = function () {
                                        w2ui.toplayout.hide('right', true);
                                        w2ui.rentablesGrid.render();
                                    };
                                form_dirty_alert(yes_callBack, no_callBack);
                                break;
                        }
                    },
                },
            },
            {
                type: 'main',
                overflow: "hidden",
                style: 'background-color: white; border: 1px solid silver; padding: 0px;',
                tabs: {
                    style: "padding-top: 10px;",
                    active: 'rentableForm',
                    tabs: [
                        {id: 'rentableForm', caption: 'Rentable Detail'},
                        {id: 'rentableStatusGrid', caption: 'Rentable Status'},
                        {id: 'rentableTypeRefGrid', caption: 'Rentable Type Ref'},
                    ],
                    onClick: function (event) {
                        if (event.target === "rentableForm") {
                            w2ui.rentableDetailLayout.html('main', w2ui.rentableForm);
                        } else if (event.target === "rentableStatusGrid") {
                            w2ui.rentableDetailLayout.html('main', w2ui.rentableStatusGrid);
                        } else if (event.target === "rentableTypeRefGrid") {
                            w2ui.rentableDetailLayout.html('main', w2ui.rentableTypeRefGrid);
                        }
                    }
                }
            },
            {
                type: 'bottom',
                size: 60,
                // style: 'background-color: white;  border-top: 1px solid silver; text-align: center; padding: 15px;',
            },
        ],
    });

    //------------------------------------------------------------------------
    //          rentableForm
    //------------------------------------------------------------------------
    $().w2form({
        name: 'rentableForm',
        style: 'border: 0px; background-color: transparent;',
        url: '/v1/rentable',
        formURL: '/webclient/html/formr.html',
        fields: [
            {field: 'recid', type: 'int', required: false, html: {page: 0, column: 0}},
            {field: 'RID', type: 'int', required: false, html: {page: 0, column: 0}},
            {field: 'BID', type: 'int', required: true, html: {page: 0, column: 0}},
            {field: 'BUD', type: 'list', required: true, options: {items: app.businesses}, html: {page: 0, column: 0}},
            {field: 'RentableName', type: 'text', required: true, html: {page: 0, column: 0}},
            {field: 'AssignmentTime', type: 'list', required: false, html: {page: 0, column: 0}},
            {field: 'Comment', type: 'text', required: false, html: {page: 0, column: 0}},
            {field: 'LastModTime', type: 'hidden', required: false},
            {field: 'LastModBy', type: 'hidden', required: false},
            {field: 'CreateTS', type: 'hidden', required: false},
            {field: 'CreateBy', type: 'hidden', required: false}
        ],
        onSubmit: function (target, data) {
            // server request form data
            delete data.postData.record.LastModTime;
            delete data.postData.record.LastModBy;
            delete data.postData.record.CreateTS;
            delete data.postData.record.CreateBy;
            getFormSubmitData(data.postData.record);
        },
        onRefresh: function (event) {
            event.onComplete = function () {
                var f = this,
                    r = f.record,
                    x = getCurrentBusiness(),
                    BID = parseInt(x.value),
                    BUD = getBUDfromBID(BID),
                    header = "";

                /*// custom header, not common one!!
                if (r.RID) {
                    header = "Edit {0} - {1} ({2})".format(app.sRentable, r.RentableName, r.RID);
                } else {
                    header = "Edit {0} ({1})".format(app.sRentable, "new");
                }*/

                // assignmentTime selected and items for w2field
                var assignmentItems = [], assignSelected = {};
                app.assignmentTimeList.forEach(function (item, index) {
                    if (index == r.AssignmentTime) {
                        assignSelected = {id: index, text: item};
                    }
                    assignmentItems.push({id: index, text: item});
                });

                f.get("AssignmentTime").options.items = assignmentItems;
                f.get("AssignmentTime").options.selected = assignSelected;

                formRefreshCallBack(f, "RID", header);
            };
        },
        onChange: function (event) {
            event.onComplete = function () {
                // formRecDiffer: 1=current record, 2=original record, 3=diff object
                var diff = formRecDiffer(this.record, app.active_form_original, {});
                // if diff == {} then make dirty flag as false, else true
                if ($.isPlainObject(diff) && $.isEmptyObject(diff)) {
                    app.form_is_dirty = false;
                } else {
                    app.form_is_dirty = true;
                }
            };
        },
    });

    //------------------------------------------------------------------------
    //          rentable Status Grid
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'rentableStatusGrid',
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
            {field: 'DtNoticeToVacate', caption: 'DtNoticeToVacate', hidden: true},
            {field: 'DtNoticeToVacateIsSet', caption: 'DtNoticeToVacateIsSet', hidden: true},
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

            // get "Unknown" status value from the map, as well as for "Inactive" from Lease Status items
            var LeaseUnknownStatus, LeaseInactiveStatus;
            app.RSLeaseStatusItems.forEach(function (status) {
                switch (status.text) {
                    case "Unknown":
                        LeaseUnknownStatus = status.id;
                        break;
                    case "Inactive":
                        LeaseInactiveStatus = status.id;
                        break;
                }
            });

            this.records.forEach(function (item, index, arr) {
                arr[index].UseStatus = parseInt(arr[index].UseStatus);
                arr[index].LeaseStatus = parseInt(arr[index].LeaseStatus);

                if (arr[index].UseStatus === UseUnknownStatus && arr[index].LeaseStatus === LeaseUnknownStatus) {
                    // if UseStatus and LeaseStatus both kept as "unknown" then it doesn't
                    // make sense to send this entry to server, remove it
                    arr.splice(index, 1);
                } else if (arr[index].UseStatus === UseInactiveStatus || arr[index].LeaseStatus === LeaseInactiveStatus) {
                    // if "Inactive" set in any of UseStatus, LeaseStatus, then set "Inactive"
                    // in both status field
                    arr[index].UseStatus = UseInactiveStatus;
                    arr[index].LeaseStatus = LeaseInactiveStatus;
                }
            });
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
                    url: "/v1/rentablestatus/" + BID + "/" + RID,
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
            var g = this,
                field = g.columns[event.column].field,
                chgRec = g.get(event.recid),
                changeIsValid = true;

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
            var g = this,
                field = g.columns[event.column].field,
                chgRec = g.get(event.recid),
                changeIsValid = true;

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

    //------------------------------------------------------------------------
    //          Rentable Form Buttons
    //------------------------------------------------------------------------
    $().w2form({
        name: 'rentableFormBtns',
        style: 'border: 0px; background-color: transparent;',
        formURL: '/webclient/html/formrbtns.html',
        url: '',
        fields: [],
        actions: {
            save: function () {
                var BID = getCurrentBID();

                // unselect record from
                w2ui.rentablesGrid.selectNone();

                // hit save
                w2ui.rentableForm.save({}, function (data) {
                    if (data.status === 'error') {
                        console.log('ERROR: ' + data.message);
                        return;
                    }

                    // in case if record is new then we've to update RID that saved on server side
                    w2ui.rentableForm.record.RID = data.recid;

                    var i;
                    // update RID in grid records (status)
                    for (i = 0; i < w2ui.rentableStatusGrid.records.length; i++) {
                        w2ui.rentableStatusGrid.records[i].RID = w2ui.rentableForm.record.RID;
                    }

                    // update RID in grid records (typeRef)
                    for (i = 0; i < w2ui.rentableTypeRefGrid.records.length; i++) {
                        w2ui.rentableTypeRefGrid.records[i].RID = w2ui.rentableForm.record.RID;
                    }

                    // now set the url of status grid so that it can save the record on server side
                    w2ui.rentableStatusGrid.url = '/v1/rentablestatus/' + BID + '/' + w2ui.rentableForm.record.RID;
                    w2ui.rentableStatusGrid.save(function (data) {
                        // no matter, if it was succeed or not, just reset it, we already setting it before save call
                        w2ui.rentableStatusGrid.url = ""; // after save, remove it

                        if (data.status == "success") {
                            // next save rentable type ref
                            // now set the url of type ref grid so that it can save the record on server side
                            w2ui.rentableTypeRefGrid.url = '/v1/rentabletyperef/' + BID + '/' + w2ui.rentableForm.record.RID;
                            w2ui.rentableTypeRefGrid.save(function (data) {
                                // no matter, if it was succeed or not, just reset it, we already setting it before save call
                                w2ui.rentableTypeRefGrid.url = ""; // after save, remove it

                                if (data.status === "success") {
                                    w2ui.toplayout.hide('right', true);
                                    w2ui.rentablesGrid.render();
                                }
                            });
                        }
                    });
                });
            },
            saveadd: function () {
                var BID = getCurrentBID(),
                    BUD = getBUDfromBID(BID);

                // clean dirty flag of form
                app.form_is_dirty = false;

                // clear the grid select recid
                app.last.grid_sel_recid = -1;

                // select none if you're going to add new record
                w2ui.rentablesGrid.selectNone();

                w2ui.rentableForm.save({}, function (data) {
                    if (data.status == 'error') {
                        console.log('ERROR: ' + data.message);
                        return;
                    }

                    // now set the url of market Rate grid so that it can save the record on server side
                    w2ui.rentableStatusGrid.url = '/v1/rentablestatus/' + BID + '/' + w2ui.rentableForm.record.RID;
                    w2ui.rentableStatusGrid.save(function (data) {
                        // no matter, if it was succeed or not, just reset it, we already setting it before save call
                        w2ui.rentableStatusGrid.url = ""; // after save, remove it

                        if (data.status === "success") {

                            // clear grid as we're going to add new Form
                            w2ui.rentableStatusGrid.clear();

                            // next save rentable type ref
                            // now set the url of type ref grid so that it can save the record on server side
                            w2ui.rentableTypeRefGrid.url = '/v1/rentabletyperef/' + BID + '/' + w2ui.rentableForm.record.RID;
                            w2ui.rentableTypeRefGrid.save(function (data) {
                                // no matter, if it was succeed or not, just reset it, we already setting it before save call
                                w2ui.rentableTypeRefGrid.url = ""; // after save, remove it

                                if (data.status === "success") {

                                    // clear the rentabletyperef grid, as we're going to add new record
                                    w2ui.rentableTypeRefGrid.clear();

                                    // JUST RENDER THE MAIN GRID ONLY
                                    w2ui.rentablesGrid.render();

                                    w2ui.rentableForm.record = getRentableInitRecord(BID, BUD, w2ui.rentableForm.record);
                                    // w2ui.rentableForm.header = "Edit {0} ({1}) as of {2}".format(app.sRentable, "new", w2uiDateControlString(w2ui.rentableForm.record.CurrentDate));
                                    w2ui.rentableForm.url = '/v1/rentable/' + BID + '/0';
                                    w2ui.rentableForm.refresh();
                                }
                            });
                        }
                    });
                });
            },
            deactivate: function () {
            }, // TODO(Sudip): deactivate action
            reactivate: function () {
            }, // TODO(Sudip): reactivate action
        },
    });
};

window.setRentableLayout = function (BID, RID) {

    // set the url for rentableForm
    w2ui.rentableForm.url = '/v1/rentable/' + BID + '/' + RID;

    // load bottom panels with action buttons panel
    w2ui.rentableDetailLayout.content("bottom", w2ui.rentableFormBtns);

    // if form has tabs then click the first one
    if (typeof w2ui.rentableForm.tabs.name == "string") {
        w2ui.rentableForm.tabs.click('tab1');
    }

    // mark this flag as is this new record
    app.new_form_rec = RID ? true : false;

    // as new content will be loaded for this form
    // mark form dirty flag as false
    app.form_is_dirty = false;

    if (RID) {

        // if rentable available then load the status grid
        w2ui.rentableStatusGrid.load('/v1/rentablestatus/' + BID + '/' + RID);
        console.log("rentable status grid load data: {0}".format('/v1/rentablestatus/' + BID + '/' + RID));

        // // if rentable available then load the type ref grid
        w2ui.rentableTypeRefGrid.load('/v1/rentabletyperef/' + BID + '/' + RID);
        console.log("rentable type ref grid load data: {0}".format('/v1/rentabletyperef/' + BID + '/' + RID));

        // change the text of form tab
        w2ui.rentableDetailLayout.get("main").tabs.get("rentableForm").text = "Rentable Details ({0})".format(RID);
        w2ui.rentableDetailLayout.get("main").tabs.refresh();

        // load form content from server
        w2ui.rentableForm.request(function (event) {
            if (event.status === "success") {
                // only render the toplayout after server has sent down data
                // so that w2ui can bind values with field's html control,
                // otherwise it is unable to find html controls
                showForm();
                return true;
            }
            else {
                showForm();
                w2ui.rentableForm.message("Could not get form data from server...!!");
                return false;
            }
        });
    }
    else {
        // if new RentableType then clear the status, type ref grid first
        w2ui.rentableStatusGrid.clear();
        w2ui.rentableTypeRefGrid.clear();

        // change the text of form tab
        w2ui.rentableDetailLayout.get("main").tabs.get("rentableForm").text = "Rentable Details ({0})".format("new");
        w2ui.rentableDetailLayout.get("main").tabs.refresh();

        // unselect the previous selected row
        var sel_recid = parseInt(w2ui.rentablesGrid.last.sel_recid);
        if (sel_recid > -1) {
            // if new record is being added then unselect {{the selected record}} from the grid
            w2ui.rentablesGrid.unselect(w2ui.rentablesGrid.last.sel_recid);
        }

        showForm();
        return true;
    }

    function showForm() {
        // SHOW the right panel now
        w2ui.toplayout.content('right', w2ui.rentableDetailLayout);
        w2ui.toplayout.sizeTo('right', 700);
        // w2ui.rentableDetailLayout.render();
        w2ui.rentableDetailLayout.get("main").tabs.click("rentableForm");
        w2ui.toplayout.show('right', true);
    }
};
