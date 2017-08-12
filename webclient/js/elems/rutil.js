/*global
    app, w2ui, $, form_dirty_alert,

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
// ChangeBusiness updates the UI to the newly selected business.
// This routine is indeed used, in spite of what JSHint thinks. It
// is embedded in a string defining the OnClick handler for a button
// in the main toolbar.
// ---------------------------------------------------------------------------------
function ChangeBusiness() {
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
        app.last.BUD = bizName;
        app.last.BID = bizVal;
        s.collapse('reports');
        return true;
    };

    // warn user if active form has been changed
    form_dirty_alert(yes_callBack, no_callBack);
}

//---------------------------------------------------------------------------------
// switchToGrid - changes the main view of the program to a grid with
//                variable name svc + 'Grid'
//
// @params  svc = prefix of grid name
//          svcOverride = name of webservice to call if the name does not
//                match the name of the svc
//
//---------------------------------------------------------------------------------
function switchToGrid(svc,svcOverride) {
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
}

//---------------------------------------------------------------------------------
// opeinInNewTab simply opens a new tab in the browser and load the provided url
//---------------------------------------------------------------------------------
function openInNewTab(url) {
    var win = window.open(url, '_blank');
    win.focus();
}


//-----------------------------------------------------------------------------
// getBUDfromBID  - given the BID return the associated BUD. Returns
//                  an empty string if BID is not found
// @params  BUD   - the BUD for the business of interest
//          PMTID - the payment type id for which we want the name
// @return  the BUD (or empty string if not found)
//-----------------------------------------------------------------------------
function getBUDfromBID(BID) {
    //
    var BUD = '';
    for (var i=0; i<app.BizMap.length; i++) {
        if (BID == app.BizMap[i].BID) {
            BUD = app.BizMap[i].BUD;
        }
    }
    return BUD;
}

//-----------------------------------------------------------------------------
// GridMoneyFormat  - format comma-delimited money amount.
// @params  x   - value to be formatted
// @return  HTML string for the amount, suitable for render in w2ui grid cells
//-----------------------------------------------------------------------------
function GridMoneyFormat(x) {
    var h = '';
    if (x !== 0) {
        h = '$ ' + number_format(x,2);
    }
    return h;
}

//-----------------------------------------------------------------------------
// getBIDfromBUD  - given the BUD return the associated BID. Returns
//                  undefined if BUD is not found
// @params  BUD   - the BUD for the business of interest
// @return  the BID (or `undefined` if not found)
//-----------------------------------------------------------------------------
function getBIDfromBUD(BUD) {

    var BID;
    for (var i=0; i<app.BizMap.length; i++) {
        if (BUD == app.BizMap[i].BUD) {
            BID = app.BizMap[i].BID;
        }
    }
    return BID;
}

//-----------------------------------------------------------------------------
// getPaymentType - searches BUD's Payment Types for PMTID.  If found the
//                  then payment type object is returned, else an empty object is returned.
// @params  BUD   - the BUD for the business of interest
//          PMTID - the payment type id for which we want the name
// @return  the Payment Type (or empty object if not found)
//-----------------------------------------------------------------------------
function getPaymentType(BUD, reqPMTID) {
    var pmt = {};
    if (typeof BUD === "undefined") {
        return pmt;
    }
    app.pmtTypes[BUD].forEach(function(item) {
        if (item.PMTID == reqPMTID) {
            pmt = { id: item.PMTID, text: item.Name };
            return pmt;
        }
    });
    return pmt;
}

//-----------------------------------------------------------------------------
// getDepMeth     - searches BUD's Deposit Methods for id.  If found the
//                  then Deposit Method object is returned, otherwise an
//                  empty object is returned.
// @params  BUD   - the BUD for the business of interest
//          id - the Deposit Method id for which we want the name
// @return  the Deposit Method (or empty object if not found)
//-----------------------------------------------------------------------------
function getDepMeth(BUD, id) {
    var dpm = {};
    if (typeof BUD === "undefined") {
        return dpm;
    }
    for (var i = 0; i < app.depmeth[BUD].length; i++) {
        if (app.depmeth[BUD][i].id == id) {
            dpm = { id: id, text: app.depmeth[BUD][i].text };
            return dpm;
        }
    }
    return dpm;
}

//-----------------------------------------------------------------------------
// getDepository - searches BUD's Depositories for id.  If found the
//                 then Depository object is returned, otherwise an
//                 empty object is returned.
// @params  BUD  - the BUD for the business of interest
//          id   - the Depository id for which we want the name
// @return  the Depository (or empty object if not found)
//-----------------------------------------------------------------------------
function getDepository(BUD, id) {
    var val = {};
    if (typeof BUD === "undefined") {
        return val;
    }
    for (var i = 0; i < app.Depositories[BUD].length; i++) {
        if (app.Depositories[BUD][i].id == id) {
            val = { id: id, text: app.Depositories[BUD][i].text };
            return val;
        }
    }
    return val;
}

//-----------------------------------------------------------------------------
// buildPaymentTypeSelectList - creates a list suitable for a dropdown menu
//                  with the payment types for the supplied BUD
// @params  BUD   - the BUD for the business of interest
// @return  the list of Payment Type Names (or empty list if BUD not found)
//-----------------------------------------------------------------------------
function buildPaymentTypeSelectList(BUD) {

    var options = [{id:0, text: " -- Select Payment Type -- "}];
    if (typeof BUD == "undefined") {
        return options;
    }
    app.pmtTypes[BUD].forEach(function(pt) {
        options.push({ id: pt.PMTID, text: pt.Name });
    });
    return options;
}

//-----------------------------------------------------------------------------
// getCurrentBusiness - return the Business Unit currently slected in the
//                      main toolbar
// @params
// @return  the BUD of the currently selected business
//-----------------------------------------------------------------------------
function getCurrentBusiness() {

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
//   doRequest = 
//-----------------------------------------------------------------------------
function setToForm(sform, url, width, doRequest) {
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

    f.url = url;
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

            /*// HACK: set the height of right panel of toplayout box div and form's box div
            // this is how w2ui set the content inside box of toplayout panel, and form's main('div.w2ui-form-box')
            var h = w2ui.toplayout.get("right").height;
            $(w2ui.toplayout.get("right").content.box).height(h);
            $(f.box).find("div.w2ui-form-box").height(h);*/
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
}

