/* This module is for updating the app datastructure */

/* global
    $,getBUDfromBID,
*/

"use strict";

//-------------------------------------------------------------------------------
// UpdateCloseInfo - This routine calls the closeinfo web service and updates
//     the global app datastruct with the latest info.
//
// INPUTS:
//     bid  - the business id
//
// RETURNS:
//     nothing at this time
//-------------------------------------------------------------------------------
window.UpdateCloseInfo = function (bid) {
    var url = "/v1/closeinfo/" + bid;
    $.get(url)
        .done( function(data, textStatus, jqXHR) {
            var BUD = getBUDfromBID(data.BID);
            app.CloseInfo[BUD] = data;
        })
        .fail( function( jqXHR, textStatus, errorThrown ) {
            console.log(jqXHR);
            console.log(textStatus);
            console.log(errorThrown);
        });
};
