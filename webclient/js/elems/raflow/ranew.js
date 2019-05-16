/*global
    InitRAFlowAjax,
    RACompConfig, w2ui,
    GetRAFlowDataAjax,
    manageParentRentableW2UIItems, managePeopleW2UIItems,
    LoadRAFlowTemplate,
    loadRAActionTemplate,
    getStringListData, initBizErrors, displayErrorDot,
    ChangeRAFlowVersionToolbar, GetRefNoByRAIDFromGrid,
    RenderRAFlowVersionData, CloseRAFlowLayout, DeleteRAFlowAJAX, HideRAFlowLoader
*/

"use strict";

//-----------------------------------------------------------------------------
// LoadRAFlowTemplate - load RA flow with data and green checkmark and
//                      necessary settings, it loads dateForm by default
//
// @params
//   FlowID = Id of the Flow
//-----------------------------------------------------------------------------
window.LoadRAFlowTemplate = function() {
    if("RAActionForm" in w2ui){
        w2ui.RAActionForm.destroy();
    }

    if("raActionLayout" in w2ui){
        w2ui.raActionLayout.destroy();
        w2ui.newraLayout.get('right').content = "";
        w2ui.newraLayout.hide('right');
    }

    // show the loader
    HideRAFlowLoader(false);

    // set the toplayout content
    w2ui.toplayout.content('right', w2ui.newraLayout);
    w2ui.toplayout.show('right', true);
    w2ui.toplayout.sizeTo('right', 950);

    $.get('/webclient/html/raflow/raflowtmpl.html', function(htmlData) {
        // set the content of template HTML into main content of layout
        w2ui.newraLayout.content('main', htmlData);

        // render the new ra layout
        w2ui.newraLayout.render();

        // reset wizard steps
        $(".ra-form-component").hide();
        $("#progressbar #steps-list li").removeClass("active done"); // remove activeClass from all li

        setTimeout(function() {
            // RENDER THE VERSION DATA
            var active_comp_id = "dates";
            RenderRAFlowVersionData(active_comp_id);

            // hide the loader
            HideRAFlowLoader(true);
        }, 0);
    });
};

