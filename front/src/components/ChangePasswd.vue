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
      axios.post('/api/user/password', {
        password: this.passwd
      }, { withCredentials: true,
        headers: { Authorization: `Bearer ${this.$cookies.get("token")}` }
      }).then(res => {
        this.$router.push({ path: '/' })
        console.log(res)
      }).catch(error => {
        console.log(error)
      })
    }
  }
}
</script>
