/*global
    w2ui, $, app, console, w2utils,
    form_dirty_alert, addDateNavToToolbar, w2uiDateControlString,
    dateFromString, taskDateRender, setToTLForm,
    taskFormDueDate,taskCompletionChange,taskFormDoneDate,
    openTaskForm,setInnerHTML,w2popup,ensureSession,dtFormatISOToW2ui,
    createNewBusiness, getBUDfromBID, exportItemReportPDF, exportItemReportCSV,
    popupNewBusinessForm, getTLDs, getCurrentBID, getNewBusinessRecord,
    closeTaskForm, setTaskButtonsState, renderTaskGridDate, localtimeToUTC, TLD,
    taskFormDueDate1, finishBizForm, setToBizForm, renderReversalIcon,
    getGridReversalSymbolHTML, openBizForm, createNewBusiness, getBusinessInitRecord,
    tlPickerRender, tlPickerDropRender, tlPickerCompare, getTLName, displayHelpContent,
*/

"use strict";

window.displayHelpContent = function(s) {
    w2ui.helpLayout.load('main', '/webclient/html/help/' + s);
};

// buildHelpElements creates the help system interface
//
// INPUTS
//
// RETURNS
//----------------------------------------------------------------------------
window.buildHelpElements = function() {

    //------------------------------------------------------------------------
    //  helpLayout - The layout to contain the help system
    //               top   - menubar
    //               left  - sidebar
    //               main  - help pages
    //------------------------------------------------------------------------
    $().w2layout({
        name: 'helpLayout',
        padding: 0,
        panels: [
            { type: 'left',    size: 200,   hidden: false, content: 'left'},
            { type: 'top',     size: 36,    hidden: false, content: 'top',  resizable: true, style: app.pstyle },
            { type: 'main',    size: '60%', hidden: false, content: '', resizable: true, style: app.pstyle },
            { type: 'preview', size: 0,     hidden: true,  content: 'PREVIEW'  },
            { type: 'bottom',  size: 50,    hidden: true, content: 'bottom', resizable: false, style: app.pstyle },
            { type: 'right',   size: 0,     hidden: true }
        ]
    });

    $().w2toolbar({
        name: 'helpToolbar',
        items: [
            {type: 'html', id: 'item5',
                html: function (item) {
                    var html =
                      '<div style="padding: 3px 10px;">'+
                      ' Search Help:'+
                      '    <input size="20" id="helpSearchText" style="padding: 3px; border-radius: 2px; border: 1px solid silver" value="" onChange="helpCB();"/>'+
                      '</div>';
                    return html;
                }}
        ],
    });


    $().w2sidebar({
        name: 'helpSidebar',
        nodes: [
            { id: 'tutorials', text: 'Tutorial Videos', img: 'icon-folder', expanded: true, group: true,
                nodes: [
                        { id: 'ASMtut',  text: plural(app.sAssessment),         /*icon: 'fas fa-video'*/ },
                        { id: 'RCPtut',  text: plural(app.sReceipt),            /*icon: 'fas fa-video'*/ },
                        { id: 'EXPtut',  text: plural(app.sExpense),            /*icon: 'fas fa-video'*/ },
                        { id: 'DEPtut',  text: 'Deposits',                      /*icon: 'fas fa-video'*/ },
                        { id: 'ALFtut',  text: 'Apply '+plural(app.sReceipt),   /*icon: 'fas fa-video'*/ },
                        { id: 'RAtut',   text: plural(app.sRentalAgreement),    icon: 'fas fa-video',  fname: 'TUTra.html' },
                        { id: 'REStut',  text: plural(app.sReservation),        /*icon: 'fas fa-video'*/ },
                        { id: 'Ttut',    text: plural(app.sTransactant),        /*icon: 'fas fa-video'*/ },
                        { id: 'RRtut',   text: 'Rent Roll',                     /*icon: 'fas fa-video'*/ },
                        { id: 'RAStut',  text: 'RA Statement',                  /*icon: 'fas fa-video'*/ },
                        { id: 'RAPtut',  text: 'Payor Statement',               /*icon: 'fas fa-video'*/ },
                        { id: 'CPtut',   text: 'Close Period',                  /*icon: 'fas fa-video'*/ },
                        { id: 'TLtut',   text: 'Task Lists',                    icon: 'fas fa-video',  fname: 'TUTtl.html' },

               ]
            },
            { id: 'helptext', text: 'Reference Manual', img: 'icon-folder', expanded: true, group: true,
                nodes: [
                        { id: 'ASMref',  text: plural(app.sAssessment),         /* icon: 'fas fa-book' */ },
                        { id: 'RCPref',  text: plural(app.sReceipt),            /* icon: 'fas fa-book' */ },
                        { id: 'EXPref',  text: plural(app.sExpense),            /* icon: 'fas fa-book' */ },
                        { id: 'DEPref',  text: 'Deposits',                      /* icon: 'fas fa-book' */ },
                        { id: 'ALFref',  text: 'Apply '+plural(app.sReceipt),   /* icon: 'fas fa-book' */ },
                        { id: 'RAref',   text: plural(app.sRentalAgreement),    icon: 'fas fa-book', fname: 'ref/RentalAgreements.html'},
                        { id: 'RESref',  text: plural(app.sReservation),        /* icon: 'fas fa-book' */ },
                        { id: 'Tref',    text: plural(app.sTransactant),        /* icon: 'fas fa-book' */ },
                        { id: 'RRref',   text: 'Rent Roll',                     /* icon: 'fas fa-book' */ },
                        { id: 'RASref',  text: 'RA Statement',                  /* icon: 'fas fa-book' */ },
                        { id: 'RAPref',  text: 'Payor Statement',               /* icon: 'fas fa-book' */ },
                        { id: 'CPref',   text: 'Close Period',                  /* icon: 'fas fa-book' */ },
                        { id: 'TLref',   text: 'Task Lists',                    /* icon: 'fas fa-book' */ },

               ]
            },
        ],
        onClick: function (event) {
            var target = event.target;
            var s = '';
            var t = event.node.text;

            if (typeof event.node.fname != "undefined") {
                displayHelpContent(event.node.fname);
            } else {
                var h = 'tutorial';
                if (event.target.substr(-3) == "ref") {
                    h = "manual";
                }
                w2ui.helpLayout.content('main','Sorry, no tutorial available yet for ' + t + '.');
            }
        },
    });
};

window.finishHelpSystem = function() {
    w2ui.helpLayout.content('left',w2ui.helpSidebar);
    w2ui.helpLayout.assignToolbar('top',w2ui.helpToolbar);
    w2ui.helpLayout.showToolbar('top');
};

window.helpCB = function() {
    var x = document.getElementById("helpSearchText");
    if (typeof x == "object") {
        console.log('search for: ' + x.value );
    }
};
