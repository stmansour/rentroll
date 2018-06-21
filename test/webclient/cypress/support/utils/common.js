"use strict";

import * as selectors from './get_selectors';
import * as constants from './constants';
import './../commands';

// Check element's existence(value) in array
export function isInArray(value, array) {
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

        if (testConfig.grid === "rrGrid") {
            switch (record.FLAGS) {
                // Normal row
                case 0:
                    testConfig.skipColumns = ["BeginReceivable", "DeltaReceivable", "EndReceivable", "BeginSecDep", "DeltaSecDep", "EndSecDep"];
                    break;
                // Main row
                case 1:
                    testConfig.skipColumns = ["BeginReceivable", "DeltaReceivable", "EndReceivable", "BeginSecDep", "DeltaSecDep", "EndSecDep"];
                    break;
                // Subtotal row
                case 2:
                    testConfig.skipColumns = [];
                    break;
                // Blank row
                case 4:
                    // Skipping tests on blank row
                    testConfig.skipColumns = [];
                    return;
            }
        }else if(testConfig.grid === "payorStmtDetailGrid"){
            switch (record.Description){
                case "*** RECEIPT SUMMARY ***":
                case "*** UNAPPLIED FUNDS ***":
                case "*** RENTAL AGREEMENT 1 ***":
                case "":
                    testConfig.skipColumns = ["Reverse", "spacer"];
                    return;
            }
        }

        // Iterate through each column in row
        w2uiGridColumns.forEach(function (w2uiGridColumn, columnNo) {

            // Skipping tests on skipColumns and
            // Perform test only if w2uiGridColumn isn't hidden
            if (!isInArray(w2uiGridColumn.field, testConfig.skipColumns) && !w2uiGridColumn.hidden) {

                // get defaultValue of cell from w2uiGrid
                let valueForCell = record[w2uiGridColumn.field];
                cy.log(w2uiGridColumn.field);
                cy.log(valueForCell);
                cy.log(record);

                if(valueForCell === null || valueForCell === undefined){
                    valueForCell = "";
                }

                // Format Value
                switch (w2uiGridColumn.render) {
                    // format money type value
                    case "money":
                        valueForCell = win.w2utils.formatters.money(valueForCell);
                        break;
                    // format float type value
                    case "float:2":
                        valueForCell = win.w2utils.formatters.float(valueForCell, 2);
                        break;
                }


                let types;
                let type;
                switch (w2uiGridColumn.field) {
                    case "ARType":
                        // Account Rules
                        valueForCell = appSettings.ARTypes[valueForCell];
                        break;
                    case "Status":
                        types = appSettings.account_stuff.statusList;
                        type = types.find(types => types.id === valueForCell);
                        valueForCell = type.text;
                        break;
                    case "AcctRule":
                    case "Payor":
                    case "Sqft":
                    case "SqFt":
                        // Chart of accounts
                        if (valueForCell === null) {
                            valueForCell = "";
                        }
                        break;
                    case "Description":
                        if (record.FLAGS === 1 && testConfig.grid === "rrGrid") {
                            if (valueForCell === null) {
                                valueForCell = "";
                            }
                        }
                        break;
                    case "AsmtAmount":
                    case "RcptAmount":
                        // --- stmtDetailGrid ---
                        if (testConfig.grid === "stmtDetailGrid") {
                            if (record.Descr !== "Closing Balance") {
                                valueForCell = "";
                            } else {
                                valueForCell = win.w2utils.formatters.float(valueForCell, 2);
                            }
                        }
                        break;
                    case "Balance": // --- stmtDetailGrid, payorStmtDetailGrid ---
                    case "UnappliedAmount": // --- payorStmtDetailGrid ---
                    case "AppliedAmount": // --- payorStmtDetailGrid ---
                    case "Assessment": // --- payorStmtDetailGrid ---
                        if (testConfig.grid !== "unpaidASMsGrid"){
                            if(valueForCell !== 0){
                                valueForCell = win.w2utils.formatters.float(valueForCell, 2);
                            }else {
                                valueForCell = "";
                            }
                        }
                        break;
                    case "RentableName":
                    case "RentableType":
                    case "Users":
                    case "Payors":
                    case "RAIDREP":
                    case "RAID":
                    case "RentCycleREP":
                    case "RentCycleGSR":
                        if (record.FLAGS === 2 && testConfig.grid === "rrGrid") {
                            if (valueForCell === null) {
                                valueForCell = "";
                            }
                        }
                        break;
                    case "UsePeriod":
                    case "RentPeriod":
                        if(testConfig.grid === "rrGrid"){
                            if (record.FLAGS === 2) {
                                valueForCell = "";
                            }else if(record.FLAGS === 1 && w2uiGridColumn.field === "UsePeriod"){
                                valueForCell = record.AgreementStart;
                            }else if(record.FLAGS === 1 && w2uiGridColumn.field === "RentPeriod"){
                                valueForCell = record.RentStart;
                            }
                        }
                        break;
                    case "RentCycle":
                    case "Proration":
                    case "GSRPC":
                    case "OverrideProrationCycle":
                    case "OverrideRentCycle":
                        cy.log(valueForCell);
                        valueForCell = appSettings.cycleFreq[valueForCell];
                        break;
                    case "ManageToBudget":
                        cy.log(valueForCell);
                        // refer /webclient/js/rt.js : rtGrid
                        if(valueForCell){
                          valueForCell = "YES (Market Rate required)";
                        }else{
                          valueForCell = "NO";
                        }
                        break;
                    case "UseStatus":
                        valueForCell = appSettings.RSUseStatus[valueForCell];
                        break;
                    case "LeaseStatus":
                        valueForCell = appSettings.RSLeaseStatus[valueForCell];
                        break;
                    case "EpochDue":
                    case "EpochPreDue":
                    case "DtPreDue":
                    case "DtPreDone":
                    case "DtDue":
                    case "DtDone":
                        valueForCell = win.dtFormatISOToW2ui(record[w2uiGridColumn.field]);
                        break;
                    
                }

                
                    cy.get(selectors.getCellSelector(testConfig.grid, rowNo, columnNo))
                        .scrollIntoView()
                        .should('be.visible')
                        .should('contain', valueForCell);
            }
        });

        // Scroll grid to left after performing tests on all columns of the grid
        if (testConfig.grid === "rrGrid" || testConfig.grid === "payorStmtDetailGrid") {
            cy.get(selectors.getGridRecordsSelector(testConfig.grid, rowNo)).scrollTo('left');
        }
    });
}

