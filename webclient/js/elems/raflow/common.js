/* global
    RACompConfig, HideSliderContent, appendNewSlider, ShowSliderContentW2UIComp, displayFormFieldsError,
    loadTargetSection, requiredFieldsFulFilled, initRAFlowAjax, getRecIDFromTMPASMID, getFeeIndex,
    saveActiveCompData, getRAFlowCompData, displayActiveComponentError, displayRAPetsGridError, dispalyRAPeopleGridError,
    lockOnGrid, getApprovals, updateFlowData, updateFlowCopy, displayErrorDot, initBizErrors,
    dispalyRARentablesGridError, dispalyRAVehiclesGridError, dispalyRAParentChildGridError, dispalyRATiePeopleGridError,
    GetCurrentFlowID, FlowFilled, ReassignPeopleGridRecords, AssignPetsGridRecords, AssignVehiclesGridRecords, AssignRentableGridRecords,
    GetGridToolbarAddButtonID, HideRAFlowLoader, toggleNonFieldsErrorDisplay, displayNonFieldsError
*/

"use strict";

//-----------------------------------------------------------------------------
// GetRefNoByRAIDFromGrid returns UserRefNo By RAID from applicantsGrid RECORDS
//-----------------------------------------------------------------------------
window.GetRefNoByRAIDFromGrid = function(RAID) {
    var RefNo = "";
    w2ui.applicantsGrid.records.forEach(function(gridRec) {
        if (gridRec.RAID == RAID) {
            RefNo = gridRec.UserRefNo;
            return;
        }
    });
    return RefNo;
};

//-----------------------------------------------------------------------------
// GetRAIDByRefNoFromGrid returns RAID By UserRefNo from applicantsGrid RECORDS
//-----------------------------------------------------------------------------
window.GetRAIDByRefNoFromGrid = function(RefNo) {
    var RAID = -1;
    w2ui.applicantsGrid.records.forEach(function(gridRec) {
        if (gridRec.UserRefNo == RefNo) {
            RAID = gridRec.RAID;
            return;
        }
    });
    return RAID;
};

//-----------------------------------------------------------------------------
// GetCurrentFlowID returns current flow ID
// which user looking at the flow currently
//-----------------------------------------------------------------------------
window.GetCurrentFlowID = function() {
    if (Object.keys(app.raflow.Flow).length != 0) { // IF NOT BLANK THEN
        return app.raflow.Flow.FlowID;
    }
    return 0;
};

//-----------------------------------------------------------------------------
// NEXT BUTTON CLICK EVENT HANDLER
//-----------------------------------------------------------------------------
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

//-----------------------------------------------------------------------------
// PREVIOUS BUTTON CLICK EVENT HANDLER
//-----------------------------------------------------------------------------
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

//-----------------------------------------------------------------------------
// Get Approvals BUTTON CLICK EVENT HANDLER
//-----------------------------------------------------------------------------
$(document).on('click', '#ra-form #save-ra-flow-btn', function () {
    getApprovals().done(function (data) {

        app.raflow.validationErrors = {
            dates: data.errors.dates.total > 0 || data.nonFieldsErrors.dates.length > 0,
            people: data.errors.people.length > 0 || data.nonFieldsErrors.people.length > 0,
            pets: data.errors.pets.length > 0 || data.nonFieldsErrors.pets.length > 0,
            vehicles: data.errors.vehicles.length > 0 || data.nonFieldsErrors.vehicles.length > 0,
            rentables: data.errors.rentables.length > 0 || data.nonFieldsErrors.rentables.length > 0,
            parentchild: data.errors.parentchild.length > 0 || data.nonFieldsErrors.parentchild.length > 0,
            tie: data.errors.tie.people.length > 0 || data.nonFieldsErrors.tie.length > 0
        };

        displayErrorDot();

        displayActiveComponentError();

        if(data.total === 0 && data.errortype === "biz"){
            alert("TODO: You'r good to go for pending first approval."); // TODO: Change its state to pending first approval. Remove this alert
        }

    });
});

