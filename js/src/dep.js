//------------------------------------------------------------------------
//          depository Grid
//------------------------------------------------------------------------
$().w2grid({
    name: 'depGrid',
    url: '/v1/dep',
    multiSelect: false,
    show: {
        toolbar        : true,
        footer         : true,
        toolbarAdd     : true,   // indicates if toolbar add new button is visible
        toolbarDelete  : false,   // indicates if toolbar delete button is visible
        toolbarSave    : false,   // indicates if toolbar save button is visible
        selectColumn   : false,
        expandColumn   : false,
        toolbarEdit    : false,
        toolbarSearch  : false,
        toolbarInput   : false,
        searchAll      : false,
        toolbarReload  : false,
        toolbarColumns : false,
    },
    columns: [
        {field: 'recid', hidden: true, caption: 'recid', size: '40px', sortable: true},
        {field: 'DEPID', hidden: true, caption: 'DEPID', size: '150px', sortable: true, style: 'text-align: center'},
        {field: 'BID', hidden: true, caption: 'BID', size: '150px', sortable: true, style: 'text-align: center'},
        {field: 'BUD', hidden: true, caption: 'BUD', size: '150px', sortable: true, style: 'text-align: center'},
        {field: 'LID', hidden: true, caption: 'LID', size: '150px', sortable: true, style: 'text-align: center'},
        {field: 'AccountNo', hidden: false, caption: 'Account Number', size: '150px', sortable: true, style: 'text-align: right'},
        {field: 'Name', hidden: false, caption: 'Name', size: '20%', sortable: true, style: 'text-align: left'},
        {field: 'LdgrName', hidden: false, caption: 'GL Acct Name', size: '20%', sortable: true, style: 'text-align: left'},
        {field: 'GLNumber', hidden: false, caption: 'GL Number', size: '150px', sortable: true, style: 'text-align: right'},
    ],
    onRefresh: function(event) {
        event.onComplete = function() {
            var sel_recid = parseInt(this.last.sel_recid);
            if (app.active_grid == this.name && sel_recid > -1) {
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

                    // get selected element as multiSelect is false
                    var rec = grid.get(recid),
                        x = getCurrentBusiness(),
                        BID=parseInt(x.value),
                        BUD = getBUDfromBID(BID);

                    // get gl account list
                    getAccountsList(rec.BID)
                    .done(function(data){
                        if (data.status != 'success') {
                            w2ui.depForm.message(data.message);
                        } else {
                            var gl_selected = {id: rec.LID, text: rec.GLNumber+" "+"(" + rec.LdgrName + ")"};
                            // get gl accounts for BUD
                            w2ui.depForm.get('LID').options.items = app.gl_accounts[BUD];
                            w2ui.depForm.get('LID').options.selected = gl_selected;
                        }
                        setToForm('depForm', '/v1/dep/' + rec.BID + '/' + rec.DEPID, 400, true);
                    })
                    .fail(function(/*data*/){
                        console.log("Failed to get glAccountList");
                    });
                };

            // warn user if form content has been changed
            form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
        };
    },
    onAdd: function (/*event*/) {
        var yes_args = [this],
            no_callBack = function() { return false; },
            yes_callBack = function(grid) {
                // reset it
                app.last.grid_sel_recid = -1;
                grid.selectNone();

                // Insert an empty record...
                var x = getCurrentBusiness(),
                    BID=parseInt(x.value),
                    BUD = getBUDfromBID(BID),
                    f = w2ui.depForm;

                // get latest gl accounts first
                getAccountsList(BID)
                .done(function(data) {

                    if (data.status != 'success') {
                        f.message(data.message);
                        w2ui.toplayout.content('right', f);
                        w2ui.toplayout.show('right', true);
                        w2ui.toplayout.sizeTo('right', 700);
                        return;
                    } else {
                        var record = {
                            recid: 0,
                            DEPID: 0,
                            BID: BID,
                            BUD: BUD,
                            LID: 0,
                            Name: "",
                            AccountNo: "",
                        };

                        var gl_accounts_pre_selected = {id: 0, text: " -- Select GL Account -- "};
                        var gl_accounts_items = [gl_accounts_pre_selected];
                        // get gl account list for BUD from `gl_accounts` key of `app`
                        gl_accounts_items = gl_accounts_items.concat(app.gl_accounts[BUD]);

                        f.get('LID').options.items = gl_accounts_items;
                        f.get('LID').options.selected = gl_accounts_pre_selected;
                        f.record = record;
                        f.refresh();
                        setToForm('depForm', '/v1/dep/' + BID + '/0', 400);
                    }
                })
                .fail( function() {
                    console.log('unable to get GLAccounts');
                    return;
                 });
            };  // yes callback

        // warn user if form content has been changed
        form_dirty_alert(yes_callBack, no_callBack, yes_args);
    },
});
