/*global
    loadRAActionTemplate,
    loadRAActionForm,
    reloadActionForm,
    submitActionForm,
    getSLStringList,
    refreshLabels,
    GetVehicleIdentity,
    dtFormatISOToW2ui,
    localtimeToUTC,
    UpdateRAFlowLocalData,
    GetCurrentFlowID, CloseRAFlowLayout,
    ChangeRAFlowVersionToolbar,
    displayErrorDot,
    displayActiveComponentError,
    w2uiUTCDateControlString, RAFlowAJAX
*/
"use strict";

var actionsUI = {
    hdrHeight: 122,
    ftrHeight: 150
};

// -------------------------------------------------------------------------------
// submitActionForm - submits the data of action form
// @params - data
// -------------------------------------------------------------------------------
window.submitActionForm = function(data) {
    var BID         = getCurrentBID(),
        FlowID      = GetCurrentFlowID();

    var url = "/v1/raactions/" + BID.toString() + "/" + FlowID.toString() + "/";

    return RAFlowAJAX(url, "POST", data, false)
    .done(function(data) {
        if (data.status !== "success") {
            console.error(data.message);
            w2ui.RAActionForm.error(data.message);
            return;
        }

        switch(true) {
            case (data.record.Flow.FlowID === -1):
                w2ui.RAActionForm.error("Flow Already Exists");
                return false;
            case (data.record.Flow.FlowID === 0):
                if (app.raflow.version === 'refno') {
                    // load ActionForm and Toolbar for raid version
                    w2ui.newraLayout.content('right', '');
                    w2ui.newraLayout.hide('right', true);
                    app.raflow.version = 'raid'; // AS IT WAS MIGRATED
                }

                // Update flow local copy and green checks
                UpdateRAFlowLocalData(data, true);
                break;
            case (data.record.Flow.FlowID > 0):

                // FlowID > 0 that means it is refno version
                app.raflow.version = 'refno';

                // Update flow local copy and green checks
                UpdateRAFlowLocalData(data, true);

                // validation errors based on validation check
                app.raflow.validationErrors = {
                    dates: app.raflow.validationCheck.errors.dates.total > 0 || app.raflow.validationCheck.nonFieldsErrors.dates.length > 0,
                    people: app.raflow.validationCheck.errors.people.total > 0 || app.raflow.validationCheck.nonFieldsErrors.people.length > 0,
                    pets: app.raflow.validationCheck.errors.pets.total > 0 || app.raflow.validationCheck.nonFieldsErrors.pets.length > 0,
                    vehicles: app.raflow.validationCheck.errors.vehicles.total > 0 || app.raflow.validationCheck.nonFieldsErrors.vehicles.length > 0,
                    rentables: app.raflow.validationCheck.errors.rentables.total > 0 || app.raflow.validationCheck.nonFieldsErrors.rentables.length > 0,
                    parentchild: app.raflow.validationCheck.errors.parentchild.total > 0 || app.raflow.validationCheck.nonFieldsErrors.parentchild.length > 0,
                    tie: app.raflow.validationCheck.errors.tie.people.total > 0 || app.raflow.validationCheck.nonFieldsErrors.tie.length > 0
                };

                displayErrorDot();

                displayActiveComponentError();

                if(app.raflow.validationCheck.total > 0){
                    w2ui.raActionLayout.get('top').toolbar.click('btnBackToRA');
                    return false;
                }

                break;
        }

        if("raActionLayout" in w2ui){
            w2ui.raActionLayout.get('main').content = "";
        }

        loadRAActionTemplate();
        setTimeout(function() {
            reloadActionForm();
        },200);

    })
    .fail(function(data) {
        if (typeof data == "object") {
            var msg = data.responseText;
            var idx = msg.indexOf('"message"');
            var l = msg.length;
            // looking for the data in the form "message":"bla bla bla"
            if (idx >= 0 && l > 11 + idx) {
                var m = msg.substr(idx+11); // starts at bla in example comment above
                idx = m.indexOf('"');
                if (idx > 0) {
                    msg = m.substring(0,idx);
                }
            }
            w2ui.RAActionForm.error(msg);
        }
    });
};

