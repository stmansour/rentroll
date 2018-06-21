/* global
    RACompConfig, sliderContentDivLength, reassignGridRecids,
    getFullName, getTCIDName,
    hideSliderContent, appendNewSlider, showSliderContentW2UIComp,
    loadTargetSection, requiredFieldsFulFilled, getRAFlowPartTypeIndex, initRAFlowAjax,
    getRAFlowAllParts, saveActiveCompData, toggleHaveCheckBoxDisablity, getRAFlowCompData,
    openNewTransactantForm, getRAAddTransactantFormInitRec,
    acceptTransactant, loadRAPeopleForm,
    setRATransactantFormHeader, showHideRATransactantFormFields,
    setNotRequiredFields, getRAPeopleGridRecord,
    updateRATransactantFormCheckboxes, getRATransactantFormInitRecord, loadRATransactantForm, ReassignPeopleGridRecords,
    manageBGInfoFormFields, addDummyBackgroundInfo, savePeopleCompData, getPeopleLocalData, setPeopleLocalData,
    getPeopleLocalDataByTCID, setTransactantDefaultRole, transactantTabs, transactantFields,
    savePetsCompData, saveVehiclesCompData, setRAFlowCompData, getStringListData, getSLStringList, updateRATransactantFormCheckboxes,
    managePeopleW2UIItems, removeRAFlowPersonAJAX, saveRAFlowPersonAJAX
*/

"use strict";

/*
* Note: Reuse method from transactant.js: getStringListData, getSLStringList, updateRATransactantFormCheckboxes
* */

