/*global
    app, w2ui, $, form_dirty_alert, jQuery, console, w2popup, number_format, getFullName, getTCIDName, finishReportCSV,
    finishReportPDF

*/

"use strict";
// ---------------------------------------------------------------------------------
// String format: https://gist.github.com/tbranyen/1049426 (if want to format object, array as well)
// Reference: https://stackoverflow.com/questions/610406/javascript-equivalent-to-printf-string-format
// ---------------------------------------------------------------------------------
// > "{0} is awesome {1}".format("javascript", "!?")
// > "javascript is awesome !?"
// ---------------------------------------------------------------------------------
String.prototype.format = function() {
    var args = arguments;
    return this.replace(/{(\d+)}/g, function(match, number) {
        return typeof args[number] != 'undefined'? args[number] : match;
    });
};

//---------------------------------------------------------------------------------
// getCookieValue - looks for a cookie with the supplied name. If found it returns
//          the cookie value. Otherwise it returns null
//
// @params  name  - name of the cookie
// @returns the value of the cookie if found, null if not found
//---------------------------------------------------------------------------------
window.getCookieValue = function (name) {
    var nameEQ = name + "=";
    var ca = document.cookie.split(';');
    for(var i=0;i < ca.length;i++) {
        var c = ca[i];
        while (c.charAt(0)==' ') c = c.substring(1,c.length);
        if (c.indexOf(nameEQ) === 0) return c.substring(nameEQ.length,c.length);
    }
    return null;
};

//---------------------------------------------------------------------------------
// deleteCookie - looks for a cookie with the supplied name. If found it returns
//          the cookie value. Otherwise it returns null
//
// @params  name  - name of the cookie
// @returns nothing at this time
//---------------------------------------------------------------------------------
window.deleteCookie = function (name) {
  document.cookie = name +'=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';
};

//---------------------------------------------------------------------------------
// ChangeBusiness updates the UI to the newly selected business.
// This routine is indeed used, in spite of what JSHint thinks. It
// is embedded in a string defining the OnClick handler for a button
// in the main toolbar.
// ---------------------------------------------------------------------------------
window.ChangeBusiness = function () {
    var bizName = $("select[name=BusinessSelect]").find(":selected").attr("name"),
        bizVal = $("select[name=BusinessSelect]").val();

    // if same business value then nothing to do
    if (bizName === app.last.BUD) {
        return;
    }

    w2ui.toplayout.content('main', ' ');
    w2ui.toplayout.hide('right',true);
    var s = w2ui.sidebarL1;
    var sel = s.selected;
    if (sel !== null) {
        s.unselect(sel);
    }

    var no_callBack = function() {
        // revert back to last one
        $("select[name=BusinessSelect] option[value="+app.last.BID+"]").prop("selected", true);
        return false;
    },
    yes_callBack = function() {
        w2ui.toplayout.content('main', ' ');
        w2ui.toplayout.hide('right',true);
        var s = w2ui.sidebarL1;
        var sel = s.selected;
        if (sel !== null) {
            s.unselect(sel);
        }
        w2ui.reportslayout.load('main','/webclient/html/blank.html');
        app.last.report = '';

        // check EDI mode for this business and set app.D2 accordingly
        // get last selected biz flags
        var d2 = dateFromString(app.D2);
        var lastBizEDIEnabled = EDIEnabledForBUD(app.last.BUD);
        var selBizEDIEnabled = EDIEnabledForBUD(bizName);

        // -> if EDI enabled for both then nothing to do
        // -> if EDI disabled for both then nothing to do
        // -> if EDI enabled for last and not for selected one then add one day in app.D2
        if (lastBizEDIEnabled && !selBizEDIEnabled) {
            d2.setDate(d2.getDate() + 1);
            app.D2 = dateControlString(d2);
        }
        // -> if EDI not enabled for last one and enabled for selected one then subtract one day in app.D2
        if (!lastBizEDIEnabled && selBizEDIEnabled) {
            d2.setDate(d2.getDate() - 1);
            app.D2 = dateControlString(d2);
        }

        app.last.BUD = bizName;
        app.last.BID = bizVal;

        s.collapse('reports');
        return true;
    };

    // warn user if active form has been changed
    form_dirty_alert(yes_callBack, no_callBack);
};

