<!-- サブミッション詳細ページ -->

<template>
  <div class="container">
    <section class="hero is-primary is-small mb-3">
      <b-tooltip
        style="position: absolute;"
        :label="submission.content"
        type="is-light"
        position="is-right"
        always
      >
        <router-link
          :to="{
            name: 'MyPage',
            params: { user_id: submission.engineer.user_id }
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
            <b-icon icon="file-upload-outline" />
            <span>提出物詳細</span>
          </template>
          <div class="content">
            <ul>
              <li>
                提出日時：
                {{
                  `${new Date(submission.createdat).toLocaleDateString()}
                   ${new Date(submission.createdat).toLocaleTimeString()}`
                }}
              </li>
              <li>
                依頼名　：
                <router-link
                  :to="{
                    name: 'RequestPage',
                    params: { request_id: submission.request.request_id }
                  }"
                >
                  {{ submission.request.requestname }}
                </router-link>
              </li>
              <li>
                <div class="is-flex is-align-items-center">
                  提出者　：
                  <router-link
                    class="is-flex is-align-items-center"
                    :to="{
                      name: 'MyPage',
                      params: { user_id: submission.engineer_id }
                    }"
                  >
                    <b-tooltip :label="submission.engineer.username">
                      <div :style="iconStyle(32, submission.engineer.icon)" />
                    </b-tooltip>
                    {{ submission.engineer.username }}
                  </router-link>
                </div>
              </li>
              <li>
                提出物　：
                <a
                  class="is-inline-flex is-align-items-center"
                  :href="submission.url"
                >
                  <b-icon icon="attachment" />
                  {{ submission.url }}
                </a>
              </li>
              <li>
                コメント：
                {{ submission.content }}
              </li>
            </ul>
          </div>
        </b-tab-item>
      </b-tabs>
    </section>
  </div>
</template>

<script>
import * as api from "@/modules/API";
import SubmissionEditor from "@/components/SubmissionEditor";

export default {
  data() {
    return {
      submission: {
        submission_id: null,
        createdat: "",
        request_id: null,
        engineer: {
          user_id: null,
          username: "",
          icon: "",
          comment: "",
          SNS: {}
        },
        content: "",
        url: "",
        request: {
          request_id: null,
          finish: null,
          createdat: "",
          requestname: "",
          client: {}
        },
        engineers: [],
        content: "",
        submissions: [],
        winner: null
      }
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
  },
  async created() {
    const submission_id = this.$route.params.submission_id;
    this.submission = await api.getsubmission(submission_id);
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped></style>
