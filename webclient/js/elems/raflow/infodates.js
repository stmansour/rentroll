/* global
    RACompConfig, sliderContentDivLength, reassignGridRecids,
    hideSliderContent, appendNewSlider, showSliderContentW2UIComp,
    loadTargetSection, requiredFieldsFulFilled, initRAFlowAjax,
    saveActiveCompData, toggleHaveCheckBoxDisablity, getRAFlowCompData,
    lockOnGrid,
*/

/* exported loadRADatesForm */

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
            formURL: '/webclient/html/formradates.html',
            fields: [
                {name: 'recid',             type: 'int',    required: true, html: {page: 0, column: 0}},
                {name: 'BID',               type: 'int',    required: true, html: {page: 0, column: 0}},
                {name: 'AgreementStart',    type: 'date',   required: true, html: {caption: "Term Start"}},
                {name: 'AgreementStop',     type: 'date',   required: true, html: {caption: "Term Stop"}},
                {name: 'RentStart',         type: 'date',   required: true, html: {caption: "Rent Start"}},
                {name: 'RentStop',          type: 'date',   required: true, html: {caption: "Rent Stop"}},
                {name: 'PossessionStart',   type: 'date',   required: true, html: {caption: "Possession Start"}},
                {name: 'PossessionStop',    type: 'date',   required: true, html: {caption: "Possession Stop"}}
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
        var compData = getRAFlowCompData("dates", app.raflow.activeFlowID);

        if (compData) {
            w2ui.RADatesForm.record = compData;
            w2ui.RADatesForm.refresh();
        } else {
            w2ui.RADatesForm.clear();
        }
    }, 500);
};