/*global
    setDefaultFormFieldAsPreviousRecord, w2uiDateControlString, $, w2ui, app, getCurrentBusiness, parseInt, getBUDfromBID,
    getRentableTypes, setToForm, form_dirty_alert, console, getFormSubmitData, addDateNavToToolbar, setRentableLayout,
    getRentableInitRecord, saveRentableLeaseStatus, buildRentableUseStatusElements, buildRentableLeaseStatusElements,
    saveRentableUseStatus, saveRentableTypeRef, buildRentableTypeRefElements, saveRentableCore, closeRentableForm,
    showRentableForm, finishRentableSaveAdd, finishRentableSave, addRentableLeaseStatus, addRentableUseStatus, addRentableTypeRef,
    getEDIAdjustedStopDate, buildRentableUseTypeElements, saveRentableUseType, addRentableUseType,
*/
/*jshint esversion: 6 */

"use strict";
var RentableEdits = {
    LeaseStatusChgList: [],         // an array of indeces to LeaseStatus changes
    UseStatusChgList: [],           // an array of indeces to UseStatus changes
    UseTypeChgList: [],             // an array of indeces to UseType changes
    RTRChgList: [],                 // an array of indeces to type ref changes
    RID: 0,                         // ID being edited
    RSID: -1,                       // ID to use for new objects
    RLID: -1,                       // ID to use for new objects
    RTRID: -1,                      // ID to use for new objects
    RMRID: -1,                      // ID to use for new objects
    LeaseStatusDataLoaded: false,   // set to true after server data is loaded
    UseStatusDataLoaded: false,     // set to true after server data is loaded
    UseTypeDataLoaded: false,       // set to true after server data is loaded
    RTRDataLoaded: false,           // set to true after server data is loaded
};

//-----------------------------------------------------------------------------
// getEDIAdjustedStopDate - This routine should only be used for date ranges
//     that are created on the webclient.  Note that date date ranges from the
//     server are already adjusted, so there is no need to call this routine
//     for data received from the server.
//
// @params
//      BUD - current business designation
//      DtStartStr - date range start time in string form
//      DtStopStr - date range stop time in string form
//
// @return  the adjusted end date as a string
//-----------------------------------------------------------------------------
window.getEDIAdjustedStopDate = function(BUD,DtStartStr,DtStopStr) {
    var d1 = new Date(DtStartStr);
    var d2 = new Date(DtStopStr);

    if (EDIEnabledForBUD(BUD)) {
        var d = d2;
        d.setDate(d.getDate()-1);
        d2 = d;
    }
    if (d1 > d2) {
        d2 = d1;
    }
    return w2uiDateControlString(d2);
};

//-----------------------------------------------------------------------------
// getRentableInitRecord - This routine returns a data record for a new
//      rentable.
//
// @params
//      BID - current business ID
//      BUD - current business designation
//      previousFormRecord - info from previously filled out form
//
// @return  the new data record
//-----------------------------------------------------------------------------
window.getRentableInitRecord = function (BID, BUD, previousFormRecord) {
    var y = new Date();
    var rec = {
        recid: 0,
        BID: BID,
        BUD: BUD,
        RID: 0,
        RentableName: "",
        RARID: 0,
        RAID: 0,
        RARDtStart: w2uiDateControlString(y),
        RARDtStop: "12/31/9999",
        RTID: {id: 0, text: ''},
        RTRID: 0,
        RTRefDtStart: w2uiDateControlString(y),
        RTRefDtStop: "12/31/9999",
        RSID: 0,
        RentableStatus: "Ready",
        RRentableLeaseStatus: "Not rented",
        RSDtStart: w2uiDateControlString(y),
        RSDtStop: "12/31/9999",
        AssignmentTime: 1,
        Comment: ""
    };

    rec.RARDtStop = getEDIAdjustedStopDate(BUD,rec.RARDtStart,rec.RARDtStop);
    rec.RTRefStop = getEDIAdjustedStopDate(BUD,rec.RTRefStart,rec.RTRefStop);
    rec.RSDtStop = getEDIAdjustedStopDate(BUD,rec.RSDtStart,rec.RSDtStop);

    //-------------------------------------------------------------------------
    // if it called after 'save and add another' action there previous form
    // record is passed as Object else it is null
    //-------------------------------------------------------------------------
    if (previousFormRecord) {
        rec = setDefaultFormFieldAsPreviousRecord(
            [], //['RentableName'], // Fields to Reset
            rec,
            previousFormRecord
        );
    }

    return rec;
};

