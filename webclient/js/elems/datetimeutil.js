"use strict";
'esversion 6';
/*global
  app, setDateControl, dateMonthBack, getDateFromDT, getTimeFromDT, dateFromString,
  dateFmtStr, zeroPad, yearFwd, dateYearFwd, yearBack, dateYearBack, applyLocaltimeDateOffset,
  UTCDateStringToW2UIValidDate,stringToDate,
*/

//-----------------------------------------------------------------------------
//
//                   TIMEZONE CONVERSION  UTILITIES
//
//-----------------------------------------------------------------------------
window.UTCstringToLocaltimeString = function(u) {
    var x = new Date(u);
    return x.toString();  // this will be a localtime value
};

// stringToDate takes a string of the form '2019-09-04' and converts it to
// a Date value in local time where the year is 2019, the month is 9 (Sep)
// and the day is 4.  The time will be 00:00:00
//
//   0123456789
//  '2019-09-04'
//
// INPUTS
//    s = a string in the format 'yyyy?mm?dd'.
//        y = digits of year
//        m = digits of month 01,02,03 ... 12
//        d = digits of day 01, 02, ... 31
//        ? = any character... typically it will be '-' or '/'
//
// RETURNS
//    localtime value with the supplied year, month, and day
//------------------------------------------------------------------------------
window.stringToDate = function(s) {
    var y = parseInt(s,10);
    var m = parseInt(s.substr(5,2),10);
    var d = parseInt(s.substr(8,2),10);
    var dt = new Date(y,m,d);
    return dt;
};

// applyLocaltimeDateOffset returns a date with the hours properly adjusted so
// that the UTC dates are at 00:00:00.  This can be useful when working with
// dates so that you don't lose or gain days due to editing across timezones.
// Example:
//    4/30/2019 00:00:00 UTC is equivalent to  4/29/2019 17:00:00 GMT-0700
//    If your browser is in localtime GMT-0700 and you load the date into
//    a UI date control that ONLY MAINTAINS THE DATE (not the time), then
//    what hours:mins:secs should you use when you get an updated date and
//    send it back to the server -- which requires the input to be in UTC.
//    Suppose you just set that date into your UI date control and didn't
//    change it.  When you retrieve it from the UI, you only get 4/29/2019
//    No hours, mins, or secs.  That is: 4/29/2019 00:00:00. Translate that
//    to UTC and you get 4/29/2019 07:00:00 which is 10 hours different from
//    its initial value. Repeat this operation several times and the date
//    will start to decrease to 28, then 27, etc.
// To address this, we adopt the standard to always make the UTC hrs, mins,
// and secs equal 00:00:00 for DATE-ONLY datetime values. This ensures that
// changes across timezones will keep the dates from drifting.
//
// INPUTS
//    d = a datevalue with the time assumed to be 00:00:00
//
// RETURNS
//    the adjusted date such that when it is converted to UTC, its hrs:mins:secs
//    will be 00:00:00
//------------------------------------------------------------------------------
window.applyLocaltimeDateOffset = function(d) {
    var m = d.getTimezoneOffset();
    if (m > 0) {
        return new Date(d.getFullYear(),d.getMonth(),d.getDate(),0,(1440-m));
    }
    return new Date(d.getFullYear(),d.getMonth(),d.getDate(),0,m);
};

//-----------------------------------------------------------------------------
// UTCDateStringToW2UIValidDate
//               - Javascript's toUTCString returns a date string like this:
//
//                 "Fri, 31 Dec 9999 00:00:00 GMT"
//
//                 This fails w2ui's validation check on dates.  This function
//                 changes the string above into m/d/y format -- which is in
//                 utc format and which will pass the validation check.  For
//                 the date above, the return value is "12/31/9999"
//
//                 Only the date is returned, not the time.
//
// @params
//   x = date in utc format
//
// @return  string representing the UTC date
//-----------------------------------------------------------------------------
window.UTCDateStringToW2UIValidDate = function(x) {
    var z = {Jan: "01", Feb: "02", Mar: "03", Apr: "04", May: "05", Jun: "06",
             Jul: "07", Aug: "08", Sep: "09", Oct: "10", Nov: "11", Dec: "12"};
    var d = x.substring(5,7);
    var mn = x.substring(8,11);
    var m = z[mn];
    var y = x.substring(12,16);
    return m + '/' + d + '/' + y;
};

