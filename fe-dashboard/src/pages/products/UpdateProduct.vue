<template>
    <div class="card border-success mb-3">
        <div class="card-header text-center">
            Sửa sản phẩm
        </div>
        <div class="cart-body">
            <form class="p-5">
                <label>Tên sản phẩm</label>
                <input type="text" class="form-control" v-model="name"  required>

                <label>Mô tả ngắn</label>
                <QuillEditor theme="snow" class="mb-2"  v-model:content="description" contentType="html" :options="options"/>

                <label>Thông tin sản phẩm</label>
                <QuillEditor theme="snow" class="mb-2" v-model:content="product_info" contentType="html" :options="options"/>

                <label>Hướng dẫn sử dụng</label>
                <QuillEditor theme="snow" class="mb-2" v-model:content="user_manual" contentType="html" :options="options" />

                <button @click="updateProduct()" class="btn btn-success">Xác nhận</button>
                <button @click="cancel()" class="btn btn-warning ml-5">Hủy bỏ</button>
            </form>
        </div>
    </div>
</template>
<script>
import { QuillEditor } from '@vueup/vue-quill'
import '@vueup/vue-quill/dist/vue-quill.snow.css';
import apiProducts from "@/services/products.service";
import { mapState } from 'vuex';
export default {
    name : 'AddProduct',
    components : {
        QuillEditor,
    },
    data(){
        return {
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
    computed: {
        ...mapState("products", ["product"]),
    },
    methods: {
        async updateProduct() {
            const product = {
                name : this.name,
                description: this.description,
                product_info: this.product_info,
                user_manual: this.user_manual,
            }
            console.log(product) 
            apiProducts.updateProductById(this.$route.params.id, product)
            await this.$router.push("/products");
        },
        async cancel() {
            await this.$router.push("/products");
        }
    },
    created() {
        this.$store.dispatch("products/getProductById", this.$route.params.id);
    },
}
</script>
