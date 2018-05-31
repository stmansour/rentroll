/*global
    w2ui,getCurrentBID,loadClosePeriodInfo,loadClosePeriodInfo,dtFormatISOToW2ui,
*/
"use strict";

var closePeriodData = {
    record: null,
    dtDone: null,
    dtLastClose: null,

};

window.switchToClosePeriod = function() {
    // w2ui.toplayout.load('main', w2ui.closePeriodLayout);
	w2ui.toplayout.load('main', '/webclient/html/cpinfo.html');
	w2ui.toplayout.hide('right',true);
    loadClosePeriodInfo();
};

//-----------------------------------------------------------------------------
// loadClosePeriodInfo - a layout in which we place an html page
// and a form.
//
// @params
//
// @returns
//-----------------------------------------------------------------------------
window.loadClosePeriodInfo = function () {
    var BID = getCurrentBID();
    var BUD = getBUDfromBID(BID);
    var params = {cmd: 'get' };
    var dat = JSON.stringify(params);

    // delete Depository request
    $.post('/v1/closeperiod/'+BID, dat, null, "json")
    .done(function(data) {
        var ctl = "";
        var lcp = "";
        var ltl = "";
        var cp = "";
        var s = "";
        var bHaveCPTLID = false;         // does the business have a ClosePeriod TaskList
        var bHaveTargetTLID = false;     // is there an instance for this close period
        var bTargetTLCompleted = false;  // is the instance marked as completed

        if (data.status === "error") {
            console.log('error = ' + data.message);
            return;
        }

        //--------------------------------------
        // Keep a local copy of the data record
        //--------------------------------------
        closePeriodData.record = data.record;
        closePeriodData.DtDone = new Date(data.record.DtDone);
        closePeriodData.DtLastClose = new Date(data.record.LastDtClose);
        var r = closePeriodData.record;

        //--------------------------------
        //  TASK LIST 
        //--------------------------------
        if (r.TLID === 0) {
            ctl = 'No TaskList defined. You must set a TaskList for ' + BUD + ' to enable Close Period.';
        } else {
            bHaveCPTLID = true;
            ctl = r.TLName + ' &nbsp;&nbsp;';
            ltl = dtFormatISOToW2ui(r.LastDtDone);
            if (ltl.length === 0) {
                ctl += "(no completed instances yet)";
            } else {
                ctl += "(last completion: " + ltl + ")";
            }
        }
        document.getElementById("closePeriodTL").innerHTML = ctl;

        //--------------------------------
        //  Last closed period 
        //--------------------------------
        lcp = dtFormatISOToW2ui(r.LastDtClose);
        document.getElementById("closePeriodLCP").innerHTML = lcp;

        //--------------------------------
        //  Target close period
        //--------------------------------
        cp = dtFormatISOToW2ui(r.CloseTarget);
        document.getElementById("closePeriodNCP").innerHTML = cp;

        //--------------------------------
        //  Target close task list
        //--------------------------------
        if (r.TLIDTarget > 0) {
            bHaveTargetTLID = true;
            s = r.TLNameTarget + ' (' + r.TLIDTarget + ')';
            //var dtDue = new Date(r.DtDueTarget);
            var dtDone = new Date(r.DtDoneTarget);
            if (dtDone.getFullYear() > 1999) {
                bTargetTLCompleted = true;
                s += '  completed ' + dtFormatISOToW2ui(r.DtDoneTarget) + ' &nbsp;&nbsp;&#9989;';
            } else {
                s += '  not completed. &nbsp;Due on ' + dtFormatISOToW2ui(r.DtDueTarget) + ' &nbsp;&nbsp;&#10060;';
            }
        } else {
            s = "No task list instance for due date " + dtFormatISOToW2ui(r.DtDueTarget) + ' &nbsp;&nbsp;&#10060;';
        }
        document.getElementById("closeTargetTL").innerHTML = s;

        //--------------------------------
        //  Submit button
        //--------------------------------
        var disable = !(bHaveCPTLID && bHaveTargetTLID && bTargetTLCompleted);
        document.getElementById("closePeriodSubmit").disabled = disable;
    })
    .fail(function(/*data*/){
        console.log("Get close period info failed.");
        return;
    });
};

//-----------------------------------------------------------------------------
// submitClosePeriod is called when all the conditions of a close period are
// met and the user clicks the buttong to close the period.
//
// @params
//
// @returns
//-----------------------------------------------------------------------------
window.submitClosePeriod = function() {
    console.log('close the period');
};