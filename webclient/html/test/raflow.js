"use strict";

// Next button handling
$("#next").click(function() {
    // get the current component (to be previous one)
    var active_comp = $(".ra-form-component:visible");

    // get the target component (to be active one)
    var target_comp = active_comp.next(".ra-form-component");

    // make sure that next component available so we can navigate onto it
    if (target_comp.length === 0) {
        return false;
    }

    // load target section
    loadTargetSection(target_comp.attr("id"), active_comp.attr("id"));
});

// Previous button handling
$("#previous").click(function() {
    // get the current component (to be previous one)
    var active_comp = $(".ra-form-component:visible");

    // get the target component (to be active one)
    var target_comp = active_comp.prev(".ra-form-component");

    // make sure that previous component available so we can navigate onto it
    if (target_comp.length === 0) {
        return false;
    }

    // load target section
    loadTargetSection(target_comp.attr("id"), active_comp.attr("id"));
});

// link click handling
$("#progressbar a").click(function() {
    var active_comp = $(".ra-form-component:visible");

    // load target form
    var target = $(this).closest("li").attr("data-target");
    target = target.split('#').join("");

    loadTargetSection(target, active_comp.attr("id"));

    // because of 'a' tag, return false
    return false;
});

//-----------------------------------------------------------------------------
// w2uiDateControlString
//           - return a date string formatted the way the w2ui dates are
//             expected, based on the supplied date that can be
//             used as the .value attribute of a date control.  That is, in
//             the format  m/d/yyyy.
// @params
//   dt = java date value
// @return string value mm-dd-yyyy
//-----------------------------------------------------------------------------
function w2uiDateControlString(dt) {
    var m = dt.getMonth() + 1;
    var d = dt.getDate();
    var s = '' + m + '/' + d+'/' + dt.getFullYear();
    return s;
}

// TODO: we should pass flowID, flowPartID here in arguments
function saveActiveCompData(record, partType) {

    var flowPartID;
    var flowParts = app.raflow.data[app.raflow.activeflowID] || [];

    for (var i = 0; i < flowParts.length; i++) {
        if(partType == flowParts[i].PartType) {
            flowPartID = flowParts[i].FlowPartID;
        }
    }

    // temporary data
    var data = {
        "cmd": "save",
        "FlowPartID": flowPartID,
        "Flow": app.raflow.name,
        "FlowID": app.raflow.activeflowID,
        "BID": 1,
        "PartType": partType,
        "Data": record,
    };

    $.ajax({
        url: "/v1/flowpart/1/0",
        method: "POST",
        contentType: "application/json",
        dataType: "json",
        data: JSON.stringify(data),
        success: function(data) {
            if (data.status != "error") {
                console.log("data has been saved for: ", app.raflow.activeflowID, ", partType: ", partType);

                $("#manage-flows #message").hide();

            } else {
                $("#manage-flows #message").text(data.message).show();
            }
        },
        error: function(data) {
            console.log(data);
        },
    });
}

function getRAFlowAllParts(flowID) {
    $.ajax({
        url: "/v1/flow/1/0",
        method: "POST",
        contentType: "application/json",
        dataType: "json",
        data: JSON.stringify({"cmd": "getFlowParts", "flowID": flowID}),
        success: function(data) {
            if (data.status != "error") {
                app.raflow.data[flowID] = data.records;
                // load form container
                $("#ra-form-container").animate({"left": "0"}, 100);
                // load first dates section
                loadRADatesForm();
                // as we load the first section
                $("#ra-form footer button#previous").addClass("disable");

                $("#manage-flows #message").hide();
            } else {
                $("#manage-flows #message").text(data.message).show();
            }
        },
        error: function(data) {
            console.log(data);
        },
    });
}

