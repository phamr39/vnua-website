<template>
    <div>
        <div class="card border-success mb-3">
            <div class="card-header text-center">
                Thông tin tài khoản
            </div>
            <div class="cart-body">
                <div class="p-5">
                    <div class="form-group">
                        <label>Họ và tên</label>
                        <input type="text" class="form-control" :value="profiles.name">
                    </div>
                    <div class="form-group">
                        <label>Email</label>
                        <input type="text" class="form-control" :value="profiles.email">
                    </div>
                    <div class="form-group">
                        <label>Điện thoại</label>
                        <input type="text" class="form-control" :value="profiles.phone_number">
                    </div>
                </div>
            </div>
        </div>

        <div class="card border-success mb-3">
            <div class="card-header text-center">
                Cập nhập thông tin
            </div>
            <div class="cart-body">
                <form class="p-5" @submit.prevent="updateProfile">
                    <div class="form-group">
                        <label>Họ và tên</label>
                        <input type="text" class="form-control" v-model="name">
                    </div>
                    <div class="form-group">
                        <label>Email</label>
                        <input type="text" class="form-control" v-model="email">
                    </div>
                    <div class="form-group">
                        <label>Điện thoại</label>
                        <input type="text" class="form-control" v-model="phone_number">
                    </div>
                    <button type="submit" class="btn btn-success">Xác nhận</button>
                </form>
            </div>
        </div>
    </div>
</template>
<script>
import { mapState } from "vuex";
export default {
    data() {
        return {
            name : "",
            email: "",
            phone_number: "",
            password : "",
            password_confirm : "",
        };
    },
    computed: {
        ...mapState("users", ["profiles"]),
    },
    created(){
        this.$store.dispatch("users/getProfiles", {});
    },
    methods : {
        async updateProfile() {
            const user = {
                name : this.name,
                email : this.email,
                phone_number : this.phone_number,
            }
            await this.$store.dispatch("users/updateProfiles", user);
        },
        async updatePassword() {
            const payload = {
                password: this.password,
                password_confirm: this.password_confirm,
            }
            await this.$store.dispatch("users/updatePassword", payload);
        },
    }
}
</script>
