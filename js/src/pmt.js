"use strict";
function buildPaymentTypeElements() {
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

    //------------------------------------------------------------------------
    //          payment types form
    //------------------------------------------------------------------------
    $().w2form({
        name: 'pmtForm',
        header: 'Payment Type Detail',
        url: '/v1/pmts',
        style: 'border: 0px; background-color: transparent;display: block;',
        formURL: '/html/formpmt.html',
        fields: [
            { field: 'recid', type: 'int', required: false, html: { caption: 'recid', page: 0, column: 0 } },
            { field: 'BID', type: 'int', required: false, html: { caption: 'BID', page: 0, column: 0 }, hidden: true },
            { field: 'BUD', type: 'list', options: { items: app.businesses }, required: true, html: { caption: 'BUD', page: 0, column: 0 } },
            { field: 'PMTID', type: 'int', required: false, html: { caption: 'PMTID', page: 0, column: 0 }, hidden: true },
            { field: 'Name', type: 'text', required: true, html: { caption: 'Name', page: 0, column: 0 }, sortable: true },
            { field: 'Description', type: 'text', required: false, html: { caption: 'Description', page: 0, column: 0 }, sortable: true },
            { field: 'LastModTime', type: 'time', required: false, html: { caption: 'LastModTime', page: 0, column: 0 } },
            { field: 'LastModBy', type: 'int', required: false, html: { caption: 'LastModBy', page: 0, column: 0 } },
        ],
        toolbar: {
            items: [
                { id: 'btnNotes', type: 'button', icon: 'fa fa-sticky-note-o' },
                { id: 'bt3', type: 'spacer' },
                { id: 'btnClose', type: 'button', icon: 'fa fa-times' },
            ],
            onClick: function(event) {
                switch(event.target) {
                    case 'btnClose':
                        var no_callBack = function() { return false; },
                            yes_callBack = function() {
                                w2ui.toplayout.hide('right',true);
                                w2ui.pmtsGrid.render();
                            };
                        form_dirty_alert(yes_callBack, no_callBack);
                        break;
                }
            }
        },
        onError: function(event) {
            console.log(event);
        },
        actions: {
            saveadd: function() {
                var f = this,
                    grid = w2ui.pmtsGrid,
                    x = getCurrentBusiness(),
                    BID=parseInt(x.value),
                    BUD=getBUDfromBID(BID);

                // clean dirty flag of form
                app.form_is_dirty = false;
                // clear the grid select recid
                app.last.grid_sel_recid  =-1;

                // select none if you're going to add new record
                grid.selectNone();

                f.save({}, function (data) {
                    if (data.status == 'error') {
                        console.log('ERROR: '+ data.message);
                        return;
                    }

                    // JUST RENDER THE GRID ONLY
                    grid.render();

                    // add new empty record and just refresh the form, don't need to do CLEAR form
                    var y = new Date();
                    var record = {
                        recid: 0,
                        PMTID: 0,
                        BID: BID,
                        BUD: BUD,
                        Name: '',
                        Description: '',
                        LastModTime: y.toISOString(),
                        LastModBy: 0
                    };
                    f.record = record;
                    f.header = "Edit Payment Type (new)"; // have to provide header here, otherwise have to call refresh method twice to get this change in form
                    f.url = '/v1/pmts/' + BID+'/0';
                    f.refresh();
                });
            },
            save: function(/*target, data*/) {
                var f = this,
                    tgrid = w2ui.pmtsGrid;

                f.save({}, function (data) {
                    if (data.status == 'error') {
                        console.log('ERROR: '+ data.message);
                        return;
                    }
                    w2ui.toplayout.hide('right',true);
                    tgrid.render();
                });
            },
            delete: function() {

                var form = this;

                w2confirm(delete_confirm_options)
                .yes(function() {
                    var tgrid = w2ui.pmtsGrid;

                    tgrid.selectNone();

                    var params = {cmd: 'delete', formname: form.name, ID: form.record.PMTID };
                    var dat = JSON.stringify(params);

                    // delete Depository request
                    $.post(form.url, dat)
                    .done(function(data) {
                        if (data.status != "success") {
                            return;
                        }

                        // update payments list for a business
                        var x = getCurrentBusiness();
                        var BID=parseInt(x.value);
                        var BUD = getBUDfromBID(BID);

                        // remove this from app.pmtTypes[BUD] if successfully removed
                        var temp_list = [];
                        app.pmtTypes[BUD].forEach(function(item){
                            if (item.PMTID != form.record.PMTID) {
                                temp_list.push({PMTID: item.PMTID, Name: item.Name});
                            }
                        });
                        app.pmtTypes[BUD] = temp_list;

                        w2ui.toplayout.hide('right',true);
                        tgrid.render();
                    })
                    .fail(function(/*data*/){
                        console.log("Delete Payment failed.");
                    });
                })
                .no(function() {
                    return;
                });
            },
        },
        onRefresh: function(event) {
            event.onComplete = function() {
                var f = this,
                    header = "Edit Payment Type ({0})";

                formRefreshCallBack(f, "PMTID", header);
            };
        },
        onChange: function(event) {
            event.onComplete = function() {
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
        onResize: function(event) {
            event.onComplete = function() {
                // HACK: set the height of right panel of toplayout box div and form's box div
                // this is how w2ui set the content inside box of toplayout panel, and form's main('div.w2ui-form-box')
                var h = w2ui.toplayout.get("right").height;
                $(w2ui.toplayout.get("right").content.box).height(h);
                $(this.box).find("div.w2ui-form-box").height(h);
            };
        },
        onSubmit: function(target, data) {
            delete data.postData.record.LastModTime;
            delete data.postData.record.LastModBy;
            delete data.postData.record.CreateTS;
            delete data.postData.record.CreateBy;
            // modify form data for server request
            getFormSubmitData(data.postData.record);
        },
    });

}