function getAllRAFlows() {
    $.ajax({
        url: "/v1/flow/1/0",
        method: "POST",
        contentType: "application/json",
        dataType: "json",
        data: JSON.stringify({"cmd": "getAllFlows", "flow": app.raflow.name}),
        success: function(data) {

            if (data.status != "error") {

                $("#flow-list").empty();
                $("#manage-flows #loader").show();

                data.records = data.records || [];
                for (var i = 0; i < data.records.length; i++) {
                    $("#flow-list").append("<li class='flowID-link' data-flow-id='"+data.records[i]+"'>"+data.records[i]+"</li>");
                }

                $("#manage-flows #message").hide();
                $("#manage-flows #loader").hide();
            } else {
                $("#manage-flows #message").text(data.message).show();
                $("#manage-flows #loader").hide();
            }
        },
        error: function(data) {
            console.log(data);
        },
    });
}

$("#back-to-flow-list").on("click", function() {
    app.raflow.activeflowID = "";
    $("#ra-form-container").animate({"left": "100%"}, 100);
});

$("#add-new-flow").on("click", function() {
    initRAFlow();
});

$(document).on("click", ".flowID-link", function() {

    // load first slide
    $("#ra-form footer button#previous").addClass("disable");
    $(".ra-form-component").hide();
    $(".ra-form-component#dates").show();
    $("#progressbar li").removeClass("active");
    $("#progressbar li[data-target='#dates']").addClass("active");

    app.raflow.activeflowID = $(this).attr("data-flow-id");
    getRAFlowAllParts(app.raflow.activeflowID);
});

function initRAFlow() {
    $.ajax({
        url: "/v1/flow/1/0",
        method: "POST",
        contentType: "application/json",
        dataType: "json",
        data: JSON.stringify({"cmd": "init", "flow": app.raflow.name}),
        success: function(data) {
            if (data.status != "error") {
                app.raflow.data[data.FlowID] = {};
                $("#flow-list").append("<li class='flowID-link' data-flow-id='"+data.FlowID+"'>"+data.FlowID+"</li>");

                $("#manage-flows #message").hide();
            } else {
                $("#manage-flows #message").text(data.message).show();
            }
        },
        error: function(data) {
            console.log(data);
        },
    });
}

// load first section in main part
$(function() {

    $.get('/v1/uilists/' + app.language + '/' + app.template)
    .done(function(data, textStatus, jqXHR) {
        if (jqXHR.status == 200) {
            for( var key in data ) {   // fit all lists, values, maps in app variable
                app[key] = data[key];
            }

            // get all flows list
            getAllRAFlows(app.raflow.name);

        } else {
            console.log( '**** YIPES! ****  status on /v1/uilists/ = ' + textStatus);
        }
    })
    .fail( function() {
        console.log('Error getting /v1/uilists');
        console.log('*** NO INTERFACE ***');
    });
});

// RACompConfig for each section
var RACompConfig = {
    "dates": {
        loader: loadRADatesForm,
        w2uiComp: "RADatesForm",
    },
    "people": {
        loader: loadRAPeopleForm,
        w2uiComp: "RAPeopleForm",
    },
    "pets": {
        loader: loadRAPetsGrid,
        w2uiComp: "RAPetsGrid",
    },
    "vehicles": {
        loader: loadRAVehiclesGrid,
        w2uiComp: "RAVehiclesGrid",
    },
    "bginfo": {
        loader: loadRABGInfoForm,
        w2uiComp: "RABGInfoForm",
    },
    "rentables": {
        loader: loadRARentablesGrid,
        w2uiComp: "RARentablesGrid",
    },
    "feesterms": {
        loader: loadRAFeesTermsGrid,
        w2uiComp: "RAFeesTermsGrid",
    },
    "final": {
        loader: null,
        w2uiComp: "",
    },
};