// -------------------------------------------------------------------------------
// reloadActionForm - reloads the data of action form according to state
// -------------------------------------------------------------------------------
window.reloadActionForm = function() {
    // if version is refno, then remove NoticeToMove and Terminate options from Dropdown
    if (app.raflow.version === 'refno') {
        var itemArrayObject = Array.from(app.w2ui.listItems.RAActions);
        itemArrayObject.pop();
        itemArrayObject.pop();

        w2ui.RAActionForm.get('RAActions').options.items = itemArrayObject;

    } else if( app.raflow.version === 'raid' ) {
        w2ui.RAActionForm.get('RAActions').options.items = app.w2ui.listItems.RAActions;
    }

    // HIDE ALL OF THE COMPONENTS, LABELS, DIV'S
    $('#RAActionRAInfo').hide();
    $('#RAActionTerminatedRAInfo').hide();
    $('#RAActionNoticeToMoveInfo').hide();

    $('button[name=RAGenerateRAForm]').hide();
    $('button[name=RAGenerateMoveInInspectionForm]').hide();
    $('button[name=RAGenerateMoveOutForm]').hide();
    $('button[name=save]').hide();

    w2ui.RAActionForm.get('RAApprovalDecision1').hidden = true;
    w2ui.RAActionForm.get('RADeclineReason1').hidden = true;
    w2ui.RAActionForm.get('RAApprovalDecision2').hidden = true;
    w2ui.RAActionForm.get('RADeclineReason2').hidden = true;
    w2ui.RAActionForm.get('RATerminationReason').hidden = true;
    w2ui.RAActionForm.get('RADocumentDate').hidden = true;
    w2ui.RAActionForm.get('RANoticeToMoveDate').hidden = true;

    var data = app.raflow.Flow.Data;
    var raFlags = data.meta.RAFLAGS;
    var state = parseInt(raFlags & 0xf);

    // DISPLAY COMPONENTS, LABELS AND DIV'S ACCORDING TO CURRENT STATE
    switch (state) {
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
            // auto load date in component if it is present in meta
            if (data.meta.DocumentDate != "1900-01-01 00:00:00 UTC"){
                var documentDate = w2uiUTCDateControlString(new Date((data.meta.DocumentDate).replace(" ", "T").replace(" UTC", "Z")));
                w2ui.RAActionForm.record.RADocumentDate = documentDate;
            }

            w2ui.RAActionForm.get('RADocumentDate').hidden = false;
            $('button[name=RAGenerateRAForm]').show();
            $('button[name=RAGenerateMoveInInspectionForm]').show();
            $('button[name=save]').show();
            break;

        // "Active"
        case 4:
            $('#RAActionRAInfo').show();
            break;

        // "Notice To Move"
        case 5:
            $('#RAActionNoticeToMoveInfo').show();
            break;

        // "Terminated"
        case 6:
            $('#RAActionTerminatedRAInfo').show();
            $('button[name=RAGenerateRAForm]').show();
            break;

        default:
    }
    w2ui.RAActionForm.refresh();
};