//---------------------------------------------------------------------------------
// getGridReversalSymbolHTML - returns the HTML to insert into a grid cell to
//          indicate that the record is reversed
//
// @params  <none>
// @returns a string with HTML
//---------------------------------------------------------------------------------
window.getGridReversalSymbolHTML = function () {
    return '<i class="fas fa-exclamation-triangle" title="reversed" aria-hidden="true" style="color: #FFA500;"></i>';
};

//---------------------------------------------------------------------------------
// get2XReversalSymbolHTML - returns the HTML to insert into a grid cell to
//          indicate that the record is reversed
//
// @params  <none>
// @returns a string with HTML
//---------------------------------------------------------------------------------
window.get2XReversalSymbolHTML = function () {
    return "<div class='reverseIconContainer'><i class='fas fa-exclamation-triangle fa-2x reverseIcon' aria-hidden='true'></i></div>";
};

//---------------------------------------------------------------------------------
// switchToGrid - changes the main view of the program to a grid with
//                variable name svc + 'Grid'
//
// @params  svc = prefix of grid name
//          svcOverride = name of webservice to call if the name does not
//                match the name of the svc
//
//---------------------------------------------------------------------------------
window.switchToGrid = function (svc, svcOverride) {
    var grid = svc + 'Grid'; // this builds the name of the w2ui grid we want
    var x = getCurrentBusiness();
    var websvc = svc;
    if (typeof svcOverride === "string") {
        websvc = svcOverride;
    }
    var url = '/v1/' + websvc + '/' + x.value;
    w2ui[grid].url = url;
    w2ui[grid].last.sel_recid = null; // whenever switch grid, erase last selected record
    app.last.grid_sel_recid = -1;
    app.active_grid = grid; // mark active grid in app.active_grid
    w2ui.toplayout.content('main', w2ui[grid]);
    w2ui.toplayout.hide('right',true);
};

//---------------------------------------------------------------------------------
// opeinInNewTab simply opens a new tab in the browser and load the provided url
//---------------------------------------------------------------------------------
window.openInNewTab = function (url) {
    var win = window.open(url, '_blank');
    win.focus();
};

//-----------------------------------------------------------------------------
// GridMoneyFormat  - format comma-delimited money amount.
// @params  x   - value to be formatted
// @return  HTML string for the amount, suitable for render in w2ui grid cells
//-----------------------------------------------------------------------------
window.GridMoneyFormat = function (x) {
    var h = '';
    if (x !== 0) {
        h = '$ ' + number_format(x,2);
    }
    return h;
};

//-----------------------------------------------------------------------------
// getBIDfromBUD  - given the BUD return the associated BID. Returns
//                  undefined if BUD is not found
// @params  BUD   - the BUD for the business of interest
// @return  the BID (or `undefined` if not found)
//-----------------------------------------------------------------------------
window.getBIDfromBUD = function (BUD) {
    var BID;
    for (var i=0; i<app.BizMap.length; i++) {
        if (BUD == app.BizMap[i].BUD) {
            BID = app.BizMap[i].BID;
        }
    }
    return BID;
};

//-----------------------------------------------------------------------------
// getDepMeth     - searches BUD's Deposit Methods for id.  If found the
//                  then Deposit Method object is returned, otherwise an
//                  empty object is returned.
// @params  BUD   - the BUD for the business of interest
//          id - the Deposit Method id for which we want the name
// @return  the Deposit Method (or empty object if not found)
//-----------------------------------------------------------------------------
window.getDepMeth = function (BUD, id) {
    var dpm = {};
    if (typeof BUD === "undefined") {
        return dpm;
    }
    if (typeof app.depmeth[BUD].length == "undefined") { return; }
    for (var i = 0; i < app.depmeth[BUD].length; i++) {
        if (app.depmeth[BUD][i].id == id) {
            dpm = { id: id, text: app.depmeth[BUD][i].text };
            return dpm;
        }
    }
    return dpm;
};

