/*global
    initRAFlowAjax,
    RACompConfig, w2ui,
    GetRAFlowDataAjax,
    manageParentRentableW2UIItems, managePeopleW2UIItems,
    LoadRAFlowTemplate,
    loadRAActionTemplate,
    getStringListData, initBizErrors, displayErrorDot,
    ChangeRAFlowVersionToolbar, GetRefNoByRAID
*/

"use strict";

//-----------------------------------------------------------------------------
// LoadRAFlowTemplate - load RA flow with data and green checkmark and
//                      necessary settings, it loads dateForm by default
//
// @params
//   FlowID = Id of the Flow
//-----------------------------------------------------------------------------
window.LoadRAFlowTemplate = function(bid) {

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
            $("#ra-form footer button#previous").prop("disabled", true);

            // mark this flag as is this new record
            // record created already
            app.new_form_rec = false;

            // as new content will be loaded for this form
            // mark form dirty flag as false
            app.form_is_dirty = false;

            // set BID in raflow settings
            app.raflow.BID = bid;

            // calculate people items
            managePeopleW2UIItems();

            // calculate parent rentable items
            manageParentRentableW2UIItems();

            // get info from local copy
            var FLAGS = app.raflow.Flow.Data.meta.RAFLAGS,
                RAID = app.raflow.Flow.ID,
                RefNo = GetRefNoByRAID(RAID);

            // render the toolbar based on raflow version
            ChangeRAFlowVersionToolbar("raid", RAID, RefNo, FLAGS);

            // clear grid, form if previously loaded in DOM
            for (var comp in app.raFlowPartTypes) {
                // reset w2ui component as well
                if(RACompConfig[comp].w2uiComp in w2ui) {
                    // clear inputs
                    w2ui[RACompConfig[comp].w2uiComp].clear();
                }
            }

            // mark first slide as active
            $(".ra-form-component#dates").show();
            $("#progressbar #steps-list li[data-target='#dates']").removeClass("done").addClass("active");
            loadRADatesForm();
        }, 0);
    });
};

window.buildRAApplicantElements = function() {
    // ------------------------------------------------------
    // applicants grid
    // ------------------------------------------------------
    $().w2grid({
        name: 'applicantsGrid',
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
            { field: 'RAID', caption: 'RAID', type: 'text' },
            { field: 'Payors', caption: 'Payor(s)', type: 'text' },
            { field: 'AgreementStart', caption: 'Agreement Start Date', type: 'date' },
            { field: 'AgreementStop', caption: 'Agreement Stop Date', type: 'date' },
            { field: 'UserRefNo', caption: 'Reference Number', type: 'text' },
        ],
        columns: [
            {
                field: 'recid',
                caption: 'recid',
                size: '40px',
                hidden: true,
                sortable: true
            },
            {
                field: 'BID',
                caption: 'BID',
                hidden: true
            },
            {
                field: 'BUD',
                caption: 'BUD',
                hidden: true
            },
            {
                field: 'RAID',
                caption: 'RAID',
                size: "60px",
                sortable: true,
                render: function (record) {
                    if (record.RAID) {
                        return "RAID - " + record.RAID;
                    }
                }
            },
            {
                field: 'Payors',
                caption: 'Payor(s)',
                size: '250px',
                sortable: true
            },
            {
                field: 'AgreementStart',
                caption: 'Agreement<br>Start',
                render: 'date',
                size: '80px',
                sortable: true,
                style: 'text-align: right'
            },
            {
                field: 'AgreementStop',
                caption: 'Agreement<br>Stop',
                render: 'date',
                size: '80px',
                sortable: true,
                style: 'text-align: right'
            },
            {
                field: 'FlowID',
                caption: 'Flow ID',
                size: '50px',
                hidden: true,
                sortable:   true
            },
            {
                field: 'UserRefNo',
                caption: 'Ref No',
                size: '200px',
                sortable:   true
            },
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

                        GetRAFlowDataAjax(rec.FlowID, rec.RAID)
                        .done(function(data) {
                            if (data.status != "success") {
                                grid.message(data.message);
                            } else {
                                LoadRAFlowTemplate(rec.BID);
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
                    initRAFlowAjax()
                    .done(function(data, textStatus, jqXHR) {
                        if (data.status === "success") {
                            var bid = getCurrentBID(),
                                bud = getBUDfromBID(bid);

                            var newRecid = grid.records.length;

                            // add new record
                            grid.add({
                                recid:  newRecid,
                                BID:    bid,
                                BUD:    bud,
                                RAID:   0,
                                FlowID: data.record.Flow.FlowID
                            });

                            grid.refresh();

                            app.last.grid_sel_recid = parseInt(newRecid);

                            // keep highlighting current row in any case
                            grid.select(app.last.grid_sel_recid);

                            var rec = grid.get(newRecid);
                            LoadRAFlowTemplate(rec.BID);

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
    addDateNavToToolbar('applicants');

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
                        { id: 'id',             type: 'html',       html: '' },
                        {                       type: 'break' },
                        { id: 'state',          type: 'html',       html: '' },
                        {                       type: 'break' },
                        { id: 'editViewBtn',    type: 'button',     text: '' },
                        { id: 'bt3',            type: 'spacer' },
                        { id: 'stateAction',    type: 'button',     icon: 'fas fa-cog', text: 'Actions' },
                        {                       type: 'break' },
                        { id: 'btnClose',       type: 'button',     icon: 'fas fa-times' }
                    ],
                    onClick: function (event) {
                        switch(event.target) {
                        case 'btnClose':
                            var no_callBack = function() { return false; },
                                yes_callBack = function() {
                                    w2ui.toplayout.hide('right',true);
                                    w2ui.applicantsGrid.render();
                                };
                            form_dirty_alert(yes_callBack, no_callBack);
                            break;
                        case 'stateAction':
                            var BID = getCurrentBID();
                            var BUD = getBUDfromBID(BID);
                            getStringListData(BID, BUD);

                            w2ui.newraLayout.lock('main');
                            // set the newralayout's right panel content
                            setTimeout(function() {
                                loadRAActionTemplate();
                            }, 500);
                            break;
                        }
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

    // get state string
    var state = app.RAStates[parseInt(FLAGS & 0xF)];
    var stateHTML = '<p>State: <span id="RAState">' + state + '</span></p>';

    var idString = "",
        editViewBtnCaption = "",
        editViewBtnIcon = "";

    switch(version) {
        case "raid":
            idString = "<p><strong>RA" + RAID + "</strong></p>";
            editViewBtnCaption = "Edit" + (RefNo ? " " + RefNo : "");
            editViewBtnIcon = "fas fa-pencil-alt fa-xs";
            break;
        case "refno":
            idString = "<p><strong>" + RefNo + "</strong></p>";
            editViewBtnCaption = "View" + (RAID ? " RA" + RAID : "");
            editViewBtnIcon = "";
            break;
    }

    // icon: 'fas fa-pencil-alt fa-sm'

    w2ui.newraLayout.get("main").toolbar.set('id', {html: idString});
    w2ui.newraLayout.get("main").toolbar.set('editViewBtn', {text: editViewBtnCaption, icon: editViewBtnIcon});
    w2ui.newraLayout.get("main").toolbar.set('state', {html: stateHTML});
    w2ui.newraLayout.get("main").toolbar.refresh();
};