window.buildRAFlowElements = function() {
    // ------------------------------------------------------
    // raflows grid
    // ------------------------------------------------------
    $().w2grid({
        name: 'raflowsGrid',
        multiSelect: false,
        show: {
            toolbar: true,
            footer: true,
            lineNumbers: false,
            selectColumn: false,
            expandColumn: false,
            toolbarAdd: true,
            toolbarDelete: false,
            toolbarSave: false,
            toolbarEdit: false,
            toolbarSearch: true,
            toolbarInput: true,
            searchAll: true,
            toolbarReload: true,
            toolbarColumns: false,
        },
        searches: [
            // { field: 'RAID', caption: 'RAID', type: 'text' },
            { field: 'Payors', caption: 'Payor(s)', type: 'text' },
            { field: 'RNames', caption: 'Rentable(s)', type: 'text' },
            // { field: 'AgreementStart', caption: 'Agreement Start Date', type: 'date' },
            // { field: 'AgreementStop', caption: 'Agreement Stop Date', type: 'date' },
            { field: 'UserRefNo', caption: 'Reference Number', type: 'text' },
        ],
        columns: [
            { field: 'recid',          caption: 'recid',              size: '40px', hidden: true, sortable: true },
            { field: 'BID',            caption: 'BID',  hidden: true },
            { field: 'BUD',            caption: 'BUD', hidden: true },
            { field: 'RAID',           caption: 'RAID',               size: "45px", sortable: true },
            { field: 'UserRefNo',      caption: 'Ref No',             size: '165px', sortable:   true },
            { field: 'AgreementStart', caption: 'Agreement<br>Start', size: '80px', render: 'date', sortable: true, style: 'text-align: right' },
            { field: 'AgreementStop',  caption: 'Agreement<br>Stop',  size: '80px', render: 'date', sortable: true, style: 'text-align: right' },
            { field: 'Payors',         caption: 'Payor(s)',           size: '250px', sortable: true },
            { field: 'RNames',         caption: 'Rentable(s)',        size: '250px', sortable: true },
            { field: 'FlowID',         caption: 'Flow ID',            size: '50px', hidden: true, sortable:   true },
        ],
        onRequest: function(event) {
            event.postData.cmd = "all";
            event.postData.FlowType = "RA";
        },
        onRefresh: function(event) {
            event.onComplete = function() {
                // var sel_recid = parseInt(this.last.sel_recid);
                if (app.active_grid == this.name) {
                    this.select(app.last.grid_sel_recid);
                    // This one is special case, you need to set last sel_recid when you're adding
                    // new record with help of onAdd event handler, so new record automatically
                    // will be selected

                    /*if (app.new_form_rec) {
                        this.unselect(sel_recid);
                    }
                    else{
                        this.select(sel_recid);
                    }*/
                }
            };
        },
        onClick: function(event) {
            event.onComplete = function () {
                var yes_args = [this, event.recid],
                    no_args = [this],
                    no_callBack = function(grid) {
                        grid.select(app.last.grid_sel_recid);
                        return false;
                    },
                    yes_callBack = function(grid, recid) {
                        app.last.grid_sel_recid = parseInt(recid);

                        // keep highlighting current row in any case
                        grid.select(app.last.grid_sel_recid);

                        // get grid record
                        var rec = grid.get(recid);
                        var version = (rec.RAID > 0) ? "raid" : "refno";

                        GetRAFlowDataAjax(rec.UserRefNo, rec.RAID, version)
                        .done(function(data) {
                            if (data.status != "success") {
                                grid.message(data.message);
                            } else {
                                LoadRAFlowTemplate();

                                // Update local copy of string list
                                var BID = getCurrentBID();
                                var BUD = getBUDfromBID(BID);
                                getStringListData(BID, BUD);
                                initBizErrors();
                            }
                        })
                        .fail(function() {
                            grid.message("Error while fetching data for selected record");
                        });
                    };

                // warn user if form content has been changed
                form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
            };
        },
        onAdd: function(/*event*/) {
            var yes_args = [this],
                no_args = [this],
                no_callBack = function(grid) {
                    grid.select(app.last.grid_sel_recid);
                    return false;
                },
                yes_callBack = function(grid, recid) {
                    InitRAFlowAjax()
                    .done(function(data, textStatus, jqXHR) {
                        if (data.status === "success") {
                            var bid = getCurrentBID(),
                                bud = getBUDfromBID(bid);

                            // Update local copy of string list
                            getStringListData(bid, bud);

                            var newRecid = grid.records.length;

                            // add new record
                            grid.add({
                                recid:          newRecid,
                                BID:            bid,
                                BUD:            bud,
                                RAID:           0,
                                Payors:         null,
                                AgreementStart: null,
                                AgreementStop:  null,
                                FlowID:         data.record.Flow.FlowID,
                                UserRefNo:      data.record.Flow.UserRefNo,
                            });

                            grid.refresh();

                            app.last.grid_sel_recid = parseInt(newRecid);

                            // keep highlighting current row in any case
                            grid.select(app.last.grid_sel_recid);

                            var rec = grid.get(newRecid);
                            LoadRAFlowTemplate();

                        } else {
                            grid.message(data.message);
                        }
                    })
                    .fail(function() {
                        grid.message("error while creating new flow ID");
                    });

                };

            // warn user if form content has been changed
            form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
        }
    });

    // add date navigation toolbar for new rental agreement form
    addDateNavToToolbar('raflows');

    //------------------------------------------------------------------------
    //          Rental Agreement Details
    //------------------------------------------------------------------------
    $().w2layout({
        name: 'newraLayout',
        padding: 0,
        panels: [
            { type: 'left',         hidden: true },
            { type: 'top',          hidden: true },
            { type: 'main',         size: '60%',    resizable: true,    style: app.pstyle,
                content: 'main',
                toolbar: {
                    items: [
                        { id: 'btnNotes',       type: 'button',     icon: 'far fa-sticky-note' },
                        { id: 'id',             type: 'html' },
                        { id: 'state',          type: 'html' },
                        { id: 'bt3',            type: 'spacer' },
                        { id: 'versionMode',    type: 'html' },
                        {                       type: 'break' },
                        { id: 'stateAction',    type: 'html' },
                        { id: 'remove-refno',   type: 'html',
                            html: '<button title="Delete" id="remove_raflow" name="remove_raflow" class="w2ui-btn" style="min-width: 30px; padding: 6px 0px;"><i class="fas fa-trash"></i></button>' },
                        {                       type: 'break' },
                        { id: 'editViewBtn',    type: 'html' },
                        {                       type: 'break' },
                        { id: 'btnClose',       type: 'button',     icon: 'fas fa-times' }
                    ],
                    onClick: function (event) {
                        console.log(event.target);
                        switch(event.target) {
                        case 'btnClose':
                            var no_callBack = function() { return false; },
                                yes_callBack = function() {
                                    // reset validationError. cause it should display error when it pressed GetApproval button
                                    initBizErrors();
                                    CloseRAFlowLayout();
                                };
                            form_dirty_alert(yes_callBack, no_callBack);
                            break;
                        }
                    },
                    onRefresh: function(event) {
                        var toolbar = this;
                        event.onComplete = function() {
                            // ADDITIONAL CHECK REQUIRED HERE - SPECIAL ONE
                            // BECAUSE WE DON'T KNOW WHEN RENDER WILL COMPLETE
                            // WHEN TOOLBAR IS COMPLETELY REFRESHED THEN ALSO CHECK FOR BUTTON
                            if (app.raflow.loading) {
                                $(toolbar.box).find("button").prop('disabled', true);
                            } else {
                                $(toolbar.box).find("button").prop('disabled', false);
                            }
                        };
                    },
                }
            },
            { type: 'preview',      hidden: true },
            { type: 'bottom',       hidden: true },
            { type: 'right',        hidden: true, size: '200', resizable: true }
        ],
        onResize: function(event) {
            event.onComplete = function() {
                $("#raflow-container .slider").width($(this.box).width());
            };
        },
        onRefresh: function(event) {
            event.onComplete = function() {
                $("button#save-ra-flow-btn").prop("disabled", (app.raflow.version === "raid"));
            };
        }
    });
};

