<template>
    <div>
        <form @submit.prevent="login">
            <h1>ログイン</h1>
            ユーザID: <input v-model="userid" type="text">
            パスワード: <input v-model="pass" type="password">
            <button type="submit" class="btn btn-primary">ログイン</button>
        </form>
    </div>
</template>

<script>
import axios from 'axios'
export default {
  data () {
    return {
      userid: '',
      pass: ''
    }
  },
  methods: {
    async login () {
      var params = new URLSearchParams()
      params.append('userid', this.userid)
      params.append('pass', this.pass)
      await axios.post(process.env.VUE_APP_API_URL_BASE + 'login', params, { withCredentials: true }).then(res => {
        if (res.status === 200) {
          const token = res.data.token
          this.$cookies.set('token', token)
          var path = '/'
          if (this.$route.query.redirect) path = path + this.$route.query.redirect
          console.log(path)
          this.$router.push({ path: path })
          this.$router.go()
        }
      }).catch(error => {
        console.log(error)
      })
    }
  }
}
</script>
