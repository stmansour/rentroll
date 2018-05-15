/* global
    RACompConfig, sliderContentDivLength, reassignGridRecids,
    getFullName, getTCIDName,
    hideSliderContent, appendNewSlider, showSliderContentW2UIComp,
    loadTargetSection, requiredFieldsFulFilled, getRAFlowPartTypeIndex, initRAFlowAJAX,
    getRAFlowAllParts, saveActiveCompData, toggleHaveCheckBoxDisablity, getRAFlowPartData,
    lockOnGrid,
    loadTransactantListingItem, openNewTransactantForm, getRAAddTransactantFormInitRec,
    acceptTransactant, findTransactantIndexByTCIDInPeopleData, loadRAPeopleForm,
    setRABGInfoFormHeader, setRABGInfoFormFields, showHideRABGInfoFormFields,
    setNotRequiredFields, getRATransanctantDetail, getRABGInfoGridRecord,
    updateRABGInfoFormCheckboxes, getRABGInfoFormInitRecord, loadRABGInfoForm, loadTransactantInRABGInfoGrid
*/

"use strict";

// -------------------------------------------------------------------------------
// Rental Agreement - People form
// -------------------------------------------------------------------------------

//-----------------------------------------------------------------------------
// loadTransactantListingItem - adds transactant into categories list
// @params
//   transactantRec = an object assumed to have a FirstName, MiddleName, LastName,
//                    IsCompany, CompanyName, IsGuarantor, IsOccupant, IsRenter
// @return - nothing
//-----------------------------------------------------------------------------
window.loadTransactantListingItem = function (transactantRec) {

    var peoplePartIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.people);
    if (peoplePartIndex < 0) {
        alert("flow data could not be found for people");
        return false;
    }

    // listing item to be appended in ul
    var s = (transactantRec.IsCompany > 0) ? transactantRec.CompanyName : getFullName(transactantRec);
    if (transactantRec.TCID > 0) {
        s += ' (TCID: ' + String(transactantRec.TCID) + ')';
    }

    var peopleListingItem = '<li data-tcid="' + transactantRec.TCID + '">';
    peopleListingItem += '<span>' + s + '</span>';
    peopleListingItem += '<i class="remove-item fas fa-times-circle fa-xs"></i>';
    peopleListingItem += '</li>';

    // add into renter list
    if (transactantRec.IsRenter) {
        // if with this tcid element exists in DOM then not append
        if ($('#renter-list .people-listing li[data-tcid="' + transactantRec.TCID + '"]').length <= 0) {
            $('#renter-list .people-listing').append(peopleListingItem);
        }
    } else {
        $('#renter-list .people-listing li[data-tcid="' + transactantRec.TCID + '"]').remove();
    }

    // add into occupant list
    if (transactantRec.IsOccupant) {
        // if with this tcid element exists in DOM then not append
        if ($('#occupant-list .people-listing li[data-tcid="' + transactantRec.TCID + '"]').length <= 0) {
            $('#occupant-list .people-listing').append(peopleListingItem);
        }
    } else {
        $('#occupant-list .people-listing li[data-tcid="' + transactantRec.TCID + '"]').remove();
    }

    // add into guarantor list
    if (transactantRec.IsGuarantor) {
        // if with this tcid element exists in DOM then not append
        if ($('#guarantor-list .people-listing li[data-tcid="' + transactantRec.TCID + '"]').length <= 0) {
            $('#guarantor-list .people-listing').append(peopleListingItem);
        }
    } else {
        $('#guarantor-list .people-listing li[data-tcid="' + transactantRec.TCID + '"]').remove();
    }
};

//--------------------------------------------------------------------
// loadTransactantInRABGInfoGrid
//--------------------------------------------------------------------
window.loadTransactantInRABGInfoGrid = function (transactantRec) {
    var peoplePartIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.people);
    if (peoplePartIndex < 0) {
        alert("flow data could not be found for people");
        return false;
    }

    var grid = w2ui.RABGInfoGrid;
    var records = grid.records;
    var isExists = false;

    for (var recordIndex = 0; recordIndex < records.length; recordIndex++){
        if(records[recordIndex].TCID === transactantRec.TCID){
            isExists = true;
            break;
        }
    }

    if(!isExists){
        grid.records.push(transactantRec);
        reassignGridRecids(grid.name);
    }else {
        grid.select(recordIndex + 1);
    }
};

//-----------------------------------------------------------------------------
// openNewTransactantForm - popup new transactant form
//-----------------------------------------------------------------------------
window.openNewTransactantForm = function() {
    var BID = getCurrentBID(),
        BUD = getBUDfromBID(BID);

    // this is new form so TCID is set to zero
    w2ui.RAAddTransactantForm.url = "/v1/person/" + BID.toString() + "/0";
    w2ui.RAAddTransactantForm.record = getRAAddTransactantFormInitRec(BID, BUD, null);
    showSliderContentW2UIComp(w2ui.RAAddTransactantForm, RACompConfig.people.sliderWidth);
    w2ui.RAAddTransactantForm.refresh(); // need to refresh for header changes
};

