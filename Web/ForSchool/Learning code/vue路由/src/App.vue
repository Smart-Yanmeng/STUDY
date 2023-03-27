<script setup></script>

<template>
  <div>
    <!--    <router-link to="/a">go A</router-link>-->
    <!--    <router-link to="/b">go B</router-link>-->
    <button @click="goA">go A</button>
    <button @click="goB">go B</button>
    <Transition>
      <router-view></router-view>
    </Transition>
  </div>
</template>

<style>
.v-enter-from {
  opacity: 0;
}
.v-enter-active {
  transition: 1s;
}
</style>

<script>
export default {
  data() {
    return {
      title: 'template'
    }
  },
  methods: {
    goA() {
      this.$router.push("/a");
    },
    goB() {
      this.$router.push("/b");
    },
  },
  computed: {}
}
</script>