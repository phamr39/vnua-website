<template>
    <div class="container-fluid pt-6 pb-6">
        <div class="row justify-content-center">
            <div class="col-md-6 col-lg-4">
              <router-link :to="{ name : 'Products', query: { category_id : 1 }}">
                <div class="banner banner-overlay text-white">
                    <img src="@/assets/images/category/2.png" alt="Banner">
                    <div class="banner-content banner-content-overlay text-center">
                        <h3 class="banner-title">Dược phẩm, thực phẩm BVSK</h3><!-- End .banner-title -->
                    </div>
                </div><!-- End .banner -->
              </router-link>
            </div><!-- End .col-lg-4 -->
            <div class="col-md-6 col-lg-4">
              <router-link :to="{ name : 'Products', query: { category_id : 2 }}">
                <div class="banner banner-overlay color-grey">
                        <img src="@/assets/images/category/6.png" alt="Banner">
                    <div class="banner-content banner-content-overlay text-center">
                        <h3 class="banner-title">Thực phẩm, thức uống dinh dưỡng</h3><!-- End .banner-title -->
                    </div>
                </div><!-- End .banner -->
              </router-link>
            </div><!-- End .col-lg-4 -->
            <div class="col-md-6 col-lg-4">
              <router-link :to="{ name : 'Products', query: { category_id : 3 }}">
                <div class="banner banner-overlay text-white">
                        <img src="@/assets/images/category/3.png" alt="Banner">
                    <div class="banner-content banner-content-overlay text-center">
                        <h3 class="banner-title">Mỹ phẩm</h3><!-- End .banner-title -->
                    </div>
                </div><!-- End .banner -->
              </router-link>
            </div><!-- End .col-lg-4 -->
        </div><!-- End .row -->
      <div class="row justify-content-center">
        <div class="col-md-6 col-lg-4">
          <router-link to="/dao-tao-chuyen-giao">
          <div class="banner banner-overlay text-white">
            <img src="@/assets/images/category/1.png" alt="Banner">
            <div class="banner-content banner-content-overlay text-center">
              <h3 class="banner-title">Nguyên liệu dược</h3><!-- End .banner-title -->
            </div>
          </div><!-- End .banner -->
          </router-link>
        </div><!-- End .col-lg-4 -->
        <div class="col-md-6 col-lg-4">
          <router-link to="/dao-tao-chuyen-giao">
          <div class="banner banner-overlay color-grey">
            <img src="@/assets/images/category/4.png" alt="Banner">
            <div class="banner-content banner-content-overlay text-center">
              <h3 class="banner-title">Đào tạo – cấp chứng chỉ</h3><!-- End .banner-title -->
            </div>
          </div><!-- End .banner -->
          </router-link>
        </div><!-- End .col-lg-4 -->
        <div class="col-md-6 col-lg-4">
          <router-link to="/dao-tao-chuyen-giao/chuyen-giao-cong-nghe">
            <div class="banner banner-overlay text-white">
              <img src="@/assets/images/category/5.png" alt="Banner">
              <div class="banner-content banner-content-overlay text-center">
                <h3 class="banner-title">Chuyển giao công nghệ</h3><!-- End .banner-title -->
              </div>
            </div>
          </router-link>
        </div><!-- End .col-lg-4 -->
      </div><!-- End .row -->
        </div><!-- End .container-fluid -->
</template>
<script>
export default {
    name :'Categories',
}
</script>