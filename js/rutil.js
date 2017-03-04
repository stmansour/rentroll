/*global
    w2ui, app, console
*/

//-----------------------------------------------------------------------------
// getCurrentBusiness - return the Business Unit currently slected in the
//                      main toolbar
// @params
// @return  the BUD of the currently selected business
//-----------------------------------------------------------------------------
function getCurrentBusiness() {
    "use strict";
    var x = document.getElementsByName("BusinessSelect");
    return x[0];
}

//-----------------------------------------------------------------------------
// setToForm -  enable form sform in toplayout.  Also, set the forms url and
//              request data from the server
// @params
//   sform = name of the form
//   url   = request URL for the form
//-----------------------------------------------------------------------------
function setToForm(sform, url) {
    "use strict";
    console.log('sform = ' + sform + '  url = ' + url);
    var f = w2ui[sform];
    w2ui.toplayout.show('right', true);
    w2ui.toplayout.content('right', f);
    w2ui.toplayout.sizeTo('right', 700);
    if (url.length > 0) {
        f.url = url;
        f.request();
    }
}

//-----------------------------------------------------------------------------
// setToRAForm -  enable the Rental Agreement form in toplayout.  Also, set
//                the forms url and request data from the server
// @params
//   bid = business id (or the BUD)
//  raid = Rental Agreement ID
//     d = date to use for time sensitive data
//-----------------------------------------------------------------------------
function setToRAForm(bid, raid, d) {
    "use strict";
    w2ui.toplayout.content('right', w2ui.raLayout);
    w2ui.toplayout.show('right', true);
    w2ui.toplayout.sizeTo('right', 900);
    w2ui.rentalagrForm.url = '/v1/rentalagr/' + bid + '/' + raid;
    w2ui.rentalagrForm.request();

    //----------------------------------------------------------------
    // Get the associated Rentables...
    //      /v1/rar/bid/raid[?d1=2017-02-1]
    //      if no date is specified, today's date is used as the default.
    //----------------------------------------------------------------
    w2ui.rarGrid.url = '/v1/rar/' + bid + '/' + raid;
    console.log('rar url = ' + w2ui.rarGrid.url);
    w2ui.rarGrid.request();
    w2ui.rarGrid.header = plural(app.sRentable) + ' as of ' + dateFmtStr(d);
    w2ui.rarGrid.show.toolbarSearch = false;

    //----------------------------------------------------------------
    // Get the associated Payors...
    //      /v1/rapeople/bid/raid[?type=payor&d1=2017-02-1]
    //      if no date is specified, today's date is used as the default.
    //      if no person type is provided, payor is assumed
    //----------------------------------------------------------------
    w2ui.rapGrid.url = '/v1/rapayor/' + bid + '/' + raid;
    console.log('rapGrid url = ' + w2ui.rapGrid.url);
    w2ui.rapGrid.request();
    w2ui.rapGrid.header = plural(app.sPayor) + ' as of ' + dateFmtStr(d);

    //----------------------------------------------------------------
    // Get the associated Users...
    //      /v1/rapeople/bid/raid[?type=user&d1=2017-02-1]
    //      if no date is specified, today's date is used as the default.
    //----------------------------------------------------------------
    w2ui.rauGrid.url = '/v1/rapeople/' + bid + '/' + raid + '?type=user';
    console.log('rauGrid url = ' + w2ui.rauGrid.url);
    w2ui.rauGrid.request();
    w2ui.rauGrid.header = plural(app.sUser) + ' as of ' + dateFmtStr(d);

    //----------------------------------------------------------------
    // Get the associated Pets...
    //      /v1/xrapets/bid/raid
    //----------------------------------------------------------------
    w2ui.raPetGrid.url = '/v1/rapets/' + bid + '/' + raid;
    console.log('xrapets url = ' + w2ui.rarGrid.url);
    w2ui.raPetGrid.request();
    w2ui.raPetGrid.header = 'Pets as of ' + dateFmtStr(d);

}

//-----------------------------------------------------------------------------
// plural - return the plural of the provided word.  Totally simplistic at
//          this point, it just adds an 's'.  It will need serious updates
//          going forward
// @params
//   s = the word to pluralize
// @return - the plural of word s
//-----------------------------------------------------------------------------
function plural(s) {
    "use strict";
    return s + 's';
}


//-----------------------------------------------------------------------------
// dateFromDC - return a java date value equal to the date in the supplied
//              date control 
// @params
//   dc = date control
// @return - java date value
//-----------------------------------------------------------------------------
function dateFromDC(dc) {
    "use strict";
    var x = new Date(dc.value);
    return new Date(x.getTime() + 24 * 60 * 60 * 1000); // for some reason we need to add 1 day to get the right value
}

//-----------------------------------------------------------------------------
// dateTodayStr - return a string with today's date in the form d/m/yyyy
// @params
//   <none>
// @return - formatted date string
//-----------------------------------------------------------------------------
function dateTodayStr() {
    "use strict";
    var today = new Date();
    return dateFmtStr(today);
}

//-----------------------------------------------------------------------------
// dateFmtStr - return a string with the supplied date in the form d/m/yyyy
// @params - date
//
// @return - formatted date string
//-----------------------------------------------------------------------------
function dateFmtStr(today) {
    "use strict";
    var dd = today.getDate();
    var mm = today.getMonth() + 1; //January is 0!
    var yyyy = today.getFullYear();
    return mm + '/' + dd + '/' + yyyy;
}

