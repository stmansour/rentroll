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
    buildClosePeriodElements,buildRAFlowElements,buildBusinessElements,
    finishBizForm, buildReservationsElements, finishReservationsForm,
    buildHelpElements, finishHelpSystem, buildAboutElements, finishAboutSystem,
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
//-----------------------------------------------------------------
window.buildPageElementsWrapper = function (uitype) {
    buildAppLayout();
    buildSidebar(uitype);
    buildAllocFundsGrid();
    buildAccountElements();
    buildAboutElements();
    buildBusinessElements();
    buildTransactElements();
    buildRentableTypeElements();
    buildRentableElements();
    buildRAFlowElements();
    buildRAPicker();
    switch (uitype) {
        case 0: buildReceiptElements(uitype); break;
        case 1: buildROVReceiptElements(uitype); break;
    }
    buildAssessmentElements();
    buildExpenseElements();
    buildReservationsElements();
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
    buildTaskListElements();
    buildTaskListDefElements();
    buildHelpElements();
    finishForms();
};

// finishForms is something that needs to be done after all the
// UI elements have been created. In particular, we use this call
// to add UI elements to w2layout objects.  I'm still not sure why
// we need to wait to do this, but I do know that if we add these
// elements to the layouts right after the elements have been
// created then it doesn't work. By waiting a little bit, it all
// seems to work.
//
// INPUTS:
//
// RETURNS:
//  nothing
//-----------------------------------------------------------------
window.finishForms = function () {
    createStmtForm();
    createPayorStmtForm();
    createDepositForm();
    finishTaskListForm();
    finishTLDForm();
    finishBizForm();
    finishReservationsForm();
    finishHelpSystem();
    finishAboutSystem();
};
