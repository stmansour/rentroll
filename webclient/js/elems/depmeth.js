/*global
    parseInt, w2ui, app, getDepMethInitRecord
*/

"use strict";

window.getDepMethInitRecord = function (BID, BUD){
    return {
        recid: 0,
        DPMID: 0,
        BID: BID,
        BUD: BUD,
        Name: '',
    };
};

window.buildDepositMethodElements = function () {
    //------------------------------------------------------------------------
    //          Deposit Methods Grid
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'depmethGrid',
        url: '/v1/depmeth',
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
            {field: 'recid',      hidden: true,  caption: 'recid',       size: '40px',  sortable: true},
            {field: 'DPMID',      hidden: true,  caption: 'DEPID',       size: '150px', sortable: true},
            {field: 'BID',        hidden: true,  caption: 'BID',         size: '150px', sortable: true},
            {field: 'Name',       hidden: false, caption: 'Name',        size: '20%',   sortable: true},
            {field: 'LastModTime',hidden: true,  caption: 'LastModTime', size: '20%',   sortable: true},
            {field: 'LastModBy',  hidden: true,  caption: 'LastModBy',   size: '150px', sortable: true},
            {field: 'CreateTS',   hidden: true,  caption: 'CreateTS',    size: '20%',   sortable: true},
            {field: 'CreateBy',   hidden: true,  caption: 'CreateBy',    size: '150px', sortable: true},
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
	                    setToForm('depmethForm', '/v1/depmeth/' + rec.BID + '/' + rec.DPMID, 400, true);
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

                    // Insert an empty record...
                    var x = getCurrentBusiness();
                    var BID=parseInt(x.value);
                    var BUD = getBUDfromBID(BID);
	                var record = getDepMethInitRecord(BID, BUD);
                    w2ui.depmethForm.record = record;
                    setToForm('depmethForm', '/v1/depmeth/' + BID + '/0', 400);

                };  // yes callback

            // warn user if form content has been changed
            form_dirty_alert(yes_callBack, no_callBack, yes_args);
        },
    });

    //------------------------------------------------------------------------
    //          Deposit Methods Form
    //------------------------------------------------------------------------
    $().w2form({
        name: 'depmethForm',
        header: 'Deposit Method Detail',
        url: '/v1/depmeth',
        style: 'border: 0px; background-color: transparent;display: block;',
        formURL: '/webclient/html/formdepmeth.html',
        fields: [
            { field: 'recid', type: 'int', required: false, html: { caption: 'recid', page: 0, column: 0 } },
            { field: 'BID', type: 'int', required: false, html: { caption: 'BID', page: 0, column: 0 }, hidden: true },
            { field: 'BUD', type: 'list', options: { items: app.businesses }, required: true, html: { caption: 'BUD', page: 0, column: 0 } },
            { field: 'DPMID', type: 'int', required: false, html: { caption: 'DPMID', page: 0, column: 0 }, hidden: true },
            { field: 'Name', type: 'text', required: true, html: { caption: 'Name', page: 0, column: 0 }, sortable: true },
            { field: 'LastModTime', type: 'time', required: false, html: { caption: 'LastModTime', page: 0, column: 0 } },
            { field: 'LastModBy', type: 'int', required: false, html: { caption: 'LastModBy', page: 0, column: 0 } },
        ],
        toolbar: {
            items: [
                { id: 'btnNotes', type: 'button', icon: 'far fa-sticky-note' },
                { id: 'bt3', type: 'spacer' },
                { id: 'btnClose', type: 'button', icon: 'fas fa-times' },
            ],
            onClick: function(event) {
                switch(event.target) {
                    case 'btnClose':
                        var no_callBack = function() { return false; },
                            yes_callBack = function() {
                                w2ui.toplayout.hide('right',true);
                                w2ui.depmethGrid.render();
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
                    grid = w2ui.depmethGrid,
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
                    var record = getDepMethInitRecord(BID, BUD);
                    f.record = record;
                    f.header = "Edit Deposit Method (new)"; // have to provide header here, otherwise have to call refresh method twice to get this change in form
                    f.url = '/v1/depmeth/' + BID+'/0';
                    f.refresh();
                });
            },
            save: function(/*target, data*/) {
                var f = this,
                    tgrid = w2ui.depmethGrid;

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

                var form = this,
                    x = getCurrentBusiness(),
                    BID=parseInt(x.value),
                    BUD = getBUDfromBID(BID);

                w2confirm(delete_confirm_options)
                .yes(function() {
                    var tgrid = w2ui.depmethGrid;
                    var params = {cmd: 'delete', formname: form.name, ID: form.record.DPMID };
                    var dat = JSON.stringify(params);

                    // delete Depository request
                    $.post(form.url, dat, null, "json")
                    .done(function(data) {
                        if (data.status === "error") {
                            form.error(w2utils.lang(data.message));
                            return;
                        }

                        // remove this from app.depmeth[BUD] if successfully removed
                        var temp_list = [];
                        app.depmeth[BUD].forEach(function(item){
                            if (item.DPMID != form.record.DPMID) {
                                temp_list.push({DPMID: item.DPMID, Name: item.Name});
                            }
                        });
                        app.depmeth[BUD] = temp_list;

                        w2ui.toplayout.hide('right',true);
                        tgrid.remove(app.last.grid_sel_recid);
                        tgrid.render();
                    })
                    .fail(function(/*data*/){
                        form.error("Delete Payment failed.");
                        return;
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
                    header = "Edit Deposit Method ({0})";

                formRefreshCallBack(f, "DPMID", header);
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
        onSubmit: function(target, data) {
            delete data.postData.record.LastModTime;
            delete data.postData.record.LastModBy;
            delete data.postData.record.CreateTS;
            delete data.postData.record.CreateBy;
            // modify form data for server request
            getFormSubmitData(data.postData.record);
        },
    });
};