window.refreshLabels = function () {
    var data = app.raflow.Flow;
    var meta = data.Data.meta;

    //------------------------------------------------------------------------
    // Header Part
    //------------------------------------------------------------------------
    var x = document.getElementById("bannerRAID");
    if (x !== null) {
        if (data.ID == 0) {
            x.innerHTML = 'New Rental Agreement';
        } else {
            x.innerHTML = '' + data.ID;
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

    //------------------------------------------------------------------------
    // Footer Part
    //------------------------------------------------------------------------
    x = document.getElementById("footerApplicationFilledBy");
    if (x !== null) {
        if (meta.ApplicationReadyUID == 0) {
            x.innerHTML = '';
        } else {
            x.innerHTML = dtFormatISOToW2ui((meta.ApplicationReadyDate).replace(" ", "T").replace(" UTC", "Z")) + ' by ' + meta.ApplicationReadyName;
        }
    }

    x = document.getElementById("footerApprover1");
    if (x !== null) {
        if (meta.Approver1 == 0) {
            x.innerHTML = '';
        } else {
            if ((meta.RAFLAGS & (1<<4)) > 0) {
                x.innerHTML = dtFormatISOToW2ui((meta.DecisionDate1).replace(" ", "T").replace(" UTC", "Z")) + '  Approved by ' + meta.Approver1Name ;
            } else{
                var reason1 = app.ApplDeny.find(function(t){if(t.id == meta.DeclineReason1){return t;}});
                var reason1Text = reason1 ? reason1.text : "";
                x.innerHTML = dtFormatISOToW2ui((meta.DecisionDate1).replace(" ", "T").replace(" UTC", "Z")) + '  Declined by ' + meta.Approver1Name + ' (' + reason1Text + ')';
            }
        }
    }

    x = document.getElementById("footerApprover2");
    if (x !== null) {
        if (meta.Approver2 == 0) {
            x.innerHTML = '';
        } else {
            if ((meta.RAFLAGS & (1<<5)) > 0) {
                x.innerHTML = dtFormatISOToW2ui((meta.DecisionDate2).replace(" ", "T").replace(" UTC", "Z")) + '  Approved by ' + meta.Approver2Name;
            } else{
                var reason2 = app.ApplDeny.find(function(t){if(t.id == meta.DeclineReason2){return t;}});
                var reason2Text = reason2 ? reason2.text : "";
                x.innerHTML = dtFormatISOToW2ui((meta.DecisionDate2).replace(" ", "T").replace(" UTC", "Z")) + '  Declined by ' + meta.Approver2Name + ' (' + reason2Text + ')';
            }
        }
    }

    x = document.getElementById("footerMoveInBy");
    if (x !== null) {
        if (meta.MoveInUID == 0) {
            x.innerHTML = '';
        } else {
            x.innerHTML = dtFormatISOToW2ui((meta.MoveInDate).replace(" ", "T").replace(" UTC", "Z")) + ' by ' + meta.MoveInName;
        }
    }

    x = document.getElementById("footerActiveBy");
    if (x !== null) {
        if (meta.ActiveUID == 0) {
            x.innerHTML = '';
        } else {
            x.innerHTML = dtFormatISOToW2ui((meta.ActiveDate).replace(" ", "T").replace(" UTC", "Z")) + ' by ' + meta.ActiveName;
        }
    }

    x = document.getElementById("footerRecievedNoticeToMoveBy");
    if (x !== null) {
        if (meta.NoticeToMoveUID == 0) {
            x.innerHTML = '';
        } else {
            var moveDate = '';
            if (meta.NoticeToMoveDate != "1900-01-01 00:00:00 UTC") {
                moveDate = w2uiUTCDateControlString(new Date((meta.NoticeToMoveDate).replace(" ", "T").replace(" UTC", "Z")));
            }
            x.innerHTML = dtFormatISOToW2ui((meta.NoticeToMoveReported).replace(" ", "T").replace(" UTC", "Z")) + ' by ' + meta.NoticeToMoveName + ' (move date: ' + moveDate + ')';
        }
    }

    x = document.getElementById("footerTerminatedBy");
    if (x !== null) {
        if (meta.LeaseTerminationReason == 0) {
            x.innerHTML = '';
        } else {
            var tReason;
            var tReasonText;
            if (meta.DeclineReason1 > 0 || meta.DeclineReason2 > 0) {
                tReason = app.RollerMsgs.find(function(t){if(t.id == meta.LeaseTerminationReason){return t;}});
            } else {
                tReason = app.WhyLeaving.find( function(t){ if(t.id == meta.LeaseTerminationReason) {return t;} } );
                if (typeof tReason === "undefined") {
                    tReason = app.RollerMsgs.find( function(t){ if(t.id == meta.LeaseTerminationReason) {return t;} } );
                }
            }
            tReasonText = tReason ? tReason.text : "";
            x.innerHTML = dtFormatISOToW2ui((meta.TerminationDate).replace(" ", "T").replace(" UTC", "Z")) + ' by '+ meta.TerminatorName + ' (' + tReasonText + ')';
        }
    }


    //------------------------------------------------------------------------
    // State Terminated Display Info
    //------------------------------------------------------------------------
    x = document.getElementById("bannerTerminatedBy");
    if (x !== null) {
        if (meta.TerminatorUID > 0) {
            x.innerHTML = dtFormatISOToW2ui((meta.TerminationDate).replace(" ", "T").replace(" UTC", "Z")) + ' by ' + meta.TerminatorName;
        } else {
            x.innerHTML = '';
        }
    }

    x = document.getElementById("bannerTerminationReason");
    if (x !== null) {
        if (meta.LeaseTerminationReason > 0) {
            var termination;
            var terminationReason;
            if (meta.DeclineReason1 > 0) {
                termination = app.RollerMsgs.find(function(t){if(t.id == meta.LeaseTerminationReason){return t;}});
                terminationReason = termination ? termination.text : "";

                // APPEND DECLINE REASON 1 IN BRACKETS
                var dreason1 = app.ApplDeny.find(function(t){if(t.id == meta.DeclineReason1){return t;}});
                var dreason1Text = dreason1 ? dreason1.text : "";

                // IN NEW LINE
                terminationReason += " ( "+ dreason1Text +" )";

            } else if (meta.DeclineReason2 > 0) {
                termination = app.RollerMsgs.find(function(t){if(t.id == meta.LeaseTerminationReason){return t;}});
                terminationReason = termination ? termination.text : "";

                // APPEND DECLINE REASON 1 IN BRACKETS
                var dreason2 = app.ApplDeny.find(function(t){if(t.id == meta.DeclineReason2){return t;}});
                var dreason2Text = dreason2 ? dreason2.text : "";

                // IN NEW LINE
                terminationReason += " ( "+ dreason2Text +" )";

            } else {
                termination = app.WhyLeaving.find(function(t){if(t.id == meta.LeaseTerminationReason){return t;}});
                // If RA is updated then the reason id will not be in WhyLeaving
                // hence we get it from RollerMsg
                if (!termination) {
                    termination = app.RollerMsgs.find(function(t){if(t.id == meta.LeaseTerminationReason){return t;}});
                }
                terminationReason = termination ? termination.text : "";
            }
            x.innerHTML = terminationReason;
        } else {
            x.innerHTML = '';
        }
    }

    // State Notice To Move Display Info
    x = document.getElementById("bannerMoveDate");
    if (x !== null) {
        if (meta.NoticeToMoveDate != "1900-01-01 00:00:00 UTC") {
            x.innerHTML = w2uiUTCDateControlString(new Date((meta.NoticeToMoveDate).replace(" ", "T").replace(" UTC", "Z")));
        } else {
            x.innerHTML = '';
        }
    }

    x = document.getElementById("bannerRecievedNoticeDate");
    if (x !== null) {
        if (meta.NoticeToMoveReported != "1900-01-01 00:00:00 UTC") {
            x.innerHTML = dtFormatISOToW2ui((meta.NoticeToMoveReported).replace(" ", "T").replace(" UTC", "Z"));
        } else {
            x.innerHTML = '';
        }
    }

    // State Active Display Info
    x = document.getElementById("bannerDocumentDate");
    if (x !== null) {
        if (meta.DocumentDate != "1900-01-01 00:00:00 UTC") {
            x.innerHTML = w2uiUTCDateControlString(new Date((meta.DocumentDate).replace(" ", "T").replace(" UTC", "Z")));
        } else {
            x.innerHTML = '';
        }
    }

    x = document.getElementById("bannerPayors");
    if (x !== null) {
        if (data.Data.people.length >0) {
            var payorList = [];
            data.Data.people.forEach(function(item) {
                if(item.IsRenter) {
                    payorList.push(item.FirstName + ' ' +item.MiddleName+ ' ' +item.LastName);
                }
            });
            x.innerHTML = payorList;
        } else {
            x.innerHTML = '';
        }
    }

    x = document.getElementById("bannerUsers");
    if (x !== null) {
        if (data.Data.people.length >0) {
            var userList = [];
            data.Data.people.forEach(function(item) {
                if(item.IsOccupant) {
                    userList.push(item.FirstName + ' ' +item.MiddleName+ ' ' +item.LastName);
                }
            });
            x.innerHTML = userList;
        } else {
            x.innerHTML = '';
        }
    }

    x = document.getElementById("bannerGuarantors");
    if (x !== null) {
        if (data.Data.people.length >0) {
            var guarantorList = [];
            data.Data.people.forEach(function(item) {
                if(item.IsGuarantor) {
                    guarantorList.push(item.FirstName + ' ' +item.MiddleName+ ' ' +item.LastName);
                }
            });
            x.innerHTML = guarantorList;
        } else {
            x.innerHTML = '';
        }
    }

    x = document.getElementById("bannerRentables");
    if (x !== null) {
        if (data.Data.rentables.length >0) {
            var rentableList = [];
            data.Data.rentables.forEach(function(item) {
                rentableList.push(item.RentableName);
            });
            x.innerHTML = rentableList;
        } else {
            x.innerHTML = '';
        }
    }

    x = document.getElementById("bannerPets");
    if (x !== null) {
            var petList = [];
        if (data.Data.pets.length >0) {
            data.Data.pets.forEach(function(item) {
                petList.push(item.Name);
            });
            x.innerHTML = petList;
        } else {
            x.innerHTML = '';
        }
    }

    x = document.getElementById("bannerVehicles");
    if (x !== null) {
            var vehicleList = [];
        if (data.Data.vehicles.length >0) {
            data.Data.vehicles.forEach(function(item) {
                vehicleList.push(GetVehicleIdentity(item));
            });
            x.innerHTML = vehicleList;
        } else {
            x.innerHTML = '';
        }
    }
};

//------------------------------------------------------------------------
// loadRAActionTemplate - It creates a layout for action forms and places
//                        it in newralayout's right panel.
//                        Top panel & bottom panel of this layout contains
//                        header & footer of action form respectively.
// -----------------------------------------------------------------------
window.loadRAActionTemplate = function() {
    if(!w2ui.raActionLayout) {
        $().w2layout({
            name: 'raActionLayout',
            padding: 0,
            panels: [
                { type: 'left', style: app.pstyle2, hidden: true },
                { type: 'top', style: app.pstyle2, content:'top', size: actionsUI.hdrHeight,
                    toolbar: {
                        items: [
                            { id: 'btnBackToRA',    type: 'button',     icon: 'fas fa-angle-left', text: '' },
                            { id: 'bt3',            type: 'spacer' },
                            { id: 'btnClose',       type: 'button',     icon: 'fas fa-times' }
                        ],
                        onClick: function (event) {
                            switch(event.target) {
                            case 'btnBackToRA':
                                var no_callBack = function() { return false; },
                                    yes_callBack = function() {
                                        w2ui.newraLayout.content('right','');
                                        w2ui.newraLayout.hide('right',true);
                                        w2ui.raActionLayout.get('main').content.destroy();
                                        w2ui.newraLayout.unlock('main');
                                        w2ui.newraLayout.get('main').toolbar.refresh();

                                        // get the current component of raflow interface (to be previous one)
                                        var active_comp = $(".ra-form-component:visible");

                                        // load target section (for refresh purpose)
                                        loadTargetSection(active_comp.attr("id"), active_comp.attr("id"));
                                    };
                                form_dirty_alert(yes_callBack, no_callBack);
                                break;
                            case 'btnClose':
                                yes_callBack = function() {
                                    CloseRAFlowLayout();
                                };
                                form_dirty_alert(yes_callBack, no_callBack);
                                break;
                            }
                        },
                    }
                },
                { type: 'main', style: app.pstyle2, content: 'main'},
                { type: 'preview', style: app.pstyle2, hidden: true },
                { type: 'bottom', style: app.pstyle2, size: actionsUI.ftrHeight,content:'bottom' },
                { type: 'right', style: app.pstyle2, hidden: true}
            ],
            onRefresh: function(event) {
                event.onComplete = function() {
                    refreshLabels();
                };
            },
            onRender: function(event) {
                event.onComplete = function() {
                    var layout = w2ui.raActionLayout;
                    var btnBackToRAText = "";
                    if (app.raflow.version === "raid") {
                        var RAID = app.raflow.Flow.ID;
                        btnBackToRAText = "<p style='font-size: 10pt; margin: 0 5px;'>Back to <strong>RA" + RAID + "</strong></p>";
                    } else if(app.raflow.version === "refno") {
                        var UserRefNo = app.raflow.Flow.UserRefNo;
                        btnBackToRAText = "<p style='font-size: 10pt; margin: 0 5px;'>Back to <strong>" + UserRefNo + "</strong></p>";
                    }
                    layout.get("top").toolbar.set('btnBackToRA', {text: btnBackToRAText});
                    // REFRESH THE TOOLBAR TO GET THE EFFECT
                    layout.get("top").toolbar.refresh();
                };
            }
        });
        w2ui.raActionLayout.load('top', '/webclient/html/raflow/formra-actionheader.html');
        w2ui.raActionLayout.load('bottom', '/webclient/html/raflow/formra-actionfooter.html');
    }
    w2ui.newraLayout.content('right', w2ui.raActionLayout);

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
            style: 'display: block;',
            focus: -1,
            formURL: '/webclient/html/raflow/formra-actionmain.html',
            fields: [
                { field: 'RAApprovalDecision1', type: 'list', width: 120, required: true, hidden: true,
                    options: {
                        items: [
                            {id: 0, text: "--Select Approve or Decline--"},
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
                            {id: 0, text: "--Select Approve or Decline--"},
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
                { field: 'RANoticeToMoveDate', type: 'date', hidden: true, options: { start: w2uiDateControlString(new Date()) } },
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
                });

                switch(event.target) {
                    case 'RAActions':
                        switch (event.value_new.id) {
                            case 5: // Received Notice-To-Move
                                w2ui.RAActionForm.get('RANoticeToMoveDate').hidden = false;

                                // auto load date in component if it is present in meta
                                if (app.raflow.Flow.Data.meta.NoticeToMoveDate != "1900-01-01 00:00:00 UTC"){
                                    var moveDate = w2uiUTCDateControlString(new Date((app.raflow.Flow.Data.meta.NoticeToMoveDate).replace(" ", "T").replace(" UTC", "Z")));
                                    this.record.RANoticeToMoveDate = moveDate;
                                }

                                w2ui.RAActionForm.get('RATerminationReason').hidden = true;
                                delete this.record.RATerminationReason;
                                $('button[name=updateAction]').attr('disabled',false);
                                break;

                            case 6: // Terminate
                                w2ui.RAActionForm.get('RATerminationReason').hidden = false;
                                $('button[name=updateAction]').attr('disabled',true);

                                delete this.record.RANoticeToMoveDate;
                                w2ui.RAActionForm.get('RANoticeToMoveDate').hidden = true;
                                break;

                            default:
                                w2ui.RAActionForm.get('RATerminationReason').hidden = true;
                                delete this.record.RATerminationReason;
                                $('button[name=updateAction]').attr('disabled',false);

                                delete this.record.RANoticeToMoveDate;
                                w2ui.RAActionForm.get('RANoticeToMoveDate').hidden = true;
                        }
                        break;

                    case 'RAApprovalDecision1':
                        if(event.value_new.text === 'Decline') {
                            $('button[name=save]').attr('disabled',true);
                            w2ui.RAActionForm.get('RADeclineReason1').hidden = false;
                        } else if(event.value_new.text === 'Approve') {
                            $('button[name=save]').attr('disabled',false);
                            w2ui.RAActionForm.get('RADeclineReason1').hidden = true;
                            delete this.record.RADeclineReason1;
                        } else {
                            $('button[name=save]').attr('disabled',true);
                            w2ui.RAActionForm.get('RADeclineReason1').hidden = true;
                            delete this.record.RADeclineReason1;
                        }
                        break;

                    case 'RAApprovalDecision2':
                        if(event.value_new.text === 'Decline') {
                            $('button[name=save]').attr('disabled',true);
                            w2ui.RAActionForm.get('RADeclineReason2').hidden = false;
                        } else if(event.value_new.text === 'Approve') {
                            $('button[name=save]').attr('disabled',false);
                            w2ui.RAActionForm.get('RADeclineReason2').hidden = true;
                            delete this.record.RADeclineReason2;
                        } else {
                            $('button[name=save]').attr('disabled',true);
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
                var data = app.raflow.Flow.Data;
                var raFlags = data.meta.RAFLAGS;
                var raStateString = app.RAStates[parseInt(raFlags & 0xf)];

                // var RAID = app.raflow.Flow.ID;
                // if(RAID > 0 &&  (raStateString === "Pending First Approval" || raStateString === "Pending Second Approval")) {
                //     raStateString = 'Modification ' + raStateString;
                // }

                $('#RAActionStateLable').text(raStateString);

                refreshLabels();
            },
            onRender: function (event) {
                w2ui.RAActionForm.record = {
                    RAActions: {id: -1, text: "--Select an Action--"},
                };

                // load sl stringlist in app
                getSLStringList(getCurrentBID(), "RollerMsgs");
            },
            actions: {
                save: function() {
                    var UserRefNo = app.raflow.Flow.UserRefNo;
                    var RAID = app.raflow.Flow.ID;
                    var Mode = "State";
                    var Version = app.raflow.version;

                    var data = app.raflow.Flow.Data;
                    var raFlags = data.meta.RAFLAGS;
                    var raState =parseInt(raFlags & 0xf);

                    var Decision1 = 0;
                    var DeclineReason1 = 0;
                    var Decision2 = 0;
                    var DeclineReason2 = 0;
                    var DocumentDate = "1/1/1900";

                    var reqData = {};
                    switch(raState) {
                        case 1:
                            Decision1 =  w2ui.RAActionForm.record.RAApprovalDecision1.id;
                            DeclineReason1 = w2ui.RAActionForm.record.RADeclineReason1.id;

                            reqData = {
                                "UserRefNo":UserRefNo,
                                "RAID":RAID,
                                "Version":Version,
                                "Mode": Mode,
                                "Decision1": Decision1,
                                "DeclineReason1": DeclineReason1
                            };
                            submitActionForm(reqData);
                            break;
                        case 2:
                            Decision2 =  w2ui.RAActionForm.record.RAApprovalDecision2.id;
                            DeclineReason2 = w2ui.RAActionForm.record.RADeclineReason2.id;

                            reqData = {
                                "UserRefNo":UserRefNo,
                                "RAID":RAID,
                                "Version":Version,
                                "Mode": Mode,
                                "Decision2": Decision2,
                                "DeclineReason2": DeclineReason2
                            };
                            submitActionForm(reqData);
                            break;
                        case 3:
                            if(w2ui.RAActionForm.record.RADocumentDate) {
                                DocumentDate = w2ui.RAActionForm.record.RADocumentDate;
                            }
                            reqData = {
                                "UserRefNo":UserRefNo,
                                "RAID":RAID,
                                "Version":Version,
                                "Mode": Mode,
                                "DocumentDate": DocumentDate
                            };
                            submitActionForm(reqData);
                            break;
                    }
                },
                updateAction: function() {
                    var UserRefNo = app.raflow.Flow.UserRefNo;
                    var RAID = app.raflow.Flow.ID;
                    var Action = this.record.RAActions.id;
                    var TerminationReason = 0;
                    var NoticeToMoveDate = "1/1/1900";
                    var Mode = "Action";
                    var Version = app.raflow.version;

                    var currentState = parseInt(app.raflow.Flow.Data.meta.RAFLAGS & (0xf));
                    //----------------------------------------------------------------
                    // if Action is to change to current state, only do this if we're
                    // in the Notice-To-Move state... in order to change the date.
                    //----------------------------------------------------------------
                    if (Action === currentState && Action != 5) {
                        return;
                    }
                    var reqData = {};
                    switch(Action) {
                        case -1:
                            break;
                        case 0:
                        case 1:
                        case 2:
                        case 3:
                        case 4:
                            reqData = {
                                "UserRefNo": UserRefNo,
                                "RAID": RAID,
                                "Version": Version,
                                "Action": Action,
                                "Mode": Mode
                            };
                            submitActionForm(reqData);
                            break;
                        case 5:
                            if(w2ui.RAActionForm.record.RANoticeToMoveDate) {
                                NoticeToMoveDate = w2ui.RAActionForm.record.RANoticeToMoveDate;
                            }

                            reqData = {
                                "UserRefNo": UserRefNo,
                                "RAID": RAID,
                                "Version": Version,
                                "Action": Action,
                                "Mode": Mode,
                                "NoticeToMoveDate": NoticeToMoveDate
                            };
                            submitActionForm(reqData);
                            break;
                        case 6:
                            if(w2ui.RAActionForm.record.RATerminationReason.id >0) {
                                TerminationReason = w2ui.RAActionForm.record.RATerminationReason.id;
                            }
                            reqData = {
                                "UserRefNo": UserRefNo,
                                "RAID": RAID,
                                "Version": Version,
                                "Action": Action,
                                "Mode": Mode,
                                "TerminationReason": TerminationReason
                            };
                            submitActionForm(reqData);
                            break;
                    }
                }
            }
        });
    }
    // now render the form in specifiec targeted panel
    w2ui.raActionLayout.content('main', w2ui.RAActionForm);
    setTimeout(function() {
        reloadActionForm();
    }, 100);
};
