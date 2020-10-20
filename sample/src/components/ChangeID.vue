<template>
    <div>
        <form @submit.prevent="regist">
            <h1>ユーザID変更</h1>
            変更後ユーザID: <input v-model="userid" type="text">
            <button type="submit" class="btn btn-primary">変更</button>
        </form>
    </div>
</template>

<script>
import axios from 'axios'
export default {
  data () {
    return {
      userid: ''
    }
  },
  methods: {
    regist () {
      var params = new URLSearchParams()
      params.append('userid', this.userid)
      axios.post(process.env.VUE_APP_API_URL_BASE + 'user/id', params, { withCredentials: true }).then(res => {
        this.$router.push({ path: '/login' })
        console.log(res)
      }).catch(error => {
        console.log(error)
      })
    }
  }
}
</script>
