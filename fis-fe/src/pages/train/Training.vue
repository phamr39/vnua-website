<template>
    <main class="main">
        <nav aria-label="breadcrumb" class="breadcrumb-nav border-0 mb-0">
            <div class="container">
                <ol class="breadcrumb">
                    <li class="breadcrumb-item"><router-link to="/">Trang chủ</router-link></li>
                    <li class="breadcrumb-item active" aria-current="page">Đào tạo chuyển giao</li>
                </ol>
            </div><!-- End .container -->
        </nav><!-- End .breadcrumb-nav -->
      <div class="container">
        <div class="page-header text-center">
          <h1 class="page-title text-white">Đào tạo chuyển giao<span class="text-white">Các Khóa đào tạo Chuyển giao công nghệ – Chuyển giao chuyên đề</span></h1>
        </div><!-- End .page-header -->
        <div class="product mb-3 mt-3">
          <div class="product product-list">
            <div class="row">
              <div class="col-6 col-lg-3">
                <figure class="product-media">
                  <router-link to="/dao-tao-chuyen-giao/chuyen-giao-cong-nghe">
                    <img src="@/assets/images/train/chuyengiao.png" alt="Product image" class="product-image">
                  </router-link>
                </figure><!-- End .product-media -->
              </div><!-- End .col-sm-6 col-lg-3 -->
              <div class="col">
                <div class="product-body product-action-inner">
                  <h3 class="title"><router-link to="/dao-tao-chuyen-giao/chuyen-giao-cong-nghe">Chuyển giao công nghệ</router-link></h3><!-- End .product-title -->
                  <div class="product-content">
                    <p>Các Khóa đào tạo Chuyển giao công nghệ – Chuyển giao chuyên đề của VNUA Phytopharma được xây dựng cho các Đối tác – Đại lý là chủ cơ sở Spa, TMV, các Trung tâm Chăm sóc sức khỏe, Phòng khám, Nhà thuốc,… và các cá nhân có nhu cầu hoạt động kinh doanh, làm việc trong lĩnh vực Da liễu, Chăm sóc sức khỏe, sắc đẹp,… </p>
                  </div><!-- End .product-content -->
                </div><!-- End .product-body -->
              </div><!-- End .col-lg-6 -->
            </div><!-- End .row -->
          </div>
          <div class="product product-list">
            <div class="row">
              <div class="col-6 col-lg-3">
                <figure class="product-media">
                  <router-link to="/lien-he">
                    <img src="@/assets/images/train/p.png" alt="Product image" class="product-image">
                  </router-link>
                </figure><!-- End .product-media -->
              </div><!-- End .col-sm-6 col-lg-3 -->
              <div class="col">
                <div class="product-body product-action-inner">
                  <h3 class="title"><router-link to="/lien-he">Lớp chuyên đề</router-link></h3><!-- End .product-title -->
                  <div class="product-content">
                    <p>Các Khóa đào tạo Chuyển giao công nghệ – Chuyển giao chuyên đề của VNUA Phytopharma được xây dựng cho các Đối tác – Đại lý là chủ cơ sở Spa, TMV, các Trung tâm Chăm sóc sức khỏe, Phòng khám, Nhà thuốc,… và các cá nhân có nhu cầu hoạt động kinh doanh, làm việc trong lĩnh vực Da liễu, Chăm sóc sức khỏe, sắc đẹp,… </p>
                  </div><!-- End .product-content -->
                </div><!-- End .product-body -->
              </div><!-- End .col-lg-6 -->
            </div><!-- End .row -->
          </div>
        </div>
      </div><!-- End .container -->
    </main>
</template>
<script>
export default {
    name : 'Training'
}
</script>
