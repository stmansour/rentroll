"use strict";

import * as constants from '../support/utils/constants';
import * as selectors from '../support/utils/get_selectors';
import * as common from '../support/utils/common';

// --- Setup --
const section = require('../support/components/pmTypes'); // Payment types

// this contain app variable of the application
let appSettings;

// holds the test configuration for the modules
let testConfig;

// -- Start Cypress UI tests for AIR Roller Application --
describe('AIR Roller UI Tests - Payment types', function () {

    // // records list of module from the API response
    let recordsAPIResponse;

    let noRecordsInAPIResponse;

    // -- Perform operation before all tests starts. It runs once before all tests in the block --
    /********************************
    * Login into application
    * Select node from left sidebar
    * Route the response for grid records
    *
    * Expect:
    * Grid records response must have status flag as success.
    ********************************/
    before(function () {

        testConfig = section.conf;

        // --- Login into Application before starting any tests ---
        // Check custom login command for more detail. File path: ./../support/commands.js
        cy.login();

        cy.visit(constants.ROLLER_APPLICATION_PATH).wait(constants.PAGE_LOAD_TIME);

        // Starting a server to begin routing responses to cy.route()
        cy.server();

        // To manage the behavior of network requests. Routing the response for the requests.
        cy.route(testConfig.methodType, common.getAPIEndPoint(testConfig.sidebarID)).as('getRecords');

        /************************
         * Select right side node
         *************************/
        // Node should be visible and selected
        cy.get(selectors.getNodeSelector(testConfig.sidebarID))
            .scrollIntoView()
            .should('be.visible')
            .click().wait(constants.WAIT_TIME)
            .should('have.class', 'w2ui-selected');

        // If have date navigation bar than change from and to Date to get in between data
        if (testConfig.haveDateValue) {
            common.changeDate(testConfig.sidebarID, testConfig.fromDate, testConfig.toDate);
        }

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

    // -- Perform operation before each test(it()) starts. It runs before each test in the block. --
    beforeEach(function () {

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
        // onSuccessful test set BID value. If above test get fail below code will not be executed.
        constants.BID = common.changeBU(appSettings);
    });

    /***********************
    * Iterate through each cell.
    *
    * Expect:
    * Cell value must be same as record's field value from API Response.
    ***********************/
    it('Grid Records', function () {
        common.testGridRecords(recordsAPIResponse, noRecordsInAPIResponse, testConfig);
    });

    /*******************************
    * Click on first record of grid
    *
    * Expect:
    * Each field must have value set same as detail record api response.
    * Button must be visible(Save, Cancel etc.)
    *
    *
    * Close the form
    ********************************/
    it('Record Detail Form', function () {
        // ----------------------------------
        // -- Tests for detail record form --
        // ----------------------------------
        // Params:
        // recordsAPIResponse: list of record from the api response,
        // testConfig: configuration for running tests
        common.testRecordDetailForm(recordsAPIResponse, testConfig);

        // -- Close the form. And assert that form isn't visible. --
        common.closeFormTests(selectors.getFormSelector(testConfig.form));
    });

    /************************************************************
    * Click Add new in toolbar
    *
    * Expect:
    * Each field must set to be its default value
    * Button must be visible(Save, Save and Add Another etc.)
    ************************************************************/
    it('Check default value of fields for new record form', function () {
        // ---------------------------------------
        // ----- Tests for add new record form ---
        // ---------------------------------------
       common.testAddNewRecordForm(testConfig);
    });

    /**************************************************
     * Click Add new bitton in toolbar
     * Fill value in the forms for each field from the fixture
     * Click save button
     *
     * Expect:
     * After saving the record, response must have status flag to be 'success'
     **************************************************/
    //TODO(Akshay): Use this Save new record tests for other modules also later on
    it('Save new record', function () {
        // Click add new button and open a form
        cy.contains('Add New', {force: true}).click().wait(constants.WAIT_TIME);

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

            // get w2ui form fields
            getW2UIFormFields = win.w2ui[formName].fields;

        });

        let fieldID;
        let field;
        let fieldValue;

        testConfig.skipFields = ['BUD'];

        cy.fixture('paymentTypes.json').then((json) => {

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
                    if (!common.isInArray(fieldID, testConfig.skipFields)) {

                        // Check visibility and match the default value of the fields.
                        cy.get(selectors.getFieldSelector(fieldID))
                            .should('be.visible').type(fieldValue)
                            .should('have.value', fieldValue);
                    }

                });
        });


        // Route request for adding new record
        cy.server();
        cy.route(testConfig.methodType, common.getDetailRecordAPIEndPoint(testConfig.module, 0)).as('addRecord');

        // Get save button and click on it
        cy.get(selectors.getButtonSelector('save')).click();

        // check response status of API end point
        cy.wait('@addRecord').its('status').should('eq', constants.HTTP_OK_STATUS);

        // get API Endpoint response
        cy.get('@addRecord').then(function (xhr) {

            // Check status flag in API Endpoint response
            expect(xhr.responseBody).to.have.property('status', constants.API_RESPONSE_SUCCESS_FLAG);
        });
    });

    // -- Perform operation after all tests finish. It runs once after all tests in the block --
    after(function () {

        // --- Logout from the Application after finishing all tests ---
        // Check custom login command for more detail. File path: ./../support/commands.js
        cy.logout();
    });
});