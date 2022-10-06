const path = require("path");
const webpack = require("webpack");
const { merge } = require("webpack-merge");
const { CleanWebpackPlugin } = require('clean-webpack-plugin');
const TerserPlugin = require("terser-webpack-plugin");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");


const config = env => require(`./webpack.${env}.js`);

module.exports = (args = { mode: "production", analyze: false }) => {
    const { mode, analyze, watch } = args;
    console.debug(`Webpack running in ${mode} mode`);
    return merge({
        mode,
        entry: {
            areaLanding: {
                import: path.resolve(__dirname, "../public/ts/area-landing.ts"),
                // dependOn: "mapboxGL"
            },
            geographyStart:  path.resolve(__dirname, "../public/ts/geography-start.ts"),
            // mapboxGL: "mapbox-gl"
        },
        module: {
            rules: [
                {
                    test: /\.ts?$/,
                    use: "ts-loader",
                    exclude: /node_modules/,
                },
                {
                    test: /\.s[ac]ss$/i,
                    use: [
                        MiniCssExtractPlugin.loader,
                        "css-loader",
                        "sass-loader"
                    ]
                }
            ],
        },
        resolve: {
            extensions: [".ts", ".js", ".json", ".css", ".scss"],
            modules: [
                path.join(__dirname, '../node_modules')
            ]
        },
        resolveLoader: {
            modules: [
                path.join(__dirname, '../node_modules')
            ]
        },
        output: {
            // filename: "[name].bundle.js",
            filename: "[name].bundle.js",
            path: path.resolve(__dirname, "../assets/dist"),
            clean: true,
        },
        plugins: [
            new webpack.ProgressPlugin(),
            new CleanWebpackPlugin({ verbose: true }),
            new TerserPlugin({ extractComments: false }),
            new MiniCssExtractPlugin({
                filename: "[name].bundle.css",
            })
        ],
        optimization: {
            splitChunks: {
                chunks: "all"
            }
       },
    }, config(mode)(args));
};
