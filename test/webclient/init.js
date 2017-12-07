"use strict";

var common = require("./common.js");

exports.init = function() {
    // ========== CASPER OPTIONS ==========
    casper.options.viewportSize = {width: pageWidth, height: pageHeight};
    casper.options.exitOnError = true;
    // casper.options.colorizerType = 'Dummy';
    casper.options.clientScripts.push("./visibility.js");
    casper.options.onError = function(casper, msg, backtrace) {
        common.capture('onError.png');
        // ------------------------------
        // NOTE: `this.log()` cause to run casper infinitely, so don't put it
        // ------------------------------
        this.echo("[{0}] onError msg: {1}".format(logSpace, msg));
        this.echo("[{0}] onError backtrace: {1}".format(logSpace, backtrace));
        this.exit();
    };
    casper.options.onLoadError = function(casper, url, status) {
        this.echo("[{0}] onLoadError URL: {1}".format(logSpace, url));
        this.echo("[{0}] onLoadError Status: {1}".format(logSpace, status));
        this.exit();
    };

    // ========== CASPER CUSTOM LOGGING ==========
    // var stderr = require('system').stderr;
    // var tsWidth = 4;
    // var tabSpace = Array(tsWidth+1).join(" "); // 4 space tab
    // casper.on('log', function onLog(entry) {
    //     stderr.write([
    //        new Date().toISOString(),
    //        entry.level,
    //        entry.message + '\n'
    //    ].join(tabSpace));
    // });

    // ========== STRING FORMAT PROTOTYPE ==========
    /*
    String format: https://gist.github.com/tbranyen/1049426 (if want to format object, array as well)
    Reference: https://stackoverflow.com/questions/610406/javascript-equivalent-to-printf-string-format
    ---------------------------------------------------------------------------------
    > "{0} is awesome {1}".format("javascript", "!?")
    > "javascript is awesome !?"
    */
    String.prototype.format = function() {
        var args = arguments;
        return this.replace(/{(\d+)}/g, function(match, number) {
            return typeof args[number] != 'undefined'? args[number] : match;
        });
    };
};