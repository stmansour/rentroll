"use strict";
function getRentableInitRecord(BID, BUD){
    var y = new Date();
    return {
        recid: 0,
        BID: BID,
        BUD: BUD,
        RID: 0,
        RentableName: "",
        RARID: 0,
        RAID: 0,
        RARDtStart: w2uiDateControlString(y),
        RARDtStop: "1/1/9999",
        RTID: {id: 0, text: ''},
        RTRID: 0,
        RTRefDtStart: w2uiDateControlString(y),
        RTRefDtStop: "1/1/9999",
        RSID: 0,
        RentableStatus: "unknown",
        RSDtStart: w2uiDateControlString(y),
        RSDtStop: "1/1/9999",
        CurrentDate: y,
        AssignmentTime: 0,
    };
}

function buildRentableElements() {
    //------------------------------------------------------------------------
    //          rentablesGrid
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'rentablesGrid',
        url: '/v1/rentable',
        multiSelect: false,
        show: {
            header: false,
            toolbar: true,
            toolbarAdd: true,
            searches: true,
            footer: true,
            lineNumbers: false,
            selectColumn: false,
            expandColumn: false
        },
        columns: [
            {field: 'recid', caption: 'recid', size: '50px', hidden: true, sortable: true},
            {field: 'RID', caption: 'RID', size: '50px', sortable: true},
            {field: 'RentableName', caption: 'Rentable Name', size: '150px', sortable: true},
            // {field: 'AssignmentTime', caption: 'Assignment Time', size: '120px', sortable: true},
            {field: 'RTID', caption: 'Rentable Type ID', hidden: true, sortable: true},
            {field: 'RentableType', caption: 'Rentable Type', size: '200px', sortable: true},
            {field: 'RentableStatus', caption: 'Rentable <br>Status', size: '100px', sortable: true},
            {field: 'RARID', caption: 'RARID', hidden: true, sortable: true},
            {field: 'RAID', caption: 'RAID', size: '70px', sortable: true},
            {field: 'RentalAgreementStart', caption: 'Rental Agreement <br>Start', size: '120px', sortable: true},
            {field: 'RentalAgreementStop', caption: 'Rental Agreement <br>Stop', size: '120px', sortable: true},
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

                        var rec = grid.get(recid),
                            x = getCurrentBusiness(),
                            BID=parseInt(x.value),
                            BUD = getBUDfromBID(BID);

                        console.log('rentable form url: ' + '/v1/rentable/' + BID + '/' + rec.RID);

                        getRentableTypes(BUD)
                        .done(function(/*data*/){
                            setToForm('rentableForm', '/v1/rentable/' + BID + '/' + rec.RID, 700, true);
                        })
                        .fail(function(){
                            console.log("Failed to get rentable type list");
                        });
                    };

                // warn user if form content has been chagned
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
                        BUD = getBUDfromBID(BID);

                    var record = getRentableInitRecord(BID, BUD);

                    getRentableTypes(BUD)
                    .done(function(/*data*/){
                        w2ui.rentableForm.record = record;
                        w2ui.rentableForm.refresh();
                        setToForm('rentableForm', '/v1/rentable/' + BID + '/0', 700);
                    })
                    .fail(function(){
                        console.log("Failed to get rentable type list");
                    });
                };

            // warn user if form content has been chagned
            form_dirty_alert(yes_callBack, no_callBack, yes_args);
        },
    });

    //------------------------------------------------------------------------
    //          rentableForm
    //------------------------------------------------------------------------
    $().w2form({
        name: 'rentableForm',
        style: 'border: 0px; background-color: transparent;',
        header: app.sRentable + ' Detail',
        url: '/v1/rentable',
        formURL: '/html/formr.html',
        fields: [
            { field: 'recid', type: 'int', required: false, html: { page: 0, column: 0 } },
            { field: 'RID', type: 'int', required: false, html: { page: 0, column: 0 } },
            { field: 'BID', type: 'int', required: true, html: { page: 0, column: 0 } },
            { field: 'BUD', type: 'list', required: true, options: {items: app.businesses}, html: { page: 0, column: 0 } },
            { field: 'RARID', type: 'int', required: false, html: { page: 0, column: 0 } },
            { field: 'RAID', type: 'int', required: false, html: { page: 0, column: 0 } },
            { field: 'RARDtStart', type: 'date', required: false, html: { page: 0, column: 0 } },
            { field: 'RARDtStop', type: 'date', required: false, html: { page: 0, column: 0 } },
            { field: 'RentableName', type: 'text', required: true, html: { page: 0, column: 0 } },
            { field: 'RTID', type: 'list', required: true, html: { page: 0, column: 0 }, options: { items: [], selected: {}, maxDropHeight: 200 } },
            // { field: 'RTID', type: 'int', required: false },
            { field: 'RTRID', type: 'int', required: false, html: { page: 0, column: 0 } },
            { field: 'RTRefDtStart', type: 'date', required: true, html: { page: 0, column: 0 } },
            { field: 'RTRefDtStop', type: 'date', required: true, html: { page: 0, column: 0 } },
            { field: 'RSID', type: 'int', required: false, html: { page: 0, column: 0 } },
            { field: 'RentableStatus', type: 'list', options: {items: app.rentableStatusList}, required: true, html: { page: 0, column: 0 } },
            { field: 'RSDtStart', type: 'date', required: true, html: { page: 0, column: 0 } },
            { field: 'RSDtStop', type: 'date', required: true, html: { page: 0, column: 0 } },
            { field: 'CurrentDate', type: 'date', required: false },
            { field: 'AssignmentTime', type: 'list', required: false, html: { page: 0, column: 0 } },
            { field: 'LastModTime',          type: 'hidden', required: false },
            { field: 'LastModBy',          type: 'hidden', required: false },
            { field: 'CreateTS',          type: 'hidden', required: false },
            { field: 'CreateBy',          type: 'hidden', required: false },
        ],
        toolbar: {
            items: [
                { id: 'btnNotes', type: 'button', icon: 'fa fa-sticky-note-o' },
                { id: 'bt3', type: 'spacer' },
                { id: 'btnClose', type: 'button', icon: 'fa fa-times' },
            ],
            onClick: function (event) {
                switch(event.target) {
                case 'btnClose':
                    var no_callBack = function() { return false; },
                        yes_callBack = function() {
                            w2ui.toplayout.hide('right',true);
                            w2ui.rentablesGrid.render();
                        };
                    form_dirty_alert(yes_callBack, no_callBack);
                    break;
                }
            },
        },
        actions: {
            saveadd: function() {
                var f = this,
                    grid = w2ui.rentablesGrid,
                    r = f.record,
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

                    getRentableTypes(BUD)
                    .done(function(/*data*/){
                        w2ui.rentableForm.record = record;
                        w2ui.rentableForm.refresh();
                        setToForm('rentableForm', '/v1/rentable/' + BID + '/0', 700);
                    })
                    .fail(function(){
                        console.log("Failed to get rentable type list");
                    });
                    // JUST RENDER THE GRID ONLY
                    grid.render();

                    var record = getRentableInitRecord(BID, BUD);
                    f.record = record;
                    f.header = "Edit {0} ({1}) as of {2}".format(app.sRentable, "new", w2uiDateControlString(r.CurrentDate));
                    f.url = '/v1/rentable/' + BID+'/0';
                    f.refresh();
                });
            },
            save: function () {
                //var obj = this;
                var tgrid = w2ui.rentablesGrid;
                tgrid.selectNone();
                this.save({}, function (data) {
                    if (data.status == 'error') {
                        console.log('ERROR: '+ data.message);
                        return;
                    }
                    w2ui.toplayout.hide('right',true);
                    tgrid.render();
                });

            },
        },
        onSubmit: function(target, data){
            // server request form data
            delete data.postData.record.RARID;
            delete data.postData.record.RAID;
            delete data.postData.record.RARDtStop;
            delete data.postData.record.RARDtStart;
            delete data.postData.record.LastModTime;
            delete data.postData.record.LastModBy;
            delete data.postData.record.CreateTS;
            delete data.postData.record.CreateBy;
            getFormSubmitData(data.postData.record);
        },
        onRefresh: function(event) {
            event.onComplete = function() {
                var f = this,
                    r = f.record,
                    x = getCurrentBusiness(),
                    BID=parseInt(x.value),
                    BUD = getBUDfromBID(BID),
                    header = "";

                // custom header, not common one!!
                if (r.RID) {
                    header = "Edit {0} - {1} ({2}) as of {3}".format(app.sRentable, r.RentableName, r.RID, r.CurrentDate);
                } else {
                    header = "Edit {0} ({1}) as of {2}".format(app.sRentable, "new", w2uiDateControlString(r.CurrentDate));
                }

                // assignmentTime selected and items for w2field
                var assignmentItems = [], assignSelected = {};
                app.assignmentTimeList.forEach(function(item, index) {
                    if (index == r.AssignmentTime) {
                        assignSelected = { id: index, text: item };
                    }
                    assignmentItems.push({ id: index, text: item });
                });

                f.get("AssignmentTime").options.items = assignmentItems;
                f.get("AssignmentTime").options.selected = assignSelected;
                f.get("RTID").options.items = app.rt_list[BUD];
                f.get("RTID").options.selected = {id: r.RTID, text: r.Name};

                formRefreshCallBack(f, "RID", header);
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
        }
    });

}
