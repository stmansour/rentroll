//------------------------------------------------------------------------
//          payment types Grid
//------------------------------------------------------------------------
$().w2grid({
    name: 'pmtsGrid',
    url: '/v1/pmts',
    multiSelect: false,
    show: {
        toolbar        : true,
        footer         : true,
        toolbarAdd     : true,   // indicates if toolbar add new button is visible
        toolbarDelete  : false,   // indicates if toolbar delete button is visible
        toolbarSave    : false,   // indicates if toolbar save button is visible
        selectColumn    : false,
        expandColumn    : false,
        toolbarEdit     : false,
        toolbarSearch   : false,
        toolbarInput    : false,
        searchAll       : false,
        toolbarReload   : false,
        toolbarColumns  : false,
    },
    columns: [
        {field: 'recid',       hidden: true,  caption: 'recid',       size: '40px',  sortable: true},
        {field: 'PMTID',       hidden: true,  caption: 'PMTID',       size: '150px', sortable: true, style: 'text-align: center'},
        {field: 'BID',         hidden: true,  caption: 'BID',         size: '150px', sortable: true, style: 'text-align: center'},
        {field: 'Name',        hidden: false, caption: 'Name',        size: '150px', sortable: true, style: 'text-align: left'},
        {field: 'Description', hidden: false, caption: 'Description', size: '60%',   sortable: true, style: 'text-align: left'},
    ],
    onLoad: function(event) {
        if (event.xhr.status == 200) {
            if (typeof data == "undefined") {
                return;
            }

            // update payments list for a business
            var x = getCurrentBusiness(),
                BID=parseInt(x.value),
                BUD = getBUDfromBID(BID),
                pmtTypesList = [],
                data = JSON.parse(event.xhr.responseText);

            // prepare list of payment and push it to app.pmtTypes[BUD]
            data.records.forEach(function(pmtRec){
                pmtTypesList.push({PMTID: pmtRec.PMTID, Name: pmtRec.Name});
            });
            app.pmtTypes[BUD] = pmtTypesList;
        }
    },
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
                    // we need to get selected previous one selected record, in case of no answer
                    // in form dirty confirmation dialog
                    grid.select(app.last.grid_sel_recid);
                    return false;
                },
                yes_callBack = function(grid, recid) {
                    app.last.grid_sel_recid = parseInt(recid);

                    // keep highlighting current row in any case
                    grid.select(app.last.grid_sel_recid);

                    // get record
                    var rec = grid.get(recid);

                    // popup the dialog form
                    setToForm('pmtForm', '/v1/pmts/' + rec.BID + '/' + rec.PMTID, 400, true);
                };

            // form alert is content has been changed
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

                var x = getCurrentBusiness(),
                    BID=parseInt(x.value),
                    BUD = getBUDfromBID(BID),
                    record = {
                        recid: 0,
                        PMTID: 0,
                        BID: BID,
                        BUD: BUD,
                        Name: '',
                        Description: '',
                    };
                w2ui.pmtForm.record = record;
                // need to call refresh once before, already refreshin in setToForm
                w2ui.pmtForm.refresh();
                setToForm('pmtForm', '/v1/pmts/' + BID + '/0', 400);
            };

        // warn user if form content has been changed
        form_dirty_alert(yes_callBack, no_callBack, yes_args);
    },
});
