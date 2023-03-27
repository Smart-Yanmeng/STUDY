<template>
  <div class="a">
    <h1>A</h1>
    <router-link to="/a/c">go C</router-link>
    <router-link to="/a/d">go D</router-link>
    <router-view></router-view>
  </div>
</template>

<style scoped>
.a {
  /* 将A组件染色 */
  border: 1px solid black;
  width: 400px;
  height: 600px;
  background-color: green;
}
</style>

<script>
export default {
  name: "A"
}
</script>