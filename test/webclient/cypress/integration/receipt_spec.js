"use strict";

import * as selectors from '../support/utils/get_selectors';
import * as constants from '../support/utils/constants';
import * as common from '../support/utils/common';

const receiptsM = require('../support/components/receipts');

// default value of field in w2ui object
let defaultValue;
let fieldValue;

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

// get api end point for grid records
function getAPIEndPoint(module) {
    return [constants.API_VERSION, module, constants.BID].join("/");
}

// get api end point for grid record detail
function getDetailRecordAPIEndPoint(module, id) {
    return [constants.API_VERSION, module, constants.BID, id].join("/");
}

// -- Close the form. And assert that form isn't visible. --
function closeFormTests(formSelector) {

    // Close the form
    cy.get(selectors.getFormCloseButtonSelector()).click().wait(constants.WAIT_TIME);

    // Check that form should not visible after closing it
    cy.get(formSelector).should('not.be.visible');
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
function unallocatedSectionTest() {

    // Check visibility and class of
    cy.get(selectors.getAllocatedSectionSelector())
        .scrollIntoView()
        .should('be.visible')
        .should('have.class', 'FLAGReportContainer');

    // Check position of allocated section in detail form
    allocatedSectionPositionTest();
}

// -- perform tests on button --
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

// -- perform test on BUD field --
function BUDFieldTest() {
    // Check Business Unit field must be disabled and have value REX
    cy.get(selectors.getBUDSelector()).should('be.disabled').and('have.value', constants.testBiz).and('be.visible');
}

// -- perform test on grid cells --
function gridCellsTest(win) {
    // Iterate through each row
    recordsAPIResponse.forEach(function (record, rowNo) {

        // Iterate through each column in row
        w2uiGridColumns.forEach(function (w2uiGridColumn, columnNo) {

            // Skipping tests on skipColumns and
            // Perform test only if w2uiGridColumn isn't hidden
            if (!common.isInArray(w2uiGridColumn.field, testConfig.skipColumns) && !w2uiGridColumn.hidden) {

                // get defaultValue of cell from w2uiGrid
                let valueForCell = record[w2uiGridColumn.field];

                // format money type value
                if (w2uiGridColumn.render === "money") {
                    valueForCell = win.w2utils.formatters.money(valueForCell);
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

// -- perform test on add new record form's field --
function addNewFormTest(formName, formSelector) {

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

            // ERentableName field manipulation
            if (fieldID === "ERentableName"){
                defaultValue = getW2UIFormRecords.RentableName;
            }

            // Check field visibility and match default value from w2ui
            if (!common.isInArray(fieldID, testConfig.skipFields)) {

                // Check visibility and match the default value of the fields.
                cy.get(selectors.getFieldSelector(fieldID))
                    .should('be.visible')
                    .should('have.value', defaultValue);
            }

        });
}

// -- perform test on detail record form's field --
function detailFormTest(formSelector, formName, recordDetailFromAPIResponse, win) {
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

            // Update fieldValue for PmtTypeName
            if (fieldID === "PmtTypeName") {
                let pmtID = recordDetailFromAPIResponse.PMTID;
                let pmtTypes = appSettings.pmtTypes[constants.testBiz];
                let pmtType = pmtTypes.find(pmtTypes => pmtTypes.PMTID === pmtID);

                fieldValue = pmtType.Name;
            }

            // ERentableName
            if (fieldID === "ERentableName"){
                fieldValue = recordDetailFromAPIResponse.RentableName;
            }

            // check fields visibility and respective value
            if (!common.isInArray(fieldID, testConfig.skipFields)) {
                // Check visibility and match the default value of the fields.
                cy.get(selectors.getFieldSelector(fieldID))
                    .should('be.visible')
                    .should('have.value', fieldValue);
            }
        });
}

// test for print receipt ui in detail record form
function printReceiptUITest() {

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

    // We are not clicking print button. Because files get downloaded.
    // cy.get(selectors.getButtonSelector(printReceiptButtons[0])).click();

    // Close the popup
    cy.get(selectors.getClosePopupButtonSelector()).click();
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
            if (item.BUD === constants.testBiz) {
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

    // -- Check export CSV and export to Print button in grid toolbar --
    it('CSV and Print button in toolbar', function () {
        // Check visibility of export to CSV button
        cy.get(selectors.getExportCSVButtonSelector()).should('be.visible');

        // Check visibility of export to PDF button
        cy.get(selectors.getExportPDFButtonSelector()).should('be.visible');
    });


    /*****************************************************
     * Tests for node
     *******************************************************/
    it('Left side node', function () {

        // It should be visible and selected
        cy.get(selectors.getNodeSelector(testConfig.sidebarID))
            .scrollIntoView()
            .should('be.visible')
            .should('have.class', 'w2ui-selected')
            .click().wait(constants.WAIT_TIME);

    });

    /******************************
     * Change date in toolbar
     * 1. Route the receipt API
     * 2. Check http status
     * 3. Check key `status` in API endpoint responseBody
     *****************************/
    it('Change date', function () {

        // Starting a server to begin routing responses to cy.route()
        cy.server();

        // To manage the behavior of network requests. Routing the response for the requests.
        cy.route(testConfig.methodType, getAPIEndPoint(testConfig.sidebarID)).as('getRecords');

        // Select From date from W2UI calender
        cy.get('[name="receiptsD1"]').click().wait(constants.WAIT_TIME);
        cy.get('[class="w2ui-calendar-title title"]').click();
        cy.get('[class="w2ui-jump-month"][name=' + constants.fromMonth +']').click();
        cy.get('[class="w2ui-jump-year"][name=' + constants.fromYear + ']').click();
        cy.get('[date="' + constants.fromDate + '"]').click().wait(constants.WAIT_TIME);

        // Select To date from W2UI calender
        cy.get('[name="receiptsD2"]').click().wait(constants.WAIT_TIME);
        cy.get('[class="w2ui-calendar-title title"]').click();
        cy.get('[class="w2ui-jump-month"][name=' + constants.toMonth +']').click();
        cy.get('[class="w2ui-jump-year"][name=' + constants.toYear + ']').click().wait(constants.WAIT_TIME);
        cy.get('[date="' + constants.toDate + '"]').click();

        // Check http status
        cy.wait('@getRecords').its('status').should('eq', constants.HTTP_OK_STATUS);

        // get API endpoint's responseBody
        cy.get('@getRecords').then(function (xhr) {

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
     * 5. Check Unallocated section visibility and position(Using Y co-ordinate of elements)
     * 6. Check print receipt ui test
     * 7. Close the detail form
     * 8. Assert that form is close.
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

                // tests for grid cells visibility and value matching with api response records
                gridCellsTest(win);

                // ----------------------------------
                // -- Tests for detail record form --
                // ----------------------------------
                cy.log("Tests for detail record form");

                // -- detail record testing --
                const id = recordsAPIResponse[0][testConfig.primaryId];

                // Routing response to detail record's api requests.
                cy.route(testConfig.methodType, getDetailRecordAPIEndPoint(testConfig.module, id)).as('getDetailRecord');

                // click on the first record of grid
                cy.get(selectors.getFirstRecordInGridSelector(testConfig.grid)).click().wait(constants.WAIT_TIME);

                // check response status of API end point
                cy.wait('@getDetailRecord').its('status').should('eq', constants.HTTP_OK_STATUS);

                // perform tests on record detail form
                cy.get('@getDetailRecord').then(function (xhr) {

                    let recordDetailFromAPIResponse = xhr.response.body.record;

                    cy.log(recordDetailFromAPIResponse);

                    // formName
                    let formName = testConfig.form;

                    // get form selector
                    let formSelector = selectors.getFormSelector(formName);

                    detailFormTest(formSelector, formName, recordDetailFromAPIResponse, win);

                    // Check Business Unit field must be disabled and have value REX
                    BUDFieldTest();

                    // -- Check buttons visibility --
                    buttonsTest(testConfig.buttonNamesInDetailForm, testConfig.notVisibleButtonNamesInForm);

                    // -- Check Unallocated section's visibility and class --
                    unallocatedSectionTest();

                    // -- Check print receipt UI --
                    printReceiptUITest();

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

    it('Test for the Add New Button', function () {

        cy.contains('Add New', {force: true}).click().wait(constants.WAIT_TIME);

        // get form name
        let formName = testConfig.form;

        // get form selector
        let formSelector = selectors.getFormSelector(formName);

        addNewFormTest(formName, formSelector);

        // Check Business Unit field must be disabled and have value REX
        BUDFieldTest();

        // Check button's visibility
        buttonsTest(testConfig.buttonNamesInForm, testConfig.notVisibleButtonNamesInForm);

        // -- Close the form. And assert that form isn't visible. --
        closeFormTests(formSelector);

    });

});