//-----------------------------------------------------------------------------
// ridRentablePickerRender - renders a name during typedown.
// @params
//   item = an object with RentableName
// @return - true if the names match, false otherwise
//-----------------------------------------------------------------------------
function ridRentablePickerRender(item) {

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


//-----------------------------------------------------------------------------
// plural - return the plural of the provided word.  Totally simplistic at
//          this point, it just adds an 's'.  It will need serious updates
//          going forward
// @params
//   s = the word to pluralize
// @return - the plural of word s
//-----------------------------------------------------------------------------
function plural(s) {

    return s + 's';
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
// getRentableTypes - return the RentableTypes list with respect of BUD
// @params
//      - BUD: current business designation
// @return  the Rentable Types List
//-----------------------------------------------------------------------------
function getRentableTypes(BUD) {

    return jQuery.ajax({
        type: "GET",
        url: "/v1/rtlist/"+BUD,
        dataType: "json",
    }).done(function(data) {
        if (data.status == "success") {
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
//      - BID: current Business ID
//      - delLID: account id which needs to be substracted from the return list
// @return the list of parent accounts excluding delLID (current account ID from accountForm)
//-----------------------------------------------------------------------------
function getParentAccounts(BID, delLID) {

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
}

//-----------------------------------------------------------------------------
// unallocAmountRemaining - based on the amounts allocated to receipts in the
// unpaid receipts list, compute the amount of funds remaining to be allocated
// and display it.
// @params
// @return
//-----------------------------------------------------------------------------
function unallocAmountRemaining() {

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

    if (i>0) {
        return true;
    } else {
        return false;
    }
}


//-----------------------------------------------------------------------------
// getPayorFund - get payor fund
// @params
// @return  the jquery promise
//-----------------------------------------------------------------------------
function getPayorFund(BID, TCID) {

    return jQuery.ajax({
        type: "GET",
        url: '/v1/payorfund/'+BID+'/'+TCID,
        dataType: "json",
    });
}

// Auto Allocate amount for each unpaid assessment
jQuery(document).on('click', '#auto_allocate_btn', function(/*event*/) {

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
// formRefreshCallBack -  callBack for form refresh event
// need to take several actions on refresh complete event
// @params
//   w2form   = w2form object
//   is_new     = true / false
//   id_name  = form's primary Id
//   form_header = header (title) of form
//-----------------------------------------------------------------------------
function formRefreshCallBack(w2frm, id_name, form_header) {

    var fname = w2frm.name,
        record = w2frm.record,
        id = record[id_name],
        header = form_header;

    if (id === undefined) {
        console.log("given id_name does not exist in form's record");
        return false;
    }

    // mark active things of form
    app.active_form = fname;
    // keep active form original record
    app.active_form_original = $.extend(true, {}, record);
    // if new record then disable delete button
    // and format the equivalent header
    if (id === 0) {
        w2frm.header = header.format("new");
        $("#"+fname).find("button[name=delete]").addClass("hidden");
        $("#"+fname).find("button[name=reverse]").addClass("hidden");
    }
    else {
        w2frm.header = header.format(id);
        $("#"+fname).find("button[name=delete]").removeClass("hidden");
        $("#"+fname).find("button[name=reverse]").removeClass("hidden");
    }

    /*// ============================
    // HACK: set the height of right panel of toplayout box div and form's box div
    // this is how w2ui set the content inside box of toplayout panel, and form's main('div.w2ui-form-box')
    // ============================
    // ALREADY HANDLED IN "setToForm"
    // ============================
    var h = w2ui.toplayout.get("right").height;
    $(w2ui.toplayout.get("right").content.box).height(h);
    $(w2frm.box).find("div.w2ui-form-box").height(h);*/
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
