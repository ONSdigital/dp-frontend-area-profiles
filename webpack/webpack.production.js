const webpack = require("webpack");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");


module.exports = ({}) => ({
    watch: false,
    // devtool: "source-map",
    module: {
        rules: [
            {
                test: /\.css$/i,
                use: [MiniCssExtractPlugin.loader, "css-loader"],
            },
        ],
    },
    plugins: [
        new webpack.DefinePlugin({
            "process.env.NODE_ENV": JSON.stringify("production"), 
            // MAPBOX_ACCESS_TOKEN should come from CI env
            "process.env.MAPBOX_ACCESS_TOKEN": JSON.stringify("pk.eyJ1Ijoiam9lZ29vc2ViYXNzIiwiYSI6ImNsMDltMGlwczBidXMzaXJxMWpleGRybm8ifQ.Xlzr3E2l2fSUxvUvt5ndkA"),
         }),
    ]
});
