"use strict";

var common = require("./common.js");
var w2ui_utils = require("./w2ui_utils.js");

function testFieldVisibilityAndValue(fieldSelector, test, formField, fieldValueInDOM, fieldValueInW2UI) {
// get visibility status  of input field in viewport
    var isVisible = casper.evaluate(function getFieldVisibility(fieldSelector) {
        return isVisibleInViewPort(document.querySelector(fieldSelector));
    }, fieldSelector);

    // Check visibility of input field
    test.assert(isVisible, "{0} input field is visible to remote screen.".format(formField.field));

    // Check default value must be blank
    test.assertEquals(fieldValueInDOM.toString(), fieldValueInW2UI.toString(), "{0} field have default value {1}".format(formField.field, fieldValueInW2UI.toString()));
}

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

    function testInputs(that, w2uiFormRecords, pageNo, haveTabs, test) {

        that.formFields.forEach(function (formField) {

            // fieldSelector to select element from the DOM
            var fieldSelector = "#{0}".format(formField.field);

            // default value of field in W2UI object
            var fieldValueInW2UI = w2uiFormRecords[formField.field];

            // value of field in DOM
            var fieldValueInDOM = casper.evaluate(function (fieldSelector) {
                return document.querySelector(fieldSelector).value;
            }, fieldSelector);


            switch (formField.type) {

                // Update inpurFieldValue of input field type is money. Because default value of money type field is $0.00.
                // Here money prefix can be any thing $, Rs, etc.
                // To make generic replace $,.,0 with blank string ""
                case "money":
                    fieldValueInDOM = w2ui_utils.getUpdatedInputFieldValueForMoneyTypeField(fieldValueInDOM);
                    break;

                // get default value for list type field from w2ui object
                case "list":
                    fieldValueInW2UI = fieldValueInW2UI.text;
                    break;

                case "checkbox":
                    // get value of field from the DOM
                    fieldValueInDOM = casper.evaluate(function (fieldSelector) {
                        return document.querySelector(fieldSelector).checked;
                    }, fieldSelector);
                    break;

                default:
            }

            if (haveTabs) {
                if ((formField.html.page === pageNo) && (!formField.isHidden)) {
                    // tests for tabs
                    testFieldVisibilityAndValue(fieldSelector, test, formField, fieldValueInDOM, fieldValueInW2UI);
                }
            } else if (!formField.isHidden) {
                // test for without tabs
                testFieldVisibilityAndValue(fieldSelector, test, formField, fieldValueInDOM, fieldValueInW2UI);
            }
        });
    }

    casper.test.begin(testName, testCount, {

        //do basic setup first
        setUp: function (/*test*/) {

            // form name
            this.form = addNewButtonConfig.form;

            // grid name
            this.grid = addNewButtonConfig.grid;

            // sidebar id to open a grid
            this.sidebarID = addNewButtonConfig.sidebarID;

            // Click on side bar node
            casper.click("#" + w2ui_utils.getSidebarID(this.sidebarID));

            casper.log('[FormTest] [{0}] sidebar node clicked with ID: "{1}"'.format(this.grid, this.sidebarID), 'debug', common.logSpace);

            // Click add new button in toolbar
            casper.wait(common.waitTime, function clickAddNewButton() {
                // It will click table cell with the text 'Add New'
                casper.clickLabel("Add New", "td");
            });


            // list of input fields
            this.formFields = casper.evaluate(function (form) {
                var formFields = w2ui[form].fields;

                // add isHidden key with default value true
                formFields.forEach(function (formField) {
                    formField.isHidden = true;
                });

                return formFields;
            }, this.form);

            // list of tabs in form
            this.tabs = addNewButtonConfig.tabs;

            // capture name
            this.capture = addNewButtonConfig.capture;

            // Button name and class
            this.buttonNames = addNewButtonConfig.buttonName;

            // Disable fields list
            this.disableFields = addNewButtonConfig.disableFields;
        },

        test: function (test) {

            var that = this;

            casper.wait(common.waitTime, function testRightPanel() {

                // Right panel rendering
                testRightPanelRendering(test);

                // Capture the screen shot of viewport
                common.capture(that.capture);

                // get W2UI from record
                var w2uiFormRecords = casper.evaluate(function (formName) {
                    return w2ui[formName].record;
                }, that.form);

                // Update isHidden key of all fields as per the field's type in DOM
                that.formFields = casper.evaluate(function (formFields) {
                    formFields.forEach(function (formField) {

                        // get field type in DOM
                        var formFieldTypeInDOM = document.querySelector("#{0}".format(formField.field)).type;

                        // Update isHidden key as per field's type in DOM
                        if (formFieldTypeInDOM !== "hidden") {
                            formField.isHidden = false;
                        }
                    });

                    return formFields;
                }, that.formFields);

                // BUD Field Test
                testBUDField(test);

                // If there are no tabs but there is only one form than perform tests on fields of the form.
                if (that.tabs.length === 0) {

                    // test all fields of w2ui form
                    testInputs(that, w2uiFormRecords, 0, false, test);
                }

                that.tabs.forEach(function (tab, pageNo) {

                    // If there are tabs in form than approach to perform test on form will be difference.
                    // W2UI Form fields have html page number. Filter records based on page number and use it to perform tests

                    // First tab has been opened already. So don't click on first tab.
                    if (pageNo !== 0) {
                        casper.clickLabel(tab, "div");
                    }

                    casper.wait(common.waitTime, function testTabs(tab, pageNo) {

                        common.capture("tab_{0}_{1}.jpg".format(tab, pageNo));

                        testInputs(that, w2uiFormRecords, pageNo, true, test);
                    }(tab, pageNo));

                });

                // Form Button rendering test
                testButtons(that.buttonNames, test);

                // Right panel after close button
                testCloseRightPanel.call(this, test);
            });
        }
    });
};
