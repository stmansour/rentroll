/*global
    w2ui, app, $, console, setToForm, form_dirty_alert, formRefreshCallBack, formRecDiffer,
    getFormSubmitData, w2confirm, delete_confirm_options, getBUDfromBID, getCurrentBusiness,
    addDateNavToToolbar,
*/
"use strict";
function getRTInitRecord(BID, BUD){
    var y = new Date();
    return {
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
}

function buildRentableTypeElements() {

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
            {field: 'LastModBy',   caption: 'LastModBy',   hidden: true},
            {field: 'RMRID',       caption: 'RMRID',       hidden: true},
            {field: 'MarketRate',  caption: 'MarketRate',  size: '100px', sortable: true, render: 'money'},
            {field: 'DtStart',     caption: 'DtStart',     size: '100px', sortable: true, style: 'text-align: right'},
            {field: 'DtStop',      caption: 'DtStop',      size: '100px', sortable: true, style: 'text-align: right'},
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
                        setRTLayout(rec.BID, rec.RTID);
                        // setToForm('rtForm', '/v1/rt/' + rec.BID + '/' + rec.RTID, 400, true);
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
                        BUD = getBUDfromBID(BID);

                    var record = getRTInitRecord(BID, BUD);
                    w2ui.rtForm.record = record;
                    w2ui.rtForm.refresh();
                    setRTLayout(BID, 0);
                    // setToForm('rtForm', '/v1/rt/' + BID + '/0', 400);
                };

            // warn user if content has been changed of form
            form_dirty_alert(yes_callBack, no_callBack, yes_args);
        },
    });

    addDateNavToToolbar('rt');

    //------------------------------------------------------------------------
    //          rentable Type Form
    //------------------------------------------------------------------------
    $().w2form({
        name: 'rtForm',
        style: 'border: 1px solid silver; background-color: transparent;',
        // header: app.sRentableType + ' Detail',
        url: '/v1/rentabletypes',
        formURL: '/webclient/html/formrt.html',
        fields: [
            { field: 'recid', type: 'int', required: false, html: { page: 0, column: 0 } },
            { field: 'RTID', type: 'int', required: true, html: { page: 0, column: 0 } },
            { field: 'BID', type: 'int', required: true, html: { page: 0, column: 0 } },
            { field: 'BUD', type: 'list', options: {items: app.businesses}, required: true, html: { page: 0, column: 0 } },
            { field: 'Style', type: 'text', required: true, html: { page: 0, column: 0 } },
            { field: 'Name', type: 'text', required: true, html: { page: 0, column: 0 } },
            { field: 'RentCycle', type: 'list', options: {items: app.cycleFreq, selected: {}}, required: true, html: { page: 0, column: 0 } },
            { field: 'Proration', type: 'list', options: {items: app.cycleFreq, selected: {}}, required: true, html: { page: 0, column: 0 } },
            { field: 'GSRPC', type: 'list', options: {items: app.cycleFreq, selected: {}}, required: true, html: { page: 0, column: 0 } },
            { field: 'ManageToBudget', type: 'list', options: {items: app.manageToBudgetList, selected: {}}, required: true, html: { page: 0, column: 0 } },
            { field: 'RMRID', type: 'int', required: true, html: { page: 0, column: 0 } },
            { field: 'LastModTime', type: 'time', required: false, html: { page: 0, column: 0 } },
            { field: 'LastModBy', type: 'int', required: false, html: { page: 0, column: 0 } },
            { field: 'CreateTS', type: 'time', required: false, html: { page: 0, column: 0 } },
            { field: 'CreateBy', type: 'int', required: false, html: { page: 0, column: 0 } },
        ],
        onValidate: function(event) {
            if (this.record.ManageToBudget.id === 1) {
                var grid = w2ui.rmrGrid;
                if (grid.records.length < 1) {
                    event.errors.push({
                        field: this.get('ManageToBudget'),
                        error: 'At least one MarketRate should be exist when Mange To Budget is Yes.\n Please checkout MarketRates tab!'
                    });
                }
            }
            if (this.record.Style === "") {
                event.errors.push({
                    field: this.get('Style'),
                    error: 'Style cannot be blank'
                });
            }
            if (this.record.Name === "") {
                event.errors.push({
                    field: this.get('Name'),
                    error: 'Name cannot be blank'
                });
            }
        },
        onSubmit: function(target, data){
            delete data.postData.record.LastModTime;
            delete data.postData.record.LastModBy;
            delete data.postData.record.CreateTS;
            delete data.postData.record.CreateBy;
            delete data.postData.record.MarketRate;
            // server request form data
            getFormSubmitData(data.postData.record);
        },
        onRefresh: function(event) {
            event.onComplete = function() {
                var f = this,
                    r = f.record,
                    header = "Edit Rentable Type ({0})";

                // dropdown list items and selected variables
                var rentCycleSel = {}, prorationSel = {}, gsrpcSel = {},
                    manageToBudgetSel = {}, cycleFreqItems = [];

                // select value for rentcycle, proration, gsrpc
                app.cycleFreq.forEach(function(itemText, itemIndex) {
                    if (itemIndex == r.RentCycle) {
                        rentCycleSel = { id: itemIndex, text: itemText };
                    }
                    if (itemIndex == r.Proration) {
                        prorationSel = { id: itemIndex, text: itemText };
                    }
                    if (itemIndex == r.GSRPC) {
                        gsrpcSel = { id: itemIndex, text: itemText };
                    }
                    cycleFreqItems.push({ id: itemIndex, text: itemText });
                });

                // select value for manage to budget
                app.manageToBudgetList.forEach(function(item) {
                    if (item.id == r.ManageToBudget) {
                        manageToBudgetSel = {id: item.id, text: item.text};
                    }
                });

                // fill the field with values
                f.get("RentCycle").options.items = cycleFreqItems;
                f.get("RentCycle").options.selected = rentCycleSel;
                f.get("Proration").options.items = cycleFreqItems;
                f.get("Proration").options.selected = prorationSel;
                f.get("GSRPC").options.items = cycleFreqItems;
                f.get("GSRPC").options.selected = gsrpcSel;
                f.get("ManageToBudget").options.items = app.manageToBudgetList;
                f.get("ManageToBudget").options.selected = manageToBudgetSel;

                formRefreshCallBack(f, "RTID", header, false);
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
    });

    //------------------------------------------------------------------------
    //          rtFormBtns
    //------------------------------------------------------------------------
    $().w2form({
        name: 'rtFormBtns',
        style: 'border: 0px; background-color: transparent;',
        formURL: '/webclient/html/formrtbtns.html',
        url: '',
        fields: [],
        actions: {
            save: function() {
                var rtG = w2ui.rtGrid,
                    rmrG = w2ui.rmrGrid,
                    rtF = w2ui.rtForm,
                    x = getCurrentBusiness(),
                    BID=parseInt(x.value);

                // unselect record from
                rtG.selectNone();

                // hit save
                rtF.save({}, function (data) {
                    if (data.status === 'error') {
                        console.log('ERROR: '+ data.message);
                        return;
                    }

                    // in case if record is new then we've to update RTID that saved on server side
                    rtF.record.RTID = data.recid;

                    // update RTID in grid records
                    for (var i = 0; i < rmrG.records.length; i++) {
                        rmrG.records[i].RTID = rtF.record.RTID;
                    }

                    // now set the url of market Rate grid so that it can save the record on server side
                    rmrG.url = '/v1/rmr/' + BID + '/' + rtF.record.RTID;
                    rmrG.save();
                    rmrG.url = ""; // after save, remove it

                    w2ui.toplayout.hide('right',true);
                    rtG.render();
                });
            },
            saveadd: function() {
                var rtF = w2ui.rtForm,
                    rtG = w2ui.rtGrid,
                    rmrG = w2ui.rmrGrid,
                    x = getCurrentBusiness(),
                    BID=parseInt(x.value),
                    BUD=getBUDfromBID(BID);

                // clean dirty flag of form
                app.form_is_dirty = false;

                // clear the grid select recid
                app.last.grid_sel_recid  =-1;

                // select none if you're going to add new record
                rtG.selectNone();

                rtF.save({}, function (data) {
                    if (data.status == 'error') {
                        console.log('ERROR: '+ data.message);
                        return;
                    }

                    // now set the url of market Rate grid so that it can save the record on server side
                    rmrG.url = '/v1/rmr/' + BID + '/' + rtF.record.RTID;
                    rmrG.save();
                    rmrG.url = ""; // after save, remove it
                    // clear grid as we're going to add new Form
                    rmrG.clear();

                    // dropdown list items and selected variables
                    var rentCycleSel = {}, prorationSel = {}, gsrpcSel = {},
                        manageToBudgetSel = {}, cycleFreqItems = [];

                    // select value for rentcycle, proration, gsrpc
                    app.cycleFreq.forEach(function(itemText, itemIndex) {
                        if (itemIndex == rtF.record.RentCycle) {
                            rentCycleSel = { id: itemIndex, text: itemText };
                        }
                        if (itemIndex == rtF.record.Proration) {
                            prorationSel = { id: itemIndex, text: itemText };
                        }
                        if (itemIndex == rtF.record.GSRPC) {
                            gsrpcSel = { id: itemIndex, text: itemText };
                        }
                        cycleFreqItems.push({ id: itemIndex, text: itemText });
                    });

                    // select value for manage to budget
                    app.manageToBudgetList.forEach(function(item) {
                        if (item.id == rtF.record.ManageToBudget) {
                            manageToBudgetSel = {id: item.id, text: item.text};
                        }
                    });

                    rtF.get("ManageToBudget").options.items = app.manageToBudgetList;
                    rtF.get("ManageToBudget").options.selected = manageToBudgetSel[0];
                    rtF.get("RentCycle").options.items = cycleFreqItems;
                    rtF.get("RentCycle").options.selected = rentCycleSel[0];
                    rtF.get("Proration").options.items = cycleFreqItems;
                    rtF.get("Proration").options.selected = prorationSel[0];
                    rtF.get("GSRPC").options.items = cycleFreqItems;
                    rtF.get("GSRPC").options.selected = gsrpcSel[0];

                    // JUST RENDER THE GRID ONLY
                    rtG.render();

                    var record = getRTInitRecord(BID, BUD);
                    rtF.record = record;
                    rtF.header = "Edit Rentable Type (new)"; // have to provide header here, otherwise have to call refresh method twice to get this change in form
                    rtF.url = '/v1/rt/' + BID+'/0';
                    rtF.refresh();
                });
            },
            delete: function() {
                var rtF = w2ui.rtForm;

                // confirm before delete
                w2confirm(delete_confirm_options)
                .yes(function() {
                    var rtG = w2ui.rtGrid;
                    var params = {cmd: 'delete', formname: rtF.name, ID: rtF.record.RTID };
                    var dat = JSON.stringify(params);

                    // delete Depository request
                    $.post(rtF.url, dat, null, "json")
                    .done(function(data) {
                        if (data.status === "error") {
                            return;
                        }

                        w2ui.toplayout.hide('right',true);
                        rtG.remove(app.last.grid_sel_recid);
                        rtG.render();
                    })
                    .fail(function(/*data*/){
                        rtF.error("Delete Payment failed.");
                        return;
                    });
                })
                .no(function() {
                    return;
                });
            },
         },
     });

    //------------------------------------------------------------------------
    //          rentable Type Detailed Layout
    //------------------------------------------------------------------------
    $().w2layout({
        name: 'rtDetailLayout',
        panels: [
            {
                type: 'top',
                size: 35,
                style: 'border: 1px solid silver;',
                content: "",
                toolbar: {
                    style: "height: 35px; background-color: #eee; border: 0px;",
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
                                    w2ui.rtGrid.render();
                                };
                            form_dirty_alert(yes_callBack, no_callBack);
                            break;
                        }
                    },
                },
            },
            {
                type: 'main',
                overflow: "hidden",
                style: 'background-color: white; border: 1px solid silver; padding: 0px;',
                tabs: {
                    style: "padding-top: 10px;",
                    active: 'rtForm',
                    tabs: [
                        { id: 'rtForm', caption: 'Rentable Type Detail' },
                        { id: 'rmrGrid', caption: 'Market Rates' },
                    ],
                    onClick: function (event) {
                        if (event.target === "rmrGrid") {
                            w2ui.rtDetailLayout.html('main', w2ui.rmrGrid);
                        }
                        if (event.target === "rtForm") {
                            w2ui.rtDetailLayout.html('main', w2ui.rtForm);
                        }
                    }
                }
            },
            {
                type: 'bottom',
                size: 60,
                // style: 'background-color: white;  border-top: 1px solid silver; text-align: center; padding: 15px;',
            },
        ],
    });

    //------------------------------------------------------------------------
    //          rentable Market Rates Grid
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'rmrGrid',
        style: 'padding: 0px',
        show: {
            header: false,
            toolbar: true,
            toolbarReload: false,
            toolbarColumns: false,
            toolbarSearch: true,
            toolbarAdd: true,
            toolbarDelete: true,
            toolbarSave: false,
            searchAll       : true,
            footer: true,
            lineNumbers: false,
            selectColumn: false,
            expandColumn: false
        },
        columns: [
            {field: 'recid', caption: 'recid', hidden: true},
            {field: 'RMRID', caption: 'RMRID', size: '150px', sortable: true},
            {field: 'RTID', caption: 'RTID', size: '150px', hidden: true},
            {field: 'BID', caption: 'BID', hidden: true},
            {field: 'BUD', caption: 'BUD', hidden: true},
            {field: 'MarketRate',  caption: 'MarketRate',  size: '100px', sortable: true, render: 'money', editable: {type: 'money'} },
            {field: 'DtStart',     caption: 'DtStart', size: "50%",    sortable: true, style: 'text-align: right', editable: {type: 'date'} },
            {field: 'DtStop',      caption: 'DtStop', size: "50%",    sortable: true, style: 'text-align: right', editable: {type: 'date'} },
        ],
        onLoad: function(event) {
            event.onComplete = function() {
                this.url = '';
            };
        },
        onAdd: function(event) {
            var x = getCurrentBusiness(),
                BID=parseInt(x.value),
                BUD = getBUDfromBID(BID),
                fr = w2ui.rtForm.record,
                g = this,
                ndStart;

            // get lastest date among all market rate object's stopDate for new MR's StartDate
            g.records.forEach(function(rec) {
                if (ndStart === undefined) {
                    ndStart = new Date(rec.DtStop);
                }
                if (rec.DtStop) {
                    var rdStop = new Date(rec.DtStop);
                    if (ndStart < rdStop) {
                        ndStart = rdStop;
                    }
                }
            });

            var newRec = { recid: g.records.length,
                BID: BID,
                BUD: BUD,
                RTID: fr.RTID,
                RMRID: 0,
                MarketRate: 0,
                DtStart: dateFmtStr(ndStart),
                DtStop: "12/31/9999" };
            g.add(newRec);
        },
        onSave: function(event) {
            event.changes = this.records;
        },
        onDelete: function(event) {
            var selected = this.getSelection(),
                RMRIDList = [],
                grid = this;

            // if not selected then return
            if (selected.length < 0) {
                return;
            }
            // collect RMRID
            selected.forEach(function(id) {
                RMRIDList.push(grid.get(id).RMRID);
            });

            event.onComplete = function() {
                var x = getCurrentBusiness(),
                    BID=parseInt(x.value),
                    BUD = getBUDfromBID(BID),
                    RTID = w2ui.rtForm.record.RTID;

                var payload = { "cmd": "delete", "RMRIDList": RMRIDList };
                $.ajax({
                    type: "POST",
                    url: "/v1/rmr/" + BID + "/" + RTID,
                    data: JSON.stringify(payload),
                    contentType: "application/json",
                    dataType: "json",
                    success: function(data) {
                        grid.reload();
                    },
                });
            };
        },
        onChange: function(event) {
            event.preventDefault();
            var g = this,
                field = g.columns[event.column].field,
                chgRec = g.get(event.recid),
                changeIsValid = true;

            if ( field === "MarketRate" ) { // if field is MarketRate
                if (event.value_new <= 0) {
                    changeIsValid = false;
                }
            }

            // if fields are DtStart or DtStop
            if ( field === "DtStart" || field === "DtStop") {

                var chgDStart = field === "DtStart" ? new Date(event.value_new) : new Date(chgRec.DtStart),
                    chgDStop = field === "DtStop" ? new Date(event.value_new) : new Date(chgRec.DtStop);

                // Stop date should not before Start Date
                if (chgDStop <= chgDStart) {
                        changeIsValid = false;
                } else {
                    // make sure date values don't overlap with other market rate dates
                    for(var i in g.records) {
                        var rec = g.records[i];
                        if (rec.recid === chgRec.recid) { // if same record then continue to next one
                            continue;
                        }

                        var rDStart = new Date(rec.DtStart),
                            rDStop = new Date(rec.DtStop);

                        // return if changed record startDate falls in other MR time span
                        if (rDStart < chgDStart && chgDStart < rDStop) {
                            changeIsValid = false;
                        } else if(rDStart < chgDStop && chgDStop < rDStop) {
                            changeIsValid = false;
                        } else if(chgDStart < rDStart && rDStop < chgDStop) {
                            changeIsValid = false;
                        }
                    }
                }
            }

            if(changeIsValid) {
                // if everything is ok, then mark this as false
                event.isCancelled = false;
            } else {
                event.isCancelled = true;
            }

            event.onComplete = function() {
                if (!event.isCancelled) { // if event not cancelled then invoke save method
                    // save automatically locally
                    this.save();
                }
            };
        }
    });
}

