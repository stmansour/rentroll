// import * as path from 'path';
// import * as webpack from 'webpack';
var path = require('path');
var webpack = require('webpack');
var glob = require('glob');

const CaseSensitivePathsPlugin = require('case-sensitive-paths-webpack-plugin');
// const HtmlWebpackPlugin = require('html-webpack-plugin');

const sourcePath = path.join(__dirname, './js');

const config = {
    context: sourcePath,
    entry: 'bundle.js',
    target: 'web',
    resolve: {
        extensions: ['.js'],
        modules: ['node_modules', '.']
    },
    module: {
        rules: [
            {
                // test: "/home/akshay/go/src/rentroll/webclient/js/bundle.js",
                test: /rovreceipt.js$/,
                loader: ['istanbul-instrumenter-loader']
            }
        ]
    },
    plugins: [
        new CaseSensitivePathsPlugin(),
        new webpack.NamedModulesPlugin()
        // new HtmlWebpackPlugin()
    ],
    devtool: 'cheap-module-source-map',
    devServer: {
        historyApiFallback: true,
        hot: true,
        stats: 'minimal'
    },
    node: {
        // workaround for webpack-dev-server issue
        // https://github.com/webpack/webpack-dev-server/issues/60#issuecomment-103411179
        fs: 'empty',
        net: 'empty'
    },
    watch: true,
    mode: 'development'
};

module.exports = config;