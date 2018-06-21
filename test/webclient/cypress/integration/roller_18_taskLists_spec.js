"use strict";

import * as constants from '../support/utils/constants';
import * as selectors from '../support/utils/get_selectors';
import * as common from '../support/utils/common';

// --- Collections ---
const section = require('../support/components/taskLists'); // Task Lists

// this contain app variable of the application
let appSettings;

// holds the test configuration for the modules
let testConfig;

// -- Start Cypress UI tests for AIR Roller Application --
describe('AIR Roller UI Tests - Task Lists', function () {

    // // records list of module from the API response
    let recordsAPIResponse;

    let noRecordsInAPIResponse;

    // -- Perform operation before all tests starts. It runs once before all tests in the block --
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

    /************************************************************
    * Click Add new in toolbar
    *
    * Expect:
    * Each field must set to be its default value
    * Button must be visible(Save, Save and Add Another etc.)
    ************************************************************/
    it('Record Detail Form', function () {
        // ----------------------------------
        // -- Tests for detail record form --
        // ----------------------------------
        
        if(noRecordsInAPIResponse >0){
            // Params:
            // recordsAPIResponse: list of record from the api response,
            // testConfig: configuration for running tests
            common.testDetailFormWithGrid(recordsAPIResponse, testConfig);

            // -- Close the form. And assert that form isn't visible. --
            common.closeFormTests(selectors.getFormSelector(testConfig.form));
        }
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
        testConfig.buttonNamesInDetailForm.splice( testConfig.buttonNamesInDetailForm.indexOf('delete'), 1 );

        cy.contains('Add New', {force: true}).click().wait(constants.WAIT_TIME);

        // record list in w2ui form
        let getW2UIFormRecords;

        // id of the field
        let fieldID;

        // default value of field in w2ui object
        let defaultValue;

        // get form name
        let formName = testConfig.formInPopUp;

        // get form selector
        let formSelector = selectors.getFormSelector(formName);

        // Check visibility of form
        cy.get(formSelector).should('be.visible');

        // get record and field list from the w2ui form object
        cy.window().then((win) => {

            // get w2ui form records
            getW2UIFormRecords = win.w2ui[formName].record;
        });

        cy.get(formSelector)
            .find('input.w2ui-input:not(:hidden)') // get all input field from the form in DOM which doesn't have type as hidden
            .each(($el, index, $list) => {

                // get id of the field
                fieldID = $el.context.id;

                cy.log(getW2UIFormRecords);

                // get default value of field
                defaultValue = getW2UIFormRecords[fieldID];

                // defaultValue type is object means it does have key value pair. get default text from the key value pair.
                if (typeof defaultValue === 'object') {
                    defaultValue = defaultValue.text;
                }

                cy.get(selectors.getFieldSelector(fieldID))
                    .should('be.visible')
                    .should('have.value', defaultValue);
            });
        // Check button's visibility
        common.buttonsTest(testConfig.buttonNamesInDetailForm, testConfig.notVisibleButtonNamesInForm);

        // Close the form
        cy.get(selectors.getClosePopupButtonSelector()).click().wait(constants.WAIT_TIME);

        // Check that form should not visible after closing it
        cy.get(formSelector).should('not.be.visible');
    });

    // -- Perform operation after all tests finish. It runs once after all tests in the block --
    after(function () {

        // --- Logout from the Application after finishing all tests ---
        // Check custom login command for more detail. File path: ./../support/commands.js
        cy.logout();
    });
});