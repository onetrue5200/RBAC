<template>
    <CardContainer>
        <el-table :data=data align="center">
            <el-table-column align="center" fixed prop="Id" label="Id" width="50" />
            <el-table-column align="center" prop="Name" label="角色名" width="100" />
            <el-table-column align="center" label="权限" width="600">
                <template #default="scope">
                    <div style="display: flex; align-items: center">
                        <span style="margin-left: 10px">
                            <el-tag class="ml-2" type="success" v-for="access in scope.row.Accesses" :key="access.Id">
                                {{access.Name}}
                            </el-tag>
                        </span>
                    </div>
                </template>
            </el-table-column>
            <el-table-column align="center" fixed="right" label="操作" width="150">
                <template #default="scope">
                    <div style="display: flex; align-items: center">
                        <span style="margin-left: 10px">
                            <el-button link type="primary" size="small" @click=edit_data(scope.row)>
                                编辑
                            </el-button>
                            <el-button link type="primary" size="small" @click=delete_data(scope.row.Id)>删除</el-button>
                        </span>
                    </div>
                </template>
            </el-table-column>
        </el-table>
        <div class="add-btn">
            <el-button type="success" :icon="Plus" @click="dialogFormVisible = true">添加</el-button>
        </div>

        <!-- 添加用户 -->
        <el-dialog v-model="dialogFormVisible" title="添加">
            <el-form :model="form">
                <el-form-item label="角色名">
                    <el-input v-model="form.name" />
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <el-button @click="dialogFormVisible = false">取消</el-button>
                    <el-button type="primary" @click=add_data>创建</el-button>
                </span>
            </template>
        </el-dialog>

        <!-- 修改用户 -->
        <el-dialog v-model="dialogFormEdit" title="修改角色">
            <el-form :model="formedit">
                <el-form-item label="角色名">
                    <el-input v-model="formedit.name" />
                </el-form-item>
                <el-form-item label="权限">
                    <el-checkbox-group v-model="checkList">
                        <el-checkbox v-for="access in accesses" :key="access.Id" :label="access.Name" border
                            @change="do_access(access.Id)" />
                    </el-checkbox-group>
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <el-button @click="dialogFormEdit = false">取消</el-button>
                    <el-button type="primary" @click=do_edit_data>修改</el-button>
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
        let accesses = ref([]);
        let role_id = ref('');
        let access_id = ref('');
        const checkList = ref([])

        let dialogFormVisible = ref(false)
        let form = reactive({
            name: '',
        })

        let dialogFormEdit = ref(false)
        let formedit = reactive({
            id: '',
            name: '',
        })

        const refresh = () => {
            form.username = '';
            $.ajax({
                url: "http://127.0.0.1:3000/api/role",
                type: "get",
                success(resp) {
                    data.value = resp.data
                },
            });
            $.ajax({
                url: "http://127.0.0.1:3000/api/access",
                type: "get",
                success(resp) {
                    accesses.value = resp.data;
                },
            });
        };
        refresh();

        const add_data = () => {
            $.ajax({
                url: "http://127.0.0.1:3000/api/role",
                type: "post",
                data: {
                    name: form.name,
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

        const edit_data = (role) => {
            dialogFormEdit.value = true;
            formedit.id = role.Id;
            formedit.name = role.Name;
            role_id = role.Id;
            checkList.value = [];
            for (let t of data.value) {
                if (t["Id"] === role.Id) {
                    for (let access of t["Accesses"]) {
                        checkList.value.push(access["Name"])
                    }
                    break;
                }
            }
        }

        const do_edit_data = () => {
            $.ajax({
                url: "http://127.0.0.1:3000/api/role",
                type: "put",
                data: {
                    id: formedit.id,
                    name: formedit.name,
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

        const delete_data = (id) => {
            $.ajax({
                url: "http://127.0.0.1:3000/api/role/" + id,
                type: "delete",
                success(resp) {
                    if (resp.message === "success") {
                        refresh();
                    } else {
                        ElMessage(resp.message)
                    }
                },
            });
        }

        const do_access = (id) => {
            access_id = id;

            let local_role_id = role_id;
            let accesses_num;
            for (let t of data.value) {
                if (parseInt(t["Id"]) === local_role_id) {
                    accesses_num = t["Accesses"].length;
                    break;
                }
            }
            
            if (accesses_num < checkList.value.length) {
                $.ajax({
                    url: "http://127.0.0.1:3000/api/role_access",
                    type: "post",
                    data: {
                        role_id: role_id,
                        access_id: access_id,
                    },
                    success(resp) {
                        if (resp.message === "success") {
                            refresh();
                        } else {
                            ElMessage(resp.message)
                        }
                    },
                });
            } else if (accesses_num > checkList.value.length) {
                $.ajax({
                    url: "http://127.0.0.1:3000/api/role_access",
                    type: "put",
                    data: {
                        role_id: role_id,
                        access_id: access_id,
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
            accesses,
            checkList,
            Plus,
            dialogFormVisible,
            form,
            dialogFormEdit,
            formedit,

            add_data,
            edit_data,
            do_edit_data,
            delete_data,
            do_access,
        }
    }
}
</script>
<style scoped>
.add-btn {
    text-align: center;
}
</style>