<template lang="">
    <div class="card border-success mb-3">
        <div class="card-header text-center">
            Danh sách nhân viên
        </div>
        <div class="cart-body">
            <div class="p-5">
                <table class="table table table-striped">
                    <thead>
                        <tr>
                            <th scope="col">STT.</th>
                            <th scope="col">Họ tên</th>
                            <th scope="col">Email</th>
                            <th scope="col">Điện thoại</th>
                            <th scope="col">Quyền</th>
                            <th scope="col"></th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="user, id in users" :key="user.id">
                            <td>{{id + 1}}</td>
                            <td>{{user.name}}</td>
                            <td>{{user.email}}</td>
                            <td>{{user.phone_number}}</td>
                            <td>{{user.role}}</td>
                            <td>
                                <button type="button" @click="deleteUser(user.id)" class="btn btn-outline-danger">Xóa</button>
                            </td>
                        </tr>
                    </tbody>
                </table>
                <Pagination
                    :length="totalItems"
                    :pageSize="20"
                    :pageIndex="pageIndex"
                    @change="changePage"
                />
            </div>
        </div>
    </div>
</template>
<script>
import Pagination from "@/components/Pagination.vue";
import apiUsers from "@/services/users.service";
import { mapState } from "vuex";
export default {
    name : 'Staffs',
    components : {
        Pagination,
    },
    computed: {
        ...mapState("users", [
            "users",
            "totalItems",
            "pageIndex",
        ]),
    },
    methods:{
        changePage(pageIndex) {
            this.$store.dispatch("users/getAllUsers", { pageIndex });
        },

        async deleteUser(userId) {
            await apiUsers.deleteUser(userId);
            console.log('user', userId);
            window.location.reload();
        },
    },
    created() {
        this.$store.dispatch("users/getAllUsers", {});
    },
}
</script>
