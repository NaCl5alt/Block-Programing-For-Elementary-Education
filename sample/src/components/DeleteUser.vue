<script>
import axios from 'axios'
export default {
  methods: {
    delete () {
      axios.delete(process.env.VUE_APP_API_URL_BASE + 'user', { withCredentials: true }).then(res => {
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
