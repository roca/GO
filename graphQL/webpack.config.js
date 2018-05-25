const  path = require('path');

module.exports = {
    entry: './resources/js/app.js',
    output: {
        path: path.join(__dirname, 'resources'),
        filename: 'bundle.js'
    },
    module: {
        loaders: [
            {
                enforce: "pre",
                test: /\.js$/,
                exclude: /node_modules/,
                loader: "eslint-loader",
            },
            {
                test: /\.js$/,
                exclude: /node_modules/,
                loader: 'babel-loader'
            }
        ]
    }
}