//-----------------------------------------------------------------------------
// getRAAddTransactantFormInitRec - returns object with default values for
//                                  fields
//-----------------------------------------------------------------------------
window.getRAAddTransactantFormInitRec = function(BID, BUD, previousFormRecord) {

    var defaultFormData = {
        TCID: 0,
        BID: BID,
        BUD: BUD,
        FirstName: "",
        MiddleName: "",
        LastName: "",
        IsCompany: false,
        CompanyName: "",
        EligibleFutureUser: "yes",
        EligibleFuturePayor: "yes",
    };

    // if it called after 'save and add another' action there previous form record is passed as Object
    // else it is null
    if ( previousFormRecord ) {
        defaultFormData = setDefaultFormFieldAsPreviousRecord(
            [ 'FirstName', 'MiddleName', 'LastName', 'IsCompany', 'CompanyName' ], // Fields to Reset
            defaultFormData,
            previousFormRecord
        );
    }

    return defaultFormData;
};

//-----------------------------------------------------------------------------
// acceptTransactant - add transactant to the list of payor/user/guarantor
//
// @params
//   item = an object assumed to have a FirstName, MiddleName, LastName,
//          IsCompany, and CompanyName.
// @return - the name to render
//-----------------------------------------------------------------------------
window.acceptTransactant = function () {
    var IsRenter = w2ui.RAPeopleForm.record.IsRenter;
    var IsOccupant = w2ui.RAPeopleForm.record.IsOccupant;
    var IsGuarantor = w2ui.RAPeopleForm.record.IsGuarantor;

    // if not set anything then alert the user to select any one of them
/*    if (!(IsRenter || IsOccupant || IsGuarantor)) {
        alert("Please, select the role");
        return false;
    }*/

    // get part type index
    var peoplePartIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.people);
    // remove entry from data
    if (peoplePartIndex < 0) {
        return false;
    }

    var transactantRec = $.extend(true, {}, w2ui.RAPeopleForm.record);
    delete transactantRec.Transactant;
    var tcidIndex = findTransactantIndexByTCIDInPeopleData(transactantRec.TCID);

    // if not found then push it in the data
    if (tcidIndex < 0) {
        app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data.push(transactantRec);
    } else {
        app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data[tcidIndex] = transactantRec;
    }

    // load item in the RABGInfoGrid grid
    loadTransactantInRABGInfoGrid(transactantRec);

    // load item in the DOM
    // loadTransactantListingItem(transactantRec);

    // clear the form
    w2ui.RAPeopleForm.actions.reset();

    // disable check boxes
    $(w2ui.RAPeopleForm.box).find("input[type=checkbox]").prop("disabled", true);
};

//-----------------------------------------------------------------------------
// findTransactantIndexByTCIDInPeopleData - finds the index of transactant data
//                in local people data of raflow by TCID
//
// @params
//   TCID = tcid
//-----------------------------------------------------------------------------
window.findTransactantIndexByTCIDInPeopleData = function(TCID) {
    var index = -1;

    // get part type index
    var peoplePartIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.people);
    // remove entry from data
    if (peoplePartIndex < 0) {
        return;
    }

    if (typeof app.raflow.data[app.raflow.activeFlowID] !== "undefined") {
        app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data.forEach(function(transactantRec, i) {
            if (transactantRec.TCID === TCID) {
                index = i;
                return false;
            }
        });
    }

    return index;
};

// click on any transactant item in people listing
$(document).on("click", ".people-listing li", function(e) {
    if($(e.target).hasClass("remove-item")) {
        return;
    }
    var TCID = parseInt($(this).attr("data-tcid"));
    var tcidIndex = findTransactantIndexByTCIDInPeopleData(TCID);
    if (tcidIndex < 0) {
        return;
    }

    // get part type index
    var peoplePartIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.people);
    // remove entry from data
    if (peoplePartIndex < 0) {
        return;
    }

    var transactantRec = app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data[tcidIndex];

    var item = {
        CompanyName: transactantRec.CompanyName,
        IsCompany: transactantRec.IsCompany,
        FirstName: transactantRec.FirstName,
        LastName: transactantRec.LastName,
        MiddleName: transactantRec.MiddleName,
        TCID: transactantRec.TCID,
        recid: 0,
    };
    w2ui.RAPeopleForm.record = transactantRec;
    w2ui.RAPeopleForm.record.Transactant = item;
    w2ui.RAPeopleForm.refresh();
    if ($(w2ui.RAPeopleForm.box).find("input[name=Transactant]").length > 0) {
        $(w2ui.RAPeopleForm.box).find("input[name=Transactant]").data('selected', [item]).data('w2field').refresh();
    }
});

