import { createRouter, createWebHistory } from "vue-router";

const routes = [
  {
    path: "/",
    name: "Chats",
    component: () => import("../views/Home"),
  },
  {
    path: "/about",
    name: "About",
    component: () => import("../views/About.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
