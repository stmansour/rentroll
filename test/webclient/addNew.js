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
    function testInputFields(formName, inputFields, test) {

        inputFields.forEach(function (inputField) {

            // console.log(JSON.stringify(inputField));
            console.log("In testInputFields");

            // get selector for the input field
            var inputFieldSelector = w2ui_utils.getInputFieldSelector(inputField.field);

            // get visibility status  of input field in viewport
            var isVisible = casper.evaluate(function inputFieldVisibility(inputFieldSelector) {
                return isVisibleInViewPort(document.querySelector(inputFieldSelector));
            }, inputFieldSelector);

            // Check visibility of input field
            test.assert(isVisible, "{0} input field is visible to remote screen.".format(inputField.field));

            // get value of input field from the DOM
            var inputFieldValue = casper.evaluate(function (inputFieldSelector) {
                return document.querySelector(inputFieldSelector).value;
            }, inputFieldSelector);

            // get default value of input field from the W2UI object
            var inputFieldValueInW2UI = casper.evaluate(function (formName, field) {
                return w2ui[formName].record[field];
            }, formName, inputField.field);

            // Update inpurFieldValue of input field type is money. Because default value of money type field is $0.00.
            // Here money prefix can be any thing $, Rs, etc.
            // To make generic replace $,.,0 with blank string ""
            if (inputField.type === "money") {
                inputFieldValue = inputFieldValue.replace(/[$.]/g, "");
            }

            // Check default value must be blank
            test.assertEquals(inputFieldValue, inputFieldValueInW2UI.toString(), "{0} field is blank".format(inputField.field));
        });
    }

    // Test visible input select fields of the form
    function testInputSelectField(formName, inputSelectField, test) {

        inputSelectField.forEach(function (inputSelectField) {

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
            }, formName, inputSelectField.field);

            // match default value with input field value in DOM
            test.assertEquals(inputSelectFieldValue, defaultValueInW2UI, "{0} have default value {1}".format(inputSelectField.field, defaultValueInW2UI));
        });
    }


    // Test buttons in form
    function testButtons(buttonNames, test) {

        buttonNames.forEach(function (btnName) {

            // get visibility status of button in viewport
            var isVisible = casper.evaluate(function formButtonVisibility(btnNode) {
                return isVisibleInViewPort(document.querySelector(btnNode));
            }, w2ui_utils.getW2UIButtonReferanceSelector(btnName));

            // Check visibility of button
            test.assert(isVisible, "[{0}] is visible to remote screen.".format(btnName));
        });
    }

    // Test checkboxes in form
    function testCheckBoxes(formName, checkboxes, test) {

        checkboxes.forEach(function (checkbox) {

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
            }, formName, checkbox.field);

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
    function testDateFields(formName, dateFields, test) {
        dateFields.forEach(function (dateField) {

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
            }, formName, dateField.field);

            test.assertEquals(defaultValue, defaultValueInW2UI, "{0} value is {1}".format(dateField.field, defaultValueInW2UI));
        });
    }

    // Disabled field rendering test
    function testDisabledFields(disableFields, test) {

        // We are expecting these fields must be disabled. Select that fields and check disable attribute
        disableFields.forEach(function (disableField) {

            var disabilityInDOM = casper.evaluate(function checkDisability(disableFieldSelector) {
                return document.querySelector(disableFieldSelector).disabled;
            }, w2ui_utils.getDisableFieldSelector(disableField));

            test.assert(disabilityInDOM, "{0} is disabled.".format(disableField));
        });
    }

    // test Right panel after close button
    function testCloseRightPanel(test) {

        casper.thenClick(w2ui_utils.getCloseButtonSelector());

        casper.wait(common.waitTime, function () {
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

            // list of tabs in form
            this.tabs = addNewButtonConfig.tabs;

            // list of input fields
            this.inputFields = this.formFields.filter(w2ui_utils.getTextTypeW2UIFields);

            // list of input select fields
            this.inputSelectField = this.formFields.filter(w2ui_utils.getInputListW2UIFields);

            // capture name
            this.capture = addNewButtonConfig.capture;

            // Button name and class
            this.buttonNames = addNewButtonConfig.buttonName;

            // Checkboxes list
            this.checkboxes = this.formFields.filter(w2ui_utils.getCheckBoxW2UIFields);

            // Date fields list
            this.dateFields = this.formFields.filter(w2ui_utils.getDateW2UIFields);

            // Disable fields list
            this.disableFields = addNewButtonConfig.disableFields;

            // updateFieldsMemberInList()

            casper.click("#" + w2ui_utils.getSidebarID(this.sidebarID));
            casper.log('[FormTest] [{0}] sidebar node clicked with ID: "{1}"'.format(this.grid, this.sidebarID), 'debug', common.logSpace);

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

                // var tabSelector = w2ui_utils.getW2UITabSelector(that.form, );

                // BUD Field Test
                testBUDField(test);

                that.tabs.forEach(function (tab, pageNo) {

                    // pageNo += 1;

                    // If there are tabs in form than approach to perform test on form will be difference.
                    // W2UI Form fields have html page number. Filter records based on page number and use it to perform tests
                    // Update members in inputFields, inputSelectFields, checkBoxes, dateFields, and other fields list

                    var inputFields = that.inputFields.filter(w2ui_utils.getFieldsForPage, pageNo);
                    var inputSelectField = that.inputSelectField.filter(w2ui_utils.getFieldsForPage, pageNo);
                    var checkboxes = that.checkboxes.filter(w2ui_utils.getFieldsForPage, pageNo);
                    var dateFields = that.dateFields.filter(w2ui_utils.getFieldsForPage, pageNo);
                    var disableFields = that.disableFields.filter(w2ui_utils.getFieldsForPage, pageNo);


                    // First tab has been opened already. So don't click on first tab.
                    if (pageNo !== 0){
                        casper.clickLabel(tab, "div");
                    }

                    casper.wait(common.waitTime, function testTabs(inputFields, tab, pageNo) {

                        common.capture("tab_{0}_{1}.jpg".format(tab, pageNo));

                        // Input fields test
                        testInputFields(that.form, inputFields, test);

                        // Drop down Input fields test
                        testInputSelectField(that.form, inputSelectField, test);

                        // Check box rendering test
                        testCheckBoxes(that.form, checkboxes, test);

                        // Date field rendering test
                        testDateFields(that.form, dateFields, test);

                        // Disabled field rendering test
                        testDisabledFields(disableFields, test);

                    }(inputFields, tab, pageNo));

                });

                // If there are no tabs but there is only one form than perform tests on fields of the form.
                if (that.tabs.length === 0){

                    console.log("No tabs!");

                    // Input fields test
                    testInputFields(that.form, that.inputFields, test);

                    // Dropdown Input fields test
                    testInputSelectField(that.form, that.inputSelectField, test);

                    // Check box rendering test
                    testCheckBoxes(that.form, that.checkboxes, test);

                    // Date field rendering test
                    testDateFields(that.form, that.dateFields, test);

                    // Disabled field rendering test
                    testDisabledFields(that.disableFields, test);
                }


                // Form Button rendering test
                testButtons(that.buttonNames, test);

                // Right panel after close button
                testCloseRightPanel.call(this, test);

                // Capture the screen shot of viewport
                common.capture(that.capture);
            });
        }
    });
};
