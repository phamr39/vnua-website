<template>
    <div class="blog-posts pt-6">
        <div class="container">
            <div class="heading text-center">
                <h2 class="title">Tin tức mới nhất</h2>
                <p class="title-desc">Những tin tức, bài viết hay nhất.</p>
            </div>
            <NewsCarousel :listNews="newestNews" />
            <div class="more-container text-center">
                <router-link to="/tin-tuc" class="btn btn-outline-primary-2 btn-more"><span>Xem thêm nhiều bài viết</span><i class="icon-long-arrow-right"></i></router-link>
            </div>
        </div>
    </div>
    
</template>
<script>
import NewsCarousel from "@/components/NewsCarousel.vue";
import {  mapState } from "vuex";
export default {
    name : 'News',
    components: {
        NewsCarousel,
    },
    computed: {
        ...mapState("posts", ["newestNews"]),
    },
    created(){
        this.$store.dispatch("posts/getNewestNews", {});
    }
}
</script>