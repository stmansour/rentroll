/*global
   $,addDateNavToToolbar,getCurrentBID,w2ui,w2utils,
   manageParentRentableW2UIItems,RACompConfig,
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
                    w2ui.toplayout.content('right', w2ui.newraLayout);
                    w2ui.toplayout.show('right', true);
                    w2ui.toplayout.sizeTo('right', 950);

                    $.get('/webclient/html/raflowtmpl.html', function(htmldata) {
                        w2ui.newraLayout.content('main', htmldata);
                        w2ui.toplayout.render();
                        $(".ra-form-component").hide();
                        $("#progressbar #steps-list li").removeClass("active done"); // remove activeClass from all li

                        setTimeout(function() {
                            $("#ra-form footer button#previous").prop("disabled", true);

                            // mark this flag as is this new record
                            // record created already
                            app.new_form_rec = false;

                            // as new content will be loaded for this form
                            // mark form dirty flag as false
                            app.form_is_dirty = false;

                            // set this flow id as in active
                            app.raflow.activeFlowID = data.record.FlowID;
                            app.raflow.data[app.raflow.activeFlowID]= data.record;

                            // set BID in raflow settings
                            app.raflow.BID = bid;

                            // calculate parent rentable items
                            manageParentRentableW2UIItems();

                            // show "done" mark on each li of navigation bar
                            for (var comp in app.raFlowPartTypes) {
                                // if required fields are fulfilled then mark this slide as done
                                if (requiredFieldsFulFilled(comp)) {
                                    // hide active component
                                    $("#progressbar #steps-list li[data-target='#" + comp + "']").addClass("done");
                                }

                                // reset w2ui component as well
                                if(RACompConfig[comp].w2uiComp in w2ui) {
                                    // clear inputs
                                    w2ui[RACompConfig[comp].w2uiComp].clear();
                                }
                            }

                            // mark first slide as active
                            $(".ra-form-component#dates").show();
                            $("#progressbar #steps-list li[data-target='#dates']").removeClass("done").addClass("active");
                            loadRADatesForm();
                        }, 0);
                    });
                })
                .fail(function(/*data*/){
                    w2ui.ra2flowGrid.error("Save Tasklist failed.");
                    return;
                });
            };
        },
    });
    addDateNavToToolbar('ra2flow');
};
