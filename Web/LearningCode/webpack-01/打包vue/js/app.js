const appTemplate = {
  template: `
    <div>{{ text }}</div>`,
  data() {
    return {
      text: "实例上的内容模板抽离成单独 js 文件"
    }
  }
}

module.exports = appTemplate;