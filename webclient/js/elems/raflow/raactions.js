/*global
	loadRAActionInProgress,
    loadRAActionFirstApproval,
    loadRAActionSecondApproval,
    loadRAActionMoveIn,
    loadRAActionActive,
    loadRAActionTerminated,
    loadRAActionNoticeToMove,
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
                { type: 'left', style: app.pstyle2, hidden: true },
                { type: 'top', style: app.pstyle2, content:'top', size:110,
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
                                        w2ui.newraLayout.unlock('main');
                                        w2ui.newraLayout.get('main').toolbar.refresh();
                                    };
                                form_dirty_alert(yes_callBack, no_callBack);
                                break;
                            }
                        },
                    }
                },
                { type: 'main', style: app.pstyle2, content: 'main'},
                { type: 'preview', style: app.pstyle2, hidden: true },
                { type: 'bottom', style: app.pstyle2, size: 40,content:'bottom' },
                { type: 'right', style: app.pstyle2, hidden: true}
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
            console.log('Inside Pending Second Approval switch case');
            loadRAActionSecondApproval();
            break;
        case "Move In" :
            console.log('Inside Move In switch case');
            loadRAActionMoveIn();
            break;
        case "Active" :
            console.log('Inside Active switch case');
            loadRAActionActive();
            break;
        case "Terminated" :
            console.log('Inside Terminated switch case');
            loadRAActionTerminated();
            break;
        case "Notice To Move" :
            console.log('Inside Notice To Move switch case');
            loadRAActionNoticeToMove();
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
                { field: 'RADeclineReason1', type: 'list', width: 120, required: true, hidden: true,
                    options: {
                        items: ['Temp1', 'Temp2']
                    }
                }
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
                        w2ui.RAActionFirstApproval.get('RADeclineReason1').hidden = true;
                        $('[name="RAUpdateStatus"]').hide();

                        this.clear();
                    }
                    w2ui.RAActionFirstApproval.refresh();
                }
                if(event.target === 'RAApprovalDecision1') {
                    if(event.value_new.text === 'Decline') {
                        w2ui.RAActionFirstApproval.get('RADeclineReason1').hidden = false;
                    } else {
                        w2ui.RAActionFirstApproval.get('RADeclineReason1').hidden = true;
                        delete this.record.RADeclineReason1;
                        $('[name="RAUpdateStatus"]').hide();
                        $('[name="save"]').show();
                    }
                    w2ui.RAActionFirstApproval.refresh();
                }

                if(event.target === 'RADeclineReason1') {
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

// -------------------------------------------------------------------------------
// Rental Agreement Action - Second Approval form
// -------------------------------------------------------------------------------
window.loadRAActionSecondApproval = function () {
    if (! w2ui.RAActionSecondApproval) {
        // SecondApproval form
        $().w2form({
            name: 'RAActionSecondApproval',
            style: 'background-color: white; display: block;',
            focus: -1,
            formURL: '/webclient/html/raflow/formra-actionsecondapproval.html',
            fields: [
                { field: 'RAActions', type: 'list', width: 120, required: true,
                    options: {
                        items: app.w2ui.listItems.RAActions
                    }
                },
                { field: 'RAApprovalDecision2', type: 'list', width: 120, required: true,
                    options: {
                        items: ['Approve', 'Decline']
                    }
                },
                { field: 'RADeclineReason2', type: 'list', width: 120, required: true, hidden: true,
                    options: {
                        items: ['Temp1', 'Temp2']
                    }
                },
            ],
            onChange: function (event) {
                event.done(function(){
                    w2ui.RAActionSecondApproval.refresh();
                });

                if(event.target === 'RAActions') {
                    if(event.value_new.text === 'Authorize Second Approval') {
                        w2ui.RAActionSecondApproval.get('RAApprovalDecision2').hidden = false;
                    } else {
                        w2ui.RAActionSecondApproval.get('RAApprovalDecision2').hidden = true;
                        w2ui.RAActionSecondApproval.get('RADeclineReason2').hidden = true;
                        $('[name="RAUpdateStatus"]').hide();

                        this.clear();
                    }
                    w2ui.RAActionSecondApproval.refresh();
                }
                if(event.target === 'RAApprovalDecision2') {
                    if(event.value_new.text === 'Decline') {
                        w2ui.RAActionSecondApproval.get('RADeclineReason2').hidden = false;
                    } else {
                        w2ui.RAActionSecondApproval.get('RADeclineReason2').hidden = true;
                        delete this.record.RADeclineReason2;
                        $('[name="RAUpdateStatus"]').hide();
                        $('[name="save"]').show();
                    }
                    w2ui.RAActionSecondApproval.refresh();
                }

                if(event.target === 'RADeclineReason2') {
                    console.log(event);
                    if(event.value_new.text != '') {
                        $('[name="RAUpdateStatus"]').show();
                        $('[name="save"]').hide();
                    } else {
                        $('[name="RAUpdateStatus"]').hide();
                    }
                    w2ui.RAActionSecondApproval.refresh();
                }
            },
            onRefresh: function (event) {
                console.log('onRefresh of RAActionSecondApproval');
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
                w2ui.RAActionSecondApproval.record = {RAActions: {id: 2, text: "Authorize Second Approval"}};
            },
            actions: {
                save: function () {
                    var actionn = w2ui.RAActionSecondApproval.record.RAActions.text;
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
                            if(this.record.RAApprovalDecision2 != undefined) {
                                raFlags = raFlags | 3;
                                app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                            }
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
    w2ui.actionLayout.content('main', w2ui.RAActionSecondApproval);
};

// -------------------------------------------------------------------------------
// Rental Agreement Action - Move-In form
// -------------------------------------------------------------------------------
window.loadRAActionMoveIn = function () {
    if (! w2ui.RAActionMoveIn) {
        // Move-In form
        $().w2form({
            name: 'RAActionMoveIn',
            style: 'background-color: white; display: block;',
            focus: -1,
            formURL: '/webclient/html/raflow/formra-actionmovein.html',
            fields: [
                { field: 'RAActions', type: 'list', width: 120, required: true,
                    options: {
                        items: app.w2ui.listItems.RAActions
                    }
                },
                { field: 'RADocumentDate', type: 'date', required: true , hidden: true}
            ],
            onChange: function (event) {
                event.done(function(){
                    w2ui.RAActionMoveIn.refresh();
                });

                if(event.target === 'RAActions') {
                    if(event.value_new.text === 'Complete Move In') {
                        $('[name="RAGenerateRAForm"]').show();
                        $('[name="RAGenerateMoveInForm"]').show();
                    } else {
                        w2ui.RAActionMoveIn.get('RADocumentDate').hidden = true;
                        $('[name="RAGenerateRAForm"]').hide();
                        $('[name="RAGenerateMoveInForm"]').hide();

                        this.clear();
                    }
                    w2ui.RAActionMoveIn.refresh();
                }

                if(event.target === 'RADocumentDate') {
                    var dateString = w2ui.RAActionMoveIn.get('RADocumentDate').el.value;
                    var yearString = dateString.substr(dateString.length - 4);
                    if (yearString <= "1999" ) {
                        this.record.RADocumentDate = w2uiDateControlString(new Date());
                        this.refresh();
                    }
                }
            },
            onRefresh: function (event) {
                console.log('onRefresh of RAActionMoveIn');
                var activeFlowID = app.raflow.activeFlowID;
                var data = app.raflow.data[activeFlowID].Data;
                var raFlags = data.meta.RAFLAGS;
                var raStateString = app.RAStates[parseInt(raFlags & 0xf)];
                $('#RAActionStateLable').text(raStateString);
            },
            onRender: function (event) {
                w2ui.RAActionMoveIn.record = {RAActions: {id: 3, text: "Complete Move In"}};
            },
            actions: {
                save: function () {
                    var actionn = w2ui.RAActionSecondApproval.record.RAActions.text;
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
                            if(this.record.RAApprovalDecision2 != undefined) {
                                raFlags = raFlags | 3;
                                app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                            }
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
                RAGenerateRAForm: function() {
                    if (w2ui.RAActionMoveIn.get('RADocumentDate').hidden) {
                        var t = new Date();
                        this.record.RADocumentDate = w2uiDateControlString(t);
                        w2ui.RAActionMoveIn.get('RADocumentDate').hidden = false;
                        this.refresh();
                    } else {

                    }
                }
            }
        });
    }
    // now render the form in specifiec targeted panel
    w2ui.actionLayout.content('main', w2ui.RAActionMoveIn);
};

// -------------------------------------------------------------------------------
// Rental Agreement Action - Active form
// -------------------------------------------------------------------------------
window.loadRAActionActive = function () {
    if (! w2ui.RAActionActive) {
        // Move-In form
        $().w2form({
            name: 'RAActionActive',
            style: 'background-color: white; display: block;',
            focus: -1,
            formURL: '/webclient/html/raflow/formra-actionactive.html',
            fields: [
                { field: 'RAActions', type: 'list', width: 120, required: true,
                    options: {
                        items: app.w2ui.listItems.RAActions
                    }
                },
                { field: 'RATerminationReason', type: 'list', width: 120, required: true, hidden: true,
                    options: {
                        items: ['Temp1', 'Temp2']
                    }
                }
            ],
            onChange: function (event) {
                event.done(function(){
                    w2ui.RAActionActive.refresh();
                });

                if(event.target === 'RAActions') {
                    if(event.value_new.text === 'Terminate') {
                        w2ui.RAActionActive.get('RATerminationReason').hidden = false;
                    } else {
                        w2ui.RAActionActive.get('RATerminationReason').hidden = true;
                        this.clear();
                    }
                }
            },
            onRefresh: function (event) {
                console.log('onRefresh of RAActionActive');
                var activeFlowID = app.raflow.activeFlowID;
                var data = app.raflow.data[activeFlowID].Data;
                var raFlags = data.meta.RAFLAGS;
                var raStateString = app.RAStates[parseInt(raFlags & 0xf)];
                $('#RAActionStateLable').text(raStateString);
            },
            actions: {
                save: function () {
                    // var actionn = w2ui.RAActionSecondApproval.record.RAActions.text;
                    // var raFlags = app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS;
                    // raFlags = raFlags & ~(0xf);

                    // switch (actionn) {
                    //     case "Edit Rental Agreement Information" :
                    //         raFlags = raFlags | 0;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    //     case "Authorize First Approval" :
                    //         if(this.record.RAApprovalDecision1 != undefined) {
                    //             raFlags = raFlags | 2;
                    //             app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         }
                    //         break;
                    //     case "Authorize Second Approval" :
                    //         if(this.record.RAApprovalDecision2 != undefined) {
                    //             raFlags = raFlags | 3;
                    //             app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         }
                    //         break;
                    //     case "Complete Move In" :
                    //         break;
                    //     case "Terminate" :
                    //         break;
                    //     case "Received Notice To Move" :
                    //         break;
                    // }
                    // w2ui.actionLayout.get('main').content.destroy();
                    // loadRAActionTemplate();
                }
            }
        });
    }
    // now render the form in specifiec targeted panel
    w2ui.actionLayout.content('main', w2ui.RAActionActive);
};

// -------------------------------------------------------------------------------
// Rental Agreement Action - Terminated form
// -------------------------------------------------------------------------------
window.loadRAActionTerminated = function () {
    if (! w2ui.RAActionTerminated) {
        // Terminated form
        $().w2form({
            name: 'RAActionTerminated',
            style: 'background-color: white; display: block;',
            focus: -1,
            formURL: '/webclient/html/raflow/formra-actionterminated.html',
            fields: [
                { field: 'RAActions', type: 'list', width: 120, required: true,
                    options: {
                        items: app.w2ui.listItems.RAActions
                    }
                }
            ],
            onRefresh: function (event) {
                console.log('onRefresh of RAActionTerminated');
                var activeFlowID = app.raflow.activeFlowID;
                var data = app.raflow.data[activeFlowID].Data;
                var raFlags = data.meta.RAFLAGS;
                var raStateString = app.RAStates[parseInt(raFlags & 0xf)];
                $('#RAActionStateLable').text(raStateString);
            }
        });
    }
    // now render the form in specifiec targeted panel
    w2ui.actionLayout.content('main', w2ui.RAActionTerminated);
};

// -------------------------------------------------------------------------------
// Rental Agreement Action - Notice To Move form
// -------------------------------------------------------------------------------
window.loadRAActionNoticeToMove = function () {
    if (! w2ui.RAActionNoticeToMove) {
        // Notice To Move form
        $().w2form({
            name: 'RAActionNoticeToMove',
            style: 'background-color: white; display: block;',
            focus: -1,
            formURL: '/webclient/html/raflow/formra-actionnoticetomove.html',
            fields: [
                { field: 'RAActions', type: 'list', width: 120, required: true,
                    options: {
                        items: app.w2ui.listItems.RAActions
                    }
                },
                { field: 'RANoticeToMoveDate', type: 'date', required: true},
                { field: 'RANoticeToMoveReported', type: 'date', required: true}
            ],
            onChange: function (event) {
                event.done(function(){
                    w2ui.RAActionMoveIn.refresh();
                });

                if(event.target === 'RANoticeToMoveDate') {
                    var date1String = w2ui.RAActionMoveIn.get('RANoticeToMoveDate').el.value;
                    var year1String = date1String.substr(date1String.length - 4);
                    if (year1String <= "1999" ) {
                        this.record.RANoticeToMoveDate = w2uiDateControlString(new Date());
                        this.refresh();
                    }
                }
                if(event.target === 'RANoticeToMoveReported') {
                    var date2String = w2ui.RAActionMoveIn.get('RANoticeToMoveReported').el.value;
                    var year2String = date2String.substr(date2String.length - 4);
                    if (year2String <= "1999" ) {
                        this.record.RANoticeToMoveReported = w2uiDateControlString(new Date());
                        this.refresh();
                    }
                }
            },
            onRefresh: function (event) {
                console.log('onRefresh of RAActionNoticeToMove');
                var activeFlowID = app.raflow.activeFlowID;
                var data = app.raflow.data[activeFlowID].Data;
                var raFlags = data.meta.RAFLAGS;
                var raStateString = app.RAStates[parseInt(raFlags & 0xf)];
                $('#RAActionStateLable').text(raStateString);
            },
            onRender: function (event) {
                w2ui.RAActionNoticeToMove.record = {RAActions: {id: 5, text: "Received Notice To Move"}};
            },
        });
    }
    // now render the form in specifiec targeted panel
    w2ui.actionLayout.content('main', w2ui.RAActionNoticeToMove);
};
