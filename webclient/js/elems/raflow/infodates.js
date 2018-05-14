/* global
    RACompConfig, sliderContentDivLength, reassignGridRecids,
    hideSliderContent, appendNewSlider, showSliderContentW2UIComp,
    loadTargetSection, requiredFieldsFulFilled, getRAFlowPartTypeIndex, initRAFlowAJAX,
    getRAFlowAllParts, saveActiveCompData, toggleHaveCheckBoxDisablity, getRAFlowPartData,
    lockOnGrid,
*/

/* exported loadRADatesForm */

"use strict";

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