//-----------------------------------------------------------------------
// ChangeRAFlowVersionToolbar - change the toolbar items content based on
//                              the requested version of raflow
//
// @params
//   version    = raflow version ("raid" / "refno")
//   RefNo      = Flow Reference No
//   RAID       = Associated Rental Agreement ID if exists (optional)
//   FLAGS      = Current version raflow FLAGS, will render the state
//-----------------------------------------------------------------------
window.ChangeRAFlowVersionToolbar = function(version, RAID, RefNo, FLAGS) {

    // RESET RAID
    if (!RAID) {
        RAID = 0;
    }

    // RESET REF NO
    if (!RefNo) {
        RefNo = "";
    }

    // GET STATE STRING USING FLAGS
    var state = app.RAStates[parseInt(FLAGS & 0xF)];
    var stateHTML = "<p style='margin:0 10px; font-size: 10pt;'>State:&nbsp;<span id='RAState'>" + state + "</span></p>";

    // STATE CHANGE ACTIONS BUTTON HTML
    var stateActionHTML = "<button class='w2ui-btn' id='raactions' name='raactions'><i class='fas fa-cog' style='margin-right: 7px;'></i>Actions</button>";

    var idString = "",
        editViewBtnHTML = "",
        versionMode = "",
        btnBackToRAText = "";

    switch(version) {
        case "raid":
            idString = "<p style='margin:0 10px; font-size: 12pt;'><strong>RA" + RAID + "</strong></p>";
            editViewBtnHTML = "<button class='w2ui-btn' id='edit_view_raflow' name='edit_view_raflow'><i class='fas fa-pencil-alt fa-sm' style='margin-right: 7px;'></i>Edit" + (RefNo ? "&nbsp;&nbsp;" + RefNo : "") + "</button>";
            versionMode = "Viewing";
            btnBackToRAText = "Back to RA" + RAID;

            // HIDE TRASH ICON
            w2ui.newraLayout.get("main").toolbar.hide('remove-refno');

            // SHOW EDIT/VIEW BUTTON AND SET THE TEXT
            w2ui.newraLayout.get("main").toolbar.show('editViewBtn');
            w2ui.newraLayout.get("main").toolbar.set('editViewBtn', {html: editViewBtnHTML});
            break;

        case "refno":
            idString = "<p style='margin:0 10px; font-size: 12pt;'><strong>" + RefNo + "</strong></p>";            editViewBtnHTML = "<button class='w2ui-btn' id='edit_view_raflow' name='edit_view_raflow'><i class='fas fa-eye fa-sm' style='margin-right: 7px;'></i>" + (RAID ? "View RA" + RAID : "") + "</button>";
            versionMode = "Editing";
            btnBackToRAText = "Back to " + RefNo;

            // SHOW TRASH ICON
            w2ui.newraLayout.get("main").toolbar.show('remove-refno');

            // SHOW EDIT/VIEW BUTTON AND SET THE TEXT BASED ON RAID
            if (RAID > 0) {
                w2ui.newraLayout.get("main").toolbar.show('editViewBtn');
                w2ui.newraLayout.get("main").toolbar.set('editViewBtn', {html: editViewBtnHTML});
            } else {
                w2ui.newraLayout.get("main").toolbar.hide('editViewBtn');
            }
            break;
    }

    // VERSION MODE
    var versionModeHTML = "<small style='font-size: 8.5pt; color: #555; vertical-align: 2px; margin-left: 10px;'>"+versionMode+"</small>";

    // SET THE ID AND STATE IN TOOLBAR
    w2ui.newraLayout.get("main").toolbar.set('id', {html: idString});
    w2ui.newraLayout.get("main").toolbar.set('state', {html: stateHTML});
    w2ui.newraLayout.get("main").toolbar.set('stateAction', {html: stateActionHTML});
    w2ui.newraLayout.get("main").toolbar.set('versionMode', {html: versionModeHTML});

    // REFRESH THE TOOLBAR TO GET THE EFFECT
    w2ui.newraLayout.get("main").toolbar.refresh();

    // TOP TOOLBAR IN ACTION LAYOUT
    if (w2ui.raActionLayout) {
        w2ui.raActionLayout.get("top").toolbar.set('btnBackToRA', {text: btnBackToRAText});
        // REFRESH THE TOOLBAR TO GET THE EFFECT
        w2ui.raActionLayout.get("top").toolbar.refresh();
    }
};