// -------------------------------------------------------------------------------
// Rental Agreement - People form, People Grid, Background information form
// -------------------------------------------------------------------------------
window.loadRAPeopleForm = function () {

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

                            // Enable Accept button
                            $(w2ui.RAPeopleForm.box).find("button[name=accept]").prop("disabled", false);

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
                        onRemove: function(event) {
                            event.onComplete = function() {
                                w2ui.RAPeopleForm.actions.reset();
                            };
                        }
                    }
                },
                {name: 'BID', type: 'int', required: true, html: {caption: "BID"}},
                {name: 'TCID', type: 'int', required: true, html: {caption: "TCID"}},
                {name: 'FirstName', type: 'text', required: false, html: {caption: "FirstName"}},
                {name: 'LastName', type: 'text', required: false, html: {caption: "LastName"}},
                {name: 'MiddleName', type: 'text', required: false, html: {caption: "MiddleName"}},
                {name: 'CompanyName', type: 'text', required: false, html: {caption: "CompanyName"}},
                {name: 'IsCompany', type: 'checkbox', required: true, html: {caption: "IsCompany"}}
            ],
            actions: {
                reset: function () {
                    w2ui.RAPeopleForm.clear();
                    $(w2ui.RAPeopleForm.box).find("button[name=accept]").prop("disabled", true);
                }
            },
            onRefresh: function (event) {
                var f = this;
                event.onComplete = function () {
                    var BID = getCurrentBID(),
                        BUD = getBUDfromBID(BID);

                    f.record.BID = BID;
                };
            }
        });

        // transanctants/people list in grid
        $().w2grid({
            name: 'RAPeopleGrid',
            header: 'Background information',
            show: {
                toolbar: true,
                toolbarSearch: false,
                toolbarAdd: true,
                toolbarReload: true,
                toolbarInput: false,
                toolbarColumns: false,
                footer: true
            },
            style: 'border: 0px solid black; display: block;',
            multiSelect: false,
            columns: [
                {
                    field: 'recid',
                    hidden: true
                },
                {
                    field: 'TMPTCID',
                    hidden: true
                },
                {
                    field: 'TCID',
                    hidden: true
                },
                {
                    field: 'FullName',
                    caption: 'Name',
                    size: '100%',
                    style: 'text-align: left;',
                    render: function (record) {
                        if (record.IsCompany) {
                            return record.CompanyName;
                        } else {
                            return getFullName(record);
                        }
                    }
                },
                {
                    field: 'IsRenter',
                    // caption: 'Renter',
                    // size: '100px',
                    hidden: true,
                    // render: function (record) {
                    //     if (record.IsRenter) {
                    //         return '<i class="fas fa-check" title="renter"></i>';
                    //     } else {
                    //         return '<i class="fas fa-times" title="renter"></i>';
                    //     }
                    // }
                },
                {
                    field: 'IsOccupant',
                    // caption: 'Occupant',
                    // size: '100px',
                    hidden: true,
                    // render: function (record) {
                    //     if (record.IsOccupant) {
                    //         return '<i class="fas fa-check" title="occupant"></i>';
                    //     } else {
                    //         return '<i class="fas fa-times" title="occupant"></i>';
                    //     }
                    // }
                },
                {
                    field: 'IsGuarantor',
                    // caption: 'Guarantor',
                    // size: '100px',
                    hidden: true,
                    // render: function (record) {
                    //     if (record.IsGuarantor) {
                    //         return '<i class="fas fa-check" title="guarantor"></i>';
                    //     } else {
                    //         return '<i class="fas fa-times" title="guarantor"></i>';
                    //     }
                    // }
                }
            ],
            onClick: function (event) {
                event.onComplete = function () {
                    var yes_args = [this, event.recid],
                        no_args = [this],
                        no_callBack = function (grid) {
                            grid.select(app.last.grid_sel_recid);
                            return false;
                        },
                        yes_callBack = function (grid, recid) {
                            var form = w2ui.RATransactantForm,
                                BID = getCurrentBID(),
                                BUD = getBUDfromBID(BID);

                            app.last.grid_sel_recid = parseInt(recid);

                            // keep highlighting current row in any case
                            grid.select(app.last.grid_sel_recid);

                            var raBGInfoGridRecord = grid.get(event.recid); // record from the w2ui grid

                            // get stringListData for list fields
                            getStringListData(BID, BUD).fail(function (data) {
                                form.message(data.message);
                            });

                            // show slider content in w2ui comp
                            showSliderContentW2UIComp(form, RACompConfig.people.sliderWidth);

                            // show/hide list of fields based on role
                            manageBGInfoFormFields(raBGInfoGridRecord);

                            form.record = getPeopleLocalData(raBGInfoGridRecord.TMPTCID);
                            form.record.recid = raBGInfoGridRecord.recid;
                            form.record.BID = BID;
                            form.record.BUD = BUD;

                            // Set the form title
                            setRATransactantFormHeader(form.record);

                            form.refresh(); // need to refresh for form changes
                        };

                    // warn user if form content has been changed
                    form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
                };
            },
            onAdd: function () {
                openNewTransactantForm();
            }
        });

        // background info form
        $().w2form({
            name: 'RATransactantForm',
            header: 'Background Information',
            style: 'border: 0px; background-color: transparent; display: block;',
            formURL: '/webclient/html/formtc.html',
            toolbar: {
                items: [
                    {id: 'bt3', type: 'spacer'},
                    // {id: 'addInfo', type: 'button', icon: 'fas fa-plus-circle'}, // TODO: Remove this in production. This button is for development purpose
                    {id: 'btnClose', type: 'button', icon: 'fas fa-times'}
                ],
                onClick: function (event) {
                    switch (event.target) {
                        case 'btnClose':
                            hideSliderContent();
                            break;
                        case 'addInfo':
                            addDummyBackgroundInfo();
                            break;
                    }
                }
            },
            fields: app.transactantFields,
            tabs: app.transactantTabs,
            actions: {
                save: function () {
                    var form = this,
                        TMPTCID = form.record.TMPTCID;

                    var errors = form.validate();
                    if (errors.length > 0) return;

                    var peopleData = getFormSubmitData(form.record, true);

                    // If transanctant role isn't selected than display error.
                    if (!(peopleData.IsRenter || peopleData.IsOccupant || peopleData.IsGuarantor)) {
                        form.message("Please select transanctant role.");
                        return;
                    }

                    // Convert integer to bool checkboxes fields
                    updateRATransactantFormCheckboxes(peopleData);

                    setPeopleLocalData(TMPTCID, peopleData);

                    // clean dirty flag of form
                    app.form_is_dirty = false;

                    // save this records in json Data
                    savePeopleCompData()
                    .done(function (data) {
                        if (data.status === 'success') {

                            form.clear();

                            // update RAPeopleGrid
                            ReassignPeopleGridRecords();

                            // close the form
                            hideSliderContent();
                        } else {
                            form.message(data.message);
                        }
                    })
                    .fail(function (data) {
                        console.error("failure " + data);
                    });
                },
                delete: function () {
                    var form = this;
                    var TMPTCID = form.record.TMPTCID;

                    removeRAFlowPersonAJAX(TMPTCID)
                    .done(function(data) {
                        if (data.status === 'success') {
                            // clear the form
                            form.clear();

                            // update RAPeopleGrid
                            ReassignPeopleGridRecords();

                            // close the form
                            hideSliderContent();
                        } else {
                            form.message(data.message);
                        }
                    })
                    .fail(function (data) {
                        console.error("failure " + data);
                    });
                },
                reset: function () {
                    w2ui.RATransactantForm.clear();
                }
            },
            onChange: function (event) {
                event.onComplete = function () {
                    $("#EvictedDes").prop("disabled", !this.record.Evicted);
                    $("#ConvictedDes").prop("disabled", !this.record.Convicted);
                    $("#BankruptcyDes").prop("disabled", !this.record.Bankruptcy);

                    manageBGInfoFormFields(this.record);

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
            onRefresh: function (event) {
                event.onComplete = function () {
                    var form = this,
                        BID = getCurrentBID(),
                        BUD = getBUDfromBID(BID);

                    // Set list field value
                    form.get('SourceSLSID').options.items = getSLStringList(BID, "HowFound");
                    form.get('DeclineReasonSLSID').options.items = getSLStringList(BID, "ApplDeny");
                    form.get('CurrentReasonForMoving').options.items = getSLStringList(BID, "WhyLeaving");
                    form.get('PriorReasonForMoving').options.items = getSLStringList(BID, "WhyLeaving");

                    // hide delete button if it is NewRecord
                    var isNewRecord = (w2ui.RAPeopleGrid.get(form.record.recid, true) === null);
                    if (isNewRecord) {
                        $(form.box).find("button[name=delete]").addClass("hidden");
                    } else {
                        $(form.box).find("button[name=delete]").removeClass("hidden");
                    }

                    $("#EvictedDes").prop("disabled", !this.record.Evicted);
                    $("#ConvictedDes").prop("disabled", !this.record.Convicted);
                    $("#BankruptcyDes").prop("disabled", !this.record.Bankruptcy);
                };
            },
            onValidate: function (event) {
                if (!this.record.IsCompany && this.record.FirstName === '') {
                    event.errors.push({
                        field: this.get('FirstName'),
                        error: 'FirstName required when "Person or Company" field is set to Person'
                    });
                }
                if (!this.record.IsCompany && this.record.LastName === '') {
                    event.errors.push({
                        field: this.get('LastName'),
                        error: 'LastName required when "Person or Company" field is set to Person'
                    });
                }
                if (this.record.IsCompany && this.record.CompanyName === '') {
                    event.errors.push({
                        field: this.get('CompanyName'),
                        error: 'Company Name required when "Person or Company" field is set to Company'
                    });
                }
            }
        });
    }

    // load form in div
    $('#ra-form #people .grid-container').w2render(w2ui.RAPeopleGrid);
    $('#ra-form #people .form-container').w2render(w2ui.RAPeopleForm);

    // load existing info in PeopleForm and PeopleGrid
    setTimeout(function () {
        // Operation on RAPeopleGrid
        ReassignPeopleGridRecords();
    }, 500);
};

