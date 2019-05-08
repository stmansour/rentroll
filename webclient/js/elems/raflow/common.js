/* global
    RACompConfig, HideSliderContent, appendNewSlider, ShowSliderContentW2UIComp, displayFormFieldsError,
    loadTargetSection, requiredFieldsFulFilled, InitRAFlowAjax, getRecIDFromTMPASMID, getFeeIndex,
    SaveCompDataAJAX, GetRAFlowCompLocalData, displayActiveComponentError, displayRAPetsGridError, dispalyRAPeopleGridError,
    lockOnGrid, GetApprovalsAJAX, UpdateRAFlowLocalData, displayErrorDot, initBizErrors,
    dispalyRARentablesGridError, dispalyRAVehiclesGridError, dispalyRAParentChildGridError, dispalyRATiePeopleGridError,
    GetCurrentFlowID, ReassignPeopleGridRecords, AssignPetsGridRecords, AssignVehiclesGridRecords, AssignRentableGridRecords,
    GetGridToolbarAddButtonID, HideRAFlowLoader, toggleNonFieldsErrorDisplay, displayErrorSummary, submitActionForm, displayGreenCircle,
    modifyFieldErrorMessage,ChangeRAFlowVersionToolbar, displayRADatesFormError, RAFlowAJAX, cleanFormError, loadRAActionTemplate,
    reloadActionForm, GetRefNoByRAIDFromGrid
*/

"use strict";

// FUNCTION FOR ELEMENT TO GIVE FLASH EFFECT
window.ElementFlash = function(el) {
    $(el).addClass("flash");
    setTimeout(function() {
        $(el).removeClass("flash");
    }, 500);
};

//-----------------------------------------------------------------------------
// RAFlowAJAX - A command ajax caller for all raflow related APIs
//              It will show loader before any request starts and
//              hides the loader when request is served
//-----------------------------------------------------------------------------
window.RAFlowAJAX = function(URL, METHOD, REQDATA, updateLocalData) {

    var DATA = null;
    if (METHOD === "POST") {
        DATA = JSON.stringify(REQDATA);
    }

    return $.ajax({
        url: URL,
        method: METHOD,
        contentType: "application/json",
        dataType: "json",
        data: DATA,
        beforeSend: function() {
            // show the loader
            HideRAFlowLoader(false);
        },
        success: function (data) {
            if (data.status !== "error") {
                if (updateLocalData) {
                    UpdateRAFlowLocalData(data);
                }
            } else {
                // alert(data.message);
                console.error(data.message);
            }
        },
        error: function (data) {
            console.error(data);
        },
        complete: function() {
            // hide the loader. GIVE UI SOME TIME TO RENDER
            setTimeout(function() {
                HideRAFlowLoader(true);
            }, 500);
        }
    });
};

//-----------------------------------------------------------------------------
// GetRefNoByRAIDFromGrid returns UserRefNo By RAID from raflowsGrid RECORDS
//-----------------------------------------------------------------------------
window.GetRefNoByRAIDFromGrid = function(RAID) {
    var RefNo = "";
    w2ui.raflowsGrid.records.forEach(function(gridRec) {
        if (gridRec.RAID == RAID) {
            RefNo = gridRec.UserRefNo;
            return;
        }
    });
    return RefNo;
};

