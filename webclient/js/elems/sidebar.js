/*global
    w2ui, app, console, $, plural, switchToGrid, showReport, form_dirty_alert,
*/
"use strict";
function buildSidebar() {
    //------------------------------------------------------------------------
    //          sidebarL1
    //------------------------------------------------------------------------
    w2ui.toplayout.content('left',$().w2sidebar({
        name: 'sidebarL1',
        nodes: [
            { id: 'workflowreceipts', text: plural(app.sAssessment)+' / '+plural(app.sReceipt), img: 'icon-folder', expanded: true, group: true,
                nodes: [
                        { id: 'asms',         text: 'Assess Charges',                icon: 'fa fa-star-o',      hint: plural(app.sAssessment) },
                        { id: 'receipts',     text: 'Receive '+plural(app.sReceipt), icon: 'fa fa-star',        hint: plural(app.sReceipt) },
                        { id: 'expense',      text: plural(app.sExpense),            icon: 'fa fa-minus-circle',hint: plural(app.sExpense) },
                        { id: 'deposit',      text: 'Deposits',                      icon: 'fa fa-plus-circle', hint: 'Make Deposit' },
                        { id: 'allocfunds',   text: 'Apply '+plural(app.sReceipt),   icon: 'fa fa-check-circle-o' },
                        // { id: 'gdssvcs',      text: 'Goods & Services',              icon: 'fa fa-coffee' },
               ]
            },
            { id: 'rentagr', text: plural(app.sRentalAgreement), img: 'icon-folder', expanded: true, group: true,
                nodes: [
                        { id: 'rentalagrs',   text: plural(app.sRentalAgreement),    icon: 'fa fa-certificate', hint: 'Rental Agreements' },
                        { id: 'transactants', text: plural(app.sTransactant),        icon: 'fa fa-users' },
                        // { id: 'assignrnt',    text: 'Assign A ' + app.sRentable,     icon: 'fa fa-check-square-o' },
                        // { id: 'movein',       text: app.sTransactant + ' Arrival',   icon: 'fa fa-sign-in' },
                        // { id: 'moveout',      text: app.sTransactant + ' Departure', icon: 'fa fa-sign-out' },
                        // { id: 'updatera',     text: 'Extend ' + app.sRentalAgreement,icon: 'fa fa-pencil' },
                ]
            },
            { id: 'collections', text: 'Collections', img: 'icon-folder', expanded: true, group: true,
                nodes: [
                        { id: 'rr',           text: 'Rent Roll',                     icon: 'fa fa-line-chart',   hint: 'Rent Roll' },
                        // { id: 'dlnq',         text: 'Delinquency Analysis',          icon: 'fa fa-pie-chart',   hint: 'Delinquency Analysis' },
                        { id: 'stmt',         text: 'RA Statements',                 icon: 'fa fa-star-half-o', hint: 'Rental Agreement Statements' },
                        { id: 'payorstmt',    text: 'Payor Statements',              icon: 'fa fa-star-half-o fa-flip-horizontal', hint: 'Payor Statements' },
                        // { id: 'prepnotice',   text: 'Prepare Notices',               icon: 'fa fa-file-text-o', hint: 'Prepare Notices' },
                ]
            },
            { id: 'acct', text: 'Accounting', img: 'icon-folder', expanded: false, group: true,
                nodes: [
                        { id: 'close',       text: 'Close Period',                   icon: 'fa fa-toggle-down', hint: 'Close Period' },
                        { id: 'adjust',      text: 'Adjust Closed Period',           icon: 'fa fa-reply',       hint: 'Adjust Closed Period' },
                ]
            },
            { id: 'facilities', text: 'Facilities Management', img: 'icon-folder', expanded: false, group: true,
                nodes: [
                        { id: 'svcreq',      text: 'Create Service Request',         icon: 'fa fa-square-o',       hint: 'Create Service Request' },
                        { id: 'svcreqcmp',   text: 'Complete Service Request',       icon: 'fa fa-check-square-o', hint: 'Complete Service Request' },
                        { id: 'housekpg',    text: 'Housekeeping',                   icon: 'fa fa-home',           hint: 'Housekeeping' },
                        { id: 'prvmaint',    text: 'Preventative Maintenance',       icon: 'fa fa-wrench',         hint: 'Preventative Maintenance' },
                        { id: 'invntory',    text: 'Inventory',                      icon: 'fa fa-shopping-cart',  hint: 'Preventative Maintenance' },
                ]
            },
            { id: 'reports', text: 'Reports', img: 'icon-folder', expanded: false, group: true,
                nodes: [
                       //{ id: 'RPTasmrpt',     text: 'Assessments',                     icon: 'fa fa-file-text-o' },
                       //{ id: 'RPTb',          text: 'Business Units',                  icon: 'fa fa-file-text-o' },
                       { id: 'RPTcoa',          text: 'Chart Of Accounts',               icon: 'fa fa-file-text-o' },
                       //{ id: 'RPTdpm',        text: 'Deposit Methods',                 icon: 'fa fa-file-text-o' },
                       //{ id: 'RPTdep',        text: 'Depository Accounts',             icon: 'fa fa-file-text-o' },
                       { id: 'RPTdelinq',       text: 'Delinquency',                     icon: 'fa fa-file-text-o' },
                       { id: 'RPTgsr',          text: 'GSR',                             icon: 'fa fa-file-text-o' },
                       { id: 'RPTj',            text: 'Journal',                         icon: 'fa fa-file-text-o' },
                       { id: 'RPTl',            text: 'Ledger',                          icon: 'fa fa-file-text-o' },
                       { id: 'RPTla',           text: 'Ledger Activity',                 icon: 'fa fa-file-text-o' },
                       { id: 'RPTpeople',       text: app.sTransactant,                  icon: 'fa fa-file-text-o' },
                       //{ id: 'RPTpmt',        text: 'Payment Types',                   icon: 'fa fa-file-text-o' },
                       //{ id: 'RPTrcpt',       text: 'Receipts',                        icon: 'fa fa-file-text-o' },
                       { id: 'RPTrcbt',         text: app.sRentable+' Type Counts',      icon: 'fa fa-file-text-o' },
                       { id: 'RPTr',            text: plural(app.sRentable),             icon: 'fa fa-file-text-o' },
                       { id: 'RPTra',           text: plural(app.sRentalAgreement),      icon: 'fa fa-file-text-o' },
                       { id: 'RPTrat',          text: app.sRentalAgreement+' Templates', icon: 'fa fa-file-text-o' },
                       { id: 'RPTrt',           text: app.sRentable+' Types',            icon: 'fa fa-file-text-o' },
                       { id: 'RPTrr',           text: 'RentRoll',                        icon: 'fa fa-file-text-o' },
                       //{ id: 'RPTstatements', text: 'Statements',                      icon: 'fa fa-file-text-o' },
                       //{ id: 'RPTsl',         text: 'String Lists',                    icon: 'fa fa-file-text-o' },
                       { id: 'RPTtb',           text: 'Trial Balance',                   icon: 'fa fa-file-text-o' },//
                ]
            },
            { id: 'setup', text: 'Setup', img: 'icon-wrench', expanded: true, group: true,
                nodes: [
                        { id: 'accounts',    text: 'Chart Of Accounts',                icon: 'fa fa-list' },
                        { id: 'pmts',        text: 'Payment Types',                    icon: 'fa fa-credit-card' },
                        { id: 'dep',         text: 'Depository Accounts',              icon: 'fa fa-university' },
                        { id: 'depmeth',     text: 'Deposit Methods',                  icon: 'fa fa-envelope-o' },
                        { id: 'ars',         text: 'Account Rules',                    icon: 'fa fa-cogs' },
                        { id: 'rt',          text: plural(app.sRentableType),          icon: 'fa fa-asterisk', hint: 'Rentable Types' },
                        { id: 'rentables',   text: plural(app.sRentable),              icon: 'fa fa-cube' },
                        // { id: 'changeRT',    text: 'Change ' + app.sRentable +' Type', icon: 'fa fa-refresh' },
                        // { id: 'permissions', text: 'Permissions',                      icon: 'fa fa-thumbs-o-up' },
                ]
            },
            { id: 'admin', text: 'Administrator', img: 'icon-wrench', expanded: false, group: true,
                nodes: [
                        { id: 'about',  text: 'Product Info',        icon: 'fa fa-info' },
                        { id: 'tws',    text: 'Timed Work Schedule', icon: 'fa fa-calendar' },
                        { id: 'ledgers',text: 'Ledgers',             icon: 'fa fa-book' },
                ]
            },
        ],
        onExpand: function(event) {
            //var x = getCurrentBusiness();
            //console.log('current biz = ' + x.value + '  name = ' + x.name );
            switch (event.target) {
                case 'reports':
                    var w = w2ui.reportslayout;
                    w2ui.toplayout.content('main', w);
                    w2ui.toplayout.hide('right',true);
                    break;
            }

        },
        onFlat: function (event) {
            console.log('event.goFlat = ' + event.goFlat );
            $('#sidebarL1').css('width', (event.goFlat ? '35px' : ''+app.sidebarWidth+'px'));
        },
        onClick: function (event) {
            console.log('event.target = ' + event.target);
            var target = event.target;
            var no_callBack = function(target) {
                console.log("sidebar active form dirty - no callBack", target);
                return false;
            },
            yes_callBack = function(target) {
                console.log("sidebar active form dirty - yes callBack", target);
                switch(target) {
                    case 'accounts':
                    case 'rt':
                    case 'rentables':
                    case 'transactants':
                    case 'rentalagrs':
                    case 'receipts':
                    case 'asms':
                    case 'pmts':
                    case 'dep':
                    case 'ars':
                    case 'tws':
                    case 'ledgers':
                    case 'stmt':
                    case 'depmeth':
                    case 'allocfunds':
                    case 'deposit':
                    case 'expense':
                    case 'payorstmt':
                    case 'rr':
                        // w2ui.sidebarL1.collapse('reports'); // close reports when jumping to a main view
                        switchToGrid(target);
                        break;
                    case 'goRatePlan':
                        // w2ui.sidebarL1.collapse('reports'); // close reports when jumping to a main view
                        w2ui.toplayout.content('main', '<h1>Sorry :-(</h1><h2>Rate Plan...  Not Available</h2><h3>But coming soon!</h3>');
                        w2ui.toplayout.hide('right',true);
                        break;
                    case 'goServiceMenu':
                        // w2ui.sidebarL1.collapse('reports'); // close reports when jumping to a main view
                        w2ui.toplayout.content('main', '<h1>Sorry :-(</h1><h2>Service Menu...  Not Available</h2><h3>But coming soon!</h3>');
                        w2ui.toplayout.hide('right',true);
                        break;
                    case 'mkdep':
                        // w2ui.sidebarL1.collapse('reports'); // close reports when jumping to a main view
                        w2ui.toplayout.load('main', '/webclient/html/formmkdep.html');
                        w2ui.toplayout.hide('right',true);
                        break;
                     // case 'reversepmt':
                     //     w2ui.sidebarL1.collapse('reports'); // close reports when jumping to a main view
                     //     w2ui.toplayout.load('main', '/webclient/html/formrevrcpt.html');
                     //     w2ui.toplayout.hide('right',true);
                     //     break;
                    case 'gdssvcs':
                        // w2ui.sidebarL1.collapse('reports'); // close reports when jumping to a main view
                        w2ui.toplayout.load('main', '/webclient/html/formgas.html');
                        w2ui.toplayout.hide('right',true);
                        break;
                    case 'assignrnt':
                        // w2ui.sidebarL1.collapse('reports'); // close reports when jumping to a main view
                        w2ui.toplayout.load('main', '/webclient/html/formaar.html');
                        w2ui.toplayout.hide('right',true);
                        break;
                    case 'movein':
                        // w2ui.sidebarL1.collapse('reports'); // close reports when jumping to a main view
                        w2ui.toplayout.load('main', '/webclient/html/formmvin.html');
                        w2ui.toplayout.hide('right',true);
                        break;
                    case 'moveout':
                        // w2ui.sidebarL1.collapse('reports'); // close reports when jumping to a main view
                        w2ui.toplayout.load('main', '/webclient/html/formmvout.html');
                        w2ui.toplayout.hide('right',true);
                        break;
                    case 'updatera':
                        // w2ui.sidebarL1.collapse('reports'); // close reports when jumping to a main view
                        w2ui.toplayout.load('main', '/webclient/html/formraextend.html');
                        w2ui.toplayout.hide('right',true);
                        break;
                    case 'dlnq':
                        // w2ui.sidebarL1.collapse('reports'); // close reports when jumping to a main view
                        w2ui.toplayout.load('main', '/webclient/html/formdlnq.html');
                        w2ui.toplayout.hide('right',true);
                        break;
                    case 'prepnotice':
                        // w2ui.sidebarL1.collapse('reports'); // close reports when jumping to a main view
                        w2ui.toplayout.load('main', '/webclient/html/formprepnotice.html');
                        w2ui.toplayout.hide('right',true);
                        break;
                    case 'close':
                        // w2ui.sidebarL1.collapse('reports'); // close reports when jumping to a main view
                        w2ui.toplayout.load('main', '/webclient/html/formclose.html');
                        w2ui.toplayout.hide('right',true);
                        break;
                    case 'adjust':
                        // w2ui.sidebarL1.collapse('reports'); // close reports when jumping to a main view
                        w2ui.toplayout.load('main', '/webclient/html/formadjust.html');
                        w2ui.toplayout.hide('right',true);
                        break;
                    case 'svcreq':
                        // w2ui.sidebarL1.collapse('reports'); // close reports when jumping to a main view
                        w2ui.toplayout.load('main', '/webclient/html/formcsr.html');
                        w2ui.toplayout.hide('right',true);
                        break;
                    case 'svcreqcmp':
                        // w2ui.sidebarL1.collapse('reports'); // close reports when jumping to a main view
                        w2ui.toplayout.load('main', '/webclient/html/formcmplsr.html');
                        w2ui.toplayout.hide('right',true);
                        break;
                    case 'housekpg':
                        // w2ui.sidebarL1.collapse('reports'); // close reports when jumping to a main view
                        w2ui.toplayout.load('main', '/webclient/html/formhskp.html');
                        w2ui.toplayout.hide('right',true);
                        break;
                    case 'prvmaint':
                        // w2ui.sidebarL1.collapse('reports'); // close reports when jumping to a main view
                        w2ui.toplayout.load('main', '/webclient/html/formprvm.html');
                        w2ui.toplayout.hide('right',true);
                        break;
                    case 'invntory':
                        // w2ui.sidebarL1.collapse('reports'); // close reports when jumping to a main view
                        w2ui.toplayout.load('main', '/webclient/html/forminv.html');
                        w2ui.toplayout.hide('right',true);
                        break;
                    case 'changeRT':
                        // w2ui.sidebarL1.collapse('reports'); // close reports when jumping to a main view
                        w2ui.toplayout.load('main', '/webclient/html/formchgrt.html');
                        w2ui.toplayout.hide('right',true);
                        break;
                    case 'permissions':
                        // w2ui.sidebarL1.collapse('reports'); // close reports when jumping to a main view
                        w2ui.toplayout.load('main', '/webclient/html/formperm.html');
                        w2ui.toplayout.hide('right',true);
                        break;
                    case 'RPTasmrpt':
                    case 'RPTb':
                    case 'RPTcoa':
                    case 'RPTdelinq':
                    case 'RPTdep':
                    case 'RPTdpm':
                    case 'RPTgsr':
                    case 'RPTj':
                    case 'RPTl':
                    case 'RPTla':
                    case 'RPTpeople':
                    case 'RPTpmt':
                    case 'RPTr':
                    case 'RPTra':
                    case 'RPTrat':
                    case 'RPTrcbt':  // rentable count by type
                    case 'RPTrcpt':
                    case 'RPTrr':
                    case 'RPTrt':
                    case 'RPTsl':
                    case 'RPTstatements':
                    case 'RPTtb':
                        showReport(target);
                        app.last.report = target;
                        break;
                    case 'about':
                        // w2ui.sidebarL1.collapse('reports'); // close reports when jumping to a main view
                        w2ui.toplayout.load('main', '/webclient/html/about.html');
                        w2ui.toplayout.hide('right',true);
                        getAboutInfo();
                        break;
                    default:
                        console.log('unhandled event target: ' + target);
                }
            };

            // warn user if form has been changed
            // also here we need to bind current event to both function for
            // use of event inside those function
            var yes_cb_args = [target],
                nb_cb_args = [target];
            form_dirty_alert(yes_callBack, no_callBack, yes_cb_args, nb_cb_args);
        }
    }));
}

//---------------------------------------------------------------------------------
// getAboutInfo - contacts the server to get info about its version, and updates
//          the version (about) html page
//
// @params  <none>
// @retunrs <none>
//---------------------------------------------------------------------------------
function getAboutInfo() {
    $.get('/v1/version/')
    .done( function(data) {
        if (typeof data == 'string') {  // it's weird, a successful data add gets parsed as an object, an error message does not
            document.getElementById("appVer").innerHTML = data;
            //w2ui.toplayout.refresh('main');
        } else {
            console.log('received response of type ' + typeof data + ' : ' + data);
        }
    })
    .fail( function() {
        console.log('Error getting /v1/version/');
    });
}