// setRATransactantFormHeader
// It set RATransactantForm header title
window.setRATransactantFormHeader = function (record) {
    if (record.IsCompany > 0) {
        w2ui.RATransactantForm.header = 'Background Information - ' + record.CompanyName;
    } else {
        w2ui.RATransactantForm.header = 'Background Information - ' + record.FirstName + ' ' + record.MiddleName + ' ' + record.LastName;
    }
};

// showHideRATransactantFormFields
// hide fields if transanctant is only user
window.showHideRATransactantFormFields = function (listOfHiddenFields, hidden) {
    if (hidden) {
        $("#cureentInfolabel").hide();
        $("#priorInfolabel").hide();
    } else {
        $("#cureentInfolabel").show();
        $("#priorInfolabel").show();
    }
    for (var fieldIndex = 0; fieldIndex < listOfHiddenFields.length; fieldIndex++) {
        w2ui.RATransactantForm.get(listOfHiddenFields[fieldIndex]).hidden = hidden;
    }
};

// setNotRequiredFields
// define fields are not required if transanctant is only user
window.setNotRequiredFields = function (listOfNotRequiredFields, required) {
    for (var fieldIndex = 0; fieldIndex < listOfNotRequiredFields.length; fieldIndex++) {
        w2ui.RATransactantForm.get(listOfNotRequiredFields[fieldIndex]).required = required;
    }
};

