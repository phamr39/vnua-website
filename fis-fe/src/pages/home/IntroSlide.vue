<template>
  <div id="carouselExampleIndicators" class="carousel slide" data-ride="carousel">
    <ol class="carousel-indicators">
      <li data-target="#carouselExampleIndicators" data-slide-to="0" class="active"></li>
      <li data-target="#carouselExampleIndicators" data-slide-to="1"></li>
      <li data-target="#carouselExampleIndicators" data-slide-to="2"></li>
      <li data-target="#carouselExampleIndicators" data-slide-to="3"></li>
      <li data-target="#carouselExampleIndicators" data-slide-to="4"></li>
      <li data-target="#carouselExampleIndicators" data-slide-to="5"></li>
    </ol>
    <div class="carousel-inner">
      <div class="carousel-item active">
        <img class="d-block w-100" :src="image" alt="First slide">
      </div>
      <div class="carousel-item">
        <img class="d-block w-100" :src="image1" alt="Second slide">
      </div>
      <div class="carousel-item">
        <img class="d-block w-100" :src="image2" alt="Third slide">
      </div>
      <div class="carousel-item">
        <img class="d-block w-100" :src="image3" alt="Third slide">
      </div>
      <div class="carousel-item">
        <img class="d-block w-100" :src="image4" alt="Third slide">
      </div>
      <div class="carousel-item">
        <img class="d-block w-100" :src="image5" alt="Third slide">
      </div>
    </div>
    <a class="carousel-control-prev" href="#carouselExampleIndicators" role="button" data-slide="prev">
      <span class="carousel-control-prev-icon" aria-hidden="true"></span>
      <span class="sr-only">Previous</span>
    </a>
    <a class="carousel-control-next" href="#carouselExampleIndicators" role="button" data-slide="next">
      <span class="carousel-control-next-icon" aria-hidden="true"></span>
      <span class="sr-only">Next</span>
    </a>
  </div>
</template>
<script>
export default {
  name: 'IntroSlide',
  data() {
    return {
      image: require("@/assets/images/banners/A.png"),
      image1: require("@/assets/images/banners/F.png"),
      image2: require("@/assets/images/banners/H.png"),
      image3: require("@/assets/images/banners/N.png"),
      image4: require("@/assets/images/banners/P.png"),
      image5: require("@/assets/images/banners/V.png"),
    };
  }
}
</script>
