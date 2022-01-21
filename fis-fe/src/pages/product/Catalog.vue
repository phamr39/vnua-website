<template>
    <aside class="col-lg-3 order-lg-first">
        <div class="sidebar sidebar-shop">
			<div class="widget widget-clean">
				<label>Bộ lọc:</label>
				<!-- <a href="#" class="sidebar-filter-clear">Xóa bộ lọc</a> -->
			</div>
			<div class="widget" v-for="category, id in rootCate" :key="category.id">
				<h3 class="widget-title">
					<a data-toggle="collapse" :href="`#`+id" @click="getProductsByCategory(category.id, $event)" role="button" aria-expanded="true" :aria-controls="id">{{category.name}}</a>
				</h3><!-- End .widget-title -->
				<!-- <div class="collapse show" :id="id">
					<div class="widget-body" v-for="cate in childCate" :key="cate.id">
						<a href="#">{{cate.name}}</a>
					</div>
				</div> -->
			</div><!-- End .widget -->
        </div><!-- End .sidebar sidebar-shop -->
    </aside><!-- End .col-lg-3 -->
</template>
<script>
import { mapState } from 'vuex';
export default {
    name : 'Catalog',
	computed: {
        ...mapState("categories", ["rootCate","childCate"]),
    },
    methods: {
        searchProducts() {
            this.$store.dispatch("products/getProducts", {
                pageIndex: 1,
                search: this.searchKeyword,
            });
        },
        setChildCate(id){
            this.$store.dispatch("categories/getCategories", id);
        },
		getProductsByCategory(id, event){
            event.preventDefault();
            this.$emit("setProducts", id);
        },
    },
    created(){
        this.$store.dispatch("categories/getCategories", "");
    }

}
</script>
