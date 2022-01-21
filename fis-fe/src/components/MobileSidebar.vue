<template >
    <div>
        <!-- Mobile Menu -->
        <div class="mobile-menu-overlay"></div><!-- End .mobil-menu-overlay -->

        <div class="mobile-menu-container mobile-menu-light">
            <div class="mobile-menu-wrapper">
                <span class="mobile-menu-close"><i class="icon-close"></i></span>

                <form action="#" method="get" class="mobile-search">
                    <label for="mobile-search" class="sr-only">Search</label>
                    <input type="search" class="form-control" name="mobile-search" id="mobile-search" placeholder="Tìm kiếm ..." required>
                    <button class="btn btn-primary" type="submit"><i class="icon-search"></i></button>
                </form>

                <nav class="mobile-nav">
                    <ul class="mobile-menu">
                        <li class="active">
                            <router-link to="/">Trang chủ</router-link>
                        </li>
                        <li>
                            <router-link to="/gioi-thieu">Giới thiệu</router-link>
                        </li>
                        <li>
                            <router-link to="/san-pham">Sản phẩm</router-link>
                        </li>
                        <!-- <li>
                            <router-link to="/dao-tao-chuyen-giao" class="sf-with-ul">Đào tạo chuyển giao</router-link>
                            <ul>
                                <li><a href="#">Chuyển giao công nghệ</a></li>
                                <li><a href="#">Lớp chuyên đề</a></li>
                            </ul>
                        </li> -->
                        <li>
                            <router-link to="/dang-ky-dai-ly" >Đăng ký đại lý</router-link>
                        </li>
                        <li>
                            <router-link to="/su-kien" >Sự kiện</router-link>
                        </li>
                        <li>
                            <router-link to="/tin-tuc" >Tin tức</router-link>
                        </li>
                        <li>
                            <router-link to="/videos" >Video</router-link>
                        </li>
                        <li>
                            <router-link to="/lien-he" >Liên hệ</router-link>
                        </li>
                    </ul>
                </nav><!-- End .mobile-nav -->
            </div><!-- End .mobile-menu-wrapper -->
        </div><!-- End .mobile-menu-container -->
    </div>
</template>
<script>
export default {
    name : 'MobileSidebar'
}
</script>