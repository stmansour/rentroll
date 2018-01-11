"use strict";

// required modules
var initM = require("./init.js");
var formM = require("./form.js");
var addNewButtonM = require("./addNew.js");
var gridRecordM = require("./gridRecords.js");
var common = require("common.js");

// required components
var asmM = require("./components/asm.js");
var receiptsM = require("./components/receipts.js");
var expensesM = require("./components/expenses");
var depositM = require("./components/deposit.js");
var allocfundsM = require("./components/allocfunds.js");
var transactantsM = require("./components/transactants.js");
var raM = require("./components/ra.js");
var acctM = require("./components/acct.js");
var pmtM = require("./components/pmt.js");
var depM = require("./components/dep.js");
var depmethM = require("./components/depmeth.js");
var arsM = require("./components/ars.js");
var rtM = require("./components/rt.js");
var rentableM = require("./components/rentable.js");

// ========== UI TEST APP OPTIONS ==========
var testBiz = "REX",
    testBizID = -1;
var appSettings;

// ========== INIT CALL ==========
initM.init();

// ========== TESTING ==========
// start webpage request
casper.start(common.pageURL);

// wait for amount of time (defined in `pageLoadWaitTime`)
// for page to be loaded
casper.wait(common.pageLoadWaitTime);

// --------------------------------------------------
// 1. get the app settings, verify it, change the business to "REX"
// --------------------------------------------------
casper.then(function afterStartAndWait() {

    // common settings parameters which are being use throughout application.
    appSettings = this.evaluate(function getAppSettings() {
        return app;
    });

    // TODO: verification pending for appsettings variable, how to do it?

    // get business id to set testBizID
    appSettings.BizMap.forEach(function (item) {
        if (item.BUD === testBiz) {
            testBizID = item.BID;
        }
    });

    // now change the business to REX
    var expBizID = this.evaluate(function selectBizToREX(bizID) {
        document.getElementsByName("BusinessSelect")[0].value = bizID;
        return parseInt(document.getElementsByName("BusinessSelect")[0].value);
    }, testBizID);
    this.log('Business "REX" => expBizID: "{0}", testBizID: "{1}"'.format(expBizID, testBizID), 'debug', common.logSpace);

    // If this test get fail than don't take other test cases in consideration. And exit casperJS
    this.test.assertEquals(expBizID, testBizID, "Business is changed to REX.");
    if (expBizID !== testBizID) {
        // Exiting casperJS
        casper.exit();
    }

    // onSuccessful test set BID value
    common.BID = testBizID;
});

// --------------------------------------------------
// 2. Page basic layout is ready or not
// --------------------------------------------------
casper.then(function pageBasicLayoutTest() {

    // check that basic layout with w2ui has been loaded in page
    var pageInitiated = this.evaluate(function evaluateBasicLayoutCheck() {

        var topToolBarSelector = "div[name=toptoolbar]";
        var sideBarSelector = "div[name=sidebarL1]";
        var topLayoutSelector = "div[name=toplayout]";
        return (
            $("#layout").attr("name") === "mainlayout" &&
            $(topToolBarSelector).length > 0 &&
            $(topToolBarSelector).hasClass("w2ui-toolbar") &&
            $(topLayoutSelector).length > 0 &&
            $(sideBarSelector).length > 0 &&
            $(sideBarSelector).hasClass("w2ui-sidebar")
        );
    });
    this.test.assertEquals(pageInitiated, true, "Page basic layout is ready.");
});

casper.then(function apiIntegrationTest() {
    gridRecordM.gridRecordsTest(pmtM.gridConf);
    gridRecordM.gridRecordsTest(depM.gridConf);
    gridRecordM.gridRecordsTest(depmethM.gridConf);

    /*
    ================================
    // Remove comment for Account Rules for grid tests will fail the test.
    // Check Tests sheet for more detail. Scrolling of records is remaining.
    =================================
    */
    // gridRecordM.gridRecordsTest(arsM.gridConf);

    gridRecordM.gridRecordsTest(rtM.gridConf);
    gridRecordM.gridRecordsTest(rentableM.gridConf);
});

