"use strict";
/*global
    GridMoneyFormat, number_format, w2ui, $, app, console,setToStmtForm,
    form_dirty_alert, addDateNavToToolbar
*/

function buildTaskListDefElements() {
    //------------------------------------------------------------------------
    //          tldsGrid  -  THE LIST OF ALL Task List Definitions
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'tldsGrid',
        url: '/v1/tlds',
        multiSelect: false,
        postData: {searchDtStart: app.D1, searchDtStop: app.D2},
        show: {
            toolbar         : true,
            footer          : true,
            toolbarAdd      : false,   // indicates if toolbar add new button is visible
            toolbarDelete   : false,   // indicates if toolbar delete button is visible
            toolbarSave     : false,   // indicates if toolbar save button is visible
            selectColumn    : false,
            expandColumn    : false,
            toolbarEdit     : false,
            toolbarSearch   : false,
            toolbarInput    : true,
            searchAll       : false,
            toolbarReload   : true,
            toolbarColumns  : true,
        },
        columns: [
            {field: 'recid',     hidden: true,  caption: 'recid',                   size: '40px',  sortable: true},
            {field: 'BID',       hidden: true,  caption: 'BID',                     size: '40px',  sortable: true},
            {field: 'Name',      hidden: false, caption: 'Name',                    size: '110px', sortable: true},
        ],
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
                        console.log( 'BID = ' + rec.BID + ',   RAID = ' + rec.RAID);
                        setToStmtForm(rec.BID, rec.RAID, app.D1, app.D2);
                    };

                // warn user if form content has been changed
                form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args);
            };
        },
    });

    addDateNavToToolbar('tlds'); // "Grid" is appended to the 
}