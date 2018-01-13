"use strict";

// ========== UI TEST APP OPTIONS ==========
// pageURL for open the application
exports.pageURL = "http://localhost:8270/home";

// Width and height of viewport
exports.pageWidth = 1855; // 1280
exports.pageHeight = 978; // 720
// exports.pageWidth = 1280; // 1280
// exports.pageHeight = 720; // 720

// wait time for page loading
exports.pageLoadWaitTime = 2000;

var fs = require("fs");

// Directory name for captured screen shot
var CAPTURES_STORE = "CAPTURES";

// Take screenshot of the viewport with filename 'fname'
exports.capture = function(fname) {
    // get absolute path to store capture images
    var fpath = fs.pathJoin(fs.workingDirectory, CAPTURES_STORE, fname);

    // now fire the capture call
    casper.capture(fpath, {
        top: 0,
        left: 0,
        width: exports.pageWidth,
        height: exports.pageHeight
    });
};

// check is column  in excluded grid columns list
exports.isColumnInExcludedGridColumns = function (column, excludedGridColumns) {
    var isColumnInExcludedGridColumns =  false;

    // forEach loop return undefined by default. So variable isColumnInExcludedGridColumns will help to return value.
    excludedGridColumns.forEach(function (excludedGridColumn) {
        if (column === excludedGridColumn){
            isColumnInExcludedGridColumns = true;
        }
    });
    return isColumnInExcludedGridColumns;
};

// Check element's existence(value) in array
exports.isInArray = function(value, array){
  return array.indexOf(value) > -1;
};

// wait time duration for function execution
exports.waitTime = 500;

// base url for API endpoints
exports.apiBaseURL = "http://localhost:8270";

// API version
exports.apiVersion = "v1";

// Unset business id
exports.BID = -1;

// Success flag to match with API response status
exports.successFlag = 'success';

// logSpace name
exports.logSpace = "rrLog";
