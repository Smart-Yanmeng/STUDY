module.exports = {
    lintOnSave: false,
    publicPath: './',
    outputDir: process.env.NODE_ENV === 'production' ? 'dist' : 'sit-dist',
    productionSourceMap: false,
    devServer: {
        open:true,
        proxy: {
            [process.env.VUE_APP_BASE_API]: {
                target: 'http://localhost:8081/',
                // target: 'http://www.codeedu.com.cn/',
                changeOrigin: true,
                pathRewrite: {
                    ['^' + process.env.VUE_APP_BASE_API]: ''
                }
            },
            [process.env.VUE_APP_IMG_BASE_API]: {
                target: process.env.IMGTARGET,
                changeOrigin: true,
                pathRewrite: {
                    ['^' + process.env.VUE_APP_BASE_API]: ''
                }
            },
        }
    }
}

