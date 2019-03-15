/*global
    w2ui, app, $, console, form_dirty_alert, formRefreshCallBack, formRecDiffer,
    getFormSubmitData, w2confirm, delete_confirm_options, getBUDfromBID, getCurrentBusiness,
    addDateNavToToolbar, setRTLayout, getRTInitRecord, getRentASMARList, showForm,
    RTEdits,
*/
/*jshint esversion: 6 */

"use strict";

// saveRentableMarketRate - creates a list of MarketRate entries that have
// been changed, then calls the webservice to save them.
//---------------------------------------------------------------------------
window.saveRentableMarketRate = function(BID,RTID) {
    var reclist = Array.from(new Set(RTEdits.MarketRateChgList));

    if (reclist.length == 0) {
        return Promise.resolve('{"status": "success"}');
    }

    var chgrec = [];
    for (var i = 0; i < reclist.length; i++) {
        var nrec =  w2ui.rmrGrid.get(reclist[i]);
        chgrec.push(nrec);
    }

    var params = {
        cmd: "save",
        selected: [],
        limit: 0,
        offset: 0,
        changes: chgrec,
        RTID: RTID
    };

    var dat = JSON.stringify(params);
    var url = '/v1/rmr/' + BID + '/' + RTID;

    return $.post(url, dat, null, "json")
    .done(function(data) {
        if (data.status === "success") {
            //------------------------------------------------------------------
            // Now that the save is complete, we can add the URL back to the
            // the grid so it can call the server to get updated rows. The
            // onLoad handler will reset the url to '' after the load completes
            // so that changes are done locally to gthe grid until the
            // rtForm save button is clicked.
            //------------------------------------------------------------------
            RTEdits.MarketRateChgList = []; // reset the change list now, because we've saved them
            w2ui.rmrGrid.url = url;
            w2ui.toplayout.hide('right', true);
        } else {
            w2ui.rentablesGrid.error('saveRentableMarketRate: ' + data.message);
        }
    })
    .fail(function(data){
        console.log("Save RentableMarketRate failed.");
    });
};
