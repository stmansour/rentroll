/*global
    getRAFlowAllParts, initRAFlowAJAX
*/

"use strict";

//-----------------------------------------------------------------------------
// setToNewRAForm -  enable the Rental Agreement form in toplayout.  Also, set
//                   the forms url and request data from the server
// @params
//   bid = business id (or the BUD)
//-----------------------------------------------------------------------------
window.setToNewRAForm = function (bid, FlowID) {

    if (FlowID.length < 1) {
        return false;
    }

    var f = w2ui.rentalagrForm;
    w2ui.toplayout.content('right', w2ui.newraLayout);
    w2ui.toplayout.show('right', true);
    w2ui.toplayout.sizeTo('right', app.WidestFormWidth);

    $.get('/webclient/html/raflowtmpl.html', function (data) {
        w2ui.newraLayout.content('main', data);
    });
    // f.url = '/v1/rentalagr/' + bid + '/' + FlowID;
    // f.request();
    w2ui.toplayout.render();

    // mark this flag as is this new record
    // record created already
    app.new_form_rec = false;

    // as new content will be loaded for this form
    // mark form dirty flag as false
    app.form_is_dirty = false;

    // click on first tab
    if (typeof f.tabs.name == "string") {
        f.tabs.click('tab1');
    }

    // load first slide
    $("#ra-form footer button#previous").addClass("disable");
    $(".ra-form-component").hide();
    $(".ra-form-component#dates").show();
    $("#progressbar li").removeClass("active");
    $("#progressbar li[data-target='#dates']").addClass("active");

    // set this flow id as in active
    app.raflow.activeFlowID = FlowID;

    // set BID in raflow settings
    app.raflow.BID = bid;

    // get all flow part related to this flow ID
    getRAFlowAllParts(app.raflow.activeFlowID);
};

window.buildNewRAElements = function () {
    // ------------------------------------------------------
    // rental agreement grid
    // ------------------------------------------------------
    $().w2grid({
        name: 'newrentalagrsGrid',
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
            {
                field: 'recid',
                hidden: true,
                caption: 'recid',
                size: '40px',
                sortable: true
            },
            {
                field: 'BID',
                caption: 'BID',
                hidden: true,
            },
            {
                field: 'BUD',
                caption: 'BUD',
                hidden: true,
            },
            {
                field: 'FlowID',
                caption: 'Flow ID',
                size: '100%',
                sortable: true
            },
        ],
        onRequest: function (event) {
            event.postData.cmd = "getAllFlows";
            event.postData.flow = "RA";
            console.log(event.postData);
        },
        onRefresh: function (event) {
            event.onComplete = function () {
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
                        var d = new Date();  // we'll use today for time-sensitive data
                        setToNewRAForm(rec.BID, rec.FlowID);
                    };

                // warn user if form content has been changed
                form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
            };
        },
        onAdd: function (/*event*/) {
            var yes_args = [this],
                no_args = [this],
                no_callBack = function (grid) {
                    grid.select(app.last.grid_sel_recid);
                    return false;
                },
                yes_callBack = function(grid, recid) {
                    initRAFlowAJAX()
                        .done(function (data, textStatus, jqXHR) {
                            var bid = getCurrentBID(),
                                bud = getBUDfromBID(bid);
                        grid.refresh();

                            console.log(data);

                            alert("refreshing the grid...");
                            grid.refresh();

                            app.last.grid_sel_recid = parseInt(newRecid);

                            // keep highlighting current row in any case
                            grid.select(app.last.grid_sel_recid);

                            var rec = grid.get(newRecid);
                            var d = new Date();  // we'll use today for time-sensitive data
                            setToNewRAForm(rec.BID, rec.FlowID);

                        })
                        .fail(function () {
                            console.log("error while creating new flow ID");
                        });

                };

            // warn user if form content has been changed
            form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
        },
    });

    // add date navigation toolbar for new rental agreement form
    addDateNavToToolbar('newrentalagrs');

    //------------------------------------------------------------------------
    //          Rental Agreement Details
    //------------------------------------------------------------------------
    $().w2layout({
        name: 'newraLayout',
        padding: 0,
        panels: [
            {type: 'left', hidden: true},
            {type: 'top', hidden: true},
            {
                type: 'main', size: '60%', resizable: true, style: app.pstyle,
                content: 'main',
                toolbar: {
                    items: [
                        {id: 'btnNotes', type: 'button', icon: 'far fa-sticky-note'},
                        {id: 'bt3', type: 'spacer'},
                        {id: 'btnClose', type: 'button', icon: 'fas fa-times'},
                    ],
                    onClick: function (event) {
                        switch(event.target) {
                        case 'btnClose':
                            var no_callBack = function() { return false; },
                                yes_callBack = function() {
                                    w2ui.toplayout.hide('right',true);
                                    w2ui.newrentalagrsGrid.render();
                                };
                            form_dirty_alert(yes_callBack, no_callBack);
                            break;
                        }
                    },
                }
            },
            {type: 'preview', hidden: true},
            {type: 'bottom', hidden: true},
            {type: 'right', hidden: true}
        ]
    });
};
