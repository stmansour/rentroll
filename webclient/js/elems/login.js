"use strict";

/*global
    $, console, app, w2ui, w2popup, setInterval,
*/
var loginTmplURL = "/webclient/html/formlogin.html";

var loginSessionChecker = {};

var loginPopupOptions = {
    body: '<div id="loginPopupForm" style="width: 100%; height: 100%;"></div>',
    style: 'padding: 15px 0px 0px 0px; overflow: auto;',
    width: 425,
    height: 210,
    showMax: true,
    modal: true,
    onOpen: function (event) {
        event.onComplete = function () {
            $('#w2ui-popup #loginPopupForm').w2render('passwordform');
        };
    },
};


function buildLoginForm() {
    $().w2form({
        name: 'passwordform',
        formURL: loginTmplURL,
        style: 'border: 0px; background-color: transparent;',
        fields: [{field: 'user', type: 'text',     required: true, html: {caption: 'User Name' /*, attr: 'readonly'*/} },
                 {field: 'pass', type: 'password', required: true, html: {caption: 'Password'} },
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
                        return;
                    }
                    else if (data.status === "success") {
                        app.uid = data.uid;
                        $(f.box).find("#LoginMessage").addClass("hidden");
                        w2popup.close();
                        w2ui.passwordform.record.pass = ""; // after closing dialog, remove password information.

                        // render the user details in web page
                        $("#user_menu_container").find("#username").text("UID: "+app.uid);
                        $("#user_menu_container").find("img").attr("src", app.userActiveImage);
                    }
                    console.log("Login service returned unexpected status: " + data.status);
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
// startSessionChecker - validate the session every 5 seconds
//
// @params  <none>
// @retunrs <none>
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
// @retunrs <none>
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
// @retunrs <none>
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
    if (found) {
        $.get('/v1/logoff/')
        .done(function(data, textStatus, jqXHR) {
            if (jqXHR.status == 200) {
                app.uid = 0;
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


