/* global
    transactantFields, transactantTabs, getSLStringList, getStringListData
*/
"use strict";

window.getTransactantInitRecord = function (BID, BUD) {
    var y = new Date();

    return {
        recid: 0,
        TMPTCID: 0,
        TCID: 0,
        BID: BID,
        BUD: BUD,
        NLID: 0,
        IsRenter: false,
        IsOccupant: true,
        IsGuarantor: false,
        FirstName: "",
        LastName: "",
        MiddleName: "",
        PreferredName: "",
        IsCompany: false,
        CompanyName: "",
        PrimaryEmail: "",
        SecondaryEmail: "",
        WorkPhone: "",
        CellPhone: "",
        Address: "",
        Address2: "",
        City: "",
        State: "",
        PostalCode: "",
        Country: "",
        Website: "",
        LastModTime: y.toISOString(),
        LastModBy: 0,
        Points: 0,
        DateofBirth: "1/1/1900",
        EmergencyContactName: "",
        EmergencyContactAddress: "",
        EmergencyContactTelephone: "",
        EmergencyContactEmail: "",
        AlternateAddress: "",
        EligibleFutureUser: true,
        Industry: "",
        SourceSLSID: 0,
        CreditLimit: 0.00,
        TaxpayorID: "",
        ThirdPartySource: 0,
        GrossIncome: 0,
        SSN: "",
        DriversLicense: "",
        EligibleFuturePayor: true,
        CompanyAddress: "",
        CompanyCity: "",
        CompanyState: "",
        CompanyPostalCode: "",
        CompanyEmail: "",
        CompanyPhone: "",
        CurrentAddress: "",
        CurrentLandLordName: "",
        CurrentLandLordPhoneNo: "",
        CurrentLengthOfResidency: "",
        CurrentReasonForMoving: 0,
        PriorAddress: "",
        PriorLandLordName: "",
        PriorLandLordPhoneNo: "",
        PriorLengthOfResidency: "",
        PriorReasonForMoving: 0,
        Evicted: false,
        EvictedDes: "",
        Convicted: false,
        ConvictedDes: "",
        Bankruptcy: false,
        BankruptcyDes: "",
        Occupation: "",
        DesiredUsageStartDate: "1/1/1900",
        RentableTypePreference: 0,
        FLAGS: 0,
        Approver: 0,
        DeclineReasonSLSID: 0,
        OtherPreferences: "",
        FollowUpDate: "1/1/1900",
        CSAgent: 0,
        OutcomeSLSID: 0,
        CommissionableThirdParty: "",
        Comment: ""
    };
};


