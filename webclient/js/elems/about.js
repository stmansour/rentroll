/*global
    w2ui, app, console, $, plural, switchToGrid, showReport, form_dirty_alert, loginPopupOptions, getAboutInfo,
    switchToClosePeriod, switchToReservations, feedbackMessage, errMsgHTML, successMsgHTML,
*/
"use strict";

window.buildAboutElements = function () {
    //------------------------------------------------------------------------
    //  aboutLayout - The layout to contain the tabbed form and buttons
    //               top - build info
    //               main - session table grid
    //------------------------------------------------------------------------
    $().w2layout({
        name: 'aboutLayout',
        padding: 0,
        panels: [
            { type: 'left',    size: 0,     hidden: true },
            { type: 'top',     size: '20%', hidden: false, content: 'top',  resizable: true, style: app.pstyle },
            { type: 'main',    size: '60%', hidden: false, content: 'main', resizable: true, style: app.pstyle },
            { type: 'preview', size: 0,     hidden: true,  content: 'PREVIEW'  },
            { type: 'bottom',  size: 0,     hidden: true,  content: 'bottom', resizable: false, style: app.pstyle },
            { type: 'right',   size: 0,     hidden: true }
        ]
    });

    //------------------------------------------------------------------------
    //          SessionGrid  -  shows the server session table
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'sessionGrid',
        url: '/v1/sessions',
        header: 'Active Sessions',
        multiSelect: false,
        show: {
            toolbar         : false,
            header          : true,
            footer          : true,
            toolbarAdd      : false,   // indicates if toolbar add new button is visible
            toolbarDelete   : false,   // indicates if toolbar delete button is visible
            toolbarSave     : false,   // indicates if toolbar save button is visible
            selectColumn    : false,
            expandColumn    : false,
            toolbarEdit     : false,
            toolbarSearch   : false,
            toolbarInput    : false,
            searchAll       : false,
            toolbarReload   : true,
            toolbarColumns  : true,
        },
        columns: [
            {field: 'recid',    caption: 'recid',    hidden: true,  size: '40px', sortable: false },
            {field: 'Token',    caption: 'Token',    hidden: false, size: '200px', sortable: false },
            {field: 'Username', caption: 'Username', hidden: false, size: '75px', sortable: false },
            {field: 'Name',     caption: 'Name',     hidden: false, size: '100px', sortable: false },
            {field: 'UID',      caption: 'UID',      hidden: false, size: '50px', sortable: false },
            {field: 'CoCode',   caption: 'CoCode',   hidden: false, size: '50px', sortable: false },
            {field: 'Expire',   caption: 'Expire',   hidden: false, size: '125px', sortable: false },
            {field: 'RoleID',   caption: 'RoleID',   hidden: false, size: '50px', sortable: false },
            {field: 'ImageURL', caption: 'ImageURL', hidden: false, size: '20%', sortable: false },
        ],
    });
};

//---------------------------------------------------------------------------------
// setToAbout - displays the About UI when the user clicks on the About command
//              in the sidebar
//
// @params  <none>
// @returns <none>
//---------------------------------------------------------------------------------
window.setToAbout = function() {
    w2ui.toplayout.content('main', w2ui.aboutLayout);
    w2ui.toplayout.hide('right',true);
    w2ui.aboutLayout.load('top','/webclient/html/about.html');
    getAboutInfo();
};

//---------------------------------------------------------------------------------
// finishAboutSystem - loads the aboutLayout with the session grid
//
// @params  <none>
// @returns <none>
//---------------------------------------------------------------------------------
window.finishAboutSystem = function () {
    w2ui.aboutLayout.content('main',w2ui.sessionGrid);
};

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
