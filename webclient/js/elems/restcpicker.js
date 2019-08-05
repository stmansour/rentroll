"use strict";
/*global
    w2ui, $, app, console, w2utils,
    ResTCDropRender,
*/

var ResSelectedTC = {};

//-----------------------------------------------------------------------------
// ResTCCompare - Compare item to the search string. Verify that the
//          supplied search string can be found in item.  If it's already
//          listed we don't want to list it again.
// @params
//   item = an object assumed to have a Name and TLID field
// @return - true if the search string is found, false otherwise
//-----------------------------------------------------------------------------
window.ResTCCompare = function (item, search) {
    var s = ResTCDropRender(item);
    s = s.toLowerCase();
    var srch = search.toLowerCase();
    var match = (s.indexOf(srch) >= 0);
    return match;
};

//-----------------------------------------------------------------------------
// ResTCDropRender - renders a name during typedown.
// @params
//   item = an object having these fields:
//    TCID
//    FirstName
//    MiddleName
//    LastName
//    CompanyName
//    IsCompany
//    PrimaryEmail
//    SecondaryEmail
//    WorkPhone
//    CellPhone
//    Address
//    Address2
//    City
//    State
//    PostalCode
//
// @return - the name to render
//-----------------------------------------------------------------------------
window.ResTCDropRender = function (item) {
    var s = item.FirstName;
    if (item.IsCompany) {
        s = '<i class="fas fa-building"></i> ' + item.CompanyName;
    }
    if (item.MiddleName.length > 0) {
        s += ' ' + item.MiddleName;
    }
    if (item.LastName.length > 0) {
        s += ' ' + item.LastName;
    }
    s += ',  ';
    if (item.Address.length > 0) {
        s += item.Address + ', ';
    }
    if (item.City.length > 0) {
        s += item.City + ', ';
    }
    if (item.State.length > 0) {
        s += item.State;
    }
    return s;
};

//-----------------------------------------------------------------------------
// ResTCRender - renders a name during typedown.
// @params
//   item = an object assumed to have a FirstName and LastName
// @return - The string to render on selection
//-----------------------------------------------------------------------------
window.ResTCRender = function (item) {
        ResSelectedTC = item;
        return ResTCDropRender(item);
};

//-----------------------------------------------------------------------------
// resUsePickedTC - copy the fields of the selected person into the
//     reservation form record.
//
// @params
// @return
//-----------------------------------------------------------------------------
window.resUsePickedTC = function() {
    var f = w2ui.resUpdateForm.record;
    f.TCID = ResSelectedTC.TCID;
    f.Amount = ResSelectedTC.Amount;
    f.FirstName = ResSelectedTC.FirstName;
    f.LastName = ResSelectedTC.LastName;
    f.IsCompany = ResSelectedTC.IsCompany;
    f.CompanyName = ResSelectedTC.CompanyName;
    f.Email = ResSelectedTC.PrimaryEmail;
    f.Phone = (ResSelectedTC.IsCompany) ? ResSelectedTC.WorkPhone : ResSelectedTC.CellPhone;
    f.Street = ResSelectedTC.Address;
    f.City = ResSelectedTC.City;
    f.Country = ResSelectedTC.Country;
    f.State = ResSelectedTC.State;
    f.PostalCode = ResSelectedTC.PostalCode;
    w2ui.resUpdateForm.render();
};