//-----------------------------------------------------------------------------
// getDepository - searches BUD's Depositories for id.  If found the
//                 then Depository object is returned, otherwise an
//                 empty object is returned.
// @params  BUD  - the BUD for the business of interest
//          id   - the Depository id for which we want the name
// @return  the Depository (or empty object if not found)
//-----------------------------------------------------------------------------
window.getDepository = function (BUD, id) {
    var val = {};
    if (typeof BUD === "undefined") {
        return val;
    }
    if (typeof app.Depositories[BUD] !== "object") {
        return val;
    }
    for (var i = 0; i < app.Depositories[BUD].length; i++) {
        if (app.Depositories[BUD][i].id == id) {
            val = { id: id, text: app.Depositories[BUD][i].text };
            return val;
        }
    }
    return val;
};

//-----------------------------------------------------------------------------
// buildPaymentTypeSelectList - creates a list suitable for a dropdown menu
//                  with the payment types for the supplied BUD
// @params  BUD   - the BUD for the business of interest
// @return  the list of Payment Type Names (or empty list if BUD not found)
//-----------------------------------------------------------------------------
window.buildPaymentTypeSelectList = function (BUD) {

    var options = [{id:0, text: " -- Select Payment Type -- "}];
    if (typeof BUD == "undefined") {
        return options;
    }
    app.pmtTypes[BUD].forEach(function(pt) {
        options.push({ id: pt.PMTID, text: pt.Name });
    });
    return options;
};

//-----------------------------------------------------------------------------
// getCurrentBusiness - return the Business Unit currently slected in the
//                      main toolbar
// @params
// @return  the HTML elements of the currently selected business
//-----------------------------------------------------------------------------
window.getCurrentBusiness = function () {
    var x = document.getElementsByName("BusinessSelect");
    return x[0];
};

//-----------------------------------------------------------------------------
// getCurrentBID - return the BID for selected Business Unit currently in the
//                 main toolbar
// @params
// @return  - the BID of the currently selected business | "-1" if not exists
//-----------------------------------------------------------------------------
window.getCurrentBID = function () {
    var x = document.getElementsByName("BusinessSelect");
    if (x.length > 0) {
        return parseInt(x[0].value);
    }
    return -1;
};

//-----------------------------------------------------------------------------
// getBUDfromBID  - given the BID return the associated BUD. Returns
//                  an empty string if BID is not found
// @params  BUD   - the BUD for the business of interest
//          PMTID - the payment type id for which we want the name
// @return  the BUD (or empty string if not found)
//-----------------------------------------------------------------------------
window.getBUDfromBID = function (BID) {
    //
    var BUD = '';
    for (var i=0; i<app.BizMap.length; i++) {
        if (BID == app.BizMap[i].BID) {
            BUD = app.BizMap[i].BUD;
        }
    }
    return BUD;
};

//-----------------------------------------------------------------------------
// setToForm -  enable form sform in toplayout.  Also, set the forms url and
//              request data from the server
// @params
//   sform   = name of the form
//   url     = request URL for the form
//   [width] = optional, if specified it is the width of the form
//   doRequest =
//-----------------------------------------------------------------------------
window.setToForm = function (sform, url, width, doRequest) {
    // if not url defined then return
    var url_len=url.length > 0;
    if (!url_len) {
        return false;
    }

    // if form not found then return
    var f = w2ui[sform];
    if (!f) {
        return false;
    }

    // if current grid not found then return
    var g = w2ui[app.active_grid];
    if (!g) {
        return false;
    }

    // if doRequest is defined then take false as default one
    if (!doRequest) {
        doRequest = false;
    }

    if (url.length > 0 ) {
        f.url = url;
    }
    if (typeof f.tabs.name == "string") {
        f.tabs.click('tab1');
    }

    // mark this flag as is this new record
    app.new_form_rec = !doRequest;

    // as new content will be loaded for this form
    // mark form dirty flag as false
    app.form_is_dirty = false;

    var right_panel_content = w2ui.toplayout.get("right").content;

    // internal function
    var showForm = function() {
        // if the same content is there, then no need to render toplayout again
        if (f !== right_panel_content) {
            w2ui.toplayout.content('right', f);
            w2ui.toplayout.sizeTo('right', width);
            w2ui.toplayout.render();
        }
        else{
            // if same form is there then just refresh the form
            f.refresh();
        }
        // NOTE: remove any error tags bound to field from previous form
        $().w2tag();
        // SHOW the right panel now
        w2ui.toplayout.show('right', true);
    };

    if (doRequest) {
        f.request(function(event) {
            if (event.status === "success") {
                // only render the toplayout after server has sent down data
                // so that w2ui can bind values with field's html control,
                // otherwise it is unable to find html controls
                showForm();
                return true;
            }
            else {
                showForm();
                f.message("Could not get form data from server...!!");
                return false;
            }
        });
    } else {
        var sel_recid = parseInt(g.last.sel_recid);
        if (sel_recid > -1) {
            // if new record is being added then unselect {{the selected record}} from the grid
            g.unselect(g.last.sel_recid);
        }
        showForm();
        return true;
    }
};

