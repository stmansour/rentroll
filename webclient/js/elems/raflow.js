/*global w2ui,
    getFullName, getTCIDName, loadTransactantListingItem,
    initRAFlowAJAX, getRAFlowAllParts, getAllRAFlows, loadRADatesForm, loadRAPeopleForm,
    loadRAPetsGrid, loadRAVehiclesGrid, loadRABGInfoForm, loadRARentablesGrid,
    loadRAFeesTermsGrid, getRAFlowPartTypeIndex, loadTargetSection, getFeesTermsGridInitalRecord,
    saveActiveCompData, loadRABGInfoForm, w2render, getVehicleFormInitalRecord,
    requiredFieldsFulFilled, getPetFormInitRecord, lockOnGrid, reassignGridRecids, getRAFlowPartData,
    openNewTransactantForm, getRAAddTransactantFormInitRec, toggleHaveCheckBoxDisablity, getRATransanctantDetail,
    setRABGInfoFormFields, getRABGInfoGridRecord, showHideRABGInfoFormFields, setNotRequiredFields,
    hideSliderContent, showSliderContentW2UIComp, getRABGInfoFormInitRecord, setRABGInfoFormHeader,
    updateRABGInfoFormCheckboxes, findTransactantIndexByTCIDInPeopleData, appendNewSlider
*/

"use strict";

// right side slider content div length
var sliderContentDivLength = 400;

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
$(document).on('click', '#ra-form #progressbar #steps-list a', function () {
    var active_comp = $(".ra-form-component:visible");

    // load target form
    var target = $(this).closest("li").attr("data-target");
    target = target.split('#').join("");

    loadTargetSection(target, active_comp.attr("id"));

    // because of 'a' tag, return false
    return false;
});

// lockOnGrid
// Lock grid if chebox is unchecked(false). Unlock grid if checkbox is checked(true).
// Lock grid when there is no record in the grid.
window.lockOnGrid = function (gridName) {
    var isChecked = $("#" + gridName + "_checkbox")[0].checked;
    var recordsLength = w2ui[gridName].records.length;

    if (!isChecked && recordsLength === 0){
        w2ui[gridName].lock();
    }else{
        w2ui[gridName].unlock();
    }

    if( recordsLength > 0 ){
        $("#" + gridName + "_checkbox")[0].disabled = true;
        $("#" + gridName + "_checkbox")[0].checked = true;
    }
};

// setRABGInfoFormHeader
// TODO(Akshay): Ask steve about form fields when transanctant is company.
window.setRABGInfoFormHeader = function(record) {
    if(record.IsCompany){
        w2ui.RABGInfoForm.header = 'Background Information - ' + record.CompanyName;
    }else{
        w2ui.RABGInfoForm.header = 'Background Information - ' + record.FirstName + ' ' + record.MiddleName + ' ' + record.LastName;
    }
};

// setRABGInfoFormFields It set default fields value from the transanctants to form.
window.setRABGInfoFormFields = function(record) {
    var formRecord = w2ui.RABGInfoForm.record; // record from the w2ui form

    formRecord.FirstName = record.FirstName;
    formRecord.MiddleName = record.MiddleName;
    formRecord.LastName = record.LastName;
    formRecord.TelephoneNo = record.CellPhone;
    formRecord.EmailAddress = record.PrimaryEmail;
    formRecord.Phone = record.WorkPhone;
    formRecord.Address = record.Address;
    formRecord.Address2 = record.Address2;
    formRecord.City = record.City;
    formRecord.Country = record.Country;
    formRecord.PostalCode = record.PostalCode;
    formRecord.State = record.State;
};

// showHideRABGInfoFormFields hide fields if transanctant is only user
window.showHideRABGInfoFormFields = function(listOfHiddenFields, hidden){
    if(hidden){
        $("#cureentInfolabel").hide();
        $("#priorInfolabel").hide();
    }else{
        $("#cureentInfolabel").show();
        $("#priorInfolabel").show();
    }
    for(var fieldIndex=0; fieldIndex < listOfHiddenFields.length; fieldIndex++){
        w2ui.RABGInfoForm.get(listOfHiddenFields[fieldIndex]).hidden = hidden;
    }
};

// setNotRequiredFields define fields are not required if transanctant is only user
window.setNotRequiredFields = function(listOfNotRequiredFields, required){
    for(var fieldIndex=0; fieldIndex < listOfNotRequiredFields.length; fieldIndex++){
        w2ui.RABGInfoForm.get(listOfNotRequiredFields[fieldIndex]).required = required;
    }
};

// toggleHaveCheckBoxDisablity
// Enable checkbox if there is no record
// lock/unlock grid based on checkbox value
window.toggleHaveCheckBoxDisablity = function (gridName) {
    var recordsLength = w2ui[gridName].records.length;
    if (recordsLength > 0){
        $("#" + gridName + "_checkbox")[0].disabled = true;
    }else if(recordsLength === 0){
        $("#" + gridName + "_checkbox")[0].disabled = false;
        lockOnGrid(gridName);
    }
};

// getRAFlowPartData
window.getRAFlowPartData = function (partType) {

    var bid = getCurrentBID();

    var flowPartID;
    var flowParts = app.raflow.data[app.raflow.activeFlowID] || [];

    for (var i = 0; i < flowParts.length; i++) {
        if (partType == flowParts[i].PartType) {
            flowPartID = flowParts[i].FlowPartID;
            break;
        }
    }

    // temporary data
    var data = {
        "cmd": "get",
        "FlowPartID": flowPartID,
        "Flow": app.raflow.name,
        "FlowID": app.raflow.activeFlowID,
        "BID": bid,
        "PartType": partType
    };


    return $.ajax({
        url: "/v1/flowpart/" + bid.toString() + "/" + flowPartID,
        method: "POST",
        contentType: "application/json",
        dataType: "json",
        data: JSON.stringify(data),
        success: function (data) {
            if (data.status != "error"){
                // app.raflow[app.raflow.activeFlowID]
                console.log("Received data for activeFlowID:", app.raflow.activeFlowID, ", partType:", partType);
            }else {
                console.error(data.message);
            }
        },
        error: function () {
            console.log("Error:" + JSON.stringify(data));
        }
    });

};

// get RATransactant detail from the server
window.getRATransanctantDetail = function(TCID){
    var bid = getCurrentBID();

    // temporary data
    var data = {
        "cmd":"get",
        "recid":0,
        "name":"transactantForm"
    };


    return $.ajax({
        url: "/v1/person/" + bid.toString() + "/" + TCID,
        method: "POST",
        contentType: "application/json",
        dataType: "json",
        data: JSON.stringify(data),
        success: function (data) {
            if (data.status != "error"){
                // console.log("Received data for transanctant:", JSON.stringify(data));
            }else {
                console.error(data.message);
            }
        },
        error: function () {
            console.log("Error:" + JSON.stringify(data));
        }
    });
};

// getRABGInfoGridRecord
window.getRABGInfoGridRecord = function(records, TCID){
    var raBGInfoGridrecord;
    for(var recordIndex=0; recordIndex < records.length; recordIndex++) {
        if(records[recordIndex].TCID == TCID){
            raBGInfoGridrecord = records[recordIndex];
            break;
        }
    }
    return raBGInfoGridrecord;
};

// updateRABGInfoFormCheckboxes
window.updateRABGInfoFormCheckboxes = function(record){
    record.Evicted = int_to_bool(record.Evicted);
    record.Bankruptcy = int_to_bool(record.Bankruptcy);
    record.Convicted = int_to_bool(record.Convicted);
};

