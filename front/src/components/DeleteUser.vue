<script>
import axios from 'axios'
export default {
  methods: {
    delete () {
      axios.delete('/api/user', { withCredentials: true,
        headers: { Authorization: `Bearer ${this.$cookies.get("token")}` }
      }).then(res => {
        this.$router.push({ path: '/' })
        console.log(res)
      }).catch(error => {
        console.log(error)
      })
    }
  },
  beforeMount () {
    this.delete()
  }
}
</script>