//-----------------------------------------------------------------------------
// ridRentablePickerRender - renders a name during typedown.
// @params
//   item = an object with RentableName
// @return - true if the names match, false otherwise
//-----------------------------------------------------------------------------
window.ridRentablePickerRender = function (item) {
    w2ui.ridRentablePicker.record.RID = item.recid;
    return item.RentableName + '  (RID: ' + item.recid + ')';
};

//-----------------------------------------------------------------------------
// asmFormRentablePickerRender - renders a name during typedown.
// @params
//   item = Object with RentableName
// @return - true if the names match, false otherwise
//-----------------------------------------------------------------------------
window.asmFormRentablePickerRender = function (item) {
    w2ui.asmEpochForm.record.RID = item.recid;
    return item.RentableName + '  (RID: ' + item.recid + ')';
};

//-----------------------------------------------------------------------------
// ridRentableDropRender - renders a name during typedown.
// @params
//   item = an object with RentableName
// @return - the name to render
//-----------------------------------------------------------------------------
window.ridRentableDropRender = function (item) {
    return item.RentableName + '  (RID: ' + item.recid + ')';
};

//-----------------------------------------------------------------------------
// ridRentableCompare - Compare two items to see if they match
// @params
//   item = an object assumed to have a RentableName
// @return - true if the names match, false otherwise
//-----------------------------------------------------------------------------
window.ridRentableCompare = function (item, search) {
    var s = item.RentableName.toLowerCase();
    return s.includes(search.toLowerCase());
};