// TODO: we should pass FlowID, flowPartID here in arguments
window.saveActiveCompData = function (record, partType) {

    var bid = getCurrentBID();

    var flowPartID;
    var flowParts = app.raflow.data[app.raflow.activeFlowID] || [];

    for (var i = 0; i < flowParts.length; i++) {
        if (partType == flowParts[i].PartType) {
            flowPartID = flowParts[i].FlowPartID;
            break;
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

    return $.ajax({
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
                        $("#progressbar #steps-list li[data-target='#" + comp + "']").addClass("done");
                    }

                    // reset w2ui component as well
                    if(RACompConfig[comp].w2uiComp in w2ui) {
                        // clear inputs
                        w2ui[RACompConfig[comp].w2uiComp].clear();
                    }
                }

                // mark first slide as active
                $(".ra-form-component#dates").show();
                $("#progressbar #steps-list li[data-target='#dates']").removeClass("done").addClass("active");
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
    var validData = true;
    var isChecked;

    switch (compID) {
        case "dates":
            data = app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data;
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
            app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data.forEach(function(item) {
                if (item.IsRenter) {
                    done = true;
                    return false;
                }
            });
            break;
        case "pets":
            data = app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data;
            isChecked = $('#RAPetsGrid_checkbox')[0].checked;
            if(!isChecked){
                done = true;
            } else {
                if (data.length > 0) {
                    done = true;
                }else{
                    done = false;
                }
            }
            break;
        case "vehicles":
            data = app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data;
            isChecked = $('#RAVehiclesGrid_checkbox')[0].checked;
            if(!isChecked){
                done = true;
            }else{
                if (data.length > 0) {
                    done = true;
                }else{
                    done = false;
                }
            }
            break;
        case "bginfo":
            // TODO(Akshay): Add for integer fields e.g., phone, gross wage.
            data = app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data;
            // list of fields which must have value and it's type string
            var listOfRequiredField = ["ApplicationDate", "MoveInDate",
                "ApartmentNo", "LeaseTerm", "FirstName", "MiddleName",
                "LastName", "BirthDate", "SSN",
                "DriverLicNo", "TelephoneNo", "EmailAddress",
                "NoPeople", "CurrentAddress", "CurrentLandLordName", "CurrentLandLordPhoneNo",
                "CurrentReasonForMoving", "Employer", "Phone", "Address",
                "Position", "EmergencyContactName", "EmergencyContactPhone", "EmergencyContactAddress"];

            listOfRequiredField.forEach(function(field) {
                if (!data[field]) {
                    validData = false;
                    return false;
                }
            });

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

    /*if ($("#progressbar #steps-list li[data-target='#" + target + "']").hasClass("done")) {
        console.log("target has been saved", target);
    } else {}*/

    // if required fields are fulfilled then mark this slide as done
    if (requiredFieldsFulFilled(activeCompID)) {
        // hide active component
        $("#progressbar #steps-list li[data-target='#" + activeCompID + "']").addClass("done");
    }

    // decide data based on type
    var data = null;
    var partTypeIndex;
    switch (activeCompID) {
        case "dates":
            data = w2ui.RADatesForm.record;
            break;
        case "people":
            partTypeIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.people);
            data = app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data;
            w2ui.RAPeopleForm.actions.reset();
            break;
        case "pets":
            data = w2ui.RAPetsGrid.records;
            break;
        case "vehicles":
            data = w2ui.RAVehiclesGrid.records;
            break;
        case "bginfo":
            partTypeIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.bginfo);
            data = app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data;
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
    var partType = $("#progressbar #steps-list li[data-target='#" + activeCompID + "']").index() + 1;
    if (data) {
        // save the content on server for active component
        saveActiveCompData(data, partType);
    }

    // hide active component
    $("#progressbar #steps-list li[data-target='#" + activeCompID + "']").removeClass("active");
    $(".ra-form-component#" + activeCompID).hide();

    // show target component
    $("#progressbar #steps-list li[data-target='#" + target + "']").removeClass("done").addClass("active");
    $(".ra-form-component#" + target).show();

    // hide previous navigation button if the target is in first section
    if ($(".ra-form-component#" + target).is($(".ra-form-component").first())) {
        $("#ra-form footer button#previous").prop("disabled", true);
    } else {
        $("#ra-form footer button#previous").prop("disabled", false);
    }

    // hide next navigation button if the target is in last section
    if ($(".ra-form-component#" + target).is($(".ra-form-component").last())) {
        $("#ra-form footer button#next").prop("disabled", true);
    } else {
        $("#ra-form footer button#next").prop("disabled", false);
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
                    // $("#progressbar #steps-list li[data-target='#" + activeCompID + "']").addClass("done");
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

    var partType = app.raFlowPartTypes.dates;

    var partTypeIndex = getRAFlowPartTypeIndex(partType);

    if (partTypeIndex < 0){
        return;
    }

    // Fetch data from the server if there is any record available.
    getRAFlowPartData(partType)
        .done(function(data){
            if(data.status === 'success'){
                app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data = data.record.Data;
            }else {
                console.log(data.message);
            }
        })
        .fail(function(data){
            console.log("failure" + data);
        });

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
//                    IsCompany, CompanyName, IsGuarantor, IsOccupant, IsRenter
// @return - nothing
//-----------------------------------------------------------------------------
window.loadTransactantListingItem = function (transactantRec) {

    var peoplePartIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.people);
    if (peoplePartIndex < 0) {
        alert("flow data could not be found for people");
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

    // add into renter list
    if (transactantRec.IsRenter) {
        // if with this tcid element exists in DOM then not append
        if ($('#renter-list .people-listing li[data-tcid="' + transactantRec.TCID + '"]').length <= 0) {
            $('#renter-list .people-listing').append(peopleListingItem);
        }
    } else {
        $('#renter-list .people-listing li[data-tcid="' + transactantRec.TCID + '"]').remove();
    }

    // add into occupant list
    if (transactantRec.IsOccupant) {
        // if with this tcid element exists in DOM then not append
        if ($('#occupant-list .people-listing li[data-tcid="' + transactantRec.TCID + '"]').length <= 0) {
            $('#occupant-list .people-listing').append(peopleListingItem);
        }
    } else {
        $('#occupant-list .people-listing li[data-tcid="' + transactantRec.TCID + '"]').remove();
    }

    // add into guarantor list
    if (transactantRec.IsGuarantor) {
        // if with this tcid element exists in DOM then not append
        if ($('#guarantor-list .people-listing li[data-tcid="' + transactantRec.TCID + '"]').length <= 0) {
            $('#guarantor-list .people-listing').append(peopleListingItem);
        }
    } else {
        $('#guarantor-list .people-listing li[data-tcid="' + transactantRec.TCID + '"]').remove();
    }
};

//-----------------------------------------------------------------------------
// openNewTransactantForm - popup new transactant form
//-----------------------------------------------------------------------------
window.openNewTransactantForm = function() {
    var BID = getCurrentBID(),
        BUD = getBUDfromBID(BID);

    // this is new form so TCID is set to zero
    w2ui.RAAddTransactantForm.url = "/v1/person/" + BID.toString() + "/0";
    w2ui.RAAddTransactantForm.record = getRAAddTransactantFormInitRec(BID, BUD, null);
    showSliderContentW2UIComp(w2ui.RAAddTransactantForm, RACompConfig.people.sliderWidth);
    w2ui.RAAddTransactantForm.refresh(); // need to refresh for header changes
};

//-----------------------------------------------------------------------------
// getRAAddTransactantFormInitRec - returns object with default values for
//                                  fields
//-----------------------------------------------------------------------------
window.getRAAddTransactantFormInitRec = function(BID, BUD, previousFormRecord) {

    var defaultFormData = {
        TCID: 0,
        BID: BID,
        BUD: BUD,
        FirstName: "",
        MiddleName: "",
        LastName: "",
        IsCompany: false,
        CompanyName: "",
        EligibleFutureUser: "yes",
        EligibleFuturePayor: "yes",
    };

    // if it called after 'save and add another' action there previous form record is passed as Object
    // else it is null
    if ( previousFormRecord ) {
        defaultFormData = setDefaultFormFieldAsPreviousRecord(
            [ 'FirstName', 'MiddleName', 'LastName', 'IsCompany', 'CompanyName' ], // Fields to Reset
            defaultFormData,
            previousFormRecord
        );
    }

    return defaultFormData;
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
    var IsRenter = w2ui.RAPeopleForm.record.IsRenter;
    var IsOccupant = w2ui.RAPeopleForm.record.IsOccupant;
    var IsGuarantor = w2ui.RAPeopleForm.record.IsGuarantor;

    // if not set anything then alert the user to select any one of them
    if (!(IsRenter || IsOccupant || IsGuarantor)) {
        alert("Please, select the role");
        return false;
    }

    // get part type index
    var peoplePartIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.people);
    // remove entry from data
    if (peoplePartIndex < 0) {
        return false;
    }

    var transactantRec = $.extend(true, {}, w2ui.RAPeopleForm.record);
    delete transactantRec.Transactant;
    var tcidIndex = findTransactantIndexByTCIDInPeopleData(transactantRec.TCID);

    // if not found then push it in the data
    if (tcidIndex < 0) {
        app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data.push(transactantRec);
    } else {
        app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data[tcidIndex] = transactantRec;
    }

    // load item in the DOM
    loadTransactantListingItem(transactantRec);

    // clear the form
    w2ui.RAPeopleForm.actions.reset();

    // disable check boxes
    $(w2ui.RAPeopleForm.box).find("input[type=checkbox]").prop("disabled", true);
};

//-----------------------------------------------------------------------------
// findTransactantIndexByTCIDInPeopleData - finds the index of transactant data
//                in local people data of raflow by TCID
//
// @params
//   TCID = tcid
//-----------------------------------------------------------------------------
window.findTransactantIndexByTCIDInPeopleData = function(TCID) {
    var index = -1;

    // get part type index
    var peoplePartIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.people);
    // remove entry from data
    if (peoplePartIndex < 0) {
        return;
    }

    if (typeof app.raflow.data[app.raflow.activeFlowID] !== "undefined") {
        app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data.forEach(function(transactantRec, i) {
            if (transactantRec.TCID === TCID) {
                index = i;
                return false;
            }
        });
    }

    return index;
};

// click on any transactant item in people listing
$(document).on("click", ".people-listing li", function(e) {
    if($(e.target).hasClass("remove-item")) {
        return;
    }
    var TCID = parseInt($(this).attr("data-tcid"));
    var tcidIndex = findTransactantIndexByTCIDInPeopleData(TCID);
    if (tcidIndex < 0) {
        return;
    }

    // get part type index
    var peoplePartIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.people);
    // remove entry from data
    if (peoplePartIndex < 0) {
        return;
    }

    var transactantRec = app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data[tcidIndex];

    var item = {
        CompanyName: transactantRec.CompanyName,
        IsCompany: transactantRec.IsCompany,
        FirstName: transactantRec.FirstName,
        LastName: transactantRec.LastName,
        MiddleName: transactantRec.MiddleName,
        TCID: transactantRec.TCID,
        recid: 0,
    };
    w2ui.RAPeopleForm.record = transactantRec;
    w2ui.RAPeopleForm.record.Transactant = item;
    w2ui.RAPeopleForm.refresh();
    if ($(w2ui.RAPeopleForm.box).find("input[name=Transactant]").length > 0) {
        $(w2ui.RAPeopleForm.box).find("input[name=Transactant]").data('selected', [item]).data('w2field').refresh();
    }
});

// remove people from the listing
$(document).on('click', '.people-listing .remove-item', function () {
    // get part type index
    var peoplePartIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.people);
    // remove entry from data
    if (peoplePartIndex < 0) {
        return;
    }

    var TCID = parseInt($(this).closest('li').attr('data-tcid'));
    var peopleType = $(this).closest('ul.people-listing').attr('data-people-type');
    var tcidIndex = findTransactantIndexByTCIDInPeopleData(TCID);
    if (tcidIndex < 0) {
        return;
    }

    // taka the reference
    var transactant = app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data[tcidIndex];
    switch (peopleType) {
        case "renters":
            transactant.IsRenter = false;
            break;
        case "occupants":
            transactant.IsOccupant = false;
            break;
        case "guarantors":
            transactant.IsGuarantor = false;
            break;
    }

    // if selected tcid record to be removed, of which all flags are false then remove it from data also
    if (!(transactant.IsGuarantor || transactant.IsRenter || transactant.IsOccupant)) {
        app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data.splice(tcidIndex, 1);

        //
        if (app.raflow.activeTransactant.TCID === transactant.TCID) {
            w2ui.RAPeopleForm.actions.reset(); // clear the form
        }
    }

    // remove item from the DOM
    $(this).closest('li').remove();

    // refresh the form
    if (w2ui.RAPeopleForm.record.TCID > 0) {
        var item = {
            CompanyName: transactant.CompanyName,
            IsCompany: transactant.IsCompany,
            FirstName: transactant.FirstName,
            LastName: transactant.LastName,
            MiddleName: transactant.MiddleName,
            TCID: transactant.TCID,
            recid: 0,
        };
        w2ui.RAPeopleForm.record = $.extend(true, transactant, w2ui.RAPeopleForm.record);
        w2ui.RAPeopleForm.record.Transactant = item;
        w2ui.RAPeopleForm.refresh();
        if ($(w2ui.RAPeopleForm.box).find("input[name=Transactant]").length > 0) {
            $(w2ui.RAPeopleForm.box).find("input[name=Transactant]").data('selected', [item]).data('w2field').refresh();
        }
    }
});

