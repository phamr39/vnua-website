<template>
    <div class="video-banner bg-light pt-5 pb-5">
        <div class="container align-items-center">
            <div class="video-banner-box bg-white">
                <div class="row align-items-center">
                    <div class="col-md-6 mb-3 mb-md-0">
                        <div class="video-box-content">
                            <h3 class="video-banner-title h1"><span class="text-primary">Video mới nhất</span><strong>{{video.title}}</strong></h3><!-- End .video-banner-title -->
                            <span v-html="video.description"></span><br>
                            <router-link to="/videos" class="btn btn-outline-primary"><span>Xem nhiều videos</span><i class="icon-long-arrow-right"></i></router-link>
                        </div><!-- End .video-box-content -->
                    </div><!-- End .col-md-6 -->
                    <div class="col-md-6">
                        <iframe id="existing-iframe-example"
                                width="574" height="316"
                                :src="video.video_url"
                                frameborder="0"
                        ></iframe>
                        
                    </div><!-- End .col-md-6 -->
                </div><!-- End .row -->
            </div><!-- End .video-banner-box -->
        </div><!-- End .container -->
    </div><!-- End .video-banner bg-image -->
</template>
<script>
import {mapState } from "vuex";
export default {
    name : 'VideoYoutube',
    computed: {
        ...mapState("videos",["video"])
    },
    created(){
        this.$store.dispatch("videos/getNewestVideo", {})
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
