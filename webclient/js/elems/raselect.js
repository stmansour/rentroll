/*global
    $, w2ui, console, app, getTCIDName, w2popup, getCurrentBusiness,
    asmFormRASelect, expFormRASelect,
*/
"use strict";

//------------------------------------------------------------------------
//          rental Agreement Finder
//------------------------------------------------------------------------

var rafinder = {
    cb: null,
};

//-----------------------------------------------------------------------------
// buildRASelect - rentalAgrFinder is a form to help the user find the Rental
//          Agreement they're looking for. It uses typedown on the payors name
//          to determine what Rental Agreements the user is responsible for
//          and lists them all in real-time so the user can pick the one they
//          want.
// @params
// @return
//-----------------------------------------------------------------------------
function buildRASelect(){
    $().w2form({
        name: 'rentalAgrFinder',
        style: 'border: 0px; background-color: transparent;',
        formURL: '/webclient/html/rentalagrfinder.html',
        focus  : 0,
        fields: [
            { field: 'TCID', type: 'int', required: true },
            // INDEX 1
            { field: 'PayorName', required: true,
                type: 'enum',
                options: {
                    url:            '/v1/rentalagrtd/' + app.RentalAgrFinder.BID,
                    // max:     1,
                    items: [],
                    openOnFocus:    true,
                    maxDropHeight:  350,
                    renderItem:     rentalAgrFinderRender,
                    renderDrop:     rentalAgrFinderDropRender,
                    compare:        rentalAgrFinderCompare,
                    onNew: function (event) {
                        console.log('++ New Item: Do not forget to submit it to the server too', event);
                        //$.extend(event.item, { FirstName: '', LastName : event.item.text });
                    }
                },
            },
            // INDEX 2
            { field: 'RentableName', type: 'list', required: true, options: { items: [] } },
            { field: 'RAID',         type: 'int',  required: true  },
            { field: 'FirstName',    type: 'text', required: false },
            { field: 'LastName',     type: 'text', required: false },
            { field: 'CompanyName',  type: 'text', required: false },
            { field: 'IsCompany',    type: 'int',  required: false },
        ],
        onRefresh: function(/*event*/) {
            w2ui.rentalAgrFinder.fields[1].options.url = '/v1/rentalagrtd/' + app.RentalAgrFinder.BID;
            w2ui.rentalAgrFinder.fields[2].options.items = app.RentalAgrFinder.RARentablesNames;
            if (app.RentalAgrFinder.RARentablesNames.length == 1) {
                w2ui.rentalAgrFinder.record.RentableName = app.RentalAgrFinder.RARentablesNames[0];
            }
        },
        actions: {
            save: function () {
                if (typeof rafinder.cb == "function" ) {
                    rafinder.cb();
                }
                w2popup.close();
            },
        },
    });
}

// popupRentalAgrPicker comes up when the user clicks on the Find... button
// while creating an assessment. It is used to locate a rental agreement by payor.
// @PARAMS
//    s - caller name
//----------------------------------------------------------------------------------
function popupRentalAgrPicker(s) {
    rafinder.caller = s;
    var x = getCurrentBusiness();
    app.RentalAgrFinder = {BID: x.value, RAID: 0, TCID: 0, RID: 0, FirstName: '', LastName: '', CompanyName: '', IsCompany: false, RAR: [], RARentablesNames: []};
    app.RentalAgrFinder.RARentablesNames = [{id: 0, text:" "}];
    w2ui.rentalAgrFinder.fields[2].options.items = app.RentalAgrFinder.RARentablesNames;
    w2ui.rentalAgrFinder.record.TCID = -1;
    w2ui.rentalAgrFinder.record.RAID = -1;
    w2ui.rentalAgrFinder.record.PayorName = '';
    w2ui.rentalAgrFinder.record.IsCompany = -1;
    w2ui.rentalAgrFinder.record.CompanyName = '';
    w2ui.rentalAgrFinder.record.FirstName = '';
    w2ui.rentalAgrFinder.record.LastName = '';
    w2ui.rentalAgrFinder.refresh();

    $().w2popup('open', {
        title   : 'Find Rental Agreement',
        body    : '<div id="form" style="width: 100%; height: 100%;"></div>',
        style   : 'padding: 15px 0px 0px 0px',
        width   : 400,
        height  : 250,
        showMax : true,
        onToggle: function (event) {
            $(w2ui.rentalAgrFinder.box).hide();
            event.onComplete = function () {
                $(w2ui.rentalAgrFinder.box).show();
                w2ui.rentalAgrFinder.resize();
            };
        },
        onOpen: function (event) {
            event.onComplete = function () {
                // specifying an onOpen handler instead would be equivalent to specifying
                // an onBeforeOpen handler, which would make this code execute too
                // early and hence not deliver.
                $('#w2ui-popup #form').w2render('rentalAgrFinder');
            };
        }
    });
}


//-----------------------------------------------------------------------------
// rentalAgrFinderCompare - Compare item to the search string. Verify that the
//          supplied search string can be found in item
// @params
//   item = an object assumed to have a FirstName and LastName
// @return - true if the search string is found, false otherwise
//-----------------------------------------------------------------------------
function rentalAgrFinderCompare(item, search) {
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

    // Try to getget the rentables associated with item.RAID.  There may not
    // be any rentables, which means they could be taking an application fee
    // from a potential renter...
    //------------------------------------------------------------------------
    var url = '/v1/rar/' + app.RentalAgrFinder.BID + '/' + item.RAID;
    $.get(url,function(data /*,status*/) {
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