window.loadRAPeopleForm = function () {
    var partType = app.raFlowPartTypes.people;
    var partTypeIndex = getRAFlowPartTypeIndex(partType);
    if (partTypeIndex < 0){
        return;
    }

    // Fetch data from the server if there is any record available.
    getRAFlowPartData(partType)
        .done(function(data){
            if(data.status === 'success'){
                app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data = data.record.Data || [];
                app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data.forEach(function (item) {
                    console.log(item);
                    loadTransactantListingItem(item);
                });
            }else {
                console.log(data.message);
            }
        })
        .fail(function(data){
            console.log("failure" + data);
        });

    // if form is loaded then return
    if (!("RAPeopleForm" in w2ui)) {

        // people form
        $().w2form({
            name: 'RAPeopleForm',
            header: 'People',
            style: 'display: block; border: none;',
            formURL: '/webclient/html/formrapeople.html',
            focus: -1,
            fields: [
                {
                    name: 'Transactant', type: 'enum', required: true, html: {caption: "Transactant"},
                    options: {
                        url: '/v1/transactantstd/' + app.raflow.BID,
                        max: 1,
                        renderItem: function (item) {
                            // enable user-role checkboxes
                            $(w2ui.RAPeopleForm.box).find("input[type=checkbox]").prop("disabled", false);
                            // mark this as transactant as an active
                            app.raflow.activeTransactant = item;
                            var s = getTCIDName(item);
                            w2ui.RAPeopleForm.record.TCID = item.TCID;
                            w2ui.RAPeopleForm.record.FirstName = item.FirstName;
                            w2ui.RAPeopleForm.record.LastName = item.LastName;
                            w2ui.RAPeopleForm.record.MiddleName = item.MiddleName;
                            w2ui.RAPeopleForm.record.CompanyName = item.CompanyName;
                            w2ui.RAPeopleForm.record.IsCompany = item.IsCompany;

                            /*// if this transactant is available in local data then try to render the bools
                            var tcidIndex = findTransactantIndexByTCIDInPeopleData(item.TCID);
                            var transactantRec;
                            if (tcidIndex >= 0) {
                                var peoplePartIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.people);
                                if (peoplePartIndex < 0){
                                    return;
                                }

                                // taka the reference
                                transactantRec = app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data[tcidIndex];

                                w2ui.RAPeopleForm.record.IsRenter = transactantRec.IsRenter;
                                w2ui.RAPeopleForm.record.IsOccupant = transactantRec.IsOccupant;
                                w2ui.RAPeopleForm.record.IsGuarantor = transactantRec.IsGuarantor;
                                // w2ui.RAPeopleForm.refresh();
                            }*/

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
                                var f = w2ui.RAPeopleForm;
                                // reset payor field related data when removed
                                f.actions.reset();
                                // disable user-role checkboxes
                                $(f.box).find("input[type=checkbox]").prop("disabled", true);

                                // NOTE: have to trigger manually, b'coz we manually change the record,
                                // otherwise it triggers the change event but it won't get change (Object: {})
                                var event = f.trigger({phase: 'before', target: f.name, type: 'change', event: event}); // event before
                                if (event.cancelled === true) return false;
                                f.trigger($.extend(event, {phase: 'after'})); // event after
                            };
                        }
                    }
                },
                {name: 'BID', type: 'int', required: true, html: {caption: "BID"}},
                {name: 'TCID', type: 'int', required: true, html: {caption: "TCID"}},
                {name: 'FirstName', type: 'text', required: true, html: {caption: "FirstName"}},
                {name: 'LastName', type: 'text', required: true, html: {caption: "LastName"}},
                {name: 'MiddleName', type: 'text', required: true, html: {caption: "MiddleName"}},
                {name: 'CompanyName', type: 'text', required: true, html: {caption: "CompanyName"}},
                {name: 'IsCompany', type: 'int', required: true, html: {caption: "IsCompany"}},
                {name: 'IsRenter', type: 'checkbox', required: false, html: {caption: "Renter"}},
                {name: 'IsOccupant', type: 'checkbox', required: false, html: {caption: "Occupant"}},
                {name: 'IsGuarantor', type: 'checkbox', required: false, html: {caption: "Guarantor"}}
            ],
            actions: {
                reset: function () {
                    w2ui.RAPeopleForm.clear();
                    // reset active Transactant to blank object
                    app.raflow.activeTransactant = {};
                }
            },
            onRefresh: function(event) {
                var f = this;
                event.onComplete = function() {
                    var BID = getCurrentBID(),
                        BUD = getBUDfromBID(BID);

                    f.record.BID = BID;
                };
            }
        });

        // new transactant form especially for this RA flow
        $().w2form({
            name: 'RAAddTransactantForm',
            header: 'Add New Transactant',
            style: 'display: block;',
            formURL: '/webclient/html/formra-addtransactant.html',
            focus: -1,
            fields: [
                {name: 'BID', type: 'int', required: true, html: {page: 0, column: 0}},
                {name: 'BUD', type: 'list', required: true, options: {items: app.businesses}, html: {page: 0, column: 0}},
                {name: 'TCID', type: 'int', hidden: false, html: { caption: 'TCID', page: 0, column: 0 } },
                {name: 'EligibleFuturePayor', type: 'text', required: true, html: {caption: "EligibleFuturePayor"}},
                {name: 'EligibleFutureUser', type: 'text', required: true, html: {caption: "EligibleFutureUser"}},
                {name: 'FirstName', type: 'text', required: true, html: {caption: "FirstName"}},
                {name: 'LastName', type: 'text', required: true, html: {caption: "LastName"}},
                {name: 'MiddleName', type: 'text', required: true, html: {caption: "MiddleName"}},
                {name: 'CompanyName', type: 'text', required: false, html: {caption: "CompanyName"}},
                {name: 'IsCompany', type: 'bool', required: false, html: {caption: "IsCompany"}}
            ],
            toolbar : {
                items: [
                    { id: 'bt3', type: 'spacer' },
                    { id: 'btnClose', type: 'button', icon: 'fas fa-times'}
                ],
                onClick: function (event) {
                    switch (event.target){
                        case 'btnClose':
                            hideSliderContent();
                            break;
                    }
                }
            },
            actions: {
                reset: function () {
                    this.clear();
                },
                save: function() {
                    var f = this;
                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    f.save({}, function(data) {
                        if (data.status === 'error') {
                            f.message(data.message);
                            return;
                        }
                        hideSliderContent();
                    });
                },
                saveadd: function() {
                    var BID = getCurrentBID(),
                        BUD = getBUDfromBID(BID),
                        f = this;

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    f.save({}, function(data) {
                        if (data.status === 'error') {
                            f.message(data.message);
                            return;
                        }

                        f.record = getRAAddTransactantFormInitRec(BID, BUD, f.record);
                        f.refresh();
                    });
                },
            },
            onSubmit: function(target, data) {
                if (data.postData.record.IsCompany) {
                    data.postData.record.IsCompany = 1;
                } else {
                    data.postData.record.IsCompany = 0;
                }
            },
            onChange: function(event) {
                event.onComplete = function() {
                    if (this.record.IsCompany) {
                        this.get("FirstName").required = false;
                        this.get("MiddleName").required = false;
                        this.get("LastName").required = false;
                        this.get("CompanyName").required = true;
                        this.get("IsCompany").required = true;
                    } else {
                        this.get("FirstName").required = true;
                        this.get("MiddleName").required = true;
                        this.get("LastName").required = true;
                        this.get("CompanyName").required = false;
                        this.get("IsCompany").required = false;
                    }
                    this.refresh();

                    // formRecDiffer: 1=current record, 2=original record, 3=diff object
                    var diff = formRecDiffer(this.record, app.active_form_original, {});
                    // if diff == {} then make dirty flag as false, else true
                    if ($.isPlainObject(diff) && $.isEmptyObject(diff)) {
                        app.form_is_dirty = false;
                    } else {
                        app.form_is_dirty = true;
                    }
                };
            },
            onRefresh: function(event) {
                var f = this;
                event.onComplete = function() {
                    var BID = getCurrentBID(),
                        BUD = getBUDfromBID(BID);

                    f.record.BID = BID;
                    f.record.BUD = BUD;

                    // there is NO PETID actually, so have to work around with recid key
                    formRefreshCallBack(f, "TCID");
                };
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
            w2ui.RAPeopleForm.actions.reset();
        }
    }, 500);
};