// initBizErrors To initialize bizError local copy for active flow
window.initBizErrors = function(){
    app.raflow.validationErrors = {
        dates: false,
        people: false,
        pets: false,
        vehicles: false,
        rentables: false,
        parentchild: false,
        tie: false
    };
};

// displayErrorDot it show red dot on each section of section contain biz logic error
window.displayErrorDot = function(){
    for (var comp in app.raFlowPartTypes) {
        if (app.raflow.validationErrors[comp]) {
            $("#progressbar #steps-list li[data-target='#" + comp + "'] .error").addClass("error-true");
        } else {
            $("#progressbar #steps-list li[data-target='#" + comp + "'] .error").removeClass("error-true");
        }
    }
};

window.getApprovals = function(){

    var bid = getCurrentBID();
    var FlowID = GetCurrentFlowID();
    var data = {
        "cmd": "get",
        "FlowID": FlowID
    };

    return $.ajax({
        url: "/v1/validate-raflow/" + bid.toString(),
        method: "POST",
        contentType: "application/json",
        dataType: "json",
        data: JSON.stringify(data),
        success: function (data) {
            console.info(data);
            // Update validationCheck error local copy
            app.raflow.validationCheck = data;
        },
        error: function (data) {
            console.error(data);
        }
    });
};

//-----------------------------------------------------------------------------
// FORM WIZARD STEP LINK CLICK EVENT HANDLER
//-----------------------------------------------------------------------------
$(document).on('click', '#ra-form #progressbar #steps-list a', function () {
    var active_comp = $(".ra-form-component:visible");

    // load target form
    var target = $(this).closest("li").attr("data-target");
    target = target.split('#').join("");

    loadTargetSection(target, active_comp.attr("id"));

    // because of 'a' tag, return false
    return false;
});

//-----------------------------------------------------------------------------
// lockOnGrid - Lock grid if checkbox is unchecked(false).
//              Unlock grid if checkbox is checked(true).
//              Lock grid when there is no record in the grid.
//
// @params
//   gridName   = name of the grid
//-----------------------------------------------------------------------------
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

//-----------------------------------------------------------------------------
// getRAFlowCompData - get the flow component data stored locally in app.raflow
//
// @params
//   key    = flow component key
//   FlowID = for which FlowID's component
//-----------------------------------------------------------------------------
window.getRAFlowCompData = function (compKey) {

    var bid = getCurrentBID();

    var flowJSON = app.raflow.Flow;
    if (flowJSON.Data) {
        return flowJSON.Data[compKey];
    }

    return null;
};

//-----------------------------------------------------------------------------
// setRAFlowCompData - set the flow component data locally in app.raflow
//
// @params
//   key    = flow component key
//   data   = data to set in the component
//-----------------------------------------------------------------------------
window.setRAFlowCompData = function (compKey, data) {

    var bid = getCurrentBID();

    var flowJSON = app.raflow.Flow;
    if (flowJSON.Data) {
        flowJSON.Data[compKey] = data;
    }
};

//-----------------------------------------------------------------------------
// saveActiveCompData - save component modified data on the server
//
// @params
//   compData   = modified latest component data
//   compID     = component key id
//-----------------------------------------------------------------------------
window.saveActiveCompData = function (compData, compID) {

    // IF RAID VERSION THEN DON"T DO ANYTHING
    if (app.raflow.version === "raid") {
        return;
    }

    var bid = getCurrentBID();
    var FlowID = GetCurrentFlowID();

    // temporary data
    var data = {
        "cmd": "save",
        "FlowType": app.raflow.name,
        "FlowID": FlowID,
        "FlowPartKey": compID,
        "BID": bid,
        "Data": compData
    };

    return $.ajax({
        url: "/v1/flow/" + bid.toString() + "/" + FlowID.toString(),
        method: "POST",
        contentType: "application/json",
        dataType: "json",
        data: JSON.stringify(data),
        success: function (data) {
            if (data.status != "error") {
                console.log("data has been saved for: ", FlowID, ", compID: ", compID);
                // Update flow local copy and green checks
                updateFlowData(data);
            } else {
                console.error(data.message);
            }
        },
        error: function (data) {
            console.error(data);
        }
    });
};

