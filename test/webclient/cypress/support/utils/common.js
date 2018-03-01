"use strict";

import * as selectors from './get_selectors';
import * as constants from './constants';

// Check element's existence(value) in array
export function isInArray(value, array){
    return array.indexOf(value) > -1;
}

// get api end point for grid records
export function getAPIEndPoint(module) {
    return [constants.API_VERSION, module, constants.BID].join("/");
}

// get api end point for grid record detail
export function getDetailRecordAPIEndPoint(module, id) {
    return [constants.API_VERSION, module, constants.BID, id].join("/");
}

// -- perform test on BUD field --
export function BUDFieldTest() {
    // Check Business Unit field must be disabled and have value REX
    cy.get(selectors.getBUDSelector()).should('be.disabled').and('have.value', constants.testBiz).and('be.visible');
}

// -- perform tests on button --
export function buttonsTest(visibleButtons, notVisibleButtons) {

    // Check visibility of buttons
    visibleButtons.forEach(function (button) {
        // Check visibility of button
        cy.get(selectors.getButtonSelector(button)).should('be.visible');
    });

    // Check buttons aren't visible
    notVisibleButtons.forEach(function (button) {
        // Check button aren't visible
        cy.get(selectors.getButtonSelector(button)).should('not.be.visible');
    });
}

// -- Close the form. And assert that form isn't visible. --
export function closeFormTests(formSelector) {

    // Close the form
    cy.get(selectors.getFormCloseButtonSelector()).click().wait(constants.WAIT_TIME);

    // Check that form should not visible after closing it
    cy.get(formSelector).should('not.be.visible');
}

// -- perform test on grid cells --
export function gridCellsTest(recordsAPIResponse, w2uiGridColumns, win, testConfig) {
    let appSettings = win.app;

    // Iterate through each row
    recordsAPIResponse.forEach(function (record, rowNo) {

        // Iterate through each column in row
        w2uiGridColumns.forEach(function (w2uiGridColumn, columnNo) {

            // Skipping tests on skipColumns and
            // Perform test only if w2uiGridColumn isn't hidden
            if (!isInArray(w2uiGridColumn.field, testConfig.skipColumns) && !w2uiGridColumn.hidden) {

                // get defaultValue of cell from w2uiGrid
                let valueForCell = record[w2uiGridColumn.field];

                // format money type value
                if (w2uiGridColumn.render === "money") {
                    valueForCell = win.w2utils.formatters.money(valueForCell);
                }

                let types;
                let type;
                switch (w2uiGridColumn.field){
                    case "ARType":
                        // Account Rules
                        valueForCell = appSettings.ARTypes[valueForCell];
                        break;
                    case "Status":
                        types = appSettings.account_stuff["statusList"];
                        type = types.find(types => types.id === valueForCell);
                        valueForCell = type.text;
                        break;
                    case "AcctRule":
                    case "Payor":
                        // Chart of accounts
                        if (valueForCell === null){
                            valueForCell = "";
                        }
                        break;
                }

                // Check visibility and default value of cell in the grid
                cy.get(selectors.getCellSelector(testConfig.grid, rowNo, columnNo))
                    .scrollIntoView()
                    .should('be.visible')
                    .should('contain', valueForCell);
            }

        });

    });
}

