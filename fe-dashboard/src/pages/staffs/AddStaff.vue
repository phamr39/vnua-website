<template>
    <div class="card border-success mb-3">
        <div class="card-header text-center">
            Thêm nhân viên
        </div>
        <div class="cart-body">
            <form class="p-5" @submit.prevent="submit">
                <div class="form-group">
                    <label>Họ và tên</label>
                    <input type="text" class="form-control" v-model="name" name="full_name">
                </div>
                <div class="form-group">
                    <label>Email</label>
                    <input type="text" class="form-control" v-model="email" name="email">
                </div>
                <div class="form-group">
                    <label>Điện thoại</label>
                    <input type="text" class="form-control" v-model="phone_number" name="phone_number">
                </div>

                <div class="form-group">
                    <label>Mật khẩu</label>
                    <input type="password" class="form-control" v-model="password" name="password">
                </div>

                <div class="form-group">
                    <label>Xác nhận mật khẩu</label>
                    <input type="password" class="form-control" v-model="password_confirm" name="password_confirm">
                </div>

                <button type="submit" class="btn btn-success">Xác nhận</button>
                <button @click="cancel()" class="btn btn-warning ml-5">Hủy bỏ</button>
            </form>
        </div>
    </div>
</template>
<script>
export default {
    name: 'AddStaff',
    data() {
        return {
            name: "",
            email: "",
            phone_number: "",
            password: "",
            password_confirm: "",
        }
    },
    methods : {
        async submit() {
            const user = {
                name : this.name,
                email : this.email,
                phone_number : this.phone_number,
                password : this.password,
                password_confirm : this.password_confirm,
            }
            await this.$store.dispatch("users/createNewUser", user)
            await this.$router.push("/staffs");
        },
        async cancel() {
            await this.$router.push("/staffs");
        }
    }
}
</script>
