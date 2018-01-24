// ***********************************************
// This example commands.js shows you how to
// create various custom commands and overwrite
// existing commands.
//
// For more comprehensive examples of custom
// commands please read more here:
// https://on.cypress.io/custom-commands
// ***********************************************
//
//
// -- This is a parent command --
// Cypress.Commands.add("login", (email, password) => { ... })
//
//
// -- This is a child command --
// Cypress.Commands.add("drag", { prevSubject: 'element'}, (subject, options) => { ... })
//
//
// -- This is a dual command --
// Cypress.Commands.add("dismiss", { prevSubject: 'optional'}, (subject, options) => { ... })
//
//
// -- This is will overwrite an existing command --
// Cypress.Commands.overwrite("visit", (originalFn, url, options) => { ... })

/*  cy
    .contains("Sign in", {log: false})
    .get("#user", {log: false}).type(email, {log: false})
    .get("#password", {log: false}).type(password, {log: false})
    .get("button", {log: false}).click({log: false}) //this should submit the form
    .get("h1", {log: false}).contains("Dashboard", {log: false}) //we should be on the dashboard now
    .url({log: false}).should("match", /dashboard/, {log: false})
    .then(function(){
      log.snapshot().end();
    });*/

// -- This is a custom command to login int AIR application --
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