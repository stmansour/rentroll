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
                    $("#flow-list").append("<li class='flowID-link' data-flow-id='"+data.records[i].FlowID+"'>"+data.records[i].FlowID+"</li>");
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

window.getRAFlowPartTypeIndex = function(partType) {
    var partTypeIndex = -1;
    if (app.raflow.activeflowID && app.raflow.data[app.raflow.activeflowID]) {
        for (var i = 0; i < app.raflow.data[app.raflow.activeflowID].length; i++) {
            if (partType == app.raflow.data[app.raflow.activeflowID][i].PartType) {
                partTypeIndex = i;
                break;
            }
        }
    }
    return partTypeIndex;
};

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
            var i = getRAFlowPartTypeIndex(app.raFlowPartTypes.people);
            data = app.raflow.data[app.raflow.activeflowID][i].Data;
            // data = w2ui.RAPeopleForm.record;
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
            },
            onRefresh: function(event) {
                var t   = new Date(),
                    nyd = new Date(new Date().setFullYear(new Date().getFullYear() + 1));

                // set default values with start=current day, stop=next year day, if record is blank
                this.record.AgreementStart = this.record.AgreementStart || w2uiDateControlString(t);
                this.record.AgreementStop = this.record.AgreementStop || w2uiDateControlString(nyd);
                this.record.RentStart = this.record.RentStart || w2uiDateControlString(t);
                this.record.RentStop = this.record.RentStop || w2uiDateControlString(nyd);
                this.record.PossessionStart = this.record.PossessionStart || w2uiDateControlString(t);
                this.record.PossessionStop = this.record.PossessionStop || w2uiDateControlString(nyd);
            }
        });

    }

    // load the existing data in dates component
    setTimeout(function() {
        var i = getRAFlowPartTypeIndex(app.raFlowPartTypes.dates);
        if (i >= 0 && app.raflow.data[app.raflow.activeflowID][i].Data) {
            w2ui.RADatesForm.record = app.raflow.data[app.raflow.activeflowID][i].Data;
            w2ui.RADatesForm.refresh();
        } else {
            w2ui.RADatesForm.clear();
        }
    }, 500);
}

// -------------------------------------------------------------------------------
// Rental Agreement - People form
// -------------------------------------------------------------------------------

//-----------------------------------------------------------------------------
// getFullName - returns a string with the full name based on the item supplied.
// @params
//   item = an object assumed to have a FirstName, MiddleName, and LastName
// @return - the full name concatenated together
//-----------------------------------------------------------------------------
window.getFullName = function (item) {

    var s = item.FirstName;
    if (item.MiddleName.length > 0) { s += ' ' + item.MiddleName; }
    if (item.LastName.length > 0 ) { s += ' ' + item.LastName; }
    return s;
};

//-----------------------------------------------------------------------------
// getTCIDName - returns an appropriate name for the supplied item. If
//          the item is a person, then the person's full name is returned.
//          If the item is a company, then the company name is returned.
// @params
//   item = an object assumed to have a FirstName, MiddleName, LastName,
//          IsCompany, and CompanyName.
// @return - the name to render
//-----------------------------------------------------------------------------
window.getTCIDName = function (item) {

    var s = (item.IsCompany > 0) ? item.CompanyName : getFullName(item);

    if (item.TCID > 0) {
        s += ' (TCID: '+ String(item.TCID) + ')';
    }
    return s;
};

