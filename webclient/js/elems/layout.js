"use strict";
window.buildAppLayout = function(){
     //------------------------------------------------------------------------
    //          toplayout
    //------------------------------------------------------------------------
    w2ui.mainlayout.content('main', $().w2layout({
        name: 'toplayout',
        padding: 2,
        panels: [
            { type: 'top',     size: 200, style: app.pstyle2,  hidden: true, resizable: true, content: w2ui.newsLayout},
            { type: 'left',    size: app.sidebarWidth, style: app.pstyle2,   resizable: true, content: 'sidebar' },
            { type: 'main',               style: app.pstyle2   },
            { type: 'preview', size: 0,   style: app.bgyellow, hidden: true, resizable: true, content: 'preview' },
            { type: 'right',   size: 200, style: app.pstyle2,  hidden: true, resizable: true },
            { type: 'bottom',  size: 0,   style: app.pstyle2,  hidden: true, resizable: true, content: 'toplayout - bottom' }
        ],
        onHide: function(event) {
            event.onComplete = function() {
                if (event.target === "right" && event.type === "hide") {
                    // get the form from active_form value
                    var f = w2ui[app.active_form];
                    if (f) {
                        // if right panel is being hidden,
                        // then make blank everything related with active form

                        // b'coz rtForm, depositForm, rentalagrForm have been
                        // filled inside a layout, so we need to consider those
                        // separately
                        if (f === this.get("right").content ||
                            f.name == "rtForm" ||
                            f.name == "rentalagrForm" ||
                            f.name === "depositForm") {

                            app.active_form = "";
                            app.active_form_original = {};
                            app.form_is_dirty = false;
                        }
                    }
                }
            };
        },
    }));
    //------------------------------------------------------------------------
    //          NEWS LAYOUT
    //------------------------------------------------------------------------
    $().w2layout({
        name: 'newsLayout',
        padding: 0,
        panels: [
            { type: 'left', hidden: false, style: app.pstyleNB, size: 20 },
            { type: 'top', hidden: true },
            { type: 'main', size: '90%', resizable: true, hidden: false, style: app.pstyleNB, content: 'Hi.  I should load w2ui.newsLayout' },
            { type: 'preview', hidden: true },
            { type: 'bottom', hidden: true },
            { type: 'right', hidden: true }
        ]
    });
};
