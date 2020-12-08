<template>
    <div>
        <form @submit.prevent="regist">
            <h1>ユーザID変更</h1>
            変更後ユーザID: <input v-model="userid" type="text">
            <button type="submit" class="btn btn-primary">変更</button>
        </form>
        <font color="red" v-show="used">そのユーザIDは既に使われています</font>
    </div>
</template>

<script>
import axios from 'axios'
export default {
  data () {
    return {
      userid: '',
      userd: false
    }
  },
  methods: {
    async check () {
      await axios.post('/api/user/check', {
        userid: this.userid,
      }).then(res => {
        console.log(res)
        this.used = res.data.exist
      }).catch(error => {
        console.log(error)
      })
    },
    async regist () {
      await this.check()
      if(this.used === false ) await axios.post('/api/user/id', {
          userid: this.userid
        }, { withCredentials: true,
          headers: { Authorization: `Bearer ${this.$cookies.get("token")}` }
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
