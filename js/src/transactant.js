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
                if (record.IsCompany === 0) {
                    s += '<span style="color:#999;font-size:16px"><i class="fa fa-handshake-o" aria-hidden="true"></i></span>';
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
                if (record.IsCompany > 0) {
                    s += '<span style="color:#999;font-size:16px"><i class="fa fa-handshake-o" aria-hidden="true"></i></span>';
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
                no_callBack = function(grid) {
                    grid.select(app.last.grid_sel_recid);
                    return false;
                },
                yes_callBack = function(grid, recid) {
                    app.last.grid_sel_recid = parseInt(recid);
                    // keep highlighting current row in any case
                    grid.select(app.last.grid_sel_recid);
                    var rec = grid.get(recid);
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
                var y = new Date();

                var record = {
                    recid: 0,
                    FirstName: "",
                    LastName: "",
                    MiddleName: "",
                    PreferredName: "",
                    PrimaryEmail: "",
                    TCID: 0,
                    BID: BID,
                    BUD: BUD,
                    NLID: 0,
                    CompanyName: "",
                    IsCompany: 0,
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
                    EmergencyEmail: "",
                    AlternateAddress: "",
                    EligibleFutureUser: "yes",
                    Industry: "",
                    SourceSLSID: 0,
                    CreditLimit: 0.00,
                    TaxpayorID: "",
                    AccountRep: 0,
                    EligibleFuturePayor: "yes",
                    EmployerName: "",
                    EmployerStreetAddress: "",
                    EmployerCity: "",
                    EmployerState: "",
                    EmployerPostalCode: "",
                    EmployerEmail: "",
                    EmployerPhone: "",
                    Occupation: "",
                    ApplicationFee: 0.00,
                    DesiredUsageStartDate: "1/1/1900",
                    RentableTypePreference: 0,
                    FLAGS: 0,
                    Approver: 0,
                    DeclineReasonSLSID: 0,
                    OtherPreferences: "",
                    FollowUpDate: "1/1/1900",
                    CSAgent: 0,
                    OutcomeSLSID: 0,
                    FloatingDeposit: 0.00,
                    RAID: 0,
                };
                w2ui.transactantForm.record = record;
                w2ui.transactantForm.refresh();
                setToForm('transactantForm', '/v1/person/' + BID + '/0', 700);
            };

        // warn user if form content has been changed
        form_dirty_alert(yes_callBack, no_callBack, yes_args);
    },
});
