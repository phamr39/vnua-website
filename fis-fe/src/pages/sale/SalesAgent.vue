<template>
    <main class="main">
        <nav aria-label="breadcrumb" class="breadcrumb-nav">
            <div class="container">
                <ol class="breadcrumb">
                    <li class="breadcrumb-item"><router-link to="/">Trang chủ</router-link></li>
                    <li class="breadcrumb-item active" aria-current="page">Đăng ký đại lý</li>
                </ol>
            </div><!-- End .container -->
        </nav><!-- End .breadcrumb-nav -->
        <div class="page-header text-center">
            <div class="elementor-background-overlay"></div>
            <div class="container">
                <h1 class="page-title text-white">Trao Giá Trị – Nhận Niềm Tin<span class="text-white">Trở thành đối tác của VNUA Phytopharma</span></h1>
            </div><!-- End .container -->
        </div><!-- End .page-header -->

        <div class="page-content pt-5 bg-light pt-5 pb-5">
            <div class="checkout">
                <div class="container bg-white p-5">
                        <div class="row p-5">
                            <div class="col">
                                <form @submit.prevent="createStore">
                                    <h2 class="page-title">Đăng ký đại lý</h2><!-- End .checkout-title -->
                                    <label>Họ và tên *</label>
                                    <input type="text" class="form-control" name="full_name" v-model="name" required>
                                    <label>Email *</label>
                                    <input type="text" class="form-control" name="email" v-model="email" required>
                                    <label>Điện thoại *</label>
                                    <input type="text" class="form-control" name="phone_number" v-model="phone_number" required>
                                    <label>Địa chỉ *</label>
                                    <input type="text" class="form-control" name="address" v-model="address" required>
                                    <label>Kinh doanh</label>
                                    <input type="text" name="business_model" v-model="business_model" class="form-control">
                                    <label>Nội dung chi tiết</label>
                                    <textarea class="form-control" cols="30" rows="4" name="message" v-model="message" placeholder="Nội dung bạn muốn gửi đến chúng tôi..."></textarea>
                                    <button type="submit" class="btn btn-outline-primary-2 btn-order">
                                        <span class="btn-text">Đăng ký</span>
                                        <span class="btn-hover-text">Tiến hành đăng ký</span>
                                    </button>
                                </form>
                            </div><!-- End .col-lg-9 -->
                        </div><!-- End .row -->
                </div>
            </div>
        </div>
    </main>
</template>
<script>
export default {
    name: 'SalesAgent',
    data() {
        return {
            name : "",
            email: "",
            phone_number: "",
            address: "",
            business_model: "",
            message: "",
        };
    },
    methods: {
        async createStore() {
            const store = {
                name : this.name,
                email : this.email,
                phone_number : this.phone_number,
                address: this.address,
                business_model : this.business_model,
                message : this.message,
            }
            await this.$store.dispatch("customers/createStore", store).then(
                alert('Đăng ký thành công')
            )
            window.location.reload();

        },
    }
}
</script>

<style scoped>
.elementor-background-overlay {
    background-image: url('../../assets/images/backgrounds/news-bg-1.png');
    background-position: center center;
    background-repeat: no-repeat;
    background-size: cover;
    opacity: 0.15;
    transition: border-radius 0.3s, opacity 0.3s;
    width: 100%;
    height: 100%;
    position: absolute;
}
</style>