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

exports.isInArray = function(value, array){
  return array.indexOf(value) > -1;
};

exports.waitTime = 500;

exports.apiBaseURL = "http://localhost:8270";
exports.apiVersion = "v1";
exports.BID = -1;

exports.successFlag = "success";