window.getPetFormInitRecord = function (BID, BUD, previousFormRecord){
    var t = new Date(),
        nyd = new Date(new Date().setFullYear(new Date().getFullYear() + 1));

    var defaultFormData = {
        recid: 0,
        PETID: 0,
        BID: BID,
        // BUD: BUD,
        Name: "",
        Breed: "",
        Type: "",
        Color: "",
        Weight: 0,
        DtStart: w2uiDateControlString(t),
        DtStop: w2uiDateControlString(nyd),
        NonRefundablePetFee: 0,
        RefundablePetDeposit: 0,
        RecurringPetFee: 0,
        LastModTime: t.toISOString(),
        LastModBy: 0,
    };

    // if it called after 'save and add another' action there previous form record is passed as Object
    // else it is null
    if ( previousFormRecord ) {
        defaultFormData = setDefaultFormFieldAsPreviousRecord(
            [ 'Name', 'Breed', 'Type', 'Color', 'Weight',
              'NonRefundablePetFee', 'RefundablePetDeposit', 'ReccurringPetFee' ], // Fields to Reset
            defaultFormData,
            previousFormRecord
        );
    }

    return defaultFormData;
};

window.loadRAPetsGrid = function () {

    var partType = app.raFlowPartTypes.pets;
    var partTypeIndex = getRAFlowPartTypeIndex(partType);

    if (partTypeIndex < 0){
        return;
    }

    // Fetch data from the server if there is any record available.
    getRAFlowPartData(partType)
        .done(function(data){
            if(data.status === 'success'){
                app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data = data.record.Data;
            }else {
                console.log(data.message);
            }
        })
        .fail(function(data){
            console.log("failure" + data);
        });

    // if form is loaded then return
    if (!("RAPetsGrid" in w2ui)) {

        // pet form
        $().w2form({
            name    : 'RAPetForm',
            header  : 'Add Pet information',
            style   : 'border: 0px; background-color: transparent; display: block;',
            formURL : '/webclient/html/formrapets.html',
            toolbar : {
                items: [
                    { id: 'bt3', type: 'spacer' },
                    { id: 'btnClose', type: 'button', icon: 'fas fa-times'}
                ],
                onClick: function (event) {
                    switch (event.target){
                        case 'btnClose':
                            hideSliderContent();
                            break;
                    }
                }
            },
            fields  : [
                { field: 'recid', type: 'int', required: false, html: { caption: 'recid', page: 0, column: 0 } },
                { field: 'BID', type: 'int', hidden: true, html: { caption: 'BID', page: 0, column: 0 } },
                // { field: 'BUD', type: 'text', hidden: false, html: { caption: 'BUD', page: 0, column: 0 } },
                { field: 'PETID', type: 'int', hidden: false, html: { caption: 'PETID', page: 0, column: 0 } },
                { field: 'Name', type: 'text', required: true},
                { field: 'Breed', type: 'text', required: true},
                { field: 'Type', type: 'text', required: true},
                { field: 'Color', type: 'text', required: true},
                { field: 'Weight', type: 'int', required: true},
                { field: 'NonRefundablePetFee', type: 'money', required: false},
                { field: 'RefundablePetDeposit', type: 'money', required: false},
                { field: 'RecurringPetFee', type: 'money', required: false},
                { field: 'DtStart', type: 'date', required: false, html: { caption: 'DtStart', page: 0, column: 0 } },
                { field: 'DtStop', type: 'date', required: false, html: { caption: 'DtStop', page: 0, column: 0 } },
                { field: 'LastModTime', type: 'time', required: false, html: { caption: 'LastModTime', page: 0, column: 0 } },
                { field: 'LastModBy', type: 'int', required: false, html: { caption: 'LastModBy', page: 0, column: 0 } },
            ],
            onRefresh: function(event) {
                event.onComplete = function() {
                    var f = w2ui.RAPetForm,
                        header = "Edit Rental Agreement Pets ({0})";

                    // there is NO PETID actually, so have to work around with recid key
                    formRefreshCallBack(f, "recid", header);

                    // hide delete button if it is NewRecord
                    var isNewRecord = (w2ui.RAPetsGrid.get(f.record.recid, true) === null);
                    if (isNewRecord) {
                        $(f.box).find("button[name=delete]").addClass("hidden");
                    } else {
                        $(f.box).find("button[name=delete]").removeClass("hidden");
                    }
                };
            },
            onChange: function(event) {
                event.onComplete = function() {
                    // formRecDiffer: 1=current record, 2=original record, 3=diff object
                    var diff = formRecDiffer(this.record, app.active_form_original, {});
                    // if diff == {} then make dirty flag as false, else true
                    if ($.isPlainObject(diff) && $.isEmptyObject(diff)) {
                        app.form_is_dirty = false;
                    } else {
                        app.form_is_dirty = true;
                    }
                };
            },
            actions: {
                save: function() {
                    var form = this;
                    var grid = w2ui.RAPetsGrid;
                    var errors = form.validate();
                    if (errors.length > 0) return;
                    var record = $.extend(true, { recid: grid.records.length + 1 }, form.record);
                    var recordsData = $.extend(true, [], grid.records);
                    var isNewRecord = (grid.get(record.recid, true) === null);

                    // if it doesn't exist then only push
                    if (isNewRecord) {
                        recordsData.push(record);
                    }

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // save this records in json Data
                    saveActiveCompData(recordsData, app.raFlowPartTypes.pets)
                    .done(function(data) {
                        if (data.status === 'success') {
                            // if null
                            if (isNewRecord) {
                                grid.add(record);
                            } else {
                                grid.set(record.recid, record);
                            }
                            form.clear();

                            // Disable "have pets?" checkbox if there is any record.
                            toggleHaveCheckBoxDisablity('RAPetsGrid');

                            // close the form
                            hideSliderContent();
                        } else {
                            form.message(data.message);
                        }
                    })
                    .fail(function(data) {
                        console.log("failure " + data);
                    });
                },
                saveadd: function() {
                    var BID = getCurrentBID(),
                        BUD = getBUDfromBID(BID);

                    var form = this;
                    var grid = w2ui.RAPetsGrid;
                    var errors = form.validate();
                    if (errors.length > 0) return;
                    var record = $.extend(true, {}, form.record);
                    var recordsData = $.extend(true, [], grid.records);
                    var isNewRecord = (grid.get(record.recid, true) === null);

                    // if it doesn't exist then only push
                    if (isNewRecord) {
                        recordsData.push(record);
                    }

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // save this records in json Data
                    saveActiveCompData(recordsData, app.raFlowPartTypes.pets)
                    .done(function(data) {
                        if (data.status === 'success') {
                            // clear the grid select recid
                            app.last.grid_sel_recid  =-1;
                            // selectNone
                            grid.selectNone();

                            // if null
                            if (isNewRecord) {
                                // add this record to grid
                                grid.add(record);
                            } else {
                                grid.set(record.recid, record);
                            }
                            // add new formatted record to current form
                            form.record = getPetFormInitRecord(BID, BUD, form.record);
                            // set record id
                            form.record.recid = grid.records.length + 1;
                            form.refresh();
                            form.refresh();
                        } else {
                            form.message(data.message);
                        }
                    })
                    .fail(function(data) {
                        console.log("failure " + data);
                    });
                },
                delete: function() {
                    var form = this;
                    var grid = w2ui.RAPetsGrid;

                    // backup the records
                    var records = $.extend(true, [], grid.records);
                    for (var i = 0; i < records.length; i++) {
                        if(records[i].recid == form.record.recid) {
                            records.splice(i, 1);
                        }
                    }

                    // save this records in json Data
                    saveActiveCompData(records, app.raFlowPartTypes.pets)
                    .done(function(data) {
                        if (data.status === 'success') {
                            // clear the grid select recid
                            app.last.grid_sel_recid  =-1;
                            // selectNone
                            grid.selectNone();

                            grid.remove(form.record.recid);
                            form.clear();

                            // Disable "have pets?" checkbox if there is any record.
                            toggleHaveCheckBoxDisablity('RAPetsGrid');

                            // need to refresh the grid as it will re-assign new recid
                            reassignGridRecids(grid.name);

                            // close the form
                            hideSliderContent();
                        } else {
                            form.message(data.message);
                        }
                    })
                    .fail(function(data) {
                        console.log("failure " + data);
                    });
                },
            },
        });

        // pets grid
        $().w2grid({
            name: 'RAPetsGrid',
            header: 'Pets',
            show: {
                toolbar: true,
                toolbarSearch: false,
                toolbarAdd: true,
                toolbarReload: true,
                toolbarInput: false,
                toolbarColumns: false,
                footer: true,
            },
            multiSelect: false,
            style: 'border: 0px solid black; display: block;',
            columns: [
                {
                    field: 'recid',
                    hidden: true
                },
                {
                    field: 'PETID',
                    hidden: true
                },
                {
                    field: 'BID',
                    hidden: true
                },
                /*{
                    field: 'BUD',
                    hidden: true
                },*/
                {
                    field: 'Name',
                    caption: 'Name',
                    size: '150px'
                },
                {
                    field: 'Type',
                    caption: 'Type',
                    size: '80px'
                },
                {
                    field: 'Breed',
                    caption: 'Breed',
                    size: '80px'
                },
                {
                    field: 'Color',
                    caption: 'Color',
                    size: '80px'
                },
                {
                    field: 'Weight',
                    caption: 'Weight',
                    size: '80px'
                },
                {
                    field: 'DtStart',
                    caption: 'DtStart',
                    size: '100px'
                },
                {
                    field: 'DtStop',
                    caption: 'DtStop',
                    size: '100px'
                },
                {
                    field: 'NonRefundablePetFee',
                    caption: 'NonRefundable<br>PetFee',
                    size: '70px',
                    render: 'money'
                },
                {
                    field: 'RefundablePetDeposit',
                    caption: 'Refundable<br>PetDeposit',
                    size: '70px',
                    render: 'money'
                },
                {
                    field: 'RecurringPetFee',
                    caption: 'Recurring<br>PetFee',
                    size: '70px',
                    render: 'money'
                }
            ],
            onChange: function (event) {
                event.onComplete = function () {
                    this.save();
                };
            },
            onClick: function(event) {
                event.onComplete = function() {
                    var yes_args = [this, event.recid],
                        no_args = [this],
                        no_callBack = function(grid) {
                            grid.select(app.last.grid_sel_recid);
                            return false;
                        },
                        yes_callBack = function(grid, recid) {
                            app.last.grid_sel_recid = parseInt(recid);

                            // keep highlighting current row in any case
                            grid.select(app.last.grid_sel_recid);
                            w2ui.RAPetForm.record = $.extend(true, {}, grid.get(app.last.grid_sel_recid));

                            showSliderContentW2UIComp(w2ui.RAPetForm, RACompConfig.pets.sliderWidth);
                            w2ui.RAPetForm.refresh(); // need to refresh for header changes
                        };

                    // warn user if form content has been changed
                    form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
                };
            },
            onAdd: function (/*event*/) {
                var yes_args = [this],
                    no_callBack = function() {
                        return false;
                    },
                    yes_callBack = function(grid) {
                        app.last.grid_sel_recid = -1;
                        grid.selectNone();

                        var BID = getCurrentBID(),
                            BUD = getBUDfromBID(BID);

                        w2ui.RAPetForm.record = getPetFormInitRecord(BID, BUD, null);
                        // set record id
                        w2ui.RAPetForm.record.recid = w2ui.RAPetsGrid.records.length + 1;
                        showSliderContentW2UIComp(w2ui.RAPetForm, RACompConfig.pets.sliderWidth);
                        w2ui.RAPetForm.refresh();
                    };

                // warn user if form content has been changed
                form_dirty_alert(yes_callBack, no_callBack, yes_args);
            }
        });
    }

    // now load grid in division
    $('#ra-form #pets .form-container').w2render(w2ui.RAPetsGrid);

    // load the existing data in pets component
    setTimeout(function () {
        var grid = w2ui.RAPetsGrid;
        var i = getRAFlowPartTypeIndex(app.raFlowPartTypes.pets);
        if (i >= 0 && app.raflow.data[app.raflow.activeFlowID][i].Data) {
            grid.records = app.raflow.data[app.raflow.activeFlowID][i].Data;
            reassignGridRecids(grid.name);

            // lock the grid until "Have pets?" checkbox checked.
            lockOnGrid(grid.name);

        } else {
            grid.clear();
        }
    }, 500);

};

