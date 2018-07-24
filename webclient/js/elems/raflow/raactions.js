/*global
    loadRAActionTemplate,
    loadRAActionForm,
    reloadActionForm,
    submitActionForm,
    getSLStringList
*/
"use strict";

// -------------------------------------------------------------------------------
// submitActionForm - submits the data of action form
// @params - FlowID, Decision, Reason, Action
// -------------------------------------------------------------------------------
window.submitActionForm = function(
                        FlowID, Action,
                        Decision1, DeclineReason1,
                        Decision2, DeclineReason2,
                        TerminationReason,
                        DocumentDate,
                        NoticeToMoveDate,
                        NoticeToMoveReported,
                        Mode
                        ) {
    var data = {
        "FlowID": FlowID,
        "Action": Action,
        "Decision1": Decision1,
        "DeclineReason1": DeclineReason1,
        "Decision2": Decision2,
        "DeclineReason2": DeclineReason2,
        "TerminationReason": TerminationReason,
        "DocumentDate": DocumentDate,
        "NoticeToMoveDate": NoticeToMoveDate,
        "NoticeToMoveReported": NoticeToMoveReported,
        "Mode": Mode
    };
    return $.ajax({
        url: "/v1/actions/",
        method: "POST",
        contentType: "application/json",
        data: JSON.stringify(data)
    }).done(function(data) {
        if (data.status === "success") {
            // update the local copy of flow for the active one
            app.raflow.data[data.record.Flow.FlowID] = data.record.Flow;
            w2ui.actionLayout.get('main').content.destroy();

            loadRAActionTemplate();
            setTimeout(function() {
                reloadActionForm();
            },200);
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
            $('button[name=save]').attr('disabled',true);
            break;

        // "Pending Second Approval"
        case 2:
            w2ui.RAActionForm.get('RAApprovalDecision2').hidden = false;
            $('button[name=save]').show();
            $('button[name=save]').attr('disabled',true);
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
                { field: 'RADeclineReason1', type: 'list', width: 120, required: true, hidden: true,
                    options: {
                        items: getSLStringList(getCurrentBID(), "ApplDeny")
                    }
                },
                { field: 'RAApprovalDecision2', type: 'list', width: 120, required: true, hidden: true, 
                    options: {
                        items: [
                            {id: 1, text: "Approve"},
                            {id: 2, text: "Decline"}
                        ]
                    }
                },
                { field: 'RADeclineReason2', type: 'list', width: 120, required: true, hidden: true,
                    options: {
                        items: getSLStringList(getCurrentBID(), "ApplDeny")
                    }
                },
                { field: 'RADocumentDate', type: 'date', hidden: true, options: { start: '01/01/2000' } },
                { field: 'RANoticeToMoveDate', type: 'date', hidden: true, options: { start: '01/01/2000' } },
                { field: 'RANoticeToMoveReported', type: 'date', hidden: true, options: { start: '01/01/2000' } },
                { field: 'RATerminationReason', type: 'list', width: 120, required: true, hidden: true,
                    options: {
                        items: getSLStringList(getCurrentBID(), "WhyLeaving")
                    }
                },
                { field: 'RAActions', type: 'list', width: 120, required: true, options: {items: app.w2ui.listItems.RAActions}}
            ],
            onChange: function (event) {
                event.done(function(){
                    this.refresh();
                    // reloadActionForm();
                });

                switch(event.target) {
                    case 'RAActions':
                        switch (event.value_new.id) {
                            case 5: // Terminate
                                w2ui.RAActionForm.get('RATerminationReason').hidden = false;
                                $('button[name=updateAction]').attr('disabled',true);

                                w2ui.RAActionForm.get('RANoticeToMoveDate').hidden = true;
                                w2ui.RAActionForm.get('RANoticeToMoveReported').hidden = true;
                                break;

                            case 6: // Received Notice-To-Move
                                w2ui.RAActionForm.get('RANoticeToMoveDate').hidden = false;
                                w2ui.RAActionForm.get('RANoticeToMoveReported').hidden = false;

                                w2ui.RAActionForm.get('RATerminationReason').hidden = true;
                                delete this.record.RATerminationReason;
                                $('button[name=updateAction]').attr('disabled',false);
                                break;

                            default:
                                w2ui.RAActionForm.get('RATerminationReason').hidden = true;
                                delete this.record.RATerminationReason;
                                $('button[name=updateAction]').attr('disabled',false);
                                

                                w2ui.RAActionForm.get('RANoticeToMoveDate').hidden = true;
                                w2ui.RAActionForm.get('RANoticeToMoveReported').hidden = true;
                        }
                        break;

                    case 'RAApprovalDecision1':
                        if(event.value_new.text === 'Decline') {
                            $('button[name=save]').attr('disabled',true);
                            w2ui.RAActionForm.get('RADeclineReason1').hidden = false;
                        } else {
                            $('button[name=save]').attr('disabled',false);
                            w2ui.RAActionForm.get('RADeclineReason1').hidden = true;
                            delete this.record.RADeclineReason1;
                        }
                        break;

                    case 'RAApprovalDecision2':
                        if(event.value_new.text === 'Decline') {
                            $('button[name=save]').attr('disabled',true);
                            w2ui.RAActionForm.get('RADeclineReason2').hidden = false;
                        } else {
                            $('button[name=save]').attr('disabled',false);
                            w2ui.RAActionForm.get('RADeclineReason2').hidden = true;
                            delete this.record.RADeclineReason2;
                        }
                        break;

                    case 'RADeclineReason1':
                        if(event.value_new.id === 0) {
                            $('button[name=save]').attr('disabled',true);
                        } else {
                            $('button[name=save]').attr('disabled',false);
                        }
                        break;

                    case 'RADeclineReason2':
                        if(event.value_new.id === 0) {
                            $('button[name=save]').attr('disabled',true);
                        } else {
                            $('button[name=save]').attr('disabled',false);
                        }
                        break;

                    case 'RATerminationReason':
                        if(event.value_new.id === 0) {
                            $('button[name=updateAction]').attr('disabled',true);
                        } else {
                            $('button[name=updateAction]').attr('disabled',false);
                        }
                        break;

                    default:
                        $('button[name=updateAction]').attr('disabled',false);
                        $('button[name=save]').attr('disabled',false);
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

                w2ui.RAActionForm.record = {
                    RAActions: {id: -1, text: "--Select an Action--"},
                };
            },
            actions: {
                save: function() {
                    var FlowID = app.raflow.activeFlowID;
                    var Action = this.record.RAActions.id;
                    var Decision1 = 0;
                    var DeclineReason1 = 0;
                    var Decision2 = 0;
                    var DeclineReason2 = 0;
                    var TerminationReason = 0;
                    var DocumentDate = "1/1/1900";
                    var NoticeToMoveDate = "1/1/1900";
                    var NoticeToMoveReported = "1/1/1900";
                    var Mode = "State";

                    submitActionForm(
                        FlowID, Action,
                        Decision1, DeclineReason1,
                        Decision2, DeclineReason2,
                        TerminationReason,
                        DocumentDate,
                        NoticeToMoveDate,
                        NoticeToMoveReported,
                        Mode
                    );
                },
                updateAction: function() {
                    var FlowID = app.raflow.activeFlowID;
                    var Action = this.record.RAActions.id;
                    var Decision1 = 0;
                    var DeclineReason1 = 0;
                    var Decision2 = 0;
                    var DeclineReason2 = 0;
                    var TerminationReason = 0;
                    var DocumentDate = "1/1/1900";
                    var NoticeToMoveDate = "1/1/1900";
                    var NoticeToMoveReported = "1/1/1900";
                    var Mode = "Action";

                    if( Action === -1 ) {
                        return;
                    }

                    if( Action === 5 ) {
                        TerminationReason = this.record.RATerminationReason.id;
                    }

                    if( Action === 5 && this.record.RATerminationReason.id === 0) {
                        return;
                    }

                    var currentState = parseInt(app.raflow.data[FlowID].Data.meta.RAFLAGS & (0xf));
                    if (Action === currentState) {
                        return;
                    }

                    submitActionForm(
                        FlowID, Action,
                        Decision1, DeclineReason1,
                        Decision2, DeclineReason2,
                        TerminationReason,
                        DocumentDate,
                        NoticeToMoveDate,
                        NoticeToMoveReported,
                        Mode
                    );
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