//-----------------------------------------------------------------------------
// getRentableTypes - return the RentableTypes list with respect of BUD
// @params
//      BUD - current business designation
//      cb  - callback routine (optional param)
// @return  the Rentable Types List
//-----------------------------------------------------------------------------
window.getRentableTypes = function (BUD,cb) {
    return jQuery.ajax({
        type: "GET",
        url: "/v1/rtlist/" + BUD,
        dataType: "json",
    }).done(function (data) {
        if (data.status == "success") {
            if (data.records) {
                app.rt_list[BUD] = data.records;
            } else {
                w2ui.rentablesGrid.error(data.message);
                app.rt_list[BUD] = [];
            }
            if (typeof cb === "function") {
                cb(BUD);
            }
        }
    });
};

//-----------------------------------------------------------------------------
// buildRentableElements - creates the UI components needed to manage rentables
//-----------------------------------------------------------------------------
window.buildRentableElements = function () {
    // inside rentable part, we need these lists
    app.cycleFreqItems = []; // cycle freq
    app.cycleFreq.forEach(function (item, index) {
        app.cycleFreqItems.push({id: index, text: item});
    });

    app.RSUseStatusItems = []; // rentable use status
    app.RSUseStatus.forEach(function (item, index) {
        app.RSUseStatusItems.push({id: index, text: item});
    });

    app.RSUseTypeItems = []; // rentable use type
    app.RSUseType.forEach(function (item, index) {
        app.RSUseTypeItems.push({id: 100 + index, text: item});
    });

    app.RSLeaseStatusItems = []; // rentable lease status
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
                field: 'UseType', caption: 'Rentable <br>Use Type', size: '100px', sortable: true,
                render: function (record/*, index, col_index*/) {
                    var text = '';
                    var offset = 100;
                    if (record) {
                        app.RSUseType.forEach(function (item, index) {
                            if (record.UseType == offset + index) {
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
                    } else {
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

                        var rec = grid.get(recid);
                        var x = getCurrentBusiness();
                        var BID = parseInt(x.value);
                        var BUD = getBUDfromBID(BID);

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

                // warn user if form content has been changed
                form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
            };
        },
        onAdd: function (/*event*/) {
            var yes_args = [this];
            var no_callBack = function () {
                    return false;
                };
            var yes_callBack = function (grid) {
                app.last.grid_sel_recid = -1; // reset it
                grid.selectNone();

                var x = getCurrentBusiness();
                var BID = parseInt(x.value);
                var BUD = getBUDfromBID(BID);

                w2ui.rentableForm.record = getRentableInitRecord(BID, BUD, null);
                RentableEdits.RTRChgList.push(w2ui.rentableForm.record.recid);
                w2ui.rentableForm.refresh();

                getRentableTypes(BUD)
                .done(function (data) {
                    if ('status' in data && data.status !== 'success') {
                        w2ui.rentableForm.message(data.message);
                    } else {
                        //-----------------------------------------------------------
                        // before we even attempt this, we must ensure that there is
                        // at least one rentable type.
                        //-----------------------------------------------------------
                        var cancel = true;
                        if (typeof app.rt_list[BUD] == "object") {
                            cancel = app.rt_list[BUD].length == 0;
                        }
                        if (cancel) {
                            w2ui.rentablesGrid.message("This business currently has no Rentable Types. Please create one or more Rentable Types and try again.");
                            return;
                        }
                        //-----------------------------------------------------------
                        // We have rentable types, we can proceed...
                        // get "RTID" column index and set rentable types list in editable items
                        //-----------------------------------------------------------
                        var RTIDColIndex = w2ui.rentableTypeRefGrid.getColumn("RTID", true);
                        w2ui.rentableTypeRefGrid.columns[RTIDColIndex].editable.items = app.rt_list[BUD];
                        setRentableLayout(BID, 0);
                    }
                })
                .fail(function () {
                    console.log('Error getting /v1/uival/' + BID + '/app.ReceiptRules');
                });
            };

            // warn user if form content has been changed
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
                        {id: 'rentableUseStatusGrid', caption: 'Use Status'},
                        {id: 'rentableUseTypeGrid', caption: 'Use Type'},
                        {id: 'rentableLeaseStatusGrid', caption: 'Lease Status'},
                        {id: 'rentableTypeRefGrid', caption: 'Rentable Type'},
                    ],
                    //---------------------------------
                    //  HANDLE THE TAB CLICKS...
                    //---------------------------------
                    onClick: function (event) {
                        if (event.target === "rentableForm") {
                            w2ui.rentableDetailLayout.html('main', w2ui.rentableForm);
                        } else if (event.target === "rentableUseStatusGrid") {
                            if (RentableEdits.UseStatusDataLoaded) {
                                w2ui.rentableUseStatusGrid.url = '';
                            }
                            w2ui.rentableDetailLayout.html('main', w2ui.rentableUseStatusGrid);
                        } else if (event.target === "rentableUseTypeGrid") {
                            if (RentableEdits.UseTypeDataLoaded) {
                                w2ui.rentableUseTypeGrid.url = '';
                            }
                            w2ui.rentableDetailLayout.html('main', w2ui.rentableUseTypeGrid);
                        } else if (event.target === "rentableLeaseStatusGrid") {
                            if (RentableEdits.LeaseStatusDataLoaded){
                                w2ui.rentableLeaseStatusGrid.url = '';
                            }
                            w2ui.rentableDetailLayout.html('main', w2ui.rentableLeaseStatusGrid);
                        } else if (event.target === "rentableTypeRefGrid") {
                            if (RentableEdits.RTRDataLoaded) {
                                w2ui.rentableTypeRefGrid.url = '';
                            }
                            w2ui.rentableDetailLayout.html('main', w2ui.rentableTypeRefGrid);
                        }
                    }
                }
            },
            {
                type: 'bottom',
                size: 60,
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
            app.form_is_dirty = false;
        },
        onRefresh: function (event) {
            event.onComplete = function () {
                var f = this;
                var r = f.record;
                var x = getCurrentBusiness();
                var BID = parseInt(x.value);
                //var BUD = getBUDfromBID(BID);
                var header = "";

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

    buildRentableUseStatusElements();
    buildRentableUseTypeElements();
    buildRentableLeaseStatusElements();
    buildRentableTypeRefElements();

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
                w2ui.rentablesGrid.selectNone();
                saveRentableCore(finishRentableSave);
            },

            saveadd: function () {
                saveRentableCore(finishRentableSaveAdd);
            }
        },
    });
};

