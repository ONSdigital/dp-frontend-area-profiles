const webpack = require("webpack");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const BundleAnalyzerPlugin = require("webpack-bundle-analyzer").BundleAnalyzerPlugin;


module.exports = ({ analyze }) => ({
    watch: true,
    devtool: "inline-source-map",
    plugins: [ ...(analyze === "true") ?
        [new BundleAnalyzerPlugin({
            analyzerHost: "localhost",
            analyzerPort: 4200,
        })] : [],
        new webpack.DefinePlugin({
           "process.env.NODE_ENV": JSON.stringify("development"), 
           // MAPBOX_ACCESS_TOKEN should come from CI env
           "process.env.MAPBOX_ACCESS_TOKEN": JSON.stringify("pk.eyJ1Ijoiam9lZ29vc2ViYXNzIiwiYSI6ImNsMDltMGlwczBidXMzaXJxMWpleGRybm8ifQ.Xlzr3E2l2fSUxvUvt5ndkA"),
        }),
        new MiniCssExtractPlugin({
            filename: "[name].bundle.css",
        }),
    ]
});
