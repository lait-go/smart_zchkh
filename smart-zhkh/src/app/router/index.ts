import { Dashboard } from '@/pages/dashboard';
import { Home } from '@/pages/home';
import { Login } from '@/pages/login';
import { Register } from '@/pages/register';
import AppLayout from '@/layouts/AppLayout.vue';
import { createRouter, createWebHistory } from 'vue-router';
import { Charges, ChargesAdd } from '@/pages/charges';
import { Profile, ProfileEdit } from '@/pages/profile';
import { AddAccount } from '@/pages/account';
import { CreateTasks } from '@/pages/tasks/intex';
import AdminLayout from '@/layouts/AdminLayout.vue';
import { AdminTasks, AdminTasksDetail } from '@/pages/admin_tasks';
import { UserHistory } from '@/pages/user';
import { AdminCharges } from '@/pages/admin_charges';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: AppLayout,
      children: [
        {
          path: '/',
          name: 'home',
          component: Home,
        },
        {
          path: '/register',
          name: 'register',
          component: Register,
        },
        {
          path: '/login',
          name: 'login',
          component: Login,
        },
        {
          path: '/dashboard',
          name: 'dashboard',
          component: Dashboard,
        },
        {
          path: '/charges',
          name: 'charges',
          component: Charges,
        },
        {
          path: '/charges/add',
          name: 'charges add',
          component: ChargesAdd,
        },
        {
          path: '/profile',
          name: 'profile',
          component: Profile,
        },
        {
          path: '/accounts/:id/edit',
          component: ProfileEdit,
        },
        {
          path: '/accounts/add',
          component: AddAccount,
        },
        {
          path: '/tasks/create',
          component: CreateTasks,
        },
        {
          path: '/profile/history',
          component: UserHistory,
        },
      ],
    },
    {
      path: '/admin',
      component: AdminLayout,
      children: [
        {
          path: 'tasks',
          component: AdminTasks,
        },
        {
          path: 'tasks/:id',
          component: AdminTasksDetail,
        },
        {
          path: 'charges',
          component: AdminCharges,
        },
      ],
    },
  ],
});

export default router;