function setRTLayout(BID, RTID) {
    var rtF = w2ui.rtForm,
        rtG = w2ui.rtGrid,
        rmrG = w2ui.rmrGrid;

    // set the url for rtForm
    rtF.url = '/v1/rt/' + BID + '/' + RTID;

    // load bottom panels with action buttons panel
    w2ui.rtDetailLayout.content("bottom", w2ui.rtFormBtns);

    // if form has tabs then click the first one
    if (typeof rtF.tabs.name == "string") {
        rtF.tabs.click('tab1');
    }

    // mark this flag as is this new record
    app.new_form_rec = RTID ? true : false;

    // as new content will be loaded for this form
    // mark form dirty flag as false
    app.form_is_dirty = false;

    if (RTID) {
        // if RentableType available then load the market rate grid
        rmrG.load('/v1/rmr/' + BID + '/' + RTID);

        // change the text of form tab
        w2ui.rtDetailLayout.get("main").tabs.get("rtForm").text = "Rentable Type Details ({0})".format(RTID);
        w2ui.rtDetailLayout.get("main").tabs.refresh();

        // load form content from server
        rtF.request(function(event) {
            if (event.status === "success") {
                // only render the toplayout after server has sent down data
                // so that w2ui can bind values with field's html control,
                // otherwise it is unable to find html controls
                showForm();
                return true;
            }
            else {
                showForm();
                rtF.message("Could not get form data from server...!!");
                return false;
            }
        });
    }
    else {
        // if new RentableType then clear the marketRate grid content first
        rmrG.clear();

        // change the text of form tab
        w2ui.rtDetailLayout.get("main").tabs.get("rtForm").text = "Rentable Type Details ({0})".format("new");
        w2ui.rtDetailLayout.get("main").tabs.refresh();

        // unselect the previous selected row
        var sel_recid = parseInt(rtG.last.sel_recid);
        if (sel_recid > -1) {
            // if new record is being added then unselect {{the selected record}} from the grid
            rtG.unselect(rtG.last.sel_recid);
        }

        showForm();
        return true;
    }

    function showForm() {
        // SHOW the right panel now
        w2ui.toplayout.content('right', w2ui.rtDetailLayout);
        w2ui.toplayout.sizeTo('right', 700);
        // w2ui.rtDetailLayout.render();
        w2ui.rtDetailLayout.get("main").tabs.click("rtForm");
        w2ui.toplayout.show('right', true);
    }
}