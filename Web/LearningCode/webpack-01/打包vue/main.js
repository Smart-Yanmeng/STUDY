// 都引入到 main.js 中

// 引入 vue
import Vue from "vue"

/* new 实例过程拆到 main.js 中，最终打包后引入 */
// new Vue({
//   el: "#app",
//   data() {
//     return {
//       text: "这是使用 webpack 引入的 vue"
//     }
//   }
// })

/* 实例上的内容模板抽离 */
// new Vue({
//   el: "#app",
//   template: '<div>{{text}}</div>',
//   data() {
//     return {
//       text: "这是使用 webpack 引入的 vue"
//     }
//   }
// })

/* 实例上的内容模板抽离成单独变量 */
// const appTemplate = {
//   template: `
//     <div>{{ text }}</div>`,
//   data() {
//     return {
//       text: "实例上的内容模板抽离成单独变量"
//     }
//   }
// }
// new Vue({
//   el: "#app",
//   template: '<app-template></app-template>',
//   components: {
//     appTemplate
//   }
// })

/* 抽离成 app.js 引入 */
// const appTemplate = require("./js/app")
// new Vue({
//   el : "#app",
//   template : "<app-template/>",
//   components : {
//       appTemplate
//   }
// })

/* 抽离成 app.vue 引入 */
import appTemplate from "./vue/app.vue"

new Vue({
  el: "#app",
  template: '<app-template/>',
  components: {
    appTemplate
  }
})