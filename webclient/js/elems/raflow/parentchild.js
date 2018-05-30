/* global
    RACompConfig, sliderContentDivLength, reassignGridRecids,
    hideSliderContent, appendNewSlider, showSliderContentW2UIComp,
    loadTargetSection, requiredFieldsFulFilled, getRAFlowPartTypeIndex, initRAFlowAjax,
    getRAFlowAllParts, saveActiveCompData, toggleHaveCheckBoxDisablity, getRAFlowCompData,
    lockOnGrid,
*/

"use strict";

//-----------------------------------------------------------------------
// loadRAPeopleChildSection -
//-----------------------------------------------------------------------
window.loadRAPeopleChildSection = function () {

    if (!("RAParentChildGrid" in w2ui)) {

        // rentables grid
        $().w2grid({
            name: 'RAParentChildGrid',
            header: 'Parent Child Rentables Relation',
            show: {
                toolbar:    false,
                footer:     true,
            },
            multiSelect: false,
            style: 'display: block;',
            columns: [
                {
                    field: 'recid',
                    hidden: true,
                },
                {
                    field: 'BID',
                    hidden: true
                },
                {
                    field: 'CRID',
                    hidden: true
                },
                {
                    field: 'PRID',
                    hidden: true
                },
                {
                    field: 'ChildRentableName',
                    caption: 'Rentable',
                    size: '40%',
                },
                {
                    field: 'ParentRentableName',
                    caption: 'Parent Rentable',
                    size: '60%',
                },
            ]
        });
    }

    // now load grid in division
    $('#ra-form #parentchild .grid-container').w2render(w2ui.RAParentChildGrid);

    // load the existing data in rentables component
    setTimeout(function () {
        var compData = getRAFlowCompData("parentchild", app.raflow.activeFlowID);
        var grid = w2ui.RAParentChildGrid;

        if (compData) {
            grid.records = compData;
            reassignGridRecids(grid.name);
        } else {
            grid.clear();
        }
    }, 500);
};