//-----------------------------------------------------------------------------
// initRAFlowAjax - will initiate new rental agreement flow and returns ajax
//                  promise
//-----------------------------------------------------------------------------
window.initRAFlowAjax = function () {
    var bid = getCurrentBID();

    return $.ajax({
        url: "/v1/flow/" + bid.toString() + "/0",
        method: "POST",
        contentType: "application/json",
        dataType: "json",
        data: JSON.stringify({"cmd": "init", "FlowType": app.raflow.name}),
        success: function (data) {
            if (data.status != "error") {
                app.raflow.version = "refno";
                // Update flow local copy and green checks
                updateFlowData(data);
            }
        },
        error: function (data) {
            console.log(data);
        }
    });
};

//-----------------------------------------------------------------------------
// GetRAFlowDataAjax - get the ajax data from the server and returns ajax promise
//
// @params
//   RefNo      = User Ref no of the raflow
//   RAID       = Rental Agreement
//   version    = which version of raflow
//-----------------------------------------------------------------------------
window.GetRAFlowDataAjax = function(UserRefNo, RAID, version) {
    var bid = getCurrentBID();

    var reqData = {
        "cmd":          "get",
        "UserRefNo":    UserRefNo,
        "RAID":         RAID,
        "Version":      version,
        "FlowType":     "RA"
    };

    return $.ajax({
        url: "/v1/flow/" + bid.toString() + "/",
        method: "POST",
        contentType: "application/json",
        dataType: "json",
        data: JSON.stringify(reqData),
        beforeSend: function() {
            // show the loader
            HideRAFlowLoader(false);
            $("#raflow-container .loader").css("display", "flex");
        },
        success: function (data) {
            if (data.status !== "error") {
                app.raflow.version = version;
                updateFlowData(data);
            }
        },
        error: function (data) {
            console.log(data);
        },
        complete: function() {
            // hide the loader
            HideRAFlowLoader(true);
        }
    });
};

// HideRAFlowLoader loader to show the progress while fetching data from the server
// which also disabled the controls in toolbar
window.HideRAFlowLoader = function(hide) {
    app.raflow.loading = !hide;
    if (hide) {
        if (w2ui.newraLayout) {
            $(w2ui.newraLayout.get("main").toolbar.box).find("button").prop('disabled', true);
        }
        $("#raflow-container .loader").hide();
    } else {
        if (w2ui.newraLayout) {
            $(w2ui.newraLayout.get("main").toolbar.box).find("button").prop('disabled', false);
        }
        $("#raflow-container .loader").show();
    }
};

// updateFlowData
window.updateFlowData = function(data){
    updateFlowCopy(data.record.Flow);
    setTimeout(function() {
        // Enable/Disable green check
        FlowFilled(data.record);
    }, 500);
};

// updateFlowCopy
window.updateFlowCopy = function(flow){
    app.raflow.Flow = flow;

    // ALSO UPDATE THE LOCAL RECORDS IN GRID
    w2ui.applicantsGrid.records.forEach(function(gridRec) {
        if (gridRec.UserRefNo === flow.UserRefNo || gridRec.RAID === flow.ID) {
            if (flow.UserRefNo) { // IF AVAILABLE THEN ONLY SET
                gridRec.UserRefNo = flow.UserRefNo;
            }
            if (flow.ID) { // IF AVAILABLE THEN ONLY SET
                gridRec.RAID = flow.ID;
            }

            w2ui.applicantsGrid.refresh();
            return;
        }
    });
};

