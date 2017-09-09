/*global
    app, w2ui, $, monthBack, monthFwd, dayBack, dayFwd, setToCurrentMonth, setToNextMonth,
    console, dateFromString, dateControlString,
*/
"use strict";
//-----------------------------------------------------------------------------
// handleDateToolbarAction
//          - based on the button selected, perform the appropriate date
//            modification, update the dates in the App structure, and update
//            the toolbar widgets.
// @params
//          event - the event that occurred on the button bar
//          prefix - the prefix of the name of the date controls.  For example,
//                  if the date control is named receiptsD1, then the prefix
//                  is 'receipts'.
// @return  <no return value>
//-----------------------------------------------------------------------------
function handleDateToolbarAction(event,prefix) {
    var xd1 = document.getElementsByName(prefix + 'D1')[0];
    var xd2 = document.getElementsByName(prefix + 'D2')[0];
    switch (event.target) {
        case 'monthback':
            app.D1 = monthBack(xd1);
            app.D2 = monthBack(xd2);
            break;
        case 'monthfwd':
            app.D1 = monthFwd(xd1);
            app.D2 = monthFwd(xd2);
            break;
        case 'today':
            app.D1 = setToCurrentMonth(xd1);
            app.D2 = setToNextMonth(xd2);
            break;
        case 'dayback':
            app.D1 = dayBack(xd1);
            app.D2 = dayBack(xd2);
            break;
        case 'dayfwd':
            app.D1 = dayFwd(xd1);
            app.D2 = dayFwd(xd2);
            break;
    }
}

//-----------------------------------------------------------------------------
// setDateControlsInToolbar
//           -  Utility routine to set the date in a toolbar date navigation
//              area to the date values in app.D1 and app.D2
// @params
//   prefix = the prefix of the name of the date controls.  For example,
//            if the date control is named receiptsD1, then the prefix is
//            'receipts'.
// @return  <no return value>
//-----------------------------------------------------------------------------
function setDateControlsInToolbar(prefix) {
    var xd1 = document.getElementsByName(prefix + 'D1')[0];
    var xd2 = document.getElementsByName(prefix + 'D2')[0];
    if (typeof xd1 != "undefined") { xd1.value = app.D1; }
    if (typeof xd2 != "undefined") { xd2.value = app.D2; }
}


//-----------------------------------------------------------------------------
// genDateRangeNavigator
//           -  Utility routine create an array of fields that form
//              a date range navigator.  The prefix is applied to the
//              <input type="date"> controls so that they can be
//              uniquely identified.
// @params
//   prefix = the prefix of the name of the date controls.  For example,
//            if the date control is named receiptsD1, then the prefix is
//            'receipts'.
// @return  an array of fields that can be passed into toolbar.add()
//-----------------------------------------------------------------------------
function genDateRangeNavigator(prefix) {
    var html1 = '<div class="w2ui-field" style="padding: 0px 5px;">From: <input name="' + prefix + 'D1"></div>';
    var html2 = '<div class="w2ui-field" style="padding: 0px 5px;">To: <input name="' + prefix + 'D2">' + '</div>';
    var tmp = [{ type: 'break', id: 'break1' },
        { type: 'button', id: 'monthback', icon: 'fa fa-backward', tooltip: 'month back' },
        { type: 'button', id: 'dayback', icon: 'fa fa-chevron-circle-left', tooltip: 'day back' },
        { type: 'html', id: 'D1', html: function() {return html1; },
        onRefresh: function(event) {
               if(event.target == 'D1'){
                   console.log('Event type: '+ event.type + ' TARGET: '+ event.target, event);

                   // w2field in toolbar must be initialized during refresh
                   // see: https://github.com/vitmalina/w2ui/issues/886
                   event.onComplete = function(/*ev*/){
                       $('input[name='+ prefix +'D1]').w2field('date', {format: 'm/d/yyyy'});
                   };
               }
            }
        },
        { type: 'button', id: 'today', icon: 'fa fa-circle-o', tooltip: 'present month' },
        { type: 'html', id: 'D2', html: function() {return html2; },
        onRefresh: function(event) {
               if(event.target == 'D2'){
                   console.log('Event type: '+ event.type + ' TARGET: '+ event.target, event);

                   // w2field in toolbar must be initialized during refresh
                   // see: https://github.com/vitmalina/w2ui/issues/886
                   event.onComplete = function(/*ev*/){
                       $('input[name='+ prefix +'D2]').w2field('date', {format: 'm/d/yyyy', start: $('input[name='+ prefix +'D1]')});
                   };
               }
            }
        },
        { type: 'button', id: 'dayfwd', icon: 'fa fa-chevron-circle-right', tooltip: 'day forward' },
        { type: 'button', id: 'monthfwd', icon: 'fa fa-forward', tooltip: 'month forward' },
    ];
    return tmp;
}

