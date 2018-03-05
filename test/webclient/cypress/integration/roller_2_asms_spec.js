"use strict";

import * as selectors from '../support/utils/get_selectors';
import * as constants from '../support/utils/constants';
import * as common from '../support/utils/common';

// --- Assessments/Receipts --
const asmsM = require('../support/components/asms'); // Assess Charges

// this contain app variable of the application
let appSettings;

// holds the test configuration for the modules
let testConfig;

// records list of module from the API response
let recordsAPIResponse;

// number of records in API response
let noRecordsInAPIResponse;

// list of columns from the grid
let w2uiGridColumns;

// -- Start Cypress UI tests for AIR Roller Application --
describe('AIR Roller UI Tests - Assessment Charges', function () {

    // -- Perform operation before all tests starts. It runs once before all tests in the block --
    before(function () {

        // --- Login into Application before starting any tests ---
        // Check custom login command for more detail. File path: ./../support/commands.js
        cy.login();

        cy.visit(constants.ROLLER_APPLICATION_PATH).wait(constants.PAGE_LOAD_TIME);

        testConfig = asmsM.conf;
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
        common.testGridRecords(testConfig);
    });

    it('Record Detail Form', function () {

    });

    it('Add new record form', function () {
        // ---------------------------------------
        // ----- Tests for add new record form ---
        // ---------------------------------------

        common.testAddNewRecordForm(testConfig);
    });

    // -- Test for Node --
    it('Grid records and record detail', function () {

        // Starting a server to begin routing responses to cy.route()
        cy.server();

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
                common.gridCellsTest(recordsAPIResponse, w2uiGridColumns, win, testConfig);

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

    // -- Perform operation after all tests finish. It runs once after all tests in the block --
    after(function () {

        // --- Logout from the Application after finishing all tests ---
        // Check custom login command for more detail. File path: ./../support/commands.js
        cy.logout();
    });
});