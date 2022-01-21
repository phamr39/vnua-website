<template>
    <div class="zalo">
        <div v-show="zaloShow" class="content">
            <h6>Hỗ trợ trực tuyến</h6>
            <a target="_blank" href="https://zalo.me/0982274784" title="Mrs Mien" rel="noopener"><p>Mrs. Mien - 0982 27 47 84</p></a>
            <a target="_blank" href="https://zalo.me/0966101519" title="Mr. Bach" rel="noopener"><p>Mr. Bach - 0966 10 15 19</p></a>
        </div>
        <img 
            @click="showData"
            src="@/assets/images/brands/zalo.png" 
            alt="zalo" />
    </div>
</template>
<script>
export default {
    name: 'Zalo',
    data() {
        return {
            zaloShow: false,
        }
    },
    methods: {
        showData() {
            this.zaloShow = !this.zaloShow;
        }
    }
}
</script>

<style scoped>
.zalo {
    display: flex;
    z-index: 10;
    position: fixed;
    top: 80%;
    left: 30px;
}

.zalo img {
    width: 100px;
    cursor: pointer;
    z-index: 10;
    position: fixed;
    top: 90%;
    left: 30px;
}

.content {
    background-color: #fff;
    padding: 5px;
    border: 1px solid #bdc3c7;
    border-radius: 5px;
}
</style>