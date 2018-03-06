"use strict";

import * as constants from '../support/utils/constants';

// -- Start Cypress UI tests for AIR Roller Application --
describe('AIR Roller UI Tests', function () {

    /**********************************
     * Assert the title of Roller application
     * 1. Visit Roller application
     * 2. Assert the title 'AIR Receipts'
     ************************************/
    it('Assert the title of Roller application', function () {

        // It visit baseUrl(from cypress.json) + applicationPath
        cy.visit(constants.RECEIPT_APPLICATION_PATH).wait(constants.PAGE_LOAD_TIME);

        // Assert application title
        cy.title().should('include', 'AIR Receipt');

    });

    it('Login into AIR Receipt Application', function () {

        /*
        * Clear cookies before login into application. Because We are preserving cookies to use it all test suit.
        * Running test suit multiple times require new session to login into application.
        */
        cy.clearCookie(constants.APPLICATION_COOKIE);


        /************************************
         * Login into application
         * 1. Fill username and password
         * 2. Click Login button
         * ***********************************/

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

            // login steps
            cy
                .get('input[name=user]').type(username).should('have.value', username) // enter username
                .get('input[name=pass]').type(password).should('have.value', password) // enter password
                .get('button[name=login]').click().wait(2000); // click on login and wait for 1s to get the dashboard page

        });
    });

    // it('Check Basic Layout', function () {
    //    // TODO(Akshay): Write tests for basic layout of application
    // });

    // -- Perform operation after all tests finish. It runs once after all tests in the block --
    after(function () {

        // --- Logout from the Application after finishing all tests ---
        // Check custom login command for more detail. File path: ./../support/commands.js
        cy.logout();
    });
});