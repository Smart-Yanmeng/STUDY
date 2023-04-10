<script setup></script>

<template>
    <el-button type="primary">Primary</el-button>
    <el-button type="primary" plain>Primary</el-button>
    <el-button type="primary" round>Primary</el-button>

    <span>num is {{ num }}</span>
    <el-slider v-model="num">slider</el-slider>
</template>

<script>
export default {
    data() {
        return {
            title: 'template',
            num: 0
        }
    },
    methods: {},
    computed: {}
}
</script>