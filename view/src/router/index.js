import Vue from 'vue'
import Router from 'vue-router'
import HomePage from '@/components/HomePage'
import MyPage from '@/components/MyPage'

Vue.use(Router)

export default new Router({
    mode: 'history',
    routes: [{
            path: '/',
            name: 'HomePage',
            component: HomePage
        },
        {
            path: '/users/:userid',
            name: 'MyPage',
            component: MyPage
        }
    ]
})