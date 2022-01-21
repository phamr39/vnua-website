<template>
    <div class="card border-success mb-3">
        <div class="card-header text-center">
            Danh sách đại lý
        </div>
        <div class="cart-body">
            <div class="p-5">
                <Table :posts="stores" />
                <button @click="exportCSV" type="button" class="btn btn-success">Xuất Excel</button>
                <Pagination
                    :length="totalItems"
                    :pageSize="10"
                    :pageIndex="pageIndex"
                    @change="changePage"
                />
            </div>
        </div>
    </div>
</template>
<script>
import Table from '@/components/Table'
import Pagination from "@/components/Pagination.vue";
import { mapState } from "vuex";
export default {
    name : 'Stores',
    components : {
        Table,
        Pagination,
    },
    computed: {
        ...mapState("customers", [
            "stores",
            "totalItems",
            "pageIndex",
        ]),
    },
    methods:{
        changePage(pageIndex) {
            this.$store.dispatch("customers/getStores", { pageIndex });
        },
        exportCSV() {
            this.$store.dispatch("customers/getStoresCSV");
        },
    },
    created() {
        this.$store.dispatch("customers/getStores", {});
    },
}
</script>
