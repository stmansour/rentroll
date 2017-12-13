"use strict";

var common = require("./common.js");
var w2ui_utils = require("./w2ui_utils.js");


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

            //
            this.gridColumns = casper.evaluate(function (grid) {
                return w2ui[grid].columns;
            }, this.grid);

            this.columns = this.gridColumns.filter(w2ui_utils.getVisibleColumnName);
            // this.columns = gridConfig.columns;

            casper.click("#" + w2ui_utils.getSidebarID(this.sidebarID));
            casper.log('[FormTest] [{0}] sidebar node clicked with ID: "{1}"'.format(this.grid, this.sidebarID), 'debug', logSpace);
        },

        test: function (test) {
            var that = this;

            casper.wait(common.waitTime, function testGridRecords() {

                //Match w2ui record length with list size in JSON
                var w2uiRecordLength = casper.evaluate(function (gridName) {
                    return w2ui[gridName].records.length;
                }, that.grid);

                test.assert(w2uiRecordLength === that.tableRecords.length, "{0} record length matched with JSON list".format(that.grid));


                that.tableRecords.forEach(function (tableRecord, rowNo) {
                    // TODO: Match database record with rendered UI

                    that.columns.forEach(function (column) {

                        // get coloumn index based on column name/field
                        var columnNo = casper.evaluate(function (gridName, column) {
                            return w2ui[gridName].getColumn(column, true);
                        }, that.grid, column.field);

                        var isVisible = casper.evaluate(function (rowColumnDataSelector) {
                            return isVisibleInViewPort(document.querySelector(rowColumnDataSelector));
                        }, w2ui_utils.getRowColumnDataSelector(that.grid, rowNo, columnNo));

                        // get data at specific cell [rowNo][columnNo]
                        var rowColumnData = casper.evaluate(function (rowColumnDataSelector) {
                            return $(rowColumnDataSelector).text();
                        }, w2ui_utils.getRowColumnDataSelector(that.grid, rowNo, columnNo));

                        // Record visibility in viewport
                        test.assert(isVisible, "{0} is visible in viewport".format(rowColumnData));

                        // JSON file record comparison with rendered w2uiGrid object
                        // test.assertEquals(rowColumnData,tableRecord[column], "{0} is matched with DOM".format(tableRecord[column]));
                    });

                });

                // Check w2ui grid record exists in JSON file

                // w2ui.arsGrid.records : It fetches the record as w2ui object
                // tableRecords: It fetch grid record from the JSON
                // Iterate through w2ui records and check that it is available in JSON tableRecords

                var w2uiGridRecords = casper.evaluate(function (gridName) {
                    return w2ui[gridName].records;
                }, that.grid);

                w2uiGridRecords.forEach(function (w2uiGridRecord) {
                    console.log(w2uiGridRecord)
                });



                // Capture the rendered image
                common.capture(that.capture);

                test.done();
            });
        }
    });
};
