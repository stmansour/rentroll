/*global
    initRAFlowAjax,
    RACompConfig, w2ui,
    getFlowDataAjax,
    manageParentRentableW2UIItems, managePeopleW2UIItems,
    LoadRAFlowTemplate,
    validateRAFlowComponents,
    getVehicleFees, getPetFees,
    renderRAStateInToolbar,
    loadRAActionTemplate,
    getStringListData
*/

"use strict";

//-----------------------------------------------------------------------------
// setToNewRAForm -  enable the Rental Agreement form in toplayout.  Also, set
//                   the forms url and request data from the server
// @params
//   bid = business id (or the BUD)
//-----------------------------------------------------------------------------
window.setToNewRAForm = function (bid, FlowID) {

    if (FlowID < 1) {
        return false;
    }

    // load ra flow template
    LoadRAFlowTemplate(bid, FlowID);
};

//-----------------------------------------------------------------------------
// LoadRAFlowTemplate - load RA flow with data and green checkmark and
//                      necessary settings, it loads dateForm by default
//
// @params
//   FlowID = Id of the Flow
//-----------------------------------------------------------------------------
window.LoadRAFlowTemplate = function(bid, FlowID) {

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

            // set this flow id as in active
            app.raflow.activeFlowID = FlowID;

            // set BID in raflow settings
            app.raflow.BID = bid;

            // calculate people items
            managePeopleW2UIItems();

            // calculate parent rentable items
            manageParentRentableW2UIItems();

            var raFlags = app.raflow.data[FlowID].Data.meta.RAFLAGS;

            // renders the Rental Agreement State in Toolbar
            renderRAStateInToolbar(raFlags);

            // show "done" mark on each li of navigation bar
            validateRAFlowComponents();

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

        // get pet and vehicle fees on loading rental agreement form
        getPetFees();
        getVehicleFees();
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
        columns: [
            {field: 'recid',     caption: 'recid',   size: '40px',   hidden: true, sortable:   true },
            {field: 'BID',       caption: 'BID',                     hidden: true,                  },
            {field: 'BUD',       caption: 'BUD',                     hidden: true,                  },
            {field: 'FlowID',    caption: 'Flow ID', size: '50px',                 sortable:   true },
            {field: 'UserRefNo', caption: 'Ref No',  size: '200px',                sortable:   true },
        ],
        onRequest: function(event) {
            event.postData.cmd = "getAllFlows";
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

                        getFlowDataAjax(rec.FlowID)
                        .done(function(data) {
                            if (data.status != "success") {
                                grid.message(data.message);
                            } else {
                                setToNewRAForm(rec.BID, rec.FlowID);
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
                                FlowID: data.record.Flow.FlowID
                            });

                            grid.refresh();

                            app.last.grid_sel_recid = parseInt(newRecid);

                            // keep highlighting current row in any case
                            grid.select(app.last.grid_sel_recid);

                            var rec = grid.get(newRecid);
                            setToNewRAForm(rec.BID, rec.FlowID);
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
        },
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
                        { id: 'btnNotes', type: 'button', icon: 'far fa-sticky-note' },
                        { id: 'raState', type: 'html',
                            html: '<span style="padding: 0 10px">State: <span id="RAState">StateText</span></span>'
                        },
                        { id: 'stateAction', type: 'button', caption: 'Actions', icon: 'fas fa-pencil-alt'},
                        { id: 'bt3', type: 'spacer' },
                        { id: 'btnClose', type: 'button', icon: 'fas fa-times' }
                    ],
                    onClick: function (event) {
                        switch(event.target) {
                        case 'btnClose':
                            var no_callBack = function() { return false; },
                                yes_callBack = function() {
                                    w2ui.toplayout.hide('right',true);
                                    w2ui.applicantsGrid.render();
                                    app.raflow.activeFlowID = "";
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
                    onRefresh: function(event) {
                        console.log("Main Panel of newralayout Refreshed");
                        if(app.raflow.activeFlowID) {
                            var raflags = app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS;
                            renderRAStateInToolbar(raflags);
                        }
                    }
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
// renderRAStateInToolbar - it selects Rental Agreement State from the
//                          string list on basis of raFlags and displays
//                          it on the toolbar.
//
// @params
//   raFlags = FLAGS of rental agreement
//-----------------------------------------------------------------------
window.renderRAStateInToolbar = function(raFlags) {
    var raStateString = app.RAStates[parseInt(raFlags & 0xf)];
    $(w2ui.newraLayout_main_toolbar.box).find('#RAState').text(raStateString);
};
