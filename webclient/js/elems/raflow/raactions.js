/*global
    loadRAActionInProgress,
    loadRAActionFirstApproval,
    loadRAActionSecondApproval,
    loadRAActionMoveIn,
    loadRAActionActive,
    loadRAActionTerminated,
    loadRAActionNoticeToMove,
    loadRAActionTemplate,
    loadActionFormByState,
    submitActionForm,
    getSLStringList,
    loadRAActionForm,
    reloadActionForm
*/
"use strict";

// -------------------------------------------------------------------------------
// submitActionForm - submits the data of action form
// @params - FlowID, Decision, Reason, Action
// -------------------------------------------------------------------------------
window.submitActionForm = function(FlowID, Decision, Reason, Action) {
    var data = {"FlowID": FlowID, "Decision": Decision, "Reason":Reason, "Action":Action};
    return $.ajax({
        url: "/v1/actions/",
        method: "POST",
        contentType: "application/json",
        data: JSON.stringify(data)
    }).done(function(data) {
        if (data.status === "success") {
            // update the local copy of flow for the active one
            app.raflow.data[data.record.FlowID] = data.record.Flow;
            w2ui.actionLayout.get('main').content.destroy();

            loadRAActionTemplate();
            reloadActionForm();
        } else {
            //Display Error
        }
    });
};

// -------------------------------------------------------------------------------
// reloadActionForm - reloads the data of action form according to state
// -------------------------------------------------------------------------------
window.reloadActionForm = function() {
    console.log('Custom Reload Function...');
    $('#RAActionRAInfo').hide();
    $('#RAActionTerminatedRAInfo').hide();
    $('#RAActionNoticeToMoveInfo').hide();

    $('button[name=RAGenerateRAForm]').hide();
    $('button[name=RAGenerateMoveInInspectionForm]').hide();
    $('button[name=RAGenerateMoveOutForm]').hide();
    $('button[name=save]').hide();

    var activeFlowID = app.raflow.activeFlowID;
    var data = app.raflow.data[activeFlowID].Data;
    var raFlags = data.meta.RAFLAGS;

    switch (parseInt(raFlags & 0xf)) {
        // "Application Being Completed"
        case 0:
            break;

        // "Pending First Approval"
        case 1:
            w2ui.RAActionForm.get('RAApprovalDecision1').hidden = false;
            $('button[name=save]').show();
            break;

        // "Pending Second Approval"
        case 2:
            w2ui.RAActionForm.get('RAApprovalDecision2').hidden = false;
            $('button[name=save]').show();
            break;

        // "Move-In / Execute Modification"
        case 3:
            w2ui.RAActionForm.get('RADocumentDate').hidden = false;
            $('button[name=RAGenerateRAForm]').show();
            $('button[name=RAGenerateMoveInInspectionForm]').show();
            $('button[name=save]').show();
            break;

        // "Active"
        case 4:
            $('#RAActionRAInfo').show();
            break;

        // "Terminated"
        case 5:
            $('#RAActionTerminatedRAInfo').show();
            $('button[name=RAGenerateRAForm]').show();
            break;

        // "Notice To Move"
        case 6:
            $('#RAActionNoticeToMoveInfo').show();
            break;

        default:
    }
    w2ui.RAActionForm.refresh();
};

