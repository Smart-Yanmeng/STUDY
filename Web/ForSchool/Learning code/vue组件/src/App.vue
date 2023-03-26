<template>
  <MyComponent v-for="(item, index) in MyComponent" :key="index" :box="item" @passNumber="receiveNumber"/>
</template>

<script>
import MyComponent from "./components/MyComponent.vue";

export default {
  components: {
    MyComponent
  },
  data() {
    return {
      MyComponent: [
        {
          title: 'card',
          name: 'York',
          number: 0
        }
      ]
    }
  },
  methods: {
    receiveNumber(nb) {
      this.number = nb;
      console.log(nb);
    }
  }
}
</script>