//-----------------------------------------------------------------------------
// loadTransactantListingItem - adds transactant into categories list
// @params
//   transactantRec = an object assumed to have a FirstName, MiddleName, LastName,
//                    IsCompany, and CompanyName.
//   IsPayor        = flag to indicate payor or not
//   IsUser         = flag to indicate user or not
//   IsGuarantor    = flag to indicate guarantor or not
// @return - nothing
//-----------------------------------------------------------------------------
window.loadTransactantListingItem = function(transactantRec, IsPayor, IsUser, IsGuarantor) {

    var peoplePartIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.people);
    if (peoplePartIndex < 0) {
        alert("flow data could not be found");
        return false;
    }

    // check that "Payors", "Users", "Guarantors" keys do exist in Data of people
    var peopleTypeKeys = Object.keys(app.raflow.data[app.raflow.activeflowID][peoplePartIndex].Data);
    var payorsIndex = peopleTypeKeys.indexOf("Payors");
    var usersIndex = peopleTypeKeys.indexOf("Users");
    var guarantorsIndex = peopleTypeKeys.indexOf("Guarantors");
    if (payorsIndex < 0 || usersIndex < 0 || guarantorsIndex < 0) {
        alert("flow data could not be found");
        return false;
    }


    // listing item to be appended in ul
    var s = (transactantRec.IsCompany > 0) ? transactantRec.CompanyName : getFullName(transactantRec);
    if (transactantRec.TCID > 0) {
        s += ' (TCID: '+ String(transactantRec.TCID) + ')';
    }

    var peopleListingItem = '<li data-tcid="' + transactantRec.TCID + '">';
    peopleListingItem += '<span>' + s + '</span>';
    peopleListingItem += '<i class="remove-item fas fa-times-circle fa-sm"></i>'
    peopleListingItem += '</li>';

    // add into payor list
    if (IsPayor) {
        // check for duplicacy
        var found = false;
        var length = app.raflow.data[app.raflow.activeflowID][peoplePartIndex].Data.Payors.length;
        for(var i = length - 1; i >= 0; i--) {
            if (app.raflow.activeTransactant.TCID == app.raflow.data[app.raflow.activeflowID][peoplePartIndex].Data.Payors[i].TCID) {
                found = true;
                break;
            }
        }
        if (!(found)) {
            if (!($.isEmptyObject(app.raflow.activeTransactant))) {
                app.raflow.data[app.raflow.activeflowID][peoplePartIndex].Data.Payors.push(app.raflow.activeTransactant);
            }

            // if with this tcid element exists in DOM then not append
            if (!($('#payor-list .people-listing li[data-tcid="'+transactantRec.TCID+'"]').length > 0)) {
                $('#payor-list .people-listing').append(peopleListingItem);
            }
        }
    }

    // add into user list
    if (IsUser) {
        var found = false;
        var length = app.raflow.data[app.raflow.activeflowID][peoplePartIndex].Data.Users.length;
        for(var i = length - 1; i >= 0; i--) {
            if (app.raflow.activeTransactant.TCID == app.raflow.data[app.raflow.activeflowID][peoplePartIndex].Data.Users[i].TCID) {
                found = true;
                break;
            }
        }
        if (!(found)) {
            if (!($.isEmptyObject(app.raflow.activeTransactant))) {
                app.raflow.data[app.raflow.activeflowID][peoplePartIndex].Data.Users.push(app.raflow.activeTransactant);
            }

            // if with this tcid element exists in DOM then not append
            if (!($('#user-list .people-listing li[data-tcid="'+transactantRec.TCID+'"]').length > 0)) {
                $('#user-list .people-listing').append(peopleListingItem);
            }
        }
    }

    // add into guarantor list
    if (IsGuarantor) {
        var found = false;
        var length = app.raflow.data[app.raflow.activeflowID][peoplePartIndex].Data.Guarantors.length;
        for(var i = length - 1; i >= 0; i--) {
            if (app.raflow.activeTransactant.TCID == app.raflow.data[app.raflow.activeflowID][peoplePartIndex].Data.Guarantors[i].TCID) {
                found = true;
                break;
            }
        }
        if (!(found)) {
            if (!($.isEmptyObject(app.raflow.activeTransactant))) {
                app.raflow.data[app.raflow.activeflowID][peoplePartIndex].Data.Guarantors.push(app.raflow.activeTransactant);
            }

            // if with this tcid element exists in DOM then not append
            if (!($('#guarantor-list .people-listing li[data-tcid="'+transactantRec.TCID+'"]').length > 0)) {
                $('#guarantor-list .people-listing').append(peopleListingItem);
            }
        }
    }
}

//-----------------------------------------------------------------------------
// acceptTransactant - add transactant to the list of payor/user/guarantor
//
// @params
//   item = an object assumed to have a FirstName, MiddleName, LastName,
//          IsCompany, and CompanyName.
// @return - the name to render
//-----------------------------------------------------------------------------
window.acceptTransactant = function() {
    var IsPayor = w2ui.RAPeopleForm.record.Payor;
    var IsUser = w2ui.RAPeopleForm.record.User;
    var IsGuarantor = w2ui.RAPeopleForm.record.Guarantor;

    // if not set anything then alert the user to select any one of them
    if (!(IsPayor || IsUser || IsGuarantor)) {
        alert("Please, select the role");
        return false;
    }

    // load item in the DOM
    loadTransactantListingItem(w2ui.RAPeopleForm.record, IsPayor, IsUser, IsGuarantor);

    // clear the form
    app.raflow.activeTransactant = {};
    w2ui.RAPeopleForm.clear();
};

