module.exports = {
    publicPath:
        process.env.NODE_ENV === "production" ? "/" : "/",
    configureWebpack: {
        performance: {
            maxEntrypointSize: 512000,
            maxAssetSize: 512000,
        },
    },
};