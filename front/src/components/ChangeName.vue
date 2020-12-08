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
      axios.post('/api/user/edit', {
        name: this.name
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
