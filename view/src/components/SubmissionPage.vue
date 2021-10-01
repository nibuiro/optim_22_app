<!-- サブミッション詳細ページ -->

<template>
  <div class="container" :request="request" :submission="submission">
    <section class="hero is-primary is-small mb-3">
      <b-tooltip
        style="position: absolute;"
        :label="submission.comment"
        type="is-light"
        position="is-right"
        always
      >
        <router-link
          :to="{
            name: 'MyPage',
            params: { user_id: submission.engineer.userid }
          }"
        >
          <div
            class="ml-3 mt-3 mb-6"
            :style="iconStyle(64, submission.engineer.icon)"
          />
        </router-link>
      </b-tooltip>
      <div class="hero-body is-flex pt-0 pb-5">
        <p class="title mb-0 pt-2" style="margin-left: 64px;">
          {{ submission.engineer.username }}さんの提出物
        </p>
        <submission-editor
          class="is-light ml-auto mt-5"
          :submission="submission"
        />
      </div>
    </section>
    <section class="mb-3">
      <b-tabs type="is-boxed">
        <b-tab-item>
          <template #header>
            <b-icon icon="file-upload-outline"></b-icon>
            <span>提出物詳細</span>
          </template>
          <div class="content">
            <ul :request="request" :submission="submission">
              <li>
                提出日時：
                {{
                  `${new Date(submission.date).toLocaleDateString()}
                   ${new Date(submission.date).toLocaleTimeString()}`
                }}
              </li>
              <li>依頼名　：{{ request.request }}</li>
              <li>
                提出物　：
                <a :href="submission.url">
                  {{ submission.url }}
                </a>
              </li>
              <li>
                コメント：
                {{ submission.comment }}
              </li>
            </ul>
          </div>
        </b-tab-item>
      </b-tabs>
    </section>
  </div>
</template>

<script>
import SubmissionEditor from "@/components/SubmissionEditor";

const submission = require("../../src/assets/sampleSubmission.json");
const request = require("../../src/assets/sampleRequest.json");

export default {
  data() {
    return {
      submission,
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
    "submission-editor": SubmissionEditor
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped></style>