//------------------------------------------------------------------------
// loadActionFormByState - It loads the Action Form on basis of the state
//                         value provided.
//
// @params
//      raState = State value according to RAFLAGS.
//------------------------------------------------------------------------
window.loadActionFormByState = function(raState) {
    switch (raState) {
        // "Application Being Completed"
        case 0:
            console.log('Inside Application Being Completed switch case');
            loadRAActionInProgress();
            break;

        // "Pending First Approval"
        case 1:
            console.log('Inside Pending First Approval switch case');
            loadRAActionFirstApproval();
            break;

        // "Pending Second Approval"
        case 2:
            console.log('Inside Pending Second Approval switch case');
            loadRAActionSecondApproval();
            break;

        // "Move-In / Execute Modification"
        case 3:
            console.log('Inside Move-In / Execute Modification switch case');
            loadRAActionMoveIn();
            break;

        // "Active"
        case 4:
            console.log('Inside Active switch case');
            loadRAActionActive();
            break;

        // "Terminated"
        case 5:
            console.log('Inside Terminated switch case');
            loadRAActionTerminated();
            break;

        // "Notice To Move"
        case 6:
            console.log('Inside Notice To Move switch case');
            loadRAActionNoticeToMove();
            break;

        default:
    }
};


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
            onRefresh: function(event) {

                event.onComplete = function() {
                    var activeFlowID = app.raflow.activeFlowID;
                    var data = app.raflow.data[activeFlowID];
                    var meta = data.Data.meta;

                    // Header Part
                    var x = document.getElementById("bannerRAID");
                    if (x !== null) {
                        if (meta.RAID == 0) {
                            x.innerHTML = 'New Rental Agreement';
                        } else {
                            x.innerHTML = '' + meta.RAID;
                        }
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

                    // Footer Part
                    x = document.getElementById("bannerApprover1");
                    if (x !== null) {
                        if (meta.Approver1 == 0) {
                            x.innerHTML = 'Pending';
                        } else {
                            if (meta.DeclineReason1 == 0) {
                                x.innerHTML = 'Approved by ' + meta.Approver1 + 'on ' + meta.DecisionDate1;
                            } else{
                                x.innerHTML = 'Declined by ' + meta.Approver1 + 'on ' + meta.DecisionDate1 + 'Reason: ' + meta.DeclineReason1;
                            }
                        }
                    }

                    x = document.getElementById("bannerApprover2");
                    if (x !== null) {
                        if (meta.Approver2 == 0) {
                            x.innerHTML = 'Pending';
                        } else {
                            if (meta.DeclineReason2 == 0) {
                                x.innerHTML = 'Approved by ' + meta.Approver2 + 'on ' + meta.DecisionDate2;
                            } else{
                                x.innerHTML = 'Declined by ' + meta.Approver2 + 'on ' + meta.DecisionDate2 + 'Reason: ' + meta.DeclineReason2;
                            }
                        }
                    }
                };
            },
        });
    }
    w2ui.newraLayout.content('right', w2ui.actionLayout);

    w2ui.actionLayout.load('top', '/webclient/html/raflow/formra-actionheader.html');
    w2ui.actionLayout.load('bottom', '/webclient/html/raflow/formra-actionfooter.html');

    var raFlags = app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS;
    var raState = parseInt(raFlags & 0xf);

    // loadActionFormByState(raState);
    loadRAActionForm();

    w2ui.newraLayout.show('right', true);
    w2ui.newraLayout.sizeTo('right', 950);
};

