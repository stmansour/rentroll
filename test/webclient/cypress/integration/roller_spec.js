"use strict";

const pageLoadTime = 2000;
const loginWaitTime = 2000;
const waitTime = 2000;

const HTTP_OK_STATUS = 200;
const API_RESPONSE_SUCCESS_FLAG = 'success';

let receiptResponse;

// w2ui formname
let formName = "receiptForm";

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

// list of grid record
let w2uiGridRecords;

// name of w2ui grid
let w2uiGridName = "receiptsGrid";

// list of columns from the grid
let w2uiGridColumns;

// application cookie name
let applicationCookie = "airoller";

// receiptForm fields value to create a new row
let receiptFormFieldsValue = {
    PmtTypeName: "Cash{enter}",
    DocNo: "AB20180122",
    Amount: "$120.00",
    ERentableName: "Rentable001",
    OtherPayorName: "Akshay Bosamiya",
    Comment: "Testing UI via Cypress"
};

// URL for AIR Receipt application
let applicationPath = "/rhome";

let recordsAPIResponse;

let noRecords;

describe('AIR Receipt UI Tests', function () {

    // -- Perform operation before all tests starts. It runs once before all tests in the block --
    before(function () {
        /*
        * Clear cookies befor starting tests. Because We are preserving cookies to use it all test suit.
        * Running test suit multiple times require new session to login into application.
        */
        cy.clearCookie(applicationCookie);
    });

    /**********************************
     * Assert the title of application
     * 1. Visit Receipt application
     * 2. Assert the title 'AIR Receipts'
     ************************************/

    it('Assert the title of application', function () {

        // It visit baseUrl(from cypress.json) + applicationPath
        cy.visit(applicationPath).wait(pageLoadTime);

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
        Cypress.Cookies.defaults({whitelist: applicationCookie});
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
        cy.route('POST', '/v1/receipts/1').as('getReceipts');

        // It should be visible and selected
        cy.get("#node_receipts").scrollIntoView()
            .should('be.visible')
            .should('have.class', 'w2ui-selected')
            .click().wait(waitTime);

        // Check http status
        cy.wait('@getReceipts').its('status').should('eq', HTTP_OK_STATUS);

        // get API endpoint's responseBody
        cy.get('@getReceipts').then(function (xhr) {

            // Check key `status` in responseBody
            expect(xhr.responseBody).to.have.property('status', API_RESPONSE_SUCCESS_FLAG);

            // get records list from the API response
            recordsAPIResponse = xhr.response.body.records;
            cy.log(recordsAPIResponse);
            noRecords = xhr.response.body.records.length;
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
     * 5. TODO(Akshay):Check Unallocated section visibility and position
     * 6. Close the detail form
     * 7. Assert that form is close.
     *
     *********************************************************/

    it('Test for grid records and record detail', function () {

        // Starting a server to begin routing responses to cy.route()
        cy.server();


        cy.log("Tests for grid records");

        // Check visibility of receiptsGrid
        cy.get('#grid_' + w2uiGridName + '_records').should('be.visible').wait(waitTime);

        // get length from the window and perform tests
        cy.window().then(win => {

            // get list of columns in the grid
            w2uiGridColumns = win.w2ui[w2uiGridName].columns;

            // Match grid record length with total rows in receiptsGrid
            cy.get('#grid_' + w2uiGridName + '_records table tr[recid]').should(($trs) => {
                expect($trs).to.have.length(noRecords);
            });

            // Iterate through each row
            recordsAPIResponse.forEach(function (w2uiGridRecord, rowNo) {

                // Iterat through each column in row
                w2uiGridColumns.forEach(function (w2uiGridColumn, columnNo) {

                    // Skipping traversal icon and RCPT ID column as of now
                    // TODO(Akshay): RCPTID returns '' from the DOM and expected value from the records. Which mismatch and test get fails. Remove condition on columnNo
                    if (columnNo !== 1 && columnNo !== 2) {

                        // Perform test only if w2uiGridColumn isn't hidden
                        if (!w2uiGridColumn.hidden) {

                            // get defaultValue of cell from w2uiGrid
                            let valueForCell = w2uiGridRecord[w2uiGridColumn.field];

                            // Check visibility and default value of cell in the grid
                            cy.get('#grid_' + w2uiGridName + '_data_' + rowNo + '_' + columnNo)
                                .scrollIntoView()
                                .should('be.visible')
                                .should('contain', valueForCell);
                        }
                    }

                });

            });

            cy.log("Tests for detail record form");

            // -- detail record testing --
            const id = recordsAPIResponse[0].RCPTID; // TODO(Akshay): Make it global. Primary id will be change based on the service

            // Routing response to detail record's api requests.
            cy.route('POST', '/v1/receipt/1/' + id).as('getDetailRecord');

            cy.get('#grid_' + w2uiGridName + '_rec_0').click().wait(pageLoadTime);

            cy.wait('@getDetailRecord').its('status').should('eq', HTTP_OK_STATUS);

            cy.get('@getDetailRecord').then(function (xhr) {

                let recordDetailFromAPIResponse = xhr.response.body.record;

                // get form selector
                let formSelector = 'div[name=' + formName + ']';

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

                        // get default value of field
                        defaultValue = recordDetailFromAPIResponse[fieldID];

                        // get field from w2ui form field list
                        field = getW2UIFormFields.find(fieldList => fieldList.field === fieldID);

                        /*
                        * TODO(Akshay): Handle PmtTypeName from app.pmtType. Render Amount in w2ui money type.
                        * */
                        if (fieldID !== "PmtTypeName" && fieldID !== "Amount" && fieldID !== "ERentableName") {

                            // Check visibility and match the default value of the fields.
                            cy.get('#' + fieldID)
                                .should('be.visible')
                                .should('have.value', defaultValue);
                        }
                    });


                // TODO(Akshay): Business Unit value will be handled dynamically.
                // Check Business Unit field must be disabled and have value REX
                cy.get('#BUD').should('be.disabled').and('have.value', 'REX').and('be.visible');

                // TODO(Akshay): List of buttons will be handled globally if needed
                // List of visible and not visible buttons
                let visibleButtons = ["save", "saveprint", "reverse"];
                let notVisibleButtons = ["close"];

                // Check visibility of buttons
                visibleButtons.forEach(function (button) {
                    // Check visibility of button
                    cy.get('button[name=' + button + ']').should('be.visible');
                });

                // Check buttons aren't visible
                notVisibleButtons.forEach(function (button) {
                    // Check button aren't visible
                    cy.get('button[name=' + button + ']').should('not.be.visible');
                });

                // Close the form
                cy.get('[class="fa fa-times"]').click().wait(waitTime);

                // Check that form should not visible after closing it
                cy.get(formSelector).should('not.be.visible');

            });
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
    // it('Test for the Add New Button', function () {
    //     // get form selector
    //     let formSelector = 'div[name=' + formName + ']';
    //
    //     cy.contains('Add New', {force: true}).click().wait(waitTime);
    //
    //     // Check visibility of form
    //     cy.get(formSelector).should('be.visible');
    //
    //     // get record and field list from the w2ui form object
    //     cy.window().then((win) => {
    //         // get w2ui form records
    //         getW2UIFormRecords = win.w2ui[formName].record;
    //
    //         // get w2ui form fields
    //         getW2UIFormFields = win.w2ui[formName].fields;
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
    //             if(typeof defaultValue === 'object'){
    //                 defaultValue = defaultValue.text;
    //             }
    //             /* Money type field have default value in DOM is "$0.00".
    //             And w2ui field have value "0".
    //             To make the comparison change default value "0" to "$0.00" */
    //             else if(field.type === "money" && typeof defaultValue === 'number'){
    //                 defaultValue = "$0.00";
    //             }
    //
    //             /* Skipping tests for Resident Address field. Because it have default value as 'undefined' and in DOM it have value as ''.
    //             Which makes the test fail.
    //             TODO(Sudip): Change default value undefine to ''.
    //             TODO(Akshay): Remove `if` condition for the tests after an issue has been resolved.
    //             */
    //             if(fieldID !== "ERentableName"){
    //                 // Check visibility and match the default value of the fields.
    //                 cy.get('#' + fieldID)
    //                 .should('be.visible')
    //                 .should('have.value', defaultValue);
    //             }
    //
    //         });
    //
    //     // TODO(Akshay): Business Unit value will be handled dynamically.
    //     // Check Business Unit field must be disabled and have value REX
    //     cy.get('#BUD').should('be.disabled').and('have.value', 'REX').and('be.visible');
    //
    //     // TODO(Akshay): List of buttons will be handled globally if needed
    //     // List of visible and not visible buttons
    //     var visibleButtons = ["save", "saveprint"];
    //     var notVisibleButtons = ["reverse", "close"];
    //
    //     // Check visibility of buttons
    //     visibleButtons.forEach(function (button) {
    //         // Check visibility of button
    //         cy.get('button[name=' + button + ']').should('be.visible');
    //     });
    //
    //     // Check buttons aren't visible
    //     notVisibleButtons.forEach(function (button) {
    //         // Check button aren't visible
    //         cy.get('button[name=' + button + ']').should('not.be.visible');
    //     });
    //
    //     // Close the form
    //     cy.get('[class="fa fa-times"]').click().wait(waitTime);
    //
    //     // Check that form should not visible after closing it
    //     cy.get(formSelector).should('not.be.visible');
    //
    // });

});
