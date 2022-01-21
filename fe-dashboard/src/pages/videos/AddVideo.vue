<template>
    <div class="card border-success mb-3">
        <div class="card-header text-center">
            Thêm video
        </div>
        <div class="cart-body">
            <form class="p-5" @submit.prevent="createVideo">

                <label>Tiêu đề</label>
                <input type="text" class="form-control" v-model="title" required>

                <label>Link video</label>
                <input type="text" class="form-control" v-model="video_url"  required>

                <label>Nội dung</label>
                <QuillEditor theme="snow" class="mb-2" v-model:content="description" contentType="html" :options="options"/>

                <button class="btn btn-success mt-2">Xác nhận</button>
                <button @click="cancel()" class="btn btn-warning ml-5">Hủy bỏ</button>
            </form>
        </div>
    </div>
</template>
<script>
import { QuillEditor } from '@vueup/vue-quill'
import '@vueup/vue-quill/dist/vue-quill.snow.css';
export default {
    name : 'AddVideo',
    components : {
        QuillEditor,
    },
    data(){
        return {
            title: '',
            description: '',
            video_url: '',
            options: {
                modules: {
                    'toolbar': [
                        [{ 'size': [] }],
                        [{ 'header': [1, 2, 3, 4, 5, 6, false] }],
                        ['bold', 'italic', 'underline', 'strike'],
                        ['blockquote', 'code-block'],
                        [{ 'list': 'ordered' }, { 'list': 'bullet' }],
                        [{ 'align': [] }],
                        ['link', 'image', 'video'],
                    ],
                },
            },
        }
    },
    methods: {
        async createVideo() {
            const video = {
                title : this.title,
                description: this.description,
                video_url : this.video_url,
            }
            await this.$store.dispatch("posts/createVideo", video);
            await this.$router.push("/videos");
        },
        async cancel() {
            await this.$router.push("/videos");
        },
    }
}
</script>

