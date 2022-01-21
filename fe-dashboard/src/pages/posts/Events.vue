<template>
<div class="card border-success mb-3">
        <div class="card-header text-center">
            Danh sách sự kiện
        </div>
        <div class="cart-body">
            <div class="p-5">
                <table class="table table table-striped">
                    <thead>
                        <tr>
                        <th scope="col">STT.</th>
                        <th scope="col">Bài viết</th>
                        <th scope="col">Tác giả</th>
                        <th scope="col">Danh mục</th>
                        <th scope="col"></th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="post, id in events" :key="post.id">
                            <td>{{id + 1}}</td>
                            <td>{{post.title}}</td>
                            <td>{{post.author}}</td>
                            <td>
                                <div v-for="topic in post.topics" :key="topic.id">
                                    <span>{{topic.name}};</span>
                                </div>
                            </td>
                            <td>
                                <button type="button" @click="deletePost(post.id)" class="btn btn-outline-danger">Xóa</button>
                            </td>
                        </tr>
                    </tbody>
                </table>
                <Pagination
                    :length="totalItems"
                    :pageSize="4"
                    :pageIndex="pageIndex"
                    @change="changePage"
                />
            </div>
        </div>
    </div>
</template>
<script>
import Pagination from "@/components/Pagination.vue";
import apiPost from "@/services/posts.service";
import {  mapState } from "vuex";
export default {
    name: 'Events',
    components : {
        Pagination,
    },
        computed: {
        ...mapState("posts", ["events", "pageIndex", "totalItems"]),
    },
    methods: {
        changePage(pageIndex){
            this.$store.dispatch("posts/getAllEvents", { pageIndex })
        },
        async deletePost(postId) {
            await apiPost.deletePost(postId);
            window.location.reload();
        },
    },
    created(){
        this.$store.dispatch("posts/getAllEvents", {});
    }
}
</script>
