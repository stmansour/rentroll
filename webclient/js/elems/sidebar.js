/*global
    w2ui, app, console, $, plural, switchToGrid, showReport, form_dirty_alert, loginPopupOptions, getAboutInfo
*/
"use strict";


// buildSidebar creates the w2ui sidebar element for the Roller interface.
//
// INPUTS:
//  uitype - 0 means build the full roller interface
//           1 means build the Receipt-only interface
//----------------------------------------------------------------------------
window.buildSidebar = function(flag) {
    var sbdata;
    if (flag === 0) {
        sbdata = {
            name: 'sidebarL1',
            nodes: [
                { id: 'workflowreceipts', text: plural(app.sAssessment)+' / '+plural(app.sReceipt), img: 'icon-folder', expanded: true, group: true,
                    nodes: [
                            { id: 'asms',         text: 'Assess Charges',                icon: 'far fa-star',      hint: plural(app.sAssessment) },
                            { id: 'receipts',     text: 'Tendered Payment ' + app.sReceipt, icon: 'fas fa-star',        hint: plural(app.sReceipt) },
                            { id: 'expense',      text: plural(app.sExpense),            icon: 'fas fa-minus-circle',hint: plural(app.sExpense) },
                            { id: 'deposit',      text: 'Deposits',                      icon: 'fas fa-plus-circle', hint: 'Make Deposit' },
                            { id: 'allocfunds',   text: 'Apply '+plural(app.sReceipt),   icon: 'far fa-check-circle' },
                            // { id: 'gdssvcs',      text: 'Goods & Services',              icon: 'fas fa-coffee' },
                   ]
                },
                { id: 'rentagr', text: plural(app.sRentalAgreement), img: 'icon-folder', expanded: true, group: true,
                    nodes: [
                            { id: 'rentalagrs',   text: plural(app.sRentalAgreement),    icon: 'fas fa-certificate', hint: 'Rental Agreements' },
                            { id: 'transactants', text: plural(app.sTransactant),        icon: 'fas fa-users' },
                            // { id: 'assignrnt',    text: 'Assign A ' + app.sRentable,     icon: 'far fa-check-square' },
                            // { id: 'movein',       text: app.sTransactant + ' Arrival',   icon: 'fas fa-sign-in-alt' },
                            // { id: 'moveout',      text: app.sTransactant + ' Departure', icon: 'fas fa-sign-out-alt' },
                            // { id: 'updatera',     text: 'Extend ' + app.sRentalAgreement,icon: 'fas fa-pencil-alt' },
                    ]
                },
                { id: 'collections', text: 'Collections', img: 'icon-folder', expanded: true, group: true,
                    nodes: [
                            { id: 'rr',           text: 'Rent Roll',                     icon: 'fas fa-chart-line',   hint: 'Rent Roll' },
                            // { id: 'dlnq',         text: 'Delinquency Analysis',          icon: 'fas fa-chart-pie',   hint: 'Delinquency Analysis' },
                            { id: 'stmt',         text: 'RA Statements',                 icon: 'fas fa-clipboard', hint: 'Rental Agreement Statements' },
                            { id: 'payorstmt',    text: 'Payor Statements',              icon: 'far fa-clipboard', hint: 'Payor Statements' },
                            // { id: 'prepnotice',   text: 'Prepare Notices',               icon: 'far fa-file-alt', hint: 'Prepare Notices' },
                    ]
                },
                { id: 'acct', text: 'Accounting', img: 'icon-folder', expanded: false, group: true,
                    nodes: [
                            { id: 'close',       text: 'Close Period',                   icon: 'far fa-caret-square-down', hint: 'Close Period' },
                            { id: 'adjust',      text: 'Adjust Closed Period',           icon: 'fas fa-reply',       hint: 'Adjust Closed Period' },
                    ]
                },
                { id: 'facilities', text: 'Facilities Management', img: 'icon-folder', expanded: false, group: true,
                    nodes: [
                            { id: 'svcreq',      text: 'Create Service Request',         icon: 'far fa-square',         hint: 'Create Service Request' },
                            { id: 'svcreqcmp',   text: 'Complete Service Request',       icon: 'far fa-check-square',   hint: 'Complete Service Request' },
                            { id: 'housekpg',    text: 'Housekeeping',                   icon: 'fas fa-home',           hint: 'Housekeeping' },
                            { id: 'prvmaint',    text: 'Preventative Maintenance',       icon: 'fas fa-wrench',         hint: 'Preventative Maintenance' },
                            { id: 'invntory',    text: 'Inventory',                      icon: 'fas fa-shopping-cart',  hint: 'Preventative Maintenance' },
                    ]
                },
                { id: 'tasks', text: 'Tasks', img: 'icon-folder', expanded: true, group: true,
                    nodes: [
                            { id: 'tlds',       text: 'Task List Definitions',           icon: 'far fa-list-ul',        hint: 'Task List Definitions' },
                            { id: 'tls',        text: 'Task Lists',                      icon: 'fas fa-list-alt',       hint: 'Task Lists' },
                    ]
                },
                 { id: 'reports', text: 'Reports', img: 'icon-folder', expanded: false, group: true,
                    nodes: [
                           //{ id: 'RPTasmrpt',     text: 'Assessments',                     icon: 'far fa-file-alt' },
                           //{ id: 'RPTb',          text: 'Business Units',                  icon: 'far fa-file-alt' },
                           { id: 'RPTar',           text: 'Account Rules',                   icon: 'far fa-file-alt' },
                           { id: 'RPTcoa',          text: 'Chart Of Accounts',               icon: 'far fa-file-alt' },
                           //{ id: 'RPTdpm',        text: 'Deposit Methods',                 icon: 'far fa-file-alt' },
                           //{ id: 'RPTdep',        text: 'Depository Accounts',             icon: 'far fa-file-alt' },
                           { id: 'RPTdelinq',       text: 'Delinquency',                     icon: 'far fa-file-alt' },
                           { id: 'RPTgsr',          text: 'GSR',                             icon: 'far fa-file-alt' },
                           { id: 'RPTj',            text: 'Journal',                         icon: 'far fa-file-alt' },
                           { id: 'RPTl',            text: 'Ledger',                          icon: 'far fa-file-alt' },
                           { id: 'RPTla',           text: 'Ledger Activity',                 icon: 'far fa-file-alt' },
                           { id: 'RPTpeople',       text: app.sTransactant,                  icon: 'far fa-file-alt' },
                           //{ id: 'RPTpmt',        text: 'Payment Types',                   icon: 'far fa-file-alt' },
                           //{ id: 'RPTrcptlist',   text: 'Receipts List',                    icon: 'far fa-file-alt' },
                           { id: 'RPTrcbt',         text: app.sRentable+' Type Counts',      icon: 'far fa-file-alt' },
                           { id: 'RPTr',            text: plural(app.sRentable),             icon: 'far fa-file-alt' },
                           { id: 'RPTra',           text: plural(app.sRentalAgreement),      icon: 'far fa-file-alt' },
                           { id: 'RPTrat',          text: app.sRentalAgreement+' Templates', icon: 'far fa-file-alt' },
                           { id: 'RPTrt',           text: app.sRentable+' Types',            icon: 'far fa-file-alt' },
                           { id: 'RPTrr',           text: 'RentRoll',                        icon: 'far fa-file-alt' },
                           //{ id: 'RPTstatements', text: 'Statements',                      icon: 'far fa-file-alt' },
                           //{ id: 'RPTsl',         text: 'String Lists',                    icon: 'far fa-file-alt' },
                           { id: 'RPTtb',           text: 'Trial Balance',                   icon: 'far fa-file-alt' },//
                    ]
                },
                { id: 'setup', text: 'Setup', img: 'icon-wrench', expanded: true, group: true,
                    nodes: [
                            { id: 'accounts',    text: 'Chart Of Accounts',                icon: 'fas fa-list' },
                            { id: 'pmts',        text: 'Payment Types',                    icon: 'far fa-credit-card' },
                            { id: 'dep',         text: 'Depository Accounts',              icon: 'fas fa-university' },
                            { id: 'depmeth',     text: 'Deposit Methods',                  icon: 'far fa-envelope' },
                            { id: 'ars',         text: 'Account Rules',                    icon: 'fas fa-cogs' },
                            { id: 'rt',          text: plural(app.sRentableType),          icon: 'fas fa-asterisk', hint: 'Rentable Types' },
                            { id: 'rentables',   text: plural(app.sRentable),              icon: 'fas fa-cube' },
                            // { id: 'changeRT',    text: 'Change ' + app.sRentable +' Type', icon: 'fas fa-sync-alt' },
                            // { id: 'permissions', text: 'Permissions',                      icon: 'far fa-thumbs-up' },
                    ]
                },
                { id: 'admin', text: 'Administrator', img: 'icon-wrench', expanded: false, group: true,
                    nodes: [
                            { id: 'about',  text: 'Product Info',        icon: 'fas fa-info' },
                            { id: 'tws',    text: 'Timed Work Schedule', icon: 'far fa-calendar-alt' },
                            { id: 'ledgers',text: 'Ledgers',             icon: 'fas fa-book' },
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

                        // if not report node then unselect any other node
                        if (!(w2ui.sidebarL1.selected && w2ui.sidebarL1.selected.startsWith("RPT"))) {
                            w2ui.sidebarL1.unselect();
                        }

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

                    // if node is other than report nodes then
                    // reset report layout content
                    if (!(target.startsWith("RPT"))) {
                        w2ui.reportslayout.content('main', '');
                        w2ui.sidebarL1.unselect();
                    }

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
                        case 'tls':
                        case 'tlds':
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
                        case 'RPTar':
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
                        case 'RPTrcptlist':
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
        };
    } else {
        sbdata = {
            name: 'sidebarL1',
            nodes: [
                { id: 'workflowreceipts', text: plural(app.sAssessment)+' / '+plural(app.sReceipt), img: 'icon-folder', expanded: true, group: true,
                    nodes: [
                            { id: 'receipts',     text: 'Tendered Payment '+app.sReceipt, icon: 'fas fa-star',        hint: plural(app.sReceipt) },
                   ]
                },
                { id: 'reports', text: 'Reports', img: 'icon-folder', expanded: false, group: true,
                    nodes: [
                          { id: 'RPTrcptlist',   text: 'Tendered Payment Log',       icon: 'far fa-file-alt' },
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
                        case 'receipts':
                            // w2ui.sidebarL1.collapse('reports'); // close reports when jumping to a main view
                            switchToGrid(target);
                            break;
                        case 'RPTrcptlist':
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
        };
    }
    w2ui.toplayout.content('left',$().w2sidebar(sbdata));
};


//---------------------------------------------------------------------------------
// getAboutInfo - contacts the server to get info about its version, and updates
//          the version (about) html page
//
// @params  <none>
// @returns <none>
//---------------------------------------------------------------------------------
window.getAboutInfo = function () {
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
};
