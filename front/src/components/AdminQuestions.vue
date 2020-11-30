<template>
  <div>
    <b-container>
      <h1>問題管理ページ</h1>
      <div>
        <b-button
          href="/admin/question/add"
          variant="outline-primary"
          id="btn_add_question"
          pill
          block
          >問題を追加</b-button
        >
      </div>
      <ul>
        <li><a href="/admin/question/1/edit">テスト(id=1)</a></li>
      </ul>
    </b-container>
  </div>
</template>

<script>
// import axios from "axios"
export default {};
</script>

<style scoped>
#btn_add_question {
  margin: 10px;
  font-size: large;
}
</style>