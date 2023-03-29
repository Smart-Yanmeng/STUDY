// 去定义某些变量、函数可以抛出到外边，让别的 js 文件来使用

// 使用 export 导出
// export default function baseAjax() {
//   const xhr = new XMLHttpRequest();
//   xhr.open('GET', 'http://www.codeedu.com.cn/xiaomi/v1/carousel');
//   xhr.send();
//   xhr.onload = function () {
//     console.log(xhr.response);
//   }
// }

// 使用 export 不默认导出
export function baseAjax() {
  const xhr = new XMLHttpRequest();
  xhr.open('GET', 'http://www.codeedu.com.cn/xiaomi/v1/carousel');
  xhr.send();
  xhr.onload = function () {
    console.log(xhr.response);
  }
}

// export 支持多个导出
export {

}