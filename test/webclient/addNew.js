"use strict";

var common = require("./common.js");
var w2ui_utils = require("./w2ui_utils.js");

exports.w2uiAddNewButtonTest = function (addNewButtonConfig) {
    var testCount = addNewButtonConfig.testCount;
    var testName = "w2ui add new button [{0}] test".format(addNewButtonConfig.form);

    // Test Right panel rendering
    function testRightPanelRendering(test) {
        test.assertExists("#" + w2ui_utils.getRightPanelID());
    }

    // Test BUD Field in form
    function testBUDField(test) {

        // Check BUD field is exists in DOM
        test.assertSelectorExists(w2ui_utils.getBUDSelector());

        // Get BUD value from the DOM
        var businessUnitValue = casper.evaluate(function getBusinessUnit(bud_selector) {
            return document.querySelector(bud_selector).value;
        }, w2ui_utils.getBUDSelector());

        // Match DOM's BUD value with testBiz
        test.assertEquals(businessUnitValue, testBiz, "Business Unit value is {0}.".format(businessUnitValue));

        // Get disable attribute of BUD from the DOM.
        var isBusinessUnitValueDisabled = casper.evaluate(function (bud_selector) {
            return document.querySelector(bud_selector).disabled;
        }, w2ui_utils.getBUDSelector());

        // Check BUD disability
        test.assert(isBusinessUnitValueDisabled, "Disability of business unit field.");
    }

    // Test visible input fields of the form
    function testInputFields(that, test) {

        that.inputFields.forEach(function (inputFieldID) {

            // get selector for the input field
            var inputFieldSelector = w2ui_utils.getInputFieldSelector(inputFieldID.field);

            // get visibility status  of input field in viewport
            var isVisible = casper.evaluate(function inputFieldVisibility(inputFieldSelector) {
                return isVisibleInViewPort(document.querySelector(inputFieldSelector));
            }, inputFieldSelector);

            // Check visibility of input field
            test.assert(isVisible, "{0} input field is visible to remote screen.".format(inputFieldID.field));

            // get value of input field from the DOM
            var inputFieldValue = casper.evaluate(function (inputFieldSelector) {
                return document.querySelector(inputFieldSelector).value;
            }, inputFieldSelector);

            // Check default value must be blank
            test.assertEquals(inputFieldValue, "", "{0} field is blank".format(inputFieldID.field));
        });
    }

    // Test visible input select fields of the form
    function testInputSelectField(that, test) {

        that.inputSelectField.forEach(function (inputSelectField) {

            // get selector for the input select field
            var inputSelectFieldSelector = w2ui_utils.getInputSelectFieldSelector(inputSelectField.field);

            // get visibility status  of input select field in viewport
            var isVisible = casper.evaluate(function selectFieldVisibility(selectField) {
                return isVisibleInViewPort(document.querySelector(selectField));
            }, inputSelectFieldSelector);

            // Check visibility of input select field
            test.assert(isVisible, "{0} field is visible to remote screen".format(inputSelectField.field));

            // get value of input select field from the DOM
            var inputSelectFieldValue = casper.evaluate(function (inputSelectFieldSelector) {
                return document.querySelector(inputSelectFieldSelector).value;
            }, inputSelectFieldSelector);

            // get value of input select field from W2UI form record
            var defaultValueInW2UI = casper.evaluate(function getDefaultValue(form, field) {
                return w2ui[form].record[field].text;
            }, that.form, inputSelectField.field);

            // match default value with input field value in DOM
            test.assertEquals(inputSelectFieldValue, defaultValueInW2UI, "{0} have default value {1}".format(inputSelectField.field, defaultValueInW2UI));
        });
    }


    // Test buttons in form
    function testButtons(that, test) {

        that.buttonName.forEach(function (btnName) {

            // get visibility status of button in viewport
            var isVisible = casper.evaluate(function formButtonVisibility(btnNode) {
                return isVisibleInViewPort(document.querySelector(btnNode));
            }, w2ui_utils.getW2UIButtonReferanceSelector(btnName));

            // Check visibility of button
            test.assert(isVisible, "[{0}] is visible to remote screen.".format(btnName));
        });
    }

    // Test checkboxes in form
    function testCheckBoxes(that, test) {

        that.checkboxes.forEach(function (checkbox) {

            // get visibility status of checkbox in viewport
            var isVisible = casper.evaluate(function checkBoxvisibility(checkboxSelector) {
                return isVisibleInViewPort(document.querySelector(checkboxSelector));
            }, w2ui_utils.getCheckBoxSelector(checkbox.field));

            // Check visibility of checkbox
            test.assert(isVisible, "[{0}] is visible to remote screen.".format(checkbox.field));

            // get status of checkbox from the DOM
            var isChecked = casper.evaluate(function isChecked(checkboxSelector) {
                return document.querySelector(checkboxSelector).checked;
            }, w2ui_utils.getCheckBoxSelector(checkbox.field));

            // get default status of checkbox from the W2UI
            var isCheckedInW2UI = casper.evaluate(function isChecked(form, field) {
                return w2ui[form].record[field];
            }, that.form, checkbox.field);

            // Match default value of checkbox with value in DOM
            test.assertEquals(isChecked, isCheckedInW2UI, "{0} checked is {1}".format(checkbox.field, isChecked));


            // TODO: Check that checkboxes are disabled/enabled as per the default value
            // var isDisable = casper.evaluate(function isChecked(checkboxSelector) {
            //     return document.querySelector(checkboxSelector).disabled;
            // }, w2ui_utils.getCheckBoxSelector(checkbox.field));
            //
            // test.assertEquals(isDisable, checkbox.el.disable, "{0} disabled: {1}".format(checkbox.id, isDisable));

        });
    }

    // Date field rendering test
    function testDateFields(that, test) {
        that.dateFields.forEach(function (dateField) {

            // Check visibility
            var isVisible = casper.evaluate(function checkDateFieldVisibility(dateFieldSelector) {
                return isVisibleInViewPort(document.querySelector(dateFieldSelector));
            }, w2ui_utils.getCheckBoxSelector(dateField.field));

            test.assert(isVisible, "[{0}] is visible to remote screen.".format(dateField.field));

            // Test default value
            var defaultValue = casper.evaluate(function getDefaultValue(dateFieldSelector) {
                return document.querySelector(dateFieldSelector).value;
            }, w2ui_utils.getCheckBoxSelector(dateField.field));

            var defaultValueInW2UI = casper.evaluate(function getDefaultValue(form, field) {
                return w2ui[form].record[field];
            }, that.form, dateField.field);

            test.assertEquals(defaultValue, defaultValueInW2UI, "{0} value is {1}".format(dateField.field, defaultValueInW2UI));
        });
    }

    // Disabled field rendering test
    function testDisabledFields(that, test) {

        // We are expecting these fields must be disabled. Select that fields and check disable attribute
        that.disableFields.forEach(function (disableField) {

            var disabilityInDOM = casper.evaluate(function checkDisability(disableFieldSelector) {
                return document.querySelector(disableFieldSelector).disabled;
            }, w2ui_utils.getDisableFieldSelector(disableField));

            test.assert(disabilityInDOM, "{0} is disabled.".format(disableField));
        });
    }

    // test Right panel after close button
    function testCloseRightPanel(test) {

        this.click(w2ui_utils.getCloseButtonSelector());

        this.wait(common.waitTime, function () {
            test.assertNotVisible("#" + w2ui_utils.getRightPanelID());
            test.done();
        });

    }

    casper.test.begin(testName, testCount, {

        //do basic setup first
        setUp: function (/*test*/) {

            //form name
            this.form = addNewButtonConfig.form;

            //grid name
            this.grid = addNewButtonConfig.grid;

            //to open a grid
            this.sidebarID = addNewButtonConfig.sidebarID;

            // list of input fields
            // this.inputFields = addNewButtonConfig.inputField;
            this.formFields = casper.evaluate(function (form) {
                return w2ui[form].fields;
            }, this.form);

            // list of input fields
            this.inputFields = this.formFields.filter(w2ui_utils.getTextTypeW2UIFields);

            // list of input select fields
            this.inputSelectField = this.formFields.filter(w2ui_utils.getInputListW2UIFields);

            // capture name
            this.capture = addNewButtonConfig.capture;

            // Button name and class
            this.buttonName = addNewButtonConfig.buttonName;

            // Checkboxes list
            this.checkboxes = this.formFields.filter(w2ui_utils.getCheckBoxW2UIFields);

            // Date fields list
            this.dateFields = this.formFields.filter(w2ui_utils.getDateW2UIFields);

            // Disable fields list
            this.disableFields = addNewButtonConfig.disableFields;

            casper.click("#" + w2ui_utils.getSidebarID(this.sidebarID));
            casper.log('[FormTest] [{0}] sidebar node clicked with ID: "{1}"'.format(this.grid, this.sidebarID), 'debug', logSpace);

            // Click add new button in toolbar
            casper.wait(common.waitTime, function clickAddNewButton() {
                //It will click table cell with the text 'Add New'
                casper.clickLabel("Add New", "td");
            });
        },

        test: function (test) {

            var that = this;

            casper.wait(common.waitTime, function testRightPanel() {

                // Right panel rendering
                testRightPanelRendering(test);

                // BUD Field Test
                testBUDField(test);

                // Input fields test
                testInputFields(that, test);

                // Dropdown Input fields test
                testInputSelectField(that, test);

                // Form Button rendering test
                testButtons(that, test);

                // Check box rendering test
                testCheckBoxes(that, test);

                // Date field rendering test
                testDateFields(that, test);

                // Disabled field rendering test
                testDisabledFields(that, test);

                // Right panel after close button
                testCloseRightPanel.call(this, test);

                // Capture the screen shot of viewport
                common.capture(that.capture);
            });
        }
    });
};