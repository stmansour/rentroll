//------------------------------------------------------------------------
//          rentable types Grid
//------------------------------------------------------------------------
$().w2grid({
    name: 'rtGrid',
    url: '/v1/rt',
    multiSelect: false,
    show: {
        header: false,
        toolbar: true,
        toolbarReload: false,
        toolbarColumns: false,
        toolbarSearch: true,
        toolbarAdd: true,
        toolbarDelete: false,
        toolbarSave: false,
        searchAll       : true,
        footer: true,
        lineNumbers: false,
        selectColumn: false,
        expandColumn: false
    },
    columns: [
        {field: 'recid', caption: 'recid', hidden: true},
        {field: 'RTID', caption: 'RTID', size: '50px', sortable: true},
        {field: 'Name', caption: 'Name', size: '150px', sortable: true},
        {field: 'Style', caption: 'Style', size: '100px', sortable: true},
        {field: 'BID', caption: 'BID', hidden: true},
        {field: 'BUD', caption: 'BUD', hidden: true},
        {field: 'RentCycle', caption: 'RentCycle', size: '75px', sortable: true,
            render: function (record/*, index, col_index*/) {
                var text = '';
                if (record) {
                    app.cycleFreq.forEach(function(itemText, itemIndex) {
                        if (record.RentCycle == itemIndex) {
                            text = itemText;
                            return false;
                        }
                    });
                }
                return text;
            },
        },
        {field: 'Proration', caption: 'Proration', size: '90px', sortable: true,
            render: function (record/*, index, col_index*/) {
                var text = '';
                if (record) {
                    app.cycleFreq.forEach(function(itemText, itemIndex) {
                        if (record.Proration == itemIndex) {
                            text = itemText;
                            return false;
                        }
                    });
                }
                return text;
            },
        },
        {field: 'GSRPC', caption: 'GSRPC', size: '65px', sortable: true,
            render: function (record/*, index, col_index*/) {
                var text = '';
                if (record) {
                    app.cycleFreq.forEach(function(itemText, itemIndex) {
                        if (record.GSRPC == itemIndex) {
                            text = itemText;
                            return false;
                        }
                    });
                }
                return text;
            },
        },
        {field: 'ManageToBudget', caption: 'ManageToBudget', size: '200px', sortable: true,
            render: function (record/*, index, col_index*/) {
                var text = '';
                if (record) {
                    app.manageToBudgetList.forEach(function(item) {
                        if (record.ManageToBudget == item.id) {
                            text = item.text;
                            return false;
                        }
                    });
                }
                return text;
            },
        },
        {field: 'LastModTime', caption: 'LastModTime', hidden: true},
        {field: 'LastModBy', caption: 'LastModBy', hidden: true},
        {field: 'RMRID', caption: 'RMRID', hidden: true},
        {field: 'MarketRate', caption: 'MarketRate', size: '100px', sortable: true, render: 'money'},
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

                    var rec = grid.get(recid);
                    console.log('rentableType form url: ' + '/v1/rt/' + rec.BID + '/' + rec.RTID);
                    setToForm('rentableTypeForm', '/v1/rt/' + rec.BID + '/' + rec.RTID, 400, true);
                };

            // warn user if content has been changed
            form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
        };
    },
    onAdd: function(/*event*/) {
        var yes_args = [this],
            no_callBack = function() { return false; },
            yes_callBack = function(grid) {
                // reset it
                app.last.grid_sel_recid = -1;
                grid.selectNone();

                var x = getCurrentBusiness(),
                    BID=parseInt(x.value),
                    BUD = getBUDfromBID(BID),
                    y = new Date();

                var record = {
                    recid: 0,
                    BID: BID,
                    BUD: BUD,
                    RTID: 0,
                    Style: "",
                    Name: "",
                    RentCycle: 0,
                    Proration: 0,
                    GSRPC: 0,
                    ManageToBudget: 0,
                    RMRID: 0,
                    MarketRate: 0.0,
                    LastModTime: y.toISOString(),
                    LastModBy: 0,
                };
                w2ui.rentableTypeForm.record = record;
                w2ui.rentableTypeForm.refresh();
                setToForm('rentableTypeForm', '/v1/rt/' + BID + '/0', 400);
            };

        // warn user if content has been changed of form
        form_dirty_alert(yes_callBack, no_callBack, yes_args);
    },
});
