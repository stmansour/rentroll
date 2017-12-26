"use strict";

// required modules
var initM = require("./init.js");
var gridM = require("./grid.js");
var formM = require("./form.js");
var addNewButtonM = require("./addNew.js");
var gridRecordM = require("./gridRecords.js");

// required components
var asmM = require("./components/asm.js");
var receiptsM = require("./components/receipts.js");
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
var pageURL = "http://localhost:8270/home",
    pageWidth = 1280,
    pageHeight = 720,
    pageLoadWaitTime = 2000;
var testBiz = "REX",
    testBizID = -1;
var logSpace = "rrLog" ;
var appSettings;


// ========== INIT CALL ==========
initM.init();

// ========== TESTING ==========
// start webpage request
casper.start(pageURL);

// wait for amount of time (defined in `pageLoadWaitTime`)
// for page to be loaded
casper.wait(pageLoadWaitTime);

// --------------------------------------------------
// 1. get the app settings, verify it, change the business to "REX"
// --------------------------------------------------
casper.then(function afterStartAndWait() {
    appSettings = this.evaluate(function getAppSettings() {
        return app;
    });

    // TODO: verification pending for appsettings variable, how to do it?

    appSettings.BizMap.forEach(function(item) {
        if (item.BUD == testBiz) {
            testBizID = item.BID;
        }
    });

    // now change the business to REX
    var expBizID = this.evaluate(function selectBizToREX(bizID) {
        document.getElementsByName("BusinessSelect")[0].value = bizID;
        return parseInt(document.getElementsByName("BusinessSelect")[0].value);
    }, testBizID);
    this.log('Business "REX" => expBizID: "{0}", testBizID: "{1}"'.format(expBizID, testBizID), 'debug', logSpace);
    this.test.assertEquals(expBizID, testBizID, "Business is changed to REX.");
});

// --------------------------------------------------
// 2. Page basic layout is ready or not
// --------------------------------------------------
/*casper.then(function pageBasicLayoutTest() {
    // check that basic layout with w2ui has been loaded in page
    var pageInitiated = this.evaluate(function evaluateBasicLayoutCheck() {
        return (
            $("#layout").attr("name") == "mainlayout" &&
            $("div[name=toptoolbar]").length > 0 &&
            $("div[name=toptoolbar]").hasClass("w2ui-toolbar") &&
            $("div[name=toplayout]").length > 0 &&
            $("div[name=sidebarL1]").length > 0 &&
            $("div[name=sidebarL1]").hasClass("w2ui-sidebar")
        );
    });
    this.test.assertEquals(pageInitiated, true, "Page basic layout is ready.");
});*/

// --------------------------------------------------
// 3. Now start all grid view UI testing
// --------------------------------------------------
/*casper.then(function gridTesting() {
    gridM.w2uiGridTest(asmM.gridConf);
    gridM.w2uiGridTest(receiptsM.gridConf);
    gridM.w2uiGridTest(depositM.gridConf);
    gridM.w2uiGridTest(allocfundsM.gridConf);
    gridM.w2uiGridTest(transactantsM.gridConf);
    gridM.w2uiGridTest(raM.gridConf);
    gridM.w2uiGridTest(acctM.gridConf);
    gridM.w2uiGridTest(pmtM.gridConf);
    gridM.w2uiGridTest(depM.gridConf);
    gridM.w2uiGridTest(depmethM.gridConf);
    gridM.w2uiGridTest(arsM.gridConf);
    gridM.w2uiGridTest(rtM.gridConf);
    gridM.w2uiGridTest(rentableM.gridConf);
});*/

// --------------------------------------------------
// 4. Now start all add new button test
// --------------------------------------------------
/*casper.then(function addNewButtonTesting() {
    addNewButtonM.w2uiAddNewButtonTest(acctM.addNewConf);
    addNewButtonM.w2uiAddNewButtonTest(pmtM.addNewConf);
    addNewButtonM.w2uiAddNewButtonTest(depM.addNewConf);
    addNewButtonM.w2uiAddNewButtonTest(depmethM.addNewConf);
    addNewButtonM.w2uiAddNewButtonTest(arsM.addNewConf);
    addNewButtonM.w2uiAddNewButtonTest(rentableM.addNewConf);
});*/

// --------------------------------------------------
// 5. Now start all right side panel view UI testing
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
// 6. Now start all grid record check test
// --------------------------------------------------

casper.then(function apiIntegrationTest() {
    gridRecordM.apiIntegrationTest(pmtM.gridConf);
    gridRecordM.apiIntegrationTest(depM.gridConf);
    gridRecordM.apiIntegrationTest(depmethM.gridConf);
    gridRecordM.apiIntegrationTest(arsM.gridConf);
    gridRecordM.apiIntegrationTest(rtM.gridConf);
    gridRecordM.apiIntegrationTest(rentableM.gridConf);
});

// ========== RUN TEST ==========
casper.run();

