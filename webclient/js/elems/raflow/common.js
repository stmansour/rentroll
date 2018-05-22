/* global
    RACompConfig, sliderContentDivLength, reassignGridRecids,
    hideSliderContent, appendNewSlider, showSliderContentW2UIComp,
    loadTargetSection, requiredFieldsFulFilled, getRAFlowPartTypeIndex, initRAFlowAJAX,
    getRAFlowAllParts, saveActiveCompData, toggleHaveCheckBoxDisablity, getRAFlowPartData,
    lockOnGrid,
    loadRADatesForm
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
// Lock grid if checkbox is unchecked(false). Unlock grid if checkbox is checked(true).
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
        "Data": record
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
        }
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
        }
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
        }
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
    if (targetLoader.length > 0) {
        window[targetLoader]();
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
