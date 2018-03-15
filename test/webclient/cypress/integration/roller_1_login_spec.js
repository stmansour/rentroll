"use strict";

import * as constants from '../support/utils/constants';

// -- Start Cypress UI tests for AIR Roller Application --
describe('AIR Roller UI Tests - Login', function () {

    // It runs before all tests get executed.
    before(function () {
        cy.clearCookie(constants.APPLICATION_COOKIE);
    });

    /**********************************
     * Assert the title of Roller application
     * 1. Visit Roller application
     * 2. Assert the title 'AIR Receipts'
     ************************************/
    it('Assert the title of Roller application', function () {

        // It visit baseUrl(from cypress.json) + applicationPath
        cy.visit(constants.ROLLER_APPLICATION_PATH).wait(constants.LOGIN_WAIT_TIME);

        // Assert application title
        cy.title().should('include', 'AIR Roller');

    });

    /************************************
     * Login into application
     * 1. Fill username and password
     * 2. Click Login button
     *
     * Expect:
     * Username and Password field must be visible.
     * It must have proper value which match with typed text.
     * Login button must be visible.
     * ***********************************/
    it('Login into AIR Roller Application', function () {
        let username;
        let password;
        let log;

        // read config.json file to get user, pass to get logged in
        cy.readFile("config.json").then((config) => {
            console.log(config);
            // bundle user, pass and return it
            return {"user": config.Tester1Name, "pass": config.Tester1Pass};
        }).then((creds) => {

            username = creds.user;
            password = creds.pass;

            // log format
            log = Cypress.log({
                name: "login",
                message: [username, password],
                consoleProps: function () {
                    return {
                        username: username,
                        password: password
                    };
                }
            });

            // Route the Login endpoint
            cy.server();
            cy.route('POST', constants.LOGIN_END_POINT).as('getLogin');

            // login steps
            cy
                .get('input[name=user]').should('be.visible').type(username).should('have.value', username) // enter username
                .get('input[name=pass]').should('be.visible').type(password).should('have.value', password) // enter password
                .get('button[name=login]').should('be.visible').click().wait(2000); // click on login and wait for 1s to get the dashboard page

            // Check Login endpoint response and status
            cy.wait('@getLogin').its('status').should('eq', constants.HTTP_OK_STATUS);
            cy.get('@getLogin').then(function (xhr) {
                expect(xhr.response.body.status).to.be.eq(constants.API_RESPONSE_SUCCESS_FLAG);
                cy.log(xhr);
            });

        });
    });

    /**********************************
     * Check Basic Layout of application after login.
     *
     * Expect:
     * Sidebar panel, top layout panel, main layout must be visible
     ************************************/
    it('Check Basic Layout', function () {
        cy.get('#layout_toplayout_panel_left').should('be.visible');
        cy.get('#layout_toplayout_panel_main').should('be.visible');
        cy.get('#layout_mainlayout_panel_top').should('be.visible');
    });

    // -- Perform operation after all tests finish. It runs once after all tests in the block --
    after(function () {

        // --- Logout from the Application after finishing all tests ---
        // Check custom login command for more detail. File path: ./../support/commands.js
        cy.logout();
    });
});