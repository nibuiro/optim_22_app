<!-- リクエスト詳細ページ -->

<template>
  <div class="container" :request="request">
    <section class="hero is-primary is-small mb-3">
      <b-tooltip
        style="position: absolute;"
        :label="request.request"
        type="is-light"
        position="is-right"
        always
      >
        <router-link
          :to="{
            name: 'MyPage',
            params: { user_id: request.client.userid }
          }"
        >
          <div
            class="ml-3 mt-3 mb-6"
            :style="iconStyle(64, request.client.icon)"
          />
        </router-link>
      </b-tooltip>
      <div class="hero-body is-flex pt-0 pb-5">
        <p class="title mb-0 pt-2" style="margin-left: 64px;">
          {{ request.client.username }}さんの依頼
        </p>
        <request-editor class="is-light ml-auto mt-5" :request="request" />
      </div>
    </section>
    <section class="mb-3">
      <b-tabs type="is-boxed">
        <b-tab-item>
          <template #header>
            <b-icon icon="account-question-outline"></b-icon>
            <span>
              依頼内容
            </span>
          </template>
          <div class="content">
            <ul :request="request">
              <li>
                依頼日時：
                {{
                  `${new Date(request.date).toLocaleDateString()}
                   ${new Date(request.date).toLocaleTimeString()}`
                }}
                <b-tag
                  :type="
                    request.accepting === true ? 'is-success' : 'is-danger'
                  "
                >
                  {{ request.accepting === true ? "受付中" : "終了" }}
                </b-tag>
              </li>
              <li>依頼名　：{{ request.request }}</li>
              <li>依頼内容：{{ request.detail }}</li>
              <li>
                <div class="is-flex is-align-items-center">
                  依頼者　：
                  <router-link
                    class="is-flex is-align-items-center"
                    :to="{
                      name: 'MyPage',
                      params: { user_id: request.client.userid }
                    }"
                  >
                    <b-tooltip :label="request.client.username">
                      <div :style="iconStyle(32, request.client.icon)" />
                    </b-tooltip>
                    {{ request.client.username }}
                  </router-link>
                </div>
              </li>
              <li>
                <div class="is-flex is-align-items-center">
                  参加者　：
                  <router-link
                    class="is-flex is-align-items-center mr-3"
                    v-for="engineer in request.engineers"
                    :key="engineer.userid"
                    :to="{
                      name: 'MyPage',
                      params: { user_id: engineer.userid }
                    }"
                  >
                    <b-tooltip :label="engineer.username">
                      <div :style="iconStyle(32, engineer.icon)" />
                    </b-tooltip>
                    {{ engineer.username }}
                  </router-link>
                </div>
              </li>
              <li>
                提出物　：
                <router-link
                  class="mr-3 is-inline-flex is-align-items-center"
                  v-for="submission in request.submissions"
                  :key="submission.submissionid"
                  :to="{
                    name: 'SubmissionPage',
                    query: { id: submission.submissionid }
                  }"
                >
                  <b-icon icon="file-upload-outline" />
                  {{ submission.engineer.username }}さんの提出
                </router-link>
              </li>
            </ul>
          </div>
        </b-tab-item>
        <b-tab-item>
          <template #header>
            <b-icon icon="forum-outline"></b-icon>
            <span>
              ディスカッション
            </span>
          </template>
          <div class="content">
            ここにディスカッションが表示されます。
          </div>
        </b-tab-item>
      </b-tabs>
    </section>
    <section class="mb-3 is-flex is-justify-content-center">
      <!-- 依頼主であり提出物が1つ以上あれば -->
      <choose-winner v-if="true" :request="request" />
      <!-- 依頼主以外であれば -->
      <request-applier v-if="true" :request="request" />
    </section>
  </div>
</template>

<script>
import RequestEditor from "@/components/RequestEditor";
import ChooseWinner from "@/components/ChooseWinner.vue";
import RequestApplier from "@/components/RequestApplier.vue";

const request = require("../../src/assets/sampleRequest.json");

export default {
  data() {
    return {
      request
    };
  },
  methods: {
    iconStyle(size, image) {
      return {
        width: `${size}px`,
        height: `${size}px`,
        backgroundImage: `url("${image}")`,
        backgroundSize: "contain",
        backgroundRepeat: "no-repeat",
        backgroundPosition: "center",
        borderRadius: "100%"
      };
    }
  },
  components: {
    "request-editor": RequestEditor,
    "choose-winner": ChooseWinner,
    "request-applier": RequestApplier
  }
};
</script>

<!-
    RequestApplier- Add "scoped" attribute to limit CSS to this component only -->
<style scoped></style>
