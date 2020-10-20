<template>
  <header>
    <b-navbar toggleable="lg">
      <b-navbar-brand href="/"><h1>サイト名</h1></b-navbar-brand>

      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav>

        <b-navbar-nav class="ml-auto">
          <b-nav-form @submit.prevent="search">
            <b-form-input size="sm" class="mr-sm-2" placeholder="Search" v-model="keyword"></b-form-input>
            <b-button size="sm" class="my-2 my-sm-0" type="submit">Search</b-button>
          </b-nav-form>

          <template class="row ml-auto" v-if="verify === false">
            <b-nav-item href="/regist">登録</b-nav-item>
            <b-nav-item href="/login">ログイン</b-nav-item>
          </template>

          <template class="row ml-auto" v-else>
            <b-nav-item href="/mypage">マイページ</b-nav-item>
            <b-nav-item href="/exhibit">出品</b-nav-item>
            <b-nav-item href="/logout">ログアウト</b-nav-item>
          </template>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>
  </header>
</template>

<script>
import axios from 'axios'
export default {
  data () {
    return {
      keyword: '',
      verify: false,
      count: 0
    }
  },
  methods: {
    search () {
      this.$router.push({
        name: 'Search',
        params: {
          keyword: encodeURIComponent(this.keyword)
        }
      }).catch(() => {})
    },
    verifyfunc () {
      this.$nextTick(() => {
        axios.post(process.env.VUE_APP_API_URL_BASE + 'token', {}, { withCredentials: true }).then(res => {
          console.log('status: ' + res.status)
          if (res.status === 200) {
            this.verify = true
          } else this.verify = false
        }).catch(err => {
          console.log(err)
        })
      })
    }
  },
  beforeMount () {
    this.verifyfunc()
  }
}
</script>

<style scoped>
.navbar {
  background-color: antiquewhite !important
}
.navbar-brand {
  color: red !important
}
</style>
