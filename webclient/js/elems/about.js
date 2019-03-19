/*global
    w2ui, app, console, $, plural, switchToGrid, showReport, form_dirty_alert, loginPopupOptions, getAboutInfo,
    switchToClosePeriod, switchToReservations, feedbackMessage, errMsgHTML, successMsgHTML,
*/
"use strict";

//---------------------------------------------------------------------------------
// getAboutInfo - contacts the server to get info about its version, and updates
//          the version (about) html page
//
// @params  <none>
// @returns <none>
//---------------------------------------------------------------------------------
window.getAboutInfo = function () {
    $.when(
        getRollerVersion(),
        getRollerBuildDate(),
        getRollerBuildMachine()
    )
    .done(function(){
        // feedbackMessage("versionPageMessage",successMsgHTML('Successfully retrieved all data'));
    })
    .fail(function(){
        feedbackMessage("versionPageMessage",errMsgHTML('Failure getting version, build date, build machine'));
    });
};

function getRollerVersion() {
    return $.get('/v1/version/')
    .done( function(data) {
        if (typeof data == 'string') {  // it's weird, a successful data add gets parsed as an object, an error message does not
            document.getElementById("appVer").innerHTML = data;
        } else if (typeof data == "object" && data.status == "error") {
            feedbackMessage("versionPageMessage",errMsgHTML('return status: ' + data.status + ' : ' + data.message));
        }
    })
    .fail( function() {
        feedbackMessage("versionPageMessage",errMsgHTML('Error getting version'));
    });
}

function getRollerBuildDate() {
    return $.get('/v1/buildtime/')
    .done( function(data) {
        if (typeof data == 'string') {  // it's weird, a successful data add gets parsed as an object, an error message does not
            document.getElementById("buildDate").innerHTML = data;
        } else if (typeof data == "object" && data.status == "error") {
            feedbackMessage("versionPageMessage",errMsgHTML('return status: ' + data.status + ' : ' + data.message));
        }
    })
    .fail( function() {
        feedbackMessage("versionPageMessage",errMsgHTML('Error getting build date'));
    });
}

function getRollerBuildMachine() {
    return $.get('/v1/buildmachine/')
    .done( function(data) {
        if (typeof data == 'string') {  // it's weird, a successful data add gets parsed as an object, an error message does not
            document.getElementById("buildMachine").innerHTML = data;
        } else if (typeof data == "object" && data.status == "error") {
            feedbackMessage("versionPageMessage",errMsgHTML('return status: ' + data.status + ' : ' + data.message));
        }
    })
    .fail( function() {
        feedbackMessage("versionPageMessage",errMsgHTML('Error getting build machine'));
    });
}
