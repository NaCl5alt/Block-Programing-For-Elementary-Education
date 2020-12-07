<script>
export default {
  beforeMount () {
    this.$cookies.remove('token')
    this.$cookies.remove('admin')
    this.$router.push('/')
    this.$router.go()
  }
}
</script>