// -------------------------------------------------------------------------------
// Rental Agreement Action Form
// -------------------------------------------------------------------------------
window.loadRAActionForm = function() {
    if(! w2ui.RAActionForm) {
        $().w2form({
            name: 'RAActionForm',
            style: 'background-color: white; display: block;',
            focus: -1,
            url: 'v1/actions/',
            formURL: '/webclient/html/raflow/formra-actionmain.html',
            fields: [
                { field: 'RAApprovalDecision1', type: 'list', width: 120, required: true, hidden: true,
                    options: {
                        items: [
                            {id: 1, text: "Approve"},
                            {id: 2, text: "Decline"}
                        ]
                    }
                },
                { field: 'RADeclineReason1', type: 'list', width: 120, required: true, hidden: true, options: {} },
                { field: 'RAApprovalDecision2', type: 'list', width: 120, required: true, hidden: true, 
                    options: {
                        items: [
                            {id: 1, text: "Approve"},
                            {id: 2, text: "Decline"}
                        ]
                    }
                },
                { field: 'RADeclineReason2', type: 'list', width: 120, required: true, hidden: true, options: {} },
                { field: 'RADocumentDate', type: 'date', hidden: true, options: { start: '01/01/2000' } },
                { field: 'RANoticeToMoveDate', type: 'date', hidden: true, options: { start: '01/01/2000' } },
                { field: 'RANoticeToMoveReported', type: 'date', hidden: true, options: { start: '01/01/2000' } },
                { field: 'RATerminationReason', type: 'list', width: 120, required: true, hidden: true, options: {} },
                { field: 'RAActions', type: 'list', width: 120, required: true, options: {items: app.w2ui.listItems.RAActions}}
            ],
            onChange: function (event) {
                event.done(function(){
                    w2ui.RAActionForm.refresh();
                });

                if(event.target === 'RAActions') {
                    switch (event.value_new.text) {
                        case 'Terminate':
                            w2ui.RAActionForm.get('RATerminationReason').hidden = false;

                            w2ui.RAActionForm.get('RANoticeToMoveDate').hidden = true;
                            w2ui.RAActionForm.get('RANoticeToMoveReported').hidden = true;
                            break;

                        case 'Received Notice To Move':
                            w2ui.RAActionForm.get('RANoticeToMoveDate').hidden = false;
                            w2ui.RAActionForm.get('RANoticeToMoveReported').hidden = false;

                            w2ui.RAActionForm.get('RATerminationReason').hidden = true;
                            delete this.record.RATerminationReason;
                            break;

                        default:
                            w2ui.RAActionForm.get('RATerminationReason').hidden = true;
                            delete this.record.RATerminationReason;

                            w2ui.RAActionForm.get('RANoticeToMoveDate').hidden = true;
                            w2ui.RAActionForm.get('RANoticeToMoveReported').hidden = true;

                            this.clear();
                    }
                }

                if(event.target === 'RAApprovalDecision1') {
                    if(event.value_new.text === 'Decline') {
                        w2ui.RAActionForm.get('RADeclineReason1').hidden = false;
                    } else {
                        w2ui.RAActionForm.get('RADeclineReason1').hidden = true;
                        delete this.record.RADeclineReason1;
                    }
                    w2ui.RAActionForm.refresh();
                }

                if(event.target === 'RAApprovalDecision2') {
                    if(event.value_new.text === 'Decline') {
                        w2ui.RAActionForm.get('RADeclineReason2').hidden = false;
                    } else {
                        w2ui.RAActionForm.get('RADeclineReason2').hidden = true;
                        delete this.record.RADeclineReason2;
                    }
                    w2ui.RAActionForm.refresh();
                }
            },
            onRefresh: function (event) {
                console.log('onRefresh of RAActionForm');
                var activeFlowID = app.raflow.activeFlowID;
                var data = app.raflow.data[activeFlowID].Data;
                var raFlags = data.meta.RAFLAGS;
                var raStateString = app.RAStates[parseInt(raFlags & 0xf)];

                // var RAID = data.meta.RAID;
                // if(RAID > 0 &&  (raStateString === "Pending First Approval" || raStateString === "Pending Second Approval")) {
                //     raStateString = 'Modification ' + raStateString;
                // }

                $('#RAActionStateLable').text(raStateString);
            },
            onRender: function (event) {
                console.log('onRender of RAActionForm');

                event.done(function(){
                    var BID = getCurrentBID();
                    this.get('RADeclineReason1').options.items = getSLStringList(BID, "ApplDeny");
                    this.get('RADeclineReason2').options.items = getSLStringList(BID, "ApplDeny");
                    this.get('RATerminationReason').options.items = getSLStringList(BID, "WhyLeaving");
                });

                w2ui.RAActionForm.record = {
                    RAActions: {id: -1, text: "--Select an Action--"},
                };
            },
            actions: {
                save: function() {
                    var FlowID = app.raflow.activeFlowID;
                    var Decision = 0;
                    var Reason = 0;
                    var Action = this.record.RAActions.id;
                    submitActionForm(FlowID, Decision, Reason, Action);
                },
                updateAction: function() {
                    var FlowID = app.raflow.activeFlowID;
                    var Decision = 0;
                    var Reason = 0;
                    var Action = this.record.RAActions.id;
                    submitActionForm(FlowID, Decision, Reason, Action);
                }
            }
        });
    }
    // now render the form in specifiec targeted panel
    w2ui.actionLayout.content('main', w2ui.RAActionForm);
    setTimeout(function() {
        reloadActionForm();
    }, 100);
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
            url: 'v1/actionss/',
            formURL: '/webclient/html/raflow/formra-actioninprogress.html',
            fields: [
                { field: 'RAActions', type: 'list', width: 120, required: true, options: {items: app.w2ui.listItems.RAActions}},
                { field: 'RATerminationReason', type: 'list', width: 120, required: true, hidden: true, 
                    options: {}
                }
            ],
            onChange: function (event) {
                event.done(function(){
                    w2ui.RAActionInProgress.refresh();
                });

                if(event.target === 'RAActions') {
                    switch (event.value_new.text) {
                        case 'Terminate':
                            w2ui.RAActionInProgress.get('RATerminationReason').hidden = false;
                            break;
                        default:
                            w2ui.RAActionInProgress.get('RATerminationReason').hidden = true;
                            delete this.record.RATerminationReason;

                            this.clear();
                    }
                }
            },
            onRefresh: function (event) {
                console.log('onRefresh of RAActionInProgress');
                var activeFlowID = app.raflow.activeFlowID;
                var data = app.raflow.data[activeFlowID].Data;
                var raFlags = data.meta.RAFLAGS;
                var raStateString = app.RAStates[parseInt(raFlags & 0xf)];
                $('#RAActionStateLable').text(raStateString);

            },
            onRender: function (event) {
                event.done(function(){
                    var BID = getCurrentBID();
                    this.get('RATerminationReason').options.items = getSLStringList(BID, "WhyLeaving");
                });
                console.log('onRender of RAActionInProgress');
                w2ui.RAActionInProgress.record = {
                    RAActions: {id: -1, text: "--Select an Action--"},
                };
            },
            actions: {
                save: function () {
                    if( this.record.RAActions.id === -1 ) {
                        return;
                    }

                    if( this.record.RATerminationReason != undefined && this.record.RATerminationReason.id === 0) {
                        return;
                    }

                    var FlowID = app.raflow.activeFlowID;
                    var Decision = 0;
                    var Reason = 0;
                    var Action = this.record.RAActions.id;
                    submitActionForm(FlowID, Decision, Reason, Action);

                    // var actionn = w2ui.RAActionInProgress.record.RAActions.text;
                    // var raFlags = app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS;
                    // raFlags = raFlags & ~(0xf);

                    // switch (actionn) {
                    //     case "Edit Rental Agreement Information" :
                    //         raFlags = raFlags | 0;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    //     case "Authorize First Approval" :
                    //         raFlags = raFlags | 1;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    //     case "Authorize Second Approval" :
                    //         raFlags = raFlags | 2;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    //     case "Complete Move In" :
                    //         raFlags = raFlags | 3;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    //     case "Terminate" :
                    //         if(this.record.RATerminationReason != undefined) {
                    //             raFlags = raFlags | 5;
                    //             app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         }else {
                    //             return;
                    //         }
                    //         break;
                    //     case "Received Notice To Move" :
                    //         raFlags = raFlags | 6;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    //     default:
                    // }

                    // loadRAActionTemplate();
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
                { field: 'RAApprovalDecision1', type: 'list', width: 120, required: true, hidden: true,
                    options: {
                        items: [{id: 1, text: "Approve"}, {id: 2, text: "Decline"}]
                    }
                },
                { field: 'RADeclineReason1', type: 'list', width: 120, required: true, hidden: true,
                    options: {
                        items: [{id: 1, text: 'Temp1'}, {id: 2, text: 'Temp2'}]
                    }
                },
                { field: 'RATerminationReason', type: 'list', width: 120, required: true, hidden: true, 
                    options: {}
                }
            ],
            onChange: function (event) {
                event.done(function(){
                    w2ui.RAActionFirstApproval.refresh();
                });

                if(event.target === 'RAActions') {
                    switch (event.value_new.text) {
                        case 'Authorize First Approval':
                            w2ui.RAActionFirstApproval.get('RAApprovalDecision1').hidden = false;
                            w2ui.RAActionFirstApproval.get('RATerminationReason').hidden = true;
                            delete this.record.RATerminationReason;
                            break;
                        case 'Terminate':
                            w2ui.RAActionFirstApproval.get('RATerminationReason').hidden = false;

                            w2ui.RAActionFirstApproval.get('RAApprovalDecision1').hidden = true;
                            w2ui.RAActionFirstApproval.get('RADeclineReason1').hidden = true;
                            delete this.record.RAApprovalDecision1;
                            delete this.record.RADeclineReason1;
                            break;
                        default:
                            w2ui.RAActionFirstApproval.get('RAApprovalDecision1').hidden = true;
                            w2ui.RAActionFirstApproval.get('RADeclineReason1').hidden = true;
                            w2ui.RAActionFirstApproval.get('RATerminationReason').hidden = true;
                            delete this.record.RAApprovalDecision1;
                            delete this.record.RADeclineReason1;
                            delete this.record.RATerminationReason;
                            this.clear();
                    }
                }
                if(event.target === 'RAApprovalDecision1') {
                    if(event.value_new.text === 'Decline') {
                        w2ui.RAActionFirstApproval.get('RADeclineReason1').hidden = false;
                    } else {
                        w2ui.RAActionFirstApproval.get('RADeclineReason1').hidden = true;
                        delete this.record.RADeclineReason1;
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
                event.done(function(){
                    var BID = getCurrentBID();
                    this.get('RADeclineReason1').options.items = getSLStringList(BID, "ApplDeny");
                    this.get('RATerminationReason').options.items = getSLStringList(BID, "WhyLeaving");
                });
                console.log('onRender of RAActionFirstApproval');
                w2ui.RAActionFirstApproval.record = {
                    RAActions: {id: -1, text: "--Select an Action--"}
                };
            },
            actions: {
                save: function () {
                    if( this.record.RAActions.id === -1) {
                        return;
                    }

                    if( this.record.RADeclineReason1 != undefined && this.record.RADeclineReason1.id === 0) {
                        return;
                    }

                    if( !(this.record.RATerminationReason && this.record.RATerminationReason.id > 0)) {
                        return;
                    }

                    var FlowID = app.raflow.activeFlowID;
                    var Decision = (this.record.RAApprovalDecision1 != undefined) ? this.record.RAApprovalDecision1 : 0;
                    var Reason = (this.record.RADeclineReason1 != undefined) ? this.record.RADeclineReason1 : 0;
                    var Action = this.record.RAActions.id;
                    submitActionForm(FlowID, Decision, Reason, Action);

                    // var actionn = w2ui.RAActionFirstApproval.record.RAActions.text;
                    // var raFlags = app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS;
                    // raFlags = raFlags & ~(0xf);

                    // switch (actionn) {
                    //     case "Edit Rental Agreement Information" :
                    //         raFlags = raFlags | 0;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    //     case "Authorize First Approval" :
                    //         if(this.record.RAApprovalDecision1 != undefined) {
                    //             if (this.record.RAApprovalDecision1.text === 'Decline' && this.record.RADeclineReason1 === undefined) {
                    //                 return;
                    //             }
                    //             raFlags = raFlags | 2;
                    //             app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         }else {
                    //             return;
                    //         }
                    //         break;
                    //     case "Authorize Second Approval" :
                    //         raFlags = raFlags | 2;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    //     case "Complete Move In" :
                    //         raFlags = raFlags | 3;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    //     case "Terminate" :
                    //         if(this.record.RATerminationReason != undefined) {
                    //             raFlags = raFlags | 5;
                    //             app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         }else {
                    //             return;
                    //         }
                    //         break;
                    //     case "Received Notice To Move" :
                    //         raFlags = raFlags | 6;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    // }
                    // w2ui.actionLayout.get('main').content.destroy();
                    // loadRAActionTemplate();
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
                { field: 'RAApprovalDecision2', type: 'list', width: 120, required: true, hidden: true,
                    options: {
                        items: [{id: 1, text: "Approve"}, {id: 2, text: "Decline"}]
                    }
                },
                { field: 'RADeclineReason2', type: 'list', width: 120, required: true, hidden: true,
                    options: {
                        items: [{id: 1, text: 'Temp1'}, {id: 2, text: 'Temp2'}]
                    }
                },
                { field: 'RATerminationReason', type: 'list', width: 120, required: true, hidden: true, 
                    options: {}
                }
            ],
            onChange: function (event) {
                event.done(function(){
                    w2ui.RAActionSecondApproval.refresh();
                });

                if(event.target === 'RAActions') {
                    switch (event.value_new.text) {
                        case 'Authorize Second Approval':
                            w2ui.RAActionSecondApproval.get('RAApprovalDecision2').hidden = false;
                            w2ui.RAActionSecondApproval.get('RATerminationReason').hidden = true;
                            delete this.record.RATerminationReason;
                            break;
                        case 'Terminate':
                            w2ui.RAActionSecondApproval.get('RATerminationReason').hidden = false;

                            w2ui.RAActionSecondApproval.get('RAApprovalDecision2').hidden = true;
                            w2ui.RAActionSecondApproval.get('RADeclineReason2').hidden = true;
                            delete this.record.RAApprovalDecision2;
                            delete this.record.RADeclineReason2;
                            break;
                        default:
                            w2ui.RAActionSecondApproval.get('RAApprovalDecision2').hidden = true;
                            w2ui.RAActionSecondApproval.get('RADeclineReason2').hidden = true;
                            w2ui.RAActionSecondApproval.get('RATerminationReason').hidden = true;
                            delete this.record.RAApprovalDecision2;
                            delete this.record.RADeclineReason2;
                            delete this.record.RATerminationReason;
                            this.clear();
                    }
                }
                if(event.target === 'RAApprovalDecision2') {
                    if(event.value_new.text === 'Decline') {
                        w2ui.RAActionSecondApproval.get('RADeclineReason2').hidden = false;
                    } else {
                        w2ui.RAActionSecondApproval.get('RADeclineReason2').hidden = true;
                        delete this.record.RADeclineReason2;
                    }
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
                event.done(function(){
                    var BID = getCurrentBID();
                    this.get('RADeclineReason2').options.items = getSLStringList(BID, "ApplDeny");
                    this.get('RATerminationReason').options.items = getSLStringList(BID, "WhyLeaving");
                });
                console.log('onRender of RAActionSecondApproval');
                w2ui.RAActionSecondApproval.record = {
                    RAActions: {id: -1, text: "--Select an Action--"}
                };
            },
            actions: {
                save: function () {
                    if( this.record.RAActions.id === -1 ) {
                        return;
                    }

                    if( this.record.RADeclineReason2 != undefined && this.record.RADeclineReason2.id === 0) {
                        return;
                    }

                    if( this.record.RATerminationReason != undefined && this.record.RATerminationReason.id === 0) {
                        return;
                    }

                    var FlowID = app.raflow.activeFlowID;
                    var Decision = (this.record.RAApprovalDecision2 != undefined) ? this.record.RAApprovalDecision2 : 0;
                    var Reason = (this.record.RADeclineReason2 != undefined) ? this.record.RADeclineReason2 : 0;
                    var Action = this.record.RAActions.id;
                    submitActionForm(FlowID, Decision, Reason, Action);

                    // var actionn = w2ui.RAActionSecondApproval.record.RAActions.text;
                    // var raFlags = app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS;
                    // raFlags = raFlags & ~(0xf);

                    // switch (actionn) {
                    //     case "Edit Rental Agreement Information" :
                    //         raFlags = raFlags | 0;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    //     case "Authorize First Approval" :
                    //         raFlags = raFlags | 1;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    //     case "Authorize Second Approval" :
                    //         if(this.record.RAApprovalDecision2 != undefined) {
                    //             if (this.record.RAApprovalDecision2.text === 'Decline' && this.record.RADeclineReason2 === undefined) {
                    //                 return;
                    //             }
                    //             raFlags = raFlags | 3;
                    //             app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         }else {
                    //             return;
                    //         }
                    //         break;
                    //     case "Complete Move In" :
                    //         raFlags = raFlags | 3;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    //     case "Terminate" :
                    //         if(this.record.RATerminationReason != undefined) {
                    //             raFlags = raFlags | 5;
                    //             app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         }else {
                    //             return;
                    //         }
                    //         break;
                    //     case "Received Notice To Move" :
                    //         raFlags = raFlags | 6;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    // }
                    // w2ui.actionLayout.get('main').content.destroy();
                    // loadRAActionTemplate();
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
                { field: 'RADocumentDate', type: 'date', hidden: true,
                    options: { start: '01/01/2000' }
                },
                { field: 'RATerminationReason', type: 'list', width: 120, required: true, hidden: true, 
                    options: {}
                }
            ],
            onChange: function (event) {
                event.done(function(){
                    w2ui.RAActionMoveIn.refresh();
                });

                if(event.target === 'RAActions') {
                    switch (event.value_new.text) {
                        case 'Complete Move In':
                            w2ui.RAActionMoveIn.get('RADocumentDate').hidden = false;
                            $('[name="RAGenerateRAForm"]').show();
                            $('[name="RAGenerateMoveInInspectionForm"]').show();

                            w2ui.RAActionMoveIn.get('RATerminationReason').hidden = true;
                            delete this.record.RATerminationReason;
                            break;
                        case 'Terminate':
                            w2ui.RAActionMoveIn.get('RATerminationReason').hidden = false;

                            w2ui.RAActionMoveIn.get('RADocumentDate').hidden = true;
                            $('[name="RAGenerateRAForm"]').hide();
                            $('[name="RAGenerateMoveInInspectionForm"]').hide();
                            break;
                        default:
                            w2ui.RAActionMoveIn.get('RADocumentDate').hidden = true;
                            $('[name="RAGenerateRAForm"]').hide();
                            $('[name="RAGenerateMoveInInspectionForm"]').hide();
                            w2ui.RAActionMoveIn.get('RATerminationReason').hidden = true;
                            delete this.record.RATerminationReason;

                            this.clear();
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
                event.done(function(){
                    var BID = getCurrentBID();
                    this.get('RATerminationReason').options.items = getSLStringList(BID, "WhyLeaving");
                    $('[name="RAGenerateRAForm"]').hide();
                    $('[name="RAGenerateMoveInInspectionForm"]').hide();
                });

                console.log('onRender of RAActionMoveIn');
                w2ui.RAActionMoveIn.record = {
                    RAActions: {id: -1, text: "--Select an Action--"}
                };
            },
            actions: {
                save: function () {
                    if( this.record.RAActions.id === -1 ) {
                        return;
                    }

                    if( this.record.RATerminationReason != undefined && this.record.RATerminationReason.id === 0) {
                        return;
                    }

                    var FlowID = app.raflow.activeFlowID;
                    var Decision = 0;
                    var Reason = 0;
                    var Action = this.record.RAActions.id;
                    submitActionForm(FlowID, Decision, Reason, Action);

                    // var actionn = w2ui.RAActionMoveIn.record.RAActions.text;
                    // var raFlags = app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS;
                    // raFlags = raFlags & ~(0xf);

                    // switch (actionn) {
                    //     case "Edit Rental Agreement Information" :
                    //         raFlags = raFlags | 0;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    //     case "Authorize First Approval" :
                    //         raFlags = raFlags | 1;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    //     case "Authorize Second Approval" :
                    //         raFlags = raFlags | 2;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    //     case "Complete Move In" :
                    //         raFlags = raFlags | 4;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    //     case "Terminate" :
                    //         if(this.record.RATerminationReason != undefined) {
                    //             raFlags = raFlags | 5;
                    //             app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         }else {
                    //             return;
                    //         }
                    //         break;
                    //     case "Received Notice To Move" :
                    //         raFlags = raFlags | 6;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    // }
                    // w2ui.actionLayout.get('main').content.destroy();
                    // loadRAActionTemplate();
                },
                RAGenerateRAForm: function() {
                },
                RAGenerateMoveInInspectionForm: function() {
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
                    options: {}
                }
            ],
            onChange: function (event) {
                event.done(function(){
                    w2ui.RAActionActive.refresh();
                });

                if(event.target === 'RAActions') {
                    switch (event.value_new.text) {
                        case 'Terminate':
                            w2ui.RAActionActive.get('RATerminationReason').hidden = false;
                            break;
                        default:
                            w2ui.RAActionActive.get('RATerminationReason').hidden = true;
                            delete this.record.RATerminationReason;

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
            onRender: function (event) {
                event.done(function(){
                    var BID = getCurrentBID();
                    this.get('RATerminationReason').options.items = getSLStringList(BID, "WhyLeaving");
                });
                console.log('onRender of RAActionActive');
                w2ui.RAActionActive.record = {
                    RAActions: {id: -1, text: "--Select an Action--"}
                };
            },
            actions: {
                save: function () {
                    if( this.record.RAActions.id === -1 ) {
                        return;
                    }

                    if( this.record.RATerminationReason != undefined && this.record.RATerminationReason.id === 0) {
                        return;
                    }

                    var FlowID = app.raflow.activeFlowID;
                    var Decision = 0;
                    var Reason = 0;
                    var Action = this.record.RAActions.id;
                    submitActionForm(FlowID, Decision, Reason, Action);

                    // var actionn = w2ui.RAActionActive.record.RAActions.text;
                    // var raFlags = app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS;
                    // raFlags = raFlags & ~(0xf);

                    // switch (actionn) {
                    //     case "Edit Rental Agreement Information" :
                    //         raFlags = raFlags | 0;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    //     case "Authorize First Approval" :
                    //         raFlags = raFlags | 1;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    //     case "Authorize Second Approval" :
                    //         raFlags = raFlags | 2;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    //     case "Complete Move In" :
                    //         raFlags = raFlags | 3;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    //     case "Terminate" :
                    //         if(this.record.RATerminationReason != undefined) {
                    //             raFlags = raFlags | 5;
                    //             app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         }else {
                    //             return;
                    //         }
                    //         break;
                    //     case "Received Notice To Move" :
                    //         raFlags = raFlags | 6;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
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
            },
            onRender: function (event) {
                console.log('onRender of RAActionTerminated');
                w2ui.RAActionTerminated.record = {
                    RAActions: {id: -1, text: "--Select an Action--"}
                };
            },
            actions: {
                save: function () {
                    if( this.record.RAActions.id === -1 ) {
                        return;
                    }

                    if( this.record.RATerminationReason != undefined && this.record.RATerminationReason.id === 0) {
                        return;
                    }

                    var FlowID = app.raflow.activeFlowID;
                    var Decision = 0;
                    var Reason = 0;
                    var Action = this.record.RAActions.id;
                    submitActionForm(FlowID, Decision, Reason, Action);

                    // var actionn = w2ui.RAActionTerminated.record.RAActions.text;
                    // var raFlags = app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS;
                    // raFlags = raFlags & ~(0xf);

                    // switch (actionn) {
                    //     case "Edit Rental Agreement Information" :
                    //         raFlags = raFlags | 0;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    //     case "Authorize First Approval" :
                    //         raFlags = raFlags | 1;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    //     case "Authorize Second Approval" :
                    //         raFlags = raFlags | 2;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    //     case "Complete Move In" :
                    //         raFlags = raFlags | 3;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    //     case "Terminate" :
                    //         return;
                    //     case "Received Notice To Move" :
                    //         raFlags = raFlags | 6;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    // }
                    // w2ui.actionLayout.get('main').content.destroy();
                    // loadRAActionTemplate();
                },
                RAGenerateMoveOutForm: function () {

                }
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
                { field: 'RANoticeToMoveDate', type: 'date', hidden: true,
                    options: { start: '01/01/2000' }
                },
                { field: 'RANoticeToMoveReported', type: 'date', hidden: true,
                    options: { start: '01/01/2000' }
                },
                { field: 'RATerminationReason', type: 'list', width: 120, required: true, hidden: true, 
                    options: {}
                }
            ],
            onChange: function (event) {
                event.done(function(){
                    w2ui.RAActionNoticeToMove.refresh();
                });

                if(event.target === 'RAActions') {
                    switch (event.value_new.text) {
                        case 'Received Notice To Move':
                            w2ui.RAActionNoticeToMove.get('RANoticeToMoveDate').hidden = false;
                            w2ui.RAActionNoticeToMove.get('RANoticeToMoveReported').hidden = false;

                            w2ui.RAActionNoticeToMove.get('RATerminationReason').hidden = true;
                            delete this.record.RATerminationReason;
                            break;
                        case 'Terminate':
                            w2ui.RAActionNoticeToMove.get('RATerminationReason').hidden = false;
                            w2ui.RAActionNoticeToMove.get('RANoticeToMoveDate').hidden = true;
                            w2ui.RAActionNoticeToMove.get('RANoticeToMoveReported').hidden = true;
                            break;
                        default:
                            w2ui.RAActionNoticeToMove.get('RANoticeToMoveDate').hidden = true;
                            w2ui.RAActionNoticeToMove.get('RANoticeToMoveReported').hidden = true;

                            w2ui.RAActionNoticeToMove.get('RATerminationReason').hidden = true;
                            delete this.record.RATerminationReason;

                            this.clear();
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
                event.done(function(){
                    var BID = getCurrentBID();
                    this.get('RATerminationReason').options.items = getSLStringList(BID, "WhyLeaving");
                });
                console.log('onRender of RAActionTerminated');
                w2ui.RAActionNoticeToMove.record = {
                    RAActions: {id: -1, text: "--Select an Action--"}
                };
            },
            actions: {
                save: function () {
                    if( this.record.RAActions.id === -1 ) {
                        return;
                    }

                    if( this.record.RATerminationReason != undefined && this.record.RATerminationReason.id === 0) {
                        return;
                    }

                    var FlowID = app.raflow.activeFlowID;
                    var Decision = 0;
                    var Reason = 0;
                    var Action = this.record.RAActions.id;
                    submitActionForm(FlowID, Decision, Reason, Action);
                    
                    // var actionn = w2ui.RAActionNoticeToMove.record.RAActions.text;
                    // var raFlags = app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS;
                    // raFlags = raFlags & ~(0xf);

                    // switch (actionn) {
                    //     case "Edit Rental Agreement Information" :
                    //         raFlags = raFlags | 0;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    //     case "Authorize First Approval" :
                    //         raFlags = raFlags | 1;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    //     case "Authorize Second Approval" :
                    //         raFlags = raFlags | 2;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    //     case "Complete Move In" :
                    //         raFlags = raFlags | 3;
                    //         app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         break;
                    //     case "Terminate" :
                    //         if(this.record.RATerminationReason != undefined) {
                    //             raFlags = raFlags | 5;
                    //             app.raflow.data[app.raflow.activeFlowID].Data.meta.RAFLAGS = raFlags;
                    //         }else {
                    //             return;
                    //         }
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
    w2ui.actionLayout.content('main', w2ui.RAActionNoticeToMove);
};
