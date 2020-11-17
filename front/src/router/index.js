import Vue from 'vue'
import VueRouter from 'vue-router'
import Top from '../components/Top'
import Login from '../components/Login'
import Regist from '../components/resist'
import Questions from '../components/Questions'
import Question from '../components/Question'
import QuestionAdd from '../components/QuestionAdd'
import QuestionEdit from '../components/QuestionEdit'
import User from '../components/User'
import Users from '../components/Users'
import Admin from '../components/Admin'
import AdminQuestions from '../components/AdminQuestions'

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
    component: Questions
  },
  {
    path: '/question/:id',
    name: 'Question',
    component: Question
  },
  {
    path: '/admin',
    name: 'Admin',
    component: Admin
  },
  {
    path: '/admin/question',
    name: 'AdminQuestions',
    component: AdminQuestions
  },
  {
    path: '/admin/question/add',
    name: 'QuestionAdd',
    component: QuestionAdd
  },
  {
    path: '/admin/question/:id/edit',
    name: 'QuestionEdit',
    component: QuestionEdit
  },
  {
    path: '/admin/user',
    name: 'Users',
    component: Users
  },
  {
    path: '/admin/user/:id',
    name: 'User',
    component: User
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
