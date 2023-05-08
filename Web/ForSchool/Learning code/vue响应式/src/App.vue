<template>
    <h1>{{ doubleNumber }}</h1>
    <h1>非响应式：{{ obj1.count }}</h1>
    <h1>响应式：{{ obj2.count }}</h1>

    <button @click="increase1">Click me to add the first</button>
    <button @click="increase2">Click me to add the second</button>

    <Son/>
</template>

<script setup>

/* 选项式 */
// export default {
//     data() {
//         return {
//             number: 5
//         }
//     },
//     computed: {
//         doubleNumber() {
//             return this.number * 2
//         }
//     }
// }

/* 响应式 */
import {
    computed,
    isProxy,
    isReactive,
    isReadonly,
    isRef,
    onMounted, provide,
    reactive,
    readonly,
    ref,
    toRef,
    unref,
    watch
} from "vue";
import Son from "./components/Son.vue";

const number = ref(5)
const doubleNumber = computed(() => number.value * 2)

/* reactive */
// 非响应式
const obj1 = {count: 0}

// 响应式
const obj2 = reactive({count: 0})

function increase1() {
    obj1.count++
}

function increase2() {
    obj2.count++
}

/* 侦听器 */
// 数据源变化时，调用所给的回调函数
const constNumber1 = ref(0)
watch(constNumber1, () => {
    console.log("constNumber changed")
})

setInterval(() => {
    constNumber1.value++
}, 1000)

// watch 中的回调函数也可以接收参数
const constNumber2 = ref(0)
watch(constNumber2, (newNumber, oldNumber) => {
    console.log(`number 从 ${oldNumber} 变成了 ${newNumber}`)
})

setInterval(() => {
    constNumber2.value++
}, 1000)

/* 响应式工具 */
// ifRef() - 检查某个参数是否为 ref，返回值是一个类型判定
const isRefNumber = ref(0)
console.log(isRef(isRefNumber))

// unRef() - 如果参数是 ref，则返回内部值，否则返回参数本身
console.log(unref(isRefNumber))

// toRef() - 把 reactive 对象中的一个属性创建为 ref，改变原属性的值将更新 ref 的值，反之亦然
const reactiveObj = reactive({age: 18, name: "York"})
const refAge = toRef(reactiveObj, "age")
console.log(refAge.value)

// isProxy() - 检查一个对象是否是由 reactive()、readonly()、shallowReactive() 或 shallowReadonly() 创建的
const arr = reactive([1, 2, 3])
const num = ref(0)
const copy = readonly(arr)
console.log("Arr is Proxy? -> " + isProxy(arr))
console.log("Num is Proxy? -> " + isProxy(num))
console.log("Copy is Proxy? -> " + isProxy(copy))

// isReactive() - 检查一个对象是否由 reactive() 或 shallowReactive() 创建的
const reactiveArr = reactive([1, 2, 3])
const refNum = ref(0)
console.log("Arr is Reactive? -> " + isReactive(reactiveArr))
console.log("Num is Reactive? -> " + isReactive(refNum))

// isReadonly() - 检查传入的值是否为只读对象
console.log("Num is Readonly? -> " + isReadonly(num))
console.log("Copy is Readonly? -> " + isReadonly(copy))

/* 生命周期函数 */
// onMounted
onMounted(() => {
    console.log("Mounted output -> York")
})

/* 依赖注入 */
const provideNumber = ref(0)

// 父组件中提供一个值
provide("number", provideNumber)

</script>