// load form according to target
function loadTargetSection(target, activeCompID) {

    // get part type from the class index
    var partType = $("#progressbar li[data-target='#" + activeCompID + "']").index() + 1;
    var data = null;
    if($("#progressbar li[data-target='#" + target + "']").hasClass("done")){
        console.log("target has been saved", target);
    } else{
        // TODO: switch cases for each part type, so that we can mark the section "done"
        // if it's completed

        // // add class "done" to mark the section tab as done
        // $("#progressbar li[data-target='#" + activeCompID + "']").addClass("done");
    }

    // decide data based on type
    switch(activeCompID) {
        case "dates":
            data = w2ui.RADatesForm.record;
            break;
        case "people":
            data = w2ui.RAPeopleForm.record;
            break;
        case "pets":
            data = w2ui.RAPetsGrid.records;
            break;
        case "vehicles":
            data = w2ui.RAVehiclesGrid.records;
            break;
        case "bginfo":
            data = w2ui.RABGInfoForm.record;
            break;
        case "rentables":
            data = w2ui.RARentablesGrid.records;
            break;
        case "feesterms":
            data = w2ui.RAFeesTermsGrid.records;
            break;
        case "final":
            data = null;
            break;
        default:
            alert("invalid active comp: ", activeCompID);
            return;
    }

    if (data) {
        // save the content on server for active component
        saveActiveCompData(data, partType);
    }

    // hide active component
    $("#progressbar li[data-target='#" + activeCompID + "']").removeClass("active");
    $(".ra-form-component#" + activeCompID).hide();

    // show target component
    $("#progressbar li[data-target='#" + target + "']").addClass("active");
    $(".ra-form-component#" + target).show();

    // hide previous navigation button if the target is in first section
    if ($(".ra-form-component#" + target).is($(".ra-form-component").first())) {
        $("#ra-form footer button#previous").addClass("disable");
    } else {
        $("#ra-form footer button#previous").removeClass("disable");
    }

    // hide next navigation button if the target is in last section
    if ($(".ra-form-component#" + target).is($(".ra-form-component").last())) {
        $("#ra-form footer button#next").addClass("disable");
    } else {
        $("#ra-form footer button#next").removeClass("disable");
    }

    // load the content in the component using loader function
    var targetLoader = RACompConfig[target].loader;
    if (typeof targetLoader === "function") {
        targetLoader();
        /*setTimeout(function() {
            var validateForm = compIDw2uiForms[activeCompID];
            if (typeof w2ui[validateForm] !== "undefined") {
                var issues = w2ui[validateForm].validate();
                if (!(Array.isArray(issues) && issues.length > 0)) {
                    // $("#progressbar li[data-target='#" + activeCompID + "']").addClass("done");
                }
            }
        }, 500);*/
    } else{
        console.log("unknown target from nav li: ", target);
    }
}

// -------------------------------------------------------------------------------
// Rental Agreement - Info Dates form
// -------------------------------------------------------------------------------
function loadRADatesForm() {

    // if form is loaded then return
    if (!("RADatesForm" in w2ui)) {

        // dates form
        $('#ra-form #dates').w2form({
            name   : 'RADatesForm',
            header : 'Dates',
            style  : 'border: 1px black solid; display: block;',
            focus  : -1,
            formURL: '/webclient/html/test/formradates.html',
            fields : [
                { name: 'AgreementStart',  type: 'date', required: true, html: { caption: "Term Start" } },
                { name: 'AgreementStop',   type: 'date', required: true, html: { caption: "Term Stop"  } },
                { name: 'RentStart',       type: 'date', required: true, html: { caption: "Rent Start" } },
                { name: 'RentStop',        type: 'date', required: true, html: { caption: "Rent Stop"  } },
                { name: 'PossessionStart', type: 'date', required: true, html: { caption: "Possession Start" } },
                { name: 'PossessionStop',  type: 'date', required: true, html: { caption: "Possession Stop"  } },
            ],
            actions: {
                reset: function () {
                    this.clear();
                },
                /*save: function () {
                    this.save();
                }*/
            }
        });

    }

    // load the existing data in dates component
    setTimeout(function() {
        var partType = app.raFlowPartTypes.dates;
        if (app.raflow.activeflowID && app.raflow.data[app.raflow.activeflowID]) {
            for (var i = 0; i < app.raflow.data[app.raflow.activeflowID].length; i++) {
                if (partType == app.raflow.data[app.raflow.activeflowID][i].PartType) {
                    if (app.raflow.data[app.raflow.activeflowID][i].Data) {
                        w2ui.RADatesForm.record = app.raflow.data[app.raflow.activeflowID][i].Data;
                        w2ui.RADatesForm.refresh();
                    } else {
                        w2ui.RADatesForm.clear();
                    }
                    break;
                }
            }
        }
    }, 500);
}

