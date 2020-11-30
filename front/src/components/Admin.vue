<style scoped>
#kinou {
  margin: auto;
}
</style>

<template>
  <div>
    <b-container>
      <h1>管理者ページ</h1>
      <div>
        <b-button
          href="/question"
          variant="outline-primary"
          id="btn_add_question"
          pill
          block
          >問題一覧へ</b-button
        >
      </div>
      <div id="kinou">
        <b-row>
          <b-col>
            <b-card border-variant="info" header="進捗管理">
              <b-card-text
                >ユーザの進捗状況を閲覧することができます。</b-card-text
              >
              <b-button
                href="/admin/user"
                pill
                variant="info"
                style="float: right"
              >
                Go ->
              </b-button>
            </b-card>
          </b-col>
          <b-col>
            <b-card border-variant="info" header="問題管理">
              <b-card-text> 問題の追加や編集ができます。 </b-card-text>
              <b-button
                href="/admin/question"
                pill
                variant="info"
                style="float: right"
              >
                Go ->
              </b-button>
            </b-card>
          </b-col>
        </b-row>
      </div>
      <!-- <h2>機能</h2>
      <b-list-group>
        <b-list-group-item to="/admin/question">問題管理</b-list-group-item>
        <b-list-group-item to="/admin/question/add"
          >問題追加画面</b-list-group-item
        >
        <b-list-group-item to="/admin/question/1/edit"
          >問題編集画面</b-list-group-item
        >
        <b-list-group-item to="/admin/user">進捗状況閲覧</b-list-group-item>
      </b-list-group> -->
    </b-container>
  </div>
</template>

<script>
export default {};
</script>