<!-- サブミッション提出フォーム -->

<template>
  <section>
    <b-button
      label="提出する"
      type="is-primary"
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
          @close="props.close"
          @displayMessage="isMessageModalActive = true"
        />
      </template>
    </b-modal>
    <b-modal v-model="isMessageModalActive">
      <b-message type="is-success" has-icon>
        提出が完了しました
        <br />
        ページを更新します
      </b-message>
    </b-modal>
  </section>
</template>

<script>
import * as api from "API";

const ModalForm = {
  data() {
    return {
      submission: {
        request_id: this.$route.params.request_id,
        engineer_id: localStorage.getItem("user_id"),
        content: "",
        url: ""
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
    async submitSubmission() {
      // 提出物のURLと内容が入力されていれば
      if (this.submission.content.length * this.submission.url.length > 0) {
        const access_token = localStorage.getItem("access_token");
        api.submitSubmission(this, this.submission, access_token);
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
          <b-button label="編集する" type="is-primary" @click="submitSubmission" />
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
    // ユーザが提出物提出成功メッセージを閉じたらページをリロードする
    isMessageModalActive(newVal, oldVal) {
      if (newVal === false && oldVal === true) {
        const request_id = this.$route.params.request_id;
        this.$router.go({
          name: "RequestPage",
          params: { request_id }
        });
      }
    }
  },
  components: {
    ModalForm
  }
};
</script>
