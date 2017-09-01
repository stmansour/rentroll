/*global
	app, w2ui, buildDepositElements, buildDepositMethodElements
    buildAppLayout, buildSidebar, buildAllocFundsGrid, buildAccountElements, 
    buildTransactElements, buildRentableTypeElements, buildRentableElements, 
    buildRAElements, buildRAPayorPicker, buildRUserPicker, buildRentablePicker, 
    buildRASelect, buildReceiptElements, buildAssessmentElements, buildExpenseElements, 
    buildARElements, buildPaymentTypeElements, buildDepositoryElements, buildDepositElements, 
    buildStatementsElements, buildReportElements, buildLedgerElements, buildTWSElements, 
    buildDepositMethodElements,buildPayorStatementElements,
*/

"use strict";

function buildPageElementsWrapper() {

    buildAppLayout();
    buildSidebar();
    buildAllocFundsGrid();
    buildAccountElements();
    buildTransactElements();
    buildRentableTypeElements();
    buildRentableElements();
    buildRAElements();
    buildRAPayorPicker();
    buildRUserPicker();
    buildRentablePicker();
    buildRASelect();
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
}