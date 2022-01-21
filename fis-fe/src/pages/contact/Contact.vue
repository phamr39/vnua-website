<template>
    <main class="main">
        <nav aria-label="breadcrumb" class="breadcrumb-nav border-0 mb-0">
            <div class="container">
                <ol class="breadcrumb">
                    <li class="breadcrumb-item"><router-link to="/">Trang chủ</router-link></li>
                    <li class="breadcrumb-item active" aria-current="page">Liên hệ</li>
                </ol>
            </div><!-- End .container -->
        </nav><!-- End .breadcrumb-nav -->
        <div class="page-content pb-0">
            <div class="container">
                <div class="row">
                    <div class="col-lg-6 mb-2 mb-lg-0">
                        <h2 class="title mb-1">Thông tin liên hệ</h2><!-- End .title mb-2 -->
                        <p class="mb-3">Công ty Cổ phần KHCN Dược liệu Việt Nam.</p>
                        <div class="row">
                            <div class="col-sm-7">
                                <div class="contact-info">
                                    <h3>Văn phòng</h3>

                                    <ul class="contact-list">
                                        <li>
                                            <i class="icon-map-marker"></i>
                                            Học Viện Nông nghiệp Việt Nam, Ngô Xuân Quảng, Trâu Quỳ, Gia Lâm, Hà Nội
                                        </li>
                                        <li>
                                            <i class="icon-phone"></i>
                                            <a href="tel:#"> 0982 27 47 84</a>
                                        </li>
                                        <li>
                                            <i class="icon-phone"></i>
                                            <a href="tel:#"> 0966 10 15 19</a>
                                        </li>
                                        <li>
                                            <i class="icon-envelope"></i>
                                            <a href="mailto:#">duoclieuvietnam@vnua.edu.vn</a>
                                        </li>
                                    </ul><!-- End .contact-list -->
                                </div><!-- End .contact-info -->
                            </div><!-- End .col-sm-7 -->

                            <div class="col-sm-5">
                                <div class="contact-info">
                                    <h3>Thời gian</h3>

                                    <ul class="contact-list">
                                        <li>
                                            <i class="icon-clock-o"></i>
                                            <span class="text-dark">Thứ 2-Thứ 7</span> <br>8am-18pm
                                        </li>
                                        <li>
                                            <i class="icon-calendar"></i>
                                            <span class="text-dark">Chủ nhật</span> <br>8am-12pm
                                        </li>
                                    </ul><!-- End .contact-list -->
                                </div><!-- End .contact-info -->
                            </div><!-- End .col-sm-5 -->
                        </div><!-- End .row -->
                    </div><!-- End .col-lg-6 -->
                    <div class="col-lg-6">
                        <h2 class="title mb-1">Đặt câu hỏi?</h2><!-- End .title mb-2 -->
                        <p class="mb-2">Gửi câu hỏi của bạn cho VNUA Phytopharma</p>

                        <form @submit.prevent="createCustomer" class="contact-form mb-3">
                            <div class="row">
                                <div class="col-sm-6">
                                    <label class="sr-only">Name</label>
                                    <input type="text" class="form-control" placeholder="Tên *" name="full_name" v-model="name" required>
                                </div><!-- End .col-sm-6 -->

                                <div class="col-sm-6">
                                    <label class="sr-only">Email</label>
                                    <input type="email" class="form-control" name="email" v-model="email" placeholder="Email *" required>
                                </div><!-- End .col-sm-6 -->
                            </div><!-- End .row -->

                            <div class="row">
                                <div class="col-sm-6">
                                    <label class="sr-only">Phone</label>
                                    <input type="tel" class="form-control" name="phone_number" v-model="phone_number" placeholder="Số điện thoại">
                                </div><!-- End .col-sm-6 -->

                                <div class="col-sm-6">
                                    <label class="sr-only">Subject</label>
                                    <input type="text" class="form-control" name="subject" v-model="subject" placeholder="Chủ đề">
                                </div><!-- End .col-sm-6 -->
                            </div><!-- End .row -->

                            <label class="sr-only">Message</label>
                            <textarea class="form-control" cols="30" rows="4" name="message" v-model="message" required placeholder="Nội dung *"></textarea>

                            <button type="submit" class="btn btn-outline-primary-2 btn-minwidth-sm">
                                <span>Gửi</span>
                                <i class="icon-long-arrow-right"></i>
                            </button>
                        </form><!-- End .contact-form -->
                    </div><!-- End .col-lg-6 -->
                </div><!-- End .row -->

                <hr class="mt-4 mb-5">

                <div class="stores mb-4 mb-lg-5">
                    <h2 class="title text-center mb-3">Văn phòng đại diện</h2><!-- End .title text-center mb-2 -->

                    <div class="row">
                        <div class="col-lg-6">
                            <div class="store">
                                <div class="row">
                                    <div class="col-sm-5 col-xl-6">
                                        <figure class="store-media mb-2 mb-lg-0">
                                            <img src="@/assets/images/stores/img-2.jpg" alt="image">
                                        </figure><!-- End .store-media -->
                                    </div><!-- End .col-xl-6 -->
                                    <div class="col-sm-7 col-xl-6">
                                        <div class="store-content">
                                            <h3 class="store-title">Trụ sở chính</h3><!-- End .store-title -->
                                            <address>Học Viện Nông nghiệp Việt Nam, Ngô Xuân Quảng, Trâu Quỳ, Gia Lâm, Hà Nội</address>
                                            <div><a href="tel:#">+84 982 27 47 84</a></div>

                                            <h4 class="store-subtitle">Thời gian làm việc:</h4><!-- End .store-subtitle -->
                                            <div>Thứ 2 - Thứ 7 8am đến 18pm</div>
                                            <div>Chủ nhật 8am đến 12pm</div>

                                            <!-- <a href="#" class="btn btn-link" target="_blank"><span>Xem bản đồ</span><i class="icon-long-arrow-right"></i></a> -->
                                        </div><!-- End .store-content -->
                                    </div><!-- End .col-xl-6 -->
                                </div><!-- End .row -->
                            </div><!-- End .store -->
                        </div><!-- End .col-lg-6 -->

                        <div class="col-lg-6">
                            <div class="store">
                                <div class="row">
                                    <iframe src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3724.6900443077316!2d105.92957561533198!3d21.005058293962083!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x3135a94c1f882977%3A0x6d016e6656923f46!2zSOG7jWMgdmnhu4duIE7DtG5nIG5naGnhu4dwIFZp4buHdCBOYW0!5e0!3m2!1svi!2s!4v1630662825067!5m2!1svi!2s" width="600" height="277" style="border:0;" allowfullscreen="" loading="lazy"></iframe>
                                </div><!-- End .row -->
                            </div><!-- End .store -->
                        </div><!-- End .col-lg-6 -->
                    </div><!-- End .row -->
                </div><!-- End .stores -->
            </div><!-- End .container -->
        </div><!-- End .page-content -->
    </main><!-- End .main -->
</template>
<script>
export default {
    name: 'Contact',
    data() {
        return {
            name : "",
            email: "",
            phone_number: "",
            product_info: "",
            subject: "",
            message: "",
        };
    },
    methods: {
        async createCustomer() {
            const customer = {
                name : this.name,
                email : this.email,
                phone_number : this.phone_number,
                product_info: this.product_info,
                subject : this.subject,
                message : this.message,
            }
            await this.$store.dispatch("customers/createCustomer", customer).then(
                alert('Đăng ký thành công')
            )
            window.location.reload();
        },
    }
}
</script>