// remove people from the listing
$(document).on('click', '.people-listing .remove-item', function () {
    // get part type index
    var peoplePartIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.people);
    // remove entry from data
    if (peoplePartIndex < 0) {
        return;
    }

    var TCID = parseInt($(this).closest('li').attr('data-tcid'));
    var peopleType = $(this).closest('ul.people-listing').attr('data-people-type');
    var tcidIndex = findTransactantIndexByTCIDInPeopleData(TCID);
    if (tcidIndex < 0) {
        return;
    }

    // taka the reference
    var transactant = app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data[tcidIndex];
    switch (peopleType) {
        case "renters":
            transactant.IsRenter = false;
            break;
        case "occupants":
            transactant.IsOccupant = false;
            break;
        case "guarantors":
            transactant.IsGuarantor = false;
            break;
    }

    // if selected tcid record to be removed, of which all flags are false then remove it from data also
    if (!(transactant.IsGuarantor || transactant.IsRenter || transactant.IsOccupant)) {
        app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data.splice(tcidIndex, 1);

        //
        if (app.raflow.activeTransactant.TCID === transactant.TCID) {
            w2ui.RAPeopleForm.actions.reset(); // clear the form
        }
    }

    // remove item from the DOM
    $(this).closest('li').remove();

    // refresh the form
    if (w2ui.RAPeopleForm.record.TCID > 0) {
        var item = {
            CompanyName: transactant.CompanyName,
            IsCompany: transactant.IsCompany,
            FirstName: transactant.FirstName,
            LastName: transactant.LastName,
            MiddleName: transactant.MiddleName,
            TCID: transactant.TCID,
            recid: 0,
        };
        w2ui.RAPeopleForm.record = $.extend(true, transactant, w2ui.RAPeopleForm.record);
        w2ui.RAPeopleForm.record.Transactant = item;
        w2ui.RAPeopleForm.refresh();
        if ($(w2ui.RAPeopleForm.box).find("input[name=Transactant]").length > 0) {
            $(w2ui.RAPeopleForm.box).find("input[name=Transactant]").data('selected', [item]).data('w2field').refresh();
        }
    }
});

