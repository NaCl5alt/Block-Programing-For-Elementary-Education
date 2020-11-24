<template>
  <div>
    <!-- <b-jumbotron> -->
      <h1>{{ this.mainData["title"] }} <span id="qId">(id: {{ this.mainData.id }})</span></h1>
      
      <hr class="my-4">

      <p>{{ this.mainData["content"] }}</p>

      <ul>
        <b-button v-b-modal="'modal-1'">{{ hint }}</b-button>
        <b-modal id="modal-1" title="ヒント1">
          <p>{{ hint }}</p>
        </b-modal>
      </ul>
    <!-- </b-jumbotron> -->

    <BlocklyComponent id="blockly1" :options="options" ref="foo"></BlocklyComponent>
    <p id="code">
      <button v-on:click="showCode()">Show JavaScript</button>
      <pre v-html="code"></pre>
    </p>
  </div>
</template>

<script>
/**
 * @license
 *
 * Copyright 2019 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/**
 * @fileoverview Main Vue component that includes the Blockly component.
 * @author samelh@google.com (Sam El-Husseini)
 */

import BlocklyComponent from "./BlocklyComponent.vue";
import "../blocks/stocks";
import "../prompt";

import BlocklyJS from "blockly/javascript";
import axios from "axios";

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
      mainData: {
        id: 0,
        title: "",
        content: "",
        hints: [],
      },
      datas: {
        qid: 1,
        title: "TEST",
        content: "ASADKSJFNJKDNFKJAN",
        hints: [{ hint: "HINT1" }, { hint: "HINT2" }],
      },
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
    window: (onload = function () {}),
    // window: (onload = function () {
    //   // axios.get("/api/question/${this.$route.params['id']}")
    //   //   .then((response) => {
    //   // })
    //   // .catch((err) => {
    //   // console.log(err);
    //   // let path = "/question";
    //   // this.$router.push({ path: path });
    //   // this.$router.go();
    //   // });
    // }),
    // mtFunc マウント時の処理
    async mtFunc() {
      this.mainData.id = this.$route.params["id"];

      axios
        .get("/api/question/$(this.questionId)")
        .then((res) => {
          this.mainData.title = res.data["title"];
          this.mainData.content = res.data["qid"];
          this.mainData.hints = res.data["hints"];
        })
        .catch((error) => {
          console.log(error);
        });
    },
    showCode() {
      this.code = BlocklyJS.workspaceToCode(this.$refs["foo"].workspace);
    },
  },
  beforeMount() {
    this.mtFunc();
    // this.questionTitle = this.datas["title"];
  },
};
</script>

<style>
html,
body {
  margin: 0;
}

#code {
  position: absolute;
  right: 0;
  bottom: 0;
  width: 50%;
  height: 50%;
  margin: 0;
  background-color: beige;
}

#blockly1 {
  position: absolute;
  left: 0;
  bottom: 0;
  width: 50%;
  height: 50%;
}

#qId {
  color: gray;
  font-size: 50%;
}
</style>
