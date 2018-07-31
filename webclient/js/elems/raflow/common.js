/* global
    RACompConfig, HideSliderContent, appendNewSlider, ShowSliderContentW2UIComp,
    loadTargetSection, requiredFieldsFulFilled, initRAFlowAjax,
    saveActiveCompData, getRAFlowCompData,
    lockOnGrid, dataFulFilled, getApprovals, updateFlowData, updateFlowCopy, displayErrorDot, initBizErrors
*/

"use strict";

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
        // Display error dot on each section if it have error
        // For business logic error only for now
        if(!(data.total > 0 && data.errortype === "biz")){
            return;
        }

        var FlowID = app.raflow.activeFlowID;

        app.raflow.bizErrors[FlowID] = {
            dates: data.errors.dates.total > 0 || data.nonFieldsErrors.dates.length > 0,
            people: data.errors.people.length > 0 || data.nonFieldsErrors.people.length > 0,
            pets: data.errors.pets.length > 0 || data.nonFieldsErrors.pets.length > 0,
            vehicles: data.errors.vehicle.length > 0 || data.nonFieldsErrors.vehicle.length > 0,
            rentables: data.errors.rentables.length > 0 || data.nonFieldsErrors.rentables.length > 0,
            parentchild: data.errors.parentchild.length > 0 || data.nonFieldsErrors.parentchild.length > 0,
            tie: data.errors.tie.people.length > 0 || data.nonFieldsErrors.tie.length > 0
        };

        displayErrorDot();

    });
});

// initBizErrors To initialize bizError local copy for active flow
window.initBizErrors = function(){

    var FlowID = app.raflow.activeFlowID;

    app.raflow.bizErrors[FlowID] = {
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
    var FlowID = app.raflow.activeFlowID;

    for (var comp in app.raFlowPartTypes) {
        if (app.raflow.bizErrors[FlowID][comp]) {
            $("#progressbar #steps-list li[data-target='#" + comp + "'] .error").addClass("error-true");
        } else {
            $("#progressbar #steps-list li[data-target='#" + comp + "'] .error").removeClass("error-true");
        }
    }
};

window.getApprovals = function(){

    var bid = getCurrentBID();
    var FlowID = app.raflow.activeFlowID;
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

            // Update bizcheck error local copy
            app.raflow.bizCheck[FlowID] = data;
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
window.getRAFlowCompData = function (compKey, FlowID) {

    var bid = getCurrentBID();

    var flowJSON = app.raflow.data[FlowID];
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
//   FlowID = for which FlowID's component
//   data   = data to set in the component
//-----------------------------------------------------------------------------
window.setRAFlowCompData = function (compKey, FlowID, data) {

    var bid = getCurrentBID();

    var flowJSON = app.raflow.data[FlowID];
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

    var bid = getCurrentBID();
    var FlowID = app.raflow.activeFlowID;

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
                console.log("data has been saved for: ", app.raflow.activeFlowID, ", compID: ", compID);
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
// getFlowDataAjax - get the ajax data from the server and returns ajax promise
//
// @params
//   FlowID = ID of the flow
//-----------------------------------------------------------------------------
window.getFlowDataAjax = function(FlowID) {
    var bid = getCurrentBID();

    return $.ajax({
        url: "/v1/flow/" + bid.toString() + "/" + FlowID.toString(),
        method: "POST",
        contentType: "application/json",
        dataType: "json",
        data: JSON.stringify({"cmd": "get", "FlowID": FlowID}),
        success: function (data) {
            if (data.status !== "error") {
                updateFlowData(data);
            }
        },
        error: function (data) {
            console.log(data);
        }
    });
};

// updateFlowData
window.updateFlowData = function(data){
    updateFlowCopy(data.record.Flow);
    setTimeout(function() {
        // Enable/Disable green check
        dataFulFilled(data.record);
    }, 500);
};

// updateFlowCopy
window.updateFlowCopy = function(flow){
    app.raflow.data[flow.FlowID] = flow;
};

// -----------------------------------------------------
// dataFulFilled:
// Enable/Disable green checks
// Enable/Disable get approvals button
// raflow parts
// -----------------------------------------------------
window.dataFulFilled = function(data) {

    // Update local copy of basicCheck and dataFulfilled
    app.raflow.basicCheck[data.Flow.FlowID] = data.BasicCheck;
    app.raflow.dataFulfilled[data.Flow.FlowID] = data.DataFulfilled;

    // Enable/Disable green check for the each section
    var active_comp = $(".ra-form-component:visible");
    var active_comp_id = active_comp.attr("id");

    for (var comp in app.raFlowPartTypes) {
        // if required fields are fulfilled then mark this slide as done

        // Apply green mark when comp is not active and when it fulfilled the requirements
        if (data.DataFulfilled[comp] && active_comp_id !== comp) {
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
    var compData = getRAFlowCompData(previousActiveCompID, app.raflow.activeFlowID);

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
// HideSliderContent - hide the slider and empty the content inside
//                     slider-content div
//
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
