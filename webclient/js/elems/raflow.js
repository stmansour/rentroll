/*global
    getFullName, getTCIDName, loadTransactantListingItem,
    initRAFlowAJAX, getRAFlowAllParts, getAllRAFlows, loadRADatesForm, loadRAPeopleForm,
    loadRAPetsGrid, loadRAVehiclesGrid, loadRABGInfoForm, loadRARentablesGrid,
    loadRAFeesTermsGrid, getRAFlowPartTypeIndex, loadTargetSection,
    getVehicleGridInitalRecord, getRentablesGridInitalRecord, getFeesTermsGridInitalRecord,
    getPetsGridInitalRecord, saveActiveCompData, loadRABGInfoForm, w2render,
    requiredFieldsFulFilled
*/

"use strict";

// Next button handling
$(document).on('click', '#ra-form #next', function () {
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
$(document).on('click', '#ra-form #previous', function () {
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
$(document).on('click', '#ra-form #progressbar a', function () {
    var active_comp = $(".ra-form-component:visible");

    // load target form
    var target = $(this).closest("li").attr("data-target");
    target = target.split('#').join("");

    loadTargetSection(target, active_comp.attr("id"));

    // because of 'a' tag, return false
    return false;
});

// TODO: we should pass FlowID, flowPartID here in arguments
window.saveActiveCompData = function (record, partType) {

    var bid = getCurrentBID();

    var flowPartID;
    var flowParts = app.raflow.data[app.raflow.activeFlowID] || [];

    for (var i = 0; i < flowParts.length; i++) {
        if (partType == flowParts[i].PartType) {
            flowPartID = flowParts[i].FlowPartID;
        }
    }

    // temporary data
    var data = {
        "cmd": "save",
        "FlowPartID": flowPartID,
        "Flow": app.raflow.name,
        "FlowID": app.raflow.activeFlowID,
        "BID": bid,
        "PartType": partType,
        "Data": record,
    };

    $.ajax({
        url: "/v1/flowpart/" + bid.toString() + "/0",
        method: "POST",
        contentType: "application/json",
        dataType: "json",
        data: JSON.stringify(data),
        success: function (data) {
            if (data.status != "error") {
                console.log("data has been saved for: ", app.raflow.activeFlowID, ", partType: ", partType);
            } else {
                console.error(data.message);
            }
        },
        error: function (data) {
            console.log(data);
        },
    });
};

window.getRAFlowAllParts = function (FlowID) {
    var bid = getCurrentBID();

    $.ajax({
        url: "/v1/flow/" + bid.toString() + "/0",
        method: "POST",
        contentType: "application/json",
        dataType: "json",
        data: JSON.stringify({"cmd": "getFlowParts", "FlowID": FlowID}),
        success: function (data) {
            if (data.status != "error") {
                app.raflow.data[FlowID] = data.records;

                // show "done" mark on each li of navigation bar
                for (var comp in app.raFlowPartTypes) {
                    // if required fields are fulfilled then mark this slide as done
                    if (requiredFieldsFulFilled(comp)) {
                        // hide active component
                        $("#progressbar li[data-target='#" + comp + "']").addClass("done");
                    }
                }

                // mark first slide as active
                $(".ra-form-component#dates").show();
                $("#progressbar li[data-target='#dates']").removeClass("done").addClass("active");
                loadRADatesForm();

            } else {
                console.error(data.message);
            }
        },
        error: function (data) {
            console.log(data);
        },
    });
};

window.initRAFlowAJAX = function () {
    var bid = getCurrentBID();

    return $.ajax({
        url: "/v1/flow/" + bid.toString() + "/0",
        method: "POST",
        contentType: "application/json",
        dataType: "json",
        data: JSON.stringify({"cmd": "init", "flow": app.raflow.name}),
        success: function (data) {
            if (data.status != "error") {
                app.raflow.data[data.FlowID] = {};
            }
        },
        error: function (data) {
            console.log(data);
        },
    });
};

window.getRAFlowPartTypeIndex = function (partType) {
    var partTypeIndex = -1;
    if (app.raflow.activeFlowID && app.raflow.data[app.raflow.activeFlowID]) {
        for (var i = 0; i < app.raflow.data[app.raflow.activeFlowID].length; i++) {
            if (partType == app.raflow.data[app.raflow.activeFlowID][i].PartType) {
                partTypeIndex = i;
                break;
            }
        }
    }
    return partTypeIndex;
};

window.requiredFieldsFulFilled = function (compID) {
    var done = false;

    // if not active flow id then return
    if (app.raflow.activeFlowID === "") {
        console.log("no active flow ID");
        return done;
    }

    // get part type index for the component
    var partType = app.raFlowPartTypes[compID];
    var partTypeIndex = getRAFlowPartTypeIndex(partType);
    if (partTypeIndex === -1) {
        console.log("no index found this part type");
        return done;
    }

    var data;

    switch (compID) {
        case "dates":
            data = app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data;
            var validData = true;
            for (var dateKey in data) {
                // if anything else then break and mark as invalid
                if (!(typeof data[dateKey] === "string" && data[dateKey] !== "")) {
                    validData = false;
                    break;
                }
            }
            // if loop passed successfully then mark it as successfully
            done = validData;
            break;
        case "people":
            data = app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data;
            if (data.Users.length > 0 && data.Payors.length > 0) {
                done = true;
            }
            break;
        case "pets":
            data = app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data;
            if (data.length > 0) {
                done = true;
            }
            break;
        case "vehicles":
            data = app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data;
            if (data.length > 0) {
                done = true;
            }
            break;
        case "bginfo":
            // TODO(Akshay): Add for integer fields e.g., phone, gross wage.
            data = app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data;
            validData = true;
            // list of fields which must have value and it's type string
            var listOfRequiredField = ["application_date", "move_in_date",
                "apt_no", "lt", "applicant_first_name", "applicant_middle_name",
                "applicant_last_name", "applicant_dob", "applicant_ssn",
                "applicant_dln", "applicant_telno", "applicant_email",
                "no_people_apt", "c_address", "cll_name", "cll_phone",
                "clr", "cresmove", "applicant_employer", "applicant_phone", "applicant_address",
                "applicant_position", "ec_name", "ec_phone", "ec_address"];
            for (var field in listOfRequiredField) {
                if (data[field] === "") {
                    validData = false;
                    break;
                }
            }
            done = validData;
            break;
        case "rentables":
            data = app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data;
            if (data.length > 0) {
                done = true;
            }
            break;
        case "feesterms":
            data = app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data;
            if (data.length > 0) {
                done = true;
            }
            break;
        case "final":
            break;
    }

    return done;
};

// load form according to target
window.loadTargetSection = function (target, activeCompID) {

    /*if ($("#progressbar li[data-target='#" + target + "']").hasClass("done")) {
        console.log("target has been saved", target);
    } else {}*/

    // if required fields are fulfilled then mark this slide as done
    if (requiredFieldsFulFilled(activeCompID)) {
        // hide active component
        $("#progressbar li[data-target='#" + activeCompID + "']").addClass("done");
    }

    // decide data based on type
    var data = null;
    switch (activeCompID) {
        case "dates":
            data = w2ui.RADatesForm.record;
            break;
        case "people":
            var i = getRAFlowPartTypeIndex(app.raFlowPartTypes.people);
            data = app.raflow.data[app.raflow.activeFlowID][i].Data;
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

    // get part type from the class index
    var partType = $("#progressbar li[data-target='#" + activeCompID + "']").index() + 1;
    if (data) {
        // save the content on server for active component
        saveActiveCompData(data, partType);
    }

    // hide active component
    $("#progressbar li[data-target='#" + activeCompID + "']").removeClass("active");
    $(".ra-form-component#" + activeCompID).hide();

    // show target component
    $("#progressbar li[data-target='#" + target + "']").removeClass("done").addClass("active");
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
    } else {
        console.log("unknown target from nav li: ", target);
    }
};

// -------------------------------------------------------------------------------
// Rental Agreement - Info Dates form
// -------------------------------------------------------------------------------
window.loadRADatesForm = function () {

    // if form is loaded then return
    if (!("RADatesForm" in w2ui)) {
        // dates form
        $().w2form({
            name: 'RADatesForm',
            header: 'Dates',
            style: 'border: 1px black solid; display: block;',
            focus: -1,
            formURL: '/webclient/html/formradates.html',
            fields: [
                {name: 'AgreementStart', type: 'date', required: true, html: {caption: "Term Start"}},
                {name: 'AgreementStop', type: 'date', required: true, html: {caption: "Term Stop"}},
                {name: 'RentStart', type: 'date', required: true, html: {caption: "Rent Start"}},
                {name: 'RentStop', type: 'date', required: true, html: {caption: "Rent Stop"}},
                {name: 'PossessionStart', type: 'date', required: true, html: {caption: "Possession Start"}},
                {name: 'PossessionStop', type: 'date', required: true, html: {caption: "Possession Stop"}}
            ],
            actions: {
                reset: function () {
                    this.clear();
                },
            },
            onRefresh: function (event) {
                var t = new Date(),
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

    // now render the form in specifiec targeted division
    $('#ra-form #dates').w2render(w2ui.RADatesForm);

    // load the existing data in dates component
    setTimeout(function () {
        var i = getRAFlowPartTypeIndex(app.raFlowPartTypes.dates);
        if (i >= 0 && app.raflow.data[app.raflow.activeFlowID][i].Data) {
            w2ui.RADatesForm.record = app.raflow.data[app.raflow.activeFlowID][i].Data;
            w2ui.RADatesForm.refresh();
        } else {
            w2ui.RADatesForm.clear();
        }
    }, 500);
};

// -------------------------------------------------------------------------------
// Rental Agreement - People form
// -------------------------------------------------------------------------------

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
window.loadTransactantListingItem = function (transactantRec, IsPayor, IsUser, IsGuarantor) {

    var peoplePartIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.people);
    if (peoplePartIndex < 0) {
        alert("flow data could not be found");
        return false;
    }

    // check that "Payors", "Users", "Guarantors" keys do exist in Data of people
    var peopleTypeKeys = Object.keys(app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data);
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
        s += ' (TCID: ' + String(transactantRec.TCID) + ')';
    }

    var peopleListingItem = '<li data-tcid="' + transactantRec.TCID + '">';
    peopleListingItem += '<span>' + s + '</span>';
    peopleListingItem += '<i class="remove-item fas fa-times-circle fa-xs"></i>';
    peopleListingItem += '</li>';

    var i, length, found = false;

    // add into payor list
    if (IsPayor) {
        // check for duplicacy
        found = false;
        length = app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data.Payors.length;
        for (i = length - 1; i >= 0; i--) {
            if (app.raflow.activeTransactant.TCID == app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data.Payors[i].TCID) {
                found = true;
                break;
            }
        }
        if (!(found)) {
            if (!($.isEmptyObject(app.raflow.activeTransactant))) {
                app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data.Payors.push(app.raflow.activeTransactant);
            }

            // if with this tcid element exists in DOM then not append
            if ($('#payor-list .people-listing li[data-tcid="' + transactantRec.TCID + '"]').length < 1) {
                $('#payor-list .people-listing').append(peopleListingItem);
            }
        }
    }

    // add into user list
    if (IsUser) {
        found = false;
        length = app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data.Users.length;
        for (i = length - 1; i >= 0; i--) {
            if (app.raflow.activeTransactant.TCID == app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data.Users[i].TCID) {
                found = true;
                break;
            }
        }
        if (!(found)) {
            if (!($.isEmptyObject(app.raflow.activeTransactant))) {
                app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data.Users.push(app.raflow.activeTransactant);
            }

            // if with this tcid element exists in DOM then not append
            if ($('#user-list .people-listing li[data-tcid="' + transactantRec.TCID + '"]').length < 1) {
                $('#user-list .people-listing').append(peopleListingItem);
            }
        }
    }

    // add into guarantor list
    if (IsGuarantor) {
        found = false;
        length = app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data.Guarantors.length;
        for (i = length - 1; i >= 0; i--) {
            if (app.raflow.activeTransactant.TCID == app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data.Guarantors[i].TCID) {
                found = true;
                break;
            }
        }
        if (!(found)) {
            if (!($.isEmptyObject(app.raflow.activeTransactant))) {
                app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data.Guarantors.push(app.raflow.activeTransactant);
            }

            // if with this tcid element exists in DOM then not append
            if ($('#guarantor-list .people-listing li[data-tcid="' + transactantRec.TCID + '"]').length < 1) {
                $('#guarantor-list .people-listing').append(peopleListingItem);
            }
        }
    }
};

//-----------------------------------------------------------------------------
// acceptTransactant - add transactant to the list of payor/user/guarantor
//
// @params
//   item = an object assumed to have a FirstName, MiddleName, LastName,
//          IsCompany, and CompanyName.
// @return - the name to render
//-----------------------------------------------------------------------------
window.acceptTransactant = function () {
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
$(document).on('click', '.people-listing .remove-item', function () {
    var tcid = parseInt($(this).closest('li').attr('data-tcid'));

    // get part type index
    var peoplePartIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.people);

    // remove entry from data
    if (peoplePartIndex >= 0) {
        // check that "Payors", "Users", "Guarantors" keys do exist in Data of people
        var peopleTypeKeys = Object.keys(app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data);
        var payorsIndex = peopleTypeKeys.indexOf("Payors");
        var usersIndex = peopleTypeKeys.indexOf("Users");
        var guarantorsIndex = peopleTypeKeys.indexOf("Guarantors");

        if (!(payorsIndex < 0 || usersIndex < 0 || guarantorsIndex < 0)) {
            var peopleType = $(this).closest('ul.people-listing').attr('data-people-type');
            var i, length;
            switch (peopleType) {
                case "payors":
                    length = app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data.Payors.length;
                    for (i = length - 1; i >= 0; i--) {
                        app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data.Payors.splice(i, 1);
                    }
                    break;
                case "users":
                    length = app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data.Users.length;
                    for (i = length - 1; i >= 0; i--) {
                        app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data.Users.splice(i, 1);
                    }
                    break;
                case "guarantors":
                    length = app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data.Guarantors.length;
                    for (i = length - 1; i >= 0; i--) {
                        app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data.Guarantors.splice(i, 1);
                    }
                    break;
            }
        }
    }


    $(this).closest('li').remove();
});

window.loadRAPeopleForm = function () {

    // have to list down all people into different categories
    var peoplePartIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.people);
    if (peoplePartIndex < 0) {
        alert("flow data could not be found");
        return false;
    }

    // check that "Payors", "Users", "Guarantors" keys do exist in Data of people
    var peopleTypeKeys = Object.keys(app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data);
    var payorsIndex = peopleTypeKeys.indexOf("Payors");
    var usersIndex = peopleTypeKeys.indexOf("Users");
    var guarantorsIndex = peopleTypeKeys.indexOf("Guarantors");
    if (!(payorsIndex < 0 || usersIndex < 0 || guarantorsIndex < 0)) { // valid then
        // load payors list
        app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data.Payors.forEach(function (item) {
            loadTransactantListingItem(item, true, false, false);
        });
        // load users list
        app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data.Users.forEach(function (item) {
            loadTransactantListingItem(item, false, true, false);
        });
        // load guarantors list
        app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data.Guarantors.forEach(function (item) {
            loadTransactantListingItem(item, false, false, true);
        });
    }


    // if form is loaded then return
    if (!("RAPeopleForm" in w2ui)) {

        // people form
        $().w2form({
            name: 'RAPeopleForm',
            header: 'People',
            style: 'display: block;',
            formURL: '/webclient/html/formrapeople.html',
            focus: -1,
            fields: [
                {
                    name: 'Transactant', type: 'enum', required: true, html: {caption: "Transactant"},
                    options: {
                        url: '/v1/transactantstd/' + app.raflow.BID,
                        max: 1,
                        renderItem: function (item) {
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
                        renderDrop: function (item) {
                            return getTCIDName(item);
                        },
                        compare: function (item, search) {
                            var s = getTCIDName(item);
                            s = s.toLowerCase();
                            var srch = search.toLowerCase();
                            var match = (s.indexOf(srch) >= 0);
                            return match;
                        },
                        onNew: function (event) {
                            //console.log('++ New Item: Do not forget to submit it to the server too', event);
                            $.extend(event.item, {FirstName: '', LastName: event.item.text});
                        },
                        onRemove: function (event) {
                            event.onComplete = function () {
                                // reset active Transactant to blank object
                                app.raflow.activeTransactant = {};

                                var f = w2ui.RAPeopleForm;
                                // reset payor field related data when removed
                                f.record.TCID = 0;

                                // NOTE: have to trigger manually, b'coz we manually change the record,
                                // otherwise it triggers the change event but it won't get change (Object: {})
                                var event = f.trigger({phase: 'before', target: f.name, type: 'change', event: event}); // event before
                                if (event.cancelled === true) return false;
                                f.trigger($.extend(event, {phase: 'after'})); // event after
                            };
                        }
                    },
                },
                {name: 'TCID', type: 'int', required: true, html: {caption: "TCID"}},
                {name: 'FirstName', type: 'text', required: true, html: {caption: "FirstName"}},
                {name: 'LastName', type: 'text', required: true, html: {caption: "LastName"}},
                {name: 'MiddleName', type: 'text', required: true, html: {caption: "MiddleName"}},
                {name: 'CompanyName', type: 'text', required: true, html: {caption: "CompanyName"}},
                {name: 'IsCompany', type: 'int', required: true, html: {caption: "IsCompany"}},
                {name: 'Payor', type: 'checkbox', required: true, html: {caption: "Payor"}},
                {name: 'User', type: 'checkbox', required: true, html: {caption: "User"}},
                {name: 'Guarantor', type: 'checkbox', required: true, html: {caption: "Guarantor"}},
            ],
            actions: {
                reset: function () {
                    this.clear();
                }
            }
        });
    }

    // load form in div
    $('#ra-form #people .form-container').w2render(w2ui.RAPeopleForm);

    // load the existing data in people component
    setTimeout(function () {
        var i = getRAFlowPartTypeIndex(app.raFlowPartTypes.people);
        if (i >= 0 && app.raflow.data[app.raflow.activeFlowID][i].Data) {
            // w2ui.RAPeopleForm.record = app.raflow.data[app.raflow.activeFlowID][i].Data;
            w2ui.RAPeopleForm.refresh();
        } else {
            w2ui.RAPeopleForm.clear();
        }
    }, 500);
};

// -------------------------------------------------------------------------------
// Rental Agreement - Pets Grid
// -------------------------------------------------------------------------------
window.getPetsGridInitalRecord = function (BID, gridLen) {
    var t = new Date(),
        nyd = new Date(new Date().setFullYear(new Date().getFullYear() + 1));

    return {
        recid: gridLen,
        PETID: 0,
        BID: BID,
        RAID: 0,
        Name: "",
        Type: "",
        Breed: "",
        Color: "",
        Weight: 0,
        DtStart: w2uiDateControlString(t),
        DtStop: w2uiDateControlString(nyd),
        RefundablePetDeposit: 0.0,
        RecurringPetFee: 0.0,
        NonRefundablePetFee: 0.0
    };
};

window.loadRAPetsGrid = function () {
    // if form is loaded then return
    if (!("RAPetsGrid" in w2ui)) {

        // pets grid
        $().w2grid({
            name: 'RAPetsGrid',
            header: 'Pets',
            show: {
                toolbar: true,
                footer: true,
            },
            style: 'border: 1px solid black; display: block;',
            toolbar: {
                items: [
                    {id: 'add', type: 'button', caption: 'Add Record', icon: 'w2ui-icon-plus'}
                ],
                onClick: function (event) {
                    var bid = getCurrentBID();
                    if (event.target == 'add') {
                        var inital = getPetsGridInitalRecord(bid, w2ui.RAPetsGrid.records.length);
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
                    field: 'PETID',
                    hidden: true
                },
                {
                    field: 'BID',
                    hidden: true
                },
                {
                    field: 'Name',
                    caption: 'Name',
                    size: '150px',
                    editable: {type: 'text'}
                },
                {
                    field: 'Type',
                    caption: 'Type',
                    size: '80px',
                    editable: {type: 'text'}
                },
                {
                    field: 'Breed',
                    caption: 'Breed',
                    size: '80px',
                    editable: {type: 'text'}
                },
                {
                    field: 'Color',
                    caption: 'Color',
                    size: '80px',
                    editable: {type: 'text'}
                },
                {
                    field: 'Weight',
                    caption: 'Weight',
                    size: '80px',
                    editable: {type: 'int'}
                },
                {
                    field: 'DtStart',
                    caption: 'DtStart',
                    size: '100px',
                    editable: {type: 'date'}
                },
                {
                    field: 'DtStop',
                    caption: 'DtStop',
                    size: '100px',
                    editable: {type: 'date'}
                },
                {
                    field: 'NonRefundablePetFee',
                    caption: 'NonRefundable<br>PetFee',
                    size: '70px',
                    render: 'money',
                    editable: {type: 'money'}
                },
                {
                    field: 'RefundablePetDeposit',
                    caption: 'Refundable<br>PetDeposit',
                    size: '70px',
                    render: 'money',
                    editable: {type: 'money'}
                },
                {
                    field: 'RecurringPetFee',
                    caption: 'Recurring<br>PetFee',
                    size: '100%',
                    render: 'money',
                    editable: {type: 'money'}
                },
            ],
            onChange: function (event) {
                event.onComplete = function () {
                    this.save();
                };
            }
        });
    }

    // now load grid in division
    $('#ra-form #pets').w2render(w2ui.RAPetsGrid);

    // load the existing data in pets component
    setTimeout(function () {
        var i = getRAFlowPartTypeIndex(app.raFlowPartTypes.pets);
        if (i >= 0 && app.raflow.data[app.raflow.activeFlowID][i].Data) {
            w2ui.RAPetsGrid.records = app.raflow.data[app.raflow.activeFlowID][i].Data;
            w2ui.RAPetsGrid.refresh();
        } else {
            w2ui.RAPetsGrid.clear();
        }
    }, 500);

};

// -------------------------------------------------------------------------------
// Rental Agreement - Vehicles Grid
// -------------------------------------------------------------------------------
window.getVehicleGridInitalRecord = function (BID, gridLen) {
    var t = new Date(),
        nyd = new Date(new Date().setFullYear(new Date().getFullYear() + 1));

    return {
        recid: gridLen,
        VID: 0,
        BID: BID,
        TCID: 0,
        VIN: "",
        Type: "",
        Make: "",
        Model: "",
        Color: "",
        LicensePlateState: "",
        LicensePlateNumber: "",
        ParkingPermitNumber: "",
        DtStart: w2uiDateControlString(t),
        DtStop: w2uiDateControlString(nyd),
    };
};

window.loadRAVehiclesGrid = function () {
    // if form is loaded then return
    if (!("RAVehiclesGrid" in w2ui)) {

        // vehicles grid
        $().w2grid({
            name: 'RAVehiclesGrid',
            header: 'Vehicles',
            show: {
                toolbar: true,
                footer: true,
            },
            style: 'border: 1px solid black; display: block;',
            toolbar: {
                items: [
                    {id: 'add', type: 'button', caption: 'Add Record', icon: 'w2ui-icon-plus'}
                ],
                onClick: function (event) {
                    var bid = getCurrentBID();
                    if (event.target == 'add') {
                        var inital = getVehicleGridInitalRecord(bid, w2ui.RAVehiclesGrid.records.length);
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
                    field: 'VID',
                    hidden: true
                },
                {
                    field: 'BID',
                    hidden: true
                },
                {
                    field: 'TCID',
                    hidden: true
                },
                {
                    field: 'Type',
                    caption: 'Type',
                    size: '80px',
                    editable: {type: 'text'}
                },
                {
                    field: 'VIN',
                    caption: 'VIN',
                    size: '80px',
                    editable: {type: 'text'}
                },
                {
                    field: 'Make',
                    caption: 'Make',
                    size: '80px',
                    editable: {type: 'text'}
                },
                {
                    field: 'Model',
                    caption: 'Model',
                    size: '80px',
                    editable: {type: 'text'}
                },
                {
                    field: 'Color',
                    caption: 'Color',
                    size: '80px',
                    editable: {type: 'text'}
                },
                {
                    field: 'LicensePlateState',
                    caption: 'License Plate<br>State',
                    size: '100px',
                    editable: {type: 'text'}
                },
                {
                    field: 'LicensePlateNumber',
                    caption: 'License Plate<br>Number',
                    size: '100px',
                    editable: {type: 'text'}
                },
                {
                    field: 'ParkingPermitNumber',
                    caption: 'Parking Permit <br>Number',
                    size: '100px',
                    editable: {type: 'text'}
                },
                {
                    field: 'DtStart',
                    caption: 'DtStart',
                    size: '100px',
                    editable: {type: 'date'}
                },
                {
                    field: 'DtStop',
                    caption: 'DtStop',
                    size: '100%',
                    editable: {type: 'date'}
                },
            ],
            onChange: function (event) {
                event.onComplete = function () {
                    this.save();
                };
            }
        });
    }

    // now load grid in target division
    $('#ra-form #vehicles').w2render(w2ui.RAVehiclesGrid);

    // load the existing data in vehicles component
    setTimeout(function () {
        var i = getRAFlowPartTypeIndex(app.raFlowPartTypes.vehicles);
        if (i >= 0 && app.raflow.data[app.raflow.activeFlowID][i].Data) {
            w2ui.RAVehiclesGrid.records = app.raflow.data[app.raflow.activeFlowID][i].Data;
            w2ui.RAVehiclesGrid.refresh();
        } else {
            w2ui.RAVehiclesGrid.clear();
        }
    }, 500);
};

// -------------------------------------------------------------------------------
// Rental Agreement - Background info form
// -------------------------------------------------------------------------------
window.loadRABGInfoForm = function () {

    // if form is loaded then return
    if (!("RABGInfoForm" in w2ui)) {

        // background info form
        $().w2form({
            name: 'RABGInfoForm',
            header: 'Background Information',
            style: 'border: 1px solid black; display: block;',
            formURL: '/webclient/html/formrabginfo.html',
            focus: -1,
            fields: [
                {field: 'application_date', type: 'date', required: true},
                {field: 'move_in_date', type: 'date', required: true},
                {field: 'apt_no', type: 'alphanumeric', required: true}, // Apartment number
                {field: 'lt', type: 'text', required: true}, // Lease term
                {field: 'applicant_first_name', type: 'text', required: true},
                {field: 'applicant_middle_name', type: 'text', required: true},
                {field: 'applicant_last_name', type: 'text', required: true},
                {field: 'applicant_dob', type: 'date', required: true}, // Date of births of applicants
                {field: 'applicant_ssn', type: 'int', required: true}, // Social security number of applicants
                {field: 'applicant_dln', type: 'alphanumeric', required: true}, // Driving licence number of applicants
                {field: 'applicant_telno', type: 'text', required: true}, // Telephone no of applicants
                {field: 'applicant_email', type: 'email', required: true}, // Email Address of applicants
                {field: 'co_applicant_first_name', type: 'text'},
                {field: 'co_applicant_middle_name', type: 'text'},
                {field: 'co_applicant_last_name', type: 'text'},
                {field: 'co_applicant_dob', type: 'date'}, // Date of births of co-applicants
                {field: 'co_applicant_ssn', type: 'int'}, // Social security number of co-applicants
                {field: 'co_applicant_dln', type: 'alphanumeric'}, // Driving licence number of co-applicants
                {field: 'co_applicant_telno', type: 'text'}, // Telephone no of co-applicants
                {field: 'co_applicant_email', type: 'email'}, // Email Address of co-applicants
                {field: 'no_people_apt', type: 'int', required: true}, // No. of people occupying apartment
                {field: 'c_address', type: 'text', required: true}, // Current Address
                {field: 'cll_name', type: 'text', required: true}, // Current landlord's name
                {field: 'cll_phone', type: 'text', required: true}, // Current landlord's phone number
                {field: 'clr', type: 'text', required: true}, // Length of residency at current address
                {field: 'cresmove', type: 'text', required: true}, // Reason of moving from current address
                {field: 'p_address', type: 'text'}, // Prior Address
                {field: 'pll_name', type: 'text'}, // Prior landlord's name
                {field: 'pll_phone', type: 'text'}, // Prior landlord's phone number
                {field: 'plr', type: 'text'}, // Length of residency at Prior address
                {field: 'presmove', type: 'text'}, // Reason of moving from Prior address
                {field: 'evicted', type: 'checkbox', required: false}, // have you ever been evicted
                {field: 'crime', type: 'checkbox', required: false}, // have you ever been Arrested or convicted of a crime
                {field: 'bankruptcy', type: 'checkbox', required: false}, // have you ever been Declared Bankruptcy
                {field: 'applicant_employer', type: 'text', required: true},
                {field: 'applicant_phone', type: 'text', required: true},
                {field: 'applicant_address', type: 'text', required: true},
                {field: 'applicant_position', type: 'text', required: true},
                {field: 'applicant_gw', type: 'money', required: true},
                {field: 'co_applicant_employer', type: 'text'},
                {field: 'co_applicant_phone', type: 'text'},
                {field: 'co_applicant_address', type: 'text'},
                {field: 'co_applicant_position', type: 'text'},
                {field: 'co_applicant_gw', type: 'money'},
                {field: 'comment', type: 'text'}, // In an effort to accommodate you, please advise us of any special needs
                {field: 'ec_name', type: 'text', required: true}, // Name of emergency contact
                {field: 'ec_phone', type: 'text', required: true}, // Phone number of emergency contact
                {field: 'ec_address', type: 'text', required: true} // Address of emergency contact
            ],
            actions: {
                reset: function () {
                    this.clear();
                }
                /*save: function () {
                    this.save();
                }*/
            }
        });
    }

    // now load form in div
    $('#ra-form #bginfo').w2render(w2ui.RABGInfoForm);

    // load the existing data in people component
    setTimeout(function () {
        var i = getRAFlowPartTypeIndex(app.raFlowPartTypes.bginfo);
        if (i >= 0 && app.raflow.data[app.raflow.activeFlowID][i].Data) {
            w2ui.RABGInfoForm.record = app.raflow.data[app.raflow.activeFlowID][i].Data;
            w2ui.RABGInfoForm.refresh();
        } else {
            w2ui.RABGInfoForm.clear();
        }
    }, 500);
};

// -------------------------------------------------------------------------------
// Rental Agreement - Rentables Grid
// -------------------------------------------------------------------------------
window.getRentablesGridInitalRecord = function (BID, gridLen) {
    return {
        recid: gridLen,
        RID: 0,
        BID: BID,
        RTID: 0,
        RentableName: "",
        ContractRent: 0.0,
        ProrateAmt: 0.0,
        TaxableAmt: 0.0,
        SalesTax: 0.0,
        TransOCC: 0.0,
    };
};

window.loadRARentablesGrid = function () {
    // if form is loaded then return
    if (!("RARentablesGrid" in w2ui)) {

        // rentables grid
        $().w2grid({
            name: 'RARentablesGrid',
            header: 'Rentables',
            show: {
                toolbar: true,
                footer: true,
            },
            style: 'border: 1px solid black; display: block;',
            toolbar: {
                items: [
                    {id: 'add', type: 'button', caption: 'Add Record', icon: 'w2ui-icon-plus'}
                ],
                onClick: function (event) {
                    var bid = getCurrentBID();
                    if (event.target == 'add') {
                        var inital = getRentablesGridInitalRecord(bid, w2ui.RARentablesGrid.records.length);
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
                    field: 'RID',
                    hidden: true
                },
                {
                    field: 'BID',
                    hidden: true
                },
                {
                    field: 'RTID',
                    hidden: true
                },
                {
                    field: 'RentableName',
                    caption: 'Rentable',
                    size: '350px',
                    editable: {type: 'text'}
                },
                {
                    field: 'ContractRent',
                    caption: 'At Signing',
                    size: '100px',
                    render: 'money',
                    editable: {type: 'money'}
                },
                {
                    field: 'ProrateAmt',
                    caption: 'Prorate',
                    size: '100px',
                    render: 'money',
                    editable: {type: 'money'}
                },
                {
                    field: 'TaxableAmt',
                    caption: 'Taxable Amt',
                    size: '100px',
                    render: 'money',
                    editable: {type: 'money'}
                },
                {
                    field: 'SalesTax',
                    caption: 'Sales Tax',
                    size: '100px',
                    render: 'money',
                    editable: {type: 'money'}
                },
                {
                    field: 'TransOCC',
                    caption: 'Trans OCC',
                    size: '100%',
                    render: 'money',
                    editable: {type: 'money'}
                }
            ],
            onChange: function (event) {
                event.onComplete = function () {
                    this.save();
                };
            }
        });
    }

    // now load grid in division
    $('#ra-form #rentables').w2render(w2ui.RARentablesGrid);

    // load the existing data in rentables component
    setTimeout(function () {
        var i = getRAFlowPartTypeIndex(app.raFlowPartTypes.rentables);
        if (i >= 0 && app.raflow.data[app.raflow.activeFlowID][i].Data) {
            w2ui.RARentablesGrid.records = app.raflow.data[app.raflow.activeFlowID][i].Data;
            w2ui.RARentablesGrid.refresh();
        } else {
            w2ui.RARentablesGrid.clear();
        }
    }, 500);
};


// -------------------------------------------------------------------------------
// Rental Agreement - Fees Terms Grid
// -------------------------------------------------------------------------------
window.getFeesTermsGridInitalRecord = function (BID, gridLen) {
    return {
        recid: gridLen,
        RID: 0,
        BID: BID,
        RTID: 0,
        RentableName: "",
        FeeName: "",
        Amount: 0.0,
        Cycle: 6,
        SigningAmt: 0.0,
        ProrateAmt: 0.0,
        TaxableAmt: 0.0,
        SalesTax: 0.0,
        TransOCC: 0.0,
    };
};

window.loadRAFeesTermsGrid = function () {
    // if form is loaded then return
    if (!("RAFeesTermsGrid" in w2ui)) {

        // feesterms grid
        $().w2grid({
            name: 'RAFeesTermsGrid',
            header: 'FeesTerms',
            show: {
                toolbar: true,
                footer: true,
            },
            style: 'border: 1px solid black; display: block;',
            toolbar: {
                items: [
                    {id: 'add', type: 'button', caption: 'Add Record', icon: 'w2ui-icon-plus'}
                ],
                onClick: function (event) {
                    var bid = getCurrentBID();
                    if (event.target == 'add') {
                        var inital = getFeesTermsGridInitalRecord(bid, w2ui.RAFeesTermsGrid.records.length);
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
                    field: 'RID',
                    hidden: true
                },
                {
                    field: 'BID',
                    hidden: true
                },
                {
                    field: 'RTID',
                    hidden: true
                },
                {
                    field: 'RentableName',
                    caption: 'Rentable',
                    size: '180px',
                    editable: {type: 'text'}
                },
                {
                    field: 'FeeName',
                    caption: 'Fee',
                    size: '120px',
                    editable: {type: 'text'}
                },
                {
                    field: 'Amount',
                    caption: 'Amount',
                    size: '80px',
                    render: 'money',
                    editable: {type: 'money'}
                },
                {
                    field: 'Cycle',
                    caption: 'Cycle',
                    size: '80px',
                    editable: {type: 'int'}
                },
                {
                    field: 'SigningAmt',
                    caption: 'At Signing',
                    size: '80px',
                    render: 'money',
                    editable: {type: 'money'}
                },
                {
                    field: 'ProrateAmt',
                    caption: 'Prorate',
                    size: '80px',
                    render: 'money',
                    editable: {type: 'money'}
                },
                {
                    field: 'TaxableAmt',
                    caption: 'Taxable Amt',
                    size: '80px',
                    render: 'money',
                    editable: {type: 'money'}
                },
                {
                    field: 'SalesTax',
                    caption: 'Sales Tax',
                    size: '80px',
                    render: 'money',
                    editable: {type: 'money'}
                },
                {
                    field: 'TransOCC',
                    caption: 'Trans OCC',
                    size: '100%',
                    render: 'money',
                    editable: {type: 'money'}
                }
            ],
            onChange: function (event) {
                event.onComplete = function () {
                    this.save();
                };
            }
        });
    }

    // load grid in division
    $('#ra-form #feesterms').w2render(w2ui.RAFeesTermsGrid);

    // load the existing data in feesterms component
    setTimeout(function () {
        var i = getRAFlowPartTypeIndex(app.raFlowPartTypes.feesterms);
        if (i >= 0 && app.raflow.data[app.raflow.activeFlowID][i].Data) {
            w2ui.RAFeesTermsGrid.records = app.raflow.data[app.raflow.activeFlowID][i].Data;
            w2ui.RAFeesTermsGrid.refresh();
        } else {
            w2ui.RAFeesTermsGrid.clear();
        }
    }, 500);
};

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