// -- perform test on detail record form's field --
export function detailFormTest(formSelector, formName, recordDetailFromAPIResponse, win, testConfig) {
    console.log(recordDetailFromAPIResponse);
    let appSettings = win.app;

    let fieldValue;

    // field
    let field;

    // id of the field
    let fieldID;


    // record list in w2ui form
    let getW2UIFormRecords;

    // field list in w2ui form
    let getW2UIFormFields;

    // Check visibility of form
    cy.get(formSelector).should('be.visible');

    // get record and field list from the w2ui form object
    cy.window().then((win) => {
        // get w2ui form records
        getW2UIFormRecords = win.w2ui[formName].record;

        // get w2ui form fields
        getW2UIFormFields = win.w2ui[formName].fields;
    });

    // perform tests on form fields
    cy.get(formSelector)
        .find('input.w2ui-input:not(:hidden)') // get all input field from the form in DOM which doesn't have type as hidden
        .each(($el, index, $list) => {

            // get id of the field
            fieldID = $el.context.id;
            cy.log(fieldID);


            // get default value of field
            fieldValue = recordDetailFromAPIResponse[fieldID];
            cy.log(fieldValue);

            // get field from w2ui form field list
            field = getW2UIFormFields.find(fieldList => fieldList.field === fieldID);

            // Convert fieldValue to w2ui money type
            if (field.type === "money") {
                fieldValue = win.w2utils.formatters.money(recordDetailFromAPIResponse[fieldID]);
            }

            let types;
            let type;

            // Get fieldValue from the win.app variable
            switch(fieldID){
                case "PmtTypeName":
                    types = appSettings.pmtTypes[constants.testBiz];
                    type = types.find(types => types.PMTID === recordDetailFromAPIResponse.PMTID);
                    fieldValue = type.Name;
                    break;
                case "PLID":
                    // Parents Account(PLID)
                    types = appSettings.parent_accounts[constants.testBiz];
                    type = types.find(types => types.id === fieldValue);
                    fieldValue = type.text;
                    break;
                case "LID":
                    // GL Account:
                    types = appSettings.gl_accounts[constants.testBiz];
                    type = types.find(types => types.id === fieldValue);
                    fieldValue = type.text;
                    break;
                case  "Status":
                    // Chart of accounts
                    types = appSettings.account_stuff["statusList"];
                    type = types.find(types => types.id === fieldValue);
                    fieldValue = type.text;
                    break;
                case  "ARType":
                    // Account Rules
                    fieldValue = appSettings.ARTypes[fieldValue];
                    break;
                case  "DebitLID":
                case  "CreditLID":
                    // Account Rules
                    types = appSettings.post_accounts[constants.testBiz];
                    type = types.find(types => types.id === fieldValue);
                    fieldValue = type.text;
                    break;
                case  "ARID":
                    let ruleName;
                    if(formName === "asmEpochForm"){
                        ruleName = "AssessmentRules";
                    }else if (formName === "receiptForm"){
                        ruleName = "ReceiptRules";
                    }
                    types = appSettings[ruleName][constants.testBiz];
                    type = types.find(types => types.id === fieldValue);
                    fieldValue = type.text;
                    break;
                case "InvoiceNo": // Assess Charges form
                case "RAID": // Assess Charges form
                case "DID": // Tendered Payment Receipt
                    fieldValue = fieldValue.toString();

            }

            // ERentableName
            // if (fieldID === "ERentableName"){
            //     fieldValue = recordDetailFromAPIResponse.RentableName;
            // }

            // check fields visibility and respective value
            if (!isInArray(fieldID, testConfig.skipFields)) {
                // Check visibility and match the default value of the fields.
                cy.get(selectors.getFieldSelector(fieldID))
                    .should('be.visible')
                    .should('have.value', fieldValue);
            }
        });
}

// change date in UI from and to date
export function changeDate(dateFieldName, fromDt, toDt) {
    let fromMonth = fromDt.getMonth(); // Month : 0-11 : Jan-Dec
    let fromYear = fromDt.getFullYear(); // Year : Return 4 digits for 4-digit year
    let fromDay = fromDt.getDate(); // day/date: 1-31
    let fromDate = [fromMonth + 1, fromDay, fromYear].join('/'); // mm/dd/yyyy

    let toMonth = toDt.getMonth();
    let toYear = toDt.getFullYear();
    let toDay = toDt.getDate();
    let toDate = [toMonth + 1, toDay, toYear].join('/'); // mm/dd/yyyy


    // Select From date from W2UI calender
    cy.get('[name="' + dateFieldName + 'D1"]').click().wait(constants.WAIT_TIME);
    cy.get('[class="w2ui-calendar-title title"]').click();
    cy.get('[class="w2ui-jump-month"][name=' + fromMonth + ']').click();
    cy.get('[class="w2ui-jump-year"][name=' + fromYear + ']').click().wait(constants.WAIT_TIME);
    cy.get('[date="' + fromDate + '"]').click().wait(constants.WAIT_TIME);

    // Select To date from W2UI calender
    cy.get('[name="' + dateFieldName + 'D2"]').click().wait(constants.WAIT_TIME);
    cy.get('[class="w2ui-calendar-title title"]').click();
    cy.get('[class="w2ui-jump-month"][name=' + toMonth + ']').click();
    cy.get('[class="w2ui-jump-year"][name=' + toYear + ']').click().wait(constants.WAIT_TIME);
    cy.get('[date="' + toDate + '"]').click();
}