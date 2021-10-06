<!-- リクエスト編集フォーム -->

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
          :requestProps="request"
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
  props: ["requestProps"],
  data() {
    return {
      request: {
        request_id: this.requestProps.request_id,
        requestname: this.requestProps.requestname,
        content: this.requestProps.content
      },
      invalid: false,
      errorMessage: ""
    };
  },
  watch: {
    submission: {
      handler() {
        // 依頼名と依頼内容が入力されていればアラートを消す
        if (this.request.requestname.length * this.request.content.length > 0) {
          this.invalid = false;
        }
      },
      deep: true
    }
  },
  methods: {
    // リクエストを編集する
    async editRequest() {
      // 依頼名と依頼内容が入力されていれば
      if (this.request.requestname.length * this.request.content.length > 0) {
        const access_token = localStorage.getItem("access_token");
        api.editRequest(this, this.request, access_token);
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
          <p class="modal-card-title has-text-centered">リクエスト内容の編集</p>
          <button type="button" class="delete" @click="$emit('close')" />
        </header>
        <section class="modal-card-body">
          <b-message v-show="invalid" type="is-danger">
            {{ errorMessage }}
          </b-message>
          <b-field label="依頼名">
            <div class="control">
              <b-input
                type="text"
                v-model="request.requestname"
                placeholder="依頼内容を分かりやすく一言で！"
                required
              />
            </div>
          </b-field>
          <b-field label="依頼内容の詳細">
            <div class="control">
              <b-input
                type="textarea"
                v-model="request.content"
                placeholder="依頼内容について具体的に説明してください。(500字以内)"
                maxlength="500"
                required
              />
            </div>
          </b-field>
        </section>
        <footer class="modal-card-foot is-flex is-justify-content-center">
          <b-button label="編集する" type="is-primary" @click="editRequest" />
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
    // ユーザがリクエスト編集成功メッセージを閉じたらページをリロードする
    isMessageModalActive(newVal, oldVal) {
      if (newVal === false && oldVal === true) {
        const request_id = this.$route.params.request_id;
        this.$router.go({ name: "RequestPage", params: { request_id } });
      }
    }
  },
  props: ["request"],
  components: {
    ModalForm
  }
};
</script>
