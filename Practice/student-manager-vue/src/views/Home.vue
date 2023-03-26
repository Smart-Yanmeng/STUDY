<template>
  <div class="home">
    <el-container style="height: 100vh; border: 1px solid #eee">
      <el-aside width="200px" style="background-color: rgb(238, 241, 246);">
        <el-menu :default-openeds="['edu']" @select="menuSelect">
          <el-submenu index="edu">
            <template slot="title"
              ><i class="el-icon-message"></i>教务管理</template
            >
              <el-menu-item index="student">学生管理</el-menu-item>
              <el-menu-item index="class">班级管理</el-menu-item>
              <el-menu-item index="course">课程管理</el-menu-item>
          </el-submenu>
          <el-submenu index="2">
            <template slot="title"
              ><i class="el-icon-setting"></i>系统管理</template
            >
            <el-menu-item index="2-1">系统设置</el-menu-item>
            <el-menu-item index="2-2">日志管理</el-menu-item>
          </el-submenu>
        </el-menu>
      </el-aside>

      <el-container>
        <el-header style="text-align: right; font-size: 12px">
          <span>王小虎</span>
        </el-header>

        <el-main>
          <router-view/>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
export default {
  name: "Home",
  components: {

  },
  methods: {
    menuSelect(index, indexPath) {
      this.$router.push("/" + indexPath.join("/")).catch(() => {})
    }
  }
};
</script>

<style lang="scss" scoped>
  .el-header {
    background-color: #B3C0D1;
    color: #333;
    line-height: 60px;
  }
  
  .el-aside {
    color: #333;
  }
</style>
