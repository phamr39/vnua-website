<template>
<div class="login-form">
    <div class="card border-success mb-3 login">
        <div class="card-header text-center">
            Đăng nhập
        </div>
        <div class="cart-body">
            <form @submit.prevent="login" class="p-5">
                <div class="form-group">
                    <label>Email</label>
                    <input type="text" class="form-control" v-model="email" placeholder="email">
                </div>
                <div class="form-group">
                    <label>Mật khẩu</label>
                    <input type="password" class="form-control" v-model="password" placeholder="password ">
                </div>
                <button type="submit" class="btn btn-success">Đăng nhập</button>
            </form>
        </div>
    </div>
</div>

</template>
<script>
import { mapState} from "vuex";
export default {
    name: 'Login',
        data() {
        return {
            email: "",
            password: "",
        };
    },
    computed: mapState("users", ["isLogined", "loginMessage"]),
    created() {
        if (this.isLogined) {
            this.$router.replace("/");
        }
    },
    methods: {
        async login() {
            const user = {
                email : this.email,
                password : this.password,
            }
            await this.$store.dispatch("users/login", user).then(() => {
                this.$router.push("/profiles");
            }).catch(function(){
                alert("Sai email hoặc mật khẩu!");
            });
        },
    },
}
</script>

<style scoped>
.login-form {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
}
.login {
    width: 500px;
}
</style>