// finishRentableSave performs the common functions for when a save is complete.
//
// @params
//
// @return
//-----------------------------------------------------------------------------
window.finishRentableSave = function() {
    app.form_is_dirty = false;
    w2ui.rentablesGrid.reload();
    w2ui.rentablesGrid.render();
    closeRentableForm();
};

// finishRentableSave performs the functions needed when a SaveAdd is complete.
//
// @params
//
// @return
//-----------------------------------------------------------------------------
window.finishRentableSaveAdd = function() {
    var BID = getCurrentBID();
    var BUD = getBUDfromBID();
    app.last.grid_sel_recid = -1;  // clear the grid select recid
    w2ui.rentablesGrid.reload();
    w2ui.rentablesGrid.render();
    w2ui.rentableForm.record = getRentableInitRecord(BID, BUD, w2ui.rentableForm.record);
    w2ui.rentableForm.url = '/v1/rentable/' + BID + '/0';
    w2ui.rentableForm.refresh();
    app.form_is_dirty = false;
};

// saveRentableCore performs the common functions to Save and SaveAdd.  It
// handles the async calls to store RentableTypeRefs, RentableUseStatus,
// RentableUseType and RentableLeaseStatus.
//
// @params
//     doneCB = callback function when all asynchronous calls complete
//
// @return
//-----------------------------------------------------------------------------
window.saveRentableCore = function (doneCB) {
    var BID = getCurrentBID();
    w2ui.rentableForm.url = '/v1/rentable/' + BID + '/' + w2ui.rentableForm.record.RID;

    var rec = w2ui.rentableForm.record;
    if (typeof rec.BUD === "object") {
        var B = rec.BUD.text;
        rec.BUD = B;
    }
    if (typeof rec.RTID === "object") {
        var R = rec.RTID.id;
        rec.RTID = R;
    }
    if (typeof rec.AssignmentTime === "object") {
        var A = rec.AssignmentTime.id;
        rec.AssignmentTime = A;
    }
    var params = {
        cmd: "save",
        selected: [],
        limit: 0,
        offset: 0,
        record: rec,
    };
    var dat = JSON.stringify(params);

    return $.post(w2ui.rentableForm.url, dat, null, "json")
    .done(function (data) {
        if (data.status === 'error') {
            console.log('ERROR: ' + data.message);
            return;
        }

        //--------------------------------------------------------------
        // if record is new then update RID that saved on server side
        // the recid is set to the RID of the newly created Rentable.
        //--------------------------------------------------------------
        w2ui.rentableForm.record.RID = data.recid;

        //--------------------------------------------------------------
        // update all associated records with the new RID
        //--------------------------------------------------------------
        var i;
        for (i = 0; i < w2ui.rentableUseStatusGrid.records.length; i++) {
            w2ui.rentableUseStatusGrid.records[i].RID = w2ui.rentableForm.record.RID;
        }
        for (i = 0; i < w2ui.rentableUseTypeGrid.records.length; i++) {
            w2ui.rentableUseTypeGrid.records[i].RID = w2ui.rentableForm.record.RID;
        }
        for (i = 0; i < w2ui.rentableLeaseStatusGrid.records.length; i++) {//add by lina
            w2ui.rentableLeaseStatusGrid.records[i].RID = w2ui.rentableForm.record.RID;
        }
        for (i = 0; i < w2ui.rentableTypeRefGrid.records.length; i++) {
            w2ui.rentableTypeRefGrid.records[i].RID = w2ui.rentableForm.record.RID;
        }

        //--------------------------------------------------------------
        // Now save all the newly created or modified records associated
        // with this rentable.  It is done asynchronously, and each of
        // the save routines implements the Promise interface
        //--------------------------------------------------------------
        var BID = w2ui.rentableForm.record.BID;
        var RID = w2ui.rentableForm.record.RID;

        $.when(
            saveRentableLeaseStatus(BID,RID),
            saveRentableUseStatus(BID,RID),
            saveRentableUseType(BID,RID),
            saveRentableTypeRef(BID,RID)
        )
        .done(function(){
            doneCB();
        })
        .fail(function(){
            var s = 'RentableSave: error reported';
            w2ui.rentablesGrid.error(s);
            doneCB();
        });
    })
    .fail( function() {
        w2ui.rentablesGrid.error('post failed to ' + w2ui.rentableForm.url);
    });
};

