"use strict";
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
        {field: 'GLNumber',  caption: "GLNumber",      size: '85px',  sortable: true, hidden: false},
        {field: 'Name',      caption: "Name",          size: '225px', sortable: true, hidden: false},
        {field: 'Active',    caption: "Active",        size: '50px',  sortable: true, hidden: false},
        // {field: 'AllowPost', caption: "Allow<br>Post", size: '50px',  sortable: true, hidden: false},
        {field: 'Balance',   caption: "Balance",       size: '100px', sortable: true, hidden: false, render: 'money'},
        {field: 'LMDate',    caption: "LM Date",       size: '170px', sortable: true, hidden: false},
        {field: 'LMState',   caption: "LM<br>State",   size: '50px',  sortable: true, hidden: false},
        {field: 'LMAmount',  caption: "LM Amount",     size: '100px', sortable: true, hidden: false, render: 'money'},
    ],
    onRequest: function(/*event*/) {
        w2ui.ledgersGrid.postData = {searchDtStart: app.D1, searchDtStop: app.D2};
    },
    // onClick: function(event) {
    //     event.onComplete = function () {
    //         app.new_form_rec = false;
    //         var rec = this.get(event.recid);
    //         setToForm('transactantForm', '/v1/person/' + rec.BID + '/' + rec.TCID, 700, true);
    //      };
    // },
});
addDateNavToToolbar('ledgers');
};
