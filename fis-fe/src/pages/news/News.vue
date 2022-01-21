<template>
    <main class="main">
        <nav aria-label="breadcrumb" class="breadcrumb-nav mb-3">
            <div class="container">
                <ol class="breadcrumb">
                    <li class="breadcrumb-item"><router-link to="/">Trang chủ</router-link></li>
                    <li class="breadcrumb-item active" aria-current="page">Tin tức</li>
                </ol>
            </div><!-- End .container -->
        </nav><!-- End .breadcrumb-nav -->

        <div class="page-content">
            <div class="container">
                <div class="row">
                    <div class="col-lg-9" >
                        <article class="entry" v-for="post in listNews" :key="post.id">
                            <figure class="entry-media">
                                <router-link :to="'/chi-tiet-tin-tuc/' + post.id" >
                                    <img :src="post.thumbnail" alt="image desc">
                                </router-link>
                            </figure><!-- End .entry-media -->

                            <div class="entry-body">
                                <div class="entry-meta">
                                    <span class="entry-author">
                                        by <a href="#">{{post.author}}</a>
                                    </span>
                                    <span class="meta-separator">|</span>
                                    <a href="#">{{post.created_at}}</a>
                                </div><!-- End .entry-meta -->

                                <h2 class="entry-title">
                                    <router-link :to="'/chi-tiet-tin-tuc/' + post.id">Bài viết : {{post.title}}</router-link>
                                </h2><!-- End .entry-title -->

                                <div class="entry-cats">
                                    in <a href="#" v-for="topic in post.topics" :key="topic.id">{{topic.name}}, </a>
                                </div><!-- End .entry-cats -->

                                <div class="entry-content">
                                    <p>{{post.description}}</p>
                                    <router-link :to="'/chi-tiet-tin-tuc/' + post.id" class="read-more">Xem thêm</router-link>
                                </div><!-- End .entry-content -->
                            </div><!-- End .entry-body -->
                        </article><!-- End .entry -->
                        <Pagination
                            :length="totalItems"
                            :pageSize="4"
                            :pageIndex="pageIndex"
                            @change="changePage"
                        />
                    </div><!-- End .col-lg-9 -->
                    <RightBar 
                        :posts="randomNews"
                        @setPost="setPostByTopic"
                        @search="searchPost"
                    />
                </div><!-- End .row -->
            </div><!-- End .container -->
        </div><!-- End .page-content -->
    </main><!-- End .main -->
</template>
<script>
import RightBar from "@/components/RightBar.vue";
import Pagination from "@/components/Pagination.vue";
import {  mapState , mapActions} from "vuex";
export default {
    name : 'News',
    components: {
        RightBar,
        Pagination,
    },
    computed: {
        ...mapState("posts", ["listNews", "order", "pageIndex", "totalItems", "randomNews"]),
    },
    methods: {
        changePage(page){
            this.$store.dispatch("posts/getAllNews", { pageIndex : page })
        },
        setPostByTopic(topicId){
            this.$store.dispatch("posts/getAllNews", { topicId : topicId })
        },
        ...mapActions("posts", ["getAllNews"]),
    },
    created(){
        if(this.$route.query.q !== undefined){
            this.$store.dispatch("posts/getAllNews", { search : this.$route.query.q });
        }
        this.$store.dispatch("posts/getAllNews", {});
        this.$store.dispatch("posts/getRandomNews", {});
    }
}
</script>