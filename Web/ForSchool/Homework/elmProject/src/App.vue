<script setup></script>

<template>
  <!-- good-list -->
  <div class="good-list-container">
    <img src="">
    <table class="good-list" v-for="(item, index) in productLists" :key="index">
      <tr>
        <td>
          <div class="good-image-container">
            <img class="good-image" :src="item.imgUrl">
          </div>
        </td>
        <td class="td-good-info">
          <div class="good-name">{{ item.name }}</div>
          <div class="good-unit">{{ item.unit }}</div>
          <div class="good-sold">月售{{ item.monthlySales }} 好评率{{ item.applauseRate }}</div>
          <div class="good-price">￥{{ item.unitPrice }}</div>
        </td>
        <td class="td-good-num">
          <div class="div-good-num">
            <button class="btn-num-minus-op" v-if="item.number!==0" @click="minusNum(item)">-</button>
            <input class="good-num" type="number" v-if="item.number!==0" v-model="item.number">
            <button class="btn-num-add-op" @click="addNum(item)">+</button>
          </div>
        </td>
      </tr>
    </table>
  </div>

  <!-- cart -->
  <div class="cart-container">
    <div class="total-price">￥{{ totalPrice }}</div>
    <button class="settlement" @click="settlement">去结算</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      productLists: [
        {
          imgUrl:
              "https://img10.360buyimg.com/n1/s150x150_jfs/t1/185921/12/21971/233386/624bef0dE8fb97149/f4293d93cc127341.jpg",
          name: "鸿运热辣鸡排",
          unit: "1块",
          monthlySales: 31,
          applauseRate: 97,
          unitPrice: 13.9,
          number: 0,
        },
        {
          imgUrl:
              "https://img2.baidu.com/it/u=2678938536,2995284458&fm=253&fmt=auto&app=138&f=JPEG?w=220&h=180",
          name: "香酥大鸡腿",
          unit: "1个",
          monthlySales: 33,
          applauseRate: 97,
          unitPrice: 13.9,
          number: 0,
        },
        {
          imgUrl:
              "https://img0.baidu.com/it/u=1572645760,2728237162&fm=253&fmt=auto&app=138&f=JPEG?w=220&h=169",
          name: "黑椒鸡块",
          unit: "1个",
          monthlySales: 29,
          applauseRate: 100,
          unitPrice: 10.9,
          number: 0,
        },
        {
          imgUrl:
              "https://img1.baidu.com/it/u=3127931307,4002335577&fm=253&fmt=auto&app=138&f=JPEG?w=222&h=180",
          name: "秘制烤翅",
          unit: "1个",
          monthlySales: "23",
          applauseRate: 98,
          unitPrice: 13.9,
          number: 0,
        },
      ]
    }
  },
  methods: {
    minusNum(item) {
      item.number--;
    },
    addNum(item) {
      item.number++;
    },
    settlement() {
      alert("恭喜您，成功消费" + this.totalPrice + "大洋！");

      this.productLists.forEach((item) => {
        item.number = 0;
      })
    }
  },
  computed: {
    totalPrice() {
      var sum = 0;

      this.productLists.forEach((item) => {
        sum += item.number * item.unitPrice;
      })

      return sum.toFixed(2);
    }
  }
}
</script>