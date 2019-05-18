/*global
    w2ui,getCurrentBID,loadClosePeriodInfo,loadClosePeriodInfo,dtFormatISOToW2ui,errMsgHTML,
    successMsgHTML,cpMsg,
*/
"use strict";


// errMsgHTML - format error message
//------------------------------------------------------------
window.errMsgHTML = function(errmsg) {
    var s;
    if (errmsg.length > 0 ) {
        s = '<p style="background-color: #ffe0e0;color: #aa2222;border-color:#aa2222;border-style:solid;border-width: 1px 1px 1px 6px;"><br>&nbsp;&nbsp;' +
            '<i class="fas fa-exclamation-circle fa-2x"></i> &nbsp;&nbsp;' +
             errmsg + "<br>&nbsp;</p>";
    } else {
        s = "";
    }
    return s;
};

// successMsgHTML - format normal message
//------------------------------------------------------------
window.successMsgHTML = function(msg) {
    var s;
    if (msg.length > 0 ) {
        s = '<p style="background-color: #e0ffe0;color: #22aa22;border-color:#22aa22;border-style:solid;border-width: 1px 1px 1px 6px;"><br>&nbsp;&nbsp;' +
            '<i class="fas fa-check-circle fa-2x"></i> &nbsp;&nbsp;' +
             msg + "<br>&nbsp;</p>";
    } else {
        s = "";
    }
    return s;
};

// warnMsgHTML - format normal message
//------------------------------------------------------------
window.warnMsgHTML = function(msg) {
    var s;
    if (msg.length > 0 ) {
        s = '<p style="background-color: #ffffa0;color: #707010;border-color:#707010;border-style:solid;border-width: 1px 1px 1px 6px;"><br>&nbsp;&nbsp;' +
            '<i class="fas fa-exclamation-triangle fa-2x"></i> &nbsp;&nbsp;' +
             msg + "<br>&nbsp;</p><p>&nbsp;</p>";
    } else {
        s = "";
    }
    return s;
};

// feedbackMessage - set the innerHTML of the supplied message
// html node to the supplied message string.
//
// INPUTS
//   msgArea = name of HTML node on which to set message
//   msg     = the message to display
//------------------------------------------------------------
window.feedbackMessage = function(msgArea,msg) {
    var x = document.getElementById(msgArea);
    if (x != null) {
        x.innerHTML = msg;
    }
};