window.loadRAPeopleForm = function () {
    var partType = app.raFlowPartTypes.people;
    var partTypeIndex = getRAFlowPartTypeIndex(partType);
    if (partTypeIndex < 0){
        return;
    }

    // Fetch data from the server if there is any record available.
    getRAFlowPartData(partType)
        .done(function(data){
            if(data.status === 'success'){
                app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data = data.record.Data || [];
                app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data.forEach(function (item) {
                    console.log(item);
                    loadTransactantListingItem(item);
                });
            }else {
                console.log(data.message);
            }
        })
        .fail(function(data){
            console.log("failure" + data);
        });

    // if form is loaded then return
    if (!("RAPeopleForm" in w2ui)) {

        // people form
        $().w2form({
            name: 'RAPeopleForm',
            header: 'People',
            style: 'display: block; border: none;',
            formURL: '/webclient/html/formrapeople.html',
            focus: -1,
            fields: [
                {
                    name: 'Transactant', type: 'enum', required: true, html: {caption: "Transactant"},
                    options: {
                        url: '/v1/transactantstd/' + app.raflow.BID,
                        max: 1,
                        renderItem: function (item) {
                            // enable user-role checkboxes
                            $(w2ui.RAPeopleForm.box).find("input[type=checkbox]").prop("disabled", false);
                            // mark this as transactant as an active
                            app.raflow.activeTransactant = item;
                            var s = getTCIDName(item);
                            w2ui.RAPeopleForm.record.TCID = item.TCID;
                            w2ui.RAPeopleForm.record.FirstName = item.FirstName;
                            w2ui.RAPeopleForm.record.LastName = item.LastName;
                            w2ui.RAPeopleForm.record.MiddleName = item.MiddleName;
                            w2ui.RAPeopleForm.record.CompanyName = item.CompanyName;
                            w2ui.RAPeopleForm.record.IsCompany = item.IsCompany;
                            return s;
                        },
                        renderDrop: function (item) {
                            return getTCIDName(item);
                        },
                        compare: function (item, search) {
                            var s = getTCIDName(item);
                            s = s.toLowerCase();
                            var srch = search.toLowerCase();
                            var match = (s.indexOf(srch) >= 0);
                            return match;
                        },
                        onAdd: function (event) {
                            console.log(event);
                            // if this transactant is available in local data then try to render the bools
                            var tcidIndex = findTransactantIndexByTCIDInPeopleData(event.item.TCID);
                            var transactantRec;
                            if (tcidIndex >= 0) {
                                var peoplePartIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.people);
                                if (peoplePartIndex < 0){
                                    return;
                                }

                                // taka the reference
                                transactantRec = app.raflow.data[app.raflow.activeFlowID][peoplePartIndex].Data[tcidIndex];
                                w2ui.RAPeopleForm.record.IsRenter = transactantRec.IsRenter;
                                w2ui.RAPeopleForm.record.IsOccupant = transactantRec.IsOccupant;
                                w2ui.RAPeopleForm.record.IsGuarantor = transactantRec.IsGuarantor;
                                w2ui.RAPeopleForm.refresh();
                            }
                        },
                        onNew: function (event) {
                            //console.log('++ New Item: Do not forget to submit it to the server too', event);
                            $.extend(event.item, {FirstName: '', LastName: event.item.text});
                        },
                        onRemove: function (event) {
                            event.onComplete = function () {
                                var f = w2ui.RAPeopleForm;
                                // reset payor field related data when removed
                                f.actions.reset();
                                // disable user-role checkboxes
                                $(f.box).find("input[type=checkbox]").prop("disabled", true);

                                // NOTE: have to trigger manually, b'coz we manually change the record,
                                // otherwise it triggers the change event but it won't get change (Object: {})
                                var event = f.trigger({phase: 'before', target: f.name, type: 'change', event: event}); // event before
                                if (event.cancelled === true) return false;
                                f.trigger($.extend(event, {phase: 'after'})); // event after
                            };
                        }
                    }
                },
                {name: 'BID', type: 'int', required: true, html: {caption: "BID"}},
                {name: 'TCID', type: 'int', required: true, html: {caption: "TCID"}},
                {name: 'FirstName', type: 'text', required: true, html: {caption: "FirstName"}},
                {name: 'LastName', type: 'text', required: true, html: {caption: "LastName"}},
                {name: 'MiddleName', type: 'text', required: true, html: {caption: "MiddleName"}},
                {name: 'CompanyName', type: 'text', required: true, html: {caption: "CompanyName"}},
                {name: 'IsCompany', type: 'int', required: true, html: {caption: "IsCompany"}},
                {name: 'IsRenter', type: 'checkbox', hidden: true, required: false, html: {caption: "Renter"}},
                {name: 'IsOccupant', type: 'checkbox', hidden: true, required: false, html: {caption: "Occupant"}},
                {name: 'IsGuarantor', type: 'checkbox', hidden: true, required: false, html: {caption: "Guarantor"}}
            ],
            actions: {
                reset: function () {
                    w2ui.RAPeopleForm.clear();
                    // reset active Transactant to blank object
                    app.raflow.activeTransactant = {};
                }
            },
            onRefresh: function(event) {
                var f = this;
                event.onComplete = function() {
                    var BID = getCurrentBID(),
                        BUD = getBUDfromBID(BID);

                    f.record.BID = BID;
                };
            }
        });

        // new transactant form especially for this RA flow
        $().w2form({
            name: 'RAAddTransactantForm',
            header: 'Add New Transactant',
            style: 'display: block;',
            formURL: '/webclient/html/formra-addtransactant.html',
            focus: -1,
            fields: [
                {name: 'BID', type: 'int', required: true, html: {page: 0, column: 0}},
                {name: 'BUD', type: 'list', required: true, options: {items: app.businesses}, html: {page: 0, column: 0}},
                {name: 'TCID', type: 'int', hidden: false, html: { caption: 'TCID', page: 0, column: 0 } },
                {name: 'EligibleFuturePayor', type: 'text', required: true, html: {caption: "EligibleFuturePayor"}},
                {name: 'EligibleFutureUser', type: 'text', required: true, html: {caption: "EligibleFutureUser"}},
                {name: 'FirstName', type: 'text', required: true, html: {caption: "FirstName"}},
                {name: 'LastName', type: 'text', required: true, html: {caption: "LastName"}},
                {name: 'MiddleName', type: 'text', required: true, html: {caption: "MiddleName"}},
                {name: 'CompanyName', type: 'text', required: false, html: {caption: "CompanyName"}},
                {name: 'IsCompany', type: 'bool', required: false, html: {caption: "IsCompany"}}
            ],
            toolbar : {
                items: [
                    { id: 'bt3', type: 'spacer' },
                    { id: 'btnClose', type: 'button', icon: 'fas fa-times'}
                ],
                onClick: function (event) {
                    switch (event.target){
                        case 'btnClose':
                            hideSliderContent();
                            break;
                    }
                }
            },
            actions: {
                reset: function () {
                    this.clear();
                },
                save: function() {
                    var f = this;
                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    f.save({}, function(data) {
                        if (data.status === 'error') {
                            f.message(data.message);
                            return;
                        }
                        hideSliderContent();
                    });
                },
                saveadd: function() {
                    var BID = getCurrentBID(),
                        BUD = getBUDfromBID(BID),
                        f = this;

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    f.save({}, function(data) {
                        if (data.status === 'error') {
                            f.message(data.message);
                            return;
                        }

                        f.record = getRAAddTransactantFormInitRec(BID, BUD, f.record);
                        f.refresh();
                    });
                },
            },
            onSubmit: function(target, data) {
                if (data.postData.record.IsCompany) {
                    data.postData.record.IsCompany = 1;
                } else {
                    data.postData.record.IsCompany = 0;
                }
            },
            onChange: function(event) {
                event.onComplete = function() {
                    if (this.record.IsCompany) {
                        this.get("FirstName").required = false;
                        this.get("MiddleName").required = false;
                        this.get("LastName").required = false;
                        this.get("CompanyName").required = true;
                        this.get("IsCompany").required = true;
                    } else {
                        this.get("FirstName").required = true;
                        this.get("MiddleName").required = true;
                        this.get("LastName").required = true;
                        this.get("CompanyName").required = false;
                        this.get("IsCompany").required = false;
                    }
                    this.refresh();

                    // formRecDiffer: 1=current record, 2=original record, 3=diff object
                    var diff = formRecDiffer(this.record, app.active_form_original, {});
                    // if diff == {} then make dirty flag as false, else true
                    if ($.isPlainObject(diff) && $.isEmptyObject(diff)) {
                        app.form_is_dirty = false;
                    } else {
                        app.form_is_dirty = true;
                    }
                };
            },
            onRefresh: function(event) {
                var f = this;
                event.onComplete = function() {
                    var BID = getCurrentBID(),
                        BUD = getBUDfromBID(BID);

                    f.record.BID = BID;
                    f.record.BUD = BUD;

                    // there is NO PETID actually, so have to work around with recid key
                    formRefreshCallBack(f, "TCID");
                };
            }
        });

        // background info form
        $().w2form({
            name: 'RABGInfoForm',
            header: 'Background Information',
            style: 'border: 0px; background-color: transparent; display: block;',
            formURL: '/webclient/html/formrabginfo.html',
            toolbar : {
                items: [
                    { id: 'bt3', type: 'spacer' },
                    { id: 'btnClose', type: 'button', icon: 'fas fa-times'}
                ],
                onClick: function (event) {
                    switch (event.target){
                        case 'btnClose':
                            hideSliderContent();
                            break;
                    }
                }
            },
            fields: [
                {name: 'BID', type: 'int', required: true, html: { caption: 'BID', page:0, column: 0}},
                {name: 'TCID', type: 'int', required: true, html: { caption: 'TCID', page:0, column:0}},
                {name: 'IsRenter', type: 'checkbox', required: false}, // will be responsible for paying rent
                {name: 'IsOccupant', type: 'checkbox', required: false}, // will reside in and/or use the items rented
                {name: 'IsGuarantor', type: 'checkbox', required: false}, // responsible for making sure all rent is paid
                {name: 'FirstName', type: 'text', required: true},
                {name: 'MiddleName', type: 'text', required: true},
                {name: 'LastName', type: 'text', required: true},
                {name: 'IsCompany', type: 'checkbox', required: false},
                {name: 'CompanyName', type: 'text', required: false},
                {name: 'BirthDate', type: 'date', required: true}, // Date of births of applicants
                {name: 'SSN', type: 'text', required: true}, // Social security number of applicants
                {name: 'DriverLicNo', type: 'text'}, // Driving licence number of applicants
                {name: 'TelephoneNo', type: 'text', required: true}, // Telephone no of applicants
                {name: 'EmailAddress', type: 'email', required: true}, // Email Address of applicants
                {name: 'CurrentAddress', type: 'text', required: true}, // Current Address
                {name: 'CurrentLandLordName', type: 'text', required: true}, // Current landlord's name
                {name: 'CurrentLandLordPhoneNo', type: 'text', required: true}, // Current landlord's phone number
                {name: 'CurrentLengthOfResidency', type: 'int', required: true}, // Length of residency at current address
                {name: 'CurrentReasonForMoving', type: 'text', required: true}, // Reason of moving from current address
                {name: 'PriorAddress', type: 'text'}, // Prior Address
                {name: 'PriorLandLordName', type: 'text'}, // Prior landlord's name
                {name: 'PriorLandLordPhoneNo', type: 'text'}, // Prior landlord's phone number
                {name: 'PriorLengthOfResidency', type: 'int'}, // Length of residency at Prior address
                {name: 'PriorReasonForMoving', type: 'text'}, // Reason of moving from Prior address
                {name: 'Evicted', type: 'checkbox', required: false}, // have you ever been Evicted
                {name: 'Convicted', type: 'checkbox', required: false}, // have you ever been Arrested or convicted of a crime
                {name: 'Bankruptcy', type: 'checkbox', required: false}, // have you ever been Declared Bankruptcy
                {name: 'Employer', type: 'text', required: true},
                {name: 'Phone', type: 'text', required: true},
                {name: 'Address', type: 'text', required: true},
                {name: 'Address2', type: 'text', required: false},
                {name: 'City', type: 'text', required: false},
                {name: 'State', type: 'list', options: {items: app.usStateAbbr}, required: false},
                {name: 'PostalCode', type: 'text', required: false},
                {name: 'Country', type: 'text', required: false},
                {name: 'Position', type: 'text', required: true},
                {name: 'GrossWages', type: 'money', required: true},
                {name: 'Comment', type: 'text'}, // In an effort to accommodate you, please advise us of any special needs
                {name: 'EmergencyContactName', type: 'text', required: true}, // Name of emergency contact
                {name: 'EmergencyContactPhone', type: 'text', required: true}, // Phone number of emergency contact
                {name: 'EmergencyContactAddress', type: 'text', required: true} // Address of emergency contact
            ],
            actions: {
                save: function () {
                    var form = this;

                    var errors = form.validate();
                    console.log(errors);
                    if (errors.length > 0) return;

                    var record = $.extend(true, {}, form.record);

                    // State filed
                    // TODO(Akshay): Think another way to modify State field if it is possible.
                    record.State = record.State.text;

                    // Convert integer to bool checkboxes fields
                    updateRABGInfoFormCheckboxes(record);

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    var partTypeIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.bginfo);
                    var bgInfoRecords = app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data || [];
                    var isTCIDMatched = false;

                    // update record if it is already exists
                    for(var recordIndex = 0; recordIndex < bgInfoRecords.length; recordIndex++){
                        if(bgInfoRecords[recordIndex].TCID === record.TCID){
                            bgInfoRecords[recordIndex] = record;
                            isTCIDMatched = true;
                            break;
                        }
                    }
                    // push record if it doesn't exists: What about last element
                    if(!isTCIDMatched){
                        bgInfoRecords.push(record);
                    }

                    var recordsData = app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data;
                    // save this records in json Data
                    saveActiveCompData(recordsData, app.raFlowPartTypes.bginfo)
                        .done(function(data) {
                            if (data.status === 'success') {

                                form.clear();

                                // close the form
                                hideSliderContent();
                            } else {
                                form.message(data.message);
                            }
                        })
                        .fail(function(data) {
                            console.log("failure " + data);
                        });
                }
            }
        });

        // transanctants info grid
        $().w2grid({
            name    : 'RABGInfoGrid',
            header  : 'Background information',
            show    : {
                toolbar: true,
                toolbarSearch: false,
                toolbarAdd: true,
                toolbarReload: true,
                toolbarInput: false,
                toolbarColumns: false,
                footer: true
            },
            style   : 'border: 0px solid black; display: block;',
            multiselect: false,
            columns : [
                {
                    field: 'recid',
                    caption: 'recid',
                    size: '50px',
                    hidden: false
                },
                {
                    field: 'TCID',
                    caption: 'TCID',
                    size: '50px',
                    hidden: false
                },
                {
                    field: 'FullName',
                    caption: 'Name',
                    size: '150px',
                    render: function (record) {
                        if(!record.IsCompany){
                            return getFullName(record);
                        }else{
                            return record.CompanyName;
                        }

                    }
                },
                {
                    field: 'IsRenter',
                    caption: 'Renter',
                    size: '100px',
                    hidden: false,
                    render: function (record) {
                        if(record.IsRenter){
                            return '<i class="fas fa-check" title="renter"></i>';
                        }else{
                            return '<i class="fas fa-times" title="renter"></i>';
                        }
                    }
                },
                {
                    field: 'IsOccupant',
                    caption: 'Occupant',
                    size: '100px',
                    hidden: false,
                    render: function (record) {
                        if(record.IsOccupant){
                            return '<i class="fas fa-check" title="occupant"></i>';
                        }else{
                            return '<i class="fas fa-times" title="occupant"></i>';
                        }
                    }
                },
                {
                    field: 'IsGuarantor',
                    caption: 'Guarantor',
                    size: '100px',
                    hidden: false,
                    render: function (record) {
                        if(record.IsGuarantor){
                            return '<i class="fas fa-check" title="guarantor"></i>';
                        }else{
                            return '<i class="fas fa-times" title="guarantor"></i>';
                        }
                    }
                }
            ],
            onClick : function(event) {
                event.onComplete = function() {

                    console.log(event.recid);
                    var raBGInfoGridRecord = w2ui.RABGInfoGrid.get(event.recid); // record from the w2ui grid
                    console.log(raBGInfoGridRecord);

                    var yes_args = [this, event.recid],
                        no_args = [this],
                        no_callBack = function(grid) {
                            grid.select(app.last.grid_sel_recid);
                            return false;
                        },
                        yes_callBack = function(grid, recid) {
                            app.last.grid_sel_recid = parseInt(recid);

                            // keep highlighting current row in any case
                            grid.select(app.last.grid_sel_recid);

                            w2ui.RABGInfoForm.record = $.extend(true, {}, grid.get(app.last.grid_sel_recid));

                            showSliderContentW2UIComp(w2ui.RABGInfoForm, RACompConfig.bginfo.sliderWidth);

                            // get transanctant information from the server
                            getRATransanctantDetail(raBGInfoGridRecord.TCID)
                                .done(function (data){
                                    // Hide these all fields when transanctant is only user.
                                    var listOfHiddenFields = ["CurrentAddress", "CurrentLandLordName",
                                        "CurrentLandLordPhoneNo", "CurrentLengthOfResidency", "CurrentReasonForMoving",
                                        "PriorAddress", "PriorLandLordName", "PriorLandLordPhoneNo",
                                        "PriorLengthOfResidency", "PriorReasonForMoving"];

                                    // These all fields are not required when transanctant is only user
                                    var listOfNotRequiredFields = ["SSN", "TelephoneNo",
                                        "Phone", "EmailAddress", "Position",
                                        "GrossWages", "CurrentAddress", "CurrentLandLordName",
                                        "CurrentLandLordPhoneNo", "CurrentReasonForMoving"];

                                    if(data.status === 'success'){
                                        var record = data.record; // record from the server response
                                        // w2ui.RABGInfoForm.record = data.record.Data;
                                        var BID = getCurrentBID();

                                        // Set the form tile
                                        setRABGInfoFormHeader(record);

                                        // Display/Required field based on transanctant type
                                        if(raBGInfoGridRecord.IsOccupant && !raBGInfoGridRecord.IsRenter && !raBGInfoGridRecord.IsGuarantor){
                                            // hide fields
                                            showHideRABGInfoFormFields(listOfHiddenFields, true);

                                            // not require fields
                                            setNotRequiredFields(listOfNotRequiredFields, false);
                                        }else{
                                            // show fields
                                            showHideRABGInfoFormFields(listOfHiddenFields, false);

                                            // require fields
                                            setNotRequiredFields(listOfNotRequiredFields, true);

                                        }

                                        // If TCID found than set form fields value else initialize it.
                                        var partTypeIndex = getRAFlowPartTypeIndex(app.raFlowPartTypes.bginfo);
                                        var bgInfoRecords = app.raflow.data[app.raflow.activeFlowID][partTypeIndex].Data || [];
                                        var isTCIDMatched = false;

                                        for(var recordIndex = 0; recordIndex < bgInfoRecords.length; recordIndex++){
                                            if(bgInfoRecords[recordIndex].TCID === raBGInfoGridRecord.TCID){
                                                w2ui.RABGInfoForm.record = bgInfoRecords[recordIndex];
                                                isTCIDMatched = true;
                                                break;
                                            }
                                        }

                                        if(!isTCIDMatched){
                                            // Assign default values to form fields
                                            w2ui.RABGInfoForm.record = getRABGInfoFormInitRecord(BID, raBGInfoGridRecord.TCID);

                                            // Set latest value for transanctant basic information from the server only
                                            setRABGInfoFormFields(record);
                                        }


                                        w2ui.RABGInfoForm.refresh(); // need to refresh for form changes
                                    }else {
                                        console.log(data.message);
                                    }
                                })
                                .fail(function (data) {
                                    console.log("failure" + data);
                                });
                        };

                    // warn user if form content has been changed
                    form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
                };
            },
            onAdd   : function () {
                openNewTransactantForm();
            },
        });
    }

    // load form in div
    $('#ra-form #people .grid-container').w2render(w2ui.RABGInfoGrid);
    $('#ra-form #people .form-container').w2render(w2ui.RAPeopleForm);

    // Do not remove code below code for now
     setTimeout(function () {
        var i = getRAFlowPartTypeIndex(app.raFlowPartTypes.people);
        if (i >= 0 && app.raflow.data[app.raflow.activeFlowID][i].Data) {
            // w2ui.RAPeopleForm.record = app.raflow.data[app.raflow.activeFlowID][i].Data;
            w2ui.RAPeopleForm.refresh();
        } else {
            w2ui.RAPeopleForm.actions.reset();
        }
    }, 500);

    // load the existing data in people component
    setTimeout(function () {
        var grid = w2ui.RABGInfoGrid;
        var i = getRAFlowPartTypeIndex(app.raFlowPartTypes.bginfo);
        if (i >= 0 && app.raflow.data[app.raflow.activeFlowID][i].Data) {
            grid.records = app.raflow.data[app.raflow.activeFlowID][i].Data;
            reassignGridRecids(grid.name);

        } else {
            grid.clear();
        }
    }, 500);
};


