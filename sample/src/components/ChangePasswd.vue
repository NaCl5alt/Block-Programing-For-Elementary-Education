<template>
    <div>
        <form @submit.prevent="regist">
            <h1>パスワード変更</h1>
            変更後パスワード: <input v-model="passwd" type="password">
            <button type="submit" class="btn btn-primary">変更</button>
        </form>
    </div>
</template>

<script>
import axios from 'axios'
export default {
  data () {
    return {
      passwd: ''
    }
  },
  methods: {
    regist () {
      var params = new URLSearchParams()
      params.append('passwd', this.passwd)
      axios.post(process.env.VUE_APP_API_URL_BASE + 'user/passwd', params, { withCredentials: true }).then(res => {
        this.$router.push({ path: '/' })
        console.log(res)
      }).catch(error => {
        console.log(error)
      })
    }
  }
}
</script>
