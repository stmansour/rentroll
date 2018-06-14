/*global
    $, w2ui, console, app, getTCIDName, w2popup, getCurrentBusiness,
    asmFormRASelect, expFormRASelect, rentalAgrPickerRender, rentalAgrPickerDropRender, rentalAgrPickerCompare,
    rafinder
*/
"use strict";

//------------------------------------------------------------------------
//          rental Agreement Finder
//------------------------------------------------------------------------

window.rafinder = {
    cb: null,
};

//-----------------------------------------------------------------------------
// buildRAPicker - rentalAgrPicker is a form to help the user find the Rental
//          Agreement they're looking for. It uses typedown on the payors name
//          to determine what Rental Agreements the user is responsible for
//          and lists them all in real-time so the user can pick the one they
//          want.
// @params
// @return
//-----------------------------------------------------------------------------
window.buildRAPicker = function (){
    $().w2form({
        name: 'rentalAgrPicker',
        style: 'border: 0px; background-color: transparent;',
        formURL: '/webclient/html/rentalagrfinder.html',
        focus  : 0,
        fields: [
            { field: 'TCID', type: 'int', required: true },
            // INDEX 1
            { field: 'PayorName', required: true,
                type: 'enum',
                options: {
                    url:            '/v1/rentalagrtd/' + app.RentalAgrPicker.BID,
                    // max:     1,
                    items: [],
                    openOnFocus:    true,
                    maxDropWidth:   350,
                    maxDropHeight:  350,
                    renderItem:     rentalAgrPickerRender,
                    renderDrop:     rentalAgrPickerDropRender,
                    compare:        rentalAgrPickerCompare,
                    onNew: function (event) {
                        console.log('++ New Item: Do not forget to submit it to the server too', event);
                        //$.extend(event.item, { FirstName: '', LastName : event.item.text });
                    }
                },
            },
            // INDEX 2
            { field: 'RentableName', type: 'list',      required: true, options: { items: [] } },
            { field: 'RAID',         type: 'int',       required: true  },
            { field: 'FirstName',    type: 'text',      required: false },
            { field: 'LastName',     type: 'text',      required: false },
            { field: 'CompanyName',  type: 'text',      required: false },
            { field: 'IsCompany',    type: 'checkbox',  required: false },
        ],
        onRefresh: function(/*event*/) {
            w2ui.rentalAgrPicker.fields[1].options.url = '/v1/rentalagrtd/' + app.RentalAgrPicker.BID;
            w2ui.rentalAgrPicker.fields[2].options.items = app.RentalAgrPicker.RARentablesNames;
            if (app.RentalAgrPicker.RARentablesNames.length == 1) {
                w2ui.rentalAgrPicker.record.RentableName = app.RentalAgrPicker.RARentablesNames[0];
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
};

// popupRentalAgrPicker comes up when the user clicks on the Find... button
// while creating an assessment. It is used to locate a rental agreement by payor.
// @PARAMS
//    s - caller name
//----------------------------------------------------------------------------------
window.popupRentalAgrPicker = function (s) {
    rafinder.caller = s;
    var x = getCurrentBusiness();
    app.RentalAgrPicker = {BID: x.value, RAID: 0, TCID: 0, RID: 0, FirstName: '', LastName: '', CompanyName: '', IsCompany: false, RAR: [], RARentablesNames: []};
    app.RentalAgrPicker.RARentablesNames = [{id: 0, text:" "}];
    w2ui.rentalAgrPicker.fields[2].options.items = app.RentalAgrPicker.RARentablesNames;
    w2ui.rentalAgrPicker.record.TCID = -1;
    w2ui.rentalAgrPicker.record.RAID = -1;
    w2ui.rentalAgrPicker.record.PayorName = '';
    w2ui.rentalAgrPicker.record.IsCompany = false;
    w2ui.rentalAgrPicker.record.CompanyName = '';
    w2ui.rentalAgrPicker.record.FirstName = '';
    w2ui.rentalAgrPicker.record.LastName = '';
    w2ui.rentalAgrPicker.refresh();

    $().w2popup('open', {
        title   : 'Find Rental Agreement',
        body    : '<div id="form" style="width: 100%; height: 100%;"></div>',
        style   : 'padding: 15px 0px 0px 0px',
        width   : 400,
        height  : 250,
        showMax : true,
        onToggle: function (event) {
            $(w2ui.rentalAgrPicker.box).hide();
            event.onComplete = function () {
                $(w2ui.rentalAgrPicker.box).show();
                w2ui.rentalAgrPicker.resize();
            };
        },
        onOpen: function (event) {
            event.onComplete = function () {
                // specifying an onOpen handler instead would be equivalent to specifying
                // an onBeforeOpen handler, which would make this code execute too
                // early and hence not deliver.
                $('#w2ui-popup #form').w2render('rentalAgrPicker');
            };
        }
    });
};


//-----------------------------------------------------------------------------
// rentalAgrPickerCompare - Compare item to the search string. Verify that the
//          supplied search string can be found in item
// @params
//   item = an object assumed to have a FirstName and LastName
// @return - true if the search string is found, false otherwise
//-----------------------------------------------------------------------------
window.rentalAgrPickerCompare = function (item, search) {
    var s = getTCIDName(item);
    s = s.toLowerCase();
    var srch = search.toLowerCase();
    var match = (s.indexOf(srch) >= 0);
    return match;
};

//-----------------------------------------------------------------------------
// rentalAgrPickerDropRender - renders a name during typedown.
// @params
//   item = an object assumed to have a FirstName and LastName
// @return - the name to render
//-----------------------------------------------------------------------------
window.rentalAgrPickerDropRender = function (item) {
    return getTCIDName(item);
};

// //-----------------------------------------------------------------------------
// // rentalAgrPickerRender - renders a name during typedown in the
// //          rentalAgrPicker. It also sets the TCID for the record.
// // @params
// //   item = an object assumed to have a FirstName and LastName
// // @return - true if the names match, false otherwise
// //-----------------------------------------------------------------------------
// function rentalAgrPickerRender(item) {
//     var s = getTCIDName(item);
//     w2ui.rentalAgrPicker.record.TCID = item.TCID;
//     w2ui.rentalAgrPicker.record.Payor = s;
//     w2ui.rentalAgrPicker.record.RAID = item.RAID;
//     return s;
// }

//-----------------------------------------------------------------------------
// rentalAgrPickerRender - renders a name during typedown.
// @params
//   item = an object assumed to have a FirstName and LastName
// @return - true if the names match, false otherwise
//-----------------------------------------------------------------------------
window.rentalAgrPickerRender = function (item) {
    var s;
    if (item.IsCompany) {
        s = item.CompanyName;
    } else {
        s = item.FirstName + ' ' + item.LastName;
    }

    w2ui.rentalAgrPicker.record = {
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
    var url = '/v1/rar/' + app.RentalAgrPicker.BID + '/' + item.RAID;
    $.get(url, null, null, "json")
    .done(function(data) {
        app.RentalAgrPicker.RAR = data;
        app.RentalAgrPicker.RARentablesNames = [];
        if (app.RentalAgrPicker.RAR.records) {
            for (var i = 0; i < app.RentalAgrPicker.RAR.records.length; i++) {
                app.RentalAgrPicker.RARentablesNames.push(
                    { id: app.RentalAgrPicker.RAR.records[i].RID, text: app.RentalAgrPicker.RAR.records[i].RentableName} );
            }
        } else {
            app.RentalAgrPicker.RARentablesNames.push({ id: 0, text: ''} );
        }
        console.log('calling rentalAgrPicker.refresh(), app.RentalAgrPicker.RARentablesNames.length = ' + app.RentalAgrPicker.RARentablesNames.length );
        w2ui.rentalAgrPicker.refresh();
    });
    return s;
};


