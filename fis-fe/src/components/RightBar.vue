<template>
    <aside class="col-lg-3">
        <div class="sidebar">
            <div class="widget widget-search">
                <h3 class="widget-title">Tìm kiếm</h3><!-- End .widget-title -->

                <form action="#">
                    <label for="ws" class="sr-only">Tìm kiếm</label>
                    <input type="search" class="form-control" name="q" id="q" placeholder="Tìm kiếm">
                    <button type="submit" class="btn"><i class="icon-search"></i><span class="sr-only">Search</span></button>
                </form>
            </div><!-- End .widget -->

            <div class="widget widget-cats">
                <h3 class="widget-title">Danh mục bài viết</h3><!-- End .widget-title -->

                <ul>
                    <li v-for="topic in topics" :key="topic.id">
                        <a href="#" @click="getPost(topic.id, $event)">{{topic.name}}</a>
                    </li>
                </ul>
            </div><!-- End .widget -->

            <div class="widget">
                <h3 class="widget-title">Bài viết nổi bật</h3><!-- End .widget-title -->

                <ul class="posts-list">
                    <li v-for="post in posts" :key="post.id">
                        <figure>
                            <router-link :to="'/chi-tiet-tin-tuc/' + post.id">
                                <img :src="post.thumbnail" alt="post">
                            </router-link>
                        </figure>

                        <div>
                            <span>{{post.created_at}}</span>
                            <h4><router-link :to="'/chi-tiet-tin-tuc/' + post.id">{{post.description}}</router-link></h4>
                        </div>
                    </li>
                </ul><!-- End .posts-list -->
            </div><!-- End .widget -->

            <!-- <div class="widget widget-banner-sidebar">
                <div class="banner-sidebar banner-overlay">
                    <a href="#">
                        <img src="@/assets/images/blog/sidebar/banner.jpg" alt="banner">
                    </a>
                </div>
            </div> -->
        </div><!-- End .sidebar -->
    </aside><!-- End .col-lg-3 -->
</template>
<script>
import {  mapState } from "vuex";
export default {
    name : 'RightBar',
    props: ["posts", "key"],
    data() {
    },
    computed:{
        ...mapState("posts", ["topics"])
    },
    methods:{
        getPost(id, event){
            event.preventDefault();
            this.$emit("setPost", id);
        },
        search(event){
            this.$emit("search", event.target.value)
        }
    },
    created(){
        this.$store.dispatch("posts/getTopics")
    }
}
</script>
