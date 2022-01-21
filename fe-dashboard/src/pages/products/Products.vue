<template lang="">
    <div class="card border-success mb-3">
        <div class="card-header text-center">
            Danh sách sản phẩm
        </div>
        <div class="cart-body">
            <div class="p-5">
                <table class="table table table-striped">
                    <thead>
                        <tr>
                        <th scope="col">STT.</th>
                        <th scope="col">Hình ảnh</th>
                        <th scope="col">Sản phẩm</th>
                        <th scope="col"></th>
                        <th scope="col"></th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="product, id in products" :key="product.id">
                            <td>{{id + 1}}</td>
                            <td>
                                <img :src="product.images[0].image_url" width="50" alt="Product image" class="product-image">
                            </td>
                            <td>{{product.name}}</td>
                            <td>
                                <router-link :to="'/update-product/' + product.id"><button type="submit" @click="updateProduct(product.id)" class="btn btn-outline-info">Sửa</button></router-link>
                            </td>
                            <td>
                                <button type="submit" @click="deleteProduct(product.id)" class="btn btn-outline-danger">Xóa</button>
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
import { mapGetters, mapState, mapActions } from "vuex";
import apiProducts from "@/services/products.service";
export default {
    name: 'Products',
    components: {
        Pagination,
    },

    computed: {
        ...mapState("products", [
            "products",
            "totalItems",
            "pageIndex",
            "limit",
        ]),
        ...mapGetters("products", [
            "sortDropdownValue",
            "itemStartIndex",
            "itemEndIndex",
        ]),
        count() {
            return this.products.length;
        }
    },

    methods:{
        sortProducts(option) {
        const options = option.value.split("-");
        const sort = options[0],
        order = options[1];
        this.$store.dispatch("products/getProducts", { sort, order });
        },
        changePage(pageIndex) {
            this.$store.dispatch("products/getProducts", { pageIndex });
        },
        async deleteProduct(productId) {
            await apiProducts.deleteProductById(productId);
            window.location.reload();
        },
        ...mapActions("products", ["getProducts"]),
    },
    created() {
        this.$store.dispatch("products/getProducts", {});
    },
}
</script>