// -------------------------------------------------------------------------------
// Rental Agreement - People form
// -------------------------------------------------------------------------------
function loadRAPeopleForm() {

    // if form is loaded then return
    if (!("RAPeopleForm" in w2ui)) {

        // people form
        $('#ra-form #people').w2form({
            name   : 'RAPeopleForm',
            header : 'People',
            style  : 'border: 1px solid black; display: block;',
            formURL: '/webclient/html/test/formrapeople.html',
            focus: -1,
            fields : [
                { name: 'Transactant', type: 'combo',    required: true, html: { caption: "Transactant" },
                    options: {
                        items: ["Captain America", "Iron Man", "Doctor Strange", "Thanos"]
                    }
                },
                { name: 'Payor',       type: 'checkbox', required: true, html: { caption: "Payor" } },
                { name: 'User',        type: 'checkbox', required: true, html: { caption: "User" } },
                { name: 'Guarantor',   type: 'checkbox', required: true, html: { caption: "Guarantor" } },
            ],
            actions: {
                reset: function () {
                    this.clear();
                },
                /*save: function () {
                    this.save();
                }*/
            }
        });
    }

    // load the existing data in people component
    setTimeout(function() {
        var partType = app.raFlowPartTypes.people;
        if (app.raflow.activeflowID && app.raflow.data[app.raflow.activeflowID]) {
            for (var i = 0; i < app.raflow.data[app.raflow.activeflowID].length; i++) {
                if (partType == app.raflow.data[app.raflow.activeflowID][i].PartType) {
                    if (app.raflow.data[app.raflow.activeflowID][i].Data) {
                        w2ui.RAPeopleForm.record = app.raflow.data[app.raflow.activeflowID][i].Data;
                        w2ui.RAPeopleForm.refresh();
                    } else {
                        w2ui.RAPeopleForm.clear();
                    }
                    break;
                }
            }
        }
    }, 500);
}

// -------------------------------------------------------------------------------
// Rental Agreement - Pets Grid
// -------------------------------------------------------------------------------
function getPetsGridInitalRecord(BID, gridLen) {
    var t   = new Date(),
        nyd = new Date(new Date().setFullYear(new Date().getFullYear() + 1));

    return {
        recid:                 gridLen,
        PETID:                 0,
        BID:                   BID,
        RAID:                  0,
        Name:                  "",
        Type:                  "",
        Breed:                 "",
        Color:                 "",
        Weight:                0,
        DtStart:               w2uiDateControlString(t),
        DtStop:                w2uiDateControlString(nyd),
        RefundablePetDeposit:  0.0,
        RecurringPetFee:       0.0,
        NonRefundablePetFee:   0.0
    };
}