//-----------------------------------------------------------------------------
// adjustARDates - a convenience function to apply the offset to 2 localtime
//                 dates just prior to sending them to the server.  This puts
//                 the dates into the appropriate value for saving in UTC.
//
// @params
//   d1 = first date
//   d2 = second date
//
// @return number of days between the dates
//-----------------------------------------------------------------------------
window.adjustARDates = function(r) {
    var d1 = applyLocaltimeDateOffset(new Date(r.DtStart));
    var d2 = applyLocaltimeDateOffset(new Date(r.DtStop));
    r.DtStart = UTCDateStringToW2UIValidDate(d1.toUTCString());
    r.DtStop = UTCDateStringToW2UIValidDate(d2.toUTCString());
};

//-----------------------------------------------------------------------------
// daysBetweenDates - returns the number of days between 2 dates.
//
// @params
//   d1 = first date
//   d2 = second date
//
// @return number of days between the dates
//-----------------------------------------------------------------------------
window.daysBetweenDates = function(d1,d2) {
    var oneDay = 24*60*60*1000; // hours*minutes*seconds*milliseconds
    var firstDate = new Date(d1.getFullYear(),d1.getMonth(),d1.getDate());
    var secondDate = new Date(d2.getFullYear(),d2.getMonth(),d2.getDate());

    return Math.round(Math.abs((firstDate.getTime() - secondDate.getTime())/(oneDay)));
};
//-----------------------------------------------------------------------------
// newDateKeepOldTime - create a new date keeping the hour and minute equal
//           to the supplied "orig" values, and using the supplied year, month,
//           and day.
// @params
//   orig = date string - probably from a date control
//   y    = year
//   m    = month
//   d    = day
//
// @return date object
//-----------------------------------------------------------------------------
window.newDateKeepOldTime = function(orig,y,m,d) {
    var dt = new Date(orig);
    var ddt = new Date(y,m,d,dt.getHours(),dt.getMinutes());
    return ddt;
};


//-----------------------------------------------------------------------------
//
//                   DATE FORMATING  UTILITIES
//
//-----------------------------------------------------------------------------

//-----------------------------------------------------------------------------
// zeroPad - if the string value of the number is < size,
//           left pad it with '0' to make it the requested size.
//           LIMITATION - at most, it will pad 10 '0's.  This is
//           extreme overkill for working with date/time strings.
// @params
//   n    = the number of interest
//   size = number of characters the output number should be
// @return string version of the number left-padded with '0's
//         to achieve a length of size.
//-----------------------------------------------------------------------------
window.zeroPad = function (n, size) {
    var s = "0000000000" + n;
    return s.substr(s.length-size);
};

//-----------------------------------------------------------------------------
// dtFormatISOToW2ui - return a w2ui datetime string from the provided
//          ISO 8601 formatted date string -- the format of JSONDateTIme
//          strings.  The returned string is in the format:
//          m/dd/yyyy H:MM {am|pm} . It is suitable for use in a w2ui
//          form control of type 'datetime'. For example:
//
//          input: "2018-04-04T23:38:00Z"
//          output: '4/04/2018 11:38 pm'.
//
//          If the year is prior to year 2000, it returns a 0 length string.
// @params
//   s    = JSONDateTime string
// @return
//         localtime string
//-----------------------------------------------------------------------------
window.dtFormatISOToW2ui = function (ds) {
    if (typeof ds != "string") {return ds;}  // handle error case of bad data type
    if (ds.indexOf('T') < 0) {return ds;}    // handle case where it's not in ISO format

    var dt = new Date(Date.parse(ds));
    if (dt.getFullYear() < 2000) {return '';}
    var hr = dt.getHours();
    var am = true;
    if (hr >= 12) {
        am = false;
    }
    if (hr > 12) {
        hr -= 12;
    }
    var s = 1+dt.getMonth() + '/' + zeroPad(dt.getDate(),2) + '/' +
            dt.getFullYear() + '   ' + hr + ':' + zeroPad(dt.getMinutes(),2) +
            ' ' + (am ? 'am' : 'pm');
    return s;
};

