"use strict";

import * as selectors from '../support/utils/get_selectors';
import * as constants from '../support/utils/constants';
import * as common from '../support/utils/common';

const receiptsM = require('../support/components/receipts');

// default value of field in w2ui object
let defaultValue;

// field
let field;

// id of the field
let fieldID;

// record list in w2ui form
let getW2UIFormRecords;

// field list in w2ui form
let getW2UIFormFields;

// list of columns from the grid
let w2uiGridColumns;

// records list of module from the API response
let recordsAPIResponse;

// number of records in API response
let noRecordsInAPIResponse;

// this contain app variable of the application
let appSettings;

// holds the test configuration for the modules
let testConfig;

function getAPIEndPoint(module) {
    return constants.API_VERSION + "/" + module + "/" + constants.BID;
}

// -- Close the form. And assert that form isn't visible. --
function closeFormTests(formSelector) {

    // Close the form
    cy.get(selectors.getFormCloseButtonSelector()).click().wait(constants.WAIT_TIME);

    // Check that form should not visible after closing it
    cy.get(formSelector).should('not.be.visible');
}

// -- Check Unallocated section's visibility and class --
function unallocatedSectionTest() {
    cy.get(selectors.getUnallocateSectionSelector())
        .scrollIntoView()
        .should('be.visible')
        .should('have.class', 'FLAGReportContainer');
}


