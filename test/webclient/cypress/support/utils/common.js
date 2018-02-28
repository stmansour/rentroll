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

                // get update value for ARType from app variable : Account Rules
                if (w2uiGridColumn.field === "ARType"){
                    valueForCell = win.app.ARTypes[record[w2uiGridColumn.field]];
                }

                // get update value for ARType from app variable : Chart of accounts
                if (w2uiGridColumn.field === "Status"){

                    let statusID = record[w2uiGridColumn.field];
                    let statusList = win.app.account_stuff["statusList"];
                    let status = statusList.find(statusList => statusList.id === statusID);

                    valueForCell = status.text;
                }

                // get update value for ARType from app variable : Chart of accounts
                if (w2uiGridColumn.field === "AcctRule" || w2uiGridColumn.field === "Payor"){
                    if (valueForCell === null){
                        valueForCell = "";
                    }
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

    cy.log("******************");
    cy.log(formSelector);
    // div[name=receiptForm]
    cy.log("******************");


    debugger;
    // perform tests on form fields
    cy.get(formSelector).debug()
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

            //TODO(Akshay): Use switch statments

            // Update fieldValue for PmtTypeName
            if (fieldID === "PmtTypeName") {
                let pmtID = recordDetailFromAPIResponse.PMTID;
                let pmtTypes = win.app.pmtTypes[constants.testBiz];
                let pmtType = pmtTypes.find(pmtTypes => pmtTypes.PMTID === pmtID);

                fieldValue = pmtType.Name;
            }

            // Update fieldValue for Parents Account(PLID)
            if(fieldID === "PLID"){
                let plid = recordDetailFromAPIResponse.PLID;
                let parentAccountsRules = win.app.parent_accounts[constants.testBiz];
                let parentAccountsRule = parentAccountsRules.find(parentAccountsRules => parentAccountsRules.id === plid);

                fieldValue = parentAccountsRule.text;
            }

            // Update fieldValue for GL Account
            if(fieldID === "LID"){
                let lid = recordDetailFromAPIResponse.LID;
                let glAccountRules = win.app.gl_accounts[constants.testBiz];
                let glAccountRule = glAccountRules.find(glAccountRules => glAccountRules.id === lid)

                fieldValue = glAccountRule.text;
            }

            // Update fieldValue for Status
            if(fieldID === "Status"){
                let statusID = recordDetailFromAPIResponse.Status;
                let statusList = win.app.account_stuff["statusList"];
                let status = statusList.find(statusList => statusList.id === statusID);

                fieldValue = status.text;
            }

            // Update fieldValue for ARType : Account Rules
            if(fieldID === "ARType"){
                let arType = recordDetailFromAPIResponse.ARType;

                fieldValue = win.app.ARTypes[arType];
            }

            // Update fieldValue for DebitLID : Account Rules
            if(fieldID === "DebitLID"){
                let lid = recordDetailFromAPIResponse.DebitLID;
                let postAccountRules = win.app.post_accounts[constants.testBiz];
                let postAccountRule = postAccountRules.find(postAccountRules => postAccountRules.id === lid)

                fieldValue = postAccountRule.text;
            }

            // Update fieldValue for CreditLID : Account Rules
            if(fieldID === "CreditLID"){
                let lid = recordDetailFromAPIResponse.CreditLID;
                let postAccountRules = win.app.post_accounts[constants.testBiz];
                let postAccountRule = postAccountRules.find(postAccountRules => postAccountRules.id === lid)

                fieldValue = postAccountRule.text;
            }

            if(fieldID === "ARID"){
                let asmRuleID = recordDetailFromAPIResponse.ARID;
                let assessmentRules = win.app.AssessmentRules[constants.testBiz];
                let assessmentRule = assessmentRules.find(assessmentRules => assessmentRules.id === asmRuleID)

                fieldValue = assessmentRule.text;
            }

            // ERentableName
            if (fieldID === "ERentableName"){
                fieldValue = recordDetailFromAPIResponse.RentableName;
            }

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