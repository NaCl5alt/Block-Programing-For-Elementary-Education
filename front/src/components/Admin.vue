<style scoped>
</style>

<template>
  <div>
    <b-container fluid>
      <h1>管理者ページ</h1>
      <h2>機能</h2>
      <b-list-group>
        <b-list-group-item to="/admin/question">問題管理</b-list-group-item>
        <b-list-group-item to="/admin/user">進捗状況閲覧</b-list-group-item>
      </b-list-group>
    </b-container>
  </div>
</template>

<script>
export default {};
</script>