// remove people from the listing
$(document).on('click', '.remove-item', function() {
    var tcid = parseInt($(this).closest('li').attr('data-tcid'));

    // get part type index
    var peoplePartIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.people);

    // remove entry from data
    if (peoplePartIndex >= 0) {
        // check that "Payors", "Users", "Guarantors" keys do exist in Data of people
        var peopleTypeKeys = Object.keys(app.raflow.data[app.raflow.activeflowID][peoplePartIndex].Data);
        var payorsIndex = peopleTypeKeys.indexOf("Payors");
        var usersIndex = peopleTypeKeys.indexOf("Users");
        var guarantorsIndex = peopleTypeKeys.indexOf("Guarantors");

        if (!(payorsIndex < 0 || usersIndex < 0 || guarantorsIndex < 0)) {
            var peopleType = $(this).closest('ul.people-listing').attr('data-people-type');
            switch(peopleType) {
                case "payors":
                    var length = app.raflow.data[app.raflow.activeflowID][peoplePartIndex].Data.Payors.length;
                    for(var i = length - 1; i >= 0; i--) {
                        app.raflow.data[app.raflow.activeflowID][peoplePartIndex].Data.Payors.splice(i, 1);
                    }
                    break;
                case "users":
                    var length = app.raflow.data[app.raflow.activeflowID][peoplePartIndex].Data.Users.length;
                    for(var i = length - 1; i >= 0; i--) {
                        app.raflow.data[app.raflow.activeflowID][peoplePartIndex].Data.Users.splice(i, 1);
                    }
                    break;
                case "guarantors":
                    var length = app.raflow.data[app.raflow.activeflowID][peoplePartIndex].Data.Guarantors.length;
                    for(var i = length - 1; i >= 0; i--) {
                        app.raflow.data[app.raflow.activeflowID][peoplePartIndex].Data.Guarantors.splice(i, 1);
                    }
                    break;
            }
        }
    }


    $(this).closest('li').remove();
});

