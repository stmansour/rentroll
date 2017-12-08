"use strict";

var common = require("./common.js");
var w2ui_utils = require("./w2ui_utils.js");

exports.w2uiGridTest = function(gridConfig) {
    var testCount = 1;
    var testName = "w2ui grid [{0}] test".format(gridConfig.grid);

    casper.test.begin(testName, testCount, {

        // do basic setup first
        setUp: function(/*test*/) {
            // grid name
            this.grid = gridConfig.grid;

            // id defined in w2ui sidebarL1 content for node
            this.sidebarID = gridConfig.sidebarID;

            // capture name
            this.capture = gridConfig.capture;

            // first click on sidebar node labeled tab with provided sidebarID so that
            // client request to load data in grid
            // casper.evaluate(function sidebarNodeClick(a, id) {
            //     $("div[name=sidebarL1]").find("#"+a(id)).click();
            // }, w2ui_utils.getSidebarID, this.sidebarID);
            casper.click('div[name=sidebarL1] #' + w2ui_utils.getSidebarID(this.sidebarID));
            casper.log('[GridTest] [{0}] sidebar node clicked with ID: "{1}"'.format(this.grid, this.sidebarID), 'debug', logSpace);
        },

        // run all the grid test cases
        test: function(test) {
            // need to store it here, so then inside function of type casper callBack,
            // it can access reference of this
            var that = this;

            // --------------------------------------------------
            // TEST 1: verify records length with DOM elements
            // --------------------------------------------------
            // need to wait for some amount of time, so meanwhile w2ui can rendered
            // grid records in DOM
            casper.wait(common.waitTime, function() {
                var recordsLen = this.evaluate(function gridRecordsLen(a, grid) {
                    return a(grid);
                }, w2ui_utils.getGridRecordsLength, that.grid);
                casper.log('[GridTest] [{0}] records length: {1}'.format(that.grid, recordsLen), 'debug', logSpace);

                // check that all rows loaded in w2ui grid object, has been rendered in DOM
                // => w2ui[grid].records.length === DOM's rendered grid table rows length
                var trsLen = this.evaluate(function gridTableRowsLen(a, b, grid) {
                        return b(a(grid));
                    },
                    w2ui_utils.getGridRecordsDivID,
                    w2ui_utils.getGridTableRowsLength,
                    that.grid
                );
                casper.log('[GridTest] [{0}] table rows length: {1}'.format(that.grid, trsLen), 'debug', logSpace);
                test.assertEquals(recordsLen, trsLen, "Grid [{0}] records (table rows) loaded in DOM".format(that.grid));

                // capture grid view at this moment
                common.capture(that.capture);

                // tests are done!!
                test.done();
            });
        }
    });
};