function loadRAPetsGrid() {
    // if form is loaded then return
    if (!("RAPetsGrid" in w2ui)) {

        // pets grid
        $('#ra-form #pets').w2grid({
            name   : 'RAPetsGrid',
            header : 'Pets',
            show   : {
                        toolbar: true,
                        footer: true,
                        // toolbarSave: true
                     },
            style  : 'border: 1px solid black; display: block;',
            toolbar: {
                items: [
                    { id: 'add', type: 'button', caption: 'Add Record', icon: 'w2ui-icon-plus' }
                ],
                onClick: function(event) {
                    if (event.target == 'add') {
                        var inital = getPetsGridInitalRecord(1, w2ui.RAPetsGrid.records.length);
                        w2ui.RAPetsGrid.add(inital);
                    }
                }
            },
            columns: [
                {
                    field: 'recid',
                    hidden: true,
                },
                {
                    field:   'PETID',
                    hidden:  true
                },
                {
                    field:   'BID',
                    hidden:  true
                },
/*                {
                    field:   'RAID',
                    hidden:  true
                },*/
                {
                    field:   'Name',
                    caption: 'Name',
                    size:    '150px',
                    editable:{ type: 'text' }
                },
                {
                    field:   'Type',
                    caption: 'Type',
                    size:    '80px',
                    editable:{ type: 'text' }
                },
                {
                    field:   'Breed',
                    caption: 'Breed',
                    size:    '80px',
                    editable:{ type: 'text' }
                },
                {
                    field:   'Color',
                    caption: 'Color',
                    size:    '80px',
                    editable:{ type: 'text' }
                },
                {
                    field:   'Weight',
                    caption: 'Weight',
                    size:    '80px',
                    editable:{ type: 'int' }
                },
                {
                    field:   'DtStart',
                    caption: 'DtStart',
                    size:    '100px',
                    editable:{ type: 'date' }
                },
                {
                    field:   'DtStop',
                    caption: 'DtStop',
                    size:    '100px',
                    editable:{ type: 'date' }
                },
                {
                    field:   'NonRefundablePetFee',
                    caption: 'NonRefundable<br>PetFee',
                    size:    '70px',
                    render:  'money',
                    editable:{ type: 'money' }
                },
                {
                    field:   'RefundablePetDeposit',
                    caption: 'Refundable<br>PetDeposit',
                    size:    '70px',
                    render:  'money',
                    editable:{ type: 'money' }
                },
                {
                    field:   'RecurringPetFee',
                    caption: 'Recurring<br>PetFee',
                    size:    '100%',
                    render:  'money',
                    editable:{ type: 'money' }
                },
            ],
            onChange: function(event) {
                event.onComplete = function() {
                    this.save();
                };
            }
        });
    }

    // load the existing data in pets component
    setTimeout(function() {
        var partType = app.raFlowPartTypes.pets;
        if (app.raflow.activeflowID && app.raflow.data[app.raflow.activeflowID]) {
            for (var i = 0; i < app.raflow.data[app.raflow.activeflowID].length; i++) {
                if (partType == app.raflow.data[app.raflow.activeflowID][i].PartType) {
                    if (app.raflow.data[app.raflow.activeflowID][i].Data) {
                        w2ui.RAPetsGrid.records = app.raflow.data[app.raflow.activeflowID][i].Data;
                        w2ui.RAPetsGrid.refresh();
                    } else {
                        w2ui.RAPetsGrid.clear();
                    }
                    break;
                }
            }
        }
    }, 500);

}

// -------------------------------------------------------------------------------
// Rental Agreement - Vehicles Grid
// -------------------------------------------------------------------------------
function getVehicleGridInitalRecord(BID, gridLen) {
    var t   = new Date(),
        nyd = new Date(new Date().setFullYear(new Date().getFullYear() + 1));

    return {
        recid:                 gridLen,
        VID:                   0,
        BID:                   BID,
        TCID:                  0,
        VIN:                   "",
        Type:                  "",
        Make:                  "",
        Model:                 "",
        Color:                 "",
        LicensePlateState:     "",
        LicensePlateNumber:    "",
        ParkingPermitNumber:   "",
        DtStart:               w2uiDateControlString(t),
        DtStop:                w2uiDateControlString(nyd),
    };
}

