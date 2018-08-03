/*global
	buildDepositElements, buildAppLayout, buildSidebar, buildAllocFundsGrid,
    buildAccountElements, buildTransactElements, buildRentableTypeElements,
    buildRentableElements, buildRAPicker, buildReceiptElements,
    buildAssessmentElements, buildExpenseElements, buildARElements,
    buildPaymentTypeElements, buildDepositoryElements, buildDepositElements,
    buildStatementsElements, buildReportElements, buildLedgerElements,
    buildTWSElements, buildDepositMethodElements, buildPayorStatementElements,
    buildRentRollElements, buildLoginForm, buildAppLayout,
    buildROVReceiptElements,buildTaskListElements,buildTaskListDefElements,
    finishTaskListForm, createDepositForm, createPayorStmtForm,
    createStmtForm, finishForms, finishTLDForm,
    buildClosePeriodElements,buildRAApplicantElements,buildRA2FlowElements,
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
window.buildPageElementsWrapper = function (uitype) {
    buildAppLayout();
    buildSidebar(uitype);
    buildAllocFundsGrid();
    buildAccountElements();
    buildTransactElements();
    buildRentableTypeElements();
    buildRentableElements();
    buildRAApplicantElements();
    buildRAPicker();
    switch (uitype) {
        case 0: buildReceiptElements(uitype); break;
        case 1: buildROVReceiptElements(uitype); break;
    }
    buildAssessmentElements();
    buildExpenseElements();
    buildARElements();
    buildRA2FlowElements();
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
    buildTaskListElements();
    buildTaskListDefElements();
    finishForms();
};

window.finishForms = function () {
    createStmtForm();
    createPayorStmtForm();
    createDepositForm();
    finishTaskListForm();
    finishTLDForm();
};
