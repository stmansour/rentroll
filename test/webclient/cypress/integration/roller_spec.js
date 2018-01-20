"use strict";

describe('Basic Cypress UI Testing Demo', function() {
    it("visiting home page", function() {

        // read config.json file to get user, pass to get logged in
        cy.readFile("./../../tmp/rentroll/config.json")
            .then((config) => {
                // bundle user, pass and return it
                return {"user": config.Tester1Name, "pass": config.Tester1Pass};
            })
            .then((creds) => {
                    // visit homepage
                    cy.visit("http://localhost:8270/home");

                    // enter username
                    cy.get('input[name=user]')
                    .type(creds.user)
                    .should('have.value', creds.user);

                    // enter password
                    cy.get('input[name=pass]')
                    .type(creds.pass)
                    .should('have.value', creds.pass);

                    // click on login and wait for 1s to get the dashboard page
                    cy.get('button[name=login]').click().wait(1000);
            });

        // click on left side node
        cy.get('#node_accounts').click().wait(500);

        // check length of grid records
        cy.window().then(win => {
            var gridRecsLength = win.w2ui.accountsGrid.records.length;
            cy.log("accountsGrid records length: ", gridRecsLength);
            cy.get('#grid_accountsGrid_records table tr[recid]').should(($trs) => {
                expect($trs).to.have.length(gridRecsLength);
            });
        });

        // now scroll the division to bottom
        // put wait so that it can be seen by human eyes!
        cy.get('#grid_accountsGrid_records').scrollTo('bottom').wait(1000);

        // click on first record from the grid, so that it loads the form
        /*cy.get('#grid_accountsGrid_records table tr[recid]:last').click();*/
        cy.get('#grid_accountsGrid_records tr[recid]:first').click();

        // get visible inputs in the DOM
        cy.get('div[name=accountForm]').find('input:not(:hidden)').should(($inputs) => {
            expect($inputs).to.have.length(11);
        });

        // check BUD field in the form, it should be visible and it should be disabled as well
        cy.get('div[class=w2ui-form-box]').find('input[name=BUD]').should('be.visible').should('be.disabled');

        // check BUD field in the form, it should be visible and it should be disabled as well
        cy.get('div[class=w2ui-form-box]').find('input[name=recid]').should('have.class', 'w2field');
    });
});