function loadRAVehiclesGrid() {
    // if form is loaded then return
    if (!("RAVehiclesGrid" in w2ui)) {

        // vehicles grid
        $('#ra-form #vehicles').w2grid({
            name   : 'RAVehiclesGrid',
            header : 'Vehicles',
            show   : {
                        toolbar: true,
                        footer: true,
                        // toolbarSave: true
                     },
            style  : 'border: 1px solid black; display: block;',
            toolbar: {
                items: [
                    { id: 'add', type: 'button', caption: 'Add Record', icon: 'w2ui-icon-plus' }
                ],
                onClick: function(event) {
                    if (event.target == 'add') {
                        var inital = getVehicleGridInitalRecord(1, w2ui.RAVehiclesGrid.records.length);
                        w2ui.RAVehiclesGrid.add(inital);
                    }
                }
            },
            columns: [
                {
                    field: 'recid',
                    hidden: true,
                },
                {
                    field:   'VID',
                    hidden:  true
                },
                {
                    field:   'BID',
                    hidden:  true
                },
                {
                    field:   'TCID',
                    hidden:  true
                },
                {
                    field:   'Type',
                    caption: 'Type',
                    size:    '80px',
                    editable:{ type: 'text' }
                },
                {
                    field:   'VIN',
                    caption: 'VIN',
                    size:    '80px',
                    editable:{ type: 'text' }
                },
                {
                    field:   'Make',
                    caption: 'Make',
                    size:    '80px',
                    editable:{ type: 'text' }
                },
                {
                    field:   'Model',
                    caption: 'Model',
                    size:    '80px',
                    editable:{ type: 'text' }
                },
                {
                    field:   'Color',
                    caption: 'Color',
                    size:    '80px',
                    editable:{ type: 'text' }
                },
                {
                    field:   'LicensePlateState',
                    caption: 'License Plate<br>State',
                    size:    '100px',
                    editable:{ type: 'text' }
                },
                {
                    field:   'LicensePlateNumber',
                    caption: 'License Plate<br>Number',
                    size:    '100px',
                    editable:{ type: 'text' }
                },
                {
                    field:   'ParkingPermitNumber',
                    caption: 'Parking Permit <br>Number',
                    size:    '100px',
                    editable:{ type: 'text' }
                },
                {
                    field:   'DtStart',
                    caption: 'DtStart',
                    size:    '100px',
                    editable:{ type: 'date' }
                },
                {
                    field:   'DtStop',
                    caption: 'DtStop',
                    size:    '100%',
                    editable:{ type: 'date' }
                },
            ],
            onChange: function(event) {
                event.onComplete = function() {
                    this.save();
                };
            }
        });
    }

    // load the existing data in vehicles component
    setTimeout(function() {
        var partType = app.raFlowPartTypes.vehicles;
        if (app.raflow.activeflowID && app.raflow.data[app.raflow.activeflowID]) {
            for (var i = 0; i < app.raflow.data[app.raflow.activeflowID].length; i++) {
                if (partType == app.raflow.data[app.raflow.activeflowID][i].PartType) {
                    if (app.raflow.data[app.raflow.activeflowID][i].Data) {
                        w2ui.RAVehiclesGrid.records = app.raflow.data[app.raflow.activeflowID][i].Data;
                        w2ui.RAVehiclesGrid.refresh();
                    } else {
                        w2ui.RAVehiclesGrid.clear();
                    }
                    break;
                }
            }
        }
    }, 500);
}

// -------------------------------------------------------------------------------
// Rental Agreement - Background info form
// -------------------------------------------------------------------------------
function loadRABGInfoForm() {

    // if form is loaded then return
    if (!("RABGInfoForm" in w2ui)) {

        // background info form
        $('#ra-form #bginfo').w2form({
            name   : 'RABGInfoForm',
            header : 'Background Information',
            style  : 'border: 1px solid black; display: block;',
            formURL: '/webclient/html/test/formrabginfo.html',
            focus: -1,
            fields : [
                { name: 'Applicant'  , type: 'text'    , required: true, html: { caption: "Applicant Name" } },
            ],
            actions: {
                reset: function () {
                    this.clear();
                },
                /*save: function () {
                    this.save();
                }*/
            }
        });
    }

    // load the existing data in people component
    setTimeout(function() {
        var partType = app.raFlowPartTypes.bginfo;
        if (app.raflow.activeflowID && app.raflow.data[app.raflow.activeflowID]) {
            for (var i = 0; i < app.raflow.data[app.raflow.activeflowID].length; i++) {
                if (partType == app.raflow.data[app.raflow.activeflowID][i].PartType) {
                    if (app.raflow.data[app.raflow.activeflowID][i].Data) {
                        w2ui.RABGInfoForm.record = app.raflow.data[app.raflow.activeflowID][i].Data;
                        w2ui.RABGInfoForm.refresh();
                    } else {
                        w2ui.RABGInfoForm.clear();
                    }
                    break;
                }
            }
        }
    }, 500);
}

