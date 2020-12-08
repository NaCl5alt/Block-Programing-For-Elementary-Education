<template>
  <div>
    <b-containter fluid>
      <b-container id="question" style="width: 100%; height: 50%">
        <b-row>
          <b-col>
            <div>
              <h2>
                {{ questionTitle }}
                <span style="color: gray; font-size: 50%"
                  >(id: {{ questionId }})</span
                >
              </h2>
              <hr class="my-4" />
              <p>{{ questionContent }}</p>
              <div role="hintlist">
                <ul style="list-style: none">
                  <div v-for="h in questionsHints" :key="h.id">
                    <li style="float: left; margin: 10px">
                      <b-button
                        pill
                        v-b-toggle.collapse
                        variant="info"
                        @click="showHint(h.id)"
                        >{{ h.title }}</b-button
                      >
                    </li>
                  </div>
                </ul>
                <p><br /><br /></p>
                <div>
                  <b-collapse id="collapse" v-model="visible">
                    <b-card>{{ visibleHint }}</b-card>
                  </b-collapse>
                </div>
              </div>
            </div>
          </b-col>
          <b-col>
            <div>
              <b-row>
                <b-col>
                  <b-form-input
                    v-model="usersAnswer"
                    type="text"
                    name="usersAnswer"
                    placeholder="ここに答えを入力"
                  ></b-form-input>
                  <br />
                  <b-button
                    pill
                    style="float: right"
                    variant="primary"
                    v-on:click="checkAnswer"
                    >答えを送信</b-button
                  >
                </b-col>
              </b-row>
              <div v-if="showAnswer">
                <b-alert variant="success" v-if="visibleAnswer" show
                  >正解!!</b-alert
                >
                <b-alert variant="danger" v-else show>間違い!!</b-alert>
              </div>
              <div v-else>
                <b-alert variant="light" show></b-alert>
              </div>
              <b-row>
                <b-button
                  pill
                  style="float: right"
                  variant="primary"
                  v-on:click="runCode()"
                  >▶プログラムを実行</b-button
                >
              </b-row>
            </div>
          </b-col>
        </b-row>
      </b-container>
      <b-container>
        <BlocklyComponent
          id="blockly1"
          :options="options"
          ref="foo"
        ></BlocklyComponent>
      </b-container>
    </b-containter>
  </div>
</template>

<script>
import BlocklyComponent from "./BlocklyComponent.vue";
import "../blocks/stocks";
import "../prompt";

import BlocklyJS from "blockly/javascript";
import Axios from "axios";