// -----------------------------------------------------
// FlowFilled:
// Enable/Disable green checks
// Enable/Disable get approvals button
// raflow parts
// -----------------------------------------------------
window.FlowFilled = function(data) {

    // Update local copy of basicCheck and FlowFilledData
    app.raflow.basicCheck = data.BasicCheck;
    app.raflow.FlowFilledData = data.DataFulfilled;

    // Enable/Disable green check for the each section
    var active_comp = $(".ra-form-component:visible");
    var active_comp_id = active_comp.attr("id");

    for (var comp in app.raFlowPartTypes) {
        // if required fields are fulfilled then mark this slide as done

        // Apply green mark when comp is not active and when it fulfilled the requirements
        if (app.raflow.FlowFilledData[comp] && active_comp_id !== comp) {
            $("#progressbar #steps-list li[data-target='#" + comp + "']").addClass("done");
        } else {
            $("#progressbar #steps-list li[data-target='#" + comp + "']").removeClass("done");
        }
    }
};

// load form according to target
window.loadTargetSection = function (target, previousActiveCompID) {

    /*if ($("#progressbar #steps-list li[data-target='#" + target + "']").hasClass("done")) {
        console.log("target has been saved", target);
    } else {}*/

    // get component data based on ID from locally
    var compData = getRAFlowCompData(previousActiveCompID);

    // default would be compData
    var modCompData = compData;

    switch (previousActiveCompID) {
        case "dates":
            modCompData = w2ui.RADatesForm.record;
            w2ui.RADatesForm.actions.reset();
            break;
        case "people":
            // modCompData = compData;
            w2ui.RAPeopleGrid.clear();
            w2ui.RAPeopleForm.actions.reset();
            break;
        case "pets":
            // modCompData = compData;
            w2ui.RAPetsGrid.clear();
            w2ui.RAPetForm.actions.reset();
            break;
        case "vehicles":
            // modCompData = compData;
            w2ui.RAVehiclesGrid.clear();
            w2ui.RAVehicleForm.actions.reset();
            break;
        case "rentables":
            // modCompData = compData;
            w2ui.RARentablesGrid.clear();
            w2ui.RARentableFeesGrid.clear();
            w2ui.RARentableFeeForm.actions.reset();
            break;
        case "parentchild":
            // modCompData = compData;
            w2ui.RAParentChildGrid.clear();
            break;
        case "tie":
            // modCompData = compData;
            w2ui.RATiePeopleGrid.clear();
            break;
        case "final":
            modCompData = null;
            w2ui.RAFinalRentablesFeesGrid.clear();
            w2ui.RAFinalPetsFeesGrid.clear();
            w2ui.RAFinalVehiclesFeesGrid.clear();
            break;
        default:
            alert("invalid active comp: " + previousActiveCompID);
            return;
    }

    // get part type from the class index
    if (modCompData) {
        // save the content on server for active component
        saveActiveCompData(modCompData, previousActiveCompID);
    }

    // hide active component
    $("#progressbar #steps-list li[data-target='#" + previousActiveCompID + "']").removeClass("active");
    $(".ra-form-component#" + previousActiveCompID).hide();

    // show target component
    $("#progressbar #steps-list li[data-target='#" + target + "']").removeClass("done").addClass("active");
    $(".ra-form-component#" + target).show();

    // display target comp non fields error
    displayNonFieldsError(target);

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
    if (targetLoader.length > 0) {
        window[targetLoader]();
        /*setTimeout(function() {
            var validateForm = compIDw2uiForms[previousActiveCompID];
            if (typeof w2ui[validateForm] !== "undefined") {
                var issues = w2ui[validateForm].validate();
                if (!(Array.isArray(issues) && issues.length > 0)) {
                    // $("#progressbar #steps-list li[data-target='#" + previousActiveCompID + "']").addClass("done");
                }
            }
        }, 500);*/
    } else {
        console.log("unknown target from nav li: ", target);
    }
};

