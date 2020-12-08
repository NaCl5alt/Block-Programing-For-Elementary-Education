import Vue from 'vue'
import VueRouter from 'vue-router'
import axios from 'axios'

import Top from '../components/Top'
import Login from '../components/Login'
import Regist from '../components/resist'
import Questions from '../components/Questions'
import Question from '../components/Question'
import QuestionAdd from '../components/QuestionAdd'
import QuestionEdit from '../components/QuestionEdit'
import User from '../components/User'
import Users from '../components/Users'
import UserPage from '../components/UserPage'
import Admin from '../components/Admin'
import AdminQuestions from '../components/AdminQuestions'
import Logout from '../components/Logout'
import ChangeID from '../components/ChangeID.vue'
import ChangeName from '../components/ChangeName.vue'
import ChangePasswd from '../components/ChangePasswd.vue'
import DeleteUser from '../components/DeleteUser.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Top',
    component: Top
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/regist',
    name: 'Regist',
    component: Regist
  },
  {
    path: '/question',
    name: 'Questions',
    component: Questions,
    meta: { requiresAuth: true }
  },
  {
    path: '/question/:id',
    name: 'Question',
    component: Question,
    meta: { requiresAuth: true }
  },
  {
    path: '/admin',
    name: 'Admin',
    component: Admin,
    meta: { requiresAuth: true }
  },
  {
    path: '/admin/question',
    name: 'AdminQuestions',
    component: AdminQuestions,
    meta: { requiresAuth: true }
  },
  {
    path: '/admin/question/add',
    name: 'QuestionAdd',
    component: QuestionAdd,
    meta: { requiresAuth: true }
  },
  {
    path: '/admin/question/:id/edit',
    name: 'QuestionEdit',
    component: QuestionEdit,
    meta: { requiresAuth: true }
  },
  {
    path: '/user',
    name: 'UserPage',
    component: UserPage,
    meta: { requiresAuth: true }
  },
  {
    path: '/admin/user',
    name: 'Users',
    component: Users,
    meta: { requiresAuth: true }
  },
  {
    path: '/admin/user/:id',
    name: 'User',
    component: User,
    meta: { requiresAuth: true }
  },
  {
    path: '/logout',
    name: 'Logout',
    component: Logout
  },
  { path: '/id', component: ChangeID, name: 'ユーザID変更', meta: { requiresAuth: true } },
  { path: '/name', component: ChangeName, name: 'ユーザ名変更', meta: { requiresAuth: true } },
  { path: '/passwd', component: ChangePasswd, name: 'パスワード変更', meta: { requiresAuth: true } },
  { path: '/delete', component: DeleteUser, name: 'ユーザ削除', meta: { requiresAuth: true } },
 
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    let token = ''
    if(window.$cookies.isKey('token')) token = window.$cookies.get('token')
    axios.post('/api/token', {}, { 
      withCredentials: true,
      headers: { Authorization: `Bearer ${token}` }
    }).then(() => {
    // if (res.data.token) this.$cookies.set('token', res.data.token)
    next()
    }).catch(() => {
      next({ path: '/login', query: { redirect: to.fullPath } })
    })
  } else next()
})

export default router
