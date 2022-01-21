<template>
    <main class="main">
        <IntroSlide />
        <Categories/>
        <FeatureProducts />
        <NewArrivals />
        <router-link to="/dang-ky-dai-ly">
            <div class="page-header text-center">
                <div class="elementor-background-overlay"></div>
                <div class="container">
                    <h1 class="page-title text-white">Trao Giá Trị – Nhận Niềm Tin<span class="text-white">Trở thành đối tác của VNUA Phytopharma</span></h1>
                </div><!-- End .container -->
                
            </div><!-- End .page-header -->
        </router-link>
        <VideoYoutube />
        <Blog />
        <CustomerForm />
    </main>
</template>
<script>
import IntroSlide from "./IntroSlide.vue";
import FeatureProducts from "./FeatureProducts.vue";
import NewArrivals from "./NewArrivals.vue";
import VideoYoutube from "./VideoYoutube.vue";
import Blog from "./Blog.vue";
import Categories from "./Categories.vue";
import CustomerForm from "@/components/CustomerForm.vue";
export default {
    name : 'HomePage',
    components: {
        IntroSlide,
        FeatureProducts,
        NewArrivals,
        VideoYoutube,
        Blog,
        CustomerForm,
        Categories
    },
}
</script>

<style scoped>
.elementor-background-overlay {
    background-image: url('../../assets/images/backgrounds/news-bg-1.png');
    background-position: center center;
    background-repeat: no-repeat;
    background-size: cover;
    opacity: 0.15;
    transition: border-radius 0.3s, opacity 0.3s;
    width: 100%;
    height: 100%;
    position: absolute;
}
</style>
