//------------------------------------------------------------------------
//          accountsGrid
//------------------------------------------------------------------------
$().w2grid({
    name: 'accountsGrid',
    url: '/v1/GLAccounts',
    multiSelect: false,
    show: {
        toolbar         : true,
        footer          : true,
        toolbarAdd      : true,   // indicates if toolbar add new button is visible
        toolbarDelete   : false,   // indicates if toolbar delete button is visible
        toolbarSave     : false,   // indicates if toolbar save button is visible
        selectColumn    : false,
        expandColumn    : false,
        toolbarEdit     : false,
        toolbarSearch   : false,
        toolbarInput    : true,
        searchAll       : false,
        toolbarReload   : true,
        toolbarColumns  : true,

    },
    columns: [
        {field: 'recid',    caption: 'recid',     size: '80px', sortable: false, hidden: true},
        {field: 'LID',      caption: 'LID',       size: '50px', sortable: false, hidden: true},
        {field: 'PLID',     caption: 'PLID',      size: '50px', sortable: false, hidden: true},
        {field: 'BID',      caption: 'BID',       size: '40px', sortable: false, hidden: true},
        {field: 'GLNumber', caption: 'GL Number', size: '100px', sortable: true},
        {field: 'Name',     caption: 'Name',      size: '450px', sortable: true},
        {field: 'Status',   caption: 'Status',    size: '90px', sortable: true,
            render: function (record/*, index, col_index*/) {
                if (typeof record === "undefined") {
                    return;
                }
                var html = '';
                for (var i=0; i < app.account_stuff.statusList.length; i++) {
                    if (record.Status == app.account_stuff.statusList[i].id) {
                        html = app.account_stuff.statusList[i].text;
                    }
                }
                return html;
            },
        },
        {field: 'AcctType', caption: 'AcctType', size: '150px', sortable: true},
    ],
    onRefresh: function(event) {
        event.onComplete = function() {
            if (app.active_grid == this.name) {
                if (app.new_form_rec) {
                    this.selectNone();
                }
                else{
                    this.select(app.last.grid_sel_recid);
                }
            }
        };
    },
    onClick: function(event) {
        event.onComplete = function() {
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

                    var x = getCurrentBusiness();
                    var BID=parseInt(x.value);
                    var BUD = getBUDfromBID(BID);
                    var rec = grid.get(recid);

                    // get latest gl accounts first
                    getParentAccounts(rec.BID, rec.LID)
                    .done(function(data) {
                        if (data.status != 'success') {
                            w2ui.accountForm.message(data.message);
                            w2ui.toplayout.content('right', w2ui.accountForm);
                            w2ui.toplayout.show('right', true);
                            w2ui.toplayout.sizeTo('right', 700);
                            return;
                        }
                        else {
                            w2ui.accountForm.get("PLID").options.items = app.parent_accounts[BUD];
                            setToForm('accountForm', '/v1/account/' + rec.BID + '/' + rec.LID, 400, true);
                        }
                    })
                    .fail(function(/*data*/) {
                        console.log("unable to get gl accounts list");
                        return;
                    });
                };

            // warn user if form content has been changed
            form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
        };
    },
    onAdd: function (/*event*/) {
        var yes_args = [this],
            no_callBack = function() {
                return false;
            },
            yes_callBack = function(grid) {
                // reset grid sel recid
                app.last.grid_sel_recid = -1;
                grid.selectNone();


                var x = getCurrentBusiness();
                var BID=parseInt(x.value);
                var BUD = getBUDfromBID(BID);

                getParentAccounts(BID, 0)
               .done(function(data) {
                    if (data.status != 'success') {
                        w2ui.accountForm.message(data.message);
                        w2ui.toplayout.content('right', w2ui.accountForm);
                        w2ui.toplayout.show('right', true);
                        w2ui.toplayout.sizeTo('right', 700);
                        return;
                    } else {
                        var y = new Date(),
                            record = {
                                recid: 0,
                                LID: 0,
                                PLID: 0,
                                BID: BID,
                                BUD: BUD,
                                RAID: 0,
                                TCID: 0,
                                GLNumber: '',
                                Status: 2,
                                Name: '',
                                AcctType: '',
                                AllowPost: 1,
                                Description: '',
                                LastModTime: y.toISOString(),
                                LastModBy: 0,
                            };

                        w2ui.accountForm.get("PLID").options.items = app.parent_accounts[BUD];
                        w2ui.accountForm.record = record;
                        // NOTE: even inside "setToForm", form is refreshing but header isn't updating,
                        // so need to call here once before
                        w2ui.accountForm.refresh();
                        // now show the form
                        setToForm('accountForm', '/v1/account/' + BID + '/0', 400);
                    }
                })
                .fail( function() {
                    console.log('unable to get GLAccounts');
                    return;
                });
            };

            // warn user if form content has been changed
            form_dirty_alert(yes_callBack, no_callBack, yes_args);
    },
});
