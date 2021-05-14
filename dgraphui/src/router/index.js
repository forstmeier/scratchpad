import Vue from 'vue';
import Router from 'vue-router';
import Home from '@/views/Home.vue';
import Database from '@/views/Database.vue';
import { authGuard } from '../auth';


Vue.use(Router);

const router = new Router({
	mode: 'history',
	// base: process.env.BASE_URL,
	routes: [
		{
			path: '/',
			name: 'Home',
			component: Home,
		},
		{
			path: '/database',
			name: 'Database',
			component: Database,
			beforeEnter: authGuard,
		},
	]
});

export default router;
