/* global
    RACompConfig, sliderContentDivLength, reassignGridRecids,
    hideSliderContent, appendNewSlider, showSliderContentW2UIComp,
    loadTargetSection, requiredFieldsFulFilled, getRAFlowPartTypeIndex, initRAFlowAJAX,
    getRAFlowAllParts, saveActiveCompData, toggleHaveCheckBoxDisablity, getRAFlowPartData,
    lockOnGrid,
    getFeesTermsGridInitalRecord
*/

"use strict";

// -------------------------------------------------------------------------------
// Rental Agreement - Fees Terms Grid
// -------------------------------------------------------------------------------
window.getFeesTermsGridInitalRecord = function (BID, gridLen) {
    return {
        recid: gridLen,
        RID: 0,
        BID: BID,
        RTID: 0,
        RentableName: "",
        FeeName: "",
        Amount: 0.0,
        Cycle: 6,
        SigningAmt: 0.0,
        ProrateAmt: 0.0,
        TaxableAmt: 0.0,
        SalesTax: 0.0,
        TransOCC: 0.0,
    };
};

window.loadRAFeesTermsGrid = function () {

    var partType = app.raFlowPartTypes.feesterms;
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
    if (!("RAFeesTermsGrid" in w2ui)) {

        // feesterms grid
        $().w2grid({
            name: 'RAFeesTermsGrid',
            header: 'FeesTerms',
            show: {
                toolbar: true,
                footer: true,
            },
            style: 'border: 1px solid black; display: block;',
            toolbar: {
                items: [
                    {id: 'add', type: 'button', caption: 'Add Record', icon: 'w2ui-icon-plus'}
                ],
                onClick: function (event) {
                    var bid = getCurrentBID();
                    if (event.target == 'add') {
                        var inital = getFeesTermsGridInitalRecord(bid, w2ui.RAFeesTermsGrid.records.length);
                        w2ui.RAFeesTermsGrid.add(inital);
                    }
                }
            },
            columns: [
                {
                    field: 'recid',
                    hidden: true,
                },
                {
                    field: 'RID',
                    hidden: true
                },
                {
                    field: 'BID',
                    hidden: true
                },
                {
                    field: 'RTID',
                    hidden: true
                },
                {
                    field: 'RentableName',
                    caption: 'Rentable',
                    size: '180px',
                    editable: {type: 'text'}
                },
                {
                    field: 'FeeName',
                    caption: 'Fee',
                    size: '120px',
                    editable: {type: 'text'}
                },
                {
                    field: 'Amount',
                    caption: 'Amount',
                    size: '80px',
                    render: 'money',
                    editable: {type: 'money'}
                },
                {
                    field: 'Cycle',
                    caption: 'Cycle',
                    size: '80px',
                    editable: {type: 'int'}
                },
                {
                    field: 'SigningAmt',
                    caption: 'At Signing',
                    size: '80px',
                    render: 'money',
                    editable: {type: 'money'}
                },
                {
                    field: 'ProrateAmt',
                    caption: 'Prorate',
                    size: '80px',
                    render: 'money',
                    editable: {type: 'money'}
                },
                {
                    field: 'TaxableAmt',
                    caption: 'Taxable Amt',
                    size: '80px',
                    render: 'money',
                    editable: {type: 'money'}
                },
                {
                    field: 'SalesTax',
                    caption: 'Sales Tax',
                    size: '80px',
                    render: 'money',
                    editable: {type: 'money'}
                },
                {
                    field: 'TransOCC',
                    caption: 'Trans OCC',
                    size: '80px',
                    render: 'money',
                    editable: {type: 'money'}
                }
            ],
            onChange: function (event) {
                event.onComplete = function () {
                    this.save();
                };
            }
        });
    }

    // load grid in division
    $('#ra-form #feesterms').w2render(w2ui.RAFeesTermsGrid);

    // load the existing data in feesterms component
    setTimeout(function () {
        var i = getRAFlowPartTypeIndex(app.raFlowPartTypes.feesterms);
        if (i >= 0 && app.raflow.data[app.raflow.activeFlowID][i].Data) {
            w2ui.RAFeesTermsGrid.records = app.raflow.data[app.raflow.activeFlowID][i].Data;
            w2ui.RAFeesTermsGrid.refresh();
        } else {
            w2ui.RAFeesTermsGrid.clear();
        }
    }, 500);
};