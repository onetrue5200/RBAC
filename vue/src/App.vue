<template>
  <div class="common-layout">
    <el-container>
      <el-header>
        <HeaderBar />
      </el-header>
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
import HeaderBar from '@/components/HeaderBar.vue'

export default {
  components: {
    NavBar,
    HeaderBar,
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
</style>
