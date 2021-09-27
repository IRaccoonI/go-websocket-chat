import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import VueNativeSock from "vue-native-websocket-vue3";

createApp(App)
  .use(store)
  .use(router)
  .use(VueNativeSock, "ws://localhost:3000/ws", {})
  .mount("#app");
