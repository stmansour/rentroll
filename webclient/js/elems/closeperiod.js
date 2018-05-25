/*global
    w2ui,
*/
"use strict";


window.switchToClosePeriod = function() {
    // w2ui.toplayout.load('main', w2ui.closePeriodLayout);
	w2ui.toplayout.load('main', '/webclient/html/cpinfo.html');
	w2ui.toplayout.hide('right',true);
	//w2ui.closePeriodLayout.load('main', '/webclient/html/cpinfo.html');

};

//-----------------------------------------------------------------------------
// buildClosePeriodElements - a layout in which we place an html page
// and a form.
//
// @params
//
// @returns
//-----------------------------------------------------------------------------
window.buildClosePeriodElements = function () {

    // //------------------------------------------------------------------------
    // //          close period layout
    // //------------------------------------------------------------------------
    // $().w2layout({
    //     name: 'closePeriodLayout',
    //     panels: [
    //         { type: "top", hidden: true },
    //         { type: "main", size: 140, style: 'border: 1px solid #cfcfcf; padding: 5px;', content: 'close period main' },
    //         { type: "left", hidden: true },
    //         { type: "right", hidden: true },
    //         { type: "bottom", hidden: true },
    //     ]
    // });

};
