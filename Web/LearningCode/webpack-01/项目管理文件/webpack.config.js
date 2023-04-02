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
      // css 相关的 loader
      {
        // 指定某类文件
        test: /\.css$/,
        // 这两块需要按顺序引入
        use: ['style-loader', 'css-loader']
      },
      // css 中图片相关的 loader
      {
        test: /\.(png|jpg|gif)$/,
        use: [{
          loader: 'file-loader',
          // 配置
          options: {
            // 指定 html 中引用图片的相对 html 文件的路径
            publicPath: './../dist/',
            // 指定文件名，一般会放在一个 img 相关的文件夹中
            name: 'imgs/[contenthash].[ext]'
          }
        }]
      },
      // es6 转 es5 loader
      {
        test: /\.js$/,
        // 除去 node_modules 文件夹，避免项目依赖包和插件被更改
        exclude: /(node_modules)/,
        use: {
          loader : 'babel-loader',
          options : {
              presets : ['@babel/preset-env']
          }
        }
      }
    ]
  }
}