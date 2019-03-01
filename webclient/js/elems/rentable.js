/*global
    setDefaultFormFieldAsPreviousRecord, w2uiDateControlString, $, w2ui, app, getCurrentBusiness, parseInt, getBUDfromBID,
    getRentableTypes, setToForm, form_dirty_alert, console, getFormSubmitData, addDateNavToToolbar, setRentableLayout,
    getRentableInitRecord, saveRentableLeaseStatus, buildRentableUseStatusElements, buildRentableLeaseStatusElements,
    saveRentableUseStatus, saveRentableTypeRef, buildRentableTypeRefElements, saveRentableCore, closeRentableForm,
*/
/*jshint esversion: 6 */

"use strict";
var RentableEdits = {
    LeaseStatusChgList: [],     // an array of indeces to LeaseStatus changes
    UseStatusChgList: [],       // an array of indeces to UseStatus changes
    RTRChgList: [],             // an array of indeces to type ref changes
    rlsDeleteInProgress: false, // indicates whether or not a delete of Rentable Lease Status is in progress
};

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
        RRentableLeaseStatus: "unknown",//add by lina
        RSDtStart: w2uiDateControlString(y),
        RSDtStop: "1/1/9999",
        AssignmentTime: 0,
        Comment: ""
    };

    // if it called after 'save and add another' action there previous form record is passed as Object
    // else it is null
    if (previousFormRecord) {
        defaultFormData = setDefaultFormFieldAsPreviousRecord(
            [],//['RentableName'], // Fields to Reset
            defaultFormData,
            previousFormRecord
        );
    }

    return defaultFormData;
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
                app.rt_list[BUD] = [];
            }
            if (typeof cb === "function") {
                cb(BUD);
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
                    RentableEdits.TypeRefChgList.push(w2ui.rentableForm.record.recid);
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
                        {id: 'rentableUseStatusGrid', caption: 'Rentable Use Status'},
                        {id: 'rentableLeaseStatusGrid', caption: 'Rentable Lease Status'},//add by lina
                        {id: 'rentableTypeRefGrid', caption: 'Rentable Type Ref'},
                    ],
                    onClick: function (event) {
                        if (event.target === "rentableForm") {
                            w2ui.rentableDetailLayout.html('main', w2ui.rentableForm);
                        } else if (event.target === "rentableUseStatusGrid") {
                            w2ui.rentableDetailLayout.html('main', w2ui.rentableUseStatusGrid);
                        } else if (event.target === "rentableLeaseStatusGrid") {//add by lina
                            w2ui.rentableDetailLayout.html('main', w2ui.rentableLeaseStatusGrid);
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
                var BID = getCurrentBID();
                var BUD = getBUDfromBID();
                w2ui.rentablesGrid.selectNone();
                saveRentableCore();
                app.form_is_dirty = false;
                closeRentableForm();
            },

            saveadd: function () {
                var BID = getCurrentBID();
                var BUD = getBUDfromBID();
                app.form_is_dirty = false;
                app.last.grid_sel_recid = -1;  // clear the grid select recid
                // w2ui.rentablesGrid.selectNone();  // select none if you're going to add new record
                saveRentableCore();
                w2ui.rentablesGrid.render();
                w2ui.rentableTypeRefGrid.reload();
                w2ui.rentableForm.record = getRentableInitRecord(BID, BUD, w2ui.rentableForm.record);
                w2ui.rentableForm.url = '/v1/rentable/' + BID + '/0';
                w2ui.rentableForm.refresh();
            }

        },
    });
};

// saveRentableCore performs the common functions to Save and SaveAdd.  It
// handles the async calls to store RentableTypeRefs, RentableUseStatus, and
// RentableLeaseStatus.
//
// @params
//     BID - current business
//     BUD - designator for current business
//
// @return
//-----------------------------------------------------------------------------
window.saveRentableCore = function (BID, BUD) {
    w2ui.rentableForm.save({}, function (data) {
        if (data.status === 'error') {
            console.log('ERROR: ' + data.message);
            return;
        }

        // if record is new then update RID that saved on server side
        w2ui.rentableForm.record.RID = data.recid;

        var i;
        // update RID in grid records (Use status)
        for (i = 0; i < w2ui.rentableUseStatusGrid.records.length; i++) {
            w2ui.rentableUseStatusGrid.records[i].RID = w2ui.rentableForm.record.RID;
        }

        // update RID in grid records (Lease status)
        for (i = 0; i < w2ui.rentableLeaseStatusGrid.records.length; i++) {//add by lina
            w2ui.rentableLeaseStatusGrid.records[i].RID = w2ui.rentableForm.record.RID;
        }

        // update RID in grid records (typeRef)
        for (i = 0; i < w2ui.rentableTypeRefGrid.records.length; i++) {
            w2ui.rentableTypeRefGrid.records[i].RID = w2ui.rentableForm.record.RID;
        }

        var BID = w2ui.rentableForm.record.BID;
        var RID = w2ui.rentableForm.record.RID;

        $.when(
            saveRentableLeaseStatus(BID,RID),
            saveRentableUseStatus(BID,RID),
            saveRentableTypeRef(BID,RID)
        )
        .done(function(){
            console.log('RentableSave: when completed, no errors');
        })
        .fail(function(){
            console.log('RentableSave: when failed.');
        });
    });
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

// setRentableLayout shows the RentableLayout.
//
// @params
//
// @return
//-----------------------------------------------------------------------------
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
        w2ui.rentableUseStatusGrid.load('/v1/rentableusestatus/' + BID + '/' + RID);
        console.log("rentable status grid load data: {0}".format('/v1/rentableusestatus/' + BID + '/' + RID));

        // if rentable available then load the Lease status grid
        w2ui.rentableLeaseStatusGrid.load('/v1/rentableleasestatus/' + BID + '/' + RID);
        console.log("rentable status grid load data: {0}".format('/v1/rentableleasestatus/' + BID + '/' + RID));//add by lina


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
        w2ui.rentableUseStatusGrid.clear();
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
        RentableEdits.LeaseStatusChgList = [];
        RentableEdits.UseStatusChgList = [];
        RentableEdits.RTRChgList = [];
        // SHOW the right panel now
        w2ui.toplayout.content('right', w2ui.rentableDetailLayout);
        w2ui.toplayout.sizeTo('right', 700);
        // w2ui.rentableDetailLayout.render();
        w2ui.rentableDetailLayout.get("main").tabs.click("rentableForm");
        w2ui.toplayout.show('right', true);
    }
};
