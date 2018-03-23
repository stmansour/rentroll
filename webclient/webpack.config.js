var path = require('path');
var webpack = require('webpack');

const CaseSensitivePathsPlugin = require('case-sensitive-paths-webpack-plugin');

const sourcePath = path.join(__dirname, './js/elems');
const outPath = path.join(__dirname, './js');

const config = {
    context: sourcePath,
    // If created new js file than add into entry list
    entry: ['account.js', 'allocFunds.js', 'ar.js', 'asms.js', 'bpwrapper.js', 'datenav.js', 'datetimeutil.js', 'depmeth.js', 'deposit.js', 'depository.js', 'dirtyforms.js', 'expenses.js', 'init.js', 'layout.js', 'ledger.js', 'login.js', 'notes.js', 'pmt.js', 'ra.js', 'rapicker.js', 'receipt.js', 'rentable.js', 'report.js', 'rovreceipt.js', 'rr.js', 'rt.js', 'rutil.js', 'sidebar.js', 'statements.js', 'stmtpayor.js', 'transactant.js', 'tws.js'],
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