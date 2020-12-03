module.exports = {
  configureWebpack: {
  devServer: {
    port: 8888,
    disableHostCheck: true,
    proxy: {
      '^/api': {
        target: 'http://35.247.90.77:8080/',
        changeOrigin: true
      }
    }
  }
  },
  transpileDependencies: [
    'vuetify'
  ]
}
