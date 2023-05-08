<template>
    <div>
        <h1>我是子组件</h1>
        <h1>{{ number }}</h1>
    </div>
</template>

<script setup>
import {inject} from "vue";

// 子组件中注入这个值
const number = inject("number")
</script>
