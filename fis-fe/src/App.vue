<template>
<div>
  <div class="page-wrapper">
    <MainHeader />
    <router-view v-slot="{ Component }">
      <transition name="fade">
        <component :is="Component" />
      </transition>
    </router-view>
    <Zalo/>
    <MainFooter />
  </div>
  <MobileSidebar/>
</div>
</template>

<script>
import MainHeader from "@/components/MainHeader.vue";
import MobileSidebar from "@/components/MobileSidebar.vue";
import MainFooter from "@/components/MainFooter.vue";
import Zalo from "@/components/Zalo.vue";

export default {
  name: "App",
  components: {
    MainHeader,
    MainFooter,
    MobileSidebar,
    Zalo
  },
};
</script>

<style>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
