import { createStore } from 'vuex'

import products from "./modules/products.module";
import posts from "./modules/posts.module";
import customers from "./modules/customers.module";
import users from "./modules/users.module";
import categories from "./modules/categories.module";
import videos from "./modules/videos.module";

export default createStore({
  modules: {
    products,
    posts,
    customers,
    users,
    categories,
    videos,
  }
})
