<template>
  <div class="wrapper d-flex align-items-stretch">
    <SideBar />
    <div id="content" class="p-4 p-md-5 pt-5">
      <router-view v-slot="{ Component }">
        <transition name="fade">
          <component :is="Component" />
        </transition>
    </router-view>
    </div>
</div>
</template>

<script>
import SideBar from "@/components/SideBar.vue";
import api from "@/services/users.service";
import {onMounted} from 'vue';
import {useRouter} from "vue-router";
import {useStore} from "vuex";
import { mapState} from "vuex";
export default {
  name: "App",
  components: {
    SideBar,
  },
  computed: mapState("users", ["isLogined", "loginMessage"]),
    created() {
      if (this.isLogined) {
          this.$router.replace("/");
      }
    },
  setup() {
    const router = useRouter();
    const store = useStore();
    onMounted(async () => {
        try {
            const {data} = await api.getProfiles();
            store.commit('users/setUser', data);
        } catch (e) {
            await router.push('/login');
        }
    });
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
