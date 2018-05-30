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

        //--------------------------------
        //  TASK LIST 
        //--------------------------------
        if (data.record.TLID === 0) {
            ctl = 'No TaskList defined. You must set a TaskList for ' + BUD + ' to enable Close Period.';
            lcp = '-';
            cp = '-';
        } else {
            ctl = data.record.TLName + ' &nbsp;&nbsp;';
            ltl = dtFormatISOToW2ui(data.record.LastDtDone);
            if (ltl.length === 0) {
                ctl += "(no completed instances yet)";
            } else {
                ctl += "(last completion: " + ltl + ")";
            }
        }

        //--------------------------------
        //  Last closed period 
        //--------------------------------
        lcp = dtFormatISOToW2ui(data.record.LastDtClose);

        //--------------------------------
        //  Target close period
        //--------------------------------
        cp = dtFormatISOToW2ui(data.record.CloseTarget);

        document.getElementById("closePeriodTL").innerHTML = ctl;
        document.getElementById("closePeriodLCP").innerHTML = lcp;
        document.getElementById("closePeriodNCP").innerHTML = cp;

        //--------------------------------
        //  Submit button
        //--------------------------------
        var disable = !(closePeriodData.DtDone !== null && closePeriodData.DtDone.getFullYear() > 1999);
        document.getElementById("closePeriodSubmit").disabled = disable;
    })
    .fail(function(/*data*/){
        console.log("Get close period info failed.");
        return;
    });
};
