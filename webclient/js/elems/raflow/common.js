/* global
    RACompConfig, HideSliderContent, appendNewSlider, ShowSliderContentW2UIComp,
    loadTargetSection, requiredFieldsFulFilled, initRAFlowAjax,
    saveActiveCompData, getRAFlowCompData,
    lockOnGrid, validateRAFlowComponents
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
                // update local data with server's response data
                app.raflow.data[data.record.Flow.FlowID] = data.record.Flow;
            } else {
                console.error(data.message);
            }
        },
        error: function (data) {
            console.log(data);
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
                app.raflow.data[data.record.Flow.FlowID] = data.record.Flow;
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
            if (data.status != "error") {
                app.raflow.data[data.record.Flow.FlowID] = data.record.Flow;
            }
        },
        error: function (data) {
            console.log(data);
        }
    });
};

//-----------------------------------------------------------------------------
// requiredFieldsFulFilled - checks whether all required fields for a component
//                           requested with compID is fulfilled or not
//
// @params
//   compID = ID of the component in the raflow
//-----------------------------------------------------------------------------
window.requiredFieldsFulFilled = function (compID) {
    var done = false;

    // if not active flow id then return
    if (app.raflow.activeFlowID === "") {
        console.log("no active flow ID");
        return done;
    }

    // get component data based on ID
    var compData = getRAFlowCompData(compID, app.raflow.activeFlowID);
    if(!compData) {
        return done;
    }

    // flag for validData, internally used
    var validData = true;

    switch (compID) {
        case "dates":
            var fields = ["AgreementStart", "AgreementStop", "RentStart",
                "RentStop", "PossessionStart", "PossessionStop"];

            // each field should not be black for green mark
            fields.forEach(function(field) {
                if (compData[field] === "") {
                    validData = false;
                    return false;
                }
            });

            // if loop passed successfully then mark it as successfully
            done = validData;
            break;

        case "people":
            compData.forEach(function(item) {
                if (item.IsRenter) {
                    done = true;
                    return false;
                }
            });
            break;

        case "pets":
            var havePetsChecked = $('#RAPetsGrid_checkbox')[0].checked;
            if(!havePetsChecked){
                done = true;
            } else {
                if (compData.length > 0) {
                    done = true;
                }else{
                    done = false;
                }
            }
            break;

        case "vehicles":
            var haveVehiclesChecked = $('#RAVehiclesGrid_checkbox')[0].checked;
            if(!haveVehiclesChecked){
                done = true;
            }else{
                if (compData.length > 0) {
                    done = true;
                }else{
                    done = false;
                }
            }
            break;

        // case "rentables":
        // Green checked by server side
        //     if (compData.length > 0) {
        //         done = true;
        //     }
        //     break;

        case "parentchild":
            // ==============================================================//
            //  ****************** VALIDATION SCENARIOS *********************//
            // ==============================================================//
            // 1.   If there are no child rentables then it is fine
            // Ex:  People only want to stay at apartment. They might
            //      no require child rentables like washing machine,
            //      car parking space etc...
            //
            // 2.   There must be at least one parent rentables in rentables
            //      section. People come to stay at rooms/apartments, so it
            //      doesn't make sense of not having any parent rentables.
            //
            // 3.   If any child rentables listed in rentables section then
            //      it must be associated with parent rentables.
            // Ex:  Washing machine (a child rentable) must be associated to
            //      an apartment(a parent rentable) where the people are
            //      living.
            // ==============================================================//

            // 1.   There must be at least one parent rentable and id of any
            //      item in this list must be > 0. If any item does not id > 0
            //      then don't mark green check.
            var pRExistForPC = false;
            app.raflow.parentRentableW2UIItems.forEach(function(item) {
                if (item.id > 0) {
                    pRExistForPC = true;
                }
            });
            if (!pRExistForPC) {
                done = false;
                break;
            }

            // If any child rentables listed then it should be associated
            // with any parent rentables.
            // If there are no child rentables then it is ok, so mark
            // validData as true.
            validData = true;
            compData.forEach(function(item) {
                if (item.PRID === 0 || item.CRID === 0) {
                    validData = false;
                    return false; // break the loop
                }
            });

            // done
            done = validData;
            break;

        case "tie":
            // ==============================================================//
            //  ****************** VALIDATION SCENARIOS *********************//
            // ==============================================================//
            // 1.   There must be at least one parent rentables in rentables
            //      section. People come to stay at rooms/apartments, so it
            //      doesn't make sense of not having any parent rentables.
            //
            // 2.   There must be at least one person with role of user.
            //      It doesn't make sense of not having any people at rooms/
            //      aprtments. At least one user must exists.
            //
            // 3.   If any user(occupant) listed in people section then
            //      it must be associated with parent rentables.
            // ==============================================================//

            // 1.   There must be at least one parent rentable and id of any
            //      item in this list must be > 0. If any item does not id > 0
            //      then don't mark green check.
            var pRExistForTie = false;
            app.raflow.parentRentableW2UIItems.forEach(function(item) {
                if (item.id > 0) {
                    pRExistForTie = true;
                }
            });
            if (!pRExistForTie) {
                done = false;
                break;
            }

            // 2. at least one occupant must exists
            var occupantExist = false;
            var peopleCompData = getRAFlowCompData("people", app.raflow.activeFlowID) || [];
            peopleCompData.forEach(function(peopleItem) {
                if (peopleItem.IsOccupant) {
                    occupantExist = true;
                    return false; // break the loop
                }
            });
            if (!occupantExist) {
                done = false;
                break;
            }

            // 3. Listed person must be associated with any parent rentables
            validData = true;
            compData.people.forEach(function(item) {
                if (item.PRID === 0 || item.TMPTCID === 0) {
                    validData = false;
                    return false; // break the loop
                }
            });

            // done
            done = validData;
            break;

        case "final":
            // TODO(Sudip): Validation rules for this slide
            break;
    }

    return done;
};

// -----------------------------------------------------
// validateRAFlowComponents:
// check validation rules for each component defined in
// raflow parts
// -----------------------------------------------------
window.validateRAFlowComponents = function() {
    for (var comp in app.raFlowPartTypes) {
        // if required fields are fulfilled then mark this slide as done
        if (requiredFieldsFulFilled(comp)) {
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

    // check validation for each slide in raflow parts
    validateRAFlowComponents();

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
            alert("invalid active comp: ", previousActiveCompID);
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
// getPetFees - will list down pet fees for a business
//-----------------------------------------------------------------------------
window.getPetFees = function () {
    var bid = getCurrentBID();

    return $.ajax({
        url: "/v1/petfees/" + bid.toString() + "/0",
        method: "GET",
        contentType: "application/json",
        dataType: "json",
        success: function (data) {
            if (data.status != "error") {
                app.petFees[bid] = data.records;
            }
        },
        error: function (data) {
            console.log(data);
        }
    });
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