// remove person with associated pets, vehicles from json data via Ajax
window.removeRAFlowPersonAJAX = function (TMPTCID) {
    var bid = getCurrentBID();

    // temporary data
    var data = {
        "cmd": "delete",
        "TMPTCID": TMPTCID,
        "FlowID": app.raflow.activeFlowID
    };

    return $.ajax({
        url: "/v1/raflow-person/" + bid.toString() + "/" + app.raflow.activeFlowID.toString(),
        method: "POST",
        contentType: "application/json",
        dataType: "json",
        data: JSON.stringify(data),
        success: function (data) {
            if (data.status != "error") {
                // update the local copy of flow for the active one
                app.raflow.data[data.record.FlowID] = data.record;
            } else {
                console.error(data.message);
            }
        },
        error: function () {
            console.error("Error:" + JSON.stringify(data));
        }
    });
};

// save Transanctant in raflow
window.saveRAFlowPersonAJAX = function (TCID) {
    var bid = getCurrentBID();

    // temporary data
    var data = {
        "cmd": "save",
        "TCID": TCID,
        "FlowID": app.raflow.activeFlowID
    };

    return $.ajax({
        url: "/v1/raflow-person/" + bid.toString() + "/" + app.raflow.activeFlowID.toString(),
        method: "POST",
        contentType: "application/json",
        dataType: "json",
        data: JSON.stringify(data),
        success: function (data) {
            if (data.status != "error") {
                // update the local copy of flow for the active one
                app.raflow.data[data.record.FlowID] = data.record;
            } else {
                console.error(data.message);
            }
        },
        error: function () {
            console.error("Error:" + JSON.stringify(data));
        }
    });
};

// getRAPeopleGridRecord
// get record from the list which match with TCID
window.getRAPeopleGridRecord = function (records, TCID) {
    var raBGInfoGridrecord;
    for (var recordIndex = 0; recordIndex < records.length; recordIndex++) {
        if (records[recordIndex].TCID === TCID) {
            raBGInfoGridrecord = records[recordIndex];
            break;
        }
    }
    return raBGInfoGridrecord;
};

//--------------------------------------------------------------------
// ReassignPeopleGridRecords
//--------------------------------------------------------------------
window.ReassignPeopleGridRecords = function () {
    var compData = getRAFlowCompData("people", app.raflow.activeFlowID);
    var grid = w2ui.RAPeopleGrid;

    if (compData) {
        grid.records = compData;
        reassignGridRecids(grid.name);

        // Operation on RAPeopleForm
        w2ui.RAPeopleForm.refresh();

        // manage people w2ui items list
        managePeopleW2UIItems();

    } else {
        // Operation on RAPeopleForm
        w2ui.RAPeopleForm.actions.reset();

        // Operation on RAPeopleGrid
        grid.clear();
    }
};

//-----------------------------------------------------------------------------
// openNewTransactantForm - popup new transactant form
//-----------------------------------------------------------------------------
window.openNewTransactantForm = function () {
    var BID = getCurrentBID(),
        BUD = getBUDfromBID(BID),
        form = w2ui.RATransactantForm;

    form.header = 'Background Information';
    form.record = getTransactantInitRecord(BID, BUD);
    setTransactantDefaultRole(form.record);

    // get stringListData for list fields
    getStringListData(BID, BUD).fail(function (data) {
        form.message(data.message);
    });

    showSliderContentW2UIComp(form, RACompConfig.people.sliderWidth);

    form.refresh(); // need to refresh for header changes
};

