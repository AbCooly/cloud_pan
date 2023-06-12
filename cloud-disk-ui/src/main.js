import Vue from 'vue'
import App from './App.vue'

// 引入路由
import Router from 'vue-router'
// 引入路由器
import router from './router'

import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';

Vue.config.productionTip = false
Vue.use(Router)
Vue.use(ElementUI)
router.beforeEach((to, from, next) => {
    /* 路由发生变化修改页面title */
    if (to.meta.title) {
        document.title = to.meta.title
    }
    next()
})
new Vue({
    router,
    render: h => h(App),
}).$mount('#app')
