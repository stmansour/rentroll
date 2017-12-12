/*global
	buildDepositElements, buildAppLayout, buildSidebar, buildAllocFundsGrid, buildAccountElements, 
    buildTransactElements, buildRentableTypeElements, buildRentableElements, 
    buildRAElements, buildRAPayorPicker, buildRUserPicker, buildRentablePicker, 
    buildRAPicker, buildReceiptElements, buildAssessmentElements, buildExpenseElements, 
    buildARElements, buildPaymentTypeElements, buildDepositoryElements, buildDepositElements, 
    buildStatementsElements, buildReportElements, buildLedgerElements, buildTWSElements, 
    buildDepositMethodElements, buildPayorStatementElements, buildRentRollElements, buildLoginForm,
    buildAppLayout,
*/

"use strict";

// buildPageElementsWrapper calls all the routines that build UI 
// elements.
//
// INPUTS:
//  flag  -  if 0, do nothing special.  If 1, build the Receipt-Only version of Roller.
//           In this version, the sidebar will reduce the number of commands it exposes.
//
// RETURNS:
//  nothing
function buildPageElementsWrapper(flag) {
    buildAppLayout();
    buildSidebar(flag);
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
    buildReceiptElements();
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