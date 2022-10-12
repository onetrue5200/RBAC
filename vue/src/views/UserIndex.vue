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
                    data.value = resp.data
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
            formedit.id = user.Id;
            formedit.username = user.Username;
            formedit.password = "";
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
            console.log(id);
            $.ajax({
                url: "http://127.0.0.1:3000/api/user/" + id,
                type: "delete",
                success(resp) {
                    console.log(resp);
                    if (resp.message === "success") {
                        refresh();
                    } else {
                        ElMessage(resp.message)
                    }
                },
            });
        }

        return {
            data,
            Plus,
            dialogFormVisible,
            form,
            dialogFormEdit,
            formedit,

            add_user,
            edit_user,
            do_edit_user,
            delete_user,
        }
    }
}
</script>
<style scoped>
.add-btn {
    text-align: center;
}
</style>