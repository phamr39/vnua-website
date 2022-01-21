<template>
    <div class="container-fluid pt-6 pb-0">
        <div class="heading heading-center mb-3">
            <h2 class="title">Sản phẩm mới</h2><!-- End .title -->
            <p class="title-desc">Những nghiên cứu mới nhất.</p>
        </div><!-- End .heading -->

        <div class="tab-content">
            <div class="tab-pane p-0 fade show active" id="new-women-tab" role="tabpanel" aria-labelledby="new-women-link">
                <div class="products">
                    <div class="row justify-content-center">
                        <div class="col-6 col-md-4 col-lg-3 col-xl-5col" v-for="product in newArrivals" :key="product.id">
                            <div class="product product-7 text-center">
                                <figure class="product-media">
                                    <span class="product-label label-circle label-new">New</span>
                                    <router-link :to="'/sp/' + product.name">
                                        <img :src="product.images[0].image_url" alt="Product image" class="product-image">
                                        <img :src="product.images[1].image_url" alt="Product image" class="product-image-hover">
                                    </router-link>

                                    <div class="product-action">
                                        <a href="#signin-modal" data-toggle="modal" class="btn-product"><span>đăng ký tư vấn</span></a>
                                    </div><!-- End .product-action -->
                                </figure><!-- End .product-media -->

                                <div class="product-body">
                                    <div class="product-cat" v-for="category in product.categories" :key="category.id">
                                        <a href="#">{{category.name}}</a>
                                    </div><!-- End .product-cat -->
                                    <h3 class="product-title"><router-link :to="'/sp/' + product.name">{{product.name}}</router-link></h3><!-- End .product-title -->
                                </div><!-- End .product-body -->
                            </div><!-- End .product -->
                        </div><!-- End .col-sm-6 col-md-4 col-lg-3 -->
                    </div><!-- End .row -->
                </div><!-- End .products -->
            </div><!-- .End .tab-pane -->
        </div><!-- End .tab-content -->

        <div class="more-container text-center mt-2">
            <router-link to="/san-pham" class="btn btn-outline-dark-3 btn-more"><span>Xem thêm sản phẩm</span><i class="icon-long-arrow-right"></i></router-link>
        </div><!-- End .more-container -->
    </div><!-- End .container-fluid -->
</template>
<script>
import { mapState, mapActions } from 'vuex'
export default {
    name: 'NewArrivals',
    computed:{
        ...mapState("products",["newArrivals"])
    },
    methods:{
        ...mapActions("products", ["getNewArrivals"]),
    },
    created(){
        this.$store.dispatch("products/getNewArrivals", {});
    }
}
</script>
