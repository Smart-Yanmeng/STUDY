<template>
  <div class="app-container">
    <div>{{ text }}</div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      text: "实例上的内容模板抽离成单独的 app 文件"
    }
  }
}
</script>