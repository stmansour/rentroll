"use strict";

var common = require("./common.js");
var w2ui_utils = require("./w2ui_utils.js");

exports.w2uiFormTest = function (formConfig) {
    var testCount = formConfig.testCount;
    var testName = "w2ui form [{0}] test".format(formConfig.form);

    casper.test.begin(testName, testCount, {
        //do basic setup first
        setUp: function (/*test*/) {
            //form name
            this.form = formConfig.form;

            //grid name
            this.grid = formConfig.grid;

            //row number
            this.row = formConfig.row;

            //to open a grid
            this.sidebarID = formConfig.sidebarID;

            //capture name
            this.capture = formConfig.capture;
            this.captureAfterClosingForm = formConfig.captureAfterClosingForm;

            //Button name and class
            this.buttonName = formConfig.buttonName;

            casper.click("#" + w2ui_utils.getSidebarID(this.sidebarID));
            casper.log('[FormTest] [{0}] sidebar node clicked with ID: "{1}"'.format(this.grid, this.sidebarID), 'debug', logSpace);
        },

        // run all the form test cases
        test: function (test) {
            var that = this;

            casper.wait(common.waitTime, function () {
                casper.click("#" + w2ui_utils.getGridRecordID(that.row, that.grid));
                casper.log('[FormTest] [{0}] grid record click with id: "{1}"'.format(that.form, that.row), 'debug', logSpace);
            });

            casper.wait(common.waitTime, function () {

                // Right panel rendering
                test.assertExists("#" + w2ui_utils.getRightPanelID());

                // Form field rendering
                common.capture(that.capture);

                // Button rendering
                that.buttonName.forEach(function (btnName) {
                    var isVisible = casper.evaluate(function formButtonVisibility(btnNode) {
                        return isVisibleInViewPort(document.querySelector(btnNode));
                    }, w2ui_utils.getW2UIButtonReferanceSelector(btnName));

                    test.assert(isVisible, "[{0}] is visible to remote screen.".format(btnName));
                });

                // Right panel after close button
                this.click(w2ui_utils.getCloseButtonSelector());
                this.wait(common.waitTime, function () {
                    common.capture(that.captureAfterClosingForm);
                    test.assertNotVisible("#" + w2ui_utils.getRightPanelID());
                    test.done();
                });
            });
        }
    });
};