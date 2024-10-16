import {createRouter, createWebHistory} from 'vue-router'

const routes = [
    {
        name: 'index',
        path: '/',
        component: () => import("@/views/main/index.vue")
    },
    {
        name: 'order',
        path: '/order',
        component: () => import("@/views/main/order.vue")
    },
    {
        name: 'product',
        path: '/product',
        component: () => import("@/views/main/product.vue")
    },
    {
        name: 'analytic',
        path: '/analytic',
        component: () => import("@/views/main/analytic.vue")
    },
    {
        name: 'porduct_details',
        path: '/porduct_details/:id',
        component: () => import("@/views/main/product_detail.vue")
    },
    {
        name: 'customer',
        path: '/customer',
        component: () => import("@/views/main/customer.vue")
    },
    {
        name: '404',
        path: '/:catchAll(.*)',
        component: () => import('@/views/404.vue') //  当路由不存在则到该路由
    },
    {
        name: 'login',
        path: '/login',
        component: () => import('@/views/user/login.vue')
    },
    {
        name: 'register',
        path: '/register',
        component: () => import('@/views/user/register.vue')
    },
    {
        name: 'PasswordReset',
        path: '/password_reset',
        component: () => import('@/views/user/PasswordReset.vue')
    },
    {
        name: 'nodeInfo',
        path: '/nodeInfo',
        component: () => import('@/views/main/nodeInfo.vue')
    },
    {
        name: 'home',
        path: '/admin/home',
        component: () => import('@/views/admin/home.vue')
    },
    {
        name: 'pay',
        path: '/admin/pay',
        component: () => import('@/views/admin/pay.vue')
    },
    {
        name: 'package',
        path: '/admin/package',
        component: () => import('@/views/admin/package.vue')
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

// 全局导守卫
router.beforeEach((to, from, next) => {
    // to 满足条件的url放过
    if (to.path == "/login" || to.path == "/register" || to.path == "/password_reset") {
        next()  // 直接放过
    } else {
        // 校验用户的，不受管理员影响
        const auth_token = localStorage.getItem("auth_token")
        // front的需要权限的url
        if (auth_token == "" || auth_token == null || auth_token == '') {
            next("/login")
        } else {
            next()
        }
    }
})


export default router
