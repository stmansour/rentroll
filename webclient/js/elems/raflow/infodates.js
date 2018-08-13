/* global
    getRAFlowCompData, EnableDisableRAFlowVersionInputs,
    HideAllSliderContent, SetlocalDataFromRADatesFormRecord,
    SetRADatesFormRecordFromLocalData, setRAFlowCompData,
    SetFormRecordFromData, SetDataFromFormRecord, SaveDatesCompData, displayRADatesFormError
*/

"use strict";

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
            formURL: '/webclient/html/raflow/formra-dates.html',
            fields: [
                {name: 'AgreementStart',    type: 'date',   required: true, html: {caption: "Term Start"}},
                {name: 'AgreementStop',     type: 'date',   required: true, html: {caption: "Term Stop"}},
                {name: 'RentStart',         type: 'date',   required: true, html: {caption: "Rent Start"}},
                {name: 'RentStop',          type: 'date',   required: true, html: {caption: "Rent Stop"}},
                {name: 'PossessionStart',   type: 'date',   required: true, html: {caption: "Possession Start"}},
                {name: 'PossessionStop',    type: 'date',   required: true, html: {caption: "Possession Stop"}},
                {name: 'CSAgent',           type: 'int',    required: true, html: {caption: "CS Agent"}}
            ],
            actions: {
                reset: function () {
                    w2ui.RADatesForm.clear();
                },
                save: function() {
                    var form = w2ui.RADatesForm;

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // validate form record
                    var errors = form.validate();
                    if (errors.length > 0) {
                        console.error("error in form validation on save action");
                        console.error(errors);
                        return;
                    }

                    // update the modified data
                    SetlocalDataFromRADatesFormRecord();

                    // save data on server side
                    SaveDatesCompData()
                    .done(function(data) {
                        if (data.status !== "success") {
                            form.message(data.message);
                        } else {
                            form.refresh();
                        }
                    });
                }
            },
            onRefresh: function (event) {
                var form = this;
                event.onComplete = function() {
                    var t = new Date(),
                    nyd = new Date(new Date().setFullYear(new Date().getFullYear() + 1));

                    // set default values with start=current day, stop=next year day, if record is blank
                   form.record.AgreementStart =form.record.AgreementStart || w2uiDateControlString(t);
                   form.record.AgreementStop =form.record.AgreementStop || w2uiDateControlString(nyd);
                   form.record.RentStart =form.record.RentStart || w2uiDateControlString(t);
                   form.record.RentStop =form.record.RentStop || w2uiDateControlString(nyd);
                   form.record.PossessionStart =form.record.PossessionStart || w2uiDateControlString(t);
                   form.record.PossessionStop =form.record.PossessionStop || w2uiDateControlString(nyd);

                   // FREEZE THE INPUTS IF VERSION IS RAID
                   EnableDisableRAFlowVersionInputs(form);
                };
            }
        });
    }

    // now render the form in specifiec targeted division
    $('#ra-form #dates .form-container').w2render(w2ui.RADatesForm);
    HideAllSliderContent();

    // load the existing data in dates component
    setTimeout(function () {
        SetRADatesFormRecordFromLocalData();
        displayRADatesFormError();
    }, 0);
};

// displayRADatesFormError If form field have error than it highlight with red border and
window.displayRADatesFormError = function(){

    // if pet section doesn't have error than return
    if(!app.raflow.validationErrors.dates){
        return;
    }

    // get list of pets
    var dates = app.raflow.validationCheck.errors.dates;

    // Iterate through fields with errors
    for(var key in dates.errors){
        var field = $("[name=RADatesForm] input#" + key);
        var error = dates.errors[key].join(", ");
        field.css("border-color", "red");
        field.after("<small class='error'>" + error + "</small>");
    }
};

// -------------------------------------------------------------
// SetlocalDataFromRADatesFormRecord
// ==================================
// will update the data from the record
// it will only update the field defined in fields list in
// form definition
// -------------------------------------------------------------
window.SetlocalDataFromRADatesFormRecord = function() {
    var form            = w2ui.RADatesForm;

    // get local data
    var localDatesData = getRAFlowCompData("dates");

    // set data from form
    // keep ID is 1 to set only records in defined fields
    var datesData = SetDataFromFormRecord(1, form, localDatesData);

    // set this modified data back
    setRAFlowCompData("dates", datesData);
};

// -------------------------------------------------------------
// SetRADatesFormRecordFromLocalData
// ================================
// will set the data in the form record
// from local vehicle data
// -------------------------------------------------------------
window.SetRADatesFormRecordFromLocalData = function() {
    var form = w2ui.RADatesForm;

    // get local data
    var localDatesData = getRAFlowCompData("dates");

    // set form record from data
    SetFormRecordFromData(form, localDatesData);

    // refresh the form after setting the record
    form.refresh();
    form.refresh();
};

//------------------------------------------------------------------------------
// SaveDatesCompData - saves the data on server side
//------------------------------------------------------------------------------
window.SaveDatesCompData = function() {
    var compData = getRAFlowCompData("dates");
    return saveActiveCompData(compData, "dates");
};

