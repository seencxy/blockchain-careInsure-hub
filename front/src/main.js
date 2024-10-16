import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import axios from 'axios'
import router from './router/router.js' // 导入路由配置
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'


const app = createApp(App)

// 使用 app.config.globalProperties 来添加全局方法或属性
const instance = axios.create({
    baseURL: 'http://127.0.0.1:9999', // 域名配置
    headers: {
        'Authorization': "Bearer "+localStorage.getItem("auth_token"), // 添加默认请求头部
        // 这里可以根据需要添加更多的默认头部
    }
});

// 添加请求拦截器
instance.interceptors.request.use(function (config) {
    // 在发送请求之前，从 localStorage 获取最新的 auth_token
    const authToken = localStorage.getItem('auth_token');
    if (authToken) {
        config.headers['Authorization'] = `Bearer ${authToken}`;
    }
    return config;
}, function (error) {
    // 对请求错误做些什么
    return Promise.reject(error);
});

// 将Axios实例添加到全局属性
app.config.globalProperties.$axios = instance;

// 使用路由配置
app.use(router)
app.use(ElementPlus)
app.mount('#app')
