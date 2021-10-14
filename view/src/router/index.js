import Vue from 'vue'
import Router from 'vue-router'
import HomePage from '@/components/HomePage'
import AboutPage from '@/components/AboutPage'
import RequestPage from '@/components/RequestPage'
import SubmissionPage from '@/components/SubmissionPage'
import MyPage from '@/components/MyPage'
import NotFound from '@/components/NotFound'

Vue.use(Router)

export default new Router({
    mode: 'history',
    routes: [
        // ホームページ
        {
            path: '/',
            name: 'HomePage',
            component: HomePage
        },
        // サービス概要ページ
        {
            path: '/about',
            name: 'AboutPage',
            component: AboutPage
        },
        // リクエスト詳細ページ
        {
            path: '/request/:request_id',
            name: 'RequestPage',
            component: RequestPage
        },
        // サブミッション詳細ページ
        {
            path: '/submission/:submission_id',
            name: 'SubmissionPage',
            component: SubmissionPage
        },
        // マイページ
        {
            path: '/user/:user_id',
            name: 'MyPage',
            component: MyPage
        },
        // NotFoundページ
        {
            path: '/not-found',
            name: 'NotFound',
            component: NotFound
        },
        // 上記以外のページをNotFoundページへリダイレクト
        {
            path: '/*',
            component: NotFound
        }
    ]
});