//-----------------------------------------------------------------------------
// localtimeToUTC - return a UTC datetime string from the localtime string
//          created by a w2ui datetime control.  That is, change a string
//          like this PDT string, "1/20/2018 1:00 am", to a UTC string
//          like this: "Sat, 20 Jan 2018 09:00:00 GMT"
// @params
//   s    = localtime string
// @return  UTC string
//-----------------------------------------------------------------------------
window.localtimeToUTC = function (s) {
    if (typeof s === "string" && s.length > 0) {
        var dt = new Date(s);
        return dt.toUTCString();
    }
    return '';
};

//-----------------------------------------------------------------------------
// dateControlString
//           - return a date string based on the supplied date that can be
//             used as the .value attribute of a date control.  That is, in
//             the format  m/d/yyyy.
// @params
//   dt = java date value
// @return string value m/d/yyyy
//-----------------------------------------------------------------------------
window.dateControlString = function (dt) {
    var m = dt.getMonth() + 1;
    var d = dt.getDate();
    // if (m < 10) { s += '0'; }
    var s = '' + m + '/';
    // if (d < 10) { s += '0'; }
    s += d;
    s += '/' + dt.getFullYear();
    return s;
};

//-----------------------------------------------------------------------------
// w2uiDateControlString
//           - return a date string formatted the way the w2ui dates are
//             expected, based on the supplied date that can be
//             used as the .value attribute of a date control.  That is, in
//             the format  m/d/yyyy.
// @params
//   dt = java date value
// @return string value m/d/yyyy
//-----------------------------------------------------------------------------
window.w2uiDateControlString = function (dt) {
    var m = dt.getMonth() + 1;
    var d = dt.getDate();
    var s = '' + m + '/' + d+'/' + dt.getFullYear();
    return s;
};

//-----------------------------------------------------------------------------
// w2uiUTCDateControlString
//           - return a UTC date string formatted the way the w2ui dates are
//             expected, based on the supplied date that can be
//             used as the .value attribute of a date control.  That is, in
//             the format  m/d/yyyy.
// @params
//   dt = java date value
// @return string value mm-dd-yyyy
//-----------------------------------------------------------------------------
window.w2uiUTCDateControlString = function (dt) {
    var m = dt.getUTCMonth() + 1;
    var d = dt.getUTCDate();
    var s = '' + m + '/' + d+'/' + dt.getUTCFullYear();
    return s;
};

//-----------------------------------------------------------------------------
// w2uiDateTimeControlString
//           - return a datetime string formatted the way the w2ui datetimes
//             are expected, based on the supplied date that can be
//             used as the .value attribute of a date control.  That is, in
//             the format  m/d/yyyy HH:MM {am|pm}.
// @params
//   dt = java date value
// @return string value mm-dd-yyyy HH:MM {am|pm}
//-----------------------------------------------------------------------------
window.w2uiDateTimeControlString = function (dt) {
    var m = dt.getMonth() + 1;
    var d = dt.getDate();
    var H = dt.getHours();
    var M = dt.getMinutes();
    var s = m + '/' + d + '/' + dt.getFullYear() + ' ';
    var bPM = H >= 12;
    if (H > 12) { H = H-12;}
    s += H + ':';
    if (M < 10) {
        s += '0';
    }
    s += M + ' ' + ((bPM) ? 'p':'a') + 'm';
    return s;
};



//-----------------------------------------------------------------------------
// dayBack - supply the date control and this function will go to the previous
//           day.
// @params
//   dc = date control
//    n = number of days to go back
// @return string value that was set in dc
//-----------------------------------------------------------------------------
window.dayBack = function (dc,n) {
    var x = dateFromString(dc.value);
    x.setDate(x.getDate() - n);
    return setDateControl(dc, x);
};

//-----------------------------------------------------------------------------
// dayFwd - supply the date control and this function will go to the next day.
// @params
//   dc = date control
//    n = number of days forward
// @return string value that was set in dc
//-----------------------------------------------------------------------------
window.dayFwd = function (dc,n) {
    var x = dateFromString(dc.value);
    x.setDate(x.getDate() + n);
    return setDateControl(dc, x);
};

