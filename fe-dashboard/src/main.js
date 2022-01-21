import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

createApp(App).use(store).use(router).mount('#app')

// router.beforeEach((to, from, next) => {
//     if (to.matched.some(record => record.meta.requiresAuth)) {
//         var token = store.getters.loggedIn;
//         if (!token ) {
//             next({
//                 path: '/login',
//             })
//             } else {
//             next()
//             }
//         }else if (to.matched.some(record => record.meta.requiresVisitor)) {
//             if (token) {
//             next({
//                 path: '/profiles',
//             })
//             } else {
//             next()
//             }
//         }
//     })