// -------------------------------------------------------------------------------
// Rental Agreement - Vehicles Grid
// -------------------------------------------------------------------------------
window.getVehicleFormInitalRecord = function (BID, BUD, previousFormRecord) {
    var t = new Date(),
        nyd = new Date(new Date().setFullYear(new Date().getFullYear() + 1));

    var defaultFormData = {
        recid: 0,
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
        ParkingPermitFee: 0,
        DtStart: w2uiDateControlString(t),
        DtStop: w2uiDateControlString(nyd)
    };

    // if it called after 'save and add another' action there previous form record is passed as Object
    // else it is null
    if ( previousFormRecord ) {
        defaultFormData = setDefaultFormFieldAsPreviousRecord(
            [ 'Type', 'Make', 'Model', 'Color', 'Year', 'LicensePlateState', 'LicensePlateNumber', 'VIN',
                'ParkingPermitNumber', 'ParkingPermitFee'], // Fields to Reset
            defaultFormData,
            previousFormRecord
        );
    }

    return defaultFormData;
};

window.loadRAVehiclesGrid = function () {

    var partType = app.raFlowPartTypes.vehicles;
    var partTypeIndex = getRAFlowPartTypeIndex(partType);

    if (partTypeIndex < 0){
        return;
    }

    // Fetch data from the server if there is any record available.
    getRAFlowPartData(partType)
        .done(function(data){
            if(data.status === 'success'){
                app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data = data.record.Data;
            }else {
                console.log(data.message);
            }
        })
        .fail(function(data){
            console.log("failure" + data);
        });

    // if form is loaded then return
    if (!("RAVehiclesGrid" in w2ui)) {

        // Add vehicle information form
        $().w2form({
            name    : 'RAVehicleForm',
            header  : 'Add Vehicle form',
            formURL : '/webclient/html/formravehicles.html',
            toolbar :{
                items: [
                    { id: 'bt3', type: 'spacer' },
                    { id: 'btnClose', type: 'button', icon: 'fas fa-times'}
                ],
                onClick: function (event) {
                    switch (event.target){
                        case 'btnClose':
                            hideSliderContent();
                            break;
                    }
                }
            },
            fields  : [
                { field: 'recid', type: 'int', required: false, html: { caption: 'recid', page: 0, column: 0 } },
                { field: 'Type', type: 'text', required: true},
                { field: 'Make', type: 'text', required: true},
                { field: 'Model', type: 'text', required: true},
                { field: 'Color', type: 'text', required: true},
                { field: 'Year', type: 'text', required: true},
                { field: 'LicensePlateState', type: 'text', required: true},
                { field: 'LicensePlateNumber', type: 'text', required: true},
                { field: 'VIN', type: 'text', required: true},
                { field: 'ParkingPermitNumber', type: 'text', required: true},
                { field: 'ParkingPermitFee', type: 'money', required: true},
                { field: 'DtStart', type: 'date', required: false, html: { caption: 'DtStart', page: 0, column: 0 } },
                { field: 'DtStop', type: 'date', required: false, html: { caption: 'DtStop', page: 0, column: 0 } },
                { field: 'LastModTime', type: 'time', required: false, html: { caption: 'LastModTime', page: 0, column: 0 } },
                { field: 'LastModBy', type: 'int', required: false, html: { caption: 'LastModBy', page: 0, column: 0 } },
            ],
            onRefresh: function(event) {
                event.onComplete = function() {
                    var f = w2ui.RAVehicleForm,
                        header = "Edit Rental Agreement Vehicles ({0})";

                    // there is NO PETID actually, so have to work around with recid key
                    formRefreshCallBack(f, "recid", header);

                    // hide delete button if it is NewRecord
                    var isNewRecord = (w2ui.RAVehiclesGrid.get(f.record.recid, true) === null);
                    if (isNewRecord) {
                        $(f.box).find("button[name=delete]").addClass("hidden");
                    } else {
                        $(f.box).find("button[name=delete]").removeClass("hidden");
                    }
                };
            },
            actions : {
                save: function () {
                    var form = this;
                    var grid = w2ui.RAVehiclesGrid;
                    var errors = form.validate();
                    if (errors.length > 0) return;
                    var record = $.extend(true, { recid: grid.records.length + 1 }, form.record);
                    var recordsData = $.extend(true, [], grid.records);
                    var isNewRecord = (grid.get(record.recid, true) === null);

                    // if it doesn't exist then only push
                    if (isNewRecord) {
                        recordsData.push(record);
                    }

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // save this records in json Data
                    saveActiveCompData(recordsData, app.raFlowPartTypes.vehicles)
                        .done(function(data) {
                            if (data.status === 'success') {
                                // if null
                                if(isNewRecord) {
                                    grid.add(record);
                                }else {
                                    grid.set(record.recid, record);
                                }
                                form.clear();

                                // Disable "have vehicles?" checkbox if there is any record.
                                toggleHaveCheckBoxDisablity('RAVehiclesGrid');

                                // close the form
                                hideSliderContent();
                            } else {
                                form.message(data.message);
                            }
                        })
                        .fail(function(data) {
                            console.log("failure " + data);
                        });
                },
                saveadd: function () {
                    var BID = getCurrentBID(),
                        BUD = getBUDfromBID(BID);

                    var form = this;
                    var grid = w2ui.RAVehiclesGrid;
                    var errors = form.validate();
                    if (errors.length > 0) return;
                    var record = $.extend(true, { recid: grid.records.length + 1 }, form.record);
                    var recordsData = $.extend(true, [], grid.records);
                    var isNewRecord = (grid.get(record.recid, true) === null);

                    if (isNewRecord) {
                        recordsData.push(record);
                    }

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // save this records in json Data
                    saveActiveCompData(recordsData, app.raFlowPartTypes.vehicles)
                        .done(function(data) {
                            if (data.status === 'success') {
                                // clear the grid select recid
                                app.last.grid_sel_recid  = -1;
                                // selectNone
                                grid.selectNone();

                                if (isNewRecord) {
                                    grid.add(record);
                                } else {
                                    grid.set(record.recid, record);
                                }
                                form.record = getVehicleFormInitalRecord(BID, BUD, form.record);
                                form.record.recid =grid.records.length + 1;
                                form.refresh();
                                form.refresh();
                            } else {
                                form.message(data.message);
                            }
                        })
                        .fail(function(data) {
                            console.log("failure " + data);
                        });
                },
                delete: function () {
                    var form = this;
                    var grid = w2ui.RAVehiclesGrid;

                    // backup the records
                    var records = $.extend(true, [], grid.records);
                    for (var i = 0; i < records.length; i++) {
                        if(records[i].recid == form.record.recid) {
                            records.splice(i, 1);
                        }
                    }

                    // save this records in json Data
                    saveActiveCompData(records, app.raFlowPartTypes.vehicles)
                        .done(function(data) {
                            if (data.status === 'success') {
                                // clear the grid select recid
                                app.last.grid_sel_recid  =-1;
                                // selectNone
                                grid.selectNone();

                                grid.remove(form.record.recid);
                                form.clear();

                                // Disable "have vehicles?" checkbox if there is any record.
                                toggleHaveCheckBoxDisablity('RAVehiclesGrid');

                                // need to refresh the grid as it will re-assign new recid
                                reassignGridRecids(grid.name);

                                // close the form
                                hideSliderContent();
                            } else {
                                form.message(data.message);
                            }
                        })
                        .fail(function(data) {
                            console.log("failure " + data);
                        });

                }
            }
        });

        // vehicles grid
        $().w2grid({
            name    : 'RAVehiclesGrid',
            header  : 'Vehicles',
            show    : {
                toolbar         : true,
                toolbarSearch   : false,
                toolbarReload   : true,
                toolbarInput    : false,
                toolbarColumns  : false,
                footer          : true,
                toolbarAdd      : true   // indicates if toolbar add new button is visible
            },
            multiSelect: false,
            style   : 'border: 0px solid black; display: block;',
            columns : [
                {
                    field: 'recid',
                    hidden: true
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
                    size: '80px'
                },
                {
                    field: 'Make',
                    caption: 'Make',
                    size: '80px'
                },
                {
                    field: 'Model',
                    caption: 'Model',
                    size: '80px'
                },
                {
                    field: 'Color',
                    caption: 'Color',
                    size: '80px'
                },
                {
                    field: 'LicensePlateState',
                    caption: 'License Plate<br>State',
                    size: '100px'
                },
                {
                    field: 'LicensePlateNumber',
                    caption: 'License Plate<br>Number',
                    size: '100px'
                },
                {
                    field: 'ParkingPermitNumber',
                    caption: 'Parking Permit <br>Number',
                    size: '100px'
                },
                {
                    field: 'ParkingPermitFee',
                    caption: 'Parking Permit <br>Fee',
                    size: '100px',
                    render: 'money'
                },
                {
                    field: 'DtStart',
                    caption: 'DtStart',
                    size: '100px'
                },
                {
                    field: 'DtStop',
                    caption: 'DtStop',
                    size: '100px'
                }
            ],
            onChange: function (event) {
                event.onComplete = function () {
                    this.save();
                };
            },
            onRefresh: function(event) {
                // have to manage recid on every refresh of this grid
                event.onComplete = function() {
                    for (var j = 0; j < w2ui.RAVehiclesGrid.records.length; j++) {
                        w2ui.RAVehiclesGrid.records[j].recid = j + 1;
                    }
                };
            },
            onClick : function (event){
                event.onComplete = function () {
                    var yes_args = [this, event.recid],
                        no_args = [this],
                        no_callBack = function(grid) {
                            grid.select(app.last.grid_sel_recid);
                            return false;
                        },
                        yes_callBack = function (grid, recid) {
                            app.last.grid_sel_recid = parseInt(recid);

                            // keep highlighting current row in any case
                            grid.select(app.last.grid_sel_recid);

                            w2ui.RAVehicleForm.record = $.extend(true, {}, grid.get(app.last.grid_sel_recid));
                            showSliderContentW2UIComp(w2ui.RAVehicleForm, RACompConfig.vehicles.sliderWidth);
                            w2ui.RAVehicleForm.refresh();

                        };

                    // warn user if form content has been changed
                    form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
                };
            },
            onAdd   : function (/*event*/) {
                var yes_args = [this],
                    no_callBack = function() {
                        return false;
                    },
                    yes_callBack = function(grid) {
                        app.last.grid_sel_recid = -1;
                        grid.selectNone();

                        var BID = getCurrentBID(),
                            BUD = getBUDfromBID(BID);

                        w2ui.RAVehicleForm.record = getVehicleFormInitalRecord(BID, BUD, null);
                        w2ui.RAVehicleForm.record.recid = w2ui.RAVehiclesGrid.records.length + 1;
                        showSliderContentW2UIComp(w2ui.RAVehicleForm, RACompConfig.vehicles.sliderWidth);
                        w2ui.RAVehicleForm.refresh();
                    };

                // warn user if form content has been changed
                form_dirty_alert(yes_callBack, no_callBack, yes_args);
            }
        });
    }

    // now load grid in target division
    $('#ra-form #vehicles .form-container').w2render(w2ui.RAVehiclesGrid);

    // load the existing data in vehicles component
    setTimeout(function () {
        var grid = w2ui.RAVehiclesGrid;
        var i = getRAFlowPartTypeIndex(app.raFlowPartTypes.vehicles);
        if (i >= 0 && app.raflow.data[app.raflow.activeFlowID][i].Data) {
            grid.records = app.raflow.data[app.raflow.activeFlowID][i].Data;
            reassignGridRecids(grid.name);

            // lock the grid until "Have vehicles?" checkbox checked.
            lockOnGrid(grid.name);

        } else {
            grid.clear();
        }
    }, 500);
};