//-----------------------------------------------------------------------------
// GetRAIDByRefNoFromGrid returns RAID By UserRefNo from raflowsGrid RECORDS
//-----------------------------------------------------------------------------
window.GetRAIDByRefNoFromGrid = function(RefNo) {
    var RAID = -1;
    w2ui.raflowsGrid.records.forEach(function(gridRec) {
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
    GetApprovalsAJAX().done(function (data) {

        if(data.status !== "success"){
            return;
        }

        app.raflow.validationErrors = {
            dates: data.record.ValidationCheck.errors.dates.total > 0 || data.record.ValidationCheck.nonFieldsErrors.dates.length > 0,
            people: data.record.ValidationCheck.errors.people.total > 0 || data.record.ValidationCheck.nonFieldsErrors.people.length > 0,
            pets: data.record.ValidationCheck.errors.pets.total > 0 || data.record.ValidationCheck.nonFieldsErrors.pets.length > 0,
            vehicles: data.record.ValidationCheck.errors.vehicles.total > 0 || data.record.ValidationCheck.nonFieldsErrors.vehicles.length > 0,
            rentables: data.record.ValidationCheck.errors.rentables.total > 0 || data.record.ValidationCheck.nonFieldsErrors.rentables.length > 0,
            parentchild: data.record.ValidationCheck.errors.parentchild.total > 0 || data.record.ValidationCheck.nonFieldsErrors.parentchild.length > 0,
            tie: data.record.ValidationCheck.errors.tie.people.total > 0 || data.record.ValidationCheck.nonFieldsErrors.tie.length > 0
        };

        displayErrorDot();

        displayActiveComponentError();

        // Display RAActionForm
        if(data.record.ValidationCheck.total === 0){

            if("raActionLayout" in w2ui){
                w2ui.raActionLayout.get('main').content = "";
            }

            loadRAActionTemplate();
            setTimeout(function() {
                reloadActionForm();
            },200);
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

//-----------------------------------------------------------------------------
// Get Approvals API AJAX CALL
//-----------------------------------------------------------------------------
window.GetApprovalsAJAX = function(){

    var BID = getCurrentBID();
    var FlowID = GetCurrentFlowID();

    var url = "/v1/validate-raflow/" + BID.toString() + "/" + FlowID.toString() + "/";
    var data = {
        "cmd": "get",
        "FlowID": FlowID
    };

    return RAFlowAJAX(url, "POST", data, true);
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
// GetRAFlowCompLocalData - get the flow component data stored locally in app.raflow
//
// @params
//   key    = flow component key
//   FlowID = for which FlowID's component
//-----------------------------------------------------------------------------
window.GetRAFlowCompLocalData = function (compKey) {

    var flowJSON = app.raflow.Flow;
    if (flowJSON.Data) {
        return flowJSON.Data[compKey];
    }

    return null;
};

//-----------------------------------------------------------------------------
// SetRAFlowCompLocalData - set the flow component data locally in app.raflow
//
// @params
//   key    = flow component key
//   data   = data to set in the component
//-----------------------------------------------------------------------------
window.SetRAFlowCompLocalData = function (compKey, data) {

    var flowJSON = app.raflow.Flow;
    if (flowJSON.Data) {
        flowJSON.Data[compKey] = data;
    }
};

//-----------------------------------------------------------------------------
// SaveCompDataAJAX - save component modified data on the server
//
// @params
//   compData   = modified latest component data
//   compID     = component key id
//-----------------------------------------------------------------------------
window.SaveCompDataAJAX = function (compData, compID) {

    // IF RAID VERSION THEN DON"T DO ANYTHING
    if (app.raflow.version === "raid") {
        return;
    }

    var BID = getCurrentBID();
    var FlowID = GetCurrentFlowID();

    var url = "/v1/flow/" + BID.toString() + "/" + FlowID.toString() + "/";
    var data = {
        "cmd": "save",
        "FlowType": app.raflow.name,
        "FlowID": FlowID,
        "FlowPartKey": compID,
        "BID": BID,
        "Data": compData
    };

    return RAFlowAJAX(url, "POST", data, true);
};

//-----------------------------------------------------------------------------
// InitRAFlowAjax - will initiate new rental agreement flow and returns ajax
//                  promise
//-----------------------------------------------------------------------------
window.InitRAFlowAjax = function () {
    var BID = getCurrentBID();

    var url = "/v1/flow/" + BID.toString() + "/0/";
    var data = {
        "cmd": "init",
        "FlowType": app.raflow.name
    };

    return RAFlowAJAX(url, "POST", data, true)
    .done(function(data) {
        if (data.status != "error") {
            // SINCE, WE'VE CREATED A BRAND NEW FLOW
            // RAFLOW VERSION MUST BE "REFNO"
            app.raflow.version = "refno";
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
    var BID = getCurrentBID();
    var FlowID = GetCurrentFlowID();

    var url = "/v1/flow/" + BID.toString() + "/" + FlowID.toString() + "/";
    var data = {
        "cmd":          "get",
        "UserRefNo":    UserRefNo,
        "RAID":         RAID,
        "Version":      version,
        "FlowType":     app.raflow.name
    };

    return RAFlowAJAX(url, "POST", data, true)
    .done(function(data) {
        if (data.status !== "error") {
            app.raflow.version = version;
        }
    });
};

// HideRAFlowLoader loader to show the progress while fetching data from the server
// which also disabled the controls in toolbar
window.HideRAFlowLoader = function(hide) {
    app.raflow.loading = !hide;
    if (hide) {
        if (w2ui.newraLayout) {
            $(w2ui.newraLayout.get("main").toolbar.box).find("button").prop('disabled', false);
        }
        $("#raflow-container .blocker").hide();
        $("#raactionform .blocker").hide();
    } else {
        if (w2ui.newraLayout) {
            $(w2ui.newraLayout.get("main").toolbar.box).find("button").prop('disabled', true);
        }
        $("#raflow-container .blocker").css("display", "flex");
        $("#raactionform .blocker").css("display", "flex");
        $("#raflow-container .blocker").show();
        $("#raactionform .blocker").show();
    }
};

// UpdateRAFlowLocalData updates the local data from the API response
window.UpdateRAFlowLocalData = function(data, reloadRequired){
    // catch RAID before app.raflow.Flow get updated
    var oldRAID = app.raflow.Flow.ID,
        newRAID = data.record.Flow.ID;

    app.raflow.Flow = data.record.Flow;

    // Update local copy of validation check
    app.raflow.validationCheck = data.record.ValidationCheck;

    // Update local copy of FlowFilledData
    app.raflow.FlowFilledData = data.record.DataFulfilled;

    // if RAID is not same the reload the grid listing
    if (reloadRequired && oldRAID && newRAID && newRAID !== oldRAID) {
        w2ui.raflowsGrid.reload();
    } else {
        // ALSO UPDATE THIS RAFLOW DATA(RAID/USERREFNO) IN THE MAIN GRID
        w2ui.raflowsGrid.records.forEach(function(gridRec) {
            if (gridRec.UserRefNo === app.raflow.Flow.UserRefNo || gridRec.RAID === app.raflow.Flow.ID) {
                if (app.raflow.Flow.UserRefNo) { // IF AVAILABLE THEN ONLY SET
                    gridRec.UserRefNo = app.raflow.Flow.UserRefNo;
                }
                if (app.raflow.Flow.ID) { // IF AVAILABLE THEN ONLY SET
                    gridRec.RAID = app.raflow.Flow.ID;
                }

                // ONCE THE RECORD UPDATE THEN ONLY REFRESH AND BREAK
                w2ui.raflowsGrid.refresh();
                return;
            }
        });
    }

    // UPDATE TOOLBAR
    if(!jQuery.isEmptyObject(app.raflow.Flow)) {
        // get info from local copy and refresh toolbar
        var VERSION = app.raflow.version,
            RAID    = app.raflow.Flow.ID,
            REFNO   = (RAID > 0) ? GetRefNoByRAIDFromGrid(RAID) : app.raflow.Flow.UserRefNo,
            FLAGS   = app.raflow.Flow.Data.meta.RAFLAGS;
        ChangeRAFlowVersionToolbar(VERSION,RAID,REFNO,FLAGS);
    }

    setTimeout(function() {
        // Enable/Disable green check
        displayGreenCircle();

        // Update error summary
        displayActiveComponentError();
    }, 500);
};

// -----------------------------------------------------
// displayGreenCircle
// -----------------------------------------------------
window.displayGreenCircle = function(){

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

    if (!target) {
        alert("no target provided to load target screen in the raflow");
        return false;
    }

    if (previousActiveCompID && previousActiveCompID !== target) {
        switch (previousActiveCompID) {
            case "dates":
                w2ui.RADatesForm.actions.reset();
                break;
            case "people":
                w2ui.RAPeopleGrid.clear();
                w2ui.RAPeopleSearchForm.actions.reset();
                break;
            case "pets":
                w2ui.RAPetsGrid.clear();
                w2ui.RAPetForm.actions.reset();
                break;
            case "vehicles":
                w2ui.RAVehiclesGrid.clear();
                w2ui.RAVehicleForm.actions.reset();
                break;
            case "rentables":
                w2ui.RARentablesGrid.clear();
                w2ui.RARentableFeesGrid.clear();
                w2ui.RARentableFeeForm.actions.reset();
                break;
            case "parentchild":
                w2ui.RAParentChildGrid.clear();
                break;
            case "tie":
                w2ui.RATiePeopleGrid.clear();
                break;
            case "final":
                w2ui.RAFinalRentablesFeesGrid.clear();
                w2ui.RAFinalPetsFeesGrid.clear();
                w2ui.RAFinalVehiclesFeesGrid.clear();
                break;
            default:
                alert("invalid active comp: " + previousActiveCompID);
                return;
        }

        // hide active component
        $("#progressbar #steps-list li[data-target='#" + previousActiveCompID + "']").removeClass("active");
        $(".ra-form-component#" + previousActiveCompID).hide();
    }

    // show target component
    $("#progressbar #steps-list li[data-target='#" + target + "']").removeClass("done").addClass("active");
    $(".ra-form-component#" + target).show();

    // display target comp fields summary
    displayErrorSummary(target);

    // display green circle based on datafulfilled flag
    displayGreenCircle();

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
    } else {
        console.error("unknown target from nav li: ", target);
        return false;
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
// displayActiveComponentError - it displays/highlight error for active component
//-----------------------------------------------------------------------------
window.displayActiveComponentError = function () {
    // get the current component (to be previous one)
    var active_comp = $(".ra-form-component:visible");

    // get active component id
    var active_comp_id = active_comp.attr("id");

    switch (active_comp_id) {
        case "dates":
            displayRADatesFormError();
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

    displayErrorSummary(active_comp_id);
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

    cleanFormError();

    // Iterate through fields with errors
    for(var key in records[index].errors){
        var field = $("[name=" + formName + "] input#" + key);
        var error = records[index].errors[key].join(", ");

        // Customize error for list input fields or if any other fields require
        var modifiedError = modifyFieldErrorMessage(key);
        if(modifiedError !== ""){
            error = modifiedError;
        }

        field.css("border-color", "red");
        field.after("<small class='error'>" + error + "</small>");
    }
};

// ---------------------------------------------------------------------
// modifyFieldErrorMessage - It modifies error message for key field
// ---------------------------------------------------------------------
window.modifyFieldErrorMessage = function(key){
    var error = "";
    switch (key){
        case "SourceSLSID":
            error = "please select a source";
            break;
        case "ARID":
            error = "please select an account rule";
            break;
        default:
            error = "";
    }
    return error;
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
        $(form.box).find("input,textarea").prop("disabled", true);
        $(form.box).find("button[class=w2ui-btn]").hide();
        $(form.box).find("div[class=w2ui-buttons]").hide();
   } else if (app.raflow.version === "refno") { // ENABLE ALL INPUTS & BUTTONS
        $(form.box).find("input,textarea").not("input[name=BUD]").prop("disabled", false);
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
    if (!UserRefNo) {
        alert("no such flow exists to delete");
        return;
    }

    var BID = getCurrentBID();
    var FlowID = GetCurrentFlowID();

    var url = "/v1/flow/" + BID.toString() + "/" + FlowID.toString() + "/";
    var data= {
        "cmd": "delete",
        "UserRefNo": UserRefNo
    };

    return RAFlowAJAX(url, "POST", data, false);
};

//----------------------------------------------------------------------------------
// It exapand/collapse non-field error summary section
//----------------------------------------------------------------------------------
$(document).on("click", "i#non-field-expandable-errors", function(event) {
    var target = event.target;

    var content = $("#non-fields-error-content");
    if (content[0].style.display === "block") {
        content[0].style.display = "none";
        $(target).removeClass("fa-caret-up").addClass("fa-caret-down");
    } else {
        content[0].style.display = "block";
        $(target).removeClass("fa-caret-down").addClass("fa-caret-up");
    }
});

//-----------------------------------------------------------------------------
// displayErrorSummary - It display error summary for active section.
//-----------------------------------------------------------------------------
window.displayErrorSummary = function (comp) {

    var error_summary_sel = "#error-summary";
    var non_field_error_dd_sel = "#error-info #non-field-expandable-errors";
    var non_field_error_content_sel = "#non-fields-error-content";


    if(app.raflow.validationErrors[comp]){
        // Display error summary
        $(error_summary_sel).css('display', 'block');

        var form_errors_count;
        if(comp !== "tie"){
            form_errors_count = app.raflow.validationCheck.errors[comp].total;
        }else{
            form_errors_count = app.raflow.validationCheck.errors[comp].people.total;
        }
        var non_fields_errors_count = app.raflow.validationCheck.nonFieldsErrors[comp].length;

        // Update error count for form error and non fields error
        $("#field-errors-count").html(form_errors_count);
        $("#non-field-errors-count").html(non_fields_errors_count);

        // If there are any non fields errors than display dropdown icon. Via it can expand non-fields-error summary
        if(non_fields_errors_count > 0){
            $(non_field_error_dd_sel).css('display', 'inline');

            var errorString = "";
            for(var i = 0; i < app.raflow.validationCheck.nonFieldsErrors[comp].length; i++){
                console.debug(app.raflow.validationCheck.nonFieldsErrors[comp][i]);
                errorString += "<li>" + app.raflow.validationCheck.nonFieldsErrors[comp][i] + "</li>";
            }

            // non fields error content
            $(non_field_error_content_sel).css('display', 'block');
            $(non_field_error_content_sel).empty();
            $(non_field_error_content_sel).append("<ul>" + errorString + "</ul>");
        }else{
            $(non_field_error_dd_sel).css('display', 'none');
            $(non_field_error_content_sel).css('display', 'none');
            $(non_field_error_content_sel).empty();
        }

    }
    else{
        // Hide error summary
        $(error_summary_sel).css('display', 'none');
    }
};

// cleanFormError It remove error small tag of current opened form if it have any
window.cleanFormError = function () {
    // Clean error
    $(".w2ui-form-box small.error").remove();
};
