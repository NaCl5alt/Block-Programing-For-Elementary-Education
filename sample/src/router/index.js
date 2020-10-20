import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../components/Home.vue'
import MyPage from '../components/MyPage.vue'
import Regist from '../components/Regist.vue'
import Login from '../components/Login.vue'
import Logout from '../components/Logout.vue'
import Exhibit from '../components/Exhibit.vue'
import ChangeID from '../components/ChangeID.vue'
import ChangeName from '../components/ChangeName.vue'
import ChangePasswd from '../components/ChangePasswd.vue'
import DeleteUser from '../components/DeleteUser.vue'
import axios from 'axios'

Vue.use(VueRouter)

const routes = [
  { path: '/', component: Home, name: 'Home' },
  { path: '/regist', component: Regist, name: 'ユーザ登録' },
  { path: '/mypage', component: MyPage, name: 'マイページ', meta: { requiresAuth: true } },
  { path: '/mypage/delete', component: DeleteUser, name: 'ユーザ削除', meta: { requiresAuth: true } },
  { path: '/change/id', component: ChangeID, name: 'ユーザID変更', meta: { requiresAuth: true } },
  { path: '/change/name', component: ChangeName, name: 'ユーザ名変更', meta: { requiresAuth: true } },
  { path: '/change/passwd', component: ChangePasswd, name: 'パスワード変更', meta: { requiresAuth: true } },
  { path: '/login', component: Login, name: 'ログイン' },
  { path: '/logout', component: Logout, name: 'ログアウト' },
  { path: '/exhibit', component: Exhibit, name: '出品' }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    axios.post(process.env.VUE_APP_API_URL_BASE + 'token', {}, { withCredentials: true }).then(res => {
      // if (res.data.token) this.$cookies.set('token', res.data.token)
      next()
    }).catch(() => {
      next({ path: '/login', query: { redirect: to.fullPath } })
    })
  } else next()
})

export default router