//-----------------------------------------------------------------------------
// dayBack - supply the date control and this function will go to the previous
//           day. 
// @params
//   dc = date control
// @return string value that was set in dc
//-----------------------------------------------------------------------------
function dayBack(dc) {
    "use strict";
    var x = dateFromDC(dc);
    var y = new Date(x.getTime() - 24 * 60 * 60 * 1000); // one day prior 
    return setDateControl(dc, y);
}

//-----------------------------------------------------------------------------
// dayFwd - supply the date control and this function will go to the next day. 
// @params
//   dc = date control
// @return string value that was set in dc
//-----------------------------------------------------------------------------
function dayFwd(dc) {
    "use strict";
    var x = dateFromDC(dc);
    var y = new Date(x.getTime() + 24 * 60 * 60 * 1000); // one day prior
    return setDateControl(dc, y);
}

//-----------------------------------------------------------------------------
// dateMonthFwd - return a date that is one month from the supplied date. It
//                will snap the date to the end of the month if the 
//                current date is the end of the month.
// @params
//   y = starting date
// @return - a date that is one month from y
//-----------------------------------------------------------------------------
function dateMonthFwd(y) {
    "use strict";
    var m = (y.getMonth() + 1) % 12; // set m to the correct next month value
    var my = (y.getMonth() + 1) / 12; // number of years to add for next month
    var d = y.getDate(); // this is the target date
    // console.log('dateMonthFwd: T1 -    d = ' + d);

    // If there is a chance that there is no such date next month, then let's make sure we 
    // do this right. If the date is > than the number of days in month m then snap as follows:
    // if d is valid in month m then use d, otherwise snap to the end of the month.
    if (d > 28) {
        var d0 = new Date(y.getFullYear() + my, m, 0, 0, 0, 0);
        var daysInCurrentMonth = d0.getDate();
        var m2 = (y.getMonth() + 2) % 12; // used to find # days in month m
        var m2y = (y.getMonth() + 2) / 12; // number of years to add for month m
        var d3 = new Date(y.getFullYear() + m2y, m2, 0, 0, 0, 0);
        var daysInNextMonth = d3.getDate();
        if (d >= daysInNextMonth || d == daysInCurrentMonth) { d = daysInNextMonth; }
    }
    // console.log('dateMonthFwd:  m = ' + m + '   d = ' + d);
    var d2 = new Date(y.getFullYear() + my, m, d, 0, 0, 0);
    return d2;
}

//-----------------------------------------------------------------------------
// monthFwd - supply the date control and this function will go to the next 
//            month. It will snap the date to the end of the month if the 
//            current date is the end of the month.
// @params
//   dc = date control
// @return string value that was set in dc
//-----------------------------------------------------------------------------
function monthFwd(dc) {
    "use strict";
    var y = dateFromDC(dc);
    var d2 = dateMonthFwd(y);
    return setDateControl(dc, d2);
}

//-----------------------------------------------------------------------------
// monthBack - supply the date control, this function will go to the previous 
//             month. It will snap the date to the end of the month if the 
//             current date is the end of the month.
// @params
//   dc = date control
// @return string value that was set in dc
//-----------------------------------------------------------------------------
function monthBack(dc) {
    "use strict";
    var y = dateFromDC(dc);
    var yb = 0; // assume same year
    var m = y.getMonth() - 1;
    if (m < 0) {
        m = 11;
        yb = 1; // we've gone back one year
    }
    var d = y.getDate();
    if (d >= 28) {
        var d0 = new Date(y.getFullYear(), ((y.getMonth() + 1) % 12), 0, 0, 0, 0); // date of last day in prev month
        var daysInCurrentMonth = d0.getDate();
        var d3 = new Date(y.getFullYear() - yb, y.getMonth(), 0, 0, 0, 0); // date() is number of days in month y.getMonth()
        var daysInPrevMonth = d3.getDate();
        if (d == daysInCurrentMonth || d >= daysInPrevMonth) { d = daysInPrevMonth; }
    }
    var d2 = new Date(y.getFullYear() - yb, m, d, 0, 0, 0);
    return setDateControl(dc, d2);
}

//-----------------------------------------------------------------------------
// dateControlString
//           - return a date string based on the supplied date that can be
//             used as the .value attribute of a date control.  That is, in
//             the format  yyyy-mm-dd.
// @params
//   dt = java date value
// @return string value yyyy-mm-dd
//-----------------------------------------------------------------------------
function dateControlString(dt) {
    "use strict";
    var m = dt.getMonth() + 1;
    var d = dt.getDate();
    var s = '' + dt.getFullYear() + '-';
    if (m < 10) { s += '0'; }
    s += '' + m + '-';
    if (d < 10) { s += '0'; }
    s += d;
    return s;
}

//-----------------------------------------------------------------------------
// w2uiDateControlString
//           - return a date string formatted the way the w2ui dates are
//             expected, based on the supplied date that can be
//             used as the .value attribute of a date control.  That is, in
//             the format  m/d/yyyy.
// @params
//   dt = java date value
// @return string value yyyy-mm-dd
//-----------------------------------------------------------------------------
function w2uiDateControlString(dt) {
    "use strict";
    var m = dt.getMonth() + 1;
    var d = dt.getDate();
    var s = '' + m + '/' + d+'/' + dt.getFullYear();
    return s;
}

//-----------------------------------------------------------------------------
// setDateControl
//           - supply the date control and the date. This function will format
//             the date as needed by the date control. Then it will set the
//             date control with that date. It also returns the date string
//             used to set the control.
// @params
//   dc = date control
//   dt = java date value to set in dc
// @return string value that was set in dc
//-----------------------------------------------------------------------------
function setDateControl(dc, dt) {
    "use strict";
    var s = dateControlString(dt);
    dc.value = s;
    return s;
}
