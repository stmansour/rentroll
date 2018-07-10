    /*global
    	loadRAActionInProgress,
        loadRAActionFirstApproval,
        loadRAActionTemplate
    */
    "use strict";

    //------------------------------------------------------------------------
    // loadRAActionTemplate - It creates a layout for action forms and places
    //                        it in newralayout's right panel.
    //                        Top panel & bottom panel of this layout contains
    //                        header & footer of action form respectively.
    // -----------------------------------------------------------------------
    window.loadRAActionTemplate = function() {
        if(! w2ui.actionLayout) {
            $().w2layout({
                name: 'actionLayout',
                padding: 0,
                panels: [
                    { type: 'left', hidden: true },
                    { type: 'top', content:'top', size:130,
                        toolbar: {
                            items: [
                                { id: 'btnNotes', type: 'button', icon: 'far fa-sticky-note' },
                                { id: 'bt3', type: 'spacer' },
                                { id: 'btnClose', type: 'button', icon: 'fas fa-times' }
                            ],
                            onClick: function (event) {
                                switch(event.target) {
                                case 'btnClose':
                                    var no_callBack = function() { return false; },
                                        yes_callBack = function() {

                                            w2ui.newraLayout.content('right','');
                                            w2ui.newraLayout.hide('right',true);
                                            w2ui.actionLayout.get('main').content.destroy();
                                            w2ui.newraLayout.get('main').toolbar.refresh();
                                        };
                                    form_dirty_alert(yes_callBack, no_callBack);
                                    break;
                                }
                            },
                        }
                    },
                    { type: 'main', style: app.pstyle, content: 'main'},
                    { type: 'preview', hidden: true },
                    { type: 'bottom', size: 56,content:'bottom' },
                    { type: 'right', hidden: true}
                ],
                onRender: function(event) {

                    event.onComplete = function() {
                        var activeFlowID = app.raflow.activeFlowID;
                        var data = app.raflow.data[activeFlowID];

                        var x = document.getElementById("bannerRAID");
                        if (x !== null) {
                            x.innerHTML = '' + data.Data.meta.RAID;
                        }
                        x = document.getElementById("bannerTermDates");
                        if (x !== null) {
                            x.innerHTML = '' + data.Data.dates.AgreementStart + ' - ' + data.Data.dates.AgreementStop;
                        }
                        x = document.getElementById("bannerPossessionDates");
                        if (x !== null) {
                            x.innerHTML = '' + data.Data.dates.PossessionStart + ' - ' + data.Data.dates.PossessionStop;
                        }
                        x = document.getElementById("bannerRentDates");
                        if (x !== null) {
                            x.innerHTML = '' + data.Data.dates.RentStart + ' - ' + data.Data.dates.RentStop;
                        }
                    };
                },
            });
        }
        w2ui.newraLayout.content('right', w2ui.actionLayout);

        w2ui.actionLayout.load('top', '/webclient/html/raflow/formra-actionheader.html');
        w2ui.actionLayout.load('bottom', '/webclient/html/raflow/formra-actionfooter.html');

        var raFlags = app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS;
        var raFlagsString = app.RAStates[parseInt(raFlags & 0xf)];

        switch (raFlagsString) {
            case "Application Being Filled In" :
                console.log('Inside Application Being Filled In switch case');
                loadRAActionInProgress();
                break;

            case "Pending First Approval" :
                console.log('Inside Pending First Approval switch case');
                loadRAActionFirstApproval();
                break;

            case "Pending Second Approval" :
                break;
            case "Move In" :
                break;
            case "Active" :
                break;
            case "Terminated" :
                break;
            case "Notice To Move" :
                break;
        }
        w2ui.newraLayout.show('right', true);
        w2ui.newraLayout.sizeTo('right', 950);
    };

    // -------------------------------------------------------------------------------
    // Rental Agreement Action - In Progress form
    // -------------------------------------------------------------------------------
    window.loadRAActionInProgress = function () {

        if (! w2ui.RAActionInProgress) {
            // InProgress form
            $().w2form({
                name: 'RAActionInProgress',
                style: 'background-color: white; display: block;',
                focus: -1,
                formURL: '/webclient/html/raflow/formra-actioninprogress.html',
                fields: [
                	{ field: 'RAActions', type: 'list', width: 120, required: true, options: {items: app.w2ui.listItems.RAActions}},
                ],
                onRefresh: function (event) {
                	console.log('onRefresh of RAActionInProgress');
                	var activeFlowID = app.raflow.activeFlowID;
                    var data = app.raflow.data[activeFlowID].Data;
                    var raFlags = data.meta.RAFLAGS;
                    var raStateString = app.RAStates[parseInt(raFlags & 0xf)];
            		$('#RAActionStateLable').text(raStateString);
                },
                onChange: function (event) {
                    event.done(function(){
                        w2ui.RAActionInProgress.refresh();
                    });
                },
                onRender: function (event) {
                    w2ui.RAActionInProgress.record = {RAActions: {id: 0, text: "Edit Rental Agreement Information"}};
                },
                actions: {
                    save: function () {
                        var actionn = w2ui.RAActionInProgress.record.RAActions.text;
                        var raFlags = app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS;
                        raFlags = raFlags & ~(0xf);

                        switch (actionn) {
                            case "Edit Rental Agreement Information" :
                                raFlags = raFlags | 0;
                                app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                                break;
                            case "Authorize First Approval" :
                                raFlags = raFlags | 1;
                                app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                                break;
                            case "Authorize Second Approval" :
                                break;
                            case "Complete Move In" :
                                break;
                            case "Terminate" :
                                break;
                            case "Received Notice To Move" :
                                break;
                        }
                        loadRAActionTemplate();
                    }
                }
            });
        }
        // now render the form in specifiec targeted panel
        w2ui.actionLayout.content('main', w2ui.RAActionInProgress);
    };

    // -------------------------------------------------------------------------------
    // Rental Agreement Action - First Approval form
    // -------------------------------------------------------------------------------
    window.loadRAActionFirstApproval = function () {
        if (! w2ui.RAActionFirstApproval) {
            // FirstApproval form
            $().w2form({
                name: 'RAActionFirstApproval',
                style: 'background-color: white; display: block;',
                focus: -1,
                formURL: '/webclient/html/raflow/formra-actionfirstapproval.html',
                fields: [
                    { field: 'RAActions', type: 'list', width: 120, required: true,
                        options: {
                            items: app.w2ui.listItems.RAActions
                        }
                    },
                    { field: 'RAApprovalDecision1', type: 'list', width: 120, required: true,
                        options: {
                            items: ['Approve', 'Decline']
                        }
                    },
                    { field: 'RADeclineReason', type: 'list', width: 120, required: true, hidden: true,
                        options: {
                            items: ['Temp1', 'Temp2']
                        }
                    },
                ],
                onChange: function (event) {
                    event.done(function(){
                        w2ui.RAActionFirstApproval.refresh();
                    });

                    if(event.target === 'RAActions') {
                        if(event.value_new.text === 'Authorize First Approval') {
                            w2ui.RAActionFirstApproval.get('RAApprovalDecision1').hidden = false;
                        } else {
                            w2ui.RAActionFirstApproval.get('RAApprovalDecision1').hidden = true;
                            w2ui.RAActionFirstApproval.get('RADeclineReason').hidden = true;
                            $('[name="RAUpdateStatus"]').hide();

                            this.clear();
                        }
                        w2ui.RAActionFirstApproval.refresh();
                    }
                    if(event.target === 'RAApprovalDecision1') {
                        if(event.value_new.text === 'Decline') {
                            w2ui.RAActionFirstApproval.get('RADeclineReason').hidden = false;
                        } else {
                            w2ui.RAActionFirstApproval.get('RADeclineReason').hidden = true;
                            delete this.record.RADeclineReason;
                            $('[name="RAUpdateStatus"]').hide();
                            $('[name="save"]').show();
                        }
                        w2ui.RAActionFirstApproval.refresh();
                    }

                    if(event.target === 'RADeclineReason') {
                        console.log(event);
                        if(event.value_new.text != '') {
                            $('[name="RAUpdateStatus"]').show();
                            $('[name="save"]').hide();
                        } else {
                            $('[name="RAUpdateStatus"]').hide();
                        }
                        w2ui.RAActionFirstApproval.refresh();
                    }
                },
                onRefresh: function (event) {
                    console.log('onRefresh of RAActionFirstApproval');
                    var activeFlowID = app.raflow.activeFlowID;
                    var data = app.raflow.data[activeFlowID].Data;
                    var raFlags = data.meta.RAFLAGS;
                    var raStateString = app.RAStates[parseInt(raFlags & 0xf)];
                    var RAID = data.meta.RAID;
                    if(RAID > 0) {
                        raStateString = 'Modification ' + raStateString;
                    }
                    $('#RAActionStateLable').text(raStateString);
                },
                onRender: function (event) {
                    w2ui.RAActionFirstApproval.record = {RAActions: {id: 1, text: "Authorize First Approval"}};
    
                },
                actions: {
                    save: function () {
                        var actionn = w2ui.RAActionFirstApproval.record.RAActions.text;
                        var raFlags = app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS;
                        raFlags = raFlags & ~(0xf);

                        switch (actionn) {
                            case "Edit Rental Agreement Information" :
                                raFlags = raFlags | 0;
                                app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                                break;
                            case "Authorize First Approval" :
                                if(this.record.RAApprovalDecision1 != undefined) {
                                    raFlags = raFlags | 2;
                                    app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                                }
                                break;
                            case "Authorize Second Approval" :
                                break;
                            case "Complete Move In" :
                                break;
                            case "Terminate" :
                                break;
                            case "Received Notice To Move" :
                                break;
                        }
                        w2ui.actionLayout.get('main').content.destroy();
                        loadRAActionTemplate();
                    },
                    RAUpdateStatus: function () {
                        console.log("Update Button Clicked");
                    }
                }
            });
        }
        // now render the form in specifiec targeted panel
        w2ui.actionLayout.content('main', w2ui.RAActionFirstApproval);
    };
