"use strict";

import * as constants from '../support/utils/constants';
import * as selectors from '../support/utils/get_selectors';
import * as common from '../support/utils/common';

// --- Assessments/Receipts --
const section = require('../support/components/applyreceipt'); // Apply receipt

// this contain app variable of the application
let appSettings;

// holds the test configuration for the modules
let testConfig;

// -- Start Cypress UI tests for AIR Roller Application --
describe('AIR Roller UI Tests - Apply Receipt', function () {

    // // records list of module from the API response
    let recordsAPIResponse;

    let noRecordsInAPIResponse;

    let allocfundsRecordLength;

    let unpaidasmsRecordsLength;

    let payorRecordsLength;

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
        // ----------------------------------
        // -- Tests for detail record form --
        // ----------------------------------

        // route the endpoint for grid records in deposit's record detail form
        cy.server();
        cy.route(testConfig.methodType, common.getAPIEndPoint(testConfig.module, 1)).as('getDetailRecord');
        cy.route(testConfig.methodType, common.getDetailRecordAPIEndPoint('unpaidasms', 1)).as('getUnpaidAsms');
        cy.route('GET', common.getDetailRecordAPIEndPoint('payorfund', 1)).as('getPayorFundResponse');

        // First record is blank. So temporary perform tests on second record.
        cy.get(selectors.getSecondRecordInGridSelector(testConfig.grid)).click().wait(constants.WAIT_TIME);

        // check response status of API end point
        cy.wait('@getDetailRecord').its('status').should('eq', constants.HTTP_OK_STATUS);
        cy.wait('@getUnpaidAsms').its('status').should('eq', constants.HTTP_OK_STATUS);
        cy.wait('@getPayorFundResponse').its('status').should('eq', constants.HTTP_OK_STATUS);

        // perform tests on record detail form
        cy.get('@getDetailRecord').then(function (xhr) {
            cy.log(xhr);

            let allocfundsRecords = xhr.response.body.records;
            cy.log(allocfundsRecords);

            if (allocfundsRecords) {
                allocfundsRecordLength = allocfundsRecords.length;
            } else {
                allocfundsRecordLength = 0;
            }

            let payorName = allocfundsRecords[1].Name;

            // Check payor name in detailed form
            cy.get('[name=unallocForm]').should('contain', payorName);

            // Check Auto-Allocate button
            cy.get('#auto_allocate_btn').should('have.text', 'Auto-Allocate');
        });

        cy.get('@getPayorFundResponse').then(function (xhr) {
            cy.log(xhr);

            let unpaidasmsRecords = xhr.response.body.records;
            cy.log(unpaidasmsRecords);

            if (unpaidasmsRecords) {
                unpaidasmsRecordsLength = unpaidasmsRecords.length;
            } else {
                unpaidasmsRecordsLength = 0;
            }

            let fund = xhr.response.body.record.fund;



            // Verify fund value
            cy.get('#total_fund_amount').should('contain', fund);
        });

        cy.get('@getUnpaidAsms').then(function (xhr) {
            cy.log(xhr);

            let payorRecords = xhr.response.body.records;
            cy.log(payorRecords);

            if (payorRecords) {
                payorRecordsLength = payorRecords.length;
            } else {
                payorRecordsLength = 0;
            }

            // assign grid name
            // testConfig.grid = 'unpaidASMsGrid';
            // Perform test on each cell of grid records
            testConfig.grid = 'unpaidASMsGrid';
            common.testGridRecords(payorRecords, payorRecordsLength, testConfig);
        });

        // -- Close the form. And assert that form isn't visible. --
        common.closeFormTests(selectors.getFormSelector(testConfig.form));
    });

    // -- Perform operation after all tests finish. It runs once after all tests in the block --
    after(function () {

        // --- Logout from the Application after finishing all tests ---
        // Check custom login command for more detail. File path: ./../support/commands.js
        cy.logout();
    });
});