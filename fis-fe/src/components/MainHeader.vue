<template>
    <header class="header header-6">
        <div class="header-middle">
            <div class="container">
                <div class="header-left">
                        <div class="widget-search">
                            <form method="get">
                                <div class="header-search-wrapper search-wrapper-wide">
                                    <button class="btn btn-primary" type="submit">
                                        <router-link :to="{ name : 'Products', query : { q: searchKeyword }}">
                                            <i class="icon-search i-search"></i>
                                        </router-link>
                                    </button>
                                    <input 
                                        type="search" 
                                        class="form-control" 
                                        placeholder="Tìm kiếm" 
                                        v-model="searchKeyword"
                                        @keyup.enter="searchProducts"
                                        required
                                        >
                                </div><!-- End .header-search-wrapper -->
                            </form>
                        </div><!-- End .header-search -->
                </div>
                <div class="header-center">
                    <router-link to="/" class="logo">
                        <img src="@/assets/images/logo-text.png" alt="Logo" width="150" height="20">
                    </router-link>
                </div><!-- End .header-left -->
                <div class="header-right">
                    <div class="social-icons">
                        <a href="https://www.facebook.com/vnuaphytopharma" class="social-icon" title="Facebook" target="_blank"><i class="icon-facebook-f"></i></a>
                        <a href="#" class="social-icon" title="Instagram" target="_blank"><i class="icon-instagram"></i></a>
                        <a href="#" class="social-icon" title="Youtube" target="_blank"><i class="icon-youtube"></i></a>
                    </div><!-- End .soial-icons -->
                </div><!-- End .header -->
            </div>
        </div>

        <div class="header-bottom sticky-header">
            <div class="container">
                <div class="header-left">
                    <nav class="main-nav">
                        <ul class="menu sf-arrows">
                            <li :class="{ 'active': $route.path == '/' }">
                                <router-link to="/">Trang chủ</router-link>
                            </li>
                            <li :class="{ 'active': $route.path == '/gioi-thieu' }">
                                <router-link to="/gioi-thieu">Giới thiệu</router-link>
                            </li>
                            <li :class="{ 'active': $route.path == '/san-pham' }">
                                <router-link to="/san-pham" class="sf-with-ul">Sản phẩm</router-link>
                                <ul>
                                  <li><router-link :to="{ name: 'Products', query: { category_id: 1, q: '' } }">Dược phẩm</router-link></li>
                                  <li><router-link :to="{ name: 'Products', query: { category_id: 2, q: '' } }">Mỹ phẩm</router-link></li>
                                  <li><router-link :to="{ name: 'Products', query: { category_id: 3, q: '' } }">Thực phẩm chức năng</router-link></li>
                                </ul>
                            </li>
                            <li :class="{ 'active': $route.path == '/dao-tao-chuyen-giao' }">
                                <router-link to="/dao-tao-chuyen-giao" class="sf-with-ul">Đào tạo chuyển giao</router-link>
                                <ul>
                                    <li><router-link to="/dao-tao-chuyen-giao/chuyen-giao-cong-nghe">Chuyển giao công nghệ</router-link></li>
                                    <li><router-link to="/lien-he" >Lớp chuyên đề</router-link></li>
                                </ul>
                            </li>
                            <li :class="{ 'active': $route.path == '/dang-ky-dai-ly' }">
                                <router-link to="/dang-ky-dai-ly" >Đăng ký đại lý</router-link>
                            </li>
                            <li :class="{ 'active': $route.path == '/su-kien' }">
                                <router-link to="/su-kien" >Sự kiện</router-link>
                            </li>
                            <li :class="{ 'active': $route.path == '/tin-tuc' }">
                                <router-link to="/tin-tuc" >Tin tức</router-link>
                            </li>
                            <li :class="{ 'active': $route.path == '/videos' }">
                                <router-link to="/videos" >Video</router-link>
                            </li>
                            <li :class="{ 'active': $route.path == '/lien-he' }">
                                <router-link to="/lien-he" >Liên hệ</router-link>
                            </li>
                        </ul><!-- End .menu -->
                    </nav><!-- End .main-nav -->
                    <button class="mobile-menu-toggler">
                        <span class="sr-only">Toggle mobile menu</span>
                        <i class="icon-bars"></i>
                    </button>
                </div><!-- End .header-left -->
            </div><!-- End .container -->
        </div><!-- End .header-bottom -->
    </header>
</template>
<script>
import {  mapState } from "vuex";
export default {
    name: 'MainHeader',
    data() {
        return {
            searchKeyword: "",
        }
    },
    computed: {
        ...mapState("categories", ["rootCate","childCate"]),
    },
    methods: {
        searchProducts(e) {
            console.log(e.target.value)
            //this.$router.push({ name:'Products' })
        },
        setChildCate(id){
            this.$store.dispatch("categories/getCategories", id);
        },
    },
    created(){
        this.$store.dispatch("categories/getCategories", "");
    }
}
</script>

<style>
.i-search {
  color: #fff;
}

.i-search:hover {
  color: #27ae60;
}
</style>