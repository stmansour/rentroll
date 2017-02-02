//-----------------------------------------------------------------------------
// setToForm -  enable form sform in toplayout.  Also, set the forms url and
//              request data from the server
// @params
//   dc = date control
// @return string value that was set in dc
function setToForm(sform,url) {
    var f = w2ui[sform];
    w2ui['toplayout'].show('right',true);
    w2ui['toplayout'].content('right', f);
    w2ui['toplayout'].sizeTo('right', 400);
    f.resize();
    if (url.length > 0) {
        f.url = url
        f.request();
    }
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
    return s + 's'
}


//-----------------------------------------------------------------------------
// dateFromDC - return a java date value equal to the date in the supplied
//              date control 
// @params
//   dc = date control
// @return - java date value
//-----------------------------------------------------------------------------
function dateFromDC(dc) {
    var x = new Date(dc.value); 
    return new Date(x.getTime() + 24*60*60*1000); // for some reason we need to add 1 day to get the right value
 }

//-----------------------------------------------------------------------------
// dayBack - supply the date control and this function will go to the previous
//           day. 
// @params
//   dc = date control
// @return string value that was set in dc
//-----------------------------------------------------------------------------
function dayBack(dc) {
    var x = dateFromDC(dc);
    var y = new Date(x.getTime() - 24*60*60*1000); // one day prior 
    return setDateControl(dc,y);
}

//-----------------------------------------------------------------------------
// dayFwd - supply the date control and this function will go to the next day. 
// @params
//   dc = date control
// @return string value that was set in dc
//-----------------------------------------------------------------------------
function dayFwd( dc ) {
    var x = dateFromDC(dc);
    var y = new Date(x.getTime() + 24*60*60*1000); // one day prior
    return setDateControl(dc,y);
}

//-----------------------------------------------------------------------------
// dateMonthFwd - return a date that is one month from the supplied date. It
//                will snap the date to the end of the month if the 
//                current date is the end of the month.
// @params
//   y = starting date
// @return - a date that is one month from y
//-----------------------------------------------------------------------------
function dateMonthFwd( y ) {
    var m = (y.getMonth() + 1) % 12;    // set m to the correct next month value
    var my = (y.getMonth() + 1) / 12;   // number of years to add for next month
    var d = y.getDate();                // this is the target date
    // console.log('dateMonthFwd: T1 -    d = ' + d);
   
    // If there is a chance that there is no such date next month, then let's make sure we 
    // do this right. If the date is > than the number of days in month m then snap as follows:
    // if d is valid in month m then use d, otherwise snap to the end of the month.
    if (d > 28) {
        var d0 = new Date(y.getFullYear()+my,m,0, 0,0,0);
        var daysInCurrentMonth = d0.getDate();
        var m2 = (y.getMonth() + 2) % 12;   // used to find # days in month m
        var m2y = (y.getMonth() + 2) / 12;  // number of years to add for month m
        var d3 = new Date(y.getFullYear() + m2y, m2, 0, 0,0,0);
        var daysInNextMonth = d3.getDate();
        if (d >= daysInNextMonth || d == daysInCurrentMonth) { d = daysInNextMonth; }
    }
    // console.log('dateMonthFwd:  m = ' + m + '   d = ' + d);
    var d2 = new Date(y.getFullYear() + my, m, d, 0,0,0);
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
function monthFwd( dc ) {
    y = dateFromDC(dc);
    var d2 = dateMonthFwd(y);
    return setDateControl(dc,d2);
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
    y = dateFromDC(dc);
    var yb = 0; // assume same year
    var m = y.getMonth() - 1;
    if (m < 0) {
    	m = 11;
    	yb = 1;	// we've gone back one year
    }
    var d = y.getDate();
    if (d >= 28) {
        var d0 = new Date(y.getFullYear(), ((y.getMonth() + 1) % 12), 0,0,0,0); // date of last day in prev month
        var daysInCurrentMonth = d0.getDate();
        var d3 = new Date(y.getFullYear() - yb, y.getMonth(), 0, 0,0,0); // date() is number of days in month y.getMonth()
        var daysInPrevMonth = d3.getDate();
        if (d == daysInCurrentMonth || d >= daysInPrevMonth) { d = daysInPrevMonth; }
    }
    var d2 = new Date(y.getFullYear() - yb, m, d, 0,0,0);
    return setDateControl(dc,d2);
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
function setDateControl(dc,dt) {
    s = dateControlString(dt)
    dc.value = s;
    return s
}