//-----------------------------------------------------------------------------
// updateGridPostDataDates
//          - if searchDtStart and searchDtStop have been defined in
//            grid.postData then update their values with the current
//            dates in the datanav controls
// @params
//   grid   the grid of interest
// @return  <no return value>
//-----------------------------------------------------------------------------
function updateGridPostDataDates(grid) {
    if (typeof grid.postData.searchDtStart === "string") {
        grid.postData.searchDtStart = app.D1;
        grid.postData.searchDtStop  = app.D2;
    }
}

//-----------------------------------------------------------------------------
// addDateNavToToolbar
//          - Utility routine create add a date navigator to a toolbar
// @params
//   prefix = the w2ui grid control prefix name that follows the naming
//            convention:  prefix + 'Grid'
// @return  <no return value>
//-----------------------------------------------------------------------------
function addDateNavToToolbar(prefix) {
    var grid = w2ui[prefix+'Grid'];
    grid.toolbar.add( genDateRangeNavigator(prefix) );
    grid.toolbar.on('click', function(event) {
        handleDateToolbarAction(event,prefix); // adjusts dates and loads into date controls
        updateGridPostDataDates(grid);
        grid.load(grid.url, function() {
            grid.refresh(); // need to refresh the grid for redraw purpose
        });
    });
    grid.toolbar.on('refresh', function (/*event*/) {
        setDateControlsInToolbar(prefix);
        updateGridPostDataDates(grid);
    });

    // bind onchange event for date input control for assessments
    var nd1 = prefix + "D1";
    var nd2 = prefix + "D2";
    $(document).on("keypress change", "input[name="+nd1+"]", function(e) {
        // if event type is keypress then
        if (e.type == 'keypress'){
            // do not procedd further untill user press the Enter key
            if (e.which != 13) {
                return;
            }
        }
        var xd1 = document.getElementsByName(nd1)[0].value;
        var xd2 = document.getElementsByName(nd2)[0].value;
        var d1 = dateFromString(xd1);
        var d2 = dateFromString(xd2);
        // check that it is valid or not
        if (isNaN(Date.parse(xd1))) {
            return;
        }
        // check that year is not behind 2000
        if (d1.getFullYear() < 2000) {
            return;
        }
        // check that from date does not have value greater then To date
        if (d1.getTime() >= d2.getTime()) {
            d1 = new Date(d2.getTime() - 24 * 60 * 60 * 1000); //one day back from To date
        }
        app.D1 = dateControlString(d1);
        app.D2 = dateControlString(d2);
        updateGridPostDataDates(grid);
        grid.load(grid.url, function() {
            grid.refresh();
            // w2ui internally loads date calender on focus event. look at the following link:
            // https://github.com/mpf82/w2ui/blob/f7f7159af5066b86a4e13bac1c1aa88e2e6b354f/src/w2fields.js#L933
            // we dont want to show calender overlay while on focus. so make it readonly temporary.
            // Ref:https://github.com/mpf82/w2ui/blob/f7f7159af5066b86a4e13bac1c1aa88e2e6b354f/src/w2fields.js#L1602
            $("input[name="+nd1+"]").prop('readonly', true).focus().prop('readonly', false);
        });
    }).on("keypress change", "input[name="+nd2+"]", function(e) {
        // if event type is keypress then
        if (e.type == 'keypress'){
            // do not procedd further untill user press the Enter key
            if (e.which != 13) {
                return;
            }
        }
        var xd1 = document.getElementsByName(nd1)[0].value;
        var xd2 = document.getElementsByName(nd2)[0].value;
        var d1 = dateFromString(xd1);
        var d2 = dateFromString(xd2);
        // check that it is valid or not
        if (isNaN(Date.parse(xd2))) {
            return;
        }
        // check that year is not behind 2000
        if (d2.getFullYear() < 2000) {
            return;
        }
        // check that from date does not have value greater then To date
        if (d2.getTime() <= d1.getTime()) {
            d2 = new Date(d1.getTime() + 24 * 60 * 60 * 1000); //one day forward from From date
        }
        app.D1 = dateControlString(d1);
        app.D2 = dateControlString(d2);
        updateGridPostDataDates(grid);
        grid.load(grid.url, function() {
            grid.refresh();
            // w2ui internally loads date calender on focus event. look at the following link:
            // https://github.com/mpf82/w2ui/blob/f7f7159af5066b86a4e13bac1c1aa88e2e6b354f/src/w2fields.js#L933
            // we dont want to show calender overlay while on focus. so make it readonly temporary.
            // Ref:https://github.com/mpf82/w2ui/blob/f7f7159af5066b86a4e13bac1c1aa88e2e6b354f/src/w2fields.js#L1602
            $("input[name="+nd2+"]").prop('readonly', true).focus().prop('readonly', false);
        });
    });
}