//-----------------------------------------------------------------------------
// tcidRAPayorPickerRender - renders a name during typedown.
// @params
//   item = an object assumed to have a FirstName and LastName
// @return - true if the names match, false otherwise
//-----------------------------------------------------------------------------
window.tcidRAPayorPickerRender = function (item) {

    var s="";
    if (item.IsCompany) {
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
};

//-----------------------------------------------------------------------------
// getFullName - returns a string with the full name based on the item supplied.
// @params
//   item = an object assumed to have a FirstName, MiddleName, and LastName
// @return - the full name concatenated together
//-----------------------------------------------------------------------------
window.getFullName = function (item) {

    var s = item.FirstName;
    if (item.MiddleName.length > 0) { s += ' ' + item.MiddleName; }
    if (item.LastName.length > 0 ) { s += ' ' + item.LastName; }
    return s;
};

//-----------------------------------------------------------------------------
// getTCIDName - returns an appropriate name for the supplied item. If
//          the item is a person, then the person's full name is returned.
//          If the item is a company, then the company name is returned.
// @params
//   item = an object assumed to have a FirstName, MiddleName, LastName,
//          IsCompany, and CompanyName.
// @return - the name to render
//-----------------------------------------------------------------------------
window.getTCIDName = function (item) {

    var s = (item.IsCompany) ? item.CompanyName : getFullName(item);

    if (item.TCID > 0) {
        s += ' (TCID: '+ String(item.TCID);
        if (typeof item.RAID == "number") {
            s += ', RAID: ' + item.RAID;
        }
        s += ')';
    }
    return s;
};

//-----------------------------------------------------------------------------
// tcidPickerCompare - Compare item to the search string. Verify that the
//          supplied search string can be found in item
// @params
//   item = an object assumed to have a FirstName and LastName
// @return - true if the search string is found, false otherwise
//-----------------------------------------------------------------------------
window.tcidPickerCompare = function (item, search) {

    var s = getTCIDName(item);
    s = s.toLowerCase();
    var srch = search.toLowerCase();
    var match = (s.indexOf(srch) >= 0);
    return match;
};

//-----------------------------------------------------------------------------
// tcidPickerDropRender - renders a name during typedown.
// @params
//   item = an object assumed to have a FirstName and LastName
// @return - the name to render
//-----------------------------------------------------------------------------
window.tcidPickerDropRender = function (item) {

    return getTCIDName(item);
};

//-----------------------------------------------------------------------------
// tcidReceiptPayorPickerRender - renders a name during typedown in the
//          receiptForm. It also sets the TCID for the record.
// @params
//   item = an object assumed to have a FirstName and LastName
// @return - true if the names match, false otherwise
//-----------------------------------------------------------------------------
window.tcidReceiptPayorPickerRender = function (item) {

    var s = getTCIDName(item);
    w2ui.receiptForm.record.TCID = item.TCID;
    w2ui.receiptForm.record.Payor = s;
    return s;
};

//-----------------------------------------------------------------------------
// tcidRUserPickerRender - renders a name during typedown.
// @params
//   item = an object assumed to have a FirstName and LastName
// @return - true if the names match, false otherwise
//-----------------------------------------------------------------------------
window.tcidRUserPickerRender = function (item) {

    var s;
    if (item.IsCompany) {
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
};


//-----------------------------------------------------------------------------
// plural - return the plural of the provided word.  Totally simplistic at
//          this point, it just adds an 's'.  It will need serious updates
//          going forward
// @params
//   s = the word to pluralize
// @return - the plural of word s
//-----------------------------------------------------------------------------
window.plural = function(s) {

    return s + 's';
};

//-----------------------------------------------------------------------------
// calcRarGridContractRent
//          - Sum the Contract Rent column of rarGrid and return the total.
//            used to set the control.
// @params
//          grid - The grid to work on
// @return  The total of the column
//-----------------------------------------------------------------------------
window.calcRarGridContractRent = function (grid) {

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
};

//-----------------------------------------------------------------------------
// getAccountsList - return the GLAccounts list with respect of BUD
// @params
// @return the list of accounts
//-----------------------------------------------------------------------------
window.getAccountsList = function (BID) {

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
};

//-----------------------------------------------------------------------------
// getPostAccounts - return the list of post accounts with respect of BUD
// @params
// @return the list of post accounts
//-----------------------------------------------------------------------------
window.getPostAccounts = function (BID) {

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
};

//-----------------------------------------------------------------------------
// getParentAccounts - return the list of Parent accounts with respect of BUD
// @params
//      - BID: current Business ID
//      - delLID: account id which needs to be substracted from the return list
// @return the list of parent accounts excluding delLID (current account ID from accountForm)
//-----------------------------------------------------------------------------
window.getParentAccounts = function (BID, delLID) {

    return jQuery.ajax({
        type: "GET",
        url: "/v1/parentaccounts/"+BID,
        dataType: "json",
    }).done(function(data) {
        if (data.status == "success") {
            var BUD = getBUDfromBID(BID);
            if (data.records) {
                var dft = {id: 0, text: ' -- No Parent LID -- '};
                var temp = [];
                data.records.forEach(function(item) {
                    if (item.id != delLID) {
                        temp.push(item);
                    }
                });
                // we don't need to exclude the default one from the list
                temp.unshift(dft);
                app.parent_accounts[BUD] = temp;
            } else{
                app.parent_accounts[BUD] = [];
            }
        }
    });
};


//-----------------------------------------------------------------------------
// int_to_bool converts int to bool.
// @params
//   i = integer to convert
// @return
//   boolean:  returns false if i == 0
//             otherwise it returns true
// This method needed to convert 1/0 value back to bool
// source: https://github.com/vitmalina/w2ui/blob/master/src/w2form.js#L368
//-----------------------------------------------------------------------------
window.int_to_bool = function (i){
    if (i>0) {
        return true;
    } else {
        return false;
    }
};


//-----------------------------------------------------------------------------
// getFormSubmitData - get form submit data
// @params, w2ui form record object
//          returnClone = true/false
// @return
// @description Helps to build form submit data, it modify record object so that each
// item in record has just a value instead of another object
//-----------------------------------------------------------------------------
window.getFormSubmitData = function (record, returnClone) {
    // check that it is typeof object or not
    if (typeof record !== "object") {
        return;
    }

    var cloneData = $.extend(true, {}, record);

    // iterate over each record
    for(var key in cloneData) {
        if (typeof cloneData[key] === "object" && cloneData[key] !== null && "id" in cloneData[key]) {
            cloneData[key] = cloneData[key].id;
        }
    }

    // if returnClone is not passed or false then
    // override cloned data into record
    if (!returnClone) {
        $.extend(record, cloneData);
        return record;
    }

    return cloneData;
};

//-----------------------------------------------------------------------------
// formRefreshCallBack -  callBack for form refresh event
// need to take several actions on refresh complete event
// @params
//   w2form   = w2form object
//   is_new     = true / false
//   id_name  = form's primary Id
//   form_header = header (title) of form
//-----------------------------------------------------------------------------
window.formRefreshCallBack = function (w2frm, primary_id, form_header, disable_header) {

    var record = w2frm.record,
        id = record[primary_id];

    // console.log(record);

    if (id === undefined) {
        console.log("given id_name '{0}' does not exist in form's '{1}' record".format(primary_id, w2frm.name));
        return false;
    }

    // mark active things of form
    app.active_form = w2frm.name;

    // keep active form original record
    app.active_form_original = $.extend(true, {}, record);

    var header = "";
    if (form_header) { // if form_header passed then
        // if new record then disable delete button
        // and format the equivalent header
        if (id > 0) {
            header = form_header.format(id);
            $(w2frm.box).find("button[name=delete]").removeClass("hidden");
            $(w2frm.box).find("button[name=reverse]").removeClass("hidden");
        } else {
            header = form_header.format("new");
            $(w2frm.box).find("button[name=delete]").addClass("hidden");
            $(w2frm.box).find("button[name=reverse]").addClass("hidden");
        }
    }

    if (typeof disable_header !== "undefined") {
        if (!disable_header) {
            w2frm.header = header;
        }
    }
};


window.number_format = function (number, decimals, dec_point, thousands_sep) {
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
};

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


//-----------------------------------------------------------------------------
// Save last form value entered as default for next form record
//
// @params
//   formFields         : array of form fields which needs to be reset, other fields are kept same as previous form record
//                        if array is empty then new form record will be same as previous form record
//                        if array is ['*'] then  new form record will be same as default(reset) form record
//   defaultFormRecord  : Object
//   previousFormRecord : Object
// @returns
//   defaultFormRecord  : Object
//-----------------------------------------------------------------------------
window.setDefaultFormFieldAsPreviousRecord = function (formFields, defaultFormRecord, previousFormRecord) {
    if (formFields.length === 0) {
        return previousFormRecord;
    }
    if (formFields[0] === '*') {
        return defaultFormRecord;
    }
    for ( var i = 0; i < formFields.length; i++) {
        previousFormRecord[formFields[i]] = defaultFormRecord[formFields[i]];
    }
    return previousFormRecord;
};

//-------------------------------------------------------------------------------
// Download the CSV report for given report name, date range
//
// @params
//   rptname    : report name to be downloaded
//   dtStart    : Start Date
//   dtStop     : Stop Date
//   returnURL  : it true then returns the url otherwise
//                downloads the report from built url in separate window
//   id         : id for the report to detail
//-------------------------------------------------------------------------------
window.exportItemReportCSV = function (rptname,id,dtStart,dtStop,returnURL) {
    var BID = getCurrentBID();
    var BUD = getBUDfromBID(BID);
    var bizEDIEnabled = EDIEnabledForBUD(BUD);
    var edi = bizEDIEnabled ? 1 : 0;

    var url = '/v1/report/' + String(BID) + '/' + id + '?r=' + rptname + '&edi=' + String(edi);
    if (returnURL) {
        return finishReportCSV(url,rptname, dtStart, dtStop, returnURL);
    }
    finishReportCSV(url,rptname, dtStart, dtStop, returnURL);
};

//-------------------------------------------------------------------------------
// Download the CSV report for given report name, date range
//
// @params
//   rptname   : report name to be downloaded
//   dtStart   : Start Date
//   dtStop    : Stop Date
//   returnURL : it true then returns the url otherwise
//               downloads the report from built url in separate window
//-------------------------------------------------------------------------------
window.exportReportCSV = function (rptname, dtStart, dtStop, returnURL){
    if (rptname === '') {
        return;
    }

    var BID = getCurrentBID();
    var BUD = getBUDfromBID(BID);
    var bizEDIEnabled = EDIEnabledForBUD(BUD);
    var edi = bizEDIEnabled ? 1 : 0;

    var url = '/v1/report/' + String(BID) + '?r=' + rptname + '&edi=' + String(edi);

    if (returnURL) { // if retrunURL is set then we need to return it
        return finishReportCSV(url,rptname, dtStart, dtStop, returnURL);
    }
    finishReportCSV(url,rptname, dtStart, dtStop, returnURL);
};

window.finishReportCSV = function (url,rptname, dtStart, dtStop, returnURL) {
    // if both dates are available then only append dtstart and dtstop in query params
    if (dtStart && dtStop) {
        url += '&dtstart=' + dtStart; // StartDate
        url += '&dtstop=' + dtStop; // stopDate
    }

    // now append the report output format
    url += '&rof=' + app.rof.csv;
    console.log('url = ' + url);

    // open separate window if returnURL is not true
    if (returnURL) {
        return url;
    } else {
        downloadMediaFromURL(url);
    }
};

//-------------------------------------------------------------------------------
// Pops up dialog to get custom width and height from user's input
//-------------------------------------------------------------------------------
window.popupPDFCustomDimensions = function () {
    w2popup.open({
        title     : 'PDF custom width and height',
        body      : '<div class="w2ui-centered">' +
            '<div class="w2ui-field"><label>Page Width (inch): </label><div><input type="text" name="custom_pdf_width" class="w2ui-input" value="'+app.pdfPageWidth+'" /></div></div>' +
            '<div class="w2ui-field"><label>Page Height (inch): </label><div><input type="text" name="custom_pdf_height"  class="w2ui-input" value="'+app.pdfPageHeight+'" /></div></div>' +
            '</div>',
        buttons   : '<button class="w2ui-btn" onclick="w2popup.close();">Close</button> '+
                    '<button class="w2ui-btn" onclick="saveCustomDims();" >Save</button>',
        width     : 500,
        height    : 200,
        overflow  : 'hidden',
        color     : '#333',
        speed     : '0.3',
        opacity   : '0.5',
        modal     : true,
        showClose : true,
    });
};

//-------------------------------------------------------------------------------
// Remembers custom dimensions set up by user locally in app variable
//-------------------------------------------------------------------------------
window.saveCustomDims = function () {
    var width = parseFloat($("input[name='custom_pdf_width']").val());
    if (!isNaN(width)) {
        app.pdfPageWidth = width;
    }
    var height = parseFloat($("input[name='custom_pdf_height']").val());
    if (!isNaN(height)) {
        app.pdfPageHeight = height;
    }
    w2popup.close();
};

//-------------------------------------------------------------------------------
// Download the PDF report for given id-focused report, date range
//
// @params
//   rptname            : report name to be downloaded
//   id                 : id of item on which report should focus
//   dtStart            : Start Date
//   dtStop             : Stop Date
//   returnURL          : it true then returns the url otherwise
//                        downloads the report from built url in separate window
//-------------------------------------------------------------------------------
window.exportItemReportPDF = function (rptname,id, dtStart, dtStop, returnURL){
    if (rptname === '') {
        return;
    }
    var BID = getCurrentBID();
    var BUD = getBUDfromBID(BID);
    var bizEDIEnabled = EDIEnabledForBUD(BUD);
    var edi = bizEDIEnabled ? 1 : 0;

    var url = '/v1/report/' + String(BID) + '/' + id + '?r=' + rptname + '&edi=' + String(edi);
    if (returnURL) {
        return finishReportPDF(url,rptname, dtStart, dtStop, returnURL);
    }
    finishReportPDF(url,rptname, dtStart, dtStop, returnURL);
};

//-------------------------------------------------------------------------------
// Download the PDF report for given report name, date range
//
// @params
//   rptname            : report name to be downloaded
//   dtStart            : Start Date
//   dtStop             : Stop Date
//   returnURL          : it true then returns the url otherwise
//                        downloads the report from built url in separate window
//-------------------------------------------------------------------------------
window.exportReportPDF = function (rptname, dtStart, dtStop, returnURL){
    if (rptname === '') {
        return;
    }

    var BID = getCurrentBID();
    var BUD = getBUDfromBID(BID);
    var bizEDIEnabled = EDIEnabledForBUD(BUD);
    var edi = bizEDIEnabled ? 1 : 0;

    var url = '/v1/report/' + String(BID) + '?r=' + rptname + '&edi=' + String(edi);
    if (returnURL) { // if retrunURL is set then we need to return it
        return finishReportPDF(url,rptname, dtStart, dtStop, returnURL);
    }
    finishReportPDF(url,rptname, dtStart, dtStop, returnURL);
};

window.finishReportPDF = function (url,rptname, dtStart, dtStop, returnURL) {
    // if both dates are available then only append dtstart and dtstop in query params
    if (dtStart && dtStop) {
        url += '&dtstart=' + dtStart; // StartDate
        url += '&dtstop=' + dtStop; // stopDate
    }

    // now append the report output format
    url += '&rof=' + app.rof.pdf;

    // need to pass page width and height
    url += '&pw=' + app.pdfPageWidth + "&ph=" + app.pdfPageHeight;
    console.log('url = ' + url);

    // open separate window if returnURL is not true
    if (returnURL) {
        return url;
    } else {
        downloadMediaFromURL(url);
    }
};

//-------------------------------------------------------------------------------
// Download the media using provided URL
//
// @params
//   url            : the url to download the media
//-------------------------------------------------------------------------------
window.downloadMediaFromURL = function (url) {
    var idown = $('#down_iframe');
    if (idown.length > 0) {
        idown.attr('src', url);
    } else {
        idown = $('<iframe>', { id: 'down_iframe', src: url }).hide().appendTo('body');
    }

    // reset the url after download after sometime
    setTimeout(function() {
        idown.attr('src', '');
    }, 1000);
};

//-------------------------------------------------------------------------------
// returns true/false tells whether EDI mode enabled for business BUD
//
// @params
//   BUD            : business designation
//-------------------------------------------------------------------------------
window.EDIEnabledForBUD = function(BUD) {
    if (app.bizFLAGS && app.bizFLAGS[BUD]) {
        return (app.bizFLAGS[BUD]&1) > 0;
    }
    return false;
};

//---------------------------------------------------------------------------------
// prepareW2UIStuff - it will prepare lists, items, other things which are
//                    required by w2ui objects. It will feed those in "w2ui"
//                    of app variable
//
// @params  app  - app variable of application
//---------------------------------------------------------------------------------
window.prepareW2UIStuff = function prepareW2UIStuff(app) {

    // cycle frequencies
    app.w2ui.listItems.cycleFreq = [];
    if (app.cycleFreq) {
        app.cycleFreq.forEach(function(freq, index) {
            app.w2ui.listItems.cycleFreq.push({ id: index, text: freq });
        });
    }
};

//-----------------------------------------------------------------------------
// reassignGridRecids -  will reassign the grid record's recid
//                       in case of record deleted within the grid
// @params
//   gridName = w2ui grid component name
//-----------------------------------------------------------------------------
window.reassignGridRecids = function(gridName) {
    if (gridName in w2ui) {
        var grid = w2ui[gridName];
        for (var j = 0; j < grid.records.length; j++) {
            grid.records[j].recid = j + 1;
        }
        // need to refresh the grid as it will assign new recid in DOM tr's attribute "recid"
        grid.refresh();
    }
};