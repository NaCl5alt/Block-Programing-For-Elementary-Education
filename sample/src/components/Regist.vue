<template>
    <div>
        <form @submit.prevent="regist">
            <h1>ユーザ登録</h1>
            ユーザID: <input v-model="userid" type="text">
            ユーザ名: <input v-model="name" type="text">
            パスワード: <input v-model="pass" type="password">
            <button type="submit" class="btn btn-primary">登録</button>
        </form>
    </div>
</template>

<script>
import axios from 'axios'
export default {
  data () {
    return {
      userid: '',
      name: '',
      pass: ''
    }
  },
  methods: {
    regist () {
      axios.post(process.env.VUE_APP_API_URL_BASE + 'user', {
        userid: this.userid,
        name: this.name,
        pass: this.pass
      }).then(res => {
        this.$router.push({ path: '/login' })
        console.log(res)
      }).catch(error => {
        console.log(error)
      })
    }
  }
}
</script>
