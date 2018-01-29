"use strict";

import * as selectors from '../support/utils/get_selectors';
import * as constants from '../support/utils/constants';

// ***********************************************
// This file contains custom build command for cypress
// ***********************************************

// -- This is a custom command to login into AIR application --
Cypress.Commands.add("login", function(){

    /************************************
     * Login into application
     * 1. Fill username and password
     * 2. Click Login button
     * ***********************************/

    let username;
    let password;
    let log;

    // read config.json file to get user, pass to get logged in
    cy.readFile("./../../tmp/rentroll/config.json").then((config) => {
        // bundle user, pass and return it
        return {"user": config.Tester1Name, "pass": config.Tester1Pass};
    }).then((creds) => {

        username    = creds.user;
        password = creds.pass;

        // log format
        log = Cypress.log({
            name: "login",
            message: [username, password],
            consoleProps: function () {
                return {
                    username: username,
                    password: password
                }
            }
        });

        // login steps
        cy
            .get('input[name=user]').type(username).should('have.value', username) // enter username
            .get('input[name=pass]').type(password).should('have.value', password) // enter password
            .get('button[name=login]').click().wait(2000) // click on login and wait for 1s to get the dashboard page
            .then(function () {
                log.snapshot().end(); //end custom command
            });
    });

});

// Cypress.Commands.add("closeForm", (formSelector) => {
//
//     /****************************************************
//      * 1. Close the form
//      * 2. Assert that form isn't visible to the screen
//      ***************************************************/
//
//     let log = Cypress.log({
//         name: "closeForm",
//         message: "Closing the form"
//     });
//
//     // Close the form
//     cy
//         .get(selectors.getFormCloseButtonSelector())
//         .click()
//         .wait(constants.WAIT_TIME);
//
//     cy.log("###########");
//     cy.log(formSelector);
//
//     // Check that form should not visible after closing it
//     cy
//         .get(formSelector)
//         .should('not.be.visible');
//
// });