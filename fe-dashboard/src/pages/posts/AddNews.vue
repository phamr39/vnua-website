<template>
    <div class="card border-success mb-3">
        <div class="card-header text-center">
            Thêm tin tức
        </div>
        <div class="cart-body">
            <form class="p-5" @submit.prevent="createNews">

                <div class="mb-3">
                    <label>Ảnh tiêu đề </label>
                    <ImageUpload 
                        @setFile="getFile"
                        @change="submitFile"
                    />
                </div>

                <label>Tiêu đề</label>
                <input type="text" class="form-control" v-model="title" required>

                <label>Mô tả ngắn</label>
                <input type="text" class="form-control" v-model="description"  required>

                <label>Nội dung</label>
                <QuillEditor theme="snow" class="mb-2" v-model:content="content" contentType="html" :options="options"/>

                <label>Tác giả</label>
                <input type="text" class="form-control" v-model="author"  required>

                <button class="btn btn-success mt-2">Xác nhận</button>
                <button @click="cancel()" class="btn btn-warning ml-5">Hủy bỏ</button>
            </form>
        </div>
    </div>
</template>
<script>
import { QuillEditor } from '@vueup/vue-quill'
import '@vueup/vue-quill/dist/vue-quill.snow.css';
import apiUpload from "@/services/upload.service";
import ImageUpload from "@/components/ImageUpload.vue"
export default {
    name : 'AddNews',
    components : {
        QuillEditor,
        ImageUpload,
    },
    data(){
        return {
            file:'',
            title: '',
            topics: [],
            thumbnail: '',
            description: '',
            content: '',
            author: '',
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
        submitFile(){
            let formData = new FormData();
            if(this.file != ''){
                formData.append('file', this.file);
                apiUpload.uploadImage(formData)
                .then(response => {
                    this.thumbnail = response;
                })
                .catch(function(){
                    console.log('FAILURE!!');
                });
            }
        },

        getFile(file){
            this.file = file;
        },

        async createNews() {
            const news = {
                title : this.title,
                topics : this.topics,
                thumbnail : this.thumbnail,
                description: this.description,
                post_type: "news",
                content : this.content,
                author : this.author,
            }
            await this.$store.dispatch("posts/createNews", news);
            await this.$router.push("/news");
        },
        async cancel() {
            await this.$router.push("/news");
        },
    }
}
</script>

