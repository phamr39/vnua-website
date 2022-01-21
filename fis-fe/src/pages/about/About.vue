<template>
    <main class="main">
        <nav aria-label="breadcrumb" class="breadcrumb-nav border-0 mb-0">
            <div class="container">
                <ol class="breadcrumb">
                    <li class="breadcrumb-item"><router-link to="/">Trang chủ</router-link></li>
                    <li class="breadcrumb-item active" aria-current="page">Giới thiệu</li>
                </ol>
            </div><!-- End .container -->
        </nav><!-- End .breadcrumb-nav -->
        <div class="container">
            <div class="page-header page-header-big text-center" :style="{'background-image': 'url('+ require('@/assets/images/about/about-banner-2.jpg')+')'}">
                <!-- <h1 class="page-title text-success">Giới thiệu<span class="text-success">Chúng tôi là ...</span></h1> -->
            </div><!-- End .page-header -->
        </div><!-- End .container -->

        <div class="page-content pb-0">
            <div class="container">
                <div class="row">
                    <div class="col-lg-6 mb-3 mb-lg-0">
                        <h2 class="title">Tầm nhìn</h2><!-- End .title -->
                        <p>Trở thành công ty hàng đầu về nghiên cứu sản xuất cung ứng nguyên liệu Dược và sản xuất mỹ phẩm, thực phẩm bảo vệ sức khỏe từ thiên nhiên.</p>
                    </div><!-- End .col-lg-6 -->
                    
                    <div class="col-lg-6">
                        <h2 class="title">Sứ mệnh</h2><!-- End .title -->
                        <p>Thực hiện hóa các giá trị nghiên cứu, chất xám của các nhà khoa học thành sản phẩm có giá trị bảo vệ sức khỏe cộng đồng. </p>
                    </div><!-- End .col-lg-6 -->
                </div><!-- End .row -->
                <div class="row">
                    <div class="col mt-3 mb-3 mb-lg-0">
                        <h2 class="title">Giá trị cốt lõi</h2><!-- End .title -->
                        <p>Chúng tôi nghiên cứu tạo sản phẩm từ nguồn thảo dược chất lượng cao được nuôi trồng, thu hái bằng chất xám của các nhà khoa học hàng đầu Việt Nam.</p>
                    </div><!-- End .col-lg-6 -->
                </div>
                <div class="mb-5"></div><!-- End .mb-4 -->
            </div><!-- End .container -->

            <div class="bg-light-2 pt-6 pb-5 mb-6 mb-lg-8">
                <div class="container">
                    <div class="row">
                        <div class="col-lg-5 mb-3 mb-lg-0">
                            <h2 class="title">Chúng tôi là ...</h2><!-- End .title -->
                            <p class="lead text-primary">Công Ty Cổ Phần Khoa Học Công Nghệ Dược Liệu Việt Nam (VNUA PHYTOPHARMA)</p>
                            <p class="mb-1">Được thành lập dựa trên nền tảng khai thác kết quả nghiên cứu của các nhà khoa học thuộc Học Viện Nông Nghiệp Việt Nam. </p>
                            <p class="mb-1">VNUA PHYTOPHARMA khai thác nguồn dược liệu bản địa vô cùng phong phú của Việt Nam. Các nhà khoa học tiến hành sàng lọc, đánh giá dược chất, bào chế theo quy trình công nghệ hiện đại và sản xuất sản phẩm tại nhà máy Dược Mỹ phẩm đạt tiêu chuẩn GMP (Dược phẩm), cGMP (Mỹ phẩm). Quá trình nghiên cứu tạo sản phẩm được luôn có sự tham vấn của các bác sĩ, lương y nhiều năm kinh nghiệm tạo ra các sản phẩm Thực phẩm bảo vệ sức khoẻ, Mỹ phẩm có chất lượng cao từ nguồn dược liệu Việt Nam. </p>
                            <p class="mb-1">Công ty Cổ phần Khoa học Công  nghệ Dược liệu Việt Nam, Học viện Nông nghiệp luôn đặt tâm huyết và trách nhiệm cao nhất trong từng sản phẩm nghiên cứu để tạo ra giá trị thực, cải thiện sức khỏe và sắc đẹp của con người Việt Nam. </p>

                            <router-link to="/su-kien" class="btn btn-sm btn-minwidth btn-outline-primary-2">
                                <span>Xem thêm về chúng tôi</span>
                                <i class="icon-long-arrow-right"></i>
                            </router-link>
                        </div><!-- End .col-lg-5 -->

                        <div class="col-lg-6 offset-lg-1">
                            <div class="about-images">
                                <img src="@/assets/images/about/1.png" alt="" class="about-img-front">
                                <img src="@/assets/images/about/3.jpg" alt="" class="about-img-back">
                            </div><!-- End .about-images -->
                        </div><!-- End .col-lg-6 -->
                    </div><!-- End .row -->
                </div><!-- End .container -->
            </div><!-- End .bg-light-2 pt-6 pb-6 -->

            <div class="mb-2"></div><!-- End .mb-2 -->

            <div class="bg-image pt-7 pb-5 pt-md-12 pb-md-9" style="background-color: #27ae60;">
                <div class="container">
                    <div class="row">
                        <div class="col-6 col-md-3">
                            <div class="count-container text-center">
                                <div class="count-wrapper text-white">
                                    <span class="count">10.000+</span>
                                </div><!-- End .count-wrapper -->
                                <h3 class="count-title text-white">Khách hàng hài lòng</h3><!-- End .count-title -->
                            </div><!-- End .count-container -->
                        </div><!-- End .col-6 col-md-3 -->

                        <div class="col-6 col-md-3">
                            <div class="count-container text-center">
                                <div class="count-wrapper text-white">
                                    <span class="count">100%</span>
                                </div><!-- End .count-wrapper -->
                                <h3 class="count-title text-white">Tăng trưởng 100% so với năm trước</h3><!-- End .count-title -->
                            </div><!-- End .count-container -->
                        </div><!-- End .col-6 col-md-3 -->

                        <div class="col-6 col-md-3">
                            <div class="count-container text-center">
                                <div class="count-wrapper text-white">
                                    <span class="count">99%</span>
                                </div><!-- End .count-wrapper -->
                                <h3 class="count-title text-white">Khách hàng quay lại</h3><!-- End .count-title -->
                            </div><!-- End .count-container -->
                        </div><!-- End .col-6 col-md-3 -->

                        <div class="col-6 col-md-3">
                            <div class="count-container text-center">
                                <div class="count-wrapper text-white">
                                    <span class="count">10</span>
                                </div><!-- End .count-wrapper -->
                                <h3 class="count-title text-white">Giải thưởng liên quan</h3><!-- End .count-title -->
                            </div><!-- End .count-container -->
                        </div><!-- End .col-6 col-md-3 -->
                    </div><!-- End .row -->
                </div><!-- End .container -->
            </div><!-- End .bg-image pt-8 pb-8 -->

            <div class="bg-light-2 pt-6 pb-7 mb-6">
                <div class="container">
                    <h2 class="title text-center mb-4">Con người của VNUA Phytopharma</h2><!-- End .title text-center mb-2 -->

                    <div class="row">
                        <div class="col-sm-6 col-lg-3">
                            <div class="member member-2 text-center">
                                <figure class="member-media">
                                    <img src="@/assets/images/about/member/1.jpg" alt="member photo">
                                </figure><!-- End .member-media -->
                                <div class="member-content">
                                    <h3 class="member-title">Ông Nguyễn Quốc Hưng<span>Chủ tịch</span><span>Tiến sĩ Khoa học</span></h3><!-- End .member-title -->
                                </div><!-- End .member-content -->
                            </div><!-- End .member -->
                        </div><!-- End .col-lg-3 -->

                        <div class="col-sm-6 col-lg-3">
                            <div class="member member-2 text-center">
                                <figure class="member-media">
                                    <img src="@/assets/images/about/member/2.jpg" alt="member photo">
                                </figure><!-- End .member-media -->
                                <div class="member-content">
                                    <h3 class="member-title">Bà Phí Thị Cẩm Miện<span>Tổng giám đốc</span><span>Tiến sĩ ngành Công nghệ sinh học</span></h3><!-- End .member-title -->
                                </div><!-- End .member-content -->
                            </div><!-- End .member -->
                        </div><!-- End .col-lg-3 -->

                        <div class="col-sm-6 col-lg-3">
                            <div class="member member-2 text-center">
                                <figure class="member-media">
                                    <img src="@/assets/images/about/member/3.jpg" alt="member photo">
                                </figure><!-- End .member-media -->
                                <div class="member-content">
                                    <h3 class="member-title">Ông Nguyễn Đức Bách<span>Giám đốc chuyên môn</span><span>Phó giáo sư, Tiến sĩ Sinh học</span></h3><!-- End .member-title -->
                                </div><!-- End .member-content -->
                            </div><!-- End .member -->
                        </div><!-- End .col-lg-3 -->

                        <div class="col-sm-6 col-lg-3">
                            <div class="member member-2 text-center">
                                <figure class="member-media">
                                    <img src="@/assets/images/about/member/4.jpg" alt="member photo">
                                </figure><!-- End .member-media -->
                                <div class="member-content">
                                    <h3 class="member-title">Ông Trần Bá Trung<span>Giám đốc kinh doanh</span><span>Thạc sĩ Quản trị Doanh nghiệp</span></h3><!-- End .member-title -->
                                </div><!-- End .member-content -->
                            </div><!-- End .member -->
                        </div><!-- End .col-lg-3 -->

                    </div><!-- End .row -->
                </div><!-- End .container -->
            </div><!-- End .bg-light-2 pt-6 pb-6 -->

            <div class="container">
                <div class="row">
                    <div class="col-lg-10 offset-lg-1">
                        <div class="brands-text text-center mx-auto mb-6">
                            <h2 class="title">Đối tác chiến lược của VNUA Phytopharma</h2><!-- End .title -->
                            <!-- <p>Phasellus hendrerit. Pellentesque aliquet nibh nec urna. In nisi neque, aliquet vel, dapibus id, mattis vel, nis</p> -->
                        </div><!-- End .brands-text -->
                        <div class="brands-display">
                            <div class="row justify-content-center">
                                <div class="col-6 col-sm-4 col-md-3">
                                    <a href="#" class="brand">
                                        <img src="@/assets/images/brands/1.jpg" alt="Brand Name">
                                    </a>
                                </div><!-- End .col-md-3 -->

                                <div class="col-6 col-sm-4 col-md-3">
                                    <a href="#" class="brand">
                                        <img src="@/assets/images/brands/2.jpg" alt="Brand Name">
                                    </a>
                                </div><!-- End .col-md-3 -->

                                <div class="col-6 col-sm-4 col-md-3">
                                    <a href="#" class="brand">
                                        <img src="@/assets/images/brands/3.jpg" alt="Brand Name">
                                    </a>
                                </div><!-- End .col-md-3 -->

                                <div class="col-6 col-sm-4 col-md-3">
                                    <a href="#" class="brand">
                                        <img src="@/assets/images/brands/4.jpg" alt="Brand Name">
                                    </a>
                                </div><!-- End .col-md-3 -->

                                <div class="col-6 col-sm-4 col-md-3">
                                    <a href="#" class="brand">
                                        <img src="@/assets/images/brands/5.jpg" alt="Brand Name">
                                    </a>
                                </div><!-- End .col-md-3 -->

                                <div class="col-6 col-sm-4 col-md-3">
                                    <a href="#" class="brand">
                                        <img src="@/assets/images/brands/6.jpg" alt="Brand Name">
                                    </a>
                                </div><!-- End .col-md-3 -->

                                <div class="col-6 col-sm-4 col-md-3">
                                    <a href="#" class="brand">
                                        <img src="@/assets/images/brands/7.jpg" alt="Brand Name">
                                    </a>
                                </div><!-- End .col-md-3 -->

                                <div class="col-6 col-sm-4 col-md-3">
                                    <a href="#" class="brand">
                                        <img src="@/assets/images/brands/8.jpg" alt="Brand Name">
                                    </a>
                                </div><!-- End .col-md-3 -->

                                <div class="col-6 col-sm-4 col-md-3">
                                    <a href="#" class="brand">
                                        <img src="@/assets/images/brands/9.jpg" alt="Brand Name">
                                    </a>
                                </div><!-- End .col-md-3 -->

                                <div class="col-6 col-sm-4 col-md-3">
                                    <a href="#" class="brand">
                                        <img src="@/assets/images/brands/10.jpg" alt="Brand Name">
                                    </a>
                                </div><!-- End .col-md-3 -->

                                <div class="col-6 col-sm-4 col-md-3">
                                    <a href="#" class="brand">
                                        <img src="@/assets/images/brands/11.jpg" alt="Brand Name">
                                    </a>
                                </div><!-- End .col-md-3 -->
                            </div><!-- End .row -->
                        </div><!-- End .brands-display -->
                    </div><!-- End .col-lg-10 offset-lg-1 -->
                </div><!-- End .row -->
            </div><!-- End .container -->
        </div><!-- End .page-content -->
    </main><!-- End .main -->
</template>
<script>
export default {
    name : 'About'
}
</script>