window.buildTransactElements = function() {

    app.transactantFields = [
        {field: 'recid',                     type: 'int',       required: false, html: {page: 0, column: 0}},
        {field: 'Address',                   type: 'text',      required: false, html: {page: 0, column: 0}},
        {field: 'Address2',                  type: 'text',      required: false, html: {page: 0, column: 0}},
        {field: 'AlternateAddress',          type: 'text',      required: false, html: {page: 1, column: 0}},
        {field: 'Approver',                  type: 'int',       required: false, html: {page: 3, column: 0}},
        {field: 'Bankruptcy',                type: 'checkbox',  required: false, html: {page: 3, column: 0}},  // have you ever been Declared Bankruptcy
        {field: 'BankruptcyDes',             type: 'text',      required: false, html: {page: 3, column: 0}},
        {field: 'BID',                       type: 'int',       required: false, html: {page: 0, column: 0}},
        {field: 'BUD',                       type: 'list',      required: false, html: {page: 0, column: 0}, options: {items: app.businesses}},
        {field: 'CellPhone',                 type: 'phone',     required: false, html: {page: 0, column: 0}},
        {field: 'City',                      type: 'text',      required: false, html: {page: 0, column: 0}},
        {field: 'Comment',                   type: 'text',      required: false, html: {page: 3, column: 0}},  // In an effort to accommodate you, please advise us of any special needs,
        {field: 'CompanyAddress',            type: 'text',      required: false, html: {page: 3, column: 0}},
        {field: 'CompanyCity',               type: 'text',      required: false, html: {page: 3, column: 0}},
        {field: 'CompanyEmail',              type: 'text',      required: false, html: {page: 3, column: 0}},
        {field: 'CompanyName',               type: 'text',      required: false, html: {page: 0, column: 0}},
        {field: 'CompanyPhone',              type: 'text',      required: false, html: {page: 3, column: 0}},
        {field: 'CompanyPostalCode',         type: 'text',      required: false, html: {page: 3, column: 0}},
        {field: 'CompanyState',              type: 'list',      required: false, html: {page: 3, column: 0}, options: {items: app.usStateAbbr}},
        {field: 'Convicted',                 type: 'checkbox',  required: false, html: {page: 3, column: 0}},  // have you ever been Arrested or convicted of a crime
        {field: 'ConvictedDes',              type: 'text',      required: false, html: {page: 3, column: 0}},
        {field: 'Country',                   type: 'text',      required: false, html: {page: 0, column: 0}},
        {field: 'CreateBy',                  type: 'int',       required: false, html: {page: 0, column: 0}},
        {field: 'CreateTS',                  type: 'time',      required: false, html: {page: 0, column: 0}},
        {field: 'CreditLimit',               type: 'money',     required: false, html: {page: 2, column: 0}},
        {field: 'CSAgent',                   type: 'text',      required: false, html: {page: 3, column: 0}},
        {field: 'CurrentAddress',            type: 'text',      required: false, html: {page: 3, column: 0}},  // Current Address
        {field: 'CurrentLandLordName',       type: 'text',      required: false, html: {page: 3, column: 0}},  // Current landlord's name
        {field: 'CurrentLandLordPhoneNo',    type: 'text',      required: false, html: {page: 3, column: 0}},  // Current landlord's phone number
        {field: 'CurrentLengthOfResidency',  type: 'text',      required: false, html: {page: 3, column: 0}},  // Length of residency at current address
        {field: 'CurrentReasonForMoving',    type: 'list',      required: false, html: {page: 3, column: 0}},  // Reason of moving from current address
        {field: 'CommissionableThirdParty',  type: 'text',      required: false, html: {page: 3, column: 0}},
        {field: 'DateofBirth',               type: 'date',      required: false, html: {page: 1, column: 0}},
        {field: 'DeclineReasonSLSID',        type: 'list',      required: false, html: {page: 3, column: 0}},  // ApplDeny String list
        {field: 'DesiredUsageStartDate',     type: 'date',      required: false, html: {page: 3, column: 0}},
        {field: 'DriversLicense',            type: 'text',      required: false, html: {page: 0, column: 0}},  // Driving licence number of applicants
        {field: 'EligibleFuturePayor',       type: 'checkbox',  required: false, html: {page: 2, column: 0}},
        {field: 'EligibleFutureUser',        type: 'checkbox',  required: false, html: {page: 1, column: 0}},
        {field: 'EmergencyContactAddress',   type: 'text',      required: false, html: {page: 1, column: 0}},
        {field: 'EmergencyContactName',      type: 'text',      required: false, html: {page: 1, column: 0}},
        {field: 'EmergencyContactTelephone', type: 'text',      required: false, html: {page: 1, column: 0}},
        {field: 'EmergencyContactEmail',     type: 'text',      required: false, html: {page: 1, column: 0}},
        {field: 'Evicted',                   type: 'checkbox',  required: false, html: {page: 3, column: 0}},  // have you ever been Evicted
        {field: 'EvictedDes',                type: 'text',      required: false, html: {page: 3, column: 0}},
        {field: 'FirstName',                 type: 'text',      required: false, html: {page: 0, column: 0}},
        {field: 'FLAGS',                     type: 'int',       required: false, html: {page: 3, column: 0}},
        {field: 'FollowUpDate',              type: 'date',      required: false, html: {page: 3, column: 0}},
        {field: 'GrossIncome',               type: 'money',     required: false, html: {page: 2, column: 0}},
        {field: 'Industry',                  type: 'text',      required: false, html: {page: 1, column: 0}},
        {field: 'IsCompany',                 type: 'checkbox',  required: true,  html: {page: 0, column: 0}},
        {field: 'IsGuarantor',               type: 'checkbox',  required: false, html: {page: 0, column: 0}},  // responsible for making sure all rent is paid
        {field: 'IsOccupant',                type: 'checkbox',  required: false, html: {page: 0, column: 0}},  // will reside in and/or use the items rented
        {field: 'IsRenter',                  type: 'checkbox',  required: false, html: {page: 0, column: 0}},  // will be responsible for paying rent
        {field: 'LastModBy',                 type: 'int',       required: false, html: {page: 0, column: 0}},
        {field: 'LastModTime',               type: 'time',      required: false, html: {page: 0, column: 0}},
        {field: 'LastName',                  type: 'text',      required: false, html: {page: 0, column: 0}},
        {field: 'MiddleName',                type: 'text',      required: false, html: {page: 0, column: 0}},
        {field: 'NLID',                      type: 'int',       required: false, html: {page: 0, column: 0}},
        {field: 'Occupation',                type: 'text',      required: false, html: {page: 3, column: 0}},
        {field: 'OtherPreferences',          type: 'text',      required: false, html: {page: 3, column: 0}},
        {field: 'OutcomeSLSID',              type: 'text',      required: false, html: {page: 3, column: 0}},
        {field: 'Points',                    type: 'text',      required: false, html: {page: 1, column: 0}},
        {field: 'PostalCode',                type: 'text',      required: false, html: {page: 0, column: 0}},
        {field: 'PreferredName',             type: 'text',      required: false, html: {page: 0, column: 0}},
        {field: 'PrimaryEmail',              type: 'email',     required: false, html: {page: 0, column: 0}},
        {field: 'PriorAddress',              type: 'text',      required: false, html: {page: 3, column: 0}},  // Prior Address
        {field: 'PriorLandLordName',         type: 'text',      required: false, html: {page: 3, column: 0}},  // Prior landlord's name
        {field: 'PriorLandLordPhoneNo',      type: 'text',      required: false, html: {page: 3, column: 0}},  // Prior landlord's phone number
        {field: 'PriorLengthOfResidency',    type: 'text',      required: false, html: {page: 3, column: 0}},  // Length of residency at Prior address
        {field: 'PriorReasonForMoving',      type: 'list',      required: false, html: {page: 3, column: 0}},  // Reason of moving from Prior address
        {field: 'RentableTypePreference',    type: 'text',      required: false, html: {page: 3, column: 0}},
        {field: 'SecondaryEmail',            type: 'email',     required: false, html: {page: 0, column: 0}},
        {field: 'SourceSLSID',               type: 'list',      required: false, html: {page: 1, column: 0}}, // "HowFound" string list
        {field: 'SSN',                       type: 'text',      required: false, html: {page: 0, column: 0}},  // Social security number of applicants
        {field: 'State',                     type: 'list',      required: false, html: {page: 0, column: 0}, options: {items: app.usStateAbbr}},
        {field: 'TaxpayorID',                type: 'text',      required: false, html: {page: 2, column: 0}},
        {field: 'TCID',                      type: 'int',       required: false, html: {page: 0, column: 0}},
        {field: 'ThirdPartySource',          type: 'text',      required: false, html: {page: 2, column: 0}},
        {field: 'TMPTCID',                   type: 'int',       required: true,  html: {page: 0, column: 0}},
        {field: 'Website',                   type: 'text',      required: false, html: {page: 0, column: 0}},
        {field: 'WorkPhone',                 type: 'phone',     required: false, html: {page: 0, column: 0}}
    ];

    app.transactantTabs = [
        {id: 'tab1', caption: app.sTransactant},
        {id: 'tab2', caption: app.sUser},
        {id: 'tab3', caption: app.sPayor},
        {id: 'tab4', caption: app.sProspect}
    ];

    //------------------------------------------------------------------------
    //          transactantsGrid
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'transactantsGrid',
        url: '/v1/transactants',
        multiSelect: false,
        show: {
            header: false,
            toolbar: true,
            footer: true,
            toolbarAdd: true,
            lineNumbers: false,
            selectColumn: false,
            expandColumn: false
        },
        columns: [
            {field: 'TCID',         caption: "TCID",          size: '50px',  sortable: true, style: 'text-align: right', hidden: false},
            {field: 'FirstName',    caption: "First Name",    size: '125px', sortable: true, hidden: false},
            {field: 'MiddleName',   caption: "Middle Name",   size: '20px',  sortable: true, hidden: true},
            {field: 'LastName',     caption: "Last Name",     size: '125px', sortable: true, hidden: false,
                render: function (record) {
                    var s = '';
                    if (typeof record === "undefined") {
                        return;
                    }
                    if (!record.IsCompany) {
                        s += '<span style="color:#999;font-size:16px"><i class="far fa-handshake" aria-hidden="true"></i></span>';
                    }
                    return s + ' ' + record.LastName;
                }
            },
            {field: 'CompanyName',  caption: "Company Name",  size: '125px', sortable: true, hidden: false,
                render: function (record) {
                    var s = '';
                    if (typeof record === "undefined") {
                        return;
                    }
                    if (record.IsCompany) {
                        s += '<span style="color:#999;font-size:16px"><i class="far fa-handshake" aria-hidden="true"></i></span>';
                    }
                    return s + ' ' + record.CompanyName;
                }
            },
            {field: 'PrimaryEmail', caption: "Primary Email", size: '175px', sortable: true, hidden: false},
            {field: 'CellPhone',    caption: "Cell Phone",    size: '100px', sortable: true, hidden: false},
            {field: 'WorkPhone',    caption: "Work Phone",    size: '100px', sortable: true, hidden: false},
        ],
        onRefresh: function(event) {
            event.onComplete = function() {
                if (app.active_grid == this.name) {
                    if (app.new_form_rec) {
                        this.selectNone();
                    }
                    else{
                        this.select(app.last.grid_sel_recid);
                    }
                }
            };
        },
        onClick: function(event) {
            event.onComplete = function () {
                var yes_args = [this, event.recid],
                    no_args = [this],
                    BID = getCurrentBID(),
                    BUD = getBUDfromBID(BID),
                    no_callBack = function(grid) {
                        grid.select(app.last.grid_sel_recid);
                        return false;
                    },
                    yes_callBack = function(grid, recid) {
                        app.last.grid_sel_recid = parseInt(recid);
                        // keep highlighting current row in any case
                        grid.select(app.last.grid_sel_recid);
                        var rec = grid.get(recid);

                        // get stringListData for list fields
                        getStringListData(BID, BUD).fail(function (data) {
                            this.message(data.message);
                        });

                        setToForm('transactantForm', '/v1/person/' + rec.BID + '/' + rec.TCID, 700, true);
                    };

                // warn user if form content has been changed
                form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
             };
        },
        onAdd: function(/*event*/) {
            var yes_args = [this],
                no_callBack = function() { return false; },
                yes_callBack = function(grid) {
                    // reset it
                    app.last.grid_sel_recid = -1;
                    grid.selectNone();

                    // insert an empty record....
                    var x = getCurrentBusiness();
                    var BID=parseInt(x.value);
                    var BUD = getBUDfromBID(BID);

                    var record = getTransactantInitRecord(BID, BUD);
                    w2ui.transactantForm.record = record;

                    // get stringListData for list fields
                    getStringListData(BID, BUD).fail(function (data) {
                        this.message(data.message);
                    });

                    w2ui.transactantForm.refresh();
                    setToForm('transactantForm', '/v1/person/' + BID + '/0', 700);
                };

            // warn user if form content has been changed
            form_dirty_alert(yes_callBack, no_callBack, yes_args);
        }
    });


    //------------------------------------------------------------------------
    //          transactantForm
    //------------------------------------------------------------------------
    $().w2form({
        name: 'transactantForm',
        style: 'border: 0px; background-color: transparent;',
        header: app.sTransactant + ' Detail',
        url: '/v1/person',
        formURL: '/webclient/html/formtc.html',
        fields: app.transactantFields,
        tabs: app.transactantTabs,
        toolbar: {
            items: [
                { id: 'btnNotes', type: 'button', icon: 'far fa-sticky-note' },
                { id: 'bt3', type: 'spacer' },
                { id: 'btnClose', type: 'button', icon: 'fas fa-times' },
            ],
            onClick: function (event) {
                if (event.target == 'btnClose') {
                    var no_callBack = function() { return false; },
                        yes_callBack = function() {
                            w2ui.toplayout.hide('right',true);
                            w2ui.transactantsGrid.render();
                        };
                    form_dirty_alert(yes_callBack, no_callBack);
                }
                if (event.target == 'btnNotes') {
                    notesPopUp();
                }
            },
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
        },
        actions: {
            save: function () {
                var tgrid = w2ui.transactantsGrid;
                console.log('before: tgrid.getSelection() = ' + tgrid.getSelection() );
                tgrid.selectNone();
                // clean dirty flag of form
                app.form_is_dirty = false;
                // clear the grid select recid
                app.last.grid_sel_recid  =-1;
                this.save({}, function (data) {
                    if (data.status == 'error') {
                        console.log('ERROR: '+ data.message);
                        return;
                    }
                    w2ui.toplayout.hide('right',true);
                    tgrid.render();
                });
            },
            saveadd: function() {
                var f = this,
                    grid = w2ui.transactantsGrid,
                    x = getCurrentBusiness(),
                    r = f.record,
                    BID=parseInt(x.value),
                    BUD=getBUDfromBID(BID);

                // clean dirty flag of form
                app.form_is_dirty = false;
                // clear the grid select recid
                app.last.grid_sel_recid  =-1;

                // select none if you're going to add new record
                grid.selectNone();

                f.save({}, function (data) {
                    if (data.status == 'error') {
                        console.log('ERROR: '+ data.message);
                        return;
                    }

                    // JUST RENDER THE GRID ONLY
                    grid.render();

                    // add new empty record and just refresh the form, don't need to do CLEAR form
                    var record = getTransactantInitRecord(BID, BUD);

                    f.record = record;
                    f.header = "Edit Transactant (new)"; // have to provide header here, otherwise have to call refresh method twice to get this change in form
                    f.url = '/v1/person/' + BID+'/0';
                    f.refresh();
                });
            },
            delete: function(/*target, data*/) {
                var form = this;
                w2confirm(delete_confirm_options)
                .yes(function() {
                    var tgrid = w2ui.transactantsGrid;
                    var params = {cmd: 'delete', formname: form.name, TCID: form.record.TCID };
                    var dat = JSON.stringify(params);

                    // delete Transactant request
                    $.post(form.url, dat, null, "json")
                    .done(function(data) {
                        if (data.status === "error") {
                            form.error(w2utils.lang(data.message));
                            return;
                        }
                        w2ui.toplayout.hide('right',true);
                        tgrid.remove(app.last.grid_sel_recid);
                        tgrid.render();
                    })
                    .fail(function(/*data*/){
                        form.error("Delete Transactant failed.");
                        return;
                    });
                })
                .no(function() {
                    return;
                });
            }
        },
        onRefresh: function(event) {
            event.onComplete = function() {
                var f = this,
                    r = f.record,
                    header="",
                    BID = getCurrentBID(),
                    BUD = getBUDfromBID(BID);

                // custom header
                if (r.TCID) {
                    if (f.original.IsCompany) {
                        header = "Edit Transactant - {0} ({1})".format(r.CompanyName, r.TCID);
                    } else {
                        header = "Edit Transactant - {0} {1} ({2})".format(r.FirstName, r.LastName, r.TCID);
                    }
                } else {
                    header = "Edit Transactant ({0})".format("new");
                }

                formRefreshCallBack(f, "TCID", header);

                // Hide Transanctant role checkboxes
                f.get("IsRenter").hidden = true;
                f.get("IsGuarantor").hidden = true;
                f.get("IsGuarantor").hidden = true;
                $("div[name=transanctant-role-tile]").hide();

                f.get('SourceSLSID').options.items = getSLStringList(BID, "HowFound");
                f.get('DeclineReasonSLSID').options.items = getSLStringList(BID, "ApplDeny");
                f.get('CurrentReasonForMoving').options.items = getSLStringList(BID, "WhyLeaving");
                f.get('PriorReasonForMoving').options.items = getSLStringList(BID, "WhyLeaving");
            };
        },
        onChange: function(event) {
            event.onComplete = function() {
                // Enable/Disable checkbox description text area
                $("#EvictedDes").prop("disabled", !this.record.Evicted);
                $("#ConvictedDes").prop("disabled", !this.record.Convicted);
                $("#BankruptcyDes").prop("disabled", !this.record.Bankruptcy);

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
        onSubmit: function(target, data){
            delete data.postData.record.LastModTime;
            delete data.postData.record.LastModBy;
            delete data.postData.record.CreateTS;
            delete data.postData.record.CreateBy;
            // server request form data
            getFormSubmitData(data.postData.record);
            data.postData.record.IsCompany = int_to_bool(data.postData.record.IsCompany);
            data.postData.record.EligibleFutureUser = int_to_bool(data.postData.record.EligibleFutureUser);
            data.postData.record.EligibleFuturePayor = int_to_bool(data.postData.record.EligibleFuturePayor);
        }
    });

};

//-----------------------------------------------------------------------------
// getStringListData - return the promise object of request to get latest
//                           string list for given BID.
//                           It updates the "app.ReceiptRules" variable for requested BUD
// @params  - BID : Business ID (expected current one)
//          - BUD : Business Unit Designation
// @return  - promise object from $.get
//-----------------------------------------------------------------------------
window.getStringListData = function (BID, BUD) {
    // if not BUD in app.ReceiptRules then initialize it with blank list
    if (!(BUD in app.StringList)) {
        app.StringList[BUD] = [];
    }

    // return promise
    return $.get("/v1/uival/" + BID + "/app.Applicants", null, null, "json").done(function(data) {
        // if it doesn't meet this condition, then save the data
        if (!('status' in data && data.status !== "success")) {
            app.StringList[BUD] = data;
        }
    });
};

// getSLStringList - It provide string list of `SLName`
window.getSLStringList = function(BID, SLName){
    var BUD = getBUDfromBID(BID);
    app[SLName] = [];
    app.StringList[BUD].forEach(function (SLObject) {
        if(SLObject.Name === SLName){
            var defaultItem;
            switch (SLName){
                case "HowFound":
                    defaultItem = {id: 0, text: " -- Select Source SLSID -- "};
                    break;
                case "WhyLeaving":
                    defaultItem = {id: 0, text: " -- Select reason for leaving -- "};
                    break;
                case "ApplDeny":
                    defaultItem = {id: 0, text: " -- Select DeclineReasonSLSID -- "};
                    break;
                default:
                    console.log("SLName doesn't exists");
            }
            app[SLName].push(defaultItem);
            for(var index = 0 ; index < SLObject.S.length ; index++){
                app[SLName].push({id: SLObject.S[index].SLSID, text: SLObject.S[index].Value});
            }
        }
    });
    return app[SLName];
};

// updateRATransactantFormCheckboxes
// Convert checkboxes w2ui int(1/0) value to bool(true/false)
window.updateRATransactantFormCheckboxes = function (record) {
    record.IsRenter = int_to_bool(record.IsRenter);
    record.IsOccupant = int_to_bool(record.IsOccupant);
    record.IsGuarantor = int_to_bool(record.IsGuarantor);
    record.IsCompany = int_to_bool(record.IsCompany);
    record.Evicted = int_to_bool(record.Evicted);
    record.Bankruptcy = int_to_bool(record.Bankruptcy);
    record.Convicted = int_to_bool(record.Convicted);
    record.EligibleFuturePayor = int_to_bool(record.EligibleFuturePayor);
    record.EligibleFutureUser = int_to_bool(record.EligibleFutureUser);
};
