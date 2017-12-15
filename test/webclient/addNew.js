"use strict";

var common = require("./common.js");
var w2ui_utils = require("./w2ui_utils.js");

exports.w2uiAddNewButtonTest = function (addNewButtonConfig) {
    var testCount = addNewButtonConfig.testCount;
    var testName = "w2ui add new button [{0}] test".format(addNewButtonConfig.form);

    casper.test.begin(testName, testCount, {
        //do basic setup first
        setUp: function (/*test*/) {
            //form name
            this.form = addNewButtonConfig.form;

            //grid name
            this.grid = addNewButtonConfig.grid;

            //to open a grid
            this.sidebarID = addNewButtonConfig.sidebarID;

            //list of input fields
            // this.inputFields = addNewButtonConfig.inputField;
            this.formFields = casper.evaluate(function (form) {
                return w2ui[form].fields;
            }, this.form);

            this.inputFields = this.formFields.filter(w2ui_utils.getTextTypeW2UIFields);
            // ["Name", "Description"]

            //list of input select fields
            // this.inputSelectField = addNewButtonConfig.inputSelectField;
            this.inputSelectField = this.formFields.filter(w2ui_utils.getInputListW2UIFields);

            //capture name
            this.capture = addNewButtonConfig.capture;

            //Button name and class
            this.buttonName = addNewButtonConfig.buttonName;

            //Checkboxes list
            this.checkboxes = this.formFields.filter(w2ui_utils.getCheckBoxW2UIFields);

            casper.click("#" + w2ui_utils.getSidebarID(this.sidebarID));
            casper.log('[FormTest] [{0}] sidebar node clicked with ID: "{1}"'.format(this.grid, this.sidebarID), 'debug', logSpace);

            casper.wait(common.waitTime, function clickAddNewButton() {
                //It will click table cell with the text 'Add New'
                casper.clickLabel("Add New", "td");
            });
        },

        test: function (test) {
            var that = this;

            casper.wait(common.waitTime, function testRightPanel() {

                // Right panel rendering
                test.assertExists("#" + w2ui_utils.getRightPanelID());

                // BUD Field Test
                test.assertSelectorExists(w2ui_utils.getBUDSelector());
                var businessUnitValue = casper.evaluate(function getBusinessUnit(bud_selector) {
                    return document.querySelector(bud_selector).value;
                }, w2ui_utils.getBUDSelector());

                if (businessUnitValue === testBiz) {
                    test.assert(true, "Business Unit value is {0}.".format(businessUnitValue))
                } else {
                    test.assert(false, "Wrong Business unit");
                }

                // BUD disability
                var isBusinessUnitValueDisabled = casper.evaluate(function (bud_selector) {
                    return document.querySelector(bud_selector).disabled;
                }, w2ui_utils.getBUDSelector());

                test.assert(isBusinessUnitValueDisabled, "Disability of business unit field.");

                // Input fields test
                that.inputFields.forEach(function (inputFieldID) {
                    var inputFieldSelector = w2ui_utils.getInputFieldSelector(inputFieldID.field);

                    var isVisible = casper.evaluate(function inputFieldVisibility(inputFieldSelector) {
                        return isVisibleInViewPort(document.querySelector(inputFieldSelector));
                    }, inputFieldSelector);

                    test.assert(isVisible, "{0} input field is visible to remote screen.".format(inputFieldID.field));

                    var inputFieldValue = casper.evaluate(function (inputFieldSelector) {
                        return document.querySelector(inputFieldSelector).value;
                    }, inputFieldSelector);

                    if (inputFieldValue === "") {
                        test.assert(true, "{0} field is blank".format(inputFieldID.field));
                    }
                    else {
                        test.assert(false, "{0} field is not blank".format(inputFieldID.field));
                    }
                });

                // Dropdown Input fields test
                that.inputSelectField.forEach(function (inputSelectField) {
                    var inputSelectFieldSelector = w2ui_utils.getInputSelectFieldSelector(inputSelectField.field);
                    // test.assertExists(inputSelectFieldSelector);

                    var isVisible = casper.evaluate(function selectFieldVisibility(selectField) {
                        return isVisibleInViewPort(document.querySelector(selectField));
                    }, inputSelectFieldSelector);

                    test.assert(isVisible, "{0} field is visible to remote screen".format(inputSelectField.field));

                    var inputSelectFieldValue = casper.evaluate(function (inputSelectFieldSelector) {
                        return document.querySelector(inputSelectFieldSelector).value;
                    }, inputSelectFieldSelector);

                    // that.formFields.record[inputSelectField.field].text

                    var defaultValueInW2UI = casper.evaluate(function getDefaultValue(form, field) {
                        return w2ui[form].record[field].text;
                    }, that.form, inputSelectField.field);

                    if (inputSelectFieldValue === defaultValueInW2UI) {
                        test.assert(true, "{0} have default value {1}".format(inputSelectField.field, defaultValueInW2UI));
                    } else {
                        test.assert(false, "{0} have different default value {1}.".format(inputSelectField.field, defaultValueInW2UI));
                    }
                });

                // Form Button rendering test
                that.buttonName.forEach(function (btnName) {

                    var isVisible = casper.evaluate(function formButtonVisibility(btnNode) {
                        return isVisibleInViewPort(document.querySelector(btnNode));
                    }, w2ui_utils.getW2UIButtonReferanceSelector(btnName));

                    test.assert(isVisible, "[{0}] is visible to remote screen.".format(btnName));
                });

                // Check box rendering test
                that.checkboxes.forEach(function (checkbox) {
                    var isVisible = casper.evaluate(function checkBoxvisibility(checkboxSelector) {
                        return isVisibleInViewPort(document.querySelector(checkboxSelector));
                    }, w2ui_utils.getCheckBoxSelector(checkbox.field));

                    test.assert(isVisible, "[{0}] is visible to remote screen.".format(checkbox.field));

                    //Test default value
                    var isChecked = casper.evaluate(function isChecked(checkboxSelector) {
                        return document.querySelector(checkboxSelector).checked;
                    }, w2ui_utils.getCheckBoxSelector(checkbox.field));

                    var isCheckedInW2UI = casper.evaluate(function isChecked(form, field) {
                        return w2ui[form].record[field];
                    }, that.form, checkbox.field);


                    test.assertEquals(isChecked, isCheckedInW2UI, "{0} checked is {1}".format(checkbox.field, isChecked));


                    // TODO: Check that checkboxes are disabled/enabled as per the default value
                    // var isDisable = casper.evaluate(function isChecked(checkboxSelector) {
                    //     return document.querySelector(checkboxSelector).disabled;
                    // }, w2ui_utils.getCheckBoxSelector(checkbox.field));
                    //
                    // test.assertEquals(isDisable, checkbox.el.disable, "{0} disabled: {1}".format(checkbox.id, isDisable));

                });

                // Form field rendering
                common.capture(that.capture);

                test.done();
            });
        }
    });
};