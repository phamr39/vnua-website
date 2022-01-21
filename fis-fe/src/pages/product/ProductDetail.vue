<template>
    <main class="main">
        <nav aria-label="breadcrumb" class="breadcrumb-nav border-0 mb-0">
            <div class="container d-flex align-items-center">
                <ol class="breadcrumb">
                    <li class="breadcrumb-item"><router-link to="/">Trang chủ</router-link></li>
                    <li class="breadcrumb-item"><router-link to="/san-pham">Sản phẩm</router-link></li>
                    <li class="breadcrumb-item active" aria-current="page">Chi tiết sản phẩm</li>
                </ol>
            </div><!-- End .container -->
        </nav><!-- End .breadcrumb-nav -->

        <div class="page-content">
            <div class="container">
                <div class="product-details-top mb-2">
                    <div class="row">
                        <div class="col-md-6">
                            <div class="product-gallery product-gallery-vertical">
                                <div class="row">
                                    <figure class="product-main-image">
                                        <img :src="mainImage" :alt="mainImage" />
                                    </figure><!-- End .product-main-image -->

                                    <div class="product-image-gallery" >
                                        <a class="product-gallery-item child-img" v-for="img, id in product.images" :key="id">
                                            <img :src="img.image_url" @click="getImageId(img)" />
                                        </a>
                                    </div><!-- End .product-image-gallery -->
                                </div><!-- End .row -->
                            </div><!-- End .product-gallery -->
                        </div><!-- End .col-md-6 -->

                        <div class="col-md-6">
                            <div class="product-details product-details-centered">
                                <h1 class="product-title">{{product.name}}</h1><!-- End .product-title -->

                                <div class="product-content">
                                    <span v-html="product.description"></span>
                                </div><!-- End .product-content -->

                                

                                <div class="product-details-action">
                                    <div class="details-action-col">
                                        <router-link to="/dang-ky-dai-ly" class="btn btn-outline-primary-2 btn-more"><span>đăng ký đại lý</span></router-link>
                                    </div><!-- End .details-action-col -->
                                    <div class="details-action-col">
                                        <a href="#signin-modal" data-toggle="modal" class="btn btn-outline-primary-2 btn-more"><span>đăng ký tư vấn</span></a>
                                    </div><!-- End .details-action-col -->
                                </div><!-- End .product-details-action -->

                                <div class="product-details-footer">
                                    <div class="social-icons social-icons-sm">
                                        <span class="social-label">Share:</span>
                                        <a :href="`https://www.facebook.com/sharer.php?u=http://vnuapharma.com.vn/sp/${product.handle}`" data-label="Facebook" onclick="window.open(this.href,this.title,'width=500,height=500,top=300px,left=300px');  return false;" rel="noopener noreferrer nofollow" target="_blank" class="social-icon" aria-label="Share on Facebook"><i class="icon-facebook-f"></i></a>
                                    </div>
                                </div><!-- End .product-details-footer -->
                            </div><!-- End .product-details -->
                        </div><!-- End .col-md-6 -->
                    </div><!-- End .row -->
                </div><!-- End .product-details-top -->

                <div class="product-details-tab">
                    <ul class="nav nav-pills justify-content-center" role="tablist">
                        <li class="nav-item">
                            <a class="nav-link active" id="product-desc-link" data-toggle="tab" href="#product-desc-tab" role="tab" aria-controls="product-desc-tab" aria-selected="true">Thông tin sản phẩm</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" id="product-info-link" data-toggle="tab" href="#product-info-tab" role="tab" aria-controls="product-info-tab" aria-selected="false">Hướng dẫn sử dụng</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" id="product-shipping-link" data-toggle="tab" href="#product-shipping-tab" role="tab" aria-controls="product-shipping-tab" aria-selected="false">Vận chuyển & Trả hàng</a>
                        </li>
                    </ul>
                    <div class="tab-content">
                        <div class="tab-pane fade show active" id="product-desc-tab" role="tabpanel" aria-labelledby="product-desc-link">
                            <div class="product-desc-content">
                                <span v-html="product.product_info"></span>
                            </div><!-- End .product-desc-content -->
                        </div><!-- .End .tab-pane -->
                        <div class="tab-pane fade" id="product-info-tab" role="tabpanel" aria-labelledby="product-info-link">
                            <div class="product-desc-content">
                                <span v-html="product.user_manual"></span>
                            </div><!-- End .product-desc-content -->
                        </div><!-- .End .tab-pane -->
                        <div class="tab-pane fade" id="product-shipping-tab" role="tabpanel" aria-labelledby="product-shipping-link">
                            <div class="product-desc-content">
                                <ul>
                                    <li>
                                        VNUA Phytopharma xác nhận và xử lý đơn hàng của các Đại lý từ thứ 2 đến thứ 6 hàng tuần, trong giờ làm việc từ 8h30 – 17h00.
                                    </li>
                                    <li>Khi nhận được đơn đặt hàng của Đại lý, chúng tôi sẽ giao hàng đến khu vực được Đại lý chỉ định.</li>
                                    <li>Các sản phẩm được kiểm tra và đóng gói đúng quy cách, đảm bảo thẩm mỹ. Chúng tôi cũng chọn lựa nhà vận chuyển phù hợp nhằm không làm ảnh hưởng đến chất lượng sản phẩm trong quá trình vận chuyển.</li>
                                    <li>VNUA Phytopharma cam kết hàng hóa được chúng tôi ủy quyền phân phối chính hãng, có đầy đủ giấy tờ pháp lý để lưu hành. Sản phẩm an toàn cho người sử dụng.</li>
                                </ul>
                                <p>Xem thêm chính sách đổi trả tại <router-link to="/doi-tra">Thông tin đổi trả</router-link></p>
                            </div><!-- End .product-desc-content -->
                        </div><!-- .End .tab-pane -->
                    </div><!-- End .tab-content -->
                </div><!-- End .product-details-tab -->
            </div><!-- End .container -->
        </div><!-- End .page-content -->
        <CustomerForm />
    </main><!-- End .main -->
</template>
<script>
import CustomerForm from "@/components/CustomerForm.vue";
import { mapMutations, mapState } from "vuex";
export default {
    name: 'ProductDetail',
    components: {
        CustomerForm,
    },
    data() { 
        return {
            imgId: 0,
        }
    },
    computed: {
        ...mapState("products", ["product","firstCate", "mainImage"]),
    },
    methods: {
        ...mapMutations("products", ["setMainImage"]),
        getImageId(img){
            this.setMainImage(img.image_url)
        },
    },
    created(){
        console.log("ProductDetail created", this.$route.params)
        this.$store.dispatch("products/getProductByHandle", this.$route.params.name);
    },
}
</script>

<style scoped>
.child-img {
    cursor: pointer;
    height: 25%;
    padding:0;
}

.child-img img:hover {
    border: 2px solid #27ae60;
}
</style>