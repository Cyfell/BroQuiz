import { createRouter, createWebHistory } from 'vue-router';
import AudienceView from './views/AudienceView.vue';
import PresenterView from './views/PresenterView.vue';

export const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/', name: 'audience', component: AudienceView },
    { path: '/presenter', name: 'presenter', component: PresenterView },
  ],
});