//-----------------------------------------------------------------------------
// ShowSliderContentW2UIComp - renders the w2ui component into slider-content
//                             and apply the given width to it
//
// @params
//   w2uiComp = w2ui component
//   width    = width to apply to slider content div
//   sliderID = slider ID (as in stack fashion)
//-----------------------------------------------------------------------------
window.ShowSliderContentW2UIComp = function(w2uiComp, width, sliderID) {
    if (!sliderID) {
        sliderID = 1;
    }

    $("#raflow-container .slider[data-slider-id="+sliderID+"]").show();
    $("#raflow-container .slider[data-slider-id="+sliderID+"] .slider-content").width(width);
    $("#raflow-container .slider[data-slider-id="+sliderID+"] .slider-content").w2render(w2uiComp);
};

//-----------------------------------------------------------------------------
// HideAllSliderContent - hides all slider and empty the content inside
//                        slider-content div
//-----------------------------------------------------------------------------
window.HideAllSliderContent = function() {
    $("#raflow-container .slider").hide();
    $("#raflow-container .slider .slider-content").width(0);
    $("#raflow-container .slider .slider-content").empty();
};

//-----------------------------------------------------------------------------
// HideSliderContent - hide the slider and empty the content inside
//                     slider-content div
//
// @params
//      sliderID = slider ID (as in stack fashion)
//-----------------------------------------------------------------------------
window.HideSliderContent = function(sliderID) {
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

//-----------------------------------------------------------------------------
// getVehicleFees - will list down vehicle fees for a business
//-----------------------------------------------------------------------------
window.getVehicleFees = function () {
    var bid = getCurrentBID();

    return $.ajax({
        url: "/v1/vehiclefees/" + bid.toString() + "/0",
        method: "GET",
        contentType: "application/json",
        dataType: "json",
        success: function (data) {
            if (data.status != "error") {
                app.vehicleFees[bid] = data.records;
            }
        },
        error: function (data) {
            console.log(data);
        }
    });
};

//-----------------------------------------------------------------------------
// displayActiveComponentError - it displays/highlight error for active component
//-----------------------------------------------------------------------------
window.displayActiveComponentError = function () {
    // get the current component (to be previous one)
    var active_comp = $(".ra-form-component:visible");

    // get active component id
    var active_comp_id = active_comp.attr("id");

    switch (active_comp_id) {
        case "dates":
            break;
        case "people":
            ReassignPeopleGridRecords();
            break;
        case "pets":
            AssignPetsGridRecords();
            break;
        case "vehicles":
            AssignVehiclesGridRecords();
            break;
        case "rentables":
            AssignRentableGridRecords();
            break;
        case "parentchild":
            w2ui.RAParentChildGrid.refresh();
            dispalyRAParentChildGridError();
            break;
        case "tie":
            w2ui.RATiePeopleGrid.refresh();
            dispalyRATiePeopleGridError();
            break;
        case "final":
            break;
        default:
            alert("invalid active comp: " + active_comp_id);
            return;
    }
};

// getRecIDFromTMPASMID It returns recid of grid record which matches TMPASMID
window.getRecIDFromTMPASMID = function(grid, TMPASMID){
    var recid;
    for (var i = 0; i < grid.records.length; i++) {
        if (grid.records[i].TMPASMID === TMPASMID) {
            recid = grid.records[i].recid;
        }
    }
    return recid;
};

// displayFormFieldsError It display form fields error  for record
window.displayFormFieldsError = function(index, records, formName){
    // Iterate through fields with errors
    for(var key in records[index].errors){
        var field = $("[name=" + formName + "] input#" + key);
        var error = records[index].errors[key].join(", ");

        field.css("border-color", "red");
        field.after("<small class='error'>" + error + "</small>");
    }
};

// getFeeIndex it return an index of fee which have TMPASMID
window.getFeeIndex = function (TMPASMID, fees) {

    var index = -1;

    for(var i = 0; i < fees.length; i++){
        // If TMPASMID doesn't match iterate for next element
        if(fees[i].TMPASMID === TMPASMID){
            index = i;
            break;
        }
    }

    return index;
};

//-----------------------------------------------------------------------
// EnableDisableRAFlowVersionInputs
//      enable/disable the inputs of form based on
//      the current version of raflow.
//      If "raid" then it'll disable else enable the inputs.
//
// @params
//   form       = w2ui form component
//-----------------------------------------------------------------------
window.EnableDisableRAFlowVersionInputs = function(form) {
    if (app.raflow.version === "raid") { // DISABLE ALL INPUTS & BUTTONS
        $(form.box).find("input").prop("disabled", true);
        $(form.box).find("button[class=w2ui-btn]").hide();
        $(form.box).find("div[class=w2ui-buttons]").hide();
   } else if (app.raflow.version === "refno") { // ENABLE ALL INPUTS & BUTTONS
        $(form.box).find("input").not("input[name=BUD]").prop("disabled", false);
        $(form.box).find("button[class=w2ui-btn]").show();
        $(form.box).find("div[class=w2ui-buttons]").show();
   }
};

//-----------------------------------------------------------------------
// EnableDisableRAFlowVersionGrid
//      lock/unlock the entire grid base on the current version of raflow.
//      If "raid" then it'll disable else enable the inputs.
//
// @params
//   grid       = w2ui grid component
//-----------------------------------------------------------------------
window.EnableDisableRAFlowVersionGrid = function(grid) {
    if (app.raflow.version === "raid") { // DISABLE ALL INPUTS & BUTTONS
        grid.lock();
   } else if (app.raflow.version === "refno") { // ENABLE ALL INPUTS & BUTTONS
        grid.unlock();
   }
};

// GetGridToolbarAddButtonID to get the DOM ID of add button in grid by gridName
window.GetGridToolbarAddButtonID = function(gridName) {
    return "tb_" + gridName +"_toolbar_item_w2ui-add";
};

// ShowHideGridToolbarAddButton shows/hides add button based on raflow version
window.ShowHideGridToolbarAddButton = function(gridName) {
    var addBtnID = GetGridToolbarAddButtonID(gridName);
    if (app.raflow.version === "raid") {
        $("#"+addBtnID).hide();
    } else if (app.raflow.version === "refno") {
        $("#"+addBtnID).show();
    }
};

//-----------------------------------------------------------------------------
// DeleteRAFlowAJAX - will request to remove ref.no version raflow
//-----------------------------------------------------------------------------
window.DeleteRAFlowAJAX = function (UserRefNo) {
    var bid = getCurrentBID();

    return $.ajax({
        url: "/v1/flow/" + bid.toString() + "/0",
        method: "POST",
        contentType: "application/json",
        dataType: "json",
        data: JSON.stringify({"cmd": "delete", "UserRefNo": UserRefNo}),
        error: function (data) {
            console.log(data);
        }
    });
};

//-----------------------------------------------------------------------------
// toggleNonFieldsErrorDisplay
//-----------------------------------------------------------------------------
window.toggleNonFieldsErrorDisplay = function () {

    $(".collapsible").toggleClass("active-collapsible");

    var content = $("#non-fields-error-content");
    if (content[0].style.display === "block") {
        content[0].style.display = "none";
    } else {
        content[0].style.display = "block";
        content[0].style.border = "1px solid red";
    }
};

//-----------------------------------------------------------------------------
// displayNonFieldsError
//-----------------------------------------------------------------------------
window.displayNonFieldsError = function (comp) {

    var errorString = "";
    if(app.raflow.validationErrors[comp]){
        $("#error-summary").css('display', 'block');

        // Display error count
        $("#field-errors-count").html(app.raflow.validationCheck.errors[comp].length);
        $("#non-field-errors-count").html(app.raflow.validationCheck.nonFieldsErrors[comp].length);

        for(var i = 0; i < app.raflow.validationCheck.nonFieldsErrors[comp].length; i++){
            console.debug(app.raflow.validationCheck.nonFieldsErrors[comp][i]);
            errorString += "<li>" + app.raflow.validationCheck.nonFieldsErrors[comp][i] + "</li>";
        }
    }else{
        $("#error-summary").css('display', 'none');
    }

    $("#non-fields-error-content").empty();
    $("#non-fields-error-content").append("<ul>" + errorString + "</ul>");

};