// -- perform test on detail record form's field --
export function detailFormTest(recordDetailFromAPIResponse, testConfig) {
    console.log(recordDetailFromAPIResponse);

    let fieldValue;

    // field
    let field;

    // id of the field
    let fieldID;

    // record list in w2ui form
    let getW2UIFormRecords;

    // field list in w2ui form
    let getW2UIFormFields;

    // formName
    let formName = testConfig.form;

    // get form selector
    let formSelector = selectors.getFormSelector(formName);

    // Check visibility of form
    cy.get(formSelector).should('be.visible');

    // get record and field list from the w2ui form object
    cy.window().then((win) => {
        // get w2ui form records
        getW2UIFormRecords = win.w2ui[formName].record;

        // get w2ui form fields
        getW2UIFormFields = win.w2ui[formName].fields;

        let appSettings = win.app;

        // perform tests on form fields
        cy.get(formSelector)
            .find('input.w2ui-input:not(:hidden)') // get all input field from the form in DOM which doesn't have type as hidden
            .each(($el, index, $list) => {

                // get id of the field
                fieldID = $el.context.id;
                cy.log($el);
                cy.log(fieldID);

                // get default value of field
                fieldValue = recordDetailFromAPIResponse[fieldID];
                cy.log(fieldValue);

                // get field from w2ui form field list
                field = getW2UIFormFields.find(fieldList => fieldList.field === fieldID);

                // Convert fieldValue to w2ui money type
                if (field.type === "money") {
                    fieldValue = win.w2utils.formatters.money(recordDetailFromAPIResponse[fieldID]);
                
                } else if(field.type === "datetime") {
                    fieldValue = win.dtFormatISOToW2ui(recordDetailFromAPIResponse[fieldID]);    
                } 

                let types;
                let type;

                // Get fieldValue from the win.app variable
                switch (fieldID) {
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
                        types = appSettings.account_stuff.statusList;
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
                        if (formName === "asmEpochForm") {
                            ruleName = "AssessmentRules";
                        } else if (formName === "receiptForm") {
                            ruleName = "ReceiptRules";
                        } else if (formName === "expenseForm") {
                            ruleName = "ExpenseRules";
                        } else if (formName === "rtForm") {
                            fieldValue = " -- No ARID -- ";
                            break;
                        }
                        types = appSettings[ruleName][constants.testBiz];
                        type = types.find(types => types.id === fieldValue);
                        fieldValue = type.text;
                        break;
                    case "InvoiceNo": // Assess Charges form
                    case "RAID": // Assess Charges form
                    case "DID": // Tendered Payment Receipt
                        fieldValue = fieldValue.toString();
                        break;
                    case "ERentableName":
                        fieldValue = recordDetailFromAPIResponse.RentableName;
                        break;
                    case "RentCycle":
                    case "ProrationCycle":
                    case "Proration":
                    case "GSRPC":
                        fieldValue = appSettings.cycleFreq[fieldValue];
                        break;
                    case "UseStatus":
                        fieldValue = appSettings.RSUseStatus[fieldValue];
                        break;
                    case "LeaseStatus":
                        fieldValue = appSettings.RSLeaseStatus[fieldValue];
                        break;
                    case "AssignmentTime":
                        types = appSettings.renewalMap;
                        type = types.find(types => types.id === fieldValue);
                        fieldValue = type.text;
                        break;
                    case "Cycle":
                        types = appSettings.w2ui.listItems.cycleFreq;
                        type = types.find(types => types.id === fieldValue);
                        fieldValue = type.text;
                        break;
                }

                // check fields visibility and respective value
                if (!isInArray(fieldID, testConfig.skipFields)) {
                    // Check visibility and match the default value of the fields.
                    switch (fieldID){
                        // Rentable Types form checkbox
                        case "ManageToBudget":
                        case "IsChildRentable":
                        // Account Rules form checkbox
                        case "ApplyRcvAccts":
                        case "RAIDrqd":
                        case "AutoPopulateToNewRA":
                        case "IsRentASM":
                        case "IsSecDepASM":
                        case "IsNonRecurCharge":
                        case "PriorToRAStop":
                        case "PriorToRAStart":
                        // Assessment Charges form checkbox
                        case "ExpandPastInst":
                        // Task List Definition form checkbox
                        case "ChkEpochPreDue":
                        case "ChkEpochDue":
                        // Task List form checkbox
                        case "ChkDtPreDue":
                        case "ChkDtPreDone":
                        case "ChkDtDue":
                        case "ChkDtDone":
                            if(fieldValue){
                                cy.get(selectors.getFieldSelector(fieldID))
                                    .should('be.visible')
                                    .should('be.checked');
                            }else{
                                cy.get(selectors.getFieldSelector(fieldID))
                                    .should('be.visible')
                                    .should('be.not.checked');
                            }
                            break;
                        default:
                            cy.get(selectors.getFieldSelector(fieldID))
                                .should('be.visible')
                                .should('have.value', fieldValue);
                            break;
                    }
                }
            });
    });

    if(testConfig.form !== "tldsInfoForm"){
        // Check Business Unit field must be disabled and have value REX
        BUDFieldTest();
    }

    // -- Check buttons visibility --
    buttonsTest(testConfig.buttonNamesInDetailForm, testConfig.notVisibleButtonNamesInForm);
}

