<template>
    <Carousel
    :items-to-show="1"
    :itemsToScroll="1"
    :breakpoints="{
        576: {
            itemsToShow: 2,
            itemsToScroll: 2,
        },
        768: {
            itemsToShow: 3,
            itemsToScroll: 3,
        },
        992: {
            itemsToShow: 5,
            itemsToScroll: 5,
        },
    }"
    >
    <Slide v-for="product in products" :key="product.id">
        <div class="products">
            <div class="justify-content-center">
                <div class="col">
                    <div class="product product-7 text-center">
                        <figure class="product-media">
                            <span class="product-label label-circle label-sale">Hot</span>
                            <router-link :to="'/sp/' + product.name">
                                <img :src="product.images[0].image_url" alt="Product image" class="product-image">
                                <img :src="product.images[1].image_url" alt="Product image" class="product-image-hover">
                            </router-link>
                            <div class="product-action">
                                <a href="#signin-modal" data-toggle="modal" class="btn-product"><span>đăng ký tư vấn</span></a>
                            </div><!-- End .product-action -->
                        </figure><!-- End .product-media -->
                        <div class="product-body">
                                <div class="product-cat" v-for="category in product.categories" :key="category.id">
                                    <a href="#">{{category.name}}</a>
                                </div><!-- End .product-cat -->
                                <h3 class="product-title"><router-link :to="'/sp/' + product.name">{{product.name}}</router-link></h3><!-- End .product-title -->
                            </div><!-- End .product-body -->
                    </div><!-- End .product -->
                </div>
            </div>
        </div>
    </Slide>

    <template #addons>
        <Navigation />
    </template>
    </Carousel>
</template>

<script>
import "vue3-carousel/dist/carousel.css";
import { Carousel, Slide, Navigation } from "vue3-carousel";

export default {
    name: "ProductsCarousel",
    props: ["products"],
    components: {
        Carousel,
        Slide,
        Navigation,
    },
};
</script>

<style>
.carousel__prev,
.carousel__next {
    background-color: transparent;
    font-size: 40px;
    top: calc((100% - 70px) / 2);
}
.carousel__icon {
    color: #ccc;
}
</style>