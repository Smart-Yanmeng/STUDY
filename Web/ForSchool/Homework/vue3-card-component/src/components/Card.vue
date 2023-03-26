<template>
  <div class="card">
    <div class="header">
      <img
          class="header-picture"
          :src="house.imgUrl"
      />
    </div>
    <div class="body">
      <p class="body-city">{{ house.city }} · {{ house.area }}</p>
      <p class="body-district">{{ house.district }}</p>
    </div>
    <div class="footer">
      <span class="footer-detail">{{ house.detail }}</span>
      <span class="footer-price">{{ house.price }}万</span>
    </div>
  </div>
</template>

<script>
export default {
  props: ["house"],
};
</script>

<style scoped>
.card {
  width: 250px;
  height: 300px;
}

.header {
  height: 200px;
}

.body {
  height: 60px;
}

.footer {
  height: 40px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-picture {
  width: 246px;
  height: 198px;
  border-radius: 10px;
}

.body p {
  font-weight: bold;
}

.footer-detail {
  color: gray;
}

.footer-price {
  color: red;
  font-weight: bold;
}
</style>