<template>
  <div class="common-layout">
    <el-container>
      <el-header class="header">简单的RBAC权限管理系统</el-header>
      <el-container>
        <el-aside width="15vw">
          <NavBar />
        </el-aside>
        <el-main>
          <router-view />
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import NavBar from '@/components/NavBar.vue'

export default {
  components: {
    NavBar,
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

.header {
  font-size: 2.5vw;
  margin: 10px;
}
</style>
