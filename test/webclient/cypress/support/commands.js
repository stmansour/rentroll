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
    cy.readFile("config.json").then((config) => {
        console.log(config);
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

// -- This is a custom command to Change Business Unit to REX --
Cypress.Commands.add("changeBU", function (appSettings) {
    let log;

    // log format
    log = Cypress.log({
        name: "Change Business to REX",
        message: [constants.testBiz, constants.testBizID],
    });

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

    cy.log("Business ID: ", constants.BID);

    log.snapshot().end(); //end custom command
});
