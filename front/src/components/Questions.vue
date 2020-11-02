<template>
  <div>
    <h2>問題一覧</h2>
    <div class="row">
      <div class="col">
        No.
      </div>
      <div class="col-8">
        問題名
      </div>
    </div>
    <div v-for="question in questions" :key="question.id" class="row">
      <div class="col">
      <i v-if="question.progress" class="far fa-check-square"></i>
      <i v-else class="far fa-square"></i>
        {{ question.id }}
      </div>
      <div class="col-8">
        {{ question.title }}
      </div>
    </div>
    <infinite-loading @infinite="infiniteHandler" spinner="spiral">
      <div slot="spinner">ロード中...</div>
      <div slot="no-more">もう登録されている問題がないよ</div>
      <div slot="no-results">登録されている問題がないよ</div>
    </infinite-loading>
  </div>
</template>

<script>
export default {
  name: 'Questions',
  data () {
    return {
      questions: [],
      end: 0,
      max: 0
    }
  },
  methods: {
    async infiniteHandler ($state) {
      if (this.end === 0) await this.firstQuestion()
      if (this.end >= this.max) {
        $state.complete()
      } else {
        await this.getQuestions()
        $state.loaded()
      }
    },
    async getCount () {
      this.max = 2
    },
    async firstQuestion() {
      this.questions.push({
        "id":1,"title": "Title","progress":false
        }
        )
      this.end = this.questions.length
      console.log(this.questions)
    },
    async getQuestions(){
      this.questions = this.questions.concat({
        "id":2,"title": "Title","progress":true
      })
      this.end = this.questions.length
      console.log(this.questions)
    }
  },
  beforeMount(){
    this.getCount()
  }
}
</script>