//-----------------------------------------------------------------------
// RenderRAFlowVersionData will load the RAID versioned data
//        in the interface from local copy
//-----------------------------------------------------------------------
window.RenderRAFlowVersionData = function(active_comp_id, prev_comp_id) {

    // mark this flag as is this new record
    // record created already
    app.new_form_rec = false;

    // as new content will be loaded for this form
    // mark form dirty flag as false
    app.form_is_dirty = false;

    // calculate people items
    managePeopleW2UIItems();

    // calculate parent rentable items
    manageParentRentableW2UIItems();

    var FLAGS   = app.raflow.Flow.Data.meta.RAFLAGS,
        RAID    = app.raflow.Flow.ID,
        RefNo   = (RAID > 0) ? GetRefNoByRAIDFromGrid(RAID) : app.raflow.Flow.UserRefNo;

    // change toolbar
    ChangeRAFlowVersionToolbar(app.raflow.version, RAID, RefNo, FLAGS);
    // hide save button
    $("button#save-ra-flow-btn").prop("disabled", (app.raflow.version === "raid"));

    // clear grid, form if previously loaded in DOM
    for (var comp in app.raFlowPartTypes) {
        // reset w2ui component as well
        if(RACompConfig[comp].w2uiComp in w2ui) {
            // clear inputs
            w2ui[RACompConfig[comp].w2uiComp].clear();
        }
    }

    // LOAD THE CURRENT COMPONENT AGAIN
    loadTargetSection(active_comp_id, prev_comp_id);
};

//-----------------------------------------------------------------------------
// EDIT/VIEW RAFLOW BUTTON CLICK EVENT HANDLER
//-----------------------------------------------------------------------------
$(document).on("click", "button#edit_view_raflow", function(e) {
    var RAID            = app.raflow.Flow.ID,
        RefNo           = (RAID > 0) ? GetRefNoByRAIDFromGrid(RAID) : app.raflow.Flow.UserRefNo,
        versionToRender = (app.raflow.version === "raid") ? "refno" : "raid";

    // GET THE DATA FROM SERVER FOR VERSION TO RENDER THE DATA
    GetRAFlowDataAjax(RefNo, RAID, versionToRender)
    .done(function(data) {
        if (data.status !== "error") {
            var active_comp = $(".ra-form-component:visible");
            var active_comp_id = active_comp.attr("id");
            RenderRAFlowVersionData(active_comp_id);
        }
    })
    .fail(function() {
        alert("Error while fetching data for selected record");
    });
});

//-----------------------------------------------------------------------------
// ACTIONS BUTTON CLICK EVENT HANDLER
//-----------------------------------------------------------------------------
$(document).on("click", "button#raactions", function(e) {
    var BID = getCurrentBID();
    var BUD = getBUDfromBID(BID);
    getStringListData(BID, BUD);

    w2ui.newraLayout.lock('main');
    // set the newralayout's right panel content
    setTimeout(function() {
        loadRAActionTemplate();
    }, 500);
});

// CloseRAFlowLayout closes the new ra layout with resetting right panel content
window.CloseRAFlowLayout = function() {
    app.raflow.version = ""; // RESET THE RAFLOW VERSION
    if (w2ui.raActionLayout) {
        w2ui.raActionLayout.content('main', '');
    }
    w2ui.newraLayout.unlock('main');
    w2ui.newraLayout.content('right', '');
    w2ui.newraLayout.hide('right', true);
    w2ui.toplayout.hide('right', true);
    w2ui.raflowsGrid.render();
    app.form_is_dirty = false;
};

//-----------------------------------------------------------------------------
// REMOVE FLOW BUTTON CLICK EVENT HANDLER
//-----------------------------------------------------------------------------
$(document).on("click", "button#remove_raflow", function(e) {
    e.preventDefault();

    var version = app.raflow.version,
        RefNo   = app.raflow.Flow.UserRefNo;

    // ONLY REF.NO CAN BE DELETED
    if (version === "raid") { // IF RAID VERSION LOADED THEN
        return;
    }

    // delete the flow
    DeleteRAFlowAJAX(RefNo)
    .done(function(data) {
        if (data.status ==="success") {
            CloseRAFlowLayout();
        } else {
            alert(data.message);
        }
    });
});
