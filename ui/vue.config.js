const publicPath = process.env.VUE_APP_BASE_URL ? process.env.VUE_APP_BASE_URL : '/'
module.exports = {
  publicPath,
  devServer: {
    proxy: {
      // proxy all requests starting with /api to jsonplaceholder
      [`${publicPath}api`]: {
        target: process.env.VUE_PROXY,
        changeOrigin: true
      }
    }
  }
}