// -------------------------------------------------------------------------------
// Rental Agreement - Background info form
// -------------------------------------------------------------------------------

window.getRABGInfoFormInitRecord = function(BID, TCID){

    return {
        recid: 0,
        TCID: TCID,
        BID: BID,
        FirstName: "",
        MiddleName: "",
        LastName: "",
        BirthDate: "",
        SSN: "",
        DriverLicNo: "",
        TelephoneNo: "",
        EmailAddress: "",
        CurrentAddress: "",
        CurrentLandLordName: "",
        CurrentLandLordPhoneNo: "",
        CurrentLengthOfResidency: 0,
        CurrentReasonForMoving: "",
        PriorAddress: "",
        PriorLandLordName: "",
        PriorLandLordPhoneNo: "",
        PriorLengthOfResidency: 0,
        PriorReasonForMoving: "",
        Evicted: false,
        Convicted: false,
        Bankruptcy: false,
        Employer: "",
        Phone: "",
        Address: "",
        Position: "",
        GrossWages: 0,
        Comment: "",
        EmergencyContactName: "",
        EmergencyContactPhone: "",
        EmergencyContactAddress: ""
    };
};

window.loadRABGInfoForm = function () {

    var partType = app.raFlowPartTypes.bginfo;
    var partTypeIndex = getRAFlowPartTypeIndex(partType);

    if (partTypeIndex < 0){
        return;
    }

    // Fetch data from the server if there is any record available.
    getRAFlowPartData(partType)
        .done(function(data){
            if(data.status === 'success'){
                app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data = data.record.Data;
            }else {
                console.log(data.message);
            }
        })
        .fail(function(data){
            console.log("failure" + data);
        });

    // if form is loaded then return
    if (!("RABGInfoForm" in w2ui)) {

        // background info form
        $().w2form({
            name: 'RABGInfoForm',
            header: 'Background Information',
            style: 'border: 0px; background-color: transparent; display: block;',
            formURL: '/webclient/html/formrabginfo.html',
            toolbar : {
                items: [
                    { id: 'bt3', type: 'spacer' },
                    { id: 'btnClose', type: 'button', icon: 'fas fa-times'}
                ],
                onClick: function (event) {
                    switch (event.target){
                        case 'btnClose':
                            hideSliderContent();
                            break;
                    }
                }
            },
            fields: [
                {name: 'BID', type: 'int', required: true, html: { caption: 'BID', page:0, column: 0}},
                {name: 'TCID', type: 'int', required: true, html: { caption: 'TCID', page:0, column:0}},
                {name: 'FirstName', type: 'text', required: true},
                {name: 'MiddleName', type: 'text', required: true},
                {name: 'LastName', type: 'text', required: true},
                {name: 'BirthDate', type: 'date', required: true}, // Date of births of applicants
                {name: 'SSN', type: 'text', required: true}, // Social security number of applicants
                {name: 'DriverLicNo', type: 'text'}, // Driving licence number of applicants
                {name: 'TelephoneNo', type: 'text', required: true}, // Telephone no of applicants
                {name: 'EmailAddress', type: 'email', required: true}, // Email Address of applicants
                {name: 'CurrentAddress', type: 'text', required: true}, // Current Address
                {name: 'CurrentLandLordName', type: 'text', required: true}, // Current landlord's name
                {name: 'CurrentLandLordPhoneNo', type: 'text', required: true}, // Current landlord's phone number
                {name: 'CurrentLengthOfResidency', type: 'int', required: true}, // Length of residency at current address
                {name: 'CurrentReasonForMoving', type: 'text', required: true}, // Reason of moving from current address
                {name: 'PriorAddress', type: 'text'}, // Prior Address
                {name: 'PriorLandLordName', type: 'text'}, // Prior landlord's name
                {name: 'PriorLandLordPhoneNo', type: 'text'}, // Prior landlord's phone number
                {name: 'PriorLengthOfResidency', type: 'int'}, // Length of residency at Prior address
                {name: 'PriorReasonForMoving', type: 'text'}, // Reason of moving from Prior address
                {name: 'Evicted', type: 'checkbox', required: false}, // have you ever been Evicted
                {name: 'Convicted', type: 'checkbox', required: false}, // have you ever been Arrested or convicted of a crime
                {name: 'Bankruptcy', type: 'checkbox', required: false}, // have you ever been Declared Bankruptcy
                {name: 'Employer', type: 'text', required: true},
                {name: 'Phone', type: 'text', required: true},
                {name: 'Address', type: 'text', required: true},
                {name: 'Address2', type: 'text', required: false},
                {name: 'City', type: 'text', required: false},
                {name: 'State', type: 'list', options: {items: app.usStateAbbr}, required: false},
                {name: 'PostalCode', type: 'text', required: false},
                {name: 'Country', type: 'text', required: false},
                {name: 'Position', type: 'text', required: true},
                {name: 'GrossWages', type: 'money', required: true},
                {name: 'Comment', type: 'text'}, // In an effort to accommodate you, please advise us of any special needs
                {name: 'EmergencyContactName', type: 'text', required: true}, // Name of emergency contact
                {name: 'EmergencyContactPhone', type: 'text', required: true}, // Phone number of emergency contact
                {name: 'EmergencyContactAddress', type: 'text', required: true} // Address of emergency contact
            ],
            actions: {
                save: function () {
                    var form = this;

                    var errors = form.validate();
                    console.log(errors);
                    if (errors.length > 0) return;

                    var record = $.extend(true, {}, form.record);

                    // State filed
                    // TODO(Akshay): Think another way to modify State field if it is possible.
                    record.State = record.State.text;

                    // Convert integer to bool checkboxes fields
                    updateRABGInfoFormCheckboxes(record);

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    var partTypeIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.bginfo);
                    var bgInfoRecords = app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data || [];
                    var isTCIDMatched = false;

                    // update record if it is already exists
                    for(var recordIndex = 0; recordIndex < bgInfoRecords.length; recordIndex++){
                        if(bgInfoRecords[recordIndex].TCID === record.TCID){
                            bgInfoRecords[recordIndex] = record;
                            isTCIDMatched = true;
                            break;
                        }
                    }
                    // push record if it doesn't exists: What about last element
                    if(!isTCIDMatched){
                        bgInfoRecords.push(record);
                    }

                    var recordsData = app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data;
                    // save this records in json Data
                    saveActiveCompData(recordsData, app.raFlowPartTypes.bginfo)
                        .done(function(data) {
                            if (data.status === 'success') {

                                form.clear();

                                // close the form
                                hideSliderContent();
                            } else {
                                form.message(data.message);
                            }
                        })
                        .fail(function(data) {
                            console.log("failure " + data);
                        });
                }
            }
        });

        // transanctants info grid
        $().w2grid({
            name    : 'RABGInfoGrid',
            header  : 'Background information',
            show    : {
                footer          : true
            },
            style   : 'border: 1px solid black; display: block;',
            multiselect: false,
            columns : [
                {
                    field: 'recid',
                    hidden: true
                },
                {
                    field: 'TCID',
                    caption: 'TCID',
                    size: '50px',
                    hidden: true
                },
                {
                    field: 'FullName',
                    caption: 'Name',
                    size: '150px',
                    render: function (record) {
                        if(!record.IsCompany){
                            return record.FirstName + " " + record.MiddleName + " " + record.LastName;
                        }else{
                            return record.CompanyName;
                        }

                    }
                },
                {
                    field: 'IsRenter',
                    caption: 'IsRenter',
                    size: '100px',
                    hidden: true
                },
                {
                    field: 'IsOccupant',
                    caption: 'IsOccupant',
                    size: '100px',
                    hidden: true
                },
                {
                    field: 'IsGuarantor',
                    caption: 'IsGuarantor',
                    size: '100px',
                    hidden: true
                }
            ],
            onClick : function(event) {
                event.onComplete = function() {
                    var raBGInfoGridRecord = w2ui.RABGInfoGrid.get(event.recid); // record from the w2ui grid

                    var yes_args = [this, event.recid],
                        no_args = [this],
                        no_callBack = function(grid) {
                            grid.select(app.last.grid_sel_recid);
                            return false;
                        },
                        yes_callBack = function(grid, recid) {
                            app.last.grid_sel_recid = parseInt(recid);

                            // keep highlighting current row in any case
                            grid.select(app.last.grid_sel_recid);

                            w2ui.RABGInfoForm.record = $.extend(true, {}, grid.get(app.last.grid_sel_recid));

                            // w2ui.RABGInfoForm.record = $.extend(true, {}, grid.get(app.last.grid_sel_recid));
                            showSliderContentW2UIComp(w2ui.RABGInfoForm, RACompConfig.bginfo.sliderWidth);

                            // get transanctant information from the server
                            getRATransanctantDetail(raBGInfoGridRecord.TCID)
                                .done(function (data){
                                    // Hide these all fields when transanctant is only user.
                                    var listOfHiddenFields = ["CurrentAddress", "CurrentLandLordName",
                                        "CurrentLandLordPhoneNo", "CurrentLengthOfResidency", "CurrentReasonForMoving",
                                        "PriorAddress", "PriorLandLordName", "PriorLandLordPhoneNo",
                                        "PriorLengthOfResidency", "PriorReasonForMoving"];

                                    // These all fields are not required when transanctant is only user
                                    var listOfNotRequiredFields = ["SSN", "TelephoneNo",
                                        "Phone", "EmailAddress", "Position",
                                        "GrossWages", "CurrentAddress", "CurrentLandLordName",
                                        "CurrentLandLordPhoneNo", "CurrentReasonForMoving"];

                                    if(data.status === 'success'){
                                        var record = data.record; // record from the server response
                                        // w2ui.RABGInfoForm.record = data.record.Data;
                                        var BID = getCurrentBID();

                                        // Set the form tile
                                        setRABGInfoFormHeader(record);

                                        // Display/Required field based on transanctant type
                                        if(raBGInfoGridRecord.IsOccupant && !raBGInfoGridRecord.IsRenter && !raBGInfoGridRecord.IsGuarantor){
                                            // hide fields
                                            showHideRABGInfoFormFields(listOfHiddenFields, true);

                                            // not require fields
                                            setNotRequiredFields(listOfNotRequiredFields, false);
                                        }else{
                                            // show fields
                                            showHideRABGInfoFormFields(listOfHiddenFields, false);

                                            // require fields
                                            setNotRequiredFields(listOfNotRequiredFields, true);

                                        }

                                        // If TCID found than set form fields value else initialize it.
                                        var partTypeIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.bginfo);
                                        var bgInfoRecords = app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data || [];
                                        var isTCIDMatched = false;

                                        for(var recordIndex = 0; recordIndex < bgInfoRecords.length; recordIndex++){
                                            if(bgInfoRecords[recordIndex].TCID === raBGInfoGridRecord.TCID){
                                                w2ui.RABGInfoForm.record = bgInfoRecords[recordIndex];
                                                isTCIDMatched = true;
                                                break;
                                            }
                                        }
                                        if(!isTCIDMatched){
                                            // Assign default values to form fields
                                            w2ui.RABGInfoForm.record = getRABGInfoFormInitRecord(BID, raBGInfoGridRecord.TCID);

                                            // Set latest value for transanctant basic information from the server only
                                            setRABGInfoFormFields(record);
                                        }


                                        w2ui.RABGInfoForm.refresh(); // need to refresh for form changes
                                    }else {
                                        console.log(data.message);
                                    }
                                })
                                .fail(function (data) {
                                    console.log("failure" + data);
                                });
                        };

                    // warn user if form content has been changed
                    form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
                };
            }
        });
    }

    // now load form in div
    $('#ra-form #bginfo').w2render(w2ui.RABGInfoGrid);

    var i = getRAFlowPartTypeIndex(app.raFlowPartTypes.people);
    var peopleData = app.raflow.data[app.raflow.activeFlowID][i].Data;

    // load the existing data in Background Info grid
    setTimeout(function (peopleData) {
        var grid = w2ui.RABGInfoGrid;

        var i = getRAFlowPartTypeIndex(app.raFlowPartTypes.bginfo);
        if (i >= 0 && app.raflow.data[app.raflow.activeFlowID][i].Data) {
            grid.records = peopleData;
            grid.refresh();
            reassignGridRecids(grid.name);
        } else {
            grid.clear();
        }
    }, 500, peopleData);
};