function buttonsTest(visibleButtons, notVisibleButtons) {

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

function BUDFieldTest() {
    // Check Business Unit field must be disabled and have value REX
    cy.get(selectors.getBUDSelector()).should('be.disabled').and('have.value', constants.testBiz).and('be.visible');
}

// -- Start Cypress UI tests --
describe('AIR Receipt UI Tests', function () {

    // -- Perform operation before all tests starts. It runs once before all tests in the block --
    before(function () {
        /*
        * Clear cookies befor starting tests. Because We are preserving cookies to use it all test suit.
        * Running test suit multiple times require new session to login into application.
        */
        cy.clearCookie(constants.APPLICATION_COOKIE);

        testConfig = receiptsM.conf;

    });

    /**********************************
     * Assert the title of application
     * 1. Visit Receipt application
     * 2. Assert the title 'AIR Receipts'
     ************************************/

    it('Assert the title of application', function () {

        // It visit baseUrl(from cypress.json) + applicationPath
        cy.visit(constants.RECEIPT_APPLICATION_PATH).wait(constants.PAGE_LOAD_TIME);

        // Assert application title
        cy.title().should('include', 'AIR Receipts');

    });

    // -- Login into application --
    it('Login into AIR Receipts', function () {
        // Check custom login command for more detail. File path: ./../support/commands.js
        cy.login();
    });

    // -- Perform operation before each test(it()) starts. It runs before each test in the block. --
    beforeEach(function () {
        /*
        * Cypress automatically clears all cookies before each test run.
        * It does make application log off.
        * To preserve cookies for entire test suit add that cookie in Cypress cookies's whitelist.
        * Link for more detail: https://docs.cypress.io/api/cypress-api/cookies.html
        */
        Cypress.Cookies.defaults({whitelist: constants.APPLICATION_COOKIE});

        // -- get app variable from the window --
        /*
        * After successfully login into application it will have fixed app variable.
        * Fetching it after successful login.
        * */
        cy.window().then((win) => {
            appSettings = win.app;
        });

        cy.log(appSettings);

    });

    // -- Change business to REX --
    it('Change business to REX', function () {

        // get business id from appSettings variable for 'REX'
        appSettings.BizMap.forEach(function (item) {
            if (item.BUD === constants.testBiz){
                constants.testBizID = item.BID;
            }
        });

        // Now change the business to REX
        cy.get('[name="BusinessSelect"]').select(constants.testBiz);

        // Check BusinessSelect value is set per the expected BID from appSettings variable
        cy.get('[name="BusinessSelect"]').should('have.value', constants.testBizID.toString());

        // onSuccessful test set BID value. If above test get fail below code will not be executed.
        constants.BID = constants.testBizID;

    });


    /*****************************************************
     * Tests for node
     * 1. Route the receipt API
     * 2. Check http status
     * 3. Check key `status` in API endpoint responseBody
     *******************************************************/
    it('Left side node', function () {

        // Starting a server to begin routing responses to cy.route()
        cy.server();

        // To manage the behavior of network requests. Routing the response for the requests.
        cy.route(testConfig.methodType, getAPIEndPoint(testConfig.sidebarID)).as('getReceipts');

        // It should be visible and selected
        cy.get(selectors.getNodeSelector(testConfig.sidebarID))
            .scrollIntoView()
            .should('be.visible')
            .should('have.class', 'w2ui-selected')
            .click().wait(constants.WAIT_TIME);

        // Check http status
        cy.wait('@getReceipts').its('status').should('eq', constants.HTTP_OK_STATUS);

        // get API endpoint's responseBody
        cy.get('@getReceipts').then(function (xhr) {

            // Check key `status` in responseBody
            expect(xhr.responseBody).to.have.property('status', constants.API_RESPONSE_SUCCESS_FLAG);

            // get records list from the API response
            recordsAPIResponse = xhr.response.body.records;

            // -- Assigning number of records to 0 if no records are available in response --
            if (recordsAPIResponse) {
                noRecordsInAPIResponse = xhr.response.body.records.length;
            } else {
                noRecordsInAPIResponse = 0;
            }

        });

    });

    /**********************************************************
     * Tests for grid records
     * 1. Iterate through each row
     * 2. Check visibility of cell in the row
     * 3. Check value of cells in the row
     *
     * Tests for grid record detail
     * 1. Click on first record
     * 2. Check visibility of detail form
     * 3. Check visibility and value of the fields
     * 4. Check button's visibility
     * 5. Check Unallocated section visibility and position(CSS Class)
     * 6. Close the detail form
     * 7. Assert that form is close.
     *
     *********************************************************/

    it('Test for grid records and record detail', function () {

        // Starting a server to begin routing responses to cy.route()
        cy.server();

        // ----------------------------
        // -- Tests for grid records --
        // ----------------------------
        cy.log("Tests for grid records");

        // Check visibility of grid
        cy.get(selectors.getGridSelector(testConfig.grid)).should('be.visible').wait(constants.WAIT_TIME);

        // get length from the window and perform tests
        cy.window().then(win => {

            // get list of columns in the grid
            w2uiGridColumns = win.w2ui[testConfig.grid].columns;

            // Match grid record length with total rows in receiptsGrid
            cy.get(selectors.getRowsInGridSelector(testConfig.grid)).should(($trs) => {
                expect($trs).to.have.length(noRecordsInAPIResponse);
            });

            // Perform test only if there is/are record(s) exists in API response.
            if (noRecordsInAPIResponse > 0) {

                // Iterate through each row
                recordsAPIResponse.forEach(function (record, rowNo) {

                    // Iterate through each column in row
                    w2uiGridColumns.forEach(function (w2uiGridColumn, columnNo) {

                        // Skipping tests on skipColumns
                        if (!common.isInArray(w2uiGridColumn.field, testConfig.skipColumns)) {

                            // Perform test only if w2uiGridColumn isn't hidden
                            if (!w2uiGridColumn.hidden) {

                                // get defaultValue of cell from w2uiGrid
                                let valueForCell = record[w2uiGridColumn.field];

                                // Check visibility and default value of cell in the grid
                                cy.get(selectors.getCellSelector(testConfig.grid, rowNo, columnNo))
                                    .scrollIntoView()
                                    .should('be.visible')
                                    .should('contain', valueForCell);
                            }
                        }

                    });

                });


                // ----------------------------------
                // -- Tests for detail record form --
                // ----------------------------------
                cy.log("Tests for detail record form");

                // -- detail record testing --
                const id = recordsAPIResponse[0][testConfig.primaryId];

                // Routing response to detail record's api requests.
                cy.route(testConfig.methodType, '/v1/receipt/1/' + id).as('getDetailRecord');

                // click on the first record of grid
                cy.get(selectors.getFirstRecordInGridSelector(testConfig.grid)).click().wait(constants.WAIT_TIME);

                // check response status of API end point
                cy.wait('@getDetailRecord').its('status').should('eq', constants.HTTP_OK_STATUS);

                // perform tests on record detail form
                cy.get('@getDetailRecord').then(function (xhr) {

                    let recordDetailFromAPIResponse = xhr.response.body.record;

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

                    });


                    // perform tests on form fields
                    cy.get(formSelector)
                        .find('input.w2ui-input:not(:hidden)') // get all input field from the form in DOM which doesn't have type as hidden
                        .each(($el, index, $list) => {

                            // get id of the field
                            fieldID = $el.context.id;

                            // get default value of field
                            defaultValue = recordDetailFromAPIResponse[fieldID];

                            // get field from w2ui form field list
                            field = getW2UIFormFields.find(fieldList => fieldList.field === fieldID);

                            /*
                            * TODO(Akshay): Handle PmtTypeName from app.pmtType. Render Amount in w2ui money type.
                            * */
                            if (fieldID !== "PmtTypeName" && fieldID !== "Amount" && fieldID !== "ERentableName") {

                                // Check visibility and match the default value of the fields.
                                cy.get(selectors.getFieldSelector(fieldID))
                                    .should('be.visible')
                                    .should('have.value', defaultValue);
                            }
                        });

                    // Check Business Unit field must be disabled and have value REX
                    BUDFieldTest();

                    buttonsTest(testConfig.buttonNamesInDetailForm, testConfig.notVisibleButtonNamesInForm);

                    // -- Check Unallocated section's visibility and class --
                    unallocatedSectionTest();

                    // -- Close the form. And assert that form isn't visible. --
                    closeFormTests(formSelector);

                });
            }

        });
    });


    /************************************************************
     * Tests for adding new record form.
     * 1. Open the form by clicking the Add New button in toolbar
     * 2. Check visibility of the form in viewport
     * 3. Find not hidden field from the DOM
     * 4. Perform visibility test on those hidden fields
     * 5. Check default value for that fields.
     ***********************************************************/
    //TODO(Akshay): Do common code for formfields
    // it('Test for the Add New Button', function () {
    //
    //     cy.contains('Add New', {force: true}).click().wait(constants.WAIT_TIME);
    //
    //     let formName = testConfig.form;
    //
    //     // get form selector
    //     let formSelector = selectors.getFormSelector(formName);
    //
    //     // Check visibility of form
    //     cy.get(formSelector).should('be.visible');
    //
    //     // get record and field list from the w2ui form object
    //     cy.window().then((win) => {
    //
    //         // get w2ui form records
    //         getW2UIFormRecords = win.w2ui[formName].record;
    //
    //         // get w2ui form fields
    //         getW2UIFormFields = win.w2ui[formName].fields;
    //
    //     });
    //
    //     cy.get(formSelector)
    //         .find('input.w2ui-input:not(:hidden)') // get all input field from the form in DOM which doesn't have type as hidden
    //         .each(($el, index, $list) => {
    //
    //             // get id of the field
    //             fieldID = $el.context.id;
    //
    //             // get default value of field
    //             defaultValue = getW2UIFormRecords[fieldID];
    //
    //             // get field from w2ui form field list
    //             field = getW2UIFormFields.find(fieldList => fieldList.field === fieldID);
    //
    //             // defaultValue type is object means it does have key value pair. get default text from the key value pair.
    //             if (typeof defaultValue === 'object') {
    //                 defaultValue = defaultValue.text;
    //             }
    //             /* Money type field have default value in DOM is "$0.00".
    //             And w2ui field have value "0".
    //             To make the comparison change default value "0" to "$0.00" */
    //             else if (field.type === "money" && typeof defaultValue === 'number') {
    //                 defaultValue = "$0.00";
    //             }
    //
    //             /* Skipping tests for Resident Address field. Because it have default value as 'undefined' and in DOM it have value as ''.
    //             Which makes the test fail.
    //             TODO(Sudip): Change default value undefine to ''.
    //             TODO(Akshay): Remove `if` condition for the tests after an issue has been resolved.
    //             */
    //             if (fieldID !== "ERentableName") {
    //                 // Check visibility and match the default value of the fields.
    //                 cy.get('#' + fieldID)
    //                     .should('be.visible')
    //                     .should('have.value', defaultValue);
    //             }
    //
    //         });
    //
    //     // Check Business Unit field must be disabled and have value REX
    //     BUDFieldTest();
    //
    //     // Check button's visibility
    //     buttonsTest(testConfig.buttonNamesInForm, testConfig.notVisibleButtonNamesInForm)
    //
    //     // -- Close the form. And assert that form isn't visible. --
    //     closeFormTests(formSelector);
    //
    // });

});

// describe('AIR Roller UI Tests', function () {
//
// });
