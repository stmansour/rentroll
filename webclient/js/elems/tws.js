"use strict";
window.buildTWSElements = function (){
//------------------------------------------------------------------------
//          twsGrid
//------------------------------------------------------------------------
$().w2grid({
    name: 'twsGrid',
    url: '/v1/tws',
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
        {field: 'recid',        caption: "recid",            size: '35px',  sortable: true, hidden: true},
        {field: 'TWSID',        caption: "TWSID",            size: '55px',  sortable: true, hidden: false},
        {field: 'Owner',        caption: "Owner",            size: '125px', sortable: true, hidden: false},
        {field: 'OwnerData',    caption: "Owner Data",       size: '20px',  sortable: true, hidden: true},
        {field: 'WorkerName',   caption: "WorkerName",       size: '125px', sortable: true, hidden: false},
        {field: 'ActivateTime', caption: "Activate Time",    size: '158px', sortable: true, hidden: false},
        {field: 'RemainingTime', caption: "Activates in...", size: '158px', sortable: true, hidden: false},
        {field: 'Node',         caption: "Node",             size: '100px', sortable: true, hidden: false},
        {field: 'FLAGS',        caption: "FLAGS",            size: '50px',  sortable: true, hidden: false},
        {field: 'DtActivated',  caption: "DtActivated",      size: '158px', sortable: true, hidden: false},
        {field: 'DtCompleted',  caption: "DtCompleted",      size: '158px', sortable: true, hidden: false},
        {field: 'DtCreate',     caption: "DtCreate",         size: '158px', sortable: true, hidden: false},
        {field: 'DtLastUpdate', caption: "DtLastUpdate",     size: '158px', sortable: true, hidden: false},
    ],
    // onClick: function(event) {
    //     event.onComplete = function () {
    //         app.new_form_rec = false;
    //         var rec = this.get(event.recid);
    //         setToForm('transactantForm', '/v1/person/' + rec.BID + '/' + rec.TCID, 700, true);
    //      };
    // },
});
};
