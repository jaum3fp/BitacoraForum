import { createRouter, createWebHistory } from 'vue-router';
import Home from './pages/Index.vue';
import Explore from './pages/Explore.vue';
import Rules from './pages/Rules.vue';

const routes = [
  { path: '/', component: Home },
  { path: '/explore', component: Explore },
  { path: '/rules', component: Rules },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;