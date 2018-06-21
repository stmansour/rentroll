"use strict";
/*global
    renderLedgerStateIcon,
*/
var adminLedger = {
    mode: 0,
};
window.buildLedgerElements = function (){
    //------------------------------------------------------------------------
    //          ledger grid
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'ledgersGrid',
        url: '/v1/ledgers',
        multiSelect: false,
        show: {
            header: false,
            toolbar: true,
            footer: true,
            toolbarReload   : true,
            toolbarColumns  : false,
            toolbarSearch   : false,
            toolbarAdd      : false,
            toolbarDelete   : false,
            toolbarInput    : false,
            searchAll       : false,
            toolbarSave     : false,
            toolbarEdit     : false,
            searches        : false,
            lineNumbers     : false,
            selectColumn    : false,
            expandColumn    : false,
        },
        columns: [
            {field: 'recid',     caption: "recid",         size: '35px',  sortable: true, hidden: true},
            {field: 'LID',       caption: "LID",           size: '55px',  sortable: true, hidden: false},
            {field: 'RAID',      caption: "RAID",          size: '55px',  sortable: true, hidden: false},
            {field: 'RID',       caption: "RID",           size: '55px',  sortable: true, hidden: false},
            {field: 'GLNumber',  caption: "GLNumber",      size: '85px',  sortable: true, hidden: false},
            {field: 'Name',      caption: "Name",          size: '225px', sortable: true, hidden: false},
            {field: 'Active',    caption: "Active",        size: '50px',  sortable: true, hidden: false},
            // {field: 'AllowPost', caption: "Allow<br>Post", size: '50px',  sortable: true, hidden: false},
            {field: 'Balance',   caption: "Balance",       size: '100px', sortable: true, hidden: false, render: 'money'},
            {field: 'LMDate',    caption: "LM Date",       size: '170px', sortable: true, hidden: false},
            {field: 'LMState',   caption: "LM<br>State",   size: '75px',  sortable: true, hidden: false, render: renderLedgerStateIcon},
            {field: 'LMAmount',  caption: "LM Amount",     size: '100px', sortable: true, hidden: false, render: 'money'},
        ],
        onRequest: function(/*event*/) {
            adminLedger.mode = document.getElementById('adminLedgerMode').options.selectedIndex;
            w2ui.ledgersGrid.postData = {
                mode : adminLedger.mode,
                searchDtStart: app.D1,
                searchDtStop: app.D2,
                client: app.client
            };
        },
        onLoad: function(event) {
            event.onComplete = function() {
                console.log('onLoad: ' + event);
                document.getElementById('adminLedgerMode').value = adminLedger.mode;
                document.getElementById("adminLedgerMode").options.selectedIndex = adminLedger.mode;
            };
        },
    });
    addDateNavToToolbar('ledgers');
    var items = [
            { type: 'html',  id: 'mode',
                html: function (/*item*/) {
                    var html =
                        '&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<i class="far fa-building"></i> &nbsp;Mode:' +
                        '<select id="adminLedgerMode" onchange="changeLedgerMode();">'+
                        '<option value="0">Initial Ledger Markers</option>'+
                        '<option value="1">GL Account Ledger Markers</option>'+
                        '<option value="2">RAID Ledger Markers</option>'+
                        '<option value="3">RID Ledger Markers</option>'+
                        '</select>&nbsp;&nbsp;&nbsp;';
                    return html;
                }
            },
    ];
    w2ui.ledgersGrid.toolbar.add( items );
};

window.renderLedgerStateIcon = function (record /*, index, col_index*/) {
    if (typeof record === "undefined") {
        return;
    }
    var s = '';
    switch (record.LMState) {
    	case "initial":
        s = '<i class="fas fa-home" style="color: #0088DD;"></i> &nbsp;';
    	break;
    	case "closed":
    	s = '<i class="fas fa-lock" style="color: #0088DD;"></i> &nbsp;';
    	break;
    	case "locked":
    	s = '<i class="fas fa-ban" style="color: #0088DD;"></i> &nbsp;';
    	break;
    }
    return s + record.LMState;
};


window.changeLedgerMode = function(){
    // console.log('hello: ' + event);
    adminLedger.mode = document.getElementById('adminLedgerMode').options.selectedIndex;
    w2ui.ledgersGrid.postData = {
        mode : adminLedger.mode,
        searchDtStart: app.D1,
        searchDtStop: app.D2,
        client: app.client
    };
    w2ui.ledgersGrid.reload();
};