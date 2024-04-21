import { createApp } from "vue";
import "./style.css";
import "primevue/resources/themes/aura-light-green/theme.css";
import "primeicons/primeicons.css";

import App from "./App.vue";
import PrimeVue from "primevue/config";
import { createWebHistory, createRouter } from "vue-router";

import LandingPage from "./pages/public/Landing.vue";
import LoginPage from "./pages/public/Login.vue";
import CoverPage from "./pages/protected/Cover.vue";
import DashboardPage from "./pages/protected/Cover.vue";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", component: LandingPage },
    { path: "/login", component: LoginPage },
    { path: "/cover", component: CoverPage },
    { path: "/dashboard", component: DashboardPage },
  ],
});

const app = createApp(App);
app.use(PrimeVue);
app.use(router);

app.mount("#app");
