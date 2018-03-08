"use strict";

import * as constants from '../support/utils/constants';

// ***********************************************
// This file contains custom build command for cypress
// ***********************************************

Cypress.Commands.add("login", function () {

    /*
    * Clear cookies before login into application. Because We are preserving cookies to use it all test suit.
    * Running test suit multiple times require new session to login into application.
    */
    cy.clearCookie(constants.APPLICATION_COOKIE);


    /************************************
     * Login into application
     * 1. Send request to endpoint: /v1/authn
     * 2. Pass the necessary parameters
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
        cy.request('POST', constants.LOGIN_END_POINT, {"user": username, "pass": password})
            .then((resp) => {
                cy.log(resp);
                expect(resp.status).to.eq(200);
            });

        // Check Cookie exists after login into application
        cy.getCookie(constants.APPLICATION_COOKIE).should('exist')
            .then(function () {
                log.snapshot().end();
            });
    });

});


// This command log off the application
Cypress.Commands.add("logout", function () {
    cy.request('GET', constants.LOGOUT_END_POINT)
        .then((resp) => {
            expect(resp.status).to.eq(200);
        });
});