// -- perform test on add new record form's field --
export function addNewFormTest(testConfig) {

    // record list in w2ui form
    let getW2UIFormRecords;

    // field list in w2ui form
    let getW2UIFormFields;

    // id of the field
    let fieldID;

    // field
    let field;

    // default value of field in w2ui object
    let defaultValue;

    // get form name
    let formName = testConfig.form;

    // get form selector
    let formSelector = selectors.getFormSelector(formName);

    // Check visibility of form
    cy.get(formSelector).should('be.visible');

    // get record and field list from the w2ui form object
    cy.window().then((win) => {

        // get w2ui form records
        getW2UIFormRecords = win.w2ui[formName].record;

        // get w2ui form fields
        getW2UIFormFields = win.w2ui[formName].fields;

    });

    cy.get(formSelector)
        .find('input.w2ui-input:not(:hidden)') // get all input field from the form in DOM which doesn't have type as hidden
        .each(($el, index, $list) => {

            // get id of the field
            fieldID = $el.context.id;

            cy.log(getW2UIFormRecords);

            // get default value of field
            defaultValue = getW2UIFormRecords[fieldID];

            // get field from w2ui form field list
            field = getW2UIFormFields.find(fieldList => fieldList.field === fieldID);

            // defaultValue type is object means it does have key value pair. get default text from the key value pair.
            if (typeof defaultValue === 'object') {
                defaultValue = defaultValue.text;
            }
            /* Money type field have default value in DOM is "$0.00".
                And w2ui field have value "0".
                To make the comparison change default value "0" to "$0.00" */
            else if (field.type === "money" && typeof defaultValue === 'number') {
                defaultValue = "$0.00";
            }
            else if (field.type === "datetime") {
                cy.window().then((win) => {
                    defaultValue = win.dtFormatISOToW2ui(defaultValue);
                });
            }

            // update defaultValue based on fieldID
            switch (fieldID) {
                case "InvoiceNo":
                case "DID":
                case "RAID":
                    defaultValue = defaultValue.toString();
                    break;
                case "ERentableName":
                    defaultValue = getW2UIFormRecords.RentableName;
                    break;
                case "ExpandPastInst":
                    if (defaultValue === true){
                        defaultValue = 'on';
                    }else {
                        defaultValue = 'off';
                    }
                    break;
            }

            // Check field visibility and match default value from w2ui
            if (!isInArray(fieldID, testConfig.skipFields)) {

                // Check visibility and match the default value of the fields.
                switch (fieldID){
                    // Rentable Types form checkbox
                    case "ManageToBudget":
                    case "IsChildRentable":
                    // Account Rules form checkbox
                    case "ApplyRcvAccts":
                    case "RAIDrqd":
                    case "AutoPopulateToNewRA":
                    case "IsRentASM":
                    case "IsSecDepASM":
                    case "IsNonRecurCharge":
                    case "PriorToRAStop":
                    case "PriorToRAStart":
                    // Assessment Charges form checkbox
                    case "ExpandPastInst":
                    // Task List Definition form checkbox
                    case "ChkEpochPreDue":
                    case "ChkEpochDue":
                        if(defaultValue){
                            cy.get(selectors.getFieldSelector(fieldID))
                                .should('be.visible')
                                .should('be.checked');
                        }else{
                            cy.get(selectors.getFieldSelector(fieldID))
                                .should('be.visible')
                                .should('be.not.checked');
                        }
                        break;
                    default:
                        cy.get(selectors.getFieldSelector(fieldID))
                            .should('be.visible')
                            .should('have.value', defaultValue);
                        break;
                }

            }

        });

    if(testConfig.form !== "tldsInfoForm"){
        // Check Business Unit field must be disabled and have value REX
        BUDFieldTest();
    }
    // Check button's visibility
    buttonsTest(testConfig.buttonNamesInDetailForm, testConfig.notVisibleButtonNamesInForm);

    // -- Close the form. And assert that form isn't visible. --
    closeFormTests(formSelector);
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

// change business unit as per the constants.testBiz
// return updated testBizID
export function changeBU(appSettings) {
    // get business id from appSettings variable for 'REX'
    appSettings.BizMap.forEach(function (item) {
        if (item.BUD === constants.testBiz) {
            constants.testBizID = item.BID;
        }
    });

    // Now change the business to REX
    cy.get('[name="BusinessSelect"]').select(constants.testBiz);

    // Check BusinessSelect value is set per the expected BID from appSettings variable
    cy.get('[name="BusinessSelect"]').should('have.value', constants.testBizID.toString());

    return constants.testBizID;
}

export function testAddNewRecordForm(testConfig) {
    cy.contains('Add New', {force: true}).click().wait(constants.WAIT_TIME);

    addNewFormTest(testConfig);
}

export function testMarketRulesDetailForm(testConfig) {

    // Open Market Rules tab
    cy.get('#tabs_rtDetailLayout_main_tabs_tab_rmrGrid').should('contain', 'Market Rates').click().wait(constants.WAIT_TIME);

    // Test on `Add New` and `Delete` button
    cy.get('#tb_rmrGrid_toolbar_item_w2ui-add').should('be.visible');
    cy.get('#tb_rmrGrid_toolbar_item_w2ui-delete').should('be.visible').should('have.class', 'disabled');


    // check response status of API end point
    cy.wait('@getMarketRulesRecords').its('status').should('eq', constants.HTTP_OK_STATUS);

    // perform tests on record detail form
    cy.get('@getMarketRulesRecords').then(function (xhr) {

        let recordsFromAPIResponse = xhr.response.body.records;
        cy.log(recordsFromAPIResponse);
        testConfig.grid = 'rmrGrid';
        if (recordsFromAPIResponse.length > 0) {
            testGridRecords(recordsFromAPIResponse, recordsFromAPIResponse.length, testConfig);

            // Click on first record and check delete button is enabled.
            cy.get(selectors.getFirstRecordInGridSelector('rmrGrid')).click();
            cy.get('#tb_rmrGrid_toolbar_item_w2ui-delete').should('be.visible').should('not.have.class', 'disabled');
        }
    });
}

export function testGridInTabbedDetailForm(gridName, layoutName, routeName, testConfig) {

    // Open Market Rules tab
    cy.get('#tabs_' + layoutName + '_main_tabs_tab_' + gridName).click().wait(constants.WAIT_TIME);
    //.should('contain', 'Market Rates')

    // Test on `Add New` and `Delete` button
    cy.get('#tb_' + gridName + '_toolbar_item_w2ui-add').should('be.visible');
    cy.get('#tb_' + gridName + '_toolbar_item_w2ui-delete').should('be.visible').should('have.class', 'disabled');


    // check response status of API end point
    cy.wait('@' + routeName).its('status').should('eq', constants.HTTP_OK_STATUS);

    // perform tests on record detail form
    cy.get('@' + routeName).then(function (xhr) {

        let recordsFromAPIResponse = xhr.response.body.records;
        cy.log(recordsFromAPIResponse);
        testConfig.grid = gridName;
        if (recordsFromAPIResponse.length > 0) {
            testGridRecords(recordsFromAPIResponse, recordsFromAPIResponse.length, testConfig);

            // Click on first record and check delete button is enabled.
            cy.get(selectors.getFirstRecordInGridSelector(gridName)).click();
            cy.get('#tb_' + gridName + '_toolbar_item_w2ui-delete').should('be.visible').should('not.have.class', 'disabled');
        }
    });
}

export function testDetailFormWithGrid(recordsAPIResponse, testConfig) {
    cy.log("Tests for detail record form");

    // -- detail record testing --
    const id = recordsAPIResponse[0][testConfig.primaryId];

    // Starting a server to begin routing responses to cy.route()
    cy.server();

    // Routing response to detail record's api requests.
    cy.route(testConfig.methodType, getDetailRecordAPIEndPoint(testConfig.module, id)).as('getDetailRecord');

    switch (testConfig.grid) {
        case "stmtGrid":
            cy.route(testConfig.methodType, getDetailRecordAPIEndPoint('stmtdetail', id)).as('getGridRecordsInDetailedForm');
            testConfig.skipColumns = ["Reverse", "dummy"];
            break;
        case "payorstmtGrid":
            cy.route(testConfig.methodType, getDetailRecordAPIEndPoint('payorstmt', id)).as('getGridRecordsInDetailedForm');
            testConfig.skipColumns = [];
            break;
        case "tldsGrid":
            cy.route(testConfig.methodType, getDetailRecordAPIEndPoint('tld', id)).as('getDetailRecord');
            cy.route(testConfig.methodType, getDetailRecordAPIEndPoint('tds', id)).as('getGridRecordsInDetailedForm');
            // testConfig.skipColumns = [];
            break;
        case "tlsGrid":
            cy.route(testConfig.methodType, getDetailRecordAPIEndPoint('tl', id)).as('getDetailRecord');
            cy.route(testConfig.methodType, getDetailRecordAPIEndPoint('tasks', id)).as('getGridRecordsInDetailedForm');
            // testConfig.skipColumns = [];
            break;
    }

    // click on the first record of grid
    cy.get(selectors.getFirstRecordInGridSelector(testConfig.grid)).click().wait(constants.WAIT_TIME);

    // perform tests on record detail form
    cy.get('@getDetailRecord').then(function (xhr) {
        let recordDetailFromAPIResponse = xhr.response.body.record;
        cy.log(recordDetailFromAPIResponse);


        if(testConfig.grid === "tldsGrid" || testConfig.grid === "tlsGrid"){
            detailFormTest(recordDetailFromAPIResponse,testConfig);
        } else {
            cy.get("#RAInfo").within(() => {
                if(testConfig.grid === "stmtGrid"){
                    cy.get('#bannerPayors').should('be.visible').should('contain', recordDetailFromAPIResponse.Payors);
                    cy.get('#RentalAgreementDates').should('be.visible').should('contain', recordDetailFromAPIResponse.AgreementStart).should('contain', recordDetailFromAPIResponse.AgreementStop);
                    cy.get('#PossessionDates').should('be.visible').should('contain', recordDetailFromAPIResponse.PossessionStart).should('contain', recordDetailFromAPIResponse.PossessionStop);
                    cy.get('#RentDates').should('be.visible').should('contain', recordDetailFromAPIResponse.RentStart).should('contain', recordDetailFromAPIResponse.RentStop);
                
                }else if(testConfig.grid === "payorstmtGrid"){
                    cy.get('#bannerTCID').should('be.visible').should('contain', recordDetailFromAPIResponse.FirstName).should('contain', recordDetailFromAPIResponse.MiddleName).should('contain', recordDetailFromAPIResponse.LastName);
                    // cy.get('#payorstmtaddr').should('be.visible').should('contain', recordDetailFromAPIResponse.Address); TODO(Akshay): Uncomment afterwards
                }
            });
        }
    });

    // check response status of API end point
    // cy.wait('@getGridRecordsInDetailedForm').its('status').should('eq', constants.HTTP_OK_STATUS);

    // perform tests on record detail form
    cy.get('@getGridRecordsInDetailedForm').then(function (xhr) {
        let recordDetailFromAPIResponse = xhr.response.body.records;
        if (recordDetailFromAPIResponse.length > 0) {
            testConfig.grid = testConfig.gridInForm;
            cy.log(testConfig.grid);
            testGridRecords(recordDetailFromAPIResponse, recordDetailFromAPIResponse.length, testConfig);
        }
        cy.log(recordDetailFromAPIResponse);

    });

}

export function testRecordDetailForm(recordsAPIResponse, testConfig) {
    cy.log("Tests for detail record form");

    // -- detail record testing --
    const id = recordsAPIResponse[0][testConfig.primaryId];

    // Starting a server to begin routing responses to cy.route()
    cy.server();

    // Routing response to detail record's api requests.
    cy.route(testConfig.methodType, getDetailRecordAPIEndPoint(testConfig.module, id)).as('getDetailRecord');

    switch (testConfig.module) {
        case "rt":
            cy.route(testConfig.methodType, getDetailRecordAPIEndPoint('rmr', id)).as('getMarketRulesRecords');
            break;
        case "rentable":
            cy.route(testConfig.methodType, getDetailRecordAPIEndPoint('rentablestatus', id)).as('getRentableStatusRecords');
            cy.route(testConfig.methodType, getDetailRecordAPIEndPoint('rentabletyperef', id)).as('getRentableTypeRef');
            break;
    }

    // Route rmr(Marketing Rates) endpoint while testing Rentable Types
    if (testConfig.module === "rt") {
        cy.route(testConfig.methodType, getDetailRecordAPIEndPoint('rmr', id)).as('getMarketRulesRecords');
    }


    // click on the first record of grid
    cy.get(selectors.getFirstRecordInGridSelector(testConfig.grid)).click().wait(constants.WAIT_TIME);

    // check response status of API end point
    cy.wait('@getDetailRecord').its('status').should('eq', constants.HTTP_OK_STATUS);

    // perform tests on record detail form
    cy.get('@getDetailRecord').then(function (xhr) {

        let recordDetailFromAPIResponse = xhr.response.body.record;

        cy.log(recordDetailFromAPIResponse);

        detailFormTest(recordDetailFromAPIResponse, testConfig);

    });

}

export function testGridRecords(recordsAPIResponse, noRecordsInAPIResponse, testConfig) {

    // list of columns from the grid
    let w2uiGridColumns;


    /**********************************************************
     * Tests for grid records
     * 1. Iterate through each row
     * 2. Check visibility of cell in the row
     * 3. Check value of cells in the row
     **********************************************************/

    // Check visibility of grid
    cy.get(selectors.getGridSelector(testConfig.grid)).should('be.visible').wait(constants.WAIT_TIME);

    // get length from the window and perform tests
    cy.window()
        .then(win => {

            // get list of columns in the grid
            w2uiGridColumns = win.w2ui[testConfig.grid].columns;
            // Match grid record length with total rows in receiptsGrid
            cy.get(selectors.getRowsInGridSelector(testConfig.grid)).should(($trs) => {
                expect($trs).to.have.length(noRecordsInAPIResponse);
            });

            // Perform test only if there is/are record(s) exists in API response.
            if (noRecordsInAPIResponse > 0) {
                // tests for grid cells visibility and value matching with api response records
                gridCellsTest(recordsAPIResponse, w2uiGridColumns, win, testConfig);
            }
        });
}

// Check position of allocated section in detail form
function allocatedSectionPositionTest() {

    // get co-ordinate of allocated section
    const allocatedSection = Cypress.$(selectors.getAllocatedSectionSelector()).get(0).getBoundingClientRect();

    // get co-ordinate of button section
    const buttonSection = Cypress.$('.w2ui-buttons').get(0).getBoundingClientRect();

    // get difference of y co-ordinate of element
    let sectionDiff = allocatedSection.y - buttonSection.y;

    // Check difference must be 1
    // expect(sectionDiff).to.equal(1);
}

// -- Check Unallocated section's visibility and class --
export function unallocatedSectionTest() {

    // Check visibility and class of
    cy.get(selectors.getAllocatedSectionSelector())
        .scrollIntoView()
        .should('be.visible')
        .should('have.class', 'FLAGReportContainer');

    // Check position of allocated section in detail form
    // allocatedSectionPositionTest();
}

// test for print receipt ui in detail record form
export function printReceiptUITest() {

    // Open print receipt UI
    cy.get(selectors.getFormPrintButtonSelector()).should('be.visible').click();

    // Check print receipt pop up should open
    cy.get(selectors.getPrintReceiptPopUpSelector()).should('be.visible').wait(constants.WAIT_TIME);

    // Check format list visibility
    cy.get(selectors.getPrintReceiptPopUpSelector())
        .find('.w2ui-field-helper').should('be.visible');

    // Check default permanent_resident radio button is checked
    cy.get(selectors.getPermanentResidentRadioButtonSelector())
        .should('be.visible')
        .should('be.checked');

    // Check hotel radio button is unchecked
    cy.get(selectors.getHotelRadioButtonSelector())
        .should('be.visible')
        .should('not.be.checked');

    // Check button visibility
    let printReceiptButtons = ["print", "close"];
    buttonsTest(printReceiptButtons, []);

    // Close the popup
    cy.get(selectors.getClosePopupButtonSelector()).click();
}

// test for Saving a new Record in Form
export function testSaveNewRecord(testConfig) {
    // record list in w2ui form
    let getW2UIFormRecords;

    // field list in w2ui form
    let getW2UIFormFields;

    // get form name
    let formName = testConfig.form;

    // get form selector
    let formSelector = selectors.getFormSelector(formName);

    // get record and field list from the w2ui form object
    cy.window().then((win) => {

        // get w2ui form records
        getW2UIFormRecords = win.w2ui[formName].record;
        cy.log(getW2UIFormRecords);
        // get w2ui form fields
        getW2UIFormFields = win.w2ui[formName].fields;

    });

    let fieldID;
    let field;
    let fieldValue;

    testConfig.skipFields.push('BUD');

    if(testConfig.form == "rtForm") {
        testConfig.skipFields.push('ManageToBudget');        
    }
    
    cy.fixture(testConfig.fixtureFile).then((json) => {

        cy.get(formSelector)
            .find('input.w2ui-input:not(:hidden)') // get all input field from the form in DOM which doesn't have type as hidden
            .each(($el, index, $list) => {

                // get id of the field
                fieldID = $el.context.id;

                cy.log(getW2UIFormRecords);

                // get default value of field
                fieldValue = json.record[fieldID];

                // get field from w2ui form field list
                field = getW2UIFormFields.find(fieldList => fieldList.field === fieldID);

                // Check field visibility and match default value from w2ui
                if (!isInArray(fieldID, testConfig.skipFields)) {

                    // Check if type of input field is list
                    if($list[index].getAttribute("type") === "list") {

                        // Get dropdown field, check visiblity and click on it
                        cy.get(selectors.getFieldSelector(fieldID)).parent().should('be.visible').click();
                        // Get dropdown value, check visiblity and click on it
                        cy.get(selectors.getDropDownValueFieldSelector(fieldValue)).should('be.visible').click();
                    }
                    else if($list[index].getAttribute("type") === "checkbox") {
                        if(!$el[0].checked){
                            if(!$el.context.disabled){
                                // Get checkbox, check visiblity and click on it
                                cy.get(selectors.getFieldSelector(fieldID))
                                    .should('be.visible')
                                    .should('not.be.disabled')
                                    .should('not.be.checked').click();
                            }
                        }
                    }
                    else if($list[index].getAttribute("type") === "text") {

                        // Check visibility and match the default value of the fields.
                        cy.get(selectors.getFieldSelector(fieldID))
                            .should('be.visible').clear().type(fieldValue)
                            .should('have.value', fieldValue);
                    }
                }
            });
    });

    // Route request for adding new record
    cy.server();
    cy.route(testConfig.methodType, getDetailRecordAPIEndPoint(testConfig.module, 0)).as('addRecord');

    // Get save button and click on it
    cy.get(selectors.getButtonSelector('save')).click();

    // check response status of API end point
    cy.wait('@addRecord').its('status').should('eq', constants.HTTP_OK_STATUS);

    // get API Endpoint response
    cy.get('@addRecord').then(function (xhr) {

        // Check status flag in API Endpoint response
        expect(xhr.responseBody).to.have.property('status', constants.API_RESPONSE_SUCCESS_FLAG);
    });
}