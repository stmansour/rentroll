const path = require('path');
var webpack = require('/usr/local/lib/node_modules/webpack');

config = {
    // define entry point
    entry: "./index.js",

    // define output point
    output: {
        path: path.resolve(__dirname, ''), // no need to create folder
        filename: 'bundle.js'
    }
};

module.exports = config;
