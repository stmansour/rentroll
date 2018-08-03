/* global
    getRAFlowCompData, EnableDisableRAFlowVersionInputs,
    HideAllSliderContent
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
                {name: 'recid',             type: 'int',    required: true, html: {page: 0, column: 0}},
                {name: 'BID',               type: 'int',    required: true, html: {page: 0, column: 0}},
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
        var compData = getRAFlowCompData("dates");

        if (compData) {
            w2ui.RADatesForm.record = compData;
            w2ui.RADatesForm.refresh();
        } else {
            w2ui.RADatesForm.clear();
        }
    }, 500);
};
