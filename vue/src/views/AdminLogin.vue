<template>
    <el-form :model="form" label-width="120px">
        <el-input v-model="form.username" placeholder="用户名" />
        <el-input v-model="form.password" placeholder="密码" />
        <el-button @click="login">登录</el-button>
    </el-form>
</template>

<script>
import { reactive } from 'vue'

export default {
    methods: {
        login() {
            console.log(this.form.username, this.form.password)
        }
    },
    setup() {
        const form = reactive({
            username: '',
            password: '',
        });
        return {
            form,
        }
    }
}
</script>

<style scoped>
.el-form {
    width: 15vw;
    margin: 0 auto;
}

.el-input {
    margin-bottom: 20px;
}

.el-button {
    width: 100%;
}
</style>