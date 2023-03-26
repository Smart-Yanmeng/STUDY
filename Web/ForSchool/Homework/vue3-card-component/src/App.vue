<script setup></script>

<template>
  <Card
      v-for="(item, index) in cards"
      :key="index"
      :house="item"
  />
</template>

<script>
import Card from "@/components/Card.vue";

export default {
  components: {
    Card,
  },
  data() {
    return {
      cards: [
        {
          imgUrl: 'https://www.google.com/url?sa=i&url=https%3A%2F%2Fimage.baidu.com%2F&psig=AOvVaw0Rmp4oDbvlcENNtnl911_3&ust=1679382626045000&source=images&cd=vfe&ved=0CA8QjRxqFwoTCMDyq_n56f0CFQAAAAAdAAAAABAE',
          city: '天津',
          district: '北京',
          detail: '大爹房',
          price: '100000'
        },
        {
          imgUrl: 'https://www.google.com/url?sa=i&url=https%3A%2F%2Fimage.baidu.com%2F&psig=AOvVaw0Rmp4oDbvlcENNtnl911_3&ust=1679382626045000&source=images&cd=vfe&ved=0CA8QjRxqFwoTCMDyq_n56f0CFQAAAAAdAAAAABAE',
          city: '天津',
          district: '北京',
          detail: '大爹房',
          price: '100000'
        }
      ]
    }
  },
}
</script>