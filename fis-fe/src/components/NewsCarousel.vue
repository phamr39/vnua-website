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
            itemsToShow: 4,
            itemsToScroll: 4,
        },
    }"
    >
    <Slide v-for="news in listNews" :key="news.id">
            <div class="container">
                <article class="entry">
                    <figure class="entry-media">
                        <router-link :to="'/chi-tiet-tin-tuc/' + news.id">
                            <img :src="news.thumbnail" alt="image desc">
                        </router-link>
                    </figure><!-- End .entry-media -->

                    <div class="entry-body">
                        <div class="entry-meta">
                            <a href="#">{{news.created_at}}</a> | 
                            <a href="#">{{news.author}}</a>
                        </div><!-- End .entry-meta -->

                        <h3 class="entry-title">
                            <router-link :to="'/chi-tiet-tin-tuc/' + news.id">Bài viết: {{news.title}}</router-link>
                        </h3><!-- End .entry-title -->

                        <div class="entry-content">
                            <p>{{news.description}}</p>
                            <router-link :to="'/chi-tiet-tin-tuc/' + news.id" class="read-more">Đọc tiếp</router-link>
                        </div><!-- End .entry-content -->
                    </div><!-- End .entry-body -->
                </article><!-- End .entry -->
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
    name: "NewsCarousel",
    props: ["listNews"],
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