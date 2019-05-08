/* This module is for updating the app datastructure */

/* global
    $,getBUDfromBID,
*/

"use strict";

//-------------------------------------------------------------------------------
// UpdateCloseInfo - This routine calls the closeinfo web service and updates
//     the global app datastruct with the latest info.  It caches the info
//     for 3 seconds before making another call to update.  This is long
//     enough to keep multiple calls from happening on a single UI screen
//     update.
//
// INPUTS:
//     bid  - the business id
//
// RETURNS:
//     nothing at this time
//-------------------------------------------------------------------------------
window.UpdateCloseInfo = function (bid) {
    var now = new Date();
    var tm = now.getTime()/1000;

    if (app.CloseInfoInProg) {
        if (tm - app.CloseInfoReqStart > 120.0) {  // if > 2min, then reset
            app.CloseInfoInProg = false;
        } else {
            // console.log('CloseInfo - query in progress');
            return;
        }
    }

    var url = "/v1/closeinfo/" + bid;
    var diff = tm - app.CloseInfoTime;

    // console.log('diff = ' + diff + ', tm = ' + tm + ', app.CloseInfoTime = ' + app.CloseInfoTime);
    if (diff < 3.00 ) {
        // console.log('cache up to date!!!! No call to server.');
        return; // we cache for 3 seconds.  That's long enough
    }
    app.CloseInfoInProg = true;
    app.CloseInfoReqStart = tm;  // mark the time we started the query
    $.get(url)
        .done( function(data, textStatus, jqXHR) {
            var BUD = getBUDfromBID(data.BID);
            app.CloseInfo[BUD] = data;
            now = new Date();
            app.CloseInfoTime = now.getTime()/1000;
            // console.log('app.CloseInfoTime updated to: ' + app.CloseInfoTime);
            app.CloseInfoInProg = false;

        })
        .fail( function( jqXHR, textStatus, errorThrown ) {
            console.log(jqXHR);
            console.log(textStatus);
            console.log(errorThrown);
            
            app.CloseInfoInProg = false;
        });
};
