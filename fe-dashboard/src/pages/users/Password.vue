<template>
    <div class="card border-success mb-3">
        <div class="card-header text-center">
            Thay đổi mật khẩu
        </div>
        <div class="cart-body">
            <form class="p-5">
                <div class="form-group">
                    <label>Mật khẩu mới</label>
                    <input type="password" class="form-control" placeholder="Password">
                </div>
                <div class="form-group">
                    <label>Nhập lại mật khẩu</label>
                    <input type="password" class="form-control" placeholder="password">
                </div>
                <button type="submit" class="btn btn-success">Xác nhận</button>
                <button @click="cancel()" class="btn btn-warning ml-5">Hủy bỏ</button>
            </form>
        </div>
    </div>
</template>
<script>
export default {
        methods : {
        async updatePassword() {
            const payload = {
                password: this.password,
                password_confirm: this.password_confirm,
            }
            await this.$store.dispatch("users/updatePassword", payload);
        },
        cancel() {
            window.history.back();
        }
    }
}
</script>
