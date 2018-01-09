"use strict";

var require = patchRequire(require);
var common = require("./common.js");

var w2ui_utils = require("./w2ui_utils.js");

// Check cell's visibility in viewport
function performRowColumnVisiblityTest(that, column, recordNo, test) {

    // get coloumn index based on column name/field
    var columnNo = casper.evaluate(function (gridName, column) {
        return w2ui[gridName].getColumn(column, true);
    }, that.grid, column.field);

    // get selector for the cell in the grid
    var rowColumnDataSelector = w2ui_utils.getRowColumnDataSelector(that.grid, recordNo, columnNo);

    // get data at specific cell [recordNo][columnNo]
    var rowColumnData = casper.evaluate(function (rowColumnDataSelector) {
        return $(rowColumnDataSelector).text();
    }, rowColumnDataSelector);

    // check visibility of data at specific cell [recordNo][columnNo]
    var isVisible = casper.evaluate(function (rowColumnDataSelector) {
        return isVisibleInViewPort(document.querySelector(rowColumnDataSelector));
    }, rowColumnDataSelector);

    // Record visibility in viewport
    test.assert(isVisible, "{0} is visible in viewport".format(rowColumnData));

    return rowColumnData;
}

exports.gridRecordsTest = function (gridConfig) {
    var testCount = gridConfig.testCount;
    var testName = "{0} record tests".format(gridConfig.grid);

    // Check status in API Response
    function testAPIResponseStatus(that, test) {
        test.assertEquals(that.apiResponse.status, common.successFlag, "API Response status is {0}".format(common.successFlag));
    }

    // Match total number of records with total number of records with W2UI object
    function testRecordLength(that, test) {
        var w2uiRecordLength = casper.evaluate(function (gridName) {
            return w2ui[gridName].records.length;
        }, that.grid);
        test.assertEquals(w2uiRecordLength, that.apiResponse.total, "{0} record length matched with response list".format(that.grid));


        // get record length from the DOM
        var trsLen = casper.evaluate(function gridTableRowsLen(a, b, grid) {
                return b(a(grid));
            },
            w2ui_utils.getGridRecordsDivID,
            w2ui_utils.getGridTableRowsLength,
            that.grid
        );
        casper.log('[GridTest] [{0}] table rows length: {1}'.format(that.grid, trsLen), 'debug', common.logSpace);

        // match api response record length with records length in DOM
        test.assertEquals(trsLen, that.apiResponse.total, "Grid [{0}] records (table rows) loaded in DOM".format(that.grid));
    }

    // Perform test on row column's data
    function testRowColoumnData(that, test) {

        // TODO: Scrolling records

        that.apiResponse.records.forEach(function (record, recordNo) {

            that.columns.forEach(function (column) {

                // Check cell's visibility in viewport
                var rowColumnData = performRowColumnVisiblityTest(that, column, recordNo, test);

                // Check cell's data exists in API Response
                test.assert(rowColumnData.indexOf(record[column.field]) > -1, "{0} DOM value matched with API response {1}".format(rowColumnData, record[column.field]));
            });

            that.excludeGridColumns.forEach(function (excludeGridColumn) {

                // Check cell's visibility in viewport
                var rowColumnData = performRowColumnVisiblityTest(that, excludeGridColumn, recordNo, test);

                // TODO: Match rowColumnData value with appSettings object

                // Making sure that displayed data length is greater than 0. Remove this test after above To do.
                // Check cell's data exists in API Response
                test.assert(rowColumnData.length > 0, "{0} length is {1}".format(rowColumnData, rowColumnData.length));
            });

        });
    }

    casper.test.begin(testName, testCount, {
        setUp: function (test) {

            // Grid name
            this.grid = gridConfig.grid;

            // Sidebar/Node id
            this.sidebarID = gridConfig.sidebarID;

            // Captured file name
            this.capture = gridConfig.capture;

            // table name
            this.tableName = gridConfig.tableName;

            // list of visible columns
            this.gridColumns = casper.evaluate(function (grid) {
                return w2ui[grid].columns;
            }, this.grid);

            // list of columns which have value in appSettings
            this.excludeGridColumnsKeyValue = gridConfig.excludeGridColumns;

            this.columns = this.gridColumns.filter(w2ui_utils.getVisibleColumnName, this.excludeGridColumnsKeyValue);
            this.excludeGridColumns = this.gridColumns.filter(w2ui_utils.getVisibleExcludedColumnName, this.excludeGridColumnsKeyValue);

            // get api end point for grid
            this.gridEndPoint = gridConfig.endPoint.format(common.apiVersion, common.BID);

            // Send api request to client and get response
            this.apiResponse = casper.evaluate(function (url, method, data) {
                return JSON.parse(__utils__.sendAJAX(url, method, data, false));
            }, this.gridEndPoint, gridConfig.methodType, gridConfig.requestData);

            require('utils').dump(this.apiResponse); // Print API Response

            casper.click("#" + w2ui_utils.getSidebarID(this.sidebarID));
            casper.log('[GridRecordTest] [{0}] sidebar node clicked with ID: "{1}"'.format(this.grid, this.sidebarID), 'debug', common.logSpace);
        },

        test: function (test) {
            var that = this;

            // test api response status
            testAPIResponseStatus(that, test);

            casper.wait(common.waitTime, function testGridRecords() {

                // Match w2ui/DOM record length with list size in API Response
                testRecordLength(that, test);

                // Check each row exist in DOM and visible in viewport
                testRowColoumnData(that, test);

                // Take screen shot of viewport
                common.capture(this.capture);

                test.done();
            })
        }
    });
};
