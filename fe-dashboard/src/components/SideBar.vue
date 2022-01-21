<template lang="">
        <nav id="sidebar">
            <div class="p-4 pt-5">
                <h2><router-link to="/" class="logo">VNUA phytopharma</router-link></h2>
                <ul class="list-unstyled components mb-5">
                    <li class="active">
                        <a href="#userSubmenu" data-toggle="collapse" aria-expanded="false" class="dropdown-toggle">Quản lý tài khoản</a>
                        <ul class="collapse list-unstyled" id="userSubmenu">
                            <li>
                                <router-link to="/profiles">Thông tin cá nhân</router-link>
                            </li>
                            <li>
                                <router-link to="/passwords">Thay đổi mật khẩu</router-link>
                            </li>
                        </ul>
                    </li>
                    <li>
                        <a href="#productsSubmenu" data-toggle="collapse" aria-expanded="false" class="dropdown-toggle">Quản lý sản phẩm</a>
                        <ul class="collapse list-unstyled" id="productsSubmenu">
                            <li>
                                <router-link to="/products">Danh sách sản phẩm</router-link>
                            </li>
                            <li>
                                <router-link to="/add-product">Thêm sản phẩm mới</router-link>
                            </li>
                        </ul>
                    </li>
                    <li>
                        <a href="#postsSubmenu" data-toggle="collapse" aria-expanded="false" class="dropdown-toggle">Quản lý bài viết</a>
                        <ul class="collapse list-unstyled" id="postsSubmenu">
                            <li>
                                <router-link to="/news">Danh sách tin tức</router-link>
                            </li>
                            <li>
                                <router-link to="/add-news">Thêm tin tức</router-link>
                            </li>
                            <li>
                                <router-link to="/events">Danh sách sự kiện</router-link>
                            </li>
                            <li>
                                <router-link to="/add-event">Thêm sự kiện</router-link>
                            </li>
                        </ul>
                    </li>
                    <li>
                        <a href="#customersSubmenu" data-toggle="collapse" aria-expanded="false" class="dropdown-toggle">Quản lý khách hàng</a>
                        <ul class="collapse list-unstyled" id="customersSubmenu">
                            <li>
                                <router-link to="/customers">Danh sách khách lẻ</router-link>
                            </li>
                            <li>
                                <router-link to="/stores">Danh sách đại lý</router-link>
                            </li>
                        </ul>
                    </li>
                    <li>
                        <a href="#videosSubmenu" data-toggle="collapse" aria-expanded="false" class="dropdown-toggle">Quản lý videos</a>
                        <ul class="collapse list-unstyled" id="videosSubmenu">
                            <li>
                                <router-link to="/videos">Danh sách videos</router-link>
                            </li>
                            <li>
                                <router-link to="/add-video">Thêm video mới</router-link>
                            </li>
                        </ul>
                    </li>
                    <li>
                        <a href="#staffsSubmenu" data-toggle="collapse" aria-expanded="false" class="dropdown-toggle">Quản lý nhân viên</a>
                        <ul class="collapse list-unstyled" id="staffsSubmenu">
                            <li>
                                <router-link to="/staffs">Danh sách nhân viên</router-link>
                            </li>
                            <li>
                                <router-link to="/add-staff">Tạo tài khoản nhân viên</router-link>
                            </li>
                        </ul>
                    </li>
                </ul>
                
                <div class="footer">
                    <button type="button" class="btn" @click="logout">
                        <img src="@/assets/logout-icon.png" alt="logout icon" width="50" />
                        <span class="btn-logout">Đăng xuất</span>
                    </button> 
                    
                    <p><!-- Link back to Colorlib can't be removed. Template is licensed under CC BY 3.0. -->
                        Copyright &copy; All rights reserved by VNUA Phytopharma
                        <!-- Link back to Colorlib can't be removed. Template is licensed under CC BY 3.0. --></p>
                </div>

            </div>
        </nav>
</template>
<script>
export default {
    name: "SideBar",
    methods: {
        async logout() {
            await this.$store.dispatch("users/logout");
            alert("Đăng xuất thành công!");
            await this.$router.push("/login")
        },
    },
}
</script>

<style scoped>
.btn-logout {
    font-weight: bold;
    color: white;
}
</style>