// -------------------------------------------------------------------------------
// Rental Agreement - Rentables Grid
// -------------------------------------------------------------------------------
function getRentablesGridInitalRecord(BID, gridLen) {
    return {
        recid:                 gridLen,
        RID:                   0,
        BID:                   BID,
        RTID:                  0,
        RentableName:          "",
        ContractRent:          0.0,
        ProrateAmt:            0.0,
        TaxableAmt:            0.0,
        SalesTax:              0.0,
        TransOCC:              0.0,
    };
}

function loadRARentablesGrid() {
    // if form is loaded then return
    if (!("RARentablesGrid" in w2ui)) {

        // rentables grid
        $('#ra-form #rentables').w2grid({
            name   : 'RARentablesGrid',
            header : 'Rentables',
            show   : {
                        toolbar: true,
                        footer: true,
                        // toolbarSave: true
                     },
            style  : 'border: 1px solid black; display: block;',
            toolbar: {
                items: [
                    { id: 'add', type: 'button', caption: 'Add Record', icon: 'w2ui-icon-plus' }
                ],
                onClick: function(event) {
                    if (event.target == 'add') {
                        var inital = getRentablesGridInitalRecord(1, w2ui.RARentablesGrid.records.length);
                        w2ui.RARentablesGrid.add(inital);
                    }
                }
            },
            columns: [
                {
                    field: 'recid',
                    hidden: true,
                },
                {
                    field:   'RID',
                    hidden:  true
                },
                {
                    field:   'BID',
                    hidden:  true
                },
                {
                    field:   'RTID',
                    hidden:  true
                },
                {
                    field:   'RentableName',
                    caption: 'Rentable',
                    size:    '350px',
                    editable:{ type: 'text' }
                },
                {
                    field:   'ContractRent',
                    caption: 'At Signing',
                    size:    '100px',
                    render:  'money',
                    editable:{ type: 'money' }
                },
                {
                    field:   'ProrateAmt',
                    caption: 'Prorate',
                    size:    '100px',
                    render:  'money',
                    editable:{ type: 'money' }
                },
                {
                    field:   'TaxableAmt',
                    caption: 'Taxable Amt',
                    size:    '100px',
                    render:  'money',
                    editable:{ type: 'money' }
                },
                {
                    field:   'SalesTax',
                    caption: 'Sales Tax',
                    size:    '100px',
                    render:  'money',
                    editable:{ type: 'money' }
                },
                {
                    field:   'TransOCC',
                    caption: 'Trans OCC',
                    size:    '100%',
                    render:  'money',
                    editable:{ type: 'money' }
                }
            ],
            onChange: function(event) {
                event.onComplete = function() {
                    this.save();
                };
            }
        });
    }

    // load the existing data in rentables component
    setTimeout(function() {
        var partType = app.raFlowPartTypes.rentables;
        if (app.raflow.activeflowID && app.raflow.data[app.raflow.activeflowID]) {
            for (var i = 0; i < app.raflow.data[app.raflow.activeflowID].length; i++) {
                if (partType == app.raflow.data[app.raflow.activeflowID][i].PartType) {
                    if (app.raflow.data[app.raflow.activeflowID][i].Data) {
                        w2ui.RARentablesGrid.records = app.raflow.data[app.raflow.activeflowID][i].Data;
                        w2ui.RARentablesGrid.refresh();
                    } else {
                        w2ui.RARentablesGrid.clear();
                    }
                    break;
                }
            }
        }
    }, 500);
}


