<template>
    <footer class="footer">
        <div class="footer-middle">
            <div class="container">
                <div class="row">
                    <div class="col-sm-6 col-lg-3">
                        <div class="widget widget-about">
                            <h4 class="widget-title">giới thiệu</h4><!-- End .widget-title -->
                            <p>Công ty Cổ phần KHCN Dược liệu Việt Nam</p>
                            <div class="social-icons">
                                <a href="https://www.facebook.com/vnuaphytopharma" class="social-icon" title="Facebook" target="_blank"><i class="icon-facebook-f"></i></a>
                                <a href="#" class="social-icon" title="Instagram" target="_blank"><i class="icon-instagram"></i></a>
                                <a href="#" class="social-icon" title="Youtube" target="_blank"><i class="icon-youtube"></i></a>
                            </div><!-- End .soial-icons -->
                        </div><!-- End .widget about-widget -->
                    </div><!-- End .col-sm-6 col-lg-3 -->
                    <div class="col-sm-6 col-lg-3">
                        <div class="widget">
                            <h4 class="widget-title">văn phòng đại diện</h4><!-- End .widget-title -->
                            <ul class="widget-list">
                                <li><a href="#"><i class="icon-map-marker"></i> Học Viện Nông nghiệp Việt Nam, Ngô Xuân Quảng, Trâu Quỳ, Gia Lâm, Hà Nội</a></li>
                            </ul><!-- End .widget-list -->
                        </div><!-- End .widget -->
                    </div><!-- End .col-sm-6 col-lg-3 -->
                    <div class="col-sm-6 col-lg-3">
                        <div class="widget">
                            <h4 class="widget-title">liên hệ</h4><!-- End .widget-title -->
                            <ul class="widget-list">
                                <li><a href="tel:#"><i class="icon-phone"></i> 0982 27 47 84</a></li>
                                <li><a href="tel:#"><i class="icon-phone"></i> 0966 10 15 19</a></li>
                                <li><a href="mailto:#"><i class="icon-envelope"></i> duoclieuvietnam@vnua.edu.vn</a></li>
                            </ul><!-- End .widget-list -->
                        </div><!-- End .widget -->
                    </div><!-- End .col-sm-6 col-lg-3 -->
                    <div class="col-sm-6 col-lg-3">
                        <div class="widget">
                            <h4 class="widget-title">quy định - chính sách</h4><!-- End .widget-title -->
                            <ul class="widget-list">
                                <!-- <li><a href="#">Chính sách hợp tác</a></li> 
                                <li><a href="#">Chính sách chuyển giao</a></li> -->
                                <li><router-link to="/mua-hang">Chính sách bán hàng</router-link></li>
                                <li><router-link to="/doi-tra">Chính sách đổi trả</router-link></li>
                                <li><router-link to="/quyen-loi">Quyền Lợi Đối tác – Đại lý</router-link></li>
                            </ul><!-- End .widget-list -->
                        </div><!-- End .widget -->
                    </div><!-- End .col-sm-6 col-lg-3 -->
                </div><!-- End .row -->
            </div><!-- End .container -->
        </div><!-- End .footer-middle -->
            <div class="footer-bottom">
                    <div class="container">
                        <div class="footer-copyright">
                            <p>Công ty Cổ phần KHCN Dược liệu Việt Nam</p>
                            <p>Mã số thuế: 0109694807</p>
                            <p>Copyright © 2021 VNUA Phytopharma. All Rights Reserved.</p>
                        </div><!-- End .footer-copyright -->
                        <figure class="footer-payments">
                            <img src="@/assets/images/logo-footer.png" alt="Payment methods" width="75">
                        </figure><!-- End .footer-payments -->
                    </div><!-- End .container -->
                </div><!-- End .footer-bottom -->
    </footer><!-- End .footer -->
</template>
<script>
export default {
    name : 'MainFooter',
}
</script>