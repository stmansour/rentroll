"use strict";
var require = patchRequire(require);
var common = require("./common.js");

var w2ui_utils = require("./w2ui_utils.js");

exports.apiIntegrationTest = function (gridConfig) {
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
        // --------- Grid record scrolling Algorithm -----------
        // Get height of row from the first record.
        // Scroll the parent div of records to (0, rowNo * height)
        // Apply scroll only if there are records available
        // -----------------------------------------------------
        var recordsParentDivSelector = w2ui_utils.getRecordsParentDivSelector(that.grid);


        for (var recordNo = 0; recordNo < that.apiResponse.total; recordNo++) {

            that.columns.forEach(function (column) {

                // get coloumn index based on column name/field
                var columnNo = casper.evaluate(function (gridName, column) {
                    return w2ui[gridName].getColumn(column, true);
                }, that.grid, column.field);

                var rowColumnDataSelector = w2ui_utils.getRowColumnDataSelector(that.grid, recordNo, columnNo);

                // get data at specific cell [recordNo][columnNo]
                var rowColumnData = casper.evaluate(function (rowColumnDataSelector) {
                    return $(rowColumnDataSelector).text();
                }, rowColumnDataSelector);

                // Get height of row
                /*var height = casper.evaluate(function (rowColumnDataSelector) {
                    return document.querySelector(rowColumnDataSelector).getBoundingClientRect().height;
                }, rowColumnDataSelector);

                casper.evaluate(function (recordsParentDivSelector, height, rowNo) {
                    document.querySelector(recordsParentDivSelector).scrollTo(0, height*rowNo);
                }, rowColumnDataSelector, height, recordNo);

                casper.then(function () {
                    common.capture("ScrollHeight.jpg");
                });*/

                // check visibility of data at specific cell [recordNo][columnNo]
                var isVisible = casper.evaluate(function (rowColumnDataSelector) {
                    return isVisibleInViewPort(document.querySelector(rowColumnDataSelector));
                }, rowColumnDataSelector);

                // Record visibility in viewport
                test.assert(isVisible, "{0} is visible in viewport".format(rowColumnData));
            });
        }
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

            this.columns = this.gridColumns.filter(w2ui_utils.getVisibleColumnName);

            // Send api request to client and get response
            this.apiResponse = casper.evaluate(function (url, method, data) {
                return JSON.parse(__utils__.sendAJAX(url, method, data, false));
            }, gridConfig.endPoint, gridConfig.methodType, gridConfig.requestData);

            // records in the table
            this.recordsInAPI = this.apiResponse.records;


            require('utils').dump(this.apiResponse); // Print API Response

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

                // TODO: Check w2ui grid record exists in API Response
                // w2ui.arsGrid.records : It fetches the record as w2ui object
                // tableRecords: It fetch grid record from the JSON
                // Iterate through w2ui records and check that it is available in JSON tableRecords


                test.done();
            });
            common.capture(this.capture);
        }
    });
};