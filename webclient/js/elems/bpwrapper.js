/*global
	buildDepositElements, buildAppLayout, buildSidebar, buildAllocFundsGrid, buildAccountElements,
    buildTransactElements, buildRentableTypeElements, buildRentableElements,
    buildRAElements, buildRAPayorPicker, buildRUserPicker, buildRentablePicker,
    buildRAPicker, buildReceiptElements, buildAssessmentElements, buildExpenseElements,
    buildARElements, buildPaymentTypeElements, buildDepositoryElements, buildDepositElements,
    buildStatementsElements, buildReportElements, buildLedgerElements, buildTWSElements,
    buildDepositMethodElements, buildPayorStatementElements, buildRentRollElements, buildLoginForm,
    buildAppLayout, buildROVReceiptElements,
*/

"use strict";

// buildPageElementsWrapper calls all the routines that build UI
// elements.
//
// INPUTS:
//  uitype - 0 - standard, full-featured, Roller interface
//           1 - the Receipt-Only version of Roller
//
// RETURNS:
//  nothing
function buildPageElementsWrapper(uitype) {
    buildAppLayout();
    buildSidebar(uitype);
    buildAllocFundsGrid();
    buildAccountElements();
    buildTransactElements();
    buildRentableTypeElements();
    buildRentableElements();
    buildRAElements();
    buildRAPayorPicker();
    buildRUserPicker();
    buildRentablePicker();
    buildRAPicker();
    switch (uitype) {
        case 0: buildReceiptElements(uitype); break;
        case 1: buildROVReceiptElements(uitype); break;
    }
    buildAssessmentElements();
    buildExpenseElements();
    buildARElements();
    buildPaymentTypeElements();
    buildDepositoryElements();
    buildDepositElements();
    buildStatementsElements();
    buildReportElements();
    buildLedgerElements();
    buildTWSElements();
    buildDepositMethodElements();
    buildPayorStatementElements();
    buildRentRollElements();
    buildLoginForm();
}
