// node 中自带的模块，用来处理文件路径
const path = require("path");

module.exports = {
  // 指定要处理的js（入口 js 文件）
  entry: "./main.js",

  // 指定文件输出的配置
  output: {
    // 指定打包输出的路径
    // __dirname：Node.js 中指向被执行 js 文件的绝对路径
    path: path.resolve(__dirname, "dist"), // path 的路径拼接方法

    // 打包输出的文件名
    filename: "bundle.js"
  },
  // 加入 loader
  module: {
    rules: [
      // vue loader
      {
        test: /\.vue$/,
        use: ['vue-loader']
      }
    ]
  },
  resolve: {
    alias: {
      // 指定使用 vue 包中的版本
      'vue$': 'vue/dist/vue.esm.js'
    }
  }
}