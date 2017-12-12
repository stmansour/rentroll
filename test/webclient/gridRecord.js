"use strict";

var common = require("./common.js");
var rrJson = require("./rentroll.json");

exports.w2uiGridRecordTest = function (gridConfig) {
    var testCount = gridConfig.testCount;
    var testName = "w2ui grid record {0} test".format(gridConfig.grid);

    casper.test.begin(testName, testCount, {
        // do basic setup first
        setUp: function (/*test*/) {
            // grid name
            this.grid = gridConfig.grid;

            // to open a grid
            this.sidebarID = gridConfig.sidebarID;

            // capture name
            this.capture = gridConfig.capture;

            // table name
            this.tableName = gridConfig.tableName;

            // records in the table
            this.tableRecords = rrJson[this.tableName];

            casper.click("#" + w2ui_utils.getSidebarID(this.sidebarID));
            casper.log('[FormTest] [{0}] sidebar node clicked with ID: "{1}"'.format(this.grid, this.sidebarID), 'debug', logSpace);
        },

        test: function (test) {
            var that = this;

            casper.wait(common.waitTime, function testGridRecords() {
                that.tableRecords.forEach(function (tableRecord) {
                    // TODO: Match database record with rendered UI
                });
            });
        }
    });
};
