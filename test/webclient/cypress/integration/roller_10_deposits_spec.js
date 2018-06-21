"use strict";

import * as constants from '../support/utils/constants';
import * as selectors from '../support/utils/get_selectors';
import * as common from '../support/utils/common';

// --- Assessments/Receipts --
const section = require('../support/components/deposits'); // Deposits

// this contain app variable of the application
let appSettings;

// holds the test configuration for the modules
let testConfig;

// -- Start Cypress UI tests for AIR Roller Application --
describe('AIR Roller UI Tests - Deposits', function () {

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

    it('Grid Records', function () {
        common.testGridRecords(recordsAPIResponse, noRecordsInAPIResponse, testConfig);
    });

    it('Record Detail Form', function () {

        // route the endpoint for grid records in deposit's record detail form
        const id = recordsAPIResponse[0][testConfig.primaryId];
        cy.server();
        cy.route(testConfig.methodType, common.getDetailRecordAPIEndPoint(testConfig.gridInForm, id)).as('getDepositListGrid');

        // ----------------------------------
        // -- Tests for detail record form --
        // ----------------------------------
        // Params:
        // recordsAPIResponse: list of record from the api response,
        // testConfig: configuration for running tests
        common.testRecordDetailForm(recordsAPIResponse, testConfig);


        // perform tests on grid record (depositListGrid) in deposit detail record form
        cy.get('@getDepositListGrid').then(function (xhr) {

            expect(xhr.responseBody).to.have.property('status', constants.API_RESPONSE_SUCCESS_FLAG);

            let recordsAPIResponse = xhr.response.body.records;

            cy.log(recordsAPIResponse);

            // -- Assigning number of records to 0 if no records are available in response --
            if (recordsAPIResponse) {
                noRecordsInAPIResponse = xhr.response.body.records.length;
            } else {
                noRecordsInAPIResponse = 0;
            }

            // assign grid name
            testConfig.grid = "depositListGrid";

            // Perform test on each cell of grid records
            common.testGridRecords(recordsAPIResponse, noRecordsInAPIResponse, testConfig);

        });

        // -- Close the form. And assert that form isn't visible. --
        common.closeFormTests(selectors.getFormSelector(testConfig.form));
    });


    it('Check default value of fields for new record form', function () {
        // ---------------------------------------
        // ----- Tests for add new record form ---
        // ---------------------------------------
        common.testAddNewRecordForm(testConfig);
    });

    // -- Perform operation after all tests finish. It runs once after all tests in the block --
    after(function () {

        // --- Logout from the Application after finishing all tests ---
        // Check custom login command for more detail. File path: ./../support/commands.js
        cy.logout();
    });
});