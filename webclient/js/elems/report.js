/*global
    app, w2ui, popupPDFCustomDimensions, getCurrentBusiness, $, console, genDateRangeNavigator,
    handleDateToolbarAction, dateFromString, dateControlString, exportReportCSV, exportReportPDF,
    w2uiDateControlString
*/
"use strict";
window.showReport = function (rptname, elToFocus) {
    if (rptname === '') {
        return;
    }

    w2ui.toplayout.content('main', w2ui.reportslayout);
    w2ui.toplayout.hide('right',true);

    // check EDI mode for this business and set app.D2 accordingly
    var BID = getCurrentBID();
    var BUD = getBUDfromBID(BID);
    var bizEDIEnabled = EDIEnabledForBUD(BUD);
    var edi = bizEDIEnabled ? 1 : 0;

    var url = '/v1/report/' + String(BID) + '?r=' + rptname +'&dtstart=' + app.D1 + '&dtstop=' + app.D2 + '&edi=' + String(edi);

    // var callBack;
    // if (elToFocus) {
    //     callBack = function() {
    //         // $("input[name="+elToFocus+"]").prop('readonly', true).focus().prop('readonly', false);
    //         // elToFocus.focus();
    //         // document.getElementsByName(elToFocus)[0].focus(); // arrr..... does not found element, WHY!!
    //     };
    // }
    w2ui.reportslayout.load('main', url, null, null /*callBack*/);
};

window.buildReportElements = function (){
    //------------------------------------------------------------------------
    //          reportslayout
    //------------------------------------------------------------------------
    $().w2layout({
        name: 'reportslayout',
        padding: 0,
        panels: [
            { type: 'top',size: 34, content: 'reports toolbar'},
            { type: 'left', size: 20, style: app.prefmt, hidden: false },
            { type: 'main',  size: 100, style: app.prefmt},
            { type: 'preview', size: 0, hidden: true, content: 'reports preview' },
            { type: 'right', size: 200, hidden: true, content: 'reports - detail' },
            { type: 'bottom', size: 20, hidden: true, content: 'reports - bottom' },
        ]
    });

    //------------------------------------------------------------------------
    //          reportstoolbar
    //------------------------------------------------------------------------
    var tmp = genDateRangeNavigator('date');
    tmp.push.apply(tmp, [
        { type: 'spacer',},
        { type: 'button', id: 'csvexport', icon: 'fas fa-table', tooltip: 'export to CSV' },
        { type: 'button', id: 'printreport', icon: 'far fa-file-pdf', tooltip: 'export to PDF' },
        { type: 'break', id: 'break2' },
        { type: 'menu-radio', id: 'page_size', icon: 'fas fa-print',
            tooltip: 'exported PDF page size',
            text: function (item) {
            //var text = item.selected;
            var el   = this.get('page_size:' + item.selected);
            if (item.selected == "Custom") {
                popupPDFCustomDimensions();
            }
            return 'Page Size: ' + el.text;
            },
            selected: 'USLetter',
            items: [
                { id: 'USLetter', text: 'US Letter (8.5 x 11 in)'},
                { id: 'Legal', text: 'Legal (8.5 x 14 in)'},
                { id: 'Ledger', text: 'Ledger (11 x 17 in)'},
                { id: 'Custom', text: 'Custom'},
            ]
        },
        { type: 'menu-radio', id: 'orientation', icon: 'far fa-clone fa-rotate-90',
            tooltip: 'exported PDF orientation',
            text: function (item) {
            //var text = item.selected;
            var el   = this.get('orientation:' + item.selected);
            var pageSize = w2ui.reportstoolbar.get('page_size').selected;
            if (pageSize != "Custom" && item.selected == "Portrait") {
                app.pdfPageWidth = app.pageSizes[pageSize].w;
                app.pdfPageHeight = app.pageSizes[pageSize].h;
            }
            else if (pageSize != "Custom" && item.selected == "LandScape") {
                app.pdfPageWidth = app.pageSizes[pageSize].h;
                app.pdfPageHeight = app.pageSizes[pageSize].w;
            }
            return 'Orientation: ' + el.text;
            },
            selected: 'LandScape',
            items: [
                { id: 'LandScape', text: 'LandScape'},
                { id: 'Portrait', text: 'Portrait'},
            ]
        },
    ]);

    w2ui.reportslayout.content('top', $().w2toolbar({
        name: 'reportstoolbar',
        items: tmp,
        onClick: function (event) {
            // var d1, d2; // start date, stop date

            if (event.target == "page_size") {
                console.log("Page size selected");
            }
            else if (event.target == "orientation") {
                console.log("orientation selected");
            }
            else if (event.target == "csvexport") {
                // now call to export csv report function with start and stop date
                exportReportCSV(app.last.report, app.D1, app.D2);
            }
            else if (event.target == "printreport") {
                // call to export pdf report function with start and stop date
                exportReportPDF(app.last.report, app.D1, app.D2);
            }
            else{
                handleDateToolbarAction(event,'date');
                showReport(app.last.report);
            }
            // TODO: prevent refresh, why toolbar needs to be refreshed when user just selects
            // paper size, orientation? That refresh must be prevented.
        },
        onRefresh: function (event) {
            if (event.target == 'monthfwd') {  // we do these tasks after monthfwd is refreshed so we know that the 2 date controls exist
                var x = document.getElementsByName("dateD1");
                x[0].value = app.D1;
                x = document.getElementsByName("dateD2");
                x[0].value = app.D2;
            }
        }
    }));

    // bind onchange event for date input control for reports
    $(document).on("keypress change", "input[name=dateD1]", function(e) {
        // if event type is keypress then
        if (e.type == 'keypress'){
            // do not procedd further untill user press the Enter key
            if (e.which != 13) {
                return;
            }
        }
        var xd1 = document.getElementsByName('dateD1')[0].value;
        var xd2 = document.getElementsByName('dateD2')[0].value;
        var d1 = dateFromString(xd1);
        var d2 = dateFromString(xd2);

        // check that it is valid or not
        if (isNaN(Date.parse(xd1))) {
            return;
        }
        // check that year is not behind 2000
        if (d1.getFullYear() < 2000) {
            return;
        }
        // check that from date does not have value greater then To date
        if (d1.getTime() >= d2.getTime()) {
            d1 = new Date(d2.getTime() - 24 * 60 * 60 * 1000); //one day back from To date
        }
        app.D1 = dateControlString(d1);
        app.D2 = dateControlString(d2);
        showReport(app.last.report, "dateD1");
    }).on("keypress change", "input[name=dateD2]", function(e) {
        // if event type is keypress then
        if (e.type == 'keypress'){
            // do not procedd further untill user press the Enter key
            if (e.which != 13) {
                return;
            }
        }
        var xd1 = document.getElementsByName('dateD1')[0].value;
        var xd2 = document.getElementsByName('dateD2')[0].value;
        var d1 = dateFromString(xd1);
        var d2 = dateFromString(xd2);

        xd2 = w2uiDateControlString(d2);
        // check that it is valid or not
        if (isNaN(Date.parse(xd2))) {
            return;
        }
        // check that year is not behind 2000
        if (d2.getFullYear() < 2000) {
            return;
        }
        // check that from date does not have value greater then To date
        if (d2.getTime() <= d1.getTime()) {
            d2 = new Date(d1.getTime() + 24 * 60 * 60 * 1000); //one day forward from From date
        }
        app.D1 = dateControlString(d1);
        app.D2 = dateControlString(d2);
        showReport(app.last.report, "dateD2");
    });
};
