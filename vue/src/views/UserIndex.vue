<template>
    <CardContainer>
        <el-table :data=data align="center">
            <el-table-column align="center" fixed prop="Id" label="Id" width="50" />
            <el-table-column align="center" prop="Username" label="用户名" width="100" />
            <el-table-column align="center" label="角色" width="400">
                <template #default="scope">
                    <div style="display: flex; align-items: center">
                        <span style="margin-left: 10px">
                            <el-tag class="ml-2" type="success" v-for="role in scope.row.Roles" :key="role.Id">
                                {{role.Name}}
                            </el-tag>
                        </span>
                    </div>
                </template>
            </el-table-column>
            <el-table-column align="center" prop="IsSuper" label="超级管理员" width="100" />
            <el-table-column align="center" fixed="right" label="操作" width="150">
                <template #default="scope">
                    <div style="display: flex; align-items: center">
                        <span style="margin-left: 10px">
                            <el-button link type="primary" size="small" @click=edit_user(scope.row)>
                                编辑
                            </el-button>
                            <el-button link type="primary" size="small" @click=delete_user(scope.row.Id)>删除</el-button>
                        </span>
                    </div>
                </template>
            </el-table-column>
        </el-table>
        <div class="add-btn">
            <el-button type="success" :icon="Plus" @click="dialogFormVisible = true">添加用户</el-button>
        </div>

        <!-- 添加用户 -->
        <el-dialog v-model="dialogFormVisible" title="添加用户">
            <el-form :model="form">
                <el-form-item label="用户名">
                    <el-input v-model="form.username" />
                </el-form-item>
                <el-form-item label="密码">
                    <el-input v-model="form.password" />
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <el-button @click="dialogFormVisible = false">取消</el-button>
                    <el-button type="primary" @click=add_user>创建</el-button>
                </span>
            </template>
        </el-dialog>

        <!-- 修改用户 -->
        <el-dialog v-model="dialogFormEdit" title="修改用户">
            <el-form :model="formedit">
                <!-- <el-input v-model="formedit.id" type="hidden" /> -->
                <el-form-item label="用户名">
                    <el-input v-model="formedit.username" />
                </el-form-item>
                <el-form-item label="密码">
                    <el-input v-model="formedit.password" />
                </el-form-item>
                <el-form-item label="角色">
                    <el-checkbox-group v-model="checkList">
                        <el-checkbox v-for="role in roles" :key="role.Id" :label="role.Name" border
                            @change="do_role(role.Id)" />
                    </el-checkbox-group>
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <el-button @click="dialogFormEdit = false">取消</el-button>
                    <el-button type="primary" @click=do_edit_user>修改</el-button>
                </span>
            </template>
        </el-dialog>
    </CardContainer>
</template>

<script>
import CardContainer from '@/components/CardContainer'
import { ref, reactive } from 'vue'
import $ from 'jquery'
import { Plus } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'

export default {
    components: {
        CardContainer,
    },
    setup() {
        let data = ref([]);
        let roles = ref([]);
        let user_id = ref('');
        let role_id = ref('');
        const checkList = ref([])

        let dialogFormVisible = ref(false)
        let form = reactive({
            username: '',
            password: '',
        })

        let dialogFormEdit = ref(false)
        let formedit = reactive({
            id: '',
            username: '',
            password: '',
        })

        const refresh = () => {
            form.username = '';
            form.password = '';
            $.ajax({
                url: "http://127.0.0.1:3000/api/user",
                type: "get",
                success(resp) {
                    data.value = resp.data;
                },
            });
            $.ajax({
                url: "http://127.0.0.1:3000/api/role",
                type: "get",
                success(resp) {
                    roles.value = resp.data;
                },
            });
        };
        refresh();

        const add_user = () => {
            $.ajax({
                url: "http://127.0.0.1:3000/api/user",
                type: "post",
                data: {
                    username: form.username,
                    password: form.password,
                },
                success(resp) {
                    if (resp.message === "success") {
                        refresh();
                        dialogFormVisible.value = false;
                    } else {
                        ElMessage(resp.message)
                    }
                },
            });
        };

        const edit_user = (user) => {
            dialogFormEdit.value = true;
            user_id = user.Id;
            formedit.id = user.Id;
            formedit.username = user.Username;
            formedit.password = "";
            checkList.value = [];
            for (let t of data.value) {
                if (t["Id"] === user.Id) {
                    for (let role of t["Roles"]) {
                        checkList.value.push(role["Name"])
                    }
                    break;
                }
            }
        }

        const do_edit_user = () => {
            $.ajax({
                url: "http://127.0.0.1:3000/api/user",
                type: "put",
                data: {
                    id: formedit.id,
                    username: formedit.username,
                    password: formedit.password,
                },
                success(resp) {
                    if (resp.message === "success") {
                        refresh();
                        dialogFormEdit.value = false;
                    } else {
                        ElMessage(resp.message)
                    }
                },
            });
        };

        const delete_user = (id) => {
            $.ajax({
                url: "http://127.0.0.1:3000/api/user/" + id,
                type: "delete",
                success(resp) {
                    if (resp.message === "success") {
                        refresh();
                    } else {
                        ElMessage(resp.message)
                    }
                },
            });
        };

        const do_role = (id) => {
            role_id = id;

            let local_user_id = user_id;
            let roles_num;
            for (let t of data.value) {
                if (parseInt(t["Id"]) === local_user_id) {
                    roles_num = t["Roles"].length;
                    break;
                }
            }
            if (roles_num < checkList.value.length) {
                $.ajax({
                    url: "http://127.0.0.1:3000/api/user_role",
                    type: "post",
                    data: {
                        user_id: user_id,
                        role_id: role_id,
                    },
                    success(resp) {
                        if (resp.message === "success") {
                            refresh();
                        } else {
                            ElMessage(resp.message)
                        }
                    },
                });
            } else if (roles_num > checkList.value.length) {
                $.ajax({
                    url: "http://127.0.0.1:3000/api/user_role",
                    type: "put",
                    data: {
                        user_id: user_id,
                        role_id: role_id,
                    },
                    success(resp) {
                        if (resp.message === "success") {
                            refresh();
                        } else {
                            ElMessage(resp.message)
                        }
                    },
                });
            }
        }

        return {
            data,
            roles,
            user_id,
            role_id,
            Plus,
            dialogFormVisible,
            form,
            dialogFormEdit,
            formedit,
            checkList,

            add_user,
            edit_user,
            do_edit_user,
            delete_user,
            do_role,
        }
    }
}
</script>
<style scoped>
.add-btn {
    text-align: center;
}
</style>