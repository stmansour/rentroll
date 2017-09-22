"use strict";
//-----------------------------------------------------------------------------
// dayBack - supply the date control and this function will go to the previous
//           day.
// @params
//   dc = date control
// @return string value that was set in dc
//-----------------------------------------------------------------------------
function dayBack(dc) {

    var x = dateFromString(dc.value);
    // set date to previous day
    x.setDate(x.getDate() - 1);
    return setDateControl(dc, x);
}

//-----------------------------------------------------------------------------
// dayFwd - supply the date control and this function will go to the next day.
// @params
//   dc = date control
// @return string value that was set in dc
//-----------------------------------------------------------------------------
function dayFwd(dc) {

    var x = dateFromString(dc.value);
    // set date to next day
    x.setDate(x.getDate() + 1);
    return setDateControl(dc, x);
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

    var y = dateFromString(dc.value);
    var d2 = dateMonthFwd(y);
    return setDateControl(dc, d2);
}

//-----------------------------------------------------------------------------
// setToCurrentMonth
//            This routine sets the supplied date control to the 1st of
//            the current month.
// @params
//   dc = date control
// @return string value that was set in dc
//-----------------------------------------------------------------------------
function setToCurrentMonth(dc) {

    var y = new Date();
    var d2 = new Date(y.getFullYear(), y.getMonth(), 1, 0, 0, 0, 0);
    return setDateControl(dc, d2);
}

//-----------------------------------------------------------------------------
// setToNextMonth
//            This routine sets the supplied date control to the 1st of
//            the next month..
// @params
//   dc = date control
// @return string value that was set in dc
//-----------------------------------------------------------------------------
function setToNextMonth(dc) {

    var y = new Date();
    var my = (y.getMonth() + 1) / 12; // number of years to add for next month
    var m = (y.getMonth() + 1) % 12;  // next month
    var d2 = new Date(y.getFullYear() + my, m, 1, 0,0,0,0);
    return setDateControl(dc, d2);
}

//-----------------------------------------------------------------------------
// dateMonthBack - return a date which is a month prior to the supplied date
// @params
//   y = input date
// @return date which is y - 1 month
//-----------------------------------------------------------------------------
function dateMonthBack(y) {

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
    return new Date(y.getFullYear() - yb, m, d, 0, 0, 0);
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

    var y = dateFromString(dc.value);
    var d2 =  dateMonthBack(y);
    return setDateControl(dc, d2);
}

//-----------------------------------------------------------------------------
// dateControlString
//           - return a date string based on the supplied date that can be
//             used as the .value attribute of a date control.  That is, in
//             the format  m/d/yyyy.
// @params
//   dt = java date value
// @return string value m/d/yyyy
//-----------------------------------------------------------------------------
function dateControlString(dt) {

    var m = dt.getMonth() + 1;
    var d = dt.getDate();
    // if (m < 10) { s += '0'; }
    var s = '' + m + '/';
    // if (d < 10) { s += '0'; }
    s += d;
    s += '/' + dt.getFullYear() + '';
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
// @return string value mm-dd-yyyy
//-----------------------------------------------------------------------------
function w2uiDateControlString(dt) {

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

    var s = w2uiDateControlString(dt);
    dc.value = s;
    return s;
}

//-----------------------------------------------------------------------------
// dateFromString - return a java date value equal to the date in the supplied
//              date control
// @params
//   ds = date string value
// @return - java date value
//-----------------------------------------------------------------------------
function dateFromString(ds) {


    // Strange thing about javascript dates
    // new Date("2017-06-28") gives a date with offset value with local timezone i.e, Wed Jun 28 2017 05:30:00 GMT+0530 (IST)
    // new Date("2017/06/28") gives a date without offset value with local timezone i.e, Wed Jun 28 2017 00:00:00 GMT+0530 (IST)

    ds = ds.replace(/-/g,"\/").replace(/T.+/, ''); // first replace `/` with `-` and also remove `hh:mm:ss` value we don't need it
    return new Date(ds);
}

//-----------------------------------------------------------------------------
// dateTodayStr - return a string with today's date in the form d/m/yyyy
// @params
//   <none>
// @return - formatted date string
//-----------------------------------------------------------------------------
function dateTodayStr() {

    var today = new Date();
    return dateFmtStr(today);
}

//-----------------------------------------------------------------------------
// dateFmtStr - return a string with the supplied date in the form d/m/yyyy
// @params
//    date
// @return - formatted date string
//-----------------------------------------------------------------------------
function dateFmtStr(today) {

    var dd = today.getDate();
    var mm = today.getMonth() + 1; //January is 0!
    var yyyy = today.getFullYear();
    return mm + '/' + dd + '/' + yyyy;
}

//-----------------------------------------------------------------------------
// isDatePriorToCurrentDate - return boolean value
// @params
//    date object
// @return - boolean
//-----------------------------------------------------------------------------
function isDatePriorToCurrentDate(date) {
    var dd = date.getDate();
    var mm = date.getMonth() + 1; //January is 0!
    var yyyy = date.getFullYear();
    var currentDateTime = new Date();
    if (currentDateTime.getTime() >= date.getTime()) {
        if (currentDateTime.getDate() == dd && currentDateTime.getMonth() == mm && currentDateTime.getFullYear() == yyyy) {
            return false;
        } else {
            return true;
        }
    }
    return false;
}



$(function() {    
     $(document).on("blur change", "input[type=us-date1], input[type=us-date2]", function(e) {   
         // replace trailing zero from date using regex   
         this.value = this.value.replace(/\b0*(?=\d)/g, '');        
     });
 });