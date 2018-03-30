"use strict";

/*global
    $, console, app, w2ui, w2popup, setInterval, getCookieValue, triggerReceiptsGrid,
    deleteCookie, userProfileToUI, handleBlankScreen, ensureSession, startSessionChecker,
    getUserInfo, startNewSession, popupLoginDialogBox
*/

var loginRoURL = "/webclient/html/formlogin.html";
var loginRcURL = "/webclient/html/formrcptlogin.html";

var loginSessionChecker = {};

var loginPopupOptions = {
    body: '<div id="loginPopupForm" style="width: 100%; height: 100%;"></div>',
    style: 'padding: 4px 0px 0px 0px; overflow: auto;',
    width: 425,
    height: 525,
    showMax: true,
    modal: true,
    onOpen: function (event) {
        event.onComplete = function () {
            $('#w2ui-popup #loginPopupForm').w2render('passwordform');
        };
    },
};

window.userProfileToUI = function() {
    var name = app.name;
    if (name.length === 0 || app.uid === 0) { name = "?";}
    $("#user_menu_container").find("#username").text(name);
    $("#user_menu_container").find("img").attr("src", app.imageurl);

    // *******************
    // ONLY FOR ROV CLIENT
    // -------------------
    if (window.location.href.endsWith("/rhome/")) {
        setTimeout(function() {
            $('#node_receipts').trigger('click');
        }, 500); // wait for some time meanwhile left sidebar render done!
    }
};

window.buildLoginForm = function() {
    var loginTmplURL = loginRoURL;
    if (app.client == "receipts") {
        loginTmplURL = loginRcURL;
    }
    $().w2form({
        name: 'passwordform',
        formURL: loginTmplURL,
        style: 'border: 0px; background-color: transparent;',
        fields: [{field: 'user', type: 'text',     required: false, html: {caption: 'User Name' /*, attr: 'readonly'*/} },
                 {field: 'pass', type: 'password', required: false, html: {caption: 'Password'} },
                ],
        actions: {
            login: function (/*event*/) {
                var f = this;
                console.log('User Name = ' + w2ui.passwordform.record.user);
                app.username = w2ui.passwordform.record.user;

                // request login only username, password entered
                if (!(app.username && w2ui.passwordform.record.pass)) {
                    console.log("Both Username and Password must be supplied");
                        $(f.box).find("#LoginMessage").find(".errors").empty();
                        var message = "Both Username and Password must be supplied";
                        $(f.box).find("#LoginMessage").find(".errors").append("<p>" + message + "</p>");
                        $(f.box).find("#LoginMessage").removeClass("hidden");
                        // w2ui.passwordform.error(w2utils.lang(data.message));
                    return;
                }

                var params = {user: app.username, pass: w2ui.passwordform.record.pass };
                var dat = JSON.stringify(params);
                $.post('/v1/authn/', dat, null, "json")
                .done(function(data) {
                    if (data.status === "error") {
                        $(f.box).find("#LoginMessage").find(".errors").empty();
                        var message = "Unrecognized Username or Password";
                        $(f.box).find("#LoginMessage").find(".errors").append("<p>" + message + "</p>");
                        $(f.box).find("#LoginMessage").removeClass("hidden");
                        // w2ui.passwordform.error(w2utils.lang(data.message));
                    }
                    else if (data.status === "success") {
                        app.uid = data.uid;
                        app.name = data.Name;
                        app.imageurl = data.ImageURL;
                        $(f.box).find("#LoginMessage").addClass("hidden");
                        w2popup.close();
                        w2ui.passwordform.record.pass = ""; // after closing dialog, remove password information.
                        userProfileToUI();

                        // remove blank screen if login successfully
                        handleBlankScreen(true);
                    } else {
                        console.log("Login service returned unexpected status: " + data.status);
                    }
                    return;
                })
                .fail(function(/*data*/){
                    w2ui.passwordform.error("Login failed");
                    return;
                });
            },
            cancel: function (/*event*/) {
                // w2popup.close();
                return;
            }
        },
        onRefresh: function(event) {
            var f = this;
            event.onComplete = function() {

                // handle enter key press event
                $(f.box).keypress(function(keypressEvent) {
                    if (keypressEvent.which === 13) {
                        // need to give time so that w2ui form have data in it's record object
                        setTimeout(function() {
                            $(f.box).find("button[name=login]").click();
                        }, 100);
                    }
                });

                // TODO: handle forgot password link
                $(f.box).find("forgot_pass_link").click(function() {
                    console.log("forgot password link clicked!");
                });
            };
        }
    });
};

//---------------------------------------------------------------------------------
// startNewSession - encapsulates the steps needed to launch a new login session
//                 and start up a session checker to have the user log in again
//                 if the session expires
//
// @params  <none>
// @returns <none>
//---------------------------------------------------------------------------------
window.startNewSession = function () {
    ensureSession(); // get the user logged in
    startSessionChecker(); // have the user log in if the session expires
};

