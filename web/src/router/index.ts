import { createRouter, createWebHistory } from 'vue-router'
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'web',
      component: () => import('../views/web/index.vue'),
    },
    {
      path: '/auth',
      name: 'loginBase',
      component: () => import('../views/login/index.vue'),
      children: [
        {
          path: '/login',
          name: 'login',
          component: () => import('../views/login/login.vue'),
        },
        {
          path: '/register',
          name: 'register',
          component: () => import('../views/login/register.vue'),
        },
      ],
    },
    {
      path: '/admin',
      name: 'admin',
      component: () => import('../views/admin/index.vue'),
      meta: {
        title: '首页',
      },
      children: [
        {
          path: 'data',
          name: 'data',
          meta: {
            title: '数据统计',
          },
          component: () => import('../views/admin/data.vue'),
        },
        {
          path: 'manage',
          name: 'manage',
          meta: {
            title: '管理',
          },
          children: [
            {
              path: 'user',
              name: 'user-list',
              component: () => import('../views/admin/user/user-list.vue'),
              meta: {
                title: '用户列表',
              },
            },
            {
              path: 'role',
              name: 'role-list',
              component: () => import('../views/admin/user/role-list.vue'),
              meta: {
                title: '角色列表',
              },
            },
          ],
        },
        {
          path: "/:all(.*)",
          component: () => import('../views/admin/404.vue'),
        }
      ],
    },
  ],
})


router.beforeEach((to, from, next) => {
  NProgress.start();//开启进度条
  next()
})
//当路由进入后：关闭进度条
router.afterEach(() => {
  // 在即将进入新的页面组件前，关闭掉进度条
  NProgress.done()//完成进度条
})

export default router
