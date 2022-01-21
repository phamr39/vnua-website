<template lang="">
    <div class="card border-success mb-3">
        <div class="card-header text-center">
            Danh sách videos
        </div>
        <div class="cart-body">
            <div class="p-5">
                <table class="table table table-striped">
                    <thead>
                        <tr>
                        <th scope="col">STT.</th>
                        <th scope="col">Tiêu đề</th>
                        <th scope="col">Link</th>
                        <th scope="col"></th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="video, id in videos" :key="video.id">
                            <td>{{id + 1}}</td>
                            <td>{{video.title}}</td>
                            <td>{{video.video_url}}</td>
                            <td>
                                <button type="submit" @click="deleteVideo(video.id)" class="btn btn-outline-danger">Xóa</button>
                            </td>
                        </tr>
                    </tbody>
                </table>
                <Pagination
                    :length="totalItems"
                    :pageSize="12"
                    :pageIndex="pageIndex"
                    @change="changePage"
                />
            </div>
        </div>
    </div>
</template>
<script>
import Pagination from "@/components/Pagination.vue";
import { mapState } from "vuex";
import apiPosts from "@/services/posts.service";
export default {
    name: 'Videos',
    components: {
        Pagination,
    },

    computed: {
        ...mapState("videos", ["totalItems","pageIndex", "videos"]),
    },

    methods: {
        changePage(pageIndex) {
            this.$store.dispatch("videos/getVideos", { pageIndex });
        },
        async deleteVideo(videoId) {
            await apiPosts.deleteVideo(videoId);
            window.location.reload();
        },
    },
    created(){
        this.$store.dispatch("videos/getVideos", {})
    }
}
</script>
