const {defineConfig} = require('@vue/cli-service')

const path = require('path')

const port = 8080 // dev port

function resolve(dir) {
    return path.join(__dirname, dir)
}

module.exports = defineConfig({
    transpileDependencies: true,
    lintOnSave: false,
    publicPath: '/',
    outputDir: 'dist',
    assetsDir: 'static',
    devServer: {
        port: port,
        // 配置代理
        proxy: {
            '/api': {
                target: 'http://127.0.0.1:8081',
                changeOrigin: true,
                ws: false,
                pathRewrite: {
                    '^/api': ''
                }
            },
        },
    }
})