//---------------------------------------------------------------------------------
// getUserInfo - get the user profile information
//
// @params  <none>
// @returns <none>
//---------------------------------------------------------------------------------
window.getUserInfo = function () {
    $.get('/v1/userprofile/')
    .done(function(data, textStatus, jqXHR) {
        if (data.status == "success") {
            app.uid = data.uid;
            app.username = data.username;
            app.name = data.Name;
            app.imageurl = data.ImageURL;
            userProfileToUI();
        }
    })
    .fail( function() {
        console.log('Error getting /v1/userprofile');
    });
};
//---------------------------------------------------------------------------------
// launchSession - if a valid sessionid exists, use it and get user profile info
//                 if not, log in.
//
// @params  <none>
// @returns <none>
//---------------------------------------------------------------------------------
window.launchSession = function () {
    var x = getCookieValue("air");
    // console.log('launchSession: getCookieValue(air) = '+x);
    if (x !== null && x.length > 0) {
        getUserInfo();
        handleBlankScreen(true);
    }
    startNewSession();
};


//---------------------------------------------------------------------------------
// startSessionChecker - validate the session every 5 seconds
//
// @params  <none>
// @returns <none>
//---------------------------------------------------------------------------------
window.startSessionChecker = function () {
    loginSessionChecker = setInterval(
    function() {
        ensureSession();
    }, 5000); // watch out for session expiring
};

//---------------------------------------------------------------------------------
// handleBlankScreen - hide the content of screen if user is not logged in as in
//                     black blank screen with login popup otherwise it will be
//                     hidden
// @params  - isLoggedIn
// @returns <none>
//---------------------------------------------------------------------------------
window.handleBlankScreen = function (isLoggedIn) {
    if (isLoggedIn) {
        $("#blank_screen").hide();
    } else {
        $("#blank_screen").show();
        popupLoginDialogBox(); // if it's not logged in then show popup
    }
};

window.popupLoginDialogBox = function () {
    $().w2popup('open', loginPopupOptions);
    var f = w2ui.passwordform;
    if (f) {
        $(f.box).find("#LoginMessage").find(".errors").empty();
        var message = "Your session hass expired. Please login again.";
        $(f.box).find("#LoginMessage").find(".errors").append("<p>" + message + "</p>");
        $(f.box).find("#LoginMessage").removeClass("hidden");
    }
};

//---------------------------------------------------------------------------------
// ensureSession - check to see if we have our session cookie.  If not, we need to
//             authenticate.
//
// @params  <none>
// @returns <none>
//---------------------------------------------------------------------------------
window.ensureSession = function () {
    if (w2popup.status == "open") {return;} // just return now if we're trying to log in

    var c = getCookieValue("air");          // Do we have an "air" cookie?
    if (c === null || c.length < 20) {   // if not...
        deleteCookie("air");
        handleBlankScreen(false);
        return;
    }

    handleBlankScreen(true);        // make sure we can see the interface
    if (app.name.length === 0) {    // if we don't have user info
        getUserInfo();              // then get it
    }
};

//---------------------------------------------------------------------------------
// logoff - sign out of the current session
//
// @params  <none>
// @returns <none>
//---------------------------------------------------------------------------------
window.logoff = function () {
    app.uid = 0;
    app.name = "";
    app.username = "";
    app.imageurl = "";
    $.get('/v1/logoff/')
    .done(function(data, textStatus, jqXHR) {
        if (jqXHR.status == 200) {
            console.log('logoff success, app.uid set to 0.');
            ensureSession();
        } else {
            console.log( '**** YIPES! ****  status on logoff = ' + textStatus);
        }
        deleteCookie("air");  // no matter what, delete the cookie after this call completes
    })
    .fail( function() {
        console.log('Error with /v1/logoff');
        deleteCookie("air");  // no matter what, delete the cookie after this call completes
    });
    handleBlankScreen(false);
};


//---------------------------------------------------------------------------------
// resetPW - reset the user's password and send it to the user's inbox
//
// @params  <none>
// @returns <none>
//---------------------------------------------------------------------------------
window.resetPW = function () {
    var f = w2ui.passwordform;
    var username = f.record.user;
    var params = {username: username };
    var dat = JSON.stringify(params);
    var message = "";
    $.post('/v1/resetpw/', dat, null, "json")
    .done(function(data) {
        if (data.status === "error") {
            $(f.box).find("#LoginMessage").find(".errors").empty();
            message = "Error changing password";
            $(f.box).find("#LoginMessage").find(".errors").append("<p>" + message + "</p>");
            $(f.box).find("#LoginMessage").removeClass("hidden");
            // w2ui.passwordform.error(w2utils.lang(data.message));
            return;
        }
        else if (data.status === "success") {
            $(f.box).find("#LoginMessage").find(".errors").empty();
            message = "An updated password has been emailed to you.";
            $(f.box).find("#LoginMessage").find(".errors").append("<p>" + message + "</p>");
            $(f.box).find("#LoginMessage").removeClass("hidden");
        }
        return;
    })
    .fail(function(/*data*/){
        w2ui.passwordform.error("Reset password failed");
        return;
    });
};