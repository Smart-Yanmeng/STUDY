const a = '来自a.js'
console.log('a', a);

const b = '来自a.js'
console.log('a', b);

// 调用 api.js 中的一个 baseAjax，发现由于模块化引入，无法访问到 api.js
// baseAjax();

// import api.js
// 导入重命名
import {baseAjax as api} from "./api.js";

api();