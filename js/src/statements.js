"use strict";

function buildStatementsElements() {

	//------------------------------------------------------------------------
	//          stmtGrid  -  THE LIST OF ALL RENTAL AGREEMENTS
	//------------------------------------------------------------------------
	$().w2grid({
	    name: 'stmtGrid',
	    url: '/v1/stmt',
	    multiSelect: false,
	    postData: {searchDtStart: app.D1, searchDtStop: app.D2},
	    show: {
	        toolbar         : true,
	        footer          : true,
	        toolbarAdd      : false,   // indicates if toolbar add new button is visible
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
	        {field: 'recid', hidden: true,  caption: 'recid',            size: '40px',  sortable: true},
	        {field: 'BID',   hidden: true,  caption: 'BID',              size: '40px',  sortable: true},
	        {field: 'RAID',  hidden: false, caption: 'Rental Agreement', size: '200px', sortable: true},
	        {field: 'Payors',hidden: false, caption: 'Payors',           size: '250px', sortable: true},
	    ],
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
                        console.log( 'BID = ' + rec.BID + ',   RAID = ' + rec.RAID);
                        var d = new Date();  // we'll use today for time-sensitive data
                        setToStmtForm(rec.BID, rec.RAID, d,d);
                    };

                // warn user if form content has been changed
                form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
            };
        },
	});

    addDateNavToToolbar('stmt');

    //------------------------------------------------------------------------
    //  stmtDetailForm
    //------------------------------------------------------------------------
    $().w2form({
        name: 'stmtDetailForm',
        style: 'border: 0px; background-color: transparent;',
        header: app.sRentalAgreement + ' Detail',
        url: '/v1/stmtInfo',
        formURL: '/html/formstmtdet.html',
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
                            w2ui.stmtGrid.render();
                        };
                    form_dirty_alert(yes_callBack, no_callBack);
                    break;
                }
            },
        },
        fields: [
            { field: 'recid', type: 'int', required: false, html: {page: 0, column: 0 } },
            { field: 'RAID', type: 'int', required: false, html: {  page: 0, column: 0 } },
            { field: 'BID', type: 'int', required: false, html: { page: 0, column: 0 } },
            { field: 'BUD', type: 'list', options: {items: app.businesses}, required: true, html: { page: 0, column: 0 } },
        ],
        save: function () {
            	console.log('Save button from stmtDetailForm');
        },
    });

	//------------------------------------------------------------------------
	//  stmtDetailGrid  -  lists all the assessments and receipts for
	//                     the selected Rental Agreement from the stmtGrid
	//------------------------------------------------------------------------
	$().w2grid({
	    name: 'stmtDetailGrid',
	    url: '/v1/stmtDetail',
	    multiSelect: false,
	    postData: {searchDtStart: app.D1, searchDtStop: app.D2},
	    show: {
	        toolbar         : true,
	        footer          : true,
	        toolbarAdd      : false,   // indicates if toolbar add new button is visible
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
	        {field: 'recid',      caption: 'recid',       size: '40px',  sortable: true, hidden: true},
	        {field: 'Date',       caption: 'Date',        size: '80px',  sortable: true},
	        {field: 'ID',         caption: 'ID',          size: '100px', sortable: true},
	        {field: 'Descr',      caption: 'Description', size: '250px', sortable: true},
	        {field: 'Assessment', caption: 'Assessment',  size: '100px', sortable: true},
	        {field: 'Receipt',    caption: 'Receipt',     size: '100px', sortable: true},
	        {field: 'Balance',    caption: 'Balance',     size: '100px', sortable: true},
	    ],
	});

    addDateNavToToolbar('stmtDetail');

	//------------------------------------------------------------------------
	//  stmtlayout - The layout to contain the stmtForm and stmtDetailGrid
	//               top  - stmtForm
	//               main - stmtDetailGrid
	//------------------------------------------------------------------------
    $().w2layout({
        name: 'stmtLayout',
        padding: 0,
        panels: [
            { type: 'left',    size: '40%', hidden: true },
            { type: 'top',     size: 300,   hidden: false, content: w2ui.stmtDetailForm,  resizable: true, style: app.pstyle },
            { type: 'main',    size: '60%', hidden: false, content: w2ui.stmtDetailGrid, resizable: true, style: app.pstyle },
            { type: 'preview', size: 0,     hidden: true,  content: 'PREVIEW'  },
            { type: 'bottom',  size: 0,     hidden: true },
            { type: 'right',   size: 0,     hidden: true }
        ]
    });
}

//-----------------------------------------------------------------------------
// setToStmtForm -  enable the Statement form in toplayout.  Also, set
//                the forms url and request data from the server
// @params
//   bid = business id (or the BUD)
//  raid = Rental Agreement ID
// d1,d2 = date range to use
//-----------------------------------------------------------------------------
function setToStmtForm(bid, raid, d1,d2) {
    if (raid > 0) {
        w2ui.stmtDetailGrid.url = '/v1/stmtDetail/' + bid + '/' + raid;
        console.log( 'w2ui.stmtDetailGrid = '+w2ui.stmtDetailGrid);
		w2ui.stmtDetailGrid.request('get',{},w2ui.stmtDetailGrid.url);
        var f = w2ui.stmtDetailForm;
        w2ui.toplayout.content('right', w2ui.stmtLayout);
        w2ui.toplayout.show('right', true);
        w2ui.toplayout.sizeTo('right', 700);
        f.url = '/v1/stmtInfo/' + bid + '/' + raid;
        f.request();
        w2ui.toplayout.render();
        app.new_form_rec = false; // mark as record exists
        app.form_is_dirty = false; // mark as no changes yet
    }
}