// -------------------------------------------------------------------------------
// Rental Agreement - Background info part
// -------------------------------------------------------------------------------

// setRABGInfoFormHeader
// TODO(Akshay): Ask steve about form fields when transanctant is company.
window.setRABGInfoFormHeader = function(record) {
    if(record.IsCompany){
        w2ui.RABGInfoForm.header = 'Background Information - ' + record.CompanyName;
    }else{
        w2ui.RABGInfoForm.header = 'Background Information - ' + record.FirstName + ' ' + record.MiddleName + ' ' + record.LastName;
    }
};

// setRABGInfoFormFields It set default fields value from the transanctants to form.
window.setRABGInfoFormFields = function(record) {
    var formRecord = w2ui.RABGInfoForm.record; // record from the w2ui form

    formRecord.FirstName = record.FirstName;
    formRecord.MiddleName = record.MiddleName;
    formRecord.LastName = record.LastName;
    formRecord.TelephoneNo = record.CellPhone;
    formRecord.EmailAddress = record.PrimaryEmail;
    formRecord.Phone = record.WorkPhone;
    formRecord.Address = record.Address;
    formRecord.Address2 = record.Address2;
    formRecord.City = record.City;
    formRecord.Country = record.Country;
    formRecord.PostalCode = record.PostalCode;
    formRecord.State = record.State;
};

