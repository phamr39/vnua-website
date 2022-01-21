<template>
    <div class="page-wrapper">
        <main class="main">
            <nav aria-label="breadcrumb" class="breadcrumb-nav mb-2">
                <div class="container">
                    <ol class="breadcrumb">
                        <li class="breadcrumb-item"><router-link to="/">Trang chủ</router-link></li>
                        <li class="breadcrumb-item active" aria-current="page">Sản phẩm</li>
                    </ol>
                </div><!-- End .container -->
            </nav><!-- End .breadcrumb-nav -->

            <div class="page-content">
                <div class="container">
                    <div class="row">
                        <div class="col">
                            <div class="toolbox">
                                <div class="toolbox-left">
                                    <div class="toolbox-info">
                                        Tổng <span>{{count}} / {{totalItems}}</span> Sản phẩm
                                    </div><!-- End .toolbox-info -->
                                </div><!-- End .toolbox-le    -->
                                <div class="toolbox-right">
                                    <div class="toolbox-sort">
                                        <label for="sortby">Sắp xếp:</label>
                                        <div class="select-custom">
                                            <select name="sortby" id="sortby" class="form-control">
                                                <option value="#" selected="selected">Mới nhất</option>
                                                <option value="#">Giá thấp đến cao</option>
                                                <option value="#">Giá cao đến thấp</option>
                                            </select>
                                        </div>
                                    </div> <!-- End .toolbox-sort -->
                                </div><!-- End .toolbox-right -->
                            </div><!-- End .toolbox -->

                            <div class="products mb-3">
                                <div class="row justify-content-center" >
                                    <div class="col-6 col-md-4 col-lg-4 col-xl-3" v-for="product in products" :key="product.id">
                                        <div class="product product-7 text-center">
                                            <figure class="product-media">
                                                <span class="product-label label-new">New</span>
                                                <router-link :to="'/sp/' + product.handle">
                                                    <img :src="product.images[0].image_url" alt="Product image" class="product-image">
                                                    <img :src="product.images[1].image_url" alt="Product image" class="product-image-hover">
                                                </router-link>

                                                <div class="product-action">
                                                    <router-link :to="'/sp/' + product.handle" class="btn-product"><span>đăng ký tư vấn</span></router-link>
                                                </div><!-- End .product-action -->
                                            </figure><!-- End .product-media -->

                                            <div class="product-body">
                                                <div class="product-cat" v-for="category in product.categories" :key="category.id">
                                                    <a href="#">{{category.name}}</a>
                                                </div><!-- End .product-cat -->
                                                <h3 class="product-title"><router-link :to="'/sp/' + product.handle">{{product.name}}</router-link></h3><!-- End .product-title -->
                                            </div><!-- End .product-body -->
                                        </div><!-- End .product -->
                                    </div><!-- End .col-sm-6 col-lg-4 col-xl-3 -->
                                </div><!-- End .row -->
                            </div><!-- End .products -->
                            <Pagination
                                :length="totalItems"
                                :pageSize="12"
                                :pageIndex="pageIndex"
                                @change="changePage"
                            />
                        </div><!-- End .col-lg-9 -->

                        <!-- <Catalog @setProducts="setProductsByCategory"/> -->
                    </div><!-- End .row -->
                </div><!-- End .container -->
            </div><!-- End .page-content -->
        </main><!-- End .main -->
    </div>
</template>
<script>
// import Catalog from "./Catalog";
import Pagination from "@/components/Pagination.vue";
import { mapGetters, mapState, mapActions } from "vuex";
export default {
    name: 'Products',
    components: {
        // Catalog,
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
        setProductsByCategory(categoryId){
            this.$store.dispatch("products/getProducts", { category : {id : categoryId} })
        },
        ...mapActions("products", ["getProducts"]),
    },
    created() {
        const payload = {
            search : this.$route.query.q ? this.$route.query.q : '',
            pageIndex : this.$route.query.page ? this.$route.query.page : 1,
            category_id : this.$route.query.category_id ? this.$route.query.category_id : null
        }
        this.$store.dispatch("products/getProducts", payload);
    },
    watch: {
        '$route.query'() {
            const payload = {
                search : this.$route.query.q ? this.$route.query.q : '',
                pageIndex : this.$route.query.page ? this.$route.query.page : 1,
                category_id: this.$route.query.category_id ? this.$route.query.category_id : ''
            }
            this.$store.dispatch("products/getProducts", payload);
        }
    },
}
</script>