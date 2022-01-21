<template>
        <main class="main">
            <nav aria-label="breadcrumb" class="breadcrumb-nav border-0 mb-0">
                <div class="container">
                    <ol class="breadcrumb">
                        <li class="breadcrumb-item"><router-link to="/">Trang chủ</router-link></li>
                        <li class="breadcrumb-item"><router-link to="/tin-tuc">Tin tức</router-link></li>
                        <li class="breadcrumb-item active" aria-current="page">Chi tiết</li>
                    </ol>
                </div><!-- End .container -->
            </nav><!-- End .breadcrumb-nav -->

            <div class="page-content">
                <figure class="entry-media">
                    <img :src="post.thumbnail" alt="image desc">
                </figure><!-- End .entry-media -->
                <div class="container">
                    <article class="entry single-entry entry-fullwidth">
                        <div class="row">
                            <div class="col-lg-11">
                                <div class="entry-body">
                                    <div class="entry-meta">
                                        <span class="entry-author">
                                            by <a href="#">{{post.author}}</a>
                                        </span>
                                        <span class="meta-separator">|</span>
                                        <a href="#">{{post.created_at}}</a>
                                    </div><!-- End .entry-meta -->

                                    <h2 class="entry-title entry-title-big">
                                        {{post.title}}
                                    </h2><!-- End .entry-title -->

                                    <div class="entry-content editor-content">
                                        <span v-html="post.content"></span>
                                    </div><!-- End .entry-content -->

                                    <div class="entry-footer row no-gutters">
                                        <div class="col">
                                            <div class="entry-tags">
                                                <span >Topics:</span> <a href="#" v-for="topic in post.topics" :key="topic.id">{{topic.name}}</a>
                                            </div><!-- End .entry-tags -->
                                        </div><!-- End .col -->
                                    </div><!-- End .entry-footer row no-gutters -->
                                </div><!-- End .entry-body -->
                            </div><!-- End .col-lg-11 -->

                            <div class="col-lg-1 order-lg-first mb-2 mb-lg-0">
                                <div class="sticky-content">
                                    <div class="social-icons social-icons-colored social-icons-vertical">
                                        <span class="social-label">SHARE:</span>
                                        <a :href="'https://www.facebook.com/sharer.php?u=http://vnuapharma.com.vn/chi-tiet-tin-tuc/'+ post.id" data-label="Facebook" onclick="window.open(this.href,this.title,'width=500,height=500,top=300px,left=300px');  return false;" rel="noopener noreferrer nofollow" target="_blank" class="social-icon social-facebook" title="Facebook" aria-label="Share on Facebook"><i class="icon-facebook-f"></i></a>
                                    </div><!-- End .soial-icons -->
                                </div><!-- End .sticky-content -->
                            </div><!-- End .col-lg-1 -->
                        </div><!-- End .row -->
                    </article><!-- End .entry -->

                </div><!-- End .container -->
            </div><!-- End .page-content -->
        </main><!-- End .main -->
</template>
<script>
import {  mapState} from "vuex";
export default {
    name : 'NewsDetail',
    computed:{
        ...mapState("posts", ["post"]),
    },
    created(){
        this.$store.dispatch("posts/getPostById", this.$route.params.id)
    }
}
</script>
