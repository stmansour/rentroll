"use strict";

var fs = require("fs");

var CAPTURES_STORE = "CAPTURES";

exports.capture = function(fname) {
    // get absolute path to store capture images
    var fpath = fs.pathJoin(fs.workingDirectory, CAPTURES_STORE, fname);

    // now fire the capture call
    casper.capture(fpath, {
        top: 0,
        left: 0,
        width: pageWidth,
        height: pageHeight
    });
};

exports.waitTime = 500;

exports.businessUnitValue = "REX";