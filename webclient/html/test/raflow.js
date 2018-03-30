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
            alert("data has been saved");
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
            app.raflow.data[flowID] = data.records;
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
            data.records = data.records || [];
            for (var i = 0; i < data.records.length; i++) {
                getRAFlowAllParts(data.records[i]);
            }

            // TODO: make it dynamic later
            if (data.records.length > 0) {
                app.raflow.activeflowID = data.records[0];
            }
        },
        error: function(data) {
            console.log(data);
        },
    });
}

function initRAFlow() {
    $.ajax({
        url: "/v1/flow/1/0",
        method: "POST",
        contentType: "application/json",
        dataType: "json",
        data: JSON.stringify({"cmd": "init", "flow": app.raflow.name}),
        success: function(data) {
            app.raflow.data[data.flowID] = {};
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

            // initiate flow
            // initRAFlow(app.raflow.name);

            setTimeout(function() {
                loadRADatesForm();
            }, 1000);

            // as we load the first section
            $("#ra-form footer button#previous").addClass("disable");

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
        loader: null,
        w2uiComp: "",
    },
    "rentables": {
        loader: null,
        w2uiComp: "",
    },
    "feesterms": {
        loader: null,
        w2uiComp: "",
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
            break
        case "people":
            data = w2ui.RAPeopleForm.record;
            break
        case "pets":
            data = w2ui.RAPetsGrid.records;
            break
        case "vehicles":
            data = w2ui.RAPetsGrid.records;
            break
        case "bginfo":
            data = w2ui.RAPetsGrid.records;
            break
        case "rentables":
            data = w2ui.RAPetsGrid.records;
            break
        case "feesterms":
            data = w2ui.RAPetsGrid.records;
            break
        default:
            alert("invalid active comp: ", activeCompID);
            return
    }

    // save the content on server for active component
    saveActiveCompData(data, partType);

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
    if ("RADatesForm" in w2ui) {
        return;
    }

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

    // load the existing data in dates component
    setTimeout(function() {
        var partType = app.raFlowPartTypes.dates;
        for (var i = 0; i < app.raflow.data[app.raflow.activeflowID].length; i++) {
            if (partType == app.raflow.data[app.raflow.activeflowID][i].PartType) {
                w2ui.RADatesForm.record = app.raflow.data[app.raflow.activeflowID][i].Data;
                w2ui.RADatesForm.refresh();
            }
        }
    }, 500);
}

// -------------------------------------------------------------------------------
// Rental Agreement - People form
// -------------------------------------------------------------------------------
function loadRAPeopleForm() {

    // if form is loaded then return
    if ("RAPeopleForm" in w2ui) {
        return;
    }

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
                    items: ["Captain America","Iron Man","Doctor Strange","Thanos"]
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
            save: function () {
                this.save();
            }
        }
    });

    // load the existing data in people component
    setTimeout(function() {
        var partType = app.raFlowPartTypes.people;
        for (var i = 0; i < app.raflow.data[app.raflow.activeflowID].length; i++) {
            if (partType == app.raflow.data[app.raflow.activeflowID][i].PartType) {
                w2ui.RAPeopleForm.record = app.raflow.data[app.raflow.activeflowID][i].Data;
                w2ui.RAPeopleForm.refresh();
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
    }
}

function loadRAPetsGrid() {
    // if form is loaded then return
    if ("RAPetsGrid" in w2ui) {
        return;
    }

    // pets grid
    $('#ra-form #pets').w2grid({
        name   : 'RAPetsGrid',
        header : 'Pets',
        show   : {
                    toolbar: true,
                    footer: true,
                    toolbarSave: true
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
            {
                field:   'RAID',
                hidden:  true
            },
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
        ]
    });

    // load the existing data in pets component
    setTimeout(function() {
        var partType = app.raFlowPartTypes.pets;
        for (var i = 0; i < app.raflow.data[app.raflow.activeflowID].length; i++) {
            if (partType == app.raflow.data[app.raflow.activeflowID][i].PartType) {
                w2ui.RAPetsGrid.records = app.raflow.data[app.raflow.activeflowID][i].Data;
                w2ui.RAPetsGrid.refresh();
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
    }
}

function loadRAVehiclesGrid() {
    // if form is loaded then return
    if ("RAVehiclesGrid" in w2ui) {
        return;
    }

    // vehicles grid
    $('#ra-form #vehicles').w2grid({
        name   : 'RAVehiclesGrid',
        header : 'Vehicles',
        show   : {
                    toolbar: true,
                    footer: true,
                    toolbarSave: true
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
            /*{
                field:   'VIN',
                caption: 'VIN',
                size:    '80px',
                editable:{ type: 'text' }
            },*/
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
                render:  'money',
                editable:{ type: 'money' }
            },
            {
                field:   'LicensePlateNumber',
                caption: 'License Plate<br>Number',
                size:    '100px',
                render:  'money',
                editable:{ type: 'money' }
            },
            {
                field:   'ParkingPermitNumber',
                caption: 'Parking Permit <br>Number',
                size:    '100px',
                render:  'money',
                editable:{ type: 'money' }
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
        ]
    });

    // load the existing data in vehicles component
    setTimeout(function() {
        var partType = app.raFlowPartTypes.vehicles;
        for (var i = 0; i < app.raflow.data[app.raflow.activeflowID].length; i++) {
            if (partType == app.raflow.data[app.raflow.activeflowID][i].PartType) {
                w2ui.RAVehiclesGrid.records = app.raflow.data[app.raflow.activeflowID][i].Data;
                w2ui.RAVehiclesGrid.refresh();
            }
        }
    }, 500);

}