function loadRAPeopleForm() {

    // have to list down all people into different categories
    var peoplePartIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.people);
    if (peoplePartIndex < 0) {
        alert("flow data could not be found");
        return false;
    }

    // check that "Payors", "Users", "Guarantors" keys do exist in Data of people
    var peopleTypeKeys = Object.keys(app.raflow.data[app.raflow.activeflowID][peoplePartIndex].Data);
    var payorsIndex = peopleTypeKeys.indexOf("Payors");
    var usersIndex = peopleTypeKeys.indexOf("Users");
    var guarantorsIndex = peopleTypeKeys.indexOf("Guarantors");
    if (!(payorsIndex < 0 || usersIndex < 0 || guarantorsIndex < 0)) { // valid then
        // load payors list
        app.raflow.data[app.raflow.activeflowID][peoplePartIndex].Data.Payors.forEach(function(item) {
            loadTransactantListingItem(item, true, false, false);
        });
        // load users list
        app.raflow.data[app.raflow.activeflowID][peoplePartIndex].Data.Users.forEach(function(item) {
            loadTransactantListingItem(item, false, true, false);
        });
        // load guarantors list
        app.raflow.data[app.raflow.activeflowID][peoplePartIndex].Data.Guarantors.forEach(function(item) {
            loadTransactantListingItem(item, false, false, true);
        });
    }


    // if form is loaded then return
    if (!("RAPeopleForm" in w2ui)) {

        // people form
        $('#ra-form #people .form-container').w2form({
            name   : 'RAPeopleForm',
            header : 'People',
            style  : 'display: block;',
            formURL: '/webclient/html/test/formrapeople.html',
            focus: -1,
            fields : [
                { name: 'Transactant', type: 'enum',     required: true, html: { caption: "Transactant" },
                    type: 'enum',
                    options: {
                        url:        '/v1/transactantstd/' + app.raflow.BID,
                        max:        1,
                        renderItem: function(item) {
                            // mark this as transactant as an active
                            app.raflow.activeTransactant = item;
                            var s = getTCIDName(item);
                            w2ui.RAPeopleForm.record.TCID = item.TCID;
                            w2ui.RAPeopleForm.record.FirstName = item.FirstName;
                            w2ui.RAPeopleForm.record.LastName = item.LastName;
                            w2ui.RAPeopleForm.record.MiddleName = item.MiddleName;
                            w2ui.RAPeopleForm.record.CompanyName = item.CompanyName;
                            w2ui.RAPeopleForm.record.IsCompany = item.IsCompany;
                            return s;
                        },
                        renderDrop: function(item) {
                            return getTCIDName(item);
                        },
                        compare:    function(item, search) {
                            var s = getTCIDName(item);
                            s = s.toLowerCase();
                            var srch = search.toLowerCase();
                            var match = (s.indexOf(srch) >= 0);
                            return match;
                        },
                        onNew:      function (event) {
                            //console.log('++ New Item: Do not forget to submit it to the server too', event);
                            $.extend(event.item, { FirstName: '', LastName : event.item.text });
                        },
                        onRemove:   function(event) {
                            event.onComplete = function() {
                                // reset active Transactant to blank object
                                app.raflow.activeTransactant = {};

                                var f = w2ui.RAPeopleForm;
                                // reset payor field related data when removed
                                f.record.TCID = 0;

                                // NOTE: have to trigger manually, b'coz we manually change the record,
                                // otherwise it triggers the change event but it won't get change (Object: {})
                                var event = f.trigger({ phase: 'before', target: f.name, type: 'change', event: event }); // event before
                                if (event.cancelled === true) return false;
                                f.trigger($.extend(event, { phase: 'after' })); // event after
                            };
                        }
                    },
                },
                { name: 'TCID',        type: 'int',      required: true, html: { caption: "TCID" } },
                { name: 'FirstName',   type: 'text',     required: true, html: { caption: "FirstName" } },
                { name: 'LastName',    type: 'text',     required: true, html: { caption: "LastName" } },
                { name: 'MiddleName',  type: 'text',     required: true, html: { caption: "MiddleName" } },
                { name: 'CompanyName', type: 'text',     required: true, html: { caption: "CompanyName" } },
                { name: 'IsCompany',   type: 'int',      required: true, html: { caption: "IsCompany" } },
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
        var i = getRAFlowPartTypeIndex(app.raFlowPartTypes.people);
        if (i >= 0 && app.raflow.data[app.raflow.activeflowID][i].Data) {
            // w2ui.RAPeopleForm.record = app.raflow.data[app.raflow.activeflowID][i].Data;
            w2ui.RAPeopleForm.refresh();
        } else {
            w2ui.RAPeopleForm.clear();
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
        var i = getRAFlowPartTypeIndex(app.raFlowPartTypes.pets);
        if (i >= 0 && app.raflow.data[app.raflow.activeflowID][i].Data) {
            w2ui.RAPetsGrid.records = app.raflow.data[app.raflow.activeflowID][i].Data;
            w2ui.RAPetsGrid.refresh();
        } else {
            w2ui.RAPetsGrid.clear();
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
        var i = getRAFlowPartTypeIndex(app.raFlowPartTypes.vehicles);
        if (i >= 0 && app.raflow.data[app.raflow.activeflowID][i].Data) {
            w2ui.RAVehiclesGrid.records = app.raflow.data[app.raflow.activeflowID][i].Data;
            w2ui.RAVehiclesGrid.refresh();
        } else {
            w2ui.RAVehiclesGrid.clear();
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
        var i = getRAFlowPartTypeIndex(app.raFlowPartTypes.bginfo);
        if (i >= 0 && app.raflow.data[app.raflow.activeflowID][i].Data) {
            w2ui.RABGInfoForm.record = app.raflow.data[app.raflow.activeflowID][i].Data;
            w2ui.RABGInfoForm.refresh();
        } else {
            w2ui.RABGInfoForm.clear();
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
        var i = getRAFlowPartTypeIndex(app.raFlowPartTypes.rentables);
        if (i >= 0 && app.raflow.data[app.raflow.activeflowID][i].Data) {
            w2ui.RARentablesGrid.records = app.raflow.data[app.raflow.activeflowID][i].Data;
            w2ui.RARentablesGrid.refresh();
        } else {
            w2ui.RARentablesGrid.clear();
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
        var i = getRAFlowPartTypeIndex(app.raFlowPartTypes.feesterms);
        if (i >= 0 && app.raflow.data[app.raflow.activeflowID][i].Data) {
            w2ui.RAFeesTermsGrid.records = app.raflow.data[app.raflow.activeflowID][i].Data;
            w2ui.RAFeesTermsGrid.refresh();
        } else {
            w2ui.RAFeesTermsGrid.clear();
        }
    }, 500);
}

