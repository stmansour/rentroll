/*global
    w2ui, app, console
*/


//-----------------------------------------------------------------------------
// getBUDfromBID  - given the BID return the associated BUD. Returns
//                  an empty string if BID is not found
// @params  BUD   - the BUD for the business of interest
//          PMTID - the payment type id for which we want the name
// @return  the BUD (or empty string if not found)
//-----------------------------------------------------------------------------
function getBUDfromBID(BID) {
    "use strict";
    var BUD = '';
    for (var i=0; i<app.BizMap.length; i++) {
        if (BID == app.BizMap[i].BID) {
            BUD = app.BizMap[i].BUD;
        }
    }
    return BUD;
}

//-----------------------------------------------------------------------------
// getBIDfromBUD  - given the BUD return the associated BID. Returns
//                  undefined if BUD is not found
// @params  BUD   - the BUD for the business of interest
// @return  the BID (or `undefined` if not found)
//-----------------------------------------------------------------------------
function getBIDfromBUD(BUD) {
    "use strict";
    var BID;
    for (var i=0; i<app.BizMap.length; i++) {
        if (BUD == app.BizMap[i].BUD) {
            BID = app.BizMap[i].BID;
        }
    }
    return BID;
}

//-----------------------------------------------------------------------------
// getPaymentTypeName - searches BUD's Payment Types for PMTID.  If found the
//                  Name is returned, else an empty string is returned.
// @params  BUD   - the BUD for the business of interest
//          PMTID - the payment type id for which we want the name
// @return  the Payment Type Name (or empty string if not found)
//-----------------------------------------------------------------------------
function getPaymentTypeName(BUD,PMTID) {
    "use strict";
    if (typeof BUD == "undefined") {
        return '';
    }
    for (var i = 0; i < app.pmtTypes[BUD].length; i++ ) {
        if (app.pmtTypes[BUD][i].PMTID == PMTID) {
            return app.pmtTypes[BUD][i].Name;
        }
    }
    return '';
}

//-----------------------------------------------------------------------------
// getPaymentTypeID - searches BUD's Payment Types for Name.  If found the
//                  PMTID is returned. Otherwise it returns -1
// @params  BUD   - the BUD for the business of interest
//          Name  - the Name of the payment type
// @return  PMTID (or -1 if not found)
//-----------------------------------------------------------------------------
function getPaymentTypeID(BUD,Name) {
    "use strict";
    if (typeof BUD == "undefined") {
        return -1;
    }
    for (var i = 0; i < app.pmtTypes[BUD].length; i++ ) {
        if (app.pmtTypes[BUD][i].Name == Name) {
            return app.pmtTypes[BUD][i].PMTID;
        }
    }
    return -1;
}

//-----------------------------------------------------------------------------
// buildPaymentTypeOptions - creates a list suitable for a dropdown menu
//                  with the payment types for the supplied BUD
// @params  BUD   - the BUD for the business of interest
// @return  the list of Payment Type Names (or empty list if BUD not found)
//-----------------------------------------------------------------------------
function buildPaymentTypeOptions(BUD) {
    "use strict";
    var options = [];
    if (typeof BUD == "undefined") {
        return options;
    }
    for (var i = 0; i < app.pmtTypes[BUD].length; i++ ) {
        options[i] = app.pmtTypes[BUD][i].Name;
    }
    return options;
}