// showHideRABGInfoFormFields hide fields if transanctant is only user
window.showHideRABGInfoFormFields = function(listOfHiddenFields, hidden){
    if(hidden){
        $("#cureentInfolabel").hide();
        $("#priorInfolabel").hide();
    }else{
        $("#cureentInfolabel").show();
        $("#priorInfolabel").show();
    }
    for(var fieldIndex=0; fieldIndex < listOfHiddenFields.length; fieldIndex++){
        w2ui.RABGInfoForm.get(listOfHiddenFields[fieldIndex]).hidden = hidden;
    }
};

// setNotRequiredFields define fields are not required if transanctant is only user
window.setNotRequiredFields = function(listOfNotRequiredFields, required){
    for(var fieldIndex=0; fieldIndex < listOfNotRequiredFields.length; fieldIndex++){
        w2ui.RABGInfoForm.get(listOfNotRequiredFields[fieldIndex]).required = required;
    }
};

// get RATransactant detail from the server
window.getRATransanctantDetail = function(TCID){
    var bid = getCurrentBID();

    // temporary data
    var data = {
        "cmd":"get",
        "recid":0,
        "name":"transactantForm"
    };


    return $.ajax({
        url: "/v1/person/" + bid.toString() + "/" + TCID,
        method: "POST",
        contentType: "application/json",
        dataType: "json",
        data: JSON.stringify(data),
        success: function (data) {
            if (data.status != "error"){
                // console.log("Received data for transanctant:", JSON.stringify(data));
            }else {
                console.error(data.message);
            }
        },
        error: function () {
            console.log("Error:" + JSON.stringify(data));
        }
    });
};