// -------------------------------------------------------------------------------
// Rental Agreement - Fees Terms Grid
// -------------------------------------------------------------------------------
function getFeesTermsGridInitalRecord(BID, gridLen) {
    return {
        recid:                 gridLen,
        RID:                   0,
        BID:                   BID,
        RTID:                  0,
        RentableName:          "",
        FeeName:                   "",
        Amount:                0.0,
        Cycle:                 6,
        SigningAmt:            0.0,
        ProrateAmt:            0.0,
        TaxableAmt:            0.0,
        SalesTax:              0.0,
        TransOCC:              0.0,
    };
}

function loadRAFeesTermsGrid() {
    // if form is loaded then return
    if (!("RAFeesTermsGrid" in w2ui)) {

        // feesterms grid
        $('#ra-form #feesterms').w2grid({
            name   : 'RAFeesTermsGrid',
            header : 'FeesTerms',
            show   : {
                        toolbar: true,
                        footer: true,
                        // toolbarSave: true
                     },
            style  : 'border: 1px solid black; display: block;',
            toolbar: {
                items: [
                    { id: 'add', type: 'button', caption: 'Add Record', icon: 'w2ui-icon-plus' }
                ],
                onClick: function(event) {
                    if (event.target == 'add') {
                        var inital = getFeesTermsGridInitalRecord(1, w2ui.RAFeesTermsGrid.records.length);
                        w2ui.RAFeesTermsGrid.add(inital);
                    }
                }
            },
            columns: [
                {
                    field: 'recid',
                    hidden: true,
                },
                {
                    field:   'RID',
                    hidden:  true
                },
                {
                    field:   'BID',
                    hidden:  true
                },
                {
                    field:   'RTID',
                    hidden:  true
                },
                {
                    field:   'RentableName',
                    caption: 'Rentable',
                    size:    '180px',
                    editable:{ type: 'text' }
                },
                {
                    field:   'FeeName',
                    caption: 'Fee',
                    size:    '120px',
                    editable:{ type: 'text' }
                },
                {
                    field:   'Amount',
                    caption: 'Amount',
                    size:    '80px',
                    render:  'money',
                    editable:{ type: 'money' }
                },
                {
                    field:   'Cycle',
                    caption: 'Cycle',
                    size:    '80px',
                    editable:{ type: 'int' }
                },
                {
                    field:   'SigningAmt',
                    caption: 'At Signing',
                    size:    '80px',
                    render:  'money',
                    editable:{ type: 'money' }
                },
                {
                    field:   'ProrateAmt',
                    caption: 'Prorate',
                    size:    '80px',
                    render:  'money',
                    editable:{ type: 'money' }
                },
                {
                    field:   'TaxableAmt',
                    caption: 'Taxable Amt',
                    size:    '80px',
                    render:  'money',
                    editable:{ type: 'money' }
                },
                {
                    field:   'SalesTax',
                    caption: 'Sales Tax',
                    size:    '80px',
                    render:  'money',
                    editable:{ type: 'money' }
                },
                {
                    field:   'TransOCC',
                    caption: 'Trans OCC',
                    size:    '100%',
                    render:  'money',
                    editable:{ type: 'money' }
                }
            ],
            onChange: function(event) {
                event.onComplete = function() {
                    this.save();
                };
            }
        });
    }

    // load the existing data in feesterms component
    setTimeout(function() {
        var partType = app.raFlowPartTypes.feesterms;
        if (app.raflow.activeflowID && app.raflow.data[app.raflow.activeflowID]) {
            for (var i = 0; i < app.raflow.data[app.raflow.activeflowID].length; i++) {
                if (partType == app.raflow.data[app.raflow.activeflowID][i].PartType) {
                    if (app.raflow.data[app.raflow.activeflowID][i].Data) {
                        w2ui.RAFeesTermsGrid.records = app.raflow.data[app.raflow.activeflowID][i].Data;
                        w2ui.RAFeesTermsGrid.refresh();
                    } else {
                        w2ui.RAFeesTermsGrid.clear();
                    }
                    break;
                }
            }
        }
    }, 500);
}

