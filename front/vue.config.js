module.exports = {
  configureWebpack: {
  devServer: {
    port: 8888,
    disableHostCheck: true,
    proxy: {
      '^/api': {
        target: 'http://192.168.1.7:8080/',
        changeOrigin: true
      }
    }
  }
  },
  transpileDependencies: [
    'vuetify'
  ]
}