// -------------------------------------------------------------------------------
// Rental Agreement - Rentables Grid
// -------------------------------------------------------------------------------
window.getRentableFeeFormInitalRecord = function (BID, gridLen) {
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

// -------------------------------------------------------------------------------
// getAutoPopulateARs - pull down all account rules which are set to auto populate
//                      to new rental agreement
// -------------------------------------------------------------------------------
window.getAutoPopulateARs = function() {
    return $.ajax({
        url: ''
    });
};

window.loadRARentablesGrid = function () {

    var partType = app.raFlowPartTypes.rentables;
    var partTypeIndex = getRAFlowPartTypeIndex(partType);

    if (partTypeIndex < 0){
        return;
    }

    // Fetch data from the server if there is any record available.
    getRAFlowPartData(partType)
        .done(function(data){
            if(data.status === 'success'){
                app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data = data.record.Data;
            }else {
                console.log(data.message);
            }
        })
        .fail(function(data){
            console.log("failure" + data);
        });

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
                    {id: 'add', type: 'button', caption: 'Add Record', icon: 'w2ui-icon-plus'},
                ],
                onClick: function (event) {
                    switch(event.target) {
                        case "add":
                            // get auto populated to new RA account rules
                            var data = {
                              "type":"FLAGS",
                              "FLAGS": app.arFLAGS.PopulateOnRA
                            };
                            var BID = getCurrentBID();
                            $.ajax({
                                url: "/v1/arslist/" + BID.toString() + "/",
                                method: "POST",
                                contentType: "application/json",
                                data: JSON.stringify(data),
                            }).done(function(data) {
                                showSliderContentW2UIComp(w2ui.RARentableFeesGrid, RACompConfig.rentables.sliderWidth);
                            });
                            break;
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
                    size: '100px',
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

        // rentables grid
        $().w2grid({
            name: 'RARentableFeesGrid',
            header: 'Rentables Fees',
            show: {
                toolbar: true,
                footer: true,
                header: true,
            },
            style: 'border: 2px solid white; display: block;',
            toolbar: {
                items: [
                    {id: 'add', type: 'button', caption: 'Add Record', icon: 'w2ui-icon-plus'},
                    {id: 'bt3', type: 'spacer'},
                    {id: 'btnClose', type: 'button', icon: 'fas fa-times'},
                ],
                onClick: function (event) {
                    switch(event.target) {
                        case "add":
                            var sliderID = 2;
                            appendNewSlider(sliderID);
                            $("#raflow-container")
                                .find(".slider[data-slider-id="+sliderID+"]")
                                .find(".slider-content")
                                .width(400)
                                .w2render(w2ui.RARentableFeesForm);
                            break;
                        case "btnClose":
                            hideSliderContent();
                            break;
                    }
                }
            },
            columns: [
                {
                    field: 'recid',
                    hidden: true,
                },
                {
                    field: 'ARID',
                    hidden: true
                },
                {
                    field: 'BID',
                    hidden: true
                },
                {
                    field: 'Name',
                    caption: 'Fee',
                    size: '250px',
                },
                {
                    field: 'Amount',
                    caption: 'Amount',
                    size: '100px',
                    render: 'money',
                },
                {
                    field: 'RentCycle',
                    caption: 'Cycle',
                    size: '100px',
                },
                {
                    field: 'Epoch',
                    caption: 'Epoch',
                    size: '100px',
                },
                {
                    field: 'RentPeriod',
                    caption: 'Rent Period',
                    size: '100px',
                },
                {
                    field: 'UsePeriod',
                    caption: 'Use Period',
                    size: '100px',
                },
                {
                    field: 'ContractRent',
                    caption: 'At Signing',
                    size: '100px',
                    render: 'money',
                },
                {
                    field: 'ProrateAmt',
                    caption: 'Prorate',
                    size: '100px',
                    render: 'money',
                },
                {
                    field: 'SalesTaxAmt',
                    caption: 'Sales Tax Amt',
                    size: '100px',
                    render: 'money',
                },
                {
                    field: 'SalesTax',
                    caption: 'Sales Tax',
                    size: '100px',
                },
                {
                    field: 'TransOccAmt',
                    caption: 'Trans Occ Amt',
                    size: '100px',
                    render: 'money',
                },
                {
                    field: 'TransOcc',
                    caption: 'Trans Occ',
                    size: '100px',
                }
            ],
            onChange: function (event) {
                event.onComplete = function () {
                    this.save();
                };
            }
        });

        // new transactant form especially for this RA flow
        $().w2form({
            name: 'RARentableFeesForm',
            header: 'Add New Rentable Fee',
            style: 'display: block;',
            formURL: '/webclient/html/formra-rentablefee.html',
            focus: -1,
            fields: [
                {name: 'BID', type: 'int', required: true, html: {page: 0, column: 0}},
                {name: 'BUD', type: 'list', required: true, options: {items: app.businesses}, html: {page: 0, column: 0}},
            ],
            toolbar : {
                items: [
                    { id: 'bt3', type: 'spacer' },
                    { id: 'btnClose', type: 'button', icon: 'fas fa-times'}
                ],
                onClick: function (event) {
                    switch (event.target){
                        case 'btnClose':
                            hideSliderContent(2);
                            break;
                    }
                }
            },
            actions: {
                reset: function () {
                    this.clear();
                },
                save: function() {
                    var f = this;
                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    f.save({}, function(data) {
                        if (data.status === 'error') {
                            f.message(data.message);
                            return;
                        }
                        hideSliderContent(2);
                    });
                },
                saveadd: function() {
                    var BID = getCurrentBID(),
                        BUD = getBUDfromBID(BID),
                        f = this;

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    f.save({}, function(data) {
                        if (data.status === 'error') {
                            f.message(data.message);
                            return;
                        }

                        // f.record = getRAAddTransactantFormInitRec(BID, BUD, f.record);
                        f.refresh();
                    });
                },
            },
            onChange: function(event) {
                event.onComplete = function() {
                       // formRecDiffer: 1=current record, 2=original record, 3=diff object
                    var diff = formRecDiffer(this.record, app.active_form_original, {});
                    // if diff == {} then make dirty flag as false, else true
                    if ($.isPlainObject(diff) && $.isEmptyObject(diff)) {
                        app.form_is_dirty = false;
                    } else {
                        app.form_is_dirty = true;
                    }
                };
            },
            onRefresh: function(event) {
                var f = this;
                event.onComplete = function() {
                    var BID = getCurrentBID(),
                        BUD = getBUDfromBID(BID);

                    f.record.BID = BID;
                    f.record.BUD = BUD;

                    // there is NO PETID actually, so have to work around with recid key
                    formRefreshCallBack(f, "recid");
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

    var partType = app.raFlowPartTypes.feesterms;
    var partTypeIndex = getRAFlowPartTypeIndex(partType);

    if (partTypeIndex < 0){
        return;
    }

    // Fetch data from the server if there is any record available.
    getRAFlowPartData(partType)
        .done(function(data){
            if(data.status === 'success'){
                app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data = data.record.Data;
            }else {
                console.log(data.message);
            }
        })
        .fail(function(data){
            console.log("failure" + data);
        });

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
                    size: '80px',
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
        sliderWidth: 0,
    },
    "people": {
        loader: loadRAPeopleForm,
        w2uiComp: "RAPeopleForm",
        sliderWidth: 400,
    },
    "pets": {
        loader: loadRAPetsGrid,
        w2uiComp: "RAPetsGrid",
        sliderWidth: 400,
    },
    "vehicles": {
        loader: loadRAVehiclesGrid,
        w2uiComp: "RAVehiclesGrid",
        sliderWidth: 400,
    },
    "bginfo": {
        loader: loadRABGInfoForm,
        w2uiComp: "RABGInfoForm",
        sliderWidth: 700,
    },
    "rentables": {
        loader: loadRARentablesGrid,
        w2uiComp: "RARentablesGrid",
        sliderWidth: 800,
    },
    "feesterms": {
        loader: loadRAFeesTermsGrid,
        w2uiComp: "RAFeesTermsGrid",
        sliderWidth: 400,
    },
    "final": {
        loader: null,
        w2uiComp: "",
        sliderWidth: 0,
    }
};

//-----------------------------------------------------------------------------
// showSliderContentW2UIComp - renders the w2ui component into slider-content
//                             and apply the given width to it
//
// @params
//   w2uiComp = w2ui component
//   width    = width to apply to slider content div
//-----------------------------------------------------------------------------
window.showSliderContentW2UIComp = function(w2uiComp, width, sliderID) {
    if (!sliderID) {
        sliderID = 1;
    }

    $("#raflow-container .slider[data-slider-id="+sliderID+"]").show();
    $("#raflow-container .slider[data-slider-id="+sliderID+"] .slider-content").width(width);
    $("#raflow-container .slider[data-slider-id="+sliderID+"] .slider-content").w2render(w2uiComp);
};

//-----------------------------------------------------------------------------
// hideSliderContent - hide the slider and empty the content inside
//                     slider-content div
//
//-----------------------------------------------------------------------------
window.hideSliderContent = function(sliderID) {
    if (!sliderID) {
        sliderID = 1;
    }

    $("#raflow-container .slider[data-slider-id="+sliderID+"]").hide();
    $("#raflow-container .slider[data-slider-id="+sliderID+"] .slider-content").width(0);
    $("#raflow-container .slider[data-slider-id="+sliderID+"] .slider-content").empty();
};

//-----------------------------------------------------------------------------
// appendNewSlider - append new right slider in the DOM dynamically
//-----------------------------------------------------------------------------
window.appendNewSlider = function(sliderID) {
    // if sliderID exists then don't append
    if ($("#raflow-container").find("div[data-slider-id="+ sliderID +"]").length > 0) {
        return;
    }

    var slidersLength = $("#raflow-container").find(".slider").length;
    var recentAddedSlider = $("#raflow-container")
        .find("div[data-slider-id="+ slidersLength +"]");

    var newSlider = recentAddedSlider.clone();
    newSlider.attr("data-slider-id", slidersLength + 1);
    recentAddedSlider.after(newSlider);
    newSlider.css("z-index", parseInt(recentAddedSlider.css("z-index")) + 10);
    newSlider.find(".slider-content").empty().width(0);
};
