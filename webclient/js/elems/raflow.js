"use strict";

// Next button handling
$("#next").click(function() {
    // get the current component (to be previous one)
    var previous_comp = $(".ra-form-component:visible");

    // get the target component (to be active one)
    var active_comp = previous_comp.next(".ra-form-component");

    // make sure that next component available so we can navigate onto it
    if (active_comp.length === 0) {
        return false;
    }

    // load target section
    loadTargetSection(active_comp.attr("id"), previous_comp.attr("id"));
});

// Previous button handling
$("#previous").click(function() {
    // get the current component (to be previous one)
    var previous_comp = $(".ra-form-component:visible");

    // get the target component (to be active one)
    var active_comp = previous_comp.prev(".ra-form-component");

    // make sure that previous component available so we can navigate onto it
    if (active_comp.length === 0) {
        return false;
    }

    // load target section
    loadTargetSection(active_comp.attr("id"), previous_comp.attr("id"));
});

// link click handling
$("#progressbar a").click(function() {
    var previous_comp = $(".ra-form-component:visible");

    // load target form
    var target = $(this).closest("li").attr("data-target");
    target = target.split('#').join("");

    loadTargetSection(target, previous_comp.attr("id"));

    // because of 'a' tag, return false
    return false;
});

// TODO: we should pass flowID, flowPartID here in arguments
function saveRAFlowPartData(record, partType) {
    var targetFlowID = app.flowIDList[0];
    var flowPartID;

    var flowParts = app.flowData[targetFlowID];

    for (var i = 0; i < flowParts.length; i++) {
        if(partType == flowParts[i].PartType) {
            flowPartID = flowParts[i].FlowPartID;
        }
    }

    // temporary data
    var data = {
        "cmd": "save",
        "FlowPartID": flowPartID,
        "Flow": app.flow,
        "FlowID": targetFlowID,
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

function getAllFlowParts(flowID) {
    $.ajax({
        url: "/v1/flow/1/0",
        method: "POST",
        contentType: "application/json",
        dataType: "json",
        data: JSON.stringify({"cmd": "getFlowParts", "flowID": flowID}),
        success: function(data) {
            app.flowData[flowID] = data.records;
        },
        error: function(data) {
            console.log(data);
        },
    });
}

function getAllFlow(flow) {
    $.ajax({
        url: "/v1/flow/1/0",
        method: "POST",
        contentType: "application/json",
        dataType: "json",
        data: JSON.stringify({"cmd": "getAllFlows", "flow": flow}),
        success: function(data) {
            app.flowIDList = data.records;
            for (var i = 0; i < app.flowIDList.length; i++) {
                getAllFlowParts(app.flowIDList[i]);
            }
        },
        error: function(data) {
            console.log(data);
        },
    });
}

function initFlow(flow) {
    $.ajax({
        url: "/v1/flow/1/0",
        method: "POST",
        contentType: "application/json",
        dataType: "json",
        data: JSON.stringify({"cmd": "init", "flow": flow}),
        success: function(data) {
            app.flowIDList.push(data.flowID);
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
        } else {
            console.log( '**** YIPES! ****  status on /v1/uilists/ = ' + textStatus);
        }
    })
    .fail( function() {
        console.log('Error getting /v1/uilists');
        console.log('*** NO INTERFACE ***');
    });

    // get all flows list
    getAllFlow(app.flow);

    // initiate flow
    // initFlow(app.flow);

    loadRADatesForm();

    // as we load the first section
    $("#ra-form footer button#previous").addClass("disable");
});

// map of component function loaders
var compIDContentLoaderMap = {
    "dates": loadRADatesForm,
    "people": loadRAPeopleForm,
    "pets": /*loadRAPetsForm*/null,
    "vehicles": /*loadRAVehiclesForm*/null,
    "bg-info": /*loadRABginfoForm*/null,
    "rentables": /*loadRARentablesForm*/null,
    "fees-terms": /*loadRAFeesTermsForm*/null,
    "final": /*loadRAFinalForm*/null,
};

// forms mapping
var compIDw2uiForms = {
    "dates": "RADatesForm",
    "people": "RAPeopleForm",
    "pets": "",
    "vehicles": "",
    "bg-info": "",
    "rentables": "",
    "fees-terms": "",
    "final": "",
};

// load form according to target
function loadTargetSection(target, previousCompID) {

    var partType = $("#progressbar li[data-target='#" + previousCompID + "']").index() + 1;

    if($("#progressbar li[data-target='#" + target + "']").hasClass("done")){
        console.log("target has been saved", target);
    } else{
        var validateForm = compIDw2uiForms[previousCompID];
        if (typeof w2ui[validateForm] !== "undefined") {
            var issues = w2ui[validateForm].validate();
            if (Array.isArray(issues) && issues.length > 0) {
                alert("form is not valid");
                return;
            } else {
                saveRAFlowPartData(w2ui[validateForm].record, partType);
                $("#progressbar li[data-target='#" + previousCompID + "']").addClass("done");
            }
        } else{
            console.log("unknown previousCompID from nav li: ", previousCompID);
            // return; // as of now just let it go
        }
    }

    $("#progressbar li[data-target='#" + previousCompID + "']").removeClass("active");
    $(".ra-form-component#" + previousCompID).hide();

    $("#progressbar li[data-target='#" + target + "']").addClass("active");
    $(".ra-form-component#" + target).show();

    // hide previous if the target is in first section
    if ($(".ra-form-component#" + target).is($(".ra-form-component").first())) {
        $("#ra-form footer button#previous").addClass("disable");
    } else {
        $("#ra-form footer button#previous").removeClass("disable");
    }

    // hide next if the target is in last section
    if ($(".ra-form-component#" + target).is($(".ra-form-component").last())) {
        $("#ra-form footer button#next").addClass("disable");
    } else {
        $("#ra-form footer button#next").removeClass("disable");
    }

    var targetLoader = compIDContentLoaderMap[target];
    if (typeof targetLoader === "function") {
        targetLoader();
        /*setTimeout(function() {
            var validateForm = compIDw2uiForms[previousCompID];
            if (typeof w2ui[validateForm] !== "undefined") {
                var issues = w2ui[validateForm].validate();
                if (!(Array.isArray(issues) && issues.length > 0)) {
                    // $("#progressbar li[data-target='#" + previousCompID + "']").addClass("done");
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
        url    : '/v1/ra/dates/',
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
            save: function () {
                this.save();
            }
        }
    });
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
        url    : '/v1/ra/people/',
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
}
