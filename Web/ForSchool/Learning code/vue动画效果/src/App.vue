<script setup></script>

<template>
  <button @click="display = !display">click</button>
  <Transition>
    <img v-if="display" src="https://img1.baidu.com/it/u=4603732,1550247081&fm=253&fmt=auto&app=120&f=JPEG?w=1280&h=800">
  </Transition>
</template>

<script>
export default {
  data() {
    return {
      display: true
    }
  },
  methods: {},
  computed: {}
}
</script>

<style scoped>
/* 图片出现之前的样式 */
.v-enter-from {
  opacity: 0;
  transform: translate(-700px);
}

/* 图片出现的动画时间 */
.v-enter-active {
  transition: 1s;
}

/* 图片消失之后的样式 */
.v-leave-to {
  opacity: 0;
  transform: scale(0.1);
}

/* 图片消失的动画时间 */
.v-leave-active {
  transition: 1s;
}
</style>