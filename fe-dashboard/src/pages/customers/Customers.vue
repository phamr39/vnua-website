<template>
    <div class="card border-success mb-3">
        <div class="card-header text-center">
            Danh sách khách hàng
        </div>
        <div class="cart-body">
            <div class="p-5">
                <Table :posts="customers" />
                <button @click="exportCSV" type="button" class="btn btn-success">Xuất Excel</button>
                <Pagination
                    :length="totalItems"
                    :pageSize="20"
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
    name : 'Customers',
    components : {
        Table,
        Pagination,
    },
    computed: {
        ...mapState("customers", [
            "customers",
            "totalItems",
            "pageIndex",
        ]),
    },
    methods:{
        changePage(pageIndex) {
            this.$store.dispatch("customers/getCustomers", { pageIndex });
        },
        exportCSV() {
            this.$store.dispatch("customers/getCustomersCSV");
        },
    },
    created() {
        this.$store.dispatch("customers/getCustomers", {});
    },
}
</script>
