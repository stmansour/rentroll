"use strict";
//-----------------------------------------------------------------------------
// formRecDiffer -  tells that form record has been changed
// **[copied from w2ui form's getChanges internal function]**
// @params
//   record = form's current record
//   original = form's initial record
//   result = returned object
// @return
//      Object with difference from `record` to `original`
//-----------------------------------------------------------------------------
window.formRecDiffer = function(record, original, result) {

    for (var i in record) {
        if (typeof record[i] == "object") {
            result[i] = formRecDiffer(record[i], original[i] || {}, {});
            if (!result[i] || $.isEmptyObject(result[i])) delete result[i];
        } else if ( record[i] !== null && record[i] != original[i] ) {
            /*** ================================================================
            || BY DEFAULT, W2UI SETS VALUE OF FIELD TO NULL IF NOTHING IS IN THERE
            || NOTE: be careful, for form record, <null> and <""> (blank string) both are same
            || it should not alert user that content has been changed !!!
            || so, for this, <undefined>, <NaN>, <null>, <""> all are same
            || NEED TO DO SOMETHING ABOUT THIS
            || HECK: it only makes sense when record[i] is not NULL (undefined, null, "", NaN)
            ================================================================ ***/
            result[i] = record[i];
        }
    }
    return result;
};

//-----------------------------------------------------------------------------
// getPersonDetailsByTCID -  returns the person details for given TCID
// @params
//   TCID = Transactant ID
// @return
//      Object with Transactant record
//-----------------------------------------------------------------------------
window.getPersonDetailsByTCID = function (BID, TCID) {


    // we need to use this structure to get person details from given TCID
    var params = {"cmd":"get","recid":0,"name":"transactantForm"},
        dat = JSON.stringify(params);

    return $.post("/v1/person/"+BID+"/"+TCID, dat, null, "json");
};

// form dirty alert confirmation dialog box options
var form_dirty_alert_options = {
    msg          : '<p>There are unsaved changes.</p><p>click Ignore Change to continue without saving your changes or click Continue Editing.</p>',
    title        : '',
    width        : 480,     // width of the dialog
    height       : 180,     // height of the dialog
    btn_yes      : {
        text     : 'Ignore Changes',   // text for yes button (or yes_text)
        class    : 'w2ui-btn w2ui-btn-red',      // class for yes button (or yes_class)
        style    : '',      // style for yes button (or yes_style)
        callBack : null     // callBack for yes button (or yes_callBack)
    },
    btn_no       : {
        text     : 'Continue Editing',    // text for no button (or no_text)
        class    : 'w2ui-btn',      // class for no button (or no_class)
        style    : '',      // style for no button (or no_style)
        callBack : null     // callBack for no button (or no_callBack)
    },
    callBack     : function(answer) {
        console.log("common callBack (Yes/No): ", answer);
    }     // common callBack
};

//-----------------------------------------------------------------------------
// form_dirty_alert - alert the user if form content has been changed and he leaves the form at five times as follows
// 1. When user change the business
// 2. When he clicks on the sidebar that load something else
// 3. When closing the form
// 4. When click on other record
// 5. When user closing the whole window
// NOTE: if form is dirty then only alert the user, otherwise always return true;
// @params
//   yes callback = what to do if user agree (Yes)
//   no callback = what to do if user disagree (No)
//   yes_args = yes callback arguments
//   no_args = no callback arguments
// @return: true or false
//-----------------------------------------------------------------------------
window.form_dirty_alert = function (yes_callBack, no_callBack, yes_args, no_args) {
    if (app.form_is_dirty) {
        w2confirm(form_dirty_alert_options)
        .yes(function() {
            if (typeof yes_callBack === "function") {
                if (Array.isArray(yes_args) && yes_args.length > 0) {
                    yes_callBack.apply(null, yes_args);
                } else{
                    yes_callBack();
                }
            }
        })
        .no(function() {
            if (typeof no_callBack === "function") {
                if (Array.isArray(no_args) && no_args.length > 0) {
                    no_callBack.apply(null, no_args);
                } else{
                    no_callBack();
                }
            }
        });
    } else {
        // if form is not dirty then simply execute yes callback which is default action
        if (typeof yes_callBack === "function") {
            // Reference: http://odetocode.com/blogs/scott/archive/2007/07/04/function-apply-and-function-call-in-javascript.aspx
            if (Array.isArray(yes_args) && yes_args.length > 0) {
                yes_callBack.apply(null, yes_args);
            } else{
                yes_callBack();
            }
        }
    }
};

// =================================================
// WINDOW BEFORE UNLOAD EVENT
// =================================================
// warn user if active form content has been changed
// for security reason you can't just popup your custom dialog
// see the thread: https://stackoverflow.com/questions/30712377/jquery-beforeunload-custom-pop-up-window-for-leaving-a-page
window.form_dirty_alert_window_unload = function (e) {
    if (app.form_is_dirty){
        if(!e) e = window.event;

        //e.cancelBubble is supported by IE - this will kill the bubbling process.
        e.cancelBubble = true;

        //e.stopPropagation works in Firefox.
        if (e.stopPropagation) {
            e.stopPropagation();
            e.preventDefault();
        }
        return "Changes in the form that you made may not be saved.";
    }
};

window.onbeforeunload=form_dirty_alert_window_unload;
// =================================================