//-----------------------------------------------------------------------------
// acceptTransactant - add transactant to the list of payor/user/guarantor
//
// @params
//   item = an object assumed to have a FirstName, MiddleName, LastName,
//          IsCompany, and Employer.
// @return - the name to render
//-----------------------------------------------------------------------------
window.acceptTransactant = function () {

    var compData = getRAFlowCompData("people", app.raflow.activeFlowID) || [];

    var peopleForm = w2ui.RAPeopleForm;

    var transactantRec = $.extend(true, {}, peopleForm.record);
    delete transactantRec.Transactant;
    var TCID = transactantRec.TCID;

    var tcidIndex = getPeopleLocalDataByTCID(TCID, true);

    // if not found then push it in the data
    if (tcidIndex < 0) {

        // save transanctant information in raflow json
        saveRAFlowPersonAJAX(TCID)
        .done(function (data) {

            if (data.status === 'success') {

                // load item in the RAPeopleGrid grid
                ReassignPeopleGridRecords();

                // clear the form
                w2ui.RAPeopleForm.actions.reset();

            } else {
                console.error(data.message);
            }
        })
        .fail(function (data) {
            console.error("failure" + data);
        });
    } else {
        var recid = compData[tcidIndex].recid;

        // Show selected row for existing transanctant record
        w2ui.RAPeopleGrid.select(recid);

        // clear the form
        w2ui.RAPeopleForm.actions.reset();
    }

};

// manageBGInfoFormFields
window.manageBGInfoFormFields = function (record) {
    // Hide these all fields when transanctant is only user.
    var listOfHiddenFields = ["CurrentAddress", "CurrentLandLordName",
        "CurrentLandLordPhoneNo", "CurrentLengthOfResidency", "CurrentReasonForMoving",
        "PriorAddress", "PriorLandLordName", "PriorLandLordPhoneNo",
        "PriorLengthOfResidency", "PriorReasonForMoving"];

    // Display/Required field based on transanctant type
    var haveToHide = record.IsOccupant && !record.IsRenter && !record.IsGuarantor; // true: hide fields, false: show fields
    // hide/show fields
    showHideRATransactantFormFields(listOfHiddenFields, haveToHide);
};

window.addDummyBackgroundInfo = function () {
    var form = w2ui.RATransactantForm;
    var record = form.record;
    record.FirstName = Math.random().toString(32).slice(2);
    record.MiddleName = Math.random().toString(32).slice(2);
    record.LastName = Math.random().toString(32).slice(2);
    record.DateofBirth = "8/30/1990";
    record.SSN = Math.random().toString(32).slice(4);
    record.DriversLicense = Math.random().toString(32).slice(2);
    record.CellPhone = Math.random().toString(32).slice(2);
    record.PrimaryEmail = Math.random().toString(32).slice(2) + "@yopmail.com";
    record.CurrentAddress = Math.random().toString(32).slice(2);
    record.CurrentLandLordName = Math.random().toString(32).slice(2);
    record.CurrentLandLordPhoneNo = Math.random().toString(32).slice(2);
    record.CurrentLengthOfResidency = 56;
    record.CurrentReasonForMoving = Math.random().toString(32).slice(2);
    record.PriorAddress = Math.random().toString(32).slice(2);
    record.PriorLandLordName = Math.random().toString(32).slice(2);
    record.PriorLandLordPhoneNo = Math.random().toString(32).slice(2);
    record.PriorLengthOfResidency = "3 years";
    record.PriorReasonForMoving = Math.random().toString(32).slice(2);
    record.CompanyName = Math.random().toString(32).slice(2);
    record.WorkPhone = Math.random().toString(32).slice(2);
    record.Address = Math.random().toString(32).slice(2);
    record.Position = Math.random().toString(32).slice(2);
    record.GrossIncome = Math.random() * 100;
    record.EmergencyContactName = Math.random().toString(32).slice(2);
    record.EmergencyContactPhone = Math.random().toString(32).slice(2);
    record.EmergencyContactAddress = Math.random().toString(32).slice(2);
    form.refresh();
};

//------------------------------------------------------------------------------
// savePetsCompData - saves the data on server side
//------------------------------------------------------------------------------
window.savePeopleCompData = function() {
	var compData = getRAFlowCompData("people", app.raflow.activeFlowID);
	return saveActiveCompData(compData, "people");
};

//-----------------------------------------------------------------------------
// getPeopleLocalDataByTCID - returns the clone of people data
//                            for requested TCID
//-----------------------------------------------------------------------------
window.getPeopleLocalDataByTCID = function(TCID, returnIndex) {
    var cloneData = {};
    var foundIndex = -1;
    var compData = getRAFlowCompData("people", app.raflow.activeFlowID) || [];
    compData.forEach(function(item, index) {
        if (item.TCID === TCID) {
            if (returnIndex) {
                foundIndex = index;
            } else {
                cloneData = $.extend(true, {}, item);
            }
            return false;
        }
    });
    if (returnIndex) {
        return foundIndex;
    }
    return cloneData;
};