// --------------------------------------------------
// 3. Now start all add new button test
// --------------------------------------------------
casper.then(function addNewButtonTesting() {

    // -------------------------------------------- //
    // Assessments / Receipts Module
    // -------------------------------------------- //

    /*  Assess Charges */
    addNewButtonM.w2uiAddNewButtonTest(asmM.addNewConf);

    /* Receive Receipts  */
    addNewButtonM.w2uiAddNewButtonTest(receiptsM.addNewConf);

    /*
    ================================
    // Remove comment for Expenses tests will fail the test.
    // Business Unit Field is enable
    // Check Tests sheet for more detail.
    =================================
    */
    /* Expenses */
    // addNewButtonM.w2uiAddNewButtonTest(expensesM.addNewConf);


    /*
    ================================
    // Remove comment for Deposits tests will fail the test.
    // Check Tests sheet for more detail.
    =================================
    */
    /* Deposits */
    // addNewButtonM.w2uiAddNewButtonTest(depositM.addNewConf);

    // ------------------------------- //
    // Rental Agreements Module
    // ------------------------------- //

    /*
    ================================
    // Remove comment for Transactants tests will fail the test.
    // Check Tests sheet for more detail.
    =================================
    */
    /* Transactants */
    // addNewButtonM.w2uiAddNewButtonTest(transactantsM.addNewConf);

    // ----------------------------- //
    // Setup Module
    // ----------------------------- //

    /*
    ================================
    // Remove comment for Chart of accounts tests will fail the test.
    // Check Tests sheet for more detail.
    =================================
    */
    /* Chart of accounts */
    // addNewButtonM.w2uiAddNewButtonTest(acctM.addNewConf);

    /* Payment Types */
    addNewButtonM.w2uiAddNewButtonTest(pmtM.addNewConf);

    /* Depository Accounts */
    addNewButtonM.w2uiAddNewButtonTest(depM.addNewConf);

    /* Deposit Methods */
    addNewButtonM.w2uiAddNewButtonTest(depmethM.addNewConf);


    /*
    ================================
    // Remove comment for Account Rules tests will fail the test.
    // Check Tests sheet for more detail.
    =================================
    */
    /* Account Rules */
    // addNewButtonM.w2uiAddNewButtonTest(arsM.addNewConf);


    /*
    ================================
    // Remove comment for Rentable Types tests will fail the test.
    // Because form is not visible properly. Check Screen shot.
    // Check Tests sheet for more detail.
    =================================
    */
    /* Rentable Types */
    // addNewButtonM.w2uiAddNewButtonTest(rtM.addNewConf);

    /* Rentables */
    addNewButtonM.w2uiAddNewButtonTest(rentableM.addNewConf);
});

// --------------------------------------------------
// 4. Now start all right side panel view UI testing
// --------------------------------------------------
/*casper.then(function formTesting() {
    formM.w2uiFormTest(asmM.formConf);
    formM.w2uiFormTest(transactantsM.formConf);
    formM.w2uiFormTest(acctM.formConf);
    formM.w2uiFormTest(pmtM.formConf);
    formM.w2uiFormTest(depM.formConf);
    formM.w2uiFormTest(depmethM.formConf);
    formM.w2uiFormTest(arsM.formConf);
    formM.w2uiFormTest(rtM.formConf);
});*/


// --------------------------------------------------
// 5. Now start all grid record check test
// --------------------------------------------------

// casper.then(function apiIntegrationTest() {
//         gridRecordM.gridRecordsTest(pmtM.gridConf);
//         gridRecordM.gridRecordsTest(depM.gridConf);
//         gridRecordM.gridRecordsTest(depmethM.gridConf);
//
//         /*
//         ================================
//         // Remove comment for Account Rules for grid tests will fail the test.
//         // Check Tests sheet for more detail. Scrolling of records is remaining.
//         =================================
//         */
//         // gridRecordM.gridRecordsTest(arsM.gridConf);
//
//         gridRecordM.gridRecordsTest(rtM.gridConf);
//         gridRecordM.gridRecordsTest(rentableM.gridConf);
// });

// ========== RUN TEST ==========
casper.run();

