<template>
    <main class="main">
        <nav aria-label="breadcrumb" class="breadcrumb-nav">
            <div class="container">
                <ol class="breadcrumb">
                    <li class="breadcrumb-item"><router-link to="/">Trang chủ</router-link></li>
                    <li class="breadcrumb-item active" aria-current="page">Videos</li>
                </ol>
            </div><!-- End .container -->
        </nav><!-- End .breadcrumb-nav -->

        <div class="page-content">

            <div class="container">
                <h2 class="title mb-3 text-center">Videos</h2><!-- End .text-center -->
            </div><!-- End .container -->

            <div class="video-banner video-banner-poster text-center">
                <div class="container">
                  <div class="row demos">
                    <div
                        class="iso-item col-sm-6 col-md-4 col-lg-3 homepages"
                        v-for="video in videos" :key="video"
                    >
                      <iframe id="existing-iframe-example"
                              width="320" height="170"
                              :src="video.video_url"
                              frameborder="0"
                              allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                              allowfullscreen
                      ></iframe>
                    </div>
                  </div>
                </div><!-- End .container -->
                <Pagination
                    :length="totalItems"
                    :pageSize="12"
                    :pageIndex="pageIndex"
                    @change="changePage"
                />
            </div><!-- End .video-banner -->

            <hr class="mt-5 mb-4">
        </div><!-- End .page-content -->
    </main><!-- End .main -->
</template>
<script>
import {mapState } from "vuex";
import Pagination from "@/components/Pagination.vue";
export default {
    name : 'Video',
    components: {
        Pagination,
    },
    computed:{
        ...mapState("videos", ["totalItems","pageIndex", "videos"])
    },
    methods: {
        changePage(pageIndex) {
            this.$store.dispatch("videos/getVideos", { pageIndex });
        },
    },
    created(){
        this.$store.dispatch("videos/getVideos", {})
    }
}
</script>

<style scoped>
iframe {
    padding-right: 30px;
}
@media only screen and (max-width: 769px) {
    iframe {
        width: 100%;
        height: 100%;
    }
}
</style>