//-----------------------------------------------------------------------------
// getPeopleLocalData - returns the clone of people data for requested TMPTCID
//-----------------------------------------------------------------------------
window.getPeopleLocalData = function(TMPTCID, returnIndex) {
	var cloneData = {};
	var foundIndex = -1;
	var compData = getRAFlowCompData("people", app.raflow.activeFlowID) || [];
	compData.forEach(function(item, index) {
		if (item.TMPTCID === TMPTCID) {
			if (returnIndex) {
				foundIndex = index;
			} else {
				cloneData = $.extend(true, {}, item);
			}
			return false;
		}
	});
	if (returnIndex) {
		return foundIndex;
	}
	return cloneData;
};

//-----------------------------------------------------------------------------
// setPeopleLocalData - save the data for requested a TMPTCID in local data
//-----------------------------------------------------------------------------
window.setPeopleLocalData = function(TMPTCID, peopleData) {
	var compData = getRAFlowCompData("people", app.raflow.activeFlowID) || [];
	var dataIndex = -1;
	compData.forEach(function(item, index) {
		if (item.TMPTCID === TMPTCID) {
			dataIndex = index;
			return false;
		}
	});
	if (dataIndex > -1) {
		compData[dataIndex] = peopleData;
	} else {
		compData.push(peopleData);
	}
};

//-----------------------------------------------------------------------------
// setTransactantDefaultRole - Assign default role for new transanctant.
//-----------------------------------------------------------------------------
window.setTransactantDefaultRole = function (transactantRec) {
	var compData = getRAFlowCompData("people", app.raflow.activeFlowID) || [];
	// If first record in the grid than transanctant will be renter by default
	if (compData.length === 0) {
		transactantRec.IsRenter = true;
	}

	// Each transactant must be occupant by default. It can be change via BGInfo detail form
	transactantRec.IsOccupant = true;
};

//-------------------------------------------------------------------------------
// managePeopleW2UIItems - maintain people w2ui items list
//                       - it will keep TMPTCID as in id of w2ui list item object
//-------------------------------------------------------------------------------
window.managePeopleW2UIItems = function() {

    // reset it first
    app.raflow.peopleW2UIItems = [];

    // inner function to push item in "app.raflow.peopleW2UIItems"
    var pushItem = function(peopleItem, atIndex) {
        var found = false;
        app.raflow.peopleW2UIItems.forEach(function(item) {
            if (item.id === peopleItem.id) {
                found = true;
                return false;
            }
        });

        // if not found the push item in app.raflow.peopleW2UIItems
        if (!found) {
            // deleting 0 item, then insert peopleItem
            app.raflow.peopleW2UIItems.splice(atIndex, 0, peopleItem);
        }
    };

    // get comp data
    var peopleCompData = getRAFlowCompData("people", app.raflow.activeFlowID) || [];

    // first build the list of rentables and sort it out in asc order of TMPTCID
    peopleCompData.forEach(function(peopleItem) {
        var TMPTCID = peopleItem.TMPTCID,
            Name = "";

        // get name
        if (peopleItem.IsCompany) {
            Name = peopleItem.CompanyName;
        } else {
            Name = getFullName(peopleItem);
        }

        // w2ui list item for people list
        var item = {id: TMPTCID, text: Name};
        pushItem(item, app.raflow.peopleW2UIItems.length);
    });

    // sort it out in asc order of TMPTCID value
    app.raflow.peopleW2UIItems.sort(function(a, b) {
        return a.id - b.id;
    });

    // if there is only one parent rentable then pre-select it for all child rentable
    // otherwise built drop down menu
    if (app.raflow.peopleW2UIItems.length != 1) {
        var item = {id: 0, text: " -- select person -- "};
        pushItem(item, 0);
    } else {
        app.raflow.peopleW2UIItems.forEach(function(item, index) {
            if (item.id === 0) {
                app.raflow.peopleW2UIItems.splice(index, 1);
            }
        });
    }
};