//-----------------------------------------------------------------------------
// dateMonthFwd - return a date that is one month from the supplied date. It
//                will snap the date to the end of the month if the
//                current date is the end of the month.
// @params
//   y = starting date
// @return - a date that is one month from y
//-----------------------------------------------------------------------------
window.dateMonthFwd = function (y) {

    var m = (y.getMonth() + 1) % 12; // set m to the correct next month value
    var my = (y.getMonth() + 1) / 12; // number of years to add for next month
    var d = y.getDate(); // this is the target date
    // console.log('dateMonthFwd: T1 -    d = ' + d);

    // If there is a chance that there is no such date next month, then let's make sure we
    // do this right. If the date is >= 28, then always snap it to the end of the month.
    if (d >= 28) {
        // var d0 = new Date(y.getFullYear() + my, m, 0, 0, 0, 0);
        // var daysInCurrentMonth = d0.getDate();
        var m2 = (y.getMonth() + 2) % 12; // used to find # days in month m
        var m2y = (y.getMonth() + 2) / 12; // number of years to add for month m
        var d3 = new Date(y.getFullYear() + m2y, m2, 0, 0, 0, 0);
        var daysInNextMonth = d3.getDate();
        //if (d >= daysInNextMonth || d == daysInCurrentMonth) { d = daysInNextMonth; }
        d = daysInNextMonth;
    }
    // console.log('dateMonthFwd:  m = ' + m + '   d = ' + d);
    var d2 = new Date(y.getFullYear() + my, m, d, 0, 0, 0);
    return d2;
};

//-----------------------------------------------------------------------------
// monthFwd - supply the date control and this function will go to the next
//            month. It will snap the date to the end of the month if the
//            current date is the end of the month.
// @params
//   dc     = date control
//   strval = if provided, it needs to the string value to use for the existing
//            date rather than the value of the supplied date control
// @return string value that was set in dc
//-----------------------------------------------------------------------------
window.monthFwd = function (dc,strval) {
    var y = dateFromString(dc.value);
    if (typeof strval == "string") {
        y = dateFromString(strval);
    }
    var d2 = dateMonthFwd(y);
    return setDateControl(dc, d2);
};

//-----------------------------------------------------------------------------
// dateYearFwd - return a date that is one year from the supplied date. It
//                will snap the date to the end of the month if the
//                current date is the end of the month.
// @params
//   y = starting date
// @return - a date that is one year from y
//-----------------------------------------------------------------------------
window.dateYearFwd = function (y) {
    var m = y.getMonth();
    var d = y.getDate();
    var d2;
    if (m == 1 /*feb*/ && d >= 28) {
        d2 = new Date(y.getFullYear() + 1, 2/*mar*/, 0, 0, 0, 0);
    } else {
        d2 = new Date(y.getFullYear() + 1, m, d, 0, 0, 0);
    }
    return d2;
};
//-----------------------------------------------------------------------------
// yearFwd - supply the date control and this function will go to the next
//            year. It will snap the date to the end of the month if the
//            current date is the end of the month.
// @params
//   dc     = date control
//   strval = if provided, it needs to the string value to use for the existing
//            date rather than the value of the supplied date control
// @return string value that was set in dc
//-----------------------------------------------------------------------------
window.yearFwd = function (dc,strval) {
    var y = dateFromString(dc.value);
    if (typeof strval == "string") {
        y = dateFromString(strval);
    }
    var d2 = dateYearFwd(y);
    return setDateControl(dc, d2);
};

//-----------------------------------------------------------------------------
// dateYearBack - return a date that is one year prior to the supplied date. It
//                will snap the date to the end of the month if the
//                current date is the end of the month.
// @params
//   y = starting date
// @return - a date that is one year from y
//-----------------------------------------------------------------------------
window.dateYearBack = function (y) {
    var m = y.getMonth();
    var d = y.getDate();
    var d2;
    if (m == 1 /*feb*/ && d >= 28) {
        d2 = new Date(y.getFullYear() - 1, 2/*mar*/, 0, 0, 0, 0);
    } else {
        d2 = new Date(y.getFullYear() - 1, m, d, 0, 0, 0);
    }
    return d2;
};
//-----------------------------------------------------------------------------
// yearBack - supply the date control and this function will go to the previous
//            year. It will snap the date to the end of the month if the
//            current date is the end of the month.
// @params
//   dc     = date control
//   strval = if provided, it needs to the string value to use for the existing
//            date rather than the value of the supplied date control
// @return string value that was set in dc
//-----------------------------------------------------------------------------
window.yearBack = function (dc,strval) {
    var y = dateFromString(dc.value);
    if (typeof strval == "string") {
        y = dateFromString(strval);
    }
    var d2 = dateYearBack(y);
    return setDateControl(dc, d2);
};

