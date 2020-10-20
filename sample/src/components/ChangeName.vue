<template>
    <div>
        <form @submit.prevent="regist">
            <h1>ユーザ名変更</h1>
            変更後ユーザ名: <input v-model="name" type="text">
            <button type="submit" class="btn btn-primary">変更</button>
        </form>
    </div>
</template>

<script>
import axios from 'axios'
export default {
  data () {
    return {
      name: ''
    }
  },
  methods: {
    regist () {
      var params = new URLSearchParams()
      params.append('name', this.name)
      axios.post(process.env.VUE_APP_API_URL_BASE + 'user/name', params, { withCredentials: true }).then(res => {
        this.$router.push({ path: '/' })
        console.log(res)
      }).catch(error => {
        console.log(error)
      })
    }
  }
}
</script>
