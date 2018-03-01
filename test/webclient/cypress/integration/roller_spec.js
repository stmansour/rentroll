"use strict";

import * as selectors from '../support/utils/get_selectors';
import * as constants from '../support/utils/constants';
import * as common from '../support/utils/common';

// --- Assessments/Receipts --
const asmsM = require('../support/components/asms'); // Assess Charges
const treceiptM = require('../support/components/tenderReceipts'); // Tendered Payment Receipt // TODO(Akshay): Detail Record
const expensesM = require('../support/components/expenses'); // Expenses

// --- Setup ---
const accountM = require('../support/components/accounts');
const pmtM = require('../support/components/pmTypes');
const depAcctM = require('../support/components/depAcct');
const depMethM = require('../support/components/depMeth');
const arsM = require('../support/components/ars');

// this contain app variable of the application
let appSettings;

// holds the test configuration for the modules
let testConfigs;

// records list of module from the API response
let recordsAPIResponse;

// number of records in API response
let noRecordsInAPIResponse;

// list of columns from the grid
let w2uiGridColumns;

// -- Start Cypress UI tests for AIR Roller Application --
describe('AIR Roller UI Tests', function () {

    // -- Perform operation before all tests starts. It runs once before all tests in the block --
    before(function () {
        /*
        * Clear cookies befor starting tests. Because We are preserving cookies to use it all test suit.
        * Running test suit multiple times require new session to login into application.
        */
        cy.clearCookie(constants.APPLICATION_COOKIE);

        // testConfigs = [asmsM.conf, treceiptM.conf, accountM.conf, pmtM.conf, depAcctM.conf, depMethM.conf, arsM.conf];
        testConfigs = [treceiptM.conf];
    });

    /**********************************
     * Assert the title of Roller application
     * 1. Visit Roller application
     * 2. Assert the title 'AIR Receipts'
     ************************************/

    it('Assert the title of Roller application', function () {

        // It visit baseUrl(from cypress.json) + applicationPath
        cy.visit(constants.ROLLER_APPLICATION_PATH).wait(constants.PAGE_LOAD_TIME);

        // Assert application title
        cy.title().should('include', 'AIR Roller');

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

    // -- Test for Node --
    it('Grid records and record detail', function () {

        testConfigs.forEach(function (testConfig) {


            // Starting a server to begin routing responses to cy.route()
            cy.server();

            // To manage the behavior of network requests. Routing the response for the requests.
            cy.route(testConfig.methodType, common.getAPIEndPoint(testConfig.sidebarID)).as('getRecords');

            // Node should be visible and selected
            /************************
            * Select right side node
            *************************/
            cy.get(selectors.getNodeSelector(testConfig.sidebarID))
                .scrollIntoView()
                .should('be.visible')
                .click().wait(constants.WAIT_TIME)
                .should('have.class', 'w2ui-selected');

            // If have date navigation bar than change from and to Date to get in between data
            if(testConfig.haveDateValue){
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
                    // common.gridCellsTest(recordsAPIResponse, w2uiGridColumns, win, testConfig);
                    // TODO(Akshay): Remove comment for above tests

                    // ----------------------------------
                    // -- Tests for detail record form --
                    // ----------------------------------
                    cy.log("Tests for detail record form");

                    // -- detail record testing --
                    const id = recordsAPIResponse[0][testConfig.primaryId];

                    // Routing response to detail record's api requests.
                    cy.route(testConfig.methodType, common.getDetailRecordAPIEndPoint(testConfig.module, id)).as('getDetailRecord');

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



                        // Check Business Unit field must be disabled and have value REX
                        common.BUDFieldTest();

                        // -- Check buttons visibility --
                        common.buttonsTest(testConfig.buttonNamesInDetailForm, testConfig.notVisibleButtonNamesInForm);

                        common.detailFormTest(formSelector, formName, recordDetailFromAPIResponse, win, testConfig);

                        // -- Close the form. And assert that form isn't visible. --
                        common.closeFormTests(formSelector);

                    });
                }
            });
        });
    });
});