// getRABGInfoGridRecord
window.getRABGInfoGridRecord = function(records, TCID){
    var raBGInfoGridrecord;
    for(var recordIndex=0; recordIndex < records.length; recordIndex++) {
        if(records[recordIndex].TCID == TCID){
            raBGInfoGridrecord = records[recordIndex];
            break;
        }
    }
    return raBGInfoGridrecord;
};

// updateRABGInfoFormCheckboxes
window.updateRABGInfoFormCheckboxes = function(record){
    record.Evicted = int_to_bool(record.Evicted);
    record.Bankruptcy = int_to_bool(record.Bankruptcy);
    record.Convicted = int_to_bool(record.Convicted);
};

window.getRABGInfoFormInitRecord = function(BID, TCID){

    return {
        recid: 0,
        TCID: TCID,
        BID: BID,
        FirstName: "",
        MiddleName: "",
        LastName: "",
        BirthDate: "",
        SSN: "",
        DriverLicNo: "",
        TelephoneNo: "",
        EmailAddress: "",
        CurrentAddress: "",
        CurrentLandLordName: "",
        CurrentLandLordPhoneNo: "",
        CurrentLengthOfResidency: 0,
        CurrentReasonForMoving: "",
        PriorAddress: "",
        PriorLandLordName: "",
        PriorLandLordPhoneNo: "",
        PriorLengthOfResidency: 0,
        PriorReasonForMoving: "",
        Evicted: false,
        Convicted: false,
        Bankruptcy: false,
        Employer: "",
        Phone: "",
        Address: "",
        Position: "",
        GrossWages: 0,
        Comment: "",
        EmergencyContactName: "",
        EmergencyContactPhone: "",
        EmergencyContactAddress: ""
    };
};
