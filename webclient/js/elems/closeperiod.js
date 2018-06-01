/*global
    w2ui,getCurrentBID,loadClosePeriodInfo,loadClosePeriodInfo,dtFormatISOToW2ui,errMsgHTML,
*/
"use strict";

var closePeriodData = {
    record: null,

};

window.switchToClosePeriod = function() {
    // w2ui.toplayout.load('main', w2ui.closePeriodLayout);
	w2ui.toplayout.load('main', '/webclient/html/cpinfo.html');
	w2ui.toplayout.hide('right',true);
    loadClosePeriodInfo();
};

window.errMsgHTML = function(errmsg) {
    var s;
    if (errmsg.length > 0 ) {
        s = '<p style="background-color: #ffe0e0;color: #ff2222;"><br>&nbsp;&nbsp;' +
            '<i class="fas fa-exclamation-circle fa-2x"></i> &nbsp;&nbsp;' +
             errmsg + "<br>&nbsp;</p>";
    } else {
        s = "";
    }
    return s;
};

//-----------------------------------------------------------------------------
// loadClosePeriodInfo - a layout in which we place an html page
// and a form.
//
// @params    errmsg - (optional) a string with an initial error message
//
// @returns
//-----------------------------------------------------------------------------
window.loadClosePeriodInfo = function (errmsg) {
    var BID = getCurrentBID();
    var BUD = getBUDfromBID(BID);
    var params = {cmd: 'get' };
    var dat = JSON.stringify(params);

    if (typeof errmsg == "undefined") {
        errmsg = "";
    }
    //------------------------------------------------------------------------
    // If we were called with an error message, let's get it up there now....
    //------------------------------------------------------------------------
    if (errmsg.length > 0 ) {
        document.getElementById("closePeriodError").innerHTML = errMsgHTML(errmsg);
    }

    // delete Depository request
    $.post('/v1/closeperiod/'+BID, dat, null, "json")
    .done(function(data) {

        var s = "";
        var bHaveCPTLID = false;         // does the business have a ClosePeriod TaskList
        var bHaveTargetTLID = false;     // is there an instance for this close period
        var bTargetTLCompleted = false;  // is the instance marked as completed

        if (data.status === "error") {
            document.getElementById("closePeriodError").innerHTML = errMsgHTML(data.message);
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
            s = 'No TaskList defined. You must set a TaskList for ' + BUD + ' to enable Close Period.';
        } else {
            bHaveCPTLID = true;
            s = r.TLName + ' &nbsp;&nbsp;';
            var ltl = dtFormatISOToW2ui(r.LastDtDone);
            if (ltl.length === 0) {
                s += "(no completed instances yet)";
            } else {
                s += "(last completion: " + ltl + ")";
            }
        }
        document.getElementById("closePeriodTL").innerHTML = s;

        //--------------------------------
        //  Last closed period 
        //--------------------------------
        s = dtFormatISOToW2ui(r.LastDtClose);
        if (s.length > 0 ) {
             s += ' &nbsp;&nbsp;<i class="fas fa-lock"></i>';
        }
        document.getElementById("closePeriodLCP").innerHTML = s;

        //--------------------------------
        //  Target close period
        //--------------------------------
        s = dtFormatISOToW2ui(r.CloseTarget);
        if (s.length > 0 ) {
             s += ' &nbsp;&nbsp;<i class="fas fa-lock-open"></i>';
        }
        document.getElementById("closePeriodNCP").innerHTML = s;

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
        var x = document.getElementById("closePeriodError");
        if (x !== null) {
            x.innerHTML = errMsgHTML("Get close period info failed.");
        }
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
    var BID = getCurrentBID();
    var params = {cmd: 'save', record: closePeriodData.record };
    var dat = JSON.stringify(params);

    // delete Depository request
    var url = '/v1/closeperiod/'+BID;
    $.post(url, dat, null, "json")
    .done( function(data) {
        if (data.status !== 'success') {
            loadClosePeriodInfo(data.message);
            return;
        }
        loadClosePeriodInfo();
    })
    .fail( function() {
        loadClosePeriodInfo('error with post to: ' + url);
    });
};



