<template>
  <h1>Hello World</h1>
  <h2>Hello {{ box.name }}</h2>
  <h3>this is {{ box.title }}</h3>
  <button @click="passNumber">click me</button>
</template>

<script>
export default {
  data() {
    return {
      number: 56
    }
  },
  methods: {
    passNumber() {
      this.$emit('passNumber', this.number)
    }
  },
  props: ['box']
}
</script>