// setRentableLayout shows the RentableLayout.
//
// @params
//
// @return
//-----------------------------------------------------------------------------
window.setRentableLayout = function (BID, RID) {
    w2ui.rentableForm.url = '/v1/rentable/' + BID + '/' + RID;
    w2ui.rentableDetailLayout.content("bottom", w2ui.rentableFormBtns); // load bottom panels with action buttons panel

    //------------------------------------------------
    // if form has tabs then click the first one
    //------------------------------------------------
    if (typeof w2ui.rentableForm.tabs.name == "string") {
        w2ui.rentableForm.tabs.click('tab1');
    }

    app.new_form_rec = RID > 0; // mark this flag as is this new record
    app.form_is_dirty = false;  // new content will be loaded, mark form dirty flag as false
    RentableEdits.RID = RID;
    if (RID) {
        //------------------------------------------------
        // change the text of form tab
        //------------------------------------------------
        w2ui.rentableDetailLayout.get("main").tabs.get("rentableForm").text = "Rentable Details ({0})".format(RID);
        w2ui.rentableDetailLayout.get("main").tabs.refresh();

        //------------------------------------------------
        // load form content from server
        //------------------------------------------------
        w2ui.rentableForm.request(function (event) {
            if (event.status === "success") {
                //------------------------------------------------
                // only render the toplayout after server has sent down data
                // so that w2ui can bind values with field's html control,
                // otherwise it is unable to find html controls
                //------------------------------------------------
                showRentableForm();
                return true;
            }
            else {
                showRentableForm();
                w2ui.rentableForm.message("Could not get form data from server...!!");
                return false;
            }
        });
    }
    else {
        //------------------------------------------------
        // change the text of form tab
        //------------------------------------------------
        w2ui.rentableDetailLayout.get("main").tabs.get("rentableForm").text = "Rentable Details ({0})".format("new");
        w2ui.rentableDetailLayout.get("main").tabs.refresh();

        //------------------------------------------------
        // unselect the previous selected row
        //------------------------------------------------
        var sel_recid = parseInt(w2ui.rentablesGrid.last.sel_recid);
        if (sel_recid > -1) {
            //------------------------------------------------
            // if new record is being added then unselect
            // {{the selected record}} from the grid
            //------------------------------------------------
            w2ui.rentablesGrid.unselect(w2ui.rentablesGrid.last.sel_recid);
        }
        showRentableForm();
        return true;
    }
};

// closeRentableForm hides the RentableLayout.
//
// @params
//
// @return
//-----------------------------------------------------------------------------
window.closeRentableForm = function() {
    var no_callBack = function () {
            return false;
        },
        yes_callBack = function () {
            w2ui.toplayout.hide('right', true);
            w2ui.rentablesGrid.render();
        };
    form_dirty_alert(yes_callBack, no_callBack);
};

// showRentableForm initializes and shows all form parts needed for the Rentable.
//
// @params
//
// @return
//-----------------------------------------------------------------------------
window.showRentableForm = function() {
    var BID = getCurrentBID();
    var RID = w2ui.rentableForm.record.RID;

    //------------------------------------------------------------------------
    // We want the grids to request server data on their initial load
    // and on a RentableForm Save.  So, we will set the urls here and clear
    // them after the grids complete their loading or after a save completes.
    //------------------------------------------------------------------------
    if (RID > 0) {
        var br = '' + BID + '/' + RID;
        w2ui.rentableLeaseStatusGrid.url = '/v1/rentableleasestatus/' + br;
        w2ui.rentableUseStatusGrid.url   = '/v1/rentableusestatus/' + br;
        w2ui.rentableUseTypeGrid.url     = '/v1/rentableusetype/' + br;
        w2ui.rentableTypeRefGrid.url     = '/v1/rentabletyperef/' + br;
        w2ui.rentableForm.url            = '/v1/rentable/' + br;
    }

    //------------------------------------------------------------------------
    // We need to reset any changes in the change lists now
    //------------------------------------------------------------------------
    RentableEdits.LeaseStatusChgList = [];
    RentableEdits.UseStatusChgList = [];
    RentableEdits.UseTypeChgList = [];
    RentableEdits.RTRChgList = [];
    RentableEdits.RSID = -1;
    RentableEdits.RTRID = -1;
    RentableEdits.UTID = -1;
    RentableEdits.LeaseStatusDataLoaded = false;
    RentableEdits.UseStatusDataLoaded = false;
    RentableEdits.UseTypeDataLoaded = false;
    RentableEdits.RTRDataLoaded = false;
    w2ui.rentableLeaseStatusGrid.records = [];
    w2ui.rentableUseStatusGrid.records = [];
    w2ui.rentableUseTypeGrid.records = [];
    w2ui.rentableTypeRefGrid.records = [];

    //------------------------------------------------------------------------
    // Initialize the statuses.  This is a big time saver.
    //------------------------------------------------------------------------
    if (RID == 0) {
        addRentableLeaseStatus();
        addRentableUseStatus();
        addRentableUseType();
        addRentableTypeRef();
    }

    //------------------------------------------------------------------------
    // Now we can sHOW the right panel
    //------------------------------------------------------------------------
    w2ui.toplayout.content('right', w2ui.rentableDetailLayout);
    w2ui.toplayout.sizeTo('right', 700);
    w2ui.rentableDetailLayout.get("main").tabs.click("rentableForm");
    w2ui.toplayout.show('right', true);
};
