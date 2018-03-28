var path = require('path');
var webpack = require('webpack');

const CaseSensitivePathsPlugin = require('case-sensitive-paths-webpack-plugin');
const glob = require('glob');

const outPath = path.join(__dirname, './js');

const config = {
    entry: glob.sync('./js/elems/*.js'),
    output: {
        path: outPath,
        publicPath: '/',
        filename: 'bundle.js' // output file name
    },
    target: 'web',
    resolve: {
        extensions: ['.js'],
        modules: ['node_modules', '.']
    },
    module: {
        rules: [
            {
                test: /\.js$/,
                loader: ['istanbul-instrumenter-loader']
            }
        ]
    },
    plugins: [
        new CaseSensitivePathsPlugin(),
        new webpack.NamedModulesPlugin()
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
    mode: 'development'
};

module.exports = config;