import { createStore } from 'vuex'

import products from "./modules/products.module";
import posts from "./modules/posts.module";
import categories from "./modules/categories.module";
import videos from "./modules/videos.module";
import customers from "./modules/customers.module";

export default createStore({
  modules: {
    products,
    posts,
    categories,
    videos,
    customers,
  }
})
