"use strict";
var require = patchRequire(require);
var common = require("./common.js");

var w2ui_utils = require("./w2ui_utils.js");

function performRowColumnVisiblityTest(that, column, recordNo, test) {
// get coloumn index based on column name/field
    var columnNo = casper.evaluate(function (gridName, column) {
        return w2ui[gridName].getColumn(column, true);
    }, that.grid, column.field);

    var rowColumnDataSelector = w2ui_utils.getRowColumnDataSelector(that.grid, recordNo, columnNo);

    // get data at specific cell [recordNo][columnNo]
    var rowColumnData = casper.evaluate(function (rowColumnDataSelector) {
        return $(rowColumnDataSelector).text();
    }, rowColumnDataSelector);

    // check visibility of data at specific cell [recordNo][columnNo]
    var isVisible = casper.evaluate(function (rowColumnDataSelector) {
        return isVisibleInViewPort(document.querySelector(rowColumnDataSelector));
    }, rowColumnDataSelector);

    console.log(rowColumnDataSelector);

    // Record visibility in viewport
    test.assert(isVisible, "{0} is visible in viewport".format(rowColumnData));
    return rowColumnData;
}

exports.gridRecordsTest = function (gridConfig) {
    var testCount = gridConfig.testCount;
    var testName = "{0} record tests".format(gridConfig.grid);

    function testAPIResponseStatus(that, test) {
        var isSuccess = that.apiResponse.status === 'success';
        test.assert(isSuccess, "API Response status is {0}".format(isSuccess));
    }

    function testRecordLength(that, test) {
        var w2uiRecordLength = casper.evaluate(function (gridName) {
            return w2ui[gridName].records.length;
        }, that.grid);
        test.assert(w2uiRecordLength === that.apiResponse.total, "{0} record length matched with response list".format(that.grid));
    }

    function testRowColoumnVisiblity(that, test) {

        // TODO: Scrolling records

        that.apiResponse.records.forEach(function (record, recordNo) {

            that.columns.forEach(function (column) {
                var rowColumnData = performRowColumnVisiblityTest(that, column, recordNo, test);
                test.assert(rowColumnData.indexOf(record[column.field]) > -1, "{0} DOM value matched with API response {1}".format(rowColumnData, record[column.field]));
            });

            that.excludeGridColumns.forEach(function (excludeGridColumn) {
                var rowColumnData = performRowColumnVisiblityTest(that, excludeGridColumn, recordNo, test);
                // TODO: Match rowColumnData value with appSettings object
                // Making sure that displayed data length is greater than 0. Remove this test after above To do.
                test.assert(rowColumnData.length > 0, "{0} length is {1}".format(rowColumnData, rowColumnData.length));
            });

        });
    }

    casper.test.begin(testName, testCount, {
        setUp: function (test) {

            this.grid = gridConfig.grid;

            this.sidebarID = gridConfig.sidebarID;

            this.capture = gridConfig.capture;

            // table name
            this.tableName = gridConfig.tableName;

            //
            this.gridColumns = casper.evaluate(function (grid) {
                return w2ui[grid].columns;
            }, this.grid);

            this.excludeGridColumnsKeyValue = gridConfig.excludeGridColumns;
            this.columns = this.gridColumns.filter(w2ui_utils.getVisibleColumnName, this.excludeGridColumnsKeyValue);
            this.excludeGridColumns = this.gridColumns.filter(w2ui_utils.getVisibleExcludedColumnName, this.excludeGridColumnsKeyValue);

            // Send api request to client and get response
            this.apiResponse = casper.evaluate(function (url, method, data) {
                return JSON.parse(__utils__.sendAJAX(url, method, data, false));
            }, gridConfig.endPoint, gridConfig.methodType, gridConfig.requestData);

            /*require('utils').dump(this.apiResponse);*/ // Print API Response

            casper.click("#" + w2ui_utils.getSidebarID(this.sidebarID));
            casper.log('[GridRecordTest] [{0}] sidebar node clicked with ID: "{1}"'.format(this.grid, this.sidebarID), 'debug', logSpace);
        },

        test: function (test) {
            var that = this;

            // test api response status
            testAPIResponseStatus(that, test);

            casper.wait(common.waitTime, function testGridRecords() {
                // Match w2ui record length with list size in API Response
                testRecordLength(that, test);

                // Check each row exist in DOM and visible in viewport
                testRowColoumnVisiblity(that, test);

                test.done();
            });
            common.capture(this.capture);
        }
    });
};