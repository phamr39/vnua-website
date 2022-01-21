<template>
    <nav aria-label="Page navigation">
        <ul class="pagination justify-content-center">
            <li class="page-item">
                <a 
                    href="#" 
                    @click="changePage(pageIndex - 1, $event)"
                    class="page-link page-link-prev" 
                    :class="{ disabled: pageIndex === 1 }"
                >
                    <span aria-hidden="true"><i class="icon-long-arrow-left"></i></span>
                </a>
            </li>
            <li class="page-item" aria-current="page">
                <a 
                    href="#"
                    class="page-link" 
                    @click="changePage(page, $event)"
                    :class="{
                        'active': pageIndex ? page === pageIndex : page === 1,
                    }"
                    v-for="page in displayPages"
                    :key="page"
                >
                    {{ page }}
                </a>
            </li>
            <li class="page-item">
                <a 
                    href="#"
                    @click="changePage(pageIndex + 1, $event)"
                    class="page-link page-link-next"
                    :class="{ disabled: pageIndex === numberOfPages }"
                >
                    <span aria-hidden="true"><i class="icon-long-arrow-right"></i></span>
                </a>
            </li>
        </ul>
    </nav>
</template>
<script>
export default {
    name: 'Pagination',
    props: ["length", "pageSize", "pageIndex", "numberOfDisplayPages"],
    emits: ["change"],
    computed: {
        numberOfPages() {
            return Math.ceil(this.length / this.pageSize);
        },
        displayPages() {
            let numberOfDisplay = this.numberOfDisplayPages || 5;
            const halfNumberOfDisplayPages = Math.floor(numberOfDisplay / 2);
            let display = [];
            let startDisplay = this.pageIndex - halfNumberOfDisplayPages;
            if (startDisplay < 0) startDisplay = 0;
            numberOfDisplay--;
            let endDisplay = startDisplay + numberOfDisplay;
            if (endDisplay > this.numberOfPages - 1) {
                endDisplay = this.numberOfPages - 1;
                startDisplay = endDisplay - numberOfDisplay;
                if (startDisplay < 0) startDisplay = 0;
            }
            for (let i = startDisplay + 1; i <= endDisplay + 1; i++) {
                display.push(i);
            }
            return display;
        },
    },
    methods: {
        changePage(page, event) {
            event.preventDefault();
            this.$emit("change", page);
        },
    },
}
</script>
