<template>
    <el-row>
        <el-col>
            <el-menu default-active="1" class="el-menu-vertical-demo">
                <router-link class="link-item" :to="{name: 'user_index'}">
                    <el-menu-item index="1">
                        <template #title>
                            <span>用户管理</span>
                        </template>
                    </el-menu-item>
                </router-link>
                <router-link class="link-item" :to="{name: 'role_index'}">
                    <el-menu-item index="2">
                        <template #title>
                            <span>角色管理</span>
                        </template>
                    </el-menu-item>
                </router-link>
                <router-link class="link-item" :to="{name: 'access_index'}">
                    <el-menu-item index="3">
                        <template #title>
                            <span>权限管理</span>
                        </template>
                    </el-menu-item>
                </router-link>
            </el-menu>
        </el-col>
    </el-row>
</template>
  
<script>
</script>

<style scoped>
.link-item {
    text-decoration: none;
}
</style>
  