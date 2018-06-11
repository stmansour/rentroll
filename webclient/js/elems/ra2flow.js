/*global
   $,addDateNavToToolbar
*/

"use strict";

window.buildRA2FlowElements = function() {
    //------------------------------------------------------------------------
    //          rentalagrsGrid
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'ra2flowGrid',
        url: '/v1/ra2flow',
        multiSelect: false,
        show: {
            toolbar: true,
            footer: true,
            lineNumbers: false,
            selectColumn: false,
            expandColumn: false,
            toolbarAdd: true,
            toolbarDelete: false,
            toolbarSave: false,
            toolbarEdit: false,
            toolbarSearch: true,
            toolbarInput: true,
            searchAll: true,
            toolbarReload: true,
            toolbarColumns: false,
        },
        multiSearch: true,
        searches: [
            { field: 'RAID', caption: 'RAID', type: 'text' },
            { field: 'Payors', caption: 'Payor(s)', type: 'text' },
            { field: 'AgreementStart', caption: 'Agreement Start Date', type: 'date' },
            { field: 'AgreementStop', caption: 'Agreement Stop Date', type: 'date' },
        ],
        columns: [
            {field: 'recid', hidden: true, caption: 'recid',  size: '40px', sortable: true},
            {field: 'BID', hidden: true, caption: 'BID',  size: '40px', sortable: false},
            {field: 'RAID', caption: 'RAID',  size: '50px', sortable: true},
            {field: 'Payors', caption: 'Payor(s)', size: '250px', sortable: true,
                // render: function (record, index, col_index) {
                //     if (record) {
                //         var icon;
                //         if (record.PayorIsCompany) {
                //             icon = 'fa-handshake-o'
                //         } else {
                //             icon = 'fa-user-o'
                //         }
                //         return '<i class="fa '+icon+'"></i>&nbsp;<span>'+record.Payors+'</span>';
                //     }
                // },
            },
            {field: 'AgreementStart', caption: 'Agreement<br>Start', render: 'date', size: '80px', sortable: true, style: 'text-align: right'},
            {field: 'AgreementStop', caption: 'Agreement<br>Stop',  render: 'date', size: '80px', sortable: true, style: 'text-align: right'},
            {field: 'PayorIsCompany', hidden: true, caption: 'IsCompany',  size: '40px', sortable: false},
        ],
    });
    addDateNavToToolbar('ra2flow');
};