<template>
    <div class="bg-light-2 pt-6 pb-6 featured">
        <div class="container-fluid">
            <div class="heading heading-center mb-3">
                <h2 class="title">Sản phẩm nổi bật</h2><!-- End .title -->
                <p class="title-desc">Các dòng sản phẩm bán chạy nhất được nhiều khách hàng tin dùng.</p>
            </div><!-- End .heading -->
            <ProductsCarousel :products="randomProducts" />
        </div><!-- End .container-fluid -->
    </div><!-- End .bg-light-2 pt-4 pb-4 -->
</template>
<script>
import ProductsCarousel from "@/components/ProductsCarousel.vue";
import { mapState, mapActions } from 'vuex'
export default {
    name : 'FeatureProducts',
    components: {
        ProductsCarousel,
    },
    computed:{
        ...mapState("products",["randomProducts"])
    },
    methods:{
        ...mapActions("products", ["getRandomProducts"]),
    },
    created(){
        this.$store.dispatch("products/getRandomProducts", {});
    }
}

</script>