//-----------------------------------------------------------------------------
// setToCurrentMonth
//            This routine sets the supplied date control to the 1st of
//            the current month.
// @params
//   dc = date control
// @return string value that was set in dc
//-----------------------------------------------------------------------------
window.setToCurrentMonth = function (dc) {
    var y = new Date();
    var d2 = new Date(y.getFullYear(), y.getMonth(), 1, 0, 0, 0, 0);
    return setDateControl(dc, d2);
};

//-----------------------------------------------------------------------------
// setToNextMonth
//            This routine sets the supplied date control to the 1st of
//            the next month.
//            NOTICE: Assumes we're setting the end date of a date range.
//                    DO NOT CALL THIS ROUTINE TO SET A START DATE
// @params
//   dc = date control
// @return string value that was set in dc
//-----------------------------------------------------------------------------
window.setToNextMonth = function (dc) {
    var y = new Date();
    var my = (y.getMonth() + 1) / 12; // number of years to add for next month
    var m = (y.getMonth() + 1) % 12;  // next month
    var d2 = new Date(y.getFullYear() + my, m, 1, 0,0,0,0);
    var s = w2uiDateControlString(d2);

    // now work out the display date:
    var dispDate = d2; // assume it's mode 0

    // check EDI mode for this business and set app.D2 accordingly
    var BID = getCurrentBID();
    var BUD = getBUDfromBID(BID);
    if (EDIEnabledForBUD(BUD)) {
        dispDate.setDate(dispDate.getDate() - 1);
    }

    dc.value = w2uiDateControlString(dispDate);

    // return s;
    return dc.value;
};

//-----------------------------------------------------------------------------
// dateMonthBack - return a date which is a month prior to the supplied date
// @params
//   y = input date
// @return date which is y - 1 month
//-----------------------------------------------------------------------------
window.dateMonthBack = function (y) {
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
        //if (d == daysInCurrentMonth || d >= daysInPrevMonth) { d = daysInPrevMonth; }
        d = daysInPrevMonth;
    }
    return new Date(y.getFullYear() - yb, m, d, 0, 0, 0);
};

//-----------------------------------------------------------------------------
// monthBack - supply the date control, this function will go to the previous
//             month. It will snap the date to the end of the month if the
//             current date is the end of the month.
// @params
//   dc = date control
// @return string value that was set in dc
//-----------------------------------------------------------------------------
window.monthBack = function (dc) {
    var y = dateFromString(dc.value);
    var d2 =  dateMonthBack(y);
    return setDateControl(dc, d2);
};

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
window.setDateControl = function (dc, dt) {
    var s = w2uiDateControlString(dt);
    dc.value = s;
    return s;
};

//-----------------------------------------------------------------------------
// getTimeFromDT
//           - If the string is a datetime string, this function will return
//             the time portion. If there is no time portion, it returns null.
//             Datetime strings come in this format: 2018-02-28T17:00:00Z
//             if the T is present it will return 17:00:00Z .
// @params
//   dt = a datetime string
// @return time portion of datetime string
//         or the original string if no time is present
//-----------------------------------------------------------------------------
window.getTimeFromDT = function (dt) {
    if (typeof dt === "undefined") { return ""; }
    var i = dt.indexOf("T");
    var l = dt.length;
    var s = dt;
    if (i >= 0 &&  i+1 < l) {
        s = dt.substr(i+1,l-1);
    }
    return s;
};

//-----------------------------------------------------------------------------
// getDateFromDT
//           - If the string is a datetime string, this function will return
//             the date portion. If there is no date portion, it returns null.
//             Datetime strings come in this format: 2018-02-28T17:00:00Z
//             if the T is present it will return 2018-02-28.
// @params
//   dt = a datetime string
// @return date portion of datetime string
//         or the original string if no date is present
//-----------------------------------------------------------------------------
window.getDateFromDT = function (dt) {
    if (typeof dt === "undefined") { return ""; }
    var i = dt.indexOf("T");
    var l = dt.length;
    if (i > 0 && l > i) {
        var s = dt.substr(0,i);
        return s;
    }
    return dt;
};

//-----------------------------------------------------------------------------
// dtTextRender - enable the Statement form in toplayout.  Also, set
//                the forms url and request data from the server
// @params
//   bid = business id (or the BUD)
//    id = Task List TLID
// d1,d2 = date range to use
//-----------------------------------------------------------------------------
window.dtTextRender = function (dt, index, col_index) {
    var d = getDateFromDT(dt);
    var t = getTimeFromDT(dt);
    if (d != t) {
        return d + ' ' + t;
    }
    return d;
};

//-----------------------------------------------------------------------------
// dateFromString - return a java date value equal to the date in the supplied
//      date control.  Datetime strings come in this format: 2018-02-28T17:00:00Z
//      if the T is present, discard everthing to the right of it before
//      doing any parsing
//
//
// @params
//   dt = date or datetime string value
// @return - java date value
//-----------------------------------------------------------------------------
window.dateFromString = function (dt) {
    if (dt === null) {
        return null;
    }

    var ds = getDateFromDT(dt);

    // Strange thing about javascript dates
    // new Date("2017-06-28") gives a date with offset value with local timezone i.e, Wed Jun 28 2017 05:30:00 GMT+0530 (IST)
    // new Date("2017/06/28") gives a date without offset value with local timezone i.e, Wed Jun 28 2017 00:00:00 GMT+0530 (IST)
    ds = ds.replace(/-/g,"\/");
    ds = ds.replace(/T.+/, ''); // first replace `/` with `-` and also remove `hh:mm:ss` value we don't need it
    return new Date(ds);
};

//-----------------------------------------------------------------------------
// dateTodayStr - return a string with today's date in the form d/m/yyyy
// @params
//   <none>
// @return - formatted date string
//-----------------------------------------------------------------------------
window.dateTodayStr = function () {
    var today = new Date();
    return dateFmtStr(today);
};

//-----------------------------------------------------------------------------
// dateFmtStr - return a string with the supplied date in the form d/m/yyyy
// @params
//    date
// @return - formatted date string
//-----------------------------------------------------------------------------
window.dateFmtStr = function (t) {
    var dd = t.getDate();
    var mm = t.getMonth() + 1; //January is 0!
    var yyyy = t.getFullYear();
    return mm + '/' + dd + '/' + yyyy;
};

//-----------------------------------------------------------------------------
// datetimeFmtStr - return a string with the supplied date in the form
//                  d/m/yyyy hh:mm am/pm
// @params
//    date
// @return - formatted date string
//-----------------------------------------------------------------------------
window.datetimeFmtStr = function (t) {
    var ap = "am";
    var hh = t.getHours();
    var MM = t.getMinutes();
    var hz = "0"; // assume zero fill
    var mz = ""; // assume no fill
    if (hh >= 12) { ap = "pm"; }
    if (hh > 12)  { hh -= 12; }
    if (hh == 0)  { hh = 12; }
    if (hh > 9)   { hz = ""; }
    if (MM <= 9)  { mz = "0"; }

    return dateFmtStr(t) + " " + hz + hh + ":" + mz + MM + " " + ap;
};

//-----------------------------------------------------------------------------
// isDatePriorToCurrentDate - return boolean value
// @params
//    date object
// @return - boolean
//-----------------------------------------------------------------------------
window.isDatePriorToCurrentDate = function (date) {
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
};

$(function() {
     $(document).on("blur change", "input[type=us-date1], input[type=us-date2]", function(e) {
         // replace trailing zero from date using regex
         this.value = this.value.replace(/\b0*(?=\d)/g, '');
         if(app.dateFormatRegex.test(this.value)){
             this.style.borderColor = '#cacaca';
         } else {
            this.style.borderColor = 'red';
         }
     });
 });