export default {
  name: "Question",
  components: {
    BlocklyComponent,
  },
  data() {
    return {
      code: "",
      questionId: 0,
      questionTitle: "",
      questionContent: "",
      questionsHints: [],
      visibleId: 0, // 表示するヒントの番号
      visibleHint: "", // 表示しているヒント
      visible: false,
      usersAnswer: "",
      visibleAnswer: false, // 正誤判定のフラグ
      showAnswer: false, // 答えのアラートを表示するかのフラグ
      options: {
        media: "media/",
        grid: {
          spacing: 25,
          length: 3,
          colour: "#ccc",
          snap: true,
        },
        toolbox: `<xml>
          <category name="Logic" colour="%{BKY_LOGIC_HUE}">
            <block type="controls_if"></block>
            <block type="logic_compare"></block>
            <block type="logic_operation"></block>
            <block type="logic_negate"></block>
            <block type="logic_boolean"></block>
          </category>
          <category name="Loops" colour="%{BKY_LOOPS_HUE}">
            <block type="controls_repeat_ext">
              <value name="TIMES">
                <block type="math_number">
                  <field name="NUM">10</field>
                </block>
              </value>
            </block>
            <block type="controls_whileUntil"></block>
          </category>
          <category name="Math" colour="%{BKY_MATH_HUE}">
            <block type="math_number">
              <field name="NUM">123</field>
            </block>
            <block type="math_arithmetic"></block>
            <block type="math_single"></block>
          </category>
          <category name="Text" colour="%{BKY_TEXTS_HUE}">
            <block type="text"></block>
            <block type="text_length"></block>
            <block type="text_print"></block>
          </category>
          <category name="Variables" custom="VARIABLE" colour="%{BKY_VARIABLES_HUE}">
            </category>
          <category name="Stocks" colour="%{BKY_LOOPS_HUE}">
            <block type="stock_buy_simple"></block>
            <block type="stock_buy_prog"></block>
            <block type="stock_fetch_price"></block>
          </category>
        </xml>`,
      },
    };
  },
  methods: {
    // mtFunc マウント時の処理
    async mtFunc() {
      this.questionId = this.$route.params["id"];

      await Axios.get(`/api/question/contents/${this.questionId}`, {
        headers: { Authorization: `Bearer ${this.$cookies.get("token")}` },
      })
        .then((res) => {
          // console.log(res.data)
          var bufHints = [];
          var bufStr0 = "";
          var bufStr1 = "";
          this.questionTitle = res.data["title"]; // タイトルを格納
          this.questionContent = res.data["content"]; // 問題文を格納
          for (let i = 0; i < res.data["hints"].length; i++) {
            bufStr0 = "collapse-" + i;
            bufStr1 = "ヒント " + i;
            bufHints.push({
              id: i,
              collapseId: bufStr0,
              title: bufStr1,
              hints: res.data["hint"][i],
            });
          }
          if (res.data["hints"][0] == null) {
            // ヒントが無い問題
            this.visibleHint = "空";
          } else {
            // ヒントがあればとりあえず最初のを入れる
            this.visibleHint = res.data["hints"][0];
          }
          this.questionsHints = bufHints;
        })
        .catch((error) => {
          // エラー処理
          this.questionTitle = "[ERROR] ページが存在しません。";
          this.questionContent = "[ERROR]";
          this.questionsHints = [
            {
              id: 0,
              collapseId: "collapse-0",
              title: "ヒント 0",
              hint: "[ERROR]",
            },
            {
              id: 1,
              collapseId: "collapse-1",
              title: "ヒント 1",
              hint: "[ERROR]",
            },
          ];
          console.log(error);
        });
    },
    async showHint(hintId) {
      this.visibleHint = this.questionsHints[hintId];
      // await Axios.get(`/api/question/contents/${this.questionId}`, {
      //   headers: { Authorization: `Bearer ${this.$cookies.get("token")}` },
      // })
      //   .then((res) => {
      //     this.visibleHint = res.data["hints"][questionId];
      //   })
      //   .catch((error) => {
      //     this.visibleHint = "[ERROR] " + questionId;
      //     console.log(error);
      //   });
      // 表示するしないの処理
      if (hintId != this.visibleId) {
        this.visibleId = !this.visibleId;
      } else {
        this.visibleId = true;
      }
    },
    runCode() {
      // コードの実行
      this.code = BlocklyJS.workspaceToCode(this.$refs["foo"].workspace);
      try {
        eval(this.code);
      } catch (e) {
        alert(e);
      }
    },
    // 回答チェック
    async checkAnswer() {
      await Axios.post(
        `/api/question/${this.questionId}`,
        {
          answer: this.usersAnswer,
        },
        {
          headers: { Authorization: `Bearer ${this.$cookies.get("token")}` },
        }
      )
        .then((res) => {
          if (res.data["accept"]) {
            // 正解だった時
            this.showAnswer = true;
            this.visibleAnswer = true;
          } else {
            // 間違いだった時
            this.showAnswer = true;
            this.visibleAnswer = false;
          }
        })
        .catch((error) => {
          console.log(error);
          this.showAnswer = true;
          this.visibleAnswer = !this.visibleAnswer;
        });
    },
  },
  beforeMount() {
    this.mtFunc();
  },
};
</script>

<style>
html,
body {
  margin: 0;
}

/* #question {
  position: absolute;
  width: 100%;
  height: 50%;
} */

#blockly1 {
  position: absolute;
  left: 0;
  bottom: 0;
  width: 100%;
  height: 45%;
}
</style>
