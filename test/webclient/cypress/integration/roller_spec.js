"use strict";

const pageLoadTime = 2000;
const loginWaitTime = 2000;
const waitTime = 2000;

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

describe('AIR Receipt UI Tests', function () {

    before(function (){
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
        // cy.visit('http://localhost:8270/rhome').wait(pageLoadTime);

        // Assert application title
        cy.title().should('include', 'AIR Receipts');

    });

    /************************************
     * Login into application
     * 1. Fill username and password
     * 2. Click Login button
     * ***********************************/

    it('Login into AIR Receipts', function () {
      // Routing thing is in WIP. It will be refactor tomorrow.
        cy.server();

        cy.route('POST', '/v1/receipts/1').as('getReceipts')
        .then((response) => {
          cy.log(response);
        });

        // read config.json file to get user, pass to get logged in
        cy.readFile("./../../tmp/rentroll/config.json").then((config) => {
            // bundle user, pass and return it
            return {"user": config.Tester1Name, "pass": config.Tester1Pass};
        }).then((creds) => {

            // enter username
            cy.get('input[name=user]')
                .type(creds.user)
                .should('have.value', creds.user);

            // enter password
            cy.get('input[name=pass]')
                .type(creds.pass)
                .should('have.value', creds.pass);

            // click on login and wait for 1s to get the dashboard page
            cy.get('button[name=login]').click().wait(waitTime);
        });

        // Check http status
        cy.wait('@getReceipts').its('status').should('eq', 200);

        // get API endpoint's responseBody
        cy.get('@getReceipts').then(function (xhr) {
            // Check key `status` in responseBody
            expect(xhr.responseBody).to.have.property('status', 'success');
        });

    });

    /*
     * Cypress automatically clears all cookies before each test run.\
     *  It does make application log off.
     * To preserve cookies for entire test suit add that cookie in Cypress cookies's whitelist.
     * Link for more detail: https://docs.cypress.io/api/cypress-api/cookies.html
     */
    beforeEach(function (){
        Cypress.Cookies.defaults({whitelist: applicationCookie});
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
    //             // Preloaded value
    //             // if(fieldID !== "BUD" && fieldID !== "Dt"){
    //             //     // Type values in field
    //             //     cy.get('#' + fieldID)
    //             //     .click()
    //             //     .type(receiptFormFieldsValue[fieldID]);
    //             // }
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


    /*
    * Tests for node
    * 1. Route the receipt API
    * 2. Check http status
    * 3. Check key `status` in API endpoint responseBody
    */
    // it('Test for receipts node', function () {
    //   // It should be visible and selected
    //   cy.get("#node_receipts").scrollIntoView()
    //       .should('be.visible')
    //       .should('have.class', 'w2ui-selected')
    //       .click().wait(2000);
    //
    //     // get API Response
    //     cy.server();
    //
    //     // Route request to the receipt API endpoint
    //     cy.route('POST', '/v1/receipts/1', {"cmd":"get","selected":[],"limit":100,"offset":0,"searchDtStart":"1/1/2018","searchDtStop":"2/1/2018","client":"receipts"}).as('getReceipts');
    //
    //     // Check http status
    //     cy.wait('@getReceipts').its('status').should('eq', 200);
    //
    //     // get API endpoint's responseBody
    //     cy.get('@getReceipts').then(function (xhr) {
    //         // Check key `status` in responseBody
    //         expect(xhr.responseBody).to.have.property('status', 'success');
    //     });
    // });


    // Temporary commented tests
    // it('Test for grid records', function () {
    //
    //       cy.server();
    //
    //       // Check visibility of receiptsGrid
    //       cy.get('#grid_receiptsGrid_records').should('be.visible').wait(waitTime);
    //
    //       // get length from the window and perform tests
    //       cy.window().then(win => {
    //
    //           // get list of columns in the grid
    //           let w2uiGridColumns = win.w2ui[w2uiGridName].columns;
    //
    //           // // get list of visible columns in the grid
    //           // let visibleW2w2uiGridColumns = w2uiGridColumns.filter(column => column.hidden === false);
    //
    //           // get list of records in grid
    //           w2uiGridRecords = win.w2ui[w2uiGridName].records;
    //
    //           var gridRecsLength = w2uiGridRecords.length;
    //
    //           cy.log(w2uiGridRecords);
    //
    //           cy.log("receiptsGrid records length: ", gridRecsLength);
    //
    //           // Match grid record length with total rows in receiptsGrid
    //           cy.get('#grid_receiptsGrid_records table tr[recid]').should(($trs) => {
    //               expect($trs).to.have.length(gridRecsLength);
    //           });
    //
    //
    //           w2uiGridRecords.forEach(function (w2uiGridRecord, rowNo){
    //
    //             w2uiGridColumns.forEach(function (w2uiGridColumn, columnNo) {
    //
    //               // Skipping traversal icon and RCPT ID column as of now
    //               if(columnNo !== 1 && columnNo !== 2){
    //
    //                 // Perform test only if w2uiGridColumn isn't hidden
    //                 if(!w2uiGridColumn.hidden){
    //                   // cy.log("filtering");
    //                   // cy.log("RowNo", row);
    //                   // cy.log(w2uiGridColumn, w2uiGridColumn.field, w2uiGridRecord);
    //                   // cy.log("ColumnNo", column);
    //                   // cy.log("Value at cell", w2uiGridRecord[w2uiGridColumn.field]);
    //
    //                   // get defaultValue of cell from w2uiGrid
    //                   var defaultValue = w2uiGridRecord[w2uiGridColumn.field];
    //
    //                   // Check visibility and default value of cell in the grid
    //                   cy.get('#grid_'+ w2uiGridName + '_data_' + rowNo + '_' + columnNo)
    //                   .scrollIntoView()
    //                   .should('be.visible')
    //                   .should('contain', defaultValue);
    //                 }
    //               }
    //
    //             });
    //
    //             cy.get('#grid_' + w2uiGridName + '_rec_' + rowNo).click().wait(pageLoadTime);
    //
    //             // cy.server();
    //
    //             // cy.debug();
    //
    //             cy.route('POST', '/v1/receipt/1/9').as('getDetailRecord');
    //
    //             cy.wait('@getDetailRecord').its('status').should('eq', 200);
    //            //  cy.route({ method: 'POST',
    //            //  url: 'v1/receipt/1/**',
    //            //    onRequest: (xhr) => {
    //            //
    //            //    },
    //            //    onResponse: (xhr) => {
    //            //      expect(xhr.responseBody).to.have.property('status', 'success');
    //            //      cy.log(xhr.responseBody);
    //            //    }
    //            // });
    //
    //
              // });


          // });
      // });

});
