/*global
   $,addDateNavToToolbar,getCurrentBID,w2ui,w2utils,
   manageParentRentableW2UIItems,RACompConfig, LoadRAFlowTemplate
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
            {field: 'Payors', caption: 'Payor(s)', size: '250px', sortable: true},
            {field: 'AgreementStart', caption: 'Agreement<br>Start', render: 'date', size: '80px', sortable: true, style: 'text-align: right'},
            {field: 'AgreementStop', caption: 'Agreement<br>Stop',  render: 'date', size: '80px', sortable: true, style: 'text-align: right'},
            {field: 'PayorIsCompany', hidden: true, caption: 'IsCompany',  size: '40px', sortable: false},
        ],
        onClick: function(event) {
            event.onComplete = function() {
                var bid = getCurrentBID();
                var r = w2ui.ra2flowGrid.records[event.recid];
                var url = "/v1/ra2flow/" + bid + '/' + r.RAID;
                var rec = { cmd: "get"};
                var dat=JSON.stringify(rec);
                $.post(url,dat)
                .done(function(data) {
                    if (data.status === "error") {
                        w2ui.ra2flowGrid.error(w2utils.lang(url + ' failed:  ' + data.message));
                        return;
                    }

                    // set the record in app raflow
                    app.raflow.data[data.record.FlowID]= data.record;

                    // load ra flow template
                    LoadRAFlowTemplate(bid, data.record.FlowID);
                })
                .fail(function(/*data*/){
                    w2ui.ra2flowGrid.error("Get Rental Agreement Flow failed.");
                    return;
                });
            };
        },
    });
    addDateNavToToolbar('ra2flow');
};
