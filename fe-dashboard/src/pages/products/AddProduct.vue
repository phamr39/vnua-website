<template>
    <div class="card border-success mb-3">
        <div class="card-header text-center">
            Thêm sản phẩm mới
        </div>
        <div class="cart-body">
            <form class="p-5">
                <div class="mb-3">
                    <label>Ảnh sản phẩm (ít nhất 2 ảnh)</label>
                        <ImageUpload 
                            @setFile="getFile"
                            @change="submitFile"
                        />
                </div>

                <label>Tên sản phẩm</label>
                <input type="text" class="form-control" v-model="name" required>

                <label>Danh mục sản phẩm</label>
                <select class="custom-select" v-model="category" >
                    <option v-for="category in categories" :key="category.id" :value="category.id">{{category.name}}</option>
                </select>

                <label>Mô tả ngắn</label>
                <QuillEditor theme="snow" class="mb-2" v-model:content="description" contentType="html" :options="options"/>

                <label>Thông tin sản phẩm</label>
                <QuillEditor theme="snow" class="mb-2" v-model:content="product_info" contentType="html" :options="options"/>

                <label>Hướng dẫn sử dụng</label>
                <QuillEditor theme="snow" class="mb-2" v-model:content="user_manual" contentType="html" :options="options" />

                <button @click="createNewProduct()" class="btn btn-success">Xác nhận</button>
                <button @click="cancel()" class="btn btn-warning ml-5">Hủy bỏ</button>
            </form>
        </div>
    </div>
</template>
<script>
import { QuillEditor } from '@vueup/vue-quill'
import '@vueup/vue-quill/dist/vue-quill.snow.css';
import apiUpload from "@/services/upload.service";
import apiProducts from "@/services/products.service";
import { mapState } from 'vuex';
import ImageUpload from "@/components/ImageUpload.vue"
export default {
    name : 'AddProduct',
    components : {
        QuillEditor,
        ImageUpload,
    },
    data(){
        return {
            file: '',
            images: [],
            category: [],
            name: '',
            description: '',
            product_info: '',
            user_manual: '',
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
    computed:{
        ...mapState("products",["categories"])
    },
    methods: {
        submitFile(){
            let formData = new FormData();
            if(this.file != ''){
                formData.append('file', this.file);
                apiUpload.uploadImage(formData)
                .then(response => {
                    this.images.push(response);
                })
                .catch(function(){
                    console.log('FAILURE!!');
                });
            }
        },
        getFile(file){
            this.file = file;
        },

        async createNewProduct() {
            let cates = [];
            for (let i = 0; i < this.category.length; i++) {
                cates.push({id: parseInt(this.category[i])})
            }
            let imgs = [];
            for (let i = 0; i < this.images.length; i++) {
                imgs.push({image_url: this.images[i]})
            }
            const product = {
                name : this.name,
                description: this.description,
                product_info: this.product_info,
                user_manual: this.user_manual,
                categories: cates,
                images: imgs,
            }
            console.log(product)
            apiProducts.createProduct(product)
            await this.$router.push("/products");
        },
        async cancel() {
            await this.$router.push("/products");
        },
    },
    created(){
        this.$store.dispatch("products/getCategories", "");
    }
}
</script>
