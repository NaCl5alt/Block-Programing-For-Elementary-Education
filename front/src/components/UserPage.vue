<template>
  <div id="root">
    <ul>
      <h3>ユーザページ</h3>
      <div class="leftspace">
        <li>
          <a href="/question">問題一覧</a>
        </li>
      </div>
      <div class="leftspace">
        <li>
          <a href="/id">ユーザID変更</a>
        </li>
        <li>
          <a href="/name">ユーザ名変更</a>
        </li>
        <li>
          <a href="/passwd">パスワード変更</a>
        </li>
        <li>
          <a href="/delete">ユーザ削除</a>
        </li>
      </div>
    </ul>
  </div>
</template>

<style scoped>
#root {
  margin-left: 15%;
}
h3 {
  margin-top: 15px;
}
.leftspace {
  margin-left: 5%;
}
</style>