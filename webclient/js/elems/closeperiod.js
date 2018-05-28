/*global
    w2ui,getCurrentBID,loadClosePeriodInfo,loadClosePeriodInfo,
*/
"use strict";

var closePeriodData = {
    record: null,
    dtDone: null,

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
        var s = "";
        if (data.status === "error") {
            console.log('error = ' + data.message);
            return;
        }

        //--------------------------------------
        // Keep a local copy of the data record
        //--------------------------------------
        closePeriodData.record = data.record;
        closePeriodData.DtDone = new Date(data.record.DtDone);

        //--------------------------------
        //  TASK LIST 
        //--------------------------------
        if (data.record.TLID === 0) {
            s = 'No TaskList defined. You must set a TaskList for ' + BUD + ' to enable Close Period.';
        } else {
            s = data.record.TLName + ' ';
        }
        document.getElementById("closePeriodTL").innerHTML = s;

        //--------------------------------
        //  Last closed period 
        //--------------------------------
        //--------------------------------
        //  Closed period 
        //--------------------------------
        var disable = !(closePeriodData.DtDone !== null && closePeriodData.DtDone.getFullYear() > 1999);
        document.getElementById("closePeriodSubmit").disabled = disable;
    })
    .fail(function(/*data*/){
        console.log("Get close period info failed.");
        return;
    });


};
