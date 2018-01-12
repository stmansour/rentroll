"use strict";

/*global
    $, console, app, w2ui, w2popup, setInterval, getCookieValue, 
*/

var loginRoURL = "/webclient/html/formlogin.html";
var loginRcURL = "/webclient/html/formrcptlogin.html";

var loginSessionChecker = {};

var loginPopupOptions = {
    body: '<div id="loginPopupForm" style="width: 100%; height: 100%;"></div>',
    style: 'padding: 15px 0px 0px 0px; overflow: auto;',
    width: 400,
    height: 435,
    showMax: true,
    modal: true,
    onOpen: function (event) {
        event.onComplete = function () {
            $('#w2ui-popup #loginPopupForm').w2render('passwordform');
        };
    },
};

function userProfileToUI() {
    var f = w2ui.passwordform;

    if (f) {
        var name = app.name;
        if (name.length === 0 || app.uid === 0) { name = "?";}
        $("#user_menu_container").find("#username").text(name);
        var imgurl = app.imageurl;
        if (imgurl.length === 0) { imgurl = app.userBlankImage; }
        $("#user_menu_container").find("img").attr("src", imgurl);
    }
}

function buildLoginForm() {
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
                        app.name = data.name;
                        app.imageurl = data.imageurl;
                        $(f.box).find("#LoginMessage").addClass("hidden");
                        w2popup.close();
                        w2ui.passwordform.record.pass = ""; // after closing dialog, remove password information.
                        userProfileToUI();
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
}

//---------------------------------------------------------------------------------
// startNewSession - encapsulates the steps needed to launch a new login session
//                 and start up a session checker to have the user log in again
//                 if the session expires
//
// @params  <none>
// @returns <none>
//---------------------------------------------------------------------------------
function startNewSession() {
    ensureSession(); // get the user logged in
    startSessionChecker(); // have the user log in if the session expires

}

//---------------------------------------------------------------------------------
// launchSession - if a valid sessionid exists, use it and get user profile info
//                 if not, log in.
//
// @params  <none>
// @returns <none>
//---------------------------------------------------------------------------------
function launchSession() {
    var x = getCookieValue("airoller");
    if (x !== null && x.length > 0) {
        $.get('/v1/userprofile/')
        .done(function(sdata, textStatus, jqXHR) {
            var data = JSON.parse(sdata);
            if (data.status == "success") {
                app.uid = data.uid;
                app.username = data.username;
                app.name = data.name;
                app.imageurl = data.imageurl;
                startSessionChecker(); // have the user log in if the session expires
                userProfileToUI();
            } else {
                startNewSession();
            }
        })
        .fail( function() {
            console.log('Error getting /v1/userprofile');
            startNewSession();
        });
    }
    startNewSession();
}


//---------------------------------------------------------------------------------
// startSessionChecker - validate the session every 5 seconds
//
// @params  <none>
// @returns <none>
//---------------------------------------------------------------------------------
function startSessionChecker() {
    loginSessionChecker = setInterval(
    function() {
        ensureSession();
    }, 5000); // watch out for session expiring
}

//---------------------------------------------------------------------------------
// ensureSession - check to see if we have our session cookie.  If not, we need to
//             authenticate.
//
// @params  <none>
// @returns <none>
//---------------------------------------------------------------------------------
function ensureSession() {
    // console.log('ensureSession');
    if (w2popup.status == "open") {return;}
    var name = "airoller=";
    var decodedCookie = decodeURIComponent(document.cookie);
    var ca = decodedCookie.split(';');
    for(var i = 0; i <ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) === 0) {
            // return c.substring(name.length, c.length);
            return; // the cookie is here, so it has not expired
        }
    }
    // The cookie was not found. We need to authenticate...
    $().w2popup('open', loginPopupOptions);
    if (app.uid !== 0) {
        var f = w2ui.passwordform;
        if (f) {
            $(f.box).find("#LoginMessage").find(".errors").empty();
            var message = "Your session hass expired. Please login again.";
            $(f.box).find("#LoginMessage").find(".errors").append("<p>" + message + "</p>");
            $(f.box).find("#LoginMessage").removeClass("hidden");
        }
    }
}

//---------------------------------------------------------------------------------
// logoff - sign out of the current session
//
// @params  <none>
// @returns <none>
//---------------------------------------------------------------------------------
function logoff() {
    var name = "airoller=";
    var decodedCookie = decodeURIComponent(document.cookie);
    var ca = decodedCookie.split(';');
    var found = false;
    for(var i = 0; i <ca.length && !found; i++) {
        var c = ca[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) === 0) {
            found = true;
            // return c.substring(name.length, c.length);

        }
    }
    app.uid = 0;
    app.name = "";
    app.username = "";
    app.imageurl = "";
    if (found) {
        $.get('/v1/logoff/')
        .done(function(data, textStatus, jqXHR) {
            if (jqXHR.status == 200) {
                console.log('logoff success, app.uid set to 0.');
                ensureSession();
            } else {
                console.log( '**** YIPES! ****  status on logoff = ' + textStatus);
            }
        })
        .fail( function() {
            console.log('Error getting /v1/uilists');
        });
    }
}


//---------------------------------------------------------------------------------
// resetPW - reset the user's password and send it to the user's inbox
//
// @params  <none>
// @returns <none>
//---------------------------------------------------------------------------------
function resetPW() {
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
}