//-----------------------------------------------------------------------------
// buildPaymentTypeSelectList - creates a list suitable for a dropdown menu
//                  with the payment types for the supplied BUD
// @params  BUD   - the BUD for the business of interest
// @return  the list of Payment Type Names (or empty list if BUD not found)
//-----------------------------------------------------------------------------
function buildPaymentTypeSelectList(BUD) {
    "use strict";
    var options = [];
    if (typeof BUD == "undefined") {
        return options;
    }
    for (var i = 0; i < app.pmtTypes[BUD].length; i++ ) {
        options[i] = {id: app.pmtTypes[BUD][i].PMTID, text: app.pmtTypes[BUD][i].Name};
    }
    return options;
}

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
//   sform   = name of the form
//   url     = request URL for the form
//   [width] = optional, if specified it is the width of the form
//-----------------------------------------------------------------------------
function setToForm(sform, url, width, doRequest) {
    "use strict";

    var f = w2ui[sform];
    if (url.length > 0) {
        f.url = url;
        if (typeof f.tabs.name == "string") {
            f.tabs.click('tab1');
        }

        if (doRequest) {
            f.request(function(/*event*/) {
                // only render the toplayout after server has sent down data
                // so that w2ui can bind values with field's html control,
                // otherwise it is unable to find html controls
                w2ui.toplayout.show('right', true);
                w2ui.toplayout.content('right', f);
                w2ui.toplayout.sizeTo('right', width);
                w2ui.toplayout.render();
            });
        }
        else {
            w2ui.toplayout.show('right', true);
            w2ui.toplayout.content('right', f);
            w2ui.toplayout.sizeTo('right', width);
            w2ui.toplayout.render();
        }
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
    if (raid > 0) {
        w2ui.toplayout.content('right', w2ui.raLayout);
        w2ui.toplayout.show('right', true);
        w2ui.toplayout.sizeTo('right', app.WidestFormWidth);
        w2ui.rentalagrForm.url = '/v1/rentalagr/' + bid + '/' + raid;
        w2ui.rentalagrForm.request();
        w2ui.toplayout.render();
    }

    //----------------------------------------------------------------
    // Get the associated Rentables...
    //      /v1/rar/bid/raid[?d1=2017-02-1]
    //      if no date is specified, today's date is used as the default.
    //----------------------------------------------------------------
    w2ui.rarGrid.url = '/v1/rar/' + bid + '/' + raid;
    // console.log('rar url = ' + w2ui.rarGrid.url);
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
    // console.log('rapGrid url = ' + w2ui.rapGrid.url);
    w2ui.rapGrid.request();
    w2ui.rapGrid.header = plural(app.sPayor) + ' as of ' + dateFmtStr(d);

    //----------------------------------------------------------------
    // Get the associated Users...
    //      /v1/ruser/bid/raid[?&d1=2017-02-1]
    //      if no date is specified, today's date is used as the default.
    //----------------------------------------------------------------
    w2ui.rauGrid.url = '/v1/ruser/' + bid + '/' + raid;
    // console.log('rauGrid url = ' + w2ui.rauGrid.url);
    w2ui.rauGrid.request();
    w2ui.rauGrid.header = plural(app.sUser) + ' as of ' + dateFmtStr(d);

    //----------------------------------------------------------------
    // Get the associated Pets...
    //      /v1/xrapets/bid/raid
    //----------------------------------------------------------------
    w2ui.raPetGrid.url = '/v1/rapets/' + bid + '/' + raid;
    // console.log('xrapets url = ' + w2ui.rarGrid.url);
    w2ui.raPetGrid.request();
    w2ui.raPetGrid.header = 'Pets as of ' + dateFmtStr(d);
}

//-----------------------------------------------------------------------------
// ridRentablePickerRender - renders a name during typedown.
// @params
//   item = an object with RentableName
// @return - true if the names match, false otherwise
//-----------------------------------------------------------------------------
function ridRentablePickerRender(item) {
    "use strict";
    w2ui.ridRentablePicker.record.RID = item.recid;
    return item.RentableName + '  (RID: ' + item.recid + ')';
}

//-----------------------------------------------------------------------------
// asmFormRentablePickerRender - renders a name during typedown.
// @params
//   item = Object with RentableName
// @return - true if the names match, false otherwise
//-----------------------------------------------------------------------------
function asmFormRentablePickerRender(item) {
    "use strict";
    w2ui.asmEpochForm.record.RID = item.recid;
    return item.RentableName + '  (RID: ' + item.recid + ')';
}

//-----------------------------------------------------------------------------
// ridRentableDropRender - renders a name during typedown.
// @params
//   item = an object with RentableName
// @return - the name to render
//-----------------------------------------------------------------------------
function ridRentableDropRender (item) {
    "use strict";
    // w2ui.ridRentablePicker.RID = item.RID;
    return item.RentableName + '  (RID: ' + item.recid + ')';
}

//-----------------------------------------------------------------------------
// ridRentableCompare - Compare two items to see if they match
// @params
//   item = an object assumed to have a RentableName
// @return - true if the names match, false otherwise
//-----------------------------------------------------------------------------
function ridRentableCompare(item, search) {
    "use strict";
    var s = item.RentableName.toLowerCase();
    return s.includes(search.toLowerCase());
}

//-----------------------------------------------------------------------------
// tcidRAPayorPickerRender - renders a name during typedown.
// @params
//   item = an object assumed to have a FirstName and LastName
// @return - true if the names match, false otherwise
//-----------------------------------------------------------------------------
function tcidRAPayorPickerRender(item) {
    "use strict";
    var s="";
    if (item.IsCompany > 0) {
        s = item.CompanyName;
    } else {
        s = item.FirstName + ' ' + item.LastName;
    }
    w2ui.tcidRAPayorPicker.record = {
        TCID: item.TCID,
        pickedName: s,
        DtStart: w2ui.tcidRAPayorPicker.record.DtStart,
        DtStop: w2ui.tcidRAPayorPicker.record.DtStop,
        FirstName: item.FirstName,
        LastName: item.LastName,
        IsCompany: item.IsCompany,
        CompanyName: item.CompanyName
    };
    return s;
}

//-----------------------------------------------------------------------------
// getFullName - returns a string with the full name based on the item supplied.
// @params
//   item = an object assumed to have a FirstName, MiddleName, and LastName
// @return - the full name concatenated together
//-----------------------------------------------------------------------------
function getFullName(item) {
    "use strict";
    var s = item.FirstName;
    if (item.MiddleName.length > 0) { s += ' ' + item.MiddleName; }
    if (item.LastName.length > 0 ) { s += ' ' + item.LastName; }
    return s;
}

//-----------------------------------------------------------------------------
// getTCIDName - returns an appropriate name for the supplied item. If
//          the item is a person, then the person's full name is returned.
//          If the item is a company, then the company name is returned.
// @params
//   item = an object assumed to have a FirstName, MiddleName, LastName,
//          IsCompany, and CompanyName.
// @return - the name to render
//-----------------------------------------------------------------------------
function getTCIDName(item) {
    "use strict";
    var s = (item.IsCompany > 0) ? item.CompanyName : getFullName(item);
    if (item.TCID > 0) { s += ' (TCID: '+ String(item.TCID) +')'; }
    return s;
}

//-----------------------------------------------------------------------------
// tcidPickerCompare - Compare item to the search string. Verify that the
//          supplied search string can be found in item
// @params
//   item = an object assumed to have a FirstName and LastName
// @return - true if the search string is found, false otherwise
//-----------------------------------------------------------------------------
function tcidPickerCompare(item, search) {
    "use strict";
    var s = getTCIDName(item);
    s = s.toLowerCase();
    var srch = search.toLowerCase();
    var match = (s.indexOf(srch) >= 0);
    return match;
}

//-----------------------------------------------------------------------------
// tcidPickerDropRender - renders a name during typedown.
// @params
//   item = an object assumed to have a FirstName and LastName
// @return - the name to render
//-----------------------------------------------------------------------------
function tcidPickerDropRender(item) {
    "use strict";
    return getTCIDName(item);
}

//-----------------------------------------------------------------------------
// tcidReceiptPayorPickerRender - renders a name during typedown in the
//          receiptForm. It also sets the TCID for the record.
// @params
//   item = an object assumed to have a FirstName and LastName
// @return - true if the names match, false otherwise
//-----------------------------------------------------------------------------
function tcidReceiptPayorPickerRender(item) {
    "use strict";
    var s = getTCIDName(item);
    w2ui.receiptForm.record.TCID = item.TCID;
    w2ui.receiptForm.record.Payor = s;
    return s;
}

//-----------------------------------------------------------------------------
// tcidRUserPickerRender - renders a name during typedown.
// @params
//   item = an object assumed to have a FirstName and LastName
// @return - true if the names match, false otherwise
//-----------------------------------------------------------------------------
function tcidRUserPickerRender(item) {
    "use strict";
    var s;
    if (item.IsCompany > 0) {
        s = item.CompanyName;
    } else {
        s = item.FirstName + ' ' + item.LastName;
    }

    w2ui.tcidRUserPicker.record = {
        TCID: item.TCID,
        pickedName: s,
        DtStart: w2ui.tcidRUserPicker.record.DtStart,
        DtStop: w2ui.tcidRUserPicker.record.DtStop,
        FirstName: item.FirstName,
        LastName: item.LastName,
        IsCompany: item.IsCompany,
        CompanyName: item.CompanyName
    };
    return s;
}

// ############################################################################


//-----------------------------------------------------------------------------
// rentalAgrFinderCompare - Compare item to the search string. Verify that the
//          supplied search string can be found in item
// @params
//   item = an object assumed to have a FirstName and LastName
// @return - true if the search string is found, false otherwise
//-----------------------------------------------------------------------------
function rentalAgrFinderCompare(item, search) {
    "use strict";
    var s = getTCIDName(item);
    s = s.toLowerCase();
    var srch = search.toLowerCase();
    var match = (s.indexOf(srch) >= 0);
    return match;
}

//-----------------------------------------------------------------------------
// rentalAgrFinderDropRender - renders a name during typedown.
// @params
//   item = an object assumed to have a FirstName and LastName
// @return - the name to render
//-----------------------------------------------------------------------------
function rentalAgrFinderDropRender(item) {
    "use strict";
    return getTCIDName(item);
}

//-----------------------------------------------------------------------------
// rentalAgrFinderRender - renders a name during typedown in the
//          rentalAgrFinder. It also sets the TCID for the record.
// @params
//   item = an object assumed to have a FirstName and LastName
// @return - true if the names match, false otherwise
//-----------------------------------------------------------------------------
function rentalAgrFinderRender(item) {
    "use strict";
    var s = getTCIDName(item);
    w2ui.rentalAgrFinder.record.TCID = item.TCID;
    w2ui.rentalAgrFinder.record.Payor = s;
    w2ui.rentalAgrFinder.record.RAID = item.RAID;
    return s;
}

//-----------------------------------------------------------------------------
// rentalAgrFinderRender - renders a name during typedown.
// @params
//   item = an object assumed to have a FirstName and LastName
// @return - true if the names match, false otherwise
//-----------------------------------------------------------------------------
function rentalAgrFinderRender(item) {
    "use strict";
    var s;
    if (item.IsCompany > 0) {
        s = item.CompanyName;
    } else {
        s = item.FirstName + ' ' + item.LastName;
    }

    w2ui.rentalAgrFinder.record = {
        TCID: item.TCID,
        RAID: item.RAID,
        PayorName: s,
        FirstName: item.FirstName,
        MiddleName: item.MiddleName,
        LastName: item.LastName,
        IsCompany: item.IsCompany,
        CompanyName: item.CompanyName,
        RID: item.RID,
    };

    // we need to get the rentables associated with item.RAID
    var url = '/v1/rar/' + app.RentalAgrFinder.BID + '/' + item.RAID;
    $.get(url,function(data,status) {
        app.RentalAgrFinder.RAR = JSON.parse(data);
        app.RentalAgrFinder.RARentablesNames = [];
        for (var i = 0; i < app.RentalAgrFinder.RAR.records.length; i++) {
            app.RentalAgrFinder.RARentablesNames.push(
                { id: app.RentalAgrFinder.RAR.records[i].RID, text: app.RentalAgrFinder.RAR.records[i].RentableName} );
        }
        console.log('calling rentalAgrFinder.refresh(), app.RentalAgrFinder.RARentablesNames.length = ' + app.RentalAgrFinder.RARentablesNames.length );
        w2ui.rentalAgrFinder.refresh();
    });


    return s;
}

// ############################################################################


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
// dateFromString - return a java date value equal to the date in the supplied
//              date control
// @params
//   ds = date string value
// @return - java date value
//-----------------------------------------------------------------------------
function dateFromString(ds) {
    "use strict";

    // Strange thing about javascript dates
    // new Date("2017-06-28") gives a date with offset value with local timezone i.e, Wed Jun 28 2017 05:30:00 GMT+0530 (IST)
    // new Date("2017/06/28") gives a date without offset value with local timezone i.e, Wed Jun 28 2017 00:00:00 GMT+0530 (IST)

    ds = ds.replace(/-/g,"\/").replace(/T.+/, ''); // first replace `/` with `-` and also remove `hh:mm:ss` value we don't need it
    return new Date(ds);

    // var re = /^([0-9]{4})[-]([0-9]{2})[-]([0-9]{2})$/; // regex pattern to satisfy date pattern `yyyy/mm/dd`

    // var valid = re.test(ds);
    // // if datestring does not satisfy the pattern then simply return null
    // if (!valid) {
    //     return null;
    // }

    // // now execute regex pattern for the string
    // var match = re.exec(ds);

    // // get year, month, date value
    // var y = match[1], m = match[2], d = match[3];
    // return new Date(y, m-1, d); // month starts from 0 index so needs to substract by 1
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
    "use strict";
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
    "use strict";
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
    "use strict";
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
    "use strict";
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
    "use strict";
    var y = dateFromString(dc.value);
    var d2 =  dateMonthBack(y);
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

//-----------------------------------------------------------------------------
// calcRarGridContractRent
//          - Sum the Contract Rent column of rarGrid and return the total.
//            used to set the control.
// @params
//          grid - The grid to work on
// @return  The total of the column
//-----------------------------------------------------------------------------
function calcRarGridContractRent(grid) {
    "use strict";
    grid = w2ui.rarGrid || grid;
    var chgs = grid.getChanges();
    var amts = [];
    //
    // Build up a list of amounts...
    //
    for (var i = 0; i < grid.records.length; i++) {
        if (typeof grid.records[i].ContractRent == "number") {
            amts.push({ recid: grid.records[i].recid, ContractRent: grid.records[i].ContractRent });
        }
    }
    //
    // Any changes override these ContractRents...
    //
    for (i = 0; i < chgs.length; i++) {
        if (typeof chgs[i].ContractRent == "number") {
            for (var j = 0; j < amts.length; j++) {
                if (chgs[i].recid == amts[j].recid) {
                    amts[j] = { recid: chgs[i].recid, ContractRent: chgs[i].ContractRent };
                    break;
                }
            }
        }
    }
    // now total everything...
    var total = 0.0;
    for (i = 0; i < amts.length; i++) {
        total += amts[i].ContractRent;
    }
    grid.set('s-1', { ContractRent: total });
}

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
    "use strict";
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
    "use strict";
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
    "use strict";
    var html1 = '<div style="padding: 0px 5px;">From: <input type="date" name="' + prefix + 'D1"></div>';
    var html2 = '<div style="padding: 0px 5px;">To: <input type="date" name="' + prefix + 'D2">' + '</div>';
    var tmp = [{ type: 'break', id: 'break1' },
        { type: 'button', id: 'monthback', icon: 'fa fa-backward', tooltip: 'month back' },
        { type: 'button', id: 'dayback', icon: 'fa fa-chevron-circle-left', tooltip: 'day back' },
        { type: 'html', id: 'D1', html: function() {return html1; } },
        { type: 'button', id: 'today', icon: 'fa fa-circle-o', tooltip: 'present month' },
        { type: 'html', id: 'D2', html: function() {return html2; } },
        { type: 'button', id: 'dayfwd', icon: 'fa fa-chevron-circle-right', tooltip: 'day forward' },
        { type: 'button', id: 'monthfwd', icon: 'fa fa-forward', tooltip: 'month forward' },
    ];
    return tmp;
}

//-----------------------------------------------------------------------------
// getRentableTypes - return the RentableTypes list with respect of BUD
// @params
// @return  the Rentable Types List
//-----------------------------------------------------------------------------
function getRentableTypes(BID) {
    "use strict";
    return jQuery.ajax({
        type: "GET",
        url: "/v1/rtlist/"+BID,
        dataType: "json",
    }).done(function(data) {
        if (data.status == "success") {
            var BUD = getBUDfromBID(BID);
            if (data.records) {
                app.rt_list[BUD] = data.records;
            } else {
                app.rt_list[BUD] = [];
            }
        }
    });
}

//-----------------------------------------------------------------------------
// getAccountsList - return the GLAccounts list with respect of BUD
// @params
// @return the list of accounts
//-----------------------------------------------------------------------------
function getAccountsList(BID) {
    "use strict";
    return jQuery.ajax({
        type: "GET",
        url: "/v1/accountlist/"+BID,
        dataType: "json",
    }).done(function(data) {
        if (data.status == "success") {
            var BUD = getBUDfromBID(BID);
            if (data.records) {
                app.gl_accounts[BUD] = data.records;
            } else{
                app.gl_accounts[BUD] = [];
            }
        }
    });
}

//-----------------------------------------------------------------------------
// getPostAccounts - return the list of post accounts with respect of BUD
// @params
// @return the list of post accounts
//-----------------------------------------------------------------------------
function getPostAccounts(BID) {
    "use strict";
    return jQuery.ajax({
        type: "GET",
        url: "/v1/postaccounts/"+BID,
        dataType: "json",
    }).done(function(data) {
        if (data.status == "success") {
            var BUD = getBUDfromBID(BID);
            if (data.records) {
                app.post_accounts[BUD] = data.records;
            } else{
                app.post_accounts[BUD] = [];
            }
        }
    });
}

//-----------------------------------------------------------------------------
// getParentAccounts - return the list of Parent accounts with respect of BUD
// @params
// @return the list of parent accounts
//-----------------------------------------------------------------------------
function getParentAccounts(BID) {
    "use strict";
    return jQuery.ajax({
        type: "GET",
        url: "/v1/parentaccounts/"+BID,
        dataType: "json",
    }).done(function(data) {
        if (data.status == "success") {
            var BUD = getBUDfromBID(BID);
            if (data.records) {
                app.parent_accounts[BUD] = data.records;
            } else{
                app.parent_accounts[BUD] = [];
            }
        }
    });
}

//-----------------------------------------------------------------------------
// unallocAmountRemaining - based on the amounts allocated to receipts in the
// unpaid receipts list, compute the amount of funds remaining to be allocated
// and display it.
// @params
// @return
//-----------------------------------------------------------------------------
function unallocAmountRemaining() {
    "use strict";
    var totalFunds = app.payor_fund; // must already be set to total unallocated receipt funds
    for (var i=0; i < w2ui.unpaidASMsGrid.records.length; i++) {
        totalFunds -= w2ui.unpaidASMsGrid.records[i].Allocate;
    }
    // var dispAmt = parseFloat(totalFunds).toFixed( 2 );
    var dispAmt = number_format(totalFunds, 2, '.', ',');
    var x = document.getElementById("total_fund_amount");
    if (x !== null) {
        x.innerHTML = dispAmt;
    }
}

//-----------------------------------------------------------------------------
// refreshUnallocAmtSummaries - This routine totals the summary columns for the
// unpaid assessments grid.
// @params
// @return
//-----------------------------------------------------------------------------
function refreshUnallocAmtSummaries() {
    "use strict";
    if (w2ui.unpaidASMsGrid.records.length === 0 ) { return; }
    var amt = 0;
    var amtPaid = 0;
    var amtOwed = 0;
    var alloc = 0;
    for (var i=0; i < w2ui.unpaidASMsGrid.records.length; i++) {
        amt += w2ui.unpaidASMsGrid.records[i].Amount;
        amtPaid += w2ui.unpaidASMsGrid.records[i].AmountPaid;
        amtOwed += w2ui.unpaidASMsGrid.records[i].AmountOwed;
        alloc += w2ui.unpaidASMsGrid.records[i].Allocate;
    }
    w2ui.unpaidASMsGrid.set('s-1', {Amount: amt, AmountPaid: amtPaid, AmountOwed: amtOwed, Allocate: alloc});
}


// int_to_bool converts int to bool. i.e, 0: false, 1: true
function int_to_bool(i){
    "use strict";
    if (i>0) {
        return true;
    } else {
        return false;
    }
}

// unallocated receipts utility literal object
var _unAllocRcpts = {
    layoutPanels: {
        top: function(unallocFund, person/*, tcid*/) {
            "use strict";
            return `<html>
                <head>
                <link rel="stylesheet" href="/html/rentroll.css">
                </head>
                <body>
                <div style="display: table; width: 100%; height: 40%;">
                    <div style="display: table-cell; vertical-align: middle;text-align: left;width: 100%;">
                        <p style="margin: 5px auto;font-size: 1.5rem;font-weight: bold;" name="unallocForm">`+person+`<br>Remaining unallocated funds:
                        <span id="total_fund_amount" data-fund="`+unallocFund+`"
                        style="padding: 10px; color: #00AA00; font-size: 1.5rem; font-weight: bold; margin: 10px auto; width: 30%;">`+
                        unallocFund+
                        `</span>&nbsp;&nbsp;&nbsp;&nbsp;<button class="w2ui-btn w2ui-btn-green" style="font-size: 1.1rem;" id="auto_allocate_btn">Auto-Allocate</button></p>
                    </div>
                </div>
                </body>
                </html>`;
        },
        bottom: function() {
            "use strict";
            return `<div style="display: table; width: 100%; height: 100%;">
                    <div style="display: table-cell; vertical-align: middle;text-align: center;width: 100%;">
                        <button class="w2ui-btn" id="alloc_fund_save_btn">Save</button>
                    </div>
                </div>`;
        }
    }
};

//-----------------------------------------------------------------------------
// getPayorFund - get payor fund
// @params
// @return  the jquery promise
//-----------------------------------------------------------------------------
function getPayorFund(BID, TCID) {
    "use strict";
    return jQuery.ajax({
        type: "GET",
        url: '/v1/payorfund/'+BID+'/'+TCID,
        dataType: "json",
    });
}

// Auto Allocate amount for each unpaid assessment
jQuery(document).on('click', '#auto_allocate_btn', function(/*event*/) {
    "use strict";
    var fund = app.payor_fund;
    var grid = w2ui.unpaidASMsGrid;

    for (var i = 0; i < grid.records.length; i++) {
        if (fund <= 0) {
            break;
        }

        // if it has already been paid, then move on to the next record
        if (grid.records[i].Amount - grid.records[i].AmountPaid <= 0) {
            continue;
        }

        // check if fully paid or not
        if (grid.records[i].Amount - grid.records[i].AmountPaid <= fund){
            grid.records[i].Allocate = grid.records[i].Amount - grid.records[i].AmountPaid;
            grid.set(grid.records[i].recid, grid.records[i]);
        } else {
            grid.records[i].Allocate = fund;
            grid.set(grid.records[i].recid, grid.records[i]);
        }

        // decrement fund value by whatever the amount allocated for each record
        fund = fund - grid.records[i].Allocate;
    }
    refreshUnallocAmtSummaries();
    unallocAmountRemaining();
    return false;
});

jQuery(document).on('click', '#alloc_fund_save_btn', function(/*event*/) {
    "use strict";
    var tgrid = w2ui.allocfundsGrid;
    var rec = tgrid.getSelection();
    if (rec.length < 0) {
        return;
    }

    // rec = tgrid.get(rec[0]);
    var tcid = app.TmpTCID,
        x = getCurrentBusiness();
    var bid = parseInt(x.value,10);


    var params = {cmd: 'save', TCID: tcid, BID: bid, records: w2ui.unpaidASMsGrid.records };
    var dat = JSON.stringify(params);

    // submit request
    $.post('/v1/allocfunds/'+bid+'/', dat)
    .done(function(data) {
        if (data.status != "success") {
            return;
        }
        w2ui.toplayout.hide('right',true);
        w2ui.toplayout.render();
        tgrid.reload();
    })
    .fail(function(/*data*/){
        console.log("Payor Fund Allocation failed.");
    });
});

//-----------------------------------------------------------------------------
// getFormSubmitData - get form submit data
// @params, w2ui form record object
// @return
// @description Helps to build form submit data, it modify record object so that each
// item in record has just a value instead of another object
//-----------------------------------------------------------------------------
function getFormSubmitData(record) {
    "use strict";

    // check that it is typeof object or not
    if (typeof record !== "object") {
        return;
    }

    // iterate over each record
    for(var key in record) {
        var item = record[key];
        if (typeof item === "object" && item !== null) {
            record[key] = item.id;
        }
    }

    return record;
}

//-----------------------------------------------------------------------------
// isNewFormRecord -  based on condition perform some actions
// 1. show / hide delete button for requested form if form has delete button
// @params
//   sform   = name of the form
//   is_new     = true / false
//-----------------------------------------------------------------------------
function isNewFormRecord(sform, is_new) {
    "use strict";
    if (is_new) {
        $("#"+sform).find("button[name=delete]").addClass("hidden");
    }
    else {
        $("#"+sform).find("button[name=delete]").removeClass("hidden");
    }
}

function number_format(number, decimals, dec_point, thousands_sep) {
    // http://kevin.vanzonneveld.net
    // +   original by: Jonas Raoni Soares Silva (http://www.jsfromhell.com)
    // +   improved by: Kevin van Zonneveld (http://kevin.vanzonneveld.net)
    // +     bugfix by: Michael White (http://getsprink.com)
    // +     bugfix by: Benjamin Lupton
    // +     bugfix by: Allan Jensen (http://www.winternet.no)
    // +    revised by: Jonas Raoni Soares Silva (http://www.jsfromhell.com)
    // +     bugfix by: Howard Yeend
    // +    revised by: Luke Smith (http://lucassmith.name)
    // +     bugfix by: Diogo Resende
    // +     bugfix by: Rival
    // +      input by: Kheang Hok Chin (http://www.distantia.ca/)
    // +   improved by: davook
    // +   improved by: Brett Zamir (http://brett-zamir.me)
    // +      input by: Jay Klehr
    // +   improved by: Brett Zamir (http://brett-zamir.me)
    // +      input by: Amir Habibi (http://www.residence-mixte.com/)
    // +     bugfix by: Brett Zamir (http://brett-zamir.me)
    // +   improved by: Theriault
    // +   improved by: Drew Noakes
    // *     example 1: number_format(1234.56);
    // *     returns 1: '1,235'
    // *     example 2: number_format(1234.56, 2, ',', ' ');
    // *     returns 2: '1 234,56'
    // *     example 3: number_format(1234.5678, 2, '.', '');
    // *     returns 3: '1234.57'
    // *     example 4: number_format(67, 2, ',', '.');
    // *     returns 4: '67,00'
    // *     example 5: number_format(1000);
    // *     returns 5: '1,000'
    // *     example 6: number_format(67.311, 2);
    // *     returns 6: '67.31'
    // *     example 7: number_format(1000.55, 1);
    // *     returns 7: '1,000.6'
    // *     example 8: number_format(67000, 5, ',', '.');
    // *     returns 8: '67.000,00000'
    // *     example 9: number_format(0.9, 0);
    // *     returns 9: '1'
    // *    example 10: number_format('1.20', 2);
    // *    returns 10: '1.20'
    // *    example 11: number_format('1.20', 4);
    // *    returns 11: '1.2000'
    // *    example 12: number_format('1.2000', 3);
    // *    returns 12: '1.200'
    var n = !isFinite(+number) ? 0 : +number,
        prec = !isFinite(+decimals) ? 0 : Math.abs(decimals),
        sep = (typeof thousands_sep === 'undefined') ? ',' : thousands_sep,
        dec = (typeof dec_point === 'undefined') ? '.' : dec_point,
        toFixedFix = function (n, prec) {
            // Fix for IE parseFloat(0.55).toFixed(0) = 0;
            var k = Math.pow(10, prec);
            return Math.round(n * k) / k;
        },
        s = (prec ? toFixedFix(n, prec) : Math.round(n)).toString().split('.');
    if (s[0].length > 3) {
        s[0] = s[0].replace(/\B(?=(?:\d{3})+(?!\d))/g, sep);
    }
    if ((s[1] || '').length < prec) {
        s[1] = s[1] || '';
        s[1] += new Array(prec - s[1].length + 1).join('0');
    }
    return s.join(dec);
}

// var exampleNumber = 1;
// function test(expected, number, decimals, dec_point, thousands_sep)
// {
//     var actual = number_format(number, decimals, dec_point, thousands_sep);
//     var result = document.createElement('div');
//     if (actual !== expected)
//     {
//         debugger;
//         result.textContent =
//             'Test case ' + exampleNumber + ' failed. ' +
//             'Expected "' + expected + '" but got "' + actual + '"';
//     } else {
//         result.textContent = 'Test case ' + exampleNumber + ' passed.'
//     }
//     document.getElementById('container').appendChild(result);
//     exampleNumber++;
// }

// test('1,235',    1234.56);
// test('1 234,56', 1234.56, 2, ',', ' ');
// test('1234.57',  1234.5678, 2, '.', '');
// test('67,00',    67, 2, ',', '.');
// test('1,000',    1000);
// test('67.31',    67.311, 2);
// test('1,000.6',  1000.55, 1);
// test('67.000,00000', 67000, 5, ',', '.');
// test('1',        0.9, 0);
// test('1.20',     '1.20', 2);
// test('1.2000',   '1.20', 4);
// test('1.200',    '1.2000', 3);
