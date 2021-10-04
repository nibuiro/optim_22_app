<template>
  <section>
    <b-button
      label="編集"
      type="is-light"
      outlined
      @click="isComponentModalActive = true"
    />
    <b-modal
      v-model="isComponentModalActive"
      has-modal-card
      trap-focus
      :destroy-on-hide="false"
      aria-role="dialog"
      aria-label="Example Modal"
      aria-modal
    >
      <template #default="props">
        <modal-form
          :submissionProps="submission"
          @close="props.close"
          @displayMessage="isMessageModalActive = true"
        />
      </template>
    </b-modal>
    <b-modal v-model="isMessageModalActive">
      <b-message type="is-success" has-icon>
        編集が完了しました
        <br />
        ページを更新します
      </b-message>
    </b-modal>
  </section>
</template>

<script>
import * as api from "@/modules/API";

const ModalForm = {
  props: ["submissionProps"],
  data() {
    return {
      submission: {
        submission_id: this.submissionProps.submission_id,
        engineer_id: this.submissionProps.engineer.engineer_id,
        content: this.submissionProps.content,
        url: this.submissionProps.url
      },
      invalid: false,
      errorMessage: ""
    };
  },
  watch: {
    submission: {
      handler() {
        // 提出物URLと内容が入力されていればアラートを消す
        if (this.submission.content.length * this.submission.url.length > 0) {
          this.invalid = false;
        }
      },
      deep: true
    }
  },
  methods: {
    // 提出物を編集する
    async editSubmission() {
      // 提出物のURLと内容が入力されていれば
      if (this.submission.content.length * this.submission.url.length > 0) {
        const access_token = localStorage.getItem("access_token");
        api.editProfile(this, this.submission, access_token);
      } else {
        this.errorMessage = "すべての項目を入力してください";
        this.invalid = true;
      }
    }
  },
  /* html */
  template: `
    <form action="">
      <div class="modal-card" style="width: auto">
        <header class="modal-card-head">
          <p class="modal-card-title has-text-centered">提出物の編集</p>
          <button type="button" class="delete" @click="$emit('close')" />
        </header>
        <section class="modal-card-body">
          <b-message v-show="invalid" type="is-danger">
            {{ errorMessage }}
          </b-message>
          <b-field label="提出物がダウンロード可能なURL">
            <div class="control">
              <b-input
                type="text"
                v-model="submission.url"
                placeholder="アップロード先URL"
                required
              />
            </div>
          </b-field>
          <b-field label="提出物の詳細">
            <div class="control">
              <b-input
                type="textarea"
                v-model="submission.content"
                placeholder="提出物について具体的に説明してください。(500字以内)"
                maxlength="500"
                required
              />
            </div>
          </b-field>
        </section>
        <footer class="modal-card-foot is-flex is-justify-content-center">
          <b-button label="編集する" type="is-primary" @click="editSubmission" />
          <b-button label="キャンセル" @click="$emit('close')" />
        </footer>
      </div>
    </form>
  `
};

export default {
  data() {
    return {
      isComponentModalActive: false,
      isMessageModalActive: false
    };
  },
  watch: {
    // ユーザが提出物編集成功メッセージを閉じたらページをリロードする
    isMessageModalActive(newVal, oldVal) {
      if (newVal === false && oldVal === true) {
        const submission_id = this.$route.params.submission_id;
        this.$router.go({ name: "SubmissionPage", params: { submission_id } });
      }
    }
  },
  props: ["submission"],
  components: {
    ModalForm
  }
};
</script>
