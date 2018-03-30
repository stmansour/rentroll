"use strict";
//------------------------------------------------------------------------
//          notesPopUp
//------------------------------------------------------------------------
window.notesPopUp = function () {
    w2popup.open({
        width   : 580,
        height  : 350,
        title   : 'Notes',
        body    : '<div class="w2ui-centered" style="line-height: 1.8">'+
                  '     This is work in progress, not the actual interface.<br><br>'+
                  '     Add Note: <textarea name="comments" type="text" style="width: 385px; height: 80px"></textarea><br>'+
                   '</div>',
        buttons : '<button class="w2ui-btn" onclick="w2popup.close()">Ok</button>'+
                  '<button class="w2ui-btn" onclick="w2popup